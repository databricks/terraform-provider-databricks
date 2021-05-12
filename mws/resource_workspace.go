package mws

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DefaultProvisionTimeout is the amount of minutes terraform will wait
// for workspace to be provisioned and DNS entry to be available. Increasing
// this may help with local DNS cache issues.
const DefaultProvisionTimeout = 20 * time.Minute

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
func (a WorkspacesAPI) Create(ws *Workspace, timeout time.Duration) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", ws.AccountID)
	err := a.client.Post(a.context, workspacesAPIPath, ws, &ws)
	if err != nil {
		return err
	}
	if err = a.WaitForRunning(*ws, timeout); err != nil {
		log.Printf("[ERROR] Deleting failed workspace: %s", err)
		if derr := a.Delete(ws.AccountID, fmt.Sprintf("%d", ws.WorkspaceID)); derr != nil {
			return fmt.Errorf("%s - %s", err, derr)
		}
		return err
	}
	return nil
}

func dial(hostAndPort, url string, timeout time.Duration) *resource.RetryError {
	conn, err := net.DialTimeout("tcp", hostAndPort, timeout)
	if err != nil {
		return resource.RetryableError(err)
	}
	log.Printf("[INFO] Workspace %s is ready to use", url)
	defer conn.Close()
	return nil
}

// WaitForRunning will wait until workspace is running, otherwise will try to explain why it failed
func (a WorkspacesAPI) WaitForRunning(ws Workspace, timeout time.Duration) error {
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
			return dial(hostAndPort, url, 10*time.Second)
		case WorkspaceStatusCanceled, WorkspaceStatusFailed:
			log.Printf("[ERROR] Cannot start workspace: %s", workspace.WorkspaceStatusMessage)
			if workspace.NetworkID == "" {
				return resource.NonRetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
			}
			network, nerr := NewNetworksAPI(a.context, a.client).Read(ws.AccountID, ws.NetworkID)
			if nerr != nil {
				return resource.NonRetryableError(fmt.Errorf(
					"failed to start workspace. Cannot read network: %s", nerr))
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
func (a WorkspacesAPI) Patch(ws Workspace, timeout time.Duration) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", ws.AccountID, ws.WorkspaceID)
	err := a.client.Patch(a.context, workspacesAPIPath, Workspace{
		AwsRegion:                           ws.AwsRegion,
		CredentialsID:                       ws.CredentialsID,
		StorageConfigurationID:              ws.StorageConfigurationID,
		IsNoPublicIPEnabled:                 ws.IsNoPublicIPEnabled,
		NetworkID:                           ws.NetworkID,
		ManagedServicesCustomerManagedKeyID: ws.ManagedServicesCustomerManagedKeyID,
		StoragexCustomerManagedKeyID:        ws.StoragexCustomerManagedKeyID,
	})
	if err != nil {
		return err
	}
	return a.WaitForRunning(ws, timeout)
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
	s := common.StructToSchema(Workspace{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
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
		s["customer_managed_key_id"].Deprecated = "Use managed_services_customer_managed_key_id instead"
		s["customer_managed_key_id"].ConflictsWith = []string{"managed_services_customer_managed_key_id", "storage_customer_managed_key_id"}
		s["managed_services_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
		s["storage_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
		return s
	})
	p := common.NewPairSeparatedID("account_id", "workspace_id", "/").Schema(
		func(_ map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema:        s,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if err := common.DataToStructPointer(d, s, &workspace); err != nil {
				return err
			}
			if len(workspace.CustomerManagedKeyID) > 0 && len(workspace.ManagedServicesCustomerManagedKeyID) == 0 {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				workspace.ManagedServicesCustomerManagedKeyID = workspace.CustomerManagedKeyID
				workspace.CustomerManagedKeyID = ""
			}
			if err := workspacesAPI.Create(&workspace, d.Timeout(schema.TimeoutCreate)); err != nil {
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
			if err = common.StructToData(workspace, s, d); err != nil {
				return err
			}
			return workspacesAPI.WaitForRunning(workspace, d.Timeout(schema.TimeoutRead))
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if err := common.DataToStructPointer(d, s, &workspace); err != nil {
				return err
			}
			if len(workspace.CustomerManagedKeyID) > 0 && len(workspace.ManagedServicesCustomerManagedKeyID) == 0 {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				workspace.ManagedServicesCustomerManagedKeyID = workspace.CustomerManagedKeyID
				workspace.CustomerManagedKeyID = ""
			}
			return workspacesAPI.Patch(workspace, d.Timeout(schema.TimeoutUpdate))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewWorkspacesAPI(ctx, c).Delete(accountID, workspaceID)
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Read:   schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}.ToResource()
}
