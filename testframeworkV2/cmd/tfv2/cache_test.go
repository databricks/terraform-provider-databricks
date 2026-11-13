package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// makeCachedZip materializes a fake packed zip at the canonical
// layout path so scanCache picks it up. The bytes are arbitrary — we
// just need a regular file with non-zero size for the size column.
func makeCachedZip(t *testing.T, root, version, target string) string {
	t.Helper()
	dir := filepath.Join(root, "registry.terraform.io", "databricks", "databricks")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	path := filepath.Join(dir, fmt.Sprintf("terraform-provider-databricks_%s_%s.zip", version, target))
	if err := os.WriteFile(path, []byte("stub-zip-content"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return path
}

// makeCachedUnpacked materializes a fake unpacked layout entry.
func makeCachedUnpacked(t *testing.T, root, version, target string) string {
	t.Helper()
	dir := filepath.Join(root, "registry.terraform.io", "databricks", "databricks", version, target)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	path := filepath.Join(dir, "terraform-provider-databricks_v"+version)
	if err := os.WriteFile(path, []byte("stub-binary"), 0o755); err != nil {
		t.Fatalf("write: %v", err)
	}
	return path
}

func TestScanCache_PackedAndUnpacked(t *testing.T) {
	root := t.TempDir()
	makeCachedZip(t, root, "1.113.0", "darwin_arm64")
	makeCachedZip(t, root, "1.114.0", "darwin_arm64")
	makeCachedUnpacked(t, root, "99.0.0-local", "darwin_arm64")

	got, err := scanCache(root)
	if err != nil {
		t.Fatalf("scanCache: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 entries, got %d: %+v", len(got), got)
	}
	// Sorted by version then target. "1.113.0" < "1.114.0" < "99.0.0-local".
	for i, want := range []struct {
		version, target, layout string
	}{
		{"1.113.0", "darwin_arm64", "packed"},
		{"1.114.0", "darwin_arm64", "packed"},
		{"99.0.0-local", "darwin_arm64", "unpacked"},
	} {
		if got[i].Version != want.version || got[i].Target != want.target || got[i].Layout != want.layout {
			t.Errorf("entry %d: got %+v want version=%s target=%s layout=%s", i, got[i], want.version, want.target, want.layout)
		}
		if got[i].SizeBytes <= 0 {
			t.Errorf("entry %d size: got %d", i, got[i].SizeBytes)
		}
	}
}

func TestScanCache_EmptyCache(t *testing.T) {
	got, err := scanCache(t.TempDir())
	if err != nil {
		t.Fatalf("scanCache: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}

func TestScanCache_MissingProviderDir(t *testing.T) {
	// Cache root exists, but the registry.terraform.io subtree does not.
	got, err := scanCache(t.TempDir())
	if err != nil {
		t.Fatalf("scanCache: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestScanCache_IgnoresUnknownFiles(t *testing.T) {
	root := t.TempDir()
	makeCachedZip(t, root, "1.114.0", "darwin_arm64")
	// Drop a stray file that doesn't match either layout.
	stray := filepath.Join(root, "registry.terraform.io", "databricks", "databricks", "README.md")
	_ = os.WriteFile(stray, []byte("ignore me"), 0o644)
	// And an unknown-layout subtree.
	_ = os.MkdirAll(filepath.Join(root, "registry.terraform.io", "databricks", "databricks", "1.113.0"), 0o755)
	_ = os.WriteFile(filepath.Join(root, "registry.terraform.io", "databricks", "databricks", "1.113.0", "spurious"), []byte("x"), 0o644)

	got, err := scanCache(root)
	if err != nil {
		t.Fatalf("scanCache: %v", err)
	}
	if len(got) != 1 {
		t.Errorf("expected 1 entry, got %d: %+v", len(got), got)
	}
}

func TestParsePackedZipName(t *testing.T) {
	for _, tc := range []struct {
		name             string
		input            string
		wantVer, wantTgt string
		wantOK           bool
	}{
		{"darwin arm64", "terraform-provider-databricks_1.114.0_darwin_arm64.zip", "1.114.0", "darwin_arm64", true},
		{"linux amd64", "terraform-provider-databricks_1.0.0_linux_amd64.zip", "1.0.0", "linux_amd64", true},
		{"prerelease version", "terraform-provider-databricks_1.10.0-rc1_linux_amd64.zip", "1.10.0-rc1", "linux_amd64", true},
		{"missing prefix", "1.114.0_darwin_arm64.zip", "", "", false},
		{"missing suffix", "terraform-provider-databricks_1.114.0_darwin_arm64", "", "", false},
		{"missing target", "terraform-provider-databricks_1.114.0.zip", "", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			v, tgt, ok := parsePackedZipName(tc.input)
			if ok != tc.wantOK || v != tc.wantVer || tgt != tc.wantTgt {
				t.Errorf("got (%q, %q, %v), want (%q, %q, %v)", v, tgt, ok, tc.wantVer, tc.wantTgt, tc.wantOK)
			}
		})
	}
}

func TestFormatBytes(t *testing.T) {
	for _, tc := range []struct {
		n    int64
		want string
	}{
		{500, "500 B"},
		{2048, "2 KB"},
		{67_000_000, "63 MB"},
		{2_147_483_648, "2.0 GB"},
	} {
		if got := formatBytes(tc.n); got != tc.want {
			t.Errorf("formatBytes(%d) = %q want %q", tc.n, got, tc.want)
		}
	}
}

func TestRunCacheList_EmptyCache(t *testing.T) {
	dir := t.TempDir()
	got := captureStdout(t, func() int { return run([]string{"cache", "list", "--cache-dir", dir}) })
	if got.code != exitCodeOK {
		t.Errorf("exit: got %d", got.code)
	}
	if !strings.Contains(got.stdout, "(empty cache") {
		t.Errorf("expected empty-cache marker, got %q", got.stdout)
	}
}

func TestRunCacheList_WithEntries(t *testing.T) {
	dir := t.TempDir()
	makeCachedZip(t, dir, "1.114.0", "darwin_arm64")
	makeCachedUnpacked(t, dir, "99.0.0-local", "darwin_arm64")

	got := captureStdout(t, func() int { return run([]string{"cache", "list", "--cache-dir", dir}) })
	if got.code != exitCodeOK {
		t.Errorf("exit: got %d", got.code)
	}
	for _, want := range []string{"1.114.0", "darwin_arm64", "packed", "99.0.0-local", "unpacked"} {
		if !strings.Contains(got.stdout, want) {
			t.Errorf("output missing %q:\n%s", want, got.stdout)
		}
	}
}

func TestRunCachePrune_RemovesCache(t *testing.T) {
	dir := t.TempDir()
	makeCachedZip(t, dir, "1.114.0", "darwin_arm64")

	got := captureStdout(t, func() int { return run([]string{"cache", "prune", "--cache-dir", dir}) })
	if got.code != exitCodeOK {
		t.Errorf("exit: got %d", got.code)
	}
	if _, err := os.Stat(dir); err == nil {
		t.Errorf("expected cache dir to be removed, but it still exists")
	}
}

func TestRunCachePrune_AlreadyEmpty(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "no-such-cache")
	got := captureStdout(t, func() int { return run([]string{"cache", "prune", "--cache-dir", dir}) })
	if got.code != exitCodeOK {
		t.Errorf("exit: got %d", got.code)
	}
	if !strings.Contains(got.stdout, "already empty") {
		t.Errorf("expected 'already empty', got %q", got.stdout)
	}
}

func TestRunCache_MissingAction(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"cache"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d", got.code)
	}
}

func TestRunCache_UnknownAction(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"cache", "frobnicate"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d", got.code)
	}
	if !strings.Contains(got.stderr, "frobnicate") {
		t.Errorf("error should name unknown action: %q", got.stderr)
	}
}

func TestRunBuild_MissingMode(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"build"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d", got.code)
	}
}

func TestRunBuild_UnknownMode(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"build", "remote"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d", got.code)
	}
	if !strings.Contains(got.stderr, "remote") {
		t.Errorf("error should name unknown mode: %q", got.stderr)
	}
}

func TestRunBuildLocal_MissingRepoFlag(t *testing.T) {
	got := captureStderr(t, func() int { return run([]string{"build", "local"}) })
	if got.code != exitCodeUsage {
		t.Errorf("exit: got %d", got.code)
	}
	if !strings.Contains(got.stderr, "--repo") {
		t.Errorf("error should mention --repo: %q", got.stderr)
	}
}

func TestRunBuildLocal_HappyPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("uses POSIX shell tooling for git init in helper")
	}
	if !haveExecOnPath(t, "go", "git") {
		t.Skip("go and git required")
	}
	repo := makeTinyGoRepoCLI(t)
	cache := t.TempDir()

	got := captureStdout(t, func() int { return run([]string{"build", "local", "--repo", repo, "--cache-dir", cache}) })
	if got.code != exitCodeOK {
		t.Errorf("exit: got %d, stdout: %s", got.code, got.stdout)
	}
	for _, want := range []string{"built 99.0.0-local", "binary:", "git_sha:"} {
		if !strings.Contains(got.stdout, want) {
			t.Errorf("output missing %q:\n%s", want, got.stdout)
		}
	}
}
