package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func contains_int64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ResourceCatalogWorkspaceBinding() *schema.Resource {
	return common.NewPairID("catalog_name", "workspace_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			createBindingRequest := catalog.UpdateWorkspaceBindings{
				Name:             catalogName,
				AssignWorkspaces: []int64{i64WorkspaceId},
			}
			_, err = catalog.NewWorkspaceBindings(c.DatabricksClient).Update(ctx, createBindingRequest)
			return err
		},
		ReadContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			bindings, err := catalog.NewWorkspaceBindings(c.DatabricksClient).GetByName(ctx, catalogName)
			if err != nil {
				return err
			}
			if !contains_int64(bindings.Workspaces, i64WorkspaceId) {
				return apierr.NotFound("Catalog has no binding to this workspace")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, catalogName, workspaceId string, c *common.DatabricksClient) error {
			i64WorkspaceId, err := strconv.ParseInt(workspaceId, 10, 64)
			if err != nil {
				return err
			}
			removeBindingRequest := catalog.UpdateWorkspaceBindings{
				Name:               catalogName,
				UnassignWorkspaces: []int64{i64WorkspaceId},
			}
			_, err = catalog.NewWorkspaceBindings(c.DatabricksClient).Update(ctx, removeBindingRequest)
			return err
		},
	})
}
