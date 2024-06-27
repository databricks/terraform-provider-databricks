// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration map[types.String]types.String `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`

	DryRun types.Bool `tfsdk:"dry_run"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `tfsdk:"filters"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition *IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *ManagedIngestionPipelineDefinition `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `tfsdk:"trigger"`
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings *PipelineSpec `tfsdk:"effective_settings"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`

	TimezoneId types.String `tfsdk:"timezone_id"`
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	Instance types.String `tfsdk:"instance"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no"`
}

// Delete a pipeline
type DeletePipelineRequest struct {
	PipelineId types.String `tfsdk:"-" url:"-"`
}

type DeletePipelineResponse struct {
}

// The deployment method that manages the pipeline: - BUNDLE: The pipeline is
// managed by a Databricks Asset Bundle.
type DeploymentKind string

const DeploymentKindBundle DeploymentKind = `BUNDLE`

// String representation for [fmt.Print]
func (f *DeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`:
		*f = DeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE"`, v)
	}
}

// Type always returns DeploymentKind to satisfy [pflag.Value] interface
func (f *DeploymentKind) Type() string {
	return "DeploymentKind"
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration map[types.String]types.String `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `tfsdk:"filters"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition *IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *ManagedIngestionPipelineDefinition `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Unique identifier for this pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" url:"-"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `tfsdk:"trigger"`
}

type EditPipelineResponse struct {
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions []SerializedException `tfsdk:"exceptions"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal"`
}

// The severity level of the event.
type EventLevel string

const EventLevelError EventLevel = `ERROR`

const EventLevelInfo EventLevel = `INFO`

const EventLevelMetrics EventLevel = `METRICS`

const EventLevelWarn EventLevel = `WARN`

// String representation for [fmt.Print]
func (f *EventLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventLevel) Set(v string) error {
	switch v {
	case `ERROR`, `INFO`, `METRICS`, `WARN`:
		*f = EventLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "INFO", "METRICS", "WARN"`, v)
	}
}

// Type always returns EventLevel to satisfy [pflag.Value] interface
func (f *EventLevel) Type() string {
	return "EventLevel"
}

type FileLibrary struct {
	// The absolute path of the file.
	Path types.String `tfsdk:"path"`
}

type Filters struct {
	// Paths to exclude.
	Exclude []types.String `tfsdk:"exclude"`
	// Paths to include.
	Include []types.String `tfsdk:"include"`
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-" url:"-"`
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PipelinePermissionsDescription `tfsdk:"permission_levels"`
}

// Get pipeline permissions
type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-" url:"-"`
}

// Get a pipeline
type GetPipelineRequest struct {
	PipelineId types.String `tfsdk:"-" url:"-"`
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause types.String `tfsdk:"cause"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// The health of a pipeline.
	Health GetPipelineResponseHealth `tfsdk:"health"`
	// The last time the pipeline settings were modified or created.
	LastModified types.Int64 `tfsdk:"last_modified"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `tfsdk:"latest_updates"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name types.String `tfsdk:"name"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *PipelineSpec `tfsdk:"spec"`
	// The pipeline state.
	State PipelineState `tfsdk:"state"`
}

// The health of a pipeline.
type GetPipelineResponseHealth string

const GetPipelineResponseHealthHealthy GetPipelineResponseHealth = `HEALTHY`

const GetPipelineResponseHealthUnhealthy GetPipelineResponseHealth = `UNHEALTHY`

// String representation for [fmt.Print]
func (f *GetPipelineResponseHealth) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetPipelineResponseHealth) Set(v string) error {
	switch v {
	case `HEALTHY`, `UNHEALTHY`:
		*f = GetPipelineResponseHealth(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HEALTHY", "UNHEALTHY"`, v)
	}
}

// Type always returns GetPipelineResponseHealth to satisfy [pflag.Value] interface
func (f *GetPipelineResponseHealth) Type() string {
	return "GetPipelineResponseHealth"
}

// Get a pipeline update
type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-" url:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-" url:"-"`
}

type GetUpdateResponse struct {
	// The current update info.
	Update *UpdateInfo `tfsdk:"update"`
}

type IngestionConfig struct {
	// Select tables from a specific source schema.
	Schema *SchemaSpec `tfsdk:"schema"`
	// Select tables from a specific source table.
	Table *TableSpec `tfsdk:"table"`
}

type IngestionGatewayPipelineDefinition struct {
	// Immutable. The Unity Catalog connection this gateway pipeline uses to
	// communicate with the source.
	ConnectionId types.String `tfsdk:"connection_id"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog types.String `tfsdk:"gateway_storage_catalog"`
	// Required. The Unity Catalog-compatible naming for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName types.String `tfsdk:"gateway_storage_name"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema types.String `tfsdk:"gateway_storage_schema"`
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
	Filter types.String `tfsdk:"-" url:"filter,omitempty"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []types.String `tfsdk:"-" url:"order_by,omitempty"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`

	PipelineId types.String `tfsdk:"-" url:"-"`
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	Events []PipelineEvent `tfsdk:"events"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
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
	Filter types.String `tfsdk:"-" url:"filter,omitempty"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy []types.String `tfsdk:"-" url:"order_by,omitempty"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of events matching the request criteria.
	Statuses []PipelineStateInfo `tfsdk:"statuses"`
}

