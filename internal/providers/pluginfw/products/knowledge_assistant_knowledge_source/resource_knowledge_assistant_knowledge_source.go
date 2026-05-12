// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledge_assistant_knowledge_source

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/knowledgeassistants_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "knowledge_assistant_knowledge_source"

var _ resource.ResourceWithConfigure = &KnowledgeSourceResource{}
var _ resource.ResourceWithModifyPlan = &KnowledgeSourceResource{}

func ResourceKnowledgeSource() resource.Resource {
	return &KnowledgeSourceResource{}
}

type KnowledgeSourceResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// KnowledgeSource extends the main model with additional fields.
type KnowledgeSource struct {
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
	// Parent resource where this source will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"parent"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType types.String `tfsdk:"source_type"`

	State          types.String `tfsdk:"state"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// KnowledgeSource struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m KnowledgeSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_table":      reflect.TypeOf(knowledgeassistants_tf.FileTableSpec{}),
		"files":           reflect.TypeOf(knowledgeassistants_tf.FilesSpec{}),
		"index":           reflect.TypeOf(knowledgeassistants_tf.IndexSpec{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeSource
// only implements ToObjectValue() and Type().
func (m KnowledgeSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"file_table":            m.FileTable,
			"files":                 m.Files,
			"id":                    m.Id,
			"index":                 m.Index,
			"knowledge_cutoff_time": m.KnowledgeCutoffTime,
			"name":                  m.Name,
			"parent":                m.Parent,
			"source_type":           m.SourceType,
			"state":                 m.State,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m KnowledgeSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"file_table":            knowledgeassistants_tf.FileTableSpec{}.Type(ctx),
			"files":                 knowledgeassistants_tf.FilesSpec{}.Type(ctx),
			"id":                    types.StringType,
			"index":                 knowledgeassistants_tf.IndexSpec{}.Type(ctx),
			"knowledge_cutoff_time": timetypes.RFC3339{}.Type(ctx),
			"name":                  types.StringType,
			"parent":                types.StringType,
			"source_type":           types.StringType,
			"state":                 types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *KnowledgeSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeSource) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				// Recursively sync the fields of FileTable
				toFileTable.SyncFieldsDuringCreateOrUpdate(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of Files
				toFiles.SyncFieldsDuringCreateOrUpdate(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				// Recursively sync the fields of Index
				toIndex.SyncFieldsDuringCreateOrUpdate(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *KnowledgeSource) SyncFieldsDuringRead(ctx context.Context, from KnowledgeSource) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				toFileTable.SyncFieldsDuringRead(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				toFiles.SyncFieldsDuringRead(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				toIndex.SyncFieldsDuringRead(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m KnowledgeSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["file_table"] = attrs["file_table"].SetOptional()
	attrs["files"] = attrs["files"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["knowledge_cutoff_time"] = attrs["knowledge_cutoff_time"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["source_type"] = attrs["source_type"].SetRequired()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetFileTable returns the value of the FileTable field in KnowledgeSource as
// a knowledgeassistants_tf.FileTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetFileTable(ctx context.Context) (knowledgeassistants_tf.FileTableSpec, bool) {
	var e knowledgeassistants_tf.FileTableSpec
	if m.FileTable.IsNull() || m.FileTable.IsUnknown() {
		return e, false
	}
	var v knowledgeassistants_tf.FileTableSpec
	d := m.FileTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileTable sets the value of the FileTable field in KnowledgeSource.
func (m *KnowledgeSource) SetFileTable(ctx context.Context, v knowledgeassistants_tf.FileTableSpec) {
	vs := v.ToObjectValue(ctx)
	m.FileTable = vs
}

// GetFiles returns the value of the Files field in KnowledgeSource as
// a knowledgeassistants_tf.FilesSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetFiles(ctx context.Context) (knowledgeassistants_tf.FilesSpec, bool) {
	var e knowledgeassistants_tf.FilesSpec
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return e, false
	}
	var v knowledgeassistants_tf.FilesSpec
	d := m.Files.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in KnowledgeSource.
func (m *KnowledgeSource) SetFiles(ctx context.Context, v knowledgeassistants_tf.FilesSpec) {
	vs := v.ToObjectValue(ctx)
	m.Files = vs
}

// GetIndex returns the value of the Index field in KnowledgeSource as
// a knowledgeassistants_tf.IndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetIndex(ctx context.Context) (knowledgeassistants_tf.IndexSpec, bool) {
	var e knowledgeassistants_tf.IndexSpec
	if m.Index.IsNull() || m.Index.IsUnknown() {
		return e, false
	}
	var v knowledgeassistants_tf.IndexSpec
	d := m.Index.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndex sets the value of the Index field in KnowledgeSource.
func (m *KnowledgeSource) SetIndex(ctx context.Context, v knowledgeassistants_tf.IndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.Index = vs
}

func (r *KnowledgeSourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *KnowledgeSourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, KnowledgeSource{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks knowledge_assistant_knowledge_source",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KnowledgeSourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *KnowledgeSourceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *KnowledgeSourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KnowledgeSource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var knowledge_source knowledgeassistants.KnowledgeSource

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &knowledge_source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := knowledgeassistants.CreateKnowledgeSourceRequest{
		KnowledgeSource: knowledge_source,
		Parent:          plan.Parent.ValueString(),
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.KnowledgeAssistants.CreateKnowledgeSource(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create knowledge_assistant_knowledge_source", err.Error())
		return
	}

	var newState KnowledgeSource

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *KnowledgeSourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState KnowledgeSource
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest knowledgeassistants.GetKnowledgeSourceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get knowledge_assistant_knowledge_source", err.Error())
		return
	}

	var newState KnowledgeSource
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *KnowledgeSourceResource) update(ctx context.Context, plan KnowledgeSource, diags *diag.Diagnostics, state *tfsdk.State) {
	var knowledge_source knowledgeassistants.KnowledgeSource

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &knowledge_source)...)
	if diags.HasError() {
		return
	}

	updateRequest := knowledgeassistants.UpdateKnowledgeSourceRequest{
		KnowledgeSource: knowledge_source,
		Name:            plan.Name.ValueString(),
		UpdateMask:      *fieldmask.New(strings.Split("description,display_name,file_table,files,index,source_type", ",")),
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.KnowledgeAssistants.UpdateKnowledgeSource(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update knowledge_assistant_knowledge_source", err.Error())
		return
	}

	var newState KnowledgeSource

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *KnowledgeSourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KnowledgeSource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *KnowledgeSourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state KnowledgeSource
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest knowledgeassistants.DeleteKnowledgeSourceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	err := client.KnowledgeAssistants.DeleteKnowledgeSource(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete knowledge_assistant_knowledge_source", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &KnowledgeSourceResource{}

func (r *KnowledgeSourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
