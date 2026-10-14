package sql

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
)

// The SDK models do not yet include statement_timeout or customer_version_label.
type realtimeWarehouseFields struct {
	StatementTimeout     int    `json:"statement_timeout" tf:"computed,optional"`
	CustomerVersionLabel string `json:"customer_version_label" tf:"computed"`
}

func writeRealtimeWarehouse(ctx context.Context, w *databricks.WorkspaceClient, path string, request any, timeout int, response any) error {
	encoded, err := json.Marshal(request)
	if err != nil {
		return err
	}
	var body map[string]any
	if err := json.Unmarshal(encoded, &body); err != nil {
		return err
	}
	body["min_num_clusters"], body["max_num_clusters"] = 1, 1
	// Zero clears an existing override and must survive JSON serialization.
	body["statement_timeout"] = timeout
	return doRealtimeWarehouseRequest(ctx, w, http.MethodPost, path, body, response)
}

func doRealtimeWarehouseRequest(ctx context.Context, w *databricks.WorkspaceClient, method, path string, request, response any) error {
	c, err := client.New(w.Config)
	if err != nil {
		return err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	if w.Config.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = w.Config.WorkspaceID
	}
	return c.Do(ctx, method, path, headers, nil, request, response)
}
