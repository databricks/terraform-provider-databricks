package clusters

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go"
)

func defaultSmallestNodeType(w *databricks.WorkspaceClient, request clusters.NodeTypeRequest) string {
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

func smallestNodeType(ctx context.Context, request clusters.NodeTypeRequest, w *databricks.WorkspaceClient) string {
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

func (a ClustersAPI) GetSmallestNodeType(request clusters.NodeTypeRequest) string {
	w, _ := a.client.WorkspaceClient()
	return smallestNodeType(a.context, request, w)
}

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() *schema.Resource {
	s := common.StructToSchema(clusters.NodeTypeRequest{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
			var this clusters.NodeTypeRequest
			common.DataToStructPointer(d, s, &this)
			client := m.(*common.DatabricksClient)
			w, err := client.WorkspaceClient()
			if err != nil {
				err = fmt.Errorf("cannot read data databricks_node_type: %w", err)
				return diag.FromErr(err)
			}
			d.SetId(smallestNodeType(ctx, this, w))
			return nil
		},
	}
}
