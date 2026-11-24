// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Per-catalog alert-related config. Singleton. Name format:
// `catalogs/{catalog}/alertConfig`.
type AlertConfig struct {
	// Window for collating Signals into security alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1h`; the
	// window in effect is reported in
	// `effective_default_security_alert_collation_window`.
	DefaultSecurityAlertCollationWindow string `json:"default_security_alert_collation_window,omitempty"`
	// Window for collating SystemEvents into system alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1d`; the
	// window in effect is reported in
	// `effective_default_system_alert_collation_window`.
	DefaultSystemAlertCollationWindow string `json:"default_system_alert_collation_window,omitempty"`
	// The security alert collation window in effect:
	// `default_security_alert_collation_window` when set, otherwise `1h`.
	EffectiveDefaultSecurityAlertCollationWindow string `json:"effective_default_security_alert_collation_window,omitempty"`
	// The system alert collation window in effect:
	// `default_system_alert_collation_window` when set, otherwise `1d`.
	EffectiveDefaultSystemAlertCollationWindow string `json:"effective_default_system_alert_collation_window,omitempty"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag string `json:"etag,omitempty"`

	Name string `json:"name,omitempty"`
	// Export destination for security alerts. When unset, export is off and
	// reads return `{"enabled": false}`.
	SecurityAlertExportConfig *ExportConfig `json:"security_alert_export_config,omitempty"`
	// Export destination for system alerts. When unset, export is off and reads
	// return `{"enabled": false}`.
	SystemAlertExportConfig *ExportConfig `json:"system_alert_export_config,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Ingests files from cloud storage with Auto Loader. Mutually exclusive with
// `unity_catalog_connection` and `existing_table`.
type AutoLoader struct {
	// Name of the workspace secret scope that holds the AWS access key ID and
	// secret access key, under the keys `awsAccessKeyId` and
	// `awsSecretAccessKey`.
	AwsSecretScope string `json:"aws_secret_scope,omitempty"`
	// Whether ingestion started with the files already present at `location`.
	// Populated on get, create, and update responses.
	EffectiveIncludeExistingFiles bool `json:"effective_include_existing_files,omitempty"`
	// The table that the raw data is written to: `raw_table` in the
	// datasource's catalog and schema, under their current Unity Catalog names.
	// Populated on get, create, and update responses.
	EffectiveRawTable *TableName `json:"effective_raw_table,omitempty"`
	// How to unwrap a transport envelope around the source records before
	// parsing. Supported only for the JSON and JSONL formats. If unset, each
	// source record is treated as one event.
	Envelope *Envelope `json:"envelope,omitempty"`
	// File format of the source files. Required.
	Format AutoLoaderFormat `json:"format"`
	// Format-specific parsing options, passed to Auto Loader. Supported for the
	// CSV, JSON, JSONL, and XML formats, each of which accepts a fixed set of
	// options. The XML format requires the `rowTag` option.
	FormatOptions map[string]string `json:"format_options,omitempty"`
	// Whether to ingest the files already present at `location` when ingestion
	// first starts. Set true to ingest the existing files and then every new
	// file. Set false, or leave unset, to ingest only files that arrive after
	// the datasource starts. Defaults to false. Cannot be changed after
	// creation. Reads return the value as supplied; the setting in effect is
	// reported in `effective_include_existing_files`.
	IncludeExistingFiles bool `json:"include_existing_files,omitempty"`
	// URI of the source files: a path in a Unity Catalog volume or external
	// location, or a cloud storage URL such as an Amazon S3 path. Required,
	// including when new files are discovered through `queue_url`. Cannot be
	// changed after creation.
	Location string `json:"location"`
	// Hive-style partition columns to derive from the source file paths. If
	// unset, no partition columns are derived.
	OverridePartitionColumns []string `json:"override_partition_columns,omitempty"`
	// URL of a customer-managed Amazon SQS queue that receives file-arrival
	// notifications. Set it to discover new files through the queue. If unset,
	// new files are discovered through the file events of the source's Unity
	// Catalog external location when it has them configured, and by directory
	// listing otherwise. Cannot be changed after creation.
	QueueUrl string `json:"queue_url,omitempty"`
	// Table that stores the raw data before parsers schematize it. Every Auto
	// Loader datasource writes its raw rows to this table before any parser
	// reads them. The table must be in the datasource's catalog and schema; an
	// omitted `catalog_name` or `schema_name` defaults to them. Cannot be
	// changed after creation. Reads return the value as supplied; the table the
	// data is written to is reported in `effective_raw_table`.
	RawTable TableName `json:"raw_table"`
	// Name of the Unity Catalog service credential to use.
	ServiceCredentials string `json:"service_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutoLoader) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoLoader) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// File format of the source files that Auto Loader reads.
type AutoLoaderFormat string

const AutoLoaderFormatAutoloaderFormatCsv AutoLoaderFormat = `AUTOLOADER_FORMAT_CSV`

const AutoLoaderFormatAutoloaderFormatJson AutoLoaderFormat = `AUTOLOADER_FORMAT_JSON`

const AutoLoaderFormatAutoloaderFormatJsonl AutoLoaderFormat = `AUTOLOADER_FORMAT_JSONL`

const AutoLoaderFormatAutoloaderFormatParquet AutoLoaderFormat = `AUTOLOADER_FORMAT_PARQUET`

const AutoLoaderFormatAutoloaderFormatXml AutoLoaderFormat = `AUTOLOADER_FORMAT_XML`

// String representation for [fmt.Print]
func (f *AutoLoaderFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AutoLoaderFormat) Set(v string) error {
	switch v {
	case `AUTOLOADER_FORMAT_CSV`, `AUTOLOADER_FORMAT_JSON`, `AUTOLOADER_FORMAT_JSONL`, `AUTOLOADER_FORMAT_PARQUET`, `AUTOLOADER_FORMAT_XML`:
		*f = AutoLoaderFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTOLOADER_FORMAT_CSV", "AUTOLOADER_FORMAT_JSON", "AUTOLOADER_FORMAT_JSONL", "AUTOLOADER_FORMAT_PARQUET", "AUTOLOADER_FORMAT_XML"`, v)
	}
}

// Values returns all possible values for AutoLoaderFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *AutoLoaderFormat) Values() []AutoLoaderFormat {
	return []AutoLoaderFormat{
		AutoLoaderFormatAutoloaderFormatCsv,
		AutoLoaderFormatAutoloaderFormatJson,
		AutoLoaderFormatAutoloaderFormatJsonl,
		AutoLoaderFormatAutoloaderFormatParquet,
		AutoLoaderFormatAutoloaderFormatXml,
	}
}

// Type always returns AutoLoaderFormat to satisfy [pflag.Value] interface
func (f *AutoLoaderFormat) Type() string {
	return "AutoLoaderFormat"
}

// A Databricks-authored parser, shared by every customer.
//
// Same container-plus-versions shape as the custom `Parser`, and addressed the
// same way, by UUID. Built-ins are global rather than catalog-scoped, so they
// carry no owning catalog or schema.
type BuiltinParser struct {
	// Newest published version number.
	LatestVersion int64 `json:"latest_version,omitempty"`
	// Resource name. Format: `builtin-parsers/{parser_id}`.
	Name string `json:"name,omitempty"`
	// The parser's UUID, authored into the built-in's published definition
	// rather than minted per customer, so every workspace addresses the same
	// built-in by the same id. It is the leaf of `name`.
	ParserId string `json:"parser_id,omitempty"`
	// The built-in's own leaf name, as published in content-marketplace. Stable
	// and human-readable, which is what a UI shows; `name` addresses the
	// resource by id.
	ParserName string `json:"parser_name,omitempty"`
	// Key/value tags describing this parser, e.g. `cloud_provider=aws` or
	// `log_source=cloudtrail`. A parser may carry up to 32. Filterable on list.
	Tags []ParserTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BuiltinParser) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BuiltinParser) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// One published version of a built-in parser, carrying the spec.
