package sdkgoutility

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// transportModularUserAgentPrefix starts the User-Agent segments that name a module of
// the modular SDK, such as sdk-go-core/0.1.0 and sdk-go-modelregistry/0.7.0.
const transportModularUserAgentPrefix = "sdk-go-"

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
//  2. adds the traceparent header unless the request has one;
//  3. sets User-Agent to the classic one for the context that the visitors
//     left, followed by the sdk-go- segments of the User-Agent that the
//     modular SDK set, which mark the call as coming from the modular SDK;
//  4. sets Accept to application/json unless the caller's request has an
//     Accept, as the classic generated calls do.
//
// It then sends the request over cfg.HTTPTransport when that is set, and
// otherwise over a transport with the proxy, connection and TLS settings of
// the classic default, including cfg.InsecureSkipVerify.
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
func Transport(cfg *config.Config) (http.RoundTripper, error) {
	if cfg == nil {
		return nil, errors.New("cannot build a transport for a nil config")
	}
	clientCfg, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &transportRoundTripper{
		base:     transportClassicBase(clientCfg),
		visitors: clientCfg.Visitors,
	}, nil
}

// transportRoundTripper is the http.RoundTripper that Transport returns.
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
	traceparent.AddTraceparent(out)
	// The context of out holds the User-Agent data that the visitors added.
	out.Header.Set("User-Agent", transportUserAgent(out.Context(), req.Header.Get("User-Agent")))
	// The classic generated call sets Accept after the visitors, so the Headers
	// hook can't change it. The modular SDK sends none.
	if req.Header.Get("Accept") == "" {
		out.Header.Set("Accept", "application/json")
	}
	return t.base.RoundTrip(out)
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

// transportClassicDefault is the transport that modular clients share when
// the config sets neither HTTPTransport nor InsecureSkipVerify, as classic
// clients share theirs.
var transportClassicDefault = transportNewClassicDefault()

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

// transportClassicBase mirrors the unexported httpTransport method of
// httpclient.ClientConfig, which picks the transport of a classic client.
func transportClassicBase(clientCfg httpclient.ClientConfig) http.RoundTripper {
	if clientCfg.Transport != nil {
		return clientCfg.Transport
	}
	if clientCfg.InsecureSkipVerify {
		transport := transportNewClassicDefault()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return transport
	}
	return transportClassicDefault
}
