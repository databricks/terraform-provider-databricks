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
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
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

// AccessRequestDestinations extends the main model with additional fields.
type AccessRequestDestinations struct {
	// Indicates whether any destinations are hidden from the caller due to a
	// lack of permissions. This value is true if the caller does not have
	// permission to see all destinations.
	AreAnyDestinationsHidden types.Bool `tfsdk:"are_any_destinations_hidden"`
	// The access request destinations for the securable.
	Destinations types.List `tfsdk:"destinations"`
	// The securable for which the access request destinations are being
	// retrieved.
	Securable types.Object `tfsdk:"securable"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccessRequestDestinations struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccessRequestDestinations) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destinations": reflect.TypeOf(catalog_tf.NotificationDestination{}),
		"securable":    reflect.TypeOf(catalog_tf.Securable{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessRequestDestinations
// only implements ToObjectValue() and Type().
func (m AccessRequestDestinations) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"are_any_destinations_hidden": m.AreAnyDestinationsHidden,
			"destinations": m.Destinations,
			"securable":    m.Securable,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"are_any_destinations_hidden": types.BoolType,
			"destinations": basetypes.ListType{
				ElemType: catalog_tf.NotificationDestination{}.Type(ctx),
			},
			"securable": catalog_tf.Securable{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *AccessRequestDestinations) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccessRequestDestinations) {
	if !from.Securable.IsNull() && !from.Securable.IsUnknown() {
		if toSecurable, ok := to.GetSecurable(ctx); ok {
			if fromSecurable, ok := from.GetSecurable(ctx); ok {
				// Recursively sync the fields of Securable
				toSecurable.SyncFieldsDuringCreateOrUpdate(ctx, fromSecurable)
				to.SetSecurable(ctx, toSecurable)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AccessRequestDestinations) SyncFieldsDuringRead(ctx context.Context, from AccessRequestDestinations) {
	if !from.Securable.IsNull() && !from.Securable.IsUnknown() {
		if toSecurable, ok := to.GetSecurable(ctx); ok {
			if fromSecurable, ok := from.GetSecurable(ctx); ok {
				toSecurable.SyncFieldsDuringRead(ctx, fromSecurable)
				to.SetSecurable(ctx, toSecurable)
			}
		}
	}
}

func (m AccessRequestDestinations) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["are_any_destinations_hidden"] = attrs["are_any_destinations_hidden"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetRequired()
	attrs["securable"] = attrs["securable"].SetRequired()

	return attrs
}

// GetDestinations returns the value of the Destinations field in AccessRequestDestinations as
// a slice of catalog_tf.NotificationDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *AccessRequestDestinations) GetDestinations(ctx context.Context) ([]catalog_tf.NotificationDestination, bool) {
	if m.Destinations.IsNull() || m.Destinations.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.NotificationDestination
	d := m.Destinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinations sets the value of the Destinations field in AccessRequestDestinations.
func (m *AccessRequestDestinations) SetDestinations(ctx context.Context, v []catalog_tf.NotificationDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Destinations = types.ListValueMust(t, vs)
}

// GetSecurable returns the value of the Securable field in AccessRequestDestinations as
// a catalog_tf.Securable value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccessRequestDestinations) GetSecurable(ctx context.Context) (catalog_tf.Securable, bool) {
	var e catalog_tf.Securable
	if m.Securable.IsNull() || m.Securable.IsUnknown() {
		return e, false
	}
	var v catalog_tf.Securable
	d := m.Securable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurable sets the value of the Securable field in AccessRequestDestinations.
func (m *AccessRequestDestinations) SetSecurable(ctx context.Context, v catalog_tf.Securable) {
	vs := v.ToObjectValue(ctx)
	m.Securable = vs
}

func (r *AccessRequestDestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AccessRequestDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, AccessRequestDestinations{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks rfa_access_request_destinations",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccessRequestDestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *AccessRequestDestinationResource) update(ctx context.Context, plan AccessRequestDestinations, diags *diag.Diagnostics, state *tfsdk.State) {
	var access_request_destinations catalog.AccessRequestDestinations

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &access_request_destinations)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateAccessRequestDestinationsRequest{
		AccessRequestDestinations: access_request_destinations,
		UpdateMask:                "destinations,securable",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Rfa.UpdateAccessRequestDestinations(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update rfa_access_request_destinations", err.Error())
		return
	}

	var newState AccessRequestDestinations
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccessRequestDestinations
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AccessRequestDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState AccessRequestDestinations
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetAccessRequestDestinationsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
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

	var newState AccessRequestDestinations
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccessRequestDestinations
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
