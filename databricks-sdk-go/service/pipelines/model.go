// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

// Top-level configuration for API Source connectors with arbitrary
// configuration.
type ApiSourceConnectorConfig struct {
	// Arbitrary key-value configuration values for the API Source connector.
	Configs map[string]string `json:"configs,omitempty"`
}

func (s *ApiSourceConnectorConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Options for API Source connectors with arbitrary configuration.
type ApiSourceConnectorOptions struct {
	// Arbitrary key-value configuration options for the API Source connector.
	Options map[string]string `json:"options,omitempty"`
}

func (s *ApiSourceConnectorOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ApplyEnvironmentRequest struct {
	// The ID of the pipeline to apply the environment to.
	PipelineId string `json:"-" url:"-"`
}

func (s *ApplyEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ApplyEnvironmentRequestResponse struct {
}

func (s *ApplyEnvironmentRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Policy for auto full refresh.
type AutoFullRefreshPolicy struct {
	// (Required, Mutable) Whether to enable auto full refresh or not.
	Enabled bool `json:"enabled"`
	// (Optional, Mutable) Specify the minimum interval in hours between the
	// timestamp at which a table was last full refreshed and the current
	// timestamp for triggering auto full If unspecified and autoFullRefresh is
	// enabled then by default min_interval_hours is 24 hours.
	MinIntervalHours int `json:"min_interval_hours,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutoFullRefreshPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoFullRefreshPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AvroTransformerOptions struct {
	// (Optional) Parse mode for Avro data. Valid values: FAILFAST, PERMISSIVE.
	// Defaults to PERMISSIVE.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Inline Avro JSON schema string.
	Schema string `json:"schema,omitempty"`
	// Path to a schema file (.avsc).
	SchemaFilePath string `json:"schema_file_path,omitempty"`
	// (Optional) Schema registry to resolve the Avro schema at runtime instead
	// of providing it inline or via a file path.
	SchemaRegistry *SchemaRegistryConfig `json:"schema_registry,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AvroTransformerOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AvroTransformerOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Enum to specify which mode of clone to execute
type CloneMode string

const CloneModeMigrateToUc CloneMode = `MIGRATE_TO_UC`

// String representation for [fmt.Print]
func (f *CloneMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CloneMode) Set(v string) error {
	switch v {
	case `MIGRATE_TO_UC`:
		*f = CloneMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MIGRATE_TO_UC"`, v)
	}
}

// Values returns all possible values for CloneMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CloneMode) Values() []CloneMode {
	return []CloneMode{
		CloneModeMigrateToUc,
	}
}

// Type always returns CloneMode to satisfy [pflag.Value] interface
func (f *CloneMode) Type() string {
	return "CloneMode"
}

type ClonePipelineRequest struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// ID (UUID) of the serverless budget policy applied to this pipeline. Only
	// supported for serverless pipelines.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. Tables
	// are published to the schema set in `schema` (or the legacy `target`)
	// inside this catalog, for example `catalog`.`schema`.`table`; source code
	// can also use fully qualified names to write to other schemas and
	// catalogs. If `catalog` is not specified, the pipeline uses the legacy
	// Hive metastore.
	Catalog string `json:"catalog,omitempty"`
	// Release channel that selects the pipeline runtime version: `CURRENT` (the
	// default) or `PREVIEW`, which runs on the upcoming runtime release.
	Channel string `json:"channel,omitempty"`
	// The type of clone to perform. Currently, only deep copies are supported
	CloneMode CloneMode `json:"clone_mode,omitempty"`
	// Cluster settings for this pipeline deployment. Applies to classic
	// (non-serverless) pipelines. Omit this field when `serverless` is `true`.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	//
	// Deprecated: wrap the pipeline in a continuous job instead, which also
	// lets you take advantage of job-level settings such as performance mode.
	// When the pipeline is started by a continuous job, the job's setting takes
	// precedence and this field is ignored.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	//
	// Deprecated: set development mode for each update instead.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition: `CORE`, `PRO`, or `ADVANCED`. Serverless
	// pipelines require `ADVANCED`.
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// clone. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Path of the pipeline parent folder in workspace file tree.
	//
	// If absent, the pipeline doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Source pipeline to clone from
	PipelineId string `json:"-" url:"-"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath string `json:"root_path,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline. Serverless is
	// the recommended compute for new pipelines; set this to `true` to run the
	// pipeline on serverless. For notebook/file pipelines, omitting this field
	// uses classic compute configured through the `clusters` field. When
	// `true`, omit `clusters`; Photon is always enabled.
	Serverless bool `json:"serverless,omitempty"`
	// Serverless compute ID specified by the user for serverless pipelines. Use
	// compute_policy_id instead. Only one of these fields may be set.
	ServerlessComputeId string `json:"serverless_compute_id,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Test configuration for this pipeline. Used only by test-only updates.
	Testing *PipelineTesting `json:"testing,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
	// Usage policy of this pipeline.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClonePipelineRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClonePipelineRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClonePipelineResponse struct {
	// The pipeline id of the cloned pipeline
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClonePipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClonePipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Confluence specific options for ingestion
type ConfluenceConnectorOptions struct {
	// (Optional) Spaces to filter Confluence data on
	IncludeConfluenceSpaces []string `json:"include_confluence_spaces,omitempty"`
}

func (s *ConfluenceConnectorOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ConfluentSchemaRegistryOptions struct {
	// Required: subject name to resolve in the registry.
	Subject string `json:"subject,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConfluentSchemaRegistryOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConfluentSchemaRegistryOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ConnectionParameters struct {
	// Source catalog for initial connection. This is necessary for schema
	// exploration in some database systems like Oracle, and optional but
	// nice-to-have in some other database systems like Postgres. For Oracle
	// databases, this maps to a service name.
	SourceCatalog string `json:"source_catalog,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConnectionParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionParameters) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Wrapper message for source-specific options to support multiple connector
// types
type ConnectorOptions struct {
	// Connector-specific options for API Source connectors.
	ApiSourceConnectorOptions *ApiSourceConnectorOptions `json:"api_source_connector_options,omitempty"`

	ConfluenceOptions *ConfluenceConnectorOptions `json:"confluence_options,omitempty"`

	GdriveOptions *GoogleDriveOptions `json:"gdrive_options,omitempty"`

	GoogleAdsOptions *GoogleAdsOptions `json:"google_ads_options,omitempty"`

	JiraOptions *JiraConnectorOptions `json:"jira_options,omitempty"`

	KafkaOptions *KafkaOptions `json:"kafka_options,omitempty"`

	LinkedinAdsOptions *LinkedInAdsOptions `json:"linkedin_ads_options,omitempty"`

	MarketoOptions *MarketoOptions `json:"marketo_options,omitempty"`

	MetaAdsOptions *MetaMarketingOptions `json:"meta_ads_options,omitempty"`

	OnedriveOptions *OneDriveOptions `json:"onedrive_options,omitempty"`

	OutlookOptions *OutlookOptions `json:"outlook_options,omitempty"`

	RabbitmqOptions *RabbitmqOptions `json:"rabbitmq_options,omitempty"`

	RedditAdsOptions *RedditAdsOptions `json:"reddit_ads_options,omitempty"`

	SharepointOptions *SharepointOptions `json:"sharepoint_options,omitempty"`

	SmartsheetOptions *SmartsheetOptions `json:"smartsheet_options,omitempty"`

	TiktokAdsOptions *TikTokAdsOptions `json:"tiktok_ads_options,omitempty"`

	ZendeskSupportOptions *ZendeskSupportOptions `json:"zendesk_support_options,omitempty"`
}

func (s *ConnectorOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// For certain database sources LakeFlow Connect offers both query based and cdc
// ingestion, ConnectorType can bse used to convey the type of ingestion. If
// connection_name is provided for database sources, we default to Query Based
// ingestion
type ConnectorType string

const ConnectorTypeCdc ConnectorType = `CDC`

const ConnectorTypeQueryBased ConnectorType = `QUERY_BASED`

// String representation for [fmt.Print]
func (f *ConnectorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectorType) Set(v string) error {
	switch v {
	case `CDC`, `QUERY_BASED`:
		*f = ConnectorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CDC", "QUERY_BASED"`, v)
	}
}

// Values returns all possible values for ConnectorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConnectorType) Values() []ConnectorType {
	return []ConnectorType{
		ConnectorTypeCdc,
		ConnectorTypeQueryBased,
	}
}

// Type always returns ConnectorType to satisfy [pflag.Value] interface
func (f *ConnectorType) Type() string {
	return "ConnectorType"
}

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	//
	// Ignored when `parent_path` is set. A pipeline in a workspace folder must
	// have a name that is unique within that folder, and creation fails if the
	// folder already contains an object with the same name.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// ID (UUID) of the serverless budget policy applied to this pipeline. Only
	// supported for serverless pipelines.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. Tables
	// are published to the schema set in `schema` (or the legacy `target`)
	// inside this catalog, for example `catalog`.`schema`.`table`; source code
	// can also use fully qualified names to write to other schemas and
	// catalogs. If `catalog` is not specified, the pipeline uses the legacy
	// Hive metastore.
	Catalog string `json:"catalog,omitempty"`
	// Release channel that selects the pipeline runtime version: `CURRENT` (the
	// default) or `PREVIEW`, which runs on the upcoming runtime release.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment. Applies to classic
	// (non-serverless) pipelines. Omit this field when `serverless` is `true`.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	//
	// Deprecated: wrap the pipeline in a continuous job instead, which also
	// lets you take advantage of job-level settings such as performance mode.
	// When the pipeline is started by a continuous job, the job's setting takes
	// precedence and this field is ignored.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	//
	// Deprecated: set development mode for each update instead.
	Development bool `json:"development,omitempty"`
	// If true, no pipeline is created. The response contains the settings that
	// would be applied, with server defaults filled in, instead of a pipeline
	// ID.
	DryRun bool `json:"dry_run,omitempty"`
	// Pipeline product edition: `CORE`, `PRO`, or `ADVANCED`. Serverless
	// pipelines require `ADVANCED`.
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Key/value map of default parameters to use for pipeline execution.
	// Maximum total size: 10k characters (JSON format)
	Parameters map[string]string `json:"parameters,omitempty"`
	// Path of the pipeline parent folder in workspace file tree.
	//
	// If absent, the pipeline doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath string `json:"root_path,omitempty"`
	// The user or service principal that the pipeline runs as. If not
	// specified, the pipeline runs as the user who created the pipeline.
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline. Serverless is
	// the recommended compute for new pipelines; set this to `true` to run the
	// pipeline on serverless. For notebook/file pipelines, omitting this field
	// uses classic compute configured through the `clusters` field. When
	// `true`, omit `clusters`; Photon is always enabled.
	Serverless bool `json:"serverless,omitempty"`
	// Serverless compute ID specified by the user for serverless pipelines. Use
	// compute_policy_id instead. Only one of these fields may be set.
	ServerlessComputeId string `json:"serverless_compute_id,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Test configuration for this pipeline. Used only by test-only updates.
	Testing *PipelineTesting `json:"testing,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
	// Usage policy of this pipeline.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings *PipelineSpec `json:"effective_settings,omitempty"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CronTrigger struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`

	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CronTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CronTrigger) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	Instance string `json:"instance,omitempty"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo int64 `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataPlaneId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Location of staged data storage
type DataStagingOptions struct {
	// (Required, Immutable) The name of the catalog for the connector's staging
	// storage location.
	CatalogName string `json:"catalog_name"`
	// (Required, Immutable) The name of the schema for the connector's staging
	// storage location.
	SchemaName string `json:"schema_name"`
	// (Optional) The Unity Catalog-compatible name for the storage location.
	// This is the volume to use for the data that is extracted by the
	// connector. Spark Declarative Pipelines system will automatically create
	// the volume under the catalog and schema. For Combined Cdc Managed
	// Ingestion pipelines default name for the volume would be :
	// __databricks_ingestion_gateway_staging_data-$pipelineId
	VolumeName string `json:"volume_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataStagingOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataStagingOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Days of week in which the window is allowed to happen. If not specified all
// days of the week will be used.
type DayOfWeek string

const DayOfWeekFriday DayOfWeek = `FRIDAY`

const DayOfWeekMonday DayOfWeek = `MONDAY`

const DayOfWeekSaturday DayOfWeek = `SATURDAY`

const DayOfWeekSunday DayOfWeek = `SUNDAY`

const DayOfWeekThursday DayOfWeek = `THURSDAY`

const DayOfWeekTuesday DayOfWeek = `TUESDAY`

const DayOfWeekWednesday DayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *DayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = DayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Values returns all possible values for DayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// Type always returns DayOfWeek to satisfy [pflag.Value] interface
func (f *DayOfWeek) Type() string {
	return "DayOfWeek"
}

type DeletePipelineRequest struct {
	// If false, pipeline deletion will not cascade to its datasets (MVs, STs,
	// Views). By default, this parameter will be true and all tables will be
	// deleted with the pipeline.
	Cascade bool `json:"-" url:"cascade,omitempty"`
	// If true, deletion will proceed even if resource cleanup fails. By
	// default, deletion will fail if resources cleanup is required but fails.
	Force bool `json:"-" url:"force,omitempty"`
	// The ID of the pipeline to delete.
	PipelineId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeletePipelineRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePipelineRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for DeploymentKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeploymentKind) Values() []DeploymentKind {
	return []DeploymentKind{
		DeploymentKindBundle,
	}
}

// Type always returns DeploymentKind to satisfy [pflag.Value] interface
func (f *DeploymentKind) Type() string {
	return "DeploymentKind"
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts with the
	// name of another pipeline.
	//
	// Ignored when the pipeline is in a workspace folder or the request sets
	// `parent_path`. Renaming or moving such a pipeline fails if the target
	// folder already contains an object with the same name.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// ID (UUID) of the serverless budget policy applied to this pipeline. Only
	// supported for serverless pipelines.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. Tables
	// are published to the schema set in `schema` (or the legacy `target`)
	// inside this catalog, for example `catalog`.`schema`.`table`; source code
	// can also use fully qualified names to write to other schemas and
	// catalogs. If `catalog` is not specified, the pipeline uses the legacy
	// Hive metastore.
	Catalog string `json:"catalog,omitempty"`
	// Release channel that selects the pipeline runtime version: `CURRENT` (the
	// default) or `PREVIEW`, which runs on the upcoming runtime release.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment. Applies to classic
	// (non-serverless) pipelines. Omit this field when `serverless` is `true`.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	//
	// Deprecated: wrap the pipeline in a continuous job instead, which also
	// lets you take advantage of job-level settings such as performance mode.
	// When the pipeline is started by a continuous job, the job's setting takes
	// precedence and this field is ignored.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	//
	// Deprecated: set development mode for each update instead.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition: `CORE`, `PRO`, or `ADVANCED`. Serverless
	// pipelines require `ADVANCED`.
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Key/value map of default parameters to use for pipeline execution.
	// Maximum total size: 10k characters (JSON format)
	Parameters map[string]string `json:"parameters,omitempty"`
	// Path of the pipeline parent folder in workspace file tree.
	//
	// If absent, the pipeline doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Unique identifier for this pipeline.
	PipelineId string `json:"-" url:"-"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath string `json:"root_path,omitempty"`
	// The user or service principal that the pipeline runs as.
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline. Serverless is
	// the recommended compute for new pipelines; set this to `true` to run the
	// pipeline on serverless. For notebook/file pipelines, omitting this field
	// uses classic compute configured through the `clusters` field. When
	// `true`, omit `clusters`; Photon is always enabled.
	Serverless bool `json:"serverless,omitempty"`
	// Serverless compute ID specified by the user for serverless pipelines. Use
	// compute_policy_id instead. Only one of these fields may be set.
	ServerlessComputeId string `json:"serverless_compute_id,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Test configuration for this pipeline. Used only by test-only updates.
	Testing *PipelineTesting `json:"testing,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
	// Usage policy of this pipeline.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditPipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditPipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions []SerializedException `json:"exceptions,omitempty"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal bool `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ErrorDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ErrorDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for EventLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *EventLevel) Values() []EventLevel {
	return []EventLevel{
		EventLevelError,
		EventLevelInfo,
		EventLevelMetrics,
		EventLevelWarn,
	}
}

// Type always returns EventLevel to satisfy [pflag.Value] interface
func (f *EventLevel) Type() string {
	return "EventLevel"
}

// Configurable event log parameters.
type EventLogSpec struct {
	// The UC catalog the event log is published under.
	Catalog string `json:"catalog,omitempty"`
	// The name the event log is published to in UC.
	Name string `json:"name,omitempty"`
	// The UC schema the event log is published under.
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EventLogSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EventLogSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileFilter struct {
	// Include files with modification times occurring after the specified time.
	// Timestamp format: YYYY-MM-DDTHH:mm:ss (e.g. 2020-06-01T13:00:00) Based on
	// https://spark.apache.org/docs/latest/sql-data-sources-generic-options.html#modification-time-path-filters
	ModifiedAfter string `json:"modified_after,omitempty"`
	// Include files with modification times occurring before the specified
	// time. Timestamp format: YYYY-MM-DDTHH:mm:ss (e.g. 2020-06-01T13:00:00)
	// Based on
	// https://spark.apache.org/docs/latest/sql-data-sources-generic-options.html#modification-time-path-filters
	ModifiedBefore string `json:"modified_before,omitempty"`
	// Include files with file names matching the pattern Based on
	// https://spark.apache.org/docs/latest/sql-data-sources-generic-options.html#path-glob-filter
	PathFilter string `json:"path_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileIngestionOptions struct {
	CorruptRecordColumn string `json:"corrupt_record_column,omitempty"`
	// Generic options
	FileFilters []FileFilter `json:"file_filters,omitempty"`
	// required for TableSpec
	Format FileIngestionOptionsFileFormat `json:"format,omitempty"`
	// Format-specific options Based on
	// https://docs.databricks.com/aws/en/ingestion/cloud-object-storage/auto-loader/options#file-format-options
	FormatOptions map[string]string `json:"format_options,omitempty"`

	IgnoreCorruptFiles bool `json:"ignore_corrupt_files,omitempty"`

	InferColumnTypes bool `json:"infer_column_types,omitempty"`
	// Column name case sensitivity
	// https://docs.databricks.com/aws/en/ingestion/cloud-object-storage/auto-loader/schema#change-case-sensitive-behavior
	ReaderCaseSensitive bool `json:"reader_case_sensitive,omitempty"`

	RescuedDataColumn string `json:"rescued_data_column,omitempty"`

	SchemaEvolutionMode FileIngestionOptionsSchemaEvolutionMode `json:"schema_evolution_mode,omitempty"`
	// Override inferred schema of specific columns Based on
	// https://docs.databricks.com/aws/en/ingestion/cloud-object-storage/auto-loader/schema#override-schema-inference-with-schema-hints
	SchemaHints string `json:"schema_hints,omitempty"`

	SingleVariantColumn string `json:"single_variant_column,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileIngestionOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileIngestionOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileIngestionOptionsFileFormat string

const FileIngestionOptionsFileFormatAvro FileIngestionOptionsFileFormat = `AVRO`

const FileIngestionOptionsFileFormatBinaryfile FileIngestionOptionsFileFormat = `BINARYFILE`

const FileIngestionOptionsFileFormatCsv FileIngestionOptionsFileFormat = `CSV`

const FileIngestionOptionsFileFormatExcel FileIngestionOptionsFileFormat = `EXCEL`

const FileIngestionOptionsFileFormatJson FileIngestionOptionsFileFormat = `JSON`

const FileIngestionOptionsFileFormatOrc FileIngestionOptionsFileFormat = `ORC`

const FileIngestionOptionsFileFormatParquet FileIngestionOptionsFileFormat = `PARQUET`

const FileIngestionOptionsFileFormatXml FileIngestionOptionsFileFormat = `XML`

// String representation for [fmt.Print]
func (f *FileIngestionOptionsFileFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileIngestionOptionsFileFormat) Set(v string) error {
	switch v {
	case `AVRO`, `BINARYFILE`, `CSV`, `EXCEL`, `JSON`, `ORC`, `PARQUET`, `XML`:
		*f = FileIngestionOptionsFileFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "BINARYFILE", "CSV", "EXCEL", "JSON", "ORC", "PARQUET", "XML"`, v)
	}
}

// Values returns all possible values for FileIngestionOptionsFileFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileIngestionOptionsFileFormat) Values() []FileIngestionOptionsFileFormat {
	return []FileIngestionOptionsFileFormat{
		FileIngestionOptionsFileFormatAvro,
		FileIngestionOptionsFileFormatBinaryfile,
		FileIngestionOptionsFileFormatCsv,
		FileIngestionOptionsFileFormatExcel,
		FileIngestionOptionsFileFormatJson,
		FileIngestionOptionsFileFormatOrc,
		FileIngestionOptionsFileFormatParquet,
		FileIngestionOptionsFileFormatXml,
	}
}

// Type always returns FileIngestionOptionsFileFormat to satisfy [pflag.Value] interface
func (f *FileIngestionOptionsFileFormat) Type() string {
	return "FileIngestionOptionsFileFormat"
}

// Based on
// https://docs.databricks.com/aws/en/ingestion/cloud-object-storage/auto-loader/schema#how-does-auto-loader-schema-evolution-work
type FileIngestionOptionsSchemaEvolutionMode string

const FileIngestionOptionsSchemaEvolutionModeAddNewColumns FileIngestionOptionsSchemaEvolutionMode = `ADD_NEW_COLUMNS`

const FileIngestionOptionsSchemaEvolutionModeAddNewColumnsWithTypeWidening FileIngestionOptionsSchemaEvolutionMode = `ADD_NEW_COLUMNS_WITH_TYPE_WIDENING`

const FileIngestionOptionsSchemaEvolutionModeFailOnNewColumns FileIngestionOptionsSchemaEvolutionMode = `FAIL_ON_NEW_COLUMNS`

const FileIngestionOptionsSchemaEvolutionModeNone FileIngestionOptionsSchemaEvolutionMode = `NONE`

const FileIngestionOptionsSchemaEvolutionModeRescue FileIngestionOptionsSchemaEvolutionMode = `RESCUE`

// String representation for [fmt.Print]
func (f *FileIngestionOptionsSchemaEvolutionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileIngestionOptionsSchemaEvolutionMode) Set(v string) error {
	switch v {
	case `ADD_NEW_COLUMNS`, `ADD_NEW_COLUMNS_WITH_TYPE_WIDENING`, `FAIL_ON_NEW_COLUMNS`, `NONE`, `RESCUE`:
		*f = FileIngestionOptionsSchemaEvolutionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD_NEW_COLUMNS", "ADD_NEW_COLUMNS_WITH_TYPE_WIDENING", "FAIL_ON_NEW_COLUMNS", "NONE", "RESCUE"`, v)
	}
}

// Values returns all possible values for FileIngestionOptionsSchemaEvolutionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileIngestionOptionsSchemaEvolutionMode) Values() []FileIngestionOptionsSchemaEvolutionMode {
	return []FileIngestionOptionsSchemaEvolutionMode{
		FileIngestionOptionsSchemaEvolutionModeAddNewColumns,
		FileIngestionOptionsSchemaEvolutionModeAddNewColumnsWithTypeWidening,
		FileIngestionOptionsSchemaEvolutionModeFailOnNewColumns,
		FileIngestionOptionsSchemaEvolutionModeNone,
		FileIngestionOptionsSchemaEvolutionModeRescue,
	}
}

// Type always returns FileIngestionOptionsSchemaEvolutionMode to satisfy [pflag.Value] interface
func (f *FileIngestionOptionsSchemaEvolutionMode) Type() string {
	return "FileIngestionOptionsSchemaEvolutionMode"
}

type FileLibrary struct {
	// The absolute path of the source code.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Filters struct {
	// Paths to exclude.
	Exclude []string `json:"exclude,omitempty"`
	// Paths to include.
	Include []string `json:"include,omitempty"`
}

func (s *Filters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func (s *GetPipelinePermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PipelinePermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetPipelinePermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func (s *GetPipelinePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPipelineRequest struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" url:"-"`
}

func (s *GetPipelineRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Publishing mode of the pipeline
	EffectivePublishingMode PublishingMode `json:"effective_publishing_mode,omitempty"`
	// Serverless compute ID resolved for the pipeline. For the equivalent
	// compute policy ID, use effective_compute_policy_id.
	EffectiveServerlessComputeId string `json:"effective_serverless_compute_id,omitempty"`
	// The health of a pipeline.
	Health GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name string `json:"name,omitempty"`
	// Key/value map of default parameters to use for pipeline execution.
	// Maximum total size: 10k characters (JSON format)
	Parameters map[string]string `json:"parameters,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The user or service principal that the pipeline runs as, if specified in
	// the request. This field indicates the explicit configuration of `run_as`
	// for the pipeline. To find the value in all cases, explicit or implicit,
	// use `run_as_user_name`.
	RunAs *RunAs `json:"run_as,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *PipelineSpec `json:"spec,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for GetPipelineResponseHealth.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetPipelineResponseHealth) Values() []GetPipelineResponseHealth {
	return []GetPipelineResponseHealth{
		GetPipelineResponseHealthHealthy,
		GetPipelineResponseHealthUnhealthy,
	}
}

// Type always returns GetPipelineResponseHealth to satisfy [pflag.Value] interface
func (f *GetPipelineResponseHealth) Type() string {
	return "GetPipelineResponseHealth"
}

type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" url:"-"`
	// The ID of the update.
	UpdateId string `json:"-" url:"-"`
}

func (s *GetUpdateRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetUpdateResponse struct {
	// The current update info.
	Update *UpdateInfo `json:"update,omitempty"`
}

func (s *GetUpdateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GoogleAdsConfig struct {
	// (Required) Manager Account ID (also called MCC Account ID) used to list
	// and access customer accounts under this manager account. This is required
	// for fetching the list of customer accounts during source selection. If
	// the same field is also set in the object-level GoogleAdsOptions
	// (connector_options), the object-level value takes precedence over this
	// top-level config.
	ManagerAccountId string `json:"manager_account_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GoogleAdsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GoogleAdsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User-defined custom report for the Google Ads connector. Mirrors the resource
// + fields + segments + metrics model that Google Ads GAQL exposes. The
// customer account this report runs against is supplied by the source schema
// (namespace), not by this message. The whole message is gated by the parent
// GoogleAdsOptions.custom_report_options stage; per-field stage annotations are
// intentionally omitted. Only supported on table-type objects: a custom report
// requires a destination table, so it cannot be specified at the schema/source
// level.
type GoogleAdsCustomReportOptions struct {
	// (Optional) Metric fields to select (e.g. "metrics.clicks",
	// "metrics.cost_micros"). Multiple values are joined into the GAQL SELECT
	// clause.
	Metrics []string `json:"metrics,omitempty"`
	// (Required) Google Ads resource to query (e.g. "ad_group_ad",
	// "keyword_view", "search_term_view"). Must be a resource that has metrics.
	// Values are validated against Google Ads' field-service catalog at
	// pipeline plan time.
	Resource string `json:"resource"`
	// (Optional) Resource fields to select, in fully-qualified GAQL form (e.g.
	// "ad_group_ad.ad.id", "ad_group_ad.status"). Multiple values are joined
	// into the GAQL SELECT clause.
	ResourceFields []string `json:"resource_fields,omitempty"`
	// (Optional) Segment fields to select (e.g. "segments.date",
	// "segments.device"). Must include at least one of segments.date,
	// segments.week, or segments.month — that segment is used as the
	// incremental cursor for the table.
	Segments []string `json:"segments,omitempty"`
}

func (s *GoogleAdsCustomReportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Google Ads specific options for ingestion (object-level). When set, these
// values override the corresponding fields in GoogleAdsConfig
// (source_configurations).
type GoogleAdsOptions struct {
	// (Optional) Custom report definition. When set, the table is treated as a
	// user-defined Google Ads custom report: the connector synthesizes a GAQL
	// query from the resource, fields, segments, and metrics specified here.
	// When unset, the table must match one of the connector's prebuilt sources.
	CustomReportOptions *GoogleAdsCustomReportOptions `json:"custom_report_options,omitempty"`
	// (Optional) Number of days to look back for report tables to capture
	// late-arriving data. If not specified, defaults to 30 days.
	LookbackWindowDays int `json:"lookback_window_days,omitempty"`
	// (Optional at this level) Manager Account ID (also called MCC Account ID)
	// used to list and access customer accounts under this manager account.
	// Overrides GoogleAdsConfig.manager_account_id from source_configurations
	// when set.
	ManagerAccountId string `json:"manager_account_id"`
	// (Optional) Start date for the initial sync of report tables in YYYY-MM-DD
	// format. This determines the earliest date from which to sync historical
	// data. If not specified, defaults to 2 years of historical data.
	SyncStartDate string `json:"sync_start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GoogleAdsOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GoogleAdsOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GoogleDriveOptions struct {
	EntityType GoogleDriveOptionsGoogleDriveEntityType `json:"entity_type,omitempty"`

	FileIngestionOptions *FileIngestionOptions `json:"file_ingestion_options,omitempty"`
	// Google Drive URL.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GoogleDriveOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GoogleDriveOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GoogleDriveOptionsGoogleDriveEntityType string

const GoogleDriveOptionsGoogleDriveEntityTypeFile GoogleDriveOptionsGoogleDriveEntityType = `FILE`

const GoogleDriveOptionsGoogleDriveEntityTypeFileMetadata GoogleDriveOptionsGoogleDriveEntityType = `FILE_METADATA`

const GoogleDriveOptionsGoogleDriveEntityTypePermission GoogleDriveOptionsGoogleDriveEntityType = `PERMISSION`

// String representation for [fmt.Print]
func (f *GoogleDriveOptionsGoogleDriveEntityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GoogleDriveOptionsGoogleDriveEntityType) Set(v string) error {
	switch v {
	case `FILE`, `FILE_METADATA`, `PERMISSION`:
		*f = GoogleDriveOptionsGoogleDriveEntityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE", "FILE_METADATA", "PERMISSION"`, v)
	}
}

// Values returns all possible values for GoogleDriveOptionsGoogleDriveEntityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GoogleDriveOptionsGoogleDriveEntityType) Values() []GoogleDriveOptionsGoogleDriveEntityType {
	return []GoogleDriveOptionsGoogleDriveEntityType{
		GoogleDriveOptionsGoogleDriveEntityTypeFile,
		GoogleDriveOptionsGoogleDriveEntityTypeFileMetadata,
		GoogleDriveOptionsGoogleDriveEntityTypePermission,
	}
}

// Type always returns GoogleDriveOptionsGoogleDriveEntityType to satisfy [pflag.Value] interface
func (f *GoogleDriveOptionsGoogleDriveEntityType) Type() string {
	return "GoogleDriveOptionsGoogleDriveEntityType"
}

type IngestionConfig struct {
	// Select a specific source report.
	Report *ReportSpec `json:"report,omitempty"`
	// Select all tables from a specific source schema.
	Schema *SchemaSpec `json:"schema,omitempty"`
	// Select a specific source table.
	Table *TableSpec `json:"table,omitempty"`
}

func (s *IngestionConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	ConnectionId string `json:"connection_id,omitempty"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName string `json:"connection_name"`
	// Optional, Internal. Parameters required to establish an initial
	// connection with the source.
	ConnectionParameters *ConnectionParameters `json:"connection_parameters,omitempty"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog string `json:"gateway_storage_catalog"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Spark Declarative Pipelines system will automatically
	// create the storage location under the catalog and schema.
	GatewayStorageName string `json:"gateway_storage_name,omitempty"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema string `json:"gateway_storage_schema"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionGatewayPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionGatewayPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IngestionPipelineDefinition struct {
	// The Unity Catalog connection that this ingestion pipeline uses to
	// communicate with the source. This is used with both connectors for
	// applications like Salesforce, Workday, and so on, and also database
	// connectors like Oracle, (connector_type = QUERY_BASED OR connector_type =
	// CDC). If connection name corresponds to database connectors like Oracle,
	// and connector_type is not provided then connector_type defaults to
	// QUERY_BASED. If connector_type is passed as CDC we use Combined Cdc
	// Managed Ingestion pipeline. Under certain conditions, this can be
	// replaced with ingestion_gateway_id to change the connector to Cdc Managed
	// Ingestion Pipeline with Gateway pipeline.
	ConnectionName string `json:"connection_name,omitempty"`
	// (Optional) Connector Type for sources. Ex: CDC, Query Based.
	ConnectorType ConnectorType `json:"connector_type,omitempty"`
	// (Optional) Location of staged data storage. This is required for
	// migration from Cdc Managed Ingestion Pipeline with Gateway pipeline to
	// Combined Cdc Managed Ingestion Pipeline. If not specified, the volume for
	// staged data will be created in catalog and schema/target specified in the
	// top level pipeline definition.
	DataStagingOptions *DataStagingOptions `json:"data_staging_options,omitempty"`
	// (Optional) A window that specifies a set of time ranges for snapshot
	// queries in CDC.
	FullRefreshWindow *OperationTimeWindow `json:"full_refresh_window,omitempty"`
	// Immutable. If set to true, the pipeline will ingest tables from the UC
	// foreign catalogs directly without the need to specify a UC connection or
	// ingestion gateway. The `source_catalog` fields in objects of
	// IngestionConfig are interpreted as the UC foreign catalogs to ingest
	// from.
	IngestFromUcForeignCatalog bool `json:"ingest_from_uc_foreign_catalog,omitempty"`
	// Identifier for the gateway that is used by this ingestion pipeline to
	// communicate with the source database. This is used with CDC connectors to
	// databases like SQL Server using a gateway pipeline (connector_type =
	// CDC). Under certain conditions, this can be replaced with connection_name
	// to change the connector to Combined Cdc Managed Ingestion Pipeline.
	IngestionGatewayId string `json:"ingestion_gateway_id,omitempty"`
	// Netsuite only configuration. When the field is set for a netsuite
	// connector, the jar stored in the field will be validated and added to the
	// classpath of pipeline's cluster.
	NetsuiteJarPath string `json:"netsuite_jar_path,omitempty"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects []IngestionConfig `json:"objects,omitempty"`
	// Top-level source configurations
	SourceConfigurations []SourceConfig `json:"source_configurations,omitempty"`
	// The type of the foreign source. The source type will be inferred from the
	// source connection or ingestion gateway. This field is output only and
	// will be ignored if provided.
	SourceType IngestionSourceType `json:"source_type,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Fanout configuration for multi-table routing from streaming sources. Routes
// each input record to a destination table based on a routing key derived from
// the record. The key value becomes the table name suffix:
// {destination_catalog}.{destination_schema}.{key_value}.
type IngestionPipelineDefinitionFanoutOptions struct {
	// Column path or SQL expression whose value determines the destination
	// table. Supports dotted paths (e.g. "value.event_name") and expressions
	// (e.g. "value:event_name::string").
	FanoutBy string `json:"fanout_by,omitempty"`
	// Optional transforms applied to each route's DataFrame before writing to
	// the destination table.
	Transforms []Transformer `json:"transforms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinitionFanoutOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinitionFanoutOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configurations that are only applicable for query-based ingestion connectors.
type IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig struct {
	// The names of the monotonically increasing columns in the source table
	// that are used to enable the table to be read and ingested incrementally
	// through structured streaming. The columns are allowed to have repeated
	// values but have to be non-decreasing. If the source data is merged into
	// the destination (e.g., using SCD Type 1 or Type 2), these columns will
	// implicitly define the `sequence_by` behavior. You can still explicitly
	// set `sequence_by` to override this default.
	CursorColumns []string `json:"cursor_columns,omitempty"`
	// Specifies a SQL WHERE condition that specifies that the source row has
	// been deleted. This is sometimes referred to as "soft-deletes". For
	// example: "Operation = 'DELETE'" or "is_deleted = true". This field is
	// orthogonal to `hard_deletion_sync_interval_in_seconds`, one for
	// soft-deletes and the other for hard-deletes. See also the
	// hard_deletion_sync_min_interval_in_seconds field for handling of "hard
	// deletes" where the source rows are physically removed from the table.
	DeletionCondition string `json:"deletion_condition,omitempty"`
	// Specifies the minimum interval (in seconds) between snapshots on primary
	// keys for detecting and synchronizing hard deletions—i.e., rows that
	// have been physically removed from the source table. This interval acts as
	// a lower bound. If ingestion runs less frequently than this value, hard
	// deletion synchronization will align with the actual ingestion frequency
	// instead of happening more often. If not set, hard deletion
	// synchronization via snapshots is disabled. This field is mutable and can
	// be updated without triggering a full snapshot.
	HardDeletionSyncMinIntervalInSeconds int64 `json:"hard_deletion_sync_min_interval_in_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IngestionPipelineDefinitionWorkdayReportParameters struct {
	// (Optional) Marks the report as incremental. This field is deprecated and
	// should not be used. Use `parameters` instead. The incremental behavior is
	// now controlled by the `parameters` field.
	Incremental bool `json:"incremental,omitempty"`
	// Parameters for the Workday report. Each key represents the parameter name
	// (e.g., "start_date", "end_date"), and the corresponding value is a
	// SQL-like expression used to compute the parameter value at runtime.
	// Example: { "start_date": "{ coalesce(current_offset(),
	// date(\"2025-02-01\")) }", "end_date": "{ current_date() - INTERVAL 1 DAY
	// }" }
	Parameters map[string]string `json:"parameters,omitempty"`
	// (Optional) Additional custom parameters for Workday Report This field is
	// deprecated and should not be used. Use `parameters` instead.
	ReportParameters []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue `json:"report_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinitionWorkdayReportParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinitionWorkdayReportParameters) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue struct {
	// Key for the report parameter, can be a column name or other metadata
	Key string `json:"key,omitempty"`
	// Value for the report parameter. Possible values it can take are these sql
	// functions: 1. coalesce(current_offset(), date("YYYY-MM-DD")) -> if
	// current_offset() is null, then the passed date, else current_offset() 2.
	// current_date() 3. date_sub(current_date(), x) -> subtract x (some
	// non-negative integer) days from current date
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of source system that a managed ingestion pipeline reads from.
type IngestionSourceType string

const IngestionSourceTypeBigquery IngestionSourceType = `BIGQUERY`

const IngestionSourceTypeConfluence IngestionSourceType = `CONFLUENCE`

const IngestionSourceTypeDynamics365 IngestionSourceType = `DYNAMICS365`

const IngestionSourceTypeForeignCatalog IngestionSourceType = `FOREIGN_CATALOG`

const IngestionSourceTypeGa4RawData IngestionSourceType = `GA4_RAW_DATA`

const IngestionSourceTypeGoogleAds IngestionSourceType = `GOOGLE_ADS`

const IngestionSourceTypeGoogleDrive IngestionSourceType = `GOOGLE_DRIVE`

const IngestionSourceTypeJira IngestionSourceType = `JIRA`

const IngestionSourceTypeManagedPostgresql IngestionSourceType = `MANAGED_POSTGRESQL`

const IngestionSourceTypeMetaMarketing IngestionSourceType = `META_MARKETING`

const IngestionSourceTypeMysql IngestionSourceType = `MYSQL`

const IngestionSourceTypeNetsuite IngestionSourceType = `NETSUITE`

const IngestionSourceTypeOnedrive IngestionSourceType = `ONEDRIVE`

const IngestionSourceTypeOracle IngestionSourceType = `ORACLE`

const IngestionSourceTypePostgresql IngestionSourceType = `POSTGRESQL`

const IngestionSourceTypeRabbitmq IngestionSourceType = `RABBITMQ`

const IngestionSourceTypeSalesforce IngestionSourceType = `SALESFORCE`

const IngestionSourceTypeServicenow IngestionSourceType = `SERVICENOW`

const IngestionSourceTypeSharepoint IngestionSourceType = `SHAREPOINT`

const IngestionSourceTypeSmartsheet IngestionSourceType = `SMARTSHEET`

const IngestionSourceTypeSqlserver IngestionSourceType = `SQLSERVER`

const IngestionSourceTypeTeradata IngestionSourceType = `TERADATA`

const IngestionSourceTypeTiktokAds IngestionSourceType = `TIKTOK_ADS`

const IngestionSourceTypeWorkdayRaas IngestionSourceType = `WORKDAY_RAAS`

const IngestionSourceTypeZendesk IngestionSourceType = `ZENDESK`

// String representation for [fmt.Print]
func (f *IngestionSourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IngestionSourceType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `CONFLUENCE`, `DYNAMICS365`, `FOREIGN_CATALOG`, `GA4_RAW_DATA`, `GOOGLE_ADS`, `GOOGLE_DRIVE`, `JIRA`, `MANAGED_POSTGRESQL`, `META_MARKETING`, `MYSQL`, `NETSUITE`, `ONEDRIVE`, `ORACLE`, `POSTGRESQL`, `RABBITMQ`, `SALESFORCE`, `SERVICENOW`, `SHAREPOINT`, `SMARTSHEET`, `SQLSERVER`, `TERADATA`, `TIKTOK_ADS`, `WORKDAY_RAAS`, `ZENDESK`:
		*f = IngestionSourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "CONFLUENCE", "DYNAMICS365", "FOREIGN_CATALOG", "GA4_RAW_DATA", "GOOGLE_ADS", "GOOGLE_DRIVE", "JIRA", "MANAGED_POSTGRESQL", "META_MARKETING", "MYSQL", "NETSUITE", "ONEDRIVE", "ORACLE", "POSTGRESQL", "RABBITMQ", "SALESFORCE", "SERVICENOW", "SHAREPOINT", "SMARTSHEET", "SQLSERVER", "TERADATA", "TIKTOK_ADS", "WORKDAY_RAAS", "ZENDESK"`, v)
	}
}

// Values returns all possible values for IngestionSourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *IngestionSourceType) Values() []IngestionSourceType {
	return []IngestionSourceType{
		IngestionSourceTypeBigquery,
		IngestionSourceTypeConfluence,
		IngestionSourceTypeDynamics365,
		IngestionSourceTypeForeignCatalog,
		IngestionSourceTypeGa4RawData,
		IngestionSourceTypeGoogleAds,
		IngestionSourceTypeGoogleDrive,
		IngestionSourceTypeJira,
		IngestionSourceTypeManagedPostgresql,
		IngestionSourceTypeMetaMarketing,
		IngestionSourceTypeMysql,
		IngestionSourceTypeNetsuite,
		IngestionSourceTypeOnedrive,
		IngestionSourceTypeOracle,
		IngestionSourceTypePostgresql,
		IngestionSourceTypeRabbitmq,
		IngestionSourceTypeSalesforce,
		IngestionSourceTypeServicenow,
		IngestionSourceTypeSharepoint,
		IngestionSourceTypeSmartsheet,
		IngestionSourceTypeSqlserver,
		IngestionSourceTypeTeradata,
		IngestionSourceTypeTiktokAds,
		IngestionSourceTypeWorkdayRaas,
		IngestionSourceTypeZendesk,
	}
}

// Type always returns IngestionSourceType to satisfy [pflag.Value] interface
func (f *IngestionSourceType) Type() string {
	return "IngestionSourceType"
}

// Jira specific options for ingestion
type JiraConnectorOptions struct {
	// (Optional) Projects to filter Jira data on
	IncludeJiraSpaces []string `json:"include_jira_spaces,omitempty"`
}

func (s *JiraConnectorOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type JsonTransformerOptions struct {
	// Parse the entire value as a single Variant column.
	AsVariant bool `json:"as_variant,omitempty"`
	// Inline schema string for JSON parsing (Spark DDL format).
	Schema string `json:"schema,omitempty"`
	// (Optional) Schema evolution mode for schema inference.
	SchemaEvolutionMode FileIngestionOptionsSchemaEvolutionMode `json:"schema_evolution_mode,omitempty"`
	// Path to a schema file (.ddl).
	SchemaFilePath string `json:"schema_file_path,omitempty"`
	// (Optional) Schema hints as a comma-separated string of "column_name type"
	// pairs.
	SchemaHints string `json:"schema_hints,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JsonTransformerOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JsonTransformerOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KafkaOptions struct {
	// Undocumented backdoor mechanism for overriding parameters to pass to the
	// Kafka client. This is not supported and may break at any time.
	ClientConfig map[string]string `json:"client_config,omitempty"`
	// (Optional) Transformer for the message key. If not specified, the key is
	// left as raw bytes.
	KeyTransformer *Transformer `json:"key_transformer,omitempty"`
	// Internal option to control the maximum number of offsets to process per
	// trigger.
	MaxOffsetsPerTrigger int64 `json:"max_offsets_per_trigger,omitempty"`
	// (Optional) Where to begin reading when no checkpoint exists. Valid
	// values: "latest" and "earliest". Defaults to "latest".
	StartingOffset string `json:"starting_offset,omitempty"`
	// Java regex pattern to subscribe to matching topics. Only one of topics or
	// topic_pattern must be specified.
	TopicPattern string `json:"topic_pattern,omitempty"`
	// Topics to subscribe to. Only one of topics or topic_pattern must be
	// specified.
	Topics []string `json:"topics,omitempty"`
	// (Optional) Transformer for the message value. If not specified, the value
	// is left as raw bytes.
	ValueTransformer *Transformer `json:"value_transformer,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KafkaOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KafkaOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// LinkedIn Ads specific options for ingestion. sync_start_date and
// lookback_window_days apply to both the prebuilt analytics tables and custom
// reports. custom_report_options defines a custom (user-defined) adAnalytics
// report and is only valid on a table object.
type LinkedInAdsOptions struct {
	// (Optional) Custom report definition. Only valid on a table object. When
	// set, the table is synthesized from /rest/adAnalytics using the finder,
	// pivots, time granularity and metrics here. When unset, the table must
	// match one of the connector's prebuilt sources.
	CustomReportOptions *LinkedInAdsOptionsLinkedInAdsCustomReportOptions `json:"custom_report_options,omitempty"`
	// (Optional) Days to look back during incremental sync for late-arriving
	// data. If not specified, defaults to 30 days.
	LookbackWindowDays int `json:"lookback_window_days,omitempty"`
	// (Optional) Start date for the initial sync of report tables, YYYY-MM-DD.
	// Earliest date from which to sync historical data; overrides the default
	// when set. For finder attributedRevenueMetrics, this must be between 30
	// and 366 days before today. If not specified, defaults to 1 year of
	// history.
	SyncStartDate string `json:"sync_start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LinkedInAdsOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LinkedInAdsOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User-defined custom report for the LinkedIn Ads connector. The destination
// table name comes from the enclosing TableSpec.destination_table, the start
// date from the enclosing LinkedInAdsOptions.sync_start_date, and the account
// it runs against from the source schema (namespace) -- none are repeated here.
type LinkedInAdsOptionsLinkedInAdsCustomReportOptions struct {
	// (Required) Entity pivots to group by; count/constraints depend on finder.
	EntityGranularity []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity `json:"entity_granularity,omitempty"`
	// (Required) adAnalytics finder. See LinkedInAdsFinder.
	Finder LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder `json:"finder"`
	// (Optional) LinkedIn metric names for the report. Open vocabulary (not an
	// enum): the valid set is large (~100) and evolves with the LinkedIn
	// adAnalytics API, so values are passed through verbatim. If empty, a
	// pivot-safe default core set is ingested: impressions, clicks,
	// costInLocalCurrency, externalWebsiteConversions (valid for every pivot).
	// Ignored for attributedRevenueMetrics (always returns the full
	// RevenueAttributionMetrics struct).
	Metrics []string `json:"metrics,omitempty"`
	// (Optional) Time aggregation. Defaults to DAILY when unspecified. Used by
	// analytics/statistics; ignored for attributedRevenueMetrics.
	TimeGranularity LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity `json:"time_granularity,omitempty"`
}

func (s *LinkedInAdsOptionsLinkedInAdsCustomReportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Entity pivot to group by.
type LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity string

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCampaign LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity = `CAMPAIGN`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCampaignGroup LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity = `CAMPAIGN_GROUP`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCreative LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity = `CREATIVE`

// String representation for [fmt.Print]
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity) Set(v string) error {
	switch v {
	case `CAMPAIGN`, `CAMPAIGN_GROUP`, `CREATIVE`:
		*f = LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAMPAIGN", "CAMPAIGN_GROUP", "CREATIVE"`, v)
	}
}

// Values returns all possible values for LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity.
//
// There is no guarantee on the order of the values in the slice.
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity) Values() []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity {
	return []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity{
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCampaign,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCampaignGroup,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularityCreative,
	}
}

// Type always returns LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity to satisfy [pflag.Value] interface
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity) Type() string {
	return "LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsEntityGranularity"
}

// adAnalytics finder. Determines call shape, valid pivots, and metric
// requirements.
type LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder string

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderAnalytics LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder = `ANALYTICS`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderAttributedRevenueMetrics LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder = `ATTRIBUTED_REVENUE_METRICS`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderStatistics LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder = `STATISTICS`

// String representation for [fmt.Print]
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder) Set(v string) error {
	switch v {
	case `ANALYTICS`, `ATTRIBUTED_REVENUE_METRICS`, `STATISTICS`:
		*f = LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ANALYTICS", "ATTRIBUTED_REVENUE_METRICS", "STATISTICS"`, v)
	}
}

// Values returns all possible values for LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder.
//
// There is no guarantee on the order of the values in the slice.
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder) Values() []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder {
	return []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder{
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderAnalytics,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderAttributedRevenueMetrics,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinderStatistics,
	}
}

// Type always returns LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder to satisfy [pflag.Value] interface
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder) Type() string {
	return "LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsFinder"
}

// Time aggregation. Used by analytics/statistics; ignored for
// attributedRevenueMetrics. Defaults to DAILY when unspecified.
type LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity string

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityAll LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity = `ALL`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityDaily LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity = `DAILY`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityMonthly LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity = `MONTHLY`

const LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityYearly LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity = `YEARLY`

// String representation for [fmt.Print]
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity) Set(v string) error {
	switch v {
	case `ALL`, `DAILY`, `MONTHLY`, `YEARLY`:
		*f = LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL", "DAILY", "MONTHLY", "YEARLY"`, v)
	}
}

// Values returns all possible values for LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity.
//
// There is no guarantee on the order of the values in the slice.
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity) Values() []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity {
	return []LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity{
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityAll,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityDaily,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityMonthly,
		LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularityYearly,
	}
}

// Type always returns LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity to satisfy [pflag.Value] interface
func (f *LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity) Type() string {
	return "LinkedInAdsOptionsLinkedInAdsCustomReportOptionsLinkedInAdsTimeGranularity"
}

type ListPipelineEventsRequest struct {
	// Criteria to select a subset of results, expressed using a SQL-like
	// syntax. The supported filters are: 1. level='INFO' (or WARN or ERROR) 2.
	// level in ('INFO', 'WARN') 3. id='[event-id]' 4. timestamp > 'TIMESTAMP'
	// (or >=,<,<=,=)
	//
	// Composite expressions are supported, for example: level in ('ERROR',
	// 'WARN') AND timestamp> '2021-07-22T06:37:33.083Z'
	Filter string `json:"-" url:"filter,omitempty"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The pipeline to return events for.
	PipelineId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelineEventsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	Events []PipelineEvent `json:"events,omitempty"`
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelineEventsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results pipelines in a response, even if there are
	// more pipelines available. This field is optional. The default value is
	// 25. The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A list of strings specifying the order of results. The supported order_by
	// field is id. The default is id asc. This field is optional.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelinesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of pipelines.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of pipelines matching the request criteria.
	Statuses []PipelineStateInfo `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelinesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The pipeline to return updates for.
	PipelineId string `json:"-" url:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId string `json:"-" url:"until_update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUpdatesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// The updates for the pipeline, ordered with the newest update first.
	Updates []UpdateInfo `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUpdatesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ManualTrigger struct {
}

func (s *ManualTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Marketo specific options for ingestion
type MarketoOptions struct {
	// (Optional) Start date for the initial sync in YYYY-MM-DD format. This
	// determines the earliest date from which to sync historical data. If not
	// specified, complete history is ingested.
	SyncStartDate string `json:"sync_start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MarketoOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MarketoOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for MaturityLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *MaturityLevel) Values() []MaturityLevel {
	return []MaturityLevel{
		MaturityLevelDeprecated,
		MaturityLevelEvolving,
		MaturityLevelStable,
	}
}

// Type always returns MaturityLevel to satisfy [pflag.Value] interface
func (f *MaturityLevel) Type() string {
	return "MaturityLevel"
}

// Meta Marketing (Meta Ads) specific options for ingestion
type MetaMarketingOptions struct {
	// (Optional, DEPRECATED — use
	// custom_report_options.action_attribution_windows) Action attribution
	// windows for insights reporting (e.g. "28d_click", "1d_view")
	ActionAttributionWindows []string `json:"action_attribution_windows,omitempty"`
	// (Optional, DEPRECATED — use custom_report_options.action_breakdowns)
	// Action breakdowns
	ActionBreakdowns []string `json:"action_breakdowns,omitempty"`
	// (Optional, DEPRECATED — use custom_report_options.action_report_time)
	// Timing used to report action statistics (impression, conversion, mixed,
	// or lifetime)
	ActionReportTime string `json:"action_report_time,omitempty"`
	// (Optional, DEPRECATED — use custom_report_options.breakdowns)
	// Breakdowns to configure
	Breakdowns []string `json:"breakdowns,omitempty"`
	// (Optional) Window in days to revisit data during sync to capture updated
	// conversion data from the API, shared by prebuilt and custom reports.
	CustomInsightsLookbackWindow int `json:"custom_insights_lookback_window,omitempty"`
	// (Optional) Per-table custom report definition. When set, defines the
	// shape of the insights call for this table
	// (level/fields/breakdowns/action_breakdowns/etc.). Supersedes the
	// deprecated flat report-shape fields above.
	CustomReportOptions *MetaMarketingOptionsMetaMarketingCustomReportOptions `json:"custom_report_options,omitempty"`
	// (Optional, DEPRECATED — use custom_report_options.level) Granularity of
	// data to pull (account, ad, adset, campaign)
	Level string `json:"level,omitempty"`
	// (Optional) Start date in yyyy-MM-dd format (e.g. 2025-01-15). Data added
	// after this date will be ingested, shared by prebuilt and custom reports.
	StartDate string `json:"start_date,omitempty"`
	// (Optional, DEPRECATED — use custom_report_options.time_increment) Value
	// in string by which to aggregate statistics (can take all_days, monthly or
	// number of days)
	TimeIncrement string `json:"time_increment,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetaMarketingOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetaMarketingOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Defines the shape of a single Meta Ads custom report (one /insights call
// shape). start_date, custom_insights_lookback_window live on
// MetaMarketingOptions, not here. Metrics are not customer-selectable; the
// connector returns a fixed standard metric set.
type MetaMarketingOptionsMetaMarketingCustomReportOptions struct {
	// (Optional) Action attribution windows for insights reporting (e.g.
	// "28d_click", "1d_view")
	ActionAttributionWindows []string `json:"action_attribution_windows,omitempty"`
	// (Optional) Action breakdowns to configure for data aggregation
	ActionBreakdowns []string `json:"action_breakdowns,omitempty"`
	// (Optional) Timing used to report action statistics (impression,
	// conversion, mixed, or lifetime)
	ActionReportTime string `json:"action_report_time,omitempty"`
	// (Optional) Breakdowns to configure for data aggregation
	Breakdowns []string `json:"breakdowns,omitempty"`
	// (Optional) Granularity of data to pull (account, ad, adset, campaign)
	Level string `json:"level,omitempty"`
	// (Optional) Value in string by which to aggregate statistics (all_days,
	// monthly or number of days)
	TimeIncrement string `json:"time_increment,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetaMarketingOptionsMetaMarketingCustomReportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetaMarketingOptionsMetaMarketingCustomReportOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotebookLibrary struct {
	// The absolute path of the source code.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	Alerts []string `json:"alerts,omitempty"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients []string `json:"email_recipients,omitempty"`
}

func (s *Notifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type OneDriveOptions struct {
	// (Optional) The type of OneDrive entity to ingest. If not specified,
	// defaults to FILE.
	EntityType OneDriveOptionsOneDriveEntityType `json:"entity_type,omitempty"`
	// (Optional) File ingestion options for processing files.
	FileIngestionOptions *FileIngestionOptions `json:"file_ingestion_options,omitempty"`
	// Required. The OneDrive URL.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OneDriveOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OneDriveOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of OneDrive entity to ingest.
type OneDriveOptionsOneDriveEntityType string

const OneDriveOptionsOneDriveEntityTypeFile OneDriveOptionsOneDriveEntityType = `FILE`

const OneDriveOptionsOneDriveEntityTypeFileMetadata OneDriveOptionsOneDriveEntityType = `FILE_METADATA`

// String representation for [fmt.Print]
func (f *OneDriveOptionsOneDriveEntityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OneDriveOptionsOneDriveEntityType) Set(v string) error {
	switch v {
	case `FILE`, `FILE_METADATA`:
		*f = OneDriveOptionsOneDriveEntityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE", "FILE_METADATA"`, v)
	}
}

// Values returns all possible values for OneDriveOptionsOneDriveEntityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OneDriveOptionsOneDriveEntityType) Values() []OneDriveOptionsOneDriveEntityType {
	return []OneDriveOptionsOneDriveEntityType{
		OneDriveOptionsOneDriveEntityTypeFile,
		OneDriveOptionsOneDriveEntityTypeFileMetadata,
	}
}

// Type always returns OneDriveOptionsOneDriveEntityType to satisfy [pflag.Value] interface
func (f *OneDriveOptionsOneDriveEntityType) Type() string {
	return "OneDriveOptionsOneDriveEntityType"
}

// Proto representing a window
type OperationTimeWindow struct {
	// Days of week in which the window is allowed to happen If not specified
	// all days of the week will be used.
	DaysOfWeek []DayOfWeek `json:"days_of_week,omitempty"`
	// An integer between 0 and 23 denoting the start hour for the window in the
	// 24-hour day.
	StartHour int `json:"start_hour"`
	// Time zone id of window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId string `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OperationTimeWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OperationTimeWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId int64 `json:"batch_id,omitempty"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud string `json:"cloud,omitempty"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId string `json:"cluster_id,omitempty"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName string `json:"dataset_name,omitempty"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId string `json:"flow_id,omitempty"`
	// The name of the flow. Not unique.
	FlowName string `json:"flow_name,omitempty"`
	// The optional host name where the event was triggered
	Host string `json:"host,omitempty"`
	// The name of the source catalog name (if known) from whose data ingestion
	// is described by this event.
	IngestionSourceCatalogName string `json:"ingestion_source_catalog_name,omitempty"`
	// The name of the source UC connection (if known) from whose data ingestion
	// is described by this event.
	IngestionSourceConnectionName string `json:"ingestion_source_connection_name,omitempty"`
	// The name of the source schema name (if known) from whose data ingestion
	// is described by this event.
	IngestionSourceSchemaName string `json:"ingestion_source_schema_name,omitempty"`
	// The name of the source table name (if known) from whose data ingestion is
	// described by this event.
	IngestionSourceTableName string `json:"ingestion_source_table_name,omitempty"`
	// An optional implementation-defined source table version of a dataset
	// being (re)ingested.
	IngestionSourceTableVersion string `json:"ingestion_source_table_version,omitempty"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId string `json:"maintenance_id,omitempty"`
	// Materialization name.
	MaterializationName string `json:"materialization_name,omitempty"`
	// The org id of the user. Unique within a cloud.
	OrgId int64 `json:"org_id,omitempty"`
	// The id of the pipeline. Globally unique.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The name of the pipeline. Not unique.
	PipelineName string `json:"pipeline_name,omitempty"`
	// The cloud region.
	Region string `json:"region,omitempty"`
	// The id of the request that caused an update.
	RequestId string `json:"request_id,omitempty"`
	// The id of a (delta) table. Globally unique.
	TableId string `json:"table_id,omitempty"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId string `json:"uc_resource_id,omitempty"`
	// The id of an execution. Globally unique.
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Origin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Origin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attachment behavior mode for Outlook ingestion
type OutlookAttachmentMode string

const OutlookAttachmentModeAll OutlookAttachmentMode = `ALL`

const OutlookAttachmentModeInlineOnly OutlookAttachmentMode = `INLINE_ONLY`

const OutlookAttachmentModeNone OutlookAttachmentMode = `NONE`

const OutlookAttachmentModeNonInlineOnly OutlookAttachmentMode = `NON_INLINE_ONLY`

// String representation for [fmt.Print]
func (f *OutlookAttachmentMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OutlookAttachmentMode) Set(v string) error {
	switch v {
	case `ALL`, `INLINE_ONLY`, `NONE`, `NON_INLINE_ONLY`:
		*f = OutlookAttachmentMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL", "INLINE_ONLY", "NONE", "NON_INLINE_ONLY"`, v)
	}
}

// Values returns all possible values for OutlookAttachmentMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *OutlookAttachmentMode) Values() []OutlookAttachmentMode {
	return []OutlookAttachmentMode{
		OutlookAttachmentModeAll,
		OutlookAttachmentModeInlineOnly,
		OutlookAttachmentModeNone,
		OutlookAttachmentModeNonInlineOnly,
	}
}

// Type always returns OutlookAttachmentMode to satisfy [pflag.Value] interface
func (f *OutlookAttachmentMode) Type() string {
	return "OutlookAttachmentMode"
}

// Body format for Outlook email content
type OutlookBodyFormat string

const OutlookBodyFormatTextHtml OutlookBodyFormat = `TEXT_HTML`

const OutlookBodyFormatTextPlain OutlookBodyFormat = `TEXT_PLAIN`

// String representation for [fmt.Print]
func (f *OutlookBodyFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OutlookBodyFormat) Set(v string) error {
	switch v {
	case `TEXT_HTML`, `TEXT_PLAIN`:
		*f = OutlookBodyFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEXT_HTML", "TEXT_PLAIN"`, v)
	}
}

// Values returns all possible values for OutlookBodyFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *OutlookBodyFormat) Values() []OutlookBodyFormat {
	return []OutlookBodyFormat{
		OutlookBodyFormatTextHtml,
		OutlookBodyFormatTextPlain,
	}
}

// Type always returns OutlookBodyFormat to satisfy [pflag.Value] interface
func (f *OutlookBodyFormat) Type() string {
	return "OutlookBodyFormat"
}

// Outlook specific options for ingestion
type OutlookOptions struct {
	// (Optional) Controls which attachments to ingest. If not specified,
	// defaults to ALL.
	AttachmentMode OutlookAttachmentMode `json:"attachment_mode,omitempty"`
	// (Optional) Defines how the body_content column is populated. TEXT_HTML:
	// Preserves full formatting, links, and styling. TEXT_PLAIN: Converts body
	// to plain text. Recommended for AI/RAG pipelines to reduce token usage and
	// noise.
	BodyFormat OutlookBodyFormat `json:"body_format,omitempty"`
	// Deprecated. Use include_folders instead.
	FolderFilter []string `json:"folder_filter,omitempty"`
	// (Optional) Filter mail folders to include in the sync. If not specified,
	// all folders will be synced. Examples: Inbox, Sent Items, Custom_Folder
	// Filter semantics: OR between different folders.
	IncludeFolders []string `json:"include_folders,omitempty"`
	// (Optional) List of mailboxes to sync (e.g. mailbox email addresses or
	// identifiers). If not specified, all accessible mailboxes are ingested.
	// Filter semantics: OR between different mailboxes.
	IncludeMailboxes []string `json:"include_mailboxes,omitempty"`
	// (Optional) Filter emails by sender address. Uses exact email match.
	// Examples: user@vendor.com, alerts@system.io, noreply@company.com If not
	// specified, emails from all senders will be synced. Filter semantics: OR
	// between different senders.
	IncludeSenders []string `json:"include_senders,omitempty"`
	// (Optional) Filter emails by subject line. Values ending with "*" use
	// prefix match (subject starts with the part before "*"); otherwise
	// substring match (subject contains the value). Examples: "Invoice"
	// (substring), "Re:*" (prefix), "Support Ticket", "URGENT*" If not
	// specified, emails with all subjects will be synced. Filter semantics: OR
	// between different subjects.
	IncludeSubjects []string `json:"include_subjects,omitempty"`
	// Deprecated. Use include_senders instead.
	SenderFilter []string `json:"sender_filter,omitempty"`
	// (Optional) Start date for the initial sync in YYYY-MM-DD format. Format:
	// YYYY-MM-DD (e.g., 2024-01-01) This determines the earliest date from
	// which to sync historical data. If not specified, complete history is
	// ingested.
	StartDate string `json:"start_date,omitempty"`
	// Deprecated. Use include_subjects instead.
	SubjectFilter []string `json:"subject_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OutlookOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OutlookOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Determines how errors encountered while deserializing records are handled.
type ParseMode string

const ParseModeFailfast ParseMode = `FAILFAST`

const ParseModePermissive ParseMode = `PERMISSIVE`

// String representation for [fmt.Print]
func (f *ParseMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ParseMode) Set(v string) error {
	switch v {
	case `FAILFAST`, `PERMISSIVE`:
		*f = ParseMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILFAST", "PERMISSIVE"`, v)
	}
}

// Values returns all possible values for ParseMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ParseMode) Values() []ParseMode {
	return []ParseMode{
		ParseModeFailfast,
		ParseModePermissive,
	}
}

// Type always returns ParseMode to satisfy [pflag.Value] interface
func (f *ParseMode) Type() string {
	return "ParseMode"
}

type PathPattern struct {
	// The source code to include for pipelines
	Include string `json:"include,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PathPattern) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PathPattern) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineAccessControlResponse struct {
	// All permissions.
	AllPermissions []PipelinePermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *PipelineClusterAutoscale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *compute.AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *compute.AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *compute.ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *compute.GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []compute.InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster settings applied to both the update and maintenance
	// clusters, `updates` to configure the update cluster, or `maintenance` to
	// configure the maintenance cluster. This field is optional. The default
	// value is `default`.
	Label string `json:"label,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
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
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
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
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	MaxWorkers int `json:"max_workers"`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int `json:"min_workers"`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode PipelineClusterAutoscaleMode `json:"mode,omitempty"`
}

func (s *PipelineClusterAutoscale) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for PipelineClusterAutoscaleMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineClusterAutoscaleMode) Values() []PipelineClusterAutoscaleMode {
	return []PipelineClusterAutoscaleMode{
		PipelineClusterAutoscaleModeEnhanced,
		PipelineClusterAutoscaleModeLegacy,
	}
}

