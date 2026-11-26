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

func ExampleWarehousesAPI_Create_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
		Tags: &sql.EndpointTags{
			CustomTags: []sql.EndpointTagPair{{
				Key:   "Owner",
				Value: "eng-dev-ecosystem-team_at_databricks.com",
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Warehouses.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_Edit_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
		Tags: &sql.EndpointTags{
			CustomTags: []sql.EndpointTagPair{{
				Key:   "Owner",
				Value: "eng-dev-ecosystem-team_at_databricks.com",
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Warehouses.Edit(ctx, sql.EditWarehouseRequest{
		Id:             created.Id,
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Warehouses.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_Get_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
		Tags: &sql.EndpointTags{
			CustomTags: []sql.EndpointTagPair{{
				Key:   "Owner",
				Value: "eng-dev-ecosystem-team_at_databricks.com",
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	wh, err := w.Warehouses.GetById(ctx, created.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", wh)

	// cleanup

	err = w.Warehouses.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_List_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
