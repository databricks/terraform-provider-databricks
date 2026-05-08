package runner

import (
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// canonicalPlanWithChanges is a representative `terraform plan` stdout
// for a non-empty plan. Used by the matcher tests to verify both the
// empty-plan check (should NOT contain "No changes") and the
// regex-anchor check (should match a "Plan: ..." pattern).
const canonicalPlanWithChanges = `Terraform used the selected providers to generate the following execution
plan. Resource actions are indicated with the following symbols:
  + create
  - destroy

Terraform will perform the following actions:

  # databricks_token.pat will be destroyed
  - resource "databricks_token" "pat" {
      - comment          = "tfv2-rollback-err-test" -> null
      - lifetime_seconds = 3600 -> null
        # forces replacement
    }

  # databricks_token.pat will be created
  + resource "databricks_token" "pat" {
      + comment          = "tfv2-rollback-err-test"
      + lifetime_seconds = 3600
    }

Plan: 1 to add, 0 to change, 1 to destroy.
`

const canonicalPlanNoOp = `databricks_token.pat: Refreshing state... [id=...]

No changes. Your infrastructure matches the configuration.
`

// TestRunPlanAssert_GateOnNonPassStatus: planassert is a no-op when
// the step is already failed by an earlier check (e.g. expect=success
// + cmdErr non-nil). PlanAssertions stays nil.
func TestRunPlanAssert_GateOnNonPassStatus(t *testing.T) {
	res := &result.StepResult{Status: result.StatusFail, Reason: "earlier failure"}
	runPlanAssert(res, config.Step{ExpectNonEmptyPlan: true}, []byte(canonicalPlanNoOp))
	if res.PlanAssertions != nil {
		t.Errorf("expected no plan assertions on already-failed step, got %v", res.PlanAssertions)
	}
}

// TestRunPlanAssert_GateOnNoMatchers: a step without either matcher
// is a no-op. PlanAssertions stays nil and Status stays Pass.
func TestRunPlanAssert_GateOnNoMatchers(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	runPlanAssert(res, config.Step{Command: config.CommandPlan}, []byte(canonicalPlanWithChanges))
	if res.PlanAssertions != nil || res.Status != result.StatusPass {
		t.Errorf("expected no-op, got status=%s assertions=%v", res.Status, res.PlanAssertions)
	}
}

// TestRunPlanAssert_ExpectNonEmptyPlan_PassesOnNonEmpty exercises
// the green path: an non-empty plan stdout passes the matcher.
func TestRunPlanAssert_ExpectNonEmptyPlan_PassesOnNonEmpty(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	runPlanAssert(res, config.Step{
		Command:            config.CommandPlan,
		ExpectNonEmptyPlan: true,
	}, []byte(canonicalPlanWithChanges))
	if res.Status != result.StatusPass {
		t.Errorf("expected pass, got %s (PlanAssertions=%v)", res.Status, res.PlanAssertions)
	}
	if len(res.PlanAssertions) != 0 {
		t.Errorf("expected no failures, got %v", res.PlanAssertions)
	}
}

// TestRunPlanAssert_ExpectNonEmptyPlan_FailsOnEmpty: the matcher
// fires when stdout contains "No changes".
func TestRunPlanAssert_ExpectNonEmptyPlan_FailsOnEmpty(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	runPlanAssert(res, config.Step{
		Command:            config.CommandPlan,
		ExpectNonEmptyPlan: true,
	}, []byte(canonicalPlanNoOp))
	if res.Status != result.StatusFail {
		t.Fatalf("expected fail, got %s", res.Status)
	}
	if len(res.PlanAssertions) != 1 {
		t.Fatalf("expected 1 failure, got %v", res.PlanAssertions)
	}
	f := res.PlanAssertions[0]
	if f.Kind != "expect_non_empty_plan" {
		t.Errorf("Kind: got %q", f.Kind)
	}
	if !strings.Contains(f.Reason, "No changes") {
		t.Errorf("Reason should mention No changes: %q", f.Reason)
	}
}

// TestRunPlanAssert_PlanMatch_PassesOnRegexMatch covers the green
// path for the regex matcher.
func TestRunPlanAssert_PlanMatch_PassesOnRegexMatch(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	re := regexp.MustCompile(`(?s)# forces replacement`)
	runPlanAssert(res, config.Step{
		Command:           config.CommandPlan,
		PlanMatch:         "# forces replacement",
		CompiledPlanMatch: re,
	}, []byte(canonicalPlanWithChanges))
	if res.Status != result.StatusPass {
		t.Errorf("expected pass, got %s", res.Status)
	}
	if len(res.PlanAssertions) != 0 {
		t.Errorf("expected no failures, got %v", res.PlanAssertions)
	}
}

// TestRunPlanAssert_PlanMatch_FailsOnRegexNoMatch covers the
// regex-not-found path.
func TestRunPlanAssert_PlanMatch_FailsOnRegexNoMatch(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	re := regexp.MustCompile(`(?s)pattern that is not present`)
	runPlanAssert(res, config.Step{
		Command:           config.CommandPlan,
		PlanMatch:         "pattern that is not present",
		CompiledPlanMatch: re,
	}, []byte(canonicalPlanWithChanges))
	if res.Status != result.StatusFail {
		t.Fatalf("expected fail, got %s", res.Status)
	}
	if len(res.PlanAssertions) != 1 {
		t.Fatalf("expected 1 failure, got %v", res.PlanAssertions)
	}
	if res.PlanAssertions[0].Kind != "plan_match" {
		t.Errorf("Kind: got %q", res.PlanAssertions[0].Kind)
	}
	if res.PlanAssertions[0].Pattern != "pattern that is not present" {
		t.Errorf("Pattern: got %q", res.PlanAssertions[0].Pattern)
	}
}

// TestRunPlanAssert_BothMatchers_BothPass exercises the AND
// composition green path: both matchers configured, both pass.
func TestRunPlanAssert_BothMatchers_BothPass(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	re := regexp.MustCompile(`(?s)Plan: 1 to add, 0 to change, 1 to destroy`)
	runPlanAssert(res, config.Step{
		Command:            config.CommandPlan,
		ExpectNonEmptyPlan: true,
		PlanMatch:          "Plan: 1 to add, 0 to change, 1 to destroy",
		CompiledPlanMatch:  re,
	}, []byte(canonicalPlanWithChanges))
	if res.Status != result.StatusPass {
		t.Errorf("expected pass, got %s (failures=%v)", res.Status, res.PlanAssertions)
	}
}

// TestRunPlanAssert_BothMatchers_CollectAllFailures pins the
// "never short-circuit" invariant: when both matchers fail, both
// failures are surfaced.
func TestRunPlanAssert_BothMatchers_CollectAllFailures(t *testing.T) {
	res := &result.StepResult{Status: result.StatusPass}
	re := regexp.MustCompile(`(?s)pattern not present`)
	runPlanAssert(res, config.Step{
		Command:            config.CommandPlan,
		ExpectNonEmptyPlan: true,
		PlanMatch:          "pattern not present",
		CompiledPlanMatch:  re,
	}, []byte(canonicalPlanNoOp))
	if res.Status != result.StatusFail {
		t.Fatalf("expected fail, got %s", res.Status)
	}
	if len(res.PlanAssertions) != 2 {
		t.Errorf("expected 2 failures (both matchers fired), got %d: %v", len(res.PlanAssertions), res.PlanAssertions)
	}
	if !strings.Contains(res.Reason, "expect_non_empty_plan") || !strings.Contains(res.Reason, "plan_match") {
		t.Errorf("Reason should mention both matchers, got %q", res.Reason)
	}
}

// TestPlanAssertionFailure_String covers the result.PlanAssertionFailure
// stringer for the with-Pattern and without-Pattern shapes.
func TestPlanAssertionFailure_String(t *testing.T) {
	withPattern := result.PlanAssertionFailure{Kind: "plan_match", Pattern: "foo", Reason: "no match"}
	if got, want := withPattern.String(), "plan_match(foo): no match"; got != want {
		t.Errorf("with pattern: got %q want %q", got, want)
	}
	noPattern := result.PlanAssertionFailure{Kind: "expect_non_empty_plan", Reason: "empty"}
	if got, want := noPattern.String(), "expect_non_empty_plan: empty"; got != want {
		t.Errorf("no pattern: got %q want %q", got, want)
	}
}
