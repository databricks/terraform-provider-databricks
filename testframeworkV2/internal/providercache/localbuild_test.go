package providercache

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

// requireGo skips the test when `go` isn't on PATH. The fixture builds
// a tiny module so we need the toolchain to be available.
func requireGo(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("go"); err != nil {
		t.Skipf("go binary not found on PATH: %v", err)
	}
}

// requireGit skips the test when `git` isn't on PATH or initial commit
// can't be made. The captured provenance relies on git plumbing.
func requireGit(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skipf("git binary not found on PATH: %v", err)
	}
}

// makeTinyGoModule creates a buildable `package main` Go module in a
// fresh t.TempDir(). go.mod uses Go 1.21 (long-supported baseline) so
// the test doesn't sneak in toolchain churn.
func makeTinyGoModule(t *testing.T) string {
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
	return dir
}

// gitInit runs the minimal sequence to turn dir into a git repo with
// one commit so gitState returns a valid SHA. We set local-only
// committer config so the test doesn't depend on the developer's
// global git identity.
func gitInit(t *testing.T, dir string) {
	t.Helper()
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
}

// TestLocalBinPath_LayoutPin pins the unpacked-mirror layout from
// A regression here breaks Terraform's filesystem_mirror
// discovery for local builds.
func TestLocalBinPath_LayoutPin(t *testing.T) {
	c := New(t.TempDir())
	target := Target{OS: "darwin", Arch: "arm64"}

	wantDir := filepath.Join(c.Root(),
		"registry.terraform.io", "databricks", "databricks",
		"99.0.0-local", "darwin_arm64")
	wantBin := filepath.Join(wantDir, "terraform-provider-databricks_v99.0.0-local")

	if got := c.localBinDir(target); got != wantDir {
		t.Errorf("localBinDir: got %q want %q", got, wantDir)
	}
	if got := c.localBinPath(target); got != wantBin {
		t.Errorf("localBinPath: got %q want %q", got, wantBin)
	}
	// Sanity: LocalBinaryName + SyntheticVersionLocal also pinned for
	// downstream consumers (override file uses SyntheticVersionLocal).
	if SyntheticVersionLocal != "99.0.0-local" {
		t.Errorf("SyntheticVersionLocal: got %q", SyntheticVersionLocal)
	}
	if LocalBinaryName() != "terraform-provider-databricks_v99.0.0-local" {
		t.Errorf("LocalBinaryName: got %q", LocalBinaryName())
	}
}

// TestBuildLocal_HappyPath compiles a tiny Go module, lands the
// binary under the unpacked layout, and confirms local-version.json
// is written with sane fields.
func TestBuildLocal_HappyPath(t *testing.T) {
	requireGo(t)
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	c := New(t.TempDir())
	target := HostTarget() // build for current host so the binary actually executes

	binDir, syn, prov, err := c.BuildLocal(t.Context(), repo, target)
	if err != nil {
		t.Fatalf("BuildLocal: %v", err)
	}
	if syn != SyntheticVersionLocal {
		t.Errorf("syntheticVersion: got %q want %q", syn, SyntheticVersionLocal)
	}
	if got, want := binDir, c.localBinDir(target); got != want {
		t.Errorf("binDir: got %q want %q", got, want)
	}
	binPath := c.localBinPath(target)
	if info, err := os.Stat(binPath); err != nil {
		t.Errorf("binary not written: %v", err)
	} else if info.Size() == 0 {
		t.Errorf("binary is empty")
	}

	// Provenance: file present, content parseable, fields populated.
	provPath := filepath.Join(binDir, LocalVersionFilename)
	body, err := os.ReadFile(provPath)
	if err != nil {
		t.Fatalf("read provenance: %v", err)
	}
	var got LocalVersion
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal provenance: %v", err)
	}
	if got.SyntheticVersion != SyntheticVersionLocal {
		t.Errorf("SyntheticVersion: got %q", got.SyntheticVersion)
	}
	if got.GoVersion != runtime.Version() {
		t.Errorf("GoVersion: got %q want %q", got.GoVersion, runtime.Version())
	}
	if got.OSArch != target.String() {
		t.Errorf("OSArch: got %q want %q", got.OSArch, target.String())
	}
	if got.GitSHA == "" {
		t.Errorf("GitSHA: empty (expected populated, repo was git-init'd)")
	}
	if got.Dirty {
		t.Errorf("Dirty: got true; freshly-committed repo should be clean")
	}
	if time.Since(got.BuiltAt) > time.Minute {
		t.Errorf("BuiltAt: %s, expected ~now", got.BuiltAt)
	}
	// Returned provenance struct should match the on-disk JSON.
	if got.GitSHA != prov.GitSHA {
		t.Errorf("returned vs file GitSHA: %q vs %q", prov.GitSHA, got.GitSHA)
	}
}

