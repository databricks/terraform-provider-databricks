package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestClustersAPI_Create(t *testing.T) {
	type args model.Cluster

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Create test",
			response: `{
						  "cluster_id": "my-cluster"
						}`,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    model.ClusterInfo{ClusterID: "my-cluster"},
			wantErr: false,
		},
		{
			name:           "Create faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    model.ClusterInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/create", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Clusters().Create(model.Cluster(tt.args))
			})
		})
	}
}

func TestClustersAPI_Get(t *testing.T) {
	type args struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Get test",
			response: `{
						"cluster_id": "11203-my-cluster",
						"driver": {
							"public_dns": "",
							"node_id": "node_id",
							"node_aws_attributes": {
								"is_spot": false
							},
							"instance_id": "i-instanceid",
							"start_timestamp": 1587334045988,
							"host_private_ip": "10.0.0.11",
							"private_ip": "10.0.0.10"
						},
						"executors": [
							{
								"public_dns": "",
								"node_id": "workernode_id",
								"node_aws_attributes": {
									"is_spot": true
								},
								"instance_id": "i-workerinstance",
								"start_timestamp": 1587334045953,
								"host_private_ip": "10.0.0.21",
								"private_ip": "10.0.0.20"
							}
						],
						"spark_context_id": 51523452250940270,
						"jdbc_port": 10000,
						"cluster_name": "unit-test-demo",
						"spark_version": "6.4.x-scala2.11",
						"aws_attributes": {
							"zone_id": "us-west-2b",
							"first_on_demand": 1,
							"availability": "SPOT_WITH_FALLBACK",
							"spot_bid_price_percent": 100,
							"ebs_volume_count": 0
						},
						"node_type_id": "i3.xlarge",
						"driver_node_type_id": "i3.xlarge",
						"spark_env_vars": {
							"PYSPARK_PYTHON": "/databricks/python3/bin/python3"
						},
						"autotermination_minutes": 120,
						"enable_elastic_disk": false,
						"cluster_source": "UI",
						"enable_local_disk_encryption": false,
						"state": "RUNNING",
						"state_message": "",
						"start_time": 1587334045566,
						"terminated_time": 0,
						"last_state_loss_time": 0,
						"last_activity_time": 1587335172192,
						"autoscale": {
							"min_workers": 2,
							"max_workers": 8
						},
						"cluster_memory_mb": 93696,
						"cluster_cores": 12.0,
						"default_tags": {
							"Vendor": "Databricks",
							"Creator": "unit.test@databricks.com",
							"ClusterName": "unit-test-demo",
							"ClusterId": "11203-my-cluster"
						},
						"creator_user_name": "unit.test@databricks.com",
						"init_scripts_safe_mode": false
					}`,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "11203-my-cluster",
			},
			wantURI: "/api/2.0/clusters/get?cluster_id=11203-my-cluster",
			want: model.ClusterInfo{
				AutoScale: &model.AutoScale{
					MinWorkers: 2,
					MaxWorkers: 8,
				},
				ClusterID:       "11203-my-cluster",
				CreatorUserName: "unit.test@databricks.com",
				Driver: &model.SparkNode{
					PrivateIP:         "10.0.0.10",
					PublicDNS:         "",
					NodeID:            "node_id",
					InstanceID:        "i-instanceid",
					StartTimestamp:    1587334045988,
					NodeAwsAttributes: &model.SparkNodeAwsAttributes{IsSpot: false},
					HostPrivateIP:     "10.0.0.11",
				},
				Executors: []model.SparkNode{
					{

						PrivateIP:         "10.0.0.20",
						PublicDNS:         "",
						NodeID:            "workernode_id",
						InstanceID:        "i-workerinstance",
						StartTimestamp:    1587334045953,
						NodeAwsAttributes: &model.SparkNodeAwsAttributes{IsSpot: true},
						HostPrivateIP:     "10.0.0.21",
					},
				},
				SparkContextID: 51523452250940270,
				JdbcPort:       10000,
				ClusterName:    "unit-test-demo",
				SparkVersion:   "6.4.x-scala2.11",
				AwsAttributes: &model.AwsAttributes{
					FirstOnDemand:       1,
					Availability:        "SPOT_WITH_FALLBACK",
					ZoneID:              "us-west-2b",
					SpotBidPricePercent: 100,
					EbsVolumeCount:      0,
				},
				NodeTypeID:       "i3.xlarge",
				DriverNodeTypeID: "i3.xlarge",
				SparkEnvVars: map[string]string{
					"PYSPARK_PYTHON": "/databricks/python3/bin/python3",
				},
				AutoterminationMinutes: 120,
				EnableElasticDisk:      false,
				ClusterSource:          "UI",
				State:                  "RUNNING",
				StateMessage:           "",
				StartTime:              1587334045566,
				TerminateTime:          0,
				LastStateLossTime:      0,
				LastActivityTime:       1587335172192,

				ClusterMemoryMb: 93696,
				ClusterCores:    12.0,
				DefaultTags: map[string]string{
					"Vendor":      "Databricks",
					"Creator":     "unit.test@databricks.com",
					"ClusterName": "unit-test-demo",
					"ClusterId":   "11203-my-cluster",
				},
			},
			wantErr: false,
		},
		{
			name:           "Get failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "11203-my-cluster",
			},
			wantURI: "/api/2.0/clusters/get?cluster_id=11203-my-cluster",
			want:    model.ClusterInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Clusters().Get(tt.args.ClusterID)
			})
		})
	}
}

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
			wantURI:        "/api/2.0/clusters/list?",
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
			wantURI:        "/api/2.0/clusters/list?",
			want:           []model.ClusterInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
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
			wantURI:        "/api/2.0/clusters/list-zones?",
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
			wantURI:        "/api/2.0/clusters/list-zones?",
			want:           model.ZonesInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
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
			wantURI:        "/api/2.0/clusters/list-node-types?",
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
			wantURI:        "/api/2.0/clusters/list-node-types?",
			want:           []model.NodeType{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, tt.wantURI, nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Clusters().ListNodeTypes()
			})
		})
	}
}

