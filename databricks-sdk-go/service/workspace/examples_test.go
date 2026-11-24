package workspace_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

func ExampleReposAPI_GetByPath_checkoutBranchByPath() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}
	// shortcut for getting RepoInfo by path
	repo, err := w.Repos.GetByPath(ctx, "/Repos/path/to/prod")
	if err != nil {
		panic(err)
	}
	// because you can update repo only by ID, not by path
	err = w.Repos.Update(ctx, workspace.UpdateRepoRequest{
		RepoId: repo.Id,
		Branch: "v1.4.18",
	})
	if err != nil {
		panic(err)
	}
}
