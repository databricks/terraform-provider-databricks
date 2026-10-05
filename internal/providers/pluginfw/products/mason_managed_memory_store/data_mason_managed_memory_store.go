// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_managed_memory_store

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/mason"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/mason_tf"
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

const dataSourceName = "mason_managed_memory_store"

var _ datasource.DataSourceWithConfigure = &ManagedMemoryStoreDataSource{}

func DataSourceManagedMemoryStore() datasource.DataSource {
	return &ManagedMemoryStoreDataSource{}
}

type ManagedMemoryStoreDataSource struct {
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

// ManagedMemoryStoreData extends the main model with additional fields.
type ManagedMemoryStoreData struct {
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
	WorkspaceId        types.Int64  `tfsdk:"workspace_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ManagedMemoryStoreData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ManagedMemoryStoreData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_backend": reflect.TypeOf(mason_tf.StorageBackend{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManagedMemoryStoreData
// only implements ToObjectValue() and Type().
func (m ManagedMemoryStoreData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":     m.CreateTime,
			"creator_user_id": m.CreatorUserId,
			"description":     m.Description,
			"display_name":    m.DisplayName,
			"name":            m.Name,
			"owner_user_id":   m.OwnerUserId,
			"storage_backend": m.StorageBackend,
			"update_time":     m.UpdateTime,
			"workspace_id":    m.WorkspaceId,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ManagedMemoryStoreData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"creator_user_id": types.StringType,
			"description":     types.StringType,
			"display_name":    types.StringType,
			"name":            types.StringType,
			"owner_user_id":   types.StringType,
			"storage_backend": mason_tf.StorageBackend{}.Type(ctx),
			"update_time":     timetypes.RFC3339{}.Type(ctx),
			"workspace_id":    types.Int64Type,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m ManagedMemoryStoreData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner_user_id"] = attrs["owner_user_id"].SetComputed()
	attrs["storage_backend"] = attrs["storage_backend"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *ManagedMemoryStoreDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ManagedMemoryStoreDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ManagedMemoryStoreData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ManagedMemoryStore",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ManagedMemoryStoreDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ManagedMemoryStoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ManagedMemoryStoreData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest mason.GetManagedMemoryStoreRequest
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

	response, err := client.Mason.GetMemoryStore(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get mason_managed_memory_store", err.Error())
		return
	}

	var newState ManagedMemoryStoreData
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
