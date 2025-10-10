// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_synced_database_table

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/database_tf"
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

const resourceName = "database_synced_database_table"

var _ resource.ResourceWithConfigure = &SyncedDatabaseTableResource{}

func ResourceSyncedDatabaseTable() resource.Resource {
	return &SyncedDatabaseTableResource{}
}

type SyncedDatabaseTableResource struct {
	Client *autogen.DatabricksClient
}

// SyncedDatabaseTable extends the main model with additional fields.
type SyncedDatabaseTable struct {
	database_tf.SyncedDatabaseTable
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// SyncedDatabaseTable struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m SyncedDatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.SyncedDatabaseTable.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable
// only implements ToObjectValue() and Type().
func (m SyncedDatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.SyncedDatabaseTable.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m SyncedDatabaseTable) Type(ctx context.Context) attr.Type {
	embeddedType := m.SyncedDatabaseTable.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *SyncedDatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan SyncedDatabaseTable) {
	m.SyncedDatabaseTable.SyncFieldsDuringCreateOrUpdate(ctx, plan.SyncedDatabaseTable)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *SyncedDatabaseTable) SyncFieldsDuringRead(ctx context.Context, existingState SyncedDatabaseTable) {
	m.SyncedDatabaseTable.SyncFieldsDuringRead(ctx, existingState.SyncedDatabaseTable)
}

func (r *SyncedDatabaseTableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SyncedDatabaseTableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, SyncedDatabaseTable{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks database_synced_database_table",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SyncedDatabaseTableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *SyncedDatabaseTableResource) update(ctx context.Context, plan SyncedDatabaseTable, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var synced_database_table database.SyncedDatabaseTable

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &synced_database_table)...)
	if diags.HasError() {
		return
	}

	updateRequest := database.UpdateSyncedDatabaseTableRequest{
		SyncedTable: synced_database_table,
		Name:        plan.Name.ValueString(),
		UpdateMask:  "database_instance_name,logical_database_name,spec",
	}

	response, err := client.Database.UpdateSyncedDatabaseTable(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update database_synced_database_table", err.Error())
		return
	}

	var newState SyncedDatabaseTable
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *SyncedDatabaseTableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan SyncedDatabaseTable
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var synced_database_table database.SyncedDatabaseTable

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &synced_database_table)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := database.CreateSyncedDatabaseTableRequest{
		SyncedTable: synced_database_table,
	}

	response, err := client.Database.CreateSyncedDatabaseTable(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create database_synced_database_table", err.Error())
		return
	}

	var newState SyncedDatabaseTable

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

func (r *SyncedDatabaseTableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState SyncedDatabaseTable
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest database.GetSyncedDatabaseTableRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.GetSyncedDatabaseTable(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get database_synced_database_table", err.Error())
		return
	}

	var newState SyncedDatabaseTable
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *SyncedDatabaseTableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan SyncedDatabaseTable
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SyncedDatabaseTableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SyncedDatabaseTable
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest database.DeleteSyncedDatabaseTableRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Database.DeleteSyncedDatabaseTable(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete database_synced_database_table", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &SyncedDatabaseTableResource{}

func (r *SyncedDatabaseTableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
