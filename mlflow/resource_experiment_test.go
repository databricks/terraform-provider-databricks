package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMLFlowExperimentCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/create",
				ExpectedRequest: MLFLowExperimentAPI{
					Name: "xyz",
					Tags: []Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
				Response: MLFLowExperimentAPI{
					ExperimentId: "123456790123456",
					Name:         "xyz",
					Tags: []Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: MLFLowExperimentsAPI{
					MLFLowExperimentAPI{
						ExperimentId: "123456790123456",
						Name:         "xyz",
						Tags: []Tag{
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
				Response: MLFLowExperimentsAPI{
					MLFLowExperimentAPI{
						ExperimentId: "123456790123456",
						Name:         "xyz",
						Tags: []Tag{
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

func TestMLFlowExperimentDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/delete",
				ExpectedRequest: MLFLowExperimentAPI{
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
