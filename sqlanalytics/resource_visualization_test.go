package sqlanalytics

import (
	"encoding/json"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
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
					ID:          12345,
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
		Resource: ResourceVisualization(),
		Create:   true,
		State: map[string]interface{}{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
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
		Resource: ResourceVisualization(),
		Read:     true,
		ID:       "12345",
		State: map[string]interface{}{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
	assert.Equal(t, "chart", d.Get("type"))
	assert.Equal(t, "My Chart", d.Get("name"))
	assert.Equal(t, "Some Description", d.Get("description"))
	assert.Less(t, 0, len(d.Get("options").(string)))
}

func TestVisualizationUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/visualizations/12345",
				ExpectedRequest: api.Visualization{
					ID:          12345,
					QueryID:     "foo",
					Type:        "CHART",
					Name:        "My Updated Chart",
					Description: "Some Updated Description",
					Options:     json.RawMessage("{}"),
				},
				Response: api.Visualization{
					// Note: "query_id" is not included in POST response.
					ID:          12345,
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
		Resource: ResourceVisualization(),
		Update:   true,
		ID:       "12345",
		InstanceState: map[string]string{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Chart",
			"description": "Some Description",
			"options":     "{}",
		},
		State: map[string]interface{}{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Updated Chart",
			"description": "Some Updated Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "12345", d.Id())
	assert.Equal(t, "foo", d.Get("query_id"))
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
		Resource: ResourceVisualization(),
		Delete:   true,
		ID:       "12345",
		State: map[string]interface{}{
			"query_id":    "foo",
			"type":        "chart",
			"name":        "My Updated Chart",
			"description": "Some Updated Description",
			"options":     "{}",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	// Delete doesn't touch schema.ResourceData, so the ID should survive.
	assert.Equal(t, "12345", d.Id())
}
