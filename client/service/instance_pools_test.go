package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

//go:generate easytags $GOFILE

func TestInstancePoolsAPI_Create(t *testing.T) {
	type args struct {
		InstancePool *model.InstancePool `json:"instance_pool"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     model.InstancePoolAndStats
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "instance_pool_id": "0101-120000-brick1-pool-ABCD1234"
						}`,
			args: args{
				InstancePool: &model.InstancePool{
					InstancePoolName:                   "",
					MinIdleInstances:                   0,
					MaxCapacity:                        10,
					NodeTypeID:                         "Standard_DS3_v2",
					IdleInstanceAutoTerminationMinutes: 60,
					EnableElasticDisk:                  false,
					DiskSpec: &model.InstancePoolDiskSpec{
						DiskType: &model.InstancePoolDiskType{
							AzureDiskVolumeType: "",
						},
						DiskCount: 1,
						DiskSize:  10,
					},
				},
			},
			want: model.InstancePoolAndStats{
				InstancePoolID: "0101-120000-brick1-pool-ABCD1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input model.InstancePool
			AssertRequestWithMockServer(t, tt.args.InstancePool, http.MethodPost, "/api/2.0/instance-pools/create", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.InstancePools().Create(*tt.args.InstancePool)
			})
		})
	}
}

func TestInstancePoolsAPI_Delete(t *testing.T) {
	type args struct {
		InstancePoolID string `json:"instance_pool_id"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Basic test",
			response: "",
			args: args{
				InstancePoolID: "0101-120000-brick1-pool-ABCD1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-pools/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.InstancePools().Delete(tt.args.InstancePoolID)
			})
		})
	}
}

func TestInstancePoolsAPI_Update(t *testing.T) {
	type args struct {
		InstancePoolInfo *model.InstancePoolAndStats `json:"instance_pool_info"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Basic test",
			response: "",
			args: args{
				InstancePoolInfo: &model.InstancePoolAndStats{
					InstancePoolID:                     "0101-120000-brick1-pool-ABCD1234",
					MinIdleInstances:                   0,
					MaxCapacity:                        10,
					NodeTypeID:                         "Standard_DS3_v2",
					IdleInstanceAutoTerminationMinutes: 60,
					EnableElasticDisk:                  false,
					DiskSpec: &model.InstancePoolDiskSpec{
						DiskType: &model.InstancePoolDiskType{
							AzureDiskVolumeType: "",
						},
						DiskCount: 1,
						DiskSize:  10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input model.InstancePoolAndStats
			AssertRequestWithMockServer(t, tt.args.InstancePoolInfo, http.MethodPost, "/api/2.0/instance-pools/edit", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.InstancePools().Update(*tt.args.InstancePoolInfo)
			})
		})
	}
}

