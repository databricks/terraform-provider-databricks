// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_project

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

const dataSourceName = "postgres_project"

var _ datasource.DataSourceWithConfigure = &ProjectDataSource{}

func DataSourceProject() datasource.DataSource {
	return &ProjectDataSource{}
}

type ProjectDataSource struct {
	Client *autogen.DatabricksClient
}

// ProjectData extends the main model with additional fields.
type ProjectData struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime timetypes.RFC3339 `tfsdk:"compute_last_active_time"`
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	DefaultEndpointSettings types.Object `tfsdk:"default_endpoint_settings"`
	// Human-readable project name.
	DisplayName types.String `tfsdk:"display_name"`

	EffectiveDefaultEndpointSettings types.Object `tfsdk:"effective_default_endpoint_settings"`

	EffectiveDisplayName types.String `tfsdk:"effective_display_name"`

	EffectiveHistoryRetentionDuration timetypes.GoDuration `tfsdk:"effective_history_retention_duration"`

	EffectivePgVersion types.Int64 `tfsdk:"effective_pg_version"`

	EffectiveSettings types.Object `tfsdk:"effective_settings"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration timetypes.GoDuration `tfsdk:"history_retention_duration"`
	// The resource name of the project. Format: projects/{project_id}
	Name types.String `tfsdk:"name"`
	// The major Postgres version number.
	PgVersion types.Int64 `tfsdk:"pg_version"`

	Settings types.Object `tfsdk:"settings"`
	// The current space occupied by the project in storage. Synthetic storage
	// size combines the logical data size and Write-Ahead Log (WAL) size for
	// all branches in a project.
	SyntheticStorageSizeBytes types.Int64 `tfsdk:"synthetic_storage_size_bytes"`
	// System generated unique ID for the project.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the project was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProjectData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ProjectData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings":           reflect.TypeOf(postgres_tf.ProjectDefaultEndpointSettings{}),
		"effective_default_endpoint_settings": reflect.TypeOf(postgres_tf.ProjectDefaultEndpointSettings{}),
		"effective_settings":                  reflect.TypeOf(postgres_tf.ProjectSettings{}),
		"settings":                            reflect.TypeOf(postgres_tf.ProjectSettings{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectData
// only implements ToObjectValue() and Type().
func (m ProjectData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_logical_size_limit_bytes":      m.BranchLogicalSizeLimitBytes,
			"compute_last_active_time":             m.ComputeLastActiveTime,
			"create_time":                          m.CreateTime,
			"default_endpoint_settings":            m.DefaultEndpointSettings,
			"display_name":                         m.DisplayName,
			"effective_default_endpoint_settings":  m.EffectiveDefaultEndpointSettings,
			"effective_display_name":               m.EffectiveDisplayName,
			"effective_history_retention_duration": m.EffectiveHistoryRetentionDuration,
			"effective_pg_version":                 m.EffectivePgVersion,
			"effective_settings":                   m.EffectiveSettings,
			"history_retention_duration":           m.HistoryRetentionDuration,
			"name":                                 m.Name,
			"pg_version":                           m.PgVersion,
			"settings":                             m.Settings,
			"synthetic_storage_size_bytes":         m.SyntheticStorageSizeBytes,
			"uid":                                  m.Uid,
			"update_time":                          m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ProjectData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes":      types.Int64Type,
			"compute_last_active_time":             timetypes.RFC3339{}.Type(ctx),
			"create_time":                          timetypes.RFC3339{}.Type(ctx),
			"default_endpoint_settings":            postgres_tf.ProjectDefaultEndpointSettings{}.Type(ctx),
			"display_name":                         types.StringType,
			"effective_default_endpoint_settings":  postgres_tf.ProjectDefaultEndpointSettings{}.Type(ctx),
			"effective_display_name":               types.StringType,
			"effective_history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"effective_pg_version":                 types.Int64Type,
			"effective_settings":                   postgres_tf.ProjectSettings{}.Type(ctx),
			"history_retention_duration":           timetypes.GoDuration{}.Type(ctx),
			"name":                                 types.StringType,
			"pg_version":                           types.Int64Type,
			"settings":                             postgres_tf.ProjectSettings{}.Type(ctx),
			"synthetic_storage_size_bytes":         types.Int64Type,
			"uid":                                  types.StringType,
			"update_time":                          timetypes.RFC3339{}.Type(ctx),
		},
	}
}

func (m ProjectData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["compute_last_active_time"] = attrs["compute_last_active_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["effective_default_endpoint_settings"] = attrs["effective_default_endpoint_settings"].SetComputed()
	attrs["effective_display_name"] = attrs["effective_display_name"].SetComputed()
	attrs["effective_history_retention_duration"] = attrs["effective_history_retention_duration"].SetComputed()
	attrs["effective_pg_version"] = attrs["effective_pg_version"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["settings"] = attrs["settings"].SetComputed()
	attrs["synthetic_storage_size_bytes"] = attrs["synthetic_storage_size_bytes"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

func (r *ProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ProjectData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Project",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ProjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ProjectData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetProjectRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.GetProject(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get postgres_project", err.Error())
		return
	}

	var newState ProjectData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
