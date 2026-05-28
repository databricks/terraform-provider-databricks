package sdkv2

import (
	"github.com/databricks/terraform-provider-databricks/access"
	"github.com/databricks/terraform-provider-databricks/apps"
	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/dashboards"
	"github.com/databricks/terraform-provider-databricks/finops"
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
	"github.com/databricks/terraform-provider-databricks/settings"
	"github.com/databricks/terraform-provider-databricks/sharing"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/tokens"
	"github.com/databricks/terraform-provider-databricks/vectorsearch"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSources returns a new map of all data sources in the provider.
func DataSources() map[string]*schema.Resource {
	return toResourcesUnion(WorkspaceDataSources, AccountDataSources, DualDataSources)
}

// Resources returns a new map of all resources in the provider.
func Resources() map[string]*schema.Resource {
	return toResourcesUnion(WorkspaceResources, AccountResources, DualResources)
}

// WorkspaceDataSources is a map of all workspace data sources in the provider.
var WorkspaceDataSources = map[string]common.Resource{
	"databricks_aws_crossaccount_policy":              aws.DataAwsCrossaccountPolicy(),
	"databricks_aws_assume_role_policy":               aws.DataAwsAssumeRolePolicy(),
	"databricks_aws_bucket_policy":                    aws.DataAwsBucketPolicy(),
	"databricks_aws_unity_catalog_assume_role_policy": aws.DataAwsUnityCatalogAssumeRolePolicy(),
	"databricks_aws_unity_catalog_policy":             aws.DataAwsUnityCatalogPolicy(),
	"databricks_cluster":                              clusters.DataSourceCluster(),
	"databricks_clusters":                             clusters.DataSourceClusters(),
	"databricks_cluster_policy":                       policies.DataSourceClusterPolicy(),
	"databricks_catalog":                              catalog.DataSourceCatalog(),
	"databricks_catalogs":                             catalog.DataSourceCatalogs(),
	"databricks_current_metastore":                    catalog.DataSourceCurrentMetastore(),
	"databricks_current_user":                         scim.DataSourceCurrentUser(),
	"databricks_dbfs_file":                            storage.DataSourceDbfsFile(),
	"databricks_dbfs_file_paths":                      storage.DataSourceDbfsFilePaths(),
	"databricks_directory":                            workspace.DataSourceDirectory(),
	"databricks_external_location":                    catalog.DataSourceExternalLocation(),
	"databricks_external_locations":                   catalog.DataSourceExternalLocations(),
	"databricks_instance_pool":                        pools.DataSourceInstancePool(),
	"databricks_instance_profiles":                    aws.DataSourceInstanceProfiles(),
	"databricks_jobs":                                 jobs.DataSourceJobs(),
	"databricks_job":                                  jobs.DataSourceJob(),
	"databricks_metastores":                           catalog.DataSourceMetastores(),
	"databricks_mlflow_experiment":                    mlflow.DataSourceExperiment(),
	"databricks_mlflow_model":                         mlflow.DataSourceModel(),
	"databricks_mlflow_models":                        mlflow.DataSourceModels(),
	"databricks_node_type":                            clusters.DataSourceNodeType(),
	"databricks_notebook":                             workspace.DataSourceNotebook(),
	"databricks_notebook_paths":                       workspace.DataSourceNotebookPaths(),
	"databricks_pipelines":                            pipelines.DataSourcePipelines(),
	"databricks_schema":                               catalog.DataSourceSchema(),
	"databricks_schemas":                              catalog.DataSourceSchemas(),
	"databricks_share":                                sharing.DataSourceShare(),
	"databricks_shares":                               sharing.DataSourceShares(),
	"databricks_spark_version":                        clusters.DataSourceSparkVersion(),
	"databricks_sql_warehouse":                        sql.DataSourceWarehouse(),
	"databricks_sql_warehouses":                       sql.DataSourceWarehouses(),
	"databricks_storage_credential":                   catalog.DataSourceStorageCredential(),
	"databricks_storage_credentials":                  catalog.DataSourceStorageCredentials(),
	"databricks_table":                                catalog.DataSourceTable(),
	"databricks_tables":                               catalog.DataSourceTables(),
	"databricks_views":                                catalog.DataSourceViews(),
	"databricks_volume":                               catalog.DataSourceVolume(),
	"databricks_volumes":                              catalog.DataSourceVolumes(),
	"databricks_zones":                                clusters.DataSourceClusterZones(),
}

// AccountDataSources is a map of all account data sources in the provider.
var AccountDataSources = map[string]common.Resource{
	"databricks_metastore":                        catalog.DataSourceMetastore(),
	"databricks_mws_credentials":                  mws.DataSourceMwsCredentials(),
	"databricks_mws_network_connectivity_config":  mws.DataSourceMwsNetworkConnectivityConfig(),
	"databricks_mws_network_connectivity_configs": mws.DataSourceMwsNetworkConnectivityConfigs(),
	"databricks_mws_workspaces":                   mws.DataSourceMwsWorkspaces(),
}

// DualDataSources is a map of all dual data sources in the provider.
var DualDataSources = map[string]common.Resource{
	"databricks_current_config":     mws.DataSourceCurrentConfiguration(),
	"databricks_group":              scim.DataSourceGroup(),
	"databricks_service_principal":  scim.DataSourceServicePrincipal(),
	"databricks_service_principals": scim.DataSourceServicePrincipals(),
	"databricks_user":               scim.DataSourceUser(),
}

// WorkspaceResources is a map of all workspace resources in the provider.
var WorkspaceResources = map[string]common.Resource{
	"databricks_alert":                                sql.ResourceAlert(),
	"databricks_artifact_allowlist":                   catalog.ResourceArtifactAllowlist(),
	"databricks_aws_s3_mount":                         storage.ResourceAWSS3Mount(),
	"databricks_azure_adls_gen1_mount":                storage.ResourceAzureAdlsGen1Mount(),
	"databricks_azure_adls_gen2_mount":                storage.ResourceAzureAdlsGen2Mount(),
	"databricks_azure_blob_mount":                     storage.ResourceAzureBlobMount(),
	"databricks_catalog":                              catalog.ResourceCatalog(),
	"databricks_catalog_workspace_binding":            catalog.ResourceCatalogWorkspaceBinding(),
	"databricks_credential":                           catalog.ResourceCredential(),
	"databricks_connection":                           catalog.ResourceConnection(),
	"databricks_cluster":                              clusters.ResourceCluster(),
	"databricks_cluster_policy":                       policies.ResourceClusterPolicy(),
	"databricks_dashboard":                            dashboards.ResourceDashboard(),
	"databricks_dbfs_file":                            storage.ResourceDbfsFile(),
	"databricks_directory":                            workspace.ResourceDirectory(),
	"databricks_entitlements":                         scim.ResourceEntitlements(),
	"databricks_external_location":                    catalog.ResourceExternalLocation(),
	"databricks_file":                                 storage.ResourceFile(),
	"databricks_git_credential":                       repos.ResourceGitCredential(),
	"databricks_global_init_script":                   workspace.ResourceGlobalInitScript(),
	"databricks_grant":                                catalog.ResourceGrant(),
	"databricks_grants":                               catalog.ResourceGrants(),
	"databricks_instance_pool":                        pools.ResourceInstancePool(),
	"databricks_instance_profile":                     aws.ResourceInstanceProfile(),
	"databricks_ip_access_list":                       access.ResourceIPAccessList(),
	"databricks_job":                                  jobs.ResourceJob(),
	"databricks_lakehouse_monitor":                    catalog.ResourceLakehouseMonitor(),
	"databricks_library":                              clusters.ResourceLibrary(),
	"databricks_mlflow_experiment":                    mlflow.ResourceMlflowExperiment(),
	"databricks_mlflow_model":                         mlflow.ResourceMlflowModel(),
	"databricks_mlflow_webhook":                       mlflow.ResourceMlflowWebhook(),
	"databricks_model_serving":                        serving.ResourceModelServing(),
	"databricks_model_serving_provisioned_throughput": serving.ResourceModelServingProvisionedThroughput(),
	"databricks_mount":                                storage.ResourceMount(),
	"databricks_notebook":                             workspace.ResourceNotebook(),
	"databricks_notification_destination":             settings.ResourceNotificationDestination(),
	"databricks_obo_token":                            tokens.ResourceOboToken(),
	"databricks_online_table":                         catalog.ResourceOnlineTable(),
	"databricks_permission_assignment":                access.ResourcePermissionAssignment(),
	"databricks_permissions":                          permissions.ResourcePermissions(),
	"databricks_pipeline":                             pipelines.ResourcePipeline(),
	"databricks_provider":                             sharing.ResourceProvider(),
	"databricks_quality_monitor":                      catalog.ResourceQualityMonitor(),
	"databricks_query":                                sql.ResourceQuery(),
	"databricks_recipient":                            sharing.ResourceRecipient(),
	"databricks_registered_model":                     catalog.ResourceRegisteredModel(),
	"databricks_repo":                                 repos.ResourceRepo(),
	"databricks_schema":                               catalog.ResourceSchema(),
	"databricks_secret":                               secrets.ResourceSecret(),
	"databricks_secret_scope":                         secrets.ResourceSecretScope(),
	"databricks_secret_acl":                           secrets.ResourceSecretACL(),
	"databricks_share":                                sharing.ResourceShare(),
	"databricks_sql_dashboard":                        sql.ResourceSqlDashboard(),
	"databricks_sql_endpoint":                         sql.ResourceSqlEndpoint(),
	"databricks_sql_global_config":                    sql.ResourceSqlGlobalConfig(),
	"databricks_sql_permissions":                      access.ResourceSqlPermissions(),
	"databricks_sql_query":                            sql.ResourceSqlQuery(),
	"databricks_sql_alert":                            sql.ResourceSqlAlert(),
	"databricks_sql_table":                            catalog.ResourceSqlTable(),
	"databricks_sql_visualization":                    sql.ResourceSqlVisualization(),
	"databricks_sql_widget":                           sql.ResourceSqlWidget(),
	"databricks_system_schema":                        catalog.ResourceSystemSchema(),
	"databricks_table":                                catalog.ResourceTable(),
	"databricks_token":                                tokens.ResourceToken(),
	"databricks_vector_search_endpoint":               vectorsearch.ResourceVectorSearchEndpoint(),
	"databricks_vector_search_index":                  vectorsearch.ResourceVectorSearchIndex(),
	"databricks_volume":                               catalog.ResourceVolume(),
	"databricks_workspace_binding":                    catalog.ResourceWorkspaceBinding(),
	"databricks_workspace_conf":                       workspace.ResourceWorkspaceConf(),
	"databricks_workspace_file":                       workspace.ResourceWorkspaceFile(),
}

// AccountResources is a map of all account resources in the provider.
var AccountResources = map[string]common.Resource{
	"databricks_budget":                          finops.ResourceBudget(),
	"databricks_custom_app_integration":          apps.ResourceCustomAppIntegration(),
	"databricks_mws_credentials":                 mws.ResourceMwsCredentials(),
	"databricks_mws_customer_managed_keys":       mws.ResourceMwsCustomerManagedKeys(),
	"databricks_mws_log_delivery":                mws.ResourceMwsLogDelivery(),
	"databricks_mws_ncc_binding":                 mws.ResourceMwsNccBinding(),
	"databricks_mws_ncc_private_endpoint_rule":   mws.ResourceMwsNccPrivateEndpointRule(),
	"databricks_mws_network_connectivity_config": mws.ResourceMwsNetworkConnectivityConfig(),
	"databricks_mws_networks":                    mws.ResourceMwsNetworks(),
	"databricks_mws_permission_assignment":       mws.ResourceMwsPermissionAssignment(),
	"databricks_mws_private_access_settings":     mws.ResourceMwsPrivateAccessSettings(),
	"databricks_mws_storage_configurations":      mws.ResourceMwsStorageConfigurations(),
	"databricks_mws_vpc_endpoint":                mws.ResourceMwsVpcEndpoint(),
	"databricks_mws_workspaces":                  mws.ResourceMwsWorkspaces(),
}

// DualResources is a map of all dual resources in the provider.
var DualResources = map[string]common.Resource{
	"databricks_access_control_rule_set":  permissions.ResourceAccessControlRuleSet(),
	"databricks_group":                    scim.ResourceGroup(),
	"databricks_group_instance_profile":   aws.ResourceGroupInstanceProfile(),
	"databricks_group_member":             scim.ResourceGroupMember(),
	"databricks_group_role":               scim.ResourceGroupRole(),
	"databricks_metastore":                catalog.ResourceMetastore(),
	"databricks_metastore_assignment":     catalog.ResourceMetastoreAssignment(),
	"databricks_metastore_data_access":    catalog.ResourceMetastoreDataAccess(),
	"databricks_service_principal":        scim.ResourceServicePrincipal(),
	"databricks_service_principal_role":   aws.ResourceServicePrincipalRole(),
	"databricks_service_principal_secret": tokens.ResourceServicePrincipalSecret(),
	"databricks_storage_credential":       catalog.ResourceStorageCredential(),
	"databricks_user":                     scim.ResourceUser(),
	"databricks_user_instance_profile":    aws.ResourceUserInstanceProfile(),
	"databricks_user_role":                aws.ResourceUserRole(),
}

func toResourcesUnion(groups ...map[string]common.Resource) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource)
	for _, g := range groups {
		for k, v := range g {
			result[k] = v.ToResource()
		}
	}
	return result
}
