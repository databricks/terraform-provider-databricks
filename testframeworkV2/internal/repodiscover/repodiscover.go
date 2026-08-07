// Package repodiscover finds the terraform-provider-databricks repo root
// by walking upward from a starting directory, looking for a go.mod whose
// `module` line matches the canonical provider module path.
//
// Used by cmd/tfv2 to fill in the --repo flag automatically when the user
// runs the framework from anywhere inside a checkout.
package repodiscover

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProviderModulePath is the canonical go.mod `module` line we match
// against. Lifted as a const so future renames touch one spot.
const ProviderModulePath = "github.com/databricks/terraform-provider-databricks"

// ErrNotFound is returned when no parent directory of `start` contains a
// go.mod with the expected module line.
var ErrNotFound = errors.New("repodiscover: no go.mod with module " +
	ProviderModulePath + " found walking up from start dir")

// Find walks upward from `start` (or os.Getwd() if empty) and returns
// the first directory containing a go.mod whose first non-comment
// `module` declaration is exactly ProviderModulePath. The framework's
// own go.mod (testframeworkV2/go.mod) is intentionally NOT a match —
// it declares a sub-module path, so the walk continues up to the parent
// provider repo.
func Find(start string) (string, error) {
	if start == "" {
		var err error
		start, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("repodiscover: getwd: %w", err)
		}
	}
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("repodiscover: abs(%q): %w", start, err)
	}
	for {
		if matchesProvider(filepath.Join(dir, "go.mod")) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", ErrNotFound
		}
		dir = parent
	}
}

// matchesProvider reports true iff path is a regular file containing a
// `module github.com/databricks/terraform-provider-databricks` line as
// its first non-blank, non-comment statement. We read just the first
// few hundred bytes — go.mod's module line is always near the top.
func matchesProvider(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		// First non-blank, non-comment line: must be the module declaration.
		return line == "module "+ProviderModulePath
	}
	return false
}
