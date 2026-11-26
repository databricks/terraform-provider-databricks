package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Represents a response.
//
// Responses must always be closed. For non-streaming responses, they are closed
// during deserialization in the client (see unmarshall()). For streaming
// responses, they are returned to the caller, who is responsible for closing
// them. When debugging is enabled, the value of DebugBytes is printed when
// logging the response body.
type ResponseWrapper struct {
	// ReadCloser is the response body that was received from the server.
	ReadCloser io.ReadCloser

	// DebugBytes is the raw response body that was received from the server, if
	// the response was not streamed.
	DebugBytes []byte

	// Response is the original http.Response object.
	Response *http.Response

	// RequestBody is the request body that was sent to the server.
	RequestBody RequestBody
}

func makeResponseWrapper(data any, response *http.Response, req RequestBody) (ResponseWrapper, error) {
	switch v := data.(type) {
	case io.ReadCloser:
		return ResponseWrapper{
			ReadCloser:  v,
			DebugBytes:  []byte("<Streaming response>"),
			Response:    response,
			RequestBody: req,
		}, nil
	case []byte:
		// We need to set the response body to a new ReadCloser, because the
		// original response body has already been read and closed.
		response.Body = io.NopCloser(bytes.NewReader(v))
		return ResponseWrapper{
			ReadCloser:  io.NopCloser(bytes.NewReader(v)),
			DebugBytes:  v,
			Response:    response,
			RequestBody: req,
		}, nil
	default:
		return ResponseWrapper{}, errors.New("newResponseBody can only be called with io.ReadCloser or []byte")
	}
}

// NewResponseWrapper creates a new ResponseWrapper from an http.Response.
//
// If the response is nil, or the response body is nil, this will return an
// error.
//
// If the response is a streaming response, the response body will be returned
// as-is. Otherwise, the response body will be read into memory and returned
// as a byte slice. Streaming responses are identified by the "Accept" request
// header being set to anything other than "application/json" and the
// "Content-Type" header being set to "application/json".
func NewResponseWrapper(r *http.Response, req RequestBody) (ResponseWrapper, error) {
	if r == nil {
		return ResponseWrapper{}, fmt.Errorf("nil response")
	}
	if r.Request == nil {
		return ResponseWrapper{}, fmt.Errorf("nil request")
	}
	// JSON media types might have more information after `application/json`, like
	// `application/json;odata.metadata=minimal;odata...=...`.
	// See https://www.rfc-editor.org/rfc/rfc9110#section-8.3.1
	isJSON := strings.HasPrefix(r.Header.Get("Content-Type"), "application/json")
	streamResponse := r.Request.Header.Get("Accept") != "application/json" && !isJSON
	if streamResponse {
		return makeResponseWrapper(r.Body, r, req)
	}
	defer r.Body.Close()
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return ResponseWrapper{}, fmt.Errorf("response body: %w", err)
	}
	return makeResponseWrapper(bs, r, req)
}
