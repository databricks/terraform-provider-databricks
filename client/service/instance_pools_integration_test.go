package service

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestInstancePools(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()

	pool := model.InstancePool{
		InstancePoolName: "my_instance_pool",
		MinIdleInstances: 0,
		MaxCapacity:      10,
		DiskSpec: &model.InstancePoolDiskSpec{
			DiskType: &model.InstancePoolDiskType{
				EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		},
		NodeTypeID:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
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

	err = client.InstancePools().Update(model.InstancePoolInfo{
		InstancePoolID:   poolReadInfo.InstancePoolID,
		InstancePoolName: "my_instance_pool",
		MinIdleInstances: 0,
		MaxCapacity:      20,
		DiskSpec: &model.InstancePoolDiskSpec{
			DiskType: &model.InstancePoolDiskType{
				EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
			},
			DiskCount: 1,
			DiskSize:  32,
		},
		NodeTypeID:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	})
	assert.NoError(t, err, err)

	poolReadInfo, err = client.InstancePools().Read(poolInfo.InstancePoolID)
	assert.NoError(t, err, err)
	assert.Equal(t, poolReadInfo.MaxCapacity, int32(20))
}
