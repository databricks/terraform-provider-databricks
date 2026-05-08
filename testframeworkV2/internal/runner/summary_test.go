package runner

import (
	"errors"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
)

// realPlanWithChanges is a representative `terraform plan` output for
// a plan that creates one resource and destroys another. Captured
// against terraform 1.5.7 + the Databricks provider; trimmed to the
// summary-relevant tail. Tests don't depend on the full output —
// just the "Plan:" line our regex looks for — but using a realistic
// chunk catches accidental over-eager regex matches in the body.
const realPlanWithChanges = `Terraform used the selected providers to generate the following execution
plan. Resource actions are indicated with the following symbols:
  + create
  - destroy

Terraform will perform the following actions:

  # databricks_token.pat will be created
  + resource "databricks_token" "pat" {
      + comment          = "tfv2-token-lifecycle-step-1"
      + lifetime_seconds = 3600
    }

  # databricks_token.old will be destroyed
  - resource "databricks_token" "old" {}

Plan: 1 to add, 0 to change, 1 to destroy.
`

const realPlanNoOp = `databricks_token.pat: Refreshing state... [id=...]

No changes. Your infrastructure matches the configuration.

Terraform has compared your real infrastructure against your configuration
and found no differences, so no changes are needed.
`

const realApply = `databricks_token.pat: Creating...
databricks_token.pat: Creation complete after 1s [id=...]

Apply complete! Resources: 2 added, 1 changed, 1 destroyed.
`

const realDestroy = `databricks_token.pat: Destroying... [id=...]
databricks_token.pat: Destruction complete after 0s

Destroy complete! Resources: 1 destroyed.
`

const realApplyNoChanges = `databricks_token.pat: Refreshing state... [id=...]

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
`

const realPlanFailureStderr = `Error: cannot populate provider_config for mws workspaces: failed to resolve workspace_id: failed to get the workspace_id: strconv.ParseInt: parsing "": invalid syntax

  on main.tf line 19, in data "databricks_mws_workspaces" "all":
  19: data "databricks_mws_workspaces" "all" {
`

func TestSummarize_PlanNoOp(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandPlan,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realPlanNoOp),
	})
	if got != "no changes" {
		t.Errorf("got %q want %q", got, "no changes")
	}
}

func TestSummarize_PlanWithChanges(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandPlan,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realPlanWithChanges),
	})
	want := "1 added, 1 destroyed"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TestSummarize_PlanWithAllThreeFields covers the case where add,
// change, and destroy all have non-zero counts.
func TestSummarize_PlanWithAllThreeFields(t *testing.T) {
	stdout := "Plan: 2 to add, 3 to change, 1 to destroy."
	got := summarize(summaryInputs{command: config.CommandPlan, stdout: []byte(stdout)})
	want := "2 added, 3 changed, 1 destroyed"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSummarize_Apply(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandApply,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realApply),
	})
	want := "2 added, 1 changed, 1 destroyed"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSummarize_ApplyNoChanges(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandApply,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realApplyNoChanges),
	})
	if got != "no changes" {
		t.Errorf("got %q want %q", got, "no changes")
	}
}

func TestSummarize_Destroy(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandDestroy,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realDestroy),
	})
	want := "1 destroyed"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSummarize_DestroyNoChanges(t *testing.T) {
	stdout := "Destroy complete! Resources: 0 destroyed."
	got := summarize(summaryInputs{command: config.CommandDestroy, stdout: []byte(stdout)})
	if got != "no changes" {
		t.Errorf("got %q want %q", got, "no changes")
	}
}

// TestSummarize_FailureAsExpected covers the expect=failure happy
// path: the parser sees a non-nil cmdErr AND expect=failure and
// renders the "failure-as-expected" prefix with the truncated stderr
// first line.
func TestSummarize_FailureAsExpected(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandPlan,
		expect:  config.ExpectFailure,
		cmdErr:  errors.New("plan failed"),
		stderr:  []byte(realPlanFailureStderr),
	})
	if !strings.HasPrefix(got, "failure-as-expected: ") {
		t.Errorf("missing prefix: got %q", got)
	}
	// First stderr line is the "Error: ..." line — confirm the
	// excerpt grabs that, not the indented context lines.
	if !strings.Contains(got, "cannot populate provider_config") {
		t.Errorf("excerpt should be the Error line, got %q", got)
	}
	// Truncation: full Error line is much longer than 80 chars.
	const prefix = "failure-as-expected: "
	excerpt := strings.TrimPrefix(got, prefix)
	if len(excerpt) > 80 {
		t.Errorf("excerpt should be ≤ 80 chars, got %d: %q", len(excerpt), excerpt)
	}
	if !strings.HasSuffix(excerpt, "...") {
		t.Errorf("long excerpt should end in ellipsis, got %q", excerpt)
	}
}

