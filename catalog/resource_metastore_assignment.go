package catalog

import (
	"context"
	"fmt"
	"strconv"

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
	WorkspaceID        int64  `json:"workspace_id"`
	MetastoreID        string `json:"metastore_id"`
	DefaultCatalogName string `json:"default_catalog_name"`
}

func (a MetastoreAssignmentAPI) createMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Put(a.context, path, ma)
}

func (a MetastoreAssignmentAPI) updateMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Patch(a.context, path, ma)
}

func ResourceMetastoreAssignment() *schema.Resource {
	s := common.StructToSchema(MetastoreAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	// TODO: fix bugs here... very unstable
	pi := common.NewPairID("workspace_id", "metastore_id")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma MetastoreAssignment
			if err := common.DataToStructPointer(d, s, &ma); err != nil {
				return err
			}
			if err := NewMetastoreAssignmentAPI(ctx, c).createMetastoreAssignment(ma); err != nil {
				return err
			}
			pi.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceID, metastoreID, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			workspaceIDint64, err := strconv.ParseInt(workspaceID, 10, 64)
			if err != nil {
				return err
			}
			metastores, err := NewMetastoresAPI(ctx, c).listMetastores()
			if err != nil {
				return err
			}
			for _, metastore := range metastores {
				if metastore.MetastoreID != metastoreID {
					continue
				}
				for _, wsID := range metastore.WorkspaceIDs {
					if wsID == int64(workspaceIDint64) {
						// workspace has a metastore assigned
						return nil
					}
				}
			}
			// metastore assignment does not exist, delete TF resource
			d.SetId("")
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma MetastoreAssignment
			if err := common.DataToStructPointer(d, s, &ma); err != nil {
				return err
			}
			// TODO: this may be not a correct way to delete assignment
			return NewMetastoreAssignmentAPI(ctx, c).updateMetastoreAssignment(MetastoreAssignment{
				WorkspaceID:        ma.WorkspaceID,
				MetastoreID:        "",
				DefaultCatalogName: "",
			})
		},
	}.ToResource()
}
