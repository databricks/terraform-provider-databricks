// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps_settings_custom_template

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
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

const resourceName = "apps_settings_custom_template"

var _ resource.ResourceWithConfigure = &CustomTemplateResource{}

func ResourceCustomTemplate() resource.Resource {
	return &CustomTemplateResource{}
}

type CustomTemplateResource struct {
	Client *autogen.DatabricksClient
}

// CustomTemplate extends the main model with additional fields.
type CustomTemplate struct {
	apps_tf.CustomTemplate
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// CustomTemplate struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m CustomTemplate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.CustomTemplate.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTemplate
// only implements ToObjectValue() and Type().
func (m CustomTemplate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.CustomTemplate.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m CustomTemplate) Type(ctx context.Context) attr.Type {
	embeddedType := m.CustomTemplate.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *CustomTemplate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan CustomTemplate) {
	m.CustomTemplate.SyncFieldsDuringCreateOrUpdate(ctx, plan.CustomTemplate)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *CustomTemplate) SyncFieldsDuringRead(ctx context.Context, existingState CustomTemplate) {
	m.CustomTemplate.SyncFieldsDuringRead(ctx, existingState.CustomTemplate)
}

func (r *CustomTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CustomTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, CustomTemplate{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks apps_settings_custom_template",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CustomTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *CustomTemplateResource) update(ctx context.Context, plan CustomTemplate, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var custom_template apps.CustomTemplate

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &custom_template)...)
	if diags.HasError() {
		return
	}

	updateRequest := apps.UpdateCustomTemplateRequest{
		Template: custom_template,
		Name:     plan.Name.ValueString(),
	}

	response, err := client.AppsSettings.UpdateCustomTemplate(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update apps_settings_custom_template", err.Error())
		return
	}

	var newState CustomTemplate
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CustomTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan CustomTemplate
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var custom_template apps.CustomTemplate

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &custom_template)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := apps.CreateCustomTemplateRequest{
		Template: custom_template,
	}

	response, err := client.AppsSettings.CreateCustomTemplate(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create apps_settings_custom_template", err.Error())
		return
	}

	var newState CustomTemplate

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

func (r *CustomTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState CustomTemplate
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest apps.GetCustomTemplateRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.AppsSettings.GetCustomTemplate(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get apps_settings_custom_template", err.Error())
		return
	}

	var newState CustomTemplate
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CustomTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan CustomTemplate
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *CustomTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomTemplate
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest apps.DeleteCustomTemplateRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.AppsSettings.DeleteCustomTemplate(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete apps_settings_custom_template", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &CustomTemplateResource{}

func (r *CustomTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
