package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids             map[string]int64  `json:"ids" tf:"computed"`
		DeploymentNames map[string]string `json:"deployment_names" tf:"computed"`
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
		data.DeploymentNames = map[string]string{}
		for _, v := range workspaces {
			data.Ids[v.WorkspaceName] = v.WorkspaceID
			data.DeploymentNames[v.WorkspaceName] = v.DeploymentName
		}
		return nil
	})
}
