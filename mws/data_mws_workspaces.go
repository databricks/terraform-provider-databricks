package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids           map[string]int64 `json:"ids" tf:"computed"`
		MwsWorkspaces []Workspace      `json:"mws_workspaces" tf:"computed"`
	}
	return common.DataResource(mwsWorkspacesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsWorkspacesData)
		if c.Config.AccountID == "" {
			return fmt.Errorf("provider block is missing `account_id` property")
		}
		workspaces, err := NewWorkspacesAPI(ctx, c).List(c.Config.AccountID)
		if err != nil {
			return err
		}
		data.Ids = map[string]int64{}
		for _, workspace := range workspaces {
			data.Ids[workspace.WorkspaceName] = workspace.WorkspaceID
			data.MwsWorkspaces = append(data.MwsWorkspaces, workspace)
		}
		return nil
	})
}
