// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleQueryHistoryAPI_List_sqlQueryHistory() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	_, err = w.QueryHistory.List(ctx, sql.ListQueryHistoryRequest{
		FilterBy: &sql.QueryFilter{
			QueryStartTimeRange: &sql.TimeRange{
				StartTimeMs: 1690243200000,
				EndTimeMs:   1690329600000,
			},
		},
	})
	if err != nil {
		panic(err)
	}

}
