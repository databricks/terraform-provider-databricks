package config

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
)

// newGroupTokenServer creates the OAuth server shared by the M2M and WIF
// cache-isolation tests and returns its token endpoint call counter.
func newGroupTokenServer(t *testing.T) (*httptest.Server, *int) {
	t.Helper()

	tokenCalls := 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch req.URL.Path {
		case "/.well-known/databricks-config":
			// Config probes host metadata before OAuth discovery. These cache tests do
			// not provide metadata, so return 404. Metadata lookup failures are non-fatal,
			// and the fake URL is inferred to be a workspace host because it does not
			// match an account-host URL pattern.
			http.NotFound(w, req)
		case "/oidc/.well-known/oauth-authorization-server":
			if err := json.NewEncoder(w).Encode(u2m.OAuthAuthorizationServer{TokenEndpoint: server.URL + "/oidc/v1/token"}); err != nil {
				t.Errorf("Could not write the OAuth discovery response: %v", err)
			}
		case "/oidc/v1/token":
			tokenCalls++

			if err := req.ParseForm(); err != nil {
				t.Fatalf("Could not parse the token request: %v", err)
			}

			groupID := req.Form.Get("assume_group")
			if groupID == "" {
				groupID = "normal"
			}

			if err := json.NewEncoder(w).Encode(map[string]any{
				"access_token": "token-" + groupID,
				"token_type":   "Bearer",
				"expires_in":   3600,
			}); err != nil {
				t.Errorf("Could not write the token response: %v", err)
			}
		default:
			t.Errorf("Received unexpected request %s %s", req.Method, req.URL.Path)
			http.NotFound(w, req)
		}
	}))
	t.Cleanup(server.Close)

	return server, &tokenCalls
}
