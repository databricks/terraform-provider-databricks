package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids map[string]int64 `json:"ids" tf:"computed"`
	}
	return common.DataResource(mwsWorkspacesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsWorkspacesData)
		if c.Config.AccountID == "" {
			return fmt.Errorf("provider block is missing `account_id` property")
		}
		a, err := c.AccountClient()
		if err != nil {
			return err
		}
		workspaces, err := a.Workspaces.List(ctx)
		if err != nil {
			return err
		}
		data.Ids = map[string]int64{}
		for _, v := range workspaces {
			data.Ids[v.WorkspaceName] = v.WorkspaceId
		}
		return nil
	})
}
