// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_managed_memory_store

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/mason"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/mason_tf"
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

const resourceName = "mason_managed_memory_store"

var _ resource.ResourceWithConfigure = &ManagedMemoryStoreResource{}
var _ resource.ResourceWithModifyPlan = &ManagedMemoryStoreResource{}

func ResourceManagedMemoryStore() resource.Resource {
	return &ManagedMemoryStoreResource{}
}

type ManagedMemoryStoreResource struct {
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

// ManagedMemoryStore extends the main model with additional fields.
type ManagedMemoryStore struct {
	// Time when the store was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Workspace-local user ID of the authenticated principal that created the
	// store. This is immutable server-set attribution and does not grant
	// access; authorization is evaluated from the authenticated request
	// context.
	CreatorUserId types.String `tfsdk:"creator_user_id"`
	// Human-readable description of the memory store.
	Description types.String `tfsdk:"description"`
	// Deprecated compatibility alias for the caller-provided managed memory
	// store ID. Canonical clients provide the ID through
	// `CreateMemoryStoreRequest.managed_memory_store_id` and use `name` as the
	// resource identifier.
	DisplayName types.String `tfsdk:"display_name"`
	// Caller-provided, workspace-unique managed memory store ID. It must be
	// 3-56 characters, begin with a lowercase letter, contain only lowercase
	// letters, digits, and hyphens, and end with a letter or digit.
	ManagedMemoryStoreId types.String `tfsdk:"managed_memory_store_id"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name types.String `tfsdk:"name"`
	// Deprecated alias for `creator_user_id`. This identifies the original
	// creator, not a transferable owner. Use `creator_user_id` instead.
	OwnerUserId types.String `tfsdk:"owner_user_id"`
	// Service-managed storage backing this memory store.
	StorageBackend types.Object `tfsdk:"storage_backend"`
	// Time when the store was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Workspace that owns the memory store.
	WorkspaceId    types.Int64  `tfsdk:"workspace_id"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ManagedMemoryStore struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ManagedMemoryStore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_backend": reflect.TypeOf(mason_tf.StorageBackend{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManagedMemoryStore
// only implements ToObjectValue() and Type().
func (m ManagedMemoryStore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"creator_user_id":         m.CreatorUserId,
			"description":             m.Description,
			"display_name":            m.DisplayName,
			"managed_memory_store_id": m.ManagedMemoryStoreId,
			"name":                    m.Name,
			"owner_user_id":           m.OwnerUserId,
			"storage_backend":         m.StorageBackend,
			"update_time":             m.UpdateTime,
			"workspace_id":            m.WorkspaceId,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ManagedMemoryStore) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"creator_user_id":         types.StringType,
			"description":             types.StringType,
			"display_name":            types.StringType,
			"managed_memory_store_id": types.StringType,
			"name":                    types.StringType,
			"owner_user_id":           types.StringType,
			"storage_backend":         mason_tf.StorageBackend{}.Type(ctx),
			"update_time":             timetypes.RFC3339{}.Type(ctx),
			"workspace_id":            types.Int64Type,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ManagedMemoryStore) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManagedMemoryStore) {
	if !from.ManagedMemoryStoreId.IsUnknown() {
		to.ManagedMemoryStoreId = from.ManagedMemoryStoreId
	}
	if !from.StorageBackend.IsNull() && !from.StorageBackend.IsUnknown() {
		if toStorageBackend, ok := to.GetStorageBackend(ctx); ok {
			if fromStorageBackend, ok := from.GetStorageBackend(ctx); ok {
				// Recursively sync the fields of StorageBackend
				toStorageBackend.SyncFieldsDuringCreateOrUpdate(ctx, fromStorageBackend)
				to.SetStorageBackend(ctx, toStorageBackend)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *ManagedMemoryStore) SyncFieldsDuringRead(ctx context.Context, from ManagedMemoryStore) {
	if !from.ManagedMemoryStoreId.IsUnknown() {
		to.ManagedMemoryStoreId = from.ManagedMemoryStoreId
	}
	if !from.StorageBackend.IsNull() && !from.StorageBackend.IsUnknown() {
		if toStorageBackend, ok := to.GetStorageBackend(ctx); ok {
			if fromStorageBackend, ok := from.GetStorageBackend(ctx); ok {
				toStorageBackend.SyncFieldsDuringRead(ctx, fromStorageBackend)
				to.SetStorageBackend(ctx, toStorageBackend)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m ManagedMemoryStore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["owner_user_id"] = attrs["owner_user_id"].SetComputed()
	attrs["storage_backend"] = attrs["storage_backend"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["managed_memory_store_id"] = attrs["managed_memory_store_id"].SetRequired()
	attrs["managed_memory_store_id"] = attrs["managed_memory_store_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["managed_memory_store_id"] = attrs["managed_memory_store_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetStorageBackend returns the value of the StorageBackend field in ManagedMemoryStore as
// a mason_tf.StorageBackend value.
// If the field is unknown or null, the boolean return value is false.
func (m *ManagedMemoryStore) GetStorageBackend(ctx context.Context) (mason_tf.StorageBackend, bool) {
	var e mason_tf.StorageBackend
	if m.StorageBackend.IsNull() || m.StorageBackend.IsUnknown() {
		return e, false
	}
	var v mason_tf.StorageBackend
	d := m.StorageBackend.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStorageBackend sets the value of the StorageBackend field in ManagedMemoryStore.
func (m *ManagedMemoryStore) SetStorageBackend(ctx context.Context, v mason_tf.StorageBackend) {
	vs := v.ToObjectValue(ctx)
	m.StorageBackend = vs
}

func (r *ManagedMemoryStoreResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ManagedMemoryStoreResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ManagedMemoryStore{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks mason_managed_memory_store",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ManagedMemoryStoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ManagedMemoryStoreResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *ManagedMemoryStoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ManagedMemoryStore
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var managed_memory_store mason.ManagedMemoryStore

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &managed_memory_store)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := mason.CreateManagedMemoryStoreRequest{
		ManagedMemoryStore:   managed_memory_store,
		ManagedMemoryStoreId: plan.ManagedMemoryStoreId.ValueString(),
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

	response, err := client.Mason.CreateMemoryStore(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create mason_managed_memory_store", err.Error())
		return
	}

	var newState ManagedMemoryStore

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

func (r *ManagedMemoryStoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ManagedMemoryStore
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest mason.GetManagedMemoryStoreRequest
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
	response, err := client.Mason.GetMemoryStore(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get mason_managed_memory_store", err.Error())
		return
	}

	var newState ManagedMemoryStore
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

func (r *ManagedMemoryStoreResource) update(ctx context.Context, plan ManagedMemoryStore, diags *diag.Diagnostics, state *tfsdk.State) {
	var managed_memory_store mason.ManagedMemoryStore

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &managed_memory_store)...)
	if diags.HasError() {
		return
	}

	updateRequest := mason.UpdateManagedMemoryStoreRequest{
		ManagedMemoryStore: managed_memory_store,
		Name:               plan.Name.ValueString(),
		UpdateMask:         *fieldmask.New(strings.Split("description", ",")),
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
	response, err := client.Mason.UpdateMemoryStore(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update mason_managed_memory_store", err.Error())
		return
	}

	var newState ManagedMemoryStore

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ManagedMemoryStoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ManagedMemoryStore
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ManagedMemoryStoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ManagedMemoryStore
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest mason.DeleteManagedMemoryStoreRequest
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

	err := client.Mason.DeleteMemoryStore(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete mason_managed_memory_store", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ManagedMemoryStoreResource{}

func (r *ManagedMemoryStoreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