// Type always returns PipelineClusterAutoscaleMode to satisfy [pflag.Value] interface
func (f *PipelineClusterAutoscaleMode) Type() string {
	return "PipelineClusterAutoscaleMode"
}

type PipelineDeployment struct {
	// ID of the deployment that manages this pipeline. Only set when `kind` is
	// `BUNDLE`. Used to look up deployment metadata from the Deployment
	// Metadata service.
	DeploymentId string `json:"deployment_id,omitempty"`
	// The deployment method that manages the pipeline.
	Kind DeploymentKind `json:"kind"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`
	// ID of the version of the deployment that produced this pipeline. Only set
	// when `kind` is `BUNDLE`. Identifies a specific snapshot of the deployment
	// in the Deployment Metadata service.
	VersionId string `json:"version_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	Error *ErrorDetail `json:"error,omitempty"`
	// The event type. Should always correspond to the details
	EventType string `json:"event_type,omitempty"`
	// A time-based, globally unique id.
	Id string `json:"id,omitempty"`
	// The severity level of the event.
	Level EventLevel `json:"level,omitempty"`
	// Maturity level for event_type.
	MaturityLevel MaturityLevel `json:"maturity_level,omitempty"`
	// The display message associated with the event.
	Message string `json:"message,omitempty"`
	// Describes where the event originates from.
	Origin *Origin `json:"origin,omitempty"`
	// A sequencing object to identify and order events.
	Sequence *Sequencing `json:"sequence,omitempty"`
	// The time of the event.
	Timestamp string `json:"timestamp,omitempty"`
	// Information about which fields were truncated from this event due to size
	// constraints. If empty or absent, no truncation occurred. See
	// https://docs.databricks.com/en/ldp/monitor-event-logs for information on
	// retrieving complete event data.
	Truncation *Truncation `json:"truncation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineEvent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineEvent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File *FileLibrary `json:"file,omitempty"`
	// The unified field to include source codes. Each entry can be a notebook
	// path, a file path, or a folder path that ends `/**`. This field cannot be
	// used together with `notebook` or `file`.
	Glob *PathPattern `json:"glob,omitempty"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed.
	Maven *compute.MavenLibrary `json:"maven,omitempty"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook *NotebookLibrary `json:"notebook,omitempty"`
	// URI of the whl to be installed.
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for PipelinePermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelinePermissionLevel) Values() []PipelinePermissionLevel {
	return []PipelinePermissionLevel{
		PipelinePermissionLevelCanManage,
		PipelinePermissionLevelCanRun,
		PipelinePermissionLevelCanView,
		PipelinePermissionLevelIsOwner,
	}
}

