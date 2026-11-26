package provisioning

import "fmt"

const resourceIdFormat = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Databricks/workspaces/%s"

// Return the AzureResourceID for the workspace, if it is an Azure workspace.
func (w Workspace) AzureResourceId() string {
	if w.AzureWorkspaceInfo == nil {
		return ""
	}
	return fmt.Sprintf(resourceIdFormat, w.AzureWorkspaceInfo.SubscriptionId, w.AzureWorkspaceInfo.ResourceGroup, w.WorkspaceName)
}
