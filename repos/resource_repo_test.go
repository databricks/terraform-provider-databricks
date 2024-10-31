package repos

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestGetGitProviderFromUrl(t *testing.T) {
	assert.Equal(t, "bitbucketCloud", GetGitProviderFromUrl("https://user@bitbucket.org/user/repo.git"))
	assert.Equal(t, "gitHub", GetGitProviderFromUrl("https://github.com//user/repo.git"))
	assert.Equal(t, "azureDevOpsServices", GetGitProviderFromUrl("https://user@dev.azure.com/user/project/_git/repo"))
	assert.Equal(t, "", GetGitProviderFromUrl("https://abc/user/repo.git"))
	assert.Equal(t, "", GetGitProviderFromUrl("ewfgwergfwe"))
	assert.Equal(t, "awsCodeCommit", GetGitProviderFromUrl("https://git-codecommit.us-east-2.amazonaws.com/v1/repos/MyDemoRepo"))
}

func TestResourceRepoRead(t *testing.T) {
	repoID := 48155820875912
	repoIDStr := fmt.Sprintf("%d", repoID)
	url := "https://user@github.com/user/repo.git"
	provider := "gitHub"
	branch := "main"
	path := "/Repos/Production/testrepo"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/repos/%d", repoID),
				Response: ReposInformation{
					ID:           int64(repoID),
					Url:          url,
					Provider:     provider,
					Branch:       branch,
					HeadCommitID: "7e0847ede61f07adede22e2bcce6050216489171",
					Path:         path,
				},
			},
		},
		Resource: ResourceRepo(),
		Read:     true,
		New:      true,
		ID:       repoIDStr,
	}.ApplyAndExpectData(t,
		map[string]any{"id": repoIDStr, "path": path, "branch": branch, "git_provider": provider,
			"url": url, "commit_hash": "7e0847ede61f07adede22e2bcce6050216489171",
			"workspace_path": "/Workspace" + path})
}

func TestResourceRepoRead_NotFound(t *testing.T) {
	repoID := "48155820875912"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/repos/%s", repoID),
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Repo could not be found",
				},
				Status: 404,
			},
		},
		Resource: ResourceRepo(),
		Read:     true,
		Removed:  true,
		ID:       repoID,
	}.Apply(t)
}

func TestResourceRepoDelete(t *testing.T) {
	repoID := "48155820875912"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/repos/%s", repoID),
				Status:   http.StatusOK,
			},
		},
		Resource: ResourceRepo(),
		Delete:   true,
		ID:       repoID,
	}.ApplyAndExpectData(t,
		map[string]any{"id": repoID})
}

func TestResourceRepoCreateNoBranch(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Branch:       "main",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: reposCreateRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
				},
				Response: resp,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url": "https://github.com/user/test.git",
		},
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{"id": resp.RepoID(), "path": resp.Path, "branch": resp.Branch,
			"git_provider": resp.Provider, "url": resp.Url, "commit_hash": resp.HeadCommitID})
}

func TestResourceRepoCreateCustomDirectory(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Branch:       "main",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: reposCreateRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
					Path:     "/Repos/Production/test/",
				},
				Response: resp,
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/Repos/Production",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/Production/test/",
		},
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{"id": resp.RepoID(), "path": resp.Path, "branch": resp.Branch,
			"git_provider": resp.Provider, "url": resp.Url, "commit_hash": resp.HeadCommitID})
}

func TestResourceRepoCreateCustomDirectoryError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/Repos/Production",
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/Production/test/",
		},
		Create: true,
	}.ExpectError(t, "Internal error happened")
}

func TestResourceRepoCreateCustomDirectoryWrongLocation(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceRepo(),
		State: map[string]any{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/Production/test/abc/",
		},
		Create: true,
	}.ExpectError(t, "invalid config supplied. [path] should have 3 components (/Repos/<directory>/<repo>), got 4. Deprecated Resource")
}

func TestResourceRepoCreateCustomDirectoryWrongPath(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceRepo(),
		State: map[string]any{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/test/",
		},
		Create: true,
	}.ExpectError(t, "invalid config supplied. [path] should have 3 components (/Repos/<directory>/<repo>), got 2. Deprecated Resource")
}

