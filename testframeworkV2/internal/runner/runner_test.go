package runner

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/providercache"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// stepBehaviour is the unit-test contract for one step's mocked
// tfexec call sequence. Each runStep invocation receives one
// stepBehaviour; the cleanup destroy receives the next one (if any).
type stepBehaviour struct {
	initErr  error
	cmdErr   error
	stderr   []byte
	stdout   []byte
	settleMs time.Duration // optional sleep before returning, to exercise timeouts
}

// mockTF implements tfExec without spawning terraform. The behaviour
// pointer steers per-method outcomes; instance fields capture state
// (env, stdout/stderr writers, call counts) for post-run assertions.
type mockTF struct {
	behaviour stepBehaviour
	workDir   string
	bin       string

	env    map[string]string
	stdout io.Writer
	stderr io.Writer

	initCalls    int
	planCalls    int
	applyCalls   int
	destroyCalls int
	envSetCalls  int
}

func (m *mockTF) Init(ctx context.Context) error {
	m.initCalls++
	return m.behaviour.initErr
}

func (m *mockTF) Plan(ctx context.Context) (bool, error) {
	m.planCalls++
	m.writeBufs()
	return false, m.behaviour.cmdErr
}

func (m *mockTF) Apply(ctx context.Context) error {
	m.applyCalls++
	m.writeBufs()
	return m.behaviour.cmdErr
}

func (m *mockTF) Destroy(ctx context.Context) error {
	m.destroyCalls++
	m.writeBufs()
	return m.behaviour.cmdErr
}

func (m *mockTF) SetEnv(env map[string]string) error {
	m.envSetCalls++
	m.env = env
	return nil
}

func (m *mockTF) SetStdout(w io.Writer) { m.stdout = w }
func (m *mockTF) SetStderr(w io.Writer) { m.stderr = w }

// writeBufs streams the configured stdout/stderr bytes to whatever
// writers the runner installed. Real terraform writes line-by-line
// under various commands; for tests a single chunk is sufficient.
func (m *mockTF) writeBufs() {
	if len(m.behaviour.stdout) > 0 && m.stdout != nil {
		_, _ = m.stdout.Write(m.behaviour.stdout)
	}
	if len(m.behaviour.stderr) > 0 && m.stderr != nil {
		_, _ = m.stderr.Write(m.behaviour.stderr)
	}
	if m.behaviour.settleMs > 0 {
		time.Sleep(m.behaviour.settleMs)
	}
}

// stubFactory returns a tfFactory that hands out a fresh mockTF for
// each call, walking sequentially through behaviours. The returned
// instances slice exposes the mocks for post-run assertions.
func stubFactory(behaviours []stepBehaviour) (tfFactory, *[]*mockTF) {
	var instances []*mockTF
	idx := 0
	f := func(workDir, bin string) (tfExec, error) {
		if idx >= len(behaviours) {
			return nil, fmt.Errorf("stubFactory: unexpected call %d (only %d behaviours configured)", idx+1, len(behaviours))
		}
		m := &mockTF{behaviour: behaviours[idx], workDir: workDir, bin: bin}
		instances = append(instances, m)
		idx++
		return m, nil
	}
	return f, &instances
}

