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
	// response, if set, returns the classic type that the classic client decodes a
	// success body into. It is unset if the classic client decodes none, as for a delete.
	response func() any
}

// classicVolumeRequest describes a classic Volumes call. Every one accepts JSON, which
// the classic client's response decoding asks for.
func classicVolumeRequest(method, path, url string) classicRequest {
	return classicRequest{method: method, path: path, url: url,
		header: http.Header{"Accept": {"application/json"}}}
}

// classicVolumeBody turns the JSON object that sdk-go built into the classic body,
// without the fields in drop. The classic Volumes structs list their fields
// alphabetically, so the classic client marshals the keys sorted, as json.Marshal does
// for a map.
func classicVolumeBody(body []byte, drop ...string) ([]byte, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, err
	}
	for _, field := range drop {
		delete(fields, field)
	}
	return json.Marshal(fields)
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

// sdkGoResponse returns resp, whose body is body, as sdk-go must read it to read it as
// the classic client does. The classic client reads every status below 400 as a
// success, so such a response reaches sdk-go as 200 OK, with the classic encoding of
// its body. Like the classic client, a body that the classic type doesn't decode fails,
// which classicError reports with the classic error and never retries.
func (a *sdkGoAttempt) sdkGoResponse(resp *http.Response,
	body []byte) (*http.Response, error) {
	if resp.StatusCode >= http.StatusBadRequest {
		return resp, nil
	}
	body, err := a.classicResponseBody(body)
	if err != nil {
		return nil, err
	}
	// The attempt keeps resp, whose status the classic request log shows.
	success := *resp
	success.StatusCode, success.Status = http.StatusOK, "200 OK"
	success.Body = io.NopCloser(bytes.NewReader(body))
	success.ContentLength = int64(len(body))
	return &success, nil
}

// classicResponseBody decodes body, a success body, into the classic response type, as
// httpclient.WithResponseUnmarshal does, and returns that value's encoding. sdk-go then
// reads the values the classic client reads, such as "7" for a string field that the
// response sets to the number 7. Like the classic client, an empty body is an empty
// response. It returns body unchanged for a call that decodes no response.
func (a *sdkGoAttempt) classicResponseBody(body []byte) ([]byte, error) {
	if a.classic.response == nil {
		return body, nil
	}
	response := a.classic.response()
	if len(body) > 0 {
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
	}
	return json.Marshal(response)
}

// classicDecodeError returns the error the classic client returns when the success body
// of the last response doesn't decode: the body's own decode error, which reports
// classic type names, or the HTML error for a body that starts with "<".
func (a *sdkGoAttempt) classicDecodeError(err error) error {
	if a == nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	if _, classicErr := a.classicResponseBody(a.responseBody); classicErr != nil {
		err = classicErr
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
