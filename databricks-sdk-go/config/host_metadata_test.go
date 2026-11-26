package config

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

const (
	testHMHost        = "https://dummy-workspace.databricks.com"
	testHMAccHost     = "https://dummy-accounts.databricks.com"
	testHMAccountID   = "00000000-0000-0000-0000-000000000001"
	testHMWorkspaceID = "111111111111111"
)

func newTestAPIClient(transport fixtures.MappingTransport) *httpclient.ApiClient {
	return httpclient.NewApiClient(httpclient.ClientConfig{Transport: transport})
}

func TestGetHostMetadata_WorkspaceStaticOIDCEndpoint(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 200,
			Response: map[string]string{
				"oidc_endpoint": testHMHost + "/oidc",
				"account_id":    testHMAccountID,
				"workspace_id":  testHMWorkspaceID,
				"cloud":         "AWS",
				"host_type":     "workspace",
			},
		},
	})
	meta, err := getHostMetadata(context.Background(), testHMHost, client)
	if err != nil {
		t.Fatal(err)
	}
	want := &HostMetadata{
		OIDCEndpoint: testHMHost + "/oidc",
		AccountID:    testHMAccountID,
		WorkspaceID:  testHMWorkspaceID,
		Cloud:        "AWS",
		HostType:     WorkspaceHost,
	}
	if diff := cmp.Diff(want, meta); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGetHostMetadata_AccountHostRawOIDCTemplate(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 200,
			Response: map[string]string{
				"oidc_endpoint": testHMAccHost + "/oidc/accounts/{account_id}",
			},
		},
	})
	meta, err := getHostMetadata(context.Background(), testHMAccHost, client)
	if err != nil {
		t.Fatal(err)
	}
	want := &HostMetadata{
		OIDCEndpoint: testHMAccHost + "/oidc/accounts/{account_id}",
	}
	if diff := cmp.Diff(want, meta); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGetHostMetadata_HTTPError(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 404,
		},
	})
	_, err := getHostMetadata(context.Background(), testHMHost, client)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "fetching host metadata from") {
		t.Errorf("expected error containing %q, got %q", "fetching host metadata from", err.Error())
	}
}

func TestGetHostMetadata_WithHostTypeField(t *testing.T) {
	tests := []struct {
		name         string
		hostType     string
		wantHostType HostType
	}{
		{
			name:         "workspace",
			hostType:     "workspace",
			wantHostType: WorkspaceHost,
		},
		{
			name:         "account",
			hostType:     "account",
			wantHostType: AccountHost,
		},
		{
			name:         "unified",
			hostType:     "unified",
			wantHostType: UnifiedHost,
		},
		{
			name:         "Workspace uppercase",
			hostType:     "WORKSPACE",
			wantHostType: WorkspaceHost,
		},
		{
			name:         "missing host_type field",
			hostType:     "",
			wantHostType: HostTypeUnknown,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := map[string]string{
				"oidc_endpoint": testHMHost + "/oidc",
				"account_id":    testHMAccountID,
			}
			if tc.hostType != "" {
				response["host_type"] = tc.hostType
			}
			client := newTestAPIClient(fixtures.MappingTransport{
				"GET /.well-known/databricks-config": {
					Status:   200,
					Response: response,
				},
			})
			meta, err := getHostMetadata(context.Background(), testHMHost, client)
			if err != nil {
				t.Fatal(err)
			}
			if meta.HostType != tc.wantHostType {
				t.Errorf("HostType field mismatch: got %q, want %q", meta.HostType, tc.wantHostType)
			}
		})
	}
}

func TestGetHostMetadata_WithTokenFederationDefaultOIDCAudiences(t *testing.T) {
	tests := []struct {
		name          string
		audiences     []string
		wantAudiences []string
	}{
		{
			name:          "workspace audience",
			audiences:     []string{testHMHost + "/oidc/v1/token"},
			wantAudiences: []string{testHMHost + "/oidc/v1/token"},
		},
		{
			name:          "account audience",
			audiences:     []string{testHMAccountID},
			wantAudiences: []string{testHMAccountID},
		},
		{
			name:          "missing field",
			audiences:     nil,
			wantAudiences: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := map[string]any{
				"oidc_endpoint": testHMHost + "/oidc",
				"account_id":    testHMAccountID,
			}
			if tc.audiences != nil {
				response["token_federation_default_oidc_audiences"] = tc.audiences
			}
			client := newTestAPIClient(fixtures.MappingTransport{
				"GET /.well-known/databricks-config": {
					Status:   200,
					Response: response,
				},
			})
			meta, err := getHostMetadata(context.Background(), testHMHost, client)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.wantAudiences, meta.TokenFederationDefaultOIDCAudiences); diff != "" {
				t.Errorf("TokenFederationDefaultOIDCAudiences mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetHostMetadata_WithCloudField(t *testing.T) {
	tests := []struct {
		name      string
		cloud     string
		wantCloud environment.Cloud
	}{
		{
			name:      "AWS",
			cloud:     "AWS",
			wantCloud: environment.CloudAWS,
		},
		{
			name:      "Azure",
			cloud:     "Azure",
			wantCloud: environment.CloudAzure,
		},
		{
			name:      "GCP",
			cloud:     "GCP",
			wantCloud: environment.CloudGCP,
		},
		{
			name:      "missing cloud field",
			cloud:     "",
			wantCloud: environment.CloudUnknown,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := map[string]string{
				"oidc_endpoint": testHMHost + "/oidc",
				"account_id":    testHMAccountID,
			}
			if tc.cloud != "" {
				response["cloud"] = tc.cloud
			}
			client := newTestAPIClient(fixtures.MappingTransport{
				"GET /.well-known/databricks-config": {
					Status:   200,
					Response: response,
				},
			})
			meta, err := getHostMetadata(context.Background(), testHMHost, client)
			if err != nil {
				t.Fatal(err)
			}
			if meta.Cloud != tc.wantCloud {
				t.Errorf("Cloud field mismatch: got %q, want %q", meta.Cloud, tc.wantCloud)
			}
		})
	}
}
