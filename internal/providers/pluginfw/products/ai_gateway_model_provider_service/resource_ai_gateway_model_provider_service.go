// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_gateway_model_provider_service

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
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

const resourceName = "ai_gateway_model_provider_service"

var _ resource.ResourceWithConfigure = &ModelProviderServiceResource{}
var _ resource.ResourceWithModifyPlan = &ModelProviderServiceResource{}

func ResourceModelProviderService() resource.Resource {
	return &ModelProviderServiceResource{}
}

type ModelProviderServiceResource struct {
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

// ModelProviderService extends the main model with additional fields.
type ModelProviderService struct {
	// User-provided description.
	Comment types.String `tfsdk:"comment"`
	// Behavioral configuration: provider connection, model catalog, and
	// passthrough policy. See `ModelProviderServiceConfig` for the per-field
	// contract. Required on CreateModelProviderService; on Update it is
	// required only when `config` (or a `config.*` subpath) appears in
	// `update_mask`.
	Config types.Object `tfsdk:"config"`
	// When the provider service was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Creator identity.
	CreatedBy types.String `tfsdk:"created_by"`
	// The resolved owner of the model provider service. Falls back to the
	// caller's identity when `owner` is not explicitly set on creation.
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// Optimistic concurrency control token. Server-generated from the entity's
	// state and returned on every read. To use it as an if-match precondition
	// on a mutation, echo the last-read value back via the dedicated `etag`
	// field on the Update / Delete request; the server rejects the mutation if
	// the stored etag differs.
	Etag types.String `tfsdk:"etag"`
	// Metastore hosting the provider service.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Leaf identifier for the provider service (the unqualified name within the
	// parent schema, e.g. "openai_prod").
	ModelProviderServiceId types.String `tfsdk:"model_provider_service_id"`
	// Resource name of the provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually.
	// Server-derived on Create from `parent` + `model_provider_service_id`;
	// required and immutable on Update/Get/Delete.
	Name types.String `tfsdk:"name"`
	// The owner of the model provider service. Write-only; read owner via
	// effective_owner.
	Owner types.String `tfsdk:"owner"`
	// Resource name of the parent schema. Format: `schemas/{catalog}.{schema}`.
	// Each `{...}` component is capped at 255 characters individually.
	Parent types.String `tfsdk:"parent"`
	// When the provider service was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Identity of the last updater.
	UpdatedBy      types.String `tfsdk:"updated_by"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ModelProviderService struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ModelProviderService) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":          reflect.TypeOf(catalog_tf.ModelProviderServiceConfig{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelProviderService
// only implements ToObjectValue() and Type().
func (m ModelProviderService) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"comment": m.Comment,
			"config":                    m.Config,
			"create_time":               m.CreateTime,
			"created_by":                m.CreatedBy,
			"effective_owner":           m.EffectiveOwner,
			"etag":                      m.Etag,
			"metastore_id":              m.MetastoreId,
			"model_provider_service_id": m.ModelProviderServiceId,
			"name":                      m.Name,
			"owner":                     m.Owner,
			"parent":                    m.Parent,
			"update_time":               m.UpdateTime,
			"updated_by":                m.UpdatedBy,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ModelProviderService) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"comment": types.StringType,
			"config":                    catalog_tf.ModelProviderServiceConfig{}.Type(ctx),
			"create_time":               timetypes.RFC3339{}.Type(ctx),
			"created_by":                types.StringType,
			"effective_owner":           types.StringType,
			"etag":                      types.StringType,
			"metastore_id":              types.StringType,
			"model_provider_service_id": types.StringType,
			"name":                      types.StringType,
			"owner":                     types.StringType,
			"parent":                    types.StringType,
			"update_time":               timetypes.RFC3339{}.Type(ctx),
			"updated_by":                types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ModelProviderService) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelProviderService) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.ModelProviderServiceId.IsUnknown() {
		to.ModelProviderServiceId = from.ModelProviderServiceId
	}
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *ModelProviderService) SyncFieldsDuringRead(ctx context.Context, from ModelProviderService) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.ModelProviderServiceId.IsUnknown() {
		to.ModelProviderServiceId = from.ModelProviderServiceId
	}
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m ModelProviderService) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["owner"] = attrs["owner"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["model_provider_service_id"] = attrs["model_provider_service_id"].SetRequired()
	attrs["model_provider_service_id"] = attrs["model_provider_service_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["model_provider_service_id"] = attrs["model_provider_service_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetConfig returns the value of the Config field in ModelProviderService as
// a catalog_tf.ModelProviderServiceConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelProviderService) GetConfig(ctx context.Context) (catalog_tf.ModelProviderServiceConfig, bool) {
	var e catalog_tf.ModelProviderServiceConfig
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v catalog_tf.ModelProviderServiceConfig
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in ModelProviderService.
func (m *ModelProviderService) SetConfig(ctx context.Context, v catalog_tf.ModelProviderServiceConfig) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

func (r *ModelProviderServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ModelProviderServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ModelProviderService{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ai_gateway_model_provider_service",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ModelProviderServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ModelProviderServiceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *ModelProviderServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ModelProviderService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var model_provider_service catalog.ModelProviderService

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &model_provider_service)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateModelProviderServiceRequest{
		ModelProviderService:   model_provider_service,
		Parent:                 plan.Parent.ValueString(),
		ModelProviderServiceId: plan.ModelProviderServiceId.ValueString(),
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

	response, err := client.AiGateway.CreateModelProviderService(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ai_gateway_model_provider_service", err.Error())
		return
	}

	var newState ModelProviderService

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

func (r *ModelProviderServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ModelProviderService
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetModelProviderServiceRequest
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
	response, err := client.AiGateway.GetModelProviderService(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get ai_gateway_model_provider_service", err.Error())
		return
	}

	var newState ModelProviderService
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

func (r *ModelProviderServiceResource) update(ctx context.Context, plan ModelProviderService, diags *diag.Diagnostics, state *tfsdk.State) {
	var model_provider_service catalog.ModelProviderService

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &model_provider_service)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateModelProviderServiceRequest{
		ModelProviderService: model_provider_service,
		Name:                 plan.Name.ValueString(),
		UpdateMask:           *fieldmask.New(strings.Split("comment,config,owner", ",")),
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
	response, err := client.AiGateway.UpdateModelProviderService(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update ai_gateway_model_provider_service", err.Error())
		return
	}

	var newState ModelProviderService

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ModelProviderServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ModelProviderService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ModelProviderServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ModelProviderService
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteModelProviderServiceRequest
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

	err := client.AiGateway.DeleteModelProviderService(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete ai_gateway_model_provider_service", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ModelProviderServiceResource{}

func (r *ModelProviderServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