// Type always returns PipelinePermissionLevel to satisfy [pflag.Value] interface
func (f *PipelinePermissionLevel) Type() string {
	return "PipelinePermissionLevel"
}

type PipelinePermissions struct {
	AccessControlList []PipelineAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermissionsRequest struct {
	AccessControlList []PipelineAccessControlRequest `json:"access_control_list,omitempty"`
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func (s *PipelinePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PipelineSpec struct {
	// ID (UUID) of the serverless budget policy applied to this pipeline. Only
	// supported for serverless pipelines.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. Tables
	// are published to the schema set in `schema` (or the legacy `target`)
	// inside this catalog, for example `catalog`.`schema`.`table`; source code
	// can also use fully qualified names to write to other schemas and
	// catalogs. If `catalog` is not specified, the pipeline uses the legacy
	// Hive metastore.
	Catalog string `json:"catalog,omitempty"`
	// Release channel that selects the pipeline runtime version: `CURRENT` (the
	// default) or `PREVIEW`, which runs on the upcoming runtime release.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment. Applies to classic
	// (non-serverless) pipelines. Omit this field when `serverless` is `true`.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	//
	// Deprecated: wrap the pipeline in a continuous job instead, which also
	// lets you take advantage of job-level settings such as performance mode.
	// When the pipeline is started by a continuous job, the job's setting takes
	// precedence and this field is ignored.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	//
	// Deprecated: set development mode for each update instead.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition: `CORE`, `PRO`, or `ADVANCED`. Serverless
	// pipelines require `ADVANCED`.
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Path of the pipeline parent folder in workspace file tree.
	//
	// If absent, the pipeline doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath string `json:"root_path,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline. Serverless is
	// the recommended compute for new pipelines; set this to `true` to run the
	// pipeline on serverless. For notebook/file pipelines, omitting this field
	// uses classic compute configured through the `clusters` field. When
	// `true`, omit `clusters`; Photon is always enabled.
	Serverless bool `json:"serverless,omitempty"`
	// Serverless compute ID specified by the user for serverless pipelines. Use
	// compute_policy_id instead. Only one of these fields may be set.
	ServerlessComputeId string `json:"serverless_compute_id,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Test configuration for this pipeline. Used only by test-only updates.
	Testing *PipelineTesting `json:"testing,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
	// Usage policy of this pipeline.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for PipelineState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineState) Values() []PipelineState {
	return []PipelineState{
		PipelineStateDeleted,
		PipelineStateDeploying,
		PipelineStateFailed,
		PipelineStateIdle,
		PipelineStateRecovering,
		PipelineStateResetting,
		PipelineStateRunning,
		PipelineStateStarting,
		PipelineStateStopping,
	}
}

// Type always returns PipelineState to satisfy [pflag.Value] interface
func (f *PipelineState) Type() string {
	return "PipelineState"
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health PipelineStateInfoHealth `json:"health,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// The user-friendly name of the pipeline.
	Name string `json:"name,omitempty"`
	// The unique identifier of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The current state of the pipeline.
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The health of a pipeline.
type PipelineStateInfoHealth string

const PipelineStateInfoHealthHealthy PipelineStateInfoHealth = `HEALTHY`

const PipelineStateInfoHealthUnhealthy PipelineStateInfoHealth = `UNHEALTHY`

// String representation for [fmt.Print]
func (f *PipelineStateInfoHealth) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineStateInfoHealth) Set(v string) error {
	switch v {
	case `HEALTHY`, `UNHEALTHY`:
		*f = PipelineStateInfoHealth(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HEALTHY", "UNHEALTHY"`, v)
	}
}

// Values returns all possible values for PipelineStateInfoHealth.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineStateInfoHealth) Values() []PipelineStateInfoHealth {
	return []PipelineStateInfoHealth{
		PipelineStateInfoHealthHealthy,
		PipelineStateInfoHealthUnhealthy,
	}
}

