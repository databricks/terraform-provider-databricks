package runner

import (
	"strings"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// runPlanAssert evaluates the step's plan-content matchers
// (`expect_non_empty_plan` and `plan_match`) against terraform plan's
// stdout. Per both fields are gated on
// `command: plan` AND `expect: success` at parse time, so this
// function only fires on success-path plan steps that previously
// passed `finalize` (Status == StatusPass).
//
// On any matcher failure: appends a structured PlanAssertionFailure
// to res.PlanAssertions, flips Status to fail, and rewrites Reason
// with the joined failure list. Multiple failures collect (don't
// short-circuit) so the operator sees both the empty-plan and
// regex-mismatch diagnostics in one shot when both fire.
//
// Pure function over (res, step, stdout) — no FS, no exec, no
// globals. Stdout is the captured terraform plan output the runner
// already plumbed through runCommand for the Summary parser.
func runPlanAssert(res *result.StepResult, step config.Step, stdout []byte) {
	if res.Status != result.StatusPass {
		return
	}
	if !step.ExpectNonEmptyPlan && step.CompiledPlanMatch == nil {
		return
	}
	failures := evaluatePlanAssertions(step, stdout)
	if len(failures) == 0 {
		return
	}
	res.Status = result.StatusFail
	res.PlanAssertions = failures
	// Stash the captured stdout so the CLI printer can render an
	// inline tail excerpt under the FAIL line. Stored only on
	// failure to keep memory bounded — the success path never holds
	// stdout bytes past the step boundary.
	res.Stdout = stdout
	parts := make([]string, len(failures))
	for i, f := range failures {
		parts[i] = f.String()
	}
	res.Reason = "plan assertion(s) failed: " + strings.Join(parts, "; ")
}

// evaluatePlanAssertions runs the matchers in YAML order
// (`expect_non_empty_plan` first, then `plan_match`). Pure for unit
// testing.
func evaluatePlanAssertions(step config.Step, stdout []byte) []result.PlanAssertionFailure {
	var failures []result.PlanAssertionFailure
	if step.ExpectNonEmptyPlan {
		// Terraform's stable phrase for empty plans is "No changes."
		// (period + capital N). The "infrastructure matches the
		// configuration" line is the verbose follow-up; either is
		// sufficient evidence. We anchor on "No changes." since it's
		// the shorter, version-stable token.
		if strings.Contains(string(stdout), "No changes") {
			failures = append(failures, result.PlanAssertionFailure{
				Kind:    "expect_non_empty_plan",
				Pattern: "true",
				Reason:  "plan was empty (stdout contains \"No changes\")",
			})
		}
	}
	if step.CompiledPlanMatch != nil {
		if !step.CompiledPlanMatch.Match(stdout) {
			failures = append(failures, result.PlanAssertionFailure{
				Kind:    "plan_match",
				Pattern: step.PlanMatch,
				Reason:  "plan stdout did not match",
			})
		}
	}
	return failures
}
