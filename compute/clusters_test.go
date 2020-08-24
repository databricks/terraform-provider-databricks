package compute

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
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
			want: []ClusterInfo{
				{
					ClusterName:  "autoscaling-cluster",
					SparkVersion: "5.3.x-scala2.11",
					NodeTypeID:   "i3.xlarge",
					AutoScale: &AutoScale{
						MinWorkers: 2,
						MaxWorkers: 50,
					},
				},
				{
					ClusterName:  "autoscaling-cluster2",
					SparkVersion: "5.3.x-scala2.11",
					NodeTypeID:   "i3.xlarge",
					AutoScale: &AutoScale{
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
			want:           []ClusterInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return NewClustersAPI(&client).List()
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
			want: ZonesInfo{

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
			want:           ZonesInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return NewClustersAPI(&client).ListZones()
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
			want: []NodeType{
				{
					NodeTypeID:         "r3.xlarge",
					MemoryMB:           31232,
					NumCores:           4.0,
					Description:        "r3.xlarge (deprecated)",
					InstanceTypeID:     "r3.xlarge",
					IsDeprecated:       false,
					Category:           "Memory Optimized",
					SupportEBSVolumes:  true,
					SupportClusterTags: true,
					NumGPUs:            0,
					NodeInstanceType: &NodeInstanceType{
						InstanceTypeID:  "r3.xlarge",
						LocalDisks:      1,
						LocalDiskSizeGB: 80,
					},
					IsHidden:              false,
					SupportPortForwarding: true,
					DisplayOrder:          1,
					IsIOCacheEnabled:      false,
				},
				{
					NodeTypeID:         "r3.2xlarge",
					MemoryMB:           62464,
					NumCores:           8.0,
					Description:        "r3.2xlarge (deprecated)",
					InstanceTypeID:     "r3.2xlarge",
					IsDeprecated:       false,
					Category:           "Memory Optimized",
					SupportEBSVolumes:  true,
					SupportClusterTags: true,
					NumGPUs:            0,
					NodeInstanceType: &NodeInstanceType{
						InstanceTypeID:  "r3.2xlarge",
						LocalDisks:      1,
						LocalDiskSizeGB: 160,
					},
					IsHidden:              false,
					SupportPortForwarding: true,
					DisplayOrder:          1,
					IsIOCacheEnabled:      false,
				},
			},
			wantErr: false,
		},
		{
			name:           "ListNodeTypes failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			wantURI:        "/api/2.0/clusters/list-node-types",
			want:           []NodeType{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return NewClustersAPI(&client).ListNodeTypes()
			})
		})
	}
}

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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/restart", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return nil, NewClustersAPI(&client).Restart(tt.args.ClusterID)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/pin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return nil, NewClustersAPI(&client).Pin(tt.args.ClusterID)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/unpin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client common.DatabricksClient) (interface{}, error) {
				return nil, NewClustersAPI(&client).Unpin(tt.args.ClusterID)
			})
		})
	}
}

func TestAccListClustersIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	randomName := qa.RandomName()

	cluster := Cluster{
		NumWorkers:             1,
		ClusterName:            "Terraform Integration Test " + randomName,
		SparkVersion:           CommonRuntimeVersion(),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "acc-list-" + randomName,
		AutoterminationMinutes: 15,
	}
	clusterReadInfo, err := NewClustersAPI(client).Create(cluster)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.NumWorkers == cluster.NumWorkers)
	assert.True(t, clusterReadInfo.ClusterName == cluster.ClusterName)
	assert.True(t, reflect.DeepEqual(clusterReadInfo.SparkEnvVars, cluster.SparkEnvVars))
	assert.True(t, clusterReadInfo.SparkVersion == cluster.SparkVersion)
	assert.True(t, clusterReadInfo.AutoterminationMinutes == cluster.AutoterminationMinutes)
	assert.True(t, clusterReadInfo.State == ClusterStateRunning)

	defer func() {
		err = NewClustersAPI(client).Terminate(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		clusterReadInfo, err = NewClustersAPI(client).Get(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
		assert.True(t, clusterReadInfo.State == ClusterStateTerminated)

		err = NewClustersAPI(client).Unpin(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)

		err = NewClustersAPI(client).PermanentDelete(clusterReadInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	err = NewClustersAPI(client).Pin(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)

	clusterReadInfo, err = NewClustersAPI(client).Get(clusterReadInfo.ClusterID)
	assert.NoError(t, err, err)
	assert.True(t, clusterReadInfo.State == ClusterStateRunning)
}

func TestClusters_SortNodeTypes_Deprecated(t *testing.T) {
	nodeTypes := []NodeType{
		{
			IsDeprecated: true,
			NodeTypeID:   "deprecated1",
		},
		{
			IsDeprecated: false,
			NodeTypeID:   "not deprecated",
		},
		{
			IsDeprecated: true,
			NodeTypeID:   "deprecated2",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "not deprecated", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_Memory(t *testing.T) {
	nodeTypes := []NodeType{
		{
			MemoryMB:   3,
			NodeTypeID: "3",
		},
		{
			MemoryMB:   1,
			NodeTypeID: "1",
		},
		{
			MemoryMB:   2,
			NodeTypeID: "2",
		},
		{
			MemoryMB:   2,
			NodeTypeID: "another 2",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_CPU(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumCores:   3,
			NodeTypeID: "3",
		},
		{
			NumCores:   1,
			NodeTypeID: "1",
		},
		{
			NumCores:   2,
			NodeTypeID: "2",
		},
		{
			NumCores:   1,
			NodeTypeID: "another 1",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_GPU(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumGPUs:    3,
			NodeTypeID: "3",
		},
		{
			NumGPUs:    1,
			NodeTypeID: "1",
		},
		{
			NumGPUs:    2,
			NodeTypeID: "2",
		},
		{
			NumGPUs:    1,
			NodeTypeID: "another 1",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_CPU_Deprecated(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NumCores:     3,
			IsDeprecated: false,
			NodeTypeID:   "3 not deprecated",
		},
		{
			NumCores:     1,
			IsDeprecated: true,
			NodeTypeID:   "1 deprecated",
		},
		{
			NumCores:     2,
			IsDeprecated: false,
			NodeTypeID:   "2 not deprecated",
		},
		{
			NumCores:     1,
			IsDeprecated: false,
			NodeTypeID:   "1 not deprecated",
		},
		{
			NumCores:     2,
			IsDeprecated: true,
			NodeTypeID:   "2 deprecated",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1 not deprecated", smallestNodeType.NodeTypeID)
}

func TestClusters_SortNodeTypes_LocalDisks(t *testing.T) {
	nodeTypes := []NodeType{
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 3,
			},
			NodeTypeID: "3",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 1,
			},
			NodeTypeID: "1",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 2,
			},
			NodeTypeID: "2",
		},
		{
			NodeInstanceType: &NodeInstanceType{
				LocalDisks: 3,
			},
			NodeTypeID: "another 3",
		},
	}

	smallestNodeType := getSmallestNodeType(nodeTypes)
	assert.Equal(t, "1", smallestNodeType.NodeTypeID)
}

func TestAwsAccSmallestNodeType(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	nodeType := NewClustersAPI(client).GetSmallestNodeTypeWithStorage()
	assert.Equal(t, "m5d.large", nodeType)
}