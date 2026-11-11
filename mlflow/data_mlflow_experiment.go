package mlflow

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
)

// likeWildcardReplacer replaces characters that have a special meaning inside a quoted
// MLflow search filter value (quotes, escapes and LIKE wildcards) with the single-character
// wildcard, so the filter stays valid. The search results are matched exactly afterwards.
var likeWildcardReplacer = strings.NewReplacer(`'`, "_", `"`, "_", `\`, "_", "%", "_")

// findExperimentByShortName resolves an experiment by its name relative to the workspace
// folder containing it, e.g. "my-experiment" for "/Users/user@example.com/my-experiment".
// It fails unless exactly one active experiment matches.
func findExperimentByShortName(ctx context.Context, w *databricks.WorkspaceClient, name string) (*ml.Experiment, error) {
	suffix := "/" + name
	experiments, err := w.Experiments.SearchExperimentsAll(ctx, ml.SearchExperiments{
		Filter: fmt.Sprintf("name LIKE '%%%s'", likeWildcardReplacer.Replace(suffix)),
	})
	if err != nil {
		return nil, err
	}
	var matches []ml.Experiment
	for _, experiment := range experiments {
		if strings.HasSuffix(experiment.Name, suffix) {
			matches = append(matches, experiment)
		}
	}
	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("experiment with name '%s' does not exist", name)
	case 1:
		return &matches[0], nil
	}
	paths := make([]string, 0, len(matches))
	for _, experiment := range matches {
		paths = append(paths, experiment.Name)
	}
	slices.Sort(paths)
	return nil, fmt.Errorf("there are %d experiments with name '%s': %s. Use the full path of the experiment instead",
		len(matches), name, strings.Join(paths, ", "))
}

func DataSourceExperiment() common.Resource {
	type experimentDataParams struct {
		common.Namespace
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

		if strings.HasPrefix(data.Name, "/") {
			experimentResponse, err := w.Experiments.GetByName(ctx, ml.GetByNameRequest{ExperimentName: data.Name})
			if err != nil {
				return nil, err
			}
			experiment.Experiment = *experimentResponse.Experiment
		} else if data.Name != "" {
			found, err := findExperimentByShortName(ctx, w, data.Name)
			if err != nil {
				return nil, err
			}
			experiment.Experiment = *found
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
