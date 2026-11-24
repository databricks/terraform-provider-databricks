// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_parser_version

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "lakewatch_parser_version"

var _ resource.ResourceWithConfigure = &ParserVersionResource{}
var _ resource.ResourceWithModifyPlan = &ParserVersionResource{}

func ResourceParserVersion() resource.Resource {
	return &ParserVersionResource{}
}

type ParserVersionResource struct {
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

// ParserVersion extends the main model with additional fields.
type ParserVersion struct {
	// When this version was appended.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The principal that appended this version.
	CreatedBy types.Object `tfsdk:"created_by"`
	// Free-form description of this version, set when the version is created.
	Description types.String `tfsdk:"description"`
	// Resource name. Format: `parsers/{parser_id}/versions/{version}`.
	Name types.String `tfsdk:"name"`
	// Resource name of the parser to append a version to. Format:
	// `parsers/{parser_id}`.
	Parent types.String `tfsdk:"parent"`
	// The parent parser's name, copied from the container on read. Denormalized
	// so a caller that needs identity alongside the spec reads the version
	// alone rather than the container as well. Populated by `GetParserVersion`;
	// List omits it, as List omits `spec`.
	ParserName types.String `tfsdk:"parser_name"`
	// The field definitions for this log source, as a YAML document. Validated
	// on write, so an invalid spec is never stored.
	Spec types.String `tfsdk:"spec"`
	// Version number, assigned by the server when the version is appended.
	Version        types.Int64  `tfsdk:"version"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ParserVersion struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ParserVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"created_by":      reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ParserVersion
// only implements ToObjectValue() and Type().
func (m ParserVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"created_by":  m.CreatedBy,
			"description": m.Description,
			"name":        m.Name,
			"parent":      m.Parent,
			"parser_name": m.ParserName,
			"spec":        m.Spec,
			"version":     m.Version,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ParserVersion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"created_by":  lakewatch_tf.Principal{}.Type(ctx),
			"description": types.StringType,
			"name":        types.StringType,
			"parent":      types.StringType,
			"parser_name": types.StringType,
			"spec":        types.StringType,
			"version":     types.Int64Type,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ParserVersion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ParserVersion) {
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				// Recursively sync the fields of CreatedBy
				toCreatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
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
func (to *ParserVersion) SyncFieldsDuringRead(ctx context.Context, from ParserVersion) {
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				toCreatedBy.SyncFieldsDuringRead(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m ParserVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["created_by"] = attrs["created_by"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["description"] = attrs["description"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["parser_name"] = attrs["parser_name"].SetComputed()
	attrs["parser_name"] = attrs["parser_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].SetRequired()
	attrs["spec"] = attrs["spec"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["version"] = attrs["version"].SetComputed()
	attrs["version"] = attrs["version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetCreatedBy returns the value of the CreatedBy field in ParserVersion as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *ParserVersion) GetCreatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.CreatedBy.IsNull() || m.CreatedBy.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.CreatedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCreatedBy sets the value of the CreatedBy field in ParserVersion.
func (m *ParserVersion) SetCreatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.CreatedBy = vs
}

func (r *ParserVersionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ParserVersionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ParserVersion{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks lakewatch_parser_version",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ParserVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ParserVersionResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *ParserVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ParserVersion
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var parser_version lakewatch.ParserVersion

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &parser_version)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := lakewatch.CreateParserVersionRequest{
		ParserVersion: parser_version,
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

	response, err := client.Lakewatch.CreateParserVersion(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create lakewatch_parser_version", err.Error())
		return
	}

	var newState ParserVersion

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
	if !resp.Diagnostics.HasError() {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, tfschema.CreationInputsAppliedKey, []byte("true"))...)
	}
}

func (r *ParserVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ParserVersion
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetParserVersionRequest
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
	response, err := client.Lakewatch.GetParserVersion(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get lakewatch_parser_version", err.Error())
		return
	}

	var newState ParserVersion
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

func (r *ParserVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	if tfschema.AdoptImportedCreateFields(ctx, req, resp, "parent") {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.Append(resp.Private.SetKey(ctx, tfschema.CreationInputsAppliedKey, []byte("true"))...)
		}
		return
	}
	// VARIANT_IMMUTABLE resources do not support updates - all changes require replacement.
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates. All changes require replacement.")
}

func (r *ParserVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ParserVersion
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest lakewatch.DeleteParserVersionRequest
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

	err := client.Lakewatch.DeleteParserVersion(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete lakewatch_parser_version", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ParserVersionResource{}

func (r *ParserVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
