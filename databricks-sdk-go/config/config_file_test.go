package config

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// withMockEnv mocks environment variables for testing config file loading
// without relying on the actual system environment or filesystem.
// getUserHomeDir falls back to the real implementation when HOME is not in
// the env map, allowing tests to optionally override HOME without breaking
// tests that don't need to control the home directory path.
func withMockEnv(t *testing.T, env map[string]string) {
	original := getenv
	originalUserHomeDir := getUserHomeDir
	t.Cleanup(func() {
		getenv = original
		getUserHomeDir = originalUserHomeDir
	})
	getenv = func(key string) string {
		return env[key]
	}
	getUserHomeDir = func() (string, error) {
		if home, ok := env["HOME"]; ok {
			return home, nil
		}
		return originalUserHomeDir()
	}
}

func TestConfigFileLoad(t *testing.T) {
	f, err := LoadFile("testdata/.databrickscfg")
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}
	if f == nil {
		t.Fatal("expected file to be non-nil")
	}

	for _, name := range []string{
		"password-with-double-quotes",
		"password-with-single-quotes",
		"password-without-quotes",
	} {
		section := f.Section(name)
		if section == nil {
			t.Fatalf("expected section %q to be non-nil", name)
		}
		if got, want := section.Key("password").String(), "%Y#X$Z"; got != want {
			t.Errorf("password mismatch for %q: got %q, want %q", name, got, want)
		}
	}
}

func TestConfigFile_Scopes(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    []string
	}{
		{
			name:    "empty defaults to all-apis",
			profile: "scope-empty",
			want:    []string{"all-apis"},
		},
		{
			name:    "single scope",
			profile: "scope-single",
			want:    []string{"clusters"},
		},
		{
			name:    "multiple scopes sorted",
			profile: "scope-multiple",
			want:    []string{"clusters", "files:read", "iam:read", "jobs", "mlflow", "model-serving", "pipelines"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withMockEnv(t, map[string]string{
				"HOME": "testdata/scopes",
			})

			cfg := &Config{Profile: tt.profile, HTTPTransport: metadataNotFoundTransport}
			err := cfg.EnsureResolved()
			if err != nil {
				t.Fatalf("EnsureResolved failed: %v", err)
			}
			if diff := cmp.Diff(tt.want, cfg.GetScopes()); diff != "" {
				t.Errorf("GetScopes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfigFile_GroupIDLoadingAndPrecedence(t *testing.T) {
	var groupIDAttribute *ConfigAttribute
	for i := range ConfigAttributes {
		if ConfigAttributes[i].Name == "group_id" {
			groupIDAttribute = &ConfigAttributes[i]
			break
		}
	}
	if groupIDAttribute == nil {
		t.Fatal("group_id configuration attribute not found")
	}

	testCases := []struct {
		name        string
		profile     string
		direct      string
		environment string
		want        string
		wantSource  SourceType
	}{
		{
			name:       "loads group ID from selected profile",
			profile:    "group",
			want:       "profile-group",
			wantSource: SourceFile,
		},
		{
			name:       "profile without group ID leaves value empty",
			profile:    "normal",
			want:       "",
			wantSource: SourceDynamicConfig,
		},
		{
			name:        "environment group ID overrides profile",
			profile:     "group",
			environment: "environment-group",
			want:        "environment-group",
			wantSource:  SourceEnv,
		},
		{
			name:        "direct group ID overrides environment and profile",
			profile:     "group",
			direct:      "direct-group",
			environment: "environment-group",
			want:        "direct-group",
			wantSource:  SourceDynamicConfig,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			withMockEnv(t, map[string]string{
				"HOME":                "testdata/group_id",
				"DATABRICKS_GROUP_ID": testCase.environment,
			})

			cfg := &Config{
				Profile: testCase.profile,
				GroupID: testCase.direct,
			}

			if err := ConfigAttributes.Configure(cfg); err != nil {
				t.Fatalf("ConfigAttributes.Configure(): %v", err)
			}
			if err := ConfigFile.Configure(cfg); err != nil {
				t.Fatalf("ConfigFile.Configure(): %v", err)
			}

			if cfg.GroupID != testCase.want {
				t.Errorf("GroupID = %q, want %q", cfg.GroupID, testCase.want)
			}
			if got := cfg.getSource(groupIDAttribute).Type; got != testCase.wantSource {
				t.Errorf("group_id source = %q, want %q", got, testCase.wantSource)
			}
		})
	}
}

// Test 1: default_profile resolves correctly (no [DEFAULT] section present)
func TestConfigFile_DefaultProfileResolves(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_no_default",
		},
		AssertAuth: "pat",
		AssertHost: "https://my-workspace.cloud.databricks.com",
	}.apply(t)
}

func TestConfigFile_DefaultProfileSetsResolvedProfileName(t *testing.T) {
	env.CleanupEnvironment(t)

	cfg, err := configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_no_default",
		},
	}.configureProviderAndReturnConfig(t)

	require.NoError(t, err)
	assert.Equal(t, "my-workspace", cfg.Profile)
}

