package terraformtest

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestRecordingNormalizesRedirectLocations(t *testing.T) {
	const host = "workspace.databricks.com"
	const absolute = "https://" + host + "/result"
	for _, tc := range []struct {
		name     string
		location string
		want     string
	}{
		{"absolute", absolute, "/result"},
		{"escaped", "https://" + host + "/a%2Fb?x=%2F#part%2Fone", "/a%2Fb?x=%2F#part%2Fone"},
		{"bare_host", "https://" + host, "/"},
		{"root", "https://" + host + "?x=1", "/?x=1"},
		{"empty_query", "https://" + host + "/result?", "/result?"},
		{"host_case", "https://WORKSPACE.DATABRICKS.COM/result", "/result"},
		{"network_path", "//" + host + "/result", "/result"},
		{"double_slash", "https://" + host + "//result", "/.//result"},
		{"relative", "../result?x=1#part", "../result?x=1#part"},
		{"other_host", "https://other.databricks.com/result", "https://other.databricks.com/result"},
		{"subdomain", "https://other." + host + "/result", "https://other." + host + "/result"},
		{"other_port", "https://" + host + ":8443/result", "https://" + host + ":8443/result"},
		{"other_scheme", "ftp://" + host + "/result", "ftp://example.databricks.com/result"},
		{"malformed", "https://" + host + "/%zz", "https://example.databricks.com/%zz"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			n, err := newRecordingNormalizer(&testDefinition{}, &config.Config{Host: "https://" + host})
			require.NoError(t, err)
			calls := []httpfixture.Call{{
				Request: httpfixture.Request{Method: "GET", Path: "/start", Headers: map[string]string{"Location": absolute}},
				Response: httpfixture.Response{Status: http.StatusFound, Headers: map[string]string{"lOcAtIoN": tc.location},
					Body: &httpfixture.Body{JSON: map[string]any{"response": map[string]any{
						"headers": map[string]any{"Location": absolute},
					}}}},
			}}
			encoded, err := yaml.Marshal(calls)
			require.NoError(t, err)
			encoded, err = n.transform(encoded, false, true)
			require.NoError(t, err)
			require.NoError(t, yaml.Unmarshal(encoded, &calls))
			require.Equal(t, tc.want, calls[0].Response.Headers["lOcAtIoN"])
			require.Equal(t, "https://example.databricks.com/result", calls[0].Request.Headers["Location"])
			require.Equal(t, map[string]any{"response": map[string]any{
				"headers": map[string]any{"Location": "https://example.databricks.com/result"},
			}}, calls[0].Response.Body.JSON)
		})
	}
}

func TestRecordingRedirectsReplayOnLocalServer(t *testing.T) {
	for _, mode := range []string{"relative", "absolute", "network_path", "double_slash"} {
		t.Run(mode, func(t *testing.T) {
			target := "/result/a%2Fb?download=1&key=%2F#section%2Fone"
			if mode == "double_slash" {
				target = "/" + target
			}
			location := target
			peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.URL.Path == "/start" {
					w.Header().Set("Location", location)
					w.WriteHeader(http.StatusFound)
					return
				}
				require.Equal(t, strings.Split(target, "#")[0], req.URL.RequestURI())
				w.Header().Set("Content-Type", "text/plain")
				_, _ = io.WriteString(w, "downloaded bytes\n")
			}))
			defer peer.Close()
			if mode != "relative" {
				location = peer.URL + target
				if mode == "network_path" {
					location = strings.TrimPrefix(location, "http:")
				}
			}
			capture := httpfixture.NewCaptureTransport(http.DefaultTransport)
			require.NoError(t, capture.StartStep(stepLabel(0, "create")))
			client := &http.Client{Transport: capture}
			response, err := client.Get(peer.URL + "/start")
			require.NoError(t, err)
			data, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			require.NoError(t, response.Body.Close())
			require.Equal(t, http.StatusOK, response.StatusCode)
			require.Equal(t, "downloaded bytes\n", string(data))
			require.NoError(t, capture.StartStep(stepLabel(1, "destroy")))
			captured, err := capture.Snapshot()
			require.NoError(t, err)
			require.Len(t, captured[0].Calls, 2)
			require.Equal(t, location, captured[0].Calls[0].Response.Headers["Location"])

			definition, err := decodeDefinition(strings.NewReader(minimalDefinition))
			require.NoError(t, err)
			normalizer, err := newRecordingNormalizer(definition, &config.Config{Host: peer.URL})
			require.NoError(t, err)
			updated, err := refreshRecording([]byte(minimalDefinition), definition, captured, normalizer)
			require.NoError(t, err)
			recorded, err := decodeDefinition(bytes.NewReader(updated))
			require.NoError(t, err)
			handler := httpfixture.NewHandler()
			require.NoError(t, handler.StartStep("replay", recorded.Steps[0].Calls))
			replay := httptest.NewServer(handler)
			defer replay.Close()
			require.NotEqual(t, peer.URL, replay.URL)
			peer.Close()

			cfg := &config.Config{}
			require.NoError(t, fixtureConfig(replay.URL)(cfg))
			client.Transport = cfg.HTTPTransport
			replayed, err := client.Get(replay.URL + "/start")
			require.NoError(t, err)
			replayedData, err := io.ReadAll(replayed.Body)
			require.NoError(t, err)
			require.NoError(t, replayed.Body.Close())
			require.Equal(t, http.StatusOK, replayed.StatusCode)
			require.Equal(t, data, replayedData)
			require.Equal(t, response.Request.URL.RequestURI(), replayed.Request.URL.RequestURI())
			require.Equal(t, response.Request.URL.Fragment, replayed.Request.URL.Fragment)
			require.NoError(t, handler.Verify())
		})
	}
}
