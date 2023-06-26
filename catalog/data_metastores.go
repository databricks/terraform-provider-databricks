package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastores() *schema.Resource {
	return common.AccountData(func(ctx context.Context, data *MetastoresData, acc *databricks.AccountClient) error {
		metastores, err := acc.Metastores.List(ctx)
		if err != nil {
			return err
		}
		data.Metastores = metastores.Metastores
		return nil
	})
}
