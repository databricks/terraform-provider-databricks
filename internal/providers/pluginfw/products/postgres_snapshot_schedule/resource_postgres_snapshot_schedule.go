// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_snapshot_schedule

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/postgres_tf"
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

const resourceName = "postgres_snapshot_schedule"

var _ resource.ResourceWithConfigure = &SnapshotScheduleResource{}
var _ resource.ResourceWithModifyPlan = &SnapshotScheduleResource{}

func ResourceSnapshotSchedule() resource.Resource {
	return &SnapshotScheduleResource{}
}

type SnapshotScheduleResource struct {
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

// SnapshotSchedule extends the main model with additional fields.
type SnapshotSchedule struct {
	// The resource name of the branch's snapshot schedule. Format:
	// projects/{project_id}/branches/{branch_id}/snapshot-schedule
	Name types.String `tfsdk:"name"`
	// The cadences at which automatic snapshots are taken. Update replaces the
	// whole set; an empty set disables automatic snapshots. Order is not
	// significant. When several cadences fire together, one snapshot is taken,
	// retained for the longest of their retentions.
	Schedule       types.Set    `tfsdk:"schedule"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// SnapshotSchedule struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m SnapshotSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule":        reflect.TypeOf(postgres_tf.ScheduleCadence{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SnapshotSchedule
// only implements ToObjectValue() and Type().
func (m SnapshotSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"name": m.Name,
			"schedule": m.Schedule,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m SnapshotSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"name": types.StringType,
			"schedule": basetypes.SetType{
				ElemType: postgres_tf.ScheduleCadence{}.Type(ctx),
			},

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *SnapshotSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SnapshotSchedule) {
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() && to.Schedule.IsNull() && len(from.Schedule.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Schedule, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Schedule = from.Schedule
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *SnapshotSchedule) SyncFieldsDuringRead(ctx context.Context, from SnapshotSchedule) {
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() && to.Schedule.IsNull() && len(from.Schedule.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Schedule, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Schedule = from.Schedule
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m SnapshotSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetComputed()
	attrs["schedule"] = attrs["schedule"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetSchedule returns the value of the Schedule field in SnapshotSchedule as
// a slice of postgres_tf.ScheduleCadence values.
// If the field is unknown or null, the boolean return value is false.
func (m *SnapshotSchedule) GetSchedule(ctx context.Context) ([]postgres_tf.ScheduleCadence, bool) {
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return nil, false
	}
	var v []postgres_tf.ScheduleCadence
	d := m.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in SnapshotSchedule.
func (m *SnapshotSchedule) SetSchedule(ctx context.Context, v []postgres_tf.ScheduleCadence) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Schedule = types.SetValueMust(t, vs)
}

func (r *SnapshotScheduleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SnapshotScheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, SnapshotSchedule{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks postgres_snapshot_schedule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SnapshotScheduleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *SnapshotScheduleResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *SnapshotScheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan SnapshotSchedule
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

func (r *SnapshotScheduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState SnapshotSchedule
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetSnapshotScheduleRequest
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
	response, err := client.Postgres.GetSnapshotSchedule(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get postgres_snapshot_schedule", err.Error())
		return
	}

	var newState SnapshotSchedule
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

func (r *SnapshotScheduleResource) update(ctx context.Context, plan SnapshotSchedule, diags *diag.Diagnostics, state *tfsdk.State) {
	var snapshot_schedule postgres.SnapshotSchedule

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &snapshot_schedule)...)
	if diags.HasError() {
		return
	}

	updateRequest := postgres.UpdateSnapshotScheduleRequest{
		SnapshotSchedule: snapshot_schedule,
		Name:             plan.Name.ValueString(),
		UpdateMask:       *fieldmask.New(strings.Split("schedule", ",")),
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
	response, err := client.Postgres.UpdateSnapshotSchedule(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update postgres_snapshot_schedule", err.Error())
		return
	}

	var newState SnapshotSchedule

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for postgres_snapshot_schedule update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *SnapshotScheduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan SnapshotSchedule
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SnapshotScheduleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &SnapshotScheduleResource{}

func (r *SnapshotScheduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
