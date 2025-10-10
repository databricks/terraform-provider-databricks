// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package online_store

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
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

const resourceName = "online_store"

var _ resource.ResourceWithConfigure = &OnlineStoreResource{}

func ResourceOnlineStore() resource.Resource {
	return &OnlineStoreResource{}
}

type OnlineStoreResource struct {
	Client *autogen.DatabricksClient
}

// OnlineStore extends the main model with additional fields.
type OnlineStore struct {
	ml_tf.OnlineStore
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// OnlineStore struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m OnlineStore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.OnlineStore.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStore
// only implements ToObjectValue() and Type().
func (m OnlineStore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.OnlineStore.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m OnlineStore) Type(ctx context.Context) attr.Type {
	embeddedType := m.OnlineStore.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *OnlineStore) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan OnlineStore) {
	m.OnlineStore.SyncFieldsDuringCreateOrUpdate(ctx, plan.OnlineStore)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *OnlineStore) SyncFieldsDuringRead(ctx context.Context, existingState OnlineStore) {
	m.OnlineStore.SyncFieldsDuringRead(ctx, existingState.OnlineStore)
}

func (r *OnlineStoreResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *OnlineStoreResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, OnlineStore{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks online_store",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *OnlineStoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *OnlineStoreResource) update(ctx context.Context, plan OnlineStore, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var online_store ml.OnlineStore

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &online_store)...)
	if diags.HasError() {
		return
	}

	updateRequest := ml.UpdateOnlineStoreRequest{
		OnlineStore: online_store,
		Name:        plan.Name.ValueString(),
		UpdateMask:  "capacity,read_replica_count",
	}

	response, err := client.FeatureStore.UpdateOnlineStore(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update online_store", err.Error())
		return
	}

	var newState OnlineStore
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *OnlineStoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan OnlineStore
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var online_store ml.OnlineStore

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &online_store)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := ml.CreateOnlineStoreRequest{
		OnlineStore: online_store,
	}

	response, err := client.FeatureStore.CreateOnlineStore(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create online_store", err.Error())
		return
	}

	var newState OnlineStore

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

func (r *OnlineStoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState OnlineStore
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetOnlineStoreRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureStore.GetOnlineStore(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get online_store", err.Error())
		return
	}

	var newState OnlineStore
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *OnlineStoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan OnlineStore
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *OnlineStoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state OnlineStore
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest ml.DeleteOnlineStoreRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.FeatureStore.DeleteOnlineStore(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete online_store", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &OnlineStoreResource{}

func (r *OnlineStoreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
