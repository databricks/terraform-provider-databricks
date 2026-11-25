// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package computemigrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just ComputeMigrations API methods
type computeMigrationsImpl struct {
	client *client.DatabricksClient
}

func (a *computeMigrationsImpl) ApproveMigrationExecution(ctx context.Context, request ApproveMigrationExecutionRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v/approveExecution", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) ApproveMigrationTranslation(ctx context.Context, request ApproveMigrationTranslationRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v/approveTranslation", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) CancelMigration(ctx context.Context, request CancelMigrationRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v/cancel", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) Create(ctx context.Context, request CreateMigrationRequest) (*Migration, error) {
	var migration Migration
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/compute-migrations/v1/migrations"
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
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Migration, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) Delete(ctx context.Context, request DeleteMigrationRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) Get(ctx context.Context, request GetMigrationRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &migration)
	return &migration, err
}

// Lists migrations in the caller's workspace, ordered by migration ID, with
// cursor-based pagination and filtering by asset type and state.
//
// Eventually consistent: results are read from a cache, and pagination is not a
// point-in-time snapshot, so a migration created during a List may not appear
// until a subsequent one.
func (a *computeMigrationsImpl) List(ctx context.Context, request ListMigrationsRequest) listing.Iterator[Migration] {

	getNextPage := func(ctx context.Context, req ListMigrationsRequest) (*ListMigrationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListMigrationsResponse) []Migration {
		return resp.Migrations
	}
	getNextReq := func(resp *ListMigrationsResponse) *ListMigrationsRequest {
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

// Lists migrations in the caller's workspace, ordered by migration ID, with
// cursor-based pagination and filtering by asset type and state.
//
// Eventually consistent: results are read from a cache, and pagination is not a
// point-in-time snapshot, so a migration created during a List may not appear
// until a subsequent one.
func (a *computeMigrationsImpl) ListAll(ctx context.Context, request ListMigrationsRequest) ([]Migration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Migration](ctx, iterator)
}

func (a *computeMigrationsImpl) internalList(ctx context.Context, request ListMigrationsRequest) (*ListMigrationsResponse, error) {
	var listMigrationsResponse ListMigrationsResponse
	path := "/api/compute-migrations/v1/migrations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMigrationsResponse)
	return &listMigrationsResponse, err
}

func (a *computeMigrationsImpl) RejectMigrationExecution(ctx context.Context, request RejectMigrationExecutionRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v/rejectExecution", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migration)
	return &migration, err
}

func (a *computeMigrationsImpl) RejectMigrationTranslation(ctx context.Context, request RejectMigrationTranslationRequest) (*Migration, error) {
	var migration Migration
	path := fmt.Sprintf("/api/compute-migrations/v1/migrations/%v/rejectTranslation", request.MigrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migration)
	return &migration, err
}
