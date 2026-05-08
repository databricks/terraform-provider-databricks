package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// printRunResult writes the run summary to w. The output deliberately
// avoids ANSI escape codes (CI consumers strip them; humans get a
// future `--color` flag in M7) and uses ASCII-only status markers so
// terminal-encoding edge cases don't garble the table.
//
// Format (per-step then summary):
//
//	[PASS] step 1 (passes_on_1_113_0): 1.113.0 plan in 1.2s   no changes
//	[FAIL] step 2 (fails_on_1_114_0): 1.114.0 plan in 0.8s — expected failure but command succeeded
//	       (full stderr at ~/.testframeworkv2/runs/<id>/step_2_fails_on_1_114_0.stderr.log)
//	...
//	test1: PASS (4/4 steps passed in 23.4s) — runDir: <path>
//
// On FAIL we always render the per-step stderr log path on the next
// line as an operator-debugging hint — useful regardless of whether
// the failure was an assertion mismatch, a regex no-match, or an
// unexpected terraform error.
func printRunResult(w io.Writer, r result.RunResult) {
	if r.Skipped {
		fmt.Fprintf(w, "%s: SKIPPED — %s\n", r.Test, r.Reason)
		return
	}
	for _, s := range r.Steps {
		fmt.Fprintln(w, formatStepLine(s))
		if hint := formatFailHint(s); hint != "" {
			fmt.Fprintln(w, hint)
		}
	}
	fmt.Fprintln(w, strings.Repeat("-", 60))
	fmt.Fprintf(w, "%s\n", r.String())
	if r.RunDir != "" {
		fmt.Fprintf(w, "run dir: %s\n", r.RunDir)
	}
}

// formatFailHint returns the indented hint block we render under
// each FAIL step. Empty string for PASS / SKIP — those don't need
// the hint, and emitting an empty line per step would dilute the
// table density.
//
// 7-space indent matches the visual width of the [FAIL] tag plus a
// trailing space, so the hint visibly threads under the tag column
// without colliding with the step number.
//
// Plan-assertion failures (PlanAssertions non-empty) get an inline
// stdout-tail excerpt before the pointer line — manual `cat`'ing of
// the stdout log is the part of the failure-triage loop we want to
// remove for plan-content matchers (DESIGN.md §17.10). Plain stderr-
// regex / unexpected-error failures keep the single-line shape.
func formatFailHint(s result.StepResult) string {
	if s.Status != result.StatusFail {
		return ""
	}
	if len(s.PlanAssertions) > 0 {
		return formatPlanFailHint(s)
	}
	if s.StderrLog == "" {
		return ""
	}
	return fmt.Sprintf("       (full stderr at %s)", s.StderrLog)
}

// planExcerptTailLines caps the inline plan-stdout excerpt at 15
// lines. 15 is enough to capture terraform's "Plan:" summary line
// plus the bulk of a typical resource diff without ballooning the
// CLI output. Tunable; widen if real-world failures show diffs
// being truncated mid-context.
const planExcerptTailLines = 15

// formatPlanFailHint renders the multi-line hint block for plan-
// assertion failures: an inline stdout-tail excerpt followed by a
// pointer line that mentions BOTH the full stdout log and stderr
// log paths. The split-pointer wording reflects the fact that the
// printer just showed the (truncated) stdout — a reader who wants
// the full thing should know exactly where to look.
func formatPlanFailHint(s result.StepResult) string {
	var b strings.Builder
	if excerpt := formatPlanStdoutExcerpt(s.Stdout); excerpt != "" {
		b.WriteString(excerpt)
		b.WriteString("\n")
	}
	switch {
	case s.StdoutLog != "" && s.StderrLog != "":
		fmt.Fprintf(&b, "       (full stdout at %s; stderr at %s)", s.StdoutLog, s.StderrLog)
	case s.StderrLog != "":
		fmt.Fprintf(&b, "       (full stderr at %s)", s.StderrLog)
	case s.StdoutLog != "":
		fmt.Fprintf(&b, "       (full stdout at %s)", s.StdoutLog)
	}
	return strings.TrimRight(b.String(), "\n")
}

// formatPlanStdoutExcerpt builds the "plan output (tail):" header +
// indented body block from raw stdout bytes. Trims leading/trailing
// blank lines so the excerpt doesn't waste vertical space on
// terraform's whitespace padding. Returns the empty-stdout
// placeholder line (no header) when stripped output is empty;
// callers should still render the pointer line afterwards.
//
// Body indent is 9 spaces (hangs 2 inside the 7-space header indent)
// for visual hierarchy: the header "plan output (tail):" sits at
// the FAIL column, the lines under it are clearly subordinate.
func formatPlanStdoutExcerpt(stdout []byte) string {
	lines := stripBlankEdges(strings.Split(string(stdout), "\n"))
	if len(lines) == 0 {
		return "       (plan stdout was empty)"
	}
	if len(lines) > planExcerptTailLines {
		lines = lines[len(lines)-planExcerptTailLines:]
	}
	var b strings.Builder
	b.WriteString("       plan output (tail):")
	for _, line := range lines {
		b.WriteString("\n         ")
		b.WriteString(line)
	}
	return b.String()
}

// stripBlankEdges drops leading and trailing whitespace-only lines
// from lines, preserving any interior blank lines. Used to keep
// plan-stdout excerpts tight without altering the meaningful diff
// shape (terraform inserts blank lines between resources, which
// readers want to keep).
func stripBlankEdges(lines []string) []string {
	isBlank := func(s string) bool { return strings.TrimSpace(s) == "" }
	start := 0
	for start < len(lines) && isBlank(lines[start]) {
		start++
	}
	end := len(lines)
	for end > start && isBlank(lines[end-1]) {
		end--
	}
	return lines[start:end]
}

// formatStepLine renders one step's outcome as a single line. The
// terraform-result Summary is appended on PASS with a 3-space gutter
// so it visually separates from the duration. On FAIL we prefer
// Reason (which carries the failure-specific diagnostic) over Summary
// — Summary's "failure-as-expected" / "error: ..." excerpt overlaps
// with Reason, and Reason is the one the user wants for a failed
// step. Skipped steps surface neither.
func formatStepLine(s result.StepResult) string {
	tag := "[PASS]"
	switch s.Status {
	case result.StatusFail:
		tag = "[FAIL]"
	case result.StatusSkipped:
		tag = "[SKIP]"
	}
	suffix := ""
	switch {
	case s.Status != result.StatusPass && s.Reason != "":
		suffix = " — " + s.Reason
	case s.Status == result.StatusPass && s.Summary != "":
		suffix = "   " + s.Summary
	}
	return fmt.Sprintf("%s step %d (%s): %s %s in %s%s",
		tag, s.Index+1, s.Name, s.Version, s.Command, formatDuration(s.Duration), suffix)
}

// formatDuration is a CLI-local twin of result.durationString. We
// don't reach into the result package's unexported helper; instead a
// tiny copy keeps the formatter self-contained.
func formatDuration(d any) string {
	type stringer interface{ String() string }
	if d == nil {
		return "0s"
	}
	if s, ok := d.(stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%v", d)
}
