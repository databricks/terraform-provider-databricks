package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fixtureProfile writes a minimal .databrickscfg containing the named
// section so config.Load's profile-existence preflight passes.
func fixtureProfile(t *testing.T, name string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, ".databrickscfg")
	body := "[" + name + "]\nhost = https://accounts.cloud.databricks.com\n"
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write databrickscfg: %v", err)
	}
	return path
}

// loadString parses a YAML body string. Most tests use this rather than
// writing fixture files because the test bodies are short and the YAML
// content is the thing under test.
func loadString(t *testing.T, body, profileName string) (*TestSpec, error) {
	t.Helper()
	dir := t.TempDir()
	yamlPath := filepath.Join(dir, "test.yaml")
	if err := os.WriteFile(yamlPath, []byte(body), 0o600); err != nil {
		t.Fatalf("write test.yaml: %v", err)
	}
	cfgPath := fixtureProfile(t, profileName)
	return LoadWithProfilePath(yamlPath, cfgPath)
}

// TestLoad_FullSchema parses the issue #5672 mission-test fixture
// verbatim. A regression here breaks the only file we point real
// terraform at in M5.
func TestLoad_FullSchema(t *testing.T) {
	body := `
name: issue_5672_mws_workspaces_account_provider_config_regression
profile: ACCOUNT_AWS
cleanup: true

requires:
  cloud: any
  level: account

steps:
  - name: passes_on_1_113_0
    version: "1.113.0"
    command: plan
    expect: success

  - name: fails_on_1_114_0
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'cannot populate provider_config for mws workspaces.*failed to resolve workspace_id'

  - name: fixed_on_1_114_1
    version: "1.114.1"
    command: plan
    expect: success

  - name: fixed_on_local
    version: "local"
    command: plan
    expect: success
`
	spec, err := loadString(t, body, "ACCOUNT_AWS")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if spec.Name != "issue_5672_mws_workspaces_account_provider_config_regression" {
		t.Errorf("Name: got %q", spec.Name)
	}
	if spec.Profile != "ACCOUNT_AWS" {
		t.Errorf("Profile: got %q", spec.Profile)
	}
	if !spec.CleanupEnabled() {
		t.Errorf("CleanupEnabled: expected true")
	}
	if spec.Requires.Cloud != CloudAny {
		t.Errorf("Requires.Cloud: got %q", spec.Requires.Cloud)
	}
	if spec.Requires.Level != LevelAccount {
		t.Errorf("Requires.Level: got %q", spec.Requires.Level)
	}
	if len(spec.Steps) != 4 {
		t.Fatalf("expected 4 steps, got %d", len(spec.Steps))
	}
	// Step 2 (failure) must have a compiled regex with the expected source.
	if spec.Steps[1].CompiledRegex == nil {
		t.Errorf("step 1 (fails_on_1_114_0): expected compiled regex")
	} else if got := spec.Steps[1].CompiledRegex.String(); got != `cannot populate provider_config for mws workspaces.*failed to resolve workspace_id` {
		t.Errorf("step 1 regex: got %q", got)
	}
	// Step 4 must use the local literal version.
	if spec.Steps[3].Version != "local" {
		t.Errorf("step 3 version: got %q want 'local'", spec.Steps[3].Version)
	}
	// Default timeout applied across all steps.
	for i, s := range spec.Steps {
		if s.Timeout != DefaultStepTimeout {
			t.Errorf("step %d timeout: got %s want %s", i, s.Timeout, DefaultStepTimeout)
		}
	}
}

// TestLoad_DefaultsApplied confirms the defaulting rules: cleanup=true,
// requires.cloud=any, requires.level=workspace, command=apply,
// expect=success, timeout=10m.
func TestLoad_DefaultsApplied(t *testing.T) {
	body := `
name: minimal
profile: P
steps:
  - name: only
    version: "1.114.0"
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !spec.CleanupEnabled() {
		t.Errorf("Cleanup default: expected true")
	}
	if spec.Requires.Cloud != CloudAny {
		t.Errorf("Cloud default: got %q want any", spec.Requires.Cloud)
	}
	if spec.Requires.Level != LevelWorkspace {
		t.Errorf("Level default: got %q want workspace", spec.Requires.Level)
	}
	s := spec.Steps[0]
	if s.Command != CommandApply {
		t.Errorf("Command default: got %q want apply", s.Command)
	}
	if s.Expect != ExpectSuccess {
		t.Errorf("Expect default: got %q want success", s.Expect)
	}
	if s.Timeout != DefaultStepTimeout {
		t.Errorf("Timeout default: got %s want %s", s.Timeout, DefaultStepTimeout)
	}
}

// TestLoad_CleanupExplicitFalse documents the pointer-vs-default trick:
// `cleanup: false` MUST stay false even though the zero value of *bool
// is nil.
func TestLoad_CleanupExplicitFalse(t *testing.T) {
	body := `
name: nocleanup
profile: P
cleanup: false
steps:
  - name: only
    version: "1.114.0"
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if spec.CleanupEnabled() {
		t.Errorf("expected CleanupEnabled=false when explicit cleanup: false")
	}
}