func TestInstancePoolsAPI_Read(t *testing.T) {
	type args struct {
		InstancePoolID string `json:"instance_pool_id"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     model.InstancePoolAndStats
		wantErr  bool
	}{
		{
			name: "Basic Test",
			response: `{
						  "instance_pool_name": "mypool",
						  "min_idle_instances": 0,
						  "aws_attributes": {
							"availability": "SPOT",
							"zone_id": "us-west-2a",
							"spot_bid_price_percent": 100
						  },
						  "node_type_id": "c4.2xlarge",
						  "idle_instance_autotermination_minutes": 60,
						  "enable_elastic_disk": false,
						  "disk_spec": {
							"disk_type": {
							  "ebs_volume_type": "GENERAL_PURPOSE_SSD"
							},
							"disk_count": 1,
							"disk_size": 100
						  },
						  "preloaded_spark_versions": [
							"5.4.x-scala2.11"
						  ],
						  "instance_pool_id": "101-120000-brick1-pool-ABCD1234",
						  "default_tags": {
							"Vendor": "Databricks",
							"DatabricksInstancePoolCreatorId": "100125",
							"DatabricksInstancePoolId": "101-120000-brick1-pool-ABCD1234"
						  },
						  "state": "ACTIVE",
						  "stats": {
							"used_count": 10,
							"idle_count": 5,
							"pending_used_count": 5,
							"pending_idle_count": 5
						  },
						  "status": {}
						}`,
			args: args{
				InstancePoolID: "101-120000-brick1-pool-ABCD1234",
			},
			want: model.InstancePoolAndStats{
				InstancePoolID:   "101-120000-brick1-pool-ABCD1234",
				InstancePoolName: "mypool",
				MinIdleInstances: 0,
				AwsAttributes: &model.InstancePoolAwsAttributes{
					Availability:        model.AwsAvailabilitySpot,
					ZoneID:              "us-west-2a",
					SpotBidPricePercent: 100,
				},
				NodeTypeID:                         "c4.2xlarge",
				IdleInstanceAutoTerminationMinutes: 60,
				EnableElasticDisk:                  false,
				DiskSpec: &model.InstancePoolDiskSpec{
					DiskType: &model.InstancePoolDiskType{
						EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
					},
					DiskCount: 1,
					DiskSize:  100,
				},
				DefaultTags: map[string]string{
					"Vendor":                          "Databricks",
					"DatabricksInstancePoolCreatorId": "100125",
					"DatabricksInstancePoolId":        "101-120000-brick1-pool-ABCD1234",
				},
				PreloadedSparkVersions: []string{
					"5.4.x-scala2.11",
				},
				State: "ACTIVE",
				Stats: &model.InstancePoolStats{
					UsedCount:        10,
					IdleCount:        5,
					PendingUsedCount: 5,
					PendingIdleCount: 5,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input model.InstancePoolAndStats
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/instance-pools/get?instance_pool_id=101-120000-brick1-pool-ABCD1234", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.InstancePools().Read(tt.args.InstancePoolID)
			})
		})
	}
}

func TestAccInstancePools(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := NewClientFromEnvironment()

	pool := model.InstancePool{
		InstancePoolName:                   "Terraform Integration Test",
		MinIdleInstances:                   0,
		MaxCapacity:                        10,
		NodeTypeID:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	}
	if !client.IsAzure() {
		pool.DiskSpec = &model.InstancePoolDiskSpec{
			DiskType: &model.InstancePoolDiskType{
				EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		}
		pool.AwsAttributes = &model.InstancePoolAwsAttributes{
			Availability: model.AwsAvailabilitySpot,
		}
	}
	poolInfo, err := client.InstancePools().Create(pool)
	assert.NoError(t, err, err)

	defer func() {
		err := client.InstancePools().Delete(poolInfo.InstancePoolID)
		assert.NoError(t, err, err)
	}()

	poolReadInfo, err := client.InstancePools().Read(poolInfo.InstancePoolID)
	assert.NoError(t, err, err)
	assert.Equal(t, poolInfo.InstancePoolID, poolReadInfo.InstancePoolID)
	assert.Equal(t, pool.InstancePoolName, poolReadInfo.InstancePoolName)
	assert.Equal(t, pool.MinIdleInstances, poolReadInfo.MinIdleInstances)
	assert.Equal(t, pool.MaxCapacity, poolReadInfo.MaxCapacity)
	assert.Equal(t, pool.NodeTypeID, poolReadInfo.NodeTypeID)
	assert.Equal(t, pool.IdleInstanceAutoTerminationMinutes, poolReadInfo.IdleInstanceAutoTerminationMinutes)

	u := model.InstancePoolAndStats{
		InstancePoolID:                     poolReadInfo.InstancePoolID,
		InstancePoolName:                   "Terraform Integration Test Updated",
		MinIdleInstances:                   0,
		MaxCapacity:                        20,
		NodeTypeID:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	}
	if !client.IsAzure() {
		u.DiskSpec = &model.InstancePoolDiskSpec{
			DiskType: &model.InstancePoolDiskType{
				EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		}
		u.AwsAttributes = &model.InstancePoolAwsAttributes{
			Availability: model.AwsAvailabilitySpot,
		}
	}
	err = client.InstancePools().Update(u)
	assert.NoError(t, err, err)

	poolReadInfo, err = client.InstancePools().Read(poolInfo.InstancePoolID)
	assert.NoError(t, err, err)
	assert.Equal(t, poolReadInfo.MaxCapacity, int32(20))
}