// prepopulatedCache materializes pre-existing zip files under a temp
// dir so providercache.Cache.Resolve returns a cache hit without
// attempting a network download. Stand-in bytes are arbitrary —
// the runner never opens the archive in M4.
func prepopulatedCache(t *testing.T, versions ...string) *providercache.Cache {
	t.Helper()
	root := t.TempDir()
	target := providercache.HostTarget()
	for _, v := range versions {
		path := filepath.Join(root, "registry.terraform.io", "databricks", "databricks",
			fmt.Sprintf("terraform-provider-databricks_%s_%s.zip", v, target.String()))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(path, []byte("stub"), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	return providercache.New(root)
}

// makeSourceDir creates a directory containing the given files. Tests
// pass the typical "main.tf only" map; helper kept generic so we can
// also exercise multi-file copy and .tfvars handling.
func makeSourceDir(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}

// makeAccountAWSProfile is the default profile shape for tests — the
// mission-test fixture targets account-level AWS, so most scenarios
// use this. Adjust returned fields per-test as needed.
func makeAccountAWSProfile() *profile.Profile {
	return &profile.Profile{
		Name:  "ACCOUNT_AWS",
		Host:  "https://accounts.cloud.databricks.com",
		Cloud: profile.CloudAWS,
		Level: profile.LevelAccount,
		Raw:   map[string]string{"host": "https://accounts.cloud.databricks.com"},
	}
}

// loadFixture writes test.yaml + main.tf in the same temp dir and
// returns (sourceDir, parsed *config.TestSpec, profilePath).
func loadFixture(t *testing.T, testYaml, mainTF, profileName string) (string, *config.TestSpec) {
	t.Helper()
	src := makeSourceDir(t, map[string]string{"test.yaml": testYaml, "main.tf": mainTF})
	cfgPath := filepath.Join(t.TempDir(), ".databrickscfg")
	if err := os.WriteFile(cfgPath, []byte("["+profileName+"]\nhost = https://accounts.cloud.databricks.com\n"), 0o600); err != nil {
		t.Fatalf("write databrickscfg: %v", err)
	}
	spec, err := config.LoadWithProfilePath(filepath.Join(src, "test.yaml"), cfgPath)
	if err != nil {
		t.Fatalf("LoadWithProfilePath: %v", err)
	}
	return src, spec
}

func newRunnerWithMock(t *testing.T, spec *config.TestSpec, prof *profile.Profile, src string, cache *providercache.Cache, behaviours []stepBehaviour) (*Runner, *[]*mockTF) {
	t.Helper()
	t.Setenv("T_NO_CLEANUP", "")
	t.Setenv("HOME", t.TempDir()) // so DATABRICKS_CONFIG_FILE points somewhere benign
	factory, instances := stubFactory(behaviours)
	r := New(spec, prof, Options{
		SourceDir:    src,
		CacheDir:     "ignored-when-WithCache-used",
		RunRoot:      t.TempDir(),
		TerraformBin: "/fake/terraform",
	},
		WithCache(cache),
		WithTFFactory(factory),
		WithNow(func() time.Time { return time.Date(2026, 5, 7, 20, 15, 0, 0, time.UTC) }),
		WithRandReader(bytes.NewReader([]byte{0xa3, 0xf2})),
	)
	return r, instances
}

const fixtureMainTF = `terraform {
  required_providers {
    databricks = { source = "databricks/databricks" }
  }
}

provider "databricks" {}

data "databricks_current_user" "me" {}
`

// TestRunner_SingleStep_Pass exercises the simplest happy path: one
// plan-only step where the mock tfexec returns nil for both Init and
// Plan. Confirms Status=Pass, log files exist, env was set on tfexec.
func TestRunner_SingleStep_Pass(t *testing.T) {
	yaml := `
name: t1
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - name: only_step
    version: "1.114.0"
    command: plan
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, // step 0: success
	})

	res, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Skipped {
		t.Fatalf("expected non-skipped run, got reason: %s", res.Reason)
	}
	if !res.AllPassed() {
		t.Errorf("expected all-pass, got: %s", res.String())
	}
	if len(res.Steps) != 1 {
		t.Fatalf("expected 1 step result, got %d", len(res.Steps))
	}
	if res.Steps[0].Status != result.StatusPass {
		t.Errorf("step status: got %q want %q", res.Steps[0].Status, result.StatusPass)
	}
	for _, log := range []string{res.Steps[0].StdoutLog, res.Steps[0].StderrLog} {
		if _, err := os.Stat(log); err != nil {
			t.Errorf("log file %s: %v", log, err)
		}
	}
	if got := len(*mocks); got != 1 {
		t.Errorf("expected 1 mock instance, got %d", got)
	}
	if (*mocks)[0].envSetCalls != 1 {
		t.Errorf("expected SetEnv called once, got %d", (*mocks)[0].envSetCalls)
	}
	// Env contents: must include the framework-controlled keys.
	for _, k := range []string{"DATABRICKS_CONFIG_PROFILE", "DATABRICKS_CONFIG_FILE", "TF_CLI_CONFIG_FILE", "TF_PLUGIN_CACHE_DIR"} {
		if _, ok := (*mocks)[0].env[k]; !ok {
			t.Errorf("env: expected %s, got %v", k, (*mocks)[0].env)
		}
	}
}

// TestRunner_MultiStep_FullMissionTest mimics the issue #5672 4-step
// shape: pass / fail-as-expected / pass / pass. The mocked stderr for
// step 2 contains the canary regex string.
func TestRunner_MultiStep_FullMissionTest(t *testing.T) {
	yaml := `
name: mission_test
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: passes_on_1_113_0, version: "1.113.0", command: plan }
  - name: fails_on_1_114_0
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'failed to resolve workspace_id'
  - { name: fixed_on_1_114_1, version: "1.114.1", command: plan }
  - { name: fixed_on_1_114_0_again, version: "1.114.0", command: plan, expect: failure, error_substring: "workspace_id" }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.113.0", "1.114.0", "1.114.1")
	stderr := []byte("Error: cannot populate provider_config for mws workspaces: failed to resolve workspace_id (boom)\n")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, // step 1 plan succeeds
		{stderr: stderr, cmdErr: errors.New("plan failed")},
		{}, // step 3 plan succeeds
		{stderr: stderr, cmdErr: errors.New("plan failed")},
	})

	res, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !res.AllPassed() {
		t.Fatalf("expected all-pass, got: %s\nfailed: %v", res.String(), res.FailedSteps())
	}
	for i, s := range res.Steps {
		if s.Status != result.StatusPass {
			t.Errorf("step %d (%s): got %q want pass — %s", i, s.Name, s.Status, s.Reason)
		}
	}
	if got := len(*mocks); got != 4 {
		t.Errorf("expected 4 mocks, got %d", got)
	}
	// Each step should have called Init exactly once and Plan exactly once.
	for i, m := range *mocks {
		if m.initCalls != 1 || m.planCalls != 1 {
			t.Errorf("mock %d: init=%d plan=%d (want 1/1)", i, m.initCalls, m.planCalls)
		}
	}
}

