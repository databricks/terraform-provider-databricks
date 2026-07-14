package exporter

import (
	"log"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
)

func listBudgetPolicies(ic *importContext) error {
	if ic.accountClient == nil {
		return nil
	}
	policies, err := ic.accountClient.BudgetPolicy.ListAll(ic.Context, billing.ListBudgetPoliciesRequest{})
	if err != nil {
		return err
	}
	for _, policy := range policies {
		if policy.PolicyId == "" {
			continue
		}
		if !ic.MatchesName(policy.PolicyName) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_budget_policy",
			ID:       policy.PolicyId,
		})
	}
	return nil
}

func importBudgetPolicy(ic *importContext, r *resource) error {
	// Get binding_workspace_ids directly from DataWrapper
	if r.DataWrapper == nil {
		log.Printf("[WARN] DataWrapper is nil for budget policy %s", r.ID)
		return nil
	}

	accountID := ic.Client.Config.AccountID
	// Emit access control rule set for the budget policy
	ic.Emit(&resource{
		Resource: "databricks_access_control_rule_set",
		ID:       "accounts/" + accountID + "/budgetPolicies/" + r.ID + "/ruleSets/default",
	})

	bindingWorkspaceIdsRaw := r.DataWrapper.Get("binding_workspace_ids")
	if bindingWorkspaceIdsRaw != nil {
		// Convert to slice of int64
		var bindingWorkspaceIds []int64
		if workspaceIdsList, ok := bindingWorkspaceIdsRaw.([]int64); ok {
			bindingWorkspaceIds = workspaceIdsList
		}
		// Emit workspace resources for each binding_workspace_id
		if !ic.Client.Config.IsAzure() {
			for _, workspaceId := range bindingWorkspaceIds {
				ic.Emit(&resource{
					Resource: "databricks_mws_workspaces",
					ID:       accountID + "/" + strconv.FormatInt(workspaceId, 10),
				})
			}
		}
	}

	return nil
}

func listBudgets(ic *importContext) error {
	updatedSinceMs := ic.getUpdatedSinceMs()
	budgets, err := ic.accountClient.Budgets.ListAll(ic.Context, billing.ListBudgetConfigurationsRequest{})
	if err != nil {
		return err
	}
	for _, budget := range budgets {
		if ic.incremental && budget.CreateTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping budget '%s' that was updated at %d (last active=%d)",
				budget.DisplayName, budget.UpdateTime, updatedSinceMs)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_budget",
			ID:       ic.accountClient.Config.AccountID + "|" + budget.BudgetConfigurationId,
			Name:     budget.DisplayName,
		})
	}
	return nil
}

func importBudget(ic *importContext, r *resource) error {
	var budget billing.BudgetConfiguration
	s := ic.Resources["databricks_budget"].Schema
	common.DataToStructPointer(r.Data, s, &budget)
	if budget.Filter != nil && budget.Filter.WorkspaceId != nil && !ic.accountClient.Config.IsAzure() {
		for _, workspaceId := range budget.Filter.WorkspaceId.Values {
			ic.Emit(&resource{
				Resource: "databricks_mws_workspaces",
				ID:       ic.accountClient.Config.AccountID + "/" + strconv.FormatInt(workspaceId, 10),
			})
		}
	}
	for _, alert := range budget.AlertConfigurations {
		for _, action := range alert.ActionConfigurations {
			if action.ActionType == billing.ActionConfigurationTypeEmailNotification {
				ic.Emit(&resource{
					Resource:  "databricks_user",
					Attribute: "user_name",
					Value:     action.Target,
				})
			}
		}
	}
	return nil
}
