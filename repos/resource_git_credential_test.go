package repos

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestResourceGitCredentialRead(t *testing.T) {
	credID := int64(48155820875912)
	credIDStr := fmt.Sprintf("%d", credID)
	provider := "gitHub"
	user := "test"

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(&workspace.GetCredentialsResponse{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}, nil)
		},
		Resource: ResourceGitCredential(),
		Read:     true,
		New:      true,
		ID:       credIDStr,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           credIDStr,
		"git_provider": provider,
		"git_username": user,
	})
}

func TestResourceGitCredentialRead_Error(t *testing.T) {
	credID := int64(48155820875912)
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(nil, errors.New("Git credential with the given ID could not be found."))
		},
		Resource: ResourceGitCredential(),
		Read:     true,
		ID:       fmt.Sprintf("%d", credID),
	}.ExpectError(t, "Git credential with the given ID could not be found.")
}

func TestResourceGitCredentialDelete(t *testing.T) {
	credID := int64(48155820875912)
	credIDStr := fmt.Sprintf("%d", credID)

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				DeleteByCredentialId(mock.Anything, credID).
				Return(nil)
		},
		Resource: ResourceGitCredential(),
		Delete:   true,
		ID:       credIDStr,
	}.ApplyAndExpectData(t, map[string]any{"id": credIDStr})
}

func TestResourceGitCredentialUpdate(t *testing.T) {
	credID := int64(121232342)
	provider := "gitHub"
	user := "test"
	token := "1234"
	resp := workspace.GetCredentialsResponse{
		CredentialId: credID,
		GitProvider:  provider,
		GitUsername:  user,
	}

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Update(mock.Anything, workspace.UpdateCredentialsRequest{
				CredentialId:         credID,
				GitProvider:          provider,
				GitUsername:          user,
				PersonalAccessToken:  token,
				IsDefaultForProvider: false,
				ForceSendFields:      []string{"IsDefaultForProvider"},
			}).
				Return(nil)
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(&resp, nil)
		},
		Resource: ResourceGitCredential(),
		InstanceState: map[string]string{
			"git_provider":            provider,
			"git_username":            user,
			"personal_access_token":   token + "56",
			"is_default_for_provider": "true",
		},
		State: map[string]any{
			"git_provider":            provider,
			"git_username":            user,
			"personal_access_token":   token,
			"is_default_for_provider": false,
		},
		ID:     "121232342",
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{"git_username": user})
}

func TestResourceGitCredentialUpdate_Error(t *testing.T) {
	credID := int64(121232342)
	provider := "gitHub"
	user := "test"
	token := "1234"

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				Update(mock.Anything, workspace.UpdateCredentialsRequest{
					CredentialId:        credID,
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				}).
				Return(errors.New("Git credential with the given ID could not be found."))
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(&resp, nil)
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: resp.CredentialId}).
				Return(&workspace.GetCredentialsResponse{
					CredentialId: resp.CredentialId,
					GitProvider:  provider,
					GitUsername:  user,
				}, nil)
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           fmt.Sprintf("%d", resp.CredentialId),
		"git_provider": provider,
		"git_username": user,
	})
}

func TestResourceGitCredentialCreate_Error_OnlyOneGitCredential(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				Create(mock.Anything, workspace.CreateCredentialsRequest{
					GitProvider:         provider,
					GitUsername:         user,
					PersonalAccessToken: token,
				}).
				Return(nil, errors.New("Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint."))
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

func TestResourceGitCredentialCreateWithForceOld(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	credID := int64(121232342)
	resp := workspace.GetCredentialsResponse{
		CredentialId: credID,
		GitProvider:  provider,
		GitUsername:  user,
	}

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(&workspace.CreateCredentialsResponse{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}, errors.New("Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint."))
			gmock.ListAll(mock.Anything).
				Return([]workspace.CredentialInfo{{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}}, nil)
			gmock.Update(mock.Anything, workspace.UpdateCredentialsRequest{
				CredentialId:        credID,
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil)
			// After Create, Terraform calls Read to populate state
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(&resp, nil)
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           fmt.Sprintf("%d", resp.CredentialId),
		"git_provider": provider,
		"git_username": user,
	})
}

func TestResourceGitCredentialCreateWithForceNew(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	credID := int64(121232342)
	resp := workspace.GetCredentialsResponse{
		CredentialId: credID,
		GitProvider:  provider,
		GitUsername:  user,
	}

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(&workspace.CreateCredentialsResponse{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}, apierr.ErrResourceConflict)
			gmock.ListAll(mock.Anything).
				Return([]workspace.CredentialInfo{{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}}, nil)
			gmock.Update(mock.Anything, workspace.UpdateCredentialsRequest{
				CredentialId:        credID,
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil)
			// After Create, Terraform calls Read to populate state
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(&resp, nil)
		},
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":          provider,
			"git_username":          user,
			"personal_access_token": token,
			"force":                 true,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           fmt.Sprintf("%d", resp.CredentialId),
		"git_provider": provider,
		"git_username": user,
	})
}
func TestResourceGitCredentialCreateWithForce_Error_List(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil, errors.New("Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint."))
			gmock.ListAll(mock.Anything).
				Return(nil, errors.New("No such endpoint"))
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil, errors.New("Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint."))
			w.GetMockGitCredentialsAPI().EXPECT().
				ListAll(mock.Anything).
				Return([]workspace.CredentialInfo{}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Create(mock.Anything, workspace.CreateCredentialsRequest{
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil, errors.New("Only one Git credential is supported at this time. If you would like to update your credential, please use the PATCH endpoint."))
			gmock.ListAll(mock.Anything).
				Return([]workspace.CredentialInfo{resp}, nil)
			gmock.Update(mock.Anything, workspace.UpdateCredentialsRequest{
				CredentialId:        resp.CredentialId,
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(errors.New("Git credential with the given ID could not be found."))
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
