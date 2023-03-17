package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastores() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *MetastoresData, w *databricks.WorkspaceClient) error {
		metastores, err := w.Metastores.ListAll(ctx)
		if err != nil {
			return err
		}
		data.Metastores = metastores

		return nil
	})
}
