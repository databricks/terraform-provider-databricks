package sdkv2

import (
	"maps"
	"slices"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
)

func TestApiField_RegistrationConsistency(t *testing.T) {
	cases := []struct {
		name    string
		entries map[string]common.Resource
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

func TestNoDuplicateKeys(t *testing.T) {
	testNoDuplicateKeys(t, WorkspaceDataSources, AccountDataSources, DualDataSources)
	testNoDuplicateKeys(t, WorkspaceResources, AccountResources, DualResources)
}

func testNoDuplicateKeys(t *testing.T, ms ...map[string]common.Resource) {
	count := count(ms...)
	for _, key := range slices.Sorted(maps.Keys(count)) {
		if c := count[key]; c > 1 {
			t.Errorf("%q is registered in multiple maps: %v", key, c)
		}
	}
}

func count(ms ...map[string]common.Resource) map[string]int {
	count := map[string]int{}
	for _, m := range ms {
		for key := range m {
			count[key]++
		}
	}
	return count
}
