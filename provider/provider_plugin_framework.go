package provider

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/pluginframework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func GetDatabricksProviderPluginFramework() provider.Provider {
	p := &DatabricksProviderPluginFramework{}
	return p
}

type DatabricksProviderPluginFramework struct {
}

var _ provider.Provider = (*DatabricksProviderPluginFramework)(nil)

func (p *DatabricksProviderPluginFramework) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		pluginframework.ResourceLakehouseMonitorPluginFramework,
	}
}

func (p *DatabricksProviderPluginFramework) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource {
			return &pluginframework.DatabricksDataSource{}
		},
	}
}

func (p *DatabricksProviderPluginFramework) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *DatabricksProviderPluginFramework) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "databricks-terraform-provider-plugin-framework"
	resp.Version = "0.0.1"
}

func (p *DatabricksProviderPluginFramework) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}
