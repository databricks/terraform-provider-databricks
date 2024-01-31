package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type tableParams struct {
	FullName string `json:"full_name" tf:"computed,optional`
}

func DataSourceTable() *schema.Resource {
	return common.WorkspaceDataWithParams(func(ctx context.Context, data tableParams, w *databricks.WorkspaceClient) (*catalog.TableInfo, error) {
		return w.Tables.Get(ctx, catalog.GetTableRequest{FullName: data.FullName})
	})
}
