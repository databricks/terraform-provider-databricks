package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CatalogWorkspaceBindingInfo struct {
	Name      string `json:"name" tf:"force_new"`
	Workspace int64  `json:"workspace" tf:"force_new"`
}

func ResourceCatalogWorkspaceBinding() *schema.Resource {
	s := common.StructToSchema(CatalogWorkspaceBindingInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			workspaces := []int64{d.Get("Workspace").(int64)}
			var createBindingRequest catalog.UpdateWorkspaceBindings
			createBindingRequest.AssignWorkspaces = workspaces
			d.Set("AssignWorkspaces", d.Get("Workspace"))
			common.DataToStructPointer(d, s, &createBindingRequest)
			w.WorkspaceBindings.Update(ctx, createBindingRequest)
			if err != nil {
				return err
			}
			// FIX: Need to set an id that can identify a unique workspace-catalog combination
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// FIX: This will return a full list of bound workspaces
			w.WorkspaceBindings.GetByName(ctx, d.Id())
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			workspaces := []int64{d.Get("Workspace").(int64)}
			var createBindingRequest catalog.UpdateWorkspaceBindings
			createBindingRequest.UnassignWorkspaces = workspaces
			d.Set("UnassignWorkspaces", d.Get("Workspace"))
			common.DataToStructPointer(d, s, &createBindingRequest)
			w.WorkspaceBindings.Update(ctx, createBindingRequest)
			if err != nil {
				return err
			}
			return nil
		},
	}.ToResource()
}
