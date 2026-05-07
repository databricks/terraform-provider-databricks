package subprocenv

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// envMap is a tiny helper that turns the []string env slice (each entry
// "KEY=VALUE") into a map for assertions. Same semantics terraform-exec
// uses, so we can test against the actual shape we hand to SetEnv.
func envMap(env []string) map[string]string {
	out := make(map[string]string, len(env))
	for _, e := range env {
		k, v, ok := strings.Cut(e, "=")
		if !ok {
			continue
		}
		out[k] = v
	}
	return out
}

// envHasKey reports whether env contains an entry with the given key,
// regardless of value. Useful for asserting strict absence (a key not
// being present in env at all is a stronger guarantee than a key being
// present with an empty value).
func envHasKey(env []string, key string) bool {
	prefix := key + "="
	for _, e := range env {
		if strings.HasPrefix(e, prefix) {
			return true
		}
	}
	return false
}

// withEnv sets env vars for the duration of the test and restores them
// after. Uses t.Setenv where applicable so cleanup happens automatically.
// For "explicit unset" we use os.Unsetenv via a t.Cleanup that restores.
func withEnv(t *testing.T, vars map[string]string) {
	t.Helper()
	for k, v := range vars {
		t.Setenv(k, v)
	}
}

// withUnset clears the named env vars for the duration of the test, even
// if they aren't set in the parent environment. t.Setenv("", "") would
// also do this but is less self-documenting.
func withUnset(t *testing.T, keys ...string) {
	t.Helper()
	for _, k := range keys {
		t.Setenv(k, "")
	}
}

// TestBuild_AllowlistedParentVarsPropagate confirms each entry in the
// curated allowlist round-trips from parent env into the subprocess env
// when set. The test sets every allowlist key to a sentinel and verifies
// each one shows up.
func TestBuild_AllowlistedParentVarsPropagate(t *testing.T) {
	want := map[string]string{}
	for _, k := range AllowedKeys() {
		want[k] = "sentinel-" + k
	}
	withEnv(t, want)

	env := Build("PROFILE", "/tmp/.tfrc", "/tmp/run", nil)
	got := envMap(env)
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s: got %q want %q", k, got[k], v)
		}
	}
}

// TestBuild_StripsDatabricksLeakageVars locks in the B5 invariant: every
// DATABRICKS_* set in the parent env must be stripped from the subprocess
// env, EXCEPT the two we set ourselves (DATABRICKS_CONFIG_PROFILE and
// DATABRICKS_CONFIG_FILE).
func TestBuild_StripsDatabricksLeakageVars(t *testing.T) {
	leakers := []string{
		"DATABRICKS_HOST",
		"DATABRICKS_TOKEN",
		"DATABRICKS_USERNAME",
		"DATABRICKS_PASSWORD",
		"DATABRICKS_CLIENT_ID",
		"DATABRICKS_CLIENT_SECRET",
		"DATABRICKS_AUTH_TYPE",
		"DATABRICKS_ACCOUNT_ID",
		"DATABRICKS_AZURE_TENANT_ID",
		"DATABRICKS_GOOGLE_SERVICE_ACCOUNT",
	}
	for _, k := range leakers {
		t.Setenv(k, "leaked-via-"+k)
	}

	env := Build("PROFILE", "/tmp/.tfrc", "/tmp/run", nil)
	for _, k := range leakers {
		if envHasKey(env, k) {
			t.Errorf("expected %s to be stripped, but it appeared in subprocess env", k)
		}
	}

	got := envMap(env)
	if got["DATABRICKS_CONFIG_PROFILE"] != "PROFILE" {
		t.Errorf("DATABRICKS_CONFIG_PROFILE: got %q want PROFILE", got["DATABRICKS_CONFIG_PROFILE"])
	}
}

