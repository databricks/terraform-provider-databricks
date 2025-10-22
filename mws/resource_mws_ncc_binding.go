package mws

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NccBinding struct {
	WorkspaceId int64  `json:"workspace_id" tf:"force_new"`
	NccId       string `json:"network_connectivity_config_id"`
}

func ResourceMwsNccBinding() common.Resource {
	s := common.StructToSchema(NccBinding{}, nil)
	p := common.NewPairSeparatedID("workspace_id", "network_connectivity_config_id", "/")
	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		acc, err := c.AccountClient()
		if err != nil {
			return err
		}
		wait, err := acc.Workspaces.Update(ctx, provisioning.UpdateWorkspaceRequest{
			CustomerFacingWorkspace: provisioning.Workspace{
				NetworkConnectivityConfigId: d.Get("network_connectivity_config_id").(string),
			},
			UpdateMask:  "network_connectivity_config_id",
			WorkspaceId: int64(d.Get("workspace_id").(int)),
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
			if d.IsNewResource() {
				log.Print("[WARN] Importing NCC binding is not supported, skipping...")
			}
			return nil
		},
		Update: createOrUpdate,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			log.Printf("[WARN] Cannot remove network connectivity config binding, only update is supported.")
			return nil
		},
	}
}
