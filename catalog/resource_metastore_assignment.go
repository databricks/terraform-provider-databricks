package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMetastoreAssignment() *schema.Resource {
	s := common.StructToSchema(catalog.MetastoreAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["default_catalog_name"].Default = "hive_metastore"
			return m
		})
	pi := common.NewPairID("workspace_id", "metastore_id").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId := int64(d.Get("workspace_id").(int))
			metastoreId := d.Get("metastore_id").(string)
			var create catalog.CreateMetastoreAssignment
			common.DataToStructPointer(d, s, &create)
			create.WorkspaceId = workspaceId
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				_, err = acc.MetastoreAssignments.Create(ctx,
					catalog.AccountsCreateMetastoreAssignment{
						WorkspaceId:         workspaceId,
						MetastoreId:         metastoreId,
						MetastoreAssignment: &create,
					})
				if err != nil {
					return err
				}
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				err = w.Metastores.Assign(ctx, create)
				if err != nil {
					return err
				}
			}
			pi.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			first, _, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			workspaceId, err := strconv.ParseInt(first, 10, 64)
			if err != nil {
				return err
			}
			// calling account-level API if using account client
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				ma, err := acc.MetastoreAssignments.GetByWorkspaceId(ctx, workspaceId)
				if err != nil {
					return err
				}
				return common.StructToData(ma, s, d)
			} else {
				//calling workspace-level API if using workspace client
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				ma, err := w.Metastores.Current(ctx)
				if err != nil {
					return err
				}
				return common.StructToData(ma, s, d)
			}
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId := int64(d.Get("workspace_id").(int))
			metastoreId := d.Get("metastore_id").(string)
			var update catalog.UpdateMetastoreAssignment
			common.DataToStructPointer(d, s, &update)
			update.WorkspaceId = workspaceId
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				err = acc.MetastoreAssignments.Update(ctx,
					catalog.AccountsUpdateMetastoreAssignment{
						WorkspaceId:         workspaceId,
						MetastoreId:         metastoreId,
						MetastoreAssignment: &update,
					})
				if err != nil {
					return err
				}
				return nil
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				return w.Metastores.UpdateAssignment(ctx, update)
			}
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			first, metastoreId, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			workspaceId, err := strconv.ParseInt(first, 10, 64)
			if err != nil {
				return err
			}
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				return acc.MetastoreAssignments.DeleteByWorkspaceIdAndMetastoreId(ctx, workspaceId, metastoreId)
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				return w.Metastores.UnassignByWorkspaceId(ctx, workspaceId)
			}
		},
	}.ToResource()
}
