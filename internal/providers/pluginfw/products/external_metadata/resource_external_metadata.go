// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_metadata

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
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

const resourceName = "external_metadata"

var _ resource.ResourceWithConfigure = &ExternalMetadataResource{}

func ResourceExternalMetadata() resource.Resource {
	return &ExternalMetadataResource{}
}

type ExternalMetadataResource struct {
	Client *autogen.DatabricksClient
}

type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(workspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

func workspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// ExternalMetadata extends the main model with additional fields.
type ExternalMetadata struct {
	// List of columns associated with the external metadata object.
	Columns types.List `tfsdk:"columns"`
	// Time at which this external metadata object was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Username of external metadata object creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// User-provided free-form text description.
	Description types.String `tfsdk:"description"`
	// Type of entity within the external system.
	EntityType types.String `tfsdk:"entity_type"`
	// Unique identifier of the external metadata object.
	Id types.String `tfsdk:"id"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of the external metadata object.
	Name types.String `tfsdk:"name"`
	// Owner of the external metadata object.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the external metadata object.
	Properties types.Map `tfsdk:"properties"`
	// Type of external system.
	SystemType types.String `tfsdk:"system_type"`
	// Time at which this external metadata object was last modified.
	UpdateTime types.String `tfsdk:"update_time"`
	// Username of user who last modified external metadata object.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// URL associated with the external metadata object.
	Url            types.String `tfsdk:"url"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ExternalMetadata struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ExternalMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":         reflect.TypeOf(types.String{}),
		"properties":      reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalMetadata
// only implements ToObjectValue() and Type().
func (m ExternalMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"columns": m.Columns,
			"create_time":  m.CreateTime,
			"created_by":   m.CreatedBy,
			"description":  m.Description,
			"entity_type":  m.EntityType,
			"id":           m.Id,
			"metastore_id": m.MetastoreId,
			"name":         m.Name,
			"owner":        m.Owner,
			"properties":   m.Properties,
			"system_type":  m.SystemType,
			"update_time":  m.UpdateTime,
			"updated_by":   m.UpdatedBy,
			"url":          m.Url,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ExternalMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"columns": basetypes.ListType{
			ElemType: types.StringType,
		},
			"create_time":  types.StringType,
			"created_by":   types.StringType,
			"description":  types.StringType,
			"entity_type":  types.StringType,
			"id":           types.StringType,
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"system_type": types.StringType,
			"update_time": types.StringType,
			"updated_by":  types.StringType,
			"url":         types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ExternalMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalMetadata) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *ExternalMetadata) SyncFieldsDuringRead(ctx context.Context, from ExternalMetadata) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m ExternalMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["system_type"] = attrs["system_type"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["url"] = attrs["url"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

// GetColumns returns the value of the Columns field in ExternalMetadata as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalMetadata) GetColumns(ctx context.Context) ([]types.String, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ExternalMetadata.
func (m *ExternalMetadata) SetColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in ExternalMetadata as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalMetadata) GetProperties(ctx context.Context) (map[string]types.String, bool) {
	if m.Properties.IsNull() || m.Properties.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Properties.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProperties sets the value of the Properties field in ExternalMetadata.
func (m *ExternalMetadata) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Properties = types.MapValueMust(t, vs)
}

func (r *ExternalMetadataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ExternalMetadataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ExternalMetadata{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks external_metadata",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalMetadataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ExternalMetadataResource) update(ctx context.Context, plan ExternalMetadata, diags *diag.Diagnostics, state *tfsdk.State) {
	var external_metadata catalog.ExternalMetadata

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &external_metadata)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateExternalMetadataRequest{
		ExternalMetadata: external_metadata,
		Name:             plan.Name.ValueString(),
		UpdateMask:       "columns,description,entity_type,owner,properties,system_type,url",
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
	response, err := client.ExternalMetadata.UpdateExternalMetadata(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ExternalMetadataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ExternalMetadata
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var external_metadata catalog.ExternalMetadata

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &external_metadata)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateExternalMetadataRequest{
		ExternalMetadata: external_metadata,
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

	response, err := client.ExternalMetadata.CreateExternalMetadata(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ExternalMetadataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ExternalMetadata
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetExternalMetadataRequest
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
	response, err := client.ExternalMetadata.GetExternalMetadata(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ExternalMetadataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ExternalMetadata
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ExternalMetadataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ExternalMetadata
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteExternalMetadataRequest
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
	err := client.ExternalMetadata.DeleteExternalMetadata(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete external_metadata", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &ExternalMetadataResource{}

func (r *ExternalMetadataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