// TestBuild_StripsTFLeakageVars confirms TF_LOG and TF_VAR_* (and other
// non-allowlisted TF vars) don't propagate. F2/G14 — TF_LOG=DEBUG would
// dilute step logs.
func TestBuild_StripsTFLeakageVars(t *testing.T) {
	leakers := []string{
		"TF_LOG",
		"TF_LOG_PATH",
		"TF_VAR_password",
		"TF_REGISTRY_DISCOVERY_RETRY",
	}
	for _, k := range leakers {
		t.Setenv(k, "leaked-via-"+k)
	}

	env := Build("p", "/tmp/.tfrc", "/tmp/run", nil)
	for _, k := range leakers {
		if envHasKey(env, k) {
			t.Errorf("expected %s to be stripped, but it appeared in subprocess env", k)
		}
	}
}

// TestBuild_StripsCloudCredentialVars verifies cloud creds aren't leaked.
// The user opts into them via passthrough_env when needed.
func TestBuild_StripsCloudCredentialVars(t *testing.T) {
	leakers := []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
		"ARM_CLIENT_ID",
		"ARM_CLIENT_SECRET",
		"ARM_TENANT_ID",
		"ARM_SUBSCRIPTION_ID",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"GOOGLE_CREDENTIALS",
	}
	for _, k := range leakers {
		t.Setenv(k, "leaked-via-"+k)
	}

	env := Build("p", "/tmp/.tfrc", "/tmp/run", nil)
	for _, k := range leakers {
		if envHasKey(env, k) {
			t.Errorf("expected %s to be stripped without passthrough, got %q", k, envMap(env)[k])
		}
	}
}

// TestBuild_FrameworkControlledVarsAreSet pins the values we always set:
// TF_CLI_CONFIG_FILE, TF_IN_AUTOMATION, TF_PLUGIN_CACHE_DIR,
// DATABRICKS_CONFIG_PROFILE, DATABRICKS_CONFIG_FILE.
func TestBuild_FrameworkControlledVarsAreSet(t *testing.T) {
	t.Setenv("HOME", "/home/testuser")

	env := Build("MY_PROFILE", "/run/.terraformrc", "/run/test-abc", nil)
	got := envMap(env)

	for _, tc := range []struct {
		key  string
		want string
	}{
		{"TF_CLI_CONFIG_FILE", "/run/.terraformrc"},
		{"TF_IN_AUTOMATION", "1"},
		{"TF_PLUGIN_CACHE_DIR", filepath.Join("/run/test-abc", "plugins")},
		{"DATABRICKS_CONFIG_PROFILE", "MY_PROFILE"},
		{"DATABRICKS_CONFIG_FILE", filepath.Join("/home/testuser", ".databrickscfg")},
	} {
		if got[tc.key] != tc.want {
			t.Errorf("%s: got %q want %q", tc.key, got[tc.key], tc.want)
		}
	}
}

// TestBuild_PassthroughEnv_OptInWorks confirms test.yaml's passthrough_env
// list adds named parent-env values into the subprocess env.
func TestBuild_PassthroughEnv_OptInWorks(t *testing.T) {
	t.Setenv("AWS_PROFILE", "saml-prod")
	t.Setenv("GCP_PROJECT", "my-gcp-project")
	t.Setenv("AZURE_CLIENT_ID", "abc-123")

	env := Build("p", "/tmp/.tfrc", "/tmp/run", []string{"AWS_PROFILE", "GCP_PROJECT", "AZURE_CLIENT_ID"})
	got := envMap(env)
	for k, want := range map[string]string{
		"AWS_PROFILE":     "saml-prod",
		"GCP_PROJECT":     "my-gcp-project",
		"AZURE_CLIENT_ID": "abc-123",
	} {
		if got[k] != want {
			t.Errorf("passthrough %s: got %q want %q", k, got[k], want)
		}
	}
}

// TestBuild_PassthroughEnv_DropsEmpty documents that names with no value
// in the parent env are silently dropped (per DESIGN.md §10 G6).
func TestBuild_PassthroughEnv_DropsEmpty(t *testing.T) {
	withUnset(t, "TFV2_TEST_NEVER_SET_VAR")
	env := Build("p", "/tmp/.tfrc", "/tmp/run", []string{"TFV2_TEST_NEVER_SET_VAR"})
	if envHasKey(env, "TFV2_TEST_NEVER_SET_VAR") {
		t.Errorf("expected unset passthrough to be dropped, got %v", envMap(env))
	}
}

