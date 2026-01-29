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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	// The source securable from which the destinations are inherited. Either
	// the same value as securable (if destination is set directly on the
	// securable) or the nearest parent securable with destinations set.
	DestinationSourceSecurable types.Object `tfsdk:"destination_source_securable"`
	// The access request destinations for the securable.
	Destinations types.List `tfsdk:"destinations"`
	// The full name of the securable. Redundant with the name in the securable
	// object, but necessary for Terraform integration
	FullName types.String `tfsdk:"full_name"`
	// The securable for which the access request destinations are being
	// modified or read.
	Securable types.Object `tfsdk:"securable"`
	// The type of the securable. Redundant with the type in the securable
	// object, but necessary for Terraform integration
	SecurableType types.String `tfsdk:"securable_type"`
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
		"destination_source_securable": reflect.TypeOf(catalog_tf.Securable{}),
		"destinations":                 reflect.TypeOf(catalog_tf.NotificationDestination{}),
		"securable":                    reflect.TypeOf(catalog_tf.Securable{}),
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
			"destination_source_securable": m.DestinationSourceSecurable,
			"destinations":                 m.Destinations,
			"full_name":                    m.FullName,
			"securable":                    m.Securable,
			"securable_type":               m.SecurableType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinations) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"are_any_destinations_hidden": types.BoolType,
			"destination_source_securable": catalog_tf.Securable{}.Type(ctx),
			"destinations": basetypes.ListType{
				ElemType: catalog_tf.NotificationDestination{}.Type(ctx),
			},
			"full_name":      types.StringType,
			"securable":      catalog_tf.Securable{}.Type(ctx),
			"securable_type": types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *AccessRequestDestinations) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccessRequestDestinations) {
	if !from.DestinationSourceSecurable.IsNull() && !from.DestinationSourceSecurable.IsUnknown() {
		if toDestinationSourceSecurable, ok := to.GetDestinationSourceSecurable(ctx); ok {
			if fromDestinationSourceSecurable, ok := from.GetDestinationSourceSecurable(ctx); ok {
				// Recursively sync the fields of DestinationSourceSecurable
				toDestinationSourceSecurable.SyncFieldsDuringCreateOrUpdate(ctx, fromDestinationSourceSecurable)
				to.SetDestinationSourceSecurable(ctx, toDestinationSourceSecurable)
			}
		}
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() && to.Destinations.IsNull() && len(from.Destinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Destinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Destinations = from.Destinations
	}
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
	if !from.DestinationSourceSecurable.IsNull() && !from.DestinationSourceSecurable.IsUnknown() {
		if toDestinationSourceSecurable, ok := to.GetDestinationSourceSecurable(ctx); ok {
			if fromDestinationSourceSecurable, ok := from.GetDestinationSourceSecurable(ctx); ok {
				toDestinationSourceSecurable.SyncFieldsDuringRead(ctx, fromDestinationSourceSecurable)
				to.SetDestinationSourceSecurable(ctx, toDestinationSourceSecurable)
			}
		}
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() && to.Destinations.IsNull() && len(from.Destinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Destinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Destinations = from.Destinations
	}
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
	attrs["destination_source_securable"] = attrs["destination_source_securable"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["securable"] = attrs["securable"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetComputed()

	attrs["securable_type"] = attrs["securable_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetDestinationSourceSecurable returns the value of the DestinationSourceSecurable field in AccessRequestDestinations as
// a catalog_tf.Securable value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccessRequestDestinations) GetDestinationSourceSecurable(ctx context.Context) (catalog_tf.Securable, bool) {
	var e catalog_tf.Securable
	if m.DestinationSourceSecurable.IsNull() || m.DestinationSourceSecurable.IsUnknown() {
		return e, false
	}
	var v catalog_tf.Securable
	d := m.DestinationSourceSecurable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinationSourceSecurable sets the value of the DestinationSourceSecurable field in AccessRequestDestinations.
func (m *AccessRequestDestinations) SetDestinationSourceSecurable(ctx context.Context, v catalog_tf.Securable) {
	vs := v.ToObjectValue(ctx)
	m.DestinationSourceSecurable = vs
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
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccessRequestDestinations
	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.SetDestinations(ctx, nil)

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

var _ resource.ResourceWithImportState = &AccessRequestDestinationResource{}

func (r *AccessRequestDestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: securable_type,full_name. Got: %q",
				req.ID,
			),
		)
		return
	}

	securableType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("securable_type"), securableType)...)
	fullName := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("full_name"), fullName)...)
}
