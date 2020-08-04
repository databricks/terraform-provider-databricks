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
				ConflictsWith: []string{
					"config_file",
					"azure_auth",
				},
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_TOKEN", nil),
				ConflictsWith: []string{
					"config_file",
					"profile",
					"basic_auth",
					"azure_auth",
				},
			},
			"basic_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ConflictsWith: []string{
					"config_file",
					"profile",
					"azure_auth",
					"token",
				},
				Deprecated: "basic_auth {} block is deprecated in favor of username & password properties " +
					"with more previctable behavior. This configuration attribute will be removed in 0.3.",
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
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_PASSWORD", nil),
			},
			"config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_CONFIG_FILE", nil),
				Description: "Location of the Databricks CLI credentials file, that is created\n" +
					"by `databricks configure --token` command. By default, it is located\n" +
					"in ~/.databrickscfg. Check  https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for docs. Config\n" +
					"file credentials will only be used when host/token are not provided.",
				ConflictsWith: []string{
					"azure_auth",
					"basic_auth",
					"token",
					"host",
				},
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_CONFIG_PROFILE", nil),
				Description: "Connection profile specified within ~/.databrickscfg. Please check\n" +
					"https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.",
				ConflictsWith: []string{
					"basic_auth",
					"token",
				},
			},
			"azure_workspace_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				// TODO: fix the naming...
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID", "AZURE_DATABRICKS_WORKSPACE_RESOURCE_ID"}, nil),
				ConflictsWith: []string{
					"azure_workspace_name",
					// "azure_resource_group",
					// "azure_subscription_id",
				},
			},
			"azure_workspace_name": {
				Type:     schema.TypeString,
				Optional: true,
				// TODO: think about MWS as well
				DefaultFunc:   schema.EnvDefaultFunc("DATABRICKS_AZURE_WORKSPACE_NAME", nil),
				ConflictsWith: []string{"azure_workspace_resource_id"},
			},
			"azure_resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_RESOURCE_GROUP", nil),
				// ConflictsWith: []string{"azure_workspace_resource_id"},
			},
			"azure_subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID"}, nil),
				// ConflictsWith: []string{"azure_workspace_resource_id"},
			},
			"azure_client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}, nil),
			},
			"azure_client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}, nil),
			},
			"azure_tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_TENANT_ID", "ARM_TENANT_ID"}, nil),
			},
			"azure_pat_token_duration_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Currently secret scopes are not accessible via AAD tokens so we will need to create a PAT token",
				Default:     durationToSecondsString(time.Hour),
			},
			"azure_auth": {
				Type:     schema.TypeMap,
				Optional: true,
				Deprecated: "azure_auth {} block is deprecated in favor of azure_* properties with more previctable behavior. " +
					"This configuration attribute will be removed in 0.3.",
				ConflictsWith: []string{
					"config_file",
					"basic_auth",
					"token",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// ADD workspace_url
						"managed_resource_group": {
							Optional:    true,
							Type:        schema.TypeString,
							Deprecated:  "This field is not used internally and will be removed in version 0.3",
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP", nil),
						},
						"azure_region": {
							Optional:    true,
							Type:        schema.TypeString,
							Deprecated:  "This field is not used internally and will be removed in version 0.3",
							DefaultFunc: schema.EnvDefaultFunc("AZURE_REGION", nil),
						},
						"workspace_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_workspace_name",
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_WORKSPACE_NAME", nil),
						},
						"resource_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_resource_group",
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_RESOURCE_GROUP", nil),
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_subscription_id",
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID"}, nil),
						},
						"client_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_client_secret",
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}, nil),
						},
						"client_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_client_id",
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}, nil),
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_tenant_id",
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_TENANT_ID", "ARM_TENANT_ID"}, nil),
						},
						"pat_token_duration_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This field is deprecated and will be removed in version 0.3. Please use azure_pat_token_duration_seconds",
							Description: "Currently secret scopes are not accessible via AAD tokens so we will need to create a PAT token",
							Default:     durationToSecondsString(time.Hour),
						},
					},
				},
			},
			"skip_verify": {
				Type:        schema.TypeBool,
				Description: "Skip SSL certificate verification for HTTP calls. Use at your own risk.",
				Optional:    true,
				Default:     false,
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
		if v, ok := d.GetOk("username"); ok {
			pc.Username = v.(string)
		}
		if v, ok := d.GetOk("password"); ok {
			pc.Password = v.(string)
		}
		if v, ok := d.GetOk("profile"); ok {
			pc.Profile = v.(string)
		}
		if _, ok := d.GetOk("basic_auth"); ok {
			username, userOk := d.GetOk("basic_auth.0.username")
			password, passOk := d.GetOk("basic_auth.0.password")
			if userOk && passOk {
				pc.Username = fmt.Sprintf("%s", username)
				pc.Password = fmt.Sprintf("%s", password)
			}
		}
		if v, ok := d.GetOk("azure_workspace_resource_id"); ok {
			pc.AzureAuth.ResourceID = v.(string)
		}
		if v, ok := d.GetOk("azure_workspace_name"); ok {
			pc.AzureAuth.WorkspaceName = v.(string)
		}
		if v, ok := d.GetOk("azure_resource_group"); ok {
			pc.AzureAuth.ResourceGroup = v.(string)
		}
		if v, ok := d.GetOk("azure_subscription_id"); ok {
			pc.AzureAuth.SubscriptionID = v.(string)
		}
		if v, ok := d.GetOk("azure_client_secret"); ok {
			pc.AzureAuth.ClientSecret = v.(string)
		}
		if v, ok := d.GetOk("azure_client_id"); ok {
			pc.AzureAuth.ClientID = v.(string)
		}
		if v, ok := d.GetOk("azure_tenant_id"); ok {
			pc.AzureAuth.TenantID = v.(string)
		}
		if v, ok := d.GetOk("azure_pat_token_duration_seconds"); ok {
			pc.AzureAuth.PATTokenDurationSeconds = v.(string)
		}
		if v, ok := d.GetOk("skip_verify"); ok {
			pc.InsecureSkipVerify = v.(bool)
		}
		if aa, ok := d.GetOk("azure_auth"); ok {
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
