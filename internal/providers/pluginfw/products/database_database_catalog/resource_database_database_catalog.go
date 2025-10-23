// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_database_catalog

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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "database_database_catalog"

var _ resource.ResourceWithConfigure = &DatabaseCatalogResource{}

func ResourceDatabaseCatalog() resource.Resource {
	return &DatabaseCatalogResource{}
}

type DatabaseCatalogResource struct {
	Client *autogen.DatabricksClient
}

// DatabaseCatalog extends the main model with additional fields.
type DatabaseCatalog struct {
	CreateDatabaseIfNotExists types.Bool `tfsdk:"create_database_if_not_exists"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName types.String `tfsdk:"database_name"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"name"`

	Uid types.String `tfsdk:"uid"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DatabaseCatalog struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DatabaseCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog
// only implements ToObjectValue() and Type().
func (m DatabaseCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_database_if_not_exists": m.CreateDatabaseIfNotExists,
			"database_instance_name": m.DatabaseInstanceName,
			"database_name":          m.DatabaseName,
			"name":                   m.Name,
			"uid":                    m.Uid,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DatabaseCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_database_if_not_exists": types.BoolType,
			"database_instance_name": types.StringType,
			"database_name":          types.StringType,
			"name":                   types.StringType,
			"uid":                    types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *DatabaseCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCatalog) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *DatabaseCatalog) SyncFieldsDuringRead(ctx context.Context, from DatabaseCatalog) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

func (m DatabaseCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetOptional()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetComputed()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["database_instance_name"] = attrs["database_instance_name"].SetRequired()
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["uid"] = attrs["uid"].SetComputed()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

func (r *DatabaseCatalogResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DatabaseCatalogResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, DatabaseCatalog{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks database_database_catalog",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatabaseCatalogResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DatabaseCatalogResource) update(ctx context.Context, plan DatabaseCatalog, diags *diag.Diagnostics, state *tfsdk.State) {
	var database_catalog database.DatabaseCatalog

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_catalog)...)
	if diags.HasError() {
		return
	}

	updateRequest := database.UpdateDatabaseCatalogRequest{
		DatabaseCatalog: database_catalog,
		Name:            plan.Name.ValueString(),
		UpdateMask:      "create_database_if_not_exists,database_instance_name,database_name",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Database.UpdateDatabaseCatalog(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update database_database_catalog", err.Error())
		return
	}

	var newState DatabaseCatalog

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *DatabaseCatalogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan DatabaseCatalog
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var database_catalog database.DatabaseCatalog

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_catalog)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := database.CreateDatabaseCatalogRequest{
		Catalog: database_catalog,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.CreateDatabaseCatalog(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create database_database_catalog", err.Error())
		return
	}

	var newState DatabaseCatalog

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

func (r *DatabaseCatalogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState DatabaseCatalog
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest database.GetDatabaseCatalogRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Database.GetDatabaseCatalog(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get database_database_catalog", err.Error())
		return
	}

	var newState DatabaseCatalog
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *DatabaseCatalogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan DatabaseCatalog
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *DatabaseCatalogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state DatabaseCatalog
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest database.DeleteDatabaseCatalogRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Database.DeleteDatabaseCatalog(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete database_database_catalog", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &DatabaseCatalogResource{}

func (r *DatabaseCatalogResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
