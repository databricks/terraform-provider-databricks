package provider

import (
	"context"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"

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
	"github.com/databricks/terraform-provider-databricks/serving"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/tokens"
	"github.com/databricks/terraform-provider-databricks/workspace"
)

func init() {
	// IMPORTANT: this line cannot be changed, because it's used for
	// internal purposes at Databricks.
	useragent.WithProduct("databricks-tf-provider", common.Version())
}

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_aws_crossaccount_policy": aws.DataAwsCrossaccountPolicy(),
			"databricks_aws_assume_role_policy":  aws.DataAwsAssumeRolePolicy(),
			"databricks_aws_bucket_policy":       aws.DataAwsBucketPolicy(),
			"databricks_cluster":                 clusters.DataSourceCluster(),
			"databricks_clusters":                clusters.DataSourceClusters(),
			"databricks_cluster_policy":          policies.DataSourceClusterPolicy(),
			"databricks_catalogs":                catalog.DataSourceCatalogs(),
			"databricks_current_user":            scim.DataSourceCurrentUser(),
			"databricks_dbfs_file":               storage.DataSourceDbfsFile(),
			"databricks_dbfs_file_paths":         storage.DataSourceDbfsFilePaths(),
			"databricks_directory":               workspace.DataSourceDirectory(),
			"databricks_group":                   scim.DataSourceGroup(),
			"databricks_instance_pool":           pools.DataSourceInstancePool(),
			"databricks_jobs":                    jobs.DataSourceJobs(),
			"databricks_job":                     jobs.DataSourceJob(),
			"databricks_metastore":               catalog.DataSourceMetastore(),
			"databricks_metastores":              catalog.DataSourceMetastores(),
			"databricks_mws_credentials":         mws.DataSourceMwsCredentials(),
			"databricks_mws_workspaces":          mws.DataSourceMwsWorkspaces(),
			"databricks_node_type":               clusters.DataSourceNodeType(),
			"databricks_notebook":                workspace.DataSourceNotebook(),
			"databricks_notebook_paths":          workspace.DataSourceNotebookPaths(),
			"databricks_pipelines":               pipelines.DataSourcePipelines(),
			"databricks_schemas":                 catalog.DataSourceSchemas(),
			"databricks_service_principal":       scim.DataSourceServicePrincipal(),
			"databricks_service_principals":      scim.DataSourceServicePrincipals(),
			"databricks_share":                   catalog.DataSourceShare(),
			"databricks_shares":                  catalog.DataSourceShares(),
			"databricks_spark_version":           clusters.DataSourceSparkVersion(),
			"databricks_sql_warehouse":           sql.DataSourceWarehouse(),
			"databricks_sql_warehouses":          sql.DataSourceWarehouses(),
			"databricks_tables":                  catalog.DataSourceTables(),
			"databricks_views":                   catalog.DataSourceViews(),
			"databricks_user":                    scim.DataSourceUser(),
			"databricks_zones":                   clusters.DataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{ // must be in alphabetical order
			"databricks_access_control_rule_set":     permissions.ResourceAccessControlRuleSet(),
			"databricks_aws_s3_mount":                storage.ResourceAWSS3Mount(),
			"databricks_azure_adls_gen1_mount":       storage.ResourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount":       storage.ResourceAzureAdlsGen2Mount(),
			"databricks_azure_blob_mount":            storage.ResourceAzureBlobMount(),
			"databricks_catalog":                     catalog.ResourceCatalog(),
			"databricks_catalog_workspace_binding":   catalog.ResourceCatalogWorkspaceBinding(),
			"databricks_connection":                  catalog.ResourceConnection(),
			"databricks_cluster":                     clusters.ResourceCluster(),
			"databricks_cluster_policy":              policies.ResourceClusterPolicy(),
			"databricks_dbfs_file":                   storage.ResourceDbfsFile(),
			"databricks_directory":                   workspace.ResourceDirectory(),
			"databricks_entitlements":                scim.ResourceEntitlements(),
			"databricks_external_location":           catalog.ResourceExternalLocation(),
			"databricks_git_credential":              repos.ResourceGitCredential(),
			"databricks_global_init_script":          workspace.ResourceGlobalInitScript(),
			"databricks_grants":                      catalog.ResourceGrants(),
			"databricks_group":                       scim.ResourceGroup(),
			"databricks_group_instance_profile":      aws.ResourceGroupInstanceProfile(),
			"databricks_group_member":                scim.ResourceGroupMember(),
			"databricks_group_role":                  scim.ResourceGroupRole(),
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
			"databricks_model_serving":               serving.ResourceModelServing(),
			"databricks_mount":                       storage.ResourceMount(),
			"databricks_mws_customer_managed_keys":   mws.ResourceMwsCustomerManagedKeys(),
			"databricks_mws_credentials":             mws.ResourceMwsCredentials(),
			"databricks_mws_log_delivery":            mws.ResourceMwsLogDelivery(),
			"databricks_mws_networks":                mws.ResourceMwsNetworks(),
			"databricks_mws_permission_assignment":   mws.ResourceMwsPermissionAssignment(),
			"databricks_mws_private_access_settings": mws.ResourceMwsPrivateAccessSettings(),
			"databricks_mws_storage_configurations":  mws.ResourceMwsStorageConfigurations(),
			"databricks_mws_vpc_endpoint":            mws.ResourceMwsVpcEndpoint(),
			"databricks_mws_workspaces":              mws.ResourceMwsWorkspaces(),
			"databricks_notebook":                    workspace.ResourceNotebook(),
			"databricks_obo_token":                   tokens.ResourceOboToken(),
			"databricks_permission_assignment":       access.ResourcePermissionAssignment(),
			"databricks_permissions":                 permissions.ResourcePermissions(),
			"databricks_pipeline":                    pipelines.ResourcePipeline(),
			"databricks_provider":                    catalog.ResourceProvider(),
			"databricks_recipient":                   catalog.ResourceRecipient(),
			"databricks_repo":                        repos.ResourceRepo(),
			"databricks_schema":                      catalog.ResourceSchema(),
			"databricks_secret":                      secrets.ResourceSecret(),
			"databricks_secret_scope":                secrets.ResourceSecretScope(),
			"databricks_secret_acl":                  secrets.ResourceSecretACL(),
			"databricks_service_principal":           scim.ResourceServicePrincipal(),
			"databricks_service_principal_role":      aws.ResourceServicePrincipalRole(),
			"databricks_service_principal_secret":    tokens.ResourceServicePrincipalSecret(),
			"databricks_share":                       catalog.ResourceShare(),
			"databricks_sql_dashboard":               sql.ResourceSqlDashboard(),
			"databricks_sql_endpoint":                sql.ResourceSqlEndpoint(),
			"databricks_sql_global_config":           sql.ResourceSqlGlobalConfig(),
			"databricks_sql_permissions":             access.ResourceSqlPermissions(),
			"databricks_sql_query":                   sql.ResourceSqlQuery(),
			"databricks_sql_alert":                   sql.ResourceSqlAlert(),
			"databricks_sql_table":                   catalog.ResourceSqlTable(),
			"databricks_sql_visualization":           sql.ResourceSqlVisualization(),
			"databricks_sql_widget":                  sql.ResourceSqlWidget(),
			"databricks_storage_credential":          catalog.ResourceStorageCredential(),
			"databricks_system_schema":               catalog.ResourceSystemSchema(),
			"databricks_table":                       catalog.ResourceTable(),
			"databricks_token":                       tokens.ResourceToken(),
			"databricks_user":                        scim.ResourceUser(),
			"databricks_user_instance_profile":       aws.ResourceUserInstanceProfile(),
			"databricks_user_role":                   aws.ResourceUserRole(),
			"databricks_volume":                      catalog.ResourceVolume(),
			"databricks_workspace_conf":              workspace.ResourceWorkspaceConf(),
			"databricks_workspace_file":              workspace.ResourceWorkspaceFile(),
		},
		Schema: providerSchema(),
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		if p.TerraformVersion != "" {
			useragent.WithUserAgentExtra("terraform", p.TerraformVersion)
		}
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
	for _, attr := range config.ConfigAttributes {
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
	// TODO: check if still relevant
	ps["rate_limit"].DefaultFunc = schema.EnvDefaultFunc("DATABRICKS_RATE_LIMIT", 15)
	ps["debug_truncate_bytes"].DefaultFunc = schema.EnvDefaultFunc("DATABRICKS_DEBUG_TRUNCATE_BYTES", 96)
	return ps
}

func configureDatabricksClient(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	cfg := &config.Config{}
	attrsUsed := []string{}
	authsUsed := map[string]bool{}
	for _, attr := range config.ConfigAttributes {
		if value, ok := d.GetOk(attr.Name); ok {
			err := attr.Set(cfg, value)
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
	if cfg.AuthType != "" {
		// mapping from previous Google authentication types
		// and current authentication types from Databricks Go SDK
		oldToNewerAuthType := map[string]string{
			"google-creds":     "google-credentials",
			"google-accounts":  "google-id",
			"google-workspace": "google-id",
		}
		newer, ok := oldToNewerAuthType[cfg.AuthType]
		if ok {
			log.Printf("[INFO] Changing required auth_type from %s to %s", cfg.AuthType, newer)
			cfg.AuthType = newer
		}
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	pc := &common.DatabricksClient{
		DatabricksClient: client,
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return pc, nil
}
