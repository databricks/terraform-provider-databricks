package app

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const (
	resourceName       = "app"
	resourceNamePlural = "apps"
)

type MetadataAttributes struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r MetadataAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	return attrs
}

func (r MetadataAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r MetadataAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)

}
func (r MetadataAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

type appResource struct {
	apps_tf.App
	NoCompute          types.Bool   `tfsdk:"no_compute"`
	MetadataAttributes types.Object `tfsdk:"metadata_attributes"`
}

func (a appResource) ApplySchemaCustomizations(s map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	s["no_compute"] = s["no_compute"].SetOptional()
	s = apps_tf.App{}.ApplySchemaCustomizations(s)
	return s
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
		exclusiveFields := []string{}
		t := reflect.TypeOf(apps_tf.AppResource{})
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Tag != "" {
				tag := f.Tag.Get("tfsdk")
				if !slices.Contains([]string{"name", "description"}, tag) {
					exclusiveFields = append(exclusiveFields, tag)
				}
			}
		}
		paths := path.Expressions{}
		for _, field := range exclusiveFields[1:] {
			paths = append(paths, path.MatchRelative().AtParent().AtName(field))
		}
		cs.AddValidator(objectvalidator.ExactlyOneOf(paths...), "resources", exclusiveFields[0])
		for _, field := range []string{
			"create_time",
			"creator",
			"service_principal_client_id",
			"service_principal_name",
			"url",
		} {
			cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), field)
		}
		cs.AddPlanModifier(int64planmodifier.UseStateForUnknown(), "service_principal_id")
		cs.SetOptional("metadata_attributes")
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
		App:             appGoSdk,
		NoCompute:       app.NoCompute.ValueBool(),
		ForceSendFields: forceSendFields,
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

	// Wait for the app to be created. If no_compute is specified, the terminal state is
	// STOPPED, otherwise it is ACTIVE.
	finalApp, err := a.waitForApp(ctx, w, appGoSdk.Name)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for app to be active or stopped", err.Error())
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

// This is copied from the retries package of the databricks-sdk-go. It should be made public,
// but for now, I'm copying it here.
func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	e := err.(*retries.Err)
	if e == nil {
		return false
	}
	return !e.Halt
}

// waitForApp waits for the app to reach the target state. The target state is either ACTIVE or STOPPED.
// Apps with no_compute set to true will reach the STOPPED state, otherwise they will reach the ACTIVE state.
func (a *resourceApp) waitForApp(ctx context.Context, w *databricks.WorkspaceClient, name string) (*apps.App, error) {
	retrier := retries.New[apps.App](retries.WithTimeout(-1), retries.WithRetryFunc(shouldRetry))
	return retrier.Run(ctx, func(ctx context.Context) (*apps.App, error) {
		app, err := w.Apps.GetByName(ctx, name)
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := app.ComputeStatus.State
		statusMessage := app.ComputeStatus.Message
		switch status {
		case apps.ComputeStateActive, apps.ComputeStateStopped:
			return app, nil
		case apps.ComputeStateError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				apps.ComputeStateActive, apps.ComputeStateStopped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

func (a *resourceApp) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var app appResource
	resp.Diagnostics.Append(req.State.Get(ctx, &app)...)
	if resp.Diagnostics.HasError() {
		return
	}

	appGoSdk, err := w.Apps.GetByName(ctx, app.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", err.Error())
		return
	}

	var newApp appResource
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	newApp.NoCompute = app.NoCompute
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
	response, err := w.Apps.Update(ctx, apps.UpdateAppRequest{App: appGoSdk, Name: app.Name.ValueString()})
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
	// Modifying no_compute after creation has no effect.
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

	var app appResource
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
