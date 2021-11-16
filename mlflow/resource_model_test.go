package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMLFlowModelCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: MLFLowModelAPI{
					Name: "xyz",
					Tags: []Tag{
						{"key1", "value1"},
						{"key2", "value2"},
					},
				},
				Response: MLFlowRegisteredModelAPI{
					MLFLowModelAPI{
						Name: "xyz",
						Tags: []Tag{
							{"key1", "value1"},
							{"key2", "value2"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/get?name=xyz",
				Response: MLFlowRegisteredModelAPI{
					MLFLowModelAPI{
						Name: "xyz",
						Tags: []Tag{
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
	assert.Equal(t, "xyz", d.Get("name"))
}

func TestMLFlowModelRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/get?name=xyz",
				Response: MLFlowRegisteredModelAPI{
					MLFLowModelAPI{
						Name: "xyz",
						Tags: []Tag{
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

// func TestDashboardUpdate(t *testing.T) {
// 	d, err := qa.ResourceFixture{
// 		Fixtures: []qa.HTTPFixture{
// 			{
// 				Method:   "POST",
// 				Resource: "/api/2.0/preview/sql/dashboards/xyz",
// 				Response: api.Dashboard{
// 					ID:   "xyz",
// 					Name: "Dashboard renamed",
// 					Tags: []string{"t2", "t3"},
// 				},
// 			},
// 			{
// 				Method:   "GET",
// 				Resource: "/api/2.0/preview/sql/dashboards/xyz",
// 				Response: api.Dashboard{
// 					ID:   "xyz",
// 					Name: "Dashboard renamed",
// 					Tags: []string{"t2", "t3"},
// 				},
// 			},
// 		},
// 		Resource: ResourceDashboard(),
// 		Update:   true,
// 		ID:       "xyz",
// 		State: map[string]interface{}{
// 			"name": "Dashboard renamed",
// 			"tags": []interface{}{"t2", "t3"},
// 		},
// 	}.Apply(t)

// 	assert.NoError(t, err, err)
// 	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
// }

func TestDashboardDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete",
				ExpectedRequest: MLFLowModelAPI{
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
