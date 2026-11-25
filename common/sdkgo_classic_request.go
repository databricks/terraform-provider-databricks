package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
