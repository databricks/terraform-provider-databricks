package workspace

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestGetGitProviderFromUrl(t *testing.T) {
	assert.Equal(t, "bitbucketCloud", GetGitProviderFromUrl("https://user@bitbucket.org/user/repo.git"))
	assert.Equal(t, "gitHub", GetGitProviderFromUrl("https://github.com//user/repo.git"))
	assert.Equal(t, "azureDevOpsServices", GetGitProviderFromUrl("https://user@dev.azure.com/user/project/_git/repo"))
	//	assert.Equal(t, "bitbucketCloud", GetGitProviderFromUrl("https://user@bitbucket.org/user/repo.git"))
	assert.Equal(t, "", GetGitProviderFromUrl("https://abc/user/repo.git"))
	assert.Equal(t, "", GetGitProviderFromUrl("ewfgwergfwe"))
}

func TestResourceRepoRead(t *testing.T) {
	repoID := 48155820875912
	repoIDStr := fmt.Sprintf("%d", repoID)
	url := "https://user@github.com/user/repo.git"
	provider := "gitHub"
	branch := "main"
	path := "/Repos/Production/testrepo"
	d, err := qa.ResourceFixture{
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
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, repoIDStr, d.Id())
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, branch, d.Get("branch"))
	assert.Equal(t, provider, d.Get("git_provider"))
	assert.Equal(t, url, d.Get("url"))
	assert.Equal(t, "7e0847ede61f07adede22e2bcce6050216489171", d.Get("commit_hash"))
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
	d, err := qa.ResourceFixture{
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
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, repoID, d.Id())
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: createRequest{
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
		State: map[string]interface{}{
			"url": "https://github.com/user/test.git",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, resp.RepoID(), d.Id())
	assert.Equal(t, resp.Branch, d.Get("branch"))
	assert.Equal(t, resp.Provider, d.Get("git_provider"))
	assert.Equal(t, resp.Path, d.Get("path"))
	assert.Equal(t, resp.HeadCommitID, d.Get("commit_hash"))
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: createRequest{
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
		State: map[string]interface{}{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/Production/test/",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, resp.RepoID(), d.Id())
	assert.Equal(t, resp.Branch, d.Get("branch"))
	assert.Equal(t, resp.Provider, d.Get("git_provider"))
	assert.Equal(t, resp.Path, d.Get("path"))
	assert.Equal(t, resp.HeadCommitID, d.Get("commit_hash"))
}

func TestResourceRepoCreateCustomDirectoryError(t *testing.T) {
	_, err := qa.ResourceFixture{
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
		State: map[string]interface{}{
			"url":  "https://github.com/user/test.git",
			"path": "/Repos/Production/test/",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestResourceRepoCreateCustomDirectoryWrongLocation(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceRepo(),
		State: map[string]interface{}{
			"url":  "https://github.com/user/test.git",
			"path": "/NotRepos/Production/test/",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "path should start with /Repos/")
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: createRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
				},
				Response: resp,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]interface{}{"branch": "releases"},
				Response:        respPatch,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: respPatch,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]interface{}{
			"url":    "https://github.com/user/test.git",
			"branch": "releases",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, resp.RepoID(), d.Id())
	assert.Equal(t, respPatch.Branch, d.Get("branch"))
	assert.Equal(t, resp.Provider, d.Get("git_provider"))
	assert.Equal(t, resp.Path, d.Get("path"))
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/repos",
				ExpectedRequest: createRequest{
					Url:      "https://github.com/user/test.git",
					Provider: "gitHub",
				},
				Response: resp,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]interface{}{"tag": "v0.1"},
				Response:        respPatch,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: respPatch,
			},
		},
		Resource: ResourceRepo(),
		State: map[string]interface{}{
			"url": "https://github.com/user/test.git",
			"tag": "v0.1",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, resp.RepoID(), d.Id())
	assert.Equal(t, respPatch.Branch, d.Get("branch"))
	assert.Equal(t, resp.Provider, d.Get("git_provider"))
	assert.Equal(t, resp.Path, d.Get("path"))
}

func TestResourceRepoCreateError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceRepo(),
		State: map[string]interface{}{
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]interface{}{"tag": "v0.1"},
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
		State: map[string]interface{}{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"tag":          "v0.1",
		},
		ID:          "121232342",
		Update:      true,
		RequiresNew: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Get("branch"))
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/repos/121232342",
				ExpectedRequest: map[string]interface{}{"branch": "releases"},
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
		State: map[string]interface{}{
			"url":          "https://github.com/user/test.git",
			"git_provider": "gitHub",
			"path":         "/Repos/user@domain/test",
			"branch":       "releases",
		},
		ID:     "121232342",
		Update: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "releases", d.Get("branch"))
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
