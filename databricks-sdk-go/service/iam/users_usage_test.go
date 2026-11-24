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

func ExampleUsersAPI_Create_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	otherOwner, err := w.Users.Create(ctx, iam.User{
		UserName: fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", otherOwner)

	// cleanup

	err = w.Users.DeleteById(ctx, otherOwner.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Create_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

}

func ExampleUsersAPI_Create_accountUsers() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	// cleanup

	err = a.Users.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Delete_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	otherOwner, err := w.Users.Create(ctx, iam.User{
		UserName: fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", otherOwner)

	err = w.Users.DeleteById(ctx, otherOwner.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Delete_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	err = w.Users.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Delete_accountUsers() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	err = a.Users.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Get_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	fetch, err := w.Users.GetById(ctx, user.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", fetch)

}

func ExampleUsersAPI_Get_accountUsers() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	byId, err := a.Users.GetById(ctx, user.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = a.Users.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_List_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	allUsers, err := w.Users.ListAll(ctx, iam.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  iam.ListSortOrderDescending,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", allUsers)

}

func ExampleUsersAPI_Patch_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	err = w.Users.Patch(ctx, iam.PartialUpdate{
		Id: user.Id,
		Operations: []iam.Patch{{
			Op:    iam.PatchOpReplace,
			Path:  "active",
			Value: "false",
		}},
		Schemas: []iam.PatchSchema{iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp},
	})
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Patch_accountUsers() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	err = a.Users.Patch(ctx, iam.PartialUpdate{
		Id:      user.Id,
		Schemas: []iam.PatchSchema{iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp},
		Operations: []iam.Patch{{
			Op: iam.PatchOpAdd,
			Value: iam.User{
				Roles: []iam.ComplexValue{{
					Value: "account_admin",
				}},
			},
		}},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = a.Users.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleUsersAPI_Update_workspaceUsers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		UserName:    fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", user)

	err = w.Users.Update(ctx, iam.User{
		Id:       user.Id,
		UserName: user.UserName,
		Active:   true,
	})
	if err != nil {
		panic(err)
	}

}
