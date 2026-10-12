package terraformtest

import (
	"context"
	"fmt"
	"maps"
	"slices"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

// emptyStateAssertion catches resources left behind by an incomplete destroy.
type emptyStateAssertion struct {
	label string
}

func (s emptyStateAssertion) CheckState(_ context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
	resp.Error = s.check(req.State)
}

func (s emptyStateAssertion) check(state *tfjson.State) error {
	actual, err := resourcesInState(state)
	if err != nil {
		return fmt.Errorf("%s: %w", s.label, err)
	}
	if len(actual) != 0 {
		return fmt.Errorf("%s: expected empty state after destroy, found resources %v", s.label, slices.Sorted(maps.Keys(actual)))
	}
	return nil
}

func resourcesInState(state *tfjson.State) (map[string]*tfjson.StateResource, error) {
	if state == nil {
		return nil, fmt.Errorf("terraform state is unavailable")
	}
	resources := map[string]*tfjson.StateResource{}
	if state.Values != nil {
		collectResources(state.Values.RootModule, resources)
	}
	return resources, nil
}

func collectResources(module *tfjson.StateModule, result map[string]*tfjson.StateResource) {
	if module == nil {
		return
	}
	for _, resource := range module.Resources {
		result[resource.Address] = resource
	}
	for _, child := range module.ChildModules {
		collectResources(child, result)
	}
}
