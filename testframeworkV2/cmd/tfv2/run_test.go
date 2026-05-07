package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// haveExecOnPath returns true when every named binary is on PATH.
// Used to skip tests that need real external commands (go, git).
func haveExecOnPath(t *testing.T, names ...string) bool {
	t.Helper()
	for _, n := range names {
		if _, err := exec.LookPath(n); err != nil {
			return false
		}
	}
	return true
}

// makeTinyGoRepoCLI is a CLI-test-local helper that produces a
// buildable Go module + git repo so `tfv2 build local` has something
// to compile against. Mirrors the runner-test helper but is duplicated
// to keep test packages independent.
func makeTinyGoRepoCLI(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "main.go"),
		[]byte("package main\n\nfunc main() {}\n"), 0o644); err != nil {
		t.Fatalf("write main.go: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"),
		[]byte("module example.com/tinybuild\n\ngo 1.21\n"), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	for _, args := range [][]string{
		{"init", "--initial-branch=main"},
		{"-c", "user.name=tester", "-c", "user.email=t@example.com", "add", "."},
		{"-c", "user.name=tester", "-c", "user.email=t@example.com", "commit", "-m", "init", "--no-gpg-sign"},
	} {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
	return dir
}

func TestParseRunFlags_RecursiveFlag(t *testing.T) {
	for _, arg := range []string{"-r", "--recursive"} {
		t.Run(arg, func(t *testing.T) {
			f, err := parseRunFlags([]string{arg, "some/dir"})
			if err != nil {
				t.Fatalf("parseRunFlags: %v", err)
			}
			if !f.recursive {
				t.Errorf("expected recursive=true, got %+v", f)
			}
			if f.testDir != "some/dir" {
				t.Errorf("testDir: got %q", f.testDir)
			}
		})
	}
}

// TestDiscoverTests_FlatLayout: a single test.yaml at root.
func TestDiscoverTests_FlatLayout(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "test1")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "test.yaml"), []byte("name: t\nprofile: P\nsteps: [{name: a, version: '1.0.0'}]\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	got, err := discoverTests(root)
	if err != nil {
		t.Fatalf("discoverTests: %v", err)
	}
	if len(got) != 1 || got[0] != dir {
		t.Errorf("got %v, want [%s]", got, dir)
	}
}

// TestDiscoverTests_NestedLayout: tests grouped by profile-level dirs.
// Walks deeper than flat layout.
func TestDiscoverTests_NestedLayout(t *testing.T) {
	root := t.TempDir()
	dirs := []string{
		filepath.Join(root, "account", "test_a"),
		filepath.Join(root, "account", "test_b"),
		filepath.Join(root, "workspace", "test_c"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(d, "test.yaml"), []byte("x"), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	got, err := discoverTests(root)
	if err != nil {
		t.Fatalf("discoverTests: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 tests, got %d: %v", len(got), got)
	}
	// sorted alphabetically.
	for i, d := range dirs {
		if got[i] != d {
			t.Errorf("entry %d: got %q want %q", i, got[i], d)
		}
	}
}

// TestDiscoverTests_SkipsHiddenDirs: .git, .terraform, etc. must be
// skipped — otherwise we descend into per-run workdirs and pick up
// test.yaml copies the framework wrote earlier.
func TestDiscoverTests_SkipsHiddenDirs(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, ".terraform", "providers"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	// A test.yaml under .terraform/ must be skipped.
	if err := os.WriteFile(filepath.Join(root, ".terraform", "test.yaml"), []byte("x"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	// One real test directory at the root.
	dir := filepath.Join(root, "real")
	_ = os.MkdirAll(dir, 0o755)
	_ = os.WriteFile(filepath.Join(dir, "test.yaml"), []byte("x"), 0o644)

	got, err := discoverTests(root)
	if err != nil {
		t.Fatalf("discoverTests: %v", err)
	}
	if len(got) != 1 || got[0] != dir {
		t.Errorf("got %v, want [%s]", got, dir)
	}
}

// TestDiscoverTests_NoTests returns an empty slice rather than an
// error — the caller decides how to handle that.
func TestDiscoverTests_NoTests(t *testing.T) {
	got, err := discoverTests(t.TempDir())
	if err != nil {
		t.Fatalf("discoverTests: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}

// TestRunRecursive_NoTestsFound exercises the friendly "no test.yaml"
// branch end-to-end via the dispatcher.
func TestRunRecursive_NoTestsFound(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"run", "-r", t.TempDir()}) })
	if got.code != exitCodeFailed {
		t.Errorf("exit: got %d want %d", got.code, exitCodeFailed)
	}
	if !strings.Contains(got.stderr, "no test.yaml found") {
		t.Errorf("stderr: %q", got.stderr)
	}
}
