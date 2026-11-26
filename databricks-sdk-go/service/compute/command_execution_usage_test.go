// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleCommandExecutionAPI_Create_commandsDirectUsage() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterId := os.Getenv("TEST_DEFAULT_CLUSTER_ID")

	context, err := w.CommandExecution.CreateAndWait(ctx, compute.CreateContext{
		ClusterId: clusterId,
		Language:  compute.LanguagePython,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", context)

	// cleanup

	err = w.CommandExecution.Destroy(ctx, compute.DestroyContext{
		ClusterId: clusterId,
		ContextId: context.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleCommandExecutionAPI_Execute_commandsDirectUsage() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterId := os.Getenv("TEST_DEFAULT_CLUSTER_ID")

	context, err := w.CommandExecution.CreateAndWait(ctx, compute.CreateContext{
		ClusterId: clusterId,
		Language:  compute.LanguagePython,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", context)

	textResults, err := w.CommandExecution.ExecuteAndWait(ctx, compute.Command{
		ClusterId: clusterId,
		ContextId: context.Id,
		Language:  compute.LanguagePython,
		Command:   "print(1)",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", textResults)

	// cleanup

	err = w.CommandExecution.Destroy(ctx, compute.DestroyContext{
		ClusterId: clusterId,
		ContextId: context.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleCommandExecutionAPI_Start_commands() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterId := os.Getenv("TEST_DEFAULT_CLUSTER_ID")

	commandContext, err := w.CommandExecution.Start(ctx, clusterId, compute.LanguagePython)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", commandContext)

}
