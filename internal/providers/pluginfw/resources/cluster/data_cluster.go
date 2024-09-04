package cluster

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceCluster() datasource.DataSource {
	return &ClusterDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ClusterDataSource{}

type ClusterDataSource struct {
	Client *common.DatabricksClient
}

type ClusterInfo struct {
	Id          types.String               `tfsdk:"id" tf:"optional,computed"`
	ClusterId   types.String               `tfsdk:"cluster_id" tf:"optional,computed"`
	Name        types.String               `tfsdk:"cluster_name" tf:"optional,computed"`
	ClusterInfo *compute_tf.ClusterDetails `tfsdk:"cluster_info" tf:"optional,computed"`
}

func (d *ClusterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_cluster_pluginframework"
}

func (d *ClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: tfschema.DataSourceStructToSchemaMap(ClusterInfo{}, nil),
	}
}

func (d *ClusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *ClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var clusterInfo ClusterInfo
	diags = req.Config.Get(ctx, &clusterInfo)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if clusterInfo.Name.ValueString() != "" {
		clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
		if err != nil {
			resp.Diagnostics.AddError("failed to list clusters", err.Error())
			return
		}
		var clustersTfSDK []compute_tf.ClusterDetails
		converters.GoSdkToTfSdkStruct(ctx, clusters, clustersTfSDK)
		namedClusters := []compute_tf.ClusterDetails{}
		for _, clst := range clustersTfSDK {
			cluster := clst
			if cluster.ClusterName == clusterInfo.Name {
				namedClusters = append(namedClusters, cluster)
			}
		}
		if len(namedClusters) == 0 {
			resp.Diagnostics.AddError(fmt.Sprintf("there is no cluster with name '%s'", clusterInfo.Name.ValueString()), "")
			return
		}
		if len(namedClusters) > 1 {
			resp.Diagnostics.AddError(fmt.Sprintf("there is more than one cluster with name '%s'", clusterInfo.Name.ValueString()), "")
			return
		}
		clusterInfo.ClusterInfo = &namedClusters[0]
		resp.State.Set(ctx, clusterInfo)
	}
}
