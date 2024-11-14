package cluster

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceNameClusters = "clusters"

func DataSourceClusters() datasource.DataSource {
	return &ClustersDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ClustersDataSource{}

type ClustersDataSource struct {
	Client *common.DatabricksClient
}

type ClustersInfo struct {
	Ids                 []types.String `tfsdk:"ids" tf:"optional,computed"`
	ClusterNameContains types.String   `tfsdk:"cluster_name_contains" tf:"optional"`
}

func (d *ClustersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(dataSourceNameClusters)
}

func (d *ClustersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ClustersInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *ClustersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *ClustersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var clustersInfo ClustersInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &clustersInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
	if err != nil {
		resp.Diagnostics.AddError("failed to list clusters", err.Error())
		return
	}

	ids := make([]types.String, 0, len(clusters))
	nameContains := strings.ToLower(clustersInfo.ClusterNameContains.ValueString())
	for _, v := range clusters {
		matchName := strings.Contains(strings.ToLower(v.ClusterName), nameContains)
		if matchName {
			ids = append(ids, types.StringValue(v.ClusterId))
		}
	}
	clustersInfo.Ids = ids
	resp.Diagnostics.Append(resp.State.Set(ctx, clustersInfo)...)
}
