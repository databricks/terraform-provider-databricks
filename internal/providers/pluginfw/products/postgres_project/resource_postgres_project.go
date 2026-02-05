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
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Output only. The full resource path of the project. Format:
	// projects/{project_id}
	Name types.String `tfsdk:"name"`
	// The ID to use for the Project. This becomes the final component of the
	// project's resource name. The ID is required and must be 1-63 characters
	// long, start with a lowercase letter, and contain only lowercase letters,
	// numbers, and hyphens. For example, `my-app` becomes `projects/my-app`.
	ProjectId types.String `tfsdk:"project_id"`
	// The spec contains the project configuration, including display_name,
	// pg_version (Postgres version), history_retention_duration, and
	// default_endpoint_settings.
	Spec types.Object `tfsdk:"spec"`
	// The current status of a Project.
	Status types.Object `tfsdk:"status"`
	// System-generated unique ID for the project.
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
		"spec":   reflect.TypeOf(postgres_tf.ProjectSpec{}),
		"status": reflect.TypeOf(postgres_tf.ProjectStatus{}),
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
		map[string]attr.Value{"create_time": m.CreateTime,
			"name":        m.Name,
			"project_id":  m.ProjectId,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Project) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"project_id":  types.StringType,
			"spec":        postgres_tf.ProjectSpec{}.Type(ctx),
			"status":      postgres_tf.ProjectStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Project) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project) {
	to.ProjectId = from.ProjectId
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
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Project) SyncFieldsDuringRead(ctx context.Context, from Project) {
	to.ProjectId = from.ProjectId
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
}

func (m Project) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetSpec returns the value of the Spec field in Project as
// a postgres_tf.ProjectSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetSpec(ctx context.Context) (postgres_tf.ProjectSpec, bool) {
	var e postgres_tf.ProjectSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Project.
func (m *Project) SetSpec(ctx context.Context, v postgres_tf.ProjectSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Project as
// a postgres_tf.ProjectStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetStatus(ctx context.Context) (postgres_tf.ProjectStatus, bool) {
	var e postgres_tf.ProjectStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v postgres_tf.ProjectStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Project.
func (m *Project) SetStatus(ctx context.Context, v postgres_tf.ProjectStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
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
		UpdateMask: *fieldmask.New(strings.Split("spec", ",")),
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

	response, err := client.Postgres.DeleteProject(ctx, deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete postgres_project", err.Error())
		return
	}

	err = response.Wait(ctx)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("error waiting for postgres_project delete", err.Error())
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
