package mlflow

import (
	"context"
	"fmt"
	"github.com/databricks/databricks-sdk-go/service/ml"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceExperiment() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		ExperimentId     string `json:"id,omitempty" tf:"computed"`
		Name             string `json:"name,omitempty" tf:"computed"`
		ArtifactLocation string `json:"artifact_location,omitempty" tf:"computed"`
		LifecycleStage   string `json:"lifecycle_stage,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		var experiment *ml.Experiment
		if data.Name != "" {
			experimentResponse, err := w.Experiments.GetByName(ctx, ml.GetByNameRequest{ExperimentName: data.Name})
			if err != nil {
				return err
			}
			experiment = experimentResponse.Experiment
			data.ExperimentId = experiment.ExperimentId
		} else if data.ExperimentId != "" {
			experimentResponse, err := w.Experiments.GetExperiment(ctx, ml.GetExperimentRequest{ExperimentId: data.ExperimentId})
			if err != nil {
				return err
			}
			experiment = experimentResponse
			data.Name = experiment.Name
		} else {
			return fmt.Errorf("you need to specify either `name` or `id`")
		}
		data.ArtifactLocation = experiment.ArtifactLocation
		data.LifecycleStage = experiment.LifecycleStage
		return nil
	})
}
