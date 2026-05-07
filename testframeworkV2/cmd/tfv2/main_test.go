package main

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

func TestRun_NoArgsPrintsUsage(t *testing.T) {
	got := captureStderr(t, func() int { return run(nil) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d want %d", got.code, exitCodeUsage)
	}
	if !strings.Contains(got.stderr, "tfv2 run") {
		t.Errorf("usage banner missing 'tfv2 run' text: %q", got.stderr)
	}
}

func TestRun_VersionSubcommand(t *testing.T) {
	for _, arg := range []string{"version", "--version"} {
		t.Run(arg, func(t *testing.T) {
			got := captureStdout(t, func() int { return run([]string{arg}) })
			if got.code != exitCodeOK {
				t.Errorf("exit: got %d want %d", got.code, exitCodeOK)
			}
			if !strings.HasPrefix(got.stdout, "tfv2 ") {
				t.Errorf("version output: got %q", got.stdout)
			}
		})
	}
}

func TestRun_HelpSubcommand(t *testing.T) {
	for _, arg := range []string{"help", "--help", "-h"} {
		t.Run(arg, func(t *testing.T) {
			got := captureStdout(t, func() int { return run([]string{arg}) })
			if got.code != exitCodeOK {
				t.Errorf("exit: got %d want %d", got.code, exitCodeOK)
			}
			if !strings.Contains(got.stdout, "tfv2 run") {
				t.Errorf("help output missing run subcommand: %q", got.stdout)
			}
		})
	}
}

func TestRun_UnknownSubcommand(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"frobnicate"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d want %d", got.code, exitCodeUsage)
	}
	if !strings.Contains(got.stderr, "frobnicate") {
		t.Errorf("error output should name unknown subcommand: %q", got.stderr)
	}
}

func TestParseRunFlags_Defaults(t *testing.T) {
	t.Setenv("TFV2_TERRAFORM_BIN", "")
	t.Setenv("TFV2_CACHE_DIR", "")
	f, err := parseRunFlags([]string{"some/test/dir"})
	if err != nil {
		t.Fatalf("parseRunFlags: %v", err)
	}
	if f.testDir != "some/test/dir" {
		t.Errorf("testDir: got %q", f.testDir)
	}
	if f.terraformBin != "" || f.cacheDir != "" || f.runDir != "" {
		t.Errorf("expected empty defaults, got %+v", f)
	}
	if f.noCleanup || f.verbose {
		t.Errorf("expected bool defaults false, got %+v", f)
	}
}

func TestParseRunFlags_AllFlags(t *testing.T) {
	args := []string{
		"--terraform-bin", "/tmp/tf",
		"--cache-dir", "/tmp/cache",
		"--run-dir", "/tmp/runs",
		"--repo", "/tmp/repo",
		"--no-cleanup",
		"--verbose",
		"some/test/dir",
	}
	f, err := parseRunFlags(args)
	if err != nil {
		t.Fatalf("parseRunFlags: %v", err)
	}
	for _, tc := range []struct{ got, want string }{
		{f.terraformBin, "/tmp/tf"},
		{f.cacheDir, "/tmp/cache"},
		{f.runDir, "/tmp/runs"},
		{f.repoRoot, "/tmp/repo"},
		{f.testDir, "some/test/dir"},
	} {
		if tc.got != tc.want {
			t.Errorf("flag value: got %q want %q", tc.got, tc.want)
		}
	}
	if !f.noCleanup || !f.verbose {
		t.Errorf("expected --no-cleanup and --verbose to be set: %+v", f)
	}
}

func TestParseRunFlags_EnvDefaults(t *testing.T) {
	t.Setenv("TFV2_TERRAFORM_BIN", "/env/tf")
	t.Setenv("TFV2_CACHE_DIR", "/env/cache")
	f, err := parseRunFlags([]string{"some/test/dir"})
	if err != nil {
		t.Fatalf("parseRunFlags: %v", err)
	}
	if f.terraformBin != "/env/tf" {
		t.Errorf("TFV2_TERRAFORM_BIN env not honored: got %q", f.terraformBin)
	}
	if f.cacheDir != "/env/cache" {
		t.Errorf("TFV2_CACHE_DIR env not honored: got %q", f.cacheDir)
	}
}

func TestParseRunFlags_FlagBeatsEnv(t *testing.T) {
	t.Setenv("TFV2_TERRAFORM_BIN", "/env/tf")
	f, err := parseRunFlags([]string{"--terraform-bin", "/flag/tf", "x"})
	if err != nil {
		t.Fatalf("parseRunFlags: %v", err)
	}
	if f.terraformBin != "/flag/tf" {
		t.Errorf("flag should win over env: got %q", f.terraformBin)
	}
}

func TestParseRunFlags_MissingTestDir(t *testing.T) {
	if _, err := parseRunFlags([]string{}); err == nil {
		t.Error("expected error for missing testDir, got nil")
	}
}