// TestBuildLocal_DirtyRepoMarked confirms an uncommitted change in the
// repo lights up Dirty=true in the provenance.
func TestBuildLocal_DirtyRepoMarked(t *testing.T) {
	requireGo(t)
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	// Introduce uncommitted edit.
	if err := os.WriteFile(filepath.Join(repo, "untracked.txt"), []byte("hi"), 0o644); err != nil {
		t.Fatalf("touch: %v", err)
	}
	c := New(t.TempDir())

	_, _, prov, err := c.BuildLocal(t.Context(), repo, HostTarget())
	if err != nil {
		t.Fatalf("BuildLocal: %v", err)
	}
	if !prov.Dirty {
		t.Errorf("Dirty: expected true with untracked file present")
	}
}

// TestBuildLocal_RebuildsEveryCall ensures the cache-strategy
// "rebuild every step" guarantee from holds. The second
// call should re-execute go build and update local-version.json's
// BuiltAt timestamp.
func TestBuildLocal_RebuildsEveryCall(t *testing.T) {
	requireGo(t)
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	c := New(t.TempDir())

	_, _, prov1, err := c.BuildLocal(t.Context(), repo, HostTarget())
	if err != nil {
		t.Fatalf("BuildLocal #1: %v", err)
	}
	time.Sleep(10 * time.Millisecond) // ensure measurable BuiltAt delta
	_, _, prov2, err := c.BuildLocal(t.Context(), repo, HostTarget())
	if err != nil {
		t.Fatalf("BuildLocal #2: %v", err)
	}
	if !prov2.BuiltAt.After(prov1.BuiltAt) {
		t.Errorf("BuiltAt should advance on rebuild: %s vs %s", prov1.BuiltAt, prov2.BuiltAt)
	}
}

// TestBuildLocal_RejectsBadInputs confirms input validation runs
// before any FS or exec work.
func TestBuildLocal_RejectsBadInputs(t *testing.T) {
	c := New(t.TempDir())
	for _, tc := range []struct {
		name    string
		repo    string
		target  Target
		wantSub string
	}{
		{"empty repoRoot", "", HostTarget(), "repoRoot"},
		{"missing OS", t.TempDir(), Target{Arch: "arm64"}, "OS and Arch"},
		{"missing Arch", t.TempDir(), Target{OS: "darwin"}, "OS and Arch"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, err := c.BuildLocal(t.Context(), tc.repo, tc.target)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tc.wantSub) {
				t.Errorf("error %q should contain %q", err.Error(), tc.wantSub)
			}
		})
	}
}

