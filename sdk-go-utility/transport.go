package sdkgoutility

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// transportModularUserAgentPrefix starts the User-Agent segments that name a module of
// the modular SDK, such as sdk-go-core/0.1.0 and sdk-go-modelregistry/0.7.0.
const transportModularUserAgentPrefix = "sdk-go-"

// transportCallHeaders are the headers that the classic generated calls set
// themselves, after the request visitors have run.
var transportCallHeaders = []string{"Accept", "Content-Type", "X-Databricks-Workspace-Id"}

// transportCredentialHeaders are the headers that http.Client leaves out of a
// request that it redirects to another domain. The list mirrors the one in
// the redirect code of net/http, which is not exported.
var transportCredentialHeaders = []string{
	"Authorization", "Www-Authenticate", "Cookie", "Cookie2", "Proxy-Authorization", "Proxy-Authenticate",
}

// Transport returns a transport that sends the requests of a modular SDK
// client with the headers that the classic client sends for cfg. Pass it to
// the modular client in the http.Client of client.WithHTTPClient.
//
// For every request the transport works on a clone, so the caller's request
// stays as it was, and it:
//
//  1. runs the request visitors of config.HTTPClientConfigFromConfig(cfg) in
//     order. They send the request to cfg.Host, add the authentication type,
//     the CI/CD provider, the runtime and the agent to the User-Agent data of
//     the context, and call the cfg.Headers hook. An error of a visitor is
//     returned wrapped, starting with "failed during request visitor: ". A
//     redirect that the http.Client follows keeps the target that the server
//     gave it, as it does with the classic client;
//  2. puts back the Accept, Content-Type and X-Databricks-Workspace-Id that
//     the caller's request had, and sets Accept to application/json when it
//     had none. The classic generated calls set these headers after the
//     visitors, so with both clients the cfg.Headers hook changes one of them
//     only when the call doesn't set it;
//  3. on a request that the http.Client made to follow a redirect, removes
//     every credential header, such as Authorization and Cookie, that the
//     request arrived without. The http.Client leaves these out of a request
//     that it redirects to another domain, and the hook, which has run again
//     on that request, must not send them there. The price is that a
//     credential header which only the hook sets is on no redirected request,
//     where the classic client sends it within the domain;
//  4. adds the traceparent header unless the request has one;
//  5. sets User-Agent to the classic one for the context that the visitors
//     left, followed by the sdk-go- segments of the User-Agent that the
//     modular SDK set, which mark the call as coming from the modular SDK.
//
// It then sends the request over cfg.HTTPTransport when that is set, and
// otherwise over a transport with the proxy, connection and TLS settings of
// the classic default, including cfg.InsecureSkipVerify. Every result of
// Transport for such a config sends over the same transport, so a caller can
// build a modular client for each call and still reuse connections.
//
// The transport returns the response as it is, with one exception. A workspace
// that refuses a request from outside its private link redirects it to its
// login page, /login.html?error=private-link-validation-error, which answers
// with a 200. The classic client reports that page as an error for every
// method. A modular client would fail to decode it, or succeed when the call
// decodes no response, as a delete does. So when apierr.GetAPIError of the
// classic SDK reports a response with a status below 400 as an error, the
// transport returns in its place the status of that error, 403, and a JSON
// body with its error code, PRIVATE_LINK_VALIDATION_ERROR, and its message,
// which a modular client reads as an API error.
//
// wrap, when it isn't nil, puts a transport of the caller between the steps
// above and the connection, which a caller can't do around the result,
// because Transport picks the transport to send over itself. Transport calls
// wrap once, with that transport, and sends every request over the one that
// wrap returns. That one is asked to send each request as it is sent: after
// the visitors, with the headers above, and never a request that a visitor
// refused. It returns the response before the exception above is made, so for
// a private link redirect it sees the login page and the caller gets the 403.
// This is where the classic client writes its debug log. An error of wrap,
// and a nil transport from it, are errors of Transport.
//
// Transport resolves cfg, as building a classic client does, and returns the
// error of that resolution.
//
// Transport leaves to the caller:
//
//   - credentials. It doesn't run the AuthVisitor of the config, so the
//     credentials of the modular client authenticate the request. They must
//     authenticate cfg before the first request, as credentials built on
//     cfg.Authenticate do, because the visitor that adds the authentication
//     type reads cfg.AuthType, which authentication sets. Without it that
//     visitor panics, and the transport returns the panic as a visitor error;
//   - retries, rate limiting, the HTTP timeout, request logging and the
//     mapping of error responses, which the classic client does around its
//     transport;
//   - the host, the workspace ID and the call headers of the modular client.
func Transport(cfg *config.Config, wrap func(base http.RoundTripper) (http.RoundTripper, error)) (http.RoundTripper, error) {
	if cfg == nil {
		return nil, errors.New("cannot build a transport for a nil config")
	}
	clientCfg, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	base := transportClassicBase(clientCfg)
	if wrap != nil {
		if base, err = wrap(base); err != nil {
			return nil, fmt.Errorf("cannot wrap the transport: %w", err)
		}
		if base == nil {
			return nil, errors.New("cannot wrap the transport: wrap returned no transport")
		}
	}
	return &transportRoundTripper{base: base, visitors: clientCfg.Visitors}, nil
}

