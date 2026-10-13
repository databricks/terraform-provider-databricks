package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
)

// fixtureProfile writes a minimal.databrickscfg containing the named
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
// present in the.databrickscfg file fails parse-time.
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
	// Pin the sentinel-wrap so callers (notably fixtures_test.go's
	// runFixture) can distinguish "profile missing in this env"
	// (→ t.Skip) from other config errors (→ t.Fatal).
	if !errors.Is(err, profile.ErrSectionNotFound) {
		t.Errorf("expected errors.Is(err, profile.ErrSectionNotFound) to be true, got: %v", err)
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
// test source). We point at a synthetic.databrickscfg matching the
// fixture's profile name so the existence check passes.
func TestLoad_FixtureFromMissionTest(t *testing.T) {
	// Path relative to this test file:../../issues-repro/issue_5672/test.yaml
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

// TestLoad_AllShippedFixturesParse loads every committed test.yaml
// under testframeworkV2/{issues-repro,tests}/ and asserts each one
// parses + validates cleanly. This catches schema drift in any
// shipped fixture without requiring per-fixture wiring — adding a
// new fixture under either tree is automatically covered.
//
// Per-fixture profile names are stubbed via fixtureProfile so the
// existence preflight passes without needing the developer's real
// ~/.databrickscfg to define WORKSPACE_AZURE_SP_UNASSIGNED etc.
func TestLoad_AllShippedFixturesParse(t *testing.T) {
	roots := []string{
		filepath.Join("..", "..", "issues-repro"),
		filepath.Join("..", "..", "tests"),
	}
	for _, root := range roots {
		entries, err := os.ReadDir(root)
		if err != nil {
			// A repo where a tree doesn't yet exist (e.g. an early
			// branch where issues-repro/ is empty) is fine — skip.
			t.Logf("skip %s: %v", root, err)
			continue
		}
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			yamlPath := filepath.Join(root, e.Name(), "test.yaml")
			if _, err := os.Stat(yamlPath); err != nil {
				continue
			}
			t.Run(filepath.Base(root)+"/"+e.Name(), func(t *testing.T) {
				// Peek at the profile name so the synthetic
				//.databrickscfg can stub the right section.
				body, err := os.ReadFile(yamlPath)
				if err != nil {
					t.Fatalf("ReadFile: %v", err)
				}
				profileName := extractYAMLField(string(body), "profile")
				if profileName == "" {
					t.Fatal("could not extract `profile` from test.yaml")
				}
				cfg := fixtureProfile(t, profileName)
				spec, err := LoadWithProfilePath(yamlPath, cfg)
				if err != nil {
					t.Fatalf("LoadWithProfilePath: %v", err)
				}
				if len(spec.Steps) == 0 {
					t.Errorf("expected at least 1 step, got 0")
				}
			})
		}
	}
}

// extractYAMLField returns the value of `key:` from a YAML document
// using a tiny line-based heuristic. We don't import yaml.v3 here
// because the only callers are the per-fixture parse tests above —
// the heuristic only needs to handle top-level scalar fields, which
// it does correctly for the shape every committed fixture uses.
func extractYAMLField(body, key string) string {
	for line := range strings.SplitSeq(body, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}
		k, v, ok := strings.Cut(trimmed, ":")
		if !ok || strings.TrimSpace(k) != key {
			continue
		}
		// Strip trailing comment if any.
		v = strings.TrimSpace(v)
		if i := strings.Index(v, "#"); i >= 0 {
			v = strings.TrimSpace(v[:i])
		}
		return v
	}
	return ""
}

// ═══════════════════════════════════════════════════════════
// v2-mode config tests
// ═══════════════════════════════════════════════════════════

// TestLoad_V2_ConfigField parses a minimal v2 spec — every step has
// a `config:` field and the runner-visible IsV2() flag flips true.
func TestLoad_V2_ConfigField(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", config: step1.tf }
  - { name: b, version: "1.0.0", config: step2.tf }
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !spec.IsV2() {
		t.Errorf("IsV2 should be true when steps declare config:")
	}
	if spec.Steps[0].Config != "step1.tf" {
		t.Errorf("Config: got %q", spec.Steps[0].Config)
	}
}

