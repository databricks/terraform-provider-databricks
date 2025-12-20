package exporter

import (
	"log"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
)

func listAccountNetworkPolicies(ic *importContext) error {
	policies, err := ic.accountClient.NetworkPolicies.ListNetworkPoliciesRpcAll(ic.Context, settings.ListNetworkPoliciesRequest{})
	if err != nil {
		return err
	}
	for _, policy := range policies {
		if policy.NetworkPolicyId == "default-policy" {
			continue
		}
		if !ic.MatchesName(policy.NetworkPolicyId) {
			log.Printf("[INFO] Skipping account_network_policy %s because it doesn't match %s", policy.NetworkPolicyId, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_account_network_policy",
			ID:       policy.NetworkPolicyId,
		})
	}
	log.Printf("[INFO] Listed %d account network policies", len(policies))
	return nil
}

func listWorkspaceNetworkOptions(ic *importContext) error {
	workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
	if err != nil {
		return err
	}

	count := 0
	for _, workspace := range workspaces {
		if !ic.MatchesName(workspace.WorkspaceName) {
			log.Printf("[INFO] Skipping workspace_network_option for workspace %s because it doesn't match %s",
				workspace.WorkspaceName, ic.match)
			continue
		}

		if workspace.WorkspaceStatus != "RUNNING" {
			log.Printf("[DEBUG] Skipping workspace_network_option for workspace %s that is not running (status: %s)",
				workspace.WorkspaceName, workspace.WorkspaceStatus)
			continue
		}

		ic.Emit(&resource{
			Resource: "databricks_workspace_network_option",
			ID:       strconv.FormatInt(workspace.WorkspaceId, 10),
		})
		count++
	}
	log.Printf("[INFO] Listed %d workspace network options", count)
	return nil
}

func importWorkspaceNetworkOption(ic *importContext, r *resource) error {
	workspaceId, err := strconv.ParseInt(r.ID, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Failed to parse workspace_id for workspace_network_option: %s", err.Error())
		return err
	}

	// Get the workspace network option to emit the associated network policy
	option, err := ic.accountClient.WorkspaceNetworkConfiguration.GetWorkspaceNetworkOptionRpc(ic.Context,
		settings.GetWorkspaceNetworkOptionRequest{
			WorkspaceId: workspaceId,
		})
	if err != nil {
		log.Printf("[WARN] Failed to get workspace network option for workspace %d: %s", workspaceId, err.Error())
		return nil
	}

	if option.NetworkPolicyId != "" && option.NetworkPolicyId != "default-policy" {
		ic.Emit(&resource{
			Resource: "databricks_account_network_policy",
			ID:       option.NetworkPolicyId,
		})
	}

	return nil
}

func listMwsWorkspaceNccBindings(ic *importContext) error {
	workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
	if err != nil {
		return err
	}
	for _, workspace := range workspaces {
		if workspace.NetworkConnectivityConfigId != "" {
			ic.emitNccBindingAndNcc(workspace.WorkspaceId, workspace.NetworkConnectivityConfigId)
			if !ic.accountClient.Config.IsAzure() {
				wsIdString := strconv.FormatInt(workspace.WorkspaceId, 10)
				ic.Emit(&resource{
					Resource: "databricks_mws_workspaces",
					ID:       ic.accountClient.Config.AccountID + "/" + wsIdString,
					Name:     workspace.WorkspaceName + "_" + wsIdString,
				})
			}
		}
	}
	return nil
}

func listMwsNetworkConnectivityConfigs(ic *importContext) error {
	updatedSinceMs := ic.getUpdatedSinceMs()
	it := ic.accountClient.NetworkConnectivity.ListNetworkConnectivityConfigurations(ic.Context,
		settings.ListNetworkConnectivityConfigurationsRequest{})
	for it.HasNext(ic.Context) {
		nc, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if !ic.MatchesName(nc.Name) {
			log.Printf("[INFO] Skipping mws_network_connectivity_config %s because it doesn't match %s", nc.Name, ic.match)
			continue
		}
		if ic.incremental && nc.UpdatedTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping mws_network_connectivity_config '%s' that was modified at %d (last active=%d)",
				nc.Name, nc.UpdatedTime, updatedSinceMs)
			continue
		}
		// TODO: technically we can create data directly from the API response
		ic.Emit(&resource{
			Resource: "databricks_mws_network_connectivity_config",
			ID:       nc.AccountId + "/" + nc.NetworkConnectivityConfigId,
		})
		if nc.EgressConfig.TargetRules != nil {
			for _, rule := range nc.EgressConfig.TargetRules.AzurePrivateEndpointRules {
				// TODO: technically we can create data directly from the API response
				resourceId := strings.ReplaceAll(rule.ResourceId, "/subscriptions/", "")
				resourceId = strings.ReplaceAll(resourceId, "/resourceGroups/", "_")
				resourceId = strings.ReplaceAll(resourceId, "/providers/Microsoft", "_")
				ic.Emit(&resource{
					Resource: "databricks_mws_ncc_private_endpoint_rule",
					ID:       nc.NetworkConnectivityConfigId + "/" + rule.RuleId,
					Name:     nc.Name + "_" + resourceId + "_" + rule.GroupId,
				})
			}
			for _, rule := range nc.EgressConfig.TargetRules.AwsPrivateEndpointRules {
				ic.Emit(&resource{
					Resource: "databricks_mws_ncc_private_endpoint_rule",
					ID:       nc.NetworkConnectivityConfigId + "/" + rule.RuleId,
					Name:     nc.Name + "_" + rule.EndpointService,
				})
			}
		}
	}
	return nil
}
