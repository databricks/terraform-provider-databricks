package app

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	resourceName       = "app"
	resourceNamePlural = "apps"
)

type appResource struct {
	apps_tf.App
	NoCompute types.Bool `tfsdk:"no_compute"`
}

func ResourceApp() resource.Resource {
	return &resourceApp{}
}

type resourceApp struct {
	client *common.DatabricksClient
}

func (a resourceApp) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (a resourceApp) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.ResourceStructToSchema(ctx, appResource{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		cs.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name")
		exclusiveFields := []string{"job", "secret", "serving_endpoint", "sql_warehouse"}
		paths := path.Expressions{}
		for _, field := range exclusiveFields[1:] {
			paths = append(paths, path.MatchRelative().AtParent().AtName(field))
		}
		cs.AddValidator(objectvalidator.ExactlyOneOf(paths...), "resources", exclusiveFields[0])
		return cs
	})
}

func (a *resourceApp) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if a.client == nil && req.ProviderData != nil {
		a.client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (a *resourceApp) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app appResource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &app)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var appGoSdk apps.App
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, app, &appGoSdk)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the app
	var forceSendFields []string
	if !app.NoCompute.IsNull() {
		forceSendFields = append(forceSendFields, "NoCompute")
	}
	waiter, err := w.Apps.Create(ctx, apps.CreateAppRequest{
		App: &appGoSdk,
		// NoCompute: app.NoCompute.ValueBool(),
		// ForceSendFields: forceSendFields,
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", err.Error())
		return
	}

	// Store the initial version of the app in state
	var newApp appResource
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waiter.Response, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	newApp.NoCompute = app.NoCompute
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Wait for the app to be created
	finalApp, err := waiter.Get()
	if err != nil {
		resp.Diagnostics.AddError("error waiting for app to be ready", err.Error())
		return
	}

	// Store the final version of the app in state
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, finalApp, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (a *resourceApp) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app apps_tf.App
	resp.Diagnostics.Append(req.State.Get(ctx, &app)...)
	if resp.Diagnostics.HasError() {
		return
	}

	appGoSdk, err := w.Apps.GetByName(ctx, app.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", err.Error())
		return
	}

	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (a *resourceApp) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app appResource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &app)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update the app
	var appGoSdk apps.App
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, app, &appGoSdk)...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := w.Apps.Update(ctx, apps.UpdateAppRequest{App: &appGoSdk, Name: app.Name.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("failed to update app", err.Error())
		return
	}

	// Store the updated version of the app in state
	var newApp appResource
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	newApp.NoCompute = app.NoCompute
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (a *resourceApp) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app apps_tf.App
	resp.Diagnostics.Append(req.State.Get(ctx, &app)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the app
	_, err := w.Apps.DeleteByName(ctx, app.Name.ValueString())
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete app", err.Error())
		return
	}
}

func (a *resourceApp) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

var _ resource.ResourceWithConfigure = &resourceApp{}
var _ resource.ResourceWithImportState = &resourceApp{}