// TestSummarize_UnexpectedError covers the success-path-but-error
// case (terraform died on a step the user expected to pass). The
// CLI renders Reason on FAIL steps, but Summary is still populated
// for use by other consumers (JSON output, future log helpers).
func TestSummarize_UnexpectedError(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandApply,
		expect:  config.ExpectSuccess,
		cmdErr:  errors.New("apply failed: provider error"),
		stderr:  []byte("Error: provider crashed\n  more context\n"),
	})
	if !strings.HasPrefix(got, "error: ") {
		t.Errorf("missing 'error:' prefix: got %q", got)
	}
	if !strings.Contains(got, "Error: provider crashed") {
		t.Errorf("expected stderr first line, got %q", got)
	}
}

// TestSummarize_UnexpectedErrorEmptyStderr covers the case where
// terraform errored without writing to stderr — the summary falls
// back to the cmdErr message.
func TestSummarize_UnexpectedErrorEmptyStderr(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandApply,
		expect:  config.ExpectSuccess,
		cmdErr:  errors.New("connection refused"),
		stderr:  []byte(""),
	})
	if !strings.HasPrefix(got, "error: ") {
		t.Errorf("missing 'error:' prefix: got %q", got)
	}
	if !strings.Contains(got, "connection refused") {
		t.Errorf("expected cmdErr text in fallback: got %q", got)
	}
}

// TestSummarize_AssertionsOKAppended covers the v2 happy path: a
// successful step with assertions that all passed gets the
// "· assertions ok" suffix.
func TestSummarize_AssertionsOKAppended(t *testing.T) {
	got := summarize(summaryInputs{
		command:       config.CommandApply,
		expect:        config.ExpectSuccess,
		stdout:        []byte(realApply),
		assertionsRan: true,
		assertionsOK:  true,
	})
	want := "2 added, 1 changed, 1 destroyed · assertions ok"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TestSummarize_AssertionsRanButFailedNoSuffix: assertionsRan=true +
// assertionsOK=false (mismatch) doesn't get the OK suffix. The
// runner separately flips Status to fail and populates Reason; the
// summary stays on the base outcome (which the CLI doesn't render
// on FAIL anyway).
func TestSummarize_AssertionsRanButFailedNoSuffix(t *testing.T) {
	got := summarize(summaryInputs{
		command:       config.CommandApply,
		expect:        config.ExpectSuccess,
		stdout:        []byte(realApply),
		assertionsRan: true,
		assertionsOK:  false,
	})
	if strings.Contains(got, "assertions ok") {
		t.Errorf("failed assertions should NOT get the ok suffix: %q", got)
	}
}

// TestSummarize_NoAssertionsNoSuffix: v1 specs (assertionsRan=false)
// don't get the suffix even on a passing step.
func TestSummarize_NoAssertionsNoSuffix(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandApply,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realApply),
	})
	if strings.Contains(got, "assertions") {
		t.Errorf("v1 spec should not mention assertions: %q", got)
	}
}

// TestSummarize_UnparseableStdout returns empty string — the CLI
// printer skips the suffix when summary == "". Tests an output that
// terraform never actually produces (defensive coverage).
func TestSummarize_UnparseableStdout(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandPlan,
		expect:  config.ExpectSuccess,
		stdout:  []byte("garbage that does not match any regex"),
	})
	if got != "" {
		t.Errorf("expected empty summary on unparseable output, got %q", got)
	}
}

// TestTruncateLine covers the helper directly to pin the ellipsis
// behaviour: short input passes through, long input gets capped with
// "..." suffix consuming the last 3 chars of the limit.
func TestTruncateLine(t *testing.T) {
	for _, tc := range []struct {
		in    string
		limit int
		want  string
	}{
		{"short", 80, "short"},
		{"exactly eighty chars long padding aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 80, "exactly eighty chars long padding aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"way too long for the eighty character limit and then some more text on the end pad", 80, "way too long for the eighty character limit and then some more text on the en..."},
		{"abcde", 3, "abc"}, // limit ≤ 3: hard cut, no ellipsis
	} {
		if got := truncateLine(tc.in, tc.limit); got != tc.want {
			t.Errorf("truncateLine(%q, %d):\n got %q\nwant %q", tc.in, tc.limit, got, tc.want)
		}
	}
}

