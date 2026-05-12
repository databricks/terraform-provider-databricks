// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledge_assistant_knowledge_source

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/knowledgeassistants_tf"
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

const dataSourceName = "knowledge_assistant_knowledge_source"

var _ datasource.DataSourceWithConfigure = &KnowledgeSourceDataSource{}

func DataSourceKnowledgeSource() datasource.DataSource {
	return &KnowledgeSourceDataSource{}
}

type KnowledgeSourceDataSource struct {
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
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
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

// KnowledgeSourceData extends the main model with additional fields.
type KnowledgeSourceData struct {
	// Timestamp when this knowledge source was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Description of the knowledge source. Required when creating a Knowledge
	// Source. When updating a Knowledge Source, optional unless included in
	// update_mask.
	Description types.String `tfsdk:"description"`
	// Human-readable display name of the knowledge source. Required when
	// creating a Knowledge Source. When updating a Knowledge Source, optional
	// unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`

	FileTable types.Object `tfsdk:"file_table"`

	Files types.Object `tfsdk:"files"`

	Id types.String `tfsdk:"id"`

	Index types.Object `tfsdk:"index"`
	// Timestamp representing the cutoff before which content in this knowledge
	// source is being ingested.
	KnowledgeCutoffTime timetypes.RFC3339 `tfsdk:"knowledge_cutoff_time"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"name"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType types.String `tfsdk:"source_type"`

	State              types.String `tfsdk:"state"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// KnowledgeSourceData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m KnowledgeSourceData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_table":      reflect.TypeOf(knowledgeassistants_tf.FileTableSpec{}),
		"files":           reflect.TypeOf(knowledgeassistants_tf.FilesSpec{}),
		"index":           reflect.TypeOf(knowledgeassistants_tf.IndexSpec{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeSourceData
// only implements ToObjectValue() and Type().
func (m KnowledgeSourceData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"file_table":            m.FileTable,
			"files":                 m.Files,
			"id":                    m.Id,
			"index":                 m.Index,
			"knowledge_cutoff_time": m.KnowledgeCutoffTime,
			"name":                  m.Name,
			"source_type":           m.SourceType,
			"state":                 m.State,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m KnowledgeSourceData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":           timetypes.RFC3339{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"file_table":            knowledgeassistants_tf.FileTableSpec{}.Type(ctx),
			"files":                 knowledgeassistants_tf.FilesSpec{}.Type(ctx),
			"id":                    types.StringType,
			"index":                 knowledgeassistants_tf.IndexSpec{}.Type(ctx),
			"knowledge_cutoff_time": timetypes.RFC3339{}.Type(ctx),
			"name":                  types.StringType,
			"source_type":           types.StringType,
			"state":                 types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m KnowledgeSourceData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["file_table"] = attrs["file_table"].SetComputed()
	attrs["files"] = attrs["files"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index"] = attrs["index"].SetComputed()
	attrs["knowledge_cutoff_time"] = attrs["knowledge_cutoff_time"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["source_type"] = attrs["source_type"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *KnowledgeSourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *KnowledgeSourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, KnowledgeSourceData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks KnowledgeSource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KnowledgeSourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *KnowledgeSourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config KnowledgeSourceData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest knowledgeassistants.GetKnowledgeSourceRequest
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

	response, err := client.KnowledgeAssistants.GetKnowledgeSource(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get knowledge_assistant_knowledge_source", err.Error())
		return
	}

	var newState KnowledgeSourceData
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
