package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastores() *schema.Resource {
	type metastoresData struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.AccountData(func(ctx context.Context, data *metastoresData, acc *databricks.AccountClient) error {
		metastores, err := acc.Metastores.List(ctx)
		if err != nil {
			return err
		}
		for _, v := range metastores.Metastores {
			data.Ids = append(data.Ids, v.MetastoreId)
		}
		return nil
	})
}
