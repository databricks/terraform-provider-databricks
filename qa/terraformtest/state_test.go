package terraformtest

import (
	"context"
	"encoding/json"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func stateWithResources(resources ...*tfjson.StateResource) *tfjson.State {
	return &tfjson.State{Values: &tfjson.StateValues{RootModule: &tfjson.StateModule{Resources: resources}}}
}

func groupState() *tfjson.State {
	return stateWithResources(&tfjson.StateResource{
		Address: "databricks_group.this",
		AttributeValues: map[string]any{
			"id": "group-id", "display_name": "name", "new_property": true,
			"settings": map[string]any{"enabled": true, "extra": "ignored"},
			"members":  []any{"one", "two"}, "number": float64(42),
			"optional": nil, "items": []any{map[string]any{"name": "one", "extra": true}},
			"labels": map[string]any{"a.b[0]": "literal", "0": "zero", "": "empty key"},
		},
	})
}

func stateAttribute(t *testing.T, path []any, value any) expectedAttribute {
	t.Helper()
	var node yaml.Node
	require.NoError(t, node.Encode(value))
	return expectedAttribute{Path: path, Equals: node}
}

func groupAssertion(attributes ...expectedAttribute) stateAssertion {
	return stateAssertion{label: "step 2 (rename)", expected: expectedState{
		Resources: map[string]expectedResource{"databricks_group.this": {Attributes: attributes}},
	}}
}

func TestSelectedAttributesTolerateNewProperties(t *testing.T) {
	s := groupAssertion(
		stateAttribute(t, []any{"id"}, "group-id"),
		stateAttribute(t, []any{"settings", "enabled"}, true),
		stateAttribute(t, []any{"number"}, 42),
		stateAttribute(t, []any{"items", 0, "name"}, "one"),
	)
	var resp statecheck.CheckStateResponse
	s.CheckState(context.Background(), statecheck.CheckStateRequest{State: groupState()}, &resp)
	require.NoError(t, resp.Error)
}

func TestSelectedAttributeFailures(t *testing.T) {
	for _, tc := range []struct {
		name, message string
		path          []any
		expected      any
	}{
		{"wrong value", "databricks_group.this.display_name", []any{"display_name"}, "other"},
		{"missing is not null", "specified key missing not found", []any{"missing"}, nil},
		{"nested wrong", "databricks_group.this.settings.enabled", []any{"settings", "enabled"}, false},
		{"list element order", "databricks_group.this.members.0", []any{"members", 0}, "two"},
		{"list index out of range", "out of range", []any{"members", 2}, "three"},
		{"wrong type", "databricks_group.this.settings", []any{"settings"}, "not an object"},
		{"map traversal on scalar", "cannot convert object at MapStep", []any{"display_name", "enabled"}, true},
		{"list traversal on map", "cannot convert object at SliceStep", []any{"settings", 0}, true},
		{"expected null", "expected nil value", []any{"id"}, nil},
		{"actual null", "expected string value", []any{"optional"}, "set"},
		{"number versus bool", "expected bool value", []any{"number"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := groupAssertion(stateAttribute(t, tc.path, tc.expected)).check(groupState())
			require.ErrorContains(t, err, "step 2 (rename)")
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestNestedPathsUseLiteralMapKeysAndListIndices(t *testing.T) {
	s := groupAssertion(
		stateAttribute(t, []any{"labels", "a.b[0]"}, "literal"),
		stateAttribute(t, []any{"labels", "0"}, "zero"),
		stateAttribute(t, []any{"labels", ""}, "empty key"),
		stateAttribute(t, []any{"members", 1}, "two"),
		stateAttribute(t, []any{"optional"}, nil),
	)
	require.NoError(t, s.check(groupState()))
}

func TestSelectedResourcesIncludeModulesAndDataSources(t *testing.T) {
	s := groupAssertion()
	state := groupState()
	state.Values.RootModule.ChildModules = []*tfjson.StateModule{{Resources: []*tfjson.StateResource{{Address: `module.child.data.databricks_group.lookup["a"]`}}}}
	require.NoError(t, s.check(state))
	state.Values.RootModule.ChildModules[0].Resources[0].AttributeValues = map[string]any{"id": "lookup-id"}
	s.expected.Resources[`module.child.data.databricks_group.lookup["a"]`] = expectedResource{
		Attributes: []expectedAttribute{stateAttribute(t, []any{"id"}, "lookup-id")},
	}
	require.NoError(t, s.check(state))
	delete(s.expected.Resources, "databricks_group.this")
	require.NoError(t, s.check(state))
	state.Values.RootModule.ChildModules = nil
	require.ErrorContains(t, s.check(state), `step 2 (rename): module.child.data.databricks_group.lookup["a"]: resource missing from state`)
	require.ErrorContains(t, groupAssertion().check(stateWithResources()), "databricks_group.this: resource missing from state")
	require.ErrorContains(t, groupAssertion().check(nil), "terraform state is unavailable")
}

func TestEmptyStepExpectationsDoNotEnforceInventory(t *testing.T) {
	for _, resources := range []map[string]expectedResource{nil, {}} {
		s := stateAssertion{expected: expectedState{Resources: resources}}
		require.NoError(t, s.check(groupState()))
		require.NoError(t, s.check(stateWithResources()))
	}
}

func TestDestroyRequiresEmptyState(t *testing.T) {
	s := emptyStateAssertion{label: "step 3 (destroy)"}
	require.NoError(t, s.check(&tfjson.State{}))
	require.NoError(t, s.check(stateWithResources()))
	require.ErrorContains(t, s.check(groupState()), "step 3 (destroy): expected empty state after destroy, found resources [databricks_group.this]")
	require.ErrorContains(t, s.check(nil), "terraform state is unavailable")
	state := stateWithResources()
	state.Values.RootModule.ChildModules = []*tfjson.StateModule{{Resources: []*tfjson.StateResource{{Address: "module.child.data.databricks_group.lookup"}}}}
	require.ErrorContains(t, s.check(state), "found resources [module.child.data.databricks_group.lookup]")
}

func TestStateNumbersPreserveLargeIntegers(t *testing.T) {
	state := groupState()
	state.Values.RootModule.Resources[0].AttributeValues["number"] = json.Number("9007199254740993")
	require.NoError(t, groupAssertion(stateAttribute(t, []any{"number"}, int64(9007199254740993))).check(state))
	require.ErrorContains(t, groupAssertion(stateAttribute(t, []any{"number"}, int64(9007199254740992))).check(state), "databricks_group.this.number")
}

func TestNestedStateNumbersRemainDistinctFromStrings(t *testing.T) {
	state := groupState()
	state.Values.RootModule.Resources[0].AttributeValues["settings"] = map[string]any{"number": "42"}
	s := groupAssertion(stateAttribute(t, []any{"settings", "number"}, 42))
	require.ErrorContains(t, s.check(state), "databricks_group.this.settings.number: expected json.Number value")
	state.Values.RootModule.Resources[0].AttributeValues["settings"] = map[string]any{"number": float64(42), "extra": true}
	require.NoError(t, s.check(state))
}

func TestStateNumberChecksUseHashicorpPrecision(t *testing.T) {
	for _, tc := range []struct {
		name     string
		expected any
		actual   json.Number
	}{
		{"fraction", 1.2, "1.2"},
		{"exponent", 1e20, "100000000000000000000"},
		{"uint64", ^uint64(0), "18446744073709551615"},
		{"negative", -42, "-42.0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := groupState()
			state.Values.RootModule.Resources[0].AttributeValues["number"] = tc.actual
			require.NoError(t, groupAssertion(stateAttribute(t, []any{"number"}, tc.expected)).check(state))
		})
	}
}
