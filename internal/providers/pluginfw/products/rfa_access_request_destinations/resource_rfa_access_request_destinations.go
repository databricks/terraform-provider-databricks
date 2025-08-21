// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package rfa_access_request_destinations

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "rfa_access_request_destinations"

var _ resource.ResourceWithConfigure = &AccessRequestDestinationResource{}

func ResourceAccessRequestDestination() resource.Resource {
	return &AccessRequestDestinationResource{}
}

type AccessRequestDestinationResource struct {
	Client *autogen.DatabricksClient
}

func (r *AccessRequestDestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AccessRequestDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, catalog_tf.AccessRequestDestinations{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
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

func (r *AccessRequestDestinationResource) update(ctx context.Context, plan catalog_tf.AccessRequestDestinations, diags *diag.Diagnostics, state *tfsdk.State) {
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

	var newState catalog_tf.AccessRequestDestinations
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan catalog_tf.AccessRequestDestinations
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

	var existingState catalog_tf.AccessRequestDestinations
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

	var newState catalog_tf.AccessRequestDestinations
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *AccessRequestDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan catalog_tf.AccessRequestDestinations
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
