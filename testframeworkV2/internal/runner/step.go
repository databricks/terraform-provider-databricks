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

// runStep executes one step end-to-end: cache resolve, override write,
// .terraform wipe, init, command, assert. The returned StepResult is
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

	if err := tfrcwriter.WriteVersionsOverride(workDir, syntheticVer); err != nil {
		return res, fmt.Errorf("runner: step %s: write version override: %w", step.Name, err)
	}
	if err := wipeTerraformState(workDir); err != nil {
		return res, fmt.Errorf("runner: step %s: %w", step.Name, err)
	}

	stdoutLog, stderrLog := r.stepLogPaths(runDir, idx, step.Name)
	res.StdoutLog, res.StderrLog = stdoutLog, stderrLog

	stderr, err := r.runCommand(ctx, step, workDir, env, stdoutLog, stderrLog)
	finalize(&res, step, err, stderr)
	return res, nil
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
// destroy) and returns the captured stderr bytes. tfexec writes
// stdout/stderr through the writers we set; we tee stderr into a
// memory buffer for assertion matching while also writing to the log
// file (DESIGN.md §7 — assertions match stderr only).
func (r *Runner) runCommand(ctx context.Context, step config.Step, workDir string, env map[string]string, stdoutLog, stderrLog string) ([]byte, error) {
	tf, err := r.tfFactory(workDir, r.opts.TerraformBin)
	if err != nil {
		return nil, fmt.Errorf("runner: step %s: tfexec factory: %w", step.Name, err)
	}
	if err := tf.SetEnv(env); err != nil {
		return nil, fmt.Errorf("runner: step %s: SetEnv: %w", step.Name, err)
	}
	stdoutF, err := os.Create(stdoutLog)
	if err != nil {
		return nil, fmt.Errorf("runner: step %s: open stdout log: %w", step.Name, err)
	}
	defer stdoutF.Close()
	stderrF, err := os.Create(stderrLog)
	if err != nil {
		return nil, fmt.Errorf("runner: step %s: open stderr log: %w", step.Name, err)
	}
	defer stderrF.Close()

	stderrBuf := &capturingWriter{}
	tf.SetStdout(stdoutF)
	tf.SetStderr(io.MultiWriter(stderrF, stderrBuf))

	if err := tf.Init(ctx); err != nil {
		return stderrBuf.Bytes(), err
	}
	cmdErr := dispatchCommand(ctx, tf, step.Command)
	return stderrBuf.Bytes(), cmdErr
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
