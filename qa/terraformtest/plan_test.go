package terraformtest

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func planFromYAML(t *testing.T, input string) *expectedPlan {
	t.Helper()
	definition, err := decodeDefinition(strings.NewReader("version: 1\nsteps:\n  - name: plan\n    config: 'provider \"databricks\" {}'\n    plan: " + input + "\n"))
	require.NoError(t, err)
	return definition.Steps[0].Plan
}

func runPlanCheck(expected *expectedPlan, plan *tfjson.Plan) error {
	var resp plancheck.CheckPlanResponse
	(planAssertion{label: "step 2 (update)", expected: expected}).CheckPlan(context.Background(), plancheck.CheckPlanRequest{Plan: plan}, &resp)
	return resp.Error
}

func examplePlan() *tfjson.Plan {
	return &tfjson.Plan{
		ResourceChanges: []*tfjson.ResourceChange{
			{Address: "databricks_group.counted[0]", Change: &tfjson.Change{
				Actions:      tfjson.Actions{tfjson.ActionUpdate},
				Before:       map[string]any{"display_name": "old", "optional": nil},
				After:        map[string]any{"display_name": "new", "optional": nil, "id": nil, "large": json.Number("9007199254740993")},
				AfterUnknown: map[string]any{"id": true, "nested": true, "items": []any{map[string]any{"id": true}}},
			}},
			{Address: `module.child.databricks_group.keyed["a"]`, Change: &tfjson.Change{Actions: tfjson.Actions{tfjson.ActionCreate, tfjson.ActionDelete}}},
		},
		OutputChanges: map[string]*tfjson.Change{
			"answer": {Actions: tfjson.Actions{tfjson.ActionCreate}, After: map[string]any{"values": []any{json.Number("42")}}},
		},
	}
}

func TestPlanAssertionsSupportSelectedAddressesValuesAndOutputs(t *testing.T) {
	expected := planFromYAML(t, `
      expect: non_empty
      changes:
        databricks_group.counted[0]:
          actions: [update]
          before: [{path: [display_name], equals: old}]
          after:
            - {path: [display_name], equals: new}
            - {path: [optional], equals: null}
            - {path: [id], unknown: true}
            - {path: [nested, field], unknown: true}
            - {path: [items, 0, id], unknown: true}
            - {path: [large], equals: 9007199254740993}
        'module.child.databricks_group.keyed["a"]':
          actions: [create, delete]
      outputs:
        answer:
          actions: [create]
          after: [{path: [values, 0], equals: 42}]`)
	require.NoError(t, runPlanCheck(expected, examplePlan()))
	delete(expected.Changes, `module.child.databricks_group.keyed["a"]`)
	require.NoError(t, runPlanCheck(expected, examplePlan()), "unselected addresses are allowed")
}

