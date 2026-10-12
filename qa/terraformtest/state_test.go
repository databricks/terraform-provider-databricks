package terraformtest

import (
	"context"
	"encoding/json"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/stretchr/testify/require"
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
		},
	})
}

func groupAssertion(attributes map[string]any) stateAssertion {
	return stateAssertion{label: "step 2 (rename)", expected: expectedState{
		Resources: map[string]expectedResource{"databricks_group.this": {Attributes: attributes}},
	}}
}

func TestSelectedAttributesTolerateNewProperties(t *testing.T) {
	s := groupAssertion(map[string]any{"id": "group-id", "settings": map[string]any{"enabled": true}, "number": 42})
	var resp statecheck.CheckStateResponse
	s.CheckState(context.Background(), statecheck.CheckStateRequest{State: groupState()}, &resp)
	require.NoError(t, resp.Error)
}

func TestSelectedAttributeFailures(t *testing.T) {
	for _, tc := range []struct {
		name, property, message string
		expected                any
	}{
		{"wrong value", "display_name", "databricks_group.this.display_name", "other"},
		{"missing", "missing", "attribute missing", nil},
		{"nested wrong", "settings", "databricks_group.this.settings.enabled", map[string]any{"enabled": false}},
		{"array order", "members", "databricks_group.this.members", []any{"two", "one"}},
		{"array extra", "members", "databricks_group.this.members", []any{"one"}},
		{"wrong type", "settings", "databricks_group.this.settings", "not an object"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := groupAssertion(map[string]any{tc.property: tc.expected}).check(groupState())
			require.ErrorContains(t, err, "step 2 (rename)")
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestResourceInventoryIncludesModulesAndDataSources(t *testing.T) {
	s := groupAssertion(nil)
	state := groupState()
	state.Values.RootModule.ChildModules = []*tfjson.StateModule{{Resources: []*tfjson.StateResource{{Address: `module.child.data.databricks_group.lookup["a"]`}}}}
	require.ErrorContains(t, s.check(state), "state resource inventory")
	s.expected.Resources[`module.child.data.databricks_group.lookup["a"]`] = expectedResource{}
	require.NoError(t, s.check(state))
	delete(s.expected.Resources, "databricks_group.this")
	require.ErrorContains(t, s.check(state), "state resource inventory")
	require.ErrorContains(t, groupAssertion(nil).check(stateWithResources()), "state resource inventory")
}

func TestDestroyRequiresEmptyState(t *testing.T) {
	s := stateAssertion{label: "step 3 (destroy)", expected: expectedState{Resources: map[string]expectedResource{}}}
	require.NoError(t, s.check(&tfjson.State{}))
	require.NoError(t, s.check(stateWithResources()))
	require.ErrorContains(t, s.check(groupState()), "step 3 (destroy): state resource inventory")
	require.ErrorContains(t, s.check(nil), "Terraform state is unavailable")
}

func TestStateNumbersPreserveLargeIntegers(t *testing.T) {
	state := groupState()
	state.Values.RootModule.Resources[0].AttributeValues["number"] = json.Number("9007199254740993")
	require.NoError(t, groupAssertion(map[string]any{"number": int64(9007199254740993)}).check(state))
	require.ErrorContains(t, groupAssertion(map[string]any{"number": int64(9007199254740992)}).check(state), "databricks_group.this.number")
}

func TestNestedStateNumbersRemainDistinctFromStrings(t *testing.T) {
	state := groupState()
	state.Values.RootModule.Resources[0].AttributeValues["settings"] = map[string]any{"number": "42"}
	s := groupAssertion(map[string]any{"settings": map[string]any{"number": 42}})
	require.ErrorContains(t, s.check(state), `databricks_group.this.settings.number: expected 42 (number), got "42"`)
	state.Values.RootModule.Resources[0].AttributeValues["settings"] = map[string]any{"number": float64(42), "extra": true}
	require.NoError(t, s.check(state))
}
