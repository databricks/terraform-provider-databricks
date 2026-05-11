package settings

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testDisableLegacyFeatures = AllSettingsResources()["disable_legacy_features"]

func TestCreateDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag1",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Create:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestReadDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag1",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: false,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Read:     true,
		HCL: `
			disable_legacy_features {
    			value = "false"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, false, d.Get("disable_legacy_features.0.value"))
}

func TestUpdateDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Update:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestUpdateDisableLegacyFeaturesWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag2",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag2",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag3",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag3",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag3",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Update:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestDeleteDisableLegacyFeatures(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything,
				settings.DeleteDisableLegacyFeaturesRequest{
					Etag: "etag1",
				}).Return(&settings.DeleteDisableLegacyFeaturesResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Delete:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestDeleteDisableLegacyFeaturesWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyFeaturesRequest{
				Etag: "etag1",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag2",
					},
				}},
			})
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDisableLegacyFeaturesResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		Delete: true,
		ID:     defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag3",
	})
}

// TestDisableLegacyFeaturesSetting_SkipProviderConfigStatePopulation is a
// regression guard for the bug where the post-Read provider_config hook ran
// against an account-only setting and failed with
// "cannot populate provider_config for disable legacy features setting:
//
//	failed to resolve workspace_id: ... Unable to load OAuth Config".
//
// Account-only settings (built via accountSetting in generic_setting.go) must
// have SkipProviderConfigStatePopulation set so the hook is short-circuited.
// See https://github.com/databricks/terraform-provider-databricks/issues/5672
// (and the parallel fix in #5689 for account-only mws data sources).
func TestDisableLegacyFeaturesSetting_SkipProviderConfigStatePopulation(t *testing.T) {
	assert.True(t, testDisableLegacyFeatures.SkipProviderConfigStatePopulation,
		"databricks_disable_legacy_features_setting must opt out of post-Read provider_config population")
	pc, ok := testDisableLegacyFeatures.Schema["provider_config"]
	assert.True(t, ok, "provider_config block must still exist in the schema (kept for state compatibility)")
	assert.NotEmpty(t, pc.Deprecated, "provider_config block must be marked deprecated for account-only settings")
}

// TestDisableLegacyFeaturesSetting_AccountLevelNoHookFailure is the end-to-end
// regression test the reviewer asked for: it drives the full Create path
// against an account-host HTTP server WITHOUT mocking the /Me endpoint
// that the post-Create provider_config hook would call.
//
// Without the fix in this PR, the post-Create hook runs
// `populateProviderConfigInState` → `CurrentWorkspaceID` → `/Me`, which has
// no fixture and so the test fails with the exact error users hit in
// production:
//
//	cannot populate provider_config for disable_legacy_features_setting:
//	failed to resolve workspace_id: ...
//
// With the fix, SkipProviderConfigStatePopulation = true short-circuits the
// hook and the Create flow returns cleanly. Mirrors the pattern from
// PR #5689's TestDataSourceMwsWorkspaces_AccountLevelNoHookFailure.
func TestDisableLegacyFeaturesSetting_AccountLevelNoHookFailure(t *testing.T) {
	apiPath := "/api/2.0/accounts/abc/settings/types/disable_legacy_features/names/default"
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "PATCH",
			Resource:     apiPath,
			ReuseRequest: true,
			Response: settings.DisableLegacyFeatures{
				Etag:        "etag1",
				SettingName: "disable_legacy_features",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
			},
		},
		{
			Method:       "GET",
			Resource:     apiPath + "?etag=etag1",
			ReuseRequest: true,
			Response: settings.DisableLegacyFeatures{
				Etag:        "etag1",
				SettingName: "disable_legacy_features",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
			},
		},
		// Deliberately NO fixture for /api/2.0/preview/scim/v2/Me — if the
		// post-Create provider_config hook fires, it will try to call /Me
		// and the test will fail.
	})
	assert.NoError(t, err)
	defer server.Close()

	// Simulate a real account-level provider: AccountID set, no cached or
	// configured workspace_id. Without the fix this is exactly the
	// configuration that fails inside populateProviderConfigInState.
	client.Config.AccountID = "abc"
	client.SetCachedWorkspaceID(0)

	r := testDisableLegacyFeatures.ToResource()
	d := r.TestResourceData()
	err = d.Set("disable_legacy_features", []any{map[string]any{"value": true}})
	assert.NoError(t, err)
	ctx := context.WithValue(context.Background(), common.ResourceName, "disable_legacy_features_setting")
	diags := r.CreateContext(ctx, d, client)

	assert.False(t, diags.HasError(),
		"post-Create provider_config hook must be skipped for account-only setting; got: %v", diags)
	assert.Equal(t, defaultSettingId, d.Id())
}