func TestPlanAssertionsFailWithUsefulContext(t *testing.T) {
	for _, tc := range []struct {
		name, yaml, message string
	}{
		{"wrong action", "{changes: {'databricks_group.counted[0]': {actions: [create]}}}", "expected actions [create], got [update]"},
		{"replacement order", `{changes: {'module.child.databricks_group.keyed["a"]': {actions: [delete, create]}}}`, "expected actions [delete create], got [create delete]"},
		{"missing address", "{changes: {databricks_group.missing: {}}}", "missing from plan"},
		{"wrong before", "{changes: {'databricks_group.counted[0]': {before: [{path: [display_name], equals: wrong}]}}}", ".before:"},
		{"wrong after", "{changes: {'databricks_group.counted[0]': {after: [{path: [display_name], equals: wrong}]}}}", ".after:"},
		{"unknown is not null", "{changes: {'databricks_group.counted[0]': {after: [{path: [id], equals: null}]}}}", "expected known value, got unknown"},
		{"null is not unknown", "{changes: {'databricks_group.counted[0]': {after: [{path: [optional], unknown: true}]}}}", "expected unknown value, got known"},
		{"missing is not null", "{changes: {'databricks_group.counted[0]': {after: [{path: [missing], equals: null}]}}}", "path not found"},
		{"no changes", "{changes: {}}", "expected no resource changes"},
		{"no output changes", "{outputs: {}}", "expected no output changes"},
		{"empty plan", "{expect: empty}", "expected empty plan"},
		{"missing output", "{outputs: {missing: {}}}", "output missing: missing from plan"},
		{"typed output", "{outputs: {answer: {after: [{path: [values, 0], equals: '42'}]}}}", "expected string"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := runPlanCheck(planFromYAML(t, tc.yaml), examplePlan())
			require.ErrorContains(t, err, "step 2 (update):")
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestEmptyResourceChangesAreNotEmptyWholePlan(t *testing.T) {
	plan := &tfjson.Plan{
		ResourceChanges: []*tfjson.ResourceChange{{Address: "unchanged", Change: &tfjson.Change{Actions: tfjson.Actions{tfjson.ActionNoop}}}},
		OutputChanges:   map[string]*tfjson.Change{"answer": {Actions: tfjson.Actions{tfjson.ActionCreate}, After: 42}},
	}
	require.NoError(t, runPlanCheck(planFromYAML(t, "{changes: {}}"), plan))
	require.NoError(t, runPlanCheck(planFromYAML(t, "{expect: non_empty}"), plan))
	require.Error(t, runPlanCheck(planFromYAML(t, "{expect: empty}"), plan))
	plan.OutputChanges = nil
	require.NoError(t, runPlanCheck(planFromYAML(t, "{expect: empty, changes: {}}"), plan))
	require.Error(t, runPlanCheck(planFromYAML(t, "{expect: non_empty}"), plan))
	require.ErrorContains(t, runPlanCheck(nil, nil), "terraform plan is unavailable")
}

func TestPlanValidationRejectsUnsupportedAssertionsBeforeTerraform(t *testing.T) {
	for _, tc := range []struct {
		input, message string
	}{
		{"{expect: anything}", "plan.expect"},
		{"{changes: {'': {}}}", "address must not be empty"},
		{"{changes: {group: {actions: []}}}", "invalid actions"},
		{"{changes: {group: {actions: [create, update]}}}", "invalid actions"},
		{"{changes: {group: {actions: [replace]}}}", "invalid actions"},
		{"{changes: {group: {after: [{path: [id]}]}}}", "equals is required"},
		{"{changes: {group: {after: [{path: [id], equals: null, unknown: true}]}}}", "without equals"},
		{"{changes: {group: {before: [{path: [id], unknown: true}]}}}", "requires an after value"},
		{"{changes: {group: {after: [{path: [items, -1], equals: 1}]}}}", "nonnegative"},
		{"{changes: {group: {after: [{path: [id], equals: {a: b}}]}}}", "collection assertions"},
		{"{outputs: {result: {after: [{equals: 42}]}}}", "path is required"},
		{"{outputs: {result: {after: [{path: [0.5], equals: 42}]}}}", "nonnegative integer"},
		{"{typo: true}", "field typo not found"},
	} {
		t.Run(tc.input, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "    calls: []", "    plan: "+tc.input+"\n    calls: []", 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
	for _, input := range []string{"    plan_only: true\n", "    plan_only: true\n    plan: {}\n"} {
		_, err := decodeDefinition(strings.NewReader(strings.Replace(minimalDefinition, "    calls: []", input+"    calls: []", 1)))
		require.ErrorContains(t, err, "plan_only requires plan.expect")
	}
}

func TestPlanChecksCannotBeHiddenByExpectedErrors(t *testing.T) {
	for _, extra := range []string{"plan: {}", "plan_only: true"} {
		input := "version: 1\nsteps:\n  - name: invalid\n    config: 'provider \"databricks\" {}'\n    expect_error: '.*'\n    " + extra + "\n"
		_, err := decodeDefinition(strings.NewReader(input))
		require.ErrorContains(t, err, "cannot be combined with expect_error")
	}
}

func TestPlanHookSelectionAndCleanup(t *testing.T) {
	d := &testDefinition{Steps: []step{
		{Name: "apply", Config: "applied", Plan: &expectedPlan{}},
		{Name: "preview", Config: "not applied", PlanOnly: true, Plan: &expectedPlan{Expect: "non_empty"}, State: &expectedState{}},
	}}
	steps := testSteps(t, d, httpfixture.NewHandler())
	require.Len(t, steps[0].ConfigPlanChecks.PreApply, 1)
	require.False(t, steps[0].ExpectNonEmptyPlan, "post-apply consistency remains mandatory")
	require.Empty(t, steps[1].ConfigPlanChecks.PreApply)
	require.Len(t, steps[1].ConfigPlanChecks.PostApplyPostRefresh, 1)
	require.Empty(t, steps[1].ConfigStateChecks)
	require.True(t, steps[1].ExpectNonEmptyPlan)
	require.Equal(t, "applied", steps[2].Config)
	d.Steps = d.Steps[1:]
	steps = testSteps(t, d, httpfixture.NewHandler())
	require.Equal(t, "not applied", steps[1].Config, "plan-only checkpoint cleanup retains provider configuration")
}

func TestPlanPrimitiveAndUnknownOutputValues(t *testing.T) {
	for _, tc := range []struct {
		yaml   string
		change *tfjson.Change
	}{
		{"{outputs: {result: {after: [{path: [], unknown: true}]}}}", &tfjson.Change{AfterUnknown: true}},
		{"{outputs: {result: {after: [{path: [nested, 0], unknown: true}]}}}", &tfjson.Change{AfterUnknown: map[string]any{"nested": true}}},
		{"{outputs: {result: {after: [{path: [], equals: null}]}}}", &tfjson.Change{After: nil}},
		{"{outputs: {result: {after: [{path: [0], equals: true}]}}}", &tfjson.Change{After: []any{true}}},
	} {
		plan := &tfjson.Plan{OutputChanges: map[string]*tfjson.Change{"result": tc.change}}
		require.NoError(t, runPlanCheck(planFromYAML(t, tc.yaml), plan))
	}
}

func TestPlanUnknownChecksRoundTripYAML(t *testing.T) {
	expected := planFromYAML(t, "{changes: {group: {after: [{path: [id], unknown: true}]}}}")
	encoded, err := yaml.Marshal(expected)
	require.NoError(t, err)
	var decoded expectedPlan
	require.NoError(t, yaml.Unmarshal(encoded, &decoded))
	require.True(t, decoded.Changes["group"].After[0].Unknown)
	require.Zero(t, decoded.Changes["group"].After[0].Equals.Kind)
	require.NoError(t, validateChanges(decoded.Changes, false))
}

func TestPlanMapAssertionsSurviveYAMLRoundTrip(t *testing.T) {
	for _, input := range []string{"{}", "{changes: {}}", "{outputs: {}}", "{changes: {}, outputs: {}}"} {
		expected := planFromYAML(t, input)
		encoded, err := yaml.Marshal(expected)
		require.NoError(t, err)
		var decoded expectedPlan
		require.NoError(t, yaml.Unmarshal(encoded, &decoded))
		require.Equal(t, expected, &decoded)
	}
}
