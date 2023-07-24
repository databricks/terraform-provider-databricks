package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastore() *schema.Resource {
	type AccountMetastoreByID struct {
		Id        string                 `json:"metastore_id"`
		Metastore *catalog.MetastoreInfo `json:"metastore_info,omitempty" tf:"computed" `
	}
	return common.AccountData(func(ctx context.Context, data *AccountMetastoreByID, acc *databricks.AccountClient) error {
		metastore, err := acc.Metastores.GetByMetastoreId(ctx, data.Id)
		if err != nil {
			return err
		}
		data.Metastore = metastore.MetastoreInfo
		return nil
	})
}