// TestLoad_V2_AssertField confirms the assert: list survives YAML
// decoding into the Assertion struct + the IsV2 detection path.
func TestLoad_V2_AssertField(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step1.tf
    command: apply
    assert:
      - resource: databricks_token.pat
        attrs:
          comment: hello
          lifetime_seconds: 1800
      - resource: databricks_token.pat
        present: false
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !spec.IsV2() {
		t.Errorf("IsV2 should be true with config: set")
	}
	if got := spec.Mode(); got != ModeV2 {
		t.Errorf("Mode: got %q want %q", got, ModeV2)
	}
	if got := len(spec.Steps[0].Assert); got != 2 {
		t.Fatalf("expected 2 assertions, got %d", got)
	}
	if spec.Steps[0].Assert[0].Resource != "databricks_token.pat" {
		t.Errorf("Assert[0].Resource: got %q", spec.Steps[0].Assert[0].Resource)
	}
	// Default Present is true (omitted YAML field → nil pointer →
	// PresentValue() returns true).
	if !spec.Steps[0].Assert[0].PresentValue() {
		t.Errorf("Assert[0].PresentValue() should default to true")
	}
	// Explicit false survives.
	if spec.Steps[0].Assert[1].PresentValue() {
		t.Errorf("Assert[1].PresentValue() with explicit false should be false")
	}
	// YAML int decodes as int (not float64).
	if got := spec.Steps[0].Assert[0].Attrs["lifetime_seconds"]; got != 1800 {
		t.Errorf("Attrs[lifetime_seconds]: got %v (%T), want 1800 (int)", got, got)
	}
	if got := spec.Steps[0].Assert[0].Attrs["comment"]; got != "hello" {
		t.Errorf("Attrs[comment]: got %v", got)
	}
}

// TestLoad_V2_RejectsMixedV1V2 enforces validateV2Consistency: if
// ANY step uses Config or Assert, ALL steps must use Config.
func TestLoad_V2_RejectsMixedV1V2(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", config: step1.tf }
  - { name: b, version: "1.0.0" }
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for v1/v2 mix, got nil")
	}
	if !strings.Contains(err.Error(), "v2 mode requires every step to set `config:`") {
		t.Errorf("error: %v", err)
	}
}

// TestLoad_V2_RejectsAssertOnPlanFailure pins the rule: state
// assertions only make sense after a successful command that
// produces state — not on expect=failure.
func TestLoad_V2_RejectsAssertOnFailure(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step1.tf
    command: apply
    expect: failure
    error_substring: oops
    assert:
      - resource: x.y
        attrs: {comment: foo}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for assert on expect=failure, got nil")
	}
}

// TestLoad_V2_RejectsAssertOnV1Spec covers rule 3:
// assert: requires v2 mode. A step without `config:` setting
// `assert:` is a parse error.
func TestLoad_V2_RejectsAssertOnV1Spec(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    command: apply
    assert:
      - resource: x.y
        attrs: {foo: bar}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error: assert without config (v1 mode) should be rejected")
	}
	if !strings.Contains(err.Error(), "v2 mode") {
		t.Errorf("error should mention v2 mode: %v", err)
	}
}

// TestLoad_V2_RejectsTfv2PrefixedConfig covers user
// config files must not collide with the framework's `_tfv2_` namespace.
func TestLoad_V2_RejectsTfv2PrefixedConfig(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", config: _tfv2_collision.tf }
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for _tfv2_-prefixed config")
	}
	if !strings.Contains(err.Error(), "_tfv2_") {
		t.Errorf("error should mention _tfv2_ prefix: %v", err)
	}
}

// TestLoad_V2_RejectsPresentFalseWithAttrs covers §17.7 rule 6:
// present:false + attrs:set is logically inconsistent.
func TestLoad_V2_RejectsPresentFalseWithAttrs(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step.tf
    command: apply
    assert:
      - resource: x.y
        present: false
        attrs: {foo: bar}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Fatal("expected error for present:false + attrs")
	}
}

// TestLoad_V2_RejectsBadResourceAddress covers §17.5: root-module
// only addresses, no module-scoped or malformed shapes.
func TestLoad_V2_RejectsBadResourceAddress(t *testing.T) {
	for _, addr := range []string{
		"module.foo.databricks_x.y", // module-scoped — deferred to v3
		"databricks_x",              // missing name
		".databricks_x.y",           // leading dot
		"data.databricks_x",         // data missing name
	} {
		body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step.tf
    command: apply
    assert:
      - resource: ` + addr + `
        attrs: {foo: bar}
`
		t.Run(addr, func(t *testing.T) {
			_, err := loadString(t, body, "P")
			if err == nil {
				t.Errorf("address %q accepted, expected rejection", addr)
			}
		})
	}
}

// TestLoad_V2_AcceptsValidResourceAddresses covers the addresses
// the regex SHOULD accept — managed + data-source root-module
// shapes.
func TestLoad_V2_AcceptsValidResourceAddresses(t *testing.T) {
	for _, addr := range []string{
		"databricks_token.pat",
		"databricks_token.MyResource_42",
		"data.databricks_mws_workspaces.all",
		"data.databricks_x.snake_case_name",
	} {
		body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step.tf
    command: apply
    assert:
      - resource: ` + addr + `
        attrs: {foo: bar}
`
		t.Run(addr, func(t *testing.T) {
			if _, err := loadString(t, body, "P"); err != nil {
				t.Errorf("address %q rejected: %v", addr, err)
			}
		})
	}
}

