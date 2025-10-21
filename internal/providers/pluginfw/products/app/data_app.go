package app

import (
	"context"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func DataSourceApp() datasource.DataSource {
	return &dataSourceApp{}
}

type dataSourceApp struct {
	client *common.DatabricksClient
}

type dataApp struct {
	Name types.String `tfsdk:"name"`
	App  types.Object `tfsdk:"app"`
	tfschema.Namespace
}

func (dataApp) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["app"] = attrs["app"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (dataApp) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":             reflect.TypeOf(apps_tf.App{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (a dataSourceApp) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (a dataSourceApp) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = tfschema.DataSourceStructToSchema(ctx, dataApp{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		return cs
	})
}

func (a *dataSourceApp) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if a.client == nil && req.ProviderData != nil {
		a.client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (a *dataSourceApp) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, resourceName)

	var config dataApp
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

	appGoSdk, err := w.Apps.GetByName(ctx, config.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", err.Error())
		return
	}

	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	dataApp := dataApp{
		Name: config.Name,
		App:  newApp.ToObjectValue(ctx),
		Namespace: tfschema.Namespace{
			ProviderConfig: config.ProviderConfig,
		},
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, dataApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

var _ datasource.DataSourceWithConfigure = &dataSourceApp{}
