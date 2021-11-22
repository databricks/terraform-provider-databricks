package provider

import (
	"context"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/terraform-provider-databricks/access"
	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/jobs"
	"github.com/databrickslabs/terraform-provider-databricks/mws"
	"github.com/databrickslabs/terraform-provider-databricks/permissions"
	"github.com/databrickslabs/terraform-provider-databricks/pipelines"
	"github.com/databrickslabs/terraform-provider-databricks/policies"
	"github.com/databrickslabs/terraform-provider-databricks/pools"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics"
	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/databrickslabs/terraform-provider-databricks/workspace"
)

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_aws_crossaccount_policy": access.DataAwsCrossAccountPolicy(),
			"databricks_aws_assume_role_policy":  access.DataAwsAssumeRolePolicy(),
			"databricks_aws_bucket_policy":       access.DataAwsBucketPolicy(),
			"databricks_current_user":            identity.DataSourceCurrentUser(),
			"databricks_dbfs_file":               storage.DataSourceDBFSFile(),
			"databricks_dbfs_file_paths":         storage.DataSourceDBFSFilePaths(),
			"databricks_group":                   identity.DataSourceGroup(),
			"databricks_node_type":               clusters.DataSourceNodeType(),
			"databricks_notebook":                workspace.DataSourceNotebook(),
			"databricks_notebook_paths":          workspace.DataSourceNotebookPaths(),
			"databricks_spark_version":           clusters.DataSourceSparkVersion(),
			"databricks_user":                    identity.DataSourceUser(),
			"databricks_zones":                   clusters.DataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_aws_s3_mount":                storage.ResourceAWSS3Mount(),
			"databricks_azure_adls_gen1_mount":       storage.ResourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount":       storage.ResourceAzureAdlsGen2Mount(),
			"databricks_azure_blob_mount":            storage.ResourceAzureBlobMount(),
			"databricks_cluster":                     clusters.ResourceCluster(),
			"databricks_cluster_policy":              policies.ResourceClusterPolicy(),
			"databricks_dbfs_file":                   storage.ResourceDBFSFile(),
			"databricks_directory":                   workspace.ResourceDirectory(),
			"databricks_global_init_script":          workspace.ResourceGlobalInitScript(),
			"databricks_group":                       identity.ResourceGroup(),
			"databricks_group_instance_profile":      identity.ResourceGroupInstanceProfile(),
			"databricks_group_member":                identity.ResourceGroupMember(),
			"databricks_instance_pool":               pools.ResourceInstancePool(),
			"databricks_instance_profile":            identity.ResourceInstanceProfile(),
			"databricks_ip_access_list":              access.ResourceIPAccessList(),
			"databricks_job":                         jobs.ResourceJob(),
			"databricks_mount":                       storage.ResourceDatabricksMount(),
			"databricks_mws_customer_managed_keys":   mws.ResourceCustomerManagedKey(),
			"databricks_mws_credentials":             mws.ResourceCredentials(),
			"databricks_mws_log_delivery":            mws.ResourceLogDelivery(),
			"databricks_mws_networks":                mws.ResourceNetwork(),
			"databricks_mws_private_access_settings": mws.ResourcePrivateAccessSettings(),
			"databricks_mws_storage_configurations":  mws.ResourceStorageConfiguration(),
			"databricks_mws_vpc_endpoint":            mws.ResourceVPCEndpoint(),
			"databricks_mws_workspaces":              mws.ResourceWorkspace(),
			"databricks_notebook":                    workspace.ResourceNotebook(),
			"databricks_obo_token":                   identity.ResourceOboToken(),
			"databricks_permissions":                 permissions.ResourcePermissions(),
			"databricks_pipeline":                    pipelines.ResourcePipeline(),
			"databricks_repo":                        workspace.ResourceRepo(),
			"databricks_secret":                      access.ResourceSecret(),
			"databricks_secret_scope":                access.ResourceSecretScope(),
			"databricks_secret_acl":                  access.ResourceSecretACL(),
			"databricks_service_principal":           identity.ResourceServicePrincipal(),
			"databricks_sql_dashboard":               sqlanalytics.ResourceDashboard(),
			"databricks_sql_endpoint":                sqlanalytics.ResourceSQLEndpoint(),
			"databricks_sql_global_config":           sqlanalytics.ResourceSQLGlobalConfig(),
			"databricks_sql_permissions":             access.ResourceSqlPermissions(),
			"databricks_sql_query":                   sqlanalytics.ResourceQuery(),
			"databricks_sql_visualization":           sqlanalytics.ResourceVisualization(),
			"databricks_sql_widget":                  sqlanalytics.ResourceWidget(),
			"databricks_token":                       identity.ResourceToken(),
			"databricks_user":                        identity.ResourceUser(),
			"databricks_user_instance_profile":       identity.ResourceUserInstanceProfile(),
			"databricks_workspace_conf":              workspace.ResourceWorkspaceConf(),
		},
		Schema: providerSchema(),
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		ctx = context.WithValue(ctx, common.Provider, p)
		return configureDatabricksClient(ctx, d)
	}
	common.AddContextToAllResources(p, "databricks")
	return p
}

