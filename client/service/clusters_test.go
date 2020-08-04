package service

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestClustersAPI_List(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		responseStatus int
		wantURI        string
		want           interface{}
		wantErr        bool
	}{
		{
			name: "List test",
			response: `{
						   "clusters":[
							  {
								 "cluster_name":"autoscaling-cluster",
								 "spark_version":"5.3.x-scala2.11",
								 "node_type_id":"i3.xlarge",
								 "autoscale":{
									"min_workers":2,
									"max_workers":50
								 }
							  },
							  {
								 "cluster_name":"autoscaling-cluster2",
								 "spark_version":"5.3.x-scala2.11",
								 "node_type_id":"i3.xlarge",
								 "autoscale":{
									"min_workers":2,
									"max_workers":50
								 }
							  }
						   ]
						}`,
			responseStatus: http.StatusOK,
			wantURI:        "/api/2.0/clusters/list",
			want: []model.ClusterInfo{
				{
					ClusterName:  "autoscaling-cluster",
					SparkVersion: "5.3.x-scala2.11",
					NodeTypeID:   "i3.xlarge",
					AutoScale: &model.AutoScale{
						MinWorkers: 2,
						MaxWorkers: 50,
					},
				},
				{
					ClusterName:  "autoscaling-cluster2",
					SparkVersion: "5.3.x-scala2.11",
					NodeTypeID:   "i3.xlarge",
					AutoScale: &model.AutoScale{
						MinWorkers: 2,
						MaxWorkers: 50,
					},
				},
			},
			wantErr: false,
		},
		{
			name:           "List failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			wantURI:        "/api/2.0/clusters/list",
			want:           []model.ClusterInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Clusters().List()
			})
		})
	}
}

func TestClustersAPI_ListZones(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		responseStatus int
		wantURI        string
		want           interface{}
		wantErr        bool
	}{
		{
			name: "ListZones test",
			response: `{
							"zones": [
								"us-west-2b",
								"us-west-2c",
								"us-west-2a"
							],
							"default_zone": "us-west-2b"
						}`,
			responseStatus: http.StatusOK,
			wantURI:        "/api/2.0/clusters/list-zones",
			want: model.ZonesInfo{

				Zones: []string{"us-west-2b",
					"us-west-2c",
					"us-west-2a"},
				DefaultZone: "us-west-2b",
			},
			wantErr: false,
		},
		{
			name:           "ListZones failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			wantURI:        "/api/2.0/clusters/list-zones",
			want:           model.ZonesInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Clusters().ListZones()
			})
		})
	}
}

func TestClustersAPI_ListNodeTypes(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		responseStatus int
		wantURI        string
		want           interface{}
		wantErr        bool
	}{
		{
			name: "ListNodeTypes test",
			response: `{
							"node_types": [
								{
									"node_type_id": "r3.xlarge",
									"memory_mb": 31232,
									"num_cores": 4.0,
									"description": "r3.xlarge (deprecated)",
									"instance_type_id": "r3.xlarge",
									"is_deprecated": false,
									"category": "Memory Optimized",
									"support_ebs_volumes": true,
									"support_cluster_tags": true,
									"num_gpus": 0,
									"node_instance_type": {
										"instance_type_id": "r3.xlarge",
										"local_disks": 1,
										"local_disk_size_gb": 80
									},
									"is_hidden": false,
									"support_port_forwarding": true,
									"display_order": 1,
									"is_io_cache_enabled": false
								},
								{
									"node_type_id": "r3.2xlarge",
									"memory_mb": 62464,
									"num_cores": 8.0,
									"description": "r3.2xlarge (deprecated)",
									"instance_type_id": "r3.2xlarge",
									"is_deprecated": false,
									"category": "Memory Optimized",
									"support_ebs_volumes": true,
									"support_cluster_tags": true,
									"num_gpus": 0,
									"node_instance_type": {
										"instance_type_id": "r3.2xlarge",
										"local_disks": 1,
										"local_disk_size_gb": 160
									},
									"is_hidden": false,
									"support_port_forwarding": true,
									"display_order": 1,
									"is_io_cache_enabled": false
								}
							]
						}`,
			responseStatus: http.StatusOK,
			wantURI:        "/api/2.0/clusters/list-node-types",
			want: []model.NodeType{
				{
					NodeTypeID:     "r3.xlarge",
					MemoryMb:       31232,
					NumCores:       4.0,
					Description:    "r3.xlarge (deprecated)",
					InstanceTypeID: "r3.xlarge",
					IsDeprecated:   false,
				},
				{
					NodeTypeID:     "r3.2xlarge",
					MemoryMb:       62464,
					NumCores:       8.0,
					Description:    "r3.2xlarge (deprecated)",
					InstanceTypeID: "r3.2xlarge",
					IsDeprecated:   false,
				},
			},
			wantErr: false,
		},
		{
			name:           "ListNodeTypes failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			wantURI:        "/api/2.0/clusters/list-node-types",
			want:           []model.NodeType{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Clusters().ListNodeTypes()
			})
		})
	}
}