func TestParseRunFlags_TooManyArgs(t *testing.T) {
	if _, err := parseRunFlags([]string{"a", "b"}); err == nil {
		t.Error("expected error for too many args, got nil")
	}
}

func TestPrintRunResult_AllPassed(t *testing.T) {
	r := result.RunResult{
		Test:     "t1",
		Steps:    []result.StepResult{{Index: 0, Name: "a", Version: "1.0.0", Command: "plan", Status: result.StatusPass, Duration: 100 * time.Millisecond}},
		Duration: 100 * time.Millisecond,
		RunDir:   "/some/run/dir",
	}
	var buf bytes.Buffer
	printRunResult(&buf, r)
	out := buf.String()
	for _, want := range []string{"[PASS]", "step 1 (a)", "1.0.0 plan", "1/1 steps passed", "/some/run/dir"} {
		if !strings.Contains(out, want) {
			t.Errorf("output should contain %q, got:\n%s", want, out)
		}
	}
}

func TestPrintRunResult_WithFailures(t *testing.T) {
	r := result.RunResult{
		Test: "t",
		Steps: []result.StepResult{
			{Index: 0, Name: "a", Version: "1.113.0", Command: "plan", Status: result.StatusPass, Duration: 100 * time.Millisecond},
			{
				Index: 1, Name: "b", Version: "1.114.0", Command: "plan",
				Status: result.StatusFail, Reason: "expected success but got error: boom",
				Duration:  200 * time.Millisecond,
				StderrLog: "/run/dir/step_2_b.stderr.log",
			},
		},
		Duration: 300 * time.Millisecond,
	}
	var buf bytes.Buffer
	printRunResult(&buf, r)
	out := buf.String()
	if !strings.Contains(out, "[PASS]") || !strings.Contains(out, "[FAIL]") {
		t.Errorf("expected both PASS and FAIL tags, got:\n%s", out)
	}
	if !strings.Contains(out, "expected success but got error: boom") {
		t.Errorf("FAIL line should include reason: %s", out)
	}
	if !strings.Contains(out, "1/2 steps passed") {
		t.Errorf("summary should show 1/2 passed: %s", out)
	}
	// Run-dir hint pointer line on FAIL — operator-debugging entry
	// point per task #23 v2 polish.
	if !strings.Contains(out, "(full stderr at /run/dir/step_2_b.stderr.log)") {
		t.Errorf("FAIL line should be followed by stderr-path hint, got:\n%s", out)
	}
}

// TestPrintRunResult_FailHintIndent pins the 7-space indent on the
// stderr-path hint line — that's the visual width of "[FAIL] " (the
// tag plus its trailing space), which threads the hint under the
// tag column. Operators eyeballing the output expect this column
// alignment.
func TestPrintRunResult_FailHintIndent(t *testing.T) {
	r := result.RunResult{
		Test: "t",
		Steps: []result.StepResult{{
			Index: 0, Name: "a", Version: "1.0.0", Command: "plan",
			Status: result.StatusFail, Reason: "boom",
			Duration:  100 * time.Millisecond,
			StderrLog: "/r/step_1_a.stderr.log",
		}},
		Duration: 100 * time.Millisecond,
	}
	var buf bytes.Buffer
	printRunResult(&buf, r)
	for line := range strings.SplitSeq(buf.String(), "\n") {
		if strings.HasPrefix(line, "(full stderr") {
			t.Errorf("hint line must be indented 7 spaces, got %q", line)
		}
		if strings.HasPrefix(line, "       (full stderr") {
			return // found the well-indented line; pass
		}
	}
	t.Errorf("did not find indented hint line in output:\n%s", buf.String())
}

// TestPrintRunResult_NoHintOnPass: passing steps don't get a hint
// line. Stderr log is harmless (and present), but the hint is FAIL-
// only.
func TestPrintRunResult_NoHintOnPass(t *testing.T) {
	r := result.RunResult{
		Test: "t",
		Steps: []result.StepResult{{
			Index: 0, Name: "a", Version: "1.0.0", Command: "plan",
			Status: result.StatusPass, Duration: 100 * time.Millisecond,
			StderrLog: "/r/step_1_a.stderr.log",
		}},
		Duration: 100 * time.Millisecond,
	}
	var buf bytes.Buffer
	printRunResult(&buf, r)
	if strings.Contains(buf.String(), "full stderr at") {
		t.Errorf("PASS step should not emit stderr hint, got:\n%s", buf.String())
	}
}

