package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceSchema() common.Resource {
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *struct {
		common.Namespace
		Id     string              `json:"id,omitempty" tf:"computed"`
		Name   string              `json:"name"`
		Schema *catalog.SchemaInfo `json:"schema_info,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		schema, err := w.Schemas.GetByFullName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Schema = schema
		data.Id = schema.FullName
		return nil
	})
}