//
// The content pipeline owns versioning, so unlike the custom version there is
// no authoring principal and no append timestamp here.
type BuiltinParserVersion struct {
	// Free-form description of this parser, authored by Databricks.
	Description string `json:"description,omitempty"`
	// Resource name. Format: `builtin-parsers/{parser_id}/versions/{version}`.
	Name string `json:"name,omitempty"`
	// The parent built-in's published leaf name, copied from the container on
	// read. Denormalized so a caller that needs identity alongside the spec
	// reads the version alone rather than the container as well. Populated by
	// `GetBuiltinParserVersion`; List omits it, as List omits `spec`.
	ParserName string `json:"parser_name,omitempty"`
	// The parser spec: the field definitions for this log source, as a YAML
	// document. Read-only; built-in specs are authored by Databricks, so no
	// client-facing size cap is enforced here.
	Spec string `json:"spec,omitempty"`
	// The parent built-in's tags, copied from the container on read. Rendering
	// reads these for connector-specific details, such as the source table a
	// Lakeflow Connect parser reads.
	Tags []ParserTag `json:"tags,omitempty"`
	// Version number.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BuiltinParserVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BuiltinParserVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Unwraps an AWS CloudWatch Logs delivery envelope: each element of `logEvents`
// becomes its own row, with the element's `message` parsed as the event body.
type CloudWatchLogsEnvelope struct {
	// Whether to keep the envelope's `owner`, `logGroup`, and `logStream`
	// values in the `data.lw_header` field of each unwrapped row whose message
	// is a JSON object. A message that is any other JSON value is kept without
	// the field. Must be set explicitly to `true` or `false`.
	HeaderRetention bool `json:"header_retention,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CloudWatchLogsEnvelope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CloudWatchLogsEnvelope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Catch-all per-catalog config — settings that don't belong to any one
// resource family. Singleton. Name format: `catalogs/{catalog}/commonConfig`.
type CommonConfig struct {
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag string `json:"etag,omitempty"`

	Name string `json:"name,omitempty"`
	// Whether users may grant permissions on individual resources rather than
	// only at the Folder level.
	PerResourcePermissionsEnabled bool `json:"per_resource_permissions_enabled,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CommonConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CommonConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reference to the ContentPack (and pack version) a resource was originally
// installed from.
type ContentPackReference struct {
	// UUID of the content pack.
	Uuid string `json:"uuid,omitempty"`
	// Pack version the resource was installed from.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContentPackReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContentPackReference) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDatasourceRequest struct {
	Datasource Datasource `json:"datasource"`
}

func (s *CreateDatasourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateParserRequest struct {
	// The parser to create. `spec` and `description` become version 1.
	Parser Parser `json:"parser"`
}

func (s *CreateParserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateParserVersionRequest struct {
	// Resource name of the parser to append a version to. Format:
	// `parsers/{parser_id}`.
	Parent string `json:"-" url:"-"`
	// The version to append.
	ParserVersion ParserVersion `json:"parser_version"`
}

func (s *CreateParserVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateRuleRequest struct {
	// The rule to create. Server-assigned fields (`name`, `rule_id`,
	// timestamps, `etag`, …) are ignored on the request.
	Rule Rule `json:"rule"`
}

func (s *CreateRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Lakewatch datasource: a pipeline that ingests security data and writes it
// to Unity Catalog tables.
//
// What a datasource ingests is defined by its ingestion method (`autoloader`,
// `unity_catalog_connection`, or `existing_table`), which sets how the data
// arrives, together with its `parsers`, which define the schemas that the
// arriving data is matched against and written as.
type Datasource struct {
	// Free-form key-value annotations on the datasource, managed by the client.
	Annotations map[string]string `json:"annotations,omitempty"`

	Autoloader *AutoLoader `json:"autoloader,omitempty"`
	// Name of the Unity Catalog catalog that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the catalog's current name is reported in
	// `effective_catalog_name`.
	CatalogName string `json:"catalog_name"`
	// Compute mode for the datasource's ingestion job. If unset, the
	// `default_compute_mode` of the catalog's datasource configuration applies,
	// or `DS_COMPUTE_PERFORMANCE_OPTIMIZED` if the catalog has none.
	ComputeMode DatasourceComputeMode `json:"compute_mode,omitempty"`
	// ID of the workspace whose compute runs this datasource's ingestion job
	// and owns its associated resources, such as the storage the ingested data
	// lands in. Server-assigned.
	ComputeWorkspaceId string `json:"compute_workspace_id,omitempty"`

	CreateTime *time.Time `json:"create_time,omitempty"`

	CreatedBy *Principal `json:"created_by,omitempty"`
	// Unity Catalog securable ID (a UUID) of the datasource. It never changes,
	// including when the datasource is renamed in Unity Catalog, and is the
	// final segment of `name`.
	DatasourceId string `json:"datasource_id,omitempty"`
	// Name of the datasource in Unity Catalog, unique within its schema.
	// Required on create and can be updated within the same catalog and schema.
	// Get synchronizes this field with the current Unity Catalog name, also
	// reported in `effective_datasource_name`.
	DatasourceName string `json:"datasource_name"`
	// Free-form description of the datasource.
	Description string `json:"description,omitempty"`
	// Current name of the Unity Catalog catalog that contains the datasource,
	// which reflects catalog renames without a change to `catalog_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveCatalogName string `json:"effective_catalog_name,omitempty"`
	// Current UC leaf name of the datasource (server-owned). Tracks UC renames
	// so callers can observe the live name without having to update the
	// client-owned `datasource_name`; unset until resolved.
	EffectiveDatasourceName string `json:"effective_datasource_name,omitempty"`
	// The identity the datasource runs as: the resolved `run_as_principal`, or
	// the requester's identity if `run_as_principal` was omitted on create or
	// cleared on update. Given as a user email or service principal application
	// ID, or as the numeric principal ID if the identity can no longer be
	// resolved. Populated on get, create, and update responses; list responses
	// omit it.
	EffectiveRunAsPrincipal string `json:"effective_run_as_principal,omitempty"`
	// Current name of the Unity Catalog schema that contains the datasource,
	// which reflects schema renames without a change to `schema_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveSchemaName string `json:"effective_schema_name,omitempty"`
	// Opaque entity tag for optimistic concurrency control. Returned by get,
	// create, and update responses and by list responses in the
	// `DATASOURCE_VIEW_FULL` view. To make an update conditional on the
	// datasource not having changed since it was read, pass this value as the
	// update request's `etag`. Changes to `status` do not change the etag.
	Etag string `json:"etag,omitempty"`

	ExistingTable *ExistingTable `json:"existing_table,omitempty"`
	// Resource name of the datasource, in the format
	// `datasources/{datasource_id}`. Unlike the datasource's Unity Catalog
	// location (`catalog_name`, `schema_name`, and `datasource_name`), the name
	// does not change when the datasource or its catalog or schema is renamed
	// in Unity Catalog.
	Name string `json:"name,omitempty"`
	// Unity Catalog owner of the datasource, assigned by the server. Create
	// responses omit it.
	Owner *Principal `json:"owner,omitempty"`
	// Parsers that schematize the ingested data. Each record is matched against
	// the parsers in order and written to the parsed table of the first parser
	// that it matches. Existing-table and Unity Catalog connection datasources
	// require at least one parser; an Auto Loader datasource with no parsers
	// writes only its raw table.
	Parsers []ParserRef `json:"parsers,omitempty"`
	// The user or service principal that the datasource runs as, given as a
	// user email or service principal application ID. When it is supplied on
	// create or update, the requester must be authorized to run the datasource
	// as that principal, or the request is rejected. When it is omitted on
	// create, the datasource runs as the requester. On update, a mask that
	// includes `run_as_principal` without a value clears it, so the datasource
	// runs as the requester; a mask that omits it keeps the current identity.
	// Reads return the value as supplied; the identity the datasource runs as
	// is reported in `effective_run_as_principal`.
	RunAsPrincipal string `json:"run_as_principal,omitempty"`
	// Schedule on which the datasource's ingestion job runs. A datasource
	// cannot be paused, so a request that sets `schedule.enabled` to `false` is
	// rejected. Continuous schedules are not supported for Unity Catalog
	// connection datasources. For an existing-table datasource, runs must be no
	// further apart than the shorter of the source table's
	// `delta.deletedFileRetentionDuration` and `delta.logRetentionDuration`.
	Schedule Schedule `json:"schedule"`
	// Name of the Unity Catalog schema that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the schema's current name is reported in
	// `effective_schema_name`.
	SchemaName string `json:"schema_name"`
	// Runtime status of the jobs and pipelines that run this datasource.
	Status *DatasourceStatus `json:"status,omitempty"`

	UnityCatalogConnection *UnityCatalogConnection `json:"unity_catalog_connection,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Version number of the datasource, incremented on each update.
	// Informational only; use `etag` for optimistic concurrency control.
	Version int64 `json:"version,omitempty"`
	// ID of the SQL warehouse that runs the checks Lakewatch performs to
	// validate a create or update request, such as source and destination
	// access checks. Required and non-blank on every create and update request
	// except a rename-only PATCH; there is no default warehouse. The requester
	// must be able to use the warehouse. It is not stored, is not returned on
	// reads, and does not affect the compute that runs the ingestion job.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Datasource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Datasource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Compute mode for the datasource's ingestion job.
type DatasourceComputeMode string

const DatasourceComputeModeDsComputePerformanceOptimized DatasourceComputeMode = `DS_COMPUTE_PERFORMANCE_OPTIMIZED`

const DatasourceComputeModeDsComputeStandard DatasourceComputeMode = `DS_COMPUTE_STANDARD`

// String representation for [fmt.Print]
func (f *DatasourceComputeMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatasourceComputeMode) Set(v string) error {
	switch v {
	case `DS_COMPUTE_PERFORMANCE_OPTIMIZED`, `DS_COMPUTE_STANDARD`:
		*f = DatasourceComputeMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DS_COMPUTE_PERFORMANCE_OPTIMIZED", "DS_COMPUTE_STANDARD"`, v)
	}
}

// Values returns all possible values for DatasourceComputeMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatasourceComputeMode) Values() []DatasourceComputeMode {
	return []DatasourceComputeMode{
		DatasourceComputeModeDsComputePerformanceOptimized,
		DatasourceComputeModeDsComputeStandard,
	}
}

// Type always returns DatasourceComputeMode to satisfy [pflag.Value] interface
func (f *DatasourceComputeMode) Type() string {
	return "DatasourceComputeMode"
}

// Per-catalog datasource-related config. Singleton. Name format:
// `catalogs/{catalog}/datasourceConfig`.
type DatasourceConfig struct {
	// Default compute mode applied to datasource ingestion jobs when the
	// individual datasource doesn't override it. v1
	// `datasources.defaultComputeMode`.
	DefaultComputeMode DatasourceComputeMode `json:"default_compute_mode,omitempty"`
	// Schema name of the registered folder (in this catalog) that a new
	// datasource is placed in when the request does not specify a schema. Must
	// name a schema registered as a folder. Omit to clear. Reads return the
	// value as supplied; the folder's current schema name is reported in
	// `effective_default_folder`.
	DefaultFolder string `json:"default_folder,omitempty"`
	// The current schema name of the `default_folder` folder. Tracks UC schema
	// renames. Unset when no default is configured, or when the folder's schema
	// no longer exists or is not visible to the caller.
	EffectiveDefaultFolder string `json:"effective_default_folder,omitempty"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag string `json:"etag,omitempty"`

	Name string `json:"name,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatasourceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatasourceConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Read-only runtime status of a datasource, with the status of each job and
// pipeline that runs it.
type DatasourceStatus struct {
	// Status of the datasource's ingestion job.
	IngestionJob *ResourceStatus `json:"ingestion_job,omitempty"`
	// Status of the managed ingestion pipeline. Populated only when
	// `unity_catalog_connection` is set, with `pipeline_id` identifying the
	// pipeline.
	IngestionPipeline *ResourceStatus `json:"ingestion_pipeline,omitempty"`
}

func (s *DatasourceStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Which fields list responses return for each datasource.
type DatasourceView string

const DatasourceViewDatasourceViewBasic DatasourceView = `DATASOURCE_VIEW_BASIC`

const DatasourceViewDatasourceViewFull DatasourceView = `DATASOURCE_VIEW_FULL`

// String representation for [fmt.Print]
func (f *DatasourceView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatasourceView) Set(v string) error {
	switch v {
	case `DATASOURCE_VIEW_BASIC`, `DATASOURCE_VIEW_FULL`:
		*f = DatasourceView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATASOURCE_VIEW_BASIC", "DATASOURCE_VIEW_FULL"`, v)
	}
}

// Values returns all possible values for DatasourceView.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatasourceView) Values() []DatasourceView {
	return []DatasourceView{
		DatasourceViewDatasourceViewBasic,
		DatasourceViewDatasourceViewFull,
	}
}

// Type always returns DatasourceView to satisfy [pflag.Value] interface
func (f *DatasourceView) Type() string {
	return "DatasourceView"
}

type DeleteDatasourceRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *DeleteDatasourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteParserRequest struct {
	// Resource name of the parser. Format: `parsers/{parser_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteParserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteParserVersionRequest struct {
	// Resource name of the version. Format:
	// `parsers/{parser_id}/versions/{version}`, where `{version}` is a version
	// number written without sign or leading zeros, such as `3`. Unlike
	// GetParserVersion, the alias `latest` and forms such as `007` or `+3` are
	// rejected, so a delete always names one fixed version.
	Name string `json:"-" url:"-"`
}

func (s *DeleteParserVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRuleRequest struct {
	// Resource name of the rule to delete. Format: `rules/{rule_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Unwraps a transport envelope from the source data so that each contained
// event becomes its own row. Applied when the source data is read, before
// parsing.
type Envelope struct {
	CloudWatchLogs *CloudWatchLogsEnvelope `json:"cloud_watch_logs,omitempty"`

	JsonArray *JsonArrayEnvelope `json:"json_array,omitempty"`
}

func (s *Envelope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Uses an existing Unity Catalog Delta table or view as the datasource's
// source. Mutually exclusive with `autoloader` and `unity_catalog_connection`.
type ExistingTable struct {
	// Name of the catalog the existing table lives in. Optional on create and
	// defaults to the datasource's own catalog. Cannot be changed. Reads return
	// the value as supplied; the catalog in use is reported in
	// `effective_catalog`.
	Catalog string `json:"catalog,omitempty"`
	// Current Unity Catalog catalog name of the source table, resolved from its
	// stable table ID, so a rename of the source table is reported here without
	// changing `catalog`.
	EffectiveCatalog string `json:"effective_catalog,omitempty"`
	// Current Unity Catalog name of the source table, resolved from its stable
	// table ID.
	EffectiveName string `json:"effective_name,omitempty"`
	// Current Unity Catalog schema name of the source table, resolved from its
	// stable table ID.
	EffectiveSchema string `json:"effective_schema,omitempty"`
	// Whether to ingest the rows already present in the source table when
	// ingestion first starts. Set true to ingest the existing rows and then
	// every new commit. Set false, or leave unset, to ingest only rows written
	// after the datasource starts. Defaults to false, and must be true for a
	// Delta Sharing source table. Takes effect only at first start and cannot
	// be changed. Reads return the value as supplied.
	IncludeExistingRows bool `json:"include_existing_rows,omitempty"`
	// Name of the source table or view. Required on create and cannot be
	// changed. Reads return the value as supplied; the table's current name is
	// reported in `effective_name`.
	Name string `json:"name"`
	// Name of the schema the existing table lives in. Optional, but required
	// when `catalog` is set; supplying `catalog` without `schema` is rejected.
	// Defaults to the datasource's own schema. Cannot be changed. Reads return
	// the value as supplied; the schema in use is reported in
	// `effective_schema`.
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExistingTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExistingTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for delivering events / alerts to an external destination.
//
// Exactly one of the destination variants is set; the server picks it via the
// `oneof`.
type ExportConfig struct {
	// Whether export is currently active.
	Enabled bool `json:"enabled,omitempty"`

	Jira *JiraExportConfig `json:"jira,omitempty"`

	ServiceNow *ServiceNowExportConfig `json:"service_now,omitempty"`

	Slack *SlackExportConfig `json:"slack,omitempty"`

	Tines *TinesExportConfig `json:"tines,omitempty"`

	Webhook *WebhookExportConfig `json:"webhook,omitempty"`
	// Deprecated and ignored: exports always egress through the request's own
	// workspace. Kept so configs stored by older builds still round-trip.
	WorkspaceId string `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExportConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAlertConfigRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetAlertConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBuiltinParserRequest struct {
	// Resource name. Format: `builtin-parsers/{parser_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetBuiltinParserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBuiltinParserVersionRequest struct {
	// Resource name of the version. Format:
	// `builtin-parsers/{parser_id}/versions/{version}`, where `{version}` is
	// either a version number or the alias `latest`.
	Name string `json:"-" url:"-"`
}

func (s *GetBuiltinParserVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCommonConfigRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetCommonConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatasourceConfigRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetDatasourceConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatasourceRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetDatasourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetParserRequest struct {
	// Resource name of the parser. Format: `parsers/{parser_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetParserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetParserVersionRequest struct {
	// Resource name of the version. Format:
	// `parsers/{parser_id}/versions/{version}`, where `{version}` is either a
	// version number or the alias `latest`. `latest` is NOT_FOUND when every
	// version of the parser has been deleted.
	Name string `json:"-" url:"-"`
}

func (s *GetParserVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRuleConfigRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetRuleConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRuleRequest struct {
	// Full resource name of the detection rule, in the form `rules/{rule_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type JiraExportConfig struct {
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Jira API token, paired with `user_email` as the Jira Cloud
	// basic-auth credential (`email:token`). Only the name is stored; the token
	// is read when an alert is exported, so rotating the secret takes effect
	// without editing this config.
	ApiTokenUcSecret string `json:"api_token_uc_secret"`
	// Default issue type for created tickets (e.g. `Task`, `Incident`). Jira
	// issue creation requires an issue type, so it is validated at config-save
	// time rather than failing the export at Jira later.
	DefaultIssueType string `json:"default_issue_type"`
	// Jira instance URL (e.g. `https://acme.atlassian.net`).
	InstanceUrl string `json:"instance_url"`
	// Jira project key the exported issues are created under.
	ProjectKey string `json:"project_key"`
	// User email paired with the API token for Jira Cloud basic auth.
	UserEmail string `json:"user_email"`
}

func (s *JiraExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Unwraps a JSON array of events into one row per array element.
type JsonArrayEnvelope struct {
	// Name of the root-level key that holds the array of events, for example
	// `Records`. Nested paths are not supported.
	Path string `json:"path"`
}

func (s *JsonArrayEnvelope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListBuiltinParserVersionsRequest struct {
	// Maximum number of results to return in one page. If unset or 0, at most
	// 100 versions are returned. The maximum allowed value is 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous `ListBuiltinParserVersions` response, used to
	// fetch the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the built-in parser whose versions are listed. Format:
	// `builtin-parsers/{parser_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBuiltinParserVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBuiltinParserVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of built-in version metadata, without the specs.
type ListBuiltinParserVersionsResponse struct {
	// Version metadata only, oldest first, without the spec.
	BuiltinParserVersions []BuiltinParserVersion `json:"builtin_parser_versions,omitempty"`
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBuiltinParserVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBuiltinParserVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBuiltinParsersRequest struct {
	// AIP-160 filter expression. Built-ins are filterable by their `tags`; the
	// exact tag-filter grammar is defined by the server.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous `ListBuiltinParsers` response, used to fetch
	// the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBuiltinParsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBuiltinParsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of built-in parsers, without their specs.
type ListBuiltinParsersResponse struct {
	// Summary rows: the container fields only, without the spec. Read a spec
	// with `GetBuiltinParserVersion`.
	BuiltinParsers []BuiltinParser `json:"builtin_parsers,omitempty"`
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty only when there are no more results. Page size is
	// best-effort, so a short or even empty page does not mean the end; keep
	// paging until this token is empty, which is the only reliable
	// end-of-results signal.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBuiltinParsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBuiltinParsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatasourcesRequest struct {
	// Name of the Unity Catalog catalog to list datasources from. Required;
	// each request lists the datasources of a single catalog.
	CatalogName string `json:"-" url:"catalog_name"`
	// Filter expression that restricts the results to datasources with a given
	// name, for example `datasource_name = "my_source"`. Only a single
	// exact-match comparison on `datasource_name` is supported, and it matches
	// the datasource's current Unity Catalog name, which reads report in
	// `effective_datasource_name`. Filtered results are returned in a single
	// page.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of a Unity Catalog schema in `catalog_name`. Optional; if set, only
	// the datasources in this schema are returned, and if omitted, the
	// datasources in all schemas of the catalog are returned.
	SchemaName string `json:"-" url:"schema_name,omitempty"`
	// Which fields to return for each datasource. Defaults to
	// `DATASOURCE_VIEW_FULL`. `DATASOURCE_VIEW_BASIC` returns only the identity
	// and Unity Catalog location fields.
	View DatasourceView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatasourcesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatasourcesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatasourcesResponse struct {
	Datasources []Datasource `json:"datasources,omitempty"`
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty only when there are no more results. Page size is
	// best-effort, so a short or even empty page does not mean the end; keep
	// paging until this token is empty, which is the only reliable
	// end-of-results signal.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatasourcesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatasourcesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListParserVersionsRequest struct {
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous `ListParserVersions` response, used to fetch
	// the next page of results. Send the same `parent` when requesting later
	// pages; a token from a request for a different parser is rejected.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the parser whose versions are listed. Format:
	// `parsers/{parser_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListParserVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListParserVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of version metadata, without the specs.
type ListParserVersionsResponse struct {
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty only when there are no more results. Page size is
	// best-effort, so a short or even empty page does not mean the end; keep
	// paging until this token is empty, which is the only reliable
	// end-of-results signal.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Version metadata only, oldest first, without the spec. Each row's
	// `create_time` and `created_by` are that version's own.
	ParserVersions []ParserVersion `json:"parser_versions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListParserVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListParserVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListParsersRequest struct {
	// Name of the parent UC catalog to list parsers from. Required: List is
	// scoped to a single catalog, mirroring the Unity Catalog tables/functions
	// APIs, which take catalog_name as a request field (bound to a query
	// parameter) rather than a catalog path segment.
	CatalogName string `json:"-" url:"catalog_name"`
	// AIP-160 filter expression. Parsers are filterable by `parser_name`.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous `ListParsers` response, used to fetch the next
	// page of results. Send the other request fields unchanged when requesting
	// later pages; a token from a request with different fields is rejected.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of the parent UC schema. Optional: when set, results are narrowed to
	// this schema; when omitted, parsers across all schemas in `catalog_name`
	// are returned.
	SchemaName string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListParsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListParsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of parsers, without their specs.
type ListParsersResponse struct {
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty only when there are no more results. Page size is
	// best-effort, so a short or even empty page does not mean the end; keep
	// paging until this token is empty, which is the only reliable
	// end-of-results signal.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Summary rows: the container fields only, without the spec. Read a spec
	// with `GetParserVersion`.
	Parsers []Parser `json:"parsers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListParsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListParsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRulesRequest struct {
	// Name of the Unity Catalog catalog to list rules from. Required.
	CatalogName string `json:"-" url:"catalog_name"`
	// Filter expression that restricts the results, such as `rule_name =
	// "my_rule"` to return only the rule with that name. The name is matched
	// against the current UC name, which reads report in `effective_rule_name`.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of rules to return. The server may return fewer. If unset
	// or 0, the server picks a default page size.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous `ListRules` response, used to fetch the next
	// page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of a Unity Catalog schema in `catalog_name`. If set, only rules in
	// this schema are returned. If unset, rules in all schemas of the catalog
	// are returned.
	SchemaName string `json:"-" url:"schema_name,omitempty"`
	// How much of each rule to return. Defaults to `RULE_VIEW_FULL`.
	// `RULE_VIEW_BASIC` returns only each rule's identity and Unity Catalog
	// location.
	View RuleView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRulesResponse struct {
	// Token to pass as `page_token` in a subsequent `ListRules` call to fetch
	// the next page, or empty when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The rules in this page of results.
	Rules []Rule `json:"rules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A customer-authored, versioned contract describing the fields of an ingested
// log source.
//
// The parser is a container: it carries the identity and a pointer to the
// newest version. Specs live on the immutable `ParserVersion` children, so one
// parser can be attached to many data sources and each attachment pins a
// version.
type Parser struct {
	// Name of the parent UC catalog, set on create. Reads return the value as
	// supplied; the current UC catalog name is reported in
	// `effective_catalog_name`.
	CatalogName string `json:"catalog_name"`
	// When the parser was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The principal that created this parser. Each version records its own
	// `created_by`.
	CreatedBy *Principal `json:"created_by,omitempty"`
	// Description of version 1. Accepted only by `CreateParser`; read a
	// version's description with `GetParserVersion`.
	Description string `json:"description,omitempty"`
	// The current name of the parser's UC catalog. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveCatalogName string `json:"effective_catalog_name,omitempty"`
	// The current name of the parser's UC schema. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveSchemaName string `json:"effective_schema_name,omitempty"`
	// Newest version of this parser. Read its spec by passing this to
	// `GetParserVersion`, or by using the `latest` alias. Absent when every
	// version has been deleted.
	LatestVersion int64 `json:"latest_version,omitempty"`
	// Resource name. Format: `parsers/{parser_id}`, where the leaf is the
	// immutable server-assigned UUID. Create-time placement names are stored
	// separately and do not affect the resource name.
	Name string `json:"name,omitempty"`
	// The principal that owns this parser. Server-set from the authenticated
	// creator, never supplied by the client. Beta reads by UUID perform no
	// authorization. Owner-gated write RPCs are outside this API's current
	// implementation scope.
	Owner *Principal `json:"owner,omitempty"`
	// The parser's server-assigned UUID. It is not currently a UC securable ID,
	// is immutable for the life of the resource, is the leaf of `name`, and is
	// returned by CreateParser.
	ParserId string `json:"parser_id,omitempty"`
	// Client-supplied parser name, stored as supplied and fixed at create time.
	// Must be a valid UC identifier. A custom parser is not currently a UC
	// securable and this field has no rename path.
	ParserName string `json:"parser_name"`
	// Name of the parent UC schema, set on create. Reads return the value as
	// supplied; the current UC schema name is reported in
	// `effective_schema_name`.
	SchemaName string `json:"schema_name"`
	// Spec of version 1, as a YAML document. Accepted only by `CreateParser`;
	// read a version's spec with `GetParserVersion`.
	Spec string `json:"spec"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Parser) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Parser) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reference to a parser that the datasource uses. Each parser defines the
// schema of the parsed table that it writes.
type ParserRef struct {
	// The table this parser's data is written to: `target_table` when supplied,
	// otherwise the table derived from the datasource's catalog and schema and
	// the parser name. Fixed once resolved, so a later parser rename does not
	// move it. Populated on get, create, and update responses with the resolved
	// table.
	EffectiveTargetTable *TableName `json:"effective_target_table,omitempty"`
	// Identifier of the parser. Stable across renames, unlike a parser name.
	Id string `json:"id"`
	// Table that this parser's data is written to. The table must be in the
	// datasource's catalog and schema. If unset, the table name is derived from
	// the parser name. Reads return the value as supplied; the table the data
	// is written to is reported in `effective_target_table`.
	TargetTable *TableName `json:"target_table,omitempty"`
	// Version of the parser to use, given as an explicit version number.
	Version string `json:"version"`
}

func (s *ParserRef) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A key/value tag on a built-in parser, e.g. `key = "cloud_provider"`, `value =
// "aws"`. Key/value pairs match the tag shape used across Databricks (MLflow,
// model-serving endpoints, clusters) rather than a bare string list.
type ParserTag struct {
	// The tag key, e.g. `cloud_provider`. Required: a keyless tag cannot be
	// filtered on and is meaningless.
	Key string `json:"key"`
	// The tag value, e.g. `aws`.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ParserTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ParserTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// One immutable version of a parser, carrying the spec.
//
// Immutable: an edit creates a new version rather than mutating an existing
// one, so a data source that pins a version is never affected. Version numbers
// are never reused, including after a delete.
type ParserVersion struct {
	// When this version was appended.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The principal that appended this version.
	CreatedBy *Principal `json:"created_by,omitempty"`
	// Free-form description of this version, set when the version is created.
	Description string `json:"description,omitempty"`
	// Resource name. Format: `parsers/{parser_id}/versions/{version}`.
	Name string `json:"name,omitempty"`
	// The parent parser's name, copied from the container on read. Denormalized
	// so a caller that needs identity alongside the spec reads the version
	// alone rather than the container as well. Populated by `GetParserVersion`;
	// List omits it, as List omits `spec`.
	ParserName string `json:"parser_name,omitempty"`
	// The field definitions for this log source, as a YAML document. Validated
	// on write, so an invalid spec is never stored.
	Spec string `json:"spec"`
	// Version number, assigned by the server when the version is appended.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ParserVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ParserVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A principal — user or service-principal — used as the value type for
// `created_by` / `updated_by` / `assignee` etc. across v2 resources.
type Principal struct {
	// Display name (user or SP name).
	DisplayName string `json:"display_name,omitempty"`
	// Email (regular users only).
	Email string `json:"email,omitempty"`
	// User / service-principal ID.
	Id string `json:"id,omitempty"`
	// Whether this principal is a regular user, service principal, or system
	// principal.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Principal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Principal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of a principal (user / SP / system).
type PrincipalType string

const PrincipalTypeRegularUser PrincipalType = `REGULAR_USER`

const PrincipalTypeServicePrincipal PrincipalType = `SERVICE_PRINCIPAL`

const PrincipalTypeSystemUser PrincipalType = `SYSTEM_USER`

// String representation for [fmt.Print]
func (f *PrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrincipalType) Set(v string) error {
	switch v {
	case `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`:
		*f = PrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "REGULAR_USER", "SERVICE_PRINCIPAL", "SYSTEM_USER"`, v)
	}
}

// Values returns all possible values for PrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeRegularUser,
		PrincipalTypeServicePrincipal,
		PrincipalTypeSystemUser,
	}
}

// Type always returns PrincipalType to satisfy [pflag.Value] interface
func (f *PrincipalType) Type() string {
	return "PrincipalType"
}

// Read-only runtime status of one job or pipeline backing a resource.
//
// Shared across resources: a Rule's single job, a Datasource's main / health /
// historical-ingestion jobs, and a Datasource's ingestion pipeline (the
// `pipeline_id` field is set only for the pipeline). For a datasource's
// historical ingestion, the shard count is simply the number of
// `historical_job_statuses` entries — there is no separate count field.
//
// Two orthogonal axes are modeled separately (v1 conflated them into a single
// `jobStatus` enum): `lifecycle` (is it configured / does it still exist) and
// `run_state` (what the most recent or in-flight run did).
//
// Every field that holds a `ResourceStatus` is OUTPUT_ONLY at its use site
// (Datasource ingest/health/historical/pipeline statuses, Rule status), so the
// inner fields don't repeat the annotation.
type ResourceStatus struct {
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Databricks Jobs API job ID. Unset for a pure ingestion pipeline.
	JobId string `json:"job_id,omitempty"`
	// When the most recent run completed. Pairs with `run_state = SUCCEEDED`.
	LastCompleteTime *time.Time `json:"last_complete_time,omitempty"`
	// Most recent error message, if any. Pairs with `run_state = FAILED`.
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// When the most recent run started. Pairs with `run_state = RUNNING`.
	LastStartTime *time.Time `json:"last_start_time,omitempty"`

	LastStatusChangeTime *time.Time `json:"last_status_change_time,omitempty"`
	// Scheduling lifecycle of the resource's underlying job/pipeline: whether
	// it is unscheduled, scheduled, stopped, or deleted. Distinct from
	// `run_state`, which reports the outcome of the most recent run.
	LifecycleState ResourceStatusLifecycle `json:"lifecycle_state,omitempty"`
	// DLT ingestion-pipeline ID. Set only when this status represents the
	// datasource's ingestion pipeline (unset for job statuses).
	PipelineId string `json:"pipeline_id,omitempty"`

	RunState ResourceStatusRunState `json:"run_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResourceStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResourceStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configured scheduling / existence of the job or pipeline.
type ResourceStatusLifecycle string

const ResourceStatusLifecycleDeleted ResourceStatusLifecycle = `DELETED`

const ResourceStatusLifecycleScheduled ResourceStatusLifecycle = `SCHEDULED`

const ResourceStatusLifecycleStopped ResourceStatusLifecycle = `STOPPED`

const ResourceStatusLifecycleUnscheduled ResourceStatusLifecycle = `UNSCHEDULED`

// String representation for [fmt.Print]
func (f *ResourceStatusLifecycle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ResourceStatusLifecycle) Set(v string) error {
	switch v {
	case `DELETED`, `SCHEDULED`, `STOPPED`, `UNSCHEDULED`:
		*f = ResourceStatusLifecycle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "SCHEDULED", "STOPPED", "UNSCHEDULED"`, v)
	}
}

// Values returns all possible values for ResourceStatusLifecycle.
//
// There is no guarantee on the order of the values in the slice.
func (f *ResourceStatusLifecycle) Values() []ResourceStatusLifecycle {
	return []ResourceStatusLifecycle{
		ResourceStatusLifecycleDeleted,
		ResourceStatusLifecycleScheduled,
		ResourceStatusLifecycleStopped,
		ResourceStatusLifecycleUnscheduled,
	}
}

// Type always returns ResourceStatusLifecycle to satisfy [pflag.Value] interface
func (f *ResourceStatusLifecycle) Type() string {
	return "ResourceStatusLifecycle"
}

// Outcome / progress of the most recent (or in-flight) run.
type ResourceStatusRunState string

const ResourceStatusRunStateCanceled ResourceStatusRunState = `CANCELED`

const ResourceStatusRunStateFailed ResourceStatusRunState = `FAILED`

const ResourceStatusRunStateRunning ResourceStatusRunState = `RUNNING`

const ResourceStatusRunStateSucceeded ResourceStatusRunState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *ResourceStatusRunState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ResourceStatusRunState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `RUNNING`, `SUCCEEDED`:
		*f = ResourceStatusRunState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for ResourceStatusRunState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ResourceStatusRunState) Values() []ResourceStatusRunState {
	return []ResourceStatusRunState{
		ResourceStatusRunStateCanceled,
		ResourceStatusRunStateFailed,
		ResourceStatusRunStateRunning,
		ResourceStatusRunStateSucceeded,
	}
}

// Type always returns ResourceStatusRunState to satisfy [pflag.Value] interface
func (f *ResourceStatusRunState) Type() string {
	return "ResourceStatusRunState"
}

// A Lakewatch detection rule: a query that evaluates ingested data and
// generates a signal for each match. Each rule is a securable in a Unity
// Catalog schema.
type Rule struct {
	// Free-form client-managed key/value annotations.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Scheduled detection (also known as batch detection) input.
	Batch *RuleInputBatch `json:"batch,omitempty"`
	// Name of the Unity Catalog catalog that contains the rule. Required.
	// Cannot be changed after creation. Reads return the value as supplied; the
	// catalog's current name is reported in `effective_catalog_name`.
	CatalogName string `json:"catalog_name"`
	// Free-form rule category.
	Category string `json:"category,omitempty"`
	// ID of the workspace whose compute runs this rule's jobs and owns its
	// associated resources, such as the SQL warehouse a scheduled detection
	// runs on. Server-assigned; a rule's warehouse can only be changed from
	// this workspace.
	ComputeWorkspaceId string `json:"compute_workspace_id,omitempty"`
	// Immutable UUIDv4 identifying the authored content resource installed as
	// this rule.
	ContentResourceId string `json:"content_resource_id,omitempty"`
	// Time the rule was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Principal that created the rule.
	CreatedBy *Principal `json:"created_by,omitempty"`
	// Signal deduplication mode.
	Deduplication RuleDeduplication `json:"deduplication,omitempty"`
	// Time window for deduplication, such as `1h`, `24h`, `7d`, or `automatic`.
	DeduplicationTimeWindow string `json:"deduplication_time_window,omitempty"`
	// Free-form description.
	Description string `json:"description,omitempty"`
	// Current name of the Unity Catalog catalog that contains the rule.
	// Reflects catalog renames, which do not change `catalog_name`. Unset until
	// resolved.
	EffectiveCatalogName string `json:"effective_catalog_name,omitempty"`
	// Current UC leaf name of the rule (server-owned). Tracks UC renames so
	// callers can observe the live name without having to update the
	// client-owned `rule_name`; unset until resolved.
	EffectiveRuleName string `json:"effective_rule_name,omitempty"`
	// The identity this rule's job runs as: the resolved `run_as_principal`, or
	// the requester's identity when it was omitted on create or cleared on
	// update. The value is a user email or service-principal application ID,
	// falling back to the numeric principal ID if that identity can no longer
	// be resolved. Populated on single-rule reads (get/create/update); list
	// responses omit it.
	EffectiveRunAsPrincipal string `json:"effective_run_as_principal,omitempty"`
	// Current name of the Unity Catalog schema that contains the rule. Reflects
	// schema renames, which do not change `schema_name`. Unset until resolved.
	EffectiveSchemaName string `json:"effective_schema_name,omitempty"`
	// Opaque entity tag for optimistic concurrency. Content-version ETags
	// change on user edits and remain stable across runtime status updates.
	Etag string `json:"etag,omitempty"`
	// Signal quality for this rule.
	Fidelity RuleFidelity `json:"fidelity"`
	// Up to 64 MITRE ATT&CK mappings, producing at most 192 flattened alert
	// tags.
	Mitre []RuleMitreMapping `json:"mitre,omitempty"`
	// Resource name, in the form `rules/{rule_id}`. Assigned by the server on
	// create; required on update.
	Name string `json:"name,omitempty"`
	// Long-form description of what this rule detects.
	Objective string `json:"objective,omitempty"`
	// Reference to the content pack this rule was originally installed from.
	OriginalContentPack *ContentPackReference `json:"original_content_pack,omitempty"`
	// What to record in each signal.
	Output RuleOutput `json:"output"`
	// Unity Catalog owner of the rule. Assigned by the server. Populated by
	// `GetRule`.
	Owner *Principal `json:"owner,omitempty"`
	// Ordered reference URLs carried from the authored rule metadata. A maximum
	// of three references of up to 512 characters each is supported. Requests
	// exceeding these limits are rejected.
	References []string `json:"references,omitempty"`
	// Unity Catalog securable UUID of the rule, and the last segment of `name`.
	// Never changes, including when the rule's catalog or schema is renamed.
	RuleId string `json:"rule_id,omitempty"`
	// Name of the rule within its Unity Catalog schema. Required on create; can
	// be updated within the same catalog and schema. Get synchronizes this
	// field with the current Unity Catalog name, also reported in
	// `effective_rule_name`.
	RuleName string `json:"rule_name"`
	// The principal whose identity this rule's job runs as, given as a user
	// email or service-principal application ID. When supplied on
	// create/update, Lakewatch resolves the principal, records it on the rule,
	// and creates the delegation to run as that principal — the requester
	// must be authorized to delegate to it, or the request is rejected. When
	// omitted on create, the requester's own identity is used. An update whose
	// mask does not include `run_as_principal` keeps the current identity, and
	// the requester must still be authorized to delegate to it; an editor who
	// is not must supply a run_as_principal they can delegate to (for example,
	// their own identity), or clear it, either of which reassigns the rule. An
	// update that only enables or disables the rule (changes `schedule.enabled`
	// and nothing else) does not require delegate permission and keeps the
	// current identity. Enabling the rule still requires the current identity
	// to have the access the rule needs to run, or the request is rejected. An
	// update whose mask includes `run_as_principal` but supplies no value
	// clears it: the rule then runs as the requester, exactly as when it is
	// omitted on create. Reads return the value as supplied; the identity the
	// rule runs as is reported in `effective_run_as_principal`.
	RunAsPrincipal string `json:"run_as_principal,omitempty"`
	// Long-form operational guidance carried from the authored rule metadata.
	// Supports up to 4096 characters. Requests exceeding this limit are
	// rejected. Supports Markdown; analysts and the investigation agent follow
	// it when triaging this rule's alerts.
	Runbook string `json:"runbook,omitempty"`
	// When the rule executes. Required for scheduled rules. A streaming rule
	// accepts only `enabled`, and an omitted schedule means enabled.
	Schedule *RuleSchedule `json:"schedule,omitempty"`
	// Name of the Unity Catalog schema that contains the rule. Required. Cannot
	// be changed after creation. Reads return the value as supplied; the
	// schema's current name is reported in `effective_schema_name`.
	SchemaName string `json:"schema_name"`
	// Free-form severity label, such as `Informational`, `Medium`, `High`, or
	// `Critical`.
	Severity string `json:"severity"`
	// Read-only lifecycle status.
	Status *ResourceStatus `json:"status,omitempty"`
	// Native streaming rule input: one source table filtered by one stateless
	// row predicate. Permanent field number; must not be renumbered.
	Streaming *RuleInputStreaming `json:"streaming,omitempty"`
	// If set, this rule's signals do not open new alerts until the given time.
	SuppressUntilTime *time.Time `json:"suppress_until_time,omitempty"`
	// Time-range configuration for scheduled detections. If unset, the
	// execution time range is extended by a fixed `15m` lookback.
	Timerange *RuleTimeRange `json:"timerange,omitempty"`
	// Time the rule was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Principal that last updated the rule.
	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Rule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Rule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Per-catalog rule-related config. Singleton. Name format:
// `catalogs/{catalog}/ruleConfig`.
type RuleConfig struct {
	// Schema name of the registered folder (in this catalog) that a new rule is
	// placed in when the request does not specify a schema. Must name a schema
	// registered as a folder. Omit to clear. Reads return the value as
	// supplied; the folder's current schema name is reported in
	// `effective_default_folder`.
	DefaultFolder string `json:"default_folder,omitempty"`
	// Default schedule for rules that don't specify one.
	DefaultSchedule *RuleSchedule `json:"default_schedule,omitempty"`
	// Default duration string for signal deduplication when a rule doesn't
	// override it — compact day/hour/minute/second components, e.g. `5d`,
	// `1h`. Reads return the value as supplied; there is no server default, so
	// it is unset until set. An empty value clears it.
	DefaultSignalCollationWindow string `json:"default_signal_collation_window,omitempty"`
	// Available detection categories a rule can be tagged with. Reads return
	// the value as supplied. When empty, rules can be tagged with the default
	// categories: `Static Signature`, `Fuzzy Signature`, `Behavioral`,
	// `Statistical`, and `Learning (AI/ML)`. The categories in effect are
	// reported in `effective_detection_categories`.
	DetectionCategories []string `json:"detection_categories,omitempty"`
	// The current schema name of the `default_folder` folder. Tracks UC schema
	// renames. Unset when no default is configured, or when the folder's schema
	// no longer exists or is not visible to the caller.
	EffectiveDefaultFolder string `json:"effective_default_folder,omitempty"`
	// The detection categories in effect: `detection_categories` when set,
	// otherwise the default categories.
	EffectiveDetectionCategories []string `json:"effective_detection_categories,omitempty"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag string `json:"etag,omitempty"`

	Name string `json:"name,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	UpdatedBy *Principal `json:"updated_by,omitempty"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Signal deduplication mode for a detection rule.
type RuleDeduplication string

const RuleDeduplicationDeduplicationDefault RuleDeduplication = `DEDUPLICATION_DEFAULT`

const RuleDeduplicationDeduplicationDisabled RuleDeduplication = `DEDUPLICATION_DISABLED`

const RuleDeduplicationDeduplicationEnabled RuleDeduplication = `DEDUPLICATION_ENABLED`

// String representation for [fmt.Print]
func (f *RuleDeduplication) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RuleDeduplication) Set(v string) error {
	switch v {
	case `DEDUPLICATION_DEFAULT`, `DEDUPLICATION_DISABLED`, `DEDUPLICATION_ENABLED`:
		*f = RuleDeduplication(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEDUPLICATION_DEFAULT", "DEDUPLICATION_DISABLED", "DEDUPLICATION_ENABLED"`, v)
	}
}

// Values returns all possible values for RuleDeduplication.
//
// There is no guarantee on the order of the values in the slice.
func (f *RuleDeduplication) Values() []RuleDeduplication {
	return []RuleDeduplication{
		RuleDeduplicationDeduplicationDefault,
		RuleDeduplicationDeduplicationDisabled,
		RuleDeduplicationDeduplicationEnabled,
	}
}

// Type always returns RuleDeduplication to satisfy [pflag.Value] interface
func (f *RuleDeduplication) Type() string {
	return "RuleDeduplication"
}

// Maturity of a detection rule.
type RuleFidelity string

const RuleFidelityFidelityDevelopment RuleFidelity = `FIDELITY_DEVELOPMENT`

const RuleFidelityFidelityHigh RuleFidelity = `FIDELITY_HIGH`

const RuleFidelityFidelityInvestigative RuleFidelity = `FIDELITY_INVESTIGATIVE`

const RuleFidelityFidelityMedium RuleFidelity = `FIDELITY_MEDIUM`

// String representation for [fmt.Print]
func (f *RuleFidelity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RuleFidelity) Set(v string) error {
	switch v {
	case `FIDELITY_DEVELOPMENT`, `FIDELITY_HIGH`, `FIDELITY_INVESTIGATIVE`, `FIDELITY_MEDIUM`:
		*f = RuleFidelity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FIDELITY_DEVELOPMENT", "FIDELITY_HIGH", "FIDELITY_INVESTIGATIVE", "FIDELITY_MEDIUM"`, v)
	}
}

// Values returns all possible values for RuleFidelity.
//
// There is no guarantee on the order of the values in the slice.
func (f *RuleFidelity) Values() []RuleFidelity {
	return []RuleFidelity{
		RuleFidelityFidelityDevelopment,
		RuleFidelityFidelityHigh,
		RuleFidelityFidelityInvestigative,
		RuleFidelityFidelityMedium,
	}
}

// Type always returns RuleFidelity to satisfy [pflag.Value] interface
func (f *RuleFidelity) Type() string {
	return "RuleFidelity"
}

// Input for a scheduled detection (also known as batch detection).
type RuleInputBatch struct {
	// Custom-notebook user function. The output is used as input data for the
	// rule.
	Custom *RuleInputCustomNotebook `json:"custom,omitempty"`
	// SQL query evaluated as the rule input.
	Sql *RuleInputBatchSql `json:"sql,omitempty"`
	// SQL warehouse that runs this scheduled detection.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleInputBatch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleInputBatch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SQL query input for a scheduled detection.
type RuleInputBatchSql struct {
	// Default catalog used to resolve unqualified identifiers in `query`. If
	// unset, the rule's `catalog_name` is used.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// Default schema used to resolve unqualified identifiers in `query`. If
	// unset, the rule's `schema_name` is used.
	DefaultSchemaName string `json:"default_schema_name,omitempty"`
	// SQL query evaluated as the rule input.
	Query string `json:"query"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleInputBatchSql) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleInputBatchSql) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reference to a custom-user-function notebook for scheduled detection input.
type RuleInputCustomNotebook struct {
	// Fully-qualified notebook path, or a path relative to the configured base
	// notebook directory.
	Notebook string `json:"notebook"`
	// Optional arguments supplied when invoking the custom function.
	Options map[string]string `json:"options,omitempty"`
}

func (s *RuleInputCustomNotebook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Native V2 streaming input for a detection rule: one Delta table and one
// stateless row predicate. The complete source row is preserved as event
// context.
type RuleInputStreaming struct {
	// Row predicate that selects matching rows.
	Predicate StreamingPredicate `json:"predicate"`
	// Source table whose rows are evaluated.
	Source StreamingSource `json:"source"`
}

func (s *RuleInputStreaming) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// MITRE ATT&CK mapping for a detection rule.
type RuleMitreMapping struct {
	// Human-readable MITRE sub-technique name.
	SubTechnique string `json:"sub_technique,omitempty"`
	// MITRE sub-technique ID, such as `T1078.001`.
	SubTechniqueId string `json:"sub_technique_id,omitempty"`
	// Adversary objective during this attack, such as `Command and Control`.
	Tactic string `json:"tactic,omitempty"`
	// Structured framework that organizes adversary behaviors, such as
	// `mitre-attack`.
	Taxonomy string `json:"taxonomy,omitempty"`
	// Human-readable MITRE technique name.
	Technique string `json:"technique,omitempty"`
	// MITRE technique ID, such as `T1078`.
	TechniqueId string `json:"technique_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleMitreMapping) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleMitreMapping) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// What to record in each signal that this rule generates.
type RuleOutput struct {
	// Key/value context to record from the input row. Values may contain SQL
	// expressions in braces.
	Context map[string]string `json:"context,omitempty"`
	// If true, also record every column of the triggering input row. Custom
	// context wins on key collision.
	DefaultContext bool `json:"default_context,omitempty"`
	// Column in the rule output that holds the event timestamp.
	EventTimeColumn string `json:"event_time_column,omitempty"`
	// Signal summary. May contain SQL expressions in braces that the server
	// substitutes from the input row.
	Summary string `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Defines when a detection rule executes.
//
// A scheduled detection sets exactly one of `at_least_every` or
// `cron_expression`. A streaming rule runs continuously and sets no timing; it
// accepts only `enabled`, and an omitted schedule means enabled. The server
// rejects any other combination. `enabled` is orthogonal to the timing.
type RuleSchedule struct {
	// Defines the minimum cadence as a duration string. It may contain
	// components for days, hours, minutes, and seconds (e.g. `5d`, `5d3h`,
	// `3d2h5m12s`, `12m16s`).
	AtLeastEvery string `json:"at_least_every,omitempty"`
	// Defines the schedule using a Quartz cron expression.
	CronExpression string `json:"cron_expression,omitempty"`
	// Sets whether the job is scheduled to run. Defaults to true.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleSchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Time-range configuration for a scheduled detection.
type RuleTimeRange struct {
	// Extends the start of the execution time range by a compact duration with
	// integer days, hours, minutes, and seconds in that order, such as `1h30m`
	// or `15m`. Zero duration disables extension. Unset or empty uses a fixed
	// `15m` at execution time. Only the start moves; the end is unchanged. Has
	// no effect on a detection query that does not use `IN_TIMERANGE`. Values
	// must fit signed 64-bit microseconds.
	ExtendLookback string `json:"extend_lookback,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RuleTimeRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleTimeRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Selects how much of each rule a list response returns.
type RuleView string

const RuleViewRuleViewBasic RuleView = `RULE_VIEW_BASIC`

const RuleViewRuleViewFull RuleView = `RULE_VIEW_FULL`

// String representation for [fmt.Print]
func (f *RuleView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RuleView) Set(v string) error {
	switch v {
	case `RULE_VIEW_BASIC`, `RULE_VIEW_FULL`:
		*f = RuleView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RULE_VIEW_BASIC", "RULE_VIEW_FULL"`, v)
	}
}

// Values returns all possible values for RuleView.
//
// There is no guarantee on the order of the values in the slice.
func (f *RuleView) Values() []RuleView {
	return []RuleView{
		RuleViewRuleViewBasic,
		RuleViewRuleViewFull,
	}
}

// Type always returns RuleView to satisfy [pflag.Value] interface
func (f *RuleView) Type() string {
	return "RuleView"
}

// Defines when a recurring job executes.
//
// Exactly one of the timing modes (`at_least_every`, `cron_expression`,
// `continuous`, `file_arrival`) must be set; combining them or setting none is
// rejected by the server. `compute_group` and `enabled` are orthogonal to the
// timing.
type Schedule struct {
	// Duration string defining the minimum cadence. May contain components for
	// days, hours, minutes, seconds (e.g. `5d`, `5d3h`, `3d2h5m12s`, `12m16s`).
	AtLeastEvery string `json:"at_least_every,omitempty"`
	// `dedicated`, `inherited`, `automatic`, or a custom group name. Reads
	// return the value as supplied. A rule whose schedule omits it runs on
	// `dedicated`.
	ComputeGroup string `json:"compute_group,omitempty"`
	// Run continuously (sentinel — must be `true` when set in this branch).
	Continuous bool `json:"continuous,omitempty"`
	// Quartz cron expression.
	CronExpression string `json:"cron_expression,omitempty"`
	// Whether the job is scheduled to run. Defaults to true.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ServiceNow destination. Alerts are delivered to the instance's Table API,
// authorized by an OAuth 2.0 client-credentials token minted from the
// registered application's client ID and secret.
type ServiceNowExportConfig struct {
	// OAuth client ID of the registered ServiceNow application. An identifier
	// rather than secret material — ServiceNow shows it on the application
	// record — so it is stored and read back verbatim.
	ClientId string `json:"client_id"`
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the OAuth client secret. Only the name is stored; the secret is
	// read when an alert is exported, so rotating it takes effect without
	// editing this config.
	ClientSecretUcSecret string `json:"client_secret_uc_secret"`
	// ServiceNow instance ID — the host label only, not the full URL. For an
	// instance reachable at `https://acme.service-now.com`, this is `acme`.
	InstanceId string `json:"instance_id"`
}

func (s *ServiceNowExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SlackExportConfig struct {
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Slack bot token. Only the name is stored; the token is read
	// when an alert is exported, so rotating the secret takes effect without
	// editing this config.
	BotTokenUcSecret string `json:"bot_token_uc_secret"`
	// Channel the alert is posted to: a channel ID (e.g. `C0123456789`) or a
	// channel name without the leading `#` (e.g. `soc-alerts`). The value is
	// stored as sent, so a leading `#`, surrounding whitespace, or a value
	// naming more than one channel is rejected. For private channels, prefer
	// the channel ID — Slack does not guarantee name resolution for those.
	//
	// The Slack app must be a member of the channel and hold the `chat:write`
	// scope.
	Channel string `json:"channel"`
}

func (s *SlackExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Stateless row predicate for a streaming detection rule.
type StreamingPredicate struct {
	// A SQL expression suitable for a WHERE clause, not a SELECT query.
	// Evaluated per input row; only stateless, row-local expressions are
	// accepted.
	Expression string `json:"expression"`
}

func (s *StreamingPredicate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Source Delta table for a streaming detection rule.
type StreamingSource struct {
	// Fully-qualified Unity Catalog table in `catalog.schema.table` form.
	Table string `json:"table"`
}

func (s *StreamingSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A fully-qualified Unity Catalog table reference by catalog / schema / table
// name.
type TableName struct {
	// Name of the parent Unity Catalog catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// Name of the parent Unity Catalog schema.
	SchemaName string `json:"schema_name,omitempty"`
	// Unity Catalog table name.
	TableName string `json:"table_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableName) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Tines destination. Alerts are delivered to a Tines webhook action, authorized
// by a secret sent in the `Authorization` header.
type TinesExportConfig struct {
	// Whether a later change to an already-exported alert is delivered too, as
	// a second `lakewatch.alert.updated` event to the same webhook action (a
	// close is such an update, carrying a terminal `alert_state`). Creation
	// export is unaffected either way.
	//
	// Opt-in, and unset means off: Tines is an event transport rather than a
	// system whose record is edited in place, so an update is an extra event
	// the receiving story has to correlate to the alert it already handled. A
	// destination configured before this event existed must not start receiving
	// more without its owner asking.
	ExportUpdates bool `json:"export_updates,omitempty"`
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Tines webhook secret. Only the name is stored; the secret is
	// read when an alert is exported, so rotating it takes effect without
	// editing this config. Required: without it the destination would be an
	// unauthenticated endpoint.
	WebhookSecretUcSecret string `json:"webhook_secret_uc_secret"`
	// Tines webhook action URL, typically of the form
	// `https://<tenant>.tines.com/webhook/<path>`. Any HTTPS host is accepted,
	// so that a Tines tenant on a custom domain works.
	//
	// The secret is sent as a header, so it must not be embedded in the URL: a
	// trailing secret path segment (`/webhook/<path>/<secret>`) and a `secret`
	// query parameter are both rejected. API-path-prefix URLs
	// (`/api/public/...`) are not supported.
	WebhookUri string `json:"webhook_uri"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TinesExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TinesExportConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Ingests data through a managed connector that uses a Unity Catalog
// connection. Mutually exclusive with `autoloader` and `existing_table`.
type UnityCatalogConnection struct {
	// Free-form key-value configuration for the managed ingestion pipeline.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Name of the Unity Catalog connection to ingest through.
	ConnectionName string `json:"connection_name"`
	// Connector type of the connection, such as `SLACK_AUDIT_LOGS`. Resolved by
	// the server from `connection_name` through Unity Catalog; not accepted as
	// input.
	Connector string `json:"connector,omitempty"`
	// The start time ingestion uses: the supplied `start_time`, or the creation
	// time when it was omitted. Always populated.
	EffectiveStartTime *time.Time `json:"effective_start_time,omitempty"`
	// Earliest source event time to ingest. Leave unset to start at creation
	// time, with no history backfilled. Rejected if in the future, beyond the
	// connector's history window, or for a connector without start-time
	// support. Takes effect only at first start and cannot be changed. Reads
	// return the value as supplied; the start time ingestion uses is reported
	// in `effective_start_time`.
	StartTime *time.Time `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UnityCatalogConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UnityCatalogConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateAlertConfigRequest struct {
	AlertConfig AlertConfig `json:"alert_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag string `json:"-" url:"etag,omitempty"`

	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAlertConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAlertConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCommonConfigRequest struct {
	CommonConfig CommonConfig `json:"common_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag string `json:"-" url:"etag,omitempty"`

	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCommonConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCommonConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateDatasourceConfigRequest struct {
	DatasourceConfig DatasourceConfig `json:"datasource_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag string `json:"-" url:"etag,omitempty"`

	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateDatasourceConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDatasourceConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateDatasourceRequest struct {
	Datasource Datasource `json:"datasource"`
	// Optimistic-concurrency precondition: if set to a non-empty value that
	// does not match the resource's current `etag`, the update is rejected as a
	// conflict (HTTP 409). Omit this field or send an empty string to update
	// unconditionally.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name of the datasource, in the format
	// `datasources/{datasource_id}`. Unlike the datasource's Unity Catalog
	// location (`catalog_name`, `schema_name`, and `datasource_name`), the name
	// does not change when the datasource or its catalog or schema is renamed
	// in Unity Catalog.
	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateDatasourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDatasourceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRuleConfigRequest struct {
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag string `json:"-" url:"etag,omitempty"`

	Name string `json:"-" url:"-"`

	RuleConfig RuleConfig `json:"rule_config"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRuleConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRuleConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRuleRequest struct {
	// Optimistic-concurrency token from the latest read. A non-empty token that
	// does not match the current `etag` rejects the update as a conflict (HTTP
	// 409). Omit this field or provide an empty string to update
	// unconditionally.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name, in the form `rules/{rule_id}`. Assigned by the server on
	// create; required on update.
	Name string `json:"-" url:"-"`
	// The rule with its new field values. `rule.name` identifies the rule to
	// update.
	Rule Rule `json:"rule"`
	// The set of fields to update. Only fields named in the mask are modified.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRuleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRuleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WebhookExportConfig struct {
	// Whether a later change to an already-exported alert is delivered too, as
	// a second `lakewatch.alert.updated` event to the same webhook URL (a close
	// is such an update, carrying a terminal `alert_state`). Creation export is
	// unaffected either way.
	//
	// Opt-in, and unset means off: a webhook is an event transport rather than
	// a system whose record is edited in place, so an update is an extra event
	// the receiving consumer has to correlate to the alert it already handled.
	// A destination configured before this event existed must not start
	// receiving more without its owner asking.
	ExportUpdates bool `json:"export_updates,omitempty"`
	// Webhook destination the server POSTs payloads to, in plaintext. Userinfo
	// (`user:password@`) is sent as HTTP Basic auth.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WebhookExportConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WebhookExportConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
