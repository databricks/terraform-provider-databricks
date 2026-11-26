package terraform

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// fakeTerraform writes a stub `terraform` shell script in t.TempDir() that
// prints the given version line and exits zero. Returns the absolute path.
// Skips the test on Windows since shebangs aren't honored.
func fakeTerraform(t *testing.T, versionLine string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("fake-terraform stub uses POSIX shebang")
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "terraform")
	body := "#!/bin/sh\nprintf '%s\\n' " + shellQuote(versionLine) + "\n"
	if err := os.WriteFile(path, []byte(body), 0o755); err != nil {
		t.Fatalf("write fake terraform: %v", err)
	}
	return path
}

// shellQuote produces a single-quoted POSIX shell string. We use it for
// the fake terraform stub above, where the version line is the only
// dynamic value we interpolate.
func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func TestLocate_OverrideFlag(t *testing.T) {
	path := fakeTerraform(t, "Terraform v1.7.0")
	got, err := Locate(path)
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if got != path {
		t.Errorf("path: got %q want %q", got, path)
	}
}

func TestLocate_OverrideTrimsWhitespace(t *testing.T) {
	path := fakeTerraform(t, "Terraform v1.7.0")
	got, err := Locate("  " + path + "  ")
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if got != path {
		t.Errorf("expected whitespace trimmed: got %q want %q", got, path)
	}
}

func TestLocate_TFV2EnvVar(t *testing.T) {
	path := fakeTerraform(t, "Terraform v1.7.0")
	t.Setenv(EnvVarBin, path)
	// Ensure PATH lookup wouldn't satisfy us — clear PATH for this test.
	t.Setenv("PATH", filepath.Dir(path))

	got, err := Locate("")
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if got != path {
		t.Errorf("path: got %q want %q", got, path)
	}
}

func TestLocate_PATHFallback(t *testing.T) {
	path := fakeTerraform(t, "Terraform v1.7.0")
	t.Setenv("PATH", filepath.Dir(path))
	t.Setenv(EnvVarBin, "")

	got, err := Locate("")
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if filepath.Base(got) != "terraform" {
		t.Errorf("expected `terraform` from PATH, got %q", got)
	}
}

// TestLocate_OverridePrecedence ensures the override beats both
// TFV2_TERRAFORM_BIN and PATH when all three are set.
func TestLocate_OverridePrecedence(t *testing.T) {
	override := fakeTerraform(t, "Terraform v1.7.0")
	envOnly := fakeTerraform(t, "Terraform v1.6.0")
	pathOnly := fakeTerraform(t, "Terraform v1.5.0")
	t.Setenv(EnvVarBin, envOnly)
	t.Setenv("PATH", filepath.Dir(pathOnly))

	got, err := Locate(override)
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if got != override {
		t.Errorf("expected override to win, got %q", got)
	}
}

// TestLocate_EnvBeatsPATH confirms the order TFV2_TERRAFORM_BIN > PATH.
func TestLocate_EnvBeatsPATH(t *testing.T) {
	envBin := fakeTerraform(t, "Terraform v1.7.0")
	pathBin := fakeTerraform(t, "Terraform v1.5.0")
	t.Setenv(EnvVarBin, envBin)
	t.Setenv("PATH", filepath.Dir(pathBin))

	got, err := Locate("")
	if err != nil {
		t.Fatalf("Locate: %v", err)
	}
	if got != envBin {
		t.Errorf("expected env to win over PATH, got %q", got)
	}
}

func TestLocate_NotFound(t *testing.T) {
	t.Setenv(EnvVarBin, "")
	t.Setenv("PATH", t.TempDir()) // empty dir — no terraform

	_, err := Locate("")
	if err == nil {
		t.Fatal("expected ErrNotFound, got nil")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected error to wrap ErrNotFound, got: %v", err)
	}
}

func TestLocate_RejectsNonexistentOverride(t *testing.T) {
	_, err := Locate("/this/path/does/not/exist")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "does not exist") {
		t.Errorf("expected 'does not exist' in error, got: %v", err)
	}
}

func TestLocate_RejectsDirectoryOverride(t *testing.T) {
	dir := t.TempDir()
	_, err := Locate(dir)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "directory") {
		t.Errorf("expected 'directory' in error, got: %v", err)
	}
}

