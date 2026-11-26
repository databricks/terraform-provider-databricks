// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleWorkspaceBindingsAPI_Get_catalogWorkspaceBindings() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	bindings, err := w.WorkspaceBindings.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", bindings)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspaceBindingsAPI_Update_catalogWorkspaceBindings() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	thisWorkspaceId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(os.Getenv("THIS_WORKSPACE_ID"))

	created, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
		Name:             created.Name,
		AssignWorkspaces: []int64{thisWorkspaceId},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
