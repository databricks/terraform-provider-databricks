package compute

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceClusterPolicyRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: ClusterPolicy{
					PolicyID:           "abc",
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimeStamp: 0,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "Dummy", d.Get("name"))
	assert.Equal(t, "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}", d.Get("definition"))
	assert.Equal(t, "abc", d.Get("policy_id"))
}

func TestResourceClusterPolicyRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceClusterPolicyRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceClusterPolicyCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/create",
				ExpectedRequest: ClusterPolicy{
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimeStamp: 0,
				},
				Response: ClusterPolicy{
					PolicyID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: ClusterPolicy{
					PolicyID:           "abc",
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimeStamp: 0,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]interface{}{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]interface{}{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestResourceClusterPolicyUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				ExpectedRequest: ClusterPolicy{
					PolicyID:           "abc",
					Name:               "Dummy Updated",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimeStamp: 0,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: ClusterPolicy{
					PolicyID:           "abc",
					Name:               "Dummy Updated",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimeStamp: 0,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]interface{}{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy Updated",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]interface{}{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy Updated",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/delete",
				ExpectedRequest: map[string]string{
					"policy_id": "abc",
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
