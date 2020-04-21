package databricks

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"os"
)

// Provider returns the entire terraform provider object
func Provider(version string) terraform.ResourceProvider {
	provider := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"databricks_notebook":        dataSourceNotebook(),
			"databricks_notebook_paths":  dataSourceNotebookPaths(),
			"databricks_dbfs_file":       dataSourceDBFSFile(),
			"databricks_dbfs_file_paths": dataSourceDBFSFilePaths(),
			"databricks_zones":           dataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"databricks_token":                 resourceToken(),
			"databricks_secret_scope":          resourceSecretScope(),
			"databricks_secret":                resourceSecret(),
			"databricks_secret_acl":            resourceSecretACL(),
			"databricks_instance_pool":         resourceInstancePool(),
			"databricks_scim_user":             resourceScimUser(),
			"databricks_scim_group":            resourceScimGroup(),
			"databricks_notebook":              resourceNotebook(),
			"databricks_cluster":               resourceCluster(),
			"databricks_job":                   resourceJob(),
			"databricks_dbfs_file":             resourceDBFSFile(),
			"databricks_dbfs_file_sync":        resourceDBFSFileSync(),
			"databricks_instance_profile":      resourceInstanceProfile(),
			"databricks_aws_s3_mount":          resourceAWSS3Mount(),
			"databricks_azure_blob_mount":      resourceAzureBlobMount(),
			"databricks_azure_adls_gen1_mount": resourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount": resourceAzureAdlsGen2Mount(),
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
		return providerConfigure(d, version)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, providerVersion string) (interface{}, error) {
	var config service.DBApiClientConfig
	if azureAuth, ok := d.GetOk("azure_auth"); !ok {
		if host, ok := d.GetOk("host"); ok {
			config.Host = host.(string)
		} else {
			config.Host = os.Getenv("HOST")
		}
		if token, ok := d.GetOk("token"); ok {
			config.Token = token.(string)
		} else {
			config.Token = os.Getenv("TOKEN")
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
			tokenPayload.SubscriptionID = subscriptionID
		} else {
			tokenPayload.SubscriptionID = os.Getenv("ARM_SUBSCRIPTION_ID")
		}
		if clientSecret, ok := azureAuthMap["client_secret"].(string); ok {
			tokenPayload.ClientSecret = clientSecret
		} else {
			tokenPayload.SubscriptionID = os.Getenv("ARM_CLIENT_SECRET")
		}
		if clientID, ok := azureAuthMap["client_id"].(string); ok {
			tokenPayload.ClientID = clientID
		} else {
			tokenPayload.SubscriptionID = os.Getenv("ARM_CLIENT_ID")
		}
		if tenantID, ok := azureAuthMap["tenant_id"].(string); ok {
			tokenPayload.TenantID = tenantID
		} else {
			tokenPayload.SubscriptionID = os.Getenv("ARM_TENANT_ID")
		}

		azureAuthSetup := AzureAuth{
			TokenPayload:           &tokenPayload,
			ManagementToken:        "",
			AdbWorkspaceResourceID: "",
			AdbAccessToken:         "",
			AdbPlatformToken:       "",
		}
		log.Println("Running Azure Auth")
		return azureAuthSetup.initWorkspaceAndGetClient(&config)
	}

	//TODO: Bake the version of the provider using -ldflags to tell the golang linker to send
	//version information from go-releaser
	config.UserAgent = fmt.Sprintf("databricks-tf-provider-%s", providerVersion)
	var dbClient service.DBApiClient
	dbClient.SetConfig(&config)
	return dbClient, nil
}