// TestLoad_V2_RejectsTraversalConfigPaths blocks "../escape.tf" and
// other shapes that would let a test reach outside its directory.
func TestLoad_V2_RejectsTraversalConfigPaths(t *testing.T) {
	for _, p := range []string{
		"../escape.tf",
		"sub/dir.tf",
		".hidden.tf",
		"no-extension",
		"wrong.txt",
	} {
		body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", config: ` + p + ` }
`
		t.Run(p, func(t *testing.T) {
			_, err := loadString(t, body, "P")
			if err == nil {
				t.Errorf("config %q accepted, expected rejection", p)
			}
		})
	}
}

// TestLoad_V2_AcceptsValidConfigPaths covers the set of basenames
// the v2ConfigPathRegexp does accept.
func TestLoad_V2_AcceptsValidConfigPaths(t *testing.T) {
	for _, p := range []string{
		"step1.tf",
		"step1_create.tf",
		"step-2.tf",
		"main.tf.json",
	} {
		body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", config: ` + p + ` }
`
		t.Run(p, func(t *testing.T) {
			if _, err := loadString(t, body, "P"); err != nil {
				t.Errorf("config %q rejected: %v", p, err)
			}
		})
	}
}

// TestLoad_V2_RejectsAssertWithoutResource enforces the minimum
// `resource:` requirement for each entry in assert:.
func TestLoad_V2_RejectsAssertWithoutResource(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    config: step.tf
    command: apply
    assert:
      - attrs: {foo: bar}
`
	_, err := loadString(t, body, "P")
	if err == nil {
		t.Errorf("expected error for missing resource: field")
	}
}

// TestIsV2_FalseForV1 confirms specs that never touch v2 fields
// remain in v1 mode.
func TestIsV2_FalseForV1(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0" }
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if spec.IsV2() {
		t.Errorf("IsV2 should be false for v1 spec")
	}
}

// TestLoadDir_RejectsMissingConfigFile pins the parse-time existence
// check added in v6.1: for v2 specs, every step's
// `config:` MUST reference an existing file under the test dir. A
// typo at step N must NOT let steps 1..N-1 execute (and mutate real
// cloud resources) before failing.
func TestLoadDir_RejectsMissingConfigFile(t *testing.T) {
	dir := t.TempDir()
	body := `name: t
profile: P
steps:
  - { name: s1, version: "1.0.0", config: nonexistent.tf, command: apply }
`
	if err := os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o600); err != nil {
		t.Fatalf("write test.yaml: %v", err)
	}
	cfg := filepath.Join(t.TempDir(), ".databrickscfg")
	_ = os.WriteFile(cfg, []byte("[P]\nhost = x\n"), 0o600)
	t.Setenv("DATABRICKS_CONFIG_FILE", cfg)

	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("expected error for missing config file, got nil")
	}
	if !strings.Contains(err.Error(), "nonexistent.tf") || !strings.Contains(err.Error(), "not found") {
		t.Errorf("error should identify the missing file: %v", err)
	}
}

// TestLoadDir_AcceptsExistingConfigFiles is the happy-path counterpart:
// when every step's `config:` resolves to a real file, LoadDir succeeds.
func TestLoadDir_AcceptsExistingConfigFiles(t *testing.T) {
	dir := t.TempDir()
	body := `name: t
profile: P
steps:
  - { name: s1, version: "1.0.0", config: step1.tf, command: apply }
  - { name: s2, version: "1.0.0", config: step2.tf, command: apply }
`
	if err := os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o600); err != nil {
		t.Fatalf("write test.yaml: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "step1.tf"), []byte("# step 1\n"), 0o600); err != nil {
		t.Fatalf("write step1.tf: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "step2.tf"), []byte("# step 2\n"), 0o600); err != nil {
		t.Fatalf("write step2.tf: %v", err)
	}
	cfg := filepath.Join(t.TempDir(), ".databrickscfg")
	_ = os.WriteFile(cfg, []byte("[P]\nhost = x\n"), 0o600)
	t.Setenv("DATABRICKS_CONFIG_FILE", cfg)

	spec, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir: %v", err)
	}
	if !spec.IsV2() {
		t.Errorf("expected v2 spec, got v1")
	}
}

// TestLoadDir_V1NoExistenceCheck confirms LoadDir does NOT stat config:
// paths for v1 specs (where steps don't set `config:`). Pure backward-
// compat: existing v1 fixtures must continue to load without any
// per-step.tf file requirement.
func TestLoadDir_V1NoExistenceCheck(t *testing.T) {
	dir := t.TempDir()
	body := `name: t
profile: P
steps:
  - { name: s1, version: "1.0.0", command: apply }
`
	if err := os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	cfg := filepath.Join(t.TempDir(), ".databrickscfg")
	_ = os.WriteFile(cfg, []byte("[P]\nhost = x\n"), 0o600)
	t.Setenv("DATABRICKS_CONFIG_FILE", cfg)

	if _, err := LoadDir(dir); err != nil {
		t.Fatalf("v1 LoadDir should not require step .tf files: %v", err)
	}
}

// ═══════════════════════════════════════════════════════════
// Plan-content matcher tests (Task #34
// ═══════════════════════════════════════════════════════════

// TestLoad_PlanMatchers_HappyPath confirms both fields decode and the
// regex compiles. Both values survive into the parsed Step.
func TestLoad_PlanMatchers_HappyPath(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    command: plan
    expect_non_empty_plan: true
    plan_match: 'forces replacement'
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	s := spec.Steps[0]
	if !s.ExpectNonEmptyPlan {
		t.Errorf("ExpectNonEmptyPlan: expected true")
	}
	if s.PlanMatch != "forces replacement" {
		t.Errorf("PlanMatch: got %q", s.PlanMatch)
	}
	if s.CompiledPlanMatch == nil {
		t.Errorf("CompiledPlanMatch should be set after Load")
	}
}

// TestLoad_PlanMatchers_DefaultsOff: a step with neither field set
// has both at zero values + nil regex.
func TestLoad_PlanMatchers_DefaultsOff(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: plan }
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	s := spec.Steps[0]
	if s.ExpectNonEmptyPlan || s.PlanMatch != "" || s.CompiledPlanMatch != nil {
		t.Errorf("expected all zero/nil, got ExpectNonEmptyPlan=%v PlanMatch=%q Compiled=%v",
			s.ExpectNonEmptyPlan, s.PlanMatch, s.CompiledPlanMatch)
	}
}

// TestLoad_PlanMatchers_RejectsApplyCommand pins the §17.10 rule:
// expect_non_empty_plan / plan_match require command: plan. An apply
// step with the field set is rejected at parse time.
func TestLoad_PlanMatchers_RejectsApplyCommand(t *testing.T) {
	for _, body := range []string{
		`
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: apply, expect_non_empty_plan: true }
`, `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: apply, plan_match: 'foo' }
`,
	} {
		_, err := loadString(t, body, "P")
		if err == nil {
			t.Errorf("expected error for plan matcher on apply step, got nil for body:\n%s", body)
		}
	}
}

// TestLoad_PlanMatchers_RejectsDestroyCommand: same gate against
// destroy.
func TestLoad_PlanMatchers_RejectsDestroyCommand(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: destroy, expect_non_empty_plan: true }
`
	if _, err := loadString(t, body, "P"); err == nil {
		t.Error("expected error for plan matcher on destroy step")
	}
}

