package useragent

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
)

func TestLookupMetaHarnessProvider(t *testing.T) {
	tests := []struct {
		name   string
		envs   map[string]string
		expect string
	}{
		{
			name:   "no harness",
			envs:   nil,
			expect: "",
		},
		{
			name:   "omnigent",
			envs:   map[string]string{"OMNIGENT": "1"},
			expect: "omnigent",
		},
		{
			name:   "omnigent empty value still counts as set",
			envs:   map[string]string{"OMNIGENT": ""},
			expect: "omnigent",
		},
		{
			name:   "agent env var does not affect harness",
			envs:   map[string]string{"CLAUDECODE": "1"},
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			ClearCache()
			for k, v := range tt.envs {
				t.Setenv(k, v)
			}
			got := lookupMetaHarnessProvider()
			if got != tt.expect {
				t.Errorf("lookupMetaHarnessProvider() = %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestMetaHarnessProviderCached(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("OMNIGENT", "1")
	got := MetaHarnessProvider()
	if got != "omnigent" {
		t.Errorf("MetaHarnessProvider() = %q, want %q", got, "omnigent")
	}
	// Change env after caching. Cached result should persist.
	t.Setenv("OMNIGENT", "")
	got = MetaHarnessProvider()
	if got != "omnigent" {
		t.Errorf("MetaHarnessProvider() after cache = %q, want %q", got, "omnigent")
	}
}
