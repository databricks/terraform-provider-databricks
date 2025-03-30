package mlflow

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func e() ml.Experiment {
	return ml.Experiment{
		Name: "xyz",
	}
}

func TestExperimentCreate(t *testing.T) {
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
				Response: ml.GetExperimentResponse{
					Experiment: &re,
				},
			},
		},
		Resource: ResourceMlflowExperiment(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, re.ExperimentId, d.Id(), "Resource ID should not be empty")
	assert.Equal(t, re.Name, d.Get("name"), "Experiment name should be set")
}

func TestExperimentCreatePostError(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/experiments/create",
				ExpectedRequest: e(),
				Response:        re,
				Status:          400,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestExperimentCreateGetError(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	_, err := qa.ResourceFixture{
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
				Response: ml.GetExperimentResponse{
					Experiment: &re,
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestExperimentRead(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: ml.GetExperimentResponse{
					Experiment: &re,
				},
			},
		},
		Resource: ResourceMlflowExperiment(),
		Read:     true,
		ID:       re.ExperimentId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, re.ExperimentId, d.Id(), "Resource ID should not be empty")
}

func TestExperimentReadGetError(t *testing.T) {
	re := e()
	re.ExperimentId = "123456790123456"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: ml.GetExperimentResponse{
					Experiment: &re,
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Read:     true,
		ID:       re.ExperimentId,
	}.Apply(t)

	assert.Error(t, err)
}

func TestExperimentUpdate(t *testing.T) {
	resPost := ml.Experiment{
		ExperimentId: "123456790123456",
		Name:         "123",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=123456790123456",
				Response: ml.GetExperimentResponse{
					Experiment: &resPost,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/update",
				Response: resPost,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Update:   true,
		ID:       resPost.ExperimentId,
		HCL: `
		name = "123"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name": resPost.Name,
		"id":   resPost.ExperimentId,
	})
}

func TestExperimentUpdatePostError(t *testing.T) {
	resPost := ml.Experiment{
		ExperimentId: "123456790123456",
		Name:         "123",
	}
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/experiments/update",
				Response: resPost,
				Status:   400,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Update:   true,
		ID:       resPost.ExperimentId,
		HCL: `
		name = "123"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestExperimentDelete(t *testing.T) {
	r := map[string]string{
		"experiment_id": "123456790123456",
	}
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/experiments/delete",
				ExpectedRequest: r,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Delete:   true,
		ID:       r["experiment_id"],
		HCL: `
		name = "xyz"
		`,
	}.ApplyNoError(t)
}

func TestExperimentDeleteError(t *testing.T) {
	r := map[string]string{
		"experiment_id": "123456790123456",
	}
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/experiments/delete",
				ExpectedRequest: r,
				Status:          400,
			},
		},
		Resource: ResourceMlflowExperiment(),
		Delete:   true,
		ID:       r["experiment_id"],
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestExperimentNameSuppressDiff(t *testing.T) {
	assert.True(t, experimentNameSuppressDiff("", "/Workspace/EXPERIMENT/", "/EXPERIMENT", nil))
	assert.True(t, experimentNameSuppressDiff("", "/Workspace/EXPERIMENT", "/EXPERIMENT", nil))
	assert.True(t, experimentNameSuppressDiff("", "/EXPERIMENT", "/EXPERIMENT", nil))
	assert.False(t, experimentNameSuppressDiff("", "/new_name", "/EXPERIMENT/", nil))
}
