package sdkv2

import (
	"maps"
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestApiField_RegistrationConsistency(t *testing.T) {
	cases := []struct {
		name    string
		entries map[string]*schema.Resource
		wantApi bool
	}{
		{"DualResources", DualResources, true},
		{"DualDataSources", DualDataSources, true},
		{"WorkspaceResources", WorkspaceResources, false},
		{"AccountResources", AccountResources, false},
		{"WorkspaceDataSources", WorkspaceDataSources, false},
		{"AccountDataSources", AccountDataSources, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range slices.Sorted(maps.Keys(tc.entries)) {
				hasApi := tc.entries[key].Schema["api"] != nil
				switch {
				case tc.wantApi && !hasApi:
					t.Errorf("%s[%q] is missing the \"api\" field", tc.name, key)
				case !tc.wantApi && hasApi:
					t.Errorf("%s[%q] has an \"api\" field; move it to DualResources/DualDataSources", tc.name, key)
				}
			}
		})
	}
}

// TestProviderConfig_AccountRegistrationConsistency asserts that account
// resources and data sources do not carry the provider_config block: they have
// no workspace context, so a workspace_id has no meaning for them.
//
// A small set of account-only data sources defined via the deprecated
// common.DataResource helper still expose a (deprecated) provider_config block
// for backward compatibility. These are pinned in exceptions so that no new
// account resource silently gains the block.
func TestProviderConfig_AccountRegistrationConsistency(t *testing.T) {
	cases := []struct {
		name    string
		entries map[string]*schema.Resource
		// exceptions are keys that retain a (deprecated) provider_config block.
		exceptions map[string]bool
	}{
		{"AccountResources", AccountResources, nil},
		{"AccountDataSources", AccountDataSources, map[string]bool{
			"databricks_mws_credentials": true,
			"databricks_mws_workspaces":  true,
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range slices.Sorted(maps.Keys(tc.entries)) {
				hasProviderConfig := tc.entries[key].Schema["provider_config"] != nil
				switch {
				case hasProviderConfig && !tc.exceptions[key]:
					t.Errorf("%s[%q] has a %q block; account resources/data sources have no workspace context", tc.name, key, "provider_config")
				case !hasProviderConfig && tc.exceptions[key]:
					t.Errorf("%s[%q] no longer has a %q block; remove it from the exceptions list", tc.name, key, "provider_config")
				}
			}
		})
	}
}

func TestNoDuplicateKeys(t *testing.T) {
	testNoDuplicateKeys(t, WorkspaceDataSources, AccountDataSources, DualDataSources)
	testNoDuplicateKeys(t, WorkspaceResources, AccountResources, DualResources)
}

func testNoDuplicateKeys(t *testing.T, ms ...map[string]*schema.Resource) {
	count := count(ms...)
	for _, key := range slices.Sorted(maps.Keys(count)) {
		if c := count[key]; c > 1 {
			t.Errorf("%q is registered in multiple maps: %v", key, c)
		}
	}
}

func count(ms ...map[string]*schema.Resource) map[string]int {
	count := map[string]int{}
	for _, m := range ms {
		for key := range m {
			count[key]++
		}
	}
	return count
}
