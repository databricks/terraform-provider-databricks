// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_project

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "postgres_project"

var _ resource.ResourceWithConfigure = &ProjectResource{}

func ResourceProject() resource.Resource {
	return &ProjectResource{}
}

type ProjectResource struct {
	Client *autogen.DatabricksClient
}

// Project extends the main model with additional fields.
type Project struct {
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
	// The ID to use for the Project, which will become the final component of
	// the project's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	ProjectId types.String `tfsdk:"project_id"`

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
// Project struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Project) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project
// only implements ToObjectValue() and Type().
func (m Project) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"branch_logical_size_limit_bytes": m.BranchLogicalSizeLimitBytes,
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
			"project_id":                           m.ProjectId,
			"settings":                             m.Settings,
			"synthetic_storage_size_bytes":         m.SyntheticStorageSizeBytes,
			"uid":                                  m.Uid,
			"update_time":                          m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Project) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"branch_logical_size_limit_bytes": types.Int64Type,
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
			"project_id":                           types.StringType,
			"settings":                             postgres_tf.ProjectSettings{}.Type(ctx),
			"synthetic_storage_size_bytes":         types.Int64Type,
			"uid":                                  types.StringType,
			"update_time":                          timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Project) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project) {
	if !from.DefaultEndpointSettings.IsUnknown() && !from.DefaultEndpointSettings.IsNull() {
		// DefaultEndpointSettings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DefaultEndpointSettings = from.DefaultEndpointSettings
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of DefaultEndpointSettings
				toDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.DisplayName.IsUnknown() && !from.DisplayName.IsNull() {
		// DisplayName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DisplayName = from.DisplayName
	}
	if !from.EffectiveDefaultEndpointSettings.IsNull() && !from.EffectiveDefaultEndpointSettings.IsUnknown() {
		if toEffectiveDefaultEndpointSettings, ok := to.GetEffectiveDefaultEndpointSettings(ctx); ok {
			if fromEffectiveDefaultEndpointSettings, ok := from.GetEffectiveDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of EffectiveDefaultEndpointSettings
				toEffectiveDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveDefaultEndpointSettings)
				to.SetEffectiveDefaultEndpointSettings(ctx, toEffectiveDefaultEndpointSettings)
			}
		}
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
	if !from.HistoryRetentionDuration.IsUnknown() && !from.HistoryRetentionDuration.IsNull() {
		// HistoryRetentionDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.HistoryRetentionDuration = from.HistoryRetentionDuration
	}
	to.ProjectId = from.ProjectId
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Project) SyncFieldsDuringRead(ctx context.Context, from Project) {
	if !from.DefaultEndpointSettings.IsUnknown() && !from.DefaultEndpointSettings.IsNull() {
		// DefaultEndpointSettings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DefaultEndpointSettings = from.DefaultEndpointSettings
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.DisplayName.IsUnknown() && !from.DisplayName.IsNull() {
		// DisplayName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DisplayName = from.DisplayName
	}
	if !from.EffectiveDefaultEndpointSettings.IsNull() && !from.EffectiveDefaultEndpointSettings.IsUnknown() {
		if toEffectiveDefaultEndpointSettings, ok := to.GetEffectiveDefaultEndpointSettings(ctx); ok {
			if fromEffectiveDefaultEndpointSettings, ok := from.GetEffectiveDefaultEndpointSettings(ctx); ok {
				toEffectiveDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromEffectiveDefaultEndpointSettings)
				to.SetEffectiveDefaultEndpointSettings(ctx, toEffectiveDefaultEndpointSettings)
			}
		}
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	if !from.HistoryRetentionDuration.IsUnknown() && !from.HistoryRetentionDuration.IsNull() {
		// HistoryRetentionDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.HistoryRetentionDuration = from.HistoryRetentionDuration
	}
	to.ProjectId = from.ProjectId
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m Project) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["compute_last_active_time"] = attrs["compute_last_active_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_default_endpoint_settings"] = attrs["effective_default_endpoint_settings"].SetComputed()
	attrs["effective_display_name"] = attrs["effective_display_name"].SetComputed()
	attrs["effective_history_retention_duration"] = attrs["effective_history_retention_duration"].SetComputed()
	attrs["effective_pg_version"] = attrs["effective_pg_version"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetOptional()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["pg_version"] = attrs["pg_version"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["synthetic_storage_size_bytes"] = attrs["synthetic_storage_size_bytes"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["project_id"] = attrs["project_id"].SetComputed()
	attrs["project_id"] = attrs["project_id"].SetOptional()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in Project as
// a postgres_tf.ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetDefaultEndpointSettings(ctx context.Context) (postgres_tf.ProjectDefaultEndpointSettings, bool) {
	var e postgres_tf.ProjectDefaultEndpointSettings
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectDefaultEndpointSettings
	d := m.DefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in Project.
func (m *Project) SetDefaultEndpointSettings(ctx context.Context, v postgres_tf.ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.DefaultEndpointSettings = vs
}

// GetEffectiveDefaultEndpointSettings returns the value of the EffectiveDefaultEndpointSettings field in Project as
// a postgres_tf.ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetEffectiveDefaultEndpointSettings(ctx context.Context) (postgres_tf.ProjectDefaultEndpointSettings, bool) {
	var e postgres_tf.ProjectDefaultEndpointSettings
	if m.EffectiveDefaultEndpointSettings.IsNull() || m.EffectiveDefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectDefaultEndpointSettings
	d := m.EffectiveDefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveDefaultEndpointSettings sets the value of the EffectiveDefaultEndpointSettings field in Project.
func (m *Project) SetEffectiveDefaultEndpointSettings(ctx context.Context, v postgres_tf.ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveDefaultEndpointSettings = vs
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Project as
// a postgres_tf.ProjectSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetEffectiveSettings(ctx context.Context) (postgres_tf.ProjectSettings, bool) {
	var e postgres_tf.ProjectSettings
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectSettings
	d := m.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Project.
func (m *Project) SetEffectiveSettings(ctx context.Context, v postgres_tf.ProjectSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveSettings = vs
}

// GetSettings returns the value of the Settings field in Project as
// a postgres_tf.ProjectSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetSettings(ctx context.Context) (postgres_tf.ProjectSettings, bool) {
	var e postgres_tf.ProjectSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in Project.
func (m *Project) SetSettings(ctx context.Context, v postgres_tf.ProjectSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Project{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks postgres_project",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Project
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var project postgres.Project

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &project)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := postgres.CreateProjectRequest{
		Project:   project,
		ProjectId: plan.ProjectId.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.CreateProject(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create postgres_project", err.Error())
		return
	}

	var newState Project

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for postgres_project to be ready", err.Error())
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

func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Project
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetProjectRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
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

	var newState Project
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ProjectResource) update(ctx context.Context, plan Project, diags *diag.Diagnostics, state *tfsdk.State) {
	var project postgres.Project

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &project)...)
	if diags.HasError() {
		return
	}

	updateRequest := postgres.UpdateProjectRequest{
		Project:    project,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("default_endpoint_settings,display_name,history_retention_duration,settings", ",")),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Postgres.UpdateProject(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update postgres_project", err.Error())
		return
	}

	var newState Project

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for postgres_project update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Project
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Project
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest postgres.DeleteProjectRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Postgres.DeleteProject(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete postgres_project", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ProjectResource{}

func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
