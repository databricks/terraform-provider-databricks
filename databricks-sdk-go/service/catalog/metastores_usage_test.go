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

func ExampleMetastoresAPI_Assign_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	workspaceId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(os.Getenv("DUMMY_WORKSPACE_ID"))

	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageRoot: fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Metastores.Assign(ctx, catalog.CreateMetastoreAssignment{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
		Id:    created.MetastoreId,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleMetastoresAPI_Create_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageRoot: fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
		Id:    created.MetastoreId,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleMetastoresAPI_Current_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	currentMetastore, err := w.Metastores.Current(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", currentMetastore)

}

func ExampleMetastoresAPI_Get_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageRoot: fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Metastores.GetById(ctx, created.MetastoreId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
		Id:    created.MetastoreId,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleMetastoresAPI_List_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Metastores.ListAll(ctx, catalog.ListMetastoresRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleMetastoresAPI_Summary_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	summary, err := w.Metastores.Summary(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", summary)

}

func ExampleMetastoresAPI_Unassign_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	workspaceId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(os.Getenv("DUMMY_WORKSPACE_ID"))

	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageRoot: fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Metastores.Unassign(ctx, catalog.UnassignRequest{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
		Id:    created.MetastoreId,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleMetastoresAPI_Update_metastores() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageRoot: fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Metastores.Update(ctx, catalog.UpdateMetastore{
		Id:      created.MetastoreId,
		NewName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
		Id:    created.MetastoreId,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
