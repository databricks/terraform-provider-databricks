// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_gateway_model_service

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

const resourceName = "ai_gateway_model_service"

var _ resource.ResourceWithConfigure = &ModelServiceResource{}
var _ resource.ResourceWithModifyPlan = &ModelServiceResource{}

func ResourceModelService() resource.Resource {
	return &ModelServiceResource{}
}

type ModelServiceResource struct {
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

// ModelService extends the main model with additional fields.
type ModelService struct {
	// User-provided description.
	Comment types.String `tfsdk:"comment"`
	// Operational configuration: destinations, routing, rate limits, inference
	// table. Required on CreateModelService; on UpdateModelService it is
	// required only when `config` (or a `config.*` subpath) appears in
	// `update_mask`.
	Config types.Object `tfsdk:"config"`
	// When the model service was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Creator identity.
	CreatedBy types.String `tfsdk:"created_by"`
	// The resolved owner of the ModelService. Falls back to the caller's
	// identity when `owner` is not explicitly set on creation.
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// Optimistic concurrency control token. Server-generated from the entity's
	// state and returned on every read. To use it as an if-match precondition
	// on a mutation, echo the last-read value back via the dedicated `etag`
	// field on the Update / Delete request; the server rejects the mutation if
	// the stored etag differs.
	Etag types.String `tfsdk:"etag"`
	// Metastore hosting the model service.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name for the model service, e.g. "my_model_service".
	ModelServiceId types.String `tfsdk:"model_service_id"`
	// Resource name of the model service. Format:
	// `model-services/{catalog}.{schema}.{model_service}`. Each `{...}`
	// component is capped at 255 characters individually. Server-derived on
	// Create from `parent` + `model_service_id`; required and immutable on
	// Update/Get/Delete.
	Name types.String `tfsdk:"name"`
	// The owner of the model service. Write-only; read owner via
	// effective_owner.
	Owner types.String `tfsdk:"owner"`
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	Parent types.String `tfsdk:"parent"`
	// Unified API types this endpoint supports (e.g. "chat", "embeddings",
	// "completions"). Derived from the destinations' backing models / providers
	// at read time.
	SupportedApiTypes types.Set `tfsdk:"supported_api_types"`
	// When the model service was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Identity of the last updater.
	UpdatedBy      types.String `tfsdk:"updated_by"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ModelService struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ModelService) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":              reflect.TypeOf(catalog_tf.ModelServiceConfig{}),
		"supported_api_types": reflect.TypeOf(types.String{}),
		"provider_config":     reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelService
// only implements ToObjectValue() and Type().
func (m ModelService) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"comment": m.Comment,
			"config":              m.Config,
			"create_time":         m.CreateTime,
			"created_by":          m.CreatedBy,
			"effective_owner":     m.EffectiveOwner,
			"etag":                m.Etag,
			"metastore_id":        m.MetastoreId,
			"model_service_id":    m.ModelServiceId,
			"name":                m.Name,
			"owner":               m.Owner,
			"parent":              m.Parent,
			"supported_api_types": m.SupportedApiTypes,
			"update_time":         m.UpdateTime,
			"updated_by":          m.UpdatedBy,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ModelService) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"comment": types.StringType,
			"config":           catalog_tf.ModelServiceConfig{}.Type(ctx),
			"create_time":      timetypes.RFC3339{}.Type(ctx),
			"created_by":       types.StringType,
			"effective_owner":  types.StringType,
			"etag":             types.StringType,
			"metastore_id":     types.StringType,
			"model_service_id": types.StringType,
			"name":             types.StringType,
			"owner":            types.StringType,
			"parent":           types.StringType,
			"supported_api_types": basetypes.SetType{
				ElemType: types.StringType,
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by":  types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ModelService) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelService) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.ModelServiceId.IsUnknown() {
		to.ModelServiceId = from.ModelServiceId
	}
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.SupportedApiTypes.IsNull() && !from.SupportedApiTypes.IsUnknown() && to.SupportedApiTypes.IsNull() && len(from.SupportedApiTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupportedApiTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupportedApiTypes = from.SupportedApiTypes
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *ModelService) SyncFieldsDuringRead(ctx context.Context, from ModelService) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.ModelServiceId.IsUnknown() {
		to.ModelServiceId = from.ModelServiceId
	}
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.SupportedApiTypes.IsNull() && !from.SupportedApiTypes.IsUnknown() && to.SupportedApiTypes.IsNull() && len(from.SupportedApiTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupportedApiTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupportedApiTypes = from.SupportedApiTypes
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m ModelService) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["supported_api_types"] = attrs["supported_api_types"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["model_service_id"] = attrs["model_service_id"].SetRequired()
	attrs["model_service_id"] = attrs["model_service_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["model_service_id"] = attrs["model_service_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetConfig returns the value of the Config field in ModelService as
// a catalog_tf.ModelServiceConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelService) GetConfig(ctx context.Context) (catalog_tf.ModelServiceConfig, bool) {
	var e catalog_tf.ModelServiceConfig
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v catalog_tf.ModelServiceConfig
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in ModelService.
func (m *ModelService) SetConfig(ctx context.Context, v catalog_tf.ModelServiceConfig) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

// GetSupportedApiTypes returns the value of the SupportedApiTypes field in ModelService as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelService) GetSupportedApiTypes(ctx context.Context) ([]types.String, bool) {
	if m.SupportedApiTypes.IsNull() || m.SupportedApiTypes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SupportedApiTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSupportedApiTypes sets the value of the SupportedApiTypes field in ModelService.
func (m *ModelService) SetSupportedApiTypes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["supported_api_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SupportedApiTypes = types.SetValueMust(t, vs)
}

func (r *ModelServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ModelServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ModelService{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ai_gateway_model_service",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ModelServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ModelServiceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *ModelServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ModelService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var model_service catalog.ModelService

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &model_service)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateModelServiceRequest{
		ModelService:   model_service,
		Parent:         plan.Parent.ValueString(),
		ModelServiceId: plan.ModelServiceId.ValueString(),
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

	response, err := client.AiGateway.CreateModelService(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ai_gateway_model_service", err.Error())
		return
	}

	var newState ModelService

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

func (r *ModelServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ModelService
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetModelServiceRequest
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
	response, err := client.AiGateway.GetModelService(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get ai_gateway_model_service", err.Error())
		return
	}

	var newState ModelService
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

func (r *ModelServiceResource) update(ctx context.Context, plan ModelService, diags *diag.Diagnostics, state *tfsdk.State) {
	var model_service catalog.ModelService

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &model_service)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateModelServiceRequest{
		ModelService: model_service,
		Name:         plan.Name.ValueString(),
		UpdateMask:   *fieldmask.New(strings.Split("comment,config,owner", ",")),
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
	response, err := client.AiGateway.UpdateModelService(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update ai_gateway_model_service", err.Error())
		return
	}

	var newState ModelService

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ModelServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ModelService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ModelServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ModelService
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteModelServiceRequest
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

	err := client.AiGateway.DeleteModelService(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete ai_gateway_model_service", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ModelServiceResource{}

func (r *ModelServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
