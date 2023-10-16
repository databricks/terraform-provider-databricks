package catalog

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var metastoreAssignmentSchema = common.StructToSchema(catalog.MetastoreAssignment{},
	func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["default_catalog_name"].Default = "hive_metastore"
		return m
	})

func ResourceMetastoreAssignment() *schema.Resource {
	pi := common.NewPairID("workspace_id", "metastore_id").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return metastoreAssignmentSchema
		})
	return common.Resource{
		Schema:        metastoreAssignmentSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 0,
				Type:    metastoreAssignmentSchemaV0(),
				Upgrade: metastoreAssignmentMigrateV0,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId := int64(d.Get("workspace_id").(int))
			metastoreId := d.Get("metastore_id").(string)
			var create catalog.CreateMetastoreAssignment
			common.DataToStructPointer(d, metastoreAssignmentSchema, &create)
			create.WorkspaceId = workspaceId

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				err := acc.MetastoreAssignments.Create(ctx,
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
				return common.StructToData(ma, metastoreAssignmentSchema, d)
			}, func(w *databricks.WorkspaceClient) error {
				//this only works when managing the metastore assigned to the current workspace.
				//plus we don't know the workspace we're logged into.
				ma, err := w.Metastores.Current(ctx)
				if err != nil {
					return err
				}
				d.Set("metastore_id", ma.MetastoreId)
				return nil
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId := int64(d.Get("workspace_id").(int))
			metastoreId := d.Get("metastore_id").(string)
			var update catalog.UpdateMetastoreAssignment
			common.DataToStructPointer(d, metastoreAssignmentSchema, &update)
			update.WorkspaceId = workspaceId

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				return acc.MetastoreAssignments.Update(ctx,
					catalog.AccountsUpdateMetastoreAssignment{
						WorkspaceId:         workspaceId,
						MetastoreId:         metastoreId,
						MetastoreAssignment: &update,
					})
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
				return acc.MetastoreAssignments.DeleteByWorkspaceIdAndMetastoreId(ctx, workspaceId, metastoreId)
			}, func(w *databricks.WorkspaceClient) error {
				return w.Metastores.Unassign(ctx, catalog.UnassignRequest{
					MetastoreId: metastoreId,
					WorkspaceId: workspaceId,
				})
			})
		},
	}.ToResource()
}

// migrate to v1 state, as the id is now changed
func metastoreAssignmentMigrateV0(ctx context.Context,
	rawState map[string]any,
	meta any) (map[string]any, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "id":
			log.Printf("[INFO] Upgrade id")
			ids := v.(string)
			if strings.Contains(ids, "/") {
				//only upgrade id for previous ones with slash
				splitIds := strings.Split(ids, "/")
				newState[k] = fmt.Sprintf("%v|%v", splitIds[1], splitIds[0])
			} else {
				newState[k] = v
			}
			continue
		default:
			newState[k] = v
		}
	}
	return newState, nil
}

func metastoreAssignmentSchemaV0() cty.Type {
	return (&schema.Resource{
		Schema: metastoreAssignmentSchema}).CoreConfigSchema().ImpliedType()
}