// List pipeline updates
type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// The pipeline to return updates for.
	PipelineId types.String `tfsdk:"-" url:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId types.String `tfsdk:"-" url:"until_update_id,omitempty"`
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token"`

	Updates []UpdateInfo `tfsdk:"updates"`
}

type ManagedIngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection this ingestion pipeline uses to
	// communicate with the source. Specify either ingestion_gateway_id or
	// connection_name.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Immutable. Identifier for the ingestion gateway used by this ingestion
	// pipeline to communicate with the source. Specify either
	// ingestion_gateway_id or connection_name.
	IngestionGatewayId types.String `tfsdk:"ingestion_gateway_id"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects []IngestionConfig `tfsdk:"objects"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration *TableSpecificConfig `tfsdk:"table_configuration"`
}

type ManualTrigger struct {
}

// Maturity level for EventDetails.
type MaturityLevel string

const MaturityLevelDeprecated MaturityLevel = `DEPRECATED`

const MaturityLevelEvolving MaturityLevel = `EVOLVING`

const MaturityLevelStable MaturityLevel = `STABLE`

// String representation for [fmt.Print]
func (f *MaturityLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MaturityLevel) Set(v string) error {
	switch v {
	case `DEPRECATED`, `EVOLVING`, `STABLE`:
		*f = MaturityLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPRECATED", "EVOLVING", "STABLE"`, v)
	}
}

// Type always returns MaturityLevel to satisfy [pflag.Value] interface
func (f *MaturityLevel) Type() string {
	return "MaturityLevel"
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	Path types.String `tfsdk:"path"`
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	Alerts []types.String `tfsdk:"alerts"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients []types.String `tfsdk:"email_recipients"`
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId types.Int64 `tfsdk:"batch_id"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud types.String `tfsdk:"cloud"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName types.String `tfsdk:"dataset_name"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId types.String `tfsdk:"flow_id"`
	// The name of the flow. Not unique.
	FlowName types.String `tfsdk:"flow_name"`
	// The optional host name where the event was triggered
	Host types.String `tfsdk:"host"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId types.String `tfsdk:"maintenance_id"`
	// Materialization name.
	MaterializationName types.String `tfsdk:"materialization_name"`
	// The org id of the user. Unique within a cloud.
	OrgId types.Int64 `tfsdk:"org_id"`
	// The id of the pipeline. Globally unique.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// The name of the pipeline. Not unique.
	PipelineName types.String `tfsdk:"pipeline_name"`
	// The cloud region.
	Region types.String `tfsdk:"region"`
	// The id of the request that caused an update.
	RequestId types.String `tfsdk:"request_id"`
	// The id of a (delta) table. Globally unique.
	TableId types.String `tfsdk:"table_id"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId types.String `tfsdk:"uc_resource_id"`
	// The id of an execution. Globally unique.
	UpdateId types.String `tfsdk:"update_id"`
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type PipelineAccessControlResponse struct {
	// All permissions.
	AllPermissions []PipelinePermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *PipelineClusterAutoscale `tfsdk:"autoscale"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *compute.AwsAttributes `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *compute.AzureAttributes `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *compute.ClusterLogConf `tfsdk:"cluster_log_conf"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[types.String]types.String `tfsdk:"custom_tags"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *compute.GcpAttributes `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []compute.InitScriptInfo `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label types.String `tfsdk:"label"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf map[types.String]types.String `tfsdk:"spark_conf"`
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
	SparkEnvVars map[types.String]types.String `tfsdk:"spark_env_vars"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys"`
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers"`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers"`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode PipelineClusterAutoscaleMode `tfsdk:"mode"`
}

// Databricks Enhanced Autoscaling optimizes cluster utilization by
// automatically allocating cluster resources based on workload volume, with
// minimal impact to the data processing latency of your pipelines. Enhanced
// Autoscaling is available for `updates` clusters only. The legacy autoscaling
// feature is used for `maintenance` clusters.
type PipelineClusterAutoscaleMode string

const PipelineClusterAutoscaleModeEnhanced PipelineClusterAutoscaleMode = `ENHANCED`

const PipelineClusterAutoscaleModeLegacy PipelineClusterAutoscaleMode = `LEGACY`

// String representation for [fmt.Print]
func (f *PipelineClusterAutoscaleMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineClusterAutoscaleMode) Set(v string) error {
	switch v {
	case `ENHANCED`, `LEGACY`:
		*f = PipelineClusterAutoscaleMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ENHANCED", "LEGACY"`, v)
	}
}

