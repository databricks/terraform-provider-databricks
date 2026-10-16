package repos

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
				Delete(mock.Anything, workspace.DeleteCredentialsRequest{
					CredentialId: credID,
				}).
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

func TestResourceGitCredentialCreateWithPrincipalId(t *testing.T) {
	provider := "gitHub"
	user := "test"
	token := "12345"
	principalID := 123456789
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
				PrincipalId:         int64(principalID),
			}).
				Return(&resp, nil)
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{
				CredentialId: resp.CredentialId,
				PrincipalId:  int64(principalID),
			}).
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
			"principal_id":          principalID,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           fmt.Sprintf("%d", resp.CredentialId),
		"git_provider": provider,
		"git_username": user,
		"principal_id": principalID,
	})
}

func TestResourceGitCredentialDeleteWithPrincipalId(t *testing.T) {
	credID := int64(48155820875912)
	credIDStr := fmt.Sprintf("%d", credID)
	principalID := int64(123456789)

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGitCredentialsAPI().EXPECT().
				Delete(mock.Anything, workspace.DeleteCredentialsRequest{
					CredentialId: credID,
					PrincipalId:  principalID,
				}).
				Return(nil)
		},
		Resource: ResourceGitCredential(),
		Delete:   true,
		ID:       credIDStr,
		HCL: fmt.Sprintf(`
		git_provider          = "gitHub"
		git_username          = "test"
		personal_access_token = "12345"
		principal_id          = %d
		`, principalID),
	}.ApplyAndExpectData(t, map[string]any{"id": credIDStr})
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
			gmock.ListAll(mock.Anything, workspace.ListCredentialsRequest{}).
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
			gmock.ListAll(mock.Anything, workspace.ListCredentialsRequest{}).
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
			gmock.ListAll(mock.Anything, workspace.ListCredentialsRequest{}).
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
				ListAll(mock.Anything, workspace.ListCredentialsRequest{}).
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
			gmock.ListAll(mock.Anything, workspace.ListCredentialsRequest{}).
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

func TestResourceGitCredentialCreate_WriteOnlyToken(t *testing.T) {
	for _, k := range []string{"GITHUB_TOKEN", "GITLAB_TOKEN", "AZDO_PERSONAL_ACCESS_TOKEN"} {
		t.Setenv(k, "")
	}
	provider := "gitHub"
	user := "test"
	token := "write-only-token"
	resp := workspace.CreateCredentialsResponse{
		CredentialId: 121232342,
		GitProvider:  provider,
		GitUsername:  user,
	}

	d, err := qa.ResourceFixture{
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
			"git_provider":                     provider,
			"git_username":                     user,
			"personal_access_token_wo":         token,
			"personal_access_token_wo_version": 1,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%d", resp.CredentialId), d.Id())
	assert.Equal(t, "", d.Get("personal_access_token"))
	assert.Equal(t, 1, d.Get("personal_access_token_wo_version"))
}

func TestResourceGitCredentialUpdate_WriteOnlyTokenVersionChange(t *testing.T) {
	for _, k := range []string{"GITHUB_TOKEN", "GITLAB_TOKEN", "AZDO_PERSONAL_ACCESS_TOKEN"} {
		t.Setenv(k, "")
	}
	credID := int64(121232342)
	provider := "gitHub"
	user := "test"
	token := "write-only-token-v2"

	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			gmock := w.GetMockGitCredentialsAPI().EXPECT()
			gmock.Update(mock.Anything, workspace.UpdateCredentialsRequest{
				CredentialId:        credID,
				GitProvider:         provider,
				GitUsername:         user,
				PersonalAccessToken: token,
			}).
				Return(nil)
			gmock.Get(mock.Anything, workspace.GetCredentialsRequest{CredentialId: credID}).
				Return(&workspace.GetCredentialsResponse{
					CredentialId: credID,
					GitProvider:  provider,
					GitUsername:  user,
				}, nil)
		},
		Resource: ResourceGitCredential(),
		InstanceState: map[string]string{
			"git_provider":                     provider,
			"git_username":                     user,
			"personal_access_token_wo_version": "1",
		},
		State: map[string]any{
			"git_provider":                     provider,
			"git_username":                     user,
			"personal_access_token_wo":         token,
			"personal_access_token_wo_version": 2,
		},
		ID:     "121232342",
		Update: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "", d.Get("personal_access_token"))
	assert.Equal(t, 2, d.Get("personal_access_token_wo_version"))
}

func TestResourceGitCredential_WriteOnlyTokenConflictsWithToken(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":                     "gitHub",
			"personal_access_token":            "token",
			"personal_access_token_wo":         "write-only-token",
			"personal_access_token_wo_version": 1,
		},
		Create: true,
	}.Apply(t)
	assert.ErrorContains(t, err, "conflicts with personal_access_token")
}

func TestResourceGitCredential_WriteOnlyTokenRequiresVersion(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceGitCredential(),
		State: map[string]any{
			"git_provider":             "gitHub",
			"personal_access_token_wo": "write-only-token",
		},
		Create: true,
	}.Apply(t)
	assert.ErrorContains(t, err, "must be specified")
}

func TestResourceGitCredential_WriteOnlyTokenSuppressesEnvDefault(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "from-env")
	t.Setenv("GITLAB_TOKEN", "")
	t.Setenv("AZDO_PERSONAL_ACCESS_TOKEN", "")

	full := ResourceGitCredential().Schema
	r := &schema.Resource{Schema: map[string]*schema.Schema{}}
	for _, k := range []string{"git_provider", "personal_access_token",
		"personal_access_token_wo", "personal_access_token_wo_version"} {
		r.Schema[k] = full[k]
	}
	plan := func(raw map[string]any) *terraform.InstanceDiff {
		// Mirror the gRPC shim, which exposes the raw configuration to DiffSuppressFunc.
		vals := map[string]cty.Value{}
		for k, v := range raw {
			switch tv := v.(type) {
			case string:
				vals[k] = cty.StringVal(tv)
			case int:
				vals[k] = cty.NumberIntVal(int64(tv))
			}
		}
		ctyVal, err := r.CoreConfigSchema().CoerceValue(cty.ObjectVal(vals))
		require.NoError(t, err)
		cfg := terraform.NewResourceConfigRaw(raw)
		cfg.CtyValue = ctyVal
		diff, err := r.Diff(context.Background(), &terraform.InstanceState{}, cfg, nil)
		require.NoError(t, err)
		return diff
	}

	diff := plan(map[string]any{
		"git_provider":                     "gitHub",
		"personal_access_token_wo":         "write-only-token",
		"personal_access_token_wo_version": 1,
	})
	assert.NotContains(t, diff.Attributes, "personal_access_token")

	diff = plan(map[string]any{"git_provider": "gitHub"})
	require.Contains(t, diff.Attributes, "personal_access_token")
	assert.Equal(t, "from-env", diff.Attributes["personal_access_token"].New)
}

func TestResourceGitCredentialPATIsSensitive(t *testing.T) {
	resource := ResourceGitCredential()
	s := resource.Schema["personal_access_token"]
	if !s.Sensitive {
		t.Error("personal_access_token should be marked as Sensitive")
	}
}

func TestGitCredentialCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceGitCredential(),
		qa.CornerCaseSkipCRUD("create"),
		qa.CornerCaseExpectError(`strconv.ParseInt: parsing "x": invalid syntax`))
}
