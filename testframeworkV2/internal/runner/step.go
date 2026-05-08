package runner

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/providercache"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/stateassert"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/tfrcwriter"
)

// runAllSteps executes every step in the spec sequentially. Per
// DESIGN.md §7, a single step's failure does NOT short-circuit the run
// — the framework's value proposition is multi-step regression where
// step N fails-as-expected and step N+1 still runs.
func (r *Runner) runAllSteps(ctx context.Context, workDir, runDir string, env map[string]string) ([]result.StepResult, error) {
	out := make([]result.StepResult, 0, len(r.spec.Steps))
	for i, s := range r.spec.Steps {
		// Per-step context honours the timeout from test.yaml. We use
		// a fresh ctx per step so a long-running step doesn't poison
		// later steps' deadlines.
		stepCtx, cancel := context.WithTimeout(ctx, s.Timeout)
		res, err := r.runStep(stepCtx, i, s, workDir, runDir, env)
		cancel()
		if err != nil {
			// Infrastructure failure (cache resolve, file IO, log
			// open). Surface as an error from Run so the CLI can
			// distinguish "framework broken" from "test failed".
			return out, err
		}
		out = append(out, res)
	}
	return out, nil
}

// runStep executes one step end-to-end: cache resolve, (v2-only)
// per-step config swap, override write, .terraform wipe, init,
// command, assert (state-level for v2). The returned StepResult is
// always populated; the error return is reserved for infrastructure
// failures (cache, FS) where we couldn't even attempt the command.
func (r *Runner) runStep(ctx context.Context, idx int, step config.Step, workDir, runDir string, env map[string]string) (result.StepResult, error) {
	res := result.StepResult{
		Index: idx, Name: step.Name, Version: step.Version,
		Command: string(step.Command), Expect: string(step.Expect),
		Started: r.nowFn(),
	}
	defer func() { res.Duration = r.nowFn().Sub(res.Started) }()

	syntheticVer, err := r.resolveStepVersion(ctx, step.Version, runDir)
	if err != nil {
		return res, err
	}
	res.SyntheticVersion = syntheticVer
	if err := r.prepareStepWorkdir(workDir, syntheticVer, step); err != nil {
		return res, err
	}

	stdoutLog, stderrLog := r.stepLogPaths(runDir, idx, step.Name)
	res.StdoutLog, res.StderrLog = stdoutLog, stderrLog

	stdout, stderr, cmdErr := r.runCommand(ctx, step, workDir, env, stdoutLog, stderrLog)
	finalize(&res, step, cmdErr, stderr)
	runPlanAssert(&res, step, stdout)
	r.runStateAssert(ctx, &res, step, workDir, runDir, env)
	res.Summary = summarizeStep(step, &res, cmdErr, stdout, stderr)
	return res, nil
}

// prepareStepWorkdir does the per-step workdir setup that happens
// before terraform init: v2-mode HCL swap, _tfv2_versions_override.tf
// write, .terraform/.terraform.lock.hcl wipe. Split from runStep so
// runStep stays under the 40-line CLAUDE.md threshold and the
// pre-init setup is one named, testable phase.
//
// v2 mode check uses spec.Mode() rather than step.Config != "" so v2
// specs with Config set (validated all-or-none at parse time) take
// this path uniformly. The swap deliberately runs BEFORE the
// override write so the wipe doesn't accidentally clobber
// _tfv2_versions_override.tf (the swap helper preserves _tfv2_* by
// name).
func (r *Runner) prepareStepWorkdir(workDir, syntheticVer string, step config.Step) error {
	if r.spec.Mode() == config.ModeV2 {
		if err := r.swapStepConfig(workDir, step.Config); err != nil {
			return fmt.Errorf("runner: step %s: swap config: %w", step.Name, err)
		}
	}
	if err := tfrcwriter.WriteVersionsOverride(workDir, syntheticVer); err != nil {
		return fmt.Errorf("runner: step %s: write version override: %w", step.Name, err)
	}
	if err := wipeTerraformState(workDir); err != nil {
		return fmt.Errorf("runner: step %s: %w", step.Name, err)
	}
	return nil
}

