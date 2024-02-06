package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceCurrentMetastore() common.Resource {
	type CurrentMetastore struct {
		Id        string                               `json:"id,omitempty" tf:"computed"`
		Metastore *catalog.GetMetastoreSummaryResponse `json:"metastore_info,omitempty" tf:"computed" `
	}
	return common.WorkspaceData(func(ctx context.Context, data *CurrentMetastore, wc *databricks.WorkspaceClient) error {
		summary, err := wc.Metastores.Summary(ctx)
		if err != nil {
			return err
		}
		data.Metastore = summary
		data.Id = summary.MetastoreId

		return nil
	})
}
