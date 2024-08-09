package mlflow

import (
	"context"
	"github.com/databricks/databricks-sdk-go/service/ml"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

type modelsData struct {
	Names []string `json:"names,omitempty" tf:"computed"`
}

func DataSourceModels() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *modelsData, w *databricks.WorkspaceClient) error {
		list, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
		if err != nil {
			return err
		}
		for _, m := range list {
			data.Names = append(data.Names, m.Name)
		}
		return nil
	})
}
