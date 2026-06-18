package sdkv2

import (
	"maps"

	"github.com/databricks/terraform-provider-databricks/access"
	"github.com/databricks/terraform-provider-databricks/apps"
	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/clusters"
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
	return merge(WorkspaceDataSources(), AccountDataSources(), DualDataSources())
}

// Resources returns a new map of all resources in the provider.
func Resources() map[string]*schema.Resource {
	return merge(WorkspaceResources(), AccountResources(), DualResources())
}

// WorkspaceDataSources returns a new map of all workspace data sources in the
// provider.
//
// NOTE: this must build and return a fresh map of fresh *schema.Resource values
// on every call. DatabricksProvider() mutates each resource in place (see
// common.AddContextToAllResources), so a shared/cached resource would be
// re-wrapped on every provider build, growing the outgoing User-Agent without
// bound. The same applies to every group below.
func WorkspaceDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_aws_crossaccount_policy":              aws.DataAwsCrossaccountPolicy().ToResource(),
		"databricks_aws_assume_role_policy":               aws.DataAwsAssumeRolePolicy().ToResource(),
		"databricks_aws_bucket_policy":                    aws.DataAwsBucketPolicy().ToResource(),
		"databricks_aws_unity_catalog_assume_role_policy": aws.DataAwsUnityCatalogAssumeRolePolicy().ToResource(),
		"databricks_aws_unity_catalog_policy":             aws.DataAwsUnityCatalogPolicy().ToResource(),
		"databricks_cluster":                              clusters.DataSourceCluster().ToResource(),
		"databricks_clusters":                             clusters.DataSourceClusters().ToResource(),
		"databricks_cluster_policy":                       policies.DataSourceClusterPolicy().ToResource(),
		"databricks_catalog":                              catalog.DataSourceCatalog().ToResource(),
		"databricks_catalogs":                             catalog.DataSourceCatalogs().ToResource(),
		"databricks_current_metastore":                    catalog.DataSourceCurrentMetastore().ToResource(),
		"databricks_current_user":                         scim.DataSourceCurrentUser().ToResource(),
		"databricks_dbfs_file":                            storage.DataSourceDbfsFile().ToResource(),
		"databricks_dbfs_file_paths":                      storage.DataSourceDbfsFilePaths().ToResource(),
		"databricks_directory":                            workspace.DataSourceDirectory().ToResource(),
		"databricks_external_location":                    catalog.DataSourceExternalLocation().ToResource(),
		"databricks_external_locations":                   catalog.DataSourceExternalLocations().ToResource(),
		"databricks_instance_pool":                        pools.DataSourceInstancePool().ToResource(),
		"databricks_instance_profiles":                    aws.DataSourceInstanceProfiles().ToResource(),
		"databricks_jobs":                                 jobs.DataSourceJobs().ToResource(),
		"databricks_job":                                  jobs.DataSourceJob().ToResource(),
		"databricks_mlflow_experiment":                    mlflow.DataSourceExperiment().ToResource(),
		"databricks_mlflow_model":                         mlflow.DataSourceModel().ToResource(),
		"databricks_mlflow_models":                        mlflow.DataSourceModels().ToResource(),
		"databricks_node_type":                            clusters.DataSourceNodeType().ToResource(),
		"databricks_notebook":                             workspace.DataSourceNotebook().ToResource(),
		"databricks_notebook_paths":                       workspace.DataSourceNotebookPaths().ToResource(),
		"databricks_pipelines":                            pipelines.DataSourcePipelines().ToResource(),
		"databricks_schema":                               catalog.DataSourceSchema().ToResource(),
		"databricks_schemas":                              catalog.DataSourceSchemas().ToResource(),
		"databricks_share":                                sharing.DataSourceShare().ToResource(),
		"databricks_shares":                               sharing.DataSourceShares().ToResource(),
		"databricks_spark_version":                        clusters.DataSourceSparkVersion().ToResource(),
		"databricks_sql_warehouse":                        sql.DataSourceWarehouse().ToResource(),
		"databricks_sql_warehouses":                       sql.DataSourceWarehouses().ToResource(),
		"databricks_storage_credential":                   catalog.DataSourceStorageCredential().ToResource(),
		"databricks_storage_credentials":                  catalog.DataSourceStorageCredentials().ToResource(),
		"databricks_table":                                catalog.DataSourceTable().ToResource(),
		"databricks_tables":                               catalog.DataSourceTables().ToResource(),
		"databricks_views":                                catalog.DataSourceViews().ToResource(),
		"databricks_volume":                               catalog.DataSourceVolume().ToResource(),
		"databricks_volumes":                              catalog.DataSourceVolumes().ToResource(),
		"databricks_zones":                                clusters.DataSourceClusterZones().ToResource(),
	}
}

// AccountDataSources returns a new map of all account data sources in the
// provider.
func AccountDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_metastore":                        catalog.DataSourceMetastore().ToResource(),
		"databricks_metastores":                       catalog.DataSourceMetastores().ToResource(),
		"databricks_mws_credentials":                  mws.DataSourceMwsCredentials().ToResource(),
		"databricks_mws_network_connectivity_config":  mws.DataSourceMwsNetworkConnectivityConfig().ToResource(),
		"databricks_mws_network_connectivity_configs": mws.DataSourceMwsNetworkConnectivityConfigs().ToResource(),
		"databricks_mws_workspaces":                   mws.DataSourceMwsWorkspaces().ToResource(),
	}
}