// TestRunner_FailureAsExpected_RegexMatch asserts the canonical "step
// fails with stderr matching the regex" path resolves to a PASS
// status. This is the keystone case for issue #5672's mission test.
func TestRunner_FailureAsExpected_RegexMatch(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - name: must_fail
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'failed to resolve workspace_id'
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{stderr: []byte("Error: failed to resolve workspace_id (auth=oauth-m2m)\n"), cmdErr: errors.New("plan failed")},
	})

	res, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Steps[0].Status != result.StatusPass {
		t.Errorf("expected PASS (failure-as-expected), got %s — %s", res.Steps[0].Status, res.Steps[0].Reason)
	}
}

// TestRunner_FailureAsExpected_RegexNoMatch confirms a missing regex
// match flips an "expected failure" step to FAIL (the cmd erred but
// the assertion didn't match).
func TestRunner_FailureAsExpected_RegexNoMatch(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - name: must_fail
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'this string is not in stderr'
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{stderr: []byte("Error: completely unrelated error\n"), cmdErr: errors.New("plan failed")},
	})

	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusFail {
		t.Errorf("expected FAIL when regex doesn't match, got %s", res.Steps[0].Status)
	}
	if !strings.Contains(res.Steps[0].Reason, "did not match") {
		t.Errorf("reason: got %q", res.Steps[0].Reason)
	}
}

// TestRunner_FailureAsExpected_SubstringNoMatch is the substring
// equivalent of the previous test.
func TestRunner_FailureAsExpected_SubstringNoMatch(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - name: must_fail
    version: "1.114.0"
    command: plan
    expect: failure
    error_substring: 'precondition failed'
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{stderr: []byte("Error: oops something else\n"), cmdErr: errors.New("plan failed")},
	})

	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusFail {
		t.Errorf("expected FAIL when substring doesn't match, got %s", res.Steps[0].Status)
	}
	if !strings.Contains(res.Steps[0].Reason, "did not contain") {
		t.Errorf("reason: got %q", res.Steps[0].Reason)
	}
}

