// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleStatementExecutionAPI_ExecuteAndWait_tables() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	tableName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	_, err = w.StatementExecution.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: os.Getenv("TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s AS SELECT 2+2 as four", tableName),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleStatementExecutionAPI_ExecuteAndWait_shares() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	tableName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	_, err = w.StatementExecution.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: os.Getenv("TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s TBLPROPERTIES (delta.enableDeletionVectors=false) AS SELECT 2+2 as four", tableName),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
