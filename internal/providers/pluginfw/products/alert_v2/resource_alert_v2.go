// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package alert_v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "alert_v2"

var _ resource.ResourceWithConfigure = &AlertV2Resource{}

func ResourceAlertV2() resource.Resource {
	return &AlertV2Resource{}
}

type AlertV2Resource struct {
	Client *autogen.DatabricksClient
}

// AlertV2 extends the main model with additional fields.
type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom description for the alert. support mustache template.
	CustomDescription types.String `tfsdk:"custom_description"`
	// Custom summary for the alert. support mustache template.
	CustomSummary types.String `tfsdk:"custom_summary"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// The actual identity that will be used to execute the alert. This is an
	// output-only field that shows the resolved run-as identity after applying
	// permissions and defaults.
	EffectiveRunAs types.Object `tfsdk:"effective_run_as"`

	Evaluation types.Object `tfsdk:"evaluation"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	ParentPath types.String `tfsdk:"parent_path"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// Specifies the identity that will be used to run the alert. This field
	// allows you to configure alerts to run as a specific user or service
	// principal. - For user identity: Set `user_name` to the email of an active
	// workspace user. Users can only set this to their own email. - For service
	// principal: Set `service_principal_name` to the application ID. Requires
	// the `servicePrincipal/user` role. If not specified, the alert will run as
	// the request user.
	RunAs types.Object `tfsdk:"run_as"`
	// The run as username or application ID of service principal. On Create and
	// Update, this field can be set to application ID of an active service
	// principal. Setting this field requires the servicePrincipal/user role.
	// Deprecated: Use `run_as` field instead. This field will be removed in a
	// future release.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`

	Schedule types.Object `tfsdk:"schedule"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AlertV2 struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AlertV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_run_as": reflect.TypeOf(sql_tf.AlertV2RunAs{}),
		"evaluation":       reflect.TypeOf(sql_tf.AlertV2Evaluation{}),
		"run_as":           reflect.TypeOf(sql_tf.AlertV2RunAs{}),
		"schedule":         reflect.TypeOf(sql_tf.CronSchedule{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2
// only implements ToObjectValue() and Type().
func (m AlertV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"custom_description": m.CustomDescription,
			"custom_summary":     m.CustomSummary,
			"display_name":       m.DisplayName,
			"effective_run_as":   m.EffectiveRunAs,
			"evaluation":         m.Evaluation,
			"id":                 m.Id,
			"lifecycle_state":    m.LifecycleState,
			"owner_user_name":    m.OwnerUserName,
			"parent_path":        m.ParentPath,
			"query_text":         m.QueryText,
			"run_as":             m.RunAs,
			"run_as_user_name":   m.RunAsUserName,
			"schedule":           m.Schedule,
			"update_time":        m.UpdateTime,
			"warehouse_id":       m.WarehouseId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AlertV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": types.StringType,
			"custom_description": types.StringType,
			"custom_summary":     types.StringType,
			"display_name":       types.StringType,
			"effective_run_as":   sql_tf.AlertV2RunAs{}.Type(ctx),
			"evaluation":         sql_tf.AlertV2Evaluation{}.Type(ctx),
			"id":                 types.StringType,
			"lifecycle_state":    types.StringType,
			"owner_user_name":    types.StringType,
			"parent_path":        types.StringType,
			"query_text":         types.StringType,
			"run_as":             sql_tf.AlertV2RunAs{}.Type(ctx),
			"run_as_user_name":   types.StringType,
			"schedule":           sql_tf.CronSchedule{}.Type(ctx),
			"update_time":        types.StringType,
			"warehouse_id":       types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *AlertV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2) {
	if !from.EffectiveRunAs.IsNull() && !from.EffectiveRunAs.IsUnknown() {
		if toEffectiveRunAs, ok := to.GetEffectiveRunAs(ctx); ok {
			if fromEffectiveRunAs, ok := from.GetEffectiveRunAs(ctx); ok {
				// Recursively sync the fields of EffectiveRunAs
				toEffectiveRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveRunAs)
				to.SetEffectiveRunAs(ctx, toEffectiveRunAs)
			}
		}
	}
	if !from.Evaluation.IsNull() && !from.Evaluation.IsUnknown() {
		if toEvaluation, ok := to.GetEvaluation(ctx); ok {
			if fromEvaluation, ok := from.GetEvaluation(ctx); ok {
				// Recursively sync the fields of Evaluation
				toEvaluation.SyncFieldsDuringCreateOrUpdate(ctx, fromEvaluation)
				to.SetEvaluation(ctx, toEvaluation)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				// Recursively sync the fields of RunAs
				toRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				// Recursively sync the fields of Schedule
				toSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AlertV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2) {
	if !from.EffectiveRunAs.IsNull() && !from.EffectiveRunAs.IsUnknown() {
		if toEffectiveRunAs, ok := to.GetEffectiveRunAs(ctx); ok {
			if fromEffectiveRunAs, ok := from.GetEffectiveRunAs(ctx); ok {
				toEffectiveRunAs.SyncFieldsDuringRead(ctx, fromEffectiveRunAs)
				to.SetEffectiveRunAs(ctx, toEffectiveRunAs)
			}
		}
	}
	if !from.Evaluation.IsNull() && !from.Evaluation.IsUnknown() {
		if toEvaluation, ok := to.GetEvaluation(ctx); ok {
			if fromEvaluation, ok := from.GetEvaluation(ctx); ok {
				toEvaluation.SyncFieldsDuringRead(ctx, fromEvaluation)
				to.SetEvaluation(ctx, toEvaluation)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				toRunAs.SyncFieldsDuringRead(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				toSchedule.SyncFieldsDuringRead(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
}

func (m AlertV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["custom_description"] = attrs["custom_description"].SetOptional()
	attrs["custom_summary"] = attrs["custom_summary"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["effective_run_as"] = attrs["effective_run_as"].SetComputed()
	attrs["evaluation"] = attrs["evaluation"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	attrs["id"] = attrs["id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetEffectiveRunAs returns the value of the EffectiveRunAs field in AlertV2 as
// a sql_tf.AlertV2RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2) GetEffectiveRunAs(ctx context.Context) (sql_tf.AlertV2RunAs, bool) {
	var e sql_tf.AlertV2RunAs
	if m.EffectiveRunAs.IsNull() || m.EffectiveRunAs.IsUnknown() {
		return e, false
	}
	var v sql_tf.AlertV2RunAs
	d := m.EffectiveRunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveRunAs sets the value of the EffectiveRunAs field in AlertV2.
func (m *AlertV2) SetEffectiveRunAs(ctx context.Context, v sql_tf.AlertV2RunAs) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveRunAs = vs
}

// GetEvaluation returns the value of the Evaluation field in AlertV2 as
// a sql_tf.AlertV2Evaluation value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2) GetEvaluation(ctx context.Context) (sql_tf.AlertV2Evaluation, bool) {
	var e sql_tf.AlertV2Evaluation
	if m.Evaluation.IsNull() || m.Evaluation.IsUnknown() {
		return e, false
	}
	var v sql_tf.AlertV2Evaluation
	d := m.Evaluation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvaluation sets the value of the Evaluation field in AlertV2.
func (m *AlertV2) SetEvaluation(ctx context.Context, v sql_tf.AlertV2Evaluation) {
	vs := v.ToObjectValue(ctx)
	m.Evaluation = vs
}

// GetRunAs returns the value of the RunAs field in AlertV2 as
// a sql_tf.AlertV2RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2) GetRunAs(ctx context.Context) (sql_tf.AlertV2RunAs, bool) {
	var e sql_tf.AlertV2RunAs
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v sql_tf.AlertV2RunAs
	d := m.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in AlertV2.
func (m *AlertV2) SetRunAs(ctx context.Context, v sql_tf.AlertV2RunAs) {
	vs := v.ToObjectValue(ctx)
	m.RunAs = vs
}

// GetSchedule returns the value of the Schedule field in AlertV2 as
// a sql_tf.CronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2) GetSchedule(ctx context.Context) (sql_tf.CronSchedule, bool) {
	var e sql_tf.CronSchedule
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v sql_tf.CronSchedule
	d := m.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in AlertV2.
func (m *AlertV2) SetSchedule(ctx context.Context, v sql_tf.CronSchedule) {
	vs := v.ToObjectValue(ctx)
	m.Schedule = vs
}

func (r *AlertV2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AlertV2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, AlertV2{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks alert_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertV2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *AlertV2Resource) update(ctx context.Context, plan AlertV2, diags *diag.Diagnostics, state *tfsdk.State) {
	var alert_v2 sql.AlertV2

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &alert_v2)...)
	if diags.HasError() {
		return
	}

	updateRequest := sql.UpdateAlertV2Request{
		Alert:      alert_v2,
		Id:         plan.Id.ValueString(),
		UpdateMask: "custom_description,custom_summary,display_name,evaluation,parent_path,query_text,run_as,run_as_user_name,schedule,warehouse_id",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.AlertsV2.UpdateAlert(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update alert_v2", err.Error())
		return
	}

	var newState AlertV2
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AlertV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AlertV2
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var alert_v2 sql.AlertV2

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &alert_v2)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := sql.CreateAlertV2Request{
		Alert: alert_v2,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.AlertsV2.CreateAlert(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create alert_v2", err.Error())
		return
	}

	var newState AlertV2

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

func (r *AlertV2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState AlertV2
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest sql.GetAlertV2Request
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.AlertsV2.GetAlert(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get alert_v2", err.Error())
		return
	}

	var newState AlertV2
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *AlertV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AlertV2
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AlertV2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state AlertV2
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest sql.TrashAlertV2Request
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := client.AlertsV2.TrashAlert(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete alert_v2", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &AlertV2Resource{}

func (r *AlertV2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: id. Got: %q",
				req.ID,
			),
		)
		return
	}

	id := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
