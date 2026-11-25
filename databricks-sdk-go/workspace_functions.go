package databricks

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// CurrentWorkspaceID returns the workspace ID of the workspace that this client is
// connected to.
//
// On unified (SPOG) hosts, Config.WorkspaceID is sent in the
// X-Databricks-Workspace-Id request header to route the call to the correct
// workspace; on unified hosts with no WorkspaceID set, the request has no
// routing information and will fail. The ID is read back from the
// X-Databricks-Org-Id response header on /api/2.0/preview/scim/v2/Me — the
// server still emits the legacy response header name during the migration.
func (w *WorkspaceClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	var workspaceIdStr string
	opts := []httpclient.DoOption{
		httpclient.WithResponseHeader("X-Databricks-Org-Id", &workspaceIdStr),
		httpclient.WithRequestData(map[string]string{"excludedAttributes": "entitlements"}),
	}
	if w.Config != nil && w.Config.WorkspaceID != "" {
		opts = append(opts, httpclient.WithRequestHeader("X-Databricks-Workspace-Id", w.Config.WorkspaceID))
	}
	err := w.apiClient.Do(ctx, "GET", "/api/2.0/preview/scim/v2/Me", opts...)
	if err != nil {
		return 0, err
	}
	workspaceId, err := strconv.ParseInt(workspaceIdStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return workspaceId, nil
}
