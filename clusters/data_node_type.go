package clusters

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NodeTypeRequest struct {
	common.Namespace
	compute.NodeTypeRequest
	Arm bool `json:"arm,omitempty"`
}

// isAws detects AWS by checking if any node_type_id contains a "." followed by "large" (e.g. "i3.xlarge").
func isAws(nodeTypes *compute.ListNodeTypesResponse) bool {
	for _, nt := range nodeTypes.NodeTypes {
		dotIdx := strings.Index(nt.NodeTypeId, ".")
		if dotIdx >= 0 && strings.Contains(nt.NodeTypeId[dotIdx:], "large") {
			return true
		}
	}
	return false
}

// isAzure detects Azure by checking if any node_type_id contains "Standard_" (e.g. "Standard_D4ds_v5").
func isAzure(nodeTypes *compute.ListNodeTypesResponse) bool {
	for _, nt := range nodeTypes.NodeTypes {
		if strings.Contains(nt.NodeTypeId, "Standard_") {
			return true
		}
	}
	return false
}

// isGcp detects GCP as a fallback when the node types match neither AWS nor Azure patterns.
func isGcp(nodeTypes *compute.ListNodeTypesResponse) bool {
	return !isAws(nodeTypes) && !isAzure(nodeTypes)
}

func defaultSmallestNodeType(nodeTypes *compute.ListNodeTypesResponse, request NodeTypeRequest) string {
	if request.Arm || request.Graviton {
		if isAws(nodeTypes) {
			if request.Fleet {
				return "rgd-fleet.xlarge"
			}
			return "m6g.xlarge"
		} else if isAzure(nodeTypes) {
			return "Standard_D4pds_v6"
		}
	}
	if isAzure(nodeTypes) {
		return "Standard_D4ds_v5"
	} else if isGcp(nodeTypes) {
		return "n1-standard-4"
	}
	if request.Fleet {
		return "md-fleet.xlarge"
	}
	return "i3.xlarge"
}

func smallestNodeType(ctx context.Context, request NodeTypeRequest, w *databricks.WorkspaceClient) (string, error) {
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	if err != nil {
		return "", fmt.Errorf("cannot determine smallest node type: %w", err)
	}
	// if arm is true, then graviton is true
	request.Graviton = request.Arm || request.Graviton
	nodeType, err := nodeTypes.Smallest(request.NodeTypeRequest)
	if err != nil {
		nodeType = defaultSmallestNodeType(nodeTypes, request)
	}
	return nodeType, nil
}

func (a ClustersAPI) GetSmallestNodeType(request NodeTypeRequest) (string, error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return "", fmt.Errorf("cannot get workspace client: %w", err)
	}
	return smallestNodeType(a.context, request, w)
}

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() common.Resource {
	return common.WorkspaceDataWithCustomizeFunc(func(ctx context.Context, data *NodeTypeRequest, w *databricks.WorkspaceClient) error {
		nodeType, err := smallestNodeType(ctx, *data, w)
		if err != nil {
			return err
		}
		data.Id = nodeType
		log.Printf("[DEBUG] smallest node: %s", data.Id)
		return nil
	}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "graviton").SetDeprecated("Use `arm` instead")
		common.NamespaceCustomizeSchemaMap(s)
		return s
	})
}
