// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_endpoint

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/postgres_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "postgres_endpoint"

var _ datasource.DataSourceWithConfigure = &EndpointDataSource{}

func DataSourceEndpoint() datasource.DataSource {
	return &EndpointDataSource{}
}

type EndpointDataSource struct {
	Client *autogen.DatabricksClient
}

// EndpointData extends the main model with additional fields.
type EndpointData struct {
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
// EndpointData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m EndpointData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(postgres_tf.EndpointSettings{}),
		"settings":           reflect.TypeOf(postgres_tf.EndpointSettings{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointData
// only implements ToObjectValue() and Type().
func (m EndpointData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu":           m.AutoscalingLimitMaxCu,
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
func (m EndpointData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu":           types.Float64Type,
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

func (m EndpointData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetComputed()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetComputed()
	attrs["effective_autoscaling_limit_max_cu"] = attrs["effective_autoscaling_limit_max_cu"].SetComputed()
	attrs["effective_autoscaling_limit_min_cu"] = attrs["effective_autoscaling_limit_min_cu"].SetComputed()
	attrs["effective_disabled"] = attrs["effective_disabled"].SetComputed()
	attrs["effective_pooler_mode"] = attrs["effective_pooler_mode"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["effective_suspend_timeout_duration"] = attrs["effective_suspend_timeout_duration"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetComputed()
	attrs["host"] = attrs["host"].SetComputed()
	attrs["last_active_time"] = attrs["last_active_time"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetComputed()
	attrs["settings"] = attrs["settings"].SetComputed()
	attrs["start_time"] = attrs["start_time"].SetComputed()
	attrs["suspend_time"] = attrs["suspend_time"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

func (r *EndpointDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *EndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EndpointData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Endpoint",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EndpointDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *EndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config EndpointData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetEndpointRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
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

	var newState EndpointData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
