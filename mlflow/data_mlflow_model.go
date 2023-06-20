package mlflow

import (
	"context"
	"fmt"
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceModel() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Name         string            `json:"name,omitempty" tf:"computed"`
		Version      string            `json:"version,omitempty" tf:"computed"`
		ModelVersion []ml.ModelVersion `json:"model_versions,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		var model []ml.ModelVersion

		if data.Name == "" {
			return fmt.Errorf("you need to specify a `name`")
		}

		// Get latest version models for each requests stage if no version is specified.
		if data.Version == "" {
			latestVersion, err := w.ModelRegistry.GetLatestVersionsAll(ctx, ml.GetLatestVersionsRequest{Name: data.Name})
			if err != nil {
				return err
			}
			model = latestVersion
		} else {
			modelResponse, err := w.ModelRegistry.GetModelVersion(ctx, ml.GetModelVersionRequest{
				Name:    data.Name,
				Version: data.Version,
			})
			if err != nil {
				return err
			}
			model = append(model, *modelResponse.ModelVersion)
		}

		data.ModelVersion = model
		return nil
	})
}
