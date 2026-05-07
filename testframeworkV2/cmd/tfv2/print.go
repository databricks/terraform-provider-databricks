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
//	[PASS] step 1 (passes_on_1_113_0): 1.113.0 plan in 1.2s
//	[FAIL] step 2 (fails_on_1_114_0): 1.114.0 plan in 0.8s — expected failure but command succeeded
//	...
//	test1: PASS (4/4 steps passed in 23.4s) — runDir: <path>
func printRunResult(w io.Writer, r result.RunResult) {
	if r.Skipped {
		fmt.Fprintf(w, "%s: SKIPPED — %s\n", r.Test, r.Reason)
		return
	}
	for _, s := range r.Steps {
		fmt.Fprintln(w, formatStepLine(s))
	}
	fmt.Fprintln(w, strings.Repeat("-", 60))
	fmt.Fprintf(w, "%s\n", r.String())
	if r.RunDir != "" {
		fmt.Fprintf(w, "run dir: %s\n", r.RunDir)
	}
}

// formatStepLine renders one step's outcome as a single line. Reason
// is appended only on FAIL; on PASS we omit it to keep the table dense.
func formatStepLine(s result.StepResult) string {
	tag := "[PASS]"
	switch s.Status {
	case result.StatusFail:
		tag = "[FAIL]"
	case result.StatusSkipped:
		tag = "[SKIP]"
	}
	suffix := ""
	if s.Status != result.StatusPass && s.Reason != "" {
		suffix = " — " + s.Reason
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
