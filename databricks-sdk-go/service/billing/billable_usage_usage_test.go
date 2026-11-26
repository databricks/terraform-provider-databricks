// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package billing_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/billing"
)

func ExampleBillableUsageAPI_Download_usageDownload() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	resp, err := a.BillableUsage.Download(ctx, billing.DownloadRequest{
		StartMonth: "2024-08",
		EndMonth:   "2024-09",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", resp)

}
