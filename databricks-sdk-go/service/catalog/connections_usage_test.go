// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleConnectionsAPI_Create_connections() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	connCreate, err := w.Connections.Create(ctx, catalog.CreateConnection{
		Comment:        "Go SDK Acceptance Test Connection",
		ConnectionType: catalog.ConnectionTypeDatabricks,
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Options:        map[string]string{"host": fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "httpPath": fmt.Sprintf("/sql/1.0/warehouses/%s", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "personalAccessToken": fmt.Sprintf("sdk-%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connCreate)

	// cleanup

	err = w.Connections.Delete(ctx, catalog.DeleteConnectionRequest{
		Name: connCreate.Name,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleConnectionsAPI_Get_connections() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	connCreate, err := w.Connections.Create(ctx, catalog.CreateConnection{
		Comment:        "Go SDK Acceptance Test Connection",
		ConnectionType: catalog.ConnectionTypeDatabricks,
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Options:        map[string]string{"host": fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "httpPath": fmt.Sprintf("/sql/1.0/warehouses/%s", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "personalAccessToken": fmt.Sprintf("sdk-%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connCreate)

	connUpdate, err := w.Connections.Update(ctx, catalog.UpdateConnection{
		Name:    connCreate.Name,
		Options: map[string]string{"host": fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "httpPath": fmt.Sprintf("/sql/1.0/warehouses/%s", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "personalAccessToken": fmt.Sprintf("sdk-%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connUpdate)

	conn, err := w.Connections.Get(ctx, catalog.GetConnectionRequest{
		Name: connUpdate.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", conn)

	// cleanup

	err = w.Connections.Delete(ctx, catalog.DeleteConnectionRequest{
		Name: connCreate.Name,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleConnectionsAPI_List_connections() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	connList, err := w.Connections.ListAll(ctx, catalog.ListConnectionsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connList)

}

func ExampleConnectionsAPI_Update_connections() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	connCreate, err := w.Connections.Create(ctx, catalog.CreateConnection{
		Comment:        "Go SDK Acceptance Test Connection",
		ConnectionType: catalog.ConnectionTypeDatabricks,
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Options:        map[string]string{"host": fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "httpPath": fmt.Sprintf("/sql/1.0/warehouses/%s", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "personalAccessToken": fmt.Sprintf("sdk-%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connCreate)

	connUpdate, err := w.Connections.Update(ctx, catalog.UpdateConnection{
		Name:    connCreate.Name,
		Options: map[string]string{"host": fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "httpPath": fmt.Sprintf("/sql/1.0/warehouses/%s", fmt.Sprintf("sdk-%x", time.Now().UnixNano())), "personalAccessToken": fmt.Sprintf("sdk-%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", connUpdate)

	// cleanup

	err = w.Connections.Delete(ctx, catalog.DeleteConnectionRequest{
		Name: connCreate.Name,
	})
	if err != nil {
		panic(err)
	}

}
