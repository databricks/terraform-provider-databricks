// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_alert_config

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "lakewatch_alert_config"

var _ datasource.DataSourceWithConfigure = &AlertConfigDataSource{}

func DataSourceAlertConfig() datasource.DataSource {
	return &AlertConfigDataSource{}
}

type AlertConfigDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// AlertConfigData extends the main model with additional fields.
type AlertConfigData struct {
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
	SecurityAlertExportConfig types.Object `tfsdk:"security_alert_export_config"`
	// Export destination for system alerts. When unset, export is off and reads
	// return `{"enabled": false}`.
	SystemAlertExportConfig types.Object `tfsdk:"system_alert_export_config"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.Object `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version            types.Int64  `tfsdk:"version"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AlertConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AlertConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"security_alert_export_config": reflect.TypeOf(lakewatch_tf.ExportConfig{}),
		"system_alert_export_config":   reflect.TypeOf(lakewatch_tf.ExportConfig{}),
		"updated_by":                   reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config":              reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfigData
// only implements ToObjectValue() and Type().
func (m AlertConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AlertConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_security_alert_collation_window":           types.StringType,
			"default_system_alert_collation_window":             types.StringType,
			"effective_default_security_alert_collation_window": types.StringType,
			"effective_default_system_alert_collation_window":   types.StringType,
			"etag":                         types.StringType,
			"name":                         types.StringType,
			"security_alert_export_config": lakewatch_tf.ExportConfig{}.Type(ctx),
			"system_alert_export_config":   lakewatch_tf.ExportConfig{}.Type(ctx),
			"update_time":                  timetypes.RFC3339{}.Type(ctx),
			"updated_by":                   lakewatch_tf.Principal{}.Type(ctx),
			"version":                      types.Int64Type,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m AlertConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_security_alert_collation_window"] = attrs["default_security_alert_collation_window"].SetComputed()
	attrs["default_system_alert_collation_window"] = attrs["default_system_alert_collation_window"].SetComputed()
	attrs["effective_default_security_alert_collation_window"] = attrs["effective_default_security_alert_collation_window"].SetComputed()
	attrs["effective_default_system_alert_collation_window"] = attrs["effective_default_system_alert_collation_window"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["security_alert_export_config"] = attrs["security_alert_export_config"].SetComputed()
	attrs["system_alert_export_config"] = attrs["system_alert_export_config"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *AlertConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AlertConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AlertConfigData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AlertConfig",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AlertConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config AlertConfigData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetAlertConfigRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Lakewatch.GetAlertConfig(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get lakewatch_alert_config", err.Error())
		return
	}

	var newState AlertConfigData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
