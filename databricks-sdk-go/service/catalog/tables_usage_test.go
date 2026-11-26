// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

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

func ExampleTablesAPI_Get_tables() {
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

	tableFullName := fmt.Sprintf("%s.%s.%s", createdCatalog.Name, createdSchema.Name, tableName)

	createdTable, err := w.Tables.GetByFullName(ctx, tableFullName)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdTable)

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
	err = w.Tables.DeleteByFullName(ctx, tableFullName)
	if err != nil {
		panic(err)
	}

}

func ExampleTablesAPI_List_tables() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

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

	allTables, err := w.Tables.ListAll(ctx, catalog.ListTablesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", allTables)

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

func ExampleTablesAPI_ListSummaries_tables() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

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

	summaries, err := w.Tables.ListSummariesAll(ctx, catalog.ListSummariesRequest{
		CatalogName:       createdCatalog.Name,
		SchemaNamePattern: createdSchema.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", summaries)

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
