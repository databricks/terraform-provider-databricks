// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package idempotencytesting

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just IdempotencyTesting API methods
type idempotencyTestingImpl struct {
	client *client.DatabricksClient
}

func (a *idempotencyTestingImpl) CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (*TestResource, error) {
	var testResource TestResource
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := "/api/2.0/idempotency-testing/resources"
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
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.TestResource, &testResource)
	return &testResource, err
}
