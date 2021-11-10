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
	WorkspaceID        int64  `json:"workspace_id"`
	MetastoreID        string `json:"metastore_id"`
}

func (a MetastoreAssignmentAPI) createMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Put(a.context, path, ma)
}

func (a MetastoreAssignmentAPI) updateMetastoreAssignment(ma MetastoreAssignment) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%d/metastore", ma.WorkspaceID)
	return a.client.Patch(a.context, path, ma)
}

func (a MetastoreAssignmentAPI) clearMetastoreAssignment(workspaceID string) error {
	path := fmt.Sprintf("/unity-catalog/workspaces/%s/metastore", workspaceID)
	return a.client.Patch(a.context, path, map[string]string{})
}

func ResourceMetastoreAssignment() *schema.Resource {
	s := common.StructToSchema(MetastoreAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["workspace_id"].ForceNew = true // because workspace_id is path param
			return m
		})
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
			// there are no working APIs at the moment to read the assignment
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma MetastoreAssignment
			if err := common.DataToStructPointer(d, s, &ma); err != nil {
				return err
			}
			return NewMetastoreAssignmentAPI(ctx, c).updateMetastoreAssignment(ma)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceID, _, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			return NewMetastoreAssignmentAPI(ctx, c).clearMetastoreAssignment(workspaceID)
		},
	}.ToResource()
}
