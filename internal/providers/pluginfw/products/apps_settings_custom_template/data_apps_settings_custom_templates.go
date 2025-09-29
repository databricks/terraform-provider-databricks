// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps_settings_custom_template

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "apps_settings_custom_templates"

var _ datasource.DataSourceWithConfigure = &CustomTemplatesDataSource{}

func DataSourceCustomTemplates() datasource.DataSource {
	return &CustomTemplatesDataSource{}
}

// CustomTemplatesData extends the main model with additional fields.
type CustomTemplatesData struct {
	AppsSettings types.List `tfsdk:"templates"`
}

func (CustomTemplatesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"templates": reflect.TypeOf(apps_tf.CustomTemplate{}),
	}
}

type CustomTemplatesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CustomTemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CustomTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CustomTemplatesData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("templates")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CustomTemplate",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CustomTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CustomTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CustomTemplatesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest apps.ListCustomTemplatesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.AppsSettings.ListCustomTemplatesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list apps_settings_custom_templates", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var custom_template apps_tf.CustomTemplate
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &custom_template)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, custom_template.ToObjectValue(ctx))
	}

	var newState CustomTemplatesData
	newState.AppsSettings = types.ListValueMust(apps_tf.CustomTemplate{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