// Type always returns PipelineClusterAutoscaleMode to satisfy [pflag.Value] interface
func (f *PipelineClusterAutoscaleMode) Type() string {
	return "PipelineClusterAutoscaleMode"
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	Kind DeploymentKind `tfsdk:"kind"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path"`
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	Error *ErrorDetail `tfsdk:"error"`
	// The event type. Should always correspond to the details
	EventType types.String `tfsdk:"event_type"`
	// A time-based, globally unique id.
	Id types.String `tfsdk:"id"`
	// The severity level of the event.
	Level EventLevel `tfsdk:"level"`
	// Maturity level for event_type.
	MaturityLevel MaturityLevel `tfsdk:"maturity_level"`
	// The display message associated with the event.
	Message types.String `tfsdk:"message"`
	// Describes where the event originates from.
	Origin *Origin `tfsdk:"origin"`
	// A sequencing object to identify and order events.
	Sequence *Sequencing `tfsdk:"sequence"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp"`
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File *FileLibrary `tfsdk:"file"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar"`
	// Specification of a maven library to be installed.
	Maven *compute.MavenLibrary `tfsdk:"maven"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook *NotebookLibrary `tfsdk:"notebook"`
}

type PipelinePermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `tfsdk:"permission_level"`
}

// Permission level
type PipelinePermissionLevel string

const PipelinePermissionLevelCanManage PipelinePermissionLevel = `CAN_MANAGE`

const PipelinePermissionLevelCanRun PipelinePermissionLevel = `CAN_RUN`

const PipelinePermissionLevelCanView PipelinePermissionLevel = `CAN_VIEW`

const PipelinePermissionLevelIsOwner PipelinePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *PipelinePermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelinePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = PipelinePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns PipelinePermissionLevel to satisfy [pflag.Value] interface
func (f *PipelinePermissionLevel) Type() string {
	return "PipelinePermissionLevel"
}

type PipelinePermissions struct {
	AccessControlList []PipelineAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type PipelinePermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `tfsdk:"permission_level"`
}

type PipelinePermissionsRequest struct {
	AccessControlList []PipelineAccessControlRequest `tfsdk:"access_control_list"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-" url:"-"`
}

type PipelineSpec struct {
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration map[types.String]types.String `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `tfsdk:"filters"`
	// The definition of a gateway pipeline to support CDC.
	GatewayDefinition *IngestionGatewayPipelineDefinition `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *ManagedIngestionPipelineDefinition `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `tfsdk:"trigger"`
}

// The pipeline state.
type PipelineState string

const PipelineStateDeleted PipelineState = `DELETED`

const PipelineStateDeploying PipelineState = `DEPLOYING`

const PipelineStateFailed PipelineState = `FAILED`

const PipelineStateIdle PipelineState = `IDLE`

const PipelineStateRecovering PipelineState = `RECOVERING`

const PipelineStateResetting PipelineState = `RESETTING`

const PipelineStateRunning PipelineState = `RUNNING`

const PipelineStateStarting PipelineState = `STARTING`

const PipelineStateStopping PipelineState = `STOPPING`

// String representation for [fmt.Print]
func (f *PipelineState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineState) Set(v string) error {
	switch v {
	case `DELETED`, `DEPLOYING`, `FAILED`, `IDLE`, `RECOVERING`, `RESETTING`, `RUNNING`, `STARTING`, `STOPPING`:
		*f = PipelineState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DEPLOYING", "FAILED", "IDLE", "RECOVERING", "RESETTING", "RUNNING", "STARTING", "STOPPING"`, v)
	}
}

// Type always returns PipelineState to satisfy [pflag.Value] interface
func (f *PipelineState) Type() string {
	return "PipelineState"
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `tfsdk:"latest_updates"`
	// The user-friendly name of the pipeline.
	Name types.String `tfsdk:"name"`
	// The unique identifier of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// The pipeline state.
	State PipelineState `tfsdk:"state"`
}