// Type always returns PipelineStateInfoHealth to satisfy [pflag.Value] interface
func (f *PipelineStateInfoHealth) Type() string {
	return "PipelineStateInfoHealth"
}

// The packages that a pipeline's tests need in addition to the pipeline's
// environment.
type PipelineTestEnvironment struct {
	// A list of files that declare the packages the tests need: a project
	// folder, a requirements file, or a wheel, in the workspace or a UC volume.
	// A project folder can take `--editable` and an extra, for example
	// `--editable /Workspace/Users/someone@example.com/orders[test]`. Package
	// names, such as `faker==33.0`, and URLs are rejected.
	Dependencies []string `json:"dependencies,omitempty"`
}

func (s *PipelineTestEnvironment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Options for selecting the tests that a test-only update runs.
type PipelineTestOptions struct {
	// A list of arguments to pass to pytest, for example `["-m", "slow"]`. Only
	// `-k` and `-m` are accepted, each followed by its value. Other arguments
	// are rejected.
	PytestArgs []string `json:"pytest_args,omitempty"`
	// A list of test files or folders to run, as absolute workspace paths
	// inside the pipeline's test sources. A path outside the test sources is
	// rejected. To run one test, append `::` and the test name to a file path,
	// for example `/Workspace/orders/tests/test_orders.py::test_backfill`. When
	// empty, every test in the test sources runs.
	TestSelection []string `json:"test_selection,omitempty"`
}

func (s *PipelineTestOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A workspace location that holds pipeline tests.
type PipelineTestSource struct {
	// An absolute workspace path to a test file, or to a folder followed by
	// `/**`.
	Glob *PathPattern `json:"glob,omitempty"`
}

func (s *PipelineTestSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Test configuration for a pipeline.
type PipelineTesting struct {
	// Packages that only the tests need. Test-only updates install them in
	// addition to the pipeline's `environment`. Other updates do not.
	Environment *PipelineTestEnvironment `json:"environment,omitempty"`
	// The workspace files and folders that hold the pipeline's tests. They must
	// not overlap the pipeline's `libraries`.
	Sources []PipelineTestSource `json:"sources,omitempty"`
}

func (s *PipelineTesting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PipelineTrigger struct {
	Cron *CronTrigger `json:"cron,omitempty"`

	Manual *ManualTrigger `json:"manual,omitempty"`
}

func (s *PipelineTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and SDP's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type PipelinesEnvironment struct {
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url>
	Dependencies []string `json:"dependencies,omitempty"`
	// The environment version of the serverless Python environment used to
	// execute customer Python code. Each environment version includes a
	// specific Python version and a curated set of pre-installed libraries with
	// defined versions, providing a stable and reproducible execution
	// environment.
	//
	// Databricks supports a three-year lifecycle for each environment version.
	// For available versions and their included packages, see
	// https://docs.databricks.com/aws/en/release-notes/serverless/environment-version/
	//
	// The value should be a string representing the environment version number,
	// for example: `"4"`.
	EnvironmentVersion string `json:"environment_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinesEnvironment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinesEnvironment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// PG-specific catalog-level configuration parameters
type PostgresCatalogConfig struct {
	// Optional. The Postgres slot configuration to use for logical replication
	SlotConfig *PostgresSlotConfig `json:"slot_config,omitempty"`
}

func (s *PostgresCatalogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// PostgresSlotConfig contains the configuration for a Postgres logical
// replication slot
type PostgresSlotConfig struct {
	// The name of the publication to use for the Postgres source
	PublicationName string `json:"publication_name,omitempty"`
	// The name of the logical replication slot to use for the Postgres source
	SlotName string `json:"slot_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PostgresSlotConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PostgresSlotConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProtobufTransformerOptions struct {
	// Required: path to the .desc file (dbfs:/... or /Volumes/...).
	DescFilePath string `json:"desc_file_path,omitempty"`
	// Required: fully-qualified message type name.
	MessageName string `json:"message_name,omitempty"`
	// (Optional) Parse mode for Protobuf data. Valid values: FAILFAST,
	// PERMISSIVE. Defaults to PERMISSIVE.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// (Optional) Maximum expansion depth for recursive protobuf fields. Spark
	// SQL does not natively support recursive types, so recursive fields are
	// expanded up to this depth and truncated beyond it. Valid values: -1
	// (disallow recursive fields), 0 (drop), 1-10.
	RecursiveFieldsMaxDepth int `json:"recursive_fields_max_depth,omitempty"`
	// (Optional) Schema registry to resolve the Protobuf schema at runtime
	// instead of providing it via desc_file_path.
	SchemaRegistry *SchemaRegistryConfig `json:"schema_registry,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProtobufTransformerOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProtobufTransformerOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Enum representing the publishing mode of a pipeline.
type PublishingMode string

const PublishingModeDefaultPublishingMode PublishingMode = `DEFAULT_PUBLISHING_MODE`

const PublishingModeLegacyPublishingMode PublishingMode = `LEGACY_PUBLISHING_MODE`

// String representation for [fmt.Print]
func (f *PublishingMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PublishingMode) Set(v string) error {
	switch v {
	case `DEFAULT_PUBLISHING_MODE`, `LEGACY_PUBLISHING_MODE`:
		*f = PublishingMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFAULT_PUBLISHING_MODE", "LEGACY_PUBLISHING_MODE"`, v)
	}
}

// Values returns all possible values for PublishingMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PublishingMode) Values() []PublishingMode {
	return []PublishingMode{
		PublishingModeDefaultPublishingMode,
		PublishingModeLegacyPublishingMode,
	}
}

// Type always returns PublishingMode to satisfy [pflag.Value] interface
func (f *PublishingMode) Type() string {
	return "PublishingMode"
}

// RabbitMQ specific options for ingestion. Performance tuning options
// (consumers_per_task, max_messages_per_fetch, etc.) are intentionally not
// exposed in the public API. The managed connector uses sensible defaults
// internally. These can be added later if user demand arises.
type RabbitmqOptions struct {
	// (Required) RabbitMQ queue name to consume from.
	Queue string `json:"queue"`
}

func (s *RabbitmqOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Reddit Ads specific options for ingestion
type RedditAdsOptions struct {
	// (Optional) Custom report definition. When set, the table is treated as a
	// user-defined Reddit Ads custom report. When unset, the table must match
	// one of the connector's prebuilt sources.
	CustomReportOptions *RedditAdsOptionsRedditAdsCustomReportOptions `json:"custom_report_options,omitempty"`
	// (Optional) Number of days to look back for report tables during
	// incremental sync to capture late-arriving conversions and attribution
	// data. If not specified, defaults to 30 days.
	LookbackWindowDays int `json:"lookback_window_days,omitempty"`
	// (Optional) Start date for the initial sync of report tables in YYYY-MM-DD
	// format. This determines the earliest date from which to sync historical
	// data. If not specified, defaults to 2 years ago.
	SyncStartDate string `json:"sync_start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RedditAdsOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RedditAdsOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User-defined custom report for the Reddit Ads connector. Applies only to the
// custom_report table — prebuilt tables ignore this.
type RedditAdsOptionsRedditAdsCustomReportOptions struct {
	// (Optional) Breakdown dimensions to group report data by. Examples:
	// CAMPAIGN_ID, DATE, COUNTRY, REGION, AD_ID. Must include at least one time
	// dimension (DATE or HOUR).
	Breakdowns []string `json:"breakdowns,omitempty"`
	// (Optional) Fields to include in the report (maps to the Reddit Ads API
	// `fields` parameter). Examples: IMPRESSIONS, CLICKS, SPEND, CPC, CTR.
	Fields []string `json:"fields,omitempty"`
}

func (s *RedditAdsOptionsRedditAdsCustomReportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Specifies a replace_where predicate override for a replace where flow.
type ReplaceWhereOverride struct {
	// Name of the flow to apply this override to.
	FlowName string `json:"flow_name,omitempty"`
	// SQL predicate string to use as replace_where condition. Example: `date =
	// '2024-10-10' AND city = 'xyz'`
	PredicateOverride string `json:"predicate_override,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ReplaceWhereOverride) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReplaceWhereOverride) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable string `json:"destination_table,omitempty"`
	// Required. Report URL in the source system.
	SourceUrl string `json:"source_url"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ReportSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReportSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	DaysOfWeek []DayOfWeek `json:"days_of_week,omitempty"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour int `json:"start_hour"`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId string `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestartWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestartWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for rewinding a specific dataset.
type RewindDatasetSpec struct {
	// Whether to cascade the rewind to dependent datasets. Must be specified.
	Cascade bool `json:"cascade,omitempty"`
	// The identifier of the dataset (e.g., "main.foo.tbl1").
	Identifier string `json:"identifier,omitempty"`
	// Whether to reset checkpoints for this dataset.
	ResetCheckpoints bool `json:"reset_checkpoints,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RewindDatasetSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RewindDatasetSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Information about a rewind being requested for this pipeline or some of the
// datasets in it.
type RewindSpec struct {
	// List of datasets to rewind with specific configuration for each. When not
	// specified, all datasets will be rewound with cascade = true and
	// reset_checkpoints = true.
	Datasets []RewindDatasetSpec `json:"datasets,omitempty"`
	// If true, this is a dry run and we should emit the RewindSummary but not
	// perform the rewind.
	DryRun bool `json:"dry_run,omitempty"`
	// The base timestamp to rewind to. Exactly one of rewind_timestamp or
	// rewind_point_id must be specified.
	RewindTimestamp string `json:"rewind_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RewindSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RewindSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Write-only setting, available only in Create/Update calls. Specifies the user
// or service principal that the pipeline runs as. If not specified, the
// pipeline runs as the user who created the pipeline.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type RunAs struct {
	// Group name of an account group assigned to the workspace. When set, the
	// pipeline runs as the group and the group's permissions are used for data
	// access. Setting this field requires being a member of the group, or
	// having the `Assume` permission on the group.
	GroupName string `json:"group_name,omitempty"`
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SchemaRegistryConfig struct {
	// Required: Confluent-compatible schema registry options.
	ConfluentOptions *ConfluentSchemaRegistryOptions `json:"confluent_options,omitempty"`
	// (Optional) UC connection for registry authentication. Specify if
	// different from the top-level source connection.
	ConnectionName string `json:"connection_name,omitempty"`
	// (Optional, Protobuf only) Selects a specific message from a schema that
	// defines multiple Protobuf messages. Simple ("Location") or
	// fully-qualified ("com.example.protos.Location"). Defaults to the first
	// message.
	ProtobufMessageName string `json:"protobuf_message_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaRegistryConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaRegistryConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SchemaSpec struct {
	// (Optional) Source Specific Connector Options
	ConnectorOptions *ConnectorOptions `json:"connector_options,omitempty"`
	// Required. Destination catalog to store tables.
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails if a table with the same name already exists.
	DestinationSchema string `json:"destination_schema"`
	// Fanout options for multi-table routing from streaming sources. When set,
	// records are routed to destination tables based on a per-record routing
	// key. The key value becomes the table name:
	// {destination_catalog}.{destination_schema}.{key_value}.
	FanoutOptions *IngestionPipelineDefinitionFanoutOptions `json:"fanout_options,omitempty"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Schema name in the source database. Optional: some source types (for
	// example streaming or message-bus connectors) do not use it, so it may be
	// absent from a pipeline's definition. Clients that assume it is always
	// present should handle its absence.
	SourceSchema string `json:"source_schema,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Sequencing struct {
	// A sequence number, unique and increasing per pipeline.
	ControlPlaneSeqNo int64 `json:"control_plane_seq_no,omitempty"`
	// the ID assigned by the data plane.
	DataPlaneId *DataPlaneId `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Sequencing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Sequencing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName string `json:"class_name,omitempty"`
	// Exception message
	Message string `json:"message,omitempty"`
	// Stack trace consisting of a list of stack frames
	Stack []StackFrame `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SerializedException) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SerializedException) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SharepointOptions struct {
	// (Optional) The type of SharePoint entity to ingest. If not specified,
	// defaults to FILE.
	EntityType SharepointOptionsSharepointEntityType `json:"entity_type,omitempty"`
	// (Optional) File ingestion options for processing files.
	FileIngestionOptions *FileIngestionOptions `json:"file_ingestion_options,omitempty"`
	// Required. The SharePoint URL.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharepointOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharepointOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SharepointOptionsSharepointEntityType string

const SharepointOptionsSharepointEntityTypeFile SharepointOptionsSharepointEntityType = `FILE`

const SharepointOptionsSharepointEntityTypeFileMetadata SharepointOptionsSharepointEntityType = `FILE_METADATA`

const SharepointOptionsSharepointEntityTypeList SharepointOptionsSharepointEntityType = `LIST`

const SharepointOptionsSharepointEntityTypePermission SharepointOptionsSharepointEntityType = `PERMISSION`

// String representation for [fmt.Print]
func (f *SharepointOptionsSharepointEntityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharepointOptionsSharepointEntityType) Set(v string) error {
	switch v {
	case `FILE`, `FILE_METADATA`, `LIST`, `PERMISSION`:
		*f = SharepointOptionsSharepointEntityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE", "FILE_METADATA", "LIST", "PERMISSION"`, v)
	}
}

// Values returns all possible values for SharepointOptionsSharepointEntityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharepointOptionsSharepointEntityType) Values() []SharepointOptionsSharepointEntityType {
	return []SharepointOptionsSharepointEntityType{
		SharepointOptionsSharepointEntityTypeFile,
		SharepointOptionsSharepointEntityTypeFileMetadata,
		SharepointOptionsSharepointEntityTypeList,
		SharepointOptionsSharepointEntityTypePermission,
	}
}

// Type always returns SharepointOptionsSharepointEntityType to satisfy [pflag.Value] interface
func (f *SharepointOptionsSharepointEntityType) Type() string {
	return "SharepointOptionsSharepointEntityType"
}

// Smartsheet specific options for ingestion
type SmartsheetOptions struct {
	// (Optional) When true, maps each column to its Smartsheet-declared type
	// (Text/Number/Date/ Checkbox/etc.). Cells that do not conform to the
	// declared type are set to NULL. When false, all columns land as STRING.
	// Use false for sheets with irregular data or columns that frequently
	// violate their own declared type. If not specified, defaults to true.
	EnforceSchema bool `json:"enforce_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SmartsheetOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SmartsheetOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SourceCatalogConfig contains catalog-level custom configuration parameters
// for each source
type SourceCatalogConfig struct {
	// Postgres-specific catalog-level configuration parameters
	Postgres *PostgresCatalogConfig `json:"postgres,omitempty"`
	// Source catalog name
	SourceCatalog string `json:"source_catalog,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SourceCatalogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SourceCatalogConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SourceConfig struct {
	// Connector-specific top-level configuration for API Source connectors.
	ApiSourceConnectorConfig *ApiSourceConnectorConfig `json:"api_source_connector_config,omitempty"`
	// Catalog-level source configuration parameters
	Catalog *SourceCatalogConfig `json:"catalog,omitempty"`

	GoogleAdsConfig *GoogleAdsConfig `json:"google_ads_config,omitempty"`
}

func (s *SourceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass string `json:"declaring_class,omitempty"`
	// File where the method is defined
	FileName string `json:"file_name,omitempty"`
	// Line from which the method was called
	LineNumber int `json:"line_number,omitempty"`
	// Name of the method which was called
	MethodName string `json:"method_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StackFrame) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StackFrame) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartUpdate struct {
	// What triggered this update, as recorded in the update's `cause`. Defaults
	// to `API_CALL`.
	Cause StartUpdateCause `json:"cause,omitempty"`
	// Whether the update is started in the development mode. This is
	// recommended for interactive development and testing. Reuses compute for
	// faster iteration and disables automatic retries. Not recommended for
	// production.
	Development bool `json:"development,omitempty"`
	// If true, this update will reset all tables before running.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// A list of file paths that identify the tables to update with a full
	// refresh.
	FullRefreshSelectionByPath []string `json:"full_refresh_selection_by_path,omitempty"`
	// Key/value map of parameters to pass to the pipeline execution
	Parameters map[string]string `json:"parameters,omitempty"`
	// The ID of the pipeline to start an update for.
	PipelineId string `json:"-" url:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// A list of file paths that identify the tables to update.
	RefreshSelectionByPath []string `json:"refresh_selection_by_path,omitempty"`
	// A list of predicate overrides for replace_where flows in this update.
	// Only replace_where flows may be specified. Flows not listed use their
	// original predicate.
	ReplaceWhereOverrides []ReplaceWhereOverride `json:"replace_where_overrides,omitempty"`
	// A list of flows for which this update should reset the streaming
	// checkpoint. This selection will not clear the data in the flow's target
	// table. Flows in this list may also appear in refresh_selection and
	// full_refresh_selection.
	ResetCheckpointSelection []string `json:"reset_checkpoint_selection,omitempty"`
	// The information about the requested rewind operation. If specified this
	// is a rewind mode update.
	RewindSpec *RewindSpec `json:"rewind_spec,omitempty"`
	// If true, this is a test-only update: it runs the pipeline's tests and
	// does not update the pipeline's datasets. A failing test does not stop the
	// other tests. The update is `COMPLETED` if at least one test runs and no
	// test fails or errors. Otherwise the update is `FAILED`, including when no
	// tests are found. Cannot be combined with `validate_only`, `explore_only`,
	// or `rewind_spec`.
	TestOnly bool `json:"test_only,omitempty"`
	// Options for selecting the tests that this test-only update runs. Requires
	// `test_only`. When absent, every test in the pipeline's test sources runs.
	TestOptions *PipelineTestOptions `json:"test_options,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StartUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// What triggered this update.
type StartUpdateCause string

const StartUpdateCauseApiCall StartUpdateCause = `API_CALL`

const StartUpdateCauseInfrastructureMaintenance StartUpdateCause = `INFRASTRUCTURE_MAINTENANCE`

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
	case `API_CALL`, `INFRASTRUCTURE_MAINTENANCE`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = StartUpdateCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "INFRASTRUCTURE_MAINTENANCE", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Values returns all possible values for StartUpdateCause.
//
// There is no guarantee on the order of the values in the slice.
func (f *StartUpdateCause) Values() []StartUpdateCause {
	return []StartUpdateCause{
		StartUpdateCauseApiCall,
		StartUpdateCauseInfrastructureMaintenance,
		StartUpdateCauseJobTask,
		StartUpdateCauseRetryOnFailure,
		StartUpdateCauseSchemaChange,
		StartUpdateCauseServiceUpgrade,
		StartUpdateCauseUserAction,
	}
}

// Type always returns StartUpdateCause to satisfy [pflag.Value] interface
func (f *StartUpdateCause) Type() string {
	return "StartUpdateCause"
}

type StartUpdateResponse struct {
	// The ID of the update that was started.
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StartUpdateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StopRequest struct {
	// The ID of the pipeline to stop.
	PipelineId string `json:"-" url:"-"`
}

func (s *StopRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TableSpec struct {
	// (Optional) Source Specific Connector Options
	ConnectorOptions *ConnectorOptions `json:"connector_options,omitempty"`
	// Required. Destination catalog to store table.
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable string `json:"destination_table,omitempty"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema string `json:"source_schema,omitempty"`
	// Table name in the source database. Optional: some source types (for
	// example streaming or message-bus connectors) do not use it, so it may be
	// absent from a pipeline's definition. Clients that assume it is always
	// present should handle its absence.
	SourceTable string `json:"source_table,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableSpecificConfig struct {
	// (Optional, Mutable) Policy for auto full refresh, if enabled pipeline
	// will automatically try to fix issues by doing a full refresh on the table
	// in the retry run. auto_full_refresh_policy in table configuration will
	// override the above level auto_full_refresh_policy. For example, {
	// "auto_full_refresh_policy": { "enabled": true, "min_interval_hours": 23,
	// } } If unspecified, auto full refresh is disabled.
	AutoFullRefreshPolicy *AutoFullRefreshPolicy `json:"auto_full_refresh_policy,omitempty"`
	// List of column names to use for clustering the destination table. When
	// specified, the destination Delta table will be clustered by these
	// columns. This can improve query performance when filtering on these
	// columns. Note: clustering_columns in table specific configuration will
	// override the pipeline definition. Note: we can only provide
	// enable_auto_clustering or clustering_columns, added as separate fields as
	// we cannot have repeated field in oneof.
	ClusteringColumns []string `json:"clustering_columns,omitempty"`
	// Whether to enable auto clustering on the destination table. When enabled,
	// Delta will automatically optimize the data layout based on the clustering
	// columns for improved query performance. Note: enable_auto_clustering in
	// table specific configuration will override the pipeline definition. Note:
	// we can only provide enable_auto_clustering or clustering_columns, added
	// as separate fields as we cannot have repeated field in oneof.
	EnableAutoClustering bool `json:"enable_auto_clustering,omitempty"`
	// A list of column names to be excluded for the ingestion. When not
	// specified, include_columns fully controls what columns to be ingested.
	// When specified, all other columns including future ones will be
	// automatically included for ingestion. This field in mutually exclusive
	// with `include_columns`.
	ExcludeColumns []string `json:"exclude_columns,omitempty"`
	// A list of column names to be included for the ingestion. When not
	// specified, all columns except ones in exclude_columns will be included.
	// Future columns will be automatically included. When specified, all other
	// future columns will be automatically excluded from ingestion. This field
	// in mutually exclusive with `exclude_columns`.
	IncludeColumns []string `json:"include_columns,omitempty"`
	// The primary key of the table used to apply changes.
	PrimaryKeys []string `json:"primary_keys,omitempty"`

	QueryBasedConnectorConfig *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig `json:"query_based_connector_config,omitempty"`
	// (Optional, Immutable) The row filter condition to be applied to the
	// table. It must not contain the WHERE keyword, only the actual filter
	// condition. It must be in DBSQL format.
	RowFilter string `json:"row_filter,omitempty"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields bool `json:"salesforce_include_formula_fields,omitempty"`

	ScdType TableSpecificConfigScdType `json:"scd_type,omitempty"`
	// The column names specifying the logical order of events in the source
	// data. Spark Declarative Pipelines uses this sequencing to handle change
	// events that arrive out of order.
	SequenceBy []string `json:"sequence_by,omitempty"`
	// (Optional) Name of the struct column added to each ingested record to
	// hold per row source metadata.
	SourceMetadataColumn string `json:"source_metadata_column,omitempty"`
	// Table properties to set on the destination table. These are key-value
	// pairs that configure various Delta table behaviors or any user defined
	// properties. Example: {"delta.feature.variantType": "supported",
	// "delta.enableTypeWidening": "true"} Note: table_properties in table
	// specific configuration will override the table_properties of the pipeline
	// definition.
	TableProperties map[string]string `json:"table_properties,omitempty"`
	// (Optional) Additional custom parameters for Workday Report
	WorkdayReportParameters *IngestionPipelineDefinitionWorkdayReportParameters `json:"workday_report_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSpecificConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpecificConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The SCD type to use to ingest the table.
type TableSpecificConfigScdType string

const TableSpecificConfigScdTypeAppendOnly TableSpecificConfigScdType = `APPEND_ONLY`

const TableSpecificConfigScdTypeScdType1 TableSpecificConfigScdType = `SCD_TYPE_1`

const TableSpecificConfigScdTypeScdType2 TableSpecificConfigScdType = `SCD_TYPE_2`

// String representation for [fmt.Print]
func (f *TableSpecificConfigScdType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableSpecificConfigScdType) Set(v string) error {
	switch v {
	case `APPEND_ONLY`, `SCD_TYPE_1`, `SCD_TYPE_2`:
		*f = TableSpecificConfigScdType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPEND_ONLY", "SCD_TYPE_1", "SCD_TYPE_2"`, v)
	}
}