// summarizeStep is a thin adapter over summarize() that picks the
// right (assertionsRan, assertionsOK, planAssertionsRan,
// planAssertionsOK) flags from the step + res. Kept in step.go
// (rather than summary.go) so summary.go stays a pure parser with
// no result-struct knowledge.
func summarizeStep(step config.Step, res *result.StepResult, cmdErr error, stdout, stderr []byte) string {
	planMatchersRan := step.ExpectNonEmptyPlan || step.CompiledPlanMatch != nil
	return summarize(summaryInputs{
		command:           step.Command,
		expect:            step.Expect,
		cmdErr:            cmdErr,
		stdout:            stdout,
		stderr:            stderr,
		assertionsRan:     len(step.Assert) > 0,
		assertionsOK:      len(step.Assert) > 0 && len(res.Assertions) == 0 && res.Status == result.StatusPass,
		planAssertionsRan: planMatchersRan,
		planAssertionsOK:  planMatchersRan && len(res.PlanAssertions) == 0 && res.Status == result.StatusPass,
	})
}

// swapStepConfig is the v2-mode workdir swap: removes any non-
// framework *.tf / *.tfvars files (keeps _tfv2_*.tf and dot-files
// intact) and copies the step's per-step config from SourceDir into
// workDir.
//
// Path safety: configBase is constrained to a slug-shaped basename
// by the config-layer validator (v2ConfigPathRegexp); the join below
// can't traverse out of SourceDir.
func (r *Runner) swapStepConfig(workDir, configBase string) error {
	entries, err := os.ReadDir(workDir)
	if err != nil {
		return fmt.Errorf("read workdir %s: %w", workDir, err)
	}
	for _, e := range entries {
		if !e.Type().IsRegular() {
			continue
		}
		name := e.Name()
		if !shouldWipeForV2Swap(name) {
			continue
		}
		if err := os.Remove(filepath.Join(workDir, name)); err != nil {
			return fmt.Errorf("remove %s: %w", name, err)
		}
	}
	src := filepath.Join(r.opts.SourceDir, configBase)
	dst := filepath.Join(workDir, configBase)
	body, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read %s: %w", src, err)
	}
	if err := os.WriteFile(dst, body, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", dst, err)
	}
	return nil
}

// shouldWipeForV2Swap returns true for files the v2 wipe pass must
// remove from the workdir before copying the next step's config in.
// We wipe user-authored *.tf and *.tfvars; the framework-generated
// _tfv2_*.tf override stays (it's regenerated per step anyway, but
// removing it here would create a window where init fails on a
// missing override). Hidden files and other extensions are
// untouched.
func shouldWipeForV2Swap(name string) bool {
	if strings.HasPrefix(name, "_tfv2_") {
		return false
	}
	if strings.HasPrefix(name, ".") {
		return false
	}
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".tf" || ext == ".tfvars"
}

// runStateAssert evaluates the step's `assert:` block (if any)
// against `terraform show -json` output, writes a per-step assert.log
// with one OK/FAIL line per assertion, and surfaces failures via
// StepResult.Assertions + StepResult.Status flip (DESIGN.md §17.5 /
// §17.8).
//
// Only fires when (a) the step has assertions to evaluate AND (b) the
// step otherwise passed (config-layer validation already rejects
// assert: on expect=failure, so the gate here is a defense-in-depth
// "don't run terraform show on a failed step where state may be
// unparseable").
func (r *Runner) runStateAssert(ctx context.Context, res *result.StepResult, step config.Step, workDir, runDir string, env map[string]string) {
	if len(step.Assert) == 0 || res.Status != result.StatusPass {
		return
	}
	failures, err := stateassert.Run(ctx, workDir, r.opts.TerraformBin, env, step.Assert)
	if err != nil {
		res.Status = result.StatusFail
		res.Reason = fmt.Sprintf("stateassert: %v", err)
		return
	}
	res.AssertLog = r.stepAssertLogPath(runDir, res.Index, step.Name)
	if logErr := writeAssertLog(res.AssertLog, step.Assert, failures); logErr != nil {
		// Log-write failures are non-fatal — surface in stderr but
		// don't flip a passing step to fail just because we couldn't
		// open a log file. The structured Assertions slice is the
		// authoritative record either way.
		fmt.Fprintf(os.Stderr, "runner: step %s: write assert log: %v\n", step.Name, logErr)
	}
	if len(failures) == 0 {
		return
	}
	res.Status = result.StatusFail
	res.Assertions = failures
	parts := make([]string, len(failures))
	for i, f := range failures {
		parts[i] = f.String()
	}
	res.Reason = "state assertion(s) failed: " + strings.Join(parts, "; ")
}

