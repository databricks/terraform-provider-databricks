// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package rfa_access_request_destinations

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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "rfa_access_request_destinations"

var _ resource.ResourceWithConfigure = &AccessRequestDestinationResource{}

func ResourceAccessRequestDestination() resource.Resource {
	return &AccessRequestDestinationResource{}
}

type AccessRequestDestinationResource struct {
	Client *autogen.DatabricksClient
}

// AccessRequestDestinationsExtended extends the main model with additional fields.
type AccessRequestDestinationsExtended struct {
	catalog_tf.AccessRequestDestinations
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccessRequestDestinationsExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccessRequestDestinationsExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.AccessRequestDestinations.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessRequestDestinationsExtended
// only implements ToObjectValue() and Type().
func (m AccessRequestDestinationsExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.AccessRequestDestinations.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinationsExtended) Type(ctx context.Context) attr.Type {
	embeddedType := m.AccessRequestDestinations.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *AccessRequestDestinationsExtended) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan AccessRequestDestinationsExtended) {
	m.AccessRequestDestinations.SyncFieldsDuringCreateOrUpdate(ctx, plan.AccessRequestDestinations)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *AccessRequestDestinationsExtended) SyncFieldsDuringRead(ctx context.Context, existingState AccessRequestDestinationsExtended) {
	m.AccessRequestDestinations.SyncFieldsDuringRead(ctx, existingState.AccessRequestDestinations)
}

func (r *AccessRequestDestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AccessRequestDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, AccessRequestDestinationsExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks rfa_access_request_destinations",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccessRequestDestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *AccessRequestDestinationResource) update(ctx context.Context, plan AccessRequestDestinationsExtended, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var access_request_destinations catalog.AccessRequestDestinations

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &access_request_destinations)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateAccessRequestDestinationsRequest{
		AccessRequestDestinations: access_request_destinations,
		UpdateMask:                "destinations,securable",
	}

	response, err := client.Rfa.UpdateAccessRequestDestinations(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update rfa_access_request_destinations", err.Error())
		return
	}

	var newState AccessRequestDestinationsExtended
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccessRequestDestinationsExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AccessRequestDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState AccessRequestDestinationsExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetAccessRequestDestinationsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Rfa.GetAccessRequestDestinations(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get rfa_access_request_destinations", err.Error())
		return
	}

	var newState AccessRequestDestinationsExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccessRequestDestinationsExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AccessRequestDestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &AccessRequestDestinationResource{}

func (r *AccessRequestDestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 0 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: . Got: %q",
				req.ID,
			),
		)
		return
	}

}
