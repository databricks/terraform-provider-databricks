// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package lakewatch_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Per-catalog alert-related config. Singleton. Name format:
// `catalogs/{catalog}/alertConfig`.
type AlertConfig_SdkV2 struct {
	// Window for collating Signals into security alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1h`; the
	// window in effect is reported in
	// `effective_default_security_alert_collation_window`.
	DefaultSecurityAlertCollationWindow types.String `tfsdk:"default_security_alert_collation_window"`
	// Window for collating SystemEvents into system alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1d`; the
	// window in effect is reported in
	// `effective_default_system_alert_collation_window`.
	DefaultSystemAlertCollationWindow types.String `tfsdk:"default_system_alert_collation_window"`
	// The security alert collation window in effect:
	// `default_security_alert_collation_window` when set, otherwise `1h`.
	EffectiveDefaultSecurityAlertCollationWindow types.String `tfsdk:"effective_default_security_alert_collation_window"`
	// The system alert collation window in effect:
	// `default_system_alert_collation_window` when set, otherwise `1d`.
	EffectiveDefaultSystemAlertCollationWindow types.String `tfsdk:"effective_default_system_alert_collation_window"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag types.String `tfsdk:"etag"`

	Name types.String `tfsdk:"name"`
	// Export destination for security alerts. When unset, export is off and
	// reads return `{"enabled": false}`.
	SecurityAlertExportConfig types.List `tfsdk:"security_alert_export_config"`
	// Export destination for system alerts. When unset, export is off and reads
	// return `{"enabled": false}`.
	SystemAlertExportConfig types.List `tfsdk:"system_alert_export_config"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.List `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version types.Int64 `tfsdk:"version"`
}

func (to *AlertConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConfig_SdkV2) {
	if !from.SecurityAlertExportConfig.IsNull() && !from.SecurityAlertExportConfig.IsUnknown() {
		if toSecurityAlertExportConfig, ok := to.GetSecurityAlertExportConfig(ctx); ok {
			if fromSecurityAlertExportConfig, ok := from.GetSecurityAlertExportConfig(ctx); ok {
				// Recursively sync the fields of SecurityAlertExportConfig
				toSecurityAlertExportConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromSecurityAlertExportConfig)
				to.SetSecurityAlertExportConfig(ctx, toSecurityAlertExportConfig)
			}
		}
	}
	if !from.SystemAlertExportConfig.IsNull() && !from.SystemAlertExportConfig.IsUnknown() {
		if toSystemAlertExportConfig, ok := to.GetSystemAlertExportConfig(ctx); ok {
			if fromSystemAlertExportConfig, ok := from.GetSystemAlertExportConfig(ctx); ok {
				// Recursively sync the fields of SystemAlertExportConfig
				toSystemAlertExportConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromSystemAlertExportConfig)
				to.SetSystemAlertExportConfig(ctx, toSystemAlertExportConfig)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (to *AlertConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertConfig_SdkV2) {
	if !from.SecurityAlertExportConfig.IsNull() && !from.SecurityAlertExportConfig.IsUnknown() {
		if toSecurityAlertExportConfig, ok := to.GetSecurityAlertExportConfig(ctx); ok {
			if fromSecurityAlertExportConfig, ok := from.GetSecurityAlertExportConfig(ctx); ok {
				toSecurityAlertExportConfig.SyncFieldsDuringRead(ctx, fromSecurityAlertExportConfig)
				to.SetSecurityAlertExportConfig(ctx, toSecurityAlertExportConfig)
			}
		}
	}
	if !from.SystemAlertExportConfig.IsNull() && !from.SystemAlertExportConfig.IsUnknown() {
		if toSystemAlertExportConfig, ok := to.GetSystemAlertExportConfig(ctx); ok {
			if fromSystemAlertExportConfig, ok := from.GetSystemAlertExportConfig(ctx); ok {
				toSystemAlertExportConfig.SyncFieldsDuringRead(ctx, fromSystemAlertExportConfig)
				to.SetSystemAlertExportConfig(ctx, toSystemAlertExportConfig)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (m AlertConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_security_alert_collation_window"] = attrs["default_security_alert_collation_window"].SetOptional()
	attrs["default_system_alert_collation_window"] = attrs["default_system_alert_collation_window"].SetOptional()
	attrs["effective_default_security_alert_collation_window"] = attrs["effective_default_security_alert_collation_window"].SetComputed()
	attrs["effective_default_system_alert_collation_window"] = attrs["effective_default_system_alert_collation_window"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["security_alert_export_config"] = attrs["security_alert_export_config"].SetOptional()
	attrs["security_alert_export_config"] = attrs["security_alert_export_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["system_alert_export_config"] = attrs["system_alert_export_config"].SetOptional()
	attrs["system_alert_export_config"] = attrs["system_alert_export_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AlertConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"security_alert_export_config": reflect.TypeOf(ExportConfig_SdkV2{}),
		"system_alert_export_config":   reflect.TypeOf(ExportConfig_SdkV2{}),
		"updated_by":                   reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_security_alert_collation_window":           m.DefaultSecurityAlertCollationWindow,
			"default_system_alert_collation_window":             m.DefaultSystemAlertCollationWindow,
			"effective_default_security_alert_collation_window": m.EffectiveDefaultSecurityAlertCollationWindow,
			"effective_default_system_alert_collation_window":   m.EffectiveDefaultSystemAlertCollationWindow,
			"etag":                         m.Etag,
			"name":                         m.Name,
			"security_alert_export_config": m.SecurityAlertExportConfig,
			"system_alert_export_config":   m.SystemAlertExportConfig,
			"update_time":                  m.UpdateTime,
			"updated_by":                   m.UpdatedBy,
			"version":                      m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_security_alert_collation_window":           types.StringType,
			"default_system_alert_collation_window":             types.StringType,
			"effective_default_security_alert_collation_window": types.StringType,
			"effective_default_system_alert_collation_window":   types.StringType,
			"etag": types.StringType,
			"name": types.StringType,
			"security_alert_export_config": basetypes.ListType{
				ElemType: ExportConfig_SdkV2{}.Type(ctx),
			},
			"system_alert_export_config": basetypes.ListType{
				ElemType: ExportConfig_SdkV2{}.Type(ctx),
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetSecurityAlertExportConfig returns the value of the SecurityAlertExportConfig field in AlertConfig_SdkV2 as
// a ExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig_SdkV2) GetSecurityAlertExportConfig(ctx context.Context) (ExportConfig_SdkV2, bool) {
	var e ExportConfig_SdkV2
	if m.SecurityAlertExportConfig.IsNull() || m.SecurityAlertExportConfig.IsUnknown() {
		return e, false
	}
	var v []ExportConfig_SdkV2
	d := m.SecurityAlertExportConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSecurityAlertExportConfig sets the value of the SecurityAlertExportConfig field in AlertConfig_SdkV2.
func (m *AlertConfig_SdkV2) SetSecurityAlertExportConfig(ctx context.Context, v ExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["security_alert_export_config"]
	m.SecurityAlertExportConfig = types.ListValueMust(t, vs)
}

// GetSystemAlertExportConfig returns the value of the SystemAlertExportConfig field in AlertConfig_SdkV2 as
// a ExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig_SdkV2) GetSystemAlertExportConfig(ctx context.Context) (ExportConfig_SdkV2, bool) {
	var e ExportConfig_SdkV2
	if m.SystemAlertExportConfig.IsNull() || m.SystemAlertExportConfig.IsUnknown() {
		return e, false
	}
	var v []ExportConfig_SdkV2
	d := m.SystemAlertExportConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSystemAlertExportConfig sets the value of the SystemAlertExportConfig field in AlertConfig_SdkV2.
func (m *AlertConfig_SdkV2) SetSystemAlertExportConfig(ctx context.Context, v ExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["system_alert_export_config"]
	m.SystemAlertExportConfig = types.ListValueMust(t, vs)
}

// GetUpdatedBy returns the value of the UpdatedBy field in AlertConfig_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in AlertConfig_SdkV2.
func (m *AlertConfig_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Ingests files from cloud storage with Auto Loader. Mutually exclusive with
// `unity_catalog_connection` and `existing_table`.
type AutoLoader_SdkV2 struct {
	// Name of the workspace secret scope that holds the AWS access key ID and
	// secret access key, under the keys `awsAccessKeyId` and
	// `awsSecretAccessKey`.
	AwsSecretScope types.String `tfsdk:"aws_secret_scope"`
	// Whether ingestion started with the files already present at `location`.
	// Populated on get, create, and update responses.
	EffectiveIncludeExistingFiles types.Bool `tfsdk:"effective_include_existing_files"`
	// The table that the raw data is written to: `raw_table` in the
	// datasource's catalog and schema, under their current Unity Catalog names.
	// Populated on get, create, and update responses.
	EffectiveRawTable types.List `tfsdk:"effective_raw_table"`
	// How to unwrap a transport envelope around the source records before
	// parsing. Supported only for the JSON and JSONL formats. If unset, each
	// source record is treated as one event.
	Envelope types.List `tfsdk:"envelope"`
	// File format of the source files. Required.
	Format types.String `tfsdk:"format"`
	// Format-specific parsing options, passed to Auto Loader. Supported for the
	// CSV, JSON, JSONL, and XML formats, each of which accepts a fixed set of
	// options. The XML format requires the `rowTag` option.
	FormatOptions types.Map `tfsdk:"format_options"`
	// Whether to ingest the files already present at `location` when ingestion
	// first starts. Set true to ingest the existing files and then every new
	// file. Set false, or leave unset, to ingest only files that arrive after
	// the datasource starts. Defaults to false. Cannot be changed after
	// creation. Reads return the value as supplied; the setting in effect is
	// reported in `effective_include_existing_files`.
	IncludeExistingFiles types.Bool `tfsdk:"include_existing_files"`
	// URI of the source files: a path in a Unity Catalog volume or external
	// location, or a cloud storage URL such as an Amazon S3 path. Required,
	// including when new files are discovered through `queue_url`. Cannot be
	// changed after creation.
	Location types.String `tfsdk:"location"`
	// Hive-style partition columns to derive from the source file paths. If
	// unset, no partition columns are derived.
	OverridePartitionColumns types.List `tfsdk:"override_partition_columns"`
	// URL of a customer-managed Amazon SQS queue that receives file-arrival
	// notifications. Set it to discover new files through the queue. If unset,
	// new files are discovered through the file events of the source's Unity
	// Catalog external location when it has them configured, and by directory
	// listing otherwise. Cannot be changed after creation.
	QueueUrl types.String `tfsdk:"queue_url"`
	// Table that stores the raw data before parsers schematize it. Every Auto
	// Loader datasource writes its raw rows to this table before any parser
	// reads them. The table must be in the datasource's catalog and schema; an
	// omitted `catalog_name` or `schema_name` defaults to them. Cannot be
	// changed after creation. Reads return the value as supplied; the table the
	// data is written to is reported in `effective_raw_table`.
	RawTable types.List `tfsdk:"raw_table"`
	// Name of the Unity Catalog service credential to use.
	ServiceCredentials types.String `tfsdk:"service_credentials"`
}

func (to *AutoLoader_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoLoader_SdkV2) {
	if !from.EffectiveRawTable.IsNull() && !from.EffectiveRawTable.IsUnknown() {
		if toEffectiveRawTable, ok := to.GetEffectiveRawTable(ctx); ok {
			if fromEffectiveRawTable, ok := from.GetEffectiveRawTable(ctx); ok {
				// Recursively sync the fields of EffectiveRawTable
				toEffectiveRawTable.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveRawTable)
				to.SetEffectiveRawTable(ctx, toEffectiveRawTable)
			}
		}
	}
	if !from.Envelope.IsNull() && !from.Envelope.IsUnknown() {
		if toEnvelope, ok := to.GetEnvelope(ctx); ok {
			if fromEnvelope, ok := from.GetEnvelope(ctx); ok {
				// Recursively sync the fields of Envelope
				toEnvelope.SyncFieldsDuringCreateOrUpdate(ctx, fromEnvelope)
				to.SetEnvelope(ctx, toEnvelope)
			}
		}
	}
	if !from.OverridePartitionColumns.IsNull() && !from.OverridePartitionColumns.IsUnknown() && to.OverridePartitionColumns.IsNull() && len(from.OverridePartitionColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OverridePartitionColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OverridePartitionColumns = from.OverridePartitionColumns
	}
	if !from.RawTable.IsNull() && !from.RawTable.IsUnknown() {
		if toRawTable, ok := to.GetRawTable(ctx); ok {
			if fromRawTable, ok := from.GetRawTable(ctx); ok {
				// Recursively sync the fields of RawTable
				toRawTable.SyncFieldsDuringCreateOrUpdate(ctx, fromRawTable)
				to.SetRawTable(ctx, toRawTable)
			}
		}
	}
}

func (to *AutoLoader_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AutoLoader_SdkV2) {
	if !from.EffectiveRawTable.IsNull() && !from.EffectiveRawTable.IsUnknown() {
		if toEffectiveRawTable, ok := to.GetEffectiveRawTable(ctx); ok {
			if fromEffectiveRawTable, ok := from.GetEffectiveRawTable(ctx); ok {
				toEffectiveRawTable.SyncFieldsDuringRead(ctx, fromEffectiveRawTable)
				to.SetEffectiveRawTable(ctx, toEffectiveRawTable)
			}
		}
	}
	if !from.Envelope.IsNull() && !from.Envelope.IsUnknown() {
		if toEnvelope, ok := to.GetEnvelope(ctx); ok {
			if fromEnvelope, ok := from.GetEnvelope(ctx); ok {
				toEnvelope.SyncFieldsDuringRead(ctx, fromEnvelope)
				to.SetEnvelope(ctx, toEnvelope)
			}
		}
	}
	if !from.OverridePartitionColumns.IsNull() && !from.OverridePartitionColumns.IsUnknown() && to.OverridePartitionColumns.IsNull() && len(from.OverridePartitionColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OverridePartitionColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OverridePartitionColumns = from.OverridePartitionColumns
	}
	if !from.RawTable.IsNull() && !from.RawTable.IsUnknown() {
		if toRawTable, ok := to.GetRawTable(ctx); ok {
			if fromRawTable, ok := from.GetRawTable(ctx); ok {
				toRawTable.SyncFieldsDuringRead(ctx, fromRawTable)
				to.SetRawTable(ctx, toRawTable)
			}
		}
	}
}

func (m AutoLoader_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_secret_scope"] = attrs["aws_secret_scope"].SetOptional()
	attrs["effective_include_existing_files"] = attrs["effective_include_existing_files"].SetComputed()
	attrs["effective_raw_table"] = attrs["effective_raw_table"].SetComputed()
	attrs["effective_raw_table"] = attrs["effective_raw_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["envelope"] = attrs["envelope"].SetOptional()
	attrs["envelope"] = attrs["envelope"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["format"] = attrs["format"].SetRequired()
	attrs["format_options"] = attrs["format_options"].SetOptional()
	attrs["include_existing_files"] = attrs["include_existing_files"].SetOptional()
	attrs["location"] = attrs["location"].SetRequired()
	attrs["override_partition_columns"] = attrs["override_partition_columns"].SetOptional()
	attrs["queue_url"] = attrs["queue_url"].SetOptional()
	attrs["raw_table"] = attrs["raw_table"].SetRequired()
	attrs["raw_table"] = attrs["raw_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["service_credentials"] = attrs["service_credentials"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoLoader.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutoLoader_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_raw_table":        reflect.TypeOf(TableName_SdkV2{}),
		"envelope":                   reflect.TypeOf(Envelope_SdkV2{}),
		"format_options":             reflect.TypeOf(types.String{}),
		"override_partition_columns": reflect.TypeOf(types.String{}),
		"raw_table":                  reflect.TypeOf(TableName_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoLoader_SdkV2
// only implements ToObjectValue() and Type().
func (m AutoLoader_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_secret_scope":                 m.AwsSecretScope,
			"effective_include_existing_files": m.EffectiveIncludeExistingFiles,
			"effective_raw_table":              m.EffectiveRawTable,
			"envelope":                         m.Envelope,
			"format":                           m.Format,
			"format_options":                   m.FormatOptions,
			"include_existing_files":           m.IncludeExistingFiles,
			"location":                         m.Location,
			"override_partition_columns":       m.OverridePartitionColumns,
			"queue_url":                        m.QueueUrl,
			"raw_table":                        m.RawTable,
			"service_credentials":              m.ServiceCredentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutoLoader_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_secret_scope":                 types.StringType,
			"effective_include_existing_files": types.BoolType,
			"effective_raw_table": basetypes.ListType{
				ElemType: TableName_SdkV2{}.Type(ctx),
			},
			"envelope": basetypes.ListType{
				ElemType: Envelope_SdkV2{}.Type(ctx),
			},
			"format": types.StringType,
			"format_options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"include_existing_files": types.BoolType,
			"location":               types.StringType,
			"override_partition_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"queue_url": types.StringType,
			"raw_table": basetypes.ListType{
				ElemType: TableName_SdkV2{}.Type(ctx),
			},
			"service_credentials": types.StringType,
		},
	}
}

// GetEffectiveRawTable returns the value of the EffectiveRawTable field in AutoLoader_SdkV2 as
// a TableName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoLoader_SdkV2) GetEffectiveRawTable(ctx context.Context) (TableName_SdkV2, bool) {
	var e TableName_SdkV2
	if m.EffectiveRawTable.IsNull() || m.EffectiveRawTable.IsUnknown() {
		return e, false
	}
	var v []TableName_SdkV2
	d := m.EffectiveRawTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveRawTable sets the value of the EffectiveRawTable field in AutoLoader_SdkV2.
func (m *AutoLoader_SdkV2) SetEffectiveRawTable(ctx context.Context, v TableName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_raw_table"]
	m.EffectiveRawTable = types.ListValueMust(t, vs)
}

// GetEnvelope returns the value of the Envelope field in AutoLoader_SdkV2 as
// a Envelope_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoLoader_SdkV2) GetEnvelope(ctx context.Context) (Envelope_SdkV2, bool) {
	var e Envelope_SdkV2
	if m.Envelope.IsNull() || m.Envelope.IsUnknown() {
		return e, false
	}
	var v []Envelope_SdkV2
	d := m.Envelope.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvelope sets the value of the Envelope field in AutoLoader_SdkV2.
func (m *AutoLoader_SdkV2) SetEnvelope(ctx context.Context, v Envelope_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["envelope"]
	m.Envelope = types.ListValueMust(t, vs)
}

// GetFormatOptions returns the value of the FormatOptions field in AutoLoader_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoLoader_SdkV2) GetFormatOptions(ctx context.Context) (map[string]types.String, bool) {
	if m.FormatOptions.IsNull() || m.FormatOptions.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.FormatOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFormatOptions sets the value of the FormatOptions field in AutoLoader_SdkV2.
func (m *AutoLoader_SdkV2) SetFormatOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["format_options"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FormatOptions = types.MapValueMust(t, vs)
}

// GetOverridePartitionColumns returns the value of the OverridePartitionColumns field in AutoLoader_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoLoader_SdkV2) GetOverridePartitionColumns(ctx context.Context) ([]types.String, bool) {
	if m.OverridePartitionColumns.IsNull() || m.OverridePartitionColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OverridePartitionColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOverridePartitionColumns sets the value of the OverridePartitionColumns field in AutoLoader_SdkV2.
func (m *AutoLoader_SdkV2) SetOverridePartitionColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["override_partition_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OverridePartitionColumns = types.ListValueMust(t, vs)
}

// GetRawTable returns the value of the RawTable field in AutoLoader_SdkV2 as
// a TableName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoLoader_SdkV2) GetRawTable(ctx context.Context) (TableName_SdkV2, bool) {
	var e TableName_SdkV2
	if m.RawTable.IsNull() || m.RawTable.IsUnknown() {
		return e, false
	}
	var v []TableName_SdkV2
	d := m.RawTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRawTable sets the value of the RawTable field in AutoLoader_SdkV2.
func (m *AutoLoader_SdkV2) SetRawTable(ctx context.Context, v TableName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["raw_table"]
	m.RawTable = types.ListValueMust(t, vs)
}

// Unwraps an AWS CloudWatch Logs delivery envelope: each element of `logEvents`
// becomes its own row, with the element's `message` parsed as the event body.
type CloudWatchLogsEnvelope_SdkV2 struct {
	// Whether to keep the envelope's `owner`, `logGroup`, and `logStream`
	// values in the `data.lw_header` field of each unwrapped row. Must be set
	// explicitly to `true` or `false`.
	HeaderRetention types.Bool `tfsdk:"header_retention"`
}

func (to *CloudWatchLogsEnvelope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloudWatchLogsEnvelope_SdkV2) {
}

func (to *CloudWatchLogsEnvelope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CloudWatchLogsEnvelope_SdkV2) {
}

func (m CloudWatchLogsEnvelope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["header_retention"] = attrs["header_retention"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudWatchLogsEnvelope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CloudWatchLogsEnvelope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudWatchLogsEnvelope_SdkV2
// only implements ToObjectValue() and Type().
func (m CloudWatchLogsEnvelope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"header_retention": m.HeaderRetention,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CloudWatchLogsEnvelope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"header_retention": types.BoolType,
		},
	}
}

// Catch-all per-catalog config — settings that don't belong to any one
// resource family. Singleton. Name format: `catalogs/{catalog}/commonConfig`.
type CommonConfig_SdkV2 struct {
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag types.String `tfsdk:"etag"`

	Name types.String `tfsdk:"name"`
	// Whether users may grant permissions on individual resources rather than
	// only at the Folder level.
	PerResourcePermissionsEnabled types.Bool `tfsdk:"per_resource_permissions_enabled"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.List `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version types.Int64 `tfsdk:"version"`
}

func (to *CommonConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CommonConfig_SdkV2) {
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (to *CommonConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CommonConfig_SdkV2) {
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (m CommonConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["per_resource_permissions_enabled"] = attrs["per_resource_permissions_enabled"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CommonConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CommonConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updated_by": reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommonConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m CommonConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                             m.Etag,
			"name":                             m.Name,
			"per_resource_permissions_enabled": m.PerResourcePermissionsEnabled,
			"update_time":                      m.UpdateTime,
			"updated_by":                       m.UpdatedBy,
			"version":                          m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CommonConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":                             types.StringType,
			"name":                             types.StringType,
			"per_resource_permissions_enabled": types.BoolType,
			"update_time":                      timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetUpdatedBy returns the value of the UpdatedBy field in CommonConfig_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CommonConfig_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in CommonConfig_SdkV2.
func (m *CommonConfig_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Reference to the ContentPack (and pack version) a resource was originally
// installed from.
type ContentPackReference_SdkV2 struct {
	// UUID of the content pack.
	Uuid types.String `tfsdk:"uuid"`
	// Pack version the resource was installed from.
	Version types.String `tfsdk:"version"`
}

func (to *ContentPackReference_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ContentPackReference_SdkV2) {
}

func (to *ContentPackReference_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ContentPackReference_SdkV2) {
}

func (m ContentPackReference_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["uuid"] = attrs["uuid"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContentPackReference.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ContentPackReference_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContentPackReference_SdkV2
// only implements ToObjectValue() and Type().
func (m ContentPackReference_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uuid":    m.Uuid,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ContentPackReference_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uuid":    types.StringType,
			"version": types.StringType,
		},
	}
}

type CreateDatasourceRequest_SdkV2 struct {
	Datasource types.List `tfsdk:"datasource"`
}

func (to *CreateDatasourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatasourceRequest_SdkV2) {
	if !from.Datasource.IsNull() && !from.Datasource.IsUnknown() {
		if toDatasource, ok := to.GetDatasource(ctx); ok {
			if fromDatasource, ok := from.GetDatasource(ctx); ok {
				// Recursively sync the fields of Datasource
				toDatasource.SyncFieldsDuringCreateOrUpdate(ctx, fromDatasource)
				to.SetDatasource(ctx, toDatasource)
			}
		}
	}
}

func (to *CreateDatasourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDatasourceRequest_SdkV2) {
	if !from.Datasource.IsNull() && !from.Datasource.IsUnknown() {
		if toDatasource, ok := to.GetDatasource(ctx); ok {
			if fromDatasource, ok := from.GetDatasource(ctx); ok {
				toDatasource.SyncFieldsDuringRead(ctx, fromDatasource)
				to.SetDatasource(ctx, toDatasource)
			}
		}
	}
}

func (m CreateDatasourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasource"] = attrs["datasource"].SetRequired()
	attrs["datasource"] = attrs["datasource"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatasourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatasourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasource": reflect.TypeOf(Datasource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatasourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDatasourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasource": m.Datasource,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatasourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasource": basetypes.ListType{
				ElemType: Datasource_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDatasource returns the value of the Datasource field in CreateDatasourceRequest_SdkV2 as
// a Datasource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatasourceRequest_SdkV2) GetDatasource(ctx context.Context) (Datasource_SdkV2, bool) {
	var e Datasource_SdkV2
	if m.Datasource.IsNull() || m.Datasource.IsUnknown() {
		return e, false
	}
	var v []Datasource_SdkV2
	d := m.Datasource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatasource sets the value of the Datasource field in CreateDatasourceRequest_SdkV2.
func (m *CreateDatasourceRequest_SdkV2) SetDatasource(ctx context.Context, v Datasource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasource"]
	m.Datasource = types.ListValueMust(t, vs)
}

type CreateRuleRequest_SdkV2 struct {
	// The rule to create. Server-assigned fields (`name`, `rule_id`,
	// timestamps, `etag`, …) are ignored on the request.
	Rule types.List `tfsdk:"rule"`
}

func (to *CreateRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRuleRequest_SdkV2) {
	if !from.Rule.IsNull() && !from.Rule.IsUnknown() {
		if toRule, ok := to.GetRule(ctx); ok {
			if fromRule, ok := from.GetRule(ctx); ok {
				// Recursively sync the fields of Rule
				toRule.SyncFieldsDuringCreateOrUpdate(ctx, fromRule)
				to.SetRule(ctx, toRule)
			}
		}
	}
}

func (to *CreateRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRuleRequest_SdkV2) {
	if !from.Rule.IsNull() && !from.Rule.IsUnknown() {
		if toRule, ok := to.GetRule(ctx); ok {
			if fromRule, ok := from.GetRule(ctx); ok {
				toRule.SyncFieldsDuringRead(ctx, fromRule)
				to.SetRule(ctx, toRule)
			}
		}
	}
}

func (m CreateRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["rule"] = attrs["rule"].SetRequired()
	attrs["rule"] = attrs["rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule": reflect.TypeOf(Rule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"rule": m.Rule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rule": basetypes.ListType{
				ElemType: Rule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRule returns the value of the Rule field in CreateRuleRequest_SdkV2 as
// a Rule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRuleRequest_SdkV2) GetRule(ctx context.Context) (Rule_SdkV2, bool) {
	var e Rule_SdkV2
	if m.Rule.IsNull() || m.Rule.IsUnknown() {
		return e, false
	}
	var v []Rule_SdkV2
	d := m.Rule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRule sets the value of the Rule field in CreateRuleRequest_SdkV2.
func (m *CreateRuleRequest_SdkV2) SetRule(ctx context.Context, v Rule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rule"]
	m.Rule = types.ListValueMust(t, vs)
}

// A Lakewatch datasource: a pipeline that ingests security data and writes it
// to Unity Catalog tables.
//
// What a datasource ingests is defined by its ingestion method (`autoloader`,
// `unity_catalog_connection`, or `existing_table`), which sets how the data
// arrives, together with its `parsers`, which define the schemas that the
// arriving data is matched against and written as.
type Datasource_SdkV2 struct {
	// Free-form key-value annotations on the datasource, managed by the client.
	Annotations types.Map `tfsdk:"annotations"`

	Autoloader types.List `tfsdk:"autoloader"`
	// Name of the Unity Catalog catalog that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the catalog's current name is reported in
	// `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Compute mode for the datasource's ingestion job. If unset, the
	// `default_compute_mode` of the catalog's datasource configuration applies,
	// or `DS_COMPUTE_PERFORMANCE_OPTIMIZED` if the catalog has none.
	ComputeMode types.String `tfsdk:"compute_mode"`
	// ID of the workspace whose compute runs this datasource's ingestion job
	// and owns its associated resources, such as the storage the ingested data
	// lands in. Server-assigned.
	ComputeWorkspaceId types.String `tfsdk:"compute_workspace_id"`

	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	CreatedBy types.List `tfsdk:"created_by"`
	// Unity Catalog securable ID (a UUID) of the datasource. It never changes,
	// including when the datasource is renamed in Unity Catalog, and is the
	// final segment of `name`.
	DatasourceId types.String `tfsdk:"datasource_id"`
	// Name of the datasource in Unity Catalog, unique within its schema.
	// Required on create and can be updated within the same catalog and schema.
	// Get synchronizes this field with the current Unity Catalog name, also
	// reported in `effective_datasource_name`.
	DatasourceName types.String `tfsdk:"datasource_name"`
	// Free-form description of the datasource.
	Description types.String `tfsdk:"description"`
	// Current name of the Unity Catalog catalog that contains the datasource,
	// which reflects catalog renames without a change to `catalog_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// Current UC leaf name of the datasource (server-owned). Tracks UC renames
	// so callers can observe the live name without having to update the
	// client-owned `datasource_name`; unset until resolved.
	EffectiveDatasourceName types.String `tfsdk:"effective_datasource_name"`
	// The identity the datasource runs as: the resolved `run_as_principal`, or
	// the requester's identity if `run_as_principal` was omitted on create or
	// cleared on update. Given as a user email or service principal application
	// ID, or as the numeric principal ID if the identity can no longer be
	// resolved. Populated on get, create, and update responses; list responses
	// omit it.
	EffectiveRunAsPrincipal types.String `tfsdk:"effective_run_as_principal"`
	// Current name of the Unity Catalog schema that contains the datasource,
	// which reflects schema renames without a change to `schema_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Opaque entity tag for optimistic concurrency control. Returned by get,
	// create, and update responses and by list responses in the
	// `DATASOURCE_VIEW_FULL` view. To make an update conditional on the
	// datasource not having changed since it was read, pass this value as the
	// update request's `etag`. Changes to `status` do not change the etag.
	Etag types.String `tfsdk:"etag"`

	ExistingTable types.List `tfsdk:"existing_table"`
	// Resource name of the datasource, in the format
	// `datasources/{datasource_id}`. Unlike the datasource's Unity Catalog
	// location (`catalog_name`, `schema_name`, and `datasource_name`), the name
	// does not change when the datasource or its catalog or schema is renamed
	// in Unity Catalog.
	Name types.String `tfsdk:"name"`
	// Unity Catalog owner of the datasource, assigned by the server. Create
	// responses omit it.
	Owner types.List `tfsdk:"owner"`
	// Parsers that schematize the ingested data. Each record is matched against
	// the parsers in order and written to the parsed table of the first parser
	// that it matches. Existing-table and Unity Catalog connection datasources
	// require at least one parser; an Auto Loader datasource with no parsers
	// writes only its raw table.
	Parsers types.List `tfsdk:"parsers"`
	// The user or service principal that the datasource runs as, given as a
	// user email or service principal application ID. When it is supplied on
	// create or update, the requester must be authorized to run the datasource
	// as that principal, or the request is rejected. When it is omitted on
	// create, the datasource runs as the requester. On update, a mask that
	// includes `run_as_principal` without a value clears it, so the datasource
	// runs as the requester; a mask that omits it keeps the current identity.
	// Reads return the value as supplied; the identity the datasource runs as
	// is reported in `effective_run_as_principal`.
	RunAsPrincipal types.String `tfsdk:"run_as_principal"`
	// Schedule on which the datasource's ingestion job runs. A datasource
	// cannot be paused, so a request that sets `schedule.enabled` to `false` is
	// rejected. Continuous schedules are not supported for Unity Catalog
	// connection datasources.
	Schedule types.List `tfsdk:"schedule"`
	// Name of the Unity Catalog schema that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the schema's current name is reported in
	// `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Runtime status of the jobs and pipelines that run this datasource.
	Status types.List `tfsdk:"status"`

	UnityCatalogConnection types.List `tfsdk:"unity_catalog_connection"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.List `tfsdk:"updated_by"`
	// Version number of the datasource, incremented on each update.
	// Informational only; use `etag` for optimistic concurrency control.
	Version types.Int64 `tfsdk:"version"`
	// ID of the SQL warehouse that runs the checks Lakewatch performs to
	// validate a create or update request, such as source and destination
	// access checks. Required and non-blank on every create and update request
	// except a rename-only PATCH; there is no default warehouse. The requester
	// must be able to use the warehouse. It is not stored, is not returned on
	// reads, and does not affect the compute that runs the ingestion job.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *Datasource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Datasource_SdkV2) {
	if !from.Autoloader.IsNull() && !from.Autoloader.IsUnknown() {
		if toAutoloader, ok := to.GetAutoloader(ctx); ok {
			if fromAutoloader, ok := from.GetAutoloader(ctx); ok {
				// Recursively sync the fields of Autoloader
				toAutoloader.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoloader)
				to.SetAutoloader(ctx, toAutoloader)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				// Recursively sync the fields of CreatedBy
				toCreatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.ExistingTable.IsNull() && !from.ExistingTable.IsUnknown() {
		if toExistingTable, ok := to.GetExistingTable(ctx); ok {
			if fromExistingTable, ok := from.GetExistingTable(ctx); ok {
				// Recursively sync the fields of ExistingTable
				toExistingTable.SyncFieldsDuringCreateOrUpdate(ctx, fromExistingTable)
				to.SetExistingTable(ctx, toExistingTable)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				// Recursively sync the fields of Owner
				toOwner.SyncFieldsDuringCreateOrUpdate(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() && to.Parsers.IsNull() && len(from.Parsers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parsers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parsers = from.Parsers
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() {
		if toParsers, ok := to.GetParsers(ctx); ok {
			if fromParsers, ok := from.GetParsers(ctx); ok {
				// Recursively sync the fields of each Parsers element by position.
				for i := range toParsers {
					if i < len(fromParsers) {
						toParsers[i].SyncFieldsDuringCreateOrUpdate(ctx, fromParsers[i])
					}
				}
				to.SetParsers(ctx, toParsers)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				// Recursively sync the fields of Schedule
				toSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.UnityCatalogConnection.IsNull() && !from.UnityCatalogConnection.IsUnknown() {
		if toUnityCatalogConnection, ok := to.GetUnityCatalogConnection(ctx); ok {
			if fromUnityCatalogConnection, ok := from.GetUnityCatalogConnection(ctx); ok {
				// Recursively sync the fields of UnityCatalogConnection
				toUnityCatalogConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromUnityCatalogConnection)
				to.SetUnityCatalogConnection(ctx, toUnityCatalogConnection)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	if !from.WarehouseId.IsUnknown() && !from.WarehouseId.IsNull() {
		// WarehouseId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.WarehouseId = from.WarehouseId
	}
}

func (to *Datasource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Datasource_SdkV2) {
	if !from.Autoloader.IsNull() && !from.Autoloader.IsUnknown() {
		if toAutoloader, ok := to.GetAutoloader(ctx); ok {
			if fromAutoloader, ok := from.GetAutoloader(ctx); ok {
				toAutoloader.SyncFieldsDuringRead(ctx, fromAutoloader)
				to.SetAutoloader(ctx, toAutoloader)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				toCreatedBy.SyncFieldsDuringRead(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.ExistingTable.IsNull() && !from.ExistingTable.IsUnknown() {
		if toExistingTable, ok := to.GetExistingTable(ctx); ok {
			if fromExistingTable, ok := from.GetExistingTable(ctx); ok {
				toExistingTable.SyncFieldsDuringRead(ctx, fromExistingTable)
				to.SetExistingTable(ctx, toExistingTable)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				toOwner.SyncFieldsDuringRead(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() && to.Parsers.IsNull() && len(from.Parsers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parsers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parsers = from.Parsers
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() {
		if toParsers, ok := to.GetParsers(ctx); ok {
			if fromParsers, ok := from.GetParsers(ctx); ok {
				for i := range toParsers {
					if i < len(fromParsers) {
						toParsers[i].SyncFieldsDuringRead(ctx, fromParsers[i])
					}
				}
				to.SetParsers(ctx, toParsers)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				toSchedule.SyncFieldsDuringRead(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.UnityCatalogConnection.IsNull() && !from.UnityCatalogConnection.IsUnknown() {
		if toUnityCatalogConnection, ok := to.GetUnityCatalogConnection(ctx); ok {
			if fromUnityCatalogConnection, ok := from.GetUnityCatalogConnection(ctx); ok {
				toUnityCatalogConnection.SyncFieldsDuringRead(ctx, fromUnityCatalogConnection)
				to.SetUnityCatalogConnection(ctx, toUnityCatalogConnection)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	if !from.WarehouseId.IsUnknown() && !from.WarehouseId.IsNull() {
		// WarehouseId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.WarehouseId = from.WarehouseId
	}
}

func (m Datasource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetOptional()
	attrs["autoloader"] = attrs["autoloader"].SetOptional()
	attrs["autoloader"] = attrs["autoloader"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["compute_mode"] = attrs["compute_mode"].SetOptional()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["created_by"] = attrs["created_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["datasource_id"] = attrs["datasource_id"].SetComputed()
	attrs["datasource_name"] = attrs["datasource_name"].SetRequired()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_datasource_name"] = attrs["effective_datasource_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["existing_table"] = attrs["existing_table"].SetOptional()
	attrs["existing_table"] = attrs["existing_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["owner"] = attrs["owner"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parsers"] = attrs["parsers"].SetOptional()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetRequired()
	attrs["schedule"] = attrs["schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["unity_catalog_connection"] = attrs["unity_catalog_connection"].SetOptional()
	attrs["unity_catalog_connection"] = attrs["unity_catalog_connection"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Datasource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Datasource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"annotations":              reflect.TypeOf(types.String{}),
		"autoloader":               reflect.TypeOf(AutoLoader_SdkV2{}),
		"created_by":               reflect.TypeOf(Principal_SdkV2{}),
		"existing_table":           reflect.TypeOf(ExistingTable_SdkV2{}),
		"owner":                    reflect.TypeOf(Principal_SdkV2{}),
		"parsers":                  reflect.TypeOf(ParserRef_SdkV2{}),
		"schedule":                 reflect.TypeOf(Schedule_SdkV2{}),
		"status":                   reflect.TypeOf(DatasourceStatus_SdkV2{}),
		"unity_catalog_connection": reflect.TypeOf(UnityCatalogConnection_SdkV2{}),
		"updated_by":               reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Datasource_SdkV2
// only implements ToObjectValue() and Type().
func (m Datasource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"annotations":                m.Annotations,
			"autoloader":                 m.Autoloader,
			"catalog_name":               m.CatalogName,
			"compute_mode":               m.ComputeMode,
			"compute_workspace_id":       m.ComputeWorkspaceId,
			"create_time":                m.CreateTime,
			"created_by":                 m.CreatedBy,
			"datasource_id":              m.DatasourceId,
			"datasource_name":            m.DatasourceName,
			"description":                m.Description,
			"effective_catalog_name":     m.EffectiveCatalogName,
			"effective_datasource_name":  m.EffectiveDatasourceName,
			"effective_run_as_principal": m.EffectiveRunAsPrincipal,
			"effective_schema_name":      m.EffectiveSchemaName,
			"etag":                       m.Etag,
			"existing_table":             m.ExistingTable,
			"name":                       m.Name,
			"owner":                      m.Owner,
			"parsers":                    m.Parsers,
			"run_as_principal":           m.RunAsPrincipal,
			"schedule":                   m.Schedule,
			"schema_name":                m.SchemaName,
			"status":                     m.Status,
			"unity_catalog_connection":   m.UnityCatalogConnection,
			"update_time":                m.UpdateTime,
			"updated_by":                 m.UpdatedBy,
			"version":                    m.Version,
			"warehouse_id":               m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Datasource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"annotations": basetypes.MapType{
				ElemType: types.StringType,
			},
			"autoloader": basetypes.ListType{
				ElemType: AutoLoader_SdkV2{}.Type(ctx),
			},
			"catalog_name":         types.StringType,
			"compute_mode":         types.StringType,
			"compute_workspace_id": types.StringType,
			"create_time":          timetypes.RFC3339{}.Type(ctx),
			"created_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"datasource_id":              types.StringType,
			"datasource_name":            types.StringType,
			"description":                types.StringType,
			"effective_catalog_name":     types.StringType,
			"effective_datasource_name":  types.StringType,
			"effective_run_as_principal": types.StringType,
			"effective_schema_name":      types.StringType,
			"etag":                       types.StringType,
			"existing_table": basetypes.ListType{
				ElemType: ExistingTable_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"owner": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"parsers": basetypes.ListType{
				ElemType: ParserRef_SdkV2{}.Type(ctx),
			},
			"run_as_principal": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule_SdkV2{}.Type(ctx),
			},
			"schema_name": types.StringType,
			"status": basetypes.ListType{
				ElemType: DatasourceStatus_SdkV2{}.Type(ctx),
			},
			"unity_catalog_connection": basetypes.ListType{
				ElemType: UnityCatalogConnection_SdkV2{}.Type(ctx),
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version":      types.Int64Type,
			"warehouse_id": types.StringType,
		},
	}
}

// GetAnnotations returns the value of the Annotations field in Datasource_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetAnnotations(ctx context.Context) (map[string]types.String, bool) {
	if m.Annotations.IsNull() || m.Annotations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Annotations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnnotations sets the value of the Annotations field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetAnnotations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["annotations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Annotations = types.MapValueMust(t, vs)
}

// GetAutoloader returns the value of the Autoloader field in Datasource_SdkV2 as
// a AutoLoader_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetAutoloader(ctx context.Context) (AutoLoader_SdkV2, bool) {
	var e AutoLoader_SdkV2
	if m.Autoloader.IsNull() || m.Autoloader.IsUnknown() {
		return e, false
	}
	var v []AutoLoader_SdkV2
	d := m.Autoloader.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoloader sets the value of the Autoloader field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetAutoloader(ctx context.Context, v AutoLoader_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["autoloader"]
	m.Autoloader = types.ListValueMust(t, vs)
}

// GetCreatedBy returns the value of the CreatedBy field in Datasource_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetCreatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.CreatedBy.IsNull() || m.CreatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.CreatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCreatedBy sets the value of the CreatedBy field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetCreatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["created_by"]
	m.CreatedBy = types.ListValueMust(t, vs)
}

// GetExistingTable returns the value of the ExistingTable field in Datasource_SdkV2 as
// a ExistingTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetExistingTable(ctx context.Context) (ExistingTable_SdkV2, bool) {
	var e ExistingTable_SdkV2
	if m.ExistingTable.IsNull() || m.ExistingTable.IsUnknown() {
		return e, false
	}
	var v []ExistingTable_SdkV2
	d := m.ExistingTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExistingTable sets the value of the ExistingTable field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetExistingTable(ctx context.Context, v ExistingTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["existing_table"]
	m.ExistingTable = types.ListValueMust(t, vs)
}

// GetOwner returns the value of the Owner field in Datasource_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetOwner(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.Owner.IsNull() || m.Owner.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.Owner.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOwner sets the value of the Owner field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetOwner(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["owner"]
	m.Owner = types.ListValueMust(t, vs)
}

// GetParsers returns the value of the Parsers field in Datasource_SdkV2 as
// a slice of ParserRef_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetParsers(ctx context.Context) ([]ParserRef_SdkV2, bool) {
	if m.Parsers.IsNull() || m.Parsers.IsUnknown() {
		return nil, false
	}
	var v []ParserRef_SdkV2
	d := m.Parsers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParsers sets the value of the Parsers field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetParsers(ctx context.Context, v []ParserRef_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parsers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parsers = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in Datasource_SdkV2 as
// a Schedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetSchedule(ctx context.Context) (Schedule_SdkV2, bool) {
	var e Schedule_SdkV2
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v []Schedule_SdkV2
	d := m.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetSchedule(ctx context.Context, v Schedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	m.Schedule = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Datasource_SdkV2 as
// a DatasourceStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetStatus(ctx context.Context) (DatasourceStatus_SdkV2, bool) {
	var e DatasourceStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []DatasourceStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetStatus(ctx context.Context, v DatasourceStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// GetUnityCatalogConnection returns the value of the UnityCatalogConnection field in Datasource_SdkV2 as
// a UnityCatalogConnection_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetUnityCatalogConnection(ctx context.Context) (UnityCatalogConnection_SdkV2, bool) {
	var e UnityCatalogConnection_SdkV2
	if m.UnityCatalogConnection.IsNull() || m.UnityCatalogConnection.IsUnknown() {
		return e, false
	}
	var v []UnityCatalogConnection_SdkV2
	d := m.UnityCatalogConnection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUnityCatalogConnection sets the value of the UnityCatalogConnection field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetUnityCatalogConnection(ctx context.Context, v UnityCatalogConnection_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["unity_catalog_connection"]
	m.UnityCatalogConnection = types.ListValueMust(t, vs)
}

// GetUpdatedBy returns the value of the UpdatedBy field in Datasource_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in Datasource_SdkV2.
func (m *Datasource_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Per-catalog datasource-related config. Singleton. Name format:
// `catalogs/{catalog}/datasourceConfig`.
type DatasourceConfig_SdkV2 struct {
	// Default compute mode applied to datasource ingestion jobs when the
	// individual datasource doesn't override it. v1
	// `datasources.defaultComputeMode`.
	DefaultComputeMode types.String `tfsdk:"default_compute_mode"`
	// Schema name of the registered folder (in this catalog) that a new
	// datasource is placed in when the request does not specify a schema. Must
	// name a schema registered as a folder. Omit to clear. Reads return the
	// value as supplied; the folder's current schema name is reported in
	// `effective_default_folder`.
	DefaultFolder types.String `tfsdk:"default_folder"`
	// The current schema name of the `default_folder` folder. Tracks UC schema
	// renames. Unset when no default is configured, or when the folder's schema
	// no longer exists or is not visible to the caller.
	EffectiveDefaultFolder types.String `tfsdk:"effective_default_folder"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag types.String `tfsdk:"etag"`

	Name types.String `tfsdk:"name"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.List `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version types.Int64 `tfsdk:"version"`
}

