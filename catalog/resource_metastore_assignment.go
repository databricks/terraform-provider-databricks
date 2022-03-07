package catalog

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MetastoreAssignmentAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewMetastoreAssignmentAPI(ctx context.Context, m interface{}) MetastoreAssignmentAPI {
	return MetastoreAssignmentAPI{m.(*common.DatabricksClient), ctx}
}

type MetastoreAssignment struct {
	WorkspaceID        int64  `json:"workspace_id" tf:"force_new"`
	MetastoreID        string `json:"metastore_id"`
	DefaultCatalogName string `json:"default_catalog_name,omitempty" tf:"default:hive_metastore"`
}

func (a MetastoreAssignmentAPI) createMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Put(a.context, path, ma)
}

func (a MetastoreAssignmentAPI) updateMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Patch(a.context, path, ma)
}

func (a MetastoreAssignmentAPI) getAssignedMetastoreID() (string, error) {
	var ma MetastoreAssignment
	err := a.client.Get(a.context, "/unity-catalog/metastore_summary", nil, &ma)
	return ma.MetastoreID, err
}

func (a MetastoreAssignmentAPI) deleteMetastoreAssignment(workspaceID, metastoreID string) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%s/metastore", workspaceID)
	return a.client.Delete(a.context, path, map[string]string{
		"metastore_id": metastoreID,
	})
}

func ResourceMetastoreAssignment() *schema.Resource {
	s := common.StructToSchema(MetastoreAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	pi := common.NewPairID("workspace_id", "metastore_id")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma MetastoreAssignment
			common.DataToStructPointer(d, s, &ma)
			if err := NewMetastoreAssignmentAPI(ctx, c).createMetastoreAssignment(ma); err != nil {
				return err
			}
			pi.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreID, err := NewMetastoreAssignmentAPI(ctx, c).getAssignedMetastoreID()
			d.Set("metastore_id", metastoreID)
			return err
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma MetastoreAssignment
			common.DataToStructPointer(d, s, &ma)
			return NewMetastoreAssignmentAPI(ctx, c).updateMetastoreAssignment(ma)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceID, metastoreID, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			return NewMetastoreAssignmentAPI(ctx, c).deleteMetastoreAssignment(workspaceID, metastoreID)
		},
	}.ToResource()
}