// TestFirstStderrLine confirms we grab the first non-empty trimmed
// line — terraform's error-block convention.
func TestFirstStderrLine(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{"", ""},
		{"\n\n  \n", ""},
		{"Error: foo\n  context\n", "Error: foo"},
		{"   leading whitespace ignored   \nrest", "leading whitespace ignored"},
	} {
		if got := firstStderrLine([]byte(tc.in)); got != tc.want {
			t.Errorf("firstStderrLine(%q):\n got %q\nwant %q", tc.in, got, tc.want)
		}
	}
}

// TestFormatCounts covers the (added, changed, destroyed) →
// human-string formatter directly. Especially the "drop zero
// counts" rule.
func TestFormatCounts(t *testing.T) {
	for _, tc := range []struct {
		a, c, d int
		want    string
	}{
		{0, 0, 0, "no changes"},
		{1, 0, 0, "1 added"},
		{0, 1, 0, "1 changed"},
		{0, 0, 1, "1 destroyed"},
		{1, 0, 1, "1 added, 1 destroyed"},
		{2, 1, 1, "2 added, 1 changed, 1 destroyed"},
	} {
		if got := formatCounts(tc.a, tc.c, tc.d); got != tc.want {
			t.Errorf("formatCounts(%d,%d,%d) = %q want %q", tc.a, tc.c, tc.d, got, tc.want)
		}
	}
}

// TestParseInt covers the panic-free Atoi helper. Garbage input
// returns 0 (the regex guarantees digits-only in callers but the
// helper handles overflow defensively).
func TestParseInt(t *testing.T) {
	if got := parseInt("42"); got != 42 {
		t.Errorf("parseInt(42): got %d", got)
	}
	if got := parseInt("garbage"); got != 0 {
		t.Errorf("parseInt(garbage): got %d, want 0", got)
	}
}

// ═══════════════════════════════════════════════════════════
// Plan-content matcher Summary suffix tests (Task #34 / §17.10)
// ═══════════════════════════════════════════════════════════

// TestSummarize_PlanMatchOK_AppendsSuffix: a successful plan step
// with at least one plan-content matcher running clean appends
// "· plan-match ok" to the base summary.
func TestSummarize_PlanMatchOK_AppendsSuffix(t *testing.T) {
	got := summarize(summaryInputs{
		command:           config.CommandPlan,
		expect:            config.ExpectSuccess,
		stdout:            []byte(realPlanWithChanges),
		planAssertionsRan: true,
		planAssertionsOK:  true,
	})
	want := "1 added, 1 destroyed · plan-match ok"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TestSummarize_PlanMatchAndStateAssertOK: when both matcher
// suffixes apply (v2 mode + plan matcher), they both render. Order
// is plan-match first, then assertions, so the visual progression
// reads "what plan said" then "what state said".
func TestSummarize_PlanMatchAndStateAssertOK(t *testing.T) {
	got := summarize(summaryInputs{
		command:           config.CommandApply,
		expect:            config.ExpectSuccess,
		stdout:            []byte(realApply),
		assertionsRan:     true,
		assertionsOK:      true,
		planAssertionsRan: false, // apply doesn't run plan matchers
		planAssertionsOK:  false,
	})
	// Apply only carries assertion suffix. Plan matchers are gated
	// on command=plan (validation rejects on apply), so there's no
	// realistic case where both fire on the same step. We test the
	// plan-only and apply-only cases separately rather than a
	// composite that the schema forbids.
	want := "2 added, 1 changed, 1 destroyed · assertions ok"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TestSummarize_PlanMatchRanButFailedNoSuffix: planAssertionsRan=true
// + planAssertionsOK=false (the matcher fired and a failure surfaced)
// does NOT append the OK suffix. The runner separately flips Status
// to fail and populates Reason; the summary stays on the base
// outcome (CLI suppresses Summary on FAIL anyway).
func TestSummarize_PlanMatchRanButFailedNoSuffix(t *testing.T) {
	got := summarize(summaryInputs{
		command:           config.CommandPlan,
		expect:            config.ExpectSuccess,
		stdout:            []byte(realPlanWithChanges),
		planAssertionsRan: true,
		planAssertionsOK:  false,
	})
	if strings.Contains(got, "plan-match ok") {
		t.Errorf("failed plan matcher should NOT get the ok suffix: %q", got)
	}
}

// TestSummarize_NoPlanMatchersNoSuffix: a plan step without any
// matchers configured doesn't get the suffix even on a passing step.
func TestSummarize_NoPlanMatchersNoSuffix(t *testing.T) {
	got := summarize(summaryInputs{
		command: config.CommandPlan,
		expect:  config.ExpectSuccess,
		stdout:  []byte(realPlanWithChanges),
	})
	if strings.Contains(got, "plan-match") {
		t.Errorf("step without matchers should not mention plan-match: %q", got)
	}
}
