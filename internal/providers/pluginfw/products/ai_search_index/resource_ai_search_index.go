// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_search_index

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/aisearch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/aisearch_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "ai_search_index"

var _ resource.ResourceWithConfigure = &IndexResource{}
var _ resource.ResourceWithModifyPlan = &IndexResource{}

func ResourceIndex() resource.Resource {
	return &IndexResource{}
}

type IndexResource struct {
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

// Index extends the main model with additional fields.
type Index struct {
	// Creator of the index.
	Creator types.String `tfsdk:"creator"`
	// Specification for a Delta Sync index. Set when `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec types.Object `tfsdk:"delta_sync_index_spec"`
	// Specification for a Direct Access index. Set when `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec types.Object `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index. Ignored on create — the
	// endpoint is taken from `CreateIndexRequest.parent`; populated only on
	// output.
	Endpoint types.String `tfsdk:"endpoint"`
	// The user-supplied Unity Catalog table name for the Index, per AIP-133.
	// The server composes the full `Index.name` as
	// `{parent}/indexes/{index_id}`. AIP-133 does not list `index_id` as a
	// fields-may-be-required entry, so we annotate it OPTIONAL on the wire; the
	// server still rejects empty values with INVALID_PARAMETER_VALUE.
	IndexId types.String `tfsdk:"index_id"`
	// The subtype of the index. Set on create and immutable thereafter.
	IndexSubtype types.String `tfsdk:"index_subtype"`
	// Type of index. Required on create and immutable thereafter.
	IndexType types.String `tfsdk:"index_type"`
	// Name of the AI Search index. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on
	// output, where `{index}` is the index's Unity Catalog table name. On
	// create, the user-supplied UC table name is conveyed via
	// `CreateIndexRequest.index_id`; the server composes the full `name` and
	// returns it on the response.
	Name types.String `tfsdk:"name"`
	// The Endpoint where this Index will be created. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent types.String `tfsdk:"parent"`
	// Primary key of the index. Set on create and immutable thereafter.
	PrimaryKey types.String `tfsdk:"primary_key"`
	// Current status of the index.
	Status         types.Object `tfsdk:"status"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Index struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Index) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(aisearch_tf.DeltaSyncIndexSpec{}),
		"direct_access_index_spec": reflect.TypeOf(aisearch_tf.DirectAccessIndexSpec{}),
		"status":                   reflect.TypeOf(aisearch_tf.IndexStatus{}),
		"provider_config":          reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Index
// only implements ToObjectValue() and Type().
func (m Index) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"creator": m.Creator,
			"delta_sync_index_spec":    m.DeltaSyncIndexSpec,
			"direct_access_index_spec": m.DirectAccessIndexSpec,
			"endpoint":                 m.Endpoint,
			"index_id":                 m.IndexId,
			"index_subtype":            m.IndexSubtype,
			"index_type":               m.IndexType,
			"name":                     m.Name,
			"parent":                   m.Parent,
			"primary_key":              m.PrimaryKey,
			"status":                   m.Status,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Index) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"creator": types.StringType,
			"delta_sync_index_spec":    aisearch_tf.DeltaSyncIndexSpec{}.Type(ctx),
			"direct_access_index_spec": aisearch_tf.DirectAccessIndexSpec{}.Type(ctx),
			"endpoint":                 types.StringType,
			"index_id":                 types.StringType,
			"index_subtype":            types.StringType,
			"index_type":               types.StringType,
			"name":                     types.StringType,
			"parent":                   types.StringType,
			"primary_key":              types.StringType,
			"status":                   aisearch_tf.IndexStatus{}.Type(ctx),

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Index) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Index) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				// Recursively sync the fields of DeltaSyncIndexSpec
				toDeltaSyncIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				// Recursively sync the fields of DirectAccessIndexSpec
				toDirectAccessIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
	if !from.IndexId.IsUnknown() {
		to.IndexId = from.IndexId
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
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
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Index) SyncFieldsDuringRead(ctx context.Context, from Index) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				toDeltaSyncIndexSpec.SyncFieldsDuringRead(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				toDirectAccessIndexSpec.SyncFieldsDuringRead(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
	if !from.IndexId.IsUnknown() {
		to.IndexId = from.IndexId
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Index) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["creator"] = attrs["creator"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["endpoint"] = attrs["endpoint"].SetComputed()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["index_subtype"] = attrs["index_subtype"].SetOptional()
	attrs["index_subtype"] = attrs["index_subtype"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["index_type"] = attrs["index_type"].SetRequired()
	attrs["index_type"] = attrs["index_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["primary_key"] = attrs["primary_key"].SetRequired()
	attrs["primary_key"] = attrs["primary_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["index_id"] = attrs["index_id"].SetComputed()
	attrs["index_id"] = attrs["index_id"].SetOptional()
	attrs["index_id"] = attrs["index_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["index_id"] = attrs["index_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in Index as
// a aisearch_tf.DeltaSyncIndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index) GetDeltaSyncIndexSpec(ctx context.Context) (aisearch_tf.DeltaSyncIndexSpec, bool) {
	var e aisearch_tf.DeltaSyncIndexSpec
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.DeltaSyncIndexSpec
	d := m.DeltaSyncIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in Index.
func (m *Index) SetDeltaSyncIndexSpec(ctx context.Context, v aisearch_tf.DeltaSyncIndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.DeltaSyncIndexSpec = vs
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in Index as
// a aisearch_tf.DirectAccessIndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index) GetDirectAccessIndexSpec(ctx context.Context) (aisearch_tf.DirectAccessIndexSpec, bool) {
	var e aisearch_tf.DirectAccessIndexSpec
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.DirectAccessIndexSpec
	d := m.DirectAccessIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in Index.
func (m *Index) SetDirectAccessIndexSpec(ctx context.Context, v aisearch_tf.DirectAccessIndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.DirectAccessIndexSpec = vs
}

// GetStatus returns the value of the Status field in Index as
// a aisearch_tf.IndexStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index) GetStatus(ctx context.Context) (aisearch_tf.IndexStatus, bool) {
	var e aisearch_tf.IndexStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.IndexStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Index.
func (m *Index) SetStatus(ctx context.Context, v aisearch_tf.IndexStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

func (r *IndexResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *IndexResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Index{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ai_search_index",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *IndexResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *IndexResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *IndexResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Index
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var index aisearch.Index

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &index)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := aisearch.CreateIndexRequest{
		Index:   index,
		Parent:  plan.Parent.ValueString(),
		IndexId: plan.IndexId.ValueString(),
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

	response, err := client.AiSearch.CreateIndex(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ai_search_index", err.Error())
		return
	}

	var newState Index

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

func (r *IndexResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Index
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest aisearch.GetIndexRequest
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
	response, err := client.AiSearch.GetIndex(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get ai_search_index", err.Error())
		return
	}

	var newState Index
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

func (r *IndexResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// VARIANT_IMMUTABLE resources do not support updates - all changes require replacement.
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates. All changes require replacement.")
}

func (r *IndexResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Index
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest aisearch.DeleteIndexRequest
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

	err := client.AiSearch.DeleteIndex(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete ai_search_index", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &IndexResource{}

func (r *IndexResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
