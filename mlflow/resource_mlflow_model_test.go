package mlflow

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func m() Model {
	return Model{
		Name: "xyz",
		Tags: []Tag{
			{Key: "key1", Value: "value1"},
			{Key: "key2", Value: "value2"},
		},
	}
}

func TestModelCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: m(),
				Response:        m(),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: registeredModel{
					RegisteredModelDatabricks: m(),
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
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: m(),
				Response:        m(),
				Status:          400,
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
				Response: registeredModel{
					RegisteredModelDatabricks: m(),
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
				Response: registeredModel{
					RegisteredModelDatabricks: m(),
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
	gm := m()
	gm.Description = "updateddescription"
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
				Response: registeredModel{
					RegisteredModelDatabricks: gm,
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
		HCL: `
		name = "xyz"
		description = "updateddescription"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "updateddescription", d.Get("description"), "Description should be updated")
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
				Response: pm,
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
