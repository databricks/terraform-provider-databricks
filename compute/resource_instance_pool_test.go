package compute

import (
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

//go:generate easytags $GOFILE

func TestInstancePoolsAPI_Create(t *testing.T) {
	type args struct {
		InstancePool *InstancePool `json:"instance_pool"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     InstancePoolAndStats
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "instance_pool_id": "0101-120000-brick1-pool-ABCD1234"
						}`,
			args: args{
				InstancePool: &InstancePool{
					InstancePoolName:                   "",
					MinIdleInstances:                   0,
					MaxCapacity:                        10,
					NodeTypeID:                         "Standard_DS3_v2",
					IdleInstanceAutoTerminationMinutes: 60,
					EnableElasticDisk:                  false,
					DiskSpec: &InstancePoolDiskSpec{
						DiskType: &InstancePoolDiskType{
							AzureDiskVolumeType: "",
						},
						DiskCount: 1,
						DiskSize:  10,
					},
				},
			},
			want: InstancePoolAndStats{
				InstancePoolID: "0101-120000-brick1-pool-ABCD1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input InstancePool
			qa.AssertRequestWithMockServer(t, tt.args.InstancePool, http.MethodPost, "/api/2.0/instance-pools/create", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewInstancePoolsAPI(&client).Create(*tt.args.InstancePool)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-pools/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewInstancePoolsAPI(&client).Delete(tt.args.InstancePoolID)
			})
		})
	}
}

