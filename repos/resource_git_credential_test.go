package repos

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceGitCredentialRead(t *testing.T) {
	credID := 48155820875912
	credIDStr := fmt.Sprintf("%d", credID)
	provider := "gitHub"
	user := "test"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d?", credID),
				Response: workspace.CredentialInfo{
					CredentialId: int64(credID),
					GitProvider:  provider,
					GitUsername:  user,
				},
			},
		},
		Resource: ResourceGitCredential(),
		Read:     true,
		New:      true,
		ID:       credIDStr,
	}.ApplyAndExpectData(t, map[string]any{"id": credIDStr, "git_provider": provider, "git_username": user})
}

func TestResourceGitCredentialRead_Error(t *testing.T) {
	credID := 48155820875912
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d?", credID),
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Git credential with the given ID could not be found.",
				},
				Status: 404,
			},
		},
		Resource: ResourceGitCredential(),
		Read:     true,
		ID:       fmt.Sprintf("%d", credID),
	}.ExpectError(t, "resource is not expected to be removed")
}

func TestResourceGitCredentialDelete(t *testing.T) {
	credID := "48155820875912"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%s?", credID),
				Status:   http.StatusOK,
			},
		},
		Resource: ResourceGitCredential(),
		Delete:   true,
		ID:       credID,
	}.ApplyAndExpectData(t, map[string]any{"id": credID})
}

func TestResourceGitCredentialUpdate(t *testing.T) {
	credID := 121232342
	provider := "gitHub"
	user := "test"
	token := "1234"
	resp := workspace.CredentialInfo{
		CredentialId: int64(credID),
		GitProvider:  provider,
		GitUsername:  user,
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d", credID),
				ExpectedRequest: workspace.UpdateCredentials{
					CredentialId:        int64(credID),
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: resp,
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d?", credID),
				Response: resp,
			},
		},
		Resource: ResourceGitCredential(),
		InstanceState: map[string]string{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token + "56",
		},
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
		},
		ID:     "121232342",
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{"git_username": user})
}

func TestResourceGitCredentialUpdate_Error(t *testing.T) {
	credID := 121232342
	provider := "gitHub"
	user := "test"
	token := "1234"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d", credID),
				ExpectedRequest: workspace.UpdateCredentials{
					CredentialId:        int64(credID),
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Git credential with the given ID could not be found.",
				},
				Status: 404,
			},
		},
		Resource: ResourceGitCredential(),
		InstanceState: map[string]string{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token + "56",
		},
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
		},
		ID:     "121232342",
		Update: true,
	}.ExpectError(t, "Git credential with the given ID could not be found.")
}

func TestResourceGitCredentialCreate(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	resp := workspace.CreateCredentialsResponse{
		CredentialId: 121232342,
		GitProvider:  provider,
		GitUsername:  user,
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: resp,
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d?", resp.CredentialId),
				Response: resp,
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": fmt.Sprintf("%d", resp.CredentialId), "git_provider": provider, "git_username": user})
}

func TestResourceGitCredentialCreate_Error(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.",
				},
				Status: 400,
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
		},
		Create: true,
	}.ExpectError(t, "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.")
}

func TestResourceGitCredentialCreateWithForce(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	resp := workspace.CredentialInfo{
		CredentialId: 121232342,
		GitProvider:  provider,
		GitUsername:  user,
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.",
				},
				Status: 400,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/git-credentials",
				Response: workspace.GetCredentialsResponse{
					Credentials: []workspace.CredentialInfo{resp},
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d", resp.CredentialId),
				ExpectedRequest: workspace.UpdateCredentials{
					CredentialId:        resp.CredentialId,
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: resp,
			},
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d?", resp.CredentialId),
				Response: resp,
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": fmt.Sprintf("%d", resp.CredentialId), "git_provider": provider, "git_username": user})
}

func TestResourceGitCredentialCreateWithForce_Error_List(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.",
				},
				Status: 400,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/git-credentials",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "No such endpoint",
				},
				Status: 404,
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ExpectError(t, "No such endpoint")
}

func TestResourceGitCredentialCreateWithForce_ErrorEmptyList(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.",
				},
				Status: 400,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/git-credentials",
				Response: workspace.GetCredentialsResponse{},
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ExpectError(t, "list of credentials is either empty or have more than one entry (0)")
}

func TestResourceGitCredentialCreateWithForce_ErrorUpdate(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	resp := workspace.CredentialInfo{
		CredentialId: 121232342,
		GitProvider:  provider,
		GitUsername:  user,
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/git-credentials",
				ExpectedRequest: workspace.CreateCredentials{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint.",
				},
				Status: 400,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/git-credentials",
				Response: workspace.GetCredentialsResponse{
					Credentials: []workspace.CredentialInfo{resp},
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: fmt.Sprintf("/api/2.0/git-credentials/%d", resp.CredentialId),
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Git credential with the given ID could not be found.",
				},
				Status: 404,
			},
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ExpectError(t, "Git credential with the given ID could not be found.")
}

func TestGitCredentialCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceGitCredential(),
		qa.CornerCaseSkipCRUD("create"),
		qa.CornerCaseExpectError(`strconv.ParseInt: parsing "x": invalid syntax`))
}
