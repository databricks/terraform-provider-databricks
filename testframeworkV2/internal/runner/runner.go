// Package runner orchestrates a single test.yaml run. It composes
// providercache, tfrcwriter, profile, subprocenv, and tfexec into the
// step execution flow described in DESIGN.md §7.
//
// The runner is deliberately mockable: pass WithTFFactory to substitute
// a fake tfExec for unit tests, WithNow / WithRandReader for
// deterministic run-dir naming, and WithCache to inject a pre-populated
// providercache (or one pointed at a fake GitHub releases server).
package runner

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/providercache"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/subprocenv"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/tfrcwriter"
)

// Options is the user-visible knob set the CLI / programmatic caller
// passes to New.
type Options struct {
	// SourceDir is the user's test directory containing test.yaml +
	// *.tf files. The runner copies *.tf into the run's workdir; it
	// never writes back to SourceDir (DESIGN.md §10 G11).
	SourceDir string

	// CacheDir is the providercache root (typically
	// ~/.testframeworkv2/providers). Created lazily.
	CacheDir string

	// RunRoot is the directory that holds per-run subdirectories
	// (typically ~/.testframeworkv2/runs). Created lazily.
	RunRoot string

	// TerraformBin is the absolute path to the terraform binary.
	// Resolved by internal/terraform.Locate before calling New.
	TerraformBin string

	// RepoRoot is the provider repo root used when a step's version
	// is "local" (DESIGN.md §8 / M6). Empty in M4.
	RepoRoot string

	// NoCleanup overrides cleanup: true per-test. Set true to keep the
	// final state — useful for debugging.
	NoCleanup bool
}

// Runner is a single test.yaml's executor. Construct with New, then call
// Run.
type Runner struct {
	spec *config.TestSpec
	prof *profile.Profile
	opts Options

	cache     *providercache.Cache
	tfFactory tfFactory

	// nowFn and randReader are seams for deterministic testing of
	// run-dir naming. Default to time.Now and crypto/rand.
	nowFn      func() time.Time
	randReader io.Reader
}

// Option configures a Runner.
type Option func(*Runner)

// WithTFFactory injects an alternative tfFactory — the linchpin of the
// unit-test mockability story.
func WithTFFactory(f tfFactory) Option {
	return func(r *Runner) {
		if f != nil {
			r.tfFactory = f
		}
	}
}

// WithCache injects a pre-built *providercache.Cache. Tests use this to
// point the runner at an httptest-backed cache without touching the
// developer's real ~/.testframeworkv2/providers.
func WithCache(c *providercache.Cache) Option {
	return func(r *Runner) {
		if c != nil {
			r.cache = c
		}
	}
}

// WithNow overrides time.Now for run-dir-naming determinism in tests.
func WithNow(now func() time.Time) Option {
	return func(r *Runner) {
		if now != nil {
			r.nowFn = now
		}
	}
}

// WithRandReader overrides the random source used to build the 4-char
// hex run-dir suffix (DESIGN.md §16/F5). Tests pass a deterministic
// reader for stable run-dir names.
func WithRandReader(rr io.Reader) Option {
	return func(r *Runner) {
		if rr != nil {
			r.randReader = rr
		}
	}
}

// New constructs a Runner. Spec and prof must already be validated by
// internal/config.Load + internal/profile.Load (or LoadFromPath).
func New(spec *config.TestSpec, prof *profile.Profile, opts Options, options ...Option) *Runner {
	r := &Runner{
		spec:       spec,
		prof:       prof,
		opts:       opts,
		tfFactory:  newRealTF,
		nowFn:      func() time.Time { return time.Now().UTC() },
		randReader: rand.Reader,
	}
	if opts.CacheDir != "" {
		r.cache = providercache.New(opts.CacheDir)
	}
	for _, o := range options {
		o(r)
	}
	return r
}

