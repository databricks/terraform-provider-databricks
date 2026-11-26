// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledgeassistants

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just KnowledgeAssistants API methods
type knowledgeAssistantsImpl struct {
	client *client.DatabricksClient
}

func (a *knowledgeAssistantsImpl) CreateExample(ctx context.Context, request CreateExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v/examples", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Example, &example)
	return &example, err
}

func (a *knowledgeAssistantsImpl) CreateKnowledgeAssistant(ctx context.Context, request CreateKnowledgeAssistantRequest) (*KnowledgeAssistant, error) {
	var knowledgeAssistant KnowledgeAssistant
	path := "/api/2.1/knowledge-assistants"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.KnowledgeAssistant, &knowledgeAssistant)
	return &knowledgeAssistant, err
}

func (a *knowledgeAssistantsImpl) CreateKnowledgeSource(ctx context.Context, request CreateKnowledgeSourceRequest) (*KnowledgeSource, error) {
	var knowledgeSource KnowledgeSource
	path := fmt.Sprintf("/api/2.1/%v/knowledge-sources", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.KnowledgeSource, &knowledgeSource)
	return &knowledgeSource, err
}

func (a *knowledgeAssistantsImpl) DeleteExample(ctx context.Context, request DeleteExampleRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
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

func (a *knowledgeAssistantsImpl) DeleteKnowledgeAssistant(ctx context.Context, request DeleteKnowledgeAssistantRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
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

func (a *knowledgeAssistantsImpl) DeleteKnowledgeSource(ctx context.Context, request DeleteKnowledgeSourceRequest) error {
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
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

func (a *knowledgeAssistantsImpl) GetExample(ctx context.Context, request GetExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &example)
	return &example, err
}

func (a *knowledgeAssistantsImpl) GetKnowledgeAssistant(ctx context.Context, request GetKnowledgeAssistantRequest) (*KnowledgeAssistant, error) {
	var knowledgeAssistant KnowledgeAssistant
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &knowledgeAssistant)
	return &knowledgeAssistant, err
}

func (a *knowledgeAssistantsImpl) GetKnowledgeSource(ctx context.Context, request GetKnowledgeSourceRequest) (*KnowledgeSource, error) {
	var knowledgeSource KnowledgeSource
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &knowledgeSource)
	return &knowledgeSource, err
}

func (a *knowledgeAssistantsImpl) GetPermissionLevels(ctx context.Context, request GetKnowledgeAssistantPermissionLevelsRequest) (*GetKnowledgeAssistantPermissionLevelsResponse, error) {
	var getKnowledgeAssistantPermissionLevelsResponse GetKnowledgeAssistantPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/knowledge-assistants/%v/permissionLevels", request.KnowledgeAssistantId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getKnowledgeAssistantPermissionLevelsResponse)
	return &getKnowledgeAssistantPermissionLevelsResponse, err
}

func (a *knowledgeAssistantsImpl) GetPermissions(ctx context.Context, request GetKnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error) {
	var knowledgeAssistantPermissions KnowledgeAssistantPermissions
	path := fmt.Sprintf("/api/2.0/permissions/knowledge-assistants/%v", request.KnowledgeAssistantId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &knowledgeAssistantPermissions)
	return &knowledgeAssistantPermissions, err
}

