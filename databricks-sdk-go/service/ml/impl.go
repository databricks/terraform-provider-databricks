// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Experiments API methods
type experimentsImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsImpl) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0/mlflow/experiments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

func (a *experimentsImpl) CreateIssue(ctx context.Context, request CreateIssueRequest) (*Issue, error) {
	var issue Issue
	path := fmt.Sprintf("/api/2.0/mlflow/%v/issues", request.Parent)
	queryParams := make(map[string]any)

	if request.IssueId != "" || slices.Contains(request.ForceSendFields, "IssueId") {
		queryParams["issue_id"] = request.IssueId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Issue, &issue)
	return &issue, err
}

func (a *experimentsImpl) CreateLoggedModel(ctx context.Context, request CreateLoggedModelRequest) (*CreateLoggedModelResponse, error) {
	var createLoggedModelResponse CreateLoggedModelResponse
	path := "/api/2.0/mlflow/logged-models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createLoggedModelResponse)
	return &createLoggedModelResponse, err
}

func (a *experimentsImpl) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0/mlflow/runs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createRunResponse)
	return &createRunResponse, err
}

func (a *experimentsImpl) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {
	path := "/api/2.0/mlflow/experiments/delete"
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

func (a *experimentsImpl) DeleteLoggedModel(ctx context.Context, request DeleteLoggedModelRequest) error {
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
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

func (a *experimentsImpl) DeleteLoggedModelTag(ctx context.Context, request DeleteLoggedModelTagRequest) error {
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags/%v", request.ModelId, request.TagKey)
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

func (a *experimentsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.0/mlflow/runs/delete"
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

func (a *experimentsImpl) DeleteRuns(ctx context.Context, request DeleteRuns) (*DeleteRunsResponse, error) {
	var deleteRunsResponse DeleteRunsResponse
	path := "/api/2.0/mlflow/databricks/runs/delete-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteRunsResponse)
	return &deleteRunsResponse, err
}

func (a *experimentsImpl) DeleteTag(ctx context.Context, request DeleteTag) error {
	path := "/api/2.0/mlflow/runs/delete-tag"
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

func (a *experimentsImpl) FinalizeLoggedModel(ctx context.Context, request FinalizeLoggedModelRequest) (*FinalizeLoggedModelResponse, error) {
	var finalizeLoggedModelResponse FinalizeLoggedModelResponse
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &finalizeLoggedModelResponse)
	return &finalizeLoggedModelResponse, err
}

func (a *experimentsImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	var getExperimentByNameResponse GetExperimentByNameResponse
	path := "/api/2.0/mlflow/experiments/get-by-name"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentByNameResponse)
	return &getExperimentByNameResponse, err
}

func (a *experimentsImpl) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0/mlflow/experiments/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

// Gets a list of all values for the specified metric for a given run.
func (a *experimentsImpl) GetHistory(ctx context.Context, request GetHistoryRequest) listing.Iterator[Metric] {

	getNextPage := func(ctx context.Context, req GetHistoryRequest) (*GetMetricHistoryResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetHistory(ctx, req)
	}
	getItems := func(resp *GetMetricHistoryResponse) []Metric {
		return resp.Metrics
	}
	getNextReq := func(resp *GetMetricHistoryResponse) *GetHistoryRequest {
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

// Gets a list of all values for the specified metric for a given run.
func (a *experimentsImpl) GetHistoryAll(ctx context.Context, request GetHistoryRequest) ([]Metric, error) {
	iterator := a.GetHistory(ctx, request)
	return listing.ToSlice[Metric](ctx, iterator)
}

func (a *experimentsImpl) internalGetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0/mlflow/metrics/get-history"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

func (a *experimentsImpl) GetIssue(ctx context.Context, request GetIssueRequest) (*Issue, error) {
	var issue Issue
	path := fmt.Sprintf("/api/2.0/mlflow/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &issue)
	return &issue, err
}

func (a *experimentsImpl) GetLoggedModel(ctx context.Context, request GetLoggedModelRequest) (*GetLoggedModelResponse, error) {
	var getLoggedModelResponse GetLoggedModelResponse
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getLoggedModelResponse)
	return &getLoggedModelResponse, err
}

func (a *experimentsImpl) GetPermissionLevels(ctx context.Context, request GetExperimentPermissionLevelsRequest) (*GetExperimentPermissionLevelsResponse, error) {
	var getExperimentPermissionLevelsResponse GetExperimentPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v/permissionLevels", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentPermissionLevelsResponse)
	return &getExperimentPermissionLevelsResponse, err
}

func (a *experimentsImpl) GetPermissions(ctx context.Context, request GetExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0/mlflow/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRunResponse)
	return &getRunResponse, err
}

// List artifacts for a run. Takes an optional `artifact_path` prefix which if
// specified, the response contains only artifacts with the specified prefix. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsImpl) ListArtifacts(ctx context.Context, request ListArtifactsRequest) listing.Iterator[FileInfo] {

	getNextPage := func(ctx context.Context, req ListArtifactsRequest) (*ListArtifactsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListArtifacts(ctx, req)
	}
	getItems := func(resp *ListArtifactsResponse) []FileInfo {
		return resp.Files
	}
	getNextReq := func(resp *ListArtifactsResponse) *ListArtifactsRequest {
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

// List artifacts for a run. Takes an optional `artifact_path` prefix which if
// specified, the response contains only artifacts with the specified prefix. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsImpl) ListArtifactsAll(ctx context.Context, request ListArtifactsRequest) ([]FileInfo, error) {
	iterator := a.ListArtifacts(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}

func (a *experimentsImpl) internalListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0/mlflow/artifacts/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

// Gets a list of all experiments, ordered by creation time with the newest
// first.
func (a *experimentsImpl) ListExperiments(ctx context.Context, request ListExperimentsRequest) listing.Iterator[Experiment] {

	getNextPage := func(ctx context.Context, req ListExperimentsRequest) (*ListExperimentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExperiments(ctx, req)
	}
	getItems := func(resp *ListExperimentsResponse) []Experiment {
		return resp.Experiments
	}
	getNextReq := func(resp *ListExperimentsResponse) *ListExperimentsRequest {
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

// Gets a list of all experiments, ordered by creation time with the newest
// first.
func (a *experimentsImpl) ListExperimentsAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	iterator := a.ListExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0/mlflow/experiments/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

func (a *experimentsImpl) LogBatch(ctx context.Context, request LogBatch) error {
	path := "/api/2.0/mlflow/runs/log-batch"
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

func (a *experimentsImpl) LogInputs(ctx context.Context, request LogInputs) error {
	path := "/api/2.0/mlflow/runs/log-inputs"
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

func (a *experimentsImpl) LogLoggedModelParams(ctx context.Context, request LogLoggedModelParamsRequest) error {
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/params", request.ModelId)
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

func (a *experimentsImpl) LogMetric(ctx context.Context, request LogMetric) error {
	path := "/api/2.0/mlflow/runs/log-metric"
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

func (a *experimentsImpl) LogModel(ctx context.Context, request LogModel) error {
	path := "/api/2.0/mlflow/runs/log-model"
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

func (a *experimentsImpl) LogOutputs(ctx context.Context, request LogOutputsRequest) error {
	path := "/api/2.0/mlflow/runs/outputs"
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

func (a *experimentsImpl) LogParam(ctx context.Context, request LogParam) error {
	path := "/api/2.0/mlflow/runs/log-parameter"
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

func (a *experimentsImpl) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {
	path := "/api/2.0/mlflow/experiments/restore"
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

func (a *experimentsImpl) RestoreRun(ctx context.Context, request RestoreRun) error {
	path := "/api/2.0/mlflow/runs/restore"
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

func (a *experimentsImpl) RestoreRuns(ctx context.Context, request RestoreRuns) (*RestoreRunsResponse, error) {
	var restoreRunsResponse RestoreRunsResponse
	path := "/api/2.0/mlflow/databricks/runs/restore-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &restoreRunsResponse)
	return &restoreRunsResponse, err
}

// Searches for experiments that satisfy specified search criteria.
func (a *experimentsImpl) SearchExperiments(ctx context.Context, request SearchExperiments) listing.Iterator[Experiment] {

	getNextPage := func(ctx context.Context, req SearchExperiments) (*SearchExperimentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchExperiments(ctx, req)
	}
	getItems := func(resp *SearchExperimentsResponse) []Experiment {
		return resp.Experiments
	}
	getNextReq := func(resp *SearchExperimentsResponse) *SearchExperiments {
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

// Searches for experiments that satisfy specified search criteria.
func (a *experimentsImpl) SearchExperimentsAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	iterator := a.SearchExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalSearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0/mlflow/experiments/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

// Searches the issues in an experiment, filtered and sorted.
func (a *experimentsImpl) SearchIssues(ctx context.Context, request SearchIssuesRequest) listing.Iterator[Issue] {

	getNextPage := func(ctx context.Context, req SearchIssuesRequest) (*SearchIssuesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchIssues(ctx, req)
	}
	getItems := func(resp *SearchIssuesResponse) []Issue {
		return resp.Issues
	}
	getNextReq := func(resp *SearchIssuesResponse) *SearchIssuesRequest {
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

// Searches the issues in an experiment, filtered and sorted.
func (a *experimentsImpl) SearchIssuesAll(ctx context.Context, request SearchIssuesRequest) ([]Issue, error) {
	iterator := a.SearchIssues(ctx, request)
	return listing.ToSlice[Issue](ctx, iterator)
}

func (a *experimentsImpl) internalSearchIssues(ctx context.Context, request SearchIssuesRequest) (*SearchIssuesResponse, error) {
	var searchIssuesResponse SearchIssuesResponse
	path := fmt.Sprintf("/api/2.0/mlflow/%v/issues:search", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchIssuesResponse)
	return &searchIssuesResponse, err
}

func (a *experimentsImpl) SearchLoggedModels(ctx context.Context, request SearchLoggedModelsRequest) (*SearchLoggedModelsResponse, error) {
	var searchLoggedModelsResponse SearchLoggedModelsResponse
	path := "/api/2.0/mlflow/logged-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchLoggedModelsResponse)
	return &searchLoggedModelsResponse, err
}

// Searches for runs that satisfy expressions.
//
// Search expressions can use metric, param, and tag keys.
func (a *experimentsImpl) SearchRuns(ctx context.Context, request SearchRuns) listing.Iterator[Run] {

	getNextPage := func(ctx context.Context, req SearchRuns) (*SearchRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchRuns(ctx, req)
	}
	getItems := func(resp *SearchRunsResponse) []Run {
		return resp.Runs
	}
	getNextReq := func(resp *SearchRunsResponse) *SearchRuns {
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

// Searches for runs that satisfy expressions.
//
// Search expressions can use metric, param, and tag keys.
func (a *experimentsImpl) SearchRunsAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	iterator := a.SearchRuns(ctx, request)
	return listing.ToSlice[Run](ctx, iterator)
}

func (a *experimentsImpl) internalSearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0/mlflow/runs/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

func (a *experimentsImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
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

func (a *experimentsImpl) SetLoggedModelTags(ctx context.Context, request SetLoggedModelTagsRequest) error {
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *experimentsImpl) SetPermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) SetTag(ctx context.Context, request SetTag) error {
	path := "/api/2.0/mlflow/runs/set-tag"
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

func (a *experimentsImpl) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {
	path := "/api/2.0/mlflow/experiments/update"
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

func (a *experimentsImpl) UpdateIssue(ctx context.Context, request UpdateIssueRequest) (*Issue, error) {
	var issue Issue
	path := fmt.Sprintf("/api/2.0/mlflow/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Issue, &issue)
	return &issue, err
}

func (a *experimentsImpl) UpdatePermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0/mlflow/runs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateRunResponse)
	return &updateRunResponse, err
}

// unexported type that holds implementations of just FeatureEngineering API methods
type featureEngineeringImpl struct {
	client *client.DatabricksClient
}

func (a *featureEngineeringImpl) BackfillFeatures(ctx context.Context, request BackfillFeaturesRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/2.0/feature-engineering/features:backfill"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) BatchCreateMaterializedFeatures(ctx context.Context, request BatchCreateMaterializedFeaturesRequest) (*BatchCreateMaterializedFeaturesResponse, error) {
	var batchCreateMaterializedFeaturesResponse BatchCreateMaterializedFeaturesResponse
	path := "/api/2.0/feature-engineering/materialized-features:batchCreate"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &batchCreateMaterializedFeaturesResponse)
	return &batchCreateMaterializedFeaturesResponse, err
}

func (a *featureEngineeringImpl) CancelOperation(ctx context.Context, request CancelOperationRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/%v:cancel", request.Name)
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

func (a *featureEngineeringImpl) CreateFeature(ctx context.Context, request CreateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := "/api/2.0/feature-engineering/features"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Feature, &feature)
	return &feature, err
}

func (a *featureEngineeringImpl) CreateFeatureSpec(ctx context.Context, request CreateFeatureSpecRequest) (*FeatureSpec, error) {
	var featureSpec FeatureSpec
	path := "/api/2.0/feature-engineering/feature-specs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.FeatureSpec, &featureSpec)
	return &featureSpec, err
}

func (a *featureEngineeringImpl) CreateKafkaConfig(ctx context.Context, request CreateKafkaConfigRequest) (*KafkaConfig, error) {
	var kafkaConfig KafkaConfig
	path := "/api/2.0/feature-engineering/features/kafka-configs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.KafkaConfig, &kafkaConfig)
	return &kafkaConfig, err
}

func (a *featureEngineeringImpl) CreateMaterializationSet(ctx context.Context, request CreateMaterializationSetRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/2.0/feature-engineering/materialization-sets"
	queryParams := make(map[string]any)

	if request.RequestId != "" || slices.Contains(request.ForceSendFields, "RequestId") {
		queryParams["request_id"] = request.RequestId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.MaterializationSet, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) CreateMaterializedFeature(ctx context.Context, request CreateMaterializedFeatureRequest) (*MaterializedFeature, error) {
	var materializedFeature MaterializedFeature
	path := "/api/2.0/feature-engineering/materialized-features"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.MaterializedFeature, &materializedFeature)
	return &materializedFeature, err
}

func (a *featureEngineeringImpl) CreateStream(ctx context.Context, request CreateStreamRequest) (*Stream, error) {
	var stream Stream
	path := "/api/2.0/feature-engineering/streams"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Stream, &stream)
	return &stream, err
}

func (a *featureEngineeringImpl) DeleteFeature(ctx context.Context, request DeleteFeatureRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/%v", request.FullName)
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

func (a *featureEngineeringImpl) DeleteFeatureSpec(ctx context.Context, request DeleteFeatureSpecRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/feature-specs/%v", request.Name)
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

func (a *featureEngineeringImpl) DeleteKafkaConfig(ctx context.Context, request DeleteKafkaConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/kafka-configs/%v", request.Name)
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

func (a *featureEngineeringImpl) DeleteMaterializationSet(ctx context.Context, request DeleteMaterializationSetRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialization-sets/%v", request.MaterializationSetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) DeleteMaterializedFeature(ctx context.Context, request DeleteMaterializedFeatureRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialized-features/%v", request.MaterializedFeatureId)
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

func (a *featureEngineeringImpl) DeleteStream(ctx context.Context, request DeleteStreamRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-engineering/streams/%v", request.Name)
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

func (a *featureEngineeringImpl) GetFeature(ctx context.Context, request GetFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &feature)
	return &feature, err
}

func (a *featureEngineeringImpl) GetFeatureSpec(ctx context.Context, request GetFeatureSpecRequest) (*FeatureSpec, error) {
	var featureSpec FeatureSpec
	path := fmt.Sprintf("/api/2.0/feature-engineering/feature-specs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &featureSpec)
	return &featureSpec, err
}

func (a *featureEngineeringImpl) GetKafkaConfig(ctx context.Context, request GetKafkaConfigRequest) (*KafkaConfig, error) {
	var kafkaConfig KafkaConfig
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/kafka-configs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &kafkaConfig)
	return &kafkaConfig, err
}

func (a *featureEngineeringImpl) GetMaterializationSet(ctx context.Context, request GetMaterializationSetRequest) (*MaterializationSet, error) {
	var materializationSet MaterializationSet
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialization-sets/%v", request.MaterializationSetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &materializationSet)
	return &materializationSet, err
}

func (a *featureEngineeringImpl) GetMaterializedFeature(ctx context.Context, request GetMaterializedFeatureRequest) (*MaterializedFeature, error) {
	var materializedFeature MaterializedFeature
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialized-features/%v", request.MaterializedFeatureId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &materializedFeature)
	return &materializedFeature, err
}

func (a *featureEngineeringImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/feature-engineering/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) GetStream(ctx context.Context, request GetStreamRequest) (*Stream, error) {
	var stream Stream
	path := fmt.Sprintf("/api/2.0/feature-engineering/streams/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &stream)
	return &stream, err
}

// List Features.
func (a *featureEngineeringImpl) ListFeatures(ctx context.Context, request ListFeaturesRequest) listing.Iterator[Feature] {

	getNextPage := func(ctx context.Context, req ListFeaturesRequest) (*ListFeaturesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListFeatures(ctx, req)
	}
	getItems := func(resp *ListFeaturesResponse) []Feature {
		return resp.Features
	}
	getNextReq := func(resp *ListFeaturesResponse) *ListFeaturesRequest {
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

// List Features.
func (a *featureEngineeringImpl) ListFeaturesAll(ctx context.Context, request ListFeaturesRequest) ([]Feature, error) {
	iterator := a.ListFeatures(ctx, request)
	return listing.ToSlice[Feature](ctx, iterator)
}

func (a *featureEngineeringImpl) internalListFeatures(ctx context.Context, request ListFeaturesRequest) (*ListFeaturesResponse, error) {
	var listFeaturesResponse ListFeaturesResponse
	path := "/api/2.0/feature-engineering/features"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFeaturesResponse)
	return &listFeaturesResponse, err
}

// List Kafka configs. During PrPr, Kafka configs can be read and used when
// creating features under the entire metastore. Only the creator of the Kafka
// config can delete it.
func (a *featureEngineeringImpl) ListKafkaConfigs(ctx context.Context, request ListKafkaConfigsRequest) listing.Iterator[KafkaConfig] {

	getNextPage := func(ctx context.Context, req ListKafkaConfigsRequest) (*ListKafkaConfigsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListKafkaConfigs(ctx, req)
	}
	getItems := func(resp *ListKafkaConfigsResponse) []KafkaConfig {
		return resp.KafkaConfigs
	}
	getNextReq := func(resp *ListKafkaConfigsResponse) *ListKafkaConfigsRequest {
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

// List Kafka configs. During PrPr, Kafka configs can be read and used when
// creating features under the entire metastore. Only the creator of the Kafka
// config can delete it.
func (a *featureEngineeringImpl) ListKafkaConfigsAll(ctx context.Context, request ListKafkaConfigsRequest) ([]KafkaConfig, error) {
	iterator := a.ListKafkaConfigs(ctx, request)
	return listing.ToSlice[KafkaConfig](ctx, iterator)
}

func (a *featureEngineeringImpl) internalListKafkaConfigs(ctx context.Context, request ListKafkaConfigsRequest) (*ListKafkaConfigsResponse, error) {
	var listKafkaConfigsResponse ListKafkaConfigsResponse
	path := "/api/2.0/feature-engineering/features/kafka-configs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listKafkaConfigsResponse)
	return &listKafkaConfigsResponse, err
}

// List materialized features.
func (a *featureEngineeringImpl) ListMaterializedFeatures(ctx context.Context, request ListMaterializedFeaturesRequest) listing.Iterator[MaterializedFeature] {

	getNextPage := func(ctx context.Context, req ListMaterializedFeaturesRequest) (*ListMaterializedFeaturesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMaterializedFeatures(ctx, req)
	}
	getItems := func(resp *ListMaterializedFeaturesResponse) []MaterializedFeature {
		return resp.MaterializedFeatures
	}
	getNextReq := func(resp *ListMaterializedFeaturesResponse) *ListMaterializedFeaturesRequest {
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

// List materialized features.
func (a *featureEngineeringImpl) ListMaterializedFeaturesAll(ctx context.Context, request ListMaterializedFeaturesRequest) ([]MaterializedFeature, error) {
	iterator := a.ListMaterializedFeatures(ctx, request)
	return listing.ToSlice[MaterializedFeature](ctx, iterator)
}

func (a *featureEngineeringImpl) internalListMaterializedFeatures(ctx context.Context, request ListMaterializedFeaturesRequest) (*ListMaterializedFeaturesResponse, error) {
	var listMaterializedFeaturesResponse ListMaterializedFeaturesResponse
	path := "/api/2.0/feature-engineering/materialized-features"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMaterializedFeaturesResponse)
	return &listMaterializedFeaturesResponse, err
}

// List Streams under a given catalog.schema parent.
func (a *featureEngineeringImpl) ListStreams(ctx context.Context, request ListStreamsRequest) listing.Iterator[Stream] {

	getNextPage := func(ctx context.Context, req ListStreamsRequest) (*ListStreamsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListStreams(ctx, req)
	}
	getItems := func(resp *ListStreamsResponse) []Stream {
		return resp.Streams
	}
	getNextReq := func(resp *ListStreamsResponse) *ListStreamsRequest {
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

// List Streams under a given catalog.schema parent.
func (a *featureEngineeringImpl) ListStreamsAll(ctx context.Context, request ListStreamsRequest) ([]Stream, error) {
	iterator := a.ListStreams(ctx, request)
	return listing.ToSlice[Stream](ctx, iterator)
}

func (a *featureEngineeringImpl) internalListStreams(ctx context.Context, request ListStreamsRequest) (*ListStreamsResponse, error) {
	var listStreamsResponse ListStreamsResponse
	path := "/api/2.0/feature-engineering/streams"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStreamsResponse)
	return &listStreamsResponse, err
}

func (a *featureEngineeringImpl) PurgeFeatureEntities(ctx context.Context, request PurgeFeatureEntitiesRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/2.0/feature-engineering/features:purgeFeatureEntities"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) TriggerMaterializationSet(ctx context.Context, request TriggerMaterializationSetRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialization-sets/%v/trigger", request.MaterializationSetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) UpdateFeature(ctx context.Context, request UpdateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/%v", request.FullName)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Feature, &feature)
	return &feature, err
}

func (a *featureEngineeringImpl) UpdateKafkaConfig(ctx context.Context, request UpdateKafkaConfigRequest) (*KafkaConfig, error) {
	var kafkaConfig KafkaConfig
	path := fmt.Sprintf("/api/2.0/feature-engineering/features/kafka-configs/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.KafkaConfig, &kafkaConfig)
	return &kafkaConfig, err
}

func (a *featureEngineeringImpl) UpdateMaterializationSet(ctx context.Context, request UpdateMaterializationSetRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialization-sets/%v", request.MaterializationSetId)
	queryParams := make(map[string]any)

	if request.RequestId != "" || slices.Contains(request.ForceSendFields, "RequestId") {
		queryParams["request_id"] = request.RequestId
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.MaterializationSet, &operation)
	return &operation, err
}

func (a *featureEngineeringImpl) UpdateMaterializedFeature(ctx context.Context, request UpdateMaterializedFeatureRequest) (*MaterializedFeature, error) {
	var materializedFeature MaterializedFeature
	path := fmt.Sprintf("/api/2.0/feature-engineering/materialized-features/%v", request.MaterializedFeatureId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.MaterializedFeature, &materializedFeature)
	return &materializedFeature, err
}

func (a *featureEngineeringImpl) UpdateStream(ctx context.Context, request UpdateStreamRequest) (*Stream, error) {
	var stream Stream
	path := fmt.Sprintf("/api/2.0/feature-engineering/streams/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Stream, &stream)
	return &stream, err
}

// unexported type that holds implementations of just FeatureStore API methods
type featureStoreImpl struct {
	client *client.DatabricksClient
}

func (a *featureStoreImpl) CreateOnlineStore(ctx context.Context, request CreateOnlineStoreRequest) (*OnlineStore, error) {
	var onlineStore OnlineStore
	path := "/api/2.0/feature-store/online-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.OnlineStore, &onlineStore)
	return &onlineStore, err
}

func (a *featureStoreImpl) DeleteOnlineStore(ctx context.Context, request DeleteOnlineStoreRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
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

func (a *featureStoreImpl) DeleteOnlineTable(ctx context.Context, request DeleteOnlineTableRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-store/online-tables/%v", request.OnlineTableName)
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

func (a *featureStoreImpl) GetOnlineStore(ctx context.Context, request GetOnlineStoreRequest) (*OnlineStore, error) {
	var onlineStore OnlineStore
	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &onlineStore)
	return &onlineStore, err
}

// List Online Feature Stores.
func (a *featureStoreImpl) ListOnlineStores(ctx context.Context, request ListOnlineStoresRequest) listing.Iterator[OnlineStore] {

	getNextPage := func(ctx context.Context, req ListOnlineStoresRequest) (*ListOnlineStoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListOnlineStores(ctx, req)
	}
	getItems := func(resp *ListOnlineStoresResponse) []OnlineStore {
		return resp.OnlineStores
	}
	getNextReq := func(resp *ListOnlineStoresResponse) *ListOnlineStoresRequest {
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

// List Online Feature Stores.
func (a *featureStoreImpl) ListOnlineStoresAll(ctx context.Context, request ListOnlineStoresRequest) ([]OnlineStore, error) {
	iterator := a.ListOnlineStores(ctx, request)
	return listing.ToSlice[OnlineStore](ctx, iterator)
}

func (a *featureStoreImpl) internalListOnlineStores(ctx context.Context, request ListOnlineStoresRequest) (*ListOnlineStoresResponse, error) {
	var listOnlineStoresResponse ListOnlineStoresResponse
	path := "/api/2.0/feature-store/online-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listOnlineStoresResponse)
	return &listOnlineStoresResponse, err
}

func (a *featureStoreImpl) PublishTable(ctx context.Context, request PublishTableRequest) (*PublishTableResponse, error) {
	var publishTableResponse PublishTableResponse
	path := fmt.Sprintf("/api/2.0/feature-store/tables/%v/publish", request.SourceTableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &publishTableResponse)
	return &publishTableResponse, err
}

func (a *featureStoreImpl) UpdateOnlineStore(ctx context.Context, request UpdateOnlineStoreRequest) (*OnlineStore, error) {
	var onlineStore OnlineStore
	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.OnlineStore, &onlineStore)
	return &onlineStore, err
}

// unexported type that holds implementations of just Forecasting API methods
type forecastingImpl struct {
	client *client.DatabricksClient
}

func (a *forecastingImpl) CreateExperiment(ctx context.Context, request CreateForecastingExperimentRequest) (*CreateForecastingExperimentResponse, error) {
	var createForecastingExperimentResponse CreateForecastingExperimentResponse
	path := "/api/2.0/automl/create-forecasting-experiment"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createForecastingExperimentResponse)
	return &createForecastingExperimentResponse, err
}

func (a *forecastingImpl) GetExperiment(ctx context.Context, request GetForecastingExperimentRequest) (*ForecastingExperiment, error) {
	var forecastingExperiment ForecastingExperiment
	path := fmt.Sprintf("/api/2.0/automl/get-forecasting-experiment/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &forecastingExperiment)
	return &forecastingExperiment, err
}

// unexported type that holds implementations of just MaterializedFeatures API methods
type materializedFeaturesImpl struct {
	client *client.DatabricksClient
}

func (a *materializedFeaturesImpl) CreateFeatureTag(ctx context.Context, request CreateFeatureTagRequest) (*FeatureTag, error) {
	var featureTag FeatureTag
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.FeatureTag, &featureTag)
	return &featureTag, err
}

func (a *materializedFeaturesImpl) DeleteFeatureTag(ctx context.Context, request DeleteFeatureTagRequest) error {
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
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

func (a *materializedFeaturesImpl) GetFeatureLineage(ctx context.Context, request GetFeatureLineageRequest) (*FeatureLineage, error) {
	var featureLineage FeatureLineage
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/lineage", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &featureLineage)
	return &featureLineage, err
}

func (a *materializedFeaturesImpl) GetFeatureTag(ctx context.Context, request GetFeatureTagRequest) (*FeatureTag, error) {
	var featureTag FeatureTag
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &featureTag)
	return &featureTag, err
}

// Lists FeatureTags.
func (a *materializedFeaturesImpl) ListFeatureTags(ctx context.Context, request ListFeatureTagsRequest) listing.Iterator[FeatureTag] {

	getNextPage := func(ctx context.Context, req ListFeatureTagsRequest) (*ListFeatureTagsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListFeatureTags(ctx, req)
	}
	getItems := func(resp *ListFeatureTagsResponse) []FeatureTag {
		return resp.FeatureTags
	}
	getNextReq := func(resp *ListFeatureTagsResponse) *ListFeatureTagsRequest {
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

// Lists FeatureTags.
func (a *materializedFeaturesImpl) ListFeatureTagsAll(ctx context.Context, request ListFeatureTagsRequest) ([]FeatureTag, error) {
	iterator := a.ListFeatureTags(ctx, request)
	return listing.ToSlice[FeatureTag](ctx, iterator)
}

func (a *materializedFeaturesImpl) internalListFeatureTags(ctx context.Context, request ListFeatureTagsRequest) (*ListFeatureTagsResponse, error) {
	var listFeatureTagsResponse ListFeatureTagsResponse
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFeatureTagsResponse)
	return &listFeatureTagsResponse, err
}

func (a *materializedFeaturesImpl) UpdateFeatureTag(ctx context.Context, request UpdateFeatureTagRequest) (*FeatureTag, error) {
	var featureTag FeatureTag
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" || slices.Contains(request.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.FeatureTag, &featureTag)
	return &featureTag, err
}

// unexported type that holds implementations of just ModelRegistry API methods
type modelRegistryImpl struct {
	client *client.DatabricksClient
}

func (a *modelRegistryImpl) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error) {
	var approveTransitionRequestResponse ApproveTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/approve"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &approveTransitionRequestResponse)
	return &approveTransitionRequestResponse, err
}

func (a *modelRegistryImpl) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {
	var createCommentResponse CreateCommentResponse
	path := "/api/2.0/mlflow/comments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCommentResponse)
	return &createCommentResponse, err
}

func (a *modelRegistryImpl) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {
	var createModelResponse CreateModelResponse
	path := "/api/2.0/mlflow/registered-models/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createModelResponse)
	return &createModelResponse, err
}

func (a *modelRegistryImpl) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

func (a *modelRegistryImpl) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {
	var createTransitionRequestResponse CreateTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createTransitionRequestResponse)
	return &createTransitionRequestResponse, err
}

func (a *modelRegistryImpl) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {
	var createWebhookResponse CreateWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createWebhookResponse)
	return &createWebhookResponse, err
}

func (a *modelRegistryImpl) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	path := "/api/2.0/mlflow/comments/delete"
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

func (a *modelRegistryImpl) DeleteModel(ctx context.Context, request DeleteModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete"
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

func (a *modelRegistryImpl) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete-tag"
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

func (a *modelRegistryImpl) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete"
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

func (a *modelRegistryImpl) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete-tag"
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

func (a *modelRegistryImpl) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) (*DeleteTransitionRequestResponse, error) {
	var deleteTransitionRequestResponse DeleteTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteTransitionRequestResponse)
	return &deleteTransitionRequestResponse, err
}

func (a *modelRegistryImpl) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/delete"
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

// Gets the latest version of a registered model.
func (a *modelRegistryImpl) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) listing.Iterator[ModelVersion] {

	getNextPage := func(ctx context.Context, req GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetLatestVersions(ctx, req)
	}
	getItems := func(resp *GetLatestVersionsResponse) []ModelVersion {
		return resp.ModelVersions
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets the latest version of a registered model.
func (a *modelRegistryImpl) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	iterator := a.GetLatestVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}

func (a *modelRegistryImpl) internalGetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *modelRegistryImpl) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {
	var getModelResponse GetModelResponse
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelResponse)
	return &getModelResponse, err
}

func (a *modelRegistryImpl) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *modelRegistryImpl) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *modelRegistryImpl) GetPermissionLevels(ctx context.Context, request GetRegisteredModelPermissionLevelsRequest) (*GetRegisteredModelPermissionLevelsResponse, error) {
	var getRegisteredModelPermissionLevelsResponse GetRegisteredModelPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v/permissionLevels", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRegisteredModelPermissionLevelsResponse)
	return &getRegisteredModelPermissionLevelsResponse, err
}

func (a *modelRegistryImpl) GetPermissions(ctx context.Context, request GetRegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryImpl) ListModels(ctx context.Context, request ListModelsRequest) listing.Iterator[Model] {

	getNextPage := func(ctx context.Context, req ListModelsRequest) (*ListModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListModels(ctx, req)
	}
	getItems := func(resp *ListModelsResponse) []Model {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *ListModelsResponse) *ListModelsRequest {
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

// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryImpl) ListModelsAll(ctx context.Context, request ListModelsRequest) ([]Model, error) {
	iterator := a.ListModels(ctx, request)
	return listing.ToSlice[Model](ctx, iterator)
}

func (a *modelRegistryImpl) internalListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error) {
	var listModelsResponse ListModelsResponse
	path := "/api/2.0/mlflow/registered-models/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelsResponse)
	return &listModelsResponse, err
}

// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryImpl) ListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) listing.Iterator[Activity] {

	getNextPage := func(ctx context.Context, req ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListTransitionRequests(ctx, req)
	}
	getItems := func(resp *ListTransitionRequestsResponse) []Activity {
		return resp.Requests
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryImpl) ListTransitionRequestsAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	iterator := a.ListTransitionRequests(ctx, request)
	return listing.ToSlice[Activity](ctx, iterator)
}

func (a *modelRegistryImpl) internalListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
	var listTransitionRequestsResponse ListTransitionRequestsResponse
	path := "/api/2.0/mlflow/transition-requests/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitionRequestsResponse)
	return &listTransitionRequestsResponse, err
}

// **NOTE:** This endpoint is in Public Preview. Lists all registry webhooks.
func (a *modelRegistryImpl) ListWebhooks(ctx context.Context, request ListWebhooksRequest) listing.Iterator[RegistryWebhook] {

	getNextPage := func(ctx context.Context, req ListWebhooksRequest) (*ListRegistryWebhooks, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWebhooks(ctx, req)
	}
	getItems := func(resp *ListRegistryWebhooks) []RegistryWebhook {
		return resp.Webhooks
	}
	getNextReq := func(resp *ListRegistryWebhooks) *ListWebhooksRequest {
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

// **NOTE:** This endpoint is in Public Preview. Lists all registry webhooks.
func (a *modelRegistryImpl) ListWebhooksAll(ctx context.Context, request ListWebhooksRequest) ([]RegistryWebhook, error) {
	iterator := a.ListWebhooks(ctx, request)
	return listing.ToSlice[RegistryWebhook](ctx, iterator)
}

func (a *modelRegistryImpl) internalListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error) {
	var listRegistryWebhooks ListRegistryWebhooks
	path := "/api/2.0/mlflow/registry-webhooks/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRegistryWebhooks)
	return &listRegistryWebhooks, err
}

func (a *modelRegistryImpl) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {
	var rejectTransitionRequestResponse RejectTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/reject"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &rejectTransitionRequestResponse)
	return &rejectTransitionRequestResponse, err
}

func (a *modelRegistryImpl) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {
	var renameModelResponse RenameModelResponse
	path := "/api/2.0/mlflow/registered-models/rename"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &renameModelResponse)
	return &renameModelResponse, err
}

// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryImpl) SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) listing.Iterator[ModelVersion] {

	getNextPage := func(ctx context.Context, req SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchModelVersions(ctx, req)
	}
	getItems := func(resp *SearchModelVersionsResponse) []ModelVersion {
		return resp.ModelVersions
	}
	getNextReq := func(resp *SearchModelVersionsResponse) *SearchModelVersionsRequest {
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

// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryImpl) SearchModelVersionsAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	iterator := a.SearchModelVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}

func (a *modelRegistryImpl) internalSearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0/mlflow/model-versions/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

// Search for registered models based on the specified __filter__.
func (a *modelRegistryImpl) SearchModels(ctx context.Context, request SearchModelsRequest) listing.Iterator[Model] {

	getNextPage := func(ctx context.Context, req SearchModelsRequest) (*SearchModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchModels(ctx, req)
	}
	getItems := func(resp *SearchModelsResponse) []Model {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *SearchModelsResponse) *SearchModelsRequest {
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

// Search for registered models based on the specified __filter__.
func (a *modelRegistryImpl) SearchModelsAll(ctx context.Context, request SearchModelsRequest) ([]Model, error) {
	iterator := a.SearchModels(ctx, request)
	return listing.ToSlice[Model](ctx, iterator)
}

func (a *modelRegistryImpl) internalSearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error) {
	var searchModelsResponse SearchModelsResponse
	path := "/api/2.0/mlflow/registered-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchModelsResponse)
	return &searchModelsResponse, err
}

func (a *modelRegistryImpl) SetModelTag(ctx context.Context, request SetModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/set-tag"
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

func (a *modelRegistryImpl) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/set-tag"
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

func (a *modelRegistryImpl) SetPermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryImpl) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/test"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

func (a *modelRegistryImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	var transitionStageResponse TransitionStageResponse
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &transitionStageResponse)
	return &transitionStageResponse, err
}

func (a *modelRegistryImpl) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {
	var updateCommentResponse UpdateCommentResponse
	path := "/api/2.0/mlflow/comments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateCommentResponse)
	return &updateCommentResponse, err
}

func (a *modelRegistryImpl) UpdateModel(ctx context.Context, request UpdateModelRequest) (*UpdateModelResponse, error) {
	var updateModelResponse UpdateModelResponse
	path := "/api/2.0/mlflow/registered-models/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateModelResponse)
	return &updateModelResponse, err
}

func (a *modelRegistryImpl) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) (*UpdateModelVersionResponse, error) {
	var updateModelVersionResponse UpdateModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateModelVersionResponse)
	return &updateModelVersionResponse, err
}

func (a *modelRegistryImpl) UpdatePermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryImpl) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) (*UpdateWebhookResponse, error) {
	var updateWebhookResponse UpdateWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateWebhookResponse)
	return &updateWebhookResponse, err
}
