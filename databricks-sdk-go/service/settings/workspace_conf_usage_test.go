// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package settings_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/settings"
)

func ExampleWorkspaceConfAPI_GetStatus_repos() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "enableWorkspaceFilesystem",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", conf)

}
