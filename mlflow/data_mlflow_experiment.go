package mlflow

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceExperiment() common.Resource {
	type experimentDataParams struct {
		ExperimentId string `json:"experiment_id" tf:"computed,optional"`
		Name         string `json:"name" tf:"computed,optional"`
	}

	type MlExperiment struct {
		ml.Experiment
		Id string `json:"id" tf:"computed,optional"`
	}

	return common.WorkspaceDataWithParams(func(ctx context.Context, data experimentDataParams, w *databricks.WorkspaceClient) (*MlExperiment, error) {
		var experiment *MlExperiment
		if data.ExperimentId == "" && data.Name == "" {
			return nil, fmt.Errorf("either 'experiment_id' or 'name' should be provided")
		}

		experiment = &MlExperiment{}

		if data.Name != "" {
			experimentResponse, err := w.Experiments.GetByName(ctx, ml.GetByNameRequest{ExperimentName: data.Name})
			if err != nil {
				return nil, err
			}
			experiment.Experiment = *experimentResponse.Experiment
		} else if data.ExperimentId != "" {
			experimentResponse, err := w.Experiments.GetExperiment(ctx, ml.GetExperimentRequest{ExperimentId: data.ExperimentId})
			if err != nil {
				return nil, err
			}
			experiment.Experiment = *experimentResponse.Experiment
		}

		experiment.Id = experiment.ExperimentId

		return experiment, nil
	})
}
