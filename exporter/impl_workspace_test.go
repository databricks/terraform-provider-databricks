package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/stretchr/testify/assert"
)

func TestEmitFilesFromMap(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FShared%2Ftest.txt&return_git_info=true",
			Response: workspace.ObjectStatus{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FShared%2Fgit%2Ftest.txt&return_git_info=true",
			Response: workspace.ObjectStatus{
				GitInfo: &sdk_workspace.RepoInfo{
					Id: 1234,
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("storage,notebooks,wsfiles,repos")
		ic.emitFilesFromMap(map[string]string{
			"k1": "dbfs:/FileStore/test.txt",
			"k2": "/Workspace/Shared/test.txt",
			"k3": "nothing",
			"k4": "/Workspace/Shared/git/test.txt",
		})
		assert.Equal(t, 3, len(ic.testEmits))
		assert.Contains(t, ic.testEmits, "databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/test.txt)")
		assert.Contains(t, ic.testEmits, "databricks_workspace_file[<unknown>] (id: /Shared/test.txt)")
		assert.Contains(t, ic.testEmits, "databricks_repo[<unknown>] (id: 1234)")
	})
}

func TestRepoName(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().ToResource().TestResourceData()
	d.SetId("12345")
	// Repo without path
	assert.Equal(t, "repo_12345", resourcesMap["databricks_repo"].Name(ic, d))
	// Repo with path
	d.Set("path", "/Repos/user/test")
	assert.Equal(t, "user_test_12345", resourcesMap["databricks_repo"].Name(ic, d))
}

func TestRepoIgnore(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().ToResource().TestResourceData()
	d.SetId("12345")
	d.Set("path", "/Repos/user/test")
	r := &resource{ID: "12345", Data: d}
	// Repo without URL
	assert.True(t, resourcesMap["databricks_repo"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
	// Repo with URL
	d.Set("url", "https://github.com/abc/abc.git")
	assert.False(t, resourcesMap["databricks_repo"].Ignore(ic, r))
}

func TestShouldOmitFoRepos(t *testing.T) {
	d := repos.ResourceRepo().ToResource().TestResourceData()
	d.SetId("1234")
	d.Set("path", "/Repos/Test/repo")
	r := &resource{
		Attribute: "databricks_repo",
		Value:     "repo",
		Data:      d,
	}
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "path",
		repos.ResourceRepo().Schema["path"], d, r))
	assert.True(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "branch",
		repos.ResourceRepo().Schema["branch"], d, r))
	assert.True(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "tag",
		repos.ResourceRepo().Schema["tag"], d, r))
	d.Set("branch", "test")
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "branch",
		repos.ResourceRepo().Schema["branch"], d, r))
	d.Set("tag", "v123")
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "tag",
		repos.ResourceRepo().Schema["tag"], d, r))
}

func TestRepoListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_repo"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func TestNotebookWorkspaceFileImportNotFound(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_notebook"].Import(ic, &resource{
			ID: "/abc",
		})
		assert.EqualError(t, err, "nope")
		assert.Contains(t, ic.ignoredResources, "databricks_notebook. path=/abc")

		err = resourcesMap["databricks_workspace_file"].Import(ic, &resource{
			ID: "/def",
		})
		assert.EqualError(t, err, "nope")
		assert.Contains(t, ic.ignoredResources, "databricks_workspace_file. path=/def")
	})
}

func TestNotebookGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						Path:       "/First/Second",
						ObjectType: "NOTEBOOK",
						ObjectID:   123,
						Language:   "PYTHON",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
		},
	}, "notebooks", false, func(ic *importContext) {
		ic.notebooksFormat = "SOURCE"
		ic.enableListing("notebooks")
		err := listWorkspaceObjects(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "first_second_123" {
		  source = "${path.module}/notebooks/First/Second_123.py"
		  path   = "/First/Second"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestNotebookGenerationJupyter(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						ObjectID:   123,
						ObjectType: "NOTEBOOK",
						Path:       "/First/Second",
						Language:   "PYTHON",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=JUPYTER&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
		},
	}, "notebooks", false, func(ic *importContext) {
		ic.notebooksFormat = "JUPYTER"
		ic.enableListing("notebooks")
		err := listWorkspaceObjects(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "first_second_123" {
		  source   = "${path.module}/notebooks/First/Second_123.ipynb"
		  path     = "/First/Second"
		  language = "PYTHON"
		  format   = "JUPYTER"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestNotebookGenerationBadCharacters(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						ObjectID:   123,
						ObjectType: "NOTEBOOK",
						Path:       "/Fir\"st\\/Second",
						Language:   "PYTHON",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FFir%22st%5C",

			Response: workspace.ObjectStatus{
				ObjectID:   124,
				ObjectType: "DIRECTORY",
				Path:       "/Fir\"st\\",
			},
			ReuseRequest: true,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFir%22st%5C%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
		},
	}, "notebooks,directories", true, func(ic *importContext) {
		ic.notebooksFormat = "SOURCE"
		ic.enableServices("notebooks")
		ic.enableListing("notebooks")
		err := listWorkspaceObjects(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "fir_st_second_123" {
		  source = "${path.module}/notebooks/Fir_st_/Second_123.py"
		  path   = "/Fir\"st\\/Second"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestDirectoryGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						ObjectID:   1234,
						Path:       "/first",
						ObjectType: "DIRECTORY",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Ffirst",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2Ffirst",
			Response: workspace.ObjectStatus{
				ObjectID:   1234,
				ObjectType: "DIRECTORY",
				Path:       "/first",
			},
		},
	}, "directories", false, func(ic *importContext) {
		ic.enableListing("directories")
		err := listWorkspaceObjects(ic)
		assert.NoError(t, err)

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_directory" "first_1234" {
		  path = "/first"
		}`), getGeneratedFile(ic, "directories"))
	})
}
