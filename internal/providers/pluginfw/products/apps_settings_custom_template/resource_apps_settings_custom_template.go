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
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
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
	Creator types.String `tfsdk:"creator"`
	// The description of the template.
	Description types.String `tfsdk:"description"`
	// The Git provider of the template.
	GitProvider types.String `tfsdk:"git_provider"`
	// The Git repository URL that the template resides in.
	GitRepo types.String `tfsdk:"git_repo"`
	// The manifest of the template. It defines fields and default values when
	// installing the template.
	Manifest types.Object `tfsdk:"manifest"`
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"name"`
	// The path to the template within the Git repository.
	Path types.String `tfsdk:"path"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// CustomTemplate struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m CustomTemplate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(apps_tf.AppManifest{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTemplate
// only implements ToObjectValue() and Type().
func (m CustomTemplate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"creator": m.Creator,
			"description":  m.Description,
			"git_provider": m.GitProvider,
			"git_repo":     m.GitRepo,
			"manifest":     m.Manifest,
			"name":         m.Name,
			"path":         m.Path,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m CustomTemplate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"creator": types.StringType,
			"description":  types.StringType,
			"git_provider": types.StringType,
			"git_repo":     types.StringType,
			"manifest":     apps_tf.AppManifest{}.Type(ctx),
			"name":         types.StringType,
			"path":         types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *CustomTemplate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTemplate) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *CustomTemplate) SyncFieldsDuringRead(ctx context.Context, from CustomTemplate) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
}

func (m CustomTemplate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_repo"] = attrs["git_repo"].SetRequired()
	attrs["manifest"] = attrs["manifest"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["path"] = attrs["path"].SetRequired()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetManifest returns the value of the Manifest field in CustomTemplate as
// a apps_tf.AppManifest value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomTemplate) GetManifest(ctx context.Context) (apps_tf.AppManifest, bool) {
	var e apps_tf.AppManifest
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v apps_tf.AppManifest
	d := m.Manifest.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManifest sets the value of the Manifest field in CustomTemplate.
func (m *CustomTemplate) SetManifest(ctx context.Context, v apps_tf.AppManifest) {
	vs := v.ToObjectValue(ctx)
	m.Manifest = vs
}

func (r *CustomTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CustomTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, CustomTemplate{}, nil)
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
	var custom_template apps.CustomTemplate

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &custom_template)...)
	if diags.HasError() {
		return
	}

	updateRequest := apps.UpdateCustomTemplateRequest{
		Template: custom_template,
		Name:     plan.Name.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
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

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
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

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
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

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
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
