// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package entity_tag_assignment

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
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

const resourceName = "entity_tag_assignment"

var _ resource.ResourceWithConfigure = &EntityTagAssignmentResource{}

func ResourceEntityTagAssignment() resource.Resource {
	return &EntityTagAssignmentResource{}
}

type EntityTagAssignmentResource struct {
	Client *autogen.DatabricksClient
}

// EntityTagAssignment extends the main model with additional fields.
type EntityTagAssignment struct {
	catalog_tf.EntityTagAssignment
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// EntityTagAssignment struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m EntityTagAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.EntityTagAssignment.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EntityTagAssignment
// only implements ToObjectValue() and Type().
func (m EntityTagAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.EntityTagAssignment.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()
	embeddedAttrs["workspace_id"] = m.WorkspaceID

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m EntityTagAssignment) Type(ctx context.Context) attr.Type {
	embeddedType := m.EntityTagAssignment.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()
	attrTypes["workspace_id"] = types.StringType

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *EntityTagAssignment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan EntityTagAssignment) {
	m.EntityTagAssignment.SyncFieldsDuringCreateOrUpdate(ctx, plan.EntityTagAssignment)
	m.WorkspaceID = plan.WorkspaceID
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *EntityTagAssignment) SyncFieldsDuringRead(ctx context.Context, existingState EntityTagAssignment) {
	m.EntityTagAssignment.SyncFieldsDuringRead(ctx, existingState.EntityTagAssignment)
	m.WorkspaceID = existingState.WorkspaceID
}

func (r *EntityTagAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *EntityTagAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, EntityTagAssignment{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "entity_type")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "entity_name")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "tag_key")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks entity_tag_assignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EntityTagAssignmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *EntityTagAssignmentResource) update(ctx context.Context, plan EntityTagAssignment, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var entity_tag_assignment catalog.EntityTagAssignment

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &entity_tag_assignment)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateEntityTagAssignmentRequest{
		TagAssignment: entity_tag_assignment,
		EntityName:    plan.EntityName.ValueString(),
		EntityType:    plan.EntityType.ValueString(),
		TagKey:        plan.TagKey.ValueString(),
		UpdateMask:    "tag_value",
	}

	response, err := client.EntityTagAssignments.Update(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update entity_tag_assignment", err.Error())
		return
	}

	var newState EntityTagAssignment
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *EntityTagAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan EntityTagAssignment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var entity_tag_assignment catalog.EntityTagAssignment

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &entity_tag_assignment)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateEntityTagAssignmentRequest{
		TagAssignment: entity_tag_assignment,
	}

	response, err := client.EntityTagAssignments.Create(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create entity_tag_assignment", err.Error())
		return
	}

	var newState EntityTagAssignment

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

func (r *EntityTagAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState EntityTagAssignment
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetEntityTagAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.EntityTagAssignments.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get entity_tag_assignment", err.Error())
		return
	}

	var newState EntityTagAssignment
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *EntityTagAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan EntityTagAssignment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *EntityTagAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state EntityTagAssignment
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteEntityTagAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.EntityTagAssignments.Delete(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete entity_tag_assignment", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &EntityTagAssignmentResource{}

func (r *EntityTagAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: entity_type,entity_name,tag_key. Got: %q",
				req.ID,
			),
		)
		return
	}

	entityType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entity_type"), entityType)...)
	entityName := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entity_name"), entityName)...)
	tagKey := parts[2]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tag_key"), tagKey)...)
}
