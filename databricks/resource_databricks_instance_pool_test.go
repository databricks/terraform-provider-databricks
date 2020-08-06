package databricks

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/stretchr/testify/assert"
)

func TestResourceInstancePoolCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/create",
			ExpectedRequest: model.InstancePool{
				InstancePoolName:                   "Shared Pool",
				MinIdleInstances:                   10,
				MaxCapacity:                        1000,
				NodeTypeID:                         "i3.xlarge",
				IdleInstanceAutoTerminationMinutes: 15,
				EnableElasticDisk:                  true,
			},
			Response: model.InstancePoolAndStats{
				InstancePoolID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
			Response: model.InstancePoolAndStats{
				InstancePoolID:                     "abc",
				InstancePoolName:                   "Shared Pool",
				MinIdleInstances:                   10,
				MaxCapacity:                        1000,
				NodeTypeID:                         "i3.xlarge",
				IdleInstanceAutoTerminationMinutes: 15,
				EnableElasticDisk:                  true,
			},
		},
	}, resourceInstancePool, map[string]interface{}{
		"idle_instance_autotermination_minutes": 15,
		"instance_pool_name":                    "Shared Pool",
		"max_capacity":                          1000,
		"min_idle_instances":                    10,
		"node_type_id":                          "i3.xlarge",
	}, resourceInstancePoolCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/create",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstancePool, map[string]interface{}{
		"idle_instance_autotermination_minutes": 15,
		"instance_pool_name":                    "Shared Pool",
		"max_capacity":                          1000,
		"min_idle_instances":                    10,
		"node_type_id":                          "i3.xlarge",
	}, resourceInstancePoolCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstancePoolRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
			Response: model.InstancePoolAndStats{
				InstancePoolID:                     "abc",
				InstancePoolName:                   "Shared Pool",
				MinIdleInstances:                   10,
				MaxCapacity:                        1000,
				NodeTypeID:                         "i3.xlarge",
				IdleInstanceAutoTerminationMinutes: 15,
				EnableElasticDisk:                  true,
			},
		},
	}, resourceInstancePool, nil, actionWithID("abc", resourceInstancePoolRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, 15, d.Get("idle_instance_autotermination_minutes"))
	assert.Equal(t, "Shared Pool", d.Get("instance_pool_name"))
	assert.Equal(t, 1000, d.Get("max_capacity"))
	assert.Equal(t, 10, d.Get("min_idle_instances"))
	assert.Equal(t, "i3.xlarge", d.Get("node_type_id"))
}

func TestResourceInstancePoolRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceInstancePool, nil, actionWithID("abc", resourceInstancePoolRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceInstancePoolRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/get?instance_pool_id=abc",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstancePool, nil, actionWithID("abc", resourceInstancePoolRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceInstancePoolUpdate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/edit",
			ExpectedRequest: model.InstancePoolAndStats{
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
			Response: model.InstancePoolAndStats{
				InstancePoolID:                     "abc",
				MaxCapacity:                        500,
				NodeTypeID:                         "i3.xlarge",
				IdleInstanceAutoTerminationMinutes: 20,
				InstancePoolName:                   "Restricted Pool",
				MinIdleInstances:                   5,
			},
		},
	}, resourceInstancePool, map[string]interface{}{
		"idle_instance_autotermination_minutes": 20,
		"instance_pool_name":                    "Restricted Pool",
		"max_capacity":                          500,
		"min_idle_instances":                    5,
		"node_type_id":                          "i3.xlarge",
	}, actionWithID("abc", resourceInstancePoolUpdate))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}
func TestResourceInstancePoolUpdate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/edit",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstancePool, map[string]interface{}{
		"idle_instance_autotermination_minutes": 20,
		"instance_pool_name":                    "Restricted Pool",
		"max_capacity":                          500,
		"min_idle_instances":                    5,
		"node_type_id":                          "i3.xlarge",
	}, actionWithID("abc", resourceInstancePoolUpdate))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/delete",
			ExpectedRequest: map[string]string{
				"instance_pool_id": "abc",
			},
		},
	}, resourceInstancePool, nil, actionWithID("abc", resourceInstancePoolDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstancePoolDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-pools/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstancePool, nil, actionWithID("abc", resourceInstancePoolDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
