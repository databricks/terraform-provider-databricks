package config

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
)

// Input: a service account to impersonate on a GCP workspace host, with the
// IAM credentials API served by the fixture.
//
// Want: generateIdToken for the workspace host with the email included,
// generateAccessToken for the cloud-platform and compute scopes, and a request
// that carries the ID token as the bearer token and the access token in
// X-Databricks-GCP-SA-Access-Token.
func TestGoogleDefaultCredentials_Configure_impersonatesTheServiceAccount(t *testing.T) {
	env.CleanupEnvironment(t)
	const (
		host           = "https://123.4.gcp.databricks.com"
		serviceAccount = "sa@project.iam.gserviceaccount.com"
		// idToken is an unsigned JWT that expires in 2100. Google's token
		// sources only decode an ID token to read its expiry, so they never
		// refresh this one.
		idToken = "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjQxMDI0NDQ4MDB9.c2ln"
	)
	iamPath := "/v1/projects/-/serviceAccounts/" + serviceAccount
	iam := fixtures.MappingTransport{
		"POST " + iamPath + ":generateIdToken": {
			ExpectedRequest: map[string]any{"audience": host, "includeEmail": true},
			Response:        `{"token": "` + idToken + `"}`,
		},
		"POST " + iamPath + ":generateAccessToken": {
			ExpectedRequest: map[string]any{
				"lifetime": "3600s",
				"scope":    []string{"https://www.googleapis.com/auth/cloud-platform", "https://www.googleapis.com/auth/compute"},
			},
			Response: `{"accessToken": "sa-access-token", "expireTime": "2100-01-01T00:00:00Z"}`,
		},
	}
	cfg := &Config{
		Host:                 host,
		GoogleServiceAccount: serviceAccount,
		Credentials: GoogleDefaultCredentials{
			opts: []option.ClientOption{option.WithHTTPClient(&http.Client{Transport: iam})},
		},
		HostMetadataResolver: func(context.Context, string) (*HostMetadata, error) {
			return nil, nil
		},
		HTTPTransport: fixtures.SliceTransport{},
	}

	req, err := authenticateRequest(cfg)

	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"Authorization":                    "Bearer " + idToken,
		"X-Databricks-Gcp-Sa-Access-Token": "sa-access-token",
	}
	got := map[string]string{}
	for name := range req.Header {
		got[name] = req.Header.Get(name)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("request headers mismatch (-want +got):\n%s", diff)
	}
}