// Lists examples under a Knowledge Assistant.
func (a *knowledgeAssistantsImpl) ListExamples(ctx context.Context, request ListExamplesRequest) listing.Iterator[Example] {

	getNextPage := func(ctx context.Context, req ListExamplesRequest) (*ListExamplesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExamples(ctx, req)
	}
	getItems := func(resp *ListExamplesResponse) []Example {
		return resp.Examples
	}
	getNextReq := func(resp *ListExamplesResponse) *ListExamplesRequest {
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

// Lists examples under a Knowledge Assistant.
func (a *knowledgeAssistantsImpl) ListExamplesAll(ctx context.Context, request ListExamplesRequest) ([]Example, error) {
	iterator := a.ListExamples(ctx, request)
	return listing.ToSlice[Example](ctx, iterator)
}

func (a *knowledgeAssistantsImpl) internalListExamples(ctx context.Context, request ListExamplesRequest) (*ListExamplesResponse, error) {
	var listExamplesResponse ListExamplesResponse
	path := fmt.Sprintf("/api/2.1/%v/examples", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExamplesResponse)
	return &listExamplesResponse, err
}

// List Knowledge Assistants
func (a *knowledgeAssistantsImpl) ListKnowledgeAssistants(ctx context.Context, request ListKnowledgeAssistantsRequest) listing.Iterator[KnowledgeAssistant] {

	getNextPage := func(ctx context.Context, req ListKnowledgeAssistantsRequest) (*ListKnowledgeAssistantsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListKnowledgeAssistants(ctx, req)
	}
	getItems := func(resp *ListKnowledgeAssistantsResponse) []KnowledgeAssistant {
		return resp.KnowledgeAssistants
	}
	getNextReq := func(resp *ListKnowledgeAssistantsResponse) *ListKnowledgeAssistantsRequest {
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

// List Knowledge Assistants
func (a *knowledgeAssistantsImpl) ListKnowledgeAssistantsAll(ctx context.Context, request ListKnowledgeAssistantsRequest) ([]KnowledgeAssistant, error) {
	iterator := a.ListKnowledgeAssistants(ctx, request)
	return listing.ToSlice[KnowledgeAssistant](ctx, iterator)
}

func (a *knowledgeAssistantsImpl) internalListKnowledgeAssistants(ctx context.Context, request ListKnowledgeAssistantsRequest) (*ListKnowledgeAssistantsResponse, error) {
	var listKnowledgeAssistantsResponse ListKnowledgeAssistantsResponse
	path := "/api/2.1/knowledge-assistants"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listKnowledgeAssistantsResponse)
	return &listKnowledgeAssistantsResponse, err
}

// Lists Knowledge Sources under a Knowledge Assistant.
func (a *knowledgeAssistantsImpl) ListKnowledgeSources(ctx context.Context, request ListKnowledgeSourcesRequest) listing.Iterator[KnowledgeSource] {

	getNextPage := func(ctx context.Context, req ListKnowledgeSourcesRequest) (*ListKnowledgeSourcesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListKnowledgeSources(ctx, req)
	}
	getItems := func(resp *ListKnowledgeSourcesResponse) []KnowledgeSource {
		return resp.KnowledgeSources
	}
	getNextReq := func(resp *ListKnowledgeSourcesResponse) *ListKnowledgeSourcesRequest {
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

// Lists Knowledge Sources under a Knowledge Assistant.
func (a *knowledgeAssistantsImpl) ListKnowledgeSourcesAll(ctx context.Context, request ListKnowledgeSourcesRequest) ([]KnowledgeSource, error) {
	iterator := a.ListKnowledgeSources(ctx, request)
	return listing.ToSlice[KnowledgeSource](ctx, iterator)
}

func (a *knowledgeAssistantsImpl) internalListKnowledgeSources(ctx context.Context, request ListKnowledgeSourcesRequest) (*ListKnowledgeSourcesResponse, error) {
	var listKnowledgeSourcesResponse ListKnowledgeSourcesResponse
	path := fmt.Sprintf("/api/2.1/%v/knowledge-sources", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listKnowledgeSourcesResponse)
	return &listKnowledgeSourcesResponse, err
}

func (a *knowledgeAssistantsImpl) SetPermissions(ctx context.Context, request KnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error) {
	var knowledgeAssistantPermissions KnowledgeAssistantPermissions
	path := fmt.Sprintf("/api/2.0/permissions/knowledge-assistants/%v", request.KnowledgeAssistantId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &knowledgeAssistantPermissions)
	return &knowledgeAssistantPermissions, err
}

func (a *knowledgeAssistantsImpl) SyncKnowledgeSources(ctx context.Context, request SyncKnowledgeSourcesRequest) error {
	path := fmt.Sprintf("/api/2.1/%v/knowledge-sources:sync", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, nil)
	return err
}

func (a *knowledgeAssistantsImpl) UpdateExample(ctx context.Context, request UpdateExampleRequest) (*Example, error) {
	var example Example
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Example, &example)
	return &example, err
}

func (a *knowledgeAssistantsImpl) UpdateKnowledgeAssistant(ctx context.Context, request UpdateKnowledgeAssistantRequest) (*KnowledgeAssistant, error) {
	var knowledgeAssistant KnowledgeAssistant
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.KnowledgeAssistant, &knowledgeAssistant)
	return &knowledgeAssistant, err
}

func (a *knowledgeAssistantsImpl) UpdateKnowledgeSource(ctx context.Context, request UpdateKnowledgeSourceRequest) (*KnowledgeSource, error) {
	var knowledgeSource KnowledgeSource
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.KnowledgeSource, &knowledgeSource)
	return &knowledgeSource, err
}

func (a *knowledgeAssistantsImpl) UpdatePermissions(ctx context.Context, request KnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error) {
	var knowledgeAssistantPermissions KnowledgeAssistantPermissions
	path := fmt.Sprintf("/api/2.0/permissions/knowledge-assistants/%v", request.KnowledgeAssistantId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &knowledgeAssistantPermissions)
	return &knowledgeAssistantPermissions, err
}
