package databricks

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"log"
	"strconv"
	"strings"
	"time"
)

func resourceMWSWorkspaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceMWSWorkspacesCreate,
		Read:   resourceMWSWorkspacesRead,
		Update: resourceMWSWorkspacePatch,
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

func resourceMWSWorkspacesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	mwsAcctID := d.Get("account_id").(string)
	workspaceName := d.Get("workspace_name").(string)
	deploymentName := d.Get("deployment_name").(string)
	awsRegion := d.Get("aws_region").(string)
	credentialsID := d.Get("credentials_id").(string)
	storageConfigurationID := d.Get("storage_configuration_id").(string)
	networkID := d.Get("network_id").(string)
	customerManagedKeyID := d.Get("customer_managed_key_id").(string)
	isNoPublicIPEnabled := d.Get("is_no_public_ip_enabled").(bool)
	var workspace model.MWSWorkspace
	var err error
	workspace, err = client.MWSWorkspaces().Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	// Sometimes workspaces api is buggy
	if err != nil {
		time.Sleep(15 * time.Second)
		workspace, err = client.MWSWorkspaces().Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
		if err != nil {
			return err
		}
	}
	workspaceResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: strconv.Itoa(int(workspace.WorkspaceID)),
	}
	d.SetId(packMWSAccountID(workspaceResourceID))
	err = client.MWSWorkspaces().WaitForWorkspaceRunning(mwsAcctID, workspace.WorkspaceID, 10, 180)
	if err != nil {
		if !reflect.ValueOf(networkID).IsZero() {
			network, networkReadErr := client.MWSNetworks().Read(mwsAcctID, networkID)
			if networkReadErr != nil {
				return fmt.Errorf("Workspace failed to create: %v, network read failure error: %v", err, networkReadErr)
			}
			return fmt.Errorf("Workspace failed to create: %v, network error message: %v", err, getNetworkErrors(network.ErrorMessages))
		}
		return err
	}
	return resourceMWSWorkspacesRead(d, m)
}

func resourceMWSWorkspacesRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	idInt64, err := strconv.ParseInt(packagedMwsID.ResourceID, 10, 64)
	if err != nil {
		return err
	}

	workspace, err := client.MWSWorkspaces().Read(packagedMwsID.MwsAcctID, idInt64)
	if err != nil {
		if isMWSWorkspaceMissing(err.Error(), id) {
			log.Printf("Missing e2 workspace with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	err = client.MWSWorkspaces().WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10, 180)
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

	err = d.Set("deployment_name", workspace.DeploymentName)
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

	if workspace.WorkspaceStatus != model.WorkspaceStatusRunning {
		network, err := client.MWSNetworks().Read(workspace.AccountID, workspace.NetworkID)
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

func resourceMWSWorkspacePatch(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	packagedMwsID, err := unpackMWSAccountID(id)
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

	err = client.MWSWorkspaces().Patch(packagedMwsID.MwsAcctID, idInt64, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID, isNoPublicIPEnabled)
	if err != nil {
		return err
	}
	err = client.MWSWorkspaces().WaitForWorkspaceRunning(packagedMwsID.MwsAcctID, idInt64, 10, 180)
	if err != nil {
		return err
	}
	return resourceMWSWorkspacesRead(d, m)
}

func resourceMWSWorkspacesDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	idInt64, err := strconv.ParseInt(packagedMwsID.ResourceID, 10, 64)
	if err != nil {
		return err
	}
	err = client.MWSWorkspaces().Delete(packagedMwsID.MwsAcctID, idInt64)
	return err
}

func getNetworkErrors(networkRespList []model.NetworkHealth) string {
	var strBuffer bytes.Buffer
	for _, networkHealth := range networkRespList {
		strBuffer.WriteString(fmt.Sprintf("error: %s;error_msg: %s;", networkHealth.ErrorType, networkHealth.ErrorMessage))
	}
	return strBuffer.String()
}

func isMWSWorkspaceMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST") &&
		strings.Contains(errorMsg, fmt.Sprintf("workspace %s does not exist", resourceID))
}
