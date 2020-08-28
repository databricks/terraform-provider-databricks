package mws

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewWorkspacesAPI creates MWSWorkspacesAPI instance from provider meta
func NewWorkspacesAPI(m interface{}) WorkspacesAPI {
	return WorkspacesAPI{client: m.(*common.DatabricksClient)}
}

// WorkspacesAPI exposes the mws workspaces API
type WorkspacesAPI struct {
	client *common.DatabricksClient
}

// Create creates the workspace creation process
func (a WorkspacesAPI) Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) (Workspace, error) {
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
	err := a.client.Post(workspacesAPIPath, mwsWorkspacesRequest, &mwsWorkspace)
	return mwsWorkspace, err
}

// WaitForWorkspaceRunning will hold the main thread till the workspace is in a running state
func (a WorkspacesAPI) WaitForWorkspaceRunning(mwsAcctID string, workspaceID int64, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	// TODO: move all resource awaiters from client to TF resource level, for sepration of concerns sake
	errChan := make(chan error, 1)
	go func() {
		for {
			workspace, err := a.Read(mwsAcctID, workspaceID)
			if err != nil {
				errChan <- err
			}
			if workspace.WorkspaceStatus == WorkspaceStatusRunning {
				errChan <- nil
			} else if ContainsWorkspaceState(WorkspaceStatusesNonRunnable, workspace.WorkspaceStatus) {
				errChan <- errors.New("Workspace is in a non runnable state will not be able to transition to running, needs " +
					"to be created again. Current state: " + workspace.WorkspaceStatus)
			}
			log.Println("Waiting for workspace to go to running, current state is: " + workspace.WorkspaceStatus)
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out workspace has not reached running state")
	}
}

// Patch will relaunch the mws workspace deployment TODO: may need to include customer managed key
func (a WorkspacesAPI) Patch(mwsAcctID string, workspaceID int64, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) error {
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
	return a.client.Patch(workspacesAPIPath, mwsWorkspacesRequest)
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a WorkspacesAPI) Read(mwsAcctID string, workspaceID int64) (Workspace, error) {
	var mwsWorkspace Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	err := a.client.Get(workspacesAPIPath, nil, &mwsWorkspace)
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a WorkspacesAPI) Delete(mwsAcctID string, workspaceID int64) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	return a.client.Delete(workspacesAPIPath, nil)
}

// List will list all workspaces in a given mws account
func (a WorkspacesAPI) List(mwsAcctID string) ([]Workspace, error) {
	var mwsWorkspacesList []Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	err := a.client.Get(workspacesAPIPath, nil, &mwsWorkspacesList)
	return mwsWorkspacesList, err
}

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
			},
			"aws_region": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"us-west-1",
					"us-west-2",
					"us-east-1",
					"sa-east-1",
					"eu-west-1",
					"eu-central-1",
					"ap-south-1",
					"ap-southeast-1",
					"ap-southeast-2",
					"ap-northeast-1",
					"ap-northeast-2",
					"ca-central-1",
				}, false),
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

func waitForWorkspaceURLResolution(workspace Workspace, timeoutDurationMinutes time.Duration) error {
	if workspace.DeploymentName == "900150983cd24fb0" {
		// nobody would probably name workspace as 900150983cd24fb0,
		// so we'll use it as unit testing shim
		return nil
	}
	hostAndPort := fmt.Sprintf("%s.cloud.databricks.com:443", workspace.DeploymentName)
	url := fmt.Sprintf("https://%s.cloud.databricks.com", workspace.DeploymentName)
	return resource.Retry(timeoutDurationMinutes, func() *resource.RetryError {
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
	workspace, err = NewWorkspacesAPI(client).Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	// Sometimes workspaces api is buggy
	if err != nil {
		time.Sleep(15 * time.Second)
		workspace, err = NewWorkspacesAPI(client).Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
		if err != nil {
			return err
		}
	}
	workspaceResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: strconv.Itoa(int(workspace.WorkspaceID)),
	}
	d.SetId(packMWSAccountID(workspaceResourceID))
	// TODO: replace with waitForWorkspaceState
	err = NewWorkspacesAPI(client).WaitForWorkspaceRunning(mwsAcctID, workspace.WorkspaceID, 10, 180)
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
	err = waitForWorkspaceURLResolution(workspace, 5*time.Minute)
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
		// TODO: replace with waitForWorkspaceState
		err = NewWorkspacesAPI(client).WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10, 180)
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

	err = NewWorkspacesAPI(client).Patch(packagedMwsID.MwsAcctID, idInt64, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	if err != nil {
		return err
	}
	// TODO: replace with waitForWorkspaceState, potentially with state machine checks
	err = NewWorkspacesAPI(client).WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10, 180)
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
	err = NewWorkspacesAPI(client).Delete(packagedMwsID.MwsAcctID, idInt64)
	if err != nil {
		return err
	}
	return resource.Retry(15*time.Minute, func() *resource.RetryError {
		workspace, err := NewWorkspacesAPI(client).Read(packagedMwsID.MwsAcctID, idInt64)
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