type PipelineTrigger struct {
	Cron *CronTrigger `tfsdk:"cron"`

	Manual *ManualTrigger `tfsdk:"manual"`
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	DestinationCatalog types.String `tfsdk:"destination_catalog"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema types.String `tfsdk:"destination_schema"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog types.String `tfsdk:"source_catalog"`
	// Required. Schema name in the source database.
	SourceSchema types.String `tfsdk:"source_schema"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the ManagedIngestionPipelineDefinition
	// object.
	TableConfiguration *TableSpecificConfig `tfsdk:"table_configuration"`
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no"`
	// the ID assigned by the data plane.
	DataPlaneId *DataPlaneId `tfsdk:"data_plane_id"`
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name"`
	// Exception message
	Message types.String `tfsdk:"message"`
	// Stack trace consisting of a list of stack frames
	Stack []StackFrame `tfsdk:"stack"`
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name"`
}

type StartUpdate struct {
	Cause StartUpdateCause `tfsdk:"cause"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []types.String `tfsdk:"full_refresh_selection"`

	PipelineId types.String `tfsdk:"-" url:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []types.String `tfsdk:"refresh_selection"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only"`
}

type StartUpdateCause string

const StartUpdateCauseApiCall StartUpdateCause = `API_CALL`

const StartUpdateCauseJobTask StartUpdateCause = `JOB_TASK`

const StartUpdateCauseRetryOnFailure StartUpdateCause = `RETRY_ON_FAILURE`

const StartUpdateCauseSchemaChange StartUpdateCause = `SCHEMA_CHANGE`

const StartUpdateCauseServiceUpgrade StartUpdateCause = `SERVICE_UPGRADE`

const StartUpdateCauseUserAction StartUpdateCause = `USER_ACTION`

