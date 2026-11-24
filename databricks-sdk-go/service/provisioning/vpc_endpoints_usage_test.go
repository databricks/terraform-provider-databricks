// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package provisioning_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleVpcEndpointsAPI_Create_vpcEndpoints() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.VpcEndpoints.Create(ctx, provisioning.CreateVpcEndpointRequest{
		AwsVpcEndpointId: os.Getenv("TEST_RELAY_VPC_ENDPOINT"),
		Region:           os.Getenv("AWS_REGION"),
		VpcEndpointName:  fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = a.VpcEndpoints.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
	if err != nil {
		panic(err)
	}

}

func ExampleVpcEndpointsAPI_Get_vpcEndpoints() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.VpcEndpoints.Create(ctx, provisioning.CreateVpcEndpointRequest{
		AwsVpcEndpointId: os.Getenv("TEST_RELAY_VPC_ENDPOINT"),
		Region:           os.Getenv("AWS_REGION"),
		VpcEndpointName:  fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.VpcEndpoints.GetByVpcEndpointId(ctx, created.VpcEndpointId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.VpcEndpoints.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
	if err != nil {
		panic(err)
	}

}

func ExampleVpcEndpointsAPI_List_vpcEndpoints() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.VpcEndpoints.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
