package app

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const (
	resourceName               = "app"
	defaultAppProvisionTimeout = 10 * time.Minute
	deleteCallTimeout          = 10 * time.Second
)

func ResourceApp() resource.Resource {
	return &resourceApp{}
}

type resourceApp struct {
	client *common.DatabricksClient
}

func (a resourceApp) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

type makeListBlockUnknownIfNotInState struct{}

func (m makeListBlockUnknownIfNotInState) Description(_ context.Context) string {
	return "Make the field unknown if not in state"
}

func (m makeListBlockUnknownIfNotInState) MarkdownDescription(_ context.Context) string {
	return "Make the field unknown if not in state"
}

func (m makeListBlockUnknownIfNotInState) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.StateValue.IsNull() {
		elem := types.ObjectUnknown(req.PlanValue.ElementType(ctx).(basetypes.ObjectType).AttrTypes)
		var d diag.Diagnostics
		resp.PlanValue, d = types.ListValue(req.PlanValue.ElementType(ctx), []attr.Value{elem})
		resp.Diagnostics.Append(d...)
		return
	}
	listplanmodifier.UseStateForUnknown().PlanModifyList(ctx, req, resp)
}

func (a resourceApp) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.ResourceStructToSchema(apps_tf.App{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		cs.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name")
		// Computed fields
		for _, p := range []string{
			"active_deployment",
			"app_status",
			"compute_status",
			"create_time",
			"creator",
			"default_source_code_path",
			"pending_deployment",
			"service_principal_client_id",
			"service_principal_id",
			"service_principal_name",
			"update_time",
			"updater",
			"url",
		} {
			cs.SetReadOnly(p)
		}
		// All pointers are treated as list blocks to be compatible with resources implemented in SDKv2.
		// The plugin framework requires that the number of blocks in the config and plan match. This means that
		// it isn't possible to have a computed list block that is not part of the config. To work around this,
		// we need to treat these blocks as attributes in the schema, which allows us to set them as computed.
		for _, p := range []string{"app_status", "compute_status"} {
			cs.Transform(func(bsb tfschema.BaseSchemaBuilder) tfschema.BaseSchemaBuilder {
				switch b := bsb.(type) {
				case tfschema.ListNestedBlockBuilder:
					return tfschema.SingleNestedBlockBuilder{
						NestedObject: b.NestedObject,
						Computed:     true,
					}
				}
				return bsb
			}, p)
		}
		exclusiveFields := []string{"job", "secret", "serving_endpoint", "sql_warehouse"}
		for _, field := range exclusiveFields {
			paths := path.Expressions{}
			for _, f := range exclusiveFields {
				if f == field {
					continue
				}
				paths = append(paths, path.MatchRelative().AtParent().AtName(f))
			}
			cs.AddValidator(listvalidator.ExactlyOneOf(paths...), "resources", field)
		}
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

	var app apps_tf.App
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
	waiter, err := w.Apps.Create(ctx, apps.CreateAppRequest{App: &appGoSdk})
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", err.Error())
		return
	}

	// Store the initial version of the app in state
	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waiter.Response, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	appStatus := newApp.AppStatus[0]
	computeStatus := newApp.ComputeStatus[0]
	newApp.AppStatus = nil
	newApp.ComputeStatus = nil
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("app_status"), appStatus)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("compute_status"), computeStatus)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Wait for the app to be created
	finalApp, err := waiter.Get()
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", err.Error())
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
		resp.Diagnostics.AddError("failed to get app", err.Error())
		return
	}

	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	appStatus := newApp.AppStatus[0]
	computeStatus := newApp.ComputeStatus[0]
	newApp.AppStatus = nil
	newApp.ComputeStatus = nil
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("app_status"), appStatus)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("compute_status"), computeStatus)...)
}

func (a *resourceApp) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app apps_tf.App
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
	_, err := w.Apps.Update(ctx, apps.UpdateAppRequest{App: &appGoSdk, Name: app.Name.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("failed to update app", err.Error())
		return
	}

	// Store the updated version of the app in state
	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	appStatus := newApp.AppStatus[0]
	computeStatus := newApp.ComputeStatus[0]
	newApp.AppStatus = nil
	newApp.ComputeStatus = nil
	resp.Diagnostics.Append(resp.State.Set(ctx, newApp)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("app_status"), appStatus)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Empty().AtName("compute_status"), computeStatus)...)
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

var _ resource.ResourceWithConfigure = &resourceApp{}
