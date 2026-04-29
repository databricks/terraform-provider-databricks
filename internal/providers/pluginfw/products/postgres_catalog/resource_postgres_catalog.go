// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_catalog

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/postgres_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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

const resourceName = "postgres_catalog"

var _ resource.ResourceWithConfigure = &CatalogResource{}
var _ resource.ResourceWithModifyPlan = &CatalogResource{}

func ResourceCatalog() resource.Resource {
	return &CatalogResource{}
}

type CatalogResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
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

// Catalog extends the main model with additional fields.
type Catalog struct {
	// The ID in the Unity Catalog. It becomes the full resource name, for
	// example "my_catalog" becomes "catalogs/my_catalog".
	CatalogId types.String `tfsdk:"catalog_id"`
	// A timestamp indicating when the catalog was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Output only. The full resource path of the catalog.
	//
	// Format: "catalogs/{catalog_id}".
	Name types.String `tfsdk:"name"`
	// The desired state of the Catalog.
	Spec types.Object `tfsdk:"spec"`
	// The observed state of the Catalog.
	Status types.Object `tfsdk:"status"`
	// System-generated unique identifier for the catalog.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the catalog was last updated.
	UpdateTime     timetypes.RFC3339 `tfsdk:"update_time"`
	ProviderConfig types.Object      `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Catalog struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Catalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":            reflect.TypeOf(postgres_tf.CatalogCatalogSpec{}),
		"status":          reflect.TypeOf(postgres_tf.CatalogCatalogStatus{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Catalog
// only implements ToObjectValue() and Type().
func (m Catalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"catalog_id": m.CatalogId,
			"create_time": m.CreateTime,
			"name":        m.Name,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Catalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"catalog_id": types.StringType,
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"spec":        postgres_tf.CatalogCatalogSpec{}.Type(ctx),
			"status":      postgres_tf.CatalogCatalogStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Catalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Catalog) {
	if !from.CatalogId.IsUnknown() {
		to.CatalogId = from.CatalogId
	}
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
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
func (to *Catalog) SyncFieldsDuringRead(ctx context.Context, from Catalog) {
	if !from.CatalogId.IsUnknown() {
		to.CatalogId = from.CatalogId
	}
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
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

func (m Catalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["uid"] = attrs["uid"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["update_time"] = attrs["update_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["catalog_id"] = attrs["catalog_id"].SetRequired()
	attrs["catalog_id"] = attrs["catalog_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["catalog_id"] = attrs["catalog_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

// GetSpec returns the value of the Spec field in Catalog as
// a postgres_tf.CatalogCatalogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Catalog) GetSpec(ctx context.Context) (postgres_tf.CatalogCatalogSpec, bool) {
	var e postgres_tf.CatalogCatalogSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v postgres_tf.CatalogCatalogSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Catalog.
func (m *Catalog) SetSpec(ctx context.Context, v postgres_tf.CatalogCatalogSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Catalog as
// a postgres_tf.CatalogCatalogStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Catalog) GetStatus(ctx context.Context) (postgres_tf.CatalogCatalogStatus, bool) {
	var e postgres_tf.CatalogCatalogStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v postgres_tf.CatalogCatalogStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Catalog.
func (m *Catalog) SetStatus(ctx context.Context, v postgres_tf.CatalogCatalogStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

func (r *CatalogResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CatalogResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Catalog{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks postgres_catalog",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CatalogResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *CatalogResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip validation on destroy plans (plan is null).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	var plan Catalog
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, validateDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())
	resp.Diagnostics.Append(validateDiags...)
}

func (r *CatalogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Catalog
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var catalog postgres.Catalog

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &catalog)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := postgres.CreateCatalogRequest{
		Catalog:   catalog,
		CatalogId: plan.CatalogId.ValueString(),
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

	response, err := client.Postgres.CreateCatalog(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create postgres_catalog", err.Error())
		return
	}

	var newState Catalog

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for postgres_catalog to be ready", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CatalogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Catalog
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetCatalogRequest
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
	response, err := client.Postgres.GetCatalog(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get postgres_catalog", err.Error())
		return
	}

	var newState Catalog
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CatalogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// VARIANT_IMMUTABLE resources do not support updates - all changes require replacement.
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates. All changes require replacement.")
}

func (r *CatalogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Catalog
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest postgres.DeleteCatalogRequest
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

	response, err := client.Postgres.DeleteCatalog(ctx, deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete postgres_catalog", err.Error())
		return
	}

	err = response.Wait(ctx)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("error waiting for postgres_catalog delete", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &CatalogResource{}

func (r *CatalogResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