func TestLoad_TimeoutOverride(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    timeout: 30s
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := spec.Steps[0].Timeout; got != 30*time.Second {
		t.Errorf("Timeout: got %s", got)
	}
}

func TestLoad_RejectsUnknownFields(t *testing.T) {
	body := `
name: t
profile: P
cleanups: true   # typo: should be 'cleanup'
steps:
  - name: a
    version: "1.114.0"
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for unknown field, got nil")
	}
	if !strings.Contains(err.Error(), "cleanups") {
		t.Errorf("expected error to mention 'cleanups', got: %v", err)
	}
}

func TestLoad_RejectsMultipleDocs(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - {name: a, version: "1.114.0"}
---
name: u
profile: P
steps:
  - {name: b, version: "1.114.0"}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for multiple docs, got nil")
	}
}

// TestLoad_ProfileMustExist checks the preflight: a profile name not
// present in the .databrickscfg file fails parse-time.
func TestLoad_ProfileMustExist(t *testing.T) {
	body := `
name: t
profile: NONEXISTENT
steps:
  - name: a
    version: "1.114.0"
`
	dir := t.TempDir()
	yamlPath := filepath.Join(dir, "test.yaml")
	_ = os.WriteFile(yamlPath, []byte(body), 0o600)

	cfg := filepath.Join(dir, ".databrickscfg")
	_ = os.WriteFile(cfg, []byte("[OTHER]\nhost = https://x.cloud.databricks.com\n"), 0o600)

	_, err := LoadWithProfilePath(yamlPath, cfg)
	if err == nil {
		t.Fatal("expected error when profile missing, got nil")
	}
	if !strings.Contains(err.Error(), "NONEXISTENT") {
		t.Errorf("expected error to mention NONEXISTENT, got: %v", err)
	}
}

func TestLoad_FailureRequiresAssertion(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    expect: failure
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for failure-without-assertion, got nil")
	}
	if !strings.Contains(err.Error(), "error_substring or error_regex") {
		t.Errorf("expected friendly assertion message, got: %v", err)
	}
}

func TestLoad_SuccessWithErrorFieldsRejected(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    expect: success
    error_substring: oops
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error: error_substring + expect=success is invalid")
	}
}

func TestLoad_RegexCompiledOrError(t *testing.T) {
	good := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    expect: failure
    error_regex: 'foo.*bar'
`
	bad := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    expect: failure
    error_regex: '['
`
	if _, err := loadString(t, good, "P"); err != nil {
		t.Errorf("good regex rejected: %v", err)
	}
	if _, err := loadString(t, bad, "P"); err == nil {
		t.Errorf("bad regex accepted")
	}
}

func TestLoad_BothErrorFieldsAllowed(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.114.0"
    expect: failure
    error_substring: 'precondition failed'
    error_regex: 'Error:.*precondition failed.*resource X'
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	s := spec.Steps[0]
	if s.ErrorSubstring == "" || s.CompiledRegex == nil {
		t.Errorf("expected both substring and regex set: %+v", s)
	}
}

func TestLoad_RejectsInvalidVersion(t *testing.T) {
	for _, v := range []string{
		"1",
		"1.0",
		"v1.0.0", // no leading v allowed in test.yaml — providercache normalizes upstream
		"1.0.0a", // not strict semver
		"latest", // not a recognized literal
		"01.2.3", // leading zero forbidden
	} {
		body := `
name: t
profile: P
steps:
  - name: a
    version: "` + v + `"
`
		t.Run(v, func(t *testing.T) {
			_, err := loadString(t, body, "P")
			if err == nil {
				t.Errorf("version %q accepted, expected rejection", v)
			}
		})
	}
}

func TestLoad_AcceptsValidVersions(t *testing.T) {
	for _, v := range []string{"1.0.0", "1.114.1", "0.1.0", "1.10.0", "99.0.0-local", "local"} {
		body := `
name: t
profile: P
steps:
  - name: a
    version: "` + v + `"
`
		t.Run(v, func(t *testing.T) {
			if _, err := loadString(t, body, "P"); err != nil {
				t.Errorf("version %q rejected: %v", v, err)
			}
		})
	}
}

func TestLoad_RejectsBadEnums(t *testing.T) {
	for _, tc := range []struct {
		name    string
		body    string
		wantSub string
	}{
		{
			"bad cloud",
			`
name: t
profile: P
requires:
  cloud: oracle
steps:
  - {name: a, version: "1.0.0"}
`, "requires.cloud",
		},
		{
			"bad level",
			`
name: t
profile: P
requires:
  level: customer
steps:
  - {name: a, version: "1.0.0"}
`, "requires.level",
		},
		{
			"bad command",
			`
name: t
profile: P
steps:
  - {name: a, version: "1.0.0", command: refresh}
`, "command",
		},
		{
			"bad expect",
			`
name: t
profile: P
steps:
  - {name: a, version: "1.0.0", expect: maybe}
`, "expect",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := loadString(t, tc.body, "P")
			if err == nil || !strings.Contains(err.Error(), tc.wantSub) {
				t.Errorf("expected error containing %q, got: %v", tc.wantSub, err)
			}
		})
	}
}

func TestLoad_RejectsDuplicateStepName(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - {name: a, version: "1.0.0"}
  - {name: a, version: "1.1.0"}
`
	_, err := loadString(t, body, "P")
	if err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("expected duplicate-step error, got: %v", err)
	}
}

