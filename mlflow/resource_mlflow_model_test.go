package mlflow

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func m() ml.Model {
	return ml.Model{
		Name: "xyz",
		Tags: []ml.ModelTag{
			{Key: "key1", Value: "value1"},
			{Key: "key2", Value: "value2"},
		},
	}
}

func TestModelCreateMVP(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: ml.CreateModelRequest{
					Name: "xyz",
				},
				Response: ml.CreateModelResponse{
					RegisteredModel: &ml.Model{
						Name: "xyz",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name: "xyz",
						Id:   "123",
					},
				},
				ReuseRequest: true,
			},
		},
		Resource: ResourceMlflowModel(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.ApplyAndExpectData(t, map[string]any{"id": "xyz", "registered_model_id": "123"})
}

func TestModelCreateWithTags(t *testing.T) {
	model := m()
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: ml.CreateModelRequest{
					Name: "xyz",
					Tags: []ml.ModelTag{
						{Key: "key1", Value: "value1"},
						{Key: "key2", Value: "value2"},
					},
				},
				Response: ml.CreateModelResponse{
					RegisteredModel: &model,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name: "xyz",
						Tags: []ml.ModelTag{
							{Key: "key1", Value: "value1"},
							{Key: "key2", Value: "value2"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name: "xyz",
						Tags: []ml.ModelTag{
							{Key: "key1", Value: "value1"},
							{Key: "key2", Value: "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMlflowModel(),
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

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "xyz", d.Get("name"), "Name should be set")
	assert.Equal(t, d.Get("name"), d.Id(), "Name and Id should match")
}

func TestModelCreatePostError(t *testing.T) {
	model := m()
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: ml.CreateModelRequest{
					Name: "xyz",
					Tags: []ml.ModelTag{
						{Key: "key1", Value: "value1"},
						{Key: "key2", Value: "value2"},
					},
				},
				Response: ml.CreateModelResponse{
					RegisteredModel: &model,
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowModel(),
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

	assert.Error(t, err)
}

func TestModelRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name: "xyz",
					},
				},
			},
		},
		Resource: ResourceMlflowModel(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestModelReadGetError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name: "xyz",
					},
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowModel(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.Error(t, err)
}

func TestModelUpdate(t *testing.T) {
	pm := m()
	pm.Description = "thedescription"
	newDescription := "updatedddescription"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				Response: pm,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: ml.GetModelResponse{
					RegisteredModelDatabricks: &ml.ModelDatabricks{
						Name:        "xyz",
						Description: newDescription,
					},
				},
			},
		},
		Resource:    ResourceMlflowModel(),
		Update:      true,
		RequiresNew: true,
		ID:          "xyz",
		State: map[string]any{
			"name": "xyz",
		},
		HCL: fmt.Sprintf(`
		name = "xyz"
		description = "%s"
		`, newDescription),
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, newDescription, d.Get("description"), "Description should be updated")
}

func TestModelUpdatePatchError(t *testing.T) {
	pm := m()
	pm.Description = "thedescription"
	gm := m()
	gm.Description = "updateddescription"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				ExpectedRequest: ml.UpdateModelRequest{
					Description: "updateddescription",
					Name:        "xyz",
				},
				Response: gm,
				Status:   400,
			},
		},
		Resource:    ResourceMlflowModel(),
		Update:      true,
		RequiresNew: true,
		ID:          "xyz",
		State: map[string]any{
			"name": "xyz",
		},
		HCL: `
		name = "xyz"
		description = "updateddescription"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestModelDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete?name=xyz",
				Response: nil,
				Status:   200,
			},
		},
		Resource: ResourceMlflowModel(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestModelDeleteError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete?name=xyz",
				Status:   400,
			},
		},
		Resource: ResourceMlflowModel(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.Error(t, err)
}
