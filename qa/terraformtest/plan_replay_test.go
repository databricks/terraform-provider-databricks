package terraformtest

import (
	"os"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestPlanReplay(t *testing.T) {
	Run(t, "testdata/plans")
}

func TestWrongPlanAssertionsFailBeforeApply(t *testing.T) {
	contents, err := os.ReadFile("testdata/plans/outputs.yaml")
	require.NoError(t, err)
	directory := t.TempDir()
	fixture := strings.Replace(string(contents), "equals: 42", "equals: 43", 1)
	writeDiscoveryFile(t, directory, "WrongPlan.yaml", fixture)
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "step 1 (preview_outputs): output answer.after")
}

func TestPlanOnlyRequiresReadAndDestroyFixtures(t *testing.T) {
	for _, missing := range []string{"read", "destroy"} {
		t.Run(missing, func(t *testing.T) {
			d, err := loadDefinition("testdata/initial_state/group_update.yaml")
			require.NoError(t, err)
			// Preview an update without applying it; the original checkpoint's
			// resource still belongs to the test and requires normal cleanup.
			d.Steps = d.Steps[:1]
			d.Steps[0].PlanOnly = true
			d.Steps[0].Plan = &expectedPlan{Expect: "non_empty"}
			d.Steps[0].State = nil
			d.Steps[0].Calls = d.Steps[0].Calls[:1]
			d.Destroy.Calls[0].Response.Body.JSON = map[string]any{"id": "group-id", "displayName": "Example group"}
			if missing == "read" {
				d.Steps[0].Calls = nil
			} else {
				d.Destroy.Calls = nil
			}
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, d.InitialState.Checkpoint, string(d.checkpoint))
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			writeDiscoveryFile(t, directory, "MissingFixture.yaml", string(contents))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			label := "step 1 (update_existing_group)"
			if missing == "destroy" {
				label = "step 2 (destroy)"
			}
			require.Contains(t, output, label+": unexpected HTTP call")
		})
	}
}

func TestPlanOnlyFirstStepDestroysInitialResource(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	d.Steps = d.Steps[:1]
	d.Steps[0].PlanOnly = true
	d.Steps[0].Plan = &expectedPlan{Expect: "non_empty"}
	d.Steps[0].State = nil
	d.Steps[0].Calls = d.Steps[0].Calls[:1]
	d.Destroy.Calls[0].Response.Body.JSON = map[string]any{"id": "group-id", "displayName": "Example group"}
	run(t, d)
}

func TestApplyPlanMismatchDoesNotRunUpdate(t *testing.T) {
	d, err := loadDefinition("testdata/plans/group_update.yaml")
	require.NoError(t, err)
	d.Steps = d.Steps[:1]
	change := d.Steps[0].Plan.Changes["databricks_group.this"]
	change.Actions = []tfjson.Action{tfjson.ActionCreate}
	d.Steps[0].Plan.Changes["databricks_group.this"] = change
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, d.InitialState.Checkpoint, string(d.checkpoint))
	contents, err := yaml.Marshal(d)
	require.NoError(t, err)
	writeDiscoveryFile(t, directory, "WrongAction.yaml", string(contents))
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "step 1 (update_existing_group): resource databricks_group.this: expected actions [create], got [update]")
	require.Contains(t, output, "unconsumed call 3: PUT", "pre-apply assertion must prevent the update")
}

func TestEmptyPlanOnlyDocumentsNativeNonRefreshRestriction(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	d.Steps = d.Steps[:1]
	d.Steps[0].PlanOnly = true
	d.Steps[0].Plan = &expectedPlan{Expect: "empty"}
	d.Steps[0].State = nil
	// The fixture would refresh the old checkpoint to the desired new value,
	// but HashiCorp rejects the non-refresh diff before making that API read.
	d.Steps[0].Calls = nil
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, d.InitialState.Checkpoint, string(d.checkpoint))
	contents, err := yaml.Marshal(d)
	require.NoError(t, err)
	writeDiscoveryFile(t, directory, "Drift.yaml", string(contents))
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "The non-refresh plan was not empty")
}