func (to *DatasourceConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatasourceConfig_SdkV2) {
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (to *DatasourceConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatasourceConfig_SdkV2) {
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (m DatasourceConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_compute_mode"] = attrs["default_compute_mode"].SetOptional()
	attrs["default_folder"] = attrs["default_folder"].SetOptional()
	attrs["effective_default_folder"] = attrs["effective_default_folder"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatasourceConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatasourceConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updated_by": reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasourceConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m DatasourceConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_compute_mode":     m.DefaultComputeMode,
			"default_folder":           m.DefaultFolder,
			"effective_default_folder": m.EffectiveDefaultFolder,
			"etag":                     m.Etag,
			"name":                     m.Name,
			"update_time":              m.UpdateTime,
			"updated_by":               m.UpdatedBy,
			"version":                  m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatasourceConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_compute_mode":     types.StringType,
			"default_folder":           types.StringType,
			"effective_default_folder": types.StringType,
			"etag":                     types.StringType,
			"name":                     types.StringType,
			"update_time":              timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetUpdatedBy returns the value of the UpdatedBy field in DatasourceConfig_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasourceConfig_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in DatasourceConfig_SdkV2.
func (m *DatasourceConfig_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Read-only runtime status of a datasource, with the status of each job and
// pipeline that runs it.
type DatasourceStatus_SdkV2 struct {
	// Status of the datasource's ingestion job.
	IngestionJob types.List `tfsdk:"ingestion_job"`
	// Status of the managed ingestion pipeline. Populated only when
	// `unity_catalog_connection` is set, with `pipeline_id` identifying the
	// pipeline.
	IngestionPipeline types.List `tfsdk:"ingestion_pipeline"`
}

func (to *DatasourceStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatasourceStatus_SdkV2) {
	if !from.IngestionJob.IsNull() && !from.IngestionJob.IsUnknown() {
		if toIngestionJob, ok := to.GetIngestionJob(ctx); ok {
			if fromIngestionJob, ok := from.GetIngestionJob(ctx); ok {
				// Recursively sync the fields of IngestionJob
				toIngestionJob.SyncFieldsDuringCreateOrUpdate(ctx, fromIngestionJob)
				to.SetIngestionJob(ctx, toIngestionJob)
			}
		}
	}
	if !from.IngestionPipeline.IsNull() && !from.IngestionPipeline.IsUnknown() {
		if toIngestionPipeline, ok := to.GetIngestionPipeline(ctx); ok {
			if fromIngestionPipeline, ok := from.GetIngestionPipeline(ctx); ok {
				// Recursively sync the fields of IngestionPipeline
				toIngestionPipeline.SyncFieldsDuringCreateOrUpdate(ctx, fromIngestionPipeline)
				to.SetIngestionPipeline(ctx, toIngestionPipeline)
			}
		}
	}
}

func (to *DatasourceStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatasourceStatus_SdkV2) {
	if !from.IngestionJob.IsNull() && !from.IngestionJob.IsUnknown() {
		if toIngestionJob, ok := to.GetIngestionJob(ctx); ok {
			if fromIngestionJob, ok := from.GetIngestionJob(ctx); ok {
				toIngestionJob.SyncFieldsDuringRead(ctx, fromIngestionJob)
				to.SetIngestionJob(ctx, toIngestionJob)
			}
		}
	}
	if !from.IngestionPipeline.IsNull() && !from.IngestionPipeline.IsUnknown() {
		if toIngestionPipeline, ok := to.GetIngestionPipeline(ctx); ok {
			if fromIngestionPipeline, ok := from.GetIngestionPipeline(ctx); ok {
				toIngestionPipeline.SyncFieldsDuringRead(ctx, fromIngestionPipeline)
				to.SetIngestionPipeline(ctx, toIngestionPipeline)
			}
		}
	}
}

func (m DatasourceStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ingestion_job"] = attrs["ingestion_job"].SetComputed()
	attrs["ingestion_job"] = attrs["ingestion_job"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["ingestion_pipeline"] = attrs["ingestion_pipeline"].SetComputed()
	attrs["ingestion_pipeline"] = attrs["ingestion_pipeline"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatasourceStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatasourceStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ingestion_job":      reflect.TypeOf(ResourceStatus_SdkV2{}),
		"ingestion_pipeline": reflect.TypeOf(ResourceStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasourceStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m DatasourceStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ingestion_job":      m.IngestionJob,
			"ingestion_pipeline": m.IngestionPipeline,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatasourceStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ingestion_job": basetypes.ListType{
				ElemType: ResourceStatus_SdkV2{}.Type(ctx),
			},
			"ingestion_pipeline": basetypes.ListType{
				ElemType: ResourceStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIngestionJob returns the value of the IngestionJob field in DatasourceStatus_SdkV2 as
// a ResourceStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasourceStatus_SdkV2) GetIngestionJob(ctx context.Context) (ResourceStatus_SdkV2, bool) {
	var e ResourceStatus_SdkV2
	if m.IngestionJob.IsNull() || m.IngestionJob.IsUnknown() {
		return e, false
	}
	var v []ResourceStatus_SdkV2
	d := m.IngestionJob.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionJob sets the value of the IngestionJob field in DatasourceStatus_SdkV2.
func (m *DatasourceStatus_SdkV2) SetIngestionJob(ctx context.Context, v ResourceStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_job"]
	m.IngestionJob = types.ListValueMust(t, vs)
}

// GetIngestionPipeline returns the value of the IngestionPipeline field in DatasourceStatus_SdkV2 as
// a ResourceStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasourceStatus_SdkV2) GetIngestionPipeline(ctx context.Context) (ResourceStatus_SdkV2, bool) {
	var e ResourceStatus_SdkV2
	if m.IngestionPipeline.IsNull() || m.IngestionPipeline.IsUnknown() {
		return e, false
	}
	var v []ResourceStatus_SdkV2
	d := m.IngestionPipeline.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionPipeline sets the value of the IngestionPipeline field in DatasourceStatus_SdkV2.
func (m *DatasourceStatus_SdkV2) SetIngestionPipeline(ctx context.Context, v ResourceStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_pipeline"]
	m.IngestionPipeline = types.ListValueMust(t, vs)
}

type DeleteDatasourceRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDatasourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatasourceRequest_SdkV2) {
}

func (to *DeleteDatasourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDatasourceRequest_SdkV2) {
}

func (m DeleteDatasourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatasourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatasourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatasourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDatasourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatasourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRuleRequest_SdkV2 struct {
	// Resource name of the rule to delete. Format: `rules/{rule_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRuleRequest_SdkV2) {
}

func (to *DeleteRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRuleRequest_SdkV2) {
}

func (m DeleteRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Unwraps a transport envelope from the source data so that each contained
// event becomes its own row. Applied when the source data is read, before
// parsing.
type Envelope_SdkV2 struct {
	CloudWatchLogs types.List `tfsdk:"cloud_watch_logs"`

	JsonArray types.List `tfsdk:"json_array"`
}

func (to *Envelope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Envelope_SdkV2) {
	if !from.CloudWatchLogs.IsNull() && !from.CloudWatchLogs.IsUnknown() {
		if toCloudWatchLogs, ok := to.GetCloudWatchLogs(ctx); ok {
			if fromCloudWatchLogs, ok := from.GetCloudWatchLogs(ctx); ok {
				// Recursively sync the fields of CloudWatchLogs
				toCloudWatchLogs.SyncFieldsDuringCreateOrUpdate(ctx, fromCloudWatchLogs)
				to.SetCloudWatchLogs(ctx, toCloudWatchLogs)
			}
		}
	}
	if !from.JsonArray.IsNull() && !from.JsonArray.IsUnknown() {
		if toJsonArray, ok := to.GetJsonArray(ctx); ok {
			if fromJsonArray, ok := from.GetJsonArray(ctx); ok {
				// Recursively sync the fields of JsonArray
				toJsonArray.SyncFieldsDuringCreateOrUpdate(ctx, fromJsonArray)
				to.SetJsonArray(ctx, toJsonArray)
			}
		}
	}
}

func (to *Envelope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Envelope_SdkV2) {
	if !from.CloudWatchLogs.IsNull() && !from.CloudWatchLogs.IsUnknown() {
		if toCloudWatchLogs, ok := to.GetCloudWatchLogs(ctx); ok {
			if fromCloudWatchLogs, ok := from.GetCloudWatchLogs(ctx); ok {
				toCloudWatchLogs.SyncFieldsDuringRead(ctx, fromCloudWatchLogs)
				to.SetCloudWatchLogs(ctx, toCloudWatchLogs)
			}
		}
	}
	if !from.JsonArray.IsNull() && !from.JsonArray.IsUnknown() {
		if toJsonArray, ok := to.GetJsonArray(ctx); ok {
			if fromJsonArray, ok := from.GetJsonArray(ctx); ok {
				toJsonArray.SyncFieldsDuringRead(ctx, fromJsonArray)
				to.SetJsonArray(ctx, toJsonArray)
			}
		}
	}
}

func (m Envelope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cloud_watch_logs"] = attrs["cloud_watch_logs"].SetOptional()
	attrs["cloud_watch_logs"] = attrs["cloud_watch_logs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["json_array"] = attrs["json_array"].SetOptional()
	attrs["json_array"] = attrs["json_array"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Envelope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Envelope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cloud_watch_logs": reflect.TypeOf(CloudWatchLogsEnvelope_SdkV2{}),
		"json_array":       reflect.TypeOf(JsonArrayEnvelope_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Envelope_SdkV2
// only implements ToObjectValue() and Type().
func (m Envelope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud_watch_logs": m.CloudWatchLogs,
			"json_array":       m.JsonArray,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Envelope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud_watch_logs": basetypes.ListType{
				ElemType: CloudWatchLogsEnvelope_SdkV2{}.Type(ctx),
			},
			"json_array": basetypes.ListType{
				ElemType: JsonArrayEnvelope_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCloudWatchLogs returns the value of the CloudWatchLogs field in Envelope_SdkV2 as
// a CloudWatchLogsEnvelope_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Envelope_SdkV2) GetCloudWatchLogs(ctx context.Context) (CloudWatchLogsEnvelope_SdkV2, bool) {
	var e CloudWatchLogsEnvelope_SdkV2
	if m.CloudWatchLogs.IsNull() || m.CloudWatchLogs.IsUnknown() {
		return e, false
	}
	var v []CloudWatchLogsEnvelope_SdkV2
	d := m.CloudWatchLogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudWatchLogs sets the value of the CloudWatchLogs field in Envelope_SdkV2.
func (m *Envelope_SdkV2) SetCloudWatchLogs(ctx context.Context, v CloudWatchLogsEnvelope_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_watch_logs"]
	m.CloudWatchLogs = types.ListValueMust(t, vs)
}

// GetJsonArray returns the value of the JsonArray field in Envelope_SdkV2 as
// a JsonArrayEnvelope_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Envelope_SdkV2) GetJsonArray(ctx context.Context) (JsonArrayEnvelope_SdkV2, bool) {
	var e JsonArrayEnvelope_SdkV2
	if m.JsonArray.IsNull() || m.JsonArray.IsUnknown() {
		return e, false
	}
	var v []JsonArrayEnvelope_SdkV2
	d := m.JsonArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJsonArray sets the value of the JsonArray field in Envelope_SdkV2.
func (m *Envelope_SdkV2) SetJsonArray(ctx context.Context, v JsonArrayEnvelope_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["json_array"]
	m.JsonArray = types.ListValueMust(t, vs)
}

// Uses an existing Unity Catalog Delta table or view as the datasource's
// source. Mutually exclusive with `autoloader` and `unity_catalog_connection`.
type ExistingTable_SdkV2 struct {
	// Name of the catalog the existing table lives in. Optional on create and
	// defaults to the datasource's own catalog. Cannot be changed. Reads return
	// the value as supplied; the catalog in use is reported in
	// `effective_catalog`.
	Catalog types.String `tfsdk:"catalog"`
	// Current Unity Catalog catalog name of the source table, resolved from its
	// stable table ID, so a rename of the source table is reported here without
	// changing `catalog`.
	EffectiveCatalog types.String `tfsdk:"effective_catalog"`
	// Current Unity Catalog name of the source table, resolved from its stable
	// table ID.
	EffectiveName types.String `tfsdk:"effective_name"`
	// Current Unity Catalog schema name of the source table, resolved from its
	// stable table ID.
	EffectiveSchema types.String `tfsdk:"effective_schema"`
	// Whether to ingest the rows already present in the source table when
	// ingestion first starts. Set true to ingest the existing rows and then
	// every new commit. Set false, or leave unset, to ingest only rows written
	// after the datasource starts. Defaults to false, and must be true for a
	// Delta Sharing source table. Takes effect only at first start and cannot
	// be changed. Reads return the value as supplied.
	IncludeExistingRows types.Bool `tfsdk:"include_existing_rows"`
	// Name of the source table or view. Required on create and cannot be
	// changed. Reads return the value as supplied; the table's current name is
	// reported in `effective_name`.
	Name types.String `tfsdk:"name"`
	// Name of the schema the existing table lives in. Optional, but required
	// when `catalog` is set; supplying `catalog` without `schema` is rejected.
	// Defaults to the datasource's own schema. Cannot be changed. Reads return
	// the value as supplied; the schema in use is reported in
	// `effective_schema`.
	Schema types.String `tfsdk:"schema"`
}

func (to *ExistingTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExistingTable_SdkV2) {
}

func (to *ExistingTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExistingTable_SdkV2) {
}

func (m ExistingTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["effective_catalog"] = attrs["effective_catalog"].SetComputed()
	attrs["effective_name"] = attrs["effective_name"].SetComputed()
	attrs["effective_schema"] = attrs["effective_schema"].SetComputed()
	attrs["include_existing_rows"] = attrs["include_existing_rows"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["schema"] = attrs["schema"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExistingTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExistingTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExistingTable_SdkV2
// only implements ToObjectValue() and Type().
func (m ExistingTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog":               m.Catalog,
			"effective_catalog":     m.EffectiveCatalog,
			"effective_name":        m.EffectiveName,
			"effective_schema":      m.EffectiveSchema,
			"include_existing_rows": m.IncludeExistingRows,
			"name":                  m.Name,
			"schema":                m.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExistingTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog":               types.StringType,
			"effective_catalog":     types.StringType,
			"effective_name":        types.StringType,
			"effective_schema":      types.StringType,
			"include_existing_rows": types.BoolType,
			"name":                  types.StringType,
			"schema":                types.StringType,
		},
	}
}

// Configuration for delivering events / alerts to an external destination.
//
// Exactly one of the destination variants is set; the server picks it via the
// `oneof`.
type ExportConfig_SdkV2 struct {
	// Whether export is currently active.
	Enabled types.Bool `tfsdk:"enabled"`

	Jira types.List `tfsdk:"jira"`

	ServiceNow types.List `tfsdk:"service_now"`

	Slack types.List `tfsdk:"slack"`

	Tines types.List `tfsdk:"tines"`

	Webhook types.List `tfsdk:"webhook"`
	// Deprecated and ignored: exports always egress through the request's own
	// workspace. Kept so configs stored by older builds still round-trip.
	WorkspaceId types.String `tfsdk:"workspace_id"`
}

func (to *ExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportConfig_SdkV2) {
	if !from.Jira.IsNull() && !from.Jira.IsUnknown() {
		if toJira, ok := to.GetJira(ctx); ok {
			if fromJira, ok := from.GetJira(ctx); ok {
				// Recursively sync the fields of Jira
				toJira.SyncFieldsDuringCreateOrUpdate(ctx, fromJira)
				to.SetJira(ctx, toJira)
			}
		}
	}
	if !from.ServiceNow.IsNull() && !from.ServiceNow.IsUnknown() {
		if toServiceNow, ok := to.GetServiceNow(ctx); ok {
			if fromServiceNow, ok := from.GetServiceNow(ctx); ok {
				// Recursively sync the fields of ServiceNow
				toServiceNow.SyncFieldsDuringCreateOrUpdate(ctx, fromServiceNow)
				to.SetServiceNow(ctx, toServiceNow)
			}
		}
	}
	if !from.Slack.IsNull() && !from.Slack.IsUnknown() {
		if toSlack, ok := to.GetSlack(ctx); ok {
			if fromSlack, ok := from.GetSlack(ctx); ok {
				// Recursively sync the fields of Slack
				toSlack.SyncFieldsDuringCreateOrUpdate(ctx, fromSlack)
				to.SetSlack(ctx, toSlack)
			}
		}
	}
	if !from.Tines.IsNull() && !from.Tines.IsUnknown() {
		if toTines, ok := to.GetTines(ctx); ok {
			if fromTines, ok := from.GetTines(ctx); ok {
				// Recursively sync the fields of Tines
				toTines.SyncFieldsDuringCreateOrUpdate(ctx, fromTines)
				to.SetTines(ctx, toTines)
			}
		}
	}
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				// Recursively sync the fields of Webhook
				toWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (to *ExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExportConfig_SdkV2) {
	if !from.Jira.IsNull() && !from.Jira.IsUnknown() {
		if toJira, ok := to.GetJira(ctx); ok {
			if fromJira, ok := from.GetJira(ctx); ok {
				toJira.SyncFieldsDuringRead(ctx, fromJira)
				to.SetJira(ctx, toJira)
			}
		}
	}
	if !from.ServiceNow.IsNull() && !from.ServiceNow.IsUnknown() {
		if toServiceNow, ok := to.GetServiceNow(ctx); ok {
			if fromServiceNow, ok := from.GetServiceNow(ctx); ok {
				toServiceNow.SyncFieldsDuringRead(ctx, fromServiceNow)
				to.SetServiceNow(ctx, toServiceNow)
			}
		}
	}
	if !from.Slack.IsNull() && !from.Slack.IsUnknown() {
		if toSlack, ok := to.GetSlack(ctx); ok {
			if fromSlack, ok := from.GetSlack(ctx); ok {
				toSlack.SyncFieldsDuringRead(ctx, fromSlack)
				to.SetSlack(ctx, toSlack)
			}
		}
	}
	if !from.Tines.IsNull() && !from.Tines.IsUnknown() {
		if toTines, ok := to.GetTines(ctx); ok {
			if fromTines, ok := from.GetTines(ctx); ok {
				toTines.SyncFieldsDuringRead(ctx, fromTines)
				to.SetTines(ctx, toTines)
			}
		}
	}
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				toWebhook.SyncFieldsDuringRead(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (m ExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["jira"] = attrs["jira"].SetOptional()
	attrs["jira"] = attrs["jira"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["service_now"] = attrs["service_now"].SetOptional()
	attrs["service_now"] = attrs["service_now"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["slack"] = attrs["slack"].SetOptional()
	attrs["slack"] = attrs["slack"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["tines"] = attrs["tines"].SetOptional()
	attrs["tines"] = attrs["tines"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["webhook"] = attrs["webhook"].SetOptional()
	attrs["webhook"] = attrs["webhook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jira":        reflect.TypeOf(JiraExportConfig_SdkV2{}),
		"service_now": reflect.TypeOf(ServiceNowExportConfig_SdkV2{}),
		"slack":       reflect.TypeOf(SlackExportConfig_SdkV2{}),
		"tines":       reflect.TypeOf(TinesExportConfig_SdkV2{}),
		"webhook":     reflect.TypeOf(WebhookExportConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m ExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":      m.Enabled,
			"jira":         m.Jira,
			"service_now":  m.ServiceNow,
			"slack":        m.Slack,
			"tines":        m.Tines,
			"webhook":      m.Webhook,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
			"jira": basetypes.ListType{
				ElemType: JiraExportConfig_SdkV2{}.Type(ctx),
			},
			"service_now": basetypes.ListType{
				ElemType: ServiceNowExportConfig_SdkV2{}.Type(ctx),
			},
			"slack": basetypes.ListType{
				ElemType: SlackExportConfig_SdkV2{}.Type(ctx),
			},
			"tines": basetypes.ListType{
				ElemType: TinesExportConfig_SdkV2{}.Type(ctx),
			},
			"webhook": basetypes.ListType{
				ElemType: WebhookExportConfig_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.StringType,
		},
	}
}

// GetJira returns the value of the Jira field in ExportConfig_SdkV2 as
// a JiraExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExportConfig_SdkV2) GetJira(ctx context.Context) (JiraExportConfig_SdkV2, bool) {
	var e JiraExportConfig_SdkV2
	if m.Jira.IsNull() || m.Jira.IsUnknown() {
		return e, false
	}
	var v []JiraExportConfig_SdkV2
	d := m.Jira.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJira sets the value of the Jira field in ExportConfig_SdkV2.
func (m *ExportConfig_SdkV2) SetJira(ctx context.Context, v JiraExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["jira"]
	m.Jira = types.ListValueMust(t, vs)
}

// GetServiceNow returns the value of the ServiceNow field in ExportConfig_SdkV2 as
// a ServiceNowExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExportConfig_SdkV2) GetServiceNow(ctx context.Context) (ServiceNowExportConfig_SdkV2, bool) {
	var e ServiceNowExportConfig_SdkV2
	if m.ServiceNow.IsNull() || m.ServiceNow.IsUnknown() {
		return e, false
	}
	var v []ServiceNowExportConfig_SdkV2
	d := m.ServiceNow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServiceNow sets the value of the ServiceNow field in ExportConfig_SdkV2.
func (m *ExportConfig_SdkV2) SetServiceNow(ctx context.Context, v ServiceNowExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_now"]
	m.ServiceNow = types.ListValueMust(t, vs)
}

// GetSlack returns the value of the Slack field in ExportConfig_SdkV2 as
// a SlackExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExportConfig_SdkV2) GetSlack(ctx context.Context) (SlackExportConfig_SdkV2, bool) {
	var e SlackExportConfig_SdkV2
	if m.Slack.IsNull() || m.Slack.IsUnknown() {
		return e, false
	}
	var v []SlackExportConfig_SdkV2
	d := m.Slack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSlack sets the value of the Slack field in ExportConfig_SdkV2.
func (m *ExportConfig_SdkV2) SetSlack(ctx context.Context, v SlackExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["slack"]
	m.Slack = types.ListValueMust(t, vs)
}

// GetTines returns the value of the Tines field in ExportConfig_SdkV2 as
// a TinesExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExportConfig_SdkV2) GetTines(ctx context.Context) (TinesExportConfig_SdkV2, bool) {
	var e TinesExportConfig_SdkV2
	if m.Tines.IsNull() || m.Tines.IsUnknown() {
		return e, false
	}
	var v []TinesExportConfig_SdkV2
	d := m.Tines.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTines sets the value of the Tines field in ExportConfig_SdkV2.
func (m *ExportConfig_SdkV2) SetTines(ctx context.Context, v TinesExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tines"]
	m.Tines = types.ListValueMust(t, vs)
}

// GetWebhook returns the value of the Webhook field in ExportConfig_SdkV2 as
// a WebhookExportConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExportConfig_SdkV2) GetWebhook(ctx context.Context) (WebhookExportConfig_SdkV2, bool) {
	var e WebhookExportConfig_SdkV2
	if m.Webhook.IsNull() || m.Webhook.IsUnknown() {
		return e, false
	}
	var v []WebhookExportConfig_SdkV2
	d := m.Webhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in ExportConfig_SdkV2.
func (m *ExportConfig_SdkV2) SetWebhook(ctx context.Context, v WebhookExportConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook"]
	m.Webhook = types.ListValueMust(t, vs)
}

type GetAlertConfigRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetAlertConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertConfigRequest_SdkV2) {
}

func (to *GetAlertConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAlertConfigRequest_SdkV2) {
}

func (m GetAlertConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAlertConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAlertConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAlertConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetCommonConfigRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetCommonConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCommonConfigRequest_SdkV2) {
}

func (to *GetCommonConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCommonConfigRequest_SdkV2) {
}

func (m GetCommonConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCommonConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCommonConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCommonConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCommonConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCommonConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatasourceConfigRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatasourceConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatasourceConfigRequest_SdkV2) {
}

func (to *GetDatasourceConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatasourceConfigRequest_SdkV2) {
}

func (m GetDatasourceConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatasourceConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatasourceConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatasourceConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatasourceConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatasourceConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatasourceRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatasourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatasourceRequest_SdkV2) {
}

func (to *GetDatasourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatasourceRequest_SdkV2) {
}

func (m GetDatasourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatasourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatasourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatasourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatasourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatasourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRuleConfigRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetRuleConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRuleConfigRequest_SdkV2) {
}

func (to *GetRuleConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRuleConfigRequest_SdkV2) {
}

func (m GetRuleConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRuleConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRuleConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRuleConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRuleConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRuleConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRuleRequest_SdkV2 struct {
	// Full resource name of the detection rule, in the form `rules/{rule_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRuleRequest_SdkV2) {
}

func (to *GetRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRuleRequest_SdkV2) {
}

func (m GetRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type JiraExportConfig_SdkV2 struct {
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Jira API token, paired with `user_email` as the Jira Cloud
	// basic-auth credential (`email:token`). Only the name is stored; the token
	// is read when an alert is exported, so rotating the secret takes effect
	// without editing this config.
	ApiTokenUcSecret types.String `tfsdk:"api_token_uc_secret"`
	// Default issue type for created tickets (e.g. `Task`, `Incident`). Jira
	// issue creation requires an issue type, so it is validated at config-save
	// time rather than failing the export at Jira later.
	DefaultIssueType types.String `tfsdk:"default_issue_type"`
	// Jira instance URL (e.g. `https://acme.atlassian.net`).
	InstanceUrl types.String `tfsdk:"instance_url"`
	// Jira project key the exported issues are created under.
	ProjectKey types.String `tfsdk:"project_key"`
	// User email paired with the API token for Jira Cloud basic auth.
	UserEmail types.String `tfsdk:"user_email"`
}

func (to *JiraExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JiraExportConfig_SdkV2) {
}

func (to *JiraExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from JiraExportConfig_SdkV2) {
}

func (m JiraExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["api_token_uc_secret"] = attrs["api_token_uc_secret"].SetRequired()
	attrs["default_issue_type"] = attrs["default_issue_type"].SetRequired()
	attrs["instance_url"] = attrs["instance_url"].SetRequired()
	attrs["project_key"] = attrs["project_key"].SetRequired()
	attrs["user_email"] = attrs["user_email"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JiraExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m JiraExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JiraExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m JiraExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"api_token_uc_secret": m.ApiTokenUcSecret,
			"default_issue_type":  m.DefaultIssueType,
			"instance_url":        m.InstanceUrl,
			"project_key":         m.ProjectKey,
			"user_email":          m.UserEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JiraExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"api_token_uc_secret": types.StringType,
			"default_issue_type":  types.StringType,
			"instance_url":        types.StringType,
			"project_key":         types.StringType,
			"user_email":          types.StringType,
		},
	}
}

// Unwraps a JSON array of events into one row per array element.
type JsonArrayEnvelope_SdkV2 struct {
	// Name of the root-level key that holds the array of events, for example
	// `Records`. Nested paths are not supported.
	Path types.String `tfsdk:"path"`
}

func (to *JsonArrayEnvelope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JsonArrayEnvelope_SdkV2) {
}

func (to *JsonArrayEnvelope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from JsonArrayEnvelope_SdkV2) {
}

func (m JsonArrayEnvelope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JsonArrayEnvelope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m JsonArrayEnvelope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JsonArrayEnvelope_SdkV2
// only implements ToObjectValue() and Type().
func (m JsonArrayEnvelope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JsonArrayEnvelope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type ListDatasourcesRequest_SdkV2 struct {
	// Name of the Unity Catalog catalog to list datasources from. Required;
	// each request lists the datasources of a single catalog.
	CatalogName types.String `tfsdk:"-"`
	// Filter expression that restricts the results to datasources with a given
	// name, for example `datasource_name = "my_source"`. Only a single
	// exact-match comparison on `datasource_name` is supported, and it matches
	// the datasource's current Unity Catalog name, which reads report in
	// `effective_datasource_name`. Filtered results are returned in a single
	// page.
	Filter types.String `tfsdk:"-"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Name of a Unity Catalog schema in `catalog_name`. Optional; if set, only
	// the datasources in this schema are returned, and if omitted, the
	// datasources in all schemas of the catalog are returned.
	SchemaName types.String `tfsdk:"-"`
	// Which fields to return for each datasource. Defaults to
	// `DATASOURCE_VIEW_FULL`. `DATASOURCE_VIEW_BASIC` returns only the identity
	// and Unity Catalog location fields.
	View types.String `tfsdk:"-"`
}

func (to *ListDatasourcesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatasourcesRequest_SdkV2) {
}

func (to *ListDatasourcesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatasourcesRequest_SdkV2) {
}

func (m ListDatasourcesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatasourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatasourcesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatasourcesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatasourcesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": m.CatalogName,
			"filter":       m.Filter,
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"schema_name":  m.SchemaName,
			"view":         m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatasourcesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"filter":       types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"schema_name":  types.StringType,
			"view":         types.StringType,
		},
	}
}

type ListDatasourcesResponse_SdkV2 struct {
	Datasources types.List `tfsdk:"datasources"`
	// Opaque token for fetching the next page; pass it as `page_token` on the
	// next call. It is empty only when there are no more results. Page size is
	// best-effort, so a short or even empty page does not mean the end; keep
	// paging until this token is empty, which is the only reliable
	// end-of-results signal.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatasourcesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatasourcesResponse_SdkV2) {
	if !from.Datasources.IsNull() && !from.Datasources.IsUnknown() && to.Datasources.IsNull() && len(from.Datasources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasources = from.Datasources
	}
	if !from.Datasources.IsNull() && !from.Datasources.IsUnknown() {
		if toDatasources, ok := to.GetDatasources(ctx); ok {
			if fromDatasources, ok := from.GetDatasources(ctx); ok {
				// Recursively sync the fields of each Datasources element by position.
				for i := range toDatasources {
					if i < len(fromDatasources) {
						toDatasources[i].SyncFieldsDuringCreateOrUpdate(ctx, fromDatasources[i])
					}
				}
				to.SetDatasources(ctx, toDatasources)
			}
		}
	}
}

func (to *ListDatasourcesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatasourcesResponse_SdkV2) {
	if !from.Datasources.IsNull() && !from.Datasources.IsUnknown() && to.Datasources.IsNull() && len(from.Datasources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasources = from.Datasources
	}
	if !from.Datasources.IsNull() && !from.Datasources.IsUnknown() {
		if toDatasources, ok := to.GetDatasources(ctx); ok {
			if fromDatasources, ok := from.GetDatasources(ctx); ok {
				for i := range toDatasources {
					if i < len(fromDatasources) {
						toDatasources[i].SyncFieldsDuringRead(ctx, fromDatasources[i])
					}
				}
				to.SetDatasources(ctx, toDatasources)
			}
		}
	}
}

func (m ListDatasourcesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasources"] = attrs["datasources"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatasourcesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatasourcesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasources": reflect.TypeOf(Datasource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatasourcesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatasourcesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasources":     m.Datasources,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatasourcesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasources": basetypes.ListType{
				ElemType: Datasource_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatasources returns the value of the Datasources field in ListDatasourcesResponse_SdkV2 as
// a slice of Datasource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatasourcesResponse_SdkV2) GetDatasources(ctx context.Context) ([]Datasource_SdkV2, bool) {
	if m.Datasources.IsNull() || m.Datasources.IsUnknown() {
		return nil, false
	}
	var v []Datasource_SdkV2
	d := m.Datasources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasources sets the value of the Datasources field in ListDatasourcesResponse_SdkV2.
func (m *ListDatasourcesResponse_SdkV2) SetDatasources(ctx context.Context, v []Datasource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Datasources = types.ListValueMust(t, vs)
}

type ListRulesRequest_SdkV2 struct {
	// Name of the Unity Catalog catalog to list rules from. Required.
	CatalogName types.String `tfsdk:"-"`
	// Filter expression that restricts the results, such as `rule_name =
	// "my_rule"` to return only the rule with that name. The name is matched
	// against the current UC name, which reads report in `effective_rule_name`.
	Filter types.String `tfsdk:"-"`
	// Maximum number of rules to return. The server may return fewer. If unset
	// or 0, the server picks a default page size.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous `ListRules` response, used to fetch the next
	// page of results.
	PageToken types.String `tfsdk:"-"`
	// Name of a Unity Catalog schema in `catalog_name`. If set, only rules in
	// this schema are returned. If unset, rules in all schemas of the catalog
	// are returned.
	SchemaName types.String `tfsdk:"-"`
	// How much of each rule to return. Defaults to `RULE_VIEW_FULL`.
	// `RULE_VIEW_BASIC` returns only each rule's identity and Unity Catalog
	// location.
	View types.String `tfsdk:"-"`
}

func (to *ListRulesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRulesRequest_SdkV2) {
}

func (to *ListRulesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRulesRequest_SdkV2) {
}

func (m ListRulesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": m.CatalogName,
			"filter":       m.Filter,
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"schema_name":  m.SchemaName,
			"view":         m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"filter":       types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"schema_name":  types.StringType,
			"view":         types.StringType,
		},
	}
}

type ListRulesResponse_SdkV2 struct {
	// Token to pass as `page_token` in a subsequent `ListRules` call to fetch
	// the next page, or empty when there are no more results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The rules in this page of results.
	Rules types.List `tfsdk:"rules"`
}

func (to *ListRulesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRulesResponse_SdkV2) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() {
		if toRules, ok := to.GetRules(ctx); ok {
			if fromRules, ok := from.GetRules(ctx); ok {
				// Recursively sync the fields of each Rules element by position.
				for i := range toRules {
					if i < len(fromRules) {
						toRules[i].SyncFieldsDuringCreateOrUpdate(ctx, fromRules[i])
					}
				}
				to.SetRules(ctx, toRules)
			}
		}
	}
}

func (to *ListRulesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRulesResponse_SdkV2) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() {
		if toRules, ok := to.GetRules(ctx); ok {
			if fromRules, ok := from.GetRules(ctx); ok {
				for i := range toRules {
					if i < len(fromRules) {
						toRules[i].SyncFieldsDuringRead(ctx, fromRules[i])
					}
				}
				to.SetRules(ctx, toRules)
			}
		}
	}
}

func (m ListRulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["rules"] = attrs["rules"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(Rule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"rules":           m.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"rules": basetypes.ListType{
				ElemType: Rule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRules returns the value of the Rules field in ListRulesResponse_SdkV2 as
// a slice of Rule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRulesResponse_SdkV2) GetRules(ctx context.Context) ([]Rule_SdkV2, bool) {
	if m.Rules.IsNull() || m.Rules.IsUnknown() {
		return nil, false
	}
	var v []Rule_SdkV2
	d := m.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in ListRulesResponse_SdkV2.
func (m *ListRulesResponse_SdkV2) SetRules(ctx context.Context, v []Rule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Rules = types.ListValueMust(t, vs)
}

// Reference to a parser that the datasource uses. Each parser defines the
// schema of the parsed table that it writes.
type ParserRef_SdkV2 struct {
	// The table this parser's data is written to: `target_table` when supplied,
	// otherwise the table derived from the datasource's catalog and schema and
	// the parser name. Fixed once resolved, so a later parser rename does not
	// move it. Populated on get, create, and update responses with the resolved
	// table.
	EffectiveTargetTable types.List `tfsdk:"effective_target_table"`
	// Identifier of the parser. Stable across renames, unlike a parser name.
	Id types.String `tfsdk:"id"`
	// Table that this parser's data is written to. The table must be in the
	// datasource's catalog and schema. If unset, the table name is derived from
	// the parser name. Reads return the value as supplied; the table the data
	// is written to is reported in `effective_target_table`.
	TargetTable types.List `tfsdk:"target_table"`
	// Version of the parser to use, given as an explicit version number.
	Version types.String `tfsdk:"version"`
}

func (to *ParserRef_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ParserRef_SdkV2) {
	if !from.EffectiveTargetTable.IsNull() && !from.EffectiveTargetTable.IsUnknown() {
		if toEffectiveTargetTable, ok := to.GetEffectiveTargetTable(ctx); ok {
			if fromEffectiveTargetTable, ok := from.GetEffectiveTargetTable(ctx); ok {
				// Recursively sync the fields of EffectiveTargetTable
				toEffectiveTargetTable.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveTargetTable)
				to.SetEffectiveTargetTable(ctx, toEffectiveTargetTable)
			}
		}
	}
	if !from.TargetTable.IsNull() && !from.TargetTable.IsUnknown() {
		if toTargetTable, ok := to.GetTargetTable(ctx); ok {
			if fromTargetTable, ok := from.GetTargetTable(ctx); ok {
				// Recursively sync the fields of TargetTable
				toTargetTable.SyncFieldsDuringCreateOrUpdate(ctx, fromTargetTable)
				to.SetTargetTable(ctx, toTargetTable)
			}
		}
	}
}

func (to *ParserRef_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ParserRef_SdkV2) {
	if !from.EffectiveTargetTable.IsNull() && !from.EffectiveTargetTable.IsUnknown() {
		if toEffectiveTargetTable, ok := to.GetEffectiveTargetTable(ctx); ok {
			if fromEffectiveTargetTable, ok := from.GetEffectiveTargetTable(ctx); ok {
				toEffectiveTargetTable.SyncFieldsDuringRead(ctx, fromEffectiveTargetTable)
				to.SetEffectiveTargetTable(ctx, toEffectiveTargetTable)
			}
		}
	}
	if !from.TargetTable.IsNull() && !from.TargetTable.IsUnknown() {
		if toTargetTable, ok := to.GetTargetTable(ctx); ok {
			if fromTargetTable, ok := from.GetTargetTable(ctx); ok {
				toTargetTable.SyncFieldsDuringRead(ctx, fromTargetTable)
				to.SetTargetTable(ctx, toTargetTable)
			}
		}
	}
}

func (m ParserRef_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_target_table"] = attrs["effective_target_table"].SetComputed()
	attrs["effective_target_table"] = attrs["effective_target_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()
	attrs["target_table"] = attrs["target_table"].SetOptional()
	attrs["target_table"] = attrs["target_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ParserRef.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ParserRef_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_target_table": reflect.TypeOf(TableName_SdkV2{}),
		"target_table":           reflect.TypeOf(TableName_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ParserRef_SdkV2
// only implements ToObjectValue() and Type().
func (m ParserRef_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_target_table": m.EffectiveTargetTable,
			"id":                     m.Id,
			"target_table":           m.TargetTable,
			"version":                m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ParserRef_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_target_table": basetypes.ListType{
				ElemType: TableName_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"target_table": basetypes.ListType{
				ElemType: TableName_SdkV2{}.Type(ctx),
			},
			"version": types.StringType,
		},
	}
}

// GetEffectiveTargetTable returns the value of the EffectiveTargetTable field in ParserRef_SdkV2 as
// a TableName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ParserRef_SdkV2) GetEffectiveTargetTable(ctx context.Context) (TableName_SdkV2, bool) {
	var e TableName_SdkV2
	if m.EffectiveTargetTable.IsNull() || m.EffectiveTargetTable.IsUnknown() {
		return e, false
	}
	var v []TableName_SdkV2
	d := m.EffectiveTargetTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveTargetTable sets the value of the EffectiveTargetTable field in ParserRef_SdkV2.
func (m *ParserRef_SdkV2) SetEffectiveTargetTable(ctx context.Context, v TableName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_target_table"]
	m.EffectiveTargetTable = types.ListValueMust(t, vs)
}

// GetTargetTable returns the value of the TargetTable field in ParserRef_SdkV2 as
// a TableName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ParserRef_SdkV2) GetTargetTable(ctx context.Context) (TableName_SdkV2, bool) {
	var e TableName_SdkV2
	if m.TargetTable.IsNull() || m.TargetTable.IsUnknown() {
		return e, false
	}
	var v []TableName_SdkV2
	d := m.TargetTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTargetTable sets the value of the TargetTable field in ParserRef_SdkV2.
func (m *ParserRef_SdkV2) SetTargetTable(ctx context.Context, v TableName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["target_table"]
	m.TargetTable = types.ListValueMust(t, vs)
}

// A principal — user or service-principal — used as the value type for
// `created_by` / `updated_by` / `assignee` etc. across v2 resources.
type Principal_SdkV2 struct {
	// Display name (user or SP name).
	DisplayName types.String `tfsdk:"display_name"`
	// Email (regular users only).
	Email types.String `tfsdk:"email"`
	// User / service-principal ID.
	Id types.String `tfsdk:"id"`
	// Whether this principal is a regular user, service principal, or system
	// principal.
	PrincipalType types.String `tfsdk:"principal_type"`
}

func (to *Principal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Principal_SdkV2) {
}

func (to *Principal_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Principal_SdkV2) {
}

func (m Principal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["email"] = attrs["email"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["principal_type"] = attrs["principal_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Principal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Principal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Principal_SdkV2
// only implements ToObjectValue() and Type().
func (m Principal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":   m.DisplayName,
			"email":          m.Email,
			"id":             m.Id,
			"principal_type": m.PrincipalType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Principal_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":   types.StringType,
			"email":          types.StringType,
			"id":             types.StringType,
			"principal_type": types.StringType,
		},
	}
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
type ResourceStatus_SdkV2 struct {
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Databricks Jobs API job ID. Unset for a pure ingestion pipeline.
	JobId types.String `tfsdk:"job_id"`
	// When the most recent run completed. Pairs with `run_state = SUCCEEDED`.
	LastCompleteTime timetypes.RFC3339 `tfsdk:"last_complete_time"`
	// Most recent error message, if any. Pairs with `run_state = FAILED`.
	LastErrorMessage types.String `tfsdk:"last_error_message"`
	// When the most recent run started. Pairs with `run_state = RUNNING`.
	LastStartTime timetypes.RFC3339 `tfsdk:"last_start_time"`

	LastStatusChangeTime timetypes.RFC3339 `tfsdk:"last_status_change_time"`
	// Scheduling lifecycle of the resource's underlying job/pipeline: whether
	// it is unscheduled, scheduled, stopped, or deleted. Distinct from
	// `run_state`, which reports the outcome of the most recent run.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// DLT ingestion-pipeline ID. Set only when this status represents the
	// datasource's ingestion pipeline (unset for job statuses).
	PipelineId types.String `tfsdk:"pipeline_id"`

	RunState types.String `tfsdk:"run_state"`
}

func (to *ResourceStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResourceStatus_SdkV2) {
}

func (to *ResourceStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResourceStatus_SdkV2) {
}

func (m ResourceStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["job_id"] = attrs["job_id"].SetComputed()
	attrs["last_complete_time"] = attrs["last_complete_time"].SetComputed()
	attrs["last_error_message"] = attrs["last_error_message"].SetComputed()
	attrs["last_start_time"] = attrs["last_start_time"].SetComputed()
	attrs["last_status_change_time"] = attrs["last_status_change_time"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetComputed()
	attrs["run_state"] = attrs["run_state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResourceStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResourceStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResourceStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m ResourceStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":             m.CreateTime,
			"job_id":                  m.JobId,
			"last_complete_time":      m.LastCompleteTime,
			"last_error_message":      m.LastErrorMessage,
			"last_start_time":         m.LastStartTime,
			"last_status_change_time": m.LastStatusChangeTime,
			"lifecycle_state":         m.LifecycleState,
			"pipeline_id":             m.PipelineId,
			"run_state":               m.RunState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResourceStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":             timetypes.RFC3339{}.Type(ctx),
			"job_id":                  types.StringType,
			"last_complete_time":      timetypes.RFC3339{}.Type(ctx),
			"last_error_message":      types.StringType,
			"last_start_time":         timetypes.RFC3339{}.Type(ctx),
			"last_status_change_time": timetypes.RFC3339{}.Type(ctx),
			"lifecycle_state":         types.StringType,
			"pipeline_id":             types.StringType,
			"run_state":               types.StringType,
		},
	}
}

// A Lakewatch detection rule: a query that evaluates ingested data and
// generates a signal for each match. Each rule is a securable in a Unity
// Catalog schema.
type Rule_SdkV2 struct {
	// Free-form client-managed key/value annotations.
	Annotations types.Map `tfsdk:"annotations"`
	// Batch SQL rule input.
	Batch types.List `tfsdk:"batch"`
	// Name of the Unity Catalog catalog that contains the rule. Required.
	// Cannot be changed after creation. Reads return the value as supplied; the
	// catalog's current name is reported in `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Free-form rule category.
	Category types.String `tfsdk:"category"`
	// ID of the workspace whose compute runs this rule's jobs and owns its
	// associated resources, such as the SQL warehouse a scheduled rule runs on.
	// Server-assigned; a rule's warehouse can only be changed from this
	// workspace.
	ComputeWorkspaceId types.String `tfsdk:"compute_workspace_id"`
	// Immutable UUIDv4 identifying the authored content resource installed as
	// this rule.
	ContentResourceId types.String `tfsdk:"content_resource_id"`
	// Time the rule was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Principal that created the rule.
	CreatedBy types.List `tfsdk:"created_by"`
	// Signal deduplication mode.
	Deduplication types.String `tfsdk:"deduplication"`
	// Time window for deduplication, such as `1h`, `24h`, `7d`, or `automatic`.
	DeduplicationTimeWindow types.String `tfsdk:"deduplication_time_window"`
	// Free-form description.
	Description types.String `tfsdk:"description"`
	// Current name of the Unity Catalog catalog that contains the rule.
	// Reflects catalog renames, which do not change `catalog_name`. Unset until
	// resolved.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// Current UC leaf name of the rule (server-owned). Tracks UC renames so
	// callers can observe the live name without having to update the
	// client-owned `rule_name`; unset until resolved.
	EffectiveRuleName types.String `tfsdk:"effective_rule_name"`
	// The identity this rule's job runs as: the resolved `run_as_principal`, or
	// the requester's identity when it was omitted on create or cleared on
	// update. The value is a user email or service-principal application ID,
	// falling back to the numeric principal ID if that identity can no longer
	// be resolved. Populated on single-rule reads (get/create/update); list
	// responses omit it.
	EffectiveRunAsPrincipal types.String `tfsdk:"effective_run_as_principal"`
	// Current name of the Unity Catalog schema that contains the rule. Reflects
	// schema renames, which do not change `schema_name`. Unset until resolved.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Opaque entity tag for optimistic concurrency. Content-version ETags
	// change on user edits and remain stable across runtime status updates.
	Etag types.String `tfsdk:"etag"`
	// Signal quality for this rule.
	Fidelity types.String `tfsdk:"fidelity"`
	// Up to 64 MITRE ATT&CK mappings, producing at most 192 flattened alert
	// tags.
	Mitre types.List `tfsdk:"mitre"`
	// Resource name, in the form `rules/{rule_id}`. Assigned by the server on
	// create; required on update.
	Name types.String `tfsdk:"name"`
	// Long-form description of what this rule detects.
	Objective types.String `tfsdk:"objective"`
	// Reference to the content pack this rule was originally installed from.
	OriginalContentPack types.List `tfsdk:"original_content_pack"`
	// What to record in each signal.
	Output types.List `tfsdk:"output"`
	// Unity Catalog owner of the rule. Assigned by the server. Populated by
	// `GetRule`.
	Owner types.List `tfsdk:"owner"`
	// Ordered reference URLs carried from the authored rule metadata. A maximum
	// of three references of up to 512 characters each is supported. Requests
	// exceeding these limits are rejected.
	References types.List `tfsdk:"references"`
	// Unity Catalog securable UUID of the rule, and the last segment of `name`.
	// Never changes, including when the rule's catalog or schema is renamed.
	RuleId types.String `tfsdk:"rule_id"`
	// Name of the rule within its Unity Catalog schema. Required on create; can
	// be updated within the same catalog and schema. Get synchronizes this
	// field with the current Unity Catalog name, also reported in
	// `effective_rule_name`.
	RuleName types.String `tfsdk:"rule_name"`
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
	// update whose mask includes `run_as_principal` but supplies no value
	// clears it: the rule then runs as the requester, exactly as when it is
	// omitted on create. Reads return the value as supplied; the identity the
	// rule runs as is reported in `effective_run_as_principal`.
	RunAsPrincipal types.String `tfsdk:"run_as_principal"`
	// Long-form operational guidance carried from the authored rule metadata.
	// Supports up to 4096 characters. Requests exceeding this limit are
	// rejected. Supports Markdown; analysts and the investigation agent follow
	// it when triaging this rule's alerts.
	Runbook types.String `tfsdk:"runbook"`
	// When the rule executes.
	Schedule types.List `tfsdk:"schedule"`
	// Name of the Unity Catalog schema that contains the rule. Required. Cannot
	// be changed after creation. Reads return the value as supplied; the
	// schema's current name is reported in `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Free-form severity label, such as `Informational`, `Medium`, `High`, or
	// `Critical`.
	Severity types.String `tfsdk:"severity"`
	// Read-only lifecycle status.
	Status types.List `tfsdk:"status"`
	// If set, this rule's signals do not open new alerts until the given time.
	SuppressUntilTime timetypes.RFC3339 `tfsdk:"suppress_until_time"`
	// Time-range configuration for batch rules. If unset, the execution time
	// range is extended by a fixed `15m` lookback.
	Timerange types.List `tfsdk:"timerange"`
	// Time the rule was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Principal that last updated the rule.
	UpdatedBy types.List `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version types.Int64 `tfsdk:"version"`
}

func (to *Rule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Rule_SdkV2) {
	if !from.Batch.IsNull() && !from.Batch.IsUnknown() {
		if toBatch, ok := to.GetBatch(ctx); ok {
			if fromBatch, ok := from.GetBatch(ctx); ok {
				// Recursively sync the fields of Batch
				toBatch.SyncFieldsDuringCreateOrUpdate(ctx, fromBatch)
				to.SetBatch(ctx, toBatch)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				// Recursively sync the fields of CreatedBy
				toCreatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() && to.Mitre.IsNull() && len(from.Mitre.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Mitre, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Mitre = from.Mitre
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() {
		if toMitre, ok := to.GetMitre(ctx); ok {
			if fromMitre, ok := from.GetMitre(ctx); ok {
				// Recursively sync the fields of each Mitre element by position.
				for i := range toMitre {
					if i < len(fromMitre) {
						toMitre[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMitre[i])
					}
				}
				to.SetMitre(ctx, toMitre)
			}
		}
	}
	if !from.OriginalContentPack.IsNull() && !from.OriginalContentPack.IsUnknown() {
		if toOriginalContentPack, ok := to.GetOriginalContentPack(ctx); ok {
			if fromOriginalContentPack, ok := from.GetOriginalContentPack(ctx); ok {
				// Recursively sync the fields of OriginalContentPack
				toOriginalContentPack.SyncFieldsDuringCreateOrUpdate(ctx, fromOriginalContentPack)
				to.SetOriginalContentPack(ctx, toOriginalContentPack)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				// Recursively sync the fields of Output
				toOutput.SyncFieldsDuringCreateOrUpdate(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				// Recursively sync the fields of Owner
				toOwner.SyncFieldsDuringCreateOrUpdate(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.References.IsNull() && !from.References.IsUnknown() && to.References.IsNull() && len(from.References.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for References, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.References = from.References
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				// Recursively sync the fields of Schedule
				toSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.Timerange.IsNull() && !from.Timerange.IsUnknown() {
		if toTimerange, ok := to.GetTimerange(ctx); ok {
			if fromTimerange, ok := from.GetTimerange(ctx); ok {
				// Recursively sync the fields of Timerange
				toTimerange.SyncFieldsDuringCreateOrUpdate(ctx, fromTimerange)
				to.SetTimerange(ctx, toTimerange)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (to *Rule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Rule_SdkV2) {
	if !from.Batch.IsNull() && !from.Batch.IsUnknown() {
		if toBatch, ok := to.GetBatch(ctx); ok {
			if fromBatch, ok := from.GetBatch(ctx); ok {
				toBatch.SyncFieldsDuringRead(ctx, fromBatch)
				to.SetBatch(ctx, toBatch)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				toCreatedBy.SyncFieldsDuringRead(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() && to.Mitre.IsNull() && len(from.Mitre.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Mitre, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Mitre = from.Mitre
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() {
		if toMitre, ok := to.GetMitre(ctx); ok {
			if fromMitre, ok := from.GetMitre(ctx); ok {
				for i := range toMitre {
					if i < len(fromMitre) {
						toMitre[i].SyncFieldsDuringRead(ctx, fromMitre[i])
					}
				}
				to.SetMitre(ctx, toMitre)
			}
		}
	}
	if !from.OriginalContentPack.IsNull() && !from.OriginalContentPack.IsUnknown() {
		if toOriginalContentPack, ok := to.GetOriginalContentPack(ctx); ok {
			if fromOriginalContentPack, ok := from.GetOriginalContentPack(ctx); ok {
				toOriginalContentPack.SyncFieldsDuringRead(ctx, fromOriginalContentPack)
				to.SetOriginalContentPack(ctx, toOriginalContentPack)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				toOutput.SyncFieldsDuringRead(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				toOwner.SyncFieldsDuringRead(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.References.IsNull() && !from.References.IsUnknown() && to.References.IsNull() && len(from.References.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for References, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.References = from.References
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				toSchedule.SyncFieldsDuringRead(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.Timerange.IsNull() && !from.Timerange.IsUnknown() {
		if toTimerange, ok := to.GetTimerange(ctx); ok {
			if fromTimerange, ok := from.GetTimerange(ctx); ok {
				toTimerange.SyncFieldsDuringRead(ctx, fromTimerange)
				to.SetTimerange(ctx, toTimerange)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (m Rule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetOptional()
	attrs["batch"] = attrs["batch"].SetOptional()
	attrs["batch"] = attrs["batch"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["category"] = attrs["category"].SetOptional()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["content_resource_id"] = attrs["content_resource_id"].SetOptional()
	attrs["content_resource_id"] = attrs["content_resource_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["created_by"] = attrs["created_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["deduplication"] = attrs["deduplication"].SetOptional()
	attrs["deduplication_time_window"] = attrs["deduplication_time_window"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_rule_name"] = attrs["effective_rule_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["fidelity"] = attrs["fidelity"].SetRequired()
	attrs["mitre"] = attrs["mitre"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["objective"] = attrs["objective"].SetOptional()
	attrs["original_content_pack"] = attrs["original_content_pack"].SetComputed()
	attrs["original_content_pack"] = attrs["original_content_pack"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["output"] = attrs["output"].SetRequired()
	attrs["output"] = attrs["output"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["owner"] = attrs["owner"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["references"] = attrs["references"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()
	attrs["rule_name"] = attrs["rule_name"].SetRequired()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetOptional()
	attrs["runbook"] = attrs["runbook"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetRequired()
	attrs["schedule"] = attrs["schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["severity"] = attrs["severity"].SetRequired()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["suppress_until_time"] = attrs["suppress_until_time"].SetOptional()
	attrs["timerange"] = attrs["timerange"].SetOptional()
	attrs["timerange"] = attrs["timerange"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Rule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Rule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"annotations":           reflect.TypeOf(types.String{}),
		"batch":                 reflect.TypeOf(RuleInputBatch_SdkV2{}),
		"created_by":            reflect.TypeOf(Principal_SdkV2{}),
		"mitre":                 reflect.TypeOf(RuleMitreMapping_SdkV2{}),
		"original_content_pack": reflect.TypeOf(ContentPackReference_SdkV2{}),
		"output":                reflect.TypeOf(RuleOutput_SdkV2{}),
		"owner":                 reflect.TypeOf(Principal_SdkV2{}),
		"references":            reflect.TypeOf(types.String{}),
		"schedule":              reflect.TypeOf(RuleSchedule_SdkV2{}),
		"status":                reflect.TypeOf(ResourceStatus_SdkV2{}),
		"timerange":             reflect.TypeOf(RuleTimeRange_SdkV2{}),
		"updated_by":            reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Rule_SdkV2
// only implements ToObjectValue() and Type().
func (m Rule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"annotations":                m.Annotations,
			"batch":                      m.Batch,
			"catalog_name":               m.CatalogName,
			"category":                   m.Category,
			"compute_workspace_id":       m.ComputeWorkspaceId,
			"content_resource_id":        m.ContentResourceId,
			"create_time":                m.CreateTime,
			"created_by":                 m.CreatedBy,
			"deduplication":              m.Deduplication,
			"deduplication_time_window":  m.DeduplicationTimeWindow,
			"description":                m.Description,
			"effective_catalog_name":     m.EffectiveCatalogName,
			"effective_rule_name":        m.EffectiveRuleName,
			"effective_run_as_principal": m.EffectiveRunAsPrincipal,
			"effective_schema_name":      m.EffectiveSchemaName,
			"etag":                       m.Etag,
			"fidelity":                   m.Fidelity,
			"mitre":                      m.Mitre,
			"name":                       m.Name,
			"objective":                  m.Objective,
			"original_content_pack":      m.OriginalContentPack,
			"output":                     m.Output,
			"owner":                      m.Owner,
			"references":                 m.References,
			"rule_id":                    m.RuleId,
			"rule_name":                  m.RuleName,
			"run_as_principal":           m.RunAsPrincipal,
			"runbook":                    m.Runbook,
			"schedule":                   m.Schedule,
			"schema_name":                m.SchemaName,
			"severity":                   m.Severity,
			"status":                     m.Status,
			"suppress_until_time":        m.SuppressUntilTime,
			"timerange":                  m.Timerange,
			"update_time":                m.UpdateTime,
			"updated_by":                 m.UpdatedBy,
			"version":                    m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Rule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"annotations": basetypes.MapType{
				ElemType: types.StringType,
			},
			"batch": basetypes.ListType{
				ElemType: RuleInputBatch_SdkV2{}.Type(ctx),
			},
			"catalog_name":         types.StringType,
			"category":             types.StringType,
			"compute_workspace_id": types.StringType,
			"content_resource_id":  types.StringType,
			"create_time":          timetypes.RFC3339{}.Type(ctx),
			"created_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"deduplication":              types.StringType,
			"deduplication_time_window":  types.StringType,
			"description":                types.StringType,
			"effective_catalog_name":     types.StringType,
			"effective_rule_name":        types.StringType,
			"effective_run_as_principal": types.StringType,
			"effective_schema_name":      types.StringType,
			"etag":                       types.StringType,
			"fidelity":                   types.StringType,
			"mitre": basetypes.ListType{
				ElemType: RuleMitreMapping_SdkV2{}.Type(ctx),
			},
			"name":      types.StringType,
			"objective": types.StringType,
			"original_content_pack": basetypes.ListType{
				ElemType: ContentPackReference_SdkV2{}.Type(ctx),
			},
			"output": basetypes.ListType{
				ElemType: RuleOutput_SdkV2{}.Type(ctx),
			},
			"owner": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"references": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rule_id":          types.StringType,
			"rule_name":        types.StringType,
			"run_as_principal": types.StringType,
			"runbook":          types.StringType,
			"schedule": basetypes.ListType{
				ElemType: RuleSchedule_SdkV2{}.Type(ctx),
			},
			"schema_name": types.StringType,
			"severity":    types.StringType,
			"status": basetypes.ListType{
				ElemType: ResourceStatus_SdkV2{}.Type(ctx),
			},
			"suppress_until_time": timetypes.RFC3339{}.Type(ctx),
			"timerange": basetypes.ListType{
				ElemType: RuleTimeRange_SdkV2{}.Type(ctx),
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetAnnotations returns the value of the Annotations field in Rule_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetAnnotations(ctx context.Context) (map[string]types.String, bool) {
	if m.Annotations.IsNull() || m.Annotations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Annotations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnnotations sets the value of the Annotations field in Rule_SdkV2.
func (m *Rule_SdkV2) SetAnnotations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["annotations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Annotations = types.MapValueMust(t, vs)
}

// GetBatch returns the value of the Batch field in Rule_SdkV2 as
// a RuleInputBatch_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetBatch(ctx context.Context) (RuleInputBatch_SdkV2, bool) {
	var e RuleInputBatch_SdkV2
	if m.Batch.IsNull() || m.Batch.IsUnknown() {
		return e, false
	}
	var v []RuleInputBatch_SdkV2
	d := m.Batch.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBatch sets the value of the Batch field in Rule_SdkV2.
func (m *Rule_SdkV2) SetBatch(ctx context.Context, v RuleInputBatch_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["batch"]
	m.Batch = types.ListValueMust(t, vs)
}

// GetCreatedBy returns the value of the CreatedBy field in Rule_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetCreatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.CreatedBy.IsNull() || m.CreatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.CreatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCreatedBy sets the value of the CreatedBy field in Rule_SdkV2.
func (m *Rule_SdkV2) SetCreatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["created_by"]
	m.CreatedBy = types.ListValueMust(t, vs)
}

// GetMitre returns the value of the Mitre field in Rule_SdkV2 as
// a slice of RuleMitreMapping_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetMitre(ctx context.Context) ([]RuleMitreMapping_SdkV2, bool) {
	if m.Mitre.IsNull() || m.Mitre.IsUnknown() {
		return nil, false
	}
	var v []RuleMitreMapping_SdkV2
	d := m.Mitre.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMitre sets the value of the Mitre field in Rule_SdkV2.
func (m *Rule_SdkV2) SetMitre(ctx context.Context, v []RuleMitreMapping_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["mitre"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Mitre = types.ListValueMust(t, vs)
}

// GetOriginalContentPack returns the value of the OriginalContentPack field in Rule_SdkV2 as
// a ContentPackReference_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetOriginalContentPack(ctx context.Context) (ContentPackReference_SdkV2, bool) {
	var e ContentPackReference_SdkV2
	if m.OriginalContentPack.IsNull() || m.OriginalContentPack.IsUnknown() {
		return e, false
	}
	var v []ContentPackReference_SdkV2
	d := m.OriginalContentPack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOriginalContentPack sets the value of the OriginalContentPack field in Rule_SdkV2.
func (m *Rule_SdkV2) SetOriginalContentPack(ctx context.Context, v ContentPackReference_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["original_content_pack"]
	m.OriginalContentPack = types.ListValueMust(t, vs)
}

// GetOutput returns the value of the Output field in Rule_SdkV2 as
// a RuleOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetOutput(ctx context.Context) (RuleOutput_SdkV2, bool) {
	var e RuleOutput_SdkV2
	if m.Output.IsNull() || m.Output.IsUnknown() {
		return e, false
	}
	var v []RuleOutput_SdkV2
	d := m.Output.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutput sets the value of the Output field in Rule_SdkV2.
func (m *Rule_SdkV2) SetOutput(ctx context.Context, v RuleOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["output"]
	m.Output = types.ListValueMust(t, vs)
}

// GetOwner returns the value of the Owner field in Rule_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetOwner(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.Owner.IsNull() || m.Owner.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.Owner.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOwner sets the value of the Owner field in Rule_SdkV2.
func (m *Rule_SdkV2) SetOwner(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["owner"]
	m.Owner = types.ListValueMust(t, vs)
}

// GetReferences returns the value of the References field in Rule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetReferences(ctx context.Context) ([]types.String, bool) {
	if m.References.IsNull() || m.References.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.References.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReferences sets the value of the References field in Rule_SdkV2.
func (m *Rule_SdkV2) SetReferences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["references"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.References = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in Rule_SdkV2 as
// a RuleSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetSchedule(ctx context.Context) (RuleSchedule_SdkV2, bool) {
	var e RuleSchedule_SdkV2
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v []RuleSchedule_SdkV2
	d := m.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in Rule_SdkV2.
func (m *Rule_SdkV2) SetSchedule(ctx context.Context, v RuleSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	m.Schedule = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Rule_SdkV2 as
// a ResourceStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetStatus(ctx context.Context) (ResourceStatus_SdkV2, bool) {
	var e ResourceStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []ResourceStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Rule_SdkV2.
func (m *Rule_SdkV2) SetStatus(ctx context.Context, v ResourceStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// GetTimerange returns the value of the Timerange field in Rule_SdkV2 as
// a RuleTimeRange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetTimerange(ctx context.Context) (RuleTimeRange_SdkV2, bool) {
	var e RuleTimeRange_SdkV2
	if m.Timerange.IsNull() || m.Timerange.IsUnknown() {
		return e, false
	}
	var v []RuleTimeRange_SdkV2
	d := m.Timerange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimerange sets the value of the Timerange field in Rule_SdkV2.
func (m *Rule_SdkV2) SetTimerange(ctx context.Context, v RuleTimeRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["timerange"]
	m.Timerange = types.ListValueMust(t, vs)
}

// GetUpdatedBy returns the value of the UpdatedBy field in Rule_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in Rule_SdkV2.
func (m *Rule_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Per-catalog rule-related config. Singleton. Name format:
// `catalogs/{catalog}/ruleConfig`.
type RuleConfig_SdkV2 struct {
	// Schema name of the registered folder (in this catalog) that a new rule is
	// placed in when the request does not specify a schema. Must name a schema
	// registered as a folder. Omit to clear. Reads return the value as
	// supplied; the folder's current schema name is reported in
	// `effective_default_folder`.
	DefaultFolder types.String `tfsdk:"default_folder"`
	// Default schedule for rules that don't specify one.
	DefaultSchedule types.List `tfsdk:"default_schedule"`
	// Default duration string for signal deduplication when a rule doesn't
	// override it — compact day/hour/minute/second components, e.g. `5d`,
	// `1h`. Reads return the value as supplied; there is no server default, so
	// it is unset until set. An empty value clears it.
	DefaultSignalCollationWindow types.String `tfsdk:"default_signal_collation_window"`
	// Available detection categories a rule can be tagged with. Reads return
	// the value as supplied. When empty, rules can be tagged with the default
	// categories: `Static Signature`, `Fuzzy Signature`, `Behavioral`,
	// `Statistical`, and `Learning (AI/ML)`. The categories in effect are
	// reported in `effective_detection_categories`.
	DetectionCategories types.List `tfsdk:"detection_categories"`
	// The current schema name of the `default_folder` folder. Tracks UC schema
	// renames. Unset when no default is configured, or when the folder's schema
	// no longer exists or is not visible to the caller.
	EffectiveDefaultFolder types.String `tfsdk:"effective_default_folder"`
	// The detection categories in effect: `detection_categories` when set,
	// otherwise the default categories.
	EffectiveDetectionCategories types.List `tfsdk:"effective_detection_categories"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag types.String `tfsdk:"etag"`

	Name types.String `tfsdk:"name"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.List `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version types.Int64 `tfsdk:"version"`
}

func (to *RuleConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleConfig_SdkV2) {
	if !from.DefaultSchedule.IsNull() && !from.DefaultSchedule.IsUnknown() {
		if toDefaultSchedule, ok := to.GetDefaultSchedule(ctx); ok {
			if fromDefaultSchedule, ok := from.GetDefaultSchedule(ctx); ok {
				// Recursively sync the fields of DefaultSchedule
				toDefaultSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultSchedule)
				to.SetDefaultSchedule(ctx, toDefaultSchedule)
			}
		}
	}
	if !from.DetectionCategories.IsNull() && !from.DetectionCategories.IsUnknown() && to.DetectionCategories.IsNull() && len(from.DetectionCategories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DetectionCategories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DetectionCategories = from.DetectionCategories
	}
	if !from.EffectiveDetectionCategories.IsNull() && !from.EffectiveDetectionCategories.IsUnknown() && to.EffectiveDetectionCategories.IsNull() && len(from.EffectiveDetectionCategories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveDetectionCategories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveDetectionCategories = from.EffectiveDetectionCategories
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (to *RuleConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleConfig_SdkV2) {
	if !from.DefaultSchedule.IsNull() && !from.DefaultSchedule.IsUnknown() {
		if toDefaultSchedule, ok := to.GetDefaultSchedule(ctx); ok {
			if fromDefaultSchedule, ok := from.GetDefaultSchedule(ctx); ok {
				toDefaultSchedule.SyncFieldsDuringRead(ctx, fromDefaultSchedule)
				to.SetDefaultSchedule(ctx, toDefaultSchedule)
			}
		}
	}
	if !from.DetectionCategories.IsNull() && !from.DetectionCategories.IsUnknown() && to.DetectionCategories.IsNull() && len(from.DetectionCategories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DetectionCategories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DetectionCategories = from.DetectionCategories
	}
	if !from.EffectiveDetectionCategories.IsNull() && !from.EffectiveDetectionCategories.IsUnknown() && to.EffectiveDetectionCategories.IsNull() && len(from.EffectiveDetectionCategories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveDetectionCategories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveDetectionCategories = from.EffectiveDetectionCategories
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
}

func (m RuleConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_folder"] = attrs["default_folder"].SetOptional()
	attrs["default_schedule"] = attrs["default_schedule"].SetOptional()
	attrs["default_schedule"] = attrs["default_schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["default_signal_collation_window"] = attrs["default_signal_collation_window"].SetOptional()
	attrs["detection_categories"] = attrs["detection_categories"].SetOptional()
	attrs["effective_default_folder"] = attrs["effective_default_folder"].SetComputed()
	attrs["effective_detection_categories"] = attrs["effective_detection_categories"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_schedule":               reflect.TypeOf(RuleSchedule_SdkV2{}),
		"detection_categories":           reflect.TypeOf(types.String{}),
		"effective_detection_categories": reflect.TypeOf(types.String{}),
		"updated_by":                     reflect.TypeOf(Principal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_folder":                  m.DefaultFolder,
			"default_schedule":                m.DefaultSchedule,
			"default_signal_collation_window": m.DefaultSignalCollationWindow,
			"detection_categories":            m.DetectionCategories,
			"effective_default_folder":        m.EffectiveDefaultFolder,
			"effective_detection_categories":  m.EffectiveDetectionCategories,
			"etag":                            m.Etag,
			"name":                            m.Name,
			"update_time":                     m.UpdateTime,
			"updated_by":                      m.UpdatedBy,
			"version":                         m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_folder": types.StringType,
			"default_schedule": basetypes.ListType{
				ElemType: RuleSchedule_SdkV2{}.Type(ctx),
			},
			"default_signal_collation_window": types.StringType,
			"detection_categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"effective_default_folder": types.StringType,
			"effective_detection_categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"etag":        types.StringType,
			"name":        types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by": basetypes.ListType{
				ElemType: Principal_SdkV2{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetDefaultSchedule returns the value of the DefaultSchedule field in RuleConfig_SdkV2 as
// a RuleSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleConfig_SdkV2) GetDefaultSchedule(ctx context.Context) (RuleSchedule_SdkV2, bool) {
	var e RuleSchedule_SdkV2
	if m.DefaultSchedule.IsNull() || m.DefaultSchedule.IsUnknown() {
		return e, false
	}
	var v []RuleSchedule_SdkV2
	d := m.DefaultSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultSchedule sets the value of the DefaultSchedule field in RuleConfig_SdkV2.
func (m *RuleConfig_SdkV2) SetDefaultSchedule(ctx context.Context, v RuleSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_schedule"]
	m.DefaultSchedule = types.ListValueMust(t, vs)
}

// GetDetectionCategories returns the value of the DetectionCategories field in RuleConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleConfig_SdkV2) GetDetectionCategories(ctx context.Context) ([]types.String, bool) {
	if m.DetectionCategories.IsNull() || m.DetectionCategories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DetectionCategories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetectionCategories sets the value of the DetectionCategories field in RuleConfig_SdkV2.
func (m *RuleConfig_SdkV2) SetDetectionCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["detection_categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DetectionCategories = types.ListValueMust(t, vs)
}

// GetEffectiveDetectionCategories returns the value of the EffectiveDetectionCategories field in RuleConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleConfig_SdkV2) GetEffectiveDetectionCategories(ctx context.Context) ([]types.String, bool) {
	if m.EffectiveDetectionCategories.IsNull() || m.EffectiveDetectionCategories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EffectiveDetectionCategories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveDetectionCategories sets the value of the EffectiveDetectionCategories field in RuleConfig_SdkV2.
func (m *RuleConfig_SdkV2) SetEffectiveDetectionCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_detection_categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveDetectionCategories = types.ListValueMust(t, vs)
}

// GetUpdatedBy returns the value of the UpdatedBy field in RuleConfig_SdkV2 as
// a Principal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleConfig_SdkV2) GetUpdatedBy(ctx context.Context) (Principal_SdkV2, bool) {
	var e Principal_SdkV2
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v []Principal_SdkV2
	d := m.UpdatedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdatedBy sets the value of the UpdatedBy field in RuleConfig_SdkV2.
func (m *RuleConfig_SdkV2) SetUpdatedBy(ctx context.Context, v Principal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updated_by"]
	m.UpdatedBy = types.ListValueMust(t, vs)
}

// Batch input for a detection rule.
type RuleInputBatch_SdkV2 struct {
	// Custom-notebook user function. The output is used as input data for the
	// rule.
	Custom types.List `tfsdk:"custom"`
	// SQL query evaluated as the rule input.
	Sql types.List `tfsdk:"sql"`
	// SQL warehouse used to run this batch rule.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *RuleInputBatch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleInputBatch_SdkV2) {
	if !from.Custom.IsNull() && !from.Custom.IsUnknown() {
		if toCustom, ok := to.GetCustom(ctx); ok {
			if fromCustom, ok := from.GetCustom(ctx); ok {
				// Recursively sync the fields of Custom
				toCustom.SyncFieldsDuringCreateOrUpdate(ctx, fromCustom)
				to.SetCustom(ctx, toCustom)
			}
		}
	}
	if !from.Sql.IsNull() && !from.Sql.IsUnknown() {
		if toSql, ok := to.GetSql(ctx); ok {
			if fromSql, ok := from.GetSql(ctx); ok {
				// Recursively sync the fields of Sql
				toSql.SyncFieldsDuringCreateOrUpdate(ctx, fromSql)
				to.SetSql(ctx, toSql)
			}
		}
	}
}

func (to *RuleInputBatch_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleInputBatch_SdkV2) {
	if !from.Custom.IsNull() && !from.Custom.IsUnknown() {
		if toCustom, ok := to.GetCustom(ctx); ok {
			if fromCustom, ok := from.GetCustom(ctx); ok {
				toCustom.SyncFieldsDuringRead(ctx, fromCustom)
				to.SetCustom(ctx, toCustom)
			}
		}
	}
	if !from.Sql.IsNull() && !from.Sql.IsUnknown() {
		if toSql, ok := to.GetSql(ctx); ok {
			if fromSql, ok := from.GetSql(ctx); ok {
				toSql.SyncFieldsDuringRead(ctx, fromSql)
				to.SetSql(ctx, toSql)
			}
		}
	}
}

func (m RuleInputBatch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom"] = attrs["custom"].SetOptional()
	attrs["custom"] = attrs["custom"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sql"] = attrs["sql"].SetOptional()
	attrs["sql"] = attrs["sql"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleInputBatch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleInputBatch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom": reflect.TypeOf(RuleInputCustomNotebook_SdkV2{}),
		"sql":    reflect.TypeOf(RuleInputBatchSql_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleInputBatch_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleInputBatch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom":       m.Custom,
			"sql":          m.Sql,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleInputBatch_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom": basetypes.ListType{
				ElemType: RuleInputCustomNotebook_SdkV2{}.Type(ctx),
			},
			"sql": basetypes.ListType{
				ElemType: RuleInputBatchSql_SdkV2{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetCustom returns the value of the Custom field in RuleInputBatch_SdkV2 as
// a RuleInputCustomNotebook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleInputBatch_SdkV2) GetCustom(ctx context.Context) (RuleInputCustomNotebook_SdkV2, bool) {
	var e RuleInputCustomNotebook_SdkV2
	if m.Custom.IsNull() || m.Custom.IsUnknown() {
		return e, false
	}
	var v []RuleInputCustomNotebook_SdkV2
	d := m.Custom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCustom sets the value of the Custom field in RuleInputBatch_SdkV2.
func (m *RuleInputBatch_SdkV2) SetCustom(ctx context.Context, v RuleInputCustomNotebook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom"]
	m.Custom = types.ListValueMust(t, vs)
}

// GetSql returns the value of the Sql field in RuleInputBatch_SdkV2 as
// a RuleInputBatchSql_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleInputBatch_SdkV2) GetSql(ctx context.Context) (RuleInputBatchSql_SdkV2, bool) {
	var e RuleInputBatchSql_SdkV2
	if m.Sql.IsNull() || m.Sql.IsUnknown() {
		return e, false
	}
	var v []RuleInputBatchSql_SdkV2
	d := m.Sql.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSql sets the value of the Sql field in RuleInputBatch_SdkV2.
func (m *RuleInputBatch_SdkV2) SetSql(ctx context.Context, v RuleInputBatchSql_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sql"]
	m.Sql = types.ListValueMust(t, vs)
}

// Batch SQL input for a detection rule.
type RuleInputBatchSql_SdkV2 struct {
	// Default catalog used to resolve unqualified identifiers in `query`. If
	// unset, the rule's `catalog_name` is used.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name"`
	// Default schema used to resolve unqualified identifiers in `query`. If
	// unset, the rule's `schema_name` is used.
	DefaultSchemaName types.String `tfsdk:"default_schema_name"`
	// SQL query evaluated as the rule input.
	Query types.String `tfsdk:"query"`
}

func (to *RuleInputBatchSql_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleInputBatchSql_SdkV2) {
}

func (to *RuleInputBatchSql_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleInputBatchSql_SdkV2) {
}

func (m RuleInputBatchSql_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_catalog_name"] = attrs["default_catalog_name"].SetOptional()
	attrs["default_schema_name"] = attrs["default_schema_name"].SetOptional()
	attrs["query"] = attrs["query"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleInputBatchSql.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleInputBatchSql_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleInputBatchSql_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleInputBatchSql_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": m.DefaultCatalogName,
			"default_schema_name":  m.DefaultSchemaName,
			"query":                m.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleInputBatchSql_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"default_schema_name":  types.StringType,
			"query":                types.StringType,
		},
	}
}

// Reference to a custom-user-function notebook for stream or batch input.
type RuleInputCustomNotebook_SdkV2 struct {
	// Fully-qualified notebook path, or a path relative to the configured base
	// notebook directory.
	Notebook types.String `tfsdk:"notebook"`
	// Optional arguments supplied when invoking the custom function.
	Options types.Map `tfsdk:"options"`
}

func (to *RuleInputCustomNotebook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleInputCustomNotebook_SdkV2) {
}

func (to *RuleInputCustomNotebook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleInputCustomNotebook_SdkV2) {
}

func (m RuleInputCustomNotebook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["notebook"] = attrs["notebook"].SetRequired()
	attrs["options"] = attrs["options"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleInputCustomNotebook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleInputCustomNotebook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleInputCustomNotebook_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleInputCustomNotebook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebook": m.Notebook,
			"options":  m.Options,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleInputCustomNotebook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notebook": types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetOptions returns the value of the Options field in RuleInputCustomNotebook_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleInputCustomNotebook_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in RuleInputCustomNotebook_SdkV2.
func (m *RuleInputCustomNotebook_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Options = types.MapValueMust(t, vs)
}

// MITRE ATT&CK mapping for a detection rule.
type RuleMitreMapping_SdkV2 struct {
	// Human-readable MITRE sub-technique name.
	SubTechnique types.String `tfsdk:"sub_technique"`
	// MITRE sub-technique ID, such as `T1078.001`.
	SubTechniqueId types.String `tfsdk:"sub_technique_id"`
	// Adversary objective during this attack, such as `Command and Control`.
	Tactic types.String `tfsdk:"tactic"`
	// Structured framework that organizes adversary behaviors, such as
	// `mitre-attack`.
	Taxonomy types.String `tfsdk:"taxonomy"`
	// Human-readable MITRE technique name.
	Technique types.String `tfsdk:"technique"`
	// MITRE technique ID, such as `T1078`.
	TechniqueId types.String `tfsdk:"technique_id"`
}

func (to *RuleMitreMapping_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleMitreMapping_SdkV2) {
}

func (to *RuleMitreMapping_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleMitreMapping_SdkV2) {
}

func (m RuleMitreMapping_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sub_technique"] = attrs["sub_technique"].SetOptional()
	attrs["sub_technique_id"] = attrs["sub_technique_id"].SetOptional()
	attrs["tactic"] = attrs["tactic"].SetOptional()
	attrs["taxonomy"] = attrs["taxonomy"].SetOptional()
	attrs["technique"] = attrs["technique"].SetOptional()
	attrs["technique_id"] = attrs["technique_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleMitreMapping.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleMitreMapping_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleMitreMapping_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleMitreMapping_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sub_technique":    m.SubTechnique,
			"sub_technique_id": m.SubTechniqueId,
			"tactic":           m.Tactic,
			"taxonomy":         m.Taxonomy,
			"technique":        m.Technique,
			"technique_id":     m.TechniqueId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleMitreMapping_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sub_technique":    types.StringType,
			"sub_technique_id": types.StringType,
			"tactic":           types.StringType,
			"taxonomy":         types.StringType,
			"technique":        types.StringType,
			"technique_id":     types.StringType,
		},
	}
}

// What to record in each signal that this rule generates.
type RuleOutput_SdkV2 struct {
	// Key/value context to record from the input row. Values may contain SQL
	// expressions in braces.
	Context types.Map `tfsdk:"context"`
	// If true, also record every column of the triggering input row. Custom
	// context wins on key collision.
	DefaultContext types.Bool `tfsdk:"default_context"`
	// Column in the rule output that holds the event timestamp.
	EventTimeColumn types.String `tfsdk:"event_time_column"`
	// Signal summary. May contain SQL expressions in braces that the server
	// substitutes from the input row.
	Summary types.String `tfsdk:"summary"`
}

func (to *RuleOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleOutput_SdkV2) {
}

func (to *RuleOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleOutput_SdkV2) {
}

func (m RuleOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["context"] = attrs["context"].SetOptional()
	attrs["default_context"] = attrs["default_context"].SetOptional()
	attrs["event_time_column"] = attrs["event_time_column"].SetOptional()
	attrs["summary"] = attrs["summary"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"context": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"context":           m.Context,
			"default_context":   m.DefaultContext,
			"event_time_column": m.EventTimeColumn,
			"summary":           m.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"context": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_context":   types.BoolType,
			"event_time_column": types.StringType,
			"summary":           types.StringType,
		},
	}
}

// GetContext returns the value of the Context field in RuleOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RuleOutput_SdkV2) GetContext(ctx context.Context) (map[string]types.String, bool) {
	if m.Context.IsNull() || m.Context.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Context.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetContext sets the value of the Context field in RuleOutput_SdkV2.
func (m *RuleOutput_SdkV2) SetContext(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["context"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Context = types.MapValueMust(t, vs)
}

// Defines when a detection rule executes.
//
// A batch rule sets exactly one of `at_least_every` or `cron_expression`. A
// streaming rule runs continuously and sets no timing; it accepts only
// `enabled`, and an omitted schedule means enabled. The server rejects any
// other combination. `enabled` is orthogonal to the timing.
type RuleSchedule_SdkV2 struct {
	// Defines the minimum cadence as a duration string. It may contain
	// components for days, hours, minutes, and seconds (e.g. `5d`, `5d3h`,
	// `3d2h5m12s`, `12m16s`).
	AtLeastEvery types.String `tfsdk:"at_least_every"`
	// Runs the rule continuously. Detection rules do not accept it: batch rules
	// reject it, and streaming rules run continuously without it.
	Continuous types.Bool `tfsdk:"continuous"`
	// Defines the schedule using a Quartz cron expression.
	CronExpression types.String `tfsdk:"cron_expression"`
	// Sets whether the job is scheduled to run. Defaults to true.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *RuleSchedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleSchedule_SdkV2) {
}

func (to *RuleSchedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleSchedule_SdkV2) {
}

func (m RuleSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["at_least_every"] = attrs["at_least_every"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["cron_expression"] = attrs["cron_expression"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"at_least_every":  m.AtLeastEvery,
			"continuous":      m.Continuous,
			"cron_expression": m.CronExpression,
			"enabled":         m.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"at_least_every":  types.StringType,
			"continuous":      types.BoolType,
			"cron_expression": types.StringType,
			"enabled":         types.BoolType,
		},
	}
}

// Time-range configuration for a batch detection rule.
type RuleTimeRange_SdkV2 struct {
	// Extends the start of the execution time range by a compact duration with
	// integer days, hours, minutes, and seconds in that order, such as `1h30m`
	// or `15m`. Zero duration disables extension. Unset or empty uses a fixed
	// `15m` at execution time. Only the start moves; the end is unchanged. Has
	// no effect on batch SQL that does not use `IN_TIMERANGE`. Values must fit
	// signed 64-bit microseconds.
	ExtendLookback types.String `tfsdk:"extend_lookback"`
}

func (to *RuleTimeRange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RuleTimeRange_SdkV2) {
}

func (to *RuleTimeRange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RuleTimeRange_SdkV2) {
}

func (m RuleTimeRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["extend_lookback"] = attrs["extend_lookback"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleTimeRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RuleTimeRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleTimeRange_SdkV2
// only implements ToObjectValue() and Type().
func (m RuleTimeRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"extend_lookback": m.ExtendLookback,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RuleTimeRange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"extend_lookback": types.StringType,
		},
	}
}

// Defines when a recurring job executes.
//
// Exactly one of the timing modes (`at_least_every`, `cron_expression`,
// `continuous`) must be set; combining them or setting none is rejected by the
// server. `compute_group` and `enabled` are orthogonal to the timing.
type Schedule_SdkV2 struct {
	// Duration string defining the minimum cadence. May contain components for
	// days, hours, minutes, seconds (e.g. `5d`, `5d3h`, `3d2h5m12s`, `12m16s`).
	AtLeastEvery types.String `tfsdk:"at_least_every"`
	// `dedicated`, `inherited`, `automatic`, or a custom group name. Reads
	// return the value as supplied. A rule whose schedule omits it runs on
	// `dedicated`.
	ComputeGroup types.String `tfsdk:"compute_group"`
	// Run continuously (sentinel — must be `true` when set in this branch).
	Continuous types.Bool `tfsdk:"continuous"`
	// Quartz cron expression.
	CronExpression types.String `tfsdk:"cron_expression"`
	// Whether the job is scheduled to run. Defaults to true.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *Schedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Schedule_SdkV2) {
}

func (to *Schedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Schedule_SdkV2) {
}

func (m Schedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["at_least_every"] = attrs["at_least_every"].SetOptional()
	attrs["compute_group"] = attrs["compute_group"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["cron_expression"] = attrs["cron_expression"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Schedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Schedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Schedule_SdkV2
// only implements ToObjectValue() and Type().
func (m Schedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"at_least_every":  m.AtLeastEvery,
			"compute_group":   m.ComputeGroup,
			"continuous":      m.Continuous,
			"cron_expression": m.CronExpression,
			"enabled":         m.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Schedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"at_least_every":  types.StringType,
			"compute_group":   types.StringType,
			"continuous":      types.BoolType,
			"cron_expression": types.StringType,
			"enabled":         types.BoolType,
		},
	}
}

// ServiceNow destination. Alerts are delivered to the instance's Table API,
// authorized by an OAuth 2.0 client-credentials token minted from the
// registered application's client ID and secret.
type ServiceNowExportConfig_SdkV2 struct {
	// OAuth client ID of the registered ServiceNow application. An identifier
	// rather than secret material — ServiceNow shows it on the application
	// record — so it is stored and read back verbatim.
	ClientId types.String `tfsdk:"client_id"`
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the OAuth client secret. Only the name is stored; the secret is
	// read when an alert is exported, so rotating it takes effect without
	// editing this config.
	ClientSecretUcSecret types.String `tfsdk:"client_secret_uc_secret"`
	// ServiceNow instance ID — the host label only, not the full URL. For an
	// instance reachable at `https://acme.service-now.com`, this is `acme`.
	InstanceId types.String `tfsdk:"instance_id"`
}

func (to *ServiceNowExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServiceNowExportConfig_SdkV2) {
}

func (to *ServiceNowExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServiceNowExportConfig_SdkV2) {
}

func (m ServiceNowExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client_id"] = attrs["client_id"].SetRequired()
	attrs["client_secret_uc_secret"] = attrs["client_secret_uc_secret"].SetRequired()
	attrs["instance_id"] = attrs["instance_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServiceNowExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServiceNowExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceNowExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m ServiceNowExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":               m.ClientId,
			"client_secret_uc_secret": m.ClientSecretUcSecret,
			"instance_id":             m.InstanceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServiceNowExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_id":               types.StringType,
			"client_secret_uc_secret": types.StringType,
			"instance_id":             types.StringType,
		},
	}
}

type SlackExportConfig_SdkV2 struct {
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Slack bot token. Only the name is stored; the token is read
	// when an alert is exported, so rotating the secret takes effect without
	// editing this config.
	BotTokenUcSecret types.String `tfsdk:"bot_token_uc_secret"`
	// Channel the alert is posted to: a channel ID (e.g. `C0123456789`) or a
	// channel name without the leading `#` (e.g. `soc-alerts`). The value is
	// stored as sent, so a leading `#`, surrounding whitespace, or a value
	// naming more than one channel is rejected. For private channels, prefer
	// the channel ID — Slack does not guarantee name resolution for those.
	//
	// The Slack app must be a member of the channel and hold the `chat:write`
	// scope.
	Channel types.String `tfsdk:"channel"`
}

func (to *SlackExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SlackExportConfig_SdkV2) {
}

func (to *SlackExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SlackExportConfig_SdkV2) {
}

func (m SlackExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bot_token_uc_secret"] = attrs["bot_token_uc_secret"].SetRequired()
	attrs["channel"] = attrs["channel"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SlackExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SlackExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SlackExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m SlackExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bot_token_uc_secret": m.BotTokenUcSecret,
			"channel":             m.Channel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SlackExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bot_token_uc_secret": types.StringType,
			"channel":             types.StringType,
		},
	}
}

// A fully-qualified Unity Catalog table reference by catalog / schema / table
// name.
type TableName_SdkV2 struct {
	// Name of the parent Unity Catalog catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Name of the parent Unity Catalog schema.
	SchemaName types.String `tfsdk:"schema_name"`
	// Unity Catalog table name.
	TableName types.String `tfsdk:"table_name"`
}

func (to *TableName_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableName_SdkV2) {
}

func (to *TableName_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TableName_SdkV2) {
}

func (m TableName_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableName.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TableName_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableName_SdkV2
// only implements ToObjectValue() and Type().
func (m TableName_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": m.CatalogName,
			"schema_name":  m.SchemaName,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TableName_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"schema_name":  types.StringType,
			"table_name":   types.StringType,
		},
	}
}

// Tines destination. Alerts are delivered to a Tines webhook action, authorized
// by a secret sent in the `Authorization` header.
type TinesExportConfig_SdkV2 struct {
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
	ExportUpdates types.Bool `tfsdk:"export_updates"`
	// Three-level Unity Catalog name (`catalog.schema.secret`) of the secret
	// holding the Tines webhook secret. Only the name is stored; the secret is
	// read when an alert is exported, so rotating it takes effect without
	// editing this config. Required: without it the destination would be an
	// unauthenticated endpoint.
	WebhookSecretUcSecret types.String `tfsdk:"webhook_secret_uc_secret"`
	// Tines webhook action URL, typically of the form
	// `https://<tenant>.tines.com/webhook/<path>`. Any HTTPS host is accepted,
	// so that a Tines tenant on a custom domain works.
	//
	// The secret is sent as a header, so it must not be embedded in the URL: a
	// trailing secret path segment (`/webhook/<path>/<secret>`) and a `secret`
	// query parameter are both rejected. API-path-prefix URLs
	// (`/api/public/...`) are not supported.
	WebhookUri types.String `tfsdk:"webhook_uri"`
}

func (to *TinesExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TinesExportConfig_SdkV2) {
}

func (to *TinesExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TinesExportConfig_SdkV2) {
}

func (m TinesExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["export_updates"] = attrs["export_updates"].SetOptional()
	attrs["webhook_secret_uc_secret"] = attrs["webhook_secret_uc_secret"].SetRequired()
	attrs["webhook_uri"] = attrs["webhook_uri"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TinesExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TinesExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TinesExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m TinesExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"export_updates":           m.ExportUpdates,
			"webhook_secret_uc_secret": m.WebhookSecretUcSecret,
			"webhook_uri":              m.WebhookUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TinesExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"export_updates":           types.BoolType,
			"webhook_secret_uc_secret": types.StringType,
			"webhook_uri":              types.StringType,
		},
	}
}

// Ingests data through a managed connector that uses a Unity Catalog
// connection. Mutually exclusive with `autoloader` and `existing_table`.
type UnityCatalogConnection_SdkV2 struct {
	// Free-form key-value configuration for the managed ingestion pipeline.
	Configuration types.Map `tfsdk:"configuration"`
	// Name of the Unity Catalog connection to ingest through.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Connector type of the connection, such as `SLACK_AUDIT_LOGS`. Resolved by
	// the server from `connection_name` through Unity Catalog; not accepted as
	// input.
	Connector types.String `tfsdk:"connector"`
	// The start time ingestion uses: the supplied `start_time`, or the creation
	// time when it was omitted. Always populated.
	EffectiveStartTime timetypes.RFC3339 `tfsdk:"effective_start_time"`
	// Earliest source event time to ingest. Leave unset to start at creation
	// time, with no history backfilled. Rejected if in the future, beyond the
	// connector's history window, or for a connector without start-time
	// support. Takes effect only at first start and cannot be changed. Reads
	// return the value as supplied; the start time ingestion uses is reported
	// in `effective_start_time`.
	StartTime timetypes.RFC3339 `tfsdk:"start_time"`
}

func (to *UnityCatalogConnection_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UnityCatalogConnection_SdkV2) {
}

func (to *UnityCatalogConnection_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UnityCatalogConnection_SdkV2) {
}

func (m UnityCatalogConnection_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetRequired()
	attrs["connector"] = attrs["connector"].SetComputed()
	attrs["effective_start_time"] = attrs["effective_start_time"].SetComputed()
	attrs["start_time"] = attrs["start_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnityCatalogConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UnityCatalogConnection_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"configuration": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnityCatalogConnection_SdkV2
// only implements ToObjectValue() and Type().
func (m UnityCatalogConnection_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"configuration":        m.Configuration,
			"connection_name":      m.ConnectionName,
			"connector":            m.Connector,
			"effective_start_time": m.EffectiveStartTime,
			"start_time":           m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UnityCatalogConnection_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"connection_name":      types.StringType,
			"connector":            types.StringType,
			"effective_start_time": timetypes.RFC3339{}.Type(ctx),
			"start_time":           timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetConfiguration returns the value of the Configuration field in UnityCatalogConnection_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UnityCatalogConnection_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if m.Configuration.IsNull() || m.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in UnityCatalogConnection_SdkV2.
func (m *UnityCatalogConnection_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

type UpdateAlertConfigRequest_SdkV2 struct {
	AlertConfig types.List `tfsdk:"alert_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
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
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateAlertConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertConfigRequest_SdkV2) {
	if !from.AlertConfig.IsNull() && !from.AlertConfig.IsUnknown() {
		if toAlertConfig, ok := to.GetAlertConfig(ctx); ok {
			if fromAlertConfig, ok := from.GetAlertConfig(ctx); ok {
				// Recursively sync the fields of AlertConfig
				toAlertConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAlertConfig)
				to.SetAlertConfig(ctx, toAlertConfig)
			}
		}
	}
}

func (to *UpdateAlertConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertConfigRequest_SdkV2) {
	if !from.AlertConfig.IsNull() && !from.AlertConfig.IsUnknown() {
		if toAlertConfig, ok := to.GetAlertConfig(ctx); ok {
			if fromAlertConfig, ok := from.GetAlertConfig(ctx); ok {
				toAlertConfig.SyncFieldsDuringRead(ctx, fromAlertConfig)
				to.SetAlertConfig(ctx, toAlertConfig)
			}
		}
	}
}

func (m UpdateAlertConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_config"] = attrs["alert_config"].SetRequired()
	attrs["alert_config"] = attrs["alert_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAlertConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_config": reflect.TypeOf(AlertConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAlertConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_config": m.AlertConfig,
			"etag":         m.Etag,
			"name":         m.Name,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAlertConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_config": basetypes.ListType{
				ElemType: AlertConfig_SdkV2{}.Type(ctx),
			},
			"etag":        types.StringType,
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetAlertConfig returns the value of the AlertConfig field in UpdateAlertConfigRequest_SdkV2 as
// a AlertConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAlertConfigRequest_SdkV2) GetAlertConfig(ctx context.Context) (AlertConfig_SdkV2, bool) {
	var e AlertConfig_SdkV2
	if m.AlertConfig.IsNull() || m.AlertConfig.IsUnknown() {
		return e, false
	}
	var v []AlertConfig_SdkV2
	d := m.AlertConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlertConfig sets the value of the AlertConfig field in UpdateAlertConfigRequest_SdkV2.
func (m *UpdateAlertConfigRequest_SdkV2) SetAlertConfig(ctx context.Context, v AlertConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_config"]
	m.AlertConfig = types.ListValueMust(t, vs)
}

type UpdateCommonConfigRequest_SdkV2 struct {
	CommonConfig types.List `tfsdk:"common_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
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
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateCommonConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCommonConfigRequest_SdkV2) {
	if !from.CommonConfig.IsNull() && !from.CommonConfig.IsUnknown() {
		if toCommonConfig, ok := to.GetCommonConfig(ctx); ok {
			if fromCommonConfig, ok := from.GetCommonConfig(ctx); ok {
				// Recursively sync the fields of CommonConfig
				toCommonConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromCommonConfig)
				to.SetCommonConfig(ctx, toCommonConfig)
			}
		}
	}
}

func (to *UpdateCommonConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCommonConfigRequest_SdkV2) {
	if !from.CommonConfig.IsNull() && !from.CommonConfig.IsUnknown() {
		if toCommonConfig, ok := to.GetCommonConfig(ctx); ok {
			if fromCommonConfig, ok := from.GetCommonConfig(ctx); ok {
				toCommonConfig.SyncFieldsDuringRead(ctx, fromCommonConfig)
				to.SetCommonConfig(ctx, toCommonConfig)
			}
		}
	}
}

func (m UpdateCommonConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["common_config"] = attrs["common_config"].SetRequired()
	attrs["common_config"] = attrs["common_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCommonConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCommonConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"common_config": reflect.TypeOf(CommonConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCommonConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCommonConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"common_config": m.CommonConfig,
			"etag":          m.Etag,
			"name":          m.Name,
			"update_mask":   m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCommonConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"common_config": basetypes.ListType{
				ElemType: CommonConfig_SdkV2{}.Type(ctx),
			},
			"etag":        types.StringType,
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetCommonConfig returns the value of the CommonConfig field in UpdateCommonConfigRequest_SdkV2 as
// a CommonConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCommonConfigRequest_SdkV2) GetCommonConfig(ctx context.Context) (CommonConfig_SdkV2, bool) {
	var e CommonConfig_SdkV2
	if m.CommonConfig.IsNull() || m.CommonConfig.IsUnknown() {
		return e, false
	}
	var v []CommonConfig_SdkV2
	d := m.CommonConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCommonConfig sets the value of the CommonConfig field in UpdateCommonConfigRequest_SdkV2.
func (m *UpdateCommonConfigRequest_SdkV2) SetCommonConfig(ctx context.Context, v CommonConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["common_config"]
	m.CommonConfig = types.ListValueMust(t, vs)
}

type UpdateDatasourceConfigRequest_SdkV2 struct {
	DatasourceConfig types.List `tfsdk:"datasource_config"`
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
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
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatasourceConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatasourceConfigRequest_SdkV2) {
	if !from.DatasourceConfig.IsNull() && !from.DatasourceConfig.IsUnknown() {
		if toDatasourceConfig, ok := to.GetDatasourceConfig(ctx); ok {
			if fromDatasourceConfig, ok := from.GetDatasourceConfig(ctx); ok {
				// Recursively sync the fields of DatasourceConfig
				toDatasourceConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromDatasourceConfig)
				to.SetDatasourceConfig(ctx, toDatasourceConfig)
			}
		}
	}
}

func (to *UpdateDatasourceConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDatasourceConfigRequest_SdkV2) {
	if !from.DatasourceConfig.IsNull() && !from.DatasourceConfig.IsUnknown() {
		if toDatasourceConfig, ok := to.GetDatasourceConfig(ctx); ok {
			if fromDatasourceConfig, ok := from.GetDatasourceConfig(ctx); ok {
				toDatasourceConfig.SyncFieldsDuringRead(ctx, fromDatasourceConfig)
				to.SetDatasourceConfig(ctx, toDatasourceConfig)
			}
		}
	}
}

func (m UpdateDatasourceConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasource_config"] = attrs["datasource_config"].SetRequired()
	attrs["datasource_config"] = attrs["datasource_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatasourceConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatasourceConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasource_config": reflect.TypeOf(DatasourceConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatasourceConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDatasourceConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasource_config": m.DatasourceConfig,
			"etag":              m.Etag,
			"name":              m.Name,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatasourceConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasource_config": basetypes.ListType{
				ElemType: DatasourceConfig_SdkV2{}.Type(ctx),
			},
			"etag":        types.StringType,
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDatasourceConfig returns the value of the DatasourceConfig field in UpdateDatasourceConfigRequest_SdkV2 as
// a DatasourceConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatasourceConfigRequest_SdkV2) GetDatasourceConfig(ctx context.Context) (DatasourceConfig_SdkV2, bool) {
	var e DatasourceConfig_SdkV2
	if m.DatasourceConfig.IsNull() || m.DatasourceConfig.IsUnknown() {
		return e, false
	}
	var v []DatasourceConfig_SdkV2
	d := m.DatasourceConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatasourceConfig sets the value of the DatasourceConfig field in UpdateDatasourceConfigRequest_SdkV2.
func (m *UpdateDatasourceConfigRequest_SdkV2) SetDatasourceConfig(ctx context.Context, v DatasourceConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasource_config"]
	m.DatasourceConfig = types.ListValueMust(t, vs)
}

type UpdateDatasourceRequest_SdkV2 struct {
	Datasource types.List `tfsdk:"datasource"`
	// Optimistic-concurrency precondition: if set to a non-empty value that
	// does not match the resource's current `etag`, the update is rejected as a
	// conflict (HTTP 409). Omit this field or send an empty string to update
	// unconditionally.
	Etag types.String `tfsdk:"-"`
	// Resource name of the datasource, in the format
	// `datasources/{datasource_id}`. Unlike the datasource's Unity Catalog
	// location (`catalog_name`, `schema_name`, and `datasource_name`), the name
	// does not change when the datasource or its catalog or schema is renamed
	// in Unity Catalog.
	Name types.String `tfsdk:"-"`
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
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatasourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatasourceRequest_SdkV2) {
	if !from.Datasource.IsNull() && !from.Datasource.IsUnknown() {
		if toDatasource, ok := to.GetDatasource(ctx); ok {
			if fromDatasource, ok := from.GetDatasource(ctx); ok {
				// Recursively sync the fields of Datasource
				toDatasource.SyncFieldsDuringCreateOrUpdate(ctx, fromDatasource)
				to.SetDatasource(ctx, toDatasource)
			}
		}
	}
}

func (to *UpdateDatasourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDatasourceRequest_SdkV2) {
	if !from.Datasource.IsNull() && !from.Datasource.IsUnknown() {
		if toDatasource, ok := to.GetDatasource(ctx); ok {
			if fromDatasource, ok := from.GetDatasource(ctx); ok {
				toDatasource.SyncFieldsDuringRead(ctx, fromDatasource)
				to.SetDatasource(ctx, toDatasource)
			}
		}
	}
}

func (m UpdateDatasourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasource"] = attrs["datasource"].SetRequired()
	attrs["datasource"] = attrs["datasource"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatasourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatasourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasource": reflect.TypeOf(Datasource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatasourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDatasourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasource":  m.Datasource,
			"etag":        m.Etag,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatasourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasource": basetypes.ListType{
				ElemType: Datasource_SdkV2{}.Type(ctx),
			},
			"etag":        types.StringType,
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDatasource returns the value of the Datasource field in UpdateDatasourceRequest_SdkV2 as
// a Datasource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatasourceRequest_SdkV2) GetDatasource(ctx context.Context) (Datasource_SdkV2, bool) {
	var e Datasource_SdkV2
	if m.Datasource.IsNull() || m.Datasource.IsUnknown() {
		return e, false
	}
	var v []Datasource_SdkV2
	d := m.Datasource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatasource sets the value of the Datasource field in UpdateDatasourceRequest_SdkV2.
func (m *UpdateDatasourceRequest_SdkV2) SetDatasource(ctx context.Context, v Datasource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasource"]
	m.Datasource = types.ListValueMust(t, vs)
}

type UpdateRuleConfigRequest_SdkV2 struct {
	// Optimistic-concurrency precondition: if set and it does not match the
	// resource's current `etag`, the update is rejected as a conflict (HTTP
	// 409).
	Etag types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`

	RuleConfig types.List `tfsdk:"rule_config"`
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
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateRuleConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRuleConfigRequest_SdkV2) {
	if !from.RuleConfig.IsNull() && !from.RuleConfig.IsUnknown() {
		if toRuleConfig, ok := to.GetRuleConfig(ctx); ok {
			if fromRuleConfig, ok := from.GetRuleConfig(ctx); ok {
				// Recursively sync the fields of RuleConfig
				toRuleConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromRuleConfig)
				to.SetRuleConfig(ctx, toRuleConfig)
			}
		}
	}
}

func (to *UpdateRuleConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRuleConfigRequest_SdkV2) {
	if !from.RuleConfig.IsNull() && !from.RuleConfig.IsUnknown() {
		if toRuleConfig, ok := to.GetRuleConfig(ctx); ok {
			if fromRuleConfig, ok := from.GetRuleConfig(ctx); ok {
				toRuleConfig.SyncFieldsDuringRead(ctx, fromRuleConfig)
				to.SetRuleConfig(ctx, toRuleConfig)
			}
		}
	}
}

func (m UpdateRuleConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["rule_config"] = attrs["rule_config"].SetRequired()
	attrs["rule_config"] = attrs["rule_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRuleConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_config": reflect.TypeOf(RuleConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRuleConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRuleConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        m.Etag,
			"name":        m.Name,
			"rule_config": m.RuleConfig,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRuleConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
			"rule_config": basetypes.ListType{
				ElemType: RuleConfig_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetRuleConfig returns the value of the RuleConfig field in UpdateRuleConfigRequest_SdkV2 as
// a RuleConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRuleConfigRequest_SdkV2) GetRuleConfig(ctx context.Context) (RuleConfig_SdkV2, bool) {
	var e RuleConfig_SdkV2
	if m.RuleConfig.IsNull() || m.RuleConfig.IsUnknown() {
		return e, false
	}
	var v []RuleConfig_SdkV2
	d := m.RuleConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRuleConfig sets the value of the RuleConfig field in UpdateRuleConfigRequest_SdkV2.
func (m *UpdateRuleConfigRequest_SdkV2) SetRuleConfig(ctx context.Context, v RuleConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rule_config"]
	m.RuleConfig = types.ListValueMust(t, vs)
}

type UpdateRuleRequest_SdkV2 struct {
	// Optimistic-concurrency token from the latest read. A non-empty token that
	// does not match the current `etag` rejects the update as a conflict (HTTP
	// 409). Omit this field or provide an empty string to update
	// unconditionally.
	Etag types.String `tfsdk:"-"`
	// Resource name, in the form `rules/{rule_id}`. Assigned by the server on
	// create; required on update.
	Name types.String `tfsdk:"-"`
	// The rule with its new field values. `rule.name` identifies the rule to
	// update.
	Rule types.List `tfsdk:"rule"`
	// The set of fields to update. Only fields named in the mask are modified.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRuleRequest_SdkV2) {
	if !from.Rule.IsNull() && !from.Rule.IsUnknown() {
		if toRule, ok := to.GetRule(ctx); ok {
			if fromRule, ok := from.GetRule(ctx); ok {
				// Recursively sync the fields of Rule
				toRule.SyncFieldsDuringCreateOrUpdate(ctx, fromRule)
				to.SetRule(ctx, toRule)
			}
		}
	}
}

func (to *UpdateRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRuleRequest_SdkV2) {
	if !from.Rule.IsNull() && !from.Rule.IsUnknown() {
		if toRule, ok := to.GetRule(ctx); ok {
			if fromRule, ok := from.GetRule(ctx); ok {
				toRule.SyncFieldsDuringRead(ctx, fromRule)
				to.SetRule(ctx, toRule)
			}
		}
	}
}

func (m UpdateRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["rule"] = attrs["rule"].SetRequired()
	attrs["rule"] = attrs["rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule": reflect.TypeOf(Rule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        m.Etag,
			"name":        m.Name,
			"rule":        m.Rule,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
			"rule": basetypes.ListType{
				ElemType: Rule_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetRule returns the value of the Rule field in UpdateRuleRequest_SdkV2 as
// a Rule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRuleRequest_SdkV2) GetRule(ctx context.Context) (Rule_SdkV2, bool) {
	var e Rule_SdkV2
	if m.Rule.IsNull() || m.Rule.IsUnknown() {
		return e, false
	}
	var v []Rule_SdkV2
	d := m.Rule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRule sets the value of the Rule field in UpdateRuleRequest_SdkV2.
func (m *UpdateRuleRequest_SdkV2) SetRule(ctx context.Context, v Rule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rule"]
	m.Rule = types.ListValueMust(t, vs)
}

type WebhookExportConfig_SdkV2 struct {
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
	ExportUpdates types.Bool `tfsdk:"export_updates"`
	// Webhook destination the server POSTs payloads to, in plaintext. Userinfo
	// (`user:password@`) is sent as HTTP Basic auth.
	Url types.String `tfsdk:"url"`
}

func (to *WebhookExportConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WebhookExportConfig_SdkV2) {
}

func (to *WebhookExportConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WebhookExportConfig_SdkV2) {
}

func (m WebhookExportConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["export_updates"] = attrs["export_updates"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WebhookExportConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WebhookExportConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WebhookExportConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m WebhookExportConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"export_updates": m.ExportUpdates,
			"url":            m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WebhookExportConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"export_updates": types.BoolType,
			"url":            types.StringType,
		},
	}
}
