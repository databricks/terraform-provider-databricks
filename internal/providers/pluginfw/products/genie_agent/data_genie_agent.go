// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genie_agent

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/genieagents"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/genieagents_tf"
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

const dataSourceName = "genie_agent"

var _ datasource.DataSourceWithConfigure = &GenieAgentDataSource{}

func DataSourceGenieAgent() datasource.DataSource {
	return &GenieAgentDataSource{}
}

type GenieAgentDataSource struct {
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

// GenieAgentData extends the main model with additional fields.
type GenieAgentData struct {
	// Time when the Genie Agent was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The definition of the Genie Agent: its sample questions, data sources,
	// instructions, and benchmarks. Must set at least one field. Returned on
	// get and on list with `GENIE_AGENT_VIEW_FULL`. Create and update return a
	// submitted definition unchanged; one the service would store differently,
	// or with a column config the workspace can't keep (unknown column, or a
	// display name while those are disabled), is rejected with
	// INVALID_PARAMETER_VALUE.
	Definition types.Object `tfsdk:"definition"`
	// Description of the Genie Agent.
	//
	// The maximum length is 4,096 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	Description types.String `tfsdk:"description"`
	// Display name of the Genie Agent. Must be unique within its folder
	// (effective_parent_path): a create, rename, or move that would collide
	// with an existing Genie Agent's name in the target folder is rejected with
	// ALREADY_EXISTS.
	//
	// The maximum length is 63 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace folder that contains the Genie Agent.
	EffectiveParentPath types.String `tfsdk:"effective_parent_path"`
	// Weak etag of `definition` only; metadata changes do not affect it. Send
	// it on update or delete to detect concurrent definition changes.
	Etag types.String `tfsdk:"etag"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name types.String `tfsdk:"name"`
	// The workspace folder to create the Genie Agent in. If omitted, the
	// caller's home folder is used. Returned exactly as submitted; see
	// `effective_parent_path` for the current folder. On update, setting it
	// moves the Genie Agent and clearing it does not. At most 4,096 characters.
	ParentPath types.String `tfsdk:"parent_path"`
	// Time when the Genie Agent was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Warehouse associated with the Genie Agent.
	WarehouseId        types.String `tfsdk:"warehouse_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// GenieAgentData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m GenieAgentData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"definition":      reflect.TypeOf(genieagents_tf.GenieAgentDefinition{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentData
// only implements ToObjectValue() and Type().
func (m GenieAgentData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"definition":            m.Definition,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"effective_parent_path": m.EffectiveParentPath,
			"etag":                  m.Etag,
			"name":                  m.Name,
			"parent_path":           m.ParentPath,
			"update_time":           m.UpdateTime,
			"warehouse_id":          m.WarehouseId,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m GenieAgentData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":           timetypes.RFC3339{}.Type(ctx),
			"definition":            genieagents_tf.GenieAgentDefinition{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"effective_parent_path": types.StringType,
			"etag":                  types.StringType,
			"name":                  types.StringType,
			"parent_path":           types.StringType,
			"update_time":           timetypes.RFC3339{}.Type(ctx),
			"warehouse_id":          types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m GenieAgentData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["definition"] = attrs["definition"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["effective_parent_path"] = attrs["effective_parent_path"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent_path"] = attrs["parent_path"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *GenieAgentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *GenieAgentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GenieAgentData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks GenieAgent",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *GenieAgentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *GenieAgentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config GenieAgentData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest genieagents.GetGenieAgentRequest
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

	response, err := client.GenieAgents.GetGenieAgent(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get genie_agent", err.Error())
		return
	}

	var newState GenieAgentData
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
