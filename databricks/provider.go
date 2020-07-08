package databricks

import (
	"fmt"
	"strconv"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns the entire terraform provider object
func Provider(version string) terraform.ResourceProvider {
	provider := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"databricks_default_user_roles": dataSourceDefaultUserRoles(),
			"databricks_notebook":           dataSourceNotebook(),
			"databricks_notebook_paths":     dataSourceNotebookPaths(),
			"databricks_dbfs_file":          dataSourceDBFSFile(),
			"databricks_dbfs_file_paths":    dataSourceDBFSFilePaths(),
			"databricks_zones":              dataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"databricks_token":         resourceToken(),
			"databricks_secret_scope":  resourceSecretScope(),
			"databricks_secret":        resourceSecret(),
			"databricks_secret_acl":    resourceSecretACL(),
			"databricks_permissions":   resourcePermissions(),
			"databricks_instance_pool": resourceInstancePool(),
			"databricks_scim_user":     resourceScimUser(),
			"databricks_scim_group":    resourceScimGroup(),
			// Scim Group is split into multiple components for flexibility to pick and choose
			"databricks_group":                  resourceGroup(),
			"databricks_group_instance_profile": resourceGroupInstanceProfile(),
			"databricks_group_member":           resourceGroupMember(),
			"databricks_notebook":               resourceNotebook(),
			"databricks_cluster":                resourceCluster(),
			"databricks_cluster_policy":         resourceClusterPolicy(),
			"databricks_job":                    resourceJob(),
			"databricks_dbfs_file":              resourceDBFSFile(),
			"databricks_instance_profile":       resourceInstanceProfile(),
			"databricks_aws_s3_mount":           resourceAWSS3Mount(),
			"databricks_azure_blob_mount":       resourceAzureBlobMount(),
			"databricks_azure_adls_gen1_mount":  resourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount":  resourceAzureAdlsGen2Mount(),
			//	MWS (multiple workspaces) resources are only limited to AWS as azure already has a built in concept of MWS
			"databricks_mws_credentials":            resourceMWSCredentials(),
			"databricks_mws_storage_configurations": resourceMWSStorageConfigurations(),
			"databricks_mws_networks":               resourceMWSNetworks(),
			"databricks_mws_workspaces":             resourceMWSWorkspaces(),
		},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_HOST", nil),
			},
			"token": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				DefaultFunc:   schema.EnvDefaultFunc("DATABRICKS_TOKEN", nil),
				ConflictsWith: []string{"basic_auth"},
			},
			"basic_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_USERNAME", nil),
						},
						"password": {
							Type:        schema.TypeString,
							Sensitive:   true,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_PASSWORD", nil),
						},
					},
				},
				ConflictsWith: []string{"token"},
			},
			"config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_CONFIG_FILE", "~/.databrickscfg"),
				Description: "Location of the Databricks CLI credentials file, that is created\n" +
					"by `databricks configure --token` command. By default, it is located\n" +
					"in ~/.databrickscfg. Check  https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for docs. Config\n" +
					"file credentials will only be used when host/token are not provided.",
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_CONFIG_PROFILE", "DEFAULT"),
				Description: "Connection profile specified within ~/.databrickscfg. Please check\n" +
					"https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.",
			},
			"azure_auth": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_resource_group": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP", nil),
						},
						"azure_region": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("AZURE_REGION", nil),
						},
						"workspace_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_WORKSPACE_NAME", nil),
						},
						"resource_group": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_RESOURCE_GROUP", nil),
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID"}, nil),
						},
						"client_secret": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}, nil),
						},
						"client_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}, nil),
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_TENANT_ID", "ARM_TENANT_ID"}, nil),
						},
						"pat_token_duration_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Currently secret scopes are not accessible via AAD tokens so we will need to create a PAT token",
							Default:     durationToSecondsString(time.Hour),
						},
					},
				},
			},
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		pc := service.DatabricksClient{}
		if host, ok := d.GetOk("host"); ok {
			pc.Host = host.(string)
		}
		if token, ok := d.GetOk("token"); ok {
			pc.Token = token.(string)
		}
		if _, ok := d.GetOk("basic_auth"); ok {
			username, userOk := d.GetOk("basic_auth.0.username")
			password, passOk := d.GetOk("basic_auth.0.password")
			if userOk && passOk {
				pc.BasicAuth.Username = fmt.Sprintf("%s", username)
				pc.BasicAuth.Password = fmt.Sprintf("%s", password)
			}
		}
		if aa, ok := d.GetOk("azure_auth"); ok { // TODO: i think this is a list here...
			// This provider takes DATABRICKS_AZURE_* for client ID etc
			// The azurerm provider uses ARM_* for the same values
			// To make it easier to use the two providers together we use the following sources in order:
			//  - provider config
			//  - DATABRICKS_AZURE_* environment variables
			//  - ARM_* environment variables
			azureAuth := aa.(map[string]interface{})
			if v, ok := azureAuth["managed_resource_group"]; ok {
				pc.AzureAuth.ManagedResourceGroup = v.(string)
			}
			if v, ok := azureAuth["azure_region"]; ok {
				pc.AzureAuth.AzureRegion = v.(string)
			}
			if v, ok := azureAuth["workspace_name"]; ok {
				pc.AzureAuth.WorkspaceName = v.(string)
			}
			if v, ok := azureAuth["resource_group"]; ok {
				pc.AzureAuth.ResourceGroup = v.(string)
			}
			if v, ok := azureAuth["subscription_id"]; ok {
				pc.AzureAuth.SubscriptionID = v.(string)
			}
			if v, ok := azureAuth["client_secret"]; ok {
				pc.AzureAuth.ClientSecret = v.(string)
			}
			if v, ok := azureAuth["client_id"]; ok {
				pc.AzureAuth.ClientID = v.(string)
			}
			if v, ok := azureAuth["tenant_id"]; ok {
				pc.AzureAuth.TenantID = v.(string)
			}
			if v, ok := azureAuth["pat_token_duration_seconds"]; ok {
				pc.AzureAuth.PATTokenDurationSeconds = v.(string)
			}
		}
		err := pc.Configure(version)
		if err != nil {
			return nil, err
		}
		return &pc, nil
	}

	return provider
}

func durationToSecondsString(duration time.Duration) string {
	return strconv.Itoa(int(duration.Seconds()))
}
