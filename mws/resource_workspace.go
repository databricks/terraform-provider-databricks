package mws

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewWorkspacesAPI creates MWSWorkspacesAPI instance from provider meta
func NewWorkspacesAPI(ctx context.Context, m interface{}) WorkspacesAPI {
	return WorkspacesAPI{m.(*common.DatabricksClient), ctx}
}

// WorkspacesAPI exposes the mws workspaces API
type WorkspacesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the workspace creation process
func (a WorkspacesAPI) Create(ws *Workspace) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", ws.AccountID)
	err := a.client.Post(a.context, workspacesAPIPath, ws, &ws)
	if err != nil {
		return err
	}
	if err = a.waitForRunning(*ws, 15*time.Minute); err != nil {
		log.Printf("[ERROR] Deleting failed workspace: %s", err)
		if derr := a.Delete(ws.AccountID, fmt.Sprintf("%d", ws.WorkspaceID)); derr != nil {
			return fmt.Errorf("%s - %s", err, derr)
		}
		return err
	}
	return nil
}

func dial(hostAndPort, url string, timeout time.Duration) *resource.RetryError {
	conn, err := net.DialTimeout("tcp", hostAndPort, 1*time.Minute)
	if err != nil {
		return resource.RetryableError(err)
	}
	log.Printf("[INFO] Workspace %s is ready to use", url)
	defer conn.Close()
	return nil
}

// waitForRunning will wait until workspace is running, otherwise will try to explain why it failed
func (a WorkspacesAPI) waitForRunning(ws Workspace, timeout time.Duration) error {
	return resource.RetryContext(a.context, timeout, func() *resource.RetryError {
		workspace, err := a.Read(ws.AccountID, fmt.Sprintf("%d", ws.WorkspaceID))
		if err != nil {
			return resource.NonRetryableError(err)
		}
		switch workspace.WorkspaceStatus {
		case WorkspaceStatusRunning:
			// wait for DNS caches to refresh, as sometimes we cannot make
			// API calls to new workspaces immediately after it's created
			hostAndPort := fmt.Sprintf("%s.cloud.databricks.com:443", workspace.DeploymentName)
			url := fmt.Sprintf("https://%s.cloud.databricks.com", workspace.DeploymentName)
			log.Printf("[INFO] Workspace is now running")
			if strings.Contains(workspace.DeploymentName, "900150983cd24fb0") {
				// nobody would probably name workspace as 900150983cd24fb0,
				// so we'll use it as unit testing shim
				return nil
			}
			return dial(hostAndPort, url, 1*time.Minute)
		case WorkspaceStatusCanceled, WorkspaceStatusFailed:
			log.Printf("[ERROR] Cannot start workspace: %s", workspace.WorkspaceStatusMessage)
			if workspace.NetworkID == "" {
				return resource.NonRetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
			}
			network, nerr := NewNetworksAPI(a.context, a.client).Read(ws.AccountID, ws.NetworkID)
			if nerr != nil {
				return resource.NonRetryableError(fmt.Errorf(
					"Failed to start workspace. Cannot read network: %s", nerr))
			}
			var strBuffer bytes.Buffer
			for _, networkHealth := range network.ErrorMessages {
				strBuffer.WriteString(fmt.Sprintf("error: %s;error_msg: %s;",
					networkHealth.ErrorType, networkHealth.ErrorMessage))
			}
			return resource.NonRetryableError(fmt.Errorf(
				"Workspace failed to create: %v, network error message: %v",
				workspace.WorkspaceStatusMessage, strBuffer.String()))
		default:
			log.Printf("[INFO] Workspace %s is %s: %s", workspace.DeploymentName,
				workspace.WorkspaceStatus, workspace.WorkspaceStatusMessage)
			return resource.RetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
		}
	})
}

// Patch will relaunch the workspace deployment
func (a WorkspacesAPI) Patch(ws Workspace) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", ws.AccountID, ws.WorkspaceID)
	err := a.client.Patch(a.context, workspacesAPIPath, Workspace{
		AwsRegion:              ws.AwsRegion,
		CredentialsID:          ws.CredentialsID,
		StorageConfigurationID: ws.StorageConfigurationID,
		IsNoPublicIPEnabled:    ws.IsNoPublicIPEnabled,
		NetworkID:              ws.NetworkID,
		CustomerManagedKeyID:   ws.CustomerManagedKeyID,
	})
	if err != nil {
		return err
	}
	return a.waitForRunning(ws, 10*time.Minute)
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a WorkspacesAPI) Read(mwsAcctID, workspaceID string) (Workspace, error) {
	var mwsWorkspace Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%s", mwsAcctID, workspaceID)
	err := a.client.Get(a.context, workspacesAPIPath, nil, &mwsWorkspace)
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a WorkspacesAPI) Delete(mwsAcctID, workspaceID string) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%s", mwsAcctID, workspaceID)
	err := a.client.Delete(a.context, workspacesAPIPath, nil)
	if err != nil {
		return err
	}
	return resource.RetryContext(a.context, 15*time.Minute, func() *resource.RetryError {
		workspace, err := a.Read(mwsAcctID, workspaceID)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] Workspace %s/%s is removed.", mwsAcctID, workspaceID)
			return nil
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		msg := fmt.Errorf("Workspace %s is not removed yet. Workspace status: %s %s",
			workspace.WorkspaceName, workspace.WorkspaceStatus, workspace.WorkspaceStatusMessage)
		log.Printf("[INFO] %s", msg)
		return resource.RetryableError(msg)
	})
}

// List will list all workspaces in a given mws account
func (a WorkspacesAPI) List(mwsAcctID string) ([]Workspace, error) {
	var mwsWorkspacesList []Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	err := a.client.Get(a.context, workspacesAPIPath, nil, &mwsWorkspacesList)
	return mwsWorkspacesList, err
}

// ResourceWorkspace manages E2 workspaces
func ResourceWorkspace() *schema.Resource {
	s := internal.StructToSchema(Workspace{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["account_id"].Sensitive = true
		s["account_id"].ForceNew = true
		s["workspace_name"].ForceNew = true
		s["deployment_name"].ForceNew = true
		s["deployment_name"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
			if old == "" && new != "" {
				return false
			}
			// Most of E2 accounts require a prefix and API returns it.
			// This is certainly a hack to get things working for Terraform operating model.
			// https://github.com/databrickslabs/terraform-provider-databricks/issues/382
			return !strings.HasSuffix(new, old)
		}
		s["is_no_public_ip_enabled"].Default = false
		return s
	})
	p := util.NewPairSeparatedID("account_id", "workspace_id", "/").Schema(
		func(_ map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return util.CommonResource{
		Schema:        s,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if err := internal.DataToStructPointer(d, s, &workspace); err != nil {
				return err
			}
			if err := workspacesAPI.Create(&workspace); err != nil {
				return err
			}
			d.Set("workspace_id", workspace.WorkspaceID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			workspace, err := workspacesAPI.Read(accountID, workspaceID)
			if err != nil {
				return err
			}
			workspace.WorkspaceURL = fmt.Sprintf("https://%s.cloud.databricks.com", workspace.DeploymentName)
			if err = internal.StructToData(workspace, s, d); err != nil {
				return err
			}
			return workspacesAPI.waitForRunning(workspace, 10*time.Minute)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if err := internal.DataToStructPointer(d, s, &workspace); err != nil {
				return err
			}
			return workspacesAPI.Patch(workspace)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewWorkspacesAPI(ctx, c).Delete(accountID, workspaceID)
		},
	}.ToResource()
}