// stepAssertLogPath returns the absolute path of the per-step
// assertion log under runDir. Naming mirrors the existing
// step_<n>_<name>.{stdout,stderr}.log convention with the
// `.assert.log` suffix (DESIGN.md §17.5).
func (r *Runner) stepAssertLogPath(runDir string, idx int, name string) string {
	return filepath.Join(runDir, fmt.Sprintf("step_%d_%s.assert.log", idx+1, name))
}

// writeAssertLog renders one OK/FAIL line per assertion into path.
// failures is the slice from stateassert.Run; assertions is the
// original Step.Assert list — we walk both to surface a "PASS" line
// for assertions that succeeded, in YAML order.
func writeAssertLog(path string, assertions []config.Assertion, failures []result.AssertionFailure) error {
	failuresByAddr := map[string][]result.AssertionFailure{}
	for _, f := range failures {
		failuresByAddr[f.Address] = append(failuresByAddr[f.Address], f)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, a := range assertions {
		fs := failuresByAddr[a.Resource]
		if len(fs) == 0 {
			fmt.Fprintf(f, "OK   %s\n", a.Resource)
			continue
		}
		for _, fail := range fs {
			fmt.Fprintf(f, "FAIL %s\n", fail.String())
		}
	}
	return nil
}

// resolveStepVersion converts a step.Version string into the
// syntheticVersion the runner pins via _tfv2_versions_override.tf.
//
// Released versions (e.g. "1.114.0") flow through providercache.Cache.
// Resolve, which downloads + atomically installs the zip on a cache
// miss. version="local" dispatches to providercache.Cache.BuildLocal,
// which (re)builds the provider from r.opts.RepoRoot every step
// (DESIGN.md §8 — "rebuild every step") and copies the provenance
// JSON into runDir.
func (r *Runner) resolveStepVersion(ctx context.Context, version, runDir string) (string, error) {
	if version == config.LocalVersion {
		return r.resolveLocalVersion(ctx, runDir)
	}
	_, syn, err := r.cache.Resolve(ctx, version, providercache.HostTarget())
	if err != nil {
		return "", fmt.Errorf("runner: cache resolve %q: %w", version, err)
	}
	return syn, nil
}

// resolveLocalVersion is the version=local arm of resolveStepVersion.
// Split out so resolveStepVersion stays small and the local path is
// single-purpose.
func (r *Runner) resolveLocalVersion(ctx context.Context, runDir string) (string, error) {
	if r.opts.RepoRoot == "" {
		return "", errMissingRepoRoot
	}
	_, syn, prov, err := r.cache.BuildLocal(ctx, r.opts.RepoRoot, providercache.HostTarget())
	if err != nil {
		return "", fmt.Errorf("runner: local build: %w", err)
	}
	// DESIGN.md §8 — write the run-dir copy of provenance so the run's
	// outcome is reproducible later. The cache-side copy is written
	// by BuildLocal itself.
	if err := providercache.CopyProvenanceTo(filepath.Join(runDir, providercache.LocalVersionFilename), prov); err != nil {
		return "", fmt.Errorf("runner: copy local-version.json into run dir: %w", err)
	}
	return syn, nil
}

// wipeTerraformState removes the per-step state from terraform's
// working directory: .terraform.lock.hcl AND the .terraform/ subdir
// (DESIGN.md §7.1.c / §10 G3). The state file (terraform.tfstate)
// lives at workdir root and is explicitly NOT touched — preserved
// across all steps.
func wipeTerraformState(workDir string) error {
	for _, name := range []string{".terraform.lock.hcl"} {
		if err := os.Remove(filepath.Join(workDir, name)); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove %s: %w", name, err)
		}
	}
	if err := os.RemoveAll(filepath.Join(workDir, ".terraform")); err != nil {
		return fmt.Errorf("remove .terraform: %w", err)
	}
	return nil
}

