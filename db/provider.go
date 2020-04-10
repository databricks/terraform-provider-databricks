package db

import (
	"github.com/databrickslabs/databricks-terraform/client"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"os"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"db_notebook":        dataSourceNotebook(),
			"db_notebook_paths":  dataSourceNotebookPaths(),
			"db_dbfs_file":       dataSourceDBFSFile(),
			"db_dbfs_file_paths": dataSourceDBFSFilePaths(),
			"db_zones":           dataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"db_token":                 resourceToken(),
			"db_secret_scope":          resourceSecretScope(),
			"db_secret":                resourceSecret(),
			"db_secret_acl":            resourceSecretAcl(),
			"db_instance_pool":         resourceInstancePool(),
			"db_scim_user":             resourceScimUser(),
			"db_scim_group":            resourceScimGroup(),
			"db_notebook":              resourceNotebook(),
			"db_cluster":               resourceCluster(),
			"db_job":                   resourceJob(),
			"db_dbfs_file":             resourceDBFSFile(),
			"db_dbfs_file_sync":        resourceDBFSFileSync(),
			"db_instance_profile":      resourceInstanceProfile(),
			"db_aws_s3_mount":          resourceAWSS3Mount(),
			"db_azure_blob_mount":      resourceAzureBlobMount(),
			"db_azure_adls_gen1_mount": resourceAzureAdlsGen1Mount(),
			"db_azure_adls_gen2_mount": resourceAzureAdlsGen2Mount(),
		},
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_auth": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_resource_group": {
							Type:     schema.TypeString,
							Required: true,
						},
						"azure_region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"workspace_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_secret": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tenant_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, s string) (interface{}, error) {
	var option client.DBClientOption
	if azureAuth, ok := d.GetOk("azure_auth"); !ok {
		if host, ok := d.GetOk("host"); ok {
			option.Host = host.(string)
		} else {
			option.Host = os.Getenv("HOST")
		}
		if token, ok := d.GetOk("token"); ok {
			option.Token = token.(string)
		} else {
			option.Token = os.Getenv("TOKEN")
		}
	} else {
		log.Println("Creating db client via azure auth!")
		azureAuthMap := azureAuth.(map[string]interface{})
		//azureAuth AzureAuth{}
		tokenPayload := TokenPayload{}
		if managedResourceGroup, ok := azureAuthMap["managed_resource_group"].(string); ok {
			tokenPayload.ManagedResourceGroup = managedResourceGroup
		}
		if azureRegion, ok := azureAuthMap["azure_region"].(string); ok {
			tokenPayload.AzureRegion = azureRegion
		}
		if resourceGroup, ok := azureAuthMap["resource_group"].(string); ok {
			tokenPayload.ResourceGroup = resourceGroup
		}
		if workspaceName, ok := azureAuthMap["workspace_name"].(string); ok {
			tokenPayload.WorkspaceName = workspaceName
		}
		if subscriptionID, ok := azureAuthMap["subscription_id"].(string); ok {
			tokenPayload.SubscriptionId = subscriptionID
		} else {
			tokenPayload.SubscriptionId = os.Getenv("ARM_SUBSCRIPTION_ID")
		}
		if clientSecret, ok := azureAuthMap["client_secret"].(string); ok {
			tokenPayload.ClientSecret = clientSecret
		} else {
			tokenPayload.SubscriptionId = os.Getenv("ARM_CLIENT_SECRET")
		}
		if clientID, ok := azureAuthMap["client_id"].(string); ok {
			tokenPayload.ClientID = clientID
		} else {
			tokenPayload.SubscriptionId = os.Getenv("ARM_CLIENT_ID")
		}
		if tenantID, ok := azureAuthMap["tenant_id"].(string); ok {
			tokenPayload.TenantID = tenantID
		} else {
			tokenPayload.SubscriptionId = os.Getenv("ARM_TENANT_ID")
		}

		azureAuthSetup := AzureAuth{
			TokenPayload:           &tokenPayload,
			ManagementToken:        "",
			AdbWorkspaceResourceID: "",
			AdbAccessToken:         "",
			AdbPlatformToken:       "",
		}
		log.Println("Running Azure Auth")
		return azureAuthSetup.initWorkspaceAndGetClient(option)
	}

	var dbClient service.DBApiClient
	dbClient.Init(option)
	return dbClient, nil
}
