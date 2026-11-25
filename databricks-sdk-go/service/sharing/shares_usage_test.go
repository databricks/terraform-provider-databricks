// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sharing_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleSharesAPI_Create_shares() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	createdShare, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdShare)

	// cleanup

	err = w.Shares.DeleteByName(ctx, createdShare.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleSharesAPI_Get_shares() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	createdShare, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdShare)

	_, err = w.Shares.GetByName(ctx, createdShare.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Shares.DeleteByName(ctx, createdShare.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleSharesAPI_List_shares() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Shares.ListAll(ctx, sharing.ListSharesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleSharesAPI_Update_shares() {
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

	tableFullName := fmt.Sprintf("%s.%s.%s", createdCatalog.Name, createdSchema.Name, tableName)

	createdShare, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdShare)

	_, err = w.Shares.Update(ctx, sharing.UpdateShare{
		Name: createdShare.Name,
		Updates: []sharing.SharedDataObjectUpdate{{
			Action: sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{
				Name:           tableFullName,
				DataObjectType: "TABLE",
			},
		}},
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
	err = w.Tables.DeleteByFullName(ctx, tableFullName)
	if err != nil {
		panic(err)
	}
	err = w.Shares.DeleteByName(ctx, createdShare.Name)
	if err != nil {
		panic(err)
	}

}
