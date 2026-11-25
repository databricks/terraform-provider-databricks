package provisioning

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestWorkspacesAPI_waitersReturnTheWorkspaceAfterCreateUpdateAndDelete(t *testing.T) {
	const accountID = "test-account"
	createRequest := CreateWorkspaceRequest{
		WorkspaceName:          "test-workspace",
		AwsRegion:              "us-east-1",
		CredentialsId:          "test-credential-original",
		StorageConfigurationId: "test-storage",
	}
	wantCreated := Workspace{
		WorkspaceId:            123,
		WorkspaceName:          createRequest.WorkspaceName,
		CredentialsId:          createRequest.CredentialsId,
		StorageConfigurationId: createRequest.StorageConfigurationId,
		WorkspaceStatus:        WorkspaceStatusRunning,
	}
	wantUpdated := wantCreated
	wantUpdated.CredentialsId = "test-credential-updated"

	accountPath := "/api/2.0/accounts/" + accountID
	workspacePath := fmt.Sprintf("%s/workspaces/%d", accountPath, wantCreated.WorkspaceId)

	httpFixtures := qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        accountPath + "/workspaces",
			ExpectedRequest: createRequest,
			Response: Workspace{
				WorkspaceId: wantCreated.WorkspaceId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: workspacePath + "?",
			Response: wantCreated,
		},
		{
			Method:   http.MethodPatch,
			Resource: workspacePath,
			ExpectedRequest: Workspace{
				CredentialsId: wantUpdated.CredentialsId,
			},
			Response: Workspace{
				WorkspaceId: wantUpdated.WorkspaceId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: workspacePath + "?",
			Response: wantUpdated,
		},
		{
			Method:   http.MethodDelete,
			Resource: workspacePath + "?",
			Response: wantUpdated,
		},
	}

	httpFixtures.Apply(t, func(ctx context.Context, cfg *config.Config) {
		cfg.AccountID = accountID
		// Pin PAT auth so credentials in the environment can't take over.
		cfg.AuthType = "pat"
		apiClient, err := client.New(cfg)
		if err != nil {
			t.Fatalf("create fixture client: %v", err)
		}
		workspaces := NewWorkspaces(apiClient)
		workspaceComparison := cmpopts.IgnoreFields(Workspace{}, "ForceSendFields")

		waiter, err := workspaces.Create(ctx, createRequest)
		if err != nil {
			t.Fatalf("create workspace: %v", err)
		}

		gotCreated, err := waiter.Get()
		if err != nil {
			t.Fatalf("wait for created workspace: %v", err)
		}
		if diff := cmp.Diff(&wantCreated, gotCreated, workspaceComparison); diff != "" {
			t.Errorf("created workspace mismatch (-want +got):\n%s", diff)
		}

		updateWaiter, err := workspaces.Update(ctx, UpdateWorkspaceRequest{
			WorkspaceId: wantUpdated.WorkspaceId,
			CustomerFacingWorkspace: Workspace{
				CredentialsId: wantUpdated.CredentialsId,
			},
		})
		if err != nil {
			t.Fatalf("update workspace: %v", err)
		}

		gotUpdated, err := updateWaiter.Get()
		if err != nil {
			t.Fatalf("wait for updated workspace: %v", err)
		}
		if diff := cmp.Diff(&wantUpdated, gotUpdated, workspaceComparison); diff != "" {
			t.Errorf("updated workspace mismatch (-want +got):\n%s", diff)
		}

		gotDeleted, err := workspaces.DeleteByWorkspaceId(ctx, wantUpdated.WorkspaceId)
		if err != nil {
			t.Fatalf("delete workspace: %v", err)
		}
		if diff := cmp.Diff(&wantUpdated, gotDeleted, workspaceComparison); diff != "" {
			t.Errorf("deleted workspace mismatch (-want +got):\n%s", diff)
		}
	})
}