func providerSchema() map[string]*schema.Schema {
	kindMap := map[reflect.Kind]schema.ValueType{
		reflect.String: schema.TypeString,
		reflect.Bool:   schema.TypeBool,
		reflect.Int:    schema.TypeInt,
		// other values will immediately fail unit tests
	}
	ps := map[string]*schema.Schema{}
	attrs := common.ClientAttributes() // TODO: pass by argument
	for _, attr := range attrs {
		fieldSchema := &schema.Schema{
			Type:     kindMap[attr.Kind],
			Optional: true,
		}
		ps[attr.Name] = fieldSchema
		if len(attr.EnvVars) > 0 {
			fieldSchema.DefaultFunc = schema.MultiEnvDefaultFunc(attr.EnvVars, nil)
		}
	}
	ps["token"].Sensitive = true
	ps["azure_client_secret"].Sensitive = true
	ps["rate_limit"].DefaultFunc = schema.EnvDefaultFunc("DATABRICKS_RATE_LIMIT",
		common.DefaultRateLimitPerSecond)
	ps["debug_truncate_bytes"].DefaultFunc = schema.EnvDefaultFunc("DATABRICKS_DEBUG_TRUNCATE_BYTES",
		common.DefaultTruncateBytes)
	return ps
}

func configureDatabricksClient(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	prov := ctx.Value(common.Provider).(*schema.Provider)
	pc := common.DatabricksClient{
		Provider: prov,
	}
	attrsUsed := []string{}
	authsUsed := map[string]bool{}
	attrs := common.ClientAttributes() // todo: pass by argument
	for _, attr := range attrs {
		if value, ok := d.GetOk(attr.Name); ok {
			err := attr.Set(&pc, value)
			if err != nil {
				return nil, diag.FromErr(err)
			}
			if attr.Kind == reflect.String {
				attrsUsed = append(attrsUsed, attr.Name)
			}
			if attr.Auth != "" {
				authsUsed[attr.Auth] = true
			}
		}
	}
	sort.Strings(attrsUsed)
	log.Printf("[INFO] Explicit and implicit attributes: %s", strings.Join(attrsUsed, ", "))
	authorizationMethodsUsed := []string{}
	for name, used := range authsUsed {
		if used {
			authorizationMethodsUsed = append(authorizationMethodsUsed, name)
		}
	}
	if len(authorizationMethodsUsed) > 1 {
		sort.Strings(authorizationMethodsUsed)
		return nil, diag.Errorf("More than one authorization method configured: %s",
			strings.Join(authorizationMethodsUsed, " and "))
	}
	if err := pc.Configure(attrsUsed...); err != nil {
		return nil, diag.FromErr(err)
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return &pc, nil
}