// transportRoundTripper is the http.RoundTripper that Transport returns. base
// is the transport that it sends over: the one of the classic client for the
// config, or what the wrap of the caller made of it.
type transportRoundTripper struct {
	base     http.RoundTripper
	visitors []httpclient.RequestVisitor
}

// RoundTrip sends a clone of req with the classic headers over the base
// transport. It closes the body of req when it fails before sending, as the
// http.RoundTripper contract asks.
func (t *transportRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	out := req.Clone(req.Context())
	if out.Header == nil {
		out.Header = http.Header{}
	}
	if err := t.visit(out); err != nil {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, fmt.Errorf("failed during request visitor: %w", err)
	}
	if out.URL != nil && req.URL != nil {
		if req.Response != nil {
			// http.Client follows a redirect with a new request, which carries
			// the response that redirected it. The classic client runs its
			// visitors before http.Client, so a redirect goes where the server
			// sent it and not to cfg.Host.
			out.URL.Scheme, out.URL.Host, out.Host = req.URL.Scheme, req.URL.Host, req.Host
		} else if out.URL.Host != req.URL.Host {
			// A classic request is built from a path, so its Host header is
			// the host that the visitors put in the URL. A modular request
			// arrives with the host of its client in both places.
			out.Host = out.URL.Host
		}
	}
	transportClassicHeaders(req, out)
	resp, err := t.base.RoundTrip(out)
	if err != nil || resp == nil {
		return resp, err
	}
	return transportPrivateLinkResponse(out, resp), nil
}

// transportClassicHeaders gives out, the clone of req that the visitors have
// run on, the headers that the classic client sends after its visitors.
func transportClassicHeaders(req, out *http.Request) {
	// The classic generated call sets its headers after the visitors, so the
	// cfg.Headers hook changes one of them only when the call doesn't set it.
	for _, name := range transportCallHeaders {
		if values := req.Header.Values(name); len(values) > 0 {
			out.Header[name] = slices.Clone(values)
		}
	}
	// The classic generated calls ask for JSON. The modular SDK sends no Accept.
	if req.Header.Get("Accept") == "" {
		out.Header.Set("Accept", "application/json")
	}
	if req.Response != nil {
		// http.Client made req to follow a redirect, without the credentials
		// when it goes to another domain. The classic client runs the hook
		// before http.Client, so there the credentials of the hook are left
		// out as well. Here the hook has just run on the redirected request.
		for _, name := range transportCredentialHeaders {
			if len(req.Header.Values(name)) == 0 {
				out.Header.Del(name)
			}
		}
	}
	traceparent.AddTraceparent(out)
	// The context of out holds the User-Agent data that the visitors added.
	out.Header.Set("User-Agent", transportUserAgent(out.Context(), req.Header.Get("User-Agent")))
}