// func TestClustersAPI_WaitForClusterRunning(t *testing.T) {
// 	type args struct {
// 		ClusterID              string        `json:"cluster_id"`
// 		SleepDurationSeconds   time.Duration `json:"sleep_duration_seconds"`
// 		TimeoutDurationMinutes time.Duration `json:"timeout_duration_minutes"`
// 	}
// 	tests := []struct {
// 		name           string
// 		response       []string
// 		responseStatus []int
// 		requestMethod  []string
// 		args           []interface{}
// 		wantURI        []string
// 		want           []model.WorkspaceObjectStatus
// 		wantErr        bool
// 	}{
// 		{
// 			name: "WaitForClusterRunning test",
// 			response: []string{`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "RUNNING"
// 								}`,
// 			},
// 			responseStatus: []int{http.StatusOK,
// 				http.StatusOK,
// 				http.StatusOK,
// 			},
// 			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
// 			args: []interface{}{
// 				&args{
// 					ClusterID:              "11203-my-cluster",
// 					SleepDurationSeconds:   0,
// 					TimeoutDurationMinutes: 1,
// 				},
// 			},
// 			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
// 			want:    nil,
// 			wantErr: false,
// 		},
// 		{
// 			name: "WaitForClusterRunning failed to get cluster info test",
// 			response: []string{`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "RUNNING"
// 								}`,
// 			},
// 			responseStatus: []int{http.StatusOK,
// 				http.StatusOK,
// 				http.StatusBadRequest,
// 			},
// 			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
// 			args: []interface{}{
// 				&args{
// 					ClusterID:              "11203-my-cluster",
// 					SleepDurationSeconds:   0,
// 					TimeoutDurationMinutes: 1,
// 				},
// 			},
// 			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "WaitForClusterRunning failed cluster invalid state test",
// 			response: []string{`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "TERMINATING"
// 								}`,
// 			},
// 			responseStatus: []int{http.StatusOK,
// 				http.StatusOK,
// 				http.StatusOK,
// 			},
// 			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
// 			args: []interface{}{
// 				&args{
// 					ClusterID:              "11203-my-cluster",
// 					SleepDurationSeconds:   0,
// 					TimeoutDurationMinutes: 1,
// 				},
// 			},
// 			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "WaitForClusterRunning failed due to timeout test",
// 			response: []string{`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "PENDING"
// 								}`,
// 				`{
// 									"state": "RUNNING"
// 								}`,
// 			},
// 			responseStatus: []int{http.StatusOK,
// 				http.StatusOK,
// 				http.StatusBadRequest,
// 			},
// 			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
// 			args: []interface{}{
// 				&args{
// 					ClusterID:              "11203-my-cluster",
// 					SleepDurationSeconds:   1,
// 					TimeoutDurationMinutes: 0,
// 				},
// 			},
// 			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
// 				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
// 				return nil, client.Clusters().WaitForClusterRunning(tt.args[0].(*args).ClusterID)
// 			})
// 		})
// 	}
// }

func TestClustersAPI_Restart(t *testing.T) {
	type args struct {
		ClusterID string `json:"cluster_id,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name:           "Restart test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Restart faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/restart", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.Clusters().Restart(tt.args.ClusterID)
			})
		})
	}
}

func TestClustersAPI_Pin(t *testing.T) {
	type args struct {
		ClusterID string `json:"cluster_id,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name:           "Pin test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Pin faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/pin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.Clusters().Pin(tt.args.ClusterID)
			})
		})
	}
}

func TestClustersAPI_Unpin(t *testing.T) {
	type args struct {
		ClusterID string `json:"cluster_id,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name:           "Unpin test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Unpin faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/unpin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.Clusters().Unpin(tt.args.ClusterID)
			})
		})
	}
}

func TestAccListClustersIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := CommonEnvironmentClient()
	randomName := randomName()

	cluster := model.Cluster{
		NumWorkers:             1,
		ClusterName:            "Terraform Integration Test " + randomName,
		SparkVersion:           CommonRuntimeVersion(),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "acc-list-" + randomName,
		AutoterminationMinutes: 15,
	}
	clusterReadInfo, err := client.Clusters().Create(cluster)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)
	assert.True(t, clusterReadInfo.State == model.ClusterStateRunning)

	defer func() {
		err = client.Clusters().Terminate(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		clusterReadInfo, err = client.Clusters().Get(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
		assert.True(t, clusterReadInfo.State == model.ClusterStateTerminated)

		err = client.Clusters().Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = client.Clusters().PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = client.Clusters().Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	clusterReadInfo, err = client.Clusters().Get(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == model.ClusterStateRunning)
}
