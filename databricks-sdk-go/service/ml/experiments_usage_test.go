// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package ml_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/ml"
)

func ExampleExperimentsAPI_CreateExperiment_experiments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleExperimentsAPI_CreateExperiment_mLflowRuns() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleExperimentsAPI_CreateRun_mLflowRuns() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	created, err := w.Experiments.CreateRun(ctx, ml.CreateRun{
		ExperimentId: experiment.ExperimentId,
		Tags: []ml.RunTag{{
			Key:   "foo",
			Value: "bar",
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}
	err = w.Experiments.DeleteRun(ctx, ml.DeleteRun{
		RunId: created.Run.Info.RunId,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleExperimentsAPI_GetExperiment_experiments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	_, err = w.Experiments.GetExperiment(ctx, ml.GetExperimentRequest{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleExperimentsAPI_ListExperiments_experiments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Experiments.ListExperimentsAll(ctx, ml.ListExperimentsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleExperimentsAPI_UpdateExperiment_experiments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	err = w.Experiments.UpdateExperiment(ctx, ml.UpdateExperiment{
		NewName:      fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleExperimentsAPI_UpdateRun_mLflowRuns() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", experiment)

	created, err := w.Experiments.CreateRun(ctx, ml.CreateRun{
		ExperimentId: experiment.ExperimentId,
		Tags: []ml.RunTag{{
			Key:   "foo",
			Value: "bar",
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Experiments.UpdateRun(ctx, ml.UpdateRun{
		RunId:  created.Run.Info.RunId,
		Status: ml.UpdateRunStatusKilled,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
		ExperimentId: experiment.ExperimentId,
	})
	if err != nil {
		panic(err)
	}
	err = w.Experiments.DeleteRun(ctx, ml.DeleteRun{
		RunId: created.Run.Info.RunId,
	})
	if err != nil {
		panic(err)
	}

}
