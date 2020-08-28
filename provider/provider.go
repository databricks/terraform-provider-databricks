package provider

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	
	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/mws"
	"github.com/databrickslabs/databricks-terraform/storage"
	"github.com/databrickslabs/databricks-terraform/workspace"
)

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"databricks_aws_crossaccount_policy": access.DataAwsCrossAccountRolicy(),
			"databricks_aws_assume_role_policy":  access.DataAwsAssumeRolePolicy(),
			"databricks_aws_bucket_policy":       access.DataAwsBucketPolicy(),
			"databricks_dbfs_file":               storage.DataSourceDBFSFile(),
			"databricks_dbfs_file_paths":         storage.DataSourceDBFSFilePaths(),
			"databricks_default_user_roles":      identity.DataSourceDefaultUserRoles(),
			"databricks_notebook":                workspace.DataSourceNotebook(),
			"databricks_notebook_paths":          workspace.DataSourceNotebookPaths(),
			"databricks_zones":                   compute.DataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"databricks_secret":       access.ResourceSecret(),
			"databricks_secret_scope": access.ResourceSecretScope(),
			"databricks_secret_acl":   access.ResourceSecretACL(),
			"databricks_permissions":  access.ResourcePermissions(),

			"databricks_cluster":        compute.ResourceCluster(),
			"databricks_cluster_policy": compute.ResourceClusterPolicy(),
			"databricks_instance_pool":  compute.ResourceInstancePool(),
			"databricks_job":            compute.ResourceJob(),

			"databricks_group":                  identity.ResourceGroup(),
			"databricks_scim_group":             identity.ResourceScimGroup(),
			"databricks_group_instance_profile": identity.ResourceGroupInstanceProfile(),
			"databricks_instance_profile":       identity.ResourceInstanceProfile(),
			"databricks_group_member":           identity.ResourceGroupMember(),
			"databricks_scim_user":              identity.ResourceScimUser(),
			"databricks_token":                  identity.ResourceToken(),

			"databricks_mws_credentials":            mws.ResourceCredentials(),
			"databricks_mws_storage_configurations": mws.ResourceStorageConfiguration(),
			"databricks_mws_networks":               mws.ResourceNetwork(),
			"databricks_mws_workspaces":             mws.ResourceWorkspace(),

			"databricks_aws_s3_mount":          storage.ResourceAWSS3Mount(),
			"databricks_azure_adls_gen1_mount": storage.ResourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount": storage.ResourceAzureAdlsGen2Mount(),
			"databricks_azure_blob_mount":      storage.ResourceAzureBlobMount(),
			"databricks_dbfs_file":             storage.ResourceDBFSFile(),

			"databricks_notebook": workspace.ResourceNotebook(),
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
					"username",
					"password",
					"config_file",
					"basic_auth",
					"profile",
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
				ConflictsWith: []string{
					"token",
				},
			},
			"password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_PASSWORD", nil),
				ConflictsWith: []string{
					"token",
				},
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
			"azure_use_pat_for_cli": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Create ephemeral PAT tokens also for AZ CLI authenticated requests",
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
			"debug_truncate_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_DEBUG_TRUNCATE_BYTES", 96),
			},
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			pc := common.DatabricksClient{}

			authsUsed := map[string]bool{}
			if host, ok := d.GetOk("host"); ok {
				pc.Host = host.(string)
			}
			if token, ok := d.GetOk("token"); ok {
				authsUsed["token"] = true
				pc.Token = token.(string)
			}
			if v, ok := d.GetOk("username"); ok {
				authsUsed["password"] = true
				pc.Username = v.(string)
			}
			if v, ok := d.GetOk("password"); ok {
				authsUsed["password"] = true
				pc.Password = v.(string)
			}
			if v, ok := d.GetOk("profile"); ok {
				authsUsed["config profile"] = true
				pc.Profile = v.(string)
			}
			if _, ok := d.GetOk("basic_auth"); ok {
				authsUsed["password"] = true
				username, userOk := d.GetOk("basic_auth.0.username")
				password, passOk := d.GetOk("basic_auth.0.password")
				if userOk && passOk {
					pc.Username = fmt.Sprintf("%s", username)
					pc.Password = fmt.Sprintf("%s", password)
				}
			}
			if v, ok := d.GetOk("azure_workspace_resource_id"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.ResourceID = v.(string)
			}
			if v, ok := d.GetOk("azure_workspace_name"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.WorkspaceName = v.(string)
			}
			if v, ok := d.GetOk("azure_resource_group"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.ResourceGroup = v.(string)
			}
			if v, ok := d.GetOk("azure_subscription_id"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.SubscriptionID = v.(string)
			}
			if v, ok := d.GetOk("azure_client_secret"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.ClientSecret = v.(string)
			}
			if v, ok := d.GetOk("azure_client_id"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.ClientID = v.(string)
			}
			if v, ok := d.GetOk("azure_tenant_id"); ok {
				authsUsed["azure"] = true
				pc.AzureAuth.TenantID = v.(string)
			}
			if v, ok := d.GetOk("azure_pat_token_duration_seconds"); ok {
				pc.AzureAuth.PATTokenDurationSeconds = v.(string)
			}
			if v, ok := d.GetOk("skip_verify"); ok {
				pc.InsecureSkipVerify = v.(bool)
			}
			if v, ok := d.GetOk("debug_truncate_bytes"); ok {
				pc.DebugTruncateBytes = v.(int)
			}
			if v, ok := d.GetOk("azure_use_pat_for_cli"); ok {
				pc.AzureAuth.UsePATForCLI = v.(bool)
			}
			if aa, ok := d.GetOk("azure_auth"); ok {
				// This provider takes DATABRICKS_AZURE_* for client ID etc
				// The azurerm provider uses ARM_* for the same values
				// To make it easier to use the two providers together we use the following sources in order:
				//  - provider config
				//  - DATABRICKS_AZURE_* environment variables
				//  - ARM_* environment variables
				azureAuth := aa.(map[string]interface{})
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

			authorizationMethodsUsed := []string{}
			for name, used := range authsUsed {
				if used {
					authorizationMethodsUsed = append(authorizationMethodsUsed, name)
				}
			}
			if len(authorizationMethodsUsed) > 1 {
				sort.Strings(authorizationMethodsUsed)
				return nil, fmt.Errorf("More than one authorization method configured: %s",
					strings.Join(authorizationMethodsUsed, " and "))
			}
			err := pc.Configure()
			if err != nil {
				return nil, err
			}
			pc.WithCommandExecutor(compute.NewCommandsAPI(&pc))
			return &pc, nil
		},
	}
}

func durationToSecondsString(duration time.Duration) string {
	return strconv.Itoa(int(duration.Seconds()))
}
