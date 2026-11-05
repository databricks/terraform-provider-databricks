// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_datasource_config

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

const dataSourceName = "lakewatch_datasource_config"

var _ datasource.DataSourceWithConfigure = &DatasourceConfigDataSource{}

func DataSourceDatasourceConfig() datasource.DataSource {
	return &DatasourceConfigDataSource{}
}

type DatasourceConfigDataSource struct {
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

// DatasourceConfigData extends the main model with additional fields.
type DatasourceConfigData struct {
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

	UpdatedBy types.Object `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version            types.Int64  `tfsdk:"version"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DatasourceConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DatasourceConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updated_by":      reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasourceConfigData
// only implements ToObjectValue() and Type().
func (m DatasourceConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DatasourceConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_compute_mode":     types.StringType,
			"default_folder":           types.StringType,
			"effective_default_folder": types.StringType,
			"etag":                     types.StringType,
			"name":                     types.StringType,
			"update_time":              timetypes.RFC3339{}.Type(ctx),
			"updated_by":               lakewatch_tf.Principal{}.Type(ctx),
			"version":                  types.Int64Type,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m DatasourceConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_compute_mode"] = attrs["default_compute_mode"].SetComputed()
	attrs["default_folder"] = attrs["default_folder"].SetComputed()
	attrs["effective_default_folder"] = attrs["effective_default_folder"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *DatasourceConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *DatasourceConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatasourceConfigData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DatasourceConfig",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatasourceConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DatasourceConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config DatasourceConfigData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetDatasourceConfigRequest
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

	response, err := client.Lakewatch.GetDatasourceConfig(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get lakewatch_datasource_config", err.Error())
		return
	}

	var newState DatasourceConfigData
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
