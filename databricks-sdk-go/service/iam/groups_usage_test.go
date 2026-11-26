// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package iam_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ExampleGroupsAPI_Create_genericPermissions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Create_groups() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Create_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Delete_genericPermissions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Delete_groups() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Delete_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleGroupsAPI_Get_groups() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	fetch, err := w.Groups.GetById(ctx, group.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", fetch)

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}
