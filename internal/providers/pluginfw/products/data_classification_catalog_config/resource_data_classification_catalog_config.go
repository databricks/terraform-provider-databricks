// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_classification_catalog_config

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/dataclassification"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/dataclassification_tf"
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

const resourceName = "data_classification_catalog_config"

var _ resource.ResourceWithConfigure = &CatalogConfigResource{}
var _ resource.ResourceWithModifyPlan = &CatalogConfigResource{}

func ResourceCatalogConfig() resource.Resource {
	return &CatalogConfigResource{}
}

type CatalogConfigResource struct {
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

// CatalogConfig extends the main model with additional fields.
type CatalogConfig struct {
	// List of auto-tagging configurations for this catalog. Empty list means no
	// auto-tagging is enabled.
	AutoTagConfigs types.List `tfsdk:"auto_tag_configs"`
	// Schemas to include in the scan. Empty list is not supported as it results
	// in a no-op scan. If `included_schemas` is not set, all schemas are
	// scanned.
	IncludedSchemas types.Object `tfsdk:"included_schemas"`
	// Resource name in the format: catalogs/{catalog_name}/config.
	Name types.String `tfsdk:"name"`
	// Parent resource in the format: catalogs/{catalog_name}
	Parent         types.String `tfsdk:"parent"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// CatalogConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m CatalogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_tag_configs": reflect.TypeOf(dataclassification_tf.AutoTaggingConfig{}),
		"included_schemas": reflect.TypeOf(dataclassification_tf.CatalogConfigSchemaNames{}),
		"provider_config":  reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CatalogConfig
// only implements ToObjectValue() and Type().
func (m CatalogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"auto_tag_configs": m.AutoTagConfigs,
			"included_schemas": m.IncludedSchemas,
			"name":             m.Name,
			"parent":           m.Parent,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m CatalogConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"auto_tag_configs": basetypes.ListType{
			ElemType: dataclassification_tf.AutoTaggingConfig{}.Type(ctx),
		},
			"included_schemas": dataclassification_tf.CatalogConfigSchemaNames{}.Type(ctx),
			"name":             types.StringType,
			"parent":           types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *CatalogConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CatalogConfig) {
	if !from.AutoTagConfigs.IsNull() && !from.AutoTagConfigs.IsUnknown() && to.AutoTagConfigs.IsNull() && len(from.AutoTagConfigs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AutoTagConfigs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AutoTagConfigs = from.AutoTagConfigs
	}
	if !from.IncludedSchemas.IsNull() && !from.IncludedSchemas.IsUnknown() {
		if toIncludedSchemas, ok := to.GetIncludedSchemas(ctx); ok {
			if fromIncludedSchemas, ok := from.GetIncludedSchemas(ctx); ok {
				// Recursively sync the fields of IncludedSchemas
				toIncludedSchemas.SyncFieldsDuringCreateOrUpdate(ctx, fromIncludedSchemas)
				to.SetIncludedSchemas(ctx, toIncludedSchemas)
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
func (to *CatalogConfig) SyncFieldsDuringRead(ctx context.Context, from CatalogConfig) {
	if !from.AutoTagConfigs.IsNull() && !from.AutoTagConfigs.IsUnknown() && to.AutoTagConfigs.IsNull() && len(from.AutoTagConfigs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AutoTagConfigs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AutoTagConfigs = from.AutoTagConfigs
	}
	if !from.IncludedSchemas.IsNull() && !from.IncludedSchemas.IsUnknown() {
		if toIncludedSchemas, ok := to.GetIncludedSchemas(ctx); ok {
			if fromIncludedSchemas, ok := from.GetIncludedSchemas(ctx); ok {
				toIncludedSchemas.SyncFieldsDuringRead(ctx, fromIncludedSchemas)
				to.SetIncludedSchemas(ctx, toIncludedSchemas)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m CatalogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_tag_configs"] = attrs["auto_tag_configs"].SetOptional()
	attrs["included_schemas"] = attrs["included_schemas"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetAutoTagConfigs returns the value of the AutoTagConfigs field in CatalogConfig as
// a slice of dataclassification_tf.AutoTaggingConfig values.
// If the field is unknown or null, the boolean return value is false.
func (m *CatalogConfig) GetAutoTagConfigs(ctx context.Context) ([]dataclassification_tf.AutoTaggingConfig, bool) {
	if m.AutoTagConfigs.IsNull() || m.AutoTagConfigs.IsUnknown() {
		return nil, false
	}
	var v []dataclassification_tf.AutoTaggingConfig
	d := m.AutoTagConfigs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoTagConfigs sets the value of the AutoTagConfigs field in CatalogConfig.
func (m *CatalogConfig) SetAutoTagConfigs(ctx context.Context, v []dataclassification_tf.AutoTaggingConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_tag_configs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AutoTagConfigs = types.ListValueMust(t, vs)
}

// GetIncludedSchemas returns the value of the IncludedSchemas field in CatalogConfig as
// a dataclassification_tf.CatalogConfigSchemaNames value.
// If the field is unknown or null, the boolean return value is false.
func (m *CatalogConfig) GetIncludedSchemas(ctx context.Context) (dataclassification_tf.CatalogConfigSchemaNames, bool) {
	var e dataclassification_tf.CatalogConfigSchemaNames
	if m.IncludedSchemas.IsNull() || m.IncludedSchemas.IsUnknown() {
		return e, false
	}
	var v dataclassification_tf.CatalogConfigSchemaNames
	d := m.IncludedSchemas.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIncludedSchemas sets the value of the IncludedSchemas field in CatalogConfig.
func (m *CatalogConfig) SetIncludedSchemas(ctx context.Context, v dataclassification_tf.CatalogConfigSchemaNames) {
	vs := v.ToObjectValue(ctx)
	m.IncludedSchemas = vs
}

func (r *CatalogConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CatalogConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, CatalogConfig{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks data_classification_catalog_config",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CatalogConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *CatalogConfigResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *CatalogConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan CatalogConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var catalog_config dataclassification.CatalogConfig

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &catalog_config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := dataclassification.CreateCatalogConfigRequest{
		CatalogConfig: catalog_config,
		Parent:        plan.Parent.ValueString(),
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

	response, err := client.DataClassification.CreateCatalogConfig(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create data_classification_catalog_config", err.Error())
		return
	}

	var newState CatalogConfig

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

func (r *CatalogConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState CatalogConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest dataclassification.GetCatalogConfigRequest
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
	response, err := client.DataClassification.GetCatalogConfig(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get data_classification_catalog_config", err.Error())
		return
	}

	var newState CatalogConfig
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

func (r *CatalogConfigResource) update(ctx context.Context, plan CatalogConfig, diags *diag.Diagnostics, state *tfsdk.State) {
	var catalog_config dataclassification.CatalogConfig

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &catalog_config)...)
	if diags.HasError() {
		return
	}

	updateRequest := dataclassification.UpdateCatalogConfigRequest{
		CatalogConfig: catalog_config,
		Name:          plan.Name.ValueString(),
		UpdateMask:    *fieldmask.New(strings.Split("auto_tag_configs,included_schemas", ",")),
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
	response, err := client.DataClassification.UpdateCatalogConfig(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update data_classification_catalog_config", err.Error())
		return
	}

	var newState CatalogConfig

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CatalogConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan CatalogConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *CatalogConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state CatalogConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest dataclassification.DeleteCatalogConfigRequest
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

	err := client.DataClassification.DeleteCatalogConfig(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete data_classification_catalog_config", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &CatalogConfigResource{}

func (r *CatalogConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