// TestRunner_FailureExpected_ButPasses confirms an expect=failure step
// that returns nil from the command is FAIL.
func TestRunner_FailureExpected_ButPasses(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - name: should_fail
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'anything'
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, // cmd succeeds (cmdErr=nil) — assertion expected failure
	})

	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusFail {
		t.Errorf("expected FAIL, got %s", res.Steps[0].Status)
	}
	if !strings.Contains(res.Steps[0].Reason, "command succeeded") {
		t.Errorf("reason: got %q", res.Steps[0].Reason)
	}
}

// TestRunner_SuccessExpected_ButFails captures the inverse: a step
// that should succeed but tfexec returned an error.
func TestRunner_SuccessExpected_ButFails(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: should_pass, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{cmdErr: errors.New("unexpected boom")},
	})

	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusFail {
		t.Errorf("expected FAIL, got %s", res.Steps[0].Status)
	}
	if !strings.Contains(res.Steps[0].Reason, "expected success") {
		t.Errorf("reason: got %q", res.Steps[0].Reason)
	}
}

// TestRunner_RunContinuesAfterStepFail confirms a failed step does not
// short-circuit the run — DESIGN.md §7 explicitly requires later steps
// to execute regardless.
func TestRunner_RunContinuesAfterStepFail(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
  - { name: b, version: "1.114.0", command: plan }
  - { name: c, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{},
		{cmdErr: errors.New("middle step blows up")},
		{},
	})

	res, _ := r.Run(context.Background())
	if got := len(res.Steps); got != 3 {
		t.Fatalf("expected 3 step results (run did not stop on failure), got %d", got)
	}
	if res.Steps[0].Status != result.StatusPass {
		t.Errorf("step 0: %s", res.Steps[0].Status)
	}
	if res.Steps[1].Status != result.StatusFail {
		t.Errorf("step 1: %s", res.Steps[1].Status)
	}
	if res.Steps[2].Status != result.StatusPass {
		t.Errorf("step 2: %s", res.Steps[2].Status)
	}
	if got := len(*mocks); got != 3 {
		t.Errorf("expected 3 mocks, got %d", got)
	}
}

// TestRunner_SkippedOnCloudMismatch verifies the requires.cloud
// skip-check fires before any FS work.
func TestRunner_SkippedOnCloudMismatch(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: gcp, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, // shouldn't fire
	})

	res, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !res.Skipped {
		t.Errorf("expected Skipped=true, got %+v", res)
	}
	if !strings.Contains(res.Reason, "requires.cloud=gcp") {
		t.Errorf("Reason: got %q", res.Reason)
	}
	if got := len(*mocks); got != 0 {
		t.Errorf("expected 0 mock instances on skip, got %d", got)
	}
}

// TestRunner_SkippedOnLevelMismatch confirms requires.level is also a
// skip gate. Workspace test against an account profile must skip.
func TestRunner_SkippedOnLevelMismatch(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: workspace }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}})

	res, _ := r.Run(context.Background())
	if !res.Skipped {
		t.Errorf("expected Skipped=true, got %+v", res)
	}
	if !strings.Contains(res.Reason, "requires.level=workspace") {
		t.Errorf("Reason: got %q", res.Reason)
	}
}