// stepLogPaths returns the (stdout, stderr) log paths for step idx /
// name under runDir. Naming follows task #9: step_<n>_<name>.{stdout,
// stderr}.log — 1-indexed because users count steps from 1, not 0.
func (r *Runner) stepLogPaths(runDir string, idx int, name string) (string, string) {
	base := fmt.Sprintf("step_%d_%s", idx+1, name)
	return filepath.Join(runDir, base+".stdout.log"), filepath.Join(runDir, base+".stderr.log")
}

// runCommand performs init + the requested command (plan/apply/
// destroy) and returns the captured (stdout, stderr) bytes plus the
// command error. tfexec writes stdout/stderr through the writers we
// set; we tee BOTH into in-memory buffers while also writing to the
// log files. stderr feeds the assertion regex match (DESIGN.md §7);
// stdout feeds the per-step Summary parser (Task #23).
//
// Init's stdout is captured separately and dropped — the summary
// parser only cares about the command's output, not init's
// boilerplate. Init failures still surface via cmdErr; the
// finalize/summarize path produces a useful error excerpt from
// stderr in that case.
func (r *Runner) runCommand(ctx context.Context, step config.Step, workDir string, env map[string]string, stdoutLog, stderrLog string) (cmdStdout, cmdStderr []byte, cmdErr error) {
	tf, err := r.tfFactory(workDir, r.opts.TerraformBin)
	if err != nil {
		return nil, nil, fmt.Errorf("runner: step %s: tfexec factory: %w", step.Name, err)
	}
	if err := tf.SetEnv(env); err != nil {
		return nil, nil, fmt.Errorf("runner: step %s: SetEnv: %w", step.Name, err)
	}
	stdoutF, err := os.Create(stdoutLog)
	if err != nil {
		return nil, nil, fmt.Errorf("runner: step %s: open stdout log: %w", step.Name, err)
	}
	defer stdoutF.Close()
	stderrF, err := os.Create(stderrLog)
	if err != nil {
		return nil, nil, fmt.Errorf("runner: step %s: open stderr log: %w", step.Name, err)
	}
	defer stderrF.Close()

	stderrBuf := &capturingWriter{}
	stdoutBuf := &capturingWriter{}
	// Init phase: stdout goes to log file only. We don't summarise
	// init output and capturing it would dilute the command's stdout
	// when the summary parser scans for "Plan: ..." / "Apply
	// complete!".
	tf.SetStdout(stdoutF)
	tf.SetStderr(io.MultiWriter(stderrF, stderrBuf))

	if err := tf.Init(ctx); err != nil {
		return stdoutBuf.Bytes(), stderrBuf.Bytes(), err
	}
	// Command phase: stdout tees to both log file AND in-memory
	// buffer for the summary parser.
	tf.SetStdout(io.MultiWriter(stdoutF, stdoutBuf))
	cmdErr = dispatchCommand(ctx, tf, step.Command)
	return stdoutBuf.Bytes(), stderrBuf.Bytes(), cmdErr
}

// dispatchCommand maps the test.yaml command enum onto the tfExec
// surface. Apply/Destroy run with auto-approve via the realTF wrapper
// (M4 mock factories don't care about that detail).
func dispatchCommand(ctx context.Context, tf tfExec, cmd config.Command) error {
	switch cmd {
	case config.CommandPlan:
		_, err := tf.Plan(ctx)
		return err
	case config.CommandApply:
		return tf.Apply(ctx)
	case config.CommandDestroy:
		return tf.Destroy(ctx)
	default:
		return fmt.Errorf("runner: unsupported command %q (config layer should have rejected this)", cmd)
	}
}

// finalize fills in Status + Reason on res based on the assertion
// outcome. Split from runStep so the assertion logic is unit-testable
// independently of FS / tfexec setup.
func finalize(res *result.StepResult, step config.Step, cmdErr error, stderr []byte) {
	switch step.Expect {
	case config.ExpectSuccess:
		if cmdErr == nil {
			res.Status = result.StatusPass
			return
		}
		res.Status = result.StatusFail
		res.Reason = fmt.Sprintf("expected success but got error: %v", cmdErr)
		return
	case config.ExpectFailure:
		finalizeFailure(res, step, cmdErr, stderr)
	}
}

