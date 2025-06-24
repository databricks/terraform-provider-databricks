package clusters

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/databricks-sdk-go"
)

func defaultSmallestNodeType(w *databricks.WorkspaceClient, request compute.NodeTypeRequest) string {
	if w.Config.IsAzure() {
		return "Standard_D3_v2"
	} else if w.Config.IsGcp() {
		return "n1-standard-4"
	}
	if request.Fleet {
		return "md-fleet.xlarge"
	}
	return "i3.xlarge"
}

func smallestNodeType(ctx context.Context, request compute.NodeTypeRequest, w *databricks.WorkspaceClient) string {
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	if err != nil {
		return defaultSmallestNodeType(w, request)
	}
	nodeType, err := nodeTypes.Smallest(request)
	if err != nil {
		nodeType = defaultSmallestNodeType(w, request)
	}
	return nodeType
}

func (a ClustersAPI) GetSmallestNodeType(request compute.NodeTypeRequest) string {
	w, _ := a.client.WorkspaceClient()
	return smallestNodeType(a.context, request, w)
}

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *compute.NodeTypeRequest, w *databricks.WorkspaceClient) error {
		data.Id = smallestNodeType(ctx, *data, w)
		log.Printf("[DEBUG] smallest node: %s", data.Id)
		return nil
	})
}
