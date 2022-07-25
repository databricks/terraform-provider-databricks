package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMwsWorkspaces() *schema.Resource {
	type mwsWorkspacesData struct {
		Ids map[string]int64 `json:"ids,omitempty" tf:"computed"`
	}
	return common.DataResource(mwsWorkspacesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsWorkspacesData)
		if c.AccountID == "" {
			return fmt.Errorf("provider block is missing `account_id` property")
		}
		workspaces, err := NewWorkspacesAPI(ctx, c).List(c.AccountID)
		if err != nil {
			return err
		}
		data.Ids = map[string]int64{}
		for _, v := range workspaces {
			data.Ids[v.WorkspaceName] = v.WorkspaceID
		}
		return nil
	})
}
