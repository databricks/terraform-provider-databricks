// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_alert_config

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
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

const resourceName = "lakewatch_alert_config"

var _ resource.ResourceWithConfigure = &AlertConfigResource{}
var _ resource.ResourceWithModifyPlan = &AlertConfigResource{}

func ResourceAlertConfig() resource.Resource {
	return &AlertConfigResource{}
}

type AlertConfigResource struct {
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

// AlertConfig extends the main model with additional fields.
type AlertConfig struct {
	// Window for collating Signals into security alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1h`; the
	// window in effect is reported in
	// `effective_default_security_alert_collation_window`.
	DefaultSecurityAlertCollationWindow types.String `tfsdk:"default_security_alert_collation_window"`
	// Window for collating SystemEvents into system alerts. Duration string —
	// compact day/hour/minute/second components, e.g. `1d`, `12h` (see
	// `Schedule.at_least_every` in common.proto for the grammar). Reads return
	// the value as supplied. When unset or empty, the window is `1d`; the
	// window in effect is reported in
	// `effective_default_system_alert_collation_window`.
	DefaultSystemAlertCollationWindow types.String `tfsdk:"default_system_alert_collation_window"`
	// The security alert collation window in effect:
	// `default_security_alert_collation_window` when set, otherwise `1h`.
	EffectiveDefaultSecurityAlertCollationWindow types.String `tfsdk:"effective_default_security_alert_collation_window"`
	// The system alert collation window in effect:
	// `default_system_alert_collation_window` when set, otherwise `1d`.
	EffectiveDefaultSystemAlertCollationWindow types.String `tfsdk:"effective_default_system_alert_collation_window"`
	// Opaque entity tag for optimistic concurrency. Returned on every read;
	// pass it back as the Update request's `etag` to make the update
	// conditional on the resource being unchanged.
	Etag types.String `tfsdk:"etag"`

	Name types.String `tfsdk:"name"`
	// The resource name of the parent.
	Parent types.String `tfsdk:"parent"`
	// Export destination for security alerts. When unset, export is off and
	// reads return `{"enabled": false}`.
	SecurityAlertExportConfig types.Object `tfsdk:"security_alert_export_config"`
	// Export destination for system alerts. When unset, export is off and reads
	// return `{"enabled": false}`.
	SystemAlertExportConfig types.Object `tfsdk:"system_alert_export_config"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.Object `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version        types.Int64  `tfsdk:"version"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AlertConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AlertConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"security_alert_export_config": reflect.TypeOf(lakewatch_tf.ExportConfig{}),
		"system_alert_export_config":   reflect.TypeOf(lakewatch_tf.ExportConfig{}),
		"updated_by":                   reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config":              reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConfig
// only implements ToObjectValue() and Type().
func (m AlertConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"default_security_alert_collation_window": m.DefaultSecurityAlertCollationWindow,
			"default_system_alert_collation_window":             m.DefaultSystemAlertCollationWindow,
			"effective_default_security_alert_collation_window": m.EffectiveDefaultSecurityAlertCollationWindow,
			"effective_default_system_alert_collation_window":   m.EffectiveDefaultSystemAlertCollationWindow,
			"etag":                         m.Etag,
			"name":                         m.Name,
			"parent":                       m.Parent,
			"security_alert_export_config": m.SecurityAlertExportConfig,
			"system_alert_export_config":   m.SystemAlertExportConfig,
			"update_time":                  m.UpdateTime,
			"updated_by":                   m.UpdatedBy,
			"version":                      m.Version,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AlertConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"default_security_alert_collation_window": types.StringType,
			"default_system_alert_collation_window":             types.StringType,
			"effective_default_security_alert_collation_window": types.StringType,
			"effective_default_system_alert_collation_window":   types.StringType,
			"etag":                         types.StringType,
			"name":                         types.StringType,
			"parent":                       types.StringType,
			"security_alert_export_config": lakewatch_tf.ExportConfig{}.Type(ctx),
			"system_alert_export_config":   lakewatch_tf.ExportConfig{}.Type(ctx),
			"update_time":                  timetypes.RFC3339{}.Type(ctx),
			"updated_by":                   lakewatch_tf.Principal{}.Type(ctx),
			"version":                      types.Int64Type,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *AlertConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConfig) {
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.SecurityAlertExportConfig.IsNull() && !from.SecurityAlertExportConfig.IsUnknown() {
		if toSecurityAlertExportConfig, ok := to.GetSecurityAlertExportConfig(ctx); ok {
			if fromSecurityAlertExportConfig, ok := from.GetSecurityAlertExportConfig(ctx); ok {
				// Recursively sync the fields of SecurityAlertExportConfig
				toSecurityAlertExportConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromSecurityAlertExportConfig)
				to.SetSecurityAlertExportConfig(ctx, toSecurityAlertExportConfig)
			}
		}
	}
	if !from.SystemAlertExportConfig.IsNull() && !from.SystemAlertExportConfig.IsUnknown() {
		if toSystemAlertExportConfig, ok := to.GetSystemAlertExportConfig(ctx); ok {
			if fromSystemAlertExportConfig, ok := from.GetSystemAlertExportConfig(ctx); ok {
				// Recursively sync the fields of SystemAlertExportConfig
				toSystemAlertExportConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromSystemAlertExportConfig)
				to.SetSystemAlertExportConfig(ctx, toSystemAlertExportConfig)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AlertConfig) SyncFieldsDuringRead(ctx context.Context, from AlertConfig) {
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.SecurityAlertExportConfig.IsNull() && !from.SecurityAlertExportConfig.IsUnknown() {
		if toSecurityAlertExportConfig, ok := to.GetSecurityAlertExportConfig(ctx); ok {
			if fromSecurityAlertExportConfig, ok := from.GetSecurityAlertExportConfig(ctx); ok {
				toSecurityAlertExportConfig.SyncFieldsDuringRead(ctx, fromSecurityAlertExportConfig)
				to.SetSecurityAlertExportConfig(ctx, toSecurityAlertExportConfig)
			}
		}
	}
	if !from.SystemAlertExportConfig.IsNull() && !from.SystemAlertExportConfig.IsUnknown() {
		if toSystemAlertExportConfig, ok := to.GetSystemAlertExportConfig(ctx); ok {
			if fromSystemAlertExportConfig, ok := from.GetSystemAlertExportConfig(ctx); ok {
				toSystemAlertExportConfig.SyncFieldsDuringRead(ctx, fromSystemAlertExportConfig)
				to.SetSystemAlertExportConfig(ctx, toSystemAlertExportConfig)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

	to.Parent = types.StringValue(strings.TrimSuffix(to.Name.ValueString(), "/alertConfig"))
}

func (m AlertConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_security_alert_collation_window"] = attrs["default_security_alert_collation_window"].SetOptional()
	attrs["default_system_alert_collation_window"] = attrs["default_system_alert_collation_window"].SetOptional()
	attrs["effective_default_security_alert_collation_window"] = attrs["effective_default_security_alert_collation_window"].SetComputed()
	attrs["effective_default_system_alert_collation_window"] = attrs["effective_default_system_alert_collation_window"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["security_alert_export_config"] = attrs["security_alert_export_config"].SetOptional()
	attrs["system_alert_export_config"] = attrs["system_alert_export_config"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetSecurityAlertExportConfig returns the value of the SecurityAlertExportConfig field in AlertConfig as
// a lakewatch_tf.ExportConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig) GetSecurityAlertExportConfig(ctx context.Context) (lakewatch_tf.ExportConfig, bool) {
	var e lakewatch_tf.ExportConfig
	if m.SecurityAlertExportConfig.IsNull() || m.SecurityAlertExportConfig.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.ExportConfig
	d := m.SecurityAlertExportConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityAlertExportConfig sets the value of the SecurityAlertExportConfig field in AlertConfig.
func (m *AlertConfig) SetSecurityAlertExportConfig(ctx context.Context, v lakewatch_tf.ExportConfig) {
	vs := v.ToObjectValue(ctx)
	m.SecurityAlertExportConfig = vs
}

// GetSystemAlertExportConfig returns the value of the SystemAlertExportConfig field in AlertConfig as
// a lakewatch_tf.ExportConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig) GetSystemAlertExportConfig(ctx context.Context) (lakewatch_tf.ExportConfig, bool) {
	var e lakewatch_tf.ExportConfig
	if m.SystemAlertExportConfig.IsNull() || m.SystemAlertExportConfig.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.ExportConfig
	d := m.SystemAlertExportConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSystemAlertExportConfig sets the value of the SystemAlertExportConfig field in AlertConfig.
func (m *AlertConfig) SetSystemAlertExportConfig(ctx context.Context, v lakewatch_tf.ExportConfig) {
	vs := v.ToObjectValue(ctx)
	m.SystemAlertExportConfig = vs
}

// GetUpdatedBy returns the value of the UpdatedBy field in AlertConfig as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertConfig) GetUpdatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.UpdatedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdatedBy sets the value of the UpdatedBy field in AlertConfig.
func (m *AlertConfig) SetUpdatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.UpdatedBy = vs
}

func (r *AlertConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AlertConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, AlertConfig{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks lakewatch_alert_config",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *AlertConfigResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *AlertConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AlertConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *AlertConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState AlertConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetAlertConfigRequest
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
	response, err := client.Lakewatch.GetAlertConfig(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get lakewatch_alert_config", err.Error())
		return
	}

	var newState AlertConfig
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

func (r *AlertConfigResource) update(ctx context.Context, plan AlertConfig, diags *diag.Diagnostics, state *tfsdk.State) {
	// AIP-156 singletons have no ID, so derive the computed name from the parent and static path segment.
	plan.Name = types.StringValue(plan.Parent.ValueString() + "/alertConfig")
	var alert_config lakewatch.AlertConfig

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &alert_config)...)
	if diags.HasError() {
		return
	}

	updateRequest := lakewatch.UpdateAlertConfigRequest{
		AlertConfig: alert_config,
		Name:        plan.Name.ValueString(),
		UpdateMask:  *fieldmask.New(strings.Split("default_security_alert_collation_window,default_system_alert_collation_window,security_alert_export_config,system_alert_export_config", ",")),
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
	response, err := client.Lakewatch.UpdateAlertConfig(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update lakewatch_alert_config", err.Error())
		return
	}

	var newState AlertConfig

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AlertConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AlertConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AlertConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &AlertConfigResource{}

func (r *AlertConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
