// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package pipelines_tf

import (
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names" tf:"optional"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]types.String `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment []PipelineDeployment `tfsdk:"deployment" tf:"optional"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`

	DryRun types.Bool `tfsdk:"dry_run" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters []Filters `tfsdk:"filters" tf:"optional"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition []IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition" tf:"optional"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition []IngestionPipelineDefinition `tfsdk:"ingestion_definition" tf:"optional"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger []PipelineTrigger `tfsdk:"trigger" tf:"optional"`
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings []PipelineSpec `tfsdk:"effective_settings" tf:"optional"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule" tf:"optional"`

	TimezoneId types.String `tfsdk:"timezone_id" tf:"optional"`
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	Instance types.String `tfsdk:"instance" tf:"optional"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no" tf:"optional"`
}

// Delete a pipeline
type DeletePipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

type DeletePipelineResponse struct {
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names" tf:"optional"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]types.String `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment []PipelineDeployment `tfsdk:"deployment" tf:"optional"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters []Filters `tfsdk:"filters" tf:"optional"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition []IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition" tf:"optional"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition []IngestionPipelineDefinition `tfsdk:"ingestion_definition" tf:"optional"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Unique identifier for this pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger []PipelineTrigger `tfsdk:"trigger" tf:"optional"`
}

type EditPipelineResponse struct {
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions []SerializedException `tfsdk:"exceptions" tf:"optional"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal" tf:"optional"`
}

type FileLibrary struct {
	// The absolute path of the file.
	Path types.String `tfsdk:"path" tf:"optional"`
}

type Filters struct {
	// Paths to exclude.
	Exclude []types.String `tfsdk:"exclude" tf:"optional"`
	// Paths to include.
	Include []types.String `tfsdk:"include" tf:"optional"`
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PipelinePermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get pipeline permissions
type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

// Get a pipeline
type GetPipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"optional"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health" tf:"optional"`
	// The last time the pipeline settings were modified or created.
	LastModified types.Int64 `tfsdk:"last_modified" tf:"optional"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `tfsdk:"latest_updates" tf:"optional"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name" tf:"optional"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec []PipelineSpec `tfsdk:"spec" tf:"optional"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

// Get a pipeline update
type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
}

type GetUpdateResponse struct {
	// The current update info.
	Update []UpdateInfo `tfsdk:"update" tf:"optional"`
}

type IngestionConfig struct {
	// Select tables from a specific source schema.
	Schema []SchemaSpec `tfsdk:"schema" tf:"optional"`
	// Select tables from a specific source table.
	Table []TableSpec `tfsdk:"table" tf:"optional"`
}

type IngestionGatewayPipelineDefinition struct {
	// Immutable. The Unity Catalog connection this gateway pipeline uses to
	// communicate with the source.
	ConnectionId types.String `tfsdk:"connection_id" tf:"optional"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog types.String `tfsdk:"gateway_storage_catalog" tf:"optional"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName types.String `tfsdk:"gateway_storage_name" tf:"optional"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema types.String `tfsdk:"gateway_storage_schema" tf:"optional"`
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection this ingestion pipeline uses to
	// communicate with the source. Specify either ingestion_gateway_id or
	// connection_name.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Immutable. Identifier for the ingestion gateway used by this ingestion
	// pipeline to communicate with the source. Specify either
	// ingestion_gateway_id or connection_name.
	IngestionGatewayId types.String `tfsdk:"ingestion_gateway_id" tf:"optional"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects []IngestionConfig `tfsdk:"objects" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration []TableSpecificConfig `tfsdk:"table_configuration" tf:"optional"`
}

// List pipeline events
type ListPipelineEventsRequest struct {
	// Criteria to select a subset of results, expressed using a SQL-like
	// syntax. The supported filters are: 1. level='INFO' (or WARN or ERROR) 2.
	// level in ('INFO', 'WARN') 3. id='[event-id]' 4. timestamp > 'TIMESTAMP'
	// (or >=,<,<=,=)
	//
	// Composite expressions are supported, for example: level in ('ERROR',
	// 'WARN') AND timestamp> '2021-07-22T06:37:33.083Z'
	Filter types.String `tfsdk:"-"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults types.Int64 `tfsdk:"-"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []types.String `tfsdk:"-"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken types.String `tfsdk:"-"`

	PipelineId types.String `tfsdk:"-"`
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	Events []PipelineEvent `tfsdk:"events" tf:"optional"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

// List pipelines
type ListPipelinesRequest struct {
	// Select a subset of results based on the specified criteria. The supported
	// filters are:
	//
	// * `notebook='<path>'` to select pipelines that reference the provided
	// notebook path. * `name LIKE '[pattern]'` to select pipelines with a name
	// that matches pattern. Wildcards are supported, for example: `name LIKE
	// '%shopping%'`
	//
	// Composite filters are not supported. This field is optional.
	Filter types.String `tfsdk:"-"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults types.Int64 `tfsdk:"-"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy []types.String `tfsdk:"-"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-"`
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// The list of events matching the request criteria.
	Statuses []PipelineStateInfo `tfsdk:"statuses" tf:"optional"`
}

