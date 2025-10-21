package app

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func DataSourceApps() datasource.DataSource {
	return &dataSourceApps{}
}

type dataSourceApps struct {
	client *common.DatabricksClient
}

type dataApps struct {
	Apps types.List `tfsdk:"app"`
	tfschema.Namespace
}

func (dataApps) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (dataApps) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":             reflect.TypeOf(apps_tf.App{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (a dataSourceApps) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceNamePlural)
}

func (a dataSourceApps) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = tfschema.DataSourceStructToSchema(ctx, dataApps{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		return cs
	})
}

func (a *dataSourceApps) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if a.client == nil && req.ProviderData != nil {
		a.client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (a *dataSourceApps) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, resourceName)

	var config dataApps
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !config.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(config.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := a.client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	appsGoSdk, err := w.Apps.ListAll(ctx, apps.ListAppsRequest{})
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", err.Error())
		return
	}

	apps := []attr.Value{}
	for _, appGoSdk := range appsGoSdk {
		app := apps_tf.App{}
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &app)...)
		if resp.Diagnostics.HasError() {
			return
		}
		apps = append(apps, app.ToObjectValue(ctx))
	}
	dataApps := dataApps{
		Apps: types.ListValueMust(apps_tf.App{}.Type(ctx), apps),
		Namespace: tfschema.Namespace{
			ProviderConfig: config.ProviderConfig,
		},
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, dataApps)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

var _ datasource.DataSourceWithConfigure = &dataSourceApp{}