// Test 2: default_profile takes precedence over [DEFAULT]
func TestConfigFile_DefaultProfileTakesPrecedenceOverDefault(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile",
		},
		AssertAuth: "pat",
		AssertHost: "https://my-workspace.cloud.databricks.com",
	}.apply(t)
}

// Test 3: Legacy fallback when no [__settings__]
func TestConfigFile_LegacyFallbackNoSettings(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata",
		},
		AssertAuth: "pat",
		AssertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
	}.apply(t)
}

// Test 4: Legacy fallback when default_profile is empty
func TestConfigFile_LegacyFallbackEmptyDefaultProfile(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_empty",
		},
		AssertAuth: "pat",
		AssertHost: "https://default.cloud.databricks.com",
	}.apply(t)
}

// Test 4a: The legacy fallback to [DEFAULT] now pins cfg.Profile to
// "DEFAULT". Callers (e.g. the Databricks CLI's per-profile OAuth cache key)
// need a non-empty cfg.Profile to derive a stable key across login and
// read flows. Previously the loader silently loaded [DEFAULT]'s values
// but left cfg.Profile empty, producing a cache-key drift between an
// explicit `--profile DEFAULT` write and a no-flag read.
func TestConfigFile_LegacyFallbackSetsResolvedProfileName(t *testing.T) {
	env.CleanupEnvironment(t)

	cfg, err := configFixture{
		Env: map[string]string{
			"HOME": "testdata",
		},
	}.configureProviderAndReturnConfig(t)

	require.NoError(t, err)
	assert.Equal(t, "DEFAULT", cfg.Profile)
	assert.Equal(t, "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com", cfg.Host)
}

// Test 4b: default_profile = DEFAULT is treated as an explicit selection,
// not as the legacy fallback. cfg.Profile is set to "DEFAULT" and a missing
// [DEFAULT] section would error (unlike the silent legacy fallback).
func TestConfigFile_ExplicitDefaultProfileViaSetting(t *testing.T) {
	env.CleanupEnvironment(t)

	cfg, err := configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_explicit_default",
		},
	}.configureProviderAndReturnConfig(t)

	require.NoError(t, err)
	assert.Equal(t, "DEFAULT", cfg.Profile)
	assert.Equal(t, "https://default.cloud.databricks.com", cfg.Host)
}

// Test 5: [__settings__] is not a profile.
// resolveProfile never returns "__settings__" as a profile name.
func TestConfigFile_SettingsSectionNotAProfile(t *testing.T) {
	f, err := LoadFile("testdata/default_profile/.databrickscfg")
	require.NoError(t, err)

	profile, _, err := resolveProfile("", f)
	require.NoError(t, err)
	assert.NotEqual(t, settingsSection, profile,
		"resolveProfile must never return the settings section name as a profile")

	// The ini library includes __settings__ as a section. Callers that
	// enumerate profiles must filter it out.
	assert.Contains(t, f.SectionStrings(), settingsSection)
}

// Test 5b: default_profile = __settings__ is rejected because __settings__
// is a reserved section name.
func TestConfigFile_DefaultProfileSettingsSelfReference(t *testing.T) {
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_self_ref",
		},
		AssertError: fmt.Sprintf("resolve: %s: __settings__ is a reserved section name and cannot be used as a profile",
			"testdata/default_profile_self_ref/.databrickscfg"),
	}.apply(t)
}

// Test 6: Explicit --profile overrides default_profile
func TestConfigFile_ExplicitProfileOverridesDefaultProfile(t *testing.T) {
	configFixture{
		Profile: "other",
		Env: map[string]string{
			"HOME": "testdata/default_profile",
		},
		AssertAuth: "pat",
		AssertHost: "https://other.cloud.databricks.com",
	}.apply(t)
}

func TestConfigFile_ExplicitSettingsSectionProfileRejected(t *testing.T) {
	env.CleanupEnvironment(t)

	_, err := configFixture{
		Profile: "__settings__",
		Env: map[string]string{
			"HOME": "testdata/default_profile",
		},
	}.configureProviderAndReturnConfig(t)

	require.Error(t, err)
	assert.ErrorContains(t, err, fmt.Sprintf("resolve: %s: __settings__ is a reserved section name and cannot be used as a profile",
		"testdata/default_profile/.databrickscfg"))
}

// Test 7: default_profile pointing to nonexistent section
func TestConfigFile_DefaultProfileNonexistentSection(t *testing.T) {
	env.CleanupEnvironment(t)
	configFixture{
		Env: map[string]string{
			"HOME": "testdata/default_profile_nonexistent",
		},
		AssertError: fmt.Sprintf("resolve: %s has no deleted-profile profile configured",
			"testdata/default_profile_nonexistent/.databrickscfg"),
	}.apply(t)
}