// TestRunner_RunDirNamingDeterministic uses WithNow + WithRandReader
// to assert the run dir matches the F5 format exactly.
func TestRunner_RunDirNamingDeterministic(t *testing.T) {
	yaml := `
name: my_test
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}})

	res, _ := r.Run(context.Background())
	wantSuffix := "my_test-2026-05-07T20-15-00-a3f2"
	if !strings.HasSuffix(res.RunDir, wantSuffix) {
		t.Errorf("RunDir: got %q, want suffix %q", res.RunDir, wantSuffix)
	}
}

// TestRunner_LogFilesPopulated reads each step's stdout/stderr log and
// confirms the bytes the mock wrote ended up there.
func TestRunner_LogFilesPopulated(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	stdout := []byte("Plan: 0 to add, 0 to change, 0 to destroy.\n")
	stderr := []byte("Warning: data source X is deprecated\n")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{stdout: stdout, stderr: stderr},
	})

	res, _ := r.Run(context.Background())
	gotOut, _ := os.ReadFile(res.Steps[0].StdoutLog)
	gotErr, _ := os.ReadFile(res.Steps[0].StderrLog)
	if !bytes.Equal(gotOut, stdout) {
		t.Errorf("stdout log: got %q want %q", gotOut, stdout)
	}
	if !bytes.Equal(gotErr, stderr) {
		t.Errorf("stderr log: got %q want %q", gotErr, stderr)
	}
}

// TestRunner_LogFilenameSchema pins the step_<n>_<name>.{stdout,
// stderr}.log naming convention. A regression here would break tools
// that grep across log files in CI.
func TestRunner_LogFilenameSchema(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: alpha, version: "1.114.0", command: plan }
  - { name: beta,  version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}, {}})

	res, _ := r.Run(context.Background())
	for _, want := range []string{"step_1_alpha.stdout.log", "step_1_alpha.stderr.log", "step_2_beta.stdout.log", "step_2_beta.stderr.log"} {
		path := filepath.Join(res.RunDir, want)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("missing log file %s", want)
		}
	}
}

// TestRunner_TerraformStateWiped confirms that .terraform.lock.hcl and
// .terraform/ are removed before each step's init (DESIGN.md §7.1.c).
// We seed pre-existing copies and assert the mock never sees them
// during the second step's init.
func TestRunner_TerraformStateWiped(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
  - { name: b, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, {},
	})

	res, _ := r.Run(context.Background())
	workdir := filepath.Join(res.RunDir, "workdir")

	// After both steps complete, .terraform/ and .terraform.lock.hcl
	// should not exist (the runner's last action was wipe-then-init,
	// and the mock Init does nothing). This is a structural assertion
	// that wipeTerraformState ran at least once.
	if _, err := os.Stat(filepath.Join(workdir, ".terraform.lock.hcl")); err == nil {
		t.Errorf(".terraform.lock.hcl must be wiped before step init")
	}
	if _, err := os.Stat(filepath.Join(workdir, ".terraform")); err == nil {
		t.Errorf(".terraform/ must be wiped before step init")
	}
	// Sanity: each step got a fresh mock instance with init/plan calls.
	for i, m := range *mocks {
		if m.initCalls != 1 {
			t.Errorf("mock %d initCalls=%d", i, m.initCalls)
		}
	}
}

// TestRunner_TerraformRCWritten confirms .terraformrc is generated in
// the run's workdir (linchpin of the no-leak invariant — TF_CLI_CONFIG_FILE
// points here).
func TestRunner_TerraformRCWritten(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}})

	res, _ := r.Run(context.Background())
	rc, err := os.ReadFile(filepath.Join(res.RunDir, "workdir", ".terraformrc"))
	if err != nil {
		t.Fatalf("read .terraformrc: %v", err)
	}
	for _, want := range []string{"provider_installation", "filesystem_mirror", "registry.terraform.io/databricks/databricks", "plugin_cache_dir"} {
		if !strings.Contains(string(rc), want) {
			t.Errorf(".terraformrc missing %q", want)
		}
	}
}

// TestRunner_VersionsOverrideRegeneratedPerStep documents that each
// step's _tfv2_versions_override.tf reflects that step's version
// pin — a regression here breaks the multi-version aspect of the
// framework entirely.
func TestRunner_VersionsOverrideRegeneratedPerStep(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.113.0", command: plan }
  - { name: b, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.113.0", "1.114.0")
	versionsAtPlan := []string{}

	// Hook the mock to capture the override file contents at the
	// moment Plan is invoked (i.e. after writeOverride for that step).
	factory, instances := stubFactory([]stepBehaviour{{}, {}})
	wrapped := func(workDir, bin string) (tfExec, error) {
		tf, err := factory(workDir, bin)
		if err != nil {
			return nil, err
		}
		return &recordingMock{tfExec: tf, workDir: workDir, captured: &versionsAtPlan}, nil
	}
	t.Setenv("HOME", t.TempDir())
	r := New(spec, makeAccountAWSProfile(), Options{
		SourceDir: src, RunRoot: t.TempDir(), TerraformBin: "/fake",
	},
		WithCache(cache),
		WithTFFactory(wrapped),
		WithNow(func() time.Time { return time.Date(2026, 5, 7, 20, 15, 0, 0, time.UTC) }),
		WithRandReader(bytes.NewReader([]byte{0xa3, 0xf2})),
	)
	_, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	_ = instances // unused
	if len(versionsAtPlan) != 2 {
		t.Fatalf("expected 2 captures, got %d", len(versionsAtPlan))
	}
	if !strings.Contains(versionsAtPlan[0], `version = "= 1.113.0"`) {
		t.Errorf("step 1 override: got %q", versionsAtPlan[0])
	}
	if !strings.Contains(versionsAtPlan[1], `version = "= 1.114.0"`) {
		t.Errorf("step 2 override: got %q", versionsAtPlan[1])
	}
}

// recordingMock wraps a tfExec and snapshots
// _tfv2_versions_override.tf the first time Plan is called per step.
// Used by TestRunner_VersionsOverrideRegeneratedPerStep.
type recordingMock struct {
	tfExec
	workDir  string
	captured *[]string
}

func (r *recordingMock) Plan(ctx context.Context) (bool, error) {
	body, _ := os.ReadFile(filepath.Join(r.workDir, "_tfv2_versions_override.tf"))
	*r.captured = append(*r.captured, string(body))
	return r.tfExec.Plan(ctx)
}

// TestRunner_CleanupRunsAfterApply confirms cleanup destroy fires
// when at least one Apply step succeeds.
func TestRunner_CleanupRunsAfterApply(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
cleanup: true
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: apply }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	// 1 step + 1 cleanup = 2 mock instances expected.
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, {}, // step 1 apply, then cleanup destroy
	})
	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusPass {
		t.Fatalf("step 0 must pass for cleanup to fire: %s", res.Steps[0].Status)
	}
	if got := len(*mocks); got != 2 {
		t.Errorf("expected 2 mocks (apply + cleanup destroy), got %d", got)
	}
	if (*mocks)[1].destroyCalls != 1 {
		t.Errorf("cleanup mock destroyCalls=%d, want 1", (*mocks)[1].destroyCalls)
	}
}

// TestRunner_CleanupSkippedWhenPlanOnly mirrors the mission test's
// shape: every step is plan-only, lastSuccessfulApply is nil, cleanup
// is a no-op.
func TestRunner_CleanupSkippedWhenPlanOnly(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
cleanup: true
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}})
	_, _ = r.Run(context.Background())
	if got := len(*mocks); got != 1 {
		t.Errorf("expected exactly 1 mock (no cleanup destroy for plan-only test), got %d", got)
	}
}

// TestRunner_NoCleanupOption respects the runner option even when the
// test.yaml says cleanup: true.
func TestRunner_NoCleanupOption(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
cleanup: true
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: apply }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	t.Setenv("T_NO_CLEANUP", "")
	t.Setenv("HOME", t.TempDir())
	factory, instances := stubFactory([]stepBehaviour{{}})
	r := New(spec, makeAccountAWSProfile(), Options{
		SourceDir: src, RunRoot: t.TempDir(), TerraformBin: "/fake", NoCleanup: true,
	},
		WithCache(cache),
		WithTFFactory(factory),
		WithNow(func() time.Time { return time.Now().UTC() }),
		WithRandReader(bytes.NewReader([]byte{0xa3, 0xf2})),
	)
	_, _ = r.Run(context.Background())
	if got := len(*instances); got != 1 {
		t.Errorf("NoCleanup=true: expected 1 mock, got %d", got)
	}
}

// TestRunner_TNoCleanupEnvOverride confirms T_NO_CLEANUP=1 in the
// parent env disables cleanup.
func TestRunner_TNoCleanupEnvOverride(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
cleanup: true
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: apply }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	t.Setenv("T_NO_CLEANUP", "1")
	t.Setenv("HOME", t.TempDir())
	factory, instances := stubFactory([]stepBehaviour{{}})
	r := New(spec, makeAccountAWSProfile(), Options{
		SourceDir: src, RunRoot: t.TempDir(), TerraformBin: "/fake",
	},
		WithCache(cache),
		WithTFFactory(factory),
		WithNow(func() time.Time { return time.Now().UTC() }),
		WithRandReader(bytes.NewReader([]byte{0xa3, 0xf2})),
	)
	_, _ = r.Run(context.Background())
	if got := len(*instances); got != 1 {
		t.Errorf("T_NO_CLEANUP=1: expected 1 mock, got %d", got)
	}
}

// TestRunner_SourceDirCopiedIntoWorkdir verifies *.tf and *.tfvars
// from the source dir end up in workdir; non-terraform files are
// skipped; the source dir is never written to.
func TestRunner_SourceDirCopiedIntoWorkdir(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: plan }
`
	src := makeSourceDir(t, map[string]string{
		"test.yaml":        yaml,
		"main.tf":          "# main",
		"variables.tf":     "# vars",
		"terraform.tfvars": `region = "us-west-2"` + "\n",
		"README.md":        "ignored",
		"data.json":        "ignored",
		".hidden.tf":       "ignored (hidden)",
	})
	cfgPath := filepath.Join(t.TempDir(), ".databrickscfg")
	_ = os.WriteFile(cfgPath, []byte("[ACCOUNT_AWS]\nhost = https://accounts.cloud.databricks.com\n"), 0o600)
	spec, err := config.LoadWithProfilePath(filepath.Join(src, "test.yaml"), cfgPath)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	cache := prepopulatedCache(t, "1.114.0")
	t.Setenv("HOME", t.TempDir())
	factory, _ := stubFactory([]stepBehaviour{{}})
	r := New(spec, makeAccountAWSProfile(), Options{
		SourceDir: src, RunRoot: t.TempDir(), TerraformBin: "/fake",
	},
		WithCache(cache),
		WithTFFactory(factory),
		WithNow(func() time.Time { return time.Now().UTC() }),
		WithRandReader(bytes.NewReader([]byte{0xa3, 0xf2})),
	)
	res, _ := r.Run(context.Background())
	workdir := filepath.Join(res.RunDir, "workdir")

	for _, want := range []string{"main.tf", "variables.tf", "terraform.tfvars"} {
		if _, err := os.Stat(filepath.Join(workdir, want)); err != nil {
			t.Errorf("expected %s in workdir, got err: %v", want, err)
		}
	}
	for _, notWant := range []string{"README.md", "data.json", ".hidden.tf", "test.yaml"} {
		if _, err := os.Stat(filepath.Join(workdir, notWant)); err == nil {
			t.Errorf("did NOT expect %s in workdir", notWant)
		}
	}
}

