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

func TestNoOverlap(t *testing.T) {
	type group struct {
		name    string
		entries map[string]*schema.Resource
	}
	cases := []struct {
		kind   string
		groups []group
	}{
		{"resources", []group{
			{"WorkspaceResources", WorkspaceResources},
			{"AccountResources", AccountResources},
			{"DualResources", DualResources},
		}},
		{"dataSources", []group{
			{"WorkspaceDataSources", WorkspaceDataSources},
			{"AccountDataSources", AccountDataSources},
			{"DualDataSources", DualDataSources},
		}},
	}
	for _, tc := range cases {
		t.Run(tc.kind, func(t *testing.T) {
			owners := map[string][]string{}
			for _, g := range tc.groups {
				for _, key := range slices.Sorted(maps.Keys(g.entries)) {
					owners[key] = append(owners[key], g.name)
				}
			}
			for _, key := range slices.Sorted(maps.Keys(owners)) {
				if len(owners[key]) > 1 {
					t.Errorf("%q is registered in multiple maps: %v", key, owners[key])
				}
			}
		})
	}
}
