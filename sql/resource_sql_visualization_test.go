package sql

import (
	"encoding/json"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/stretchr/testify/assert"
)

func TestVisualizationCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/visualizations",
				ExpectedRequest: api.Visualization{
					QueryID:     "foo",
					Type:        "CHART",
					Name:        "My Chart",
					Description: "Some Description",
					Options:     json.RawMessage("{}"),
				},
				Response: api.Visualization{
					// Note: "query_id" is not included in POST response.
					ID:          "12345",
					Type:        "CHART",
					Name:        "My Chart",
					Description: "Some Description",
					Options:     json.RawMessage("{}"),
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12344
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART",
								"name": "My Chart",
								"description": "Some Description",
								"options": {}
							}
						`),
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlVisualization(),
		Create:   true,
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo/12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
	assert.Equal(t, "12345", d.Get("visualization_id"))
	assert.Equal(t, "chart", d.Get("type"))
	assert.Equal(t, "My Chart", d.Get("name"))
	assert.Equal(t, "Some Description", d.Get("description"))
	assert.Less(t, 0, len(d.Get("options").(string)))
}

func TestVisualizationRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12344
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART",
								"name": "My Chart",
								"description": "Some Description",
								"options": {}
							}
						`),
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlVisualization(),
		Read:     true,
		ID:       "foo/12345",
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo/12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
	assert.Equal(t, "12345", d.Get("visualization_id"))
	assert.Equal(t, "chart", d.Get("type"))
	assert.Equal(t, "My Chart", d.Get("name"))
	assert.Equal(t, "Some Description", d.Get("description"))
	assert.Less(t, 0, len(d.Get("options").(string)))
}

func TestVisualizationReadWithQueryPlan(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12344
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART",
								"name": "My Chart",
								"description": "Some Description",
								"options": {},
								"query_plan": {"foo":"qux"}
							}
						`),
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlVisualization(),
		Read:     true,
		ID:       "foo/12345",
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
			"query_plan":  `{"foo":"bar"}`,
		},
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, `{"foo":"qux"}`, d.Get("query_plan").(string))
}

func TestVisualizationReadWithNullQueryPlan(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12344
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART",
								"name": "My Chart",
								"description": "Some Description",
								"options": {},
								"query_plan": null
							}
						`),
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlVisualization(),
		Read:     true,
		New:      true,
		ID:       "foo/12345",
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "", d.Get("query_plan").(string))
}

func TestVisualizationReadNotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART"
							}
						`),
					},
				},
			},
		},
		Resource:    ResourceSqlVisualization(),
		Read:        true,
		Removed:     true,
		RequiresNew: true,
		ID:          "foo/1234",
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Resource ID should be empty")
}

func TestVisualizationUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/visualizations/12345",
				ExpectedRequest: api.Visualization{
					QueryID:     "foo",
					Type:        "CHART",
					Name:        "My Updated Chart",
					Description: "Some Updated Description",
					Options:     json.RawMessage("{}"),
				},
				Response: api.Visualization{
					// Note: "query_id" is not included in POST response.
					ID:          "12345",
					Type:        "CHART",
					Name:        "My Updated Chart",
					Description: "Some Updated Description",
					Options:     json.RawMessage("{}"),
				},
			},
			// This is executed AFTER the update.
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12344
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"type": "CHART",
								"name": "My Updated Chart",
								"description": "Some Updated Description",
								"options": {}
							}
						`),
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlVisualization(),
		Update:   true,
		ID:       "foo/12345",
		InstanceState: map[string]string{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
		State: map[string]any{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Updated Chart",
			"description": "Some Updated Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo/12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
	assert.Equal(t, "12345", d.Get("visualization_id"))
	assert.Equal(t, "chart", d.Get("type"))
	assert.Equal(t, "My Updated Chart", d.Get("name"))
	assert.Equal(t, "Some Updated Description", d.Get("description"))
	assert.Less(t, 0, len(d.Get("options").(string)))
}

func TestVisualizationDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/visualizations/12345",
			},
		},
		Resource: ResourceSqlVisualization(),
		Delete:   true,
		ID:       "foo/12345",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "foo/12345", d.Id(), "Resource ID should not be empty")
}

func TestResourceVisualizationCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlVisualization(), qa.CornerCaseID("foo/bar"))
}
