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

	// AssertLog is the absolute path to the per-step state-assertion
	// log file. Populated only on v2-mode steps that declared an
	// `assert:` block (DESIGN.md §17.5). Empty otherwise.
	AssertLog string `json:",omitempty"`

	// Assertions surfaces structured per-attribute assertion failures
	// for v2-mode steps. Populated only when at least one assertion
	// failed; nil/omitted on v1 runs and on passing v2 steps so the
	// JSON shape stays backwards-compatible (DESIGN.md §17.5 / §17.8).
	Assertions []AssertionFailure `json:",omitempty"`

	// Summary is a short human-readable phrase rendered at the end of
	// the step's CLI line — e.g. "no changes", "1 added, 1 destroyed",
	// "failure-as-expected: cannot populate provider_config...". The
	// runner's summary parser fills this in from terraform's stdout +
	// stderr + step.Expect + assertion outcome. omitempty so v1 JSON
	// output stays unchanged for runs that don't (or can't) extract a
	// summary.
	Summary string `json:",omitempty"`

	// PlanAssertions surfaces structured plan-content matcher failures
	// from `expect_non_empty_plan` / `plan_match` fields on the step
	// (DESIGN.md §17.10). Populated only when at least one matcher
	// failed; nil/omitted on plain plan steps and on passing steps so
	// the JSON shape stays backwards-compatible.
	PlanAssertions []PlanAssertionFailure `json:",omitempty"`

	// Stdout holds the captured terraform stdout bytes. Populated only
	// when PlanAssertions is non-empty, so the CLI printer can render
	// an inline excerpt under the FAIL line without re-reading
	// StdoutLog. Empty otherwise. Always omitted from JSON (the file
	// at StdoutLog is the canonical archive; this field is an
	// in-memory crutch for the printer).
	Stdout []byte `json:"-"`
}

// PlanAssertionFailure is one plan-content matcher mismatch. Kind is
// either "expect_non_empty_plan" or "plan_match" (the YAML field name
// that fired). Pattern carries the raw matcher value for diagnostic
// rendering — the regex source for plan_match, or "true" for
// expect_non_empty_plan. Reason is a short human-readable
// explanation.
type PlanAssertionFailure struct {
	Kind    string
	Pattern string `json:",omitempty"`
	Reason  string
}

// String returns a one-line human-readable form. Used by the runner
// to populate StepResult.Reason and to render in the CLI summary.
func (f PlanAssertionFailure) String() string {
	if f.Pattern != "" {
		return fmt.Sprintf("%s(%s): %s", f.Kind, f.Pattern, f.Reason)
	}
	return fmt.Sprintf("%s: %s", f.Kind, f.Reason)
}

// AssertionFailure is one structured per-attribute (or per-resource-
// presence) failure produced by stateassert.Run during a v2-mode step.
// Lives in the result package (rather than internal/stateassert) so
// it appears alongside StepResult in the JSON shape and external
// consumers (cmd/tfv2's printer, future programmatic callers) don't
// need to import stateassert just to render a step outcome.
//
// Address is the Terraform resource address from the test.yaml
// `assert[].resource:` field (e.g. `databricks_token.pat`). Reason
// is a short human-readable explanation. Field is non-empty only for
// per-attribute mismatches (Reason="value mismatch"); it's the
// dot-walked attribute key that didn't match. Expected and Actual
// are the YAML-decoded expected value and the JSON-decoded actual
// value respectively (both `any`).
type AssertionFailure struct {
	Address  string
	Reason   string
	Field    string `json:",omitempty"`
	Expected any    `json:",omitempty"`
	Actual   any    `json:",omitempty"`
}

// String returns a one-line human-readable form. Used by the runner
// to populate StepResult.Reason and to write per-line entries in the
// step's assert.log file.
func (f AssertionFailure) String() string {
	if f.Field != "" {
		return fmt.Sprintf("%s.%s: %s (expected=%v, actual=%v)",
			f.Address, f.Field, f.Reason, f.Expected, f.Actual)
	}
	return fmt.Sprintf("%s: %s", f.Address, f.Reason)
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