// TestLoad_PlanMatchers_RejectsExpectFailure: per §17.10,
// plan-content matchers also require expect: success. A failed
// plan's stdout shape is undefined (terraform may abort early), so
// content assertions there are noise.
func TestLoad_PlanMatchers_RejectsExpectFailure(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - name: a
    version: "1.0.0"
    command: plan
    expect: failure
    error_substring: oops
    expect_non_empty_plan: true
`
	if _, err := loadString(t, body, "P"); err == nil {
		t.Error("expected error for plan matcher on expect=failure step")
	}
}

// TestLoad_PlanMatchers_RejectsBadRegex pins parse-time regex
// compilation. A `(` with no closing `)` is RE2-invalid.
func TestLoad_PlanMatchers_RejectsBadRegex(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: plan, plan_match: '(' }
`
	if _, err := loadString(t, body, "P"); err == nil {
		t.Error("expected error for invalid plan_match regex")
	}
}

// TestLoad_PlanMatchers_RegexMultiline confirms the implicit `(?s)`
// flag — `.` matches across newlines, so a plan_match like
// "Plan:.*destroy" works against multiline plan output.
func TestLoad_PlanMatchers_RegexMultiline(t *testing.T) {
	body := `
name: t
profile: P
steps:
  - { name: a, version: "1.0.0", command: plan, plan_match: 'Plan:.*destroy' }
`
	spec, err := loadString(t, body, "P")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	re := spec.Steps[0].CompiledPlanMatch
	multilineInput := "Plan: 1 to add, 0 to change, 1 to destroy.\nMore output here."
	if !re.MatchString(multilineInput) {
		t.Errorf("compiled regex should match multiline input via implicit (?s)")
	}
}
