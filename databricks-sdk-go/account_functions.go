package databricks

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

// This is a workaround to the fact that legacy account hosts don't support workspace operations yet.
// If we detect a dns zone for which we cannot construct a workspace host,
// we assume we are in a unified host and return the account host itself.
func workspaceHost(env environment.DatabricksEnvironment, accountHost string, deploymentName string) string {
	// New hosts like SPOG may have an Environment without a dns zone.
	if env.DnsZone == "" {
		return accountHost
	}
	// Legacy hosts have very specific patterns for the account host.
	// In particular, they end with the dns zone in the Environment.
	normalized := accountHost
	if normalized != "" && !strings.Contains(normalized, "://") {
		normalized = "https://" + normalized
	}
	if u, err := url.Parse(normalized); err == nil && u.Hostname() != "" && strings.HasSuffix(u.Hostname(), env.DnsZone) {
		return env.DeploymentURL(deploymentName)
	}
	// New hosts like SPOG hosts use the same host for the account and workspace.
	return accountHost
}

// GetWorkspaceClient returns a WorkspaceClient for the given workspace. The
// workspace can be fetched by calling w.Workspaces.Get() or w.Workspaces.List().
//
// The config used for the workspace is identical to that used for the account,
// except that the host is set to the workspace host, and the account ID is
// not set.
//
// Deprecated: Use NewWorkspaceClient directly instead.
//
// Example:
//
//	a, err := databricks.NewAccountClient()
//	if err != nil {
//		panic(err)
//	}
//	ctx := context.Background()
//	workspaces, err := a.Workspaces.List(ctx)
//	if err != nil {
//		panic(err)
//	}
//	w, err := a.GetWorkspaceClient(workspaces[0])
//	if err != nil {
//		panic(err)
//	}
//	me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
func (c *AccountClient) GetWorkspaceClient(ws provisioning.Workspace) (*WorkspaceClient, error) {
	host := workspaceHost(c.Config.Environment(), c.Config.Host, ws.DeploymentName)
	cfg, err := c.Config.NewWithWorkspaceHost(host)
	if err != nil {
		return nil, err
	}
	cfg.AzureResourceID = ws.AzureResourceId()
	cfg.WorkspaceID = fmt.Sprintf("%d", ws.WorkspaceId)
	w, err := NewWorkspaceClient((*Config)(cfg))
	if err != nil {
		return nil, err
	}
	return w, nil
}
