package common

import (
	"context"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a mock plan with resource changes
func createMockPlan(changes []*tfjson.ResourceChange) *tfjson.Plan {
	return &tfjson.Plan{
		ResourceChanges: changes,
	}
}

// Helper function to create a resource change
func createResourceChange(address string, actions ...tfjson.Action) *tfjson.ResourceChange {
	return &tfjson.ResourceChange{
		Address: address,
		Change: &tfjson.Change{
			Actions: actions,
		},
	}
}

func TestFindResourceChange(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		searchAddress   string
		expectError     bool
		expectedAddress string
		errorContains   []string
	}{
		{
			name: "finds existing resource",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionCreate),
				createResourceChange("resource.bar", tfjson.ActionUpdate),
			},
			searchAddress:   "resource.bar",
			expectError:     false,
			expectedAddress: "resource.bar",
		},
		{
			name: "returns error when resource not found",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionCreate),
				createResourceChange("resource.bar", tfjson.ActionUpdate),
			},
			searchAddress: "resource.nonexistent",
			expectError:   true,
			errorContains: []string{"address resource.nonexistent not found", "resource.foo, resource.bar"},
		},
		{
			name:            "handles empty plan",
			resourceChanges: []*tfjson.ResourceChange{},
			searchAddress:   "resource.foo",
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}

			change, err := findResourceChange(req, tt.searchAddress)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, change)
				for _, substr := range tt.errorContains {
					assert.Contains(t, err.Error(), substr)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, change)
				assert.Equal(t, tt.expectedAddress, change.Address)
			}
		})
	}
}

func TestGetPlannedActions(t *testing.T) {
	tests := []struct {
		name            string
		actions         []tfjson.Action
		expectedActions []string
	}{
		{
			name:            "returns single action",
			actions:         []tfjson.Action{tfjson.ActionCreate},
			expectedActions: []string{"create"},
		},
		{
			name:            "returns multiple actions",
			actions:         []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate},
			expectedActions: []string{"delete", "create"},
		},
		{
			name:            "handles no actions",
			actions:         []tfjson.Action{},
			expectedActions: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			change := createResourceChange("resource.foo", tt.actions...)
			actions := getPlannedActions(change)
			assert.Equal(t, tt.expectedActions, actions)
		})
	}
}

