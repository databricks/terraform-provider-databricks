// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package iam_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ExampleWorkspaceAssignmentAPI_List_workspaceAssignmentOnAws() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	workspaceId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(os.Getenv("DUMMY_WORKSPACE_ID"))

	all, err := a.WorkspaceAssignment.ListByWorkspaceId(ctx, workspaceId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleWorkspaceAssignmentAPI_Update_workspaceAssignmentOnAws() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	spn, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", spn)

	spnId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(spn.Id)

	workspaceId := func(v string) int64 {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
		}
		return i
	}(os.Getenv("DUMMY_WORKSPACE_ID"))

	_, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PrincipalId: spnId,
		Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
	})
	if err != nil {
		panic(err)
	}

}
