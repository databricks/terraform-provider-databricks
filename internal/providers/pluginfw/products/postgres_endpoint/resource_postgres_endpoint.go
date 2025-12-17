// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_endpoint

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
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "postgres_endpoint"

var _ resource.ResourceWithConfigure = &EndpointResource{}

func ResourceEndpoint() resource.Resource {
	return &EndpointResource{}
}

type EndpointResource struct {
	Client *autogen.DatabricksClient
}

// Endpoint extends the main model with additional fields.
type Endpoint struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// A timestamp indicating when the compute endpoint was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	CurrentState types.String `tfsdk:"current_state"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled types.Bool `tfsdk:"disabled"`
	// The maximum number of Compute Units.
	EffectiveAutoscalingLimitMaxCu types.Float64 `tfsdk:"effective_autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	EffectiveAutoscalingLimitMinCu types.Float64 `tfsdk:"effective_autoscaling_limit_min_cu"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	EffectiveDisabled types.Bool `tfsdk:"effective_disabled"`

	EffectivePoolerMode types.String `tfsdk:"effective_pooler_mode"`

	EffectiveSettings types.Object `tfsdk:"effective_settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	EffectiveSuspendTimeoutDuration timetypes.GoDuration `tfsdk:"effective_suspend_timeout_duration"`
	// The ID to use for the Endpoint, which will become the final component of
	// the endpoint's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// The endpoint type. There could be only one READ_WRITE endpoint per
	// branch.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host types.String `tfsdk:"host"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime timetypes.RFC3339 `tfsdk:"last_active_time"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"name"`
	// The branch containing this endpoint. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"parent"`

	PendingState types.String `tfsdk:"pending_state"`

	PoolerMode types.String `tfsdk:"pooler_mode"`

	Settings types.Object `tfsdk:"settings"`
	// A timestamp indicating when the compute endpoint was last started.
	StartTime timetypes.RFC3339 `tfsdk:"start_time"`
	// A timestamp indicating when the compute endpoint was last suspended.
	SuspendTime timetypes.RFC3339 `tfsdk:"suspend_time"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
	// System generated unique ID for the endpoint.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Endpoint struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Endpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(postgres_tf.EndpointSettings{}),
		"settings":           reflect.TypeOf(postgres_tf.EndpointSettings{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint
// only implements ToObjectValue() and Type().
func (m Endpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu":           m.AutoscalingLimitMinCu,
			"create_time":                        m.CreateTime,
			"current_state":                      m.CurrentState,
			"disabled":                           m.Disabled,
			"effective_autoscaling_limit_max_cu": m.EffectiveAutoscalingLimitMaxCu,
			"effective_autoscaling_limit_min_cu": m.EffectiveAutoscalingLimitMinCu,
			"effective_disabled":                 m.EffectiveDisabled,
			"effective_pooler_mode":              m.EffectivePoolerMode,
			"effective_settings":                 m.EffectiveSettings,
			"effective_suspend_timeout_duration": m.EffectiveSuspendTimeoutDuration,
			"endpoint_id":                        m.EndpointId,
			"endpoint_type":                      m.EndpointType,
			"host":                               m.Host,
			"last_active_time":                   m.LastActiveTime,
			"name":                               m.Name,
			"parent":                             m.Parent,
			"pending_state":                      m.PendingState,
			"pooler_mode":                        m.PoolerMode,
			"settings":                           m.Settings,
			"start_time":                         m.StartTime,
			"suspend_time":                       m.SuspendTime,
			"suspend_timeout_duration":           m.SuspendTimeoutDuration,
			"uid":                                m.Uid,
			"update_time":                        m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Endpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu":           types.Float64Type,
			"create_time":                        timetypes.RFC3339{}.Type(ctx),
			"current_state":                      types.StringType,
			"disabled":                           types.BoolType,
			"effective_autoscaling_limit_max_cu": types.Float64Type,
			"effective_autoscaling_limit_min_cu": types.Float64Type,
			"effective_disabled":                 types.BoolType,
			"effective_pooler_mode":              types.StringType,
			"effective_settings":                 postgres_tf.EndpointSettings{}.Type(ctx),
			"effective_suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
			"endpoint_id":                        types.StringType,
			"endpoint_type":                      types.StringType,
			"host":                               types.StringType,
			"last_active_time":                   timetypes.RFC3339{}.Type(ctx),
			"name":                               types.StringType,
			"parent":                             types.StringType,
			"pending_state":                      types.StringType,
			"pooler_mode":                        types.StringType,
			"settings":                           postgres_tf.EndpointSettings{}.Type(ctx),
			"start_time":                         timetypes.RFC3339{}.Type(ctx),
			"suspend_time":                       timetypes.RFC3339{}.Type(ctx),
			"suspend_timeout_duration":           timetypes.GoDuration{}.Type(ctx),
			"uid":                                types.StringType,
			"update_time":                        timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Endpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint) {
	if !from.AutoscalingLimitMaxCu.IsUnknown() && !from.AutoscalingLimitMaxCu.IsNull() {
		// AutoscalingLimitMaxCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMaxCu = from.AutoscalingLimitMaxCu
	}
	if !from.AutoscalingLimitMinCu.IsUnknown() && !from.AutoscalingLimitMinCu.IsNull() {
		// AutoscalingLimitMinCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMinCu = from.AutoscalingLimitMinCu
	}
	if !from.Disabled.IsUnknown() && !from.Disabled.IsNull() {
		// Disabled is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Disabled = from.Disabled
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				// Recursively sync the fields of EffectiveSettings
				toEffectiveSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	to.EndpointId = from.EndpointId
	if !from.PoolerMode.IsUnknown() && !from.PoolerMode.IsNull() {
		// PoolerMode is an input only field and not returned by the service, so we keep the value from the prior state.
		to.PoolerMode = from.PoolerMode
	}
	if !from.Settings.IsUnknown() && !from.Settings.IsNull() {
		// Settings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Settings = from.Settings
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
	if !from.SuspendTimeoutDuration.IsUnknown() && !from.SuspendTimeoutDuration.IsNull() {
		// SuspendTimeoutDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SuspendTimeoutDuration = from.SuspendTimeoutDuration
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
	if !from.AutoscalingLimitMaxCu.IsUnknown() && !from.AutoscalingLimitMaxCu.IsNull() {
		// AutoscalingLimitMaxCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMaxCu = from.AutoscalingLimitMaxCu
	}
	if !from.AutoscalingLimitMinCu.IsUnknown() && !from.AutoscalingLimitMinCu.IsNull() {
		// AutoscalingLimitMinCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMinCu = from.AutoscalingLimitMinCu
	}
	if !from.Disabled.IsUnknown() && !from.Disabled.IsNull() {
		// Disabled is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Disabled = from.Disabled
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	to.EndpointId = from.EndpointId
	if !from.PoolerMode.IsUnknown() && !from.PoolerMode.IsNull() {
		// PoolerMode is an input only field and not returned by the service, so we keep the value from the prior state.
		to.PoolerMode = from.PoolerMode
	}
	if !from.Settings.IsUnknown() && !from.Settings.IsNull() {
		// Settings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Settings = from.Settings
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
	if !from.SuspendTimeoutDuration.IsUnknown() && !from.SuspendTimeoutDuration.IsNull() {
		// SuspendTimeoutDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SuspendTimeoutDuration = from.SuspendTimeoutDuration
	}
}

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetComputed()
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].(tfschema.Float64AttributeBuilder).AddPlanModifier(float64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetComputed()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].(tfschema.Float64AttributeBuilder).AddPlanModifier(float64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetOptional()
	attrs["disabled"] = attrs["disabled"].SetComputed()
	attrs["disabled"] = attrs["disabled"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_autoscaling_limit_max_cu"] = attrs["effective_autoscaling_limit_max_cu"].SetComputed()
	attrs["effective_autoscaling_limit_min_cu"] = attrs["effective_autoscaling_limit_min_cu"].SetComputed()
	attrs["effective_disabled"] = attrs["effective_disabled"].SetComputed()
	attrs["effective_pooler_mode"] = attrs["effective_pooler_mode"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["effective_suspend_timeout_duration"] = attrs["effective_suspend_timeout_duration"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["host"] = attrs["host"].SetComputed()
	attrs["last_active_time"] = attrs["last_active_time"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetOptional()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["settings"] = attrs["settings"].SetComputed()
	attrs["settings"] = attrs["settings"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["start_time"] = attrs["start_time"].SetComputed()
	attrs["suspend_time"] = attrs["suspend_time"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetComputed()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Endpoint as
// a postgres_tf.EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetEffectiveSettings(ctx context.Context) (postgres_tf.EndpointSettings, bool) {
	var e postgres_tf.EndpointSettings
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.EndpointSettings
	d := m.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Endpoint.
func (m *Endpoint) SetEffectiveSettings(ctx context.Context, v postgres_tf.EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveSettings = vs
}

// GetSettings returns the value of the Settings field in Endpoint as
// a postgres_tf.EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetSettings(ctx context.Context) (postgres_tf.EndpointSettings, bool) {
	var e postgres_tf.EndpointSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.EndpointSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in Endpoint.
func (m *Endpoint) SetSettings(ctx context.Context, v postgres_tf.EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

func (r *EndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *EndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Endpoint{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks postgres_endpoint",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *EndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Endpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var endpoint postgres.Endpoint

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &endpoint)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := postgres.CreateEndpointRequest{
		Endpoint:   endpoint,
		Parent:     plan.Parent.ValueString(),
		EndpointId: plan.EndpointId.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.CreateEndpoint(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create postgres_endpoint", err.Error())
		return
	}

	var newState Endpoint

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for postgres_endpoint to be ready", err.Error())
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

func (r *EndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Endpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetEndpointRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Postgres.GetEndpoint(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get postgres_endpoint", err.Error())
		return
	}

	var newState Endpoint
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *EndpointResource) update(ctx context.Context, plan Endpoint, diags *diag.Diagnostics, state *tfsdk.State) {
	var endpoint postgres.Endpoint

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &endpoint)...)
	if diags.HasError() {
		return
	}

	updateRequest := postgres.UpdateEndpointRequest{
		Endpoint:   endpoint,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("autoscaling_limit_max_cu,autoscaling_limit_min_cu,disabled,pooler_mode,settings,suspend_timeout_duration", ",")),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Postgres.UpdateEndpoint(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update postgres_endpoint", err.Error())
		return
	}

	var newState Endpoint

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for postgres_endpoint update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *EndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Endpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *EndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Endpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest postgres.DeleteEndpointRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Postgres.DeleteEndpoint(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete postgres_endpoint", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &EndpointResource{}

func (r *EndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
