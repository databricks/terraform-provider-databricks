package pluginframework

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceVolumes() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &VolumesDataSource{}
	}
}

var _ datasource.DataSource = &VolumesDataSource{}

type VolumesDataSource struct {
	Client *common.DatabricksClient
}

type VolumesList struct {
	CatalogName types.String   `tfsdk:"catalog_name"`
	SchemaName  types.String   `tfsdk:"schema_name"`
	Ids         []types.String `tfsdk:"ids" tf:"optional"`
}

func (d *VolumesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_volumes_pluginframework"
}

func (d *VolumesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// TODO: Use StructToSchemaMap to generate the schema once it supports schema for data sources
		Attributes: map[string]schema.Attribute{
			"catalog_name": schema.StringAttribute{
				Required: true,
			},
			"schema_name": schema.StringAttribute{
				Required: true,
			},
			"ids": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (d *VolumesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.Client = client
}

func (d *VolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	client := d.Client
	w, err := client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var volumesList VolumesList
	diags := req.Config.Get(ctx, &volumesList)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	volumes, err := w.Volumes.ListAll(ctx, catalog.ListVolumesRequest{
		CatalogName: volumesList.CatalogName.ValueString(),
		SchemaName:  volumesList.SchemaName.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get volumes for the catalog and schema", err.Error())
		return
	}
	for _, v := range volumes {
		volumesList.Ids = append(volumesList.Ids, types.StringValue(v.FullName))
	}
	resp.State.Set(ctx, volumesList)
}
