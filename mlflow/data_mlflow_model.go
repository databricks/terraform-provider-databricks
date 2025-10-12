package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceModel() common.Resource {
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *struct {
		common.Namespace
		Name            string             `json:"name"`
		UserId          string             `json:"user_id,omitempty" tf:"computed"`
		Description     string             `json:"description,omitempty" tf:"computed"`
		LatestVersions  []ml.ModelVersion  `json:"latest_versions,omitempty" tf:"computed"`
		Id              string             `json:"id" tf:"computed"`
		PermissionLevel ml.PermissionLevel `json:"permission_level,omitempty" tf:"computed"`
		Tags            []ml.ModelTag      `json:"tags,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		getModel, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: data.Name})
		if err != nil {
			return err
		}
		model := getModel.RegisteredModelDatabricks
		data.UserId = model.UserId
		data.Description = model.Description
		data.LatestVersions = model.LatestVersions
		data.Id = model.Id
		data.PermissionLevel = model.PermissionLevel
		data.Tags = model.Tags
		return nil
	})
}
