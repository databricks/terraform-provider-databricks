// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleQueriesAPI_Create_queries() {
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

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SHOW TABLES",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Create_alerts() {
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

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Get_queries() {
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

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SHOW TABLES",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	byId, err := w.Queries.GetById(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Update_queries() {
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

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SHOW TABLES",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	updated, err := w.Queries.Update(ctx, sql.UpdateQueryRequest{
		Id: query.Id,
		Query: &sql.UpdateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Description: "UPDATED: test query from Go SDK",
			QueryText:   "SELECT 2+2",
		},
		UpdateMask: "display_name,description,query_text",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", updated)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}
