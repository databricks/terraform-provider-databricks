// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package provisioning_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleNetworksAPI_Create_networks() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	netw, err := a.Networks.Create(ctx, provisioning.CreateNetworkRequest{
		NetworkName:      fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		VpcId:            fmt.Sprintf("%x", time.Now().UnixNano()),
		SubnetIds:        []string{fmt.Sprintf("%x", time.Now().UnixNano()), fmt.Sprintf("%x", time.Now().UnixNano())},
		SecurityGroupIds: []string{fmt.Sprintf("%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", netw)

}

func ExampleNetworksAPI_Get_networks() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	netw, err := a.Networks.Create(ctx, provisioning.CreateNetworkRequest{
		NetworkName:      fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		VpcId:            fmt.Sprintf("%x", time.Now().UnixNano()),
		SubnetIds:        []string{fmt.Sprintf("%x", time.Now().UnixNano()), fmt.Sprintf("%x", time.Now().UnixNano())},
		SecurityGroupIds: []string{fmt.Sprintf("%x", time.Now().UnixNano())},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", netw)

	byId, err := a.Networks.GetByNetworkId(ctx, netw.NetworkId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

}

func ExampleNetworksAPI_List_networks() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	configs, err := a.Networks.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", configs)

}