// Run executes every step in the spec and returns the aggregate result.
// Per DESIGN.md §7, the run continues even after a step fails — the
// whole point of the framework is multi-step regression tests where a
// middle step fails-as-expected.
//
// The returned RunResult.AllPassed() reports overall success. The
// returned error is non-nil only on infrastructure failures
// (cache-resolve, run-dir creation, etc.) — assertion failures are
// surfaced via StepResult.Status, not as an error.
func (r *Runner) Run(ctx context.Context) (result.RunResult, error) {
	started := r.nowFn()
	if reason, skip := r.skipReason(); skip {
		return result.RunResult{
			Test: r.spec.Name, Profile: r.spec.Profile, Skipped: true,
			Reason: reason, Started: started, Duration: r.nowFn().Sub(started),
		}, nil
	}
	prep, err := r.prepareRun()
	if err != nil {
		return result.RunResult{}, err
	}
	steps, err := r.runAllSteps(ctx, prep.workDir, prep.runDir, prep.envMap)
	if err != nil {
		return result.RunResult{}, err
	}
	res := result.RunResult{
		Test: r.spec.Name, Profile: r.spec.Profile, RunDir: prep.runDir,
		Steps: steps, Started: started, Duration: r.nowFn().Sub(started),
	}
	if r.cleanupEligible() {
		r.runCleanup(ctx, prep.workDir, prep.runDir, prep.envMap, steps)
	}
	return res, nil
}

// runPrep bundles the directories and env map produced during the
// pre-step setup phase. Keeps Run small while letting prepareRun
// return everything in a single shot.
type runPrep struct {
	runDir  string
	workDir string
	envMap  map[string]string
}

// prepareRun creates the run directory, copies user HCL into workdir,
// writes the per-run .terraformrc, and computes the curated subprocess
// env map. Errors here are infrastructure failures and abort the run.
func (r *Runner) prepareRun() (runPrep, error) {
	runDir, err := r.makeRunDir()
	if err != nil {
		return runPrep{}, err
	}
	workDir := filepath.Join(runDir, "workdir")
	pluginDir := filepath.Join(runDir, "plugins")
	if err := os.MkdirAll(pluginDir, 0o755); err != nil {
		return runPrep{}, fmt.Errorf("runner: mkdir %s: %w", pluginDir, err)
	}
	// In v1 mode the runner copies every *.tf / *.tfvars from SourceDir
	// into workDir once at run start. In v2 mode each step brings its
	// own per-step config file (Step.Config), so the bulk-copy is
	// skipped here and runStep does a wipe-and-copy before each init.
	// The workdir is still created so .terraformrc + override file
	// writes have somewhere to land.
	if !r.spec.IsV2() {
		if err := copyTerraformFiles(r.opts.SourceDir, workDir); err != nil {
			return runPrep{}, err
		}
	} else {
		if err := os.MkdirAll(workDir, 0o755); err != nil {
			return runPrep{}, fmt.Errorf("runner: mkdir %s: %w", workDir, err)
		}
	}
	tfrcPath, err := tfrcwriter.WriteTerraformRC(workDir, r.cache.Root(), pluginDir)
	if err != nil {
		return runPrep{}, fmt.Errorf("runner: write .terraformrc: %w", err)
	}
	env := subprocenv.Build(r.spec.Profile, tfrcPath, runDir, r.spec.PassthroughEnv)
	return runPrep{runDir: runDir, workDir: workDir, envMap: envSliceToMap(env)}, nil
}

// skipReason returns the (reason, true) pair when requires.cloud /
// requires.level don't match the loaded profile. requires.cloud == any
// always matches, regardless of profile.Cloud — including
// profile.CloudUnknown.
func (r *Runner) skipReason() (string, bool) {
	if r.spec.Requires.Cloud != config.CloudAny && string(r.spec.Requires.Cloud) != string(r.prof.Cloud) {
		return fmt.Sprintf("requires.cloud=%s but profile is %s", r.spec.Requires.Cloud, r.prof.Cloud), true
	}
	// requires.level always matches verbatim (no `any` value for level
	// — the schema doesn't allow it).
	if string(r.spec.Requires.Level) != string(r.prof.Level) {
		return fmt.Sprintf("requires.level=%s but profile is %s", r.spec.Requires.Level, r.prof.Level), true
	}
	return "", false
}

