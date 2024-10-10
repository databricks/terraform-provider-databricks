package cluster

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "cluster"

func DataSourceCluster() datasource.DataSource {
	return &ClusterDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ClusterDataSource{}

type ClusterDataSource struct {
	Client *common.DatabricksClient
}

type ClusterInfo struct {
	ClusterId   types.String               `tfsdk:"cluster_id" tf:"optional,computed"`
	Name        types.String               `tfsdk:"cluster_name" tf:"optional,computed"`
	ClusterInfo *compute_tf.ClusterDetails `tfsdk:"cluster_info" tf:"optional,computed"`
}

func (d *ClusterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(dataSourceName)
}

func (d *ClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ClusterInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *ClusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func validateClustersList(ctx context.Context, clusters []compute_tf.ClusterDetails, clusterName string) diag.Diagnostics {
	if len(clusters) == 0 {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no cluster with name '%s'", clusterName), "")}
	}
	if len(clusters) > 1 {
		clusterIDs := []string{}
		for _, cluster := range clusters {
			clusterIDs = append(clusterIDs, cluster.ClusterId.ValueString())
		}
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is more than one cluster with name '%s'", clusterName), fmt.Sprintf("The IDs of those clusters are: %s. When specifying a cluster name, the name must be unique. Alternatively, specify the cluster by ID using the cluster_id attribute.", strings.Join(clusterIDs, ", ")))}
	}
	return nil
}

func (d *ClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcommon.SetDataSourceNameInContext(ctx, dataSourceName)
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var clusterInfo ClusterInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &clusterInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	clusterName := clusterInfo.Name.ValueString()
	clusterId := clusterInfo.ClusterId.ValueString()
	if clusterName != "" {
		clustersGoSDk, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
		if err != nil {
			resp.Diagnostics.AddError("failed to list clusters", err.Error())
			return
		}
		var clustersTfSDK []compute_tf.ClusterDetails
		for _, cluster := range clustersGoSDk {
			var clusterDetails compute_tf.ClusterDetails
			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, cluster, &clusterDetails)...)
			if resp.Diagnostics.HasError() {
				return
			}
			clustersTfSDK = append(clustersTfSDK, clusterDetails)
		}
		namedClusters := []compute_tf.ClusterDetails{}
		for _, cluster := range clustersTfSDK {
			if cluster.ClusterName == clusterInfo.Name {
				namedClusters = append(namedClusters, cluster)
			}
		}
		resp.Diagnostics.Append(validateClustersList(ctx, namedClusters, clusterName)...)
		if resp.Diagnostics.HasError() {
			return
		}
		clusterInfo.ClusterInfo = &namedClusters[0]
	} else if clusterId != "" {
		cluster, err := w.Clusters.GetByClusterId(ctx, clusterId)
		if err != nil {
			if apierr.IsMissing(err) {
				resp.State.RemoveResource(ctx)
			}
			resp.Diagnostics.AddError(fmt.Sprintf("failed to get cluster with cluster id: %s", clusterId), err.Error())
			return
		}
		var clusterDetails compute_tf.ClusterDetails
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, cluster, &clusterDetails)...)
		if resp.Diagnostics.HasError() {
			return
		}
		clusterInfo.ClusterInfo = &clusterDetails
	} else {
		resp.Diagnostics.AddError("you need to specify either `cluster_name` or `cluster_id`", "")
		return
	}
	clusterInfo.ClusterId = clusterInfo.ClusterInfo.ClusterId
	clusterInfo.Name = clusterInfo.ClusterInfo.ClusterName
	resp.Diagnostics.Append(resp.State.Set(ctx, clusterInfo)...)
}