// finalizeFailure handles the expect=failure path: the command MUST
// error, AND any error_substring / error_regex must match stderr
// (AND semantics when both present — DESIGN.md §4 / §7).
func finalizeFailure(res *result.StepResult, step config.Step, cmdErr error, stderr []byte) {
	if cmdErr == nil {
		res.Status = result.StatusFail
		res.Reason = "expected failure but command succeeded"
		return
	}
	if step.ErrorSubstring != "" && !strings.Contains(string(stderr), step.ErrorSubstring) {
		res.Status = result.StatusFail
		res.Reason = fmt.Sprintf("stderr did not contain %q", step.ErrorSubstring)
		return
	}
	if step.CompiledRegex != nil && !step.CompiledRegex.Match(stderr) {
		res.Status = result.StatusFail
		res.Reason = fmt.Sprintf("stderr did not match /%s/", step.CompiledRegex.String())
		return
	}
	res.Status = result.StatusPass
}

// runCleanup attempts a single destroy with the last successful
// Apply step's version. Failures are loud-logged but never
// fatal — destroy retries are out of scope (DESIGN.md §10 G12).
//
// In M4 the cleanup path is exercised via mocked tfexec to confirm we
// invoke the right sequence (writeOverride → wipe → init → destroy).
func (r *Runner) runCleanup(ctx context.Context, workDir, runDir string, env map[string]string, steps []result.StepResult) {
	last := lastSuccessfulApply(steps, r.spec.Steps)
	if last == nil {
		return
	}
	if err := tfrcwriter.WriteVersionsOverride(workDir, last.SyntheticVersion); err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup write override: %v\n", err)
		return
	}
	if err := wipeTerraformState(workDir); err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup wipe: %v\n", err)
		return
	}
	tf, err := r.tfFactory(workDir, r.opts.TerraformBin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup factory: %v\n", err)
		return
	}
	if err := tf.SetEnv(env); err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup SetEnv: %v\n", err)
		return
	}
	stdoutLog := filepath.Join(runDir, "cleanup.stdout.log")
	stderrLog := filepath.Join(runDir, "cleanup.stderr.log")
	stdoutF, _ := os.Create(stdoutLog)
	defer stdoutF.Close()
	stderrF, _ := os.Create(stderrLog)
	defer stderrF.Close()
	tf.SetStdout(stdoutF)
	tf.SetStderr(stderrF)

	if err := tf.Init(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup init: %v\n", err)
		return
	}
	if err := tf.Destroy(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "runner: cleanup destroy: %v — leaving resources for manual cleanup\n", err)
	}
}

// lastSuccessfulApply walks the executed step results in reverse and
// returns the most recent one whose corresponding spec step had
// Command=apply AND Status=pass. Returns nil when no Apply step
// succeeded — cleanup is then a no-op (DESIGN.md §7).
func lastSuccessfulApply(executed []result.StepResult, spec []config.Step) *result.StepResult {
	for i := len(executed) - 1; i >= 0; i-- {
		s := executed[i]
		if s.Status != result.StatusPass {
			continue
		}
		// Match back to the spec by index for the original Command
		// (StepResult.Command stores the rendered string).
		if i < len(spec) && spec[i].Command == config.CommandApply {
			return &s
		}
	}
	return nil
}

// capturingWriter is a tiny io.Writer that captures every Write into a
// growable byte buffer. We use it instead of bytes.Buffer so the
// Write method is concurrency-friendly with io.MultiWriter (which
// calls each underlying writer in sequence — bytes.Buffer would be
// fine here, but capturingWriter makes the intent explicit and lets us
// add bounds-checking later without API churn).
type capturingWriter struct{ buf []byte }

func (c *capturingWriter) Write(p []byte) (int, error) {
	c.buf = append(c.buf, p...)
	return len(p), nil
}

// Bytes returns a borrowed slice of the captured bytes. Callers must
// not modify the returned slice or use it after additional Writes.
func (c *capturingWriter) Bytes() []byte { return c.buf }
