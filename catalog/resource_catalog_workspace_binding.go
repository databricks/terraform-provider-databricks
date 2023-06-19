package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCatalogWorkspaceBinding() *schema.Resource {
	return common.NewPairID("catalog_name", "workspace_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			_, err = w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
				Name:             catalogName,
				AssignWorkspaces: []int64{i64WorkspaceId},
			})
			return err
		},
		ReadContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			bindings, err := w.WorkspaceBindings.GetByName(ctx, catalogName)
			if err != nil {
				return err
			}
			if !contains(bindings.Workspaces, i64WorkspaceId) {
				return apierr.NotFound("Catalog has no binding to this workspace")
			}
			return nil
		},
		DeleteContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			_, err = w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
				Name:               catalogName,
				UnassignWorkspaces: []int64{i64WorkspaceId},
			})
			return err
		},
	})
}
