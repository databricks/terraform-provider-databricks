// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

import (
	"context"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleAccountMetastoreAssignmentsAPI_ListAll_metastoreAssignments() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	ws, err := a.MetastoreAssignments.ListAll(ctx, catalog.ListAccountMetastoreAssignmentsRequest{
		MetastoreId: os.Getenv("TEST_METASTORE_ID"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", ws)

}
