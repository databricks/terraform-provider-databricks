package terraformtest

import (
	"context"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

const failedDestroyDefinition = `version: 1
steps:
  - name: create
    config: 'resource "databricks_group" "this" { display_name = "name" }'
    calls: []
destroy:
  refresh: false
  expect_error: 'cannot delete group'
  state:
    resources:
      databricks_group.this:
        attributes:
          - path: [display_name]
            equals: name
  calls: []
`

func TestDecodeDestroyExpectedErrors(t *testing.T) {
	kept := "\n  state:\n    resources:\n      ADDRESS:\n        attributes: []\n"
	for _, tc := range []struct {
		name, destroy, message string
	}{
		{"valid", "  expect_error: 'rejected'" + strings.Replace(kept, "ADDRESS", "databricks_group.this", 1), ""},
		{"state without error", strings.TrimPrefix(strings.Replace(kept, "ADDRESS", "databricks_group.this", 1), "\n"), "state requires expect_error"},
		{"error without state", "  expect_error: 'rejected'\n", "requires state listing every resource"},
		{"empty pattern", "  expect_error: ' '" + strings.Replace(kept, "ADDRESS", "databricks_group.this", 1), "nonempty regular expression"},
		{"invalid pattern", "  expect_error: '['" + strings.Replace(kept, "ADDRESS", "databricks_group.this", 1), "invalid expect_error regular expression"},
		{"data source", "  expect_error: 'rejected'" + strings.Replace(kept, "ADDRESS", "data.databricks_group.this", 1), "root-module managed resource without an index"},
		{"indexed", "  expect_error: 'rejected'" + strings.Replace(kept, "ADDRESS", `'databricks_group.this["a"]'`, 1), "root-module managed resource without an index"},
		{"module", "  expect_error: 'rejected'" + strings.Replace(kept, "ADDRESS", "module.m.databricks_group.this", 1), "root-module managed resource without an index"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := "version: 1\nsteps:\n  - name: create\n    config: 'provider \"databricks\" {}'\n    calls: []\ndestroy:\n" + tc.destroy + "  calls: []\n"
			_, err := decodeDefinition(strings.NewReader(input))
			if tc.message == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.message)
			require.True(t, strings.HasPrefix(err.Error(), "destroy: "), err.Error())
		})
	}
}

func TestFailedDestroyForgetsKeptResourcesWithoutRefresh(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(failedDestroyDefinition))
	require.NoError(t, err)
	tc := terraformCase(t, d, httpfixture.NewHandler(), fixtureConfig("http://127.0.0.1"), nil)
	require.Len(t, tc.Steps, 3)
	destroy, forget := tc.Steps[1], tc.Steps[2]
	require.True(t, destroy.Destroy)
	require.True(t, destroy.ExpectError.MatchString("cannot delete group"))
	require.Empty(t, destroy.ConfigStateChecks)

	require.False(t, tc.AdditionalCLIOptions.Plan.NoRefresh)
	destroy.PreConfig()
	require.True(t, tc.AdditionalCLIOptions.Plan.NoRefresh)
	forget.PreConfig()
	require.True(t, tc.AdditionalCLIOptions.Plan.NoRefresh)

	require.False(t, forget.Destroy)
	require.Nil(t, forget.ExpectError)
	require.Contains(t, forget.Config, "removed {\n  from = databricks_group.this\n  lifecycle {\n    destroy = false\n  }\n}\n")
	var plan plancheck.CheckPlanResponse
	forget.ConfigPlanChecks.PreApply[0].CheckPlan(context.Background(), plancheck.CheckPlanRequest{Plan: &tfjson.Plan{PriorState: groupState()}}, &plan)
	require.NoError(t, plan.Error)
	forget.ConfigPlanChecks.PreApply[0].CheckPlan(context.Background(), plancheck.CheckPlanRequest{Plan: &tfjson.Plan{PriorState: stateWithResources()}}, &plan)
	require.ErrorContains(t, plan.Error, "step 2 (destroy): databricks_group.this: resource missing from state")
	var state statecheck.CheckStateResponse
	forget.ConfigStateChecks[0].CheckState(context.Background(), statecheck.CheckStateRequest{State: groupState()}, &state)
	require.ErrorContains(t, state.Error, "step 3 (forget): expected empty state after destroy")
}

func TestDestroyWithoutExpectedErrorAddsNoForgetStep(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(strings.Replace(minimalDefinition, "destroy:\n", "destroy:\n  refresh: false\n", 1)))
	require.NoError(t, err)
	tc := terraformCase(t, d, httpfixture.NewHandler(), fixtureConfig("http://127.0.0.1"), nil)
	require.Len(t, tc.Steps, 2)
	tc.Steps[1].PreConfig()
	require.True(t, tc.AdditionalCLIOptions.Plan.NoRefresh)
	require.Len(t, tc.Steps[1].ConfigStateChecks, 1)

	d, err = decodeDefinition(strings.NewReader(minimalDefinition))
	require.NoError(t, err)
	require.Nil(t, terraformCase(t, d, httpfixture.NewHandler(), fixtureConfig("http://127.0.0.1"), nil).AdditionalCLIOptions)
}

func TestRecordingRejectsExpectedDestroyFailure(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(failedDestroyDefinition))
	require.NoError(t, err)
	require.ErrorContains(t, validateRecordableCalls(d), "recording cannot leave the resources a failed destroy keeps")
}
