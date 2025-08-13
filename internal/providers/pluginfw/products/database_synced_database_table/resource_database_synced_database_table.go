// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_synced_database_table

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/database_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "database_synced_database_table"

var _ resource.ResourceWithConfigure = &SyncedDatabaseTableResource{}

func ResourceSyncedDatabaseTable() resource.Resource {
	return &SyncedDatabaseTableResource{}
}

type SyncedDatabaseTableResource struct {
	Client *autogen.DatabricksClient
}

func (r *SyncedDatabaseTableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SyncedDatabaseTableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, database_tf.SyncedDatabaseTable{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
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

func (r *SyncedDatabaseTableResource) update(ctx context.Context, plan database_tf.SyncedDatabaseTable, diags *diag.Diagnostics, state *tfsdk.State) {
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

	var newState database_tf.SyncedDatabaseTable
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
	var plan database_tf.SyncedDatabaseTable
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

	var newState database_tf.SyncedDatabaseTable

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

	var existingState database_tf.SyncedDatabaseTable
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

	var newState database_tf.SyncedDatabaseTable
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *SyncedDatabaseTableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan database_tf.SyncedDatabaseTable
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

	var state database_tf.SyncedDatabaseTable
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
