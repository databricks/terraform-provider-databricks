package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/mlflow/api"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func e() api.Experiment {
	return api.Experiment{
		Name: "xyz",
	}
}

func TestMLFlowExperimentCreate(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/experiments/create",
				ExpectedRequest: e(),
				Response:        re,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					Experiment: re,
				},
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, re.ExperimentId, d.Id(), "Resource ID should not be empty")
	assert.Equal(t, re.Name, d.Get("name"), "Experiment name should be set")
}
func TestMLFlowExperimentRead(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					Experiment: re,
				},
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Read:     true,
		ID:       re.ExperimentId,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, re.ExperimentId, d.Id(), "Resource ID should not be empty")
}

func TestMLFlowExperimentUpdate(t *testing.T) {
	resPost := api.Experiment{
		ExperimentId: "123456790123456",
		Name:         "123",
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					Experiment: resPost,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/update",
				Response: resPost,
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Update:   true,
		ID:       resPost.ExperimentId,
		HCL: `
		name = "123"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, resPost.ExperimentId, d.Id(), "Resource ID should not be empty")
	assert.Equal(t, resPost.Name, d.Get("name"), "Name should be updated")
}

func TestMLFlowExperimentDelete(t *testing.T) {
	r := api.Experiment{
		ExperimentId: "123456790123456",
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/experiments/delete",
				ExpectedRequest: r,
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Delete:   true,
		ID:       r.ExperimentId,
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, r.ExperimentId, d.Id(), "Resource ID should not be empty")
}