// DualDataSources returns a new map of all dual data sources in the provider.
func DualDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_current_config":     mws.DataSourceCurrentConfiguration().ToResource(),
		"databricks_group":              scim.DataSourceGroup().ToResource(),
		"databricks_service_principal":  scim.DataSourceServicePrincipal().ToResource(),
		"databricks_service_principals": scim.DataSourceServicePrincipals().ToResource(),
		"databricks_user":               scim.DataSourceUser().ToResource(),
	}
}

// WorkspaceResources returns a new map of all workspace resources in the
// provider.
func WorkspaceResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_alert":                                sql.ResourceAlert().ToResource(),
		"databricks_artifact_allowlist":                   catalog.ResourceArtifactAllowlist().ToResource(),
		"databricks_aws_s3_mount":                         storage.ResourceAWSS3Mount().ToResource(),
		"databricks_azure_adls_gen1_mount":                storage.ResourceAzureAdlsGen1Mount().ToResource(),
		"databricks_azure_adls_gen2_mount":                storage.ResourceAzureAdlsGen2Mount().ToResource(),
		"databricks_azure_blob_mount":                     storage.ResourceAzureBlobMount().ToResource(),
		"databricks_catalog":                              catalog.ResourceCatalog().ToResource(),
		"databricks_catalog_workspace_binding":            catalog.ResourceCatalogWorkspaceBinding().ToResource(),
		"databricks_credential":                           catalog.ResourceCredential().ToResource(),
		"databricks_connection":                           catalog.ResourceConnection().ToResource(),
		"databricks_cluster":                              clusters.ResourceCluster().ToResource(),
		"databricks_cluster_policy":                       policies.ResourceClusterPolicy().ToResource(),
		"databricks_dashboard":                            dashboards.ResourceDashboard().ToResource(),
		"databricks_dbfs_file":                            storage.ResourceDbfsFile().ToResource(),
		"databricks_directory":                            workspace.ResourceDirectory().ToResource(),
		"databricks_entitlements":                         scim.ResourceEntitlements().ToResource(),
		"databricks_external_location":                    catalog.ResourceExternalLocation().ToResource(),
		"databricks_file":                                 storage.ResourceFile().ToResource(),
		"databricks_git_credential":                       repos.ResourceGitCredential().ToResource(),
		"databricks_global_init_script":                   workspace.ResourceGlobalInitScript().ToResource(),
		"databricks_grant":                                catalog.ResourceGrant().ToResource(),
		"databricks_grants":                               catalog.ResourceGrants().ToResource(),
		"databricks_instance_pool":                        pools.ResourceInstancePool().ToResource(),
		"databricks_instance_profile":                     aws.ResourceInstanceProfile().ToResource(),
		"databricks_ip_access_list":                       access.ResourceIPAccessList().ToResource(),
		"databricks_job":                                  jobs.ResourceJob().ToResource(),
		"databricks_lakehouse_monitor":                    catalog.ResourceLakehouseMonitor().ToResource(),
		"databricks_library":                              clusters.ResourceLibrary().ToResource(),
		"databricks_mlflow_experiment":                    mlflow.ResourceMlflowExperiment().ToResource(),
		"databricks_mlflow_model":                         mlflow.ResourceMlflowModel().ToResource(),
		"databricks_mlflow_webhook":                       mlflow.ResourceMlflowWebhook().ToResource(),
		"databricks_model_serving":                        serving.ResourceModelServing().ToResource(),
		"databricks_model_serving_provisioned_throughput": serving.ResourceModelServingProvisionedThroughput().ToResource(),
		"databricks_mount":                                storage.ResourceMount().ToResource(),
		"databricks_notebook":                             workspace.ResourceNotebook().ToResource(),
		"databricks_notification_destination":             settings.ResourceNotificationDestination().ToResource(),
		"databricks_obo_token":                            tokens.ResourceOboToken().ToResource(),
		"databricks_online_table":                         catalog.ResourceOnlineTable().ToResource(),
		"databricks_permission_assignment":                access.ResourcePermissionAssignment().ToResource(),
		"databricks_permissions":                          permissions.ResourcePermissions().ToResource(),
		"databricks_pipeline":                             pipelines.ResourcePipeline().ToResource(),
		"databricks_provider":                             sharing.ResourceProvider().ToResource(),
		"databricks_quality_monitor":                      catalog.ResourceQualityMonitor().ToResource(),
		"databricks_query":                                sql.ResourceQuery().ToResource(),
		"databricks_recipient":                            sharing.ResourceRecipient().ToResource(),
		"databricks_registered_model":                     catalog.ResourceRegisteredModel().ToResource(),
		"databricks_repo":                                 repos.ResourceRepo().ToResource(),
		"databricks_schema":                               catalog.ResourceSchema().ToResource(),
		"databricks_secret":                               secrets.ResourceSecret().ToResource(),
		"databricks_secret_scope":                         secrets.ResourceSecretScope().ToResource(),
		"databricks_secret_acl":                           secrets.ResourceSecretACL().ToResource(),
		"databricks_share":                                sharing.ResourceShare().ToResource(),
		"databricks_sql_dashboard":                        sql.ResourceSqlDashboard().ToResource(),
		"databricks_sql_endpoint":                         sql.ResourceSqlEndpoint().ToResource(),
		"databricks_sql_global_config":                    sql.ResourceSqlGlobalConfig().ToResource(),
		"databricks_sql_permissions":                      access.ResourceSqlPermissions().ToResource(),
		"databricks_sql_query":                            sql.ResourceSqlQuery().ToResource(),
		"databricks_sql_alert":                            sql.ResourceSqlAlert().ToResource(),
		"databricks_sql_table":                            catalog.ResourceSqlTable().ToResource(),
		"databricks_sql_visualization":                    sql.ResourceSqlVisualization().ToResource(),
		"databricks_sql_widget":                           sql.ResourceSqlWidget().ToResource(),
		"databricks_system_schema":                        catalog.ResourceSystemSchema().ToResource(),
		"databricks_table":                                catalog.ResourceTable().ToResource(),
		"databricks_token":                                tokens.ResourceToken().ToResource(),
		"databricks_vector_search_endpoint":               vectorsearch.ResourceVectorSearchEndpoint().ToResource(),
		"databricks_vector_search_index":                  vectorsearch.ResourceVectorSearchIndex().ToResource(),
		"databricks_volume":                               catalog.ResourceVolume().ToResource(),
		"databricks_workspace_binding":                    catalog.ResourceWorkspaceBinding().ToResource(),
		"databricks_workspace_conf":                       workspace.ResourceWorkspaceConf().ToResource(),
		"databricks_workspace_file":                       workspace.ResourceWorkspaceFile().ToResource(),
	}
}