// TestFormatFailHint covers the helper directly: FAIL + stderr →
// hint, every other shape → empty.
func TestFormatFailHint(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   result.StepResult
		want string
	}{
		{"fail with stderr", result.StepResult{Status: result.StatusFail, StderrLog: "/x.log"}, "       (full stderr at /x.log)"},
		{"fail without stderr", result.StepResult{Status: result.StatusFail}, ""},
		{"pass with stderr", result.StepResult{Status: result.StatusPass, StderrLog: "/x.log"}, ""},
		{"skipped", result.StepResult{Status: result.StatusSkipped, StderrLog: "/x.log"}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatFailHint(tc.in); got != tc.want {
				t.Errorf("got %q\nwant %q", got, tc.want)
			}
		})
	}
}

func TestPrintRunResult_Skipped(t *testing.T) {
	r := result.RunResult{Test: "t", Skipped: true, Reason: "no GCP profile"}
	var buf bytes.Buffer
	printRunResult(&buf, r)
	out := buf.String()
	if !strings.Contains(out, "SKIPPED") || !strings.Contains(out, "no GCP profile") {
		t.Errorf("skipped output: got %q", out)
	}
}

// TestRunOnce_LocatesFakeTerraformAndAssertsBinaryFails is a thin
// integration test that exercises terraform.LocateAndCheck via
// runOnce. We use a fake terraform script that returns a too-old
// version so we can assert the flow without spinning up real
// terraform — and runOnce surfaces a useful error.
func TestRunOnce_LocatesFakeTerraformAndAssertsBinaryFails(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake terraform stub uses POSIX shebang")
	}
	dir := t.TempDir()
	// Old version → SanityCheckVersion will reject.
	stub := filepath.Join(dir, "terraform")
	if err := os.WriteFile(stub, []byte("#!/bin/sh\necho 'Terraform v1.4.0'\n"), 0o755); err != nil {
		t.Fatalf("write stub: %v", err)
	}
	// We don't actually need a valid test dir because the flow errors
	// at SanityCheckVersion before reaching config.LoadDir.
	f := runFlags{terraformBin: stub, testDir: t.TempDir()}
	_, err := runOnce(t.Context(), f)
	if err == nil {
		t.Fatal("expected error for old terraform version, got nil")
	}
	if !strings.Contains(err.Error(), "below minimum") {
		t.Errorf("expected 'below minimum' error, got: %v", err)
	}
}

// TestOptionsFromFlags_DefaultsApplied verifies CLI flag → runner.Options translation.
func TestOptionsFromFlags_DefaultsApplied(t *testing.T) {
	t.Setenv("HOME", "/home/u")
	o := optionsFromFlags(runFlags{}, "/tf", "/src")
	if o.SourceDir != "/src" {
		t.Errorf("SourceDir: got %q", o.SourceDir)
	}
	if o.TerraformBin != "/tf" {
		t.Errorf("TerraformBin: got %q", o.TerraformBin)
	}
	if !strings.HasSuffix(o.CacheDir, ".testframeworkv2/providers") {
		t.Errorf("CacheDir default: got %q", o.CacheDir)
	}
	if !strings.HasSuffix(o.RunRoot, ".testframeworkv2/runs") {
		t.Errorf("RunRoot default: got %q", o.RunRoot)
	}
	if o.NoCleanup {
		t.Errorf("NoCleanup default: expected false")
	}
}

func TestOptionsFromFlags_FlagsOverrideDefaults(t *testing.T) {
	o := optionsFromFlags(runFlags{
		cacheDir: "/c", runDir: "/r", repoRoot: "/repo", noCleanup: true,
	}, "/tf", "/src")
	if o.CacheDir != "/c" || o.RunRoot != "/r" || o.RepoRoot != "/repo" || !o.NoCleanup {
		t.Errorf("flag overrides not applied: %+v", o)
	}
}

// captureResult bundles the captured CLI output for easy assertion.
type captureResult struct {
	code   int
	stdout string
	stderr string
}

// captureStdout invokes fn while capturing os.Stdout.
func captureStdout(t *testing.T, fn func() int) captureResult {
	t.Helper()
	return capture(t, fn, true, false)
}

// captureStderr invokes fn while capturing os.Stderr.
func captureStderr(t *testing.T, fn func() int) captureResult {
	t.Helper()
	return capture(t, fn, false, true)
}

// capture redirects os.Stdout / os.Stderr to pipes, calls fn, then
// restores. Returns the captured strings + the function's return code.
func capture(t *testing.T, fn func() int, captureOut, captureErr bool) captureResult {
	t.Helper()
	origOut, origErr := os.Stdout, os.Stderr
	defer func() { os.Stdout, os.Stderr = origOut, origErr }()

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	if captureOut {
		os.Stdout = wOut
	}
	if captureErr {
		os.Stderr = wErr
	}

	code := fn()
	wOut.Close()
	wErr.Close()

	var outBuf, errBuf bytes.Buffer
	if captureOut {
		_, _ = outBuf.ReadFrom(rOut)
	}
	if captureErr {
		_, _ = errBuf.ReadFrom(rErr)
	}
	return captureResult{code: code, stdout: outBuf.String(), stderr: errBuf.String()}
}
