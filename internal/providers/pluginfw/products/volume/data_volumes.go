package volume

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "volumes"

func DataSourceVolumes() datasource.DataSource {
	return &VolumesDataSource{}
}

var _ datasource.DataSourceWithConfigure = &VolumesDataSource{}

type VolumesDataSource struct {
	Client *common.DatabricksClient
}

type VolumesList struct {
	CatalogName types.String   `tfsdk:"catalog_name"`
	SchemaName  types.String   `tfsdk:"schema_name"`
	Ids         []types.String `tfsdk:"ids" tf:"optional,computed"`
}

func (d *VolumesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *VolumesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(VolumesList{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *VolumesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *VolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var volumesList VolumesList
	diags = req.Config.Get(ctx, &volumesList)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var listVolumesRequest catalog.ListVolumesRequest
	converters.TfSdkToGoSdkStruct(ctx, volumesList, &listVolumesRequest)
	volumes, err := w.Volumes.ListAll(ctx, listVolumesRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get volumes for the catalog:%s and schema%s", listVolumesRequest.CatalogName, listVolumesRequest.SchemaName), err.Error())
		return
	}
	for _, v := range volumes {
		volumesList.Ids = append(volumesList.Ids, types.StringValue(v.FullName))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, volumesList)...)
}
