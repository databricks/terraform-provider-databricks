package terraformtest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"math/big"
	"reflect"
	"slices"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

type stateAssertion struct {
	label    string
	expected expectedState
}

func (s stateAssertion) CheckState(_ context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
	resp.Error = s.check(req.State)
}

// check only requires listed resources and attributes, keeping assertions
// focused on the behavior under test rather than the entire state shape.
func (s stateAssertion) check(state *tfjson.State) error {
	actual, err := resourcesInState(state)
	if err != nil {
		return fmt.Errorf("%s: %w", s.label, err)
	}
	for _, address := range slices.Sorted(maps.Keys(s.expected.Resources)) {
		resource, exists := actual[address]
		if !exists {
			return fmt.Errorf("%s: %s: resource missing from state", s.label, address)
		}
		if err := checkAttributes(address, s.expected.Resources[address].Attributes, resource.AttributeValues); err != nil {
			return fmt.Errorf("%s: %w", s.label, err)
		}
	}
	return nil
}

// emptyStateAssertion is mandatory cleanup validation, independent of optional
// per-step assertions supplied by a test author.
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

func checkAttributes(address string, expected, actual map[string]any) error {
	for _, name := range slices.Sorted(maps.Keys(expected)) {
		value, present := actual[name]
		if !present {
			return fmt.Errorf("%s.%s: attribute missing", address, name)
		}
		want, err := jsonValue(expected[name])
		if err != nil {
			return err
		}
		got, err := jsonValue(value)
		if err != nil {
			return err
		}
		if err := checkValue(address+"."+name, want, got); err != nil {
			return err
		}
	}
	return nil
}

// jsonValue normalizes YAML expectations and Terraform attribute values without
// rounding large numbers. Only attribute values are converted, not the state file.
func jsonValue(value any) (any, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var result any
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()
	err = decoder.Decode(&result)
	return normalizeStateNumbers(result), err
}

func normalizeStateNumbers(value any) any {
	switch v := value.(type) {
	case json.Number:
		n, _ := new(big.Rat).SetString(string(v))
		return n
	case map[string]any:
		for key, item := range v {
			v[key] = normalizeStateNumbers(item)
		}
	case []any:
		for i, item := range v {
			v[i] = normalizeStateNumbers(item)
		}
	}
	return value
}

func checkValue(path string, expected, actual any) error {
	if object, ok := expected.(map[string]any); ok {
		if actualObject, ok := actual.(map[string]any); ok {
			return checkObject(path, object, actualObject)
		}
	}
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("%s: expected %s, got %s", path, describeValue(expected), describeValue(actual))
	}
	return nil
}

func describeValue(value any) string {
	if number, ok := value.(*big.Rat); ok {
		return number.RatString() + " (number)"
	}
	return fmt.Sprintf("%#v", value)
}

// checkObject applies selected-property checks recursively; extra object keys
// are allowed for the same reason as unasserted top-level attributes.
func checkObject(path string, expected, actual map[string]any) error {
	for _, name := range slices.Sorted(maps.Keys(expected)) {
		value, present := actual[name]
		if !present {
			return fmt.Errorf("%s.%s: attribute missing", path, name)
		}
		if err := checkValue(path+"."+name, expected[name], value); err != nil {
			return err
		}
	}
	return nil
}
