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

func ExampleVolumesAPI_Create_volumes() {
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

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	createdVolume, err := w.Volumes.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName:     createdCatalog.Name,
		SchemaName:      createdSchema.Name,
		Name:            fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageLocation: externalLocation.Url,
		VolumeType:      catalog.VolumeTypeExternal,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdVolume)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, externalLocation.Name)
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Volumes.DeleteByName(ctx, createdVolume.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleVolumesAPI_List_volumes() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	allVolumes, err := w.Volumes.ListAll(ctx, catalog.ListVolumesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", allVolumes)

	// cleanup

	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleVolumesAPI_Read_volumes() {
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

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	createdVolume, err := w.Volumes.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName:     createdCatalog.Name,
		SchemaName:      createdSchema.Name,
		Name:            fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageLocation: externalLocation.Url,
		VolumeType:      catalog.VolumeTypeExternal,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdVolume)

	loadedVolume, err := w.Volumes.ReadByName(ctx, createdVolume.FullName)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", loadedVolume)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, externalLocation.Name)
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Volumes.DeleteByName(ctx, createdVolume.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleVolumesAPI_Update_volumes() {
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

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdCatalog)

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: createdCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdSchema)

	createdVolume, err := w.Volumes.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName:     createdCatalog.Name,
		SchemaName:      createdSchema.Name,
		Name:            fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		StorageLocation: externalLocation.Url,
		VolumeType:      catalog.VolumeTypeExternal,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdVolume)

	loadedVolume, err := w.Volumes.ReadByName(ctx, createdVolume.FullName)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", loadedVolume)

	_, err = w.Volumes.Update(ctx, catalog.UpdateVolumeRequestContent{
		Name:    loadedVolume.FullName,
		Comment: "Updated volume comment",
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, externalLocation.Name)
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}
	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Volumes.DeleteByName(ctx, createdVolume.FullName)
	if err != nil {
		panic(err)
	}

}
