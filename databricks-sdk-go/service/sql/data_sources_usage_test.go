// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
)

func ExampleDataSourcesAPI_List_queries() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

}

func ExampleDataSourcesAPI_List_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

}
