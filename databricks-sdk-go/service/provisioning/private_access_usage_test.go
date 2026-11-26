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

func ExamplePrivateAccessAPI_Create_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.CreatePrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}

func ExamplePrivateAccessAPI_Get_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.CreatePrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}

func ExamplePrivateAccessAPI_List_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.PrivateAccess.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExamplePrivateAccessAPI_Replace_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.CreatePrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = a.PrivateAccess.Replace(ctx, provisioning.ReplacePrivateAccessSettingsRequest{
		PrivateAccessSettingsId: created.PrivateAccessSettingsId,
		CustomerFacingPrivateAccessSettings: provisioning.PrivateAccessSettings{
			PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Region:                    os.Getenv("AWS_REGION"),
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	_, err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}