func TestParseVersion(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want string
	}{
		{"plain", "Terraform v1.7.0", "1.7.0"},
		{"trailing newline", "Terraform v1.7.0\n", "1.7.0"},
		{"with platform line", "Terraform v1.5.0\non darwin_arm64\n", "1.5.0"},
		{"trailing whitespace", "Terraform v1.7.0  \n", "1.7.0"},
		{"trailing build info", "Terraform v1.7.0 (foo)\n", "1.7.0"},
		{"prerelease", "Terraform v1.10.0-rc1", "1.10.0-rc1"},
		{"leading whitespace", "  Terraform v1.7.0", "1.7.0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseVersion(tc.in)
			if err != nil {
				t.Fatalf("ParseVersion: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestParseVersion_Errors(t *testing.T) {
	for _, in := range []string{
		"",
		"not terraform output",
		"terraform v1.7.0", // lowercase 't' — must be capitalized per spec
		"Terraform 1.7.0",  // missing 'v' prefix
	} {
		t.Run(in, func(t *testing.T) {
			_, err := ParseVersion(in)
			if err == nil {
				t.Errorf("ParseVersion(%q): expected error, got nil", in)
			}
		})
	}
}

func TestSanityCheckVersion_Pass(t *testing.T) {
	bin := fakeTerraform(t, "Terraform v1.7.0")
	v, err := SanityCheckVersion(context.Background(), bin)
	if err != nil {
		t.Fatalf("SanityCheckVersion: %v", err)
	}
	if v != "1.7.0" {
		t.Errorf("version: got %q want 1.7.0", v)
	}
}

func TestSanityCheckVersion_Minimum(t *testing.T) {
	// MinVersion is currently 1.5.0; pass a binary at exactly the floor.
	bin := fakeTerraform(t, "Terraform v"+MinVersion)
	v, err := SanityCheckVersion(context.Background(), bin)
	if err != nil {
		t.Fatalf("SanityCheckVersion at floor: %v", err)
	}
	if v != MinVersion {
		t.Errorf("version: got %q want %q", v, MinVersion)
	}
}

func TestSanityCheckVersion_BelowMinimum(t *testing.T) {
	bin := fakeTerraform(t, "Terraform v1.4.0")
	_, err := SanityCheckVersion(context.Background(), bin)
	if err == nil {
		t.Fatal("expected error for old version, got nil")
	}
	if !strings.Contains(err.Error(), "below minimum") {
		t.Errorf("expected 'below minimum' in error, got: %v", err)
	}
}

func TestSanityCheckVersion_Unparseable(t *testing.T) {
	bin := fakeTerraform(t, "garbage output")
	_, err := SanityCheckVersion(context.Background(), bin)
	if err == nil {
		t.Fatal("expected error for unparseable version, got nil")
	}
}

// TestLocateAndCheck_HappyPath exercises the combined flow used by the
// runner at startup.
func TestLocateAndCheck_HappyPath(t *testing.T) {
	bin := fakeTerraform(t, "Terraform v1.7.0")
	path, version, err := LocateAndCheck(context.Background(), bin)
	if err != nil {
		t.Fatalf("LocateAndCheck: %v", err)
	}
	if path != bin {
		t.Errorf("path: got %q want %q", path, bin)
	}
	if version != "1.7.0" {
		t.Errorf("version: got %q want 1.7.0", version)
	}
}

// TestLocateAndCheck_OldVersion documents that LocateAndCheck still
// returns the path even when the version check fails — useful for
// callers that want to print "you have terraform vX.Y.Z; please upgrade".
func TestLocateAndCheck_OldVersion(t *testing.T) {
	bin := fakeTerraform(t, "Terraform v1.4.0")
	path, version, err := LocateAndCheck(context.Background(), bin)
	if err == nil {
		t.Fatal("expected error for old version, got nil")
	}
	if path != bin {
		t.Errorf("path: got %q want %q", path, bin)
	}
	if version != "1.4.0" {
		t.Errorf("version: got %q want 1.4.0", version)
	}
}

// TestSplitSemver_PrereleaseStripped covers the explicit choice to
// compare release versions ignoring -rc / -beta suffixes.
func TestSplitSemver_PrereleaseStripped(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want [3]int
	}{
		{"1.5.0", [3]int{1, 5, 0}},
		{"1.5.0-rc1", [3]int{1, 5, 0}},
		{"1.5.0+build", [3]int{1, 5, 0}},
	} {
		got, err := splitSemver(tc.in)
		if err != nil {
			t.Errorf("splitSemver(%q): %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("splitSemver(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestCompareSemver(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want int
	}{
		{"1.5.0", "1.5.0", 0},
		{"1.4.9", "1.5.0", -1},
		{"1.5.1", "1.5.0", 1},
		{"2.0.0", "1.99.99", 1},
		{"1.10.0", "1.9.0", 1},
	} {
		got, err := compareSemver(tc.a, tc.b)
		if err != nil {
			t.Errorf("compareSemver(%s, %s): %v", tc.a, tc.b, err)
			continue
		}
		if got != tc.want {
			t.Errorf("compareSemver(%s, %s) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}
