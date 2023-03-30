package mlflow

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestJobsAPIList(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/mlflow/registered-models/search?max_results=100",
			Response: ModelListResponse{
				Models: []Model{
					{
						Name:              "test",
						RegisteredModelID: "model_id",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewModelsAPI(ctx, client)
		l, err := a.List()
		require.NoError(t, err)
		assert.Len(t, l, 1)
	})
}

func TestJobsAPIListMultiplePages(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/mlflow/registered-models/search?max_results=100",
			Response: ModelListResponse{
				Models: []Model{
					{
						Name:              "test",
						RegisteredModelID: "model_id",
					},
				},
				NextPageToken: "TOKEN",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/mlflow/registered-models/search?max_results=100&page_token=TOKEN",
			Response: ModelListResponse{
				Models: []Model{
					{
						Name:              "test2",
						RegisteredModelID: "model_id2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewModelsAPI(ctx, client)
		l, err := a.List()
		require.NoError(t, err)
		assert.Len(t, l, 2)
	})
}

func TestModelsAPIListByName(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/mlflow/registered-models/search?filter=name%20ilike%20%27model%27&max_results=100",
			Response: ModelListResponse{
				Models: []Model{
					{
						Name:              "test",
						RegisteredModelID: "model_id",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewModelsAPI(ctx, client)
		l, err := a.ListByFilter("name ilike 'model'", nil, "")
		require.NoError(t, err)
		assert.Len(t, l, 1)
	})
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
				Resource: "/api/2.0/mlflow/registered-models/delete",
				ExpectedRequest: Model{
					Name: "xyz",
				},
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
				Resource: "/api/2.0/mlflow/registered-models/delete",
				ExpectedRequest: Model{
					Name: "xyz",
				},
				Status: 400,
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
