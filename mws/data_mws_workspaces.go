package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids map[string]int64 `json:"ids" tf:"computed"`
	}
	s := common.StructToSchema(mwsWorkspacesData{}, nil)
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var data mwsWorkspacesData
			common.DataToStructPointer(d, s, &data)
			if newClient.Config.AccountID == "" {
				return fmt.Errorf("provider block is missing `account_id` property")
			}
			workspaces, err := NewWorkspacesAPI(ctx, newClient).List(newClient.Config.AccountID)
			if err != nil {
				return err
			}
			data.Ids = map[string]int64{}
			for _, v := range workspaces {
				data.Ids[v.WorkspaceName] = v.WorkspaceID
			}
			err = common.StructToData(data, s, d)
			if err != nil {
				return err
			}
			d.SetId("_")
			return nil
		},
	}
}
