package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/mlflow/api"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMLFlowModelCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: api.Model{
					Name: "xyz",
					Tags: []api.Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
				Response: api.RegisteredModel{
					api.Model{
						Name: "xyz",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/get?name=xyz",
				Response: api.RegisteredModel{
					api.Model{
						Name: "xyz",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowModel(),
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
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "xyz", d.Get("name"), "Name should be set")
	assert.Equal(t, d.Get("name"), d.Id(), "Name and Id should match")

}

func TestMLFlowModelRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/get?name=xyz",
				Response: api.RegisteredModel{
					api.Model{
						Name: "xyz",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowModel(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestMLFlowModelUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				Response: api.RegisteredModel{
					api.Model{
						Name:        "xyz",
						Description: "thedescription",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/get?name=xyz",
				Response: api.RegisteredModel{
					api.Model{
						Name:        "xyz",
						Description: "updateddescription",
						Tags: []api.Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowModel(),
		Update:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		description = "updateddescription"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "updateddescription", d.Get("description"), "Description should be updated")

}

func TestDashboardDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete",
				ExpectedRequest: api.Model{
					Name: "xyz",
				},
			},
		},
		Resource: ResourceMLFlowModel(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}
