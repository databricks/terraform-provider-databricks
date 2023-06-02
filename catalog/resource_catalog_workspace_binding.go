package catalog

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type catalogWorkspaceBindingInfo struct {
	Name      string `json:"name" tf:"force_new"`
	Workspace int64  `json:"workspace" tf:"force_new"`
}

func contains_int64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ResourceCatalogWorkspaceBinding() *schema.Resource {
	s := common.StructToSchema(catalogWorkspaceBindingInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	pi := common.NewPairID("name", "workspace")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			workspaces := []int64{int64(d.Get("workspace").(int))}
			var createBindingRequest catalog.UpdateWorkspaceBindings
			createBindingRequest.AssignWorkspaces = workspaces
			_, err = w.WorkspaceBindings.Update(ctx, createBindingRequest)
			if err != nil {
				return err
			}
			pi.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			name, workspace, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			workspaceId, err := strconv.ParseInt(workspace, 10, 64)
			if err != nil {
				return err
			}
			bindings, err := w.WorkspaceBindings.GetByName(ctx, name)
			if err != nil {
				return err
			}
			cwbInfo := catalogWorkspaceBindingInfo{
				Name:      name,
				Workspace: workspaceId,
			}
			if contains_int64(bindings.Workspaces, workspaceId) {
				return common.StructToData(&cwbInfo, s, d)
			} else {
				return nil
			}
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			workspaces := []int64{int64(d.Get("workspace").(int))}
			var createBindingRequest catalog.UpdateWorkspaceBindings
			createBindingRequest.UnassignWorkspaces = workspaces
			_, err = w.WorkspaceBindings.Update(ctx, createBindingRequest)
			if err != nil {
				return err
			}
			return nil
		},
	}.ToResource()
}