// TestBuild_PassthroughEnv_RejectsDatabricksPrefix locks in the
// defense-in-depth check: even if a user puts DATABRICKS_HOST in their
// passthrough_env, it must NOT leak. The config layer also rejects this
// at parse time, but subprocenv is the last line of defense.
func TestBuild_PassthroughEnv_RejectsDatabricksPrefix(t *testing.T) {
	t.Setenv("DATABRICKS_HOST", "https://leaked.example.com")
	t.Setenv("DATABRICKS_TOKEN", "leaked")

	env := Build("p", "/tmp/.tfrc", "/tmp/run", []string{"DATABRICKS_HOST", "DATABRICKS_TOKEN"})
	got := envMap(env)
	if got["DATABRICKS_HOST"] != "" {
		t.Errorf("DATABRICKS_HOST leaked via passthrough: %q", got["DATABRICKS_HOST"])
	}
	if got["DATABRICKS_TOKEN"] != "" {
		t.Errorf("DATABRICKS_TOKEN leaked via passthrough: %q", got["DATABRICKS_TOKEN"])
	}
	// Sanity: our framework-controlled DATABRICKS_CONFIG_PROFILE is
	// still set. The reject-prefix logic must not over-reach.
	if got["DATABRICKS_CONFIG_PROFILE"] != "p" {
		t.Errorf("framework-controlled DATABRICKS_CONFIG_PROFILE clobbered: %q", got["DATABRICKS_CONFIG_PROFILE"])
	}
}

// TestBuild_PassthroughEnv_PreservesOrder ensures the passthrough entries
// appear in the order test.yaml declared them. Order matters for log
// readability — surprising re-ordering would make env diffs noisy.
func TestBuild_PassthroughEnv_PreservesOrder(t *testing.T) {
	t.Setenv("ZZZ_FIRST", "1")
	t.Setenv("AAA_SECOND", "2")
	t.Setenv("MMM_THIRD", "3")

	env := Build("p", "/tmp/.tfrc", "/tmp/run", []string{"ZZZ_FIRST", "AAA_SECOND", "MMM_THIRD"})

	// Find indices of the three passthroughs in the resulting env slice.
	var idx [3]int
	for i, e := range env {
		switch {
		case strings.HasPrefix(e, "ZZZ_FIRST="):
			idx[0] = i
		case strings.HasPrefix(e, "AAA_SECOND="):
			idx[1] = i
		case strings.HasPrefix(e, "MMM_THIRD="):
			idx[2] = i
		}
	}
	if !(idx[0] < idx[1] && idx[1] < idx[2]) {
		t.Errorf("passthrough order not preserved: %v", idx)
	}
}

// TestBuild_AllowedKeysReturnsCopy confirms AllowedKeys returns a defensive
// copy so callers can't mutate the package-level allowlist.
func TestBuild_AllowedKeysReturnsCopy(t *testing.T) {
	a := AllowedKeys()
	b := AllowedKeys()
	if !slices.Equal(a, b) {
		t.Fatalf("two AllowedKeys() calls disagree: %v vs %v", a, b)
	}
	a[0] = "MUTATED"
	c := AllowedKeys()
	if c[0] == "MUTATED" {
		t.Errorf("AllowedKeys must return a copy, got mutation through caller: %v", c)
	}
}

// TestBuild_SkipsUnsetAllowlistEntries documents that PATH not being set
// in the parent env results in PATH not being in the subprocess env (we
// do NOT propagate empty values — exec resolves binaries against the
// host's PATH config in that case which is the correct fallback).
func TestBuild_SkipsUnsetAllowlistEntries(t *testing.T) {
	withUnset(t, "TMPDIR")
	env := Build("p", "/tmp/.tfrc", "/tmp/run", nil)
	if envHasKey(env, "TMPDIR") {
		t.Errorf("expected unset TMPDIR to NOT appear in subprocess env, got %q", envMap(env)["TMPDIR"])
	}
}
