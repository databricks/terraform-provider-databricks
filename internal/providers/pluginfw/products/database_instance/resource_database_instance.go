// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_instance

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

const resourceName = "database_instance"

var _ resource.ResourceWithConfigure = &DatabaseInstanceResource{}

func ResourceDatabaseInstance() resource.Resource {
	return &DatabaseInstanceResource{}
}

type DatabaseInstanceResource struct {
	Client *autogen.DatabricksClient
}

func (r *DatabaseInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DatabaseInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, database_tf.DatabaseInstance{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks database_instance",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatabaseInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DatabaseInstanceResource) update(ctx context.Context, plan database_tf.DatabaseInstance, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var database_instance database.DatabaseInstance
	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_instance)...)
	if diags.HasError() {
		return
	}

	updateRequest := database.UpdateDatabaseInstanceRequest{
		DatabaseInstance: database_instance,
		Name:             plan.Name.ValueString(),
		UpdateMask:       "capacity,enable_readable_secondaries,node_count,parent_instance_ref,retention_window_in_days,stopped",
	}

	response, err := client.Database.UpdateDatabaseInstance(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update database_instance", err.Error())
		return
	}

	var newState database_tf.DatabaseInstance
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *DatabaseInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan database_tf.DatabaseInstance
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var database_instance database.DatabaseInstance
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_instance)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := database.CreateDatabaseInstanceRequest{
		DatabaseInstance: database_instance,
	}

	response, err := client.Database.CreateDatabaseInstance(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create database_instance", err.Error())
		return
	}

	var newState database_tf.DatabaseInstance

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DatabaseInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState database_tf.DatabaseInstance
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest database.GetDatabaseInstanceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.GetDatabaseInstance(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get database_instance", err.Error())
		return
	}

	var newState database_tf.DatabaseInstance
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringRead(existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *DatabaseInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan database_tf.DatabaseInstance
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *DatabaseInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state database_tf.DatabaseInstance
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest database.DeleteDatabaseInstanceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Database.DeleteDatabaseInstance(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete database_instance", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &DatabaseInstanceResource{}

func (r *DatabaseInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
