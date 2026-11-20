// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_parser

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

const dataSourceName = "lakewatch_parser"

var _ datasource.DataSourceWithConfigure = &ParserDataSource{}

func DataSourceParser() datasource.DataSource {
	return &ParserDataSource{}
}

type ParserDataSource struct {
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

// ParserData extends the main model with additional fields.
type ParserData struct {
	// Name of the parent UC catalog, set on create. Reads return the value as
	// supplied; the current UC catalog name is reported in
	// `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// When the parser was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The principal that created this parser. Each version records its own
	// `created_by`.
	CreatedBy types.Object `tfsdk:"created_by"`
	// Description of version 1. Accepted only by `CreateParser`; read a
	// version's description with `GetParserVersion`.
	Description types.String `tfsdk:"description"`
	// The current name of the parser's UC catalog. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// The current name of the parser's UC schema. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Newest version of this parser. Read its spec by passing this to
	// `GetParserVersion`, or by using the `latest` alias. Absent when every
	// version has been deleted.
	LatestVersion types.Int64 `tfsdk:"latest_version"`
	// Resource name. Format: `parsers/{parser_id}`, where the leaf is the
	// immutable server-assigned UUID. Create-time placement names are stored
	// separately and do not affect the resource name.
	Name types.String `tfsdk:"name"`
	// The principal that owns this parser. Server-set from the authenticated
	// creator, never supplied by the client. Beta reads by UUID perform no
	// authorization. Owner-gated write RPCs are outside this API's current
	// implementation scope.
	Owner types.Object `tfsdk:"owner"`
	// The parser's server-assigned UUID. It is not currently a UC securable ID,
	// is immutable for the life of the resource, is the leaf of `name`, and is
	// returned by CreateParser.
	ParserId types.String `tfsdk:"parser_id"`
	// Client-supplied parser name, stored as supplied and fixed at create time.
	// Must be a valid UC identifier. A custom parser is not currently a UC
	// securable and this field has no rename path.
	ParserName types.String `tfsdk:"parser_name"`
	// Name of the parent UC schema, set on create. Reads return the value as
	// supplied; the current UC schema name is reported in
	// `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Spec of version 1, as a YAML document. Accepted only by `CreateParser`;
	// read a version's spec with `GetParserVersion`.
	Spec               types.String `tfsdk:"spec"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ParserData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ParserData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"created_by":      reflect.TypeOf(lakewatch_tf.Principal{}),
		"owner":           reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ParserData
// only implements ToObjectValue() and Type().
func (m ParserData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":           m.CatalogName,
			"create_time":            m.CreateTime,
			"created_by":             m.CreatedBy,
			"description":            m.Description,
			"effective_catalog_name": m.EffectiveCatalogName,
			"effective_schema_name":  m.EffectiveSchemaName,
			"latest_version":         m.LatestVersion,
			"name":                   m.Name,
			"owner":                  m.Owner,
			"parser_id":              m.ParserId,
			"parser_name":            m.ParserName,
			"schema_name":            m.SchemaName,
			"spec":                   m.Spec,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ParserData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":           types.StringType,
			"create_time":            timetypes.RFC3339{}.Type(ctx),
			"created_by":             lakewatch_tf.Principal{}.Type(ctx),
			"description":            types.StringType,
			"effective_catalog_name": types.StringType,
			"effective_schema_name":  types.StringType,
			"latest_version":         types.Int64Type,
			"name":                   types.StringType,
			"owner":                  lakewatch_tf.Principal{}.Type(ctx),
			"parser_id":              types.StringType,
			"parser_name":            types.StringType,
			"schema_name":            types.StringType,
			"spec":                   types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m ParserData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["latest_version"] = attrs["latest_version"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["parser_id"] = attrs["parser_id"].SetComputed()
	attrs["parser_name"] = attrs["parser_name"].SetComputed()
	attrs["schema_name"] = attrs["schema_name"].SetComputed()
	attrs["spec"] = attrs["spec"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *ParserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ParserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ParserData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Parser",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ParserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ParserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ParserData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetParserRequest
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

	response, err := client.Lakewatch.GetParser(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get lakewatch_parser", err.Error())
		return
	}

	var newState ParserData
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
