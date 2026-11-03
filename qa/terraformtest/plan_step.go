package terraformtest

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// configurePlan uses PreApply to check the saved plan that Terraform will apply.
// Native PlanOnly instead runs a non-refresh plan and a refresh plan; assertions
// belong on the latter so they observe the API fixtures' current values.
func configurePlan(native *resource.TestStep, s step, label string) {
	native.PlanOnly = s.PlanOnly
	check := planAssertion{label: label, expected: s.Plan}
	if s.PlanOnly {
		native.ExpectNonEmptyPlan = s.Plan.Expect == "non_empty"
		if s.State != nil && len(s.State.Resources) != 0 {
			check.state = s.State
		}
		native.ConfigStateChecks = nil
		native.ConfigPlanChecks.PostApplyPostRefresh = []plancheck.PlanCheck{check}
	} else if s.Plan != nil {
		native.ConfigPlanChecks.PreApply = []plancheck.PlanCheck{check}
	}
}
