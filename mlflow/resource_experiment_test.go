package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/mlflow/api"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMLFlowExperimentCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/create",
				ExpectedRequest: api.Experiment{
					Name: "xyz",
					Tags: []api.Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
				Response: api.Experiment{
					ExperimentId: "123456790123456",
					Name:         "xyz",
					Tags: []api.Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					api.Experiment{
						ExperimentId: "123456790123456",
						Name:         "xyz",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Create:   true,
		HCL: `
		name = "xyz"
		tags {
				key = "key1"
				value = "value1"
		    }
		tags {
				key = "key2"
				value = "value2"
			  }
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "123456790123456", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "xyz", d.Get("name"), "Experiment name should be set")
}
func TestMLFlowExperimentRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					api.Experiment{
						ExperimentId: "123456790123456",
						Name:         "xyz",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Read:     true,
		ID:       "123456790123456",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "123456790123456", d.Id(), "Resource ID should not be empty")
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
					api.Experiment{
						ExperimentId: resPost.ExperimentId,
						Name:         resPost.Name,
					},
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

func TestMLFlowExperimentUpdateTagsNeedsRecreate(t *testing.T) {
	resPost := api.Experiment{
		ExperimentId: "123456790123456",
		Name:         "123",
		Tags: []api.Tag{
			{"key1", "value1"},
		},
	}
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: api.Experiments{
					api.Experiment{
						ExperimentId: resPost.ExperimentId,
						Name:         resPost.Name,
						Tags: []api.Tag{
							{"key1", "value1"},
						},
					},
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
		tags {
			key = "key2"
			value = "value2"
		}
		`,
	}.Apply(t)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "changes require new: tags")
}

func TestMLFlowExperimentDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/delete",
				ExpectedRequest: api.Experiment{
					ExperimentId: "123456790123456",
					Name:         "xyz",
				},
			},
		},
		Resource: ResourceMLFlowExperiment(),
		Delete:   true,
		ID:       "123456790123456",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "123456790123456", d.Id(), "Resource ID should not be empty")
}
