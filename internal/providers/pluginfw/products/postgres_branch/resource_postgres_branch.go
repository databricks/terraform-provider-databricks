// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_branch

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

const resourceName = "postgres_branch"

var _ resource.ResourceWithConfigure = &BranchResource{}

func ResourceBranch() resource.Resource {
	return &BranchResource{}
}

type BranchResource struct {
	Client *autogen.DatabricksClient
}

// Branch extends the main model with additional fields.
type Branch struct {
	// The ID to use for the Branch, which will become the final component of
	// the branch's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	BranchId types.String `tfsdk:"branch_id"`
	// A timestamp indicating when the branch was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"name"`
	// The project containing this branch. Format: projects/{project_id}
	Parent types.String `tfsdk:"parent"`
	// The desired state of a Branch.
	Spec types.Object `tfsdk:"spec"`
	// The current status of a Branch.
	Status types.Object `tfsdk:"status"`
	// System generated unique ID for the branch.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Branch struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Branch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(postgres_tf.BranchSpec{}),
		"status": reflect.TypeOf(postgres_tf.BranchStatus{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Branch
// only implements ToObjectValue() and Type().
func (m Branch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"branch_id": m.BranchId,
			"create_time": m.CreateTime,
			"name":        m.Name,
			"parent":      m.Parent,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Branch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"branch_id": types.StringType,
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec":        postgres_tf.BranchSpec{}.Type(ctx),
			"status":      postgres_tf.BranchStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Branch) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Branch) {
	to.BranchId = from.BranchId
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
func (to *Branch) SyncFieldsDuringRead(ctx context.Context, from Branch) {
	to.BranchId = from.BranchId
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

func (m Branch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetSpec returns the value of the Spec field in Branch as
// a postgres_tf.BranchSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch) GetSpec(ctx context.Context) (postgres_tf.BranchSpec, bool) {
	var e postgres_tf.BranchSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v postgres_tf.BranchSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Branch.
func (m *Branch) SetSpec(ctx context.Context, v postgres_tf.BranchSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Branch as
// a postgres_tf.BranchStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch) GetStatus(ctx context.Context) (postgres_tf.BranchStatus, bool) {
	var e postgres_tf.BranchStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v postgres_tf.BranchStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Branch.
func (m *Branch) SetStatus(ctx context.Context, v postgres_tf.BranchStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

func (r *BranchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *BranchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Branch{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks postgres_branch",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BranchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *BranchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Branch
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var branch postgres.Branch

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &branch)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := postgres.CreateBranchRequest{
		Branch:   branch,
		Parent:   plan.Parent.ValueString(),
		BranchId: plan.BranchId.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.CreateBranch(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create postgres_branch", err.Error())
		return
	}

	var newState Branch

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for postgres_branch to be ready", err.Error())
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

func (r *BranchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Branch
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetBranchRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Postgres.GetBranch(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get postgres_branch", err.Error())
		return
	}

	var newState Branch
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *BranchResource) update(ctx context.Context, plan Branch, diags *diag.Diagnostics, state *tfsdk.State) {
	var branch postgres.Branch

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &branch)...)
	if diags.HasError() {
		return
	}

	updateRequest := postgres.UpdateBranchRequest{
		Branch:     branch,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("spec", ",")),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Postgres.UpdateBranch(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update postgres_branch", err.Error())
		return
	}

	var newState Branch

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for postgres_branch update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *BranchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Branch
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *BranchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Branch
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest postgres.DeleteBranchRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.DeleteBranch(ctx, deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete postgres_branch", err.Error())
		return
	}

	err = response.Wait(ctx)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("error waiting for postgres_branch delete", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &BranchResource{}

func (r *BranchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
