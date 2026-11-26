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

func ExampleSchemasAPI_Create_volumes() {
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

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Create_tables() {
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

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Create_schemas() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", newCatalog)

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Create_shares() {
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

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  createdCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Get_schemas() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", newCatalog)

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Schemas.GetByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_List_schemas() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", newCatalog)

	all, err := w.Schemas.ListAll(ctx, catalog.ListSchemasRequest{
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Update_schemas() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", newCatalog)

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Schemas.Update(ctx, catalog.UpdateSchema{
		FullName: created.FullName,
		Comment:  fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}
