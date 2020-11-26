package mws

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewWorkspacesAPI creates MWSWorkspacesAPI instance from provider meta
func NewWorkspacesAPI(m interface{}) WorkspacesAPI {
	return WorkspacesAPI{m.(*common.DatabricksClient), context.TODO()}
}

// WorkspacesAPI exposes the mws workspaces API
type WorkspacesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the workspace creation process
func (a WorkspacesAPI) Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID,
	storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) (Workspace, error) {
	var mwsWorkspace Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	mwsWorkspacesRequest := Workspace{
		WorkspaceName:          workspaceName,
		DeploymentName:         deploymentName,
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIPEnabled:    isNoPublicIPEnabled,
	}
	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}
	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}
	err := a.client.Post(a.context, workspacesAPIPath, mwsWorkspacesRequest, &mwsWorkspace)
	return mwsWorkspace, err
}

// WaitForWorkspaceRunning will hold the main thread till the workspace is in a running state
func (a WorkspacesAPI) WaitForWorkspaceRunning(mwsAcctID string, workspaceID int64, timeout time.Duration) error {
	return resource.RetryContext(a.context, timeout, func() *resource.RetryError {
		workspace, err := a.Read(mwsAcctID, workspaceID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		switch workspace.WorkspaceStatus {
		case WorkspaceStatusRunning:
			log.Printf("[INFO] Workspace is now running")
			return nil
		case WorkspaceStatusCanceled, WorkspaceStatusFailed:
			log.Printf("[ERROR] Cannot start workspace: %s", workspace.WorkspaceStatusMessage)
			return resource.NonRetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
		default:
			log.Printf("[INFO] Workspace %s is %s: %s", workspace.DeploymentName,
				workspace.WorkspaceStatus, workspace.WorkspaceStatusMessage)
			return resource.RetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
		}
	})
}

// Patch will relaunch the mws workspace deployment TODO: may need to include customer managed key
func (a WorkspacesAPI) Patch(mwsAcctID string, workspaceID int64, awsRegion, credentialsID,
	storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	mwsWorkspacesRequest := Workspace{
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIPEnabled:    isNoPublicIPEnabled,
	}
	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}
	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}
	return a.client.Patch(a.context, workspacesAPIPath, mwsWorkspacesRequest)
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a WorkspacesAPI) Read(mwsAcctID string, workspaceID int64) (Workspace, error) {
	var mwsWorkspace Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	err := a.client.Get(a.context, workspacesAPIPath, nil, &mwsWorkspace)
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a WorkspacesAPI) Delete(mwsAcctID string, workspaceID int64) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	return a.client.Delete(a.context, workspacesAPIPath, nil)
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
	return &schema.Resource{
		Create: resourceMWSWorkspacesCreate,
		Read:   resourceMWSWorkspacesRead,
		Update: resourceMWSWorkspacesUpdate,
		Delete: resourceMWSWorkspacesDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
				ForceNew:  true,
			},
			"workspace_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "" && new != "" {
						return false
					}
					// Most of E2 accounts require a prefix and API returns it.
					// This is certainly a hack to get things working for Terraform operating model.
					// https://github.com/databrickslabs/terraform-provider-databricks/issues/382
					return !strings.HasSuffix(new, old)
				},
			},
			"aws_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credentials_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"storage_configuration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"verify_workspace_runnning": {
				Deprecated: "`verify_workspace_runnning` is deprecated and are going to be removed in 0.3. " +
					"All workspaces would be verified to get into runnable state or cleaned up upon failure.",
				Type:     schema.TypeBool,
				Required: true,
			},
			"customer_managed_key_id": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_no_public_ip_enabled": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"workspace_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workspace_status_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"workspace_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"workspace_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_error_messages": {
				Deprecated: "`network_error_messages` are deprecated and are going to be removed in 0.3. " +
					"Any VPC failures would simply return an error directly.",
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func waitForWorkspaceURLResolution(ctx context.Context, workspace Workspace, timeoutDurationMinutes time.Duration) error {
	if workspace.DeploymentName == "900150983cd24fb0" {
		// nobody would probably name workspace as 900150983cd24fb0,
		// so we'll use it as unit testing shim
		return nil
	}
	hostAndPort := fmt.Sprintf("%s.cloud.databricks.com:443", workspace.DeploymentName)
	url := fmt.Sprintf("https://%s.cloud.databricks.com", workspace.DeploymentName)
	return resource.RetryContext(ctx, timeoutDurationMinutes, func() *resource.RetryError {
		conn, err := net.DialTimeout("tcp", hostAndPort, 1*time.Minute)
		if err != nil {
			log.Printf("Cannot yet reach %s", url)
			return resource.RetryableError(err)
		}
		log.Printf("Workspace %s is ready to use", url)
		defer conn.Close()
		return nil
	})
}

func resourceMWSWorkspacesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	mwsAcctID := d.Get("account_id").(string)
	workspaceName := d.Get("workspace_name").(string)
	deploymentName := d.Get("deployment_name").(string)
	awsRegion := d.Get("aws_region").(string)
	credentialsID := d.Get("credentials_id").(string)
	storageConfigurationID := d.Get("storage_configuration_id").(string)
	networkID := d.Get("network_id").(string)
	customerManagedKeyID := d.Get("customer_managed_key_id").(string)
	isNoPublicIPEnabled := d.Get("is_no_public_ip_enabled").(bool)
	var workspace Workspace
	var err error

	workspacesAPI := NewWorkspacesAPI(client)

	workspace, err = workspacesAPI.Create(mwsAcctID, workspaceName, deploymentName,
		awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	// Sometimes workspaces api is buggy
	if err != nil {
		time.Sleep(15 * time.Second)
		workspace, err = workspacesAPI.Create(mwsAcctID, workspaceName, deploymentName,
			awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
		if err != nil {
			return err
		}
	}
	workspaceResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: strconv.Itoa(int(workspace.WorkspaceID)),
	}
	d.SetId(packMWSAccountID(workspaceResourceID))
	err = workspacesAPI.WaitForWorkspaceRunning(mwsAcctID, workspace.WorkspaceID, 10*time.Minute)
	if err != nil {
		if !reflect.ValueOf(networkID).IsZero() {
			network, networkReadErr := NewNetworksAPI(client).Read(mwsAcctID, networkID)
			if networkReadErr != nil {
				return fmt.Errorf("Workspace failed to create: %v, network read failure error: %v", err, networkReadErr)
			}
			return fmt.Errorf("Workspace failed to create: %v, network error message: %v", err, getNetworkErrors(network.ErrorMessages))
		}
		return err
	}
	// wait maximum 5 minute for DNS caches to refresh, as
	// sometimes we cannot make API calls to new workspaces
	// immediately after it's created
	err = waitForWorkspaceURLResolution(workspacesAPI.context, workspace, 5*time.Minute)
	if err != nil {
		return err
	}
	return resourceMWSWorkspacesRead(d, m)
}

func resourceMWSWorkspacesRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	idInt64, err := strconv.ParseInt(packagedMwsID.ResourceID, 10, 64)
	if err != nil {
		return err
	}

	workspace, err := NewWorkspacesAPI(client).Read(packagedMwsID.MwsAcctID, idInt64)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	if workspace.WorkspaceStatus != WorkspaceStatusRunning {
		err = NewWorkspacesAPI(client).WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10*time.Minute)
		if err != nil {
			log.Println("WORKSPACE IS NOT RUNNING")
			err2 := d.Set("verify_workspace_runnning", false)
			if err2 != nil {
				return err2
			}
		} else {
			err2 := d.Set("verify_workspace_runnning", true)
			if err2 != nil {
				return err2
			}
		}
	}
	// TODO: account id property
	err = d.Set("deployment_name", workspace.DeploymentName)
	if err != nil {
		return err
	}
	err = d.Set("workspace_name", workspace.WorkspaceName)
	if err != nil {
		return err
	}
	err = d.Set("aws_region", workspace.AwsRegion)
	if err != nil {
		return err
	}
	err = d.Set("credentials_id", workspace.CredentialsID)
	if err != nil {
		return err
	}
	err = d.Set("storage_configuration_id", workspace.StorageConfigurationID)
	if err != nil {
		return err
	}
	err = d.Set("network_id", workspace.NetworkID)
	if err != nil {
		return err
	}
	err = d.Set("customer_managed_key_id", workspace.CustomerManagedKeyID)
	if err != nil {
		return err
	}
	err = d.Set("account_id", workspace.AccountID)
	if err != nil {
		return err
	}
	err = d.Set("workspace_status", workspace.WorkspaceStatus)
	if err != nil {
		return err
	}

	if workspace.WorkspaceStatus != WorkspaceStatusRunning {
		network, err := NewNetworksAPI(client).Read(workspace.AccountID, workspace.NetworkID)
		if err == nil && !reflect.ValueOf(network.ErrorMessages).IsZero() {
			err = d.Set("network_error_messages", convertErrorMessagesToListOfMaps(network.ErrorMessages))
			if err != nil {
				return err
			}
		}
	} else {
		err = d.Set("network_error_messages", []map[string]string{{"error_type": "", "error_message": ""}})
		if err != nil {
			return err
		}
	}

	err = d.Set("workspace_status_message", workspace.WorkspaceStatusMessage)
	if err != nil {
		return err
	}
	err = d.Set("creation_time", workspace.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("workspace_id", workspace.WorkspaceID)
	if err != nil {
		return err
	}
	err = d.Set("workspace_url", fmt.Sprintf("https://%s.cloud.databricks.com", workspace.DeploymentName))
	if err != nil {
		return err
	}
	return nil
}

func resourceMWSWorkspacesUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	idInt64, err := strconv.ParseInt(packagedMwsID.ResourceID, 10, 64)
	if err != nil {
		return err
	}
	awsRegion := d.Get("aws_region").(string)
	credentialsID := d.Get("credentials_id").(string)
	storageConfigurationID := d.Get("storage_configuration_id").(string)
	networkID := d.Get("network_id").(string)
	customerManagedKeyID := d.Get("customer_managed_key_id").(string)
	isNoPublicIPEnabled := d.Get("is_no_public_ip_enabled").(bool)

	err = NewWorkspacesAPI(client).Patch(packagedMwsID.MwsAcctID, idInt64, awsRegion, credentialsID,
		storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	if err != nil {
		return err
	}
	err = NewWorkspacesAPI(client).WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10*time.Minute)
	if err != nil {
		return err
	}
	return resourceMWSWorkspacesRead(d, m)
}

func resourceMWSWorkspacesDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	idInt64, err := strconv.ParseInt(packagedMwsID.ResourceID, 10, 64)
	if err != nil {
		return err
	}
	workspacesAPI := NewWorkspacesAPI(client)
	if err = workspacesAPI.Delete(packagedMwsID.MwsAcctID, idInt64); err != nil {
		return err
	}
	return resource.RetryContext(workspacesAPI.context, 15*time.Minute, func() *resource.RetryError {
		workspace, err := workspacesAPI.Read(packagedMwsID.MwsAcctID, idInt64)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] Workspace %s is removed.", packagedMwsID.ResourceID)
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

func getNetworkErrors(networkRespList []NetworkHealth) string {
	var strBuffer bytes.Buffer
	for _, networkHealth := range networkRespList {
		strBuffer.WriteString(fmt.Sprintf("error: %s;error_msg: %s;", networkHealth.ErrorType, networkHealth.ErrorMessage))
	}
	return strBuffer.String()
}
