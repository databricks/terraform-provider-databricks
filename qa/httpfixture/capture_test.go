package httpfixture

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// A local HTTP peer exercises wire capture without recording live credentials
// or depending on an API deployment. The captured calls must replay unchanged.
func TestCaptureRoundTripReplaysBodies(t *testing.T) {
	for _, body := range []string{`{"id":9223372036854775807,"items":[null,true,1.5]}`, "plain text", "\x00\xff", "null", ""} {
		t.Run(body, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				data, err := io.ReadAll(req.Body)
				require.NoError(t, err)
				require.Equal(t, body, string(data))
				w.Header().Set("Set-Cookie", "secret-cookie")
				w.Header().Set("Content-Type", "application/octet-stream")
				w.WriteHeader(201)
				_, _ = w.Write(data)
			}))
			defer server.Close()
			capture := NewCaptureTransport(http.DefaultTransport)
			require.NoError(t, capture.StartStep("create"))
			req, err := http.NewRequest("POST", server.URL+"/items/a%2Fb?x=1&x=2", strings.NewReader(body))
			require.NoError(t, err)
			req.Header.Set("Authorization", "Bearer secret-token")
			resp, err := capture.RoundTrip(req)
			require.NoError(t, err)
			data, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			require.Equal(t, body, string(data))
			snapshot, err := capture.Snapshot()
			require.NoError(t, err)
			encoded, err := yaml.Marshal(snapshot)
			require.NoError(t, err)
			require.NotContains(t, string(encoded), "secret-token")
			require.NotContains(t, string(encoded), "secret-cookie")
			handler := NewHandler()
			require.NoError(t, handler.StartStep("replay", snapshot[0].Calls))
			writer := httptest.NewRecorder()
			handler.ServeHTTP(writer, httptest.NewRequest("POST", "/items/a%2Fb?x=1&x=2", strings.NewReader(body)))
			require.Equal(t, 201, writer.Code)
			require.NoError(t, handler.Verify())
		})
	}
}

func TestCaptureMultipartIgnoresBoundaryAndPreservesBinary(t *testing.T) {
	request := func() *http.Request {
		var body bytes.Buffer
		writer := multipart.NewWriter(&body)
		part, err := writer.CreateFormFile("file", "sample.bin")
		require.NoError(t, err)
		_, err = part.Write([]byte{0xff, 0x00, 0x01})
		require.NoError(t, err)
		require.NoError(t, writer.WriteField("path", "/example"))
		require.NoError(t, writer.Close())
		req := httptest.NewRequest("POST", "/upload", &body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		return req
	}
	capture := NewCaptureTransport(roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		body, err := parseRequestBody(req)
		require.NoError(t, err)
		require.Len(t, body.parts, 2)
		return &http.Response{StatusCode: 204, Header: http.Header{}}, nil
	}))
	require.NoError(t, capture.StartStep("upload"))
	_, err := capture.RoundTrip(request())
	require.NoError(t, err)
	steps, err := capture.Snapshot()
	require.NoError(t, err)
	require.NotNil(t, steps[0].Calls[0].Request.Body.Multipart[0].Body.Base64)
	handler := NewHandler()
	require.NoError(t, handler.StartStep("replay", steps[0].Calls))
	handler.ServeHTTP(httptest.NewRecorder(), request())
	require.NoError(t, handler.Verify())
}

func TestCaptureRetainsDispatchOrderAcrossConcurrentReplies(t *testing.T) {
	started, release, done := make(chan struct{}), make(chan struct{}), make(chan error, 1)
	capture := NewCaptureTransport(roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method == "GET" && req.URL.Path == "/first" {
			close(started)
			<-release
		}
		return &http.Response{StatusCode: 204}, nil
	}))
	require.NoError(t, capture.StartStep("concurrent"))
	go func() {
		_, err := capture.RoundTrip(httptest.NewRequest("GET", "/first", nil))
		done <- err
	}()
	<-started
	_, snapshotErr := capture.Snapshot()
	transitionErr := capture.StartStep("too early")
	_, secondErr := capture.RoundTrip(httptest.NewRequest("GET", "/second", nil))
	close(release)
	require.NoError(t, <-done)
	require.ErrorContains(t, snapshotErr, "still in flight")
	require.ErrorContains(t, transitionErr, "still in flight")
	require.NoError(t, secondErr)
	steps, err := capture.Snapshot()
	require.NoError(t, err)
	require.Equal(t, "/first", steps[0].Calls[0].Request.Path)
	require.Equal(t, "/second", steps[0].Calls[1].Request.Path)
	steps[0].Calls[0].Request.Path = "/edited"
	fresh, err := capture.Snapshot()
	require.NoError(t, err)
	require.Equal(t, "/first", fresh[0].Calls[0].Request.Path)
	require.NoError(t, capture.StartStep("destroy"))
	_, err = capture.RoundTrip(httptest.NewRequest("DELETE", "/first", nil))
	require.NoError(t, err)
	steps, err = capture.Snapshot()
	require.NoError(t, err)
	require.Equal(t, "destroy", steps[1].Label)
	require.Equal(t, "DELETE", steps[1].Calls[0].Request.Method)
}

func TestCaptureRejectsIncompleteRecordingsWithoutLeakingErrors(t *testing.T) {
	capture := NewCaptureTransport(roundTripperFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("URL contains secret-token")
	}))
	_, err := capture.RoundTrip(httptest.NewRequest("GET", "/", nil))
	require.ErrorContains(t, err, "requires StartStep")
	require.NoError(t, capture.StartStep("create"))
	_, err = capture.RoundTrip(httptest.NewRequest("GET", "/", nil))
	require.Error(t, err)
	_, err = capture.Snapshot()
	require.ErrorContains(t, err, "create: could not capture call 1")
	require.NotContains(t, err.Error(), "secret-token")
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// Unified hosts distinguish a bodyless HTTP/2 GET from an unknown-length body.
// Capturing must retain the SDK's http.NoBody sentinel and wire framing.
func TestCapturePreservesBodylessRequestFraming(t *testing.T) {
	length := make(chan int64, 1)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		length <- req.ContentLength
		w.WriteHeader(http.StatusNoContent)
	}))
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()
	capture := NewCaptureTransport(server.Client().Transport)
	require.NoError(t, capture.StartStep("read"))
	req, err := http.NewRequest(http.MethodGet, server.URL+"?path=%2FShared", strings.NewReader(""))
	require.NoError(t, err)
	response, err := capture.RoundTrip(req)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	require.Equal(t, "HTTP/2.0", response.Proto)
	require.Equal(t, int64(0), <-length)
	require.NoError(t, capture.Verify())
}
