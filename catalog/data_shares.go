package catalog

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceShares() *schema.Resource {
	type sharesData struct {
		Shares []string `json:"shares,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(sharesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*sharesData)
		sharesAPI := NewSharesAPI(ctx, c)
		shares, err := sharesAPI.list()
		if err != nil {
			return err
		}
		for _, share := range shares.Shares {
			data.Shares = append(data.Shares, share.Name)
		}
		return nil
	})
}
