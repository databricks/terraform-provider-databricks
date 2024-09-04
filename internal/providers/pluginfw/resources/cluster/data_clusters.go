package cluster

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
	resp.TypeName = "databricks_clusters_pluginframework"
}

func (d *ClustersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: tfschema.DataSourceStructToSchemaMap(ClustersInfo{}, nil),
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
	diags = req.Config.Get(ctx, &clustersInfo)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError("failed to list clusters", err.Error())
		return
	}

	ids := make([]types.String, 0, len(clusters))
	name_contains := strings.ToLower(clustersInfo.ClusterNameContains.ValueString())
	for _, v := range clusters {
		match_name := strings.Contains(strings.ToLower(v.ClusterName), name_contains)
		if name_contains != "" && !match_name {
			continue
		}
		ids = append(ids, types.StringValue(v.ClusterId))
	}
	clustersInfo.Ids = ids
	resp.State.Set(ctx, clustersInfo)
}
