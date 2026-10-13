// Package testframeworkv2_test — go-test integration for the framework's
// own fixtures.
//
// TestFixtures discovers every test.yaml under issues-repro/ and tests/,
// then runs each one programmatically through internal/runner under a
// dedicated t.Run subtest. This gives developers a single `go test
// ./...` invocation to validate every fixture without going through the
// CLI — useful for IDE green/red dots and CI integration.
//
// Gating: TFV2_RUN=1 is required to run. Otherwise TestFixtures skips.
// The fixtures hit real cloud APIs (terraform apply against ~/.databrickscfg
// profiles); we don't want them firing on every `go test./...` developer
// run.
package testframeworkv2_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/repodiscover"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/runner"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/terraform"
)

// fixtureRoots are the directories TestFixtures walks looking for
// `test.yaml`. Order is stable for deterministic test ordering.
var fixtureRoots = []string{"issues-repro", "tests"}

// TestFixtures walks every test.yaml under fixtureRoots and runs each
// one in a t.Run subtest. Each fixture is independent: a failure in one
// doesn't short-circuit the others (Go's testing.T.Run already isolates
// subtests).
//
// Subtest names preserve the directory tree (e.g.
// "issues-repro/issue_5672") so `go test -run 'TestFixtures/issues-repro/'`
// filters to a whole subtree, and tree-position information shows up in
// IDE / CI test output.
//
// Skip conditions:
// - TFV2_RUN != "1" — the test fires real cloud-auth flows; keep
// `go test./...` cheap by default.
// - The fixture's `requires:` block doesn't match the active profile —
// the runner returns Skipped=true; we surface that via t.Skip.
func TestFixtures(t *testing.T) {
	if os.Getenv("TFV2_RUN") != "1" {
		t.Skip("set TFV2_RUN=1 to run the live fixtures")
	}
	specs, err := discoverFixtures(fixtureRoots)
	if err != nil {
		t.Fatalf("discover fixtures: %v", err)
	}
	if len(specs) == 0 {
		t.Fatalf("no test.yaml files found under %v", fixtureRoots)
	}

	// Resolve the provider repo root once for all fixtures — every test
	// shares the same checkout. Failure is non-fatal here; per-fixture
	// `version: local` steps will surface a clear error if RepoRoot is
	// empty.
	repoRoot, _ := repodiscover.Find("")
	terraformBin, _, err := terraform.LocateAndCheck(context.Background(), os.Getenv("TFV2_TERRAFORM_BIN"))
	if err != nil {
		t.Fatalf("terraform locate: %v", err)
	}

	for _, dir := range specs {
		// Subtest name preserves the dir tree (e.g.
		// "issues-repro/issue_5672") and uses forward slashes regardless
		// of host OS so `-run TestFixtures/issues-repro/issue_5672`
		// matches on every platform. Trim a leading "./" if filepath
		// added one.
		name := filepath.ToSlash(strings.TrimPrefix(dir, "./"))
		t.Run(name, func(t *testing.T) {
			runFixture(t, dir, repoRoot, terraformBin)
		})
	}
}

// discoverFixtures returns directory paths (relative to the package
// directory, which is the cwd of `go test`) for every directory
// containing a test.yaml under any of `roots`. Roots that don't exist
// are silently skipped — not an error.
func discoverFixtures(roots []string) ([]string, error) {
	var out []string
	for _, root := range roots {
		matches, err := filepath.Glob(filepath.Join(root, "*", "test.yaml"))
		if err != nil {
			return nil, err
		}
		for _, m := range matches {
			out = append(out, filepath.Dir(m))
		}
	}
	sort.Strings(out)
	return out, nil
}

// runFixture executes a single fixture under a dedicated t.Run subtest.
// Skipped runs (per `requires:` mismatch) propagate as t.Skip; failed
// steps propagate as t.Errorf so go-test's report shows which step
// inside which fixture failed.
func runFixture(t *testing.T, dir, repoRoot, terraformBin string) {
	t.Helper()
	absDir, err := filepath.Abs(dir)
	if err != nil {
		t.Fatalf("abs(%q): %v", dir, err)
	}
	spec, err := config.LoadDir(absDir)
	if err != nil {
		// `config.LoadDir` validates profile-section existence in
		// `~/.databrickscfg` at parse time. When the
		// fixture's named profile isn't on this machine, the wrapped
		// error carries `profile.ErrSectionNotFound` — surface as Skip
		// with the profile name (intent: developer adds the profile and
		// reruns), not as Fail (the test isn't broken; the environment
		// just doesn't have what it needs).
		if errors.Is(err, profile.ErrSectionNotFound) {
			t.Skipf("load %s: %v", dir, err)
		}
		t.Fatalf("load %s: %v", dir, err)
	}
	prof, err := profile.Load(spec.Profile)
	if err != nil {
		// Belt-and-suspenders: if profile.Load disagrees with config
		// (~/.databrickscfg edited between parse and Load), still skip
		// rather than fail.
		if errors.Is(err, profile.ErrSectionNotFound) {
			t.Skipf("profile %q unavailable: %v", spec.Profile, err)
		}
		t.Fatalf("profile %q load: %v", spec.Profile, err)
	}
	opts := runner.Options{
		SourceDir:    absDir,
		CacheDir:     defaultCacheDirForTest(),
		RunRoot:      defaultRunDirForTest(),
		TerraformBin: terraformBin,
		RepoRoot:     repoRoot,
	}
	r := runner.New(spec, prof, opts)
	res, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("runner: %v", err)
	}
	if res.Skipped {
		t.Skipf("requires-skip: %s", res.Reason)
	}
	for _, step := range res.Steps {
		if step.Status != "pass" {
			t.Errorf("step %d (%s): %s — %s", step.Index+1, step.Name, step.Status, step.Reason)
		}
	}
}

// defaultCacheDirForTest mirrors cmd/tfv2/main.go's defaultCacheDir
// (no exported helper exists; the duplication is small enough to live
// alongside this single test for now).
func defaultCacheDirForTest() string {
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		return filepath.Join(home, ".testframeworkv2", "providers")
	}
	return filepath.Join(".testframeworkv2", "providers")
}

// defaultRunDirForTest is the analogous default for run roots.
func defaultRunDirForTest() string {
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		return filepath.Join(home, ".testframeworkv2", "runs")
	}
	return filepath.Join(".testframeworkv2", "runs")
}
