package terraformtest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"math/big"
	"slices"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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

func checkAttributes(address string, expected []expectedAttribute, actual map[string]any) error {
	for _, attribute := range expected {
		path, err := attribute.attributePath()
		if err != nil {
			return err
		}
		check, err := attribute.valueCheck()
		if err != nil {
			return err
		}
		value, err := tfjsonpath.Traverse(actual, path)
		if err != nil {
			return fmt.Errorf("%s.%s: %w", address, path, err)
		}
		got, err := jsonValue(value)
		if err != nil {
			return fmt.Errorf("%s.%s: %w", address, path, err)
		}
		if err := check.CheckValue(got); err != nil {
			return fmt.Errorf("%s.%s: %w", address, path, err)
		}
	}
	return nil
}

// attributePath keeps literal map keys distinct from list indices, without
// parsing dots or brackets in attribute names.
func (a expectedAttribute) attributePath() (tfjsonpath.Path, error) {
	if len(a.Path) == 0 {
		return tfjsonpath.Path{}, fmt.Errorf("path must start with an attribute name")
	}
	name, ok := a.Path[0].(string)
	if !ok || name == "" {
		return tfjsonpath.Path{}, fmt.Errorf("path must start with an attribute name")
	}
	path := tfjsonpath.New(name)
	for _, segment := range a.Path[1:] {
		switch v := segment.(type) {
		case string:
			path = path.AtMapKey(v)
		case int:
			if v < 0 {
				return tfjsonpath.Path{}, fmt.Errorf("path indices must be nonnegative")
			}
			path = path.AtSliceIndex(v)
		default:
			return tfjsonpath.Path{}, fmt.Errorf("path segments must be string keys or nonnegative integer indices")
		}
	}
	return path, nil
}

// valueCheck delegates typed comparisons to HashiCorp instead of assuming
// Terraform collections have an exact, ordered JSON representation.
func (a expectedAttribute) valueCheck() (knownvalue.Check, error) {
	if a.Equals.Kind == 0 {
		return nil, fmt.Errorf("equals is required; use equals: null to assert null")
	}
	var expected any
	if err := a.Equals.Decode(&expected); err != nil {
		return nil, err
	}
	value, err := jsonValue(expected)
	if err != nil {
		return nil, fmt.Errorf("equals must be a JSON primitive or null: %w", err)
	}
	switch v := value.(type) {
	case nil:
		return knownvalue.Null(), nil
	case string:
		// Non-JSON YAML scalars (such as timestamps) must not become strings silently.
		if _, ok := expected.(string); !ok {
			break
		}
		return knownvalue.StringExact(v), nil
	case bool:
		return knownvalue.Bool(v), nil
	case json.Number:
		number, _, err := big.ParseFloat(v.String(), 10, 512, big.ToNearestEven)
		return knownvalue.NumberExact(number), err
	}
	return nil, fmt.Errorf("equals must be a JSON primitive or null; collection assertions are not supported")
}

// jsonValue preserves numeric types and large integers for knownvalue helpers.
// It converts only a selected value, never reads or rewrites a state file.
func jsonValue(value any) (any, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var result any
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()
	err = decoder.Decode(&result)
	return result, err
}