func TestLoad_RejectsEmptySteps(t *testing.T) {
	body := `
name: t
profile: P
steps: []
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for empty steps")
	}
}

func TestLoad_RejectsBadSlugs(t *testing.T) {
	// Note: "trailing-space " gets normalized by the YAML parser (plain
	// scalars strip trailing whitespace per the YAML spec), so we don't
	// test that case — it's the parser's job, not ours.
	for _, n := range []string{"Has Space", "UPPER", "has.dot", ""} {
		body := `
name: ` + n + `
profile: P
steps:
  - {name: a, version: "1.0.0"}
`
		t.Run(n, func(t *testing.T) {
			_, err := loadString(t, body, "P")
			if err == nil {
				t.Errorf("name %q accepted, expected rejection", n)
			}
		})
	}
}

// TestLoad_RejectsDatabricksPassthrough locks in the §10/G6 invariant in
// the config layer (subprocenv has its own defense-in-depth check).
func TestLoad_RejectsDatabricksPassthrough(t *testing.T) {
	body := `
name: t
profile: P
passthrough_env:
  - DATABRICKS_HOST
steps:
  - {name: a, version: "1.0.0"}
`
	_, err := loadString(t, body, "P")
	if err == nil || !strings.Contains(err.Error(), "DATABRICKS_") {
		t.Errorf("expected DATABRICKS_-rejection error, got: %v", err)
	}
}

func TestLoad_RejectsEmptyPassthroughName(t *testing.T) {
	body := `
name: t
profile: P
passthrough_env:
  - ""
  - AWS_PROFILE
steps:
  - {name: a, version: "1.0.0"}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for empty passthrough name")
	}
}

// TestLoad_PassthroughEnvAccepted exercises the typical good-shape case.
func TestLoad_PassthroughEnvAccepted(t *testing.T) {
	body := `
name: t
profile: P
passthrough_env:
  - AWS_PROFILE
  - AZURE_CLIENT_ID
  - GCP_PROJECT
steps:
  - {name: a, version: "1.0.0"}
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(spec.PassthroughEnv) != 3 {
		t.Errorf("expected 3 passthrough names, got %v", spec.PassthroughEnv)
	}
}

// TestLoadDir confirms the convenience wrapper resolves <dir>/test.yaml.
func TestLoadDir(t *testing.T) {
	dir := t.TempDir()
	body := `
name: t
profile: P
steps:
  - {name: a, version: "1.0.0"}
`
	if err := os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	cfg := filepath.Join(t.TempDir(), ".databrickscfg")
	_ = os.WriteFile(cfg, []byte("[P]\nhost = x\n"), 0o600)

	t.Setenv("DATABRICKS_CONFIG_FILE", cfg)
	spec, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir: %v", err)
	}
	if spec.Name != "t" {
		t.Errorf("Name: got %q", spec.Name)
	}
}

func TestLoad_RejectsZeroOrNegativeTimeout(t *testing.T) {
	for _, d := range []string{"0s", "-5m"} {
		body := `
name: t
profile: P
steps:
  - {name: a, version: "1.0.0", timeout: ` + d + `}
`
		t.Run(d, func(t *testing.T) {
			_, err := loadString(t, body, "P")
			if err == nil {
				t.Errorf("timeout %q accepted, expected rejection", d)
			}
		})
	}
}

// TestLoad_FixtureFromMissionTest exercises the actual on-disk
// `issues-repro/issue_5672/test.yaml` (resolved relative to the
// test source). We point at a synthetic .databrickscfg matching the
// fixture's profile name so the existence check passes.
func TestLoad_FixtureFromMissionTest(t *testing.T) {
	// Path relative to this test file: ../../issues-repro/issue_5672/test.yaml
	yamlPath := filepath.Join("..", "..", "issues-repro", "issue_5672", "test.yaml")
	if _, err := os.Stat(yamlPath); err != nil {
		t.Skipf("mission-test fixture not available: %v", err)
	}
	cfg := fixtureProfile(t, "ACCOUNT_AWS")
	spec, err := LoadWithProfilePath(yamlPath, cfg)
	if err != nil {
		t.Fatalf("LoadWithProfilePath: %v", err)
	}
	if got := len(spec.Steps); got != 4 {
		t.Errorf("expected 4 steps in mission test, got %d", got)
	}
}
