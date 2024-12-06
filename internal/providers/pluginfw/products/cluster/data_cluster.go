package cluster

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	ClusterId   types.String `tfsdk:"cluster_id" tf:"optional,computed"`
	Name        types.String `tfsdk:"cluster_name" tf:"optional,computed"`
	ClusterInfo types.List   `tfsdk:"cluster_info" tf:"optional,computed"`
}

func (ClusterInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_info": reflect.TypeOf(compute_tf.ClusterDetails{}),
	}
}

func (d *ClusterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(dataSourceName)
}

func (d *ClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ClusterInfo{}, nil)
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

func (d *ClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)
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
	cluster, diag := d.getClusterDetails(ctx, w, clusterName, clusterId)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfCluster compute_tf.ClusterDetails
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, cluster, &tfCluster)...)

	clusterInfo.ClusterId = tfCluster.ClusterId
	clusterInfo.Name = tfCluster.ClusterName
	clusterInfo.ClusterInfo = types.ListValueMust(tfCluster.Type(ctx), []attr.Value{tfCluster.ToObjectValue(ctx)})
	resp.Diagnostics.Append(resp.State.Set(ctx, clusterInfo)...)
}

func validateClustersList(_ context.Context, clusters []compute.ClusterDetails, clusterName string) diag.Diagnostics {
	if len(clusters) == 0 {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no cluster with name '%s'", clusterName), "")}
	}
	if len(clusters) > 1 {
		clusterIDs := []string{}
		for _, cluster := range clusters {
			clusterIDs = append(clusterIDs, cluster.ClusterId)
		}
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is more than one cluster with name '%s'", clusterName), fmt.Sprintf("The IDs of those clusters are: %s. When specifying a cluster name, the name must be unique. Alternatively, specify the cluster by ID using the cluster_id attribute.", strings.Join(clusterIDs, ", ")))}
	}
	return nil
}

func (d *ClusterDataSource) getClusterDetails(ctx context.Context, w *databricks.WorkspaceClient, clusterName, clusterId string) (c compute.ClusterDetails, dd diag.Diagnostics) {
	if clusterName != "" {
		clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
		if err != nil {
			dd.AddError("failed to list clusters", err.Error())
			return
		}
		cc := []compute.ClusterDetails{}
		for _, cluster := range clusters {
			if cluster.ClusterName == clusterName {
				cc = append(cc, cluster)
			}
		}
		dd.Append(validateClustersList(ctx, cc, clusterName)...)
		if dd.HasError() {
			return
		}
		return cc[0], dd
	}
	if clusterId != "" {
		cluster, err := w.Clusters.GetByClusterId(ctx, clusterId)
		if err != nil {
			dd.AddError(fmt.Sprintf("failed to get cluster with cluster id: %s", clusterId), err.Error())
			return
		}
		return *cluster, dd
	}

	dd.AddError("you need to specify either `cluster_name` or `cluster_id`", "")
	return
}
