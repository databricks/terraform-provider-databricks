package oidc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

// defaultEnv is the default environment variables that must be set to
// successfully create an Azure DevOps ID token source.
var defaultEnv = map[string]string{
	"SYSTEM_ACCESSTOKEN":                 "test-access-token",
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI": "https://dev.azure.com/myorg",
	"SYSTEM_PLANID":                      "test-plan-id",
	"SYSTEM_JOBID":                       "test-job-id",
	"SYSTEM_TEAMPROJECTID":               "test-team-project-id",
	"SYSTEM_HOSTTYPE":                    "test-host-type",
}

func TestNewAzureDevOpsIDTokenSource_success(t *testing.T) {
	for k, v := range defaultEnv {
		t.Setenv(k, v)
	}

	want := &azureDevOpsIDTokenSource{
		AccessToken:                 "test-access-token",
		TeamFoundationCollectionURI: "https://dev.azure.com/myorg",
		PlanID:                      "test-plan-id",
		JobID:                       "test-job-id",
		TeamProjectID:               "test-team-project-id",
		HostType:                    "test-host-type",
	}

	got, gotErr := NewAzureDevOpsIDTokenSource(nil)
	if gotErr != nil {
		t.Fatalf("NewAzureDevOpsIDTokenSource() want no error, got error: %v", gotErr)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewAzureDevOpsIDTokenSource() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewAzureDevOpsIDTokenSource_missingEnv(t *testing.T) {
	// Guarantee that the tests are run in a deterministic order.
	sortedKeys := make([]string, 0, len(defaultEnv))
	for k := range defaultEnv {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, env := range sortedKeys {
		t.Run(fmt.Sprintf("missing %s", env), func(t *testing.T) {
			// Set all env vars except the missing one.
			for k, v := range defaultEnv {
				if k != env {
					t.Setenv(k, v)
				}
			}

			wantErr := errNotCallingFromAzureDevOps
			if env == "SYSTEM_ACCESSTOKEN" {
				wantErr = errMissingAccessToken
			}

			got, gotErr := NewAzureDevOpsIDTokenSource(nil)

			if !errors.Is(gotErr, wantErr) {
				t.Fatalf("NewAzureDevOpsIDTokenSource() want error %v, got %v", wantErr, gotErr)
			}
			if got != nil {
				t.Errorf("NewAzureDevOpsIDTokenSource() want nil, got %v", got)
			}
		})
	}
}

func TestNewAzureDevOpsIDTokenSource_IDToken(t *testing.T) {
	// Set up default environment variables for all test cases.
	for k, v := range defaultEnv {
		t.Setenv(k, v)
	}

	testCases := []struct {
		desc          string
		httpTransport http.RoundTripper
		want          *IDToken
		wantErr       bool
	}{
		{
			desc: "server error response",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/test-team-project-id/_apis/distributedtask/hubs/test-host-type/plans/test-plan-id/jobs/test-job-id/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer test-access-token",
						"Accept":        "application/json",
					},
				},
			},
			wantErr: true,
		},
		{
			desc: "success",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/test-team-project-id/_apis/distributedtask/hubs/test-host-type/plans/test-plan-id/jobs/test-job-id/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer test-access-token",
						"Accept":        "application/json",
					},
					Response: `{"oidcToken": "azure-devops-id-token-42"}`,
				},
			},
			want: &IDToken{
				Value: "azure-devops-id-token-42",
			},
			wantErr: false,
		},
		{
			desc: "empty oidc token in response",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/test-team-project-id/_apis/distributedtask/hubs/test-host-type/plans/test-plan-id/jobs/test-job-id/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer test-access-token",
						"Accept":        "application/json",
					},
					Response: `{"oidcToken": ""}`,
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			client := httpclient.NewApiClient(httpclient.ClientConfig{
				Transport: tc.httpTransport,
			})

			tokenSource, err := NewAzureDevOpsIDTokenSource(client)
			if err != nil {
				t.Fatalf("NewAzureDevOpsIDTokenSource() failed: %v", err)
			}

			got, gotErr := tokenSource.IDToken(context.Background(), "test-audience")

			if tc.wantErr && gotErr == nil {
				t.Errorf("IDToken(): expected error, got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Errorf("IDToken(): got error %q, want none", gotErr)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("IDToken(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
