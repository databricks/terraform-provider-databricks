package service

import (
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstancePools(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()

	pool := model.InstancePool{
		InstancePoolName:                   "my_instance_pool",
		MinIdleInstances:                   0,
		MaxCapacity:                        10,
		NodeTypeId:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	}
	poolInfo, err := client.InstancePools().Create(pool)
	assert.NoError(t, err, err)

	defer func() {
		err := client.InstancePools().Delete(poolInfo.InstancePoolId)
		assert.NoError(t, err, err)
	}()

	poolReadInfo, err := client.InstancePools().Read(poolInfo.InstancePoolId)
	assert.NoError(t, err, err)
	assert.Equal(t, poolInfo.InstancePoolId, poolReadInfo.InstancePoolId)
	assert.Equal(t, pool.InstancePoolName, poolReadInfo.InstancePoolName)
	assert.Equal(t, pool.MinIdleInstances, poolReadInfo.MinIdleInstances)
	assert.Equal(t, pool.MaxCapacity, poolReadInfo.MaxCapacity)
	assert.Equal(t, pool.NodeTypeId, poolReadInfo.NodeTypeId)
	assert.Equal(t, pool.IdleInstanceAutoTerminationMinutes, poolReadInfo.IdleInstanceAutoTerminationMinutes)

	err = client.InstancePools().Update(model.InstancePoolInfo{
		InstancePoolId:                     poolReadInfo.InstancePoolId,
		InstancePoolName:                   "my_instance_pool",
		MinIdleInstances:                   0,
		MaxCapacity:                        20,
		NodeTypeId:                         GetCloudInstanceType(client),
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	})
	assert.NoError(t, err, err)

	poolReadInfo, err = client.InstancePools().Read(poolInfo.InstancePoolId)
	assert.NoError(t, err, err)
	assert.Equal(t, poolReadInfo.MaxCapacity, int32(20))

}
