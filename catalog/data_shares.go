package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceShares() common.Resource {
	dataSource := common.WorkspaceData(func(ctx context.Context, data *struct {
		Shares []string `json:"shares,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {

		shares, err := w.Shares.ListAll(ctx)
		if err != nil {
			return err
		}
		for _, share := range shares {
			data.Shares = append(data.Shares, share.Name)
		}
		return nil
	})
	dataSource.WorkspaceIdField = common.ManagementWorkspaceId
	return dataSource
}
