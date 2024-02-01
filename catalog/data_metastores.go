package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMetastores() common.Resource {
	type metastoresData struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.AccountData(func(ctx context.Context, data *metastoresData, acc *databricks.AccountClient) error {
		metastores, err := acc.Metastores.ListAll(ctx)
		if err != nil {
			return err
		}
		data.Ids = map[string]string{}
		for _, v := range metastores {
			name := v.Name
			_, duplicateName := data.Ids[name]
			if duplicateName {
				return fmt.Errorf("duplicate metastore name detected: %s", name)
			}
			data.Ids[name] = v.MetastoreId
		}
		return nil
	})
}
