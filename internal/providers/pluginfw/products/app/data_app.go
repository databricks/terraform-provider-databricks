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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
}

func (dataApp) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["app"] = attrs["app"].SetComputed()

	return attrs
}

func (dataApp) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(apps_tf.App{}),
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
	w, diags := a.client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var name types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("name"), &name)...)
	if resp.Diagnostics.HasError() {
		return
	}

	appGoSdk, err := w.Apps.GetByName(ctx, name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", err.Error())
		return
	}

	var newApp apps_tf.App
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, appGoSdk, &newApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	dataApp := dataApp{Name: name, App: newApp.ToObjectValue(ctx)}
	resp.Diagnostics.Append(resp.State.Set(ctx, dataApp)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

var _ datasource.DataSourceWithConfigure = &dataSourceApp{}