// makeRunDir creates ~/.testframeworkv2/runs/<test>-<ts>-<rand>/ and
// returns its absolute path. Naming follows DESIGN.md §16/F5:
// `<testName>-<RFC3339-with-colons-as-dashes>-<4-char-hex>`.
func (r *Runner) makeRunDir() (string, error) {
	ts := r.nowFn().UTC().Format("2006-01-02T15-04-05")
	suffix, err := r.randHex()
	if err != nil {
		return "", fmt.Errorf("runner: random suffix: %w", err)
	}
	dirName := fmt.Sprintf("%s-%s-%s", r.spec.Name, ts, suffix)
	full := filepath.Join(r.opts.RunRoot, dirName)
	if err := os.MkdirAll(filepath.Join(full, "workdir"), 0o755); err != nil {
		return "", fmt.Errorf("runner: mkdir %s: %w", full, err)
	}
	return full, nil
}

// randHex returns a 4-character lowercase hex string from the
// configured random reader (DESIGN.md §16/F5). Two bytes encode to
// exactly four hex characters.
func (r *Runner) randHex() (string, error) {
	b := make([]byte, 2)
	if _, err := io.ReadFull(r.randReader, b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// cleanupEligible reports whether we should run the post-run destroy
// pass. The conditions are:
//
//   - test.yaml's cleanup field is true (or unset, defaulting to true),
//   - --no-cleanup wasn't passed (Options.NoCleanup),
//   - T_NO_CLEANUP=1 wasn't set in the parent env (DESIGN.md §12.4),
//   - at least one Apply step succeeded (otherwise destroy is a no-op
//     — DESIGN.md §7).
//
// The lastSuccessfulApply check happens inside runCleanup itself; here
// we only check the flags.
func (r *Runner) cleanupEligible() bool {
	if !r.spec.CleanupEnabled() {
		return false
	}
	if r.opts.NoCleanup {
		return false
	}
	if os.Getenv("T_NO_CLEANUP") == "1" {
		return false
	}
	return true
}

// envSliceToMap converts the os/exec-style env slice subprocenv.Build
// returns into the map shape tfexec.SetEnv requires.
func envSliceToMap(env []string) map[string]string {
	out := make(map[string]string, len(env))
	for _, e := range env {
		k, v, ok := strings.Cut(e, "=")
		if ok {
			out[k] = v
		}
	}
	return out
}

// copyTerraformFiles copies every regular *.tf file from src into dst.
// Subdirectories are NOT recursed (test layouts are flat by
// convention); hidden files (e.g. .terraformrc) are skipped to avoid
// accidentally clobbering the framework's generated files. *.tfvars
// files copy too — they're occasionally useful for parameterized
// tests.
func copyTerraformFiles(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return fmt.Errorf("runner: mkdir %s: %w", dst, err)
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("runner: read source dir %s: %w", src, err)
	}
	copied := 0
	for _, e := range entries {
		if !shouldCopy(e) {
			continue
		}
		if err := copyFile(filepath.Join(src, e.Name()), filepath.Join(dst, e.Name())); err != nil {
			return err
		}
		copied++
	}
	if copied == 0 {
		return fmt.Errorf("runner: no .tf files found in %s", src)
	}
	return nil
}

// shouldCopy returns true for the file entries we copy from src to
// workdir. We accept *.tf and *.tfvars; subdirectories, hidden files,
// and other extensions are skipped.
func shouldCopy(e fs.DirEntry) bool {
	if !e.Type().IsRegular() {
		return false
	}
	name := e.Name()
	if strings.HasPrefix(name, ".") {
		return false
	}
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".tf" || ext == ".tfvars"
}

// copyFile reads src and writes it to dst with 0o644 permissions.
// Uses os.ReadFile / os.WriteFile rather than io.Copy because the
// individual *.tf files are small (KB) and the simpler API trades
// negligible memory for clarity.
func copyFile(src, dst string) error {
	body, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("runner: read %s: %w", src, err)
	}
	if err := os.WriteFile(dst, body, 0o644); err != nil {
		return fmt.Errorf("runner: write %s: %w", dst, err)
	}
	return nil
}

// errMissingRepoRoot is returned when a step requests version=local but
// the runner wasn't configured with a RepoRoot (M6 territory).
var errMissingRepoRoot = errors.New("runner: version=local requires Options.RepoRoot — not yet implemented (DESIGN.md §15 M6)")