func TestResourceRepoCreateWithBranch(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Branch:       "main",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	respPatch := resp
	respPatch.Branch = "releases"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: reposCreateRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
				},
				Response: resp,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]any{"branch": "releases"},
				Response:        respPatch,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: respPatch,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url":    "https://github.com/user/test.git",
			"branch": "releases",
		},
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{"id": resp.RepoID(), "path": resp.Path, "branch": respPatch.Branch,
			"git_provider": resp.Provider, "url": resp.Url, "commit_hash": resp.HeadCommitID})
}

func TestResourceRepoCreateWithTag(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Branch:       "main",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	respPatch := resp
	respPatch.Branch = ""
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: reposCreateRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
				},
				Response: resp,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]any{"tag": "v0.1"},
				Response:        respPatch,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: respPatch,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url": "https://github.com/user/test.git",
			"tag": "v0.1",
		},
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{"id": resp.RepoID(), "path": resp.Path,
			"git_provider": resp.Provider, "url": resp.Url, "commit_hash": resp.HeadCommitID})
}

func TestResourceRepoCreateError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceRepo(),
		State: map[string]any{
			"url": "https://somegit.com/user/test.git",
		},
		Create: true,
	}.ExpectError(t, "git_provider isn't specified and we can't detect provider from URL")
}

func TestResourceReposUpdateSwitchToTag(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]any{"tag": "v0.1"},
				Response:        resp,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
		},
		Resource: ResourceRepo(),
		InstanceState: map[string]string{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"branch":       "main",
		},
		State: map[string]any{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"tag":          "v0.1",
		},
		ID:          "121232342",
		Update:      true,
		RequiresNew: true,
	}.ApplyAndExpectData(t, map[string]any{"branch": ""})
}

func TestResourceReposUpdateSwitchToBranch(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
		Branch:       "releases",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]any{"branch": "releases"},
				Response:        resp,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
		},
		Resource: ResourceRepo(),
		InstanceState: map[string]string{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"branch":       "main",
		},
		State: map[string]any{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"branch":       "releases",
		},
		ID:     "121232342",
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{"branch": "releases"})
}

func TestResourceReposUpdateSparseCheckout(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]any{"branch": "main",
					"sparse_checkout": map[string]any{"patterns": []string{"abc", "def"}},
				},
				Response: resp,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
		},
		Resource: ResourceRepo(),
		InstanceState: map[string]string{
			"url":                        "https://github.com/user/test.git",
			"git_provider":               "gitHub",
			"path":                       "/Repos/user@domain/test",
			"branch":                     "main",
			"sparse_checkout.0.patterns": `["abc"]`,
		},
		HCL: `
			url = "https://github.com/user/test.git"
			git_provider = "gitHub",
			path = "/Repos/user@domain/test"
			branch = "main"
			sparse_checkout{
				patterns = ["abc", "def"]
			}
			`,
		ID:          "121232342",
		Update:      true,
		RequiresNew: true,
	}.ApplyAndExpectData(t, map[string]any{"branch": "main"})
}

func TestReposListAll(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
		Branch:       "releases",
	}

	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/repos?",
			Response: ReposListResponse{
				NextPageToken: "12312423442343242343",
				Repos: []ReposInformation{
					resp,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/repos?next_page_token=12312423442343242343",
			Response: ReposListResponse{
				Repos: []ReposInformation{
					resp,
				},
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	reposList, err := NewReposAPI(ctx, client).ListAll()
	require.NoError(t, err)
	assert.Equal(t, len(reposList), 2)
	assert.Equal(t, resp.Branch, reposList[1].Branch)
}

func TestReposListWithPrefix(t *testing.T) {
	resp := ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
		Branch:       "releases",
	}

	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/repos?path_prefix=%2FRepos%2Fabc",
			Response: ReposListResponse{
				Repos: []ReposInformation{
					resp,
				},
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	reposList, err := NewReposAPI(ctx, client).List("/Repos/abc")
	require.NoError(t, err)
	assert.Equal(t, len(reposList), 1)
	assert.Equal(t, resp.Branch, reposList[0].Branch)
}