func TestInstancePoolsAPI_Update(t *testing.T) {
	type args struct {
		InstancePoolInfo *InstancePoolAndStats `json:"instance_pool_info"`
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
				InstancePoolInfo: &InstancePoolAndStats{
					InstancePoolID:                     "0101-120000-brick1-pool-ABCD1234",
					MinIdleInstances:                   0,
					MaxCapacity:                        10,
					NodeTypeID:                         "Standard_DS3_v2",
					IdleInstanceAutoTerminationMinutes: 60,
					EnableElasticDisk:                  false,
					DiskSpec: &InstancePoolDiskSpec{
						DiskType: &InstancePoolDiskType{
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
			var input InstancePoolAndStats
			qa.AssertRequestWithMockServer(t, tt.args.InstancePoolInfo, http.MethodPost, "/api/2.0/instance-pools/edit", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewInstancePoolsAPI(&client).Update(*tt.args.InstancePoolInfo)
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
		want     InstancePoolAndStats
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
			want: InstancePoolAndStats{
				InstancePoolID:   "101-120000-brick1-pool-ABCD1234",
				InstancePoolName: "mypool",
				MinIdleInstances: 0,
				AwsAttributes: &InstancePoolAwsAttributes{
					Availability:        AwsAvailabilitySpot,
					ZoneID:              "us-west-2a",
					SpotBidPricePercent: 100,
				},
				NodeTypeID:                         "c4.2xlarge",
				IdleInstanceAutoTerminationMinutes: 60,
				EnableElasticDisk:                  false,
				DiskSpec: &InstancePoolDiskSpec{
					DiskType: &InstancePoolDiskType{
						EbsVolumeType: EbsVolumeTypeGeneralPurposeSsd,
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
				Stats: &InstancePoolStats{
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
			var input InstancePoolAndStats
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/instance-pools/get?instance_pool_id=101-120000-brick1-pool-ABCD1234", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewInstancePoolsAPI(&client).Read(tt.args.InstancePoolID)
			})
		})
	}
}

func TestAccInstancePools(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	pool := InstancePool{
		InstancePoolName:                   "Terraform Integration Test",
		MinIdleInstances:                   0,
		NodeTypeID:                         qa.GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"7.1.x-scala2.12",
		},
	}
	if !client.IsUsingAzureAuth() {
		pool.DiskSpec = &InstancePoolDiskSpec{
			DiskType: &InstancePoolDiskType{
				EbsVolumeType: EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		}
		pool.AwsAttributes = &InstancePoolAwsAttributes{
			Availability: AwsAvailabilitySpot,
		}
	}
	poolInfo, err := NewInstancePoolsAPI(client).Create(pool)
	assert.NoError(t, err, err)

	defer func() {
		err := NewInstancePoolsAPI(client).Delete(poolInfo.InstancePoolID)
		assert.NoError(t, err, err)
	}()

	poolReadInfo, err := NewInstancePoolsAPI(client).Read(poolInfo.InstancePoolID)
	assert.NoError(t, err, err)
	assert.Equal(t, poolInfo.InstancePoolID, poolReadInfo.InstancePoolID)
	assert.Equal(t, pool.InstancePoolName, poolReadInfo.InstancePoolName)
	assert.Equal(t, pool.MinIdleInstances, poolReadInfo.MinIdleInstances)
	assert.Equal(t, pool.MaxCapacity, poolReadInfo.MaxCapacity)
	assert.Equal(t, pool.NodeTypeID, poolReadInfo.NodeTypeID)
	assert.Equal(t, pool.IdleInstanceAutoTerminationMinutes, poolReadInfo.IdleInstanceAutoTerminationMinutes)

	u := InstancePoolAndStats{
		InstancePoolID:                     poolReadInfo.InstancePoolID,
		InstancePoolName:                   "Terraform Integration Test Updated",
		MinIdleInstances:                   0,
		MaxCapacity:                        20,
		NodeTypeID:                         qa.GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"7.1.x-scala2.12",
		},
	}
	if !client.IsUsingAzureAuth() {
		u.DiskSpec = &InstancePoolDiskSpec{
			DiskType: &InstancePoolDiskType{
				EbsVolumeType: EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		}
		u.AwsAttributes = &InstancePoolAwsAttributes{
			Availability: AwsAvailabilitySpot,
		}
	}
	err = NewInstancePoolsAPI(client).Update(u)
	assert.NoError(t, err, err)

	poolReadInfo, err = NewInstancePoolsAPI(client).Read(poolInfo.InstancePoolID)
	assert.NoError(t, err, err)
	assert.Equal(t, poolReadInfo.MaxCapacity, int32(20))
}

func TestResourceInstancePoolCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/create",
				ExpectedRequest: InstancePool{
					InstancePoolName:                   "Shared Pool",
					MinIdleInstances:                   10,
					MaxCapacity:                        1000,
					NodeTypeID:                         "i3.xlarge",
					IdleInstanceAutoTerminationMinutes: 15,
					EnableElasticDisk:                  true,
				},
				Response: InstancePoolAndStats{
					InstancePoolID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
				Response: InstancePoolAndStats{
					InstancePoolID:                     "abc",
					InstancePoolName:                   "Shared Pool",
					MinIdleInstances:                   10,
					MaxCapacity:                        1000,
					NodeTypeID:                         "i3.xlarge",
					IdleInstanceAutoTerminationMinutes: 15,
					EnableElasticDisk:                  true,
				},
			},
		},
		Resource: ResourceInstancePool(),
		State: map[string]interface{}{
			"idle_instance_autotermination_minutes": 15,
			"instance_pool_name":                    "Shared Pool",
			"max_capacity":                          1000,
			"min_idle_instances":                    10,
			"node_type_id":                          "i3.xlarge",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstancePool(),
		State: map[string]interface{}{
			"idle_instance_autotermination_minutes": 15,
			"instance_pool_name":                    "Shared Pool",
			"max_capacity":                          1000,
			"min_idle_instances":                    10,
			"node_type_id":                          "i3.xlarge",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstancePoolRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
				Response: InstancePoolAndStats{
					InstancePoolID:                     "abc",
					InstancePoolName:                   "Shared Pool",
					MinIdleInstances:                   10,
					MaxCapacity:                        1000,
					NodeTypeID:                         "i3.xlarge",
					IdleInstanceAutoTerminationMinutes: 15,
					EnableElasticDisk:                  true,
				},
			},
		},
		Resource: ResourceInstancePool(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, 15, d.Get("idle_instance_autotermination_minutes"))
	assert.Equal(t, "Shared Pool", d.Get("instance_pool_name"))
	assert.Equal(t, 1000, d.Get("max_capacity"))
	assert.Equal(t, 10, d.Get("min_idle_instances"))
	assert.Equal(t, "i3.xlarge", d.Get("node_type_id"))
}

func TestResourceInstancePoolRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceInstancePool(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceInstancePoolRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstancePool(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceInstancePoolUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/edit",
				ExpectedRequest: InstancePoolAndStats{
					InstancePoolID:                     "abc",
					MaxCapacity:                        500,
					NodeTypeID:                         "i3.xlarge",
					IdleInstanceAutoTerminationMinutes: 20,
					InstancePoolName:                   "Restricted Pool",
					MinIdleInstances:                   5,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
				Response: InstancePoolAndStats{
					InstancePoolID:                     "abc",
					MaxCapacity:                        500,
					NodeTypeID:                         "i3.xlarge",
					IdleInstanceAutoTerminationMinutes: 20,
					InstancePoolName:                   "Restricted Pool",
					MinIdleInstances:                   5,
				},
			},
		},
		Resource: ResourceInstancePool(),
		State: map[string]interface{}{
			"idle_instance_autotermination_minutes": 20,
			"instance_pool_name":                    "Restricted Pool",
			"max_capacity":                          500,
			"min_idle_instances":                    5,
			"node_type_id":                          "i3.xlarge",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}
func TestResourceInstancePoolUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/edit",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstancePool(),
		State: map[string]interface{}{
			"idle_instance_autotermination_minutes": 20,
			"instance_pool_name":                    "Restricted Pool",
			"max_capacity":                          500,
			"min_idle_instances":                    5,
			"node_type_id":                          "i3.xlarge",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/delete",
				ExpectedRequest: map[string]string{
					"instance_pool_id": "abc",
				},
			},
		},
		Resource: ResourceInstancePool(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-pools/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstancePool(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
