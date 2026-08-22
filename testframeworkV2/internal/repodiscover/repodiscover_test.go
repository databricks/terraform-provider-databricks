package repodiscover

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// TestFind_FromRepoSubdir lays out a synthetic provider repo + a few
// nested subdirs and confirms Find walks upward correctly.
func TestFind_FromRepoSubdir(t *testing.T) {
	root := t.TempDir()
	writeGoMod(t, root, "module "+ProviderModulePath)

	// nested/deeply/buried — Find from here should return root.
	deep := filepath.Join(root, "nested", "deeply", "buried")
	if err := os.MkdirAll(deep, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	got, err := Find(deep)
	if err != nil {
		t.Fatalf("Find: %v", err)
	}
	wantRoot, _ := filepath.EvalSymlinks(root)
	gotRoot, _ := filepath.EvalSymlinks(got)
	if gotRoot != wantRoot {
		t.Fatalf("Find returned %q; want %q", gotRoot, wantRoot)
	}
}

// TestFind_SkipsSubmodule confirms a sub-module's go.mod (different
// module path) does NOT short-circuit the walk — Find must keep
// climbing until it finds the canonical provider module.
func TestFind_SkipsSubmodule(t *testing.T) {
	root := t.TempDir()
	writeGoMod(t, root, "module "+ProviderModulePath)

	// A submodule with its own go.mod (different module path), like the
	// framework's testframeworkV2/go.mod.
	sub := filepath.Join(root, "testframeworkV2")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeGoMod(t, sub, "module "+ProviderModulePath+"/testframeworkV2")

	// Find from inside the submodule should walk past the submodule's
	// go.mod (wrong module path) and return the parent root.
	got, err := Find(sub)
	if err != nil {
		t.Fatalf("Find: %v", err)
	}
	wantRoot, _ := filepath.EvalSymlinks(root)
	gotRoot, _ := filepath.EvalSymlinks(got)
	if gotRoot != wantRoot {
		t.Fatalf("Find returned %q; want %q (skipped past submodule)", gotRoot, wantRoot)
	}
}

// TestFind_NotFound exercises the failure path: a tree with no
// matching go.mod anywhere in the upward walk.
func TestFind_NotFound(t *testing.T) {
	tmp := t.TempDir()
	deep := filepath.Join(tmp, "a", "b", "c")
	if err := os.MkdirAll(deep, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	_, err := Find(deep)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Find: want ErrNotFound, got %v", err)
	}
}

// TestFind_DefaultsToCwd confirms Find("") uses the current working
// directory. We chdir into a fixture root and verify Find returns it.
func TestFind_DefaultsToCwd(t *testing.T) {
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(orig) })

	root := t.TempDir()
	writeGoMod(t, root, "module "+ProviderModulePath)
	if err := os.Chdir(root); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	got, err := Find("")
	if err != nil {
		t.Fatalf("Find(\"\"): %v", err)
	}
	wantRoot, _ := filepath.EvalSymlinks(root)
	gotRoot, _ := filepath.EvalSymlinks(got)
	if gotRoot != wantRoot {
		t.Fatalf("Find(\"\") returned %q; want %q", gotRoot, wantRoot)
	}
}

// TestMatchesProvider_IgnoresComments confirms leading // comments and
// blank lines don't trip the match — the first non-blank, non-comment
// line is what counts.
func TestMatchesProvider_IgnoresComments(t *testing.T) {
	dir := t.TempDir()
	writeGoMod(t, dir, "// SPDX-License-Identifier: Apache-2.0\n\n"+
		"module "+ProviderModulePath+"\n\ngo 1.25\n")
	if !matchesProvider(filepath.Join(dir, "go.mod")) {
		t.Fatalf("matchesProvider: expected true for go.mod with leading comments")
	}
}

// TestMatchesProvider_RejectsSubmodulePath confirms a go.mod whose
// module line is a sub-path of the provider module is NOT matched.
// (Otherwise Find would short-circuit on testframeworkV2/go.mod.)
func TestMatchesProvider_RejectsSubmodulePath(t *testing.T) {
	dir := t.TempDir()
	writeGoMod(t, dir, "module "+ProviderModulePath+"/testframeworkV2")
	if matchesProvider(filepath.Join(dir, "go.mod")) {
		t.Fatalf("matchesProvider: expected false for submodule path")
	}
}

func writeGoMod(t *testing.T, dir, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(body), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
}