// visit runs the visitors on req in order and stops at the first error. It
// returns a panic of a visitor as an error: the classic client authenticates
// before it runs the visitors, so they never see a config without an
// authentication type there, and here they can.
func (t *transportRoundTripper) visit(req *http.Request) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("panic: %v", recovered)
		}
	}()
	for _, visitor := range t.visitors {
		if err := visitor(req); err != nil {
			return err
		}
	}
	return nil
}

// transportPrivateLinkResponse returns resp, the answer to sent, unless the
// classic client reports it as an error although its status is below 400.
// That is the login page of a workspace that refused a request from outside
// its private link, which apierr.GetAPIError recognizes by the path and the
// query of the request. In its place the function returns that error as an
// API sends one: with its status, and its error code and message in a JSON
// body. It closes the page unread.
func transportPrivateLinkResponse(sent *http.Request, resp *http.Response) *http.Response {
	// The classic parser reads the URL from the request of the response, which
	// http.Transport sets and another base transport may leave out.
	page := &http.Response{StatusCode: resp.StatusCode, Request: resp.Request}
	if page.Request == nil {
		page.Request = sent
	}
	// It reads a status from 400 as an error response, which stays as it is.
	if page.StatusCode >= 400 || page.Request.URL == nil {
		return resp
	}
	classicErr, ok := apierr.GetAPIError(sent.Context(), common.ResponseWrapper{Response: page}).(*apierr.APIError)
	if !ok {
		return resp
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
	body, _ := json.Marshal(map[string]string{"error_code": classicErr.ErrorCode, "message": classicErr.Message})
	return &http.Response{
		Status:        fmt.Sprintf("%d %s", classicErr.StatusCode, http.StatusText(classicErr.StatusCode)),
		StatusCode:    classicErr.StatusCode,
		Proto:         resp.Proto,
		ProtoMajor:    resp.ProtoMajor,
		ProtoMinor:    resp.ProtoMinor,
		Header:        http.Header{"Content-Type": {"application/json"}},
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Request:       resp.Request,
	}
}

// transportUserAgent returns the classic User-Agent for ctx, followed by the
// segments of modularUserAgent that name a module of the modular SDK.
func transportUserAgent(ctx context.Context, modularUserAgent string) string {
	segments := []string{useragent.FromContext(ctx)}
	for _, segment := range strings.Fields(modularUserAgent) {
		if strings.HasPrefix(segment, transportModularUserAgentPrefix) {
			segments = append(segments, segment)
		}
	}
	return strings.Join(segments, " ")
}

// transportClassicDefault and transportClassicSkipVerify are the transports
// that modular clients share when the config sets no HTTPTransport, without
// and with InsecureSkipVerify. Classic clients share their default as well,
// but each classic client with InsecureSkipVerify makes its own transport.
// Modular clients are built far more often than classic ones, up to one for
// every operation, so a transport for each would cost a TLS handshake per
// client and leave its idle connection open.
var (
	transportClassicDefault    = transportNewClassicDefault()
	transportClassicSkipVerify = transportNewClassicSkipVerify()
)

// transportNewClassicDefault mirrors makeDefaultTransport of the classic SDK
// package httpclient, which is unexported. Keep the two in step when the
// classic SDK is upgraded: TestTransport_MirrorsClassicDefault fails when
// they differ.
func transportNewClassicDefault() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		IdleConnTimeout:       180 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// transportNewClassicSkipVerify mirrors the transport that the classic SDK
// makes for a client with InsecureSkipVerify.
func transportNewClassicSkipVerify() *http.Transport {
	transport := transportNewClassicDefault()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return transport
}

// transportClassicBase picks the transport as the unexported httpTransport
// method of httpclient.ClientConfig does for a classic client.
func transportClassicBase(clientCfg httpclient.ClientConfig) http.RoundTripper {
	if clientCfg.Transport != nil {
		return clientCfg.Transport
	}
	if clientCfg.InsecureSkipVerify {
		return transportClassicSkipVerify
	}
	return transportClassicDefault
}
