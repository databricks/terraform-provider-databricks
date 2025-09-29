package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
)

type SharesData struct {
	Shares []string `json:"shares,omitempty" tf:"computed,slice_set"`
	common.ProviderConfig
}

func DataSourceShares() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *SharesData, w *databricks.WorkspaceClient) error {

		shares, err := w.Shares.ListAll(ctx, sharing.ListSharesRequest{})
		if err != nil {
			return err
		}
		for _, share := range shares {
			data.Shares = append(data.Shares, share.Name)
		}
		return nil
	})
}
