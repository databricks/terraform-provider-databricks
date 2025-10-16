package common

import (
	"context"
	"fmt"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// CheckResourceCreate implements the PlanCheck interface defined in terraform-plugin-testing/plancheck
// This can be added as part of ConfigPlanChecks in tests to check whether the resource
// is planned to becreated.
// Example:
//
//	acceptance.Step{
//		ConfigPlanChecks: resource.ConfigPlanChecks{
//			PreApply: []plancheck.PlanCheck{
//				common.CheckResourceCreate{Address: "databricks_resource.name"},
//			},
//		},
//	},
type CheckResourceCreate struct {
	Address string
}

func (c CheckResourceCreate) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	checkActionPresence(req, resp, c.Address, tfjson.ActionCreate, true)
}

// CheckResourceUpdate implements the PlanCheck interface defined in terraform-plugin-testing/plancheck
// This can be added as part of ConfigPlanChecks in tests to check whether the resource
// is planned to be updated.
// Example:
//
//	acceptance.Step{
//		ConfigPlanChecks: resource.ConfigPlanChecks{
//			PreApply: []plancheck.PlanCheck{
//				common.CheckResourceUpdate{Address: "databricks_resource.name"},
//			},
//		},
//	},
type CheckResourceUpdate struct {
	Address string
}

func (c CheckResourceUpdate) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	checkActionPresence(req, resp, c.Address, tfjson.ActionUpdate, true)
}

// CheckResourceDelete implements the PlanCheck interface defined in terraform-plugin-testing/plancheck
// This can be added as part of ConfigPlanChecks in tests to check whether the resource
// is planned to be deleted.
// Example:
//
//	acceptance.Step{
//		ConfigPlanChecks: resource.ConfigPlanChecks{
//			PreApply: []plancheck.PlanCheck{
//				common.CheckResourceDelete{Address: "databricks_resource.name"},
//			},
//		},
//	},
type CheckResourceDelete struct {
	Address string
}

func (c CheckResourceDelete) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	checkActionPresence(req, resp, c.Address, tfjson.ActionDelete, true)
}

// CheckResourceNoCreate implements the PlanCheck interface defined in terraform-plugin-testing/plancheck
// This can be added as part of ConfigPlanChecks in tests to check whether the resource
// is not planned to be created.
// Example:
//
//	acceptance.Step{
//		ConfigPlanChecks: resource.ConfigPlanChecks{
//			PreApply: []plancheck.PlanCheck{
//				common.CheckResourceNoCreate{Address: "databricks_resource.name"},
//			},
//		},
//	},
type CheckResourceNoCreate struct {
	Address string
}

func (c CheckResourceNoCreate) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	checkActionPresence(req, resp, c.Address, tfjson.ActionCreate, false)
}

// CheckResourceNoDelete implements the PlanCheck interface defined in terraform-plugin-testing/plancheck
// This can be added as part of ConfigPlanChecks in tests to check whether the resource
// is not planned to be deleted.
// Example:
//
//	acceptance.Step{
//		ConfigPlanChecks: resource.ConfigPlanChecks{
//			PreApply: []plancheck.PlanCheck{
//				common.CheckResourceNoDelete{Address: "databricks_resource.name"},
//			},
//		},
//	},
type CheckResourceNoDelete struct {
	Address string
}

func (c CheckResourceNoDelete) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	checkActionPresence(req, resp, c.Address, tfjson.ActionDelete, false)
}

// findResourceChange finds a resource change by address in the plan
func findResourceChange(req plancheck.CheckPlanRequest, address string) (*tfjson.ResourceChange, error) {
	for _, resourceChange := range req.Plan.ResourceChanges {
		if resourceChange.Address == address {
			return resourceChange, nil
		}
	}

	addressesWithPlannedChanges := make([]string, 0, len(req.Plan.ResourceChanges))
	for _, change := range req.Plan.ResourceChanges {
		addressesWithPlannedChanges = append(addressesWithPlannedChanges, change.Address)
	}
	return nil, fmt.Errorf("address %s not found in resource changes; only planned changes for addresses %s",
		address, strings.Join(addressesWithPlannedChanges, ", "))
}

// getPlannedActions returns a string slice of all planned actions
func getPlannedActions(change *tfjson.ResourceChange) []string {
	plannedActions := make([]string, 0, len(change.Change.Actions))
	for _, action := range change.Change.Actions {
		plannedActions = append(plannedActions, string(action))
	}
	return plannedActions
}

// hasAction checks if a specific action is in the planned actions
func hasAction(change *tfjson.ResourceChange, targetAction tfjson.Action) bool {
	for _, action := range change.Change.Actions {
		if action == targetAction {
			return true
		}
	}
	return false
}

// checkActionPresence checks if an action is present (or absent) in the plan
func checkActionPresence(req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse,
	address string, action tfjson.Action, shouldBePresent bool) {

	change, err := findResourceChange(req, address)
	if err != nil {
		resp.Error = err
		return
	}

	actionPresent := hasAction(change, action)

	if shouldBePresent && !actionPresent {
		resp.Error = fmt.Errorf("no %s is planned for %s; planned actions are: %s",
			action, address, strings.Join(getPlannedActions(change), ", "))
	} else if !shouldBePresent && actionPresent {
		resp.Error = fmt.Errorf("%s is planned for %s; planned actions are: %s",
			action, address, strings.Join(getPlannedActions(change), ", "))
	}
}
