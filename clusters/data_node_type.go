package clusters

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go"
)

type NodeTypeRequest struct {
	compute.NodeTypeRequest
	Arm bool `json:"arm,omitempty"`
}

func defaultSmallestNodeType(w *databricks.WorkspaceClient, request NodeTypeRequest) string {
	if request.Arm || request.Graviton {
		if w.Config.IsAws() {
			if request.Fleet {
				return "rgd-fleet.xlarge"
			}
			return "m6g.xlarge"
		} else if w.Config.IsAzure() {
			return "Standard_D4pds_v6"
		}
	}
	if w.Config.IsAzure() {
		return "Standard_D4ds_v5"
	} else if w.Config.IsGcp() {
		return "n1-standard-4"
	}
	if request.Fleet {
		return "md-fleet.xlarge"
	}
	return "i3.xlarge"
}

func smallestNodeType(ctx context.Context, request NodeTypeRequest, w *databricks.WorkspaceClient) string {
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	if err != nil {
		return defaultSmallestNodeType(w, request)
	}
	// if arm is true, then graviton is true
	request.Graviton = request.Arm || request.Graviton
	nodeType, err := nodeTypes.Smallest(request.NodeTypeRequest)
	if err != nil {
		nodeType = defaultSmallestNodeType(w, request)
	}
	return nodeType
}

func (a ClustersAPI) GetSmallestNodeType(request NodeTypeRequest) string {
	w, _ := a.client.WorkspaceClient()
	return smallestNodeType(a.context, request, w)
}

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() common.Resource {
	return common.WorkspaceDataWithUnifiedProviderWithCustomizeFunc(func(ctx context.Context, data *NodeTypeRequest, w *databricks.WorkspaceClient) error {
		data.Id = smallestNodeType(ctx, *data, w)
		log.Printf("[DEBUG] smallest node: %s", data.Id)
		return nil
	}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "graviton").SetDeprecated("Use `arm` instead")
		return s
	})
}