// Values returns all possible values for TableSpecificConfigScdType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableSpecificConfigScdType) Values() []TableSpecificConfigScdType {
	return []TableSpecificConfigScdType{
		TableSpecificConfigScdTypeAppendOnly,
		TableSpecificConfigScdTypeScdType1,
		TableSpecificConfigScdTypeScdType2,
	}
}

// Type always returns TableSpecificConfigScdType to satisfy [pflag.Value] interface
func (f *TableSpecificConfigScdType) Type() string {
	return "TableSpecificConfigScdType"
}

// TikTok Ads specific options for ingestion
type TikTokAdsOptions struct {
	// (Optional) Custom report definition. When set, the table is treated as a
	// user-defined TikTok Ads custom report: the connector synthesizes a report
	// request from the dimensions, metrics, report type, and data level
	// specified here. Supersedes the deprecated top-level
	// dimensions/metrics/report_type/ data_level/query_lifetime fields above.
	CustomReportOptions *TikTokAdsOptionsTikTokAdsCustomReportOptions `json:"custom_report_options,omitempty"`
	// Deprecated. Use custom_report_options.data_level instead.
	DataLevel TikTokAdsOptionsTikTokDataLevel `json:"data_level,omitempty"`
	// Deprecated. Use custom_report_options.dimensions instead.
	Dimensions []string `json:"dimensions,omitempty"`
	// (Optional) Number of days to look back for report tables during
	// incremental sync to capture late-arriving conversions and attribution
	// data.
	LookbackWindowDays int `json:"lookback_window_days,omitempty"`
	// Deprecated. Use custom_report_options.metrics instead.
	Metrics []string `json:"metrics,omitempty"`
	// Deprecated. Use custom_report_options.query_lifetime instead.
	QueryLifetime bool `json:"query_lifetime,omitempty"`
	// Deprecated. Use custom_report_options.report_type instead.
	ReportType TikTokAdsOptionsTikTokReportType `json:"report_type,omitempty"`
	// (Optional) Start date for the initial sync of report tables in YYYY-MM-DD
	// format. This determines the earliest date from which to sync historical
	// data.
	SyncStartDate string `json:"sync_start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TikTokAdsOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TikTokAdsOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User-defined custom report for the TikTok Ads connector. Groups the
// dimensions + metrics + report type + data level that define a TikTok Ads
// custom report request.
type TikTokAdsOptionsTikTokAdsCustomReportOptions struct {
	// (Optional) Data level for the report. If not specified, defaults to
	// AUCTION_CAMPAIGN.
	DataLevel TikTokAdsOptionsTikTokDataLevel `json:"data_level,omitempty"`
	// (Optional) Dimensions to include in the report (e.g. "campaign_id",
	// "adgroup_id", "ad_id", "stat_time_day", "stat_time_hour").
	Dimensions []string `json:"dimensions,omitempty"`
	// (Optional) Metrics to include in the report (e.g. "spend", "impressions",
	// "clicks", "conversion", "cpc").
	Metrics []string `json:"metrics,omitempty"`
	// (Optional) Whether to request lifetime metrics (all-time aggregated
	// data). When true, the report returns all-time data. If not specified,
	// defaults to false.
	QueryLifetime bool `json:"query_lifetime,omitempty"`
	// (Optional) Report type for the TikTok Ads API. If not specified, defaults
	// to BASIC.
	ReportType TikTokAdsOptionsTikTokReportType `json:"report_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TikTokAdsOptionsTikTokAdsCustomReportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TikTokAdsOptionsTikTokAdsCustomReportOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Data level for TikTok Ads report aggregation.
type TikTokAdsOptionsTikTokDataLevel string

const TikTokAdsOptionsTikTokDataLevelAuctionAd TikTokAdsOptionsTikTokDataLevel = `AUCTION_AD`

const TikTokAdsOptionsTikTokDataLevelAuctionAdgroup TikTokAdsOptionsTikTokDataLevel = `AUCTION_ADGROUP`

const TikTokAdsOptionsTikTokDataLevelAuctionAdvertiser TikTokAdsOptionsTikTokDataLevel = `AUCTION_ADVERTISER`

const TikTokAdsOptionsTikTokDataLevelAuctionCampaign TikTokAdsOptionsTikTokDataLevel = `AUCTION_CAMPAIGN`

// String representation for [fmt.Print]
func (f *TikTokAdsOptionsTikTokDataLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TikTokAdsOptionsTikTokDataLevel) Set(v string) error {
	switch v {
	case `AUCTION_AD`, `AUCTION_ADGROUP`, `AUCTION_ADVERTISER`, `AUCTION_CAMPAIGN`:
		*f = TikTokAdsOptionsTikTokDataLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUCTION_AD", "AUCTION_ADGROUP", "AUCTION_ADVERTISER", "AUCTION_CAMPAIGN"`, v)
	}
}

// Values returns all possible values for TikTokAdsOptionsTikTokDataLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *TikTokAdsOptionsTikTokDataLevel) Values() []TikTokAdsOptionsTikTokDataLevel {
	return []TikTokAdsOptionsTikTokDataLevel{
		TikTokAdsOptionsTikTokDataLevelAuctionAd,
		TikTokAdsOptionsTikTokDataLevelAuctionAdgroup,
		TikTokAdsOptionsTikTokDataLevelAuctionAdvertiser,
		TikTokAdsOptionsTikTokDataLevelAuctionCampaign,
	}
}

