package mws

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMwsNccBinding() common.Resource {
	type binding struct {
		WorkspaceId int64  `json:"workspace_id"`
		NccId       string `json:"ncc_id"`
	}
	s := common.StructToSchema(binding{}, common.NoCustomize)
	p := common.NewPairSeparatedID("workspace_id", "ncc_id", "/")
	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		acc, err := c.AccountClient()
		if err != nil {
			return err
		}
		wait, err := acc.Workspaces.Update(ctx, provisioning.UpdateWorkspaceRequest{
			NetworkConnectivityConfigId: d.Get("ncc_id").(string),
			WorkspaceId:                 int64(d.Get("workspace_id").(int)),
		})
		if err != nil {
			return err
		}
		_, err = wait.Get()
		if err != nil {
			return err
		}
		p.Pack(d)
		return nil
	}
	return common.Resource{
		Schema: s,
		Create: createOrUpdate,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			log.Printf("[WARN] Cannot check existing network connectivity config binding, only update is supported.")
			return nil
		},
		Update: createOrUpdate,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			log.Printf("[WARN] Cannot remove network connectivity config binding, only update is supported.")
			return nil
		},
	}
}
