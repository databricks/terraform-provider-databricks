package runner

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
)

// summaryInputs is the bag of data the summary parser consumes. Kept
// as a struct (rather than a long parameter list) so callers can pass
// it positionally without coupling to the runner.
type summaryInputs struct {
	command           config.Command
	expect            config.Expect
	cmdErr            error  // non-nil iff the terraform subprocess returned an error
	stdout            []byte // captured terraform stdout (combined logs)
	stderr            []byte // captured terraform stderr
	assertionsRan     bool   // true when the step had >= 1 v2-mode assertion
	assertionsOK      bool   // true when assertionsRan AND no failures
	planAssertionsRan bool   // true when the step had ≥1 plan-content matcher
	planAssertionsOK  bool   // true when planAssertionsRan AND none failed
}

// summarize produces the short human-readable phrase that lands in
// StepResult.Summary. Pure function over its inputs — no FS, no exec,
// no globals. Expected output shapes per DESIGN dispatch:
//
// - Plan no-op: "no changes"
// - Plan with changes: "1 added, 1 destroyed"
// - Apply: "2 added, 1 changed"
// - Destroy: "1 destroyed"
// - expect=failure (regex match): "failure-as-expected: <80-char excerpt>"
// - unexpected error: "error: <80-char excerpt>"
// - v2 + passing state asserts: " · assertions ok" appended
// - plan-content matchers ok: " · plan-match ok" appended (in
// addition to any state-assertion suffix above)
//
// The empty-string return is the "no summary available" sentinel —
// the CLI printer skips the suffix entirely when summary == "".
func summarize(in summaryInputs) string {
	base := summarizeOutcome(in)
	if base == "" {
		return ""
	}
	if in.planAssertionsRan && in.planAssertionsOK {
		base += " · plan-match ok"
	}
	if in.assertionsRan && in.assertionsOK {
		base += " · assertions ok"
	}
	return base
}

// summarizeOutcome dispatches on (command, expect, cmdErr) to produce
// the core phrase. Assertion-OK suffixing happens in summarize.
func summarizeOutcome(in summaryInputs) string {
	if in.expect == config.ExpectFailure && in.cmdErr != nil {
		return "failure-as-expected: " + truncateLine(firstStderrLine(in.stderr), summaryExcerptLen)
	}
	if in.cmdErr != nil {
		// Unexpected error on a success-path step. The runner has
		// already populated res.Reason with the full message; here we
		// give the CLI a short one-liner.
		excerpt := truncateLine(firstStderrLine(in.stderr), summaryExcerptLen)
		if excerpt == "" {
			return "error: " + truncateLine(in.cmdErr.Error(), summaryExcerptLen)
		}
		return "error: " + excerpt
	}
	switch in.command {
	case config.CommandPlan:
		return summarizePlan(in.stdout)
	case config.CommandApply:
		return summarizeApply(in.stdout)
	case config.CommandDestroy:
		return summarizeDestroy(in.stdout)
	default:
		return ""
	}
}

// summaryExcerptLen caps stderr-derived excerpts. 80 chars matches the
// task spec and keeps the summary readable on standard 100-column
// terminals after the "[PASS] step N (name): version command in Xs"
// prefix.
const summaryExcerptLen = 80

// firstStderrLine returns the first non-empty trimmed line of stderr.
// terraform's `Error:` block is always the first line of a failure
// (followed by indented context), so this is the line a user wants
// to see. Returns empty string when stderr is empty / whitespace.
func firstStderrLine(stderr []byte) string {
	for line := range strings.SplitSeq(string(stderr), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

// truncateLine returns s capped at limit characters, with an ellipsis
// when truncation occurred. Preserves the structure of "..." callouts
// readers expect from CLI output.
func truncateLine(s string, limit int) string {
	if len(s) <= limit {
		return s
	}
	if limit <= 3 {
		return s[:limit]
	}
	return s[:limit-3] + "..."
}

// planRegexp matches terraform plan's summary line:
//
//	"Plan: 2 to add, 1 to change, 1 to destroy."
var planRegexp = regexp.MustCompile(`Plan: (\d+) to add, (\d+) to change, (\d+) to destroy`)

// summarizePlan parses `terraform plan` output. No-op plans contain
// "No changes." or "Your infrastructure matches the configuration";
// non-no-op plans have a "Plan: X to add, Y to change, Z to destroy"
// summary line.
func summarizePlan(stdout []byte) string {
	out := string(stdout)
	if strings.Contains(out, "No changes") || strings.Contains(out, "infrastructure matches the configuration") {
		return "no changes"
	}
	if m := planRegexp.FindStringSubmatch(out); m != nil {
		return formatCounts(parseInt(m[1]), parseInt(m[2]), parseInt(m[3]))
	}
	return ""
}

// applyRegexp matches terraform apply's summary line:
//
//	"Apply complete! Resources: 2 added, 1 changed, 1 destroyed."
var applyRegexp = regexp.MustCompile(`Apply complete! Resources: (\d+) added, (\d+) changed, (\d+) destroyed`)

// summarizeApply parses the `Apply complete!` summary. Apply output
// can describe a destroy (when the previous state had resources and
// the current config doesn't) by reporting destroyed > 0; the
// formatter handles that case naturally.
func summarizeApply(stdout []byte) string {
	if m := applyRegexp.FindStringSubmatch(string(stdout)); m != nil {
		return formatCounts(parseInt(m[1]), parseInt(m[2]), parseInt(m[3]))
	}
	return ""
}

// destroyRegexp matches terraform destroy's summary line:
//
//	"Destroy complete! Resources: 1 destroyed."
var destroyRegexp = regexp.MustCompile(`Destroy complete! Resources: (\d+) destroyed`)

func summarizeDestroy(stdout []byte) string {
	if m := destroyRegexp.FindStringSubmatch(string(stdout)); m != nil {
		n := parseInt(m[1])
		if n == 0 {
			return "no changes"
		}
		return fmt.Sprintf("%d destroyed", n)
	}
	return ""
}

// formatCounts renders (added, changed, destroyed) in human-friendly
// short form: drop zero-counts entirely, comma-separate the rest. An
// all-zero call returns "no changes" (the apply equivalent of plan's
// no-op shape).
func formatCounts(added, changed, destroyed int) string {
	parts := make([]string, 0, 3)
	if added > 0 {
		parts = append(parts, fmt.Sprintf("%d added", added))
	}
	if changed > 0 {
		parts = append(parts, fmt.Sprintf("%d changed", changed))
	}
	if destroyed > 0 {
		parts = append(parts, fmt.Sprintf("%d destroyed", destroyed))
	}
	if len(parts) == 0 {
		return "no changes"
	}
	return strings.Join(parts, ", ")
}

// parseInt is a panic-free strconv.Atoi alias — the regex already
// guaranteed the captured group is digits, so an error from Atoi here
// would only fire on integer overflow (>= 2^63). Treat that as zero
// rather than complicating the caller.
func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
