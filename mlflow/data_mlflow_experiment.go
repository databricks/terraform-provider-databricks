package mlflow

import (
	"context"
	"github.com/databricks/databricks-sdk-go/service/ml"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceExperiment() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		ExperimentId     string `json:"experiment_id,omitempty" tf:"computed"`
		Name             string `json:"name,omitempty" tf:"computed"`
		ArtifactLocation string `json:"artifact_location,omitempty" tf:"computed"`
		LifecycleStage   string `json:"lifecycle_stage,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		experimentResponse, err := w.Experiments.GetByName(ctx, ml.GetByNameRequest{ExperimentName: data.Name})
		if err != nil {
			return err
		}
		experiment := experimentResponse.Experiment
		data.Name = experiment.Name
		data.ExperimentId = experiment.ExperimentId
		data.ArtifactLocation = experiment.ArtifactLocation
		data.LifecycleStage = experiment.LifecycleStage
		return nil
	})
}