// Type always returns TikTokAdsOptionsTikTokDataLevel to satisfy [pflag.Value] interface
func (f *TikTokAdsOptionsTikTokDataLevel) Type() string {
	return "TikTokAdsOptionsTikTokDataLevel"
}

// Report type for TikTok Ads API.
type TikTokAdsOptionsTikTokReportType string

const TikTokAdsOptionsTikTokReportTypeAudience TikTokAdsOptionsTikTokReportType = `AUDIENCE`

const TikTokAdsOptionsTikTokReportTypeBasic TikTokAdsOptionsTikTokReportType = `BASIC`

const TikTokAdsOptionsTikTokReportTypeBusinessCenter TikTokAdsOptionsTikTokReportType = `BUSINESS_CENTER`

const TikTokAdsOptionsTikTokReportTypeDsa TikTokAdsOptionsTikTokReportType = `DSA`

const TikTokAdsOptionsTikTokReportTypeGmvMax TikTokAdsOptionsTikTokReportType = `GMV_MAX`

const TikTokAdsOptionsTikTokReportTypePlayableAd TikTokAdsOptionsTikTokReportType = `PLAYABLE_AD`

// String representation for [fmt.Print]
func (f *TikTokAdsOptionsTikTokReportType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TikTokAdsOptionsTikTokReportType) Set(v string) error {
	switch v {
	case `AUDIENCE`, `BASIC`, `BUSINESS_CENTER`, `DSA`, `GMV_MAX`, `PLAYABLE_AD`:
		*f = TikTokAdsOptionsTikTokReportType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUDIENCE", "BASIC", "BUSINESS_CENTER", "DSA", "GMV_MAX", "PLAYABLE_AD"`, v)
	}
}