func TestClustersAPI_WaitForClusterRunning(t *testing.T) {
	type args struct {
		ClusterID              string        `json:"cluster_id"`
		SleepDurationSeconds   time.Duration `json:"sleep_duration_seconds"`
		TimeoutDurationMinutes time.Duration `json:"timeout_duration_minutes"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		requestMethod  []string
		args           []interface{}
		wantURI        []string
		want           []model.NotebookInfo
		wantErr        bool
	}{
		{
			name: "WaitForClusterRunning test",
			response: []string{`{
									"state": "PENDING"				
								}`,
				`{
									"state": "PENDING"				
								}`,
				`{
									"state": "RUNNING"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusOK,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: false,
		},
		{
			name: "WaitForClusterRunning failed to get cluster info test",
			response: []string{`{
									"state": "PENDING"				
								}`,
				`{
									"state": "PENDING"				
								}`,
				`{
									"state": "RUNNING"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusBadRequest,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WaitForClusterRunning failed cluster invalid state test",
			response: []string{`{
									"state": "PENDING"				
								}`,
				`{
									"state": "PENDING"				
								}`,
				`{
									"state": "TERMINATING"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusOK,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WaitForClusterRunning failed due to timeout test",
			response: []string{`{
									"state": "PENDING"				
								}`,
				`{
									"state": "PENDING"				
								}`,
				`{
									"state": "RUNNING"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusBadRequest,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   1,
					TimeoutDurationMinutes: 0,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().WaitForClusterRunning(tt.args[0].(*args).ClusterID, tt.args[0].(*args).SleepDurationSeconds, tt.args[0].(*args).TimeoutDurationMinutes)
			})
		})
	}
}

func TestClustersAPI_WaitForClusterTerminated(t *testing.T) {
	type args struct {
		ClusterID              string        `json:"cluster_id"`
		SleepDurationSeconds   time.Duration `json:"sleep_duration_seconds"`
		TimeoutDurationMinutes time.Duration `json:"timeout_duration_minutes"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		requestMethod  []string
		args           []interface{}
		wantURI        []string
		want           []model.NotebookInfo
		wantErr        bool
	}{
		{
			name: "WaitForClusterTerminated test",
			response: []string{`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATED"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusOK,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: false,
		},
		{
			name: "WaitForClusterTerminated failed to get cluster info test",
			response: []string{`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATED"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusBadRequest,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WaitForClusterTerminated failed cluster invalid state test",
			response: []string{`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "UNKNOWN"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusOK,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   0,
					TimeoutDurationMinutes: 1,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WaitForClusterTerminated failed due to timeout test",
			response: []string{`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATING"				
								}`,
				`{
									"state": "TERMINATED"				
								}`,
			},
			responseStatus: []int{http.StatusOK,
				http.StatusOK,
				http.StatusBadRequest,
			},
			requestMethod: []string{http.MethodGet, http.MethodGet, http.MethodGet},
			args: []interface{}{
				&args{
					ClusterID:              "11203-my-cluster",
					SleepDurationSeconds:   1,
					TimeoutDurationMinutes: 0,
				},
			},
			wantURI: []string{"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster",
				"/api/2.0/clusters/get?cluster_id=11203-my-cluster"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().WaitForClusterTerminated(tt.args[0].(*args).ClusterID, tt.args[0].(*args).SleepDurationSeconds, tt.args[0].(*args).TimeoutDurationMinutes)
			})
		})
	}
}

func TestClustersAPI_Edit(t *testing.T) {
	type args model.Cluster

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Edit test",
			response: `{
						  "cluster_name": "my-cluster"
						}`,
			responseStatus: http.StatusOK,
			args: args{
				ClusterName: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Edit faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterName: "my-cluster-id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/edit", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().Edit(model.Cluster(tt.args))
			})
		})
	}
}

func TestClustersAPI_Start(t *testing.T) {
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
			name:           "Start test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Start faulure test",
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/start", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().Start(tt.args.ClusterID)
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/restart", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/pin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/unpin", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().Unpin(tt.args.ClusterID)
			})
		})
	}
}

func TestClustersAPI_Delete(t *testing.T) {
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
			name:           "Delete test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Delete faulure test",
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/delete", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().Delete(tt.args.ClusterID)
			})
		})
	}
}

func TestClustersAPI_PermanentDelete(t *testing.T) {
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
			name:           "PermanentDelete test",
			response:       ``,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "PermanentDelete faulure test",
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/clusters/permanent-delete", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Clusters().PermanentDelete(tt.args.ClusterID)
			})
		})
	}
}
