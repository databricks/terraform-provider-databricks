package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMetastoreAssignment() *schema.Resource {
	s := common.StructToSchema(unitycatalog.MetastoreAssignment{},
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
			var create unitycatalog.CreateMetastoreAssignment
			common.DataToStructPointer(d, s, &create)
			s := d.Get("workspace_id").(string)
			workspaceId, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				return err
			}
			create.WorkspaceId = workspaceId

			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				_, err = acc.AccountMetastoreAssignments.Create(ctx, create)
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
			workspaceId, err := strconv.ParseInt(first, 10, 0)
			if err != nil {
				return err
			}
			var ma *unitycatalog.MetastoreAssignment
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				ma, err = acc.AccountMetastoreAssignments.GetByWorkspaceId(ctx, workspaceId)
				if err != nil {
					return err
				}
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				ma, err = w.Metastores.Current(ctx)
				if err != nil {
					return err
				}
			}
			return common.StructToData(ma, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ma unitycatalog.UpdateMetastoreAssignment
			common.DataToStructPointer(d, s, &ma)
			s := d.Get("workspace_id").(string)
			workspaceId, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				return err
			}
			ma.WorkspaceId = workspaceId

			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				_, err = acc.AccountMetastoreAssignments.Update(ctx, ma)
				if err != nil {
					return err
				}
				return nil
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				return w.Metastores.UpdateAssignment(ctx, ma)
			}
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			first, metastoreId, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			workspaceId, err := strconv.ParseInt(first, 10, 0)
			if err != nil {
				return err
			}
			if c.Config.IsAccountClient() && c.Config.AccountID != "" {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				return acc.AccountMetastoreAssignments.DeleteByWorkspaceIdAndMetastoreId(ctx, workspaceId, metastoreId)
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
