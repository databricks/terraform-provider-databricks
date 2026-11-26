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

func ExampleReposAPI_Create_repos() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	root := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	ri, err := w.Repos.Create(ctx, workspace.CreateRepoRequest{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", ri)

	// cleanup

	err = w.Repos.DeleteByRepoId(ctx, ri.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleReposAPI_Get_repos() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	root := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	ri, err := w.Repos.Create(ctx, workspace.CreateRepoRequest{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", ri)

	byId, err := w.Repos.GetByRepoId(ctx, ri.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Repos.DeleteByRepoId(ctx, ri.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleReposAPI_List_repos() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Repos.ListAll(ctx, workspace.ListReposRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleReposAPI_Update_repos() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	root := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	ri, err := w.Repos.Create(ctx, workspace.CreateRepoRequest{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", ri)

	err = w.Repos.Update(ctx, workspace.UpdateRepoRequest{
		RepoId: ri.Id,
		Branch: "foo",
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Repos.DeleteByRepoId(ctx, ri.Id)
	if err != nil {
		panic(err)
	}

}