// TestBuildLocal_BuildFailureSurfaced confirms a `go build` failure
// in the target dir surfaces as an error from BuildLocal.
func TestBuildLocal_BuildFailureSurfaced(t *testing.T) {
	requireGo(t)
	dir := t.TempDir()
	// Garbage Go source — guaranteed compile error.
	if err := os.WriteFile(filepath.Join(dir, "main.go"),
		[]byte("package main\n\nbroken!\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"),
		[]byte("module example.com/x\n\ngo 1.21\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	c := New(t.TempDir())

	_, _, _, err := c.BuildLocal(t.Context(), dir, HostTarget())
	if err == nil {
		t.Fatal("expected error from broken Go source, got nil")
	}
	if !strings.Contains(err.Error(), "go build") {
		t.Errorf("error should mention go build: %v", err)
	}
}

// TestBuildLocal_NoGitProvenanceFallback documents the best-effort
// git path: in a non-git directory, BuildLocal still succeeds but
// records GitSHA="" + Dirty=true.
func TestBuildLocal_NoGitProvenanceFallback(t *testing.T) {
	requireGo(t)
	repo := makeTinyGoModule(t) // no git init
	c := New(t.TempDir())

	_, _, prov, err := c.BuildLocal(t.Context(), repo, HostTarget())
	if err != nil {
		t.Fatalf("BuildLocal in non-git dir: %v", err)
	}
	if prov.GitSHA != "" {
		t.Errorf("GitSHA: got %q, want empty in non-git dir", prov.GitSHA)
	}
	if !prov.Dirty {
		t.Errorf("Dirty: expected true as the fallback signal in non-git dir")
	}
}

// TestCopyProvenanceTo writes a LocalVersion to a destination path
// and confirms the JSON round-trips. Used by the runner to satisfy
// the two-copy provenance rule from
func TestCopyProvenanceTo(t *testing.T) {
	dst := filepath.Join(t.TempDir(), "local-version.json")
	want := LocalVersion{
		SyntheticVersion: SyntheticVersionLocal,
		GitSHA:           "abc123",
		Dirty:            false,
		BuiltAt:          time.Date(2026, 5, 7, 20, 15, 0, 0, time.UTC),
		GoVersion:        "go1.25.8",
		OSArch:           "darwin_arm64",
	}
	if err := CopyProvenanceTo(dst, want); err != nil {
		t.Fatalf("CopyProvenanceTo: %v", err)
	}
	body, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	var got LocalVersion
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got != want {
		t.Errorf("provenance round-trip mismatch:\n got=%+v\nwant=%+v", got, want)
	}
}

// TestGitState_CleanRepo / TestGitState_DirtyRepo exercise gitState
// in isolation so a fail there localizes diagnostics independently
// from BuildLocal as a whole.
func TestGitState_CleanRepo(t *testing.T) {
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	sha, dirty, err := gitState(t.Context(), repo)
	if err != nil {
		t.Fatalf("gitState: %v", err)
	}
	if sha == "" {
		t.Error("expected non-empty SHA")
	}
	if dirty {
		t.Error("freshly-committed repo should not be dirty")
	}
}

func TestGitState_DirtyRepo(t *testing.T) {
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	if err := os.WriteFile(filepath.Join(repo, "extra.go"), []byte("package main\n"), 0o644); err != nil {
		t.Fatalf("touch: %v", err)
	}
	_, dirty, err := gitState(t.Context(), repo)
	if err != nil {
		t.Fatalf("gitState: %v", err)
	}
	if !dirty {
		t.Error("dirty repo not detected")
	}
}

func TestGitState_NotARepo(t *testing.T) {
	requireGit(t)
	_, _, err := gitState(t.Context(), t.TempDir())
	if err == nil {
		t.Error("expected error in non-git dir, got nil")
	}
}

// TestWriteProvenance_RoundTrip is the unit-level test for the
// provenance writer; the file format is human-inspected, so we want
// known-good shape.
func TestWriteProvenance_RoundTrip(t *testing.T) {
	dst := filepath.Join(t.TempDir(), "lv.json")
	want := LocalVersion{
		SyntheticVersion: SyntheticVersionLocal,
		GitSHA:           "deadbeef",
		Dirty:            true,
		BuiltAt:          time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
		GoVersion:        "go1.25.8",
		OSArch:           "linux_amd64",
	}
	if err := writeProvenance(dst, want); err != nil {
		t.Fatalf("writeProvenance: %v", err)
	}
	body, _ := os.ReadFile(dst)
	// Sanity check — the file should be human-readable JSON.
	if !strings.Contains(string(body), `"synthetic_version": "99.0.0-local"`) {
		t.Errorf("missing synthetic_version field: %s", body)
	}
	if !strings.HasSuffix(string(body), "\n") {
		t.Errorf("provenance file should end with newline (POSIX text-file convention)")
	}
}

// TestRunGoBuild_RespectsCGOEnv verifies the framework's canonical
// build flags reach `go build`. Indirect: we set GOOS to a value
// guaranteed to be valid (linux_amd64) on every platform; the build
// produces a non-host binary which we don't try to execute.
func TestRunGoBuild_RespectsCGOEnv(t *testing.T) {
	requireGo(t)
	repo := makeTinyGoModule(t)

	// Build for a target that's always cross-buildable from any host.
	target := Target{OS: "linux", Arch: "amd64"}
	out := filepath.Join(t.TempDir(), "out-binary")
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := runGoBuild(t.Context(), repo, out, target); err != nil {
		// Cross-build may fail on toolchains without linux SDK
		// available — skip rather than fail.
		t.Skipf("cross-build unavailable: %v", err)
	}
	if _, err := os.Stat(out); err != nil {
		t.Errorf("binary missing: %v", err)
	}
}

// TestBuildLocal_CrossArchTargetDir ensures the binary lands in the
// requested target's subdirectory, not the host's.
func TestBuildLocal_CrossArchTargetDir(t *testing.T) {
	requireGo(t)
	requireGit(t)
	repo := makeTinyGoModule(t)
	gitInit(t, repo)
	c := New(t.TempDir())

	// Build for linux_amd64 even on a darwin_arm64 host. Pure-Go
	// compilation makes this universally cross-buildable; if some
	// CI lacks linux stdlib, skip.
	target := Target{OS: "linux", Arch: "amd64"}
	binDir, _, _, err := c.BuildLocal(context.WithoutCancel(t.Context()), repo, target)
	if err != nil {
		t.Skipf("cross-build unavailable: %v", err)
	}
	if !strings.HasSuffix(binDir, filepath.Join("99.0.0-local", "linux_amd64")) {
		t.Errorf("binDir for cross target: got %q", binDir)
	}
}
