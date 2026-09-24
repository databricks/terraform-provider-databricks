package providercache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// LocalVersion captures provenance for a local provider build
// . Two copies are persisted per step:
//
// 1. <cache-root>/registry.terraform.io/databricks/databricks/99.0.0-local/<target>/local-version.json
// — overwritten each rebuild; reflects the most recent build.
// 2. <run-dir>/local-version.json — preserved with the run's logs so a
// test result is reproducible later.
//
// JSON field names match example.
type LocalVersion struct {
	SyntheticVersion string    `json:"synthetic_version"`
	GitSHA           string    `json:"git_sha"`
	Dirty            bool      `json:"dirty"`
	BuiltAt          time.Time `json:"built_at"`
	GoVersion        string    `json:"go_version"`
	OSArch           string    `json:"os_arch"`
}

// LocalVersionFilename is the basename used for both the cache-side
// and run-side provenance files. Constant so the runner can name
// the run-side file deterministically.
const LocalVersionFilename = "local-version.json"

// LocalBinaryName is the filename of the compiled provider binary
// inside the unpacked layout. The "v" prefix matches Terraform's
// expectation for provider binaries (it has historical baggage from
// pre-1.0 release naming).
func LocalBinaryName() string {
	return "terraform-provider-databricks_v" + SyntheticVersionLocal
}

// localBinDir returns the unpacked-layout target directory for the
// local synthetic version under the cache root. Exposed for tests.
func (c *Cache) localBinDir(target Target) string {
	return filepath.Join(c.providerDir(), SyntheticVersionLocal, target.String())
}

// localBinPath returns the absolute path the local-build binary will
// be written to.
func (c *Cache) localBinPath(target Target) string {
	return filepath.Join(c.localBinDir(target), LocalBinaryName())
}

// BuildLocal compiles the provider from repoRoot and installs it into
// the cache's unpacked layout. Per design's
// "rebuild every step" policy, BuildLocal does not consult any cache —
// it always invokes `go build`.
//
// Returns:
//
// - binDir unpacked layout directory containing the
// binary (the path Cache.Resolve also returns
// for local versions).
// - syntheticVersion always SyntheticVersionLocal ("99.0.0-local").
// - provenance the captured LocalVersion record. The runner
// uses this to also write a copy into the run dir
// .
//
// Provenance is captured BEFORE the build (so timestamps reflect when
// the inputs were sampled, not when the build finished). The git
// SHA / dirty flag are best-effort: a failed `git rev-parse` results in
// an empty SHA + Dirty=true so operators know the provenance is
// unreliable, rather than failing the build outright.
func (c *Cache) BuildLocal(ctx context.Context, repoRoot string, target Target) (binDir, syntheticVersion string, provenance LocalVersion, err error) {
	if repoRoot == "" {
		return "", "", LocalVersion{}, errors.New("providercache: BuildLocal requires repoRoot")
	}
	if target.OS == "" || target.Arch == "" {
		return "", "", LocalVersion{}, fmt.Errorf("providercache: BuildLocal target requires both OS and Arch (got %q)", target.String())
	}

	syntheticVersion = SyntheticVersionLocal
	binDir = c.localBinDir(target)
	binPath := c.localBinPath(target)
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return "", "", LocalVersion{}, fmt.Errorf("providercache: mkdir %s: %w", binDir, err)
	}

	provenance = captureProvenance(ctx, repoRoot, target)

	if err := runGoBuild(ctx, repoRoot, binPath, target); err != nil {
		return "", "", provenance, err
	}
	if err := writeProvenance(filepath.Join(binDir, LocalVersionFilename), provenance); err != nil {
		return "", "", provenance, err
	}
	return binDir, syntheticVersion, provenance, nil
}

// captureProvenance assembles the LocalVersion record. Best-effort on
// git fields — see BuildLocal's doc for the dirty=true fallback.
func captureProvenance(ctx context.Context, repoRoot string, target Target) LocalVersion {
	prov := LocalVersion{
		SyntheticVersion: SyntheticVersionLocal,
		GoVersion:        runtime.Version(),
		OSArch:           target.String(),
		BuiltAt:          time.Now().UTC(),
	}
	sha, dirty, err := gitState(ctx, repoRoot)
	if err != nil {
		// Non-fatal: record empty SHA + dirty=true so the operator
		// knows provenance is unreliable.
		prov.Dirty = true
		return prov
	}
	prov.GitSHA = sha
	prov.Dirty = dirty
	return prov
}

// runGoBuild invokes `go build -o <binPath>./` with the framework's
// canonical env (CGO_ENABLED=0 + GOOS/GOARCH for the target). Stdout
// and stderr go directly to the calling process's streams so users
// see compile errors without an extra log-file hop.
func runGoBuild(ctx context.Context, repoRoot, binPath string, target Target) error {
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binPath, "./")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(),
		"GOOS="+target.OS,
		"GOARCH="+target.Arch,
		"CGO_ENABLED=0", // matches goreleaser config
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("providercache: go build in %s: %w", repoRoot, err)
	}
	return nil
}

// gitState returns the (HEAD SHA, dirty flag) for repoRoot. dirty is
// true when `git status --porcelain` produces any output. Either git
// command failing is reported as an error; BuildLocal treats it as a
// non-fatal best-effort signal.
func gitState(ctx context.Context, repoRoot string) (string, bool, error) {
	sha, err := runGit(ctx, repoRoot, "rev-parse", "HEAD")
	if err != nil {
		return "", false, fmt.Errorf("git rev-parse: %w", err)
	}
	out, err := runGit(ctx, repoRoot, "status", "--porcelain")
	if err != nil {
		return "", false, fmt.Errorf("git status: %w", err)
	}
	return strings.TrimSpace(sha), strings.TrimSpace(out) != "", nil
}

// runGit is a thin exec wrapper used by gitState. Splitting it out
// keeps each command at one line in the caller.
func runGit(ctx context.Context, dir string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return string(out), err
}

// writeProvenance JSON-serializes v and writes it to path with 0644
// permissions. We marshal with indentation because the file is human-
// inspected during forensic debugging — disk usage is negligible.
func writeProvenance(path string, v LocalVersion) error {
	body, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("providercache: marshal provenance: %w", err)
	}
	body = append(body, '\n')
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return fmt.Errorf("providercache: write %s: %w", path, err)
	}
	return nil
}

// CopyProvenanceTo writes the provenance v to dst (typically the run
// dir's local-version.json). Used by the runner to satisfy the
// two-copy provenance rule from — the cache-side copy is
// written by BuildLocal directly; the run-dir copy is the runner's
// responsibility because providercache doesn't know about run dirs.
func CopyProvenanceTo(dst string, v LocalVersion) error {
	return writeProvenance(dst, v)
}
