package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
)

// classicRequest is how the classic client sends a call. sdkGoTransport sends every
// attempt of the sdk-go call with its method, URL, headers and body.
type classicRequest struct {
	method string
	// path is formatted unescaped from the names the resource passed, as the classic
	// client formats it. sdk-go must request the same unescaped path.
	path string
	// url is the string the classic client parses: path and the query it adds.
	url    string
	header http.Header
	// body, if set, turns the body sdk-go built into the body the classic client sends.
	body func([]byte) ([]byte, error)
	// decode, if set, decodes a success body into the classic response type, whose
	// error the classic client returns when the body doesn't decode.
	decode func([]byte) error
}

// classicVolumeRequest describes a classic Volumes call. Every one accepts JSON, which
// the classic client's response decoding asks for.
func classicVolumeRequest(method, path, url string) classicRequest {
	return classicRequest{method: method, path: path, url: url,
		header: http.Header{"Accept": {"application/json"}}}
}

// useClassicRequest gives r, an sdk-go attempt, the classic URL and body. The classic
// visitors then add the host to the URL, as they do for the classic client.
func (a *sdkGoAttempt) useClassicRequest(r *http.Request) error {
	if r.Method != a.classic.method || r.URL.Path != a.classic.path {
		return fmt.Errorf("sdk-go requested %s %s instead of the classic %s %s",
			r.Method, r.URL.Path, a.classic.method, a.classic.path)
	}
	classicURL := *a.classicURL
	r.URL = &classicURL
	r.Host = ""
	if a.classic.body == nil {
		return nil
	}
	return rewriteBody(r, a.classic.body)
}

// setClassicHeaders sets the headers that the classic client sets after its visitors.
func (a *sdkGoAttempt) setClassicHeaders(r *http.Request) {
	for key := range a.classic.header {
		r.Header.Set(key, a.classic.header.Get(key))
	}
}

// rewriteBody replaces the body of r, its length and GetBody with rewrite's result for
// the body r has.
func rewriteBody(r *http.Request, rewrite func([]byte) ([]byte, error)) error {
	body := r.Body
	if r.GetBody != nil {
		var err error
		if body, err = r.GetBody(); err != nil {
			return err
		}
	}
	var data []byte
	if body != nil {
		var err error
		data, err = io.ReadAll(body)
		body.Close()
		if err != nil {
			return err
		}
	}
	data, err := rewrite(data)
	if err != nil {
		return err
	}
	r.ContentLength = int64(len(data))
	r.GetBody = func() (io.ReadCloser, error) {
		if len(data) == 0 {
			return http.NoBody, nil
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}
	r.Body, _ = r.GetBody()
	return nil
}

// gotSuccess reports whether the last attempt got a success response.
func (a *sdkGoAttempt) gotSuccess() bool {
	return a != nil && a.response != nil && a.response.StatusCode < http.StatusBadRequest
}

// classicDecodeError returns the error the classic client returns when the success body
// of the last response doesn't decode: the body's own decode error, which reports
// classic type names, or the HTML error for a body that starts with "<".
func (a *sdkGoAttempt) classicDecodeError(err error) error {
	if a == nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	if a.classic.decode != nil {
		if classicErr := a.classic.decode(a.responseBody); classicErr != nil {
			err = classicErr
		}
	}
	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) && strings.HasPrefix(string(a.responseBody), "<") {
		return httpclient.ErrHTMLContent
	}
	return fmt.Errorf("failed to unmarshal response body: %w. %s", err,
		classicUnexpectedResponse(a.response, a.requestBody, a.responseBody))
}

// classicUnexpectedResponse mirrors the unexported httpclient.makeUnexpectedResponse.
func classicUnexpectedResponse(resp *http.Response, requestBody,
	responseBody []byte) string {
	var req *http.Request
	if resp != nil {
		req = resp.Request
	}
	log := httplog.RoundTripStringer{
		Request:                  req,
		Response:                 resp,
		RequestBody:              requestBody,
		ResponseBody:             responseBody,
		DebugHeaders:             true,
		DebugTruncateBytes:       10 * 1024,
		DebugAuthorizationHeader: false,
	}
	return "This is likely a bug in the Databricks SDK for Go or the underlying REST " +
		"API. " +
		"Please report this issue with the following debugging information to the SDK " +
		"issue tracker at https://github.com/databricks/databricks-sdk-go/issues. " +
		"Request log:\n```\n" + log.String() + "\n```"
}
