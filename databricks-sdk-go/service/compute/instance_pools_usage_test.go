// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleInstancePoolsAPI_Create_instancePools() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	smallest, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", smallest)

	created, err := w.InstancePools.Create(ctx, compute.CreateInstancePool{
		InstancePoolName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		NodeTypeId:       smallest,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)
	if err != nil {
		panic(err)
	}

}

func ExampleInstancePoolsAPI_Edit_instancePools() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	smallest, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", smallest)

	created, err := w.InstancePools.Create(ctx, compute.CreateInstancePool{
		InstancePoolName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		NodeTypeId:       smallest,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.InstancePools.Edit(ctx, compute.EditInstancePool{
		InstancePoolId:   created.InstancePoolId,
		InstancePoolName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		NodeTypeId:       smallest,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)
	if err != nil {
		panic(err)
	}

}

func ExampleInstancePoolsAPI_Get_instancePools() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	smallest, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", smallest)

	created, err := w.InstancePools.Create(ctx, compute.CreateInstancePool{
		InstancePoolName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		NodeTypeId:       smallest,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.InstancePools.GetByInstancePoolId(ctx, created.InstancePoolId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)
	if err != nil {
		panic(err)
	}

}

func ExampleInstancePoolsAPI_List_instancePools() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.InstancePools.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
