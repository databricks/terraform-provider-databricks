// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package iam_test

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ExamplePermissionsAPI_Get_genericPermissions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obj)

	_, err = w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	if err != nil {
		panic(err)
	}

}

func ExamplePermissionsAPI_GetPermissionLevels_genericPermissions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obj)

	levels, err := w.Permissions.GetPermissionLevels(ctx, iam.GetPermissionLevelsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", levels)

}

func ExamplePermissionsAPI_Set_genericPermissions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obj)

	_, err = w.Permissions.Set(ctx, iam.SetObjectPermissions{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
		AccessControlList: []iam.AccessControlRequest{{
			GroupName:       group.DisplayName,
			PermissionLevel: iam.PermissionLevelCanRun,
		}},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}

}
