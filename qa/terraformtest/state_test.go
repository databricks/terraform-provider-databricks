package terraformtest

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/stretchr/testify/require"
)

func stateWithResources(resources ...*tfjson.StateResource) *tfjson.State {
	return &tfjson.State{Values: &tfjson.StateValues{RootModule: &tfjson.StateModule{Resources: resources}}}
}

func groupState() *tfjson.State {
	return stateWithResources(&tfjson.StateResource{
		Address: "databricks_group.this",
	})
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
