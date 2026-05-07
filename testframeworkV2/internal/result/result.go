// Package result holds the result types the runner returns. Kept in its
// own package per DESIGN.md §3 so the CLI (cmd/tfv2) and any future
// programmatic callers can depend on the result shape without pulling in
// the full runner (and its tfexec dependency).
package result

import (
	"fmt"
	"strings"
	"time"
)

// Status reports whether a step or run passed, failed, or was skipped.
type Status string

const (
	StatusPass    Status = "pass"
	StatusFail    Status = "fail"
	StatusSkipped Status = "skipped"
)

// StepResult captures the outcome of a single test.yaml step. The runner
// fills it in incrementally as the step progresses; on a fail or skip,
// Reason is populated with a human-readable explanation.
type StepResult struct {
	Index            int
	Name             string
	Version          string
	SyntheticVersion string // "99.0.0-local" for local builds; otherwise == Version
	Command          string
	Expect           string

	Status   Status
	Reason   string
	Started  time.Time
	Duration time.Duration

	// StdoutLog and StderrLog are absolute paths to the per-step log
	// files. Populated regardless of pass/fail so debugging is
	// uniform.
	StdoutLog string
	StderrLog string
}

// RunResult is the top-level result Runner.Run returns.
type RunResult struct {
	Test    string
	Profile string
	RunDir  string

	// Skipped is true when the test's `requires` block didn't match the
	// host's profile (DESIGN.md §10/G9). Skipped runs return early
	// before any step executes; Steps is empty in that case.
	Skipped bool
	Reason  string

	Steps    []StepResult
	Started  time.Time
	Duration time.Duration
}

// AllPassed reports whether every step has Status=StatusPass. A skipped
// run trivially returns true (no steps were attempted, so nothing
// failed). A run with zero steps but Skipped=false also returns true,
// which is harmless since the config layer rejects step-less tests.
func (r RunResult) AllPassed() bool {
	if r.Skipped {
		return true
	}
	for _, s := range r.Steps {
		if s.Status != StatusPass {
			return false
		}
	}
	return true
}

// FailedSteps returns the indices of steps with Status != StatusPass.
// Useful for summary output and exit-code computation.
func (r RunResult) FailedSteps() []int {
	var out []int
	for _, s := range r.Steps {
		if s.Status != StatusPass {
			out = append(out, s.Index)
		}
	}
	return out
}

// String returns a single-line summary of the run, suitable for log
// banners. The full per-step breakdown is the caller's responsibility
// (e.g. a future result.Print helper in M7).
func (r RunResult) String() string {
	if r.Skipped {
		return fmt.Sprintf("%s: SKIPPED (%s)", r.Test, r.Reason)
	}
	pass := 0
	fail := 0
	for _, s := range r.Steps {
		switch s.Status {
		case StatusPass:
			pass++
		case StatusFail:
			fail++
		}
	}
	overall := "PASS"
	if fail > 0 {
		overall = "FAIL"
	}
	return fmt.Sprintf("%s: %s (%d/%d steps passed in %s)",
		r.Test, overall, pass, len(r.Steps), durationString(r.Duration))
}

// durationString rounds a Duration to ms for human display while keeping
// sub-second runs distinguishable. We avoid time.Duration's default
// formatting because for short runs it produces "734.2ms" which is
// noisier than "0.7s".
func durationString(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	s := d.Truncate(100 * time.Millisecond).String()
	// Drop trailing zeros in fractional seconds for readability:
	// "1.500s" → "1.5s", "2.000s" → "2s".
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
