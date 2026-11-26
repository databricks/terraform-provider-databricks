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

func ExampleExternalLocationsAPI_Create_volumes() {
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

	externalLocation, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: storageCredential.Name,
		Comment:        "created via SDK",
		Url:            "s3://" + os.Getenv("TEST_BUCKET") + "/" + fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", externalLocation)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, externalLocation.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleExternalLocationsAPI_Create_externalLocationsOnAws() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleExternalLocationsAPI_Get_externalLocationsOnAws() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ExternalLocations.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleExternalLocationsAPI_List_externalLocationsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ExternalLocations.ListAll(ctx, catalog.ListExternalLocationsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleExternalLocationsAPI_Update_externalLocationsOnAws() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ExternalLocations.Update(ctx, catalog.UpdateExternalLocation{
		Name:           created.Name,
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
