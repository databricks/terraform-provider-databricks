package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMetastoreAssignment() common.Resource {
	s := common.StructToSchema(catalog.MetastoreAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["default_catalog_name"].Computed = true
			m["default_catalog_name"].Deprecated = "Use databricks_default_namespace_setting resource instead"
			m["workspace_id"].ForceNew = true
			m["metastore_id"].ForceNew = true
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

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				_, err := acc.MetastoreAssignments.Create(ctx,
					catalog.AccountsCreateMetastoreAssignment{
						WorkspaceId:         workspaceId,
						MetastoreId:         metastoreId,
						MetastoreAssignment: &create,
					})
				if err != nil {
					return err
				}
				pi.Pack(d)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				err := w.Metastores.Assign(ctx, create)
				if err != nil {
					return err
				}
				pi.Pack(d)
				return nil
			})
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

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				ma, err := acc.MetastoreAssignments.GetByWorkspaceId(ctx, workspaceId)
				if err != nil {
					return err
				}
				return common.StructToData(ma, s, d)
			}, func(w *databricks.WorkspaceClient) error {
				//this only works when managing the metastore assigned to the current workspace.
				//plus we don't know the workspace we're logged into.
				ma, err := w.Metastores.Current(ctx)
				if err != nil {
					return err
				}
				d.Set("metastore_id", ma.MetastoreId)
				d.Set("default_catalog_name", ma.DefaultCatalogName)
				d.Set("workspace_id", workspaceId)
				return nil
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId := int64(d.Get("workspace_id").(int))
			metastoreId := d.Get("metastore_id").(string)
			var update catalog.UpdateMetastoreAssignment
			common.DataToStructPointer(d, s, &update)
			update.WorkspaceId = workspaceId

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				_, err := acc.MetastoreAssignments.Update(ctx,
					catalog.AccountsUpdateMetastoreAssignment{
						WorkspaceId:         workspaceId,
						MetastoreId:         metastoreId,
						MetastoreAssignment: &update,
					})
				return err
			}, func(w *databricks.WorkspaceClient) error {
				return w.Metastores.UpdateAssignment(ctx, update)
			})
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
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				_, err := acc.MetastoreAssignments.DeleteByWorkspaceIdAndMetastoreId(ctx, workspaceId, metastoreId)
				return err
			}, func(w *databricks.WorkspaceClient) error {
				return w.Metastores.Unassign(ctx, catalog.UnassignRequest{
					MetastoreId: metastoreId,
					WorkspaceId: workspaceId,
				})
			})
		},
	}
}
