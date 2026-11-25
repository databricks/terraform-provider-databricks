// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleStorageCredentialsAPI_Create_volumes() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	storageCredential, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
		Comment: "created via SDK",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storageCredential)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_Create_storageCredentialsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_Create_externalLocationsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	credential, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", credential)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_Get_storageCredentialsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byName, err := w.StorageCredentials.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byName)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_List_storageCredentialsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.StorageCredentials.ListAll(ctx, catalog.ListStorageCredentialsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleStorageCredentialsAPI_Update_storageCredentialsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.StorageCredentials.Update(ctx, catalog.UpdateStorageCredential{
		Name:    created.Name,
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