func TestHasAction(t *testing.T) {
	tests := []struct {
		name         string
		actions      []tfjson.Action
		targetAction tfjson.Action
		expected     bool
	}{
		{
			name:         "finds action in single action list",
			actions:      []tfjson.Action{tfjson.ActionCreate},
			targetAction: tfjson.ActionCreate,
			expected:     true,
		},
		{
			name:         "action not in single action list",
			actions:      []tfjson.Action{tfjson.ActionCreate},
			targetAction: tfjson.ActionUpdate,
			expected:     false,
		},
		{
			name:         "finds delete in multiple actions",
			actions:      []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate},
			targetAction: tfjson.ActionDelete,
			expected:     true,
		},
		{
			name:         "finds create in multiple actions",
			actions:      []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate},
			targetAction: tfjson.ActionCreate,
			expected:     true,
		},
		{
			name:         "action not in multiple actions",
			actions:      []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate},
			targetAction: tfjson.ActionUpdate,
			expected:     false,
		},
		{
			name:         "returns false for no actions",
			actions:      []tfjson.Action{},
			targetAction: tfjson.ActionCreate,
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			change := createResourceChange("resource.foo", tt.actions...)
			result := hasAction(change, tt.targetAction)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCheckActionPresence(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		action          tfjson.Action
		shouldBePresent bool
		expectError     bool
		errorContains   []string
	}{
		{
			name: "succeeds when expected action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionCreate),
			},
			address:         "resource.foo",
			action:          tfjson.ActionCreate,
			shouldBePresent: true,
			expectError:     false,
		},
		{
			name: "fails when expected action is missing",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionUpdate),
			},
			address:         "resource.foo",
			action:          tfjson.ActionCreate,
			shouldBePresent: true,
			expectError:     true,
			errorContains:   []string{"no create is planned for resource.foo", "planned actions are: update"},
		},
		{
			name: "succeeds when action should be absent and is absent",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionUpdate),
			},
			address:         "resource.foo",
			action:          tfjson.ActionCreate,
			shouldBePresent: false,
			expectError:     false,
		},
		{
			name: "fails when action should be absent but is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionCreate),
			},
			address:         "resource.foo",
			action:          tfjson.ActionCreate,
			shouldBePresent: false,
			expectError:     true,
			errorContains:   []string{"create is planned for resource.foo", "planned actions are: create"},
		},
		{
			name: "sets error when resource not found",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("resource.foo", tfjson.ActionCreate),
			},
			address:         "resource.nonexistent",
			action:          tfjson.ActionCreate,
			shouldBePresent: true,
			expectError:     true,
			errorContains:   []string{"address resource.nonexistent not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checkActionPresence(req, resp, tt.address, tt.action, tt.shouldBePresent)

			if tt.expectError {
				assert.Error(t, resp.Error)
				for _, substr := range tt.errorContains {
					assert.Contains(t, resp.Error.Error(), substr)
				}
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}

func TestCheckResourceCreate(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		expectError     bool
		errorContains   string
	}{
		{
			name: "succeeds when create action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionCreate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "fails when create action is missing",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionUpdate),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "no create is planned",
		},
		{
			name:            "fails when resource not found",
			resourceChanges: []*tfjson.ResourceChange{},
			address:         "databricks_cluster.test",
			expectError:     true,
			errorContains:   "not found in resource changes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checker := CheckResourceCreate{Address: tt.address}
			checker.CheckPlan(context.Background(), req, resp)

			if tt.expectError {
				assert.Error(t, resp.Error)
				assert.Contains(t, resp.Error.Error(), tt.errorContains)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}

func TestCheckResourceUpdate(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		expectError     bool
		errorContains   string
	}{
		{
			name: "succeeds when update action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionUpdate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "fails when update action is missing",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionCreate),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "no update is planned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checker := CheckResourceUpdate{Address: tt.address}
			checker.CheckPlan(context.Background(), req, resp)

			if tt.expectError {
				assert.Error(t, resp.Error)
				assert.Contains(t, resp.Error.Error(), tt.errorContains)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}

func TestCheckResourceDelete(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		expectError     bool
		errorContains   string
	}{
		{
			name: "succeeds when delete action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionDelete),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "succeeds with replace actions",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionDelete, tfjson.ActionCreate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "fails when delete action is missing",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionUpdate),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "no delete is planned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checker := CheckResourceDelete{Address: tt.address}
			checker.CheckPlan(context.Background(), req, resp)

			if tt.expectError {
				assert.Error(t, resp.Error)
				assert.Contains(t, resp.Error.Error(), tt.errorContains)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}

func TestCheckResourceNoCreate(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		expectError     bool
		errorContains   string
	}{
		{
			name: "succeeds when create action is absent",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionUpdate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "succeeds when no actions",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test"),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "fails when create action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionCreate),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "create is planned for databricks_cluster.test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checker := CheckResourceNoCreate{Address: tt.address}
			checker.CheckPlan(context.Background(), req, resp)

			if tt.expectError {
				assert.Error(t, resp.Error)
				assert.Contains(t, resp.Error.Error(), tt.errorContains)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}

func TestCheckResourceNoDelete(t *testing.T) {
	tests := []struct {
		name            string
		resourceChanges []*tfjson.ResourceChange
		address         string
		expectError     bool
		errorContains   string
	}{
		{
			name: "succeeds when delete action is absent",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionUpdate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "succeeds when only create action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionCreate),
			},
			address:     "databricks_cluster.test",
			expectError: false,
		},
		{
			name: "fails when delete action is present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionDelete),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "delete is planned for databricks_cluster.test",
		},
		{
			name: "fails when replace actions are present",
			resourceChanges: []*tfjson.ResourceChange{
				createResourceChange("databricks_cluster.test", tfjson.ActionDelete, tfjson.ActionCreate),
			},
			address:       "databricks_cluster.test",
			expectError:   true,
			errorContains: "delete is planned for databricks_cluster.test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := plancheck.CheckPlanRequest{
				Plan: createMockPlan(tt.resourceChanges),
			}
			resp := &plancheck.CheckPlanResponse{}

			checker := CheckResourceNoDelete{Address: tt.address}
			checker.CheckPlan(context.Background(), req, resp)

			if tt.expectError {
				assert.Error(t, resp.Error)
				assert.Contains(t, resp.Error.Error(), tt.errorContains)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}
