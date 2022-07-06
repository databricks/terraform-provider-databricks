package provider

import (
	"context"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/access"
	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/mlflow"
	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/policies"
	"github.com/databricks/terraform-provider-databricks/pools"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/tokens"
	"github.com/databricks/terraform-provider-databricks/workspace"
)

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_aws_crossaccount_policy": aws.DataAwsCrossaccountPolicy(),
			"databricks_aws_assume_role_policy":  aws.DataAwsAssumeRolePolicy(),
			"databricks_aws_bucket_policy":       aws.DataAwsBucketPolicy(),
			"databricks_clusters":                clusters.DataSourceClusters(),
			"databricks_catalogs":                catalog.DataSourceCatalogs(),
			"databricks_current_user":            scim.DataSourceCurrentUser(),
			"databricks_dbfs_file":               storage.DataSourceDbfsFile(),
			"databricks_dbfs_file_paths":         storage.DataSourceDbfsFilePaths(),
			"databricks_group":                   scim.DataSourceGroup(),
			"databricks_jobs":                    jobs.DataSourceJobs(),
			"databricks_node_type":               clusters.DataSourceNodeType(),
			"databricks_notebook":                workspace.DataSourceNotebook(),
			"databricks_notebook_paths":          workspace.DataSourceNotebookPaths(),
			"databricks_schemas":                 catalog.DataSourceSchemas(),
			"databricks_service_principal":       scim.DataSourceServicePrincipal(),
			"databricks_service_principals":      scim.DataSourceServicePrincipals(),
			"databricks_spark_version":           clusters.DataSourceSparkVersion(),
			"databricks_tables":                  catalog.DataSourceTables(),
			"databricks_views":                   catalog.DataSourceViews(),
			"databricks_user":                    scim.DataSourceUser(),
			"databricks_zones":                   clusters.DataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_aws_s3_mount":                storage.ResourceAWSS3Mount(),
			"databricks_azure_adls_gen1_mount":       storage.ResourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount":       storage.ResourceAzureAdlsGen2Mount(),
			"databricks_azure_blob_mount":            storage.ResourceAzureBlobMount(),
			"databricks_catalog":                     catalog.ResourceCatalog(),
			"databricks_cluster":                     clusters.ResourceCluster(),
			"databricks_cluster_policy":              policies.ResourceClusterPolicy(),
			"databricks_dbfs_file":                   storage.ResourceDbfsFile(),
			"databricks_directory":                   workspace.ResourceDirectory(),
			"databricks_external_location":           catalog.ResourceExternalLocation(),
			"databricks_git_credential":              repos.ResourceGitCredential(),
			"databricks_global_init_script":          workspace.ResourceGlobalInitScript(),
			"databricks_grants":                      catalog.ResourceGrants(),
			"databricks_group":                       scim.ResourceGroup(),
			"databricks_group_instance_profile":      aws.ResourceGroupInstanceProfile(),
			"databricks_group_member":                scim.ResourceGroupMember(),
			"databricks_instance_pool":               pools.ResourceInstancePool(),
			"databricks_instance_profile":            aws.ResourceInstanceProfile(),
			"databricks_ip_access_list":              access.ResourceIPAccessList(),
			"databricks_job":                         jobs.ResourceJob(),
			"databricks_library":                     clusters.ResourceLibrary(),
			"databricks_metastore":                   catalog.ResourceMetastore(),
			"databricks_metastore_assignment":        catalog.ResourceMetastoreAssignment(),
			"databricks_metastore_data_access":       catalog.ResourceMetastoreDataAccess(),
			"databricks_mlflow_experiment":           mlflow.ResourceMlflowExperiment(),
			"databricks_mlflow_model":                mlflow.ResourceMlflowModel(),
			"databricks_mlflow_webhook":              mlflow.ResourceMlflowWebhook(),
			"databricks_mount":                       storage.ResourceMount(),
			"databricks_mws_customer_managed_keys":   mws.ResourceMwsCustomerManagedKeys(),
			"databricks_mws_credentials":             mws.ResourceMwsCredentials(),
			"databricks_mws_log_delivery":            mws.ResourceMwsLogDelivery(),
			"databricks_mws_networks":                mws.ResourceMwsNetworks(),
			"databricks_mws_private_access_settings": mws.ResourceMwsPrivateAccessSettings(),
			"databricks_mws_storage_configurations":  mws.ResourceMwsStorageConfigurations(),
			"databricks_mws_vpc_endpoint":            mws.ResourceMwsVpcEndpoint(),
			"databricks_mws_workspaces":              mws.ResourceMwsWorkspaces(),
			"databricks_notebook":                    workspace.ResourceNotebook(),
			"databricks_obo_token":                   tokens.ResourceOboToken(),
			"databricks_permissions":                 permissions.ResourcePermissions(),
			"databricks_pipeline":                    pipelines.ResourcePipeline(),
			"databricks_repo":                        repos.ResourceRepo(),
			"databricks_schema":                      catalog.ResourceSchema(),
			"databricks_secret":                      secrets.ResourceSecret(),
			"databricks_secret_scope":                secrets.ResourceSecretScope(),
			"databricks_secret_acl":                  secrets.ResourceSecretACL(),
			"databricks_service_principal":           scim.ResourceServicePrincipal(),
			"databricks_service_principal_role":      aws.ResourceServicePrincipalRole(),
			"databricks_sql_dashboard":               sql.ResourceSqlDashboard(),
			"databricks_sql_endpoint":                sql.ResourceSqlEndpoint(),
			"databricks_sql_global_config":           sql.ResourceSqlGlobalConfig(),
			"databricks_sql_permissions":             access.ResourceSqlPermissions(),
			"databricks_sql_query":                   sql.ResourceSqlQuery(),
			"databricks_sql_visualization":           sql.ResourceSqlVisualization(),
			"databricks_sql_widget":                  sql.ResourceSqlWidget(),
			"databricks_storage_credential":          catalog.ResourceStorageCredential(),
			"databricks_table":                       catalog.ResourceTable(),
			"databricks_token":                       tokens.ResourceToken(),
			"databricks_user":                        scim.ResourceUser(),
			"databricks_user_instance_profile":       aws.ResourceUserInstanceProfile(),
			"databricks_user_role":                   aws.ResourceUserRole(),
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
			Type:      kindMap[attr.Kind],
			Optional:  true,
			Sensitive: attr.Sensitive,
		}
		ps[attr.Name] = fieldSchema
		if len(attr.EnvVars) > 0 {
			fieldSchema.DefaultFunc = schema.MultiEnvDefaultFunc(attr.EnvVars, nil)
		}
	}
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
	if pc.AuthType == "" && len(authorizationMethodsUsed) > 1 {
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
