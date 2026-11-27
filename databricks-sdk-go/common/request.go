package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Represents a request body.
//
// If the provided request data is an io.Reader, DebugBytes is set to
// "<io.Reader>". Otherwise, DebugBytes is set to the marshaled JSON
// representation of the request data, and ReadCloser is set to a new
// io.ReadCloser that reads from DebugBytes.
//
// Request bodies are never closed by the client, hence only accepting
// io.Reader.
type RequestBody struct {
	Reader      io.Reader
	ContentType string
	DebugBytes  []byte
}

func NewRequestBody(data any) (RequestBody, error) {
	switch v := data.(type) {
	case io.Reader:
		return RequestBody{
			Reader:     v,
			DebugBytes: []byte("<io.Reader>"),
		}, nil
	case string:
		return RequestBody{
			Reader:     strings.NewReader(v),
			DebugBytes: []byte(v),
		}, nil
	case []byte:
		return RequestBody{
			Reader:     bytes.NewReader(v),
			DebugBytes: v,
		}, nil
	default:
		bs, err := json.Marshal(data)
		if err != nil {
			return RequestBody{}, fmt.Errorf("request marshal failure: %w", err)
		}
		return RequestBody{
			Reader:      bytes.NewReader(bs),
			ContentType: "application/json",
			DebugBytes:  bs,
		}, nil
	}
}

// Reset a request body to its initial state.
//
// This is used to retry requests with a body that has already been read.
// If the request body is not resettable (i.e. not nil and of type other than
// strings.Reader or bytes.Reader), this will return an error.
func (r RequestBody) Reset() error {
	if r.Reader == nil {
		return nil
	}
	if v, ok := r.Reader.(io.Seeker); ok {
		_, err := v.Seek(0, io.SeekStart)
		return err
	} else {
		return fmt.Errorf("cannot reset reader of type %T", r.Reader)
	}
}
