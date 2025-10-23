// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_refresh

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "data_quality_refresh"

var _ resource.ResourceWithConfigure = &RefreshResource{}

func ResourceRefresh() resource.Resource {
	return &RefreshResource{}
}

type RefreshResource struct {
	Client *autogen.DatabricksClient
}

// Refresh extends the main model with additional fields.
type Refresh struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message"`
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType types.String `tfsdk:"object_type"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id"`
	// Time when the refresh started (milliseconds since 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
	// The current state of the refresh.
	State types.String `tfsdk:"state"`
	// What triggered the refresh.
	Trigger types.String `tfsdk:"trigger"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Refresh struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Refresh) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Refresh
// only implements ToObjectValue() and Type().
func (m Refresh) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"end_time_ms": m.EndTimeMs,
			"message":       m.Message,
			"object_id":     m.ObjectId,
			"object_type":   m.ObjectType,
			"refresh_id":    m.RefreshId,
			"start_time_ms": m.StartTimeMs,
			"state":         m.State,
			"trigger":       m.Trigger,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Refresh) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"end_time_ms": types.Int64Type,
			"message":       types.StringType,
			"object_id":     types.StringType,
			"object_type":   types.StringType,
			"refresh_id":    types.Int64Type,
			"start_time_ms": types.Int64Type,
			"state":         types.StringType,
			"trigger":       types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Refresh) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Refresh) {
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Refresh) SyncFieldsDuringRead(ctx context.Context, from Refresh) {
}

func (m Refresh) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time_ms"] = attrs["end_time_ms"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetComputed()
	attrs["start_time_ms"] = attrs["start_time_ms"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["trigger"] = attrs["trigger"].SetComputed()

	attrs["object_type"] = attrs["object_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["object_id"] = attrs["object_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["refresh_id"] = attrs["refresh_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

func (r *RefreshResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *RefreshResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Refresh{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks data_quality_refresh",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RefreshResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *RefreshResource) update(ctx context.Context, plan Refresh, diags *diag.Diagnostics, state *tfsdk.State) {
	var refresh dataquality.Refresh

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &refresh)...)
	if diags.HasError() {
		return
	}

	updateRequest := dataquality.UpdateRefreshRequest{
		Refresh:    refresh,
		ObjectId:   plan.ObjectId.ValueString(),
		ObjectType: plan.ObjectType.ValueString(),
		RefreshId:  plan.RefreshId.ValueInt64(),
		UpdateMask: "",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.DataQuality.UpdateRefresh(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update data_quality_refresh", err.Error())
		return
	}

	var newState Refresh

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *RefreshResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Refresh
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var refresh dataquality.Refresh

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &refresh)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := dataquality.CreateRefreshRequest{
		Refresh:    refresh,
		ObjectId:   plan.ObjectId.ValueString(),
		ObjectType: plan.ObjectType.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.CreateRefresh(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create data_quality_refresh", err.Error())
		return
	}

	var newState Refresh

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

func (r *RefreshResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Refresh
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest dataquality.GetRefreshRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.DataQuality.GetRefresh(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get data_quality_refresh", err.Error())
		return
	}

	var newState Refresh
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *RefreshResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Refresh
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *RefreshResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Refresh
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest dataquality.DeleteRefreshRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.DataQuality.DeleteRefresh(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete data_quality_refresh", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &RefreshResource{}

func (r *RefreshResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: object_type,object_id,refresh_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	objectType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_type"), objectType)...)
	objectId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_id"), objectId)...)
	refreshId, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("refresh_id"), refreshId)...)
}
