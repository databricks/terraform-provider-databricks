// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pluginfw

import (
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_federation_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_network_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_setting_user_preference_v2"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_setting_v2"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/alert_v2"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/app_space"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/apps_settings_custom_template"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/budget_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/data_classification_catalog_config"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/data_quality_monitor"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/data_quality_refresh"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_database_catalog"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_instance"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_synced_database_table"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/endpoint"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/entity_tag_assignment"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/environments_default_workspace_base_environment"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/environments_workspace_base_environment"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/external_metadata"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/feature_engineering_feature"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/feature_engineering_kafka_config"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/feature_engineering_materialized_feature"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/knowledge_assistant"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/knowledge_assistant_knowledge_source"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/materialized_features_feature_tag"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/online_store"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/policy_info"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_branch"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_catalog"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_database"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_endpoint"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_project"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_role"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_synced_table"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/quality_monitor_v2"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/rfa_access_request_destinations"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/secret_uc"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/service_principal_federation_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/supervisor_agent"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/supervisor_agent_tool"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/tag_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/warehouses_default_warehouse_override"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/workspace_entity_tag_assignment"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/workspace_network_option"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/workspace_setting_v2"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// List of resources that are auto generated based on service OpenAPI specs.
var autoGeneratedResources = []func() resource.Resource{
	account_federation_policy.ResourceFederationPolicy,
	account_network_policy.ResourceAccountNetworkPolicy,
	account_setting_user_preference_v2.ResourceUserPreference,
	account_setting_v2.ResourceSetting,
	alert_v2.ResourceAlertV2,
	app_space.ResourceSpace,
	apps_settings_custom_template.ResourceCustomTemplate,
	budget_policy.ResourceBudgetPolicy,
	data_classification_catalog_config.ResourceCatalogConfig,
	data_quality_monitor.ResourceMonitor,
	data_quality_refresh.ResourceRefresh,
	database_database_catalog.ResourceDatabaseCatalog,
	database_instance.ResourceDatabaseInstance,
	database_synced_database_table.ResourceSyncedDatabaseTable,
	endpoint.ResourceEndpoint,
	entity_tag_assignment.ResourceEntityTagAssignment,
	environments_default_workspace_base_environment.ResourceDefaultWorkspaceBaseEnvironment,
	environments_workspace_base_environment.ResourceWorkspaceBaseEnvironment,
	external_metadata.ResourceExternalMetadata,
	feature_engineering_feature.ResourceFeature,
	feature_engineering_kafka_config.ResourceKafkaConfig,
	feature_engineering_materialized_feature.ResourceMaterializedFeature,
	knowledge_assistant.ResourceKnowledgeAssistant,
	knowledge_assistant_knowledge_source.ResourceKnowledgeSource,
	materialized_features_feature_tag.ResourceFeatureTag,
	online_store.ResourceOnlineStore,
	policy_info.ResourcePolicyInfo,
	postgres_branch.ResourceBranch,
	postgres_catalog.ResourceCatalog,
	postgres_database.ResourceDatabase,
	postgres_endpoint.ResourceEndpoint,
	postgres_project.ResourceProject,
	postgres_role.ResourceRole,
	postgres_synced_table.ResourceSyncedTable,
	quality_monitor_v2.ResourceQualityMonitor,
	rfa_access_request_destinations.ResourceAccessRequestDestination,
	secret_uc.ResourceSecret,
	service_principal_federation_policy.ResourceFederationPolicy,
	supervisor_agent.ResourceSupervisorAgent,
	supervisor_agent_tool.ResourceTool,
	tag_policy.ResourceTagPolicy,
	warehouses_default_warehouse_override.ResourceDefaultWarehouseOverride,
	workspace_entity_tag_assignment.ResourceTagAssignment,
	workspace_network_option.ResourceWorkspaceNetworkOption,
	workspace_setting_v2.ResourceSetting,
}

// List of data sources that are auto generated based on service OpenAPI specs.
var autoGeneratedDataSources = []func() datasource.DataSource{
	account_federation_policy.DataSourceFederationPolicy,
	account_network_policy.DataSourceAccountNetworkPolicy,
	account_setting_user_preference_v2.DataSourceUserPreference,
	account_setting_v2.DataSourceSetting,
	alert_v2.DataSourceAlertV2,
	app_space.DataSourceSpace,
	apps_settings_custom_template.DataSourceCustomTemplate,
	budget_policy.DataSourceBudgetPolicy,
	data_classification_catalog_config.DataSourceCatalogConfig,
	data_quality_monitor.DataSourceMonitor,
	data_quality_refresh.DataSourceRefresh,
	database_database_catalog.DataSourceDatabaseCatalog,
	database_instance.DataSourceDatabaseInstance,
	database_synced_database_table.DataSourceSyncedDatabaseTable,
	endpoint.DataSourceEndpoint,
	entity_tag_assignment.DataSourceEntityTagAssignment,
	environments_default_workspace_base_environment.DataSourceDefaultWorkspaceBaseEnvironment,
	environments_workspace_base_environment.DataSourceWorkspaceBaseEnvironment,
	external_metadata.DataSourceExternalMetadata,
	feature_engineering_feature.DataSourceFeature,
	feature_engineering_kafka_config.DataSourceKafkaConfig,
	feature_engineering_materialized_feature.DataSourceMaterializedFeature,
	knowledge_assistant.DataSourceKnowledgeAssistant,
	knowledge_assistant_knowledge_source.DataSourceKnowledgeSource,
	materialized_features_feature_tag.DataSourceFeatureTag,
	online_store.DataSourceOnlineStore,
	policy_info.DataSourcePolicyInfo,
	postgres_branch.DataSourceBranch,
	postgres_catalog.DataSourceCatalog,
	postgres_database.DataSourceDatabase,
	postgres_endpoint.DataSourceEndpoint,
	postgres_project.DataSourceProject,
	postgres_role.DataSourceRole,
	postgres_synced_table.DataSourceSyncedTable,
	quality_monitor_v2.DataSourceQualityMonitor,
	rfa_access_request_destinations.DataSourceAccessRequestDestination,
	secret_uc.DataSourceSecret,
	service_principal_federation_policy.DataSourceFederationPolicy,
	supervisor_agent.DataSourceSupervisorAgent,
	supervisor_agent_tool.DataSourceTool,
	tag_policy.DataSourceTagPolicy,
	warehouses_default_warehouse_override.DataSourceDefaultWarehouseOverride,
	workspace_entity_tag_assignment.DataSourceTagAssignment,
	workspace_network_option.DataSourceWorkspaceNetworkOption,
	workspace_setting_v2.DataSourceSetting,
	account_federation_policy.DataSourceFederationPolicies,
	account_network_policy.DataSourceAccountNetworkPolicies,
	alert_v2.DataSourceAlertsV2,
	app_space.DataSourceSpaces,
	apps_settings_custom_template.DataSourceCustomTemplates,
	budget_policy.DataSourceBudgetPolicies,
	data_quality_monitor.DataSourceMonitors,
	data_quality_refresh.DataSourceRefreshes,
	database_database_catalog.DataSourceDatabaseCatalogs,
	database_instance.DataSourceDatabaseInstances,
	database_synced_database_table.DataSourceSyncedDatabaseTables,
	endpoint.DataSourceEndpoints,
	entity_tag_assignment.DataSourceEntityTagAssignments,
	environments_workspace_base_environment.DataSourceWorkspaceBaseEnvironments,
	external_metadata.DataSourceExternalMetadatas,
	feature_engineering_feature.DataSourceFeatures,
	feature_engineering_kafka_config.DataSourceKafkaConfigs,
	feature_engineering_materialized_feature.DataSourceMaterializedFeatures,
	knowledge_assistant.DataSourceKnowledgeAssistants,
	knowledge_assistant_knowledge_source.DataSourceKnowledgeSources,
	materialized_features_feature_tag.DataSourceFeatureTags,
	online_store.DataSourceOnlineStores,
	policy_info.DataSourcePolicyInfos,
	postgres_branch.DataSourceBranches,
	postgres_database.DataSourceDatabases,
	postgres_endpoint.DataSourceEndpoints,
	postgres_project.DataSourceProjects,
	postgres_role.DataSourceRoles,
	quality_monitor_v2.DataSourceQualityMonitors,
	secret_uc.DataSourceSecrets,
	service_principal_federation_policy.DataSourceFederationPolicies,
	supervisor_agent.DataSourceSupervisorAgents,
	supervisor_agent_tool.DataSourceTools,
	tag_policy.DataSourceTagPolicies,
	warehouses_default_warehouse_override.DataSourceDefaultWarehouseOverrides,
	workspace_entity_tag_assignment.DataSourceTagAssignments,
}