// List pipeline updates
type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults types.Int64 `tfsdk:"-"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-"`
	// The pipeline to return updates for.
	PipelineId types.String `tfsdk:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId types.String `tfsdk:"-"`
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`

	Updates []UpdateInfo `tfsdk:"updates" tf:"optional"`
}

type ManualTrigger struct {
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	Path types.String `tfsdk:"path" tf:"optional"`
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	Alerts []types.String `tfsdk:"alerts" tf:"optional"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients []types.String `tfsdk:"email_recipients" tf:"optional"`
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId types.Int64 `tfsdk:"batch_id" tf:"optional"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName types.String `tfsdk:"dataset_name" tf:"optional"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId types.String `tfsdk:"flow_id" tf:"optional"`
	// The name of the flow. Not unique.
	FlowName types.String `tfsdk:"flow_name" tf:"optional"`
	// The optional host name where the event was triggered
	Host types.String `tfsdk:"host" tf:"optional"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId types.String `tfsdk:"maintenance_id" tf:"optional"`
	// Materialization name.
	MaterializationName types.String `tfsdk:"materialization_name" tf:"optional"`
	// The org id of the user. Unique within a cloud.
	OrgId types.Int64 `tfsdk:"org_id" tf:"optional"`
	// The id of the pipeline. Globally unique.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// The name of the pipeline. Not unique.
	PipelineName types.String `tfsdk:"pipeline_name" tf:"optional"`
	// The cloud region.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The id of the request that caused an update.
	RequestId types.String `tfsdk:"request_id" tf:"optional"`
	// The id of a (delta) table. Globally unique.
	TableId types.String `tfsdk:"table_id" tf:"optional"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId types.String `tfsdk:"uc_resource_id" tf:"optional"`
	// The id of an execution. Globally unique.
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type PipelineAccessControlResponse struct {
	// All permissions.
	AllPermissions []PipelinePermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale []PipelineClusterAutoscale `tfsdk:"autoscale" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes compute.AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes compute.AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf compute.ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes compute.GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts compute.InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label types.String `tfsdk:"label" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers" tf:""`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers" tf:""`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode types.String `tfsdk:"mode" tf:"optional"`
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	Kind types.String `tfsdk:"kind" tf:"optional"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path" tf:"optional"`
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	Error []ErrorDetail `tfsdk:"error" tf:"optional"`
	// The event type. Should always correspond to the details
	EventType types.String `tfsdk:"event_type" tf:"optional"`
	// A time-based, globally unique id.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The severity level of the event.
	Level types.String `tfsdk:"level" tf:"optional"`
	// Maturity level for event_type.
	MaturityLevel types.String `tfsdk:"maturity_level" tf:"optional"`
	// The display message associated with the event.
	Message types.String `tfsdk:"message" tf:"optional"`
	// Describes where the event originates from.
	Origin []Origin `tfsdk:"origin" tf:"optional"`
	// A sequencing object to identify and order events.
	Sequence []Sequencing `tfsdk:"sequence" tf:"optional"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File []FileLibrary `tfsdk:"file" tf:"optional"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed.
	Maven compute.MavenLibrary `tfsdk:"maven" tf:"optional"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook []NotebookLibrary `tfsdk:"notebook" tf:"optional"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

type PipelinePermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type PipelinePermissions struct {
	AccessControlList []PipelineAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type PipelinePermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type PipelinePermissionsRequest struct {
	AccessControlList []PipelineAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]types.String `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment []PipelineDeployment `tfsdk:"deployment" tf:"optional"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters []Filters `tfsdk:"filters" tf:"optional"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition []IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition" tf:"optional"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition []IngestionPipelineDefinition `tfsdk:"ingestion_definition" tf:"optional"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger []PipelineTrigger `tfsdk:"trigger" tf:"optional"`
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health" tf:"optional"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `tfsdk:"latest_updates" tf:"optional"`
	// The user-friendly name of the pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The unique identifier of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName types.String `tfsdk:"run_as_user_name" tf:"optional"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

type PipelineTrigger struct {
	Cron []CronTrigger `tfsdk:"cron" tf:"optional"`

	Manual []ManualTrigger `tfsdk:"manual" tf:"optional"`
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	DestinationCatalog types.String `tfsdk:"destination_catalog" tf:"optional"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema types.String `tfsdk:"destination_schema" tf:"optional"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog types.String `tfsdk:"source_catalog" tf:"optional"`
	// Required. Schema name in the source database.
	SourceSchema types.String `tfsdk:"source_schema" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration []TableSpecificConfig `tfsdk:"table_configuration" tf:"optional"`
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no" tf:"optional"`
	// the ID assigned by the data plane.
	DataPlaneId []DataPlaneId `tfsdk:"data_plane_id" tf:"optional"`
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name" tf:"optional"`
	// Exception message
	Message types.String `tfsdk:"message" tf:"optional"`
	// Stack trace consisting of a list of stack frames
	Stack []StackFrame `tfsdk:"stack" tf:"optional"`
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class" tf:"optional"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name" tf:"optional"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number" tf:"optional"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name" tf:"optional"`
}

type StartUpdate struct {
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []types.String `tfsdk:"full_refresh_selection" tf:"optional"`

	PipelineId types.String `tfsdk:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []types.String `tfsdk:"refresh_selection" tf:"optional"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

type StartUpdateResponse struct {
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

type StopPipelineResponse struct {
}

// Stop a pipeline
type StopRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog types.String `tfsdk:"destination_catalog" tf:"optional"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema" tf:"optional"`
	// Optional. Destination table name. The pipeline fails If a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable types.String `tfsdk:"destination_table" tf:"optional"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog types.String `tfsdk:"source_catalog" tf:"optional"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema types.String `tfsdk:"source_schema" tf:"optional"`
	// Required. Table name in the source database.
	SourceTable types.String `tfsdk:"source_table" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration []TableSpecificConfig `tfsdk:"table_configuration" tf:"optional"`
}

type TableSpecificConfig struct {
	// The primary key of the table used to apply changes.
	PrimaryKeys []types.String `tfsdk:"primary_keys" tf:"optional"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields types.Bool `tfsdk:"salesforce_include_formula_fields" tf:"optional"`
	// The SCD type to use to ingest the table.
	ScdType types.String `tfsdk:"scd_type" tf:"optional"`
}

type UpdateInfo struct {
	// What triggered this update.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config []PipelineSpec `tfsdk:"config" tf:"optional"`
	// The time when this update was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []types.String `tfsdk:"full_refresh_selection" tf:"optional"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []types.String `tfsdk:"refresh_selection" tf:"optional"`
	// The update state.
	State types.String `tfsdk:"state" tf:"optional"`
	// The ID of this update.
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

type UpdateStateInfo struct {
	CreationTime types.String `tfsdk:"creation_time" tf:"optional"`

	State types.String `tfsdk:"state" tf:"optional"`

	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}
