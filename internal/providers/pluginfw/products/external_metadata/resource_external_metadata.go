// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_metadata

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

const resourceName = "external_metadata"

var _ resource.ResourceWithConfigure = &ExternalMetadataResource{}

func ResourceExternalMetadata() resource.Resource {
	return &ExternalMetadataResource{}
}

type ExternalMetadataResource struct {
	Client *autogen.DatabricksClient
}

// ExternalMetadata extends the main model with additional fields.
type ExternalMetadata struct {
	catalog_tf.ExternalMetadata
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ExternalMetadata struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ExternalMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.ExternalMetadata.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalMetadata
// only implements ToObjectValue() and Type().
func (m ExternalMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.ExternalMetadata.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()
	embeddedAttrs["workspace_id"] = m.WorkspaceID

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ExternalMetadata) Type(ctx context.Context) attr.Type {
	embeddedType := m.ExternalMetadata.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()
	attrTypes["workspace_id"] = types.StringType

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *ExternalMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan ExternalMetadata) {
	m.ExternalMetadata.SyncFieldsDuringCreateOrUpdate(ctx, plan.ExternalMetadata)
	m.WorkspaceID = plan.WorkspaceID
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *ExternalMetadata) SyncFieldsDuringRead(ctx context.Context, existingState ExternalMetadata) {
	m.ExternalMetadata.SyncFieldsDuringRead(ctx, existingState.ExternalMetadata)
	m.WorkspaceID = existingState.WorkspaceID
}

func (r *ExternalMetadataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ExternalMetadataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ExternalMetadata{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks external_metadata",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalMetadataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ExternalMetadataResource) update(ctx context.Context, plan ExternalMetadata, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var external_metadata catalog.ExternalMetadata

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &external_metadata)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateExternalMetadataRequest{
		ExternalMetadata: external_metadata,
		Name:             plan.Name.ValueString(),
		UpdateMask:       "columns,description,entity_type,owner,properties,system_type,url",
	}

	response, err := client.ExternalMetadata.UpdateExternalMetadata(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ExternalMetadataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan ExternalMetadata
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var external_metadata catalog.ExternalMetadata

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &external_metadata)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateExternalMetadataRequest{
		ExternalMetadata: external_metadata,
	}

	response, err := client.ExternalMetadata.CreateExternalMetadata(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata

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

func (r *ExternalMetadataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState ExternalMetadata
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetExternalMetadataRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.ExternalMetadata.GetExternalMetadata(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get external_metadata", err.Error())
		return
	}

	var newState ExternalMetadata
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ExternalMetadataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ExternalMetadata
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ExternalMetadataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ExternalMetadata
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteExternalMetadataRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.ExternalMetadata.DeleteExternalMetadata(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete external_metadata", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &ExternalMetadataResource{}

func (r *ExternalMetadataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