// Values returns all possible values for TikTokAdsOptionsTikTokReportType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TikTokAdsOptionsTikTokReportType) Values() []TikTokAdsOptionsTikTokReportType {
	return []TikTokAdsOptionsTikTokReportType{
		TikTokAdsOptionsTikTokReportTypeAudience,
		TikTokAdsOptionsTikTokReportTypeBasic,
		TikTokAdsOptionsTikTokReportTypeBusinessCenter,
		TikTokAdsOptionsTikTokReportTypeDsa,
		TikTokAdsOptionsTikTokReportTypeGmvMax,
		TikTokAdsOptionsTikTokReportTypePlayableAd,
	}
}

// Type always returns TikTokAdsOptionsTikTokReportType to satisfy [pflag.Value] interface
func (f *TikTokAdsOptionsTikTokReportType) Type() string {
	return "TikTokAdsOptionsTikTokReportType"
}

// One step that transforms a column: either decodes data into structured data
// (`format`) or reshapes an already-decoded column (`operation`).
type Transformer struct {
	AvroOptions *AvroTransformerOptions `json:"avro_options,omitempty"`
	// The format to decode the input column from.
	Format TransformerFormat `json:"format,omitempty"`
	// Optional input column to transform. When set, the transformer reads from
	// this column instead of the default source column.
	InputColumn string `json:"input_column,omitempty"`

	JsonOptions *JsonTransformerOptions `json:"json_options,omitempty"`
	// Optional output column name. When set, the transformed result is written
	// to this column instead of replacing the input column.
	OutputColumn string `json:"output_column,omitempty"`

	ProtobufOptions *ProtobufTransformerOptions `json:"protobuf_options,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Transformer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Transformer) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TransformerFormat string

const TransformerFormatAvro TransformerFormat = `AVRO`

const TransformerFormatJson TransformerFormat = `JSON`

const TransformerFormatProtobuf TransformerFormat = `PROTOBUF`

const TransformerFormatString TransformerFormat = `STRING`

// String representation for [fmt.Print]
func (f *TransformerFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TransformerFormat) Set(v string) error {
	switch v {
	case `AVRO`, `JSON`, `PROTOBUF`, `STRING`:
		*f = TransformerFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "JSON", "PROTOBUF", "STRING"`, v)
	}
}

