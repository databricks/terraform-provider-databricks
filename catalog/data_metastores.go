package catalog

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastores() *schema.Resource {
	type MetastoresData struct {
		Metastores []MetastoreInfo `json:"objects,omitempty" tf:"computed,slice_set,alias:object"`
	}
	return common.DataResource(MetastoresData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*MetastoresData)
		metastoresAPI := NewMetastoresAPI(ctx, c)
		metastores, err := metastoresAPI.listMetastores()
		if err != nil {
			return err
		}
		data.Metastores = append(data.Metastores, metastores.Metastores...)

		return nil
	})
}