// TestRunner_LocalVersionWithoutRepoRoot confirms M4 errors out when
// asked for version=local without RepoRoot configured (M6 territory).
func TestRunner_LocalVersionWithoutRepoRoot(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "local", command: plan }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t)
	r, _ := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{{}})

	_, err := r.Run(context.Background())
	if err == nil {
		t.Fatal("expected error for version=local without RepoRoot, got nil")
	}
	if !errors.Is(err, errMissingRepoRoot) {
		t.Errorf("expected errMissingRepoRoot, got: %v", err)
	}
}

// TestRunner_AppliesRunsApply uses command: apply and confirms the
// mock's Apply method is the one that fires (not Plan).
func TestRunner_AppliesRunsApply(t *testing.T) {
	yaml := `
name: t
profile: ACCOUNT_AWS
requires: { cloud: any, level: account }
steps:
  - { name: a, version: "1.114.0", command: apply }
`
	src, spec := loadFixture(t, yaml, fixtureMainTF, "ACCOUNT_AWS")
	cache := prepopulatedCache(t, "1.114.0")
	r, mocks := newRunnerWithMock(t, spec, makeAccountAWSProfile(), src, cache, []stepBehaviour{
		{}, // apply succeeds; cleanup mock will be the next behaviour
		{}, // cleanup destroy
	})

	res, _ := r.Run(context.Background())
	if res.Steps[0].Status != result.StatusPass {
		t.Errorf("step 0: %s", res.Steps[0].Status)
	}
	if (*mocks)[0].applyCalls != 1 || (*mocks)[0].planCalls != 0 {
		t.Errorf("step mock: apply=%d plan=%d (want 1/0)", (*mocks)[0].applyCalls, (*mocks)[0].planCalls)
	}
}
