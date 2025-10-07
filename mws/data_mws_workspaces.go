package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids map[string]int64 `json:"ids" tf:"computed"`
	}
	return common.AccountData(func(ctx context.Context, data *mwsWorkspacesData, acc *databricks.AccountClient) error {
		workspaces, err := acc.Workspaces.List(ctx)
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
