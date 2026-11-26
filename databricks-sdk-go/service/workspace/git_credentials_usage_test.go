// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package workspace_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/workspace"
)

func ExampleGitCredentialsAPI_Create_gitCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentialsRequest{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", cr)

	// cleanup

	err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
	if err != nil {
		panic(err)
	}

}

func ExampleGitCredentialsAPI_Get_gitCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentialsRequest{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", cr)

	byId, err := w.GitCredentials.GetByCredentialId(ctx, cr.CredentialId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
	if err != nil {
		panic(err)
	}

}

func ExampleGitCredentialsAPI_List_gitCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	list, err := w.GitCredentials.ListAll(ctx, workspace.ListCredentialsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", list)

}

func ExampleGitCredentialsAPI_Update_gitCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentialsRequest{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", cr)

	err = w.GitCredentials.Update(ctx, workspace.UpdateCredentialsRequest{
		CredentialId:        cr.CredentialId,
		GitProvider:         "gitHub",
		GitUsername:         fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
		PersonalAccessToken: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
	if err != nil {
		panic(err)
	}

}
