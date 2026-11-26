// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genieagents

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just GenieAgents API methods
type genieAgentsImpl struct {
	client *client.DatabricksClient
}

func (a *genieAgentsImpl) CreateGenieAgent(ctx context.Context, request CreateGenieAgentRequest) (*GenieAgent, error) {
	var genieAgent GenieAgent
	path := "/api/2.0/genie/agents"
	queryParams := make(map[string]any)

	if request.GenieAgentId != "" || slices.Contains(request.ForceSendFields, "GenieAgentId") {
		queryParams["genie_agent_id"] = request.GenieAgentId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.GenieAgent, &genieAgent)
	return &genieAgent, err
}

func (a *genieAgentsImpl) DeleteGenieAgent(ctx context.Context, request DeleteGenieAgentRequest) error {
	path := fmt.Sprintf("/api/2.0/genie/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *genieAgentsImpl) GetGenieAgent(ctx context.Context, request GetGenieAgentRequest) (*GenieAgent, error) {
	var genieAgent GenieAgent
	path := fmt.Sprintf("/api/2.0/genie/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieAgent)
	return &genieAgent, err
}

// Lists the Genie Agents in the workspace. Genie Agents created from an
// integration (for example, from a dashboard) are not listed.
func (a *genieAgentsImpl) ListGenieAgents(ctx context.Context, request ListGenieAgentsRequest) listing.Iterator[GenieAgent] {

	getNextPage := func(ctx context.Context, req ListGenieAgentsRequest) (*ListGenieAgentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListGenieAgents(ctx, req)
	}
	getItems := func(resp *ListGenieAgentsResponse) []GenieAgent {
		return resp.GenieAgents
	}
	getNextReq := func(resp *ListGenieAgentsResponse) *ListGenieAgentsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the Genie Agents in the workspace. Genie Agents created from an
// integration (for example, from a dashboard) are not listed.
func (a *genieAgentsImpl) ListGenieAgentsAll(ctx context.Context, request ListGenieAgentsRequest) ([]GenieAgent, error) {
	iterator := a.ListGenieAgents(ctx, request)
	return listing.ToSlice[GenieAgent](ctx, iterator)
}

func (a *genieAgentsImpl) internalListGenieAgents(ctx context.Context, request ListGenieAgentsRequest) (*ListGenieAgentsResponse, error) {
	var listGenieAgentsResponse ListGenieAgentsResponse
	path := "/api/2.0/genie/agents"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGenieAgentsResponse)
	return &listGenieAgentsResponse, err
}

func (a *genieAgentsImpl) UpdateGenieAgent(ctx context.Context, request UpdateGenieAgentRequest) (*GenieAgent, error) {
	var genieAgent GenieAgent
	path := fmt.Sprintf("/api/2.0/genie/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.GenieAgent, &genieAgent)
	return &genieAgent, err
}
