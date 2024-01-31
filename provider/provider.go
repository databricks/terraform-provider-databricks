package provider

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	tflogger "github.com/databricks/terraform-provider-databricks/logger"
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
	"github.com/databricks/terraform-provider-databricks/workspace"
)

func init() {
	// IMPORTANT: this line cannot be changed, because it's used for
	// internal purposes at Databricks.
	useragent.WithProduct("databricks-tf-provider", common.Version())
}

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider() *schema.Provider {
	// must be in alphabetical order
	dataSources := map[string]common.DatabricksTerraformResource{
		"databricks_aws_crossaccount_policy": {Resource: aws.DataAwsCrossaccountPolicy()},
		"databricks_aws_assume_role_policy":  {Resource: aws.DataAwsAssumeRolePolicy()},
		"databricks_aws_bucket_policy":       {Resource: aws.DataAwsBucketPolicy()},
		"databricks_cluster":                 {Resource: clusters.DataSourceCluster(), IsWorkspaceLevel: true},
		"databricks_clusters":                {Resource: clusters.DataSourceClusters(), IsWorkspaceLevel: true},
		"databricks_cluster_policy":          {Resource: policies.DataSourceClusterPolicy(), IsWorkspaceLevel: true},
		"databricks_catalogs":                {Resource: catalog.DataSourceCatalogs(), IsWorkspaceLevel: true},
		"databricks_current_config":          {Resource: mws.DataSourceCurrentConfiguration(), IsWorkspaceLevel: true},
		"databricks_current_metastore":       {Resource: catalog.DataSourceCurrentMetastore(), IsWorkspaceLevel: true},
		"databricks_current_user":            {Resource: scim.DataSourceCurrentUser(), IsWorkspaceLevel: true},
		"databricks_dbfs_file":               {Resource: storage.DataSourceDbfsFile(), IsWorkspaceLevel: true},
		"databricks_dbfs_file_paths":         {Resource: storage.DataSourceDbfsFilePaths(), IsWorkspaceLevel: true},
		"databricks_directory":               {Resource: workspace.DataSourceDirectory(), IsWorkspaceLevel: true},
		"databricks_group":                   {Resource: scim.DataSourceGroup(), IsWorkspaceLevel: true},
		"databricks_instance_pool":           {Resource: pools.DataSourceInstancePool(), IsWorkspaceLevel: true},
		"databricks_instance_profiles":       {Resource: aws.DataSourceInstanceProfiles(), IsWorkspaceLevel: true},
		"databricks_jobs":                    {Resource: jobs.DataSourceJobs(), IsWorkspaceLevel: true},
		"databricks_job":                     {Resource: jobs.DataSourceJob(), IsWorkspaceLevel: true},
		"databricks_metastore":               {Resource: catalog.DataSourceMetastore(), IsWorkspaceLevel: true},
		"databricks_metastores":              {Resource: catalog.DataSourceMetastores(), IsWorkspaceLevel: true},
		"databricks_mlflow_model":            {Resource: mlflow.DataSourceModel(), IsWorkspaceLevel: true},
		"databricks_mws_credentials":         {Resource: mws.DataSourceMwsCredentials()},
		"databricks_mws_workspaces":          {Resource: mws.DataSourceMwsWorkspaces()},
		"databricks_node_type":               {Resource: clusters.DataSourceNodeType(), IsWorkspaceLevel: true},
		"databricks_notebook":                {Resource: workspace.DataSourceNotebook(), IsWorkspaceLevel: true},
		"databricks_notebook_paths":          {Resource: workspace.DataSourceNotebookPaths(), IsWorkspaceLevel: true},
		"databricks_pipelines":               {Resource: pipelines.DataSourcePipelines(), IsWorkspaceLevel: true},
		"databricks_schemas":                 {Resource: catalog.DataSourceSchemas(), IsWorkspaceLevel: true},
		"databricks_service_principal":       {Resource: scim.DataSourceServicePrincipal(), IsWorkspaceLevel: true},
		"databricks_service_principals":      {Resource: scim.DataSourceServicePrincipals(), IsWorkspaceLevel: true},
		"databricks_share":                   {Resource: catalog.DataSourceShare(), IsWorkspaceLevel: true},
		"databricks_shares":                  {Resource: catalog.DataSourceShares(), IsWorkspaceLevel: true},
		"databricks_spark_version":           {Resource: clusters.DataSourceSparkVersion(), IsWorkspaceLevel: true},
		"databricks_sql_warehouse":           {Resource: sql.DataSourceWarehouse(), IsWorkspaceLevel: true},
		"databricks_sql_warehouses":          {Resource: sql.DataSourceWarehouses(), IsWorkspaceLevel: true},
		"databricks_tables":                  {Resource: catalog.DataSourceTables(), IsWorkspaceLevel: true},
		"databricks_views":                   {Resource: catalog.DataSourceViews(), IsWorkspaceLevel: true},
		"databricks_user":                    {Resource: scim.DataSourceUser(), IsWorkspaceLevel: true},
		"databricks_zones":                   {Resource: clusters.DataSourceClusterZones(), IsWorkspaceLevel: true},
	}
	// must be in alphabetical order
	resources := map[string]common.DatabricksTerraformResource{
		"databricks_access_control_rule_set":     {Resource: permissions.ResourceAccessControlRuleSet(), IsWorkspaceLevel: true},
		"databricks_artifact_allowlist":          {Resource: catalog.ResourceArtifactAllowlist(), IsWorkspaceLevel: true},
		"databricks_aws_s3_mount":                {Resource: storage.ResourceAWSS3Mount(), IsWorkspaceLevel: true},
		"databricks_azure_adls_gen1_mount":       {Resource: storage.ResourceAzureAdlsGen1Mount(), IsWorkspaceLevel: true},
		"databricks_azure_adls_gen2_mount":       {Resource: storage.ResourceAzureAdlsGen2Mount(), IsWorkspaceLevel: true},
		"databricks_azure_blob_mount":            {Resource: storage.ResourceAzureBlobMount(), IsWorkspaceLevel: true},
		"databricks_catalog":                     {Resource: catalog.ResourceCatalog(), IsWorkspaceLevel: true},
		"databricks_catalog_workspace_binding":   {Resource: catalog.ResourceCatalogWorkspaceBinding(), IsWorkspaceLevel: true},
		"databricks_connection":                  {Resource: catalog.ResourceConnection(), IsWorkspaceLevel: true},
		"databricks_cluster":                     {Resource: clusters.ResourceCluster(), IsWorkspaceLevel: true},
		"databricks_cluster_policy":              {Resource: policies.ResourceClusterPolicy(), IsWorkspaceLevel: true},
		"databricks_dbfs_file":                   {Resource: storage.ResourceDbfsFile(), IsWorkspaceLevel: true},
		"databricks_directory":                   {Resource: workspace.ResourceDirectory(), IsWorkspaceLevel: true},
		"databricks_entitlements":                {Resource: scim.ResourceEntitlements(), IsWorkspaceLevel: true},
		"databricks_external_location":           {Resource: catalog.ResourceExternalLocation(), IsWorkspaceLevel: true},
		"databricks_git_credential":              {Resource: repos.ResourceGitCredential(), IsWorkspaceLevel: true},
		"databricks_global_init_script":          {Resource: workspace.ResourceGlobalInitScript(), IsWorkspaceLevel: true},
		"databricks_grant":                       {Resource: catalog.ResourceGrant(), IsWorkspaceLevel: true},
		"databricks_grants":                      {Resource: catalog.ResourceGrants(), IsWorkspaceLevel: true},
		"databricks_group":                       {Resource: scim.ResourceGroup(), IsWorkspaceLevel: true},
		"databricks_group_instance_profile":      {Resource: aws.ResourceGroupInstanceProfile(), IsWorkspaceLevel: true},
		"databricks_group_member":                {Resource: scim.ResourceGroupMember(), IsWorkspaceLevel: true},
		"databricks_group_role":                  {Resource: scim.ResourceGroupRole(), IsWorkspaceLevel: true},
		"databricks_instance_pool":               {Resource: pools.ResourceInstancePool(), IsWorkspaceLevel: true},
		"databricks_instance_profile":            {Resource: aws.ResourceInstanceProfile(), IsWorkspaceLevel: true},
		"databricks_ip_access_list":              {Resource: access.ResourceIPAccessList(), IsWorkspaceLevel: true},
		"databricks_job":                         {Resource: jobs.ResourceJob(), IsWorkspaceLevel: true},
		"databricks_library":                     {Resource: clusters.ResourceLibrary(), IsWorkspaceLevel: true},
		"databricks_metastore":                   {Resource: catalog.ResourceMetastore(), IsWorkspaceLevel: true},
		"databricks_metastore_assignment":        {Resource: catalog.ResourceMetastoreAssignment(), IsWorkspaceLevel: true},
		"databricks_metastore_data_access":       {Resource: catalog.ResourceMetastoreDataAccess(), IsWorkspaceLevel: true},
		"databricks_mlflow_experiment":           {Resource: mlflow.ResourceMlflowExperiment(), IsWorkspaceLevel: true},
		"databricks_mlflow_model":                {Resource: mlflow.ResourceMlflowModel(), IsWorkspaceLevel: true},
		"databricks_mlflow_webhook":              {Resource: mlflow.ResourceMlflowWebhook(), IsWorkspaceLevel: true},
		"databricks_model_serving":               {Resource: serving.ResourceModelServing(), IsWorkspaceLevel: true},
		"databricks_mount":                       {Resource: storage.ResourceMount(), IsWorkspaceLevel: true},
		"databricks_mws_customer_managed_keys":   {Resource: mws.ResourceMwsCustomerManagedKeys()},
		"databricks_mws_credentials":             {Resource: mws.ResourceMwsCredentials()},
		"databricks_mws_log_delivery":            {Resource: mws.ResourceMwsLogDelivery()},
		"databricks_mws_networks":                {Resource: mws.ResourceMwsNetworks()},
		"databricks_mws_permission_assignment":   {Resource: mws.ResourceMwsPermissionAssignment()},
		"databricks_mws_private_access_settings": {Resource: mws.ResourceMwsPrivateAccessSettings()},
		"databricks_mws_storage_configurations":  {Resource: mws.ResourceMwsStorageConfigurations()},
		"databricks_mws_vpc_endpoint":            {Resource: mws.ResourceMwsVpcEndpoint()},
		"databricks_mws_workspaces":              {Resource: mws.ResourceMwsWorkspaces()},
		"databricks_notebook":                    {Resource: workspace.ResourceNotebook(), IsWorkspaceLevel: true},
		"databricks_obo_token":                   {Resource: tokens.ResourceOboToken(), IsWorkspaceLevel: true},
		"databricks_permission_assignment":       {Resource: access.ResourcePermissionAssignment(), IsWorkspaceLevel: true},
		"databricks_permissions":                 {Resource: permissions.ResourcePermissions(), IsWorkspaceLevel: true},
		"databricks_pipeline":                    {Resource: pipelines.ResourcePipeline(), IsWorkspaceLevel: true},
		"databricks_provider":                    {Resource: catalog.ResourceProvider(), IsWorkspaceLevel: true},
		"databricks_recipient":                   {Resource: sharing.ResourceRecipient(), IsWorkspaceLevel: true},
		"databricks_registered_model":            {Resource: catalog.ResourceRegisteredModel(), IsWorkspaceLevel: true},
		"databricks_repo":                        {Resource: repos.ResourceRepo(), IsWorkspaceLevel: true},
		"databricks_schema":                      {Resource: catalog.ResourceSchema(), IsWorkspaceLevel: true},
		"databricks_secret":                      {Resource: secrets.ResourceSecret(), IsWorkspaceLevel: true},
		"databricks_secret_scope":                {Resource: secrets.ResourceSecretScope(), IsWorkspaceLevel: true},
		"databricks_secret_acl":                  {Resource: secrets.ResourceSecretACL(), IsWorkspaceLevel: true},
		"databricks_service_principal":           {Resource: scim.ResourceServicePrincipal(), IsWorkspaceLevel: true},
		"databricks_service_principal_role":      {Resource: aws.ResourceServicePrincipalRole(), IsWorkspaceLevel: true},
		"databricks_service_principal_secret":    {Resource: tokens.ResourceServicePrincipalSecret(), IsWorkspaceLevel: true},
		"databricks_share":                       {Resource: catalog.ResourceShare(), IsWorkspaceLevel: true},
		"databricks_sql_dashboard":               {Resource: sql.ResourceSqlDashboard(), IsWorkspaceLevel: true},
		"databricks_sql_endpoint":                {Resource: sql.ResourceSqlEndpoint(), IsWorkspaceLevel: true},
		"databricks_sql_global_config":           {Resource: sql.ResourceSqlGlobalConfig(), IsWorkspaceLevel: true},
		"databricks_sql_permissions":             {Resource: access.ResourceSqlPermissions(), IsWorkspaceLevel: true},
		"databricks_sql_query":                   {Resource: sql.ResourceSqlQuery(), IsWorkspaceLevel: true},
		"databricks_sql_alert":                   {Resource: sql.ResourceSqlAlert(), IsWorkspaceLevel: true},
		"databricks_sql_table":                   {Resource: catalog.ResourceSqlTable(), IsWorkspaceLevel: true},
		"databricks_sql_visualization":           {Resource: sql.ResourceSqlVisualization(), IsWorkspaceLevel: true},
		"databricks_sql_widget":                  {Resource: sql.ResourceSqlWidget(), IsWorkspaceLevel: true},
		"databricks_storage_credential":          {Resource: catalog.ResourceStorageCredential(), IsWorkspaceLevel: true},
		"databricks_system_schema":               {Resource: catalog.ResourceSystemSchema(), IsWorkspaceLevel: true},
		"databricks_table":                       {Resource: catalog.ResourceTable(), IsWorkspaceLevel: true},
		"databricks_token":                       {Resource: tokens.ResourceToken(), IsWorkspaceLevel: true},
		"databricks_user":                        {Resource: scim.ResourceUser(), IsWorkspaceLevel: true},
		"databricks_user_instance_profile":       {Resource: aws.ResourceUserInstanceProfile(), IsWorkspaceLevel: true},
		"databricks_user_role":                   {Resource: aws.ResourceUserRole(), IsWorkspaceLevel: true},
		"databricks_volume":                      {Resource: catalog.ResourceVolume(), IsWorkspaceLevel: true},
		"databricks_workspace_conf":              {Resource: workspace.ResourceWorkspaceConf(), IsWorkspaceLevel: true},
		"databricks_workspace_file":              {Resource: workspace.ResourceWorkspaceFile(), IsWorkspaceLevel: true},
	}
	for name, resource := range settings.AllSettingsResources() {
		resources[fmt.Sprintf("databricks_%s_setting", name)] = resource
	}

	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
		Schema:         providerSchema(),
	}
	for name, resource := range dataSources {
		p.DataSourcesMap[name] = resource.Resource
		if resource.IsWorkspaceLevel {
			common.AddWorkspaceIdField(resource.Resource.Schema)
		}
	}
	for name, resource := range resources {
		p.ResourcesMap[name] = resource.Resource
		if resource.IsWorkspaceLevel {
			common.AddWorkspaceIdField(resource.Resource.Schema)
		}
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		if p.TerraformVersion != "" {
			useragent.WithUserAgentExtra("terraform", p.TerraformVersion)
		}
		tflogger.SetLogger()
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
	tflog.Info(ctx, fmt.Sprintf("Explicit and implicit attributes: %s", strings.Join(attrsUsed, ", ")))
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
