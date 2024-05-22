package pluginframework

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceVolumesPluginFramework() datasource.DataSource {
	return &VolumesDataSource{}
}

type VolumesDataSource struct{}

type VolumesList struct {
	CatalogName string   `json:"catalog_name"`
	SchemaName  string   `json:"schema_name"`
	Ids         []string `json:"ids,omitempty"`
}

func (d *VolumesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumes_plugin_framework"
}

func (d *VolumesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			},
		},
	}
}

func (d *VolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var volumeInfo catalog.VolumeInfo
	diags := req.Config.Get(ctx, &volumeInfo)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	volumes, err := w.Volumes.ListAll(ctx, catalog.ListVolumesRequest{
		CatalogName: volumeInfo.CatalogName,
		SchemaName:  volumeInfo.SchemaName,
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get volumes for the catalog and schema", err.Error())
		return
	}
	volumeList := VolumesList{}
	for _, v := range volumes {
		volumeList.Ids = append(volumeList.Ids, v.FullName)
	}
	resp.State.Set(ctx, volumeList)
}
