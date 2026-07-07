package pluginfw

import (
	"testing"
)

// TestRolloutLists_NoDuplicateResources guards against a contributor adding a
// resource name to both migratedResources and pluginFwOptInResources, which
// would produce duplicate registrations and undefined mux behaviour.
func TestRolloutLists_NoDuplicateResources(t *testing.T) {
	migrated := make(map[string]struct{}, len(migratedResources))
	for _, f := range migratedResources {
		migrated[getResourceName(f)] = struct{}{}
	}
	for _, f := range pluginFwOptInResources {
		name := getResourceName(f)
		if _, dup := migrated[name]; dup {
			t.Errorf("resource %q appears in both migratedResources and pluginFwOptInResources; pick one rollout stage", name)
		}
	}
}

// TestRolloutLists_NoDuplicateDataSources is the data source counterpart of
// TestRolloutLists_NoDuplicateResources.
func TestRolloutLists_NoDuplicateDataSources(t *testing.T) {
	migrated := make(map[string]struct{}, len(migratedDataSources))
	for _, f := range migratedDataSources {
		migrated[getDataSourceName(f)] = struct{}{}
	}
	for _, f := range pluginFwOptInDataSources {
		name := getDataSourceName(f)
		if _, dup := migrated[name]; dup {
			t.Errorf("data source %q appears in both migratedDataSources and pluginFwOptInDataSources; pick one rollout stage", name)
		}
	}
}

// The third invariant (opt-in entries must exist in the SDKv2 provider map)
// lives in internal/providers/sdkv2/sdkv2_test.go to avoid an import cycle.
