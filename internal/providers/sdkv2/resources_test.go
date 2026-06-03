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
