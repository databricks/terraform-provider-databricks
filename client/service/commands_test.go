package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// Test interface compliance
var _ CommandExecutor = (*CommandsAPI)(nil)

func TestCommandsAPI_Execute(t *testing.T) {
	type context struct {
		Language  string `json:"language,omitempty"`
		ClusterID string `json:"clusterId,omitempty"`
	}
	type command struct {
		Language  string `json:"language,omitempty"`
		ClusterID string `json:"clusterId,omitempty"`
		ContextID string `json:"contextId,omitempty"`
		Command   string `json:"command,omitempty"`
	}
	type contextDelete struct {
		ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
		ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	}

	type params struct {
		ClusterID  string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Language   string `json:"language,omitempty" url:"language,omitempty"`
		CommandStr string `json:"command_str,omitempty" url:"command_str,omitempty"`
	}
	tests := []struct {
		params           params
		name             string
		response         []string
		responseStatus   []int
		requestMethod    []string
		postStructExpect []interface{}
		args             []interface{}
		wantURI          []string
		want             interface{}
		wantErr          bool
	}{
		{
			name: "Execute test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want: model.Command{
				ID:     "my-command-id",
				Status: "Finished",
				Results: &model.CommandResults{
					Data: "Hello world",
				},
			},
			wantErr: false,
		},
		{
			name: "Execute context failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusBadRequest, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute context status failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute command create failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusBadRequest, http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute command status failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusBadRequest, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute command results fetch failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusBadRequest, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute context close failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want: model.Command{
				ID:     "my-command-id",
				Status: "Finished",
				Results: &model.CommandResults{
					Data: "Hello world",
				},
			},
			wantErr: true,
		},
		{
			name: "Execute context invalid state failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Error"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
		{
			name: "Execute command invalid state failure test",
			params: params{
				ClusterID:  "my-cluster-id",
				Language:   "python",
				CommandStr: `print("hello world")`,
			},
			response: []string{
				`{
					"id": "my-context-id"
				}`, `
				{
					"status": "Running"
				}`,
				`{
					"id": "my-command-id"
				}`,
				`{
					"id": "my-command-id",
					"status": "Error"
				}`,
				`{
					"id": "my-command-id",
					"status": "Finished",
					"results": {
						"data": "Hello world"
					}
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodGet, http.MethodGet, http.MethodPost},
			args: []interface{}{
				&context{
					Language:  "python",
					ClusterID: "my-cluster-id",
				},
				nil,
				&command{
					Language:  "python",
					ClusterID: "my-cluster-id",
					ContextID: "my-context-id",
					Command:   `print("hello world")`,
				},
				nil,
				nil,
				&contextDelete{
					ContextID: "my-context-id",
					ClusterID: "my-cluster-id",
				},
			},
			postStructExpect: []interface{}{
				&context{},
				nil,
				&command{},
				nil,
				nil,
				&contextDelete{},
			},
			wantURI: []string{"/api/1.2/contexts/create",
				"/api/1.2/contexts/status?clusterId=my-cluster-id&contextId=my-context-id",
				"/api/1.2/commands/execute",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/commands/status?clusterId=my-cluster-id&commandId=my-command-id&contextId=my-context-id",
				"/api/1.2/contexts/destroy"},
			want:    model.Command{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, tt.postStructExpect, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Commands().Execute(tt.params.ClusterID, tt.params.Language, tt.params.CommandStr)
			})
		})
	}
}