// AccountResources returns a new map of all account resources in the provider.
func AccountResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_budget":                          finops.ResourceBudget().ToResource(),
		"databricks_custom_app_integration":          apps.ResourceCustomAppIntegration().ToResource(),
		"databricks_mws_credentials":                 mws.ResourceMwsCredentials().ToResource(),
		"databricks_mws_customer_managed_keys":       mws.ResourceMwsCustomerManagedKeys().ToResource(),
		"databricks_mws_log_delivery":                mws.ResourceMwsLogDelivery().ToResource(),
		"databricks_mws_ncc_binding":                 mws.ResourceMwsNccBinding().ToResource(),
		"databricks_mws_ncc_private_endpoint_rule":   mws.ResourceMwsNccPrivateEndpointRule().ToResource(),
		"databricks_mws_network_connectivity_config": mws.ResourceMwsNetworkConnectivityConfig().ToResource(),
		"databricks_mws_networks":                    mws.ResourceMwsNetworks().ToResource(),
		"databricks_mws_permission_assignment":       mws.ResourceMwsPermissionAssignment().ToResource(),
		"databricks_mws_private_access_settings":     mws.ResourceMwsPrivateAccessSettings().ToResource(),
		"databricks_mws_storage_configurations":      mws.ResourceMwsStorageConfigurations().ToResource(),
		"databricks_mws_vpc_endpoint":                mws.ResourceMwsVpcEndpoint().ToResource(),
		"databricks_mws_workspaces":                  mws.ResourceMwsWorkspaces().ToResource(),
	}
}

// DualResources returns a new map of all dual resources in the provider.
func DualResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"databricks_access_control_rule_set":  permissions.ResourceAccessControlRuleSet().ToResource(),
		"databricks_group":                    scim.ResourceGroup().ToResource(),
		"databricks_group_instance_profile":   aws.ResourceGroupInstanceProfile().ToResource(),
		"databricks_group_member":             scim.ResourceGroupMember().ToResource(),
		"databricks_group_role":               scim.ResourceGroupRole().ToResource(),
		"databricks_metastore":                catalog.ResourceMetastore().ToResource(),
		"databricks_metastore_assignment":     catalog.ResourceMetastoreAssignment().ToResource(),
		"databricks_metastore_data_access":    catalog.ResourceMetastoreDataAccess().ToResource(),
		"databricks_service_principal":        scim.ResourceServicePrincipal().ToResource(),
		"databricks_service_principal_role":   aws.ResourceServicePrincipalRole().ToResource(),
		"databricks_service_principal_secret": tokens.ResourceServicePrincipalSecret().ToResource(),
		"databricks_storage_credential":       catalog.ResourceStorageCredential().ToResource(),
		"databricks_user":                     scim.ResourceUser().ToResource(),
		"databricks_user_instance_profile":    aws.ResourceUserInstanceProfile().ToResource(),
		"databricks_user_role":                aws.ResourceUserRole().ToResource(),
	}
}

func merge(groups ...map[string]*schema.Resource) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource)
	for _, g := range groups {
		maps.Copy(result, g)
	}
	return result
}
