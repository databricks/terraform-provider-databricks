package provider

import (
	"errors"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
)

type mockProviderConfig struct {
	accountClient      *databricks.AccountClient
	accountClientErr   error
	workspaceClient    *databricks.WorkspaceClient
	workspaceClientErr error
}

func (m *mockProviderConfig) AccountClient() (*databricks.AccountClient, error) {
	return m.accountClient, m.accountClientErr
}

func (m *mockProviderConfig) WorkspaceClient() (*databricks.WorkspaceClient, error) {
	return m.workspaceClient, m.workspaceClientErr
}

func TestAccountClient(t *testing.T) {
	client := &databricks.AccountClient{}

	// wantErrSubs is a slice so the wrap-and-cause case can assert that the
	// diagnostic contains both the wrap prefix (in the summary) and the
	// original error (in the detail). Empty slice means no error expected.
	tests := []struct {
		name         string
		providerData any
		wantClient   *databricks.AccountClient
		wantErrSubs  []string
	}{
		{
			name:         "nil ProviderData returns nil with no diagnostics",
			providerData: nil,
			wantClient:   nil,
		},
		{
			name:         "ProviderData not implementing providerConfig surfaces an error diagnostic",
			providerData: 42, // not a providerConfig
			wantClient:   nil,
			wantErrSubs:  []string{"unexpected provider data type"},
		},
		{
			name:         "providerConfig.AccountClient success passes the client through",
			providerData: &mockProviderConfig{accountClient: client},
			wantClient:   client,
		},
		{
			name: "providerConfig.AccountClient error wraps the cause in a diagnostic",
			providerData: &mockProviderConfig{
				accountClientErr: errors.New("test-error"),
			},
			wantClient:  nil,
			wantErrSubs: []string{"failed to resolve account client", "test-error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, gotDiags := AccountClient(tt.providerData)

			if gotClient != tt.wantClient {
				t.Errorf("client: got %p, want %p", gotClient, tt.wantClient)
			}
			if len(tt.wantErrSubs) == 0 {
				if gotDiags.HasError() {
					t.Errorf("unexpected error diagnostics: %v", gotDiags)
				}
				return
			}
			if !gotDiags.HasError() {
				t.Fatalf("expected error diagnostic, got none")
			}
			msg := gotDiags.Errors()[0].Summary() + " " + gotDiags.Errors()[0].Detail()
			for _, want := range tt.wantErrSubs {
				if !strings.Contains(msg, want) {
					t.Errorf("diagnostic %q does not contain %q", msg, want)
				}
			}
		})
	}
}

func TestWorkspaceClient(t *testing.T) {
	client := &databricks.WorkspaceClient{}

	tests := []struct {
		name         string
		providerData any
		wantClient   *databricks.WorkspaceClient
		wantErrSubs  []string
	}{
		{
			name:         "nil ProviderData returns nil with no diagnostics",
			providerData: nil,
			wantClient:   nil,
		},
		{
			name:         "ProviderData not implementing providerConfig surfaces an error diagnostic",
			providerData: 42, // not a providerConfig
			wantClient:   nil,
			wantErrSubs:  []string{"unexpected provider data type"},
		},
		{
			name:         "providerConfig.WorkspaceClient success passes the client through",
			providerData: &mockProviderConfig{workspaceClient: client},
			wantClient:   client,
		},
		{
			name: "providerConfig.WorkspaceClient error wraps the cause in a diagnostic",
			providerData: &mockProviderConfig{
				workspaceClientErr: errors.New("test-error"),
			},
			wantClient:  nil,
			wantErrSubs: []string{"failed to resolve workspace client", "test-error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, gotDiags := WorkspaceClient(tt.providerData)

			if gotClient != tt.wantClient {
				t.Errorf("client: got %p, want %p", gotClient, tt.wantClient)
			}
			if len(tt.wantErrSubs) == 0 {
				if gotDiags.HasError() {
					t.Errorf("unexpected error diagnostics: %v", gotDiags)
				}
				return
			}
			if !gotDiags.HasError() {
				t.Fatalf("expected error diagnostic, got none")
			}
			msg := gotDiags.Errors()[0].Summary() + " " + gotDiags.Errors()[0].Detail()
			for _, want := range tt.wantErrSubs {
				if !strings.Contains(msg, want) {
					t.Errorf("diagnostic %q does not contain %q", msg, want)
				}
			}
		})
	}
}
