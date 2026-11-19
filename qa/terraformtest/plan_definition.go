package terraformtest

import (
	"fmt"
	"maps"
	"slices"

	tfjson "github.com/hashicorp/terraform-json"
)

type expectedPlan struct {
	Expect  string                    `yaml:"expect,omitempty"`
	Changes map[string]expectedChange `yaml:"changes,omitempty"`
	Outputs map[string]expectedChange `yaml:"outputs,omitempty"`
}

// MarshalYAML preserves empty-map assertions instead of dropping them through
// omitempty or turning an omitted map into an assertion during fixture refresh.
func (p expectedPlan) MarshalYAML() (any, error) {
	result := map[string]any{}
	if p.Expect != "" {
		result["expect"] = p.Expect
	}
	if p.Changes != nil {
		result["changes"] = p.Changes
	}
	if p.Outputs != nil {
		result["outputs"] = p.Outputs
	}
	return result, nil
}

type expectedChange struct {
	Actions tfjson.Actions          `yaml:"actions,omitempty"`
	Before  []expectedPlanAttribute `yaml:"before,omitempty"`
	After   []expectedPlanAttribute `yaml:"after,omitempty"`
}

type expectedPlanAttribute struct {
	expectedAttribute `yaml:",inline"`
	Unknown           bool `yaml:"unknown,omitempty"`
}

func (s step) validatePlan() error {
	if s.Plan == nil && !s.PlanOnly {
		return nil
	}
	if s.ExpectError != nil {
		return fmt.Errorf("plan assertions and plan_only cannot be combined with expect_error; use a diagnostic step without plan assertions")
	}
	if s.PlanOnly && (s.Plan == nil || s.Plan.Expect == "") {
		return fmt.Errorf("plan_only requires plan.expect: empty or non_empty")
	}
	if s.Plan == nil {
		return nil
	}
	if s.Plan.Expect != "" && s.Plan.Expect != "empty" && s.Plan.Expect != "non_empty" {
		return fmt.Errorf("plan.expect must be empty or non_empty, or omitted")
	}
	if err := validateChanges(s.Plan.Changes, false); err != nil {
		return fmt.Errorf("plan.changes: %w", err)
	}
	if err := validateChanges(s.Plan.Outputs, true); err != nil {
		return fmt.Errorf("plan.outputs: %w", err)
	}
	return nil
}

func validateChanges(changes map[string]expectedChange, output bool) error {
	for _, address := range slices.Sorted(maps.Keys(changes)) {
		if address == "" {
			return fmt.Errorf("address must not be empty")
		}
		change := changes[address]
		if a := change.Actions; a != nil && !(a.NoOp() || a.Create() || a.Read() || a.Update() || a.Delete() || a.Replace() || a.Forget()) {
			return fmt.Errorf("%s: invalid actions %v", address, a)
		}
		for _, side := range []struct {
			name       string
			attributes []expectedPlanAttribute
		}{{"before", change.Before}, {"after", change.After}} {
			if err := validatePlanAttributes(side.attributes, output, side.name == "after"); err != nil {
				return fmt.Errorf("%s.%s: %w", address, side.name, err)
			}
		}
	}
	return nil
}

func validatePlanAttributes(attributes []expectedPlanAttribute, output, after bool) error {
	for i, attribute := range attributes {
		// Outputs may be scalar or indexed at the root. Reuse the normal path
		// validator by placing them under a temporary object key.
		path := attribute.expectedAttribute
		if output {
			if path.Path == nil {
				return fmt.Errorf("check %d: path is required; use [] for the whole output", i+1)
			}
			path.Path = append([]any{"output"}, path.Path...)
		}
		if _, err := path.attributePath(); err != nil {
			return fmt.Errorf("check %d: %w", i+1, err)
		}
		if attribute.Unknown {
			if !after || attribute.Equals.Kind != 0 {
				return fmt.Errorf("check %d: unknown: true requires an after value without equals", i+1)
			}
		} else if _, err := attribute.valueCheck(); err != nil {
			return fmt.Errorf("check %d: %w", i+1, err)
		}
	}
	return nil
}
