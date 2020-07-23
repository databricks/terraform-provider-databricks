package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
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
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, tt.postStructExpect, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Commands().Execute(tt.params.ClusterID, tt.params.Language, tt.params.CommandStr)
			})
		})
	}
}

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := NewClientFromEnvironment()
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	cluster := model.Cluster{
		NumWorkers:  1,
		ClusterName: "Terraform Integration Test " + randomName,
		SparkEnvVars: map[string]string{
			"PYSPARK_PYTHON": "/databricks/python3/bin/python3",
		},
		SparkVersion:           "6.2.x-scala2.11",
		NodeTypeID:             GetCloudInstanceType(client),
		DriverNodeTypeID:       GetCloudInstanceType(client),
		IdempotencyToken:       "commands-" + randomName,
		AutoterminationMinutes: 20,
	}

	if cloud == "AWS" {
		cluster.AwsAttributes = &model.AwsAttributes{
			EbsVolumeType:  model.EbsVolumeTypeGeneralPurposeSsd,
			EbsVolumeCount: 1,
			EbsVolumeSize:  32,
		}
	}

	clusterInfo, err := client.Clusters().Create(cluster)
	assert.NoError(t, err, err)
	defer func() {
		err := client.Clusters().PermanentDelete(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	clusterID := clusterInfo.ClusterID

	// TODO: Cluster is in a non runnable state will not be able to transition to running, needs to be started again. Current state: TERMINATED
	err = client.Clusters().WaitForClusterRunning(clusterID, 10, 20)
	assert.NoError(t, err, err)

	context, err := client.Commands().createContext("python", clusterID)
	assert.NoError(t, err, err)
	t.Log(context)

	err = client.Commands().waitForContextReady(context, clusterID, 1, 1)
	assert.NoError(t, err, err)

	status, err := client.Commands().getContext(context, clusterID)
	assert.NoError(t, err, err)
	assert.True(t, status == "Running")
	t.Log(status)

	commandID, err := client.Commands().createCommand(context, clusterID, "python", "print('hello world')")
	assert.NoError(t, err, err)

	err = client.Commands().waitForCommandFinished(commandID, context, clusterID, 5, 20)
	assert.NoError(t, err, err)

	resp, err := client.Commands().getCommand(commandID, context, clusterID)
	assert.NoError(t, err, err)
	assert.NotNil(t, resp.Results.Data)

	// Testing the public api Execute
	command, err := client.Commands().Execute(clusterID, "python", "print('hello world')")
	assert.NoError(t, err, err)
	assert.NotNil(t, command.Results.Data)
}
