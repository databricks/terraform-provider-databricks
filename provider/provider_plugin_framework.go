package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*DatabricksProviderPluginFramework)(nil)

func GetDatabricksProviderPluginFramework() provider.Provider {
	p := &DatabricksProviderPluginFramework{}
	providerserver.NewProtocol6(p)
	return p
}

type DatabricksProviderPluginFramework struct {
}

func (p *DatabricksProviderPluginFramework) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource {
			return &DatabricksResource{}
		},
	}
}

func (p *DatabricksProviderPluginFramework) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource {
			return &DatabricksDataSource{}
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

// Data Source
type DatabricksDataSource struct {
}

func (d *DatabricksDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
}

func (d *DatabricksDataSource) Read(_ context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}

func (d *DatabricksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
}

// Resource
type DatabricksResource struct {
}

func (r *DatabricksResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
}

func (r *DatabricksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
}

func (r *DatabricksResource) Create(_ context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

func (r *DatabricksResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *DatabricksResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *DatabricksResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