// String representation for [fmt.Print]
func (f *StartUpdateCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StartUpdateCause) Set(v string) error {
	switch v {
	case `API_CALL`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = StartUpdateCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Type always returns StartUpdateCause to satisfy [pflag.Value] interface
func (f *StartUpdateCause) Type() string {
	return "StartUpdateCause"
}

type StartUpdateResponse struct {
	UpdateId types.String `tfsdk:"update_id"`
}

type StopPipelineResponse struct {
}

// Stop a pipeline
type StopRequest struct {
	PipelineId types.String `tfsdk:"-" url:"-"`
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog types.String `tfsdk:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema"`
	// Optional. Destination table name. The pipeline fails If a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable types.String `tfsdk:"destination_table"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog types.String `tfsdk:"source_catalog"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema types.String `tfsdk:"source_schema"`
	// Required. Table name in the source database.
	SourceTable types.String `tfsdk:"source_table"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// ManagedIngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration *TableSpecificConfig `tfsdk:"table_configuration"`
}

type TableSpecificConfig struct {
	// The primary key of the table used to apply changes.
	PrimaryKeys []types.String `tfsdk:"primary_keys"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields types.Bool `tfsdk:"salesforce_include_formula_fields"`
	// The SCD type to use to ingest the table.
	ScdType TableSpecificConfigScdType `tfsdk:"scd_type"`
}

// The SCD type to use to ingest the table.
type TableSpecificConfigScdType string

const TableSpecificConfigScdTypeScdType1 TableSpecificConfigScdType = `SCD_TYPE_1`

const TableSpecificConfigScdTypeScdType2 TableSpecificConfigScdType = `SCD_TYPE_2`

// String representation for [fmt.Print]
func (f *TableSpecificConfigScdType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableSpecificConfigScdType) Set(v string) error {
	switch v {
	case `SCD_TYPE_1`, `SCD_TYPE_2`:
		*f = TableSpecificConfigScdType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SCD_TYPE_1", "SCD_TYPE_2"`, v)
	}
}

// Type always returns TableSpecificConfigScdType to satisfy [pflag.Value] interface
func (f *TableSpecificConfigScdType) Type() string {
	return "TableSpecificConfigScdType"
}

type UpdateInfo struct {
	// What triggered this update.
	Cause UpdateInfoCause `tfsdk:"cause"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *PipelineSpec `tfsdk:"config"`
	// The time when this update was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []types.String `tfsdk:"full_refresh_selection"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []types.String `tfsdk:"refresh_selection"`
	// The update state.
	State UpdateInfoState `tfsdk:"state"`
	// The ID of this update.
	UpdateId types.String `tfsdk:"update_id"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only"`
}

// What triggered this update.
type UpdateInfoCause string

const UpdateInfoCauseApiCall UpdateInfoCause = `API_CALL`

const UpdateInfoCauseJobTask UpdateInfoCause = `JOB_TASK`

const UpdateInfoCauseRetryOnFailure UpdateInfoCause = `RETRY_ON_FAILURE`

const UpdateInfoCauseSchemaChange UpdateInfoCause = `SCHEMA_CHANGE`

const UpdateInfoCauseServiceUpgrade UpdateInfoCause = `SERVICE_UPGRADE`

const UpdateInfoCauseUserAction UpdateInfoCause = `USER_ACTION`

// String representation for [fmt.Print]
func (f *UpdateInfoCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateInfoCause) Set(v string) error {
	switch v {
	case `API_CALL`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = UpdateInfoCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Type always returns UpdateInfoCause to satisfy [pflag.Value] interface
func (f *UpdateInfoCause) Type() string {
	return "UpdateInfoCause"
}

// The update state.
type UpdateInfoState string

const UpdateInfoStateCanceled UpdateInfoState = `CANCELED`

const UpdateInfoStateCompleted UpdateInfoState = `COMPLETED`

const UpdateInfoStateCreated UpdateInfoState = `CREATED`

const UpdateInfoStateFailed UpdateInfoState = `FAILED`

const UpdateInfoStateInitializing UpdateInfoState = `INITIALIZING`

const UpdateInfoStateQueued UpdateInfoState = `QUEUED`

const UpdateInfoStateResetting UpdateInfoState = `RESETTING`

const UpdateInfoStateRunning UpdateInfoState = `RUNNING`

const UpdateInfoStateSettingUpTables UpdateInfoState = `SETTING_UP_TABLES`

const UpdateInfoStateStopping UpdateInfoState = `STOPPING`

const UpdateInfoStateWaitingForResources UpdateInfoState = `WAITING_FOR_RESOURCES`

// String representation for [fmt.Print]
func (f *UpdateInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPLETED`, `CREATED`, `FAILED`, `INITIALIZING`, `QUEUED`, `RESETTING`, `RUNNING`, `SETTING_UP_TABLES`, `STOPPING`, `WAITING_FOR_RESOURCES`:
		*f = UpdateInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPLETED", "CREATED", "FAILED", "INITIALIZING", "QUEUED", "RESETTING", "RUNNING", "SETTING_UP_TABLES", "STOPPING", "WAITING_FOR_RESOURCES"`, v)
	}
}

// Type always returns UpdateInfoState to satisfy [pflag.Value] interface
func (f *UpdateInfoState) Type() string {
	return "UpdateInfoState"
}

type UpdateStateInfo struct {
	CreationTime types.String `tfsdk:"creation_time"`

	State UpdateStateInfoState `tfsdk:"state"`

	UpdateId types.String `tfsdk:"update_id"`
}

type UpdateStateInfoState string

const UpdateStateInfoStateCanceled UpdateStateInfoState = `CANCELED`

const UpdateStateInfoStateCompleted UpdateStateInfoState = `COMPLETED`

const UpdateStateInfoStateCreated UpdateStateInfoState = `CREATED`

const UpdateStateInfoStateFailed UpdateStateInfoState = `FAILED`

const UpdateStateInfoStateInitializing UpdateStateInfoState = `INITIALIZING`

const UpdateStateInfoStateQueued UpdateStateInfoState = `QUEUED`

const UpdateStateInfoStateResetting UpdateStateInfoState = `RESETTING`

const UpdateStateInfoStateRunning UpdateStateInfoState = `RUNNING`

const UpdateStateInfoStateSettingUpTables UpdateStateInfoState = `SETTING_UP_TABLES`

const UpdateStateInfoStateStopping UpdateStateInfoState = `STOPPING`

const UpdateStateInfoStateWaitingForResources UpdateStateInfoState = `WAITING_FOR_RESOURCES`

// String representation for [fmt.Print]
func (f *UpdateStateInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateStateInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPLETED`, `CREATED`, `FAILED`, `INITIALIZING`, `QUEUED`, `RESETTING`, `RUNNING`, `SETTING_UP_TABLES`, `STOPPING`, `WAITING_FOR_RESOURCES`:
		*f = UpdateStateInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPLETED", "CREATED", "FAILED", "INITIALIZING", "QUEUED", "RESETTING", "RUNNING", "SETTING_UP_TABLES", "STOPPING", "WAITING_FOR_RESOURCES"`, v)
	}
}

// Type always returns UpdateStateInfoState to satisfy [pflag.Value] interface
func (f *UpdateStateInfoState) Type() string {
	return "UpdateStateInfoState"
}
