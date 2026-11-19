package terraformtest

import (
	"context"
	"fmt"
	"maps"
	"slices"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

type planAssertion struct {
	label    string
	expected *expectedPlan
	state    *expectedState
}

// CheckPlan inspects Terraform's supported plan representation, never a raw
// checkpoint. Plan-only state checks inspect the refreshed prior state, not
// planned values that have not been applied.
func (p planAssertion) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	if req.Plan == nil {
		resp.Error = fmt.Errorf("%s: terraform plan is unavailable", p.label)
		return
	}
	if err := p.check(ctx, req); err != nil {
		resp.Error = fmt.Errorf("%s: %w", p.label, err)
	}
}

func (p planAssertion) check(ctx context.Context, req plancheck.CheckPlanRequest) error {
	if p.expected != nil {
		var check plancheck.PlanCheck
		switch p.expected.Expect {
		case "empty":
			check = plancheck.ExpectEmptyPlan()
		case "non_empty":
			check = plancheck.ExpectNonEmptyPlan()
		}
		if check != nil {
			var response plancheck.CheckPlanResponse
			check.CheckPlan(ctx, req, &response)
			if response.Error != nil {
				return response.Error
			}
		}
		if err := p.checkChanges(req.Plan); err != nil {
			return err
		}
	}
	if p.state != nil {
		return (stateAssertion{label: "refreshed state", expected: *p.state}).check(req.Plan.PriorState)
	}
	return nil
}

func (p planAssertion) checkChanges(plan *tfjson.Plan) error {
	resources := map[string]*tfjson.Change{}
	for _, resource := range plan.ResourceChanges {
		resources[resource.Address] = resource.Change
	}
	if err := checkChangeMap("resource", p.expected.Changes, resources); err != nil {
		return err
	}
	return checkChangeMap("output", p.expected.Outputs, plan.OutputChanges)
}

// An explicitly empty map asserts no changes of that kind. Nonempty maps
// select addresses to check; unrelated resources, outputs, and fields are allowed.
func checkChangeMap(kind string, expected map[string]expectedChange, actual map[string]*tfjson.Change) error {
	if expected == nil {
		return nil
	}
	if len(expected) == 0 {
		for _, address := range slices.Sorted(maps.Keys(actual)) {
			if change := actual[address]; change != nil && !change.Actions.NoOp() {
				return fmt.Errorf("expected no %s changes, but %s has actions %v", kind, address, change.Actions)
			}
		}
	}
	for _, address := range slices.Sorted(maps.Keys(expected)) {
		change := actual[address]
		if change == nil {
			return fmt.Errorf("%s %s: missing from plan", kind, address)
		}
		want := expected[address]
		if want.Actions != nil && !slices.Equal(want.Actions, change.Actions) {
			return fmt.Errorf("%s %s: expected actions %v, got %v", kind, address, want.Actions, change.Actions)
		}
		if err := checkPlanAttributes(want.Before, change.Before, nil); err != nil {
			return fmt.Errorf("%s %s.before: %w", kind, address, err)
		}
		if err := checkPlanAttributes(want.After, change.After, change.AfterUnknown); err != nil {
			return fmt.Errorf("%s %s.after: %w", kind, address, err)
		}
	}
	return nil
}

func checkPlanAttributes(attributes []expectedPlanAttribute, actual, unknown any) error {
	for _, attribute := range attributes {
		path, _ := (expectedAttribute{Path: append([]any{"value"}, attribute.Path...)}).attributePath()
		isUnknown := unknownAtPath(unknown, attribute.Path)
		if attribute.Unknown {
			if !isUnknown {
				return fmt.Errorf("path %v: expected unknown value, got known value", attribute.Path)
			}
			continue
		}
		if isUnknown {
			return fmt.Errorf("path %v: expected known value, got unknown value", attribute.Path)
		}
		value, err := tfjsonpath.Traverse(map[string]any{"value": actual}, path)
		if err != nil {
			return err
		}
		if err := checkPlanValue(attribute, value); err != nil {
			return fmt.Errorf("path %v: %w", attribute.Path, err)
		}
	}
	return nil
}

func checkPlanValue(attribute expectedPlanAttribute, value any) error {
	check, err := attribute.valueCheck()
	if err != nil {
		return err
	}
	value, err = jsonValue(value)
	if err != nil {
		return err
	}
	return check.CheckValue(value)
}

// Unknown masks may mark an entire ancestor collection unknown, especially
// with SDKv2. A child of that collection is unknown too, not a missing/null value.
func unknownAtPath(mask any, path []any) bool {
	if unknown, ok := mask.(bool); ok {
		return unknown
	}
	if len(path) == 0 {
		return false
	}
	switch key := path[0].(type) {
	case string:
		if object, ok := mask.(map[string]any); ok {
			return unknownAtPath(object[key], path[1:])
		}
	case int:
		if list, ok := mask.([]any); ok && key >= 0 && key < len(list) {
			return unknownAtPath(list[key], path[1:])
		}
	}
	return false
}
