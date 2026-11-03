package httpfixture

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestCaptureBodyRespectsContentType(t *testing.T) {
	const formattedJSON = "{ \"b\":2, \"a\":1 }\n"
	for _, tc := range []struct {
		name        string
		contentType string
		body        string
		kind        string
	}{
		{"json", "application/json", formattedJSON, "json"},
		{"json_parameters", "Application/JSON; charset=utf-8", formattedJSON, "json"},
		{"scim", "application/scim+json; charset=utf-8", formattedJSON, "json"},
		{"problem", "application/problem+json", formattedJSON, "json"},
		{"large_id", "application/json", `{"id":9223372036854775807}`, "json"},
		{"large_unsigned_id", "application/json", `{"id":18446744073709551615}`, "json"},
		{"json_null", "application/json", "null", "json"},
		{"raw_download", "application/octet-stream", formattedJSON, "text"},
		{"plain_text", "text/plain", formattedJSON, "text"},
		{"missing_type", "", formattedJSON, "text"},
		{"malformed_type", "application/json; charset", formattedJSON, "text"},
		{"non_application_type", "text/example+json", formattedJSON, "text"},
		{"malformed_json", "application/json", "{broken\n", "text"},
		{"trailing_data", "application/json", "null null", "text"},
		{"binary", "application/octet-stream", "\x00\xff", "base64"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				data, err := io.ReadAll(req.Body)
				require.NoError(t, err)
				require.Equal(t, tc.body, string(data))
				// Suppress automatic sniffing when the response has no media type.
				w.Header()["Content-Type"] = nil
				if tc.contentType != "" {
					w.Header().Set("Content-Type", tc.contentType)
				}
				_, _ = w.Write(data)
			}))
			defer server.Close()
			capture := NewCaptureTransport(http.DefaultTransport)
			require.NoError(t, capture.StartStep("capture"))
			req, err := http.NewRequest(http.MethodPost, server.URL+"/body", strings.NewReader(tc.body))
			require.NoError(t, err)
			if tc.contentType != "" {
				req.Header.Set("Content-Type", tc.contentType)
			}
			resp, err := capture.RoundTrip(req)
			require.NoError(t, err)
			data, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			require.Equal(t, tc.body, string(data))

			steps, err := capture.Snapshot()
			require.NoError(t, err)
			encoded, err := yaml.Marshal(steps)
			require.NoError(t, err)
			require.NoError(t, yaml.Unmarshal(encoded, &steps))
			for _, body := range []*Body{steps[0].Calls[0].Request.Body, steps[0].Calls[0].Response.Body} {
				kind, err := body.kind()
				require.NoError(t, err)
				require.Equal(t, tc.kind, kind)
			}
			handler := NewHandler()
			require.NoError(t, handler.StartStep("replay", steps[0].Calls))
			replayBody := tc.body
			if tc.body == formattedJSON && tc.kind == "json" {
				replayBody = `{"a":1,"b":2}`
			}
			replay := httptest.NewRequest(http.MethodPost, "/body", strings.NewReader(replayBody))
			replay.Header = req.Header.Clone()
			writer := httptest.NewRecorder()
			handler.ServeHTTP(writer, replay)
			require.Equal(t, http.StatusOK, writer.Code)
			require.Equal(t, replayBody, writer.Body.String())
			require.NoError(t, handler.Verify())
		})
	}
}

func TestCaptureRetainsRedirectLocations(t *testing.T) {
	for _, absolute := range []bool{false, true} {
		name := "relative"
		if absolute {
			name = "absolute"
		}
		t.Run(name, func(t *testing.T) {
			location := "/result?download=1"
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.URL.Path == "/redirect" {
					w.Header().Set("Location", location)
					w.WriteHeader(http.StatusFound)
					return
				}
				require.Equal(t, "/result?download=1", req.URL.RequestURI())
				w.Header().Set("Content-Type", "text/plain")
				_, _ = io.WriteString(w, "downloaded bytes\n")
			}))
			defer server.Close()
			if absolute {
				location = server.URL + location
			}
			capture := NewCaptureTransport(http.DefaultTransport)
			require.NoError(t, capture.StartStep("redirect"))
			client := &http.Client{Transport: capture}
			resp, err := client.Get(server.URL + "/redirect")
			require.NoError(t, err)
			data, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			require.Equal(t, http.StatusOK, resp.StatusCode)
			require.Equal(t, "downloaded bytes\n", string(data))
			steps, err := capture.Snapshot()
			require.NoError(t, err)
			require.Len(t, steps[0].Calls, 2)
			require.Equal(t, location, steps[0].Calls[0].Response.Headers["Location"])
		})
	}
}
