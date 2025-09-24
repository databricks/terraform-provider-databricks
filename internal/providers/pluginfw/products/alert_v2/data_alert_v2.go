// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package alert_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "alert_v2"

var _ datasource.DataSourceWithConfigure = &AlertV2DataSource{}

func DataSourceAlertV2() datasource.DataSource {
	return &AlertV2DataSource{}
}

type AlertV2DataSource struct {
	Client *autogen.DatabricksClient
}

// AlertV2Data extends the main model with additional fields.
type AlertV2Data struct {
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
// AlertV2Data struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AlertV2Data) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Data
// only implements ToObjectValue() and Type().
func (m AlertV2Data) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m AlertV2Data) Type(ctx context.Context) attr.Type {
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

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AlertV2Data) SyncFieldsDuringRead(ctx context.Context, from AlertV2Data) {
}

func (m AlertV2Data) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

	return attrs
}

func (r *AlertV2DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AlertV2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AlertV2Data{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AlertV2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertV2DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AlertV2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AlertV2Data
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest sql.GetAlertV2Request
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
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

	var newState AlertV2Data
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