// Values returns all possible values for TransformerFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *TransformerFormat) Values() []TransformerFormat {
	return []TransformerFormat{
		TransformerFormatAvro,
		TransformerFormatJson,
		TransformerFormatProtobuf,
		TransformerFormatString,
	}
}

// Type always returns TransformerFormat to satisfy [pflag.Value] interface
func (f *TransformerFormat) Type() string {
	return "TransformerFormat"
}

// Information about truncations applied to this event.
type Truncation struct {
	// List of fields that were truncated from this event. If empty or absent,
	// no truncation occurred.
	TruncatedFields []TruncationTruncationDetail `json:"truncated_fields,omitempty"`
}

func (s *Truncation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details about a specific field that was truncated.
type TruncationTruncationDetail struct {
	// The name of the truncated field (e.g., "error"). Corresponds to field
	// names in PipelineEvent.
	FieldName string `json:"field_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TruncationTruncationDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TruncationTruncationDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateInfo struct {
	// What triggered this update.
	Cause UpdateInfoCause `json:"cause,omitempty"`
	// The ID of the cluster that the update is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *PipelineSpec `json:"config,omitempty"`
	// The time when this update was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// If true, this update will reset all tables before running.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// A list of file paths that identify the tables to update with a full
	// refresh.
	FullRefreshSelectionByPath []string `json:"full_refresh_selection_by_path,omitempty"`
	// Indicates whether the update is either part of a continuous job run, or
	// running in legacy continuous pipeline mode. Returned only for GetUpdate;
	// not populated in ListUpdates responses.
	Mode UpdateMode `json:"mode,omitempty"`
	// Key/value map of parameters used to initiate the update
	Parameters map[string]string `json:"parameters,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// A list of file paths that identify the tables to update.
	RefreshSelectionByPath []string `json:"refresh_selection_by_path,omitempty"`
	// The update state.
	State UpdateInfoState `json:"state,omitempty"`
	// If true, this update is a test-only update. Absent for other updates.
	TestOnly bool `json:"test_only,omitempty"`
	// The ID of this update.
	UpdateId string `json:"update_id,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// What triggered this update.
type UpdateInfoCause string

const UpdateInfoCauseApiCall UpdateInfoCause = `API_CALL`

const UpdateInfoCauseInfrastructureMaintenance UpdateInfoCause = `INFRASTRUCTURE_MAINTENANCE`

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
	case `API_CALL`, `INFRASTRUCTURE_MAINTENANCE`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = UpdateInfoCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "INFRASTRUCTURE_MAINTENANCE", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Values returns all possible values for UpdateInfoCause.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateInfoCause) Values() []UpdateInfoCause {
	return []UpdateInfoCause{
		UpdateInfoCauseApiCall,
		UpdateInfoCauseInfrastructureMaintenance,
		UpdateInfoCauseJobTask,
		UpdateInfoCauseRetryOnFailure,
		UpdateInfoCauseSchemaChange,
		UpdateInfoCauseServiceUpgrade,
		UpdateInfoCauseUserAction,
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

// Values returns all possible values for UpdateInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateInfoState) Values() []UpdateInfoState {
	return []UpdateInfoState{
		UpdateInfoStateCanceled,
		UpdateInfoStateCompleted,
		UpdateInfoStateCreated,
		UpdateInfoStateFailed,
		UpdateInfoStateInitializing,
		UpdateInfoStateQueued,
		UpdateInfoStateResetting,
		UpdateInfoStateRunning,
		UpdateInfoStateSettingUpTables,
		UpdateInfoStateStopping,
		UpdateInfoStateWaitingForResources,
	}
}

// Type always returns UpdateInfoState to satisfy [pflag.Value] interface
func (f *UpdateInfoState) Type() string {
	return "UpdateInfoState"
}

type UpdateMode string

const UpdateModeContinuous UpdateMode = `CONTINUOUS`

const UpdateModeDefault UpdateMode = `DEFAULT`

// String representation for [fmt.Print]
func (f *UpdateMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateMode) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `DEFAULT`:
		*f = UpdateMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "DEFAULT"`, v)
	}
}

// Values returns all possible values for UpdateMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateMode) Values() []UpdateMode {
	return []UpdateMode{
		UpdateModeContinuous,
		UpdateModeDefault,
	}
}

// Type always returns UpdateMode to satisfy [pflag.Value] interface
func (f *UpdateMode) Type() string {
	return "UpdateMode"
}

type UpdateStateInfo struct {
	// The time when the update was created, in ISO 8601 format.
	CreationTime string `json:"creation_time,omitempty"`
	// The state of the update.
	State UpdateStateInfoState `json:"state,omitempty"`
	// The ID of the update.
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The update state.
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

// Values returns all possible values for UpdateStateInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateStateInfoState) Values() []UpdateStateInfoState {
	return []UpdateStateInfoState{
		UpdateStateInfoStateCanceled,
		UpdateStateInfoStateCompleted,
		UpdateStateInfoStateCreated,
		UpdateStateInfoStateFailed,
		UpdateStateInfoStateInitializing,
		UpdateStateInfoStateQueued,
		UpdateStateInfoStateResetting,
		UpdateStateInfoStateRunning,
		UpdateStateInfoStateSettingUpTables,
		UpdateStateInfoStateStopping,
		UpdateStateInfoStateWaitingForResources,
	}
}

// Type always returns UpdateStateInfoState to satisfy [pflag.Value] interface
func (f *UpdateStateInfoState) Type() string {
	return "UpdateStateInfoState"
}

// Zendesk Support specific options for ingestion
type ZendeskSupportOptions struct {
	// (Optional) Start date in YYYY-MM-DD format for the initial sync. This
	// determines the earliest date from which to sync historical data.
	StartDate string `json:"start_date,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ZendeskSupportOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ZendeskSupportOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
