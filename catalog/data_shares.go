package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceShares() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
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
}
