package sql

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestWidgetCreateWithVisualization(t *testing.T) {
	i678 := api.NewStringOrInt("678")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
					Options: api.WidgetOptions{
						Title:       "title",
						Description: "description",
					},
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
					Options: api.WidgetOptions{
						Title:       "title",
						Description: "description",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": 678,
								"options": {
									"title": "title",
									"description": "description"
								}
							}
						`),
						json.RawMessage(`
							{
								"id": "12346",
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id     = "some-uuid"
			visualization_id = "678"

			title       = "title"
			description = "description"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))
	assert.Equal(t, "title", d.Get("title"))
	assert.Equal(t, "description", d.Get("description"))
}

func TestWidgetCreateWithVisualizationByResourceID(t *testing.T) {
	i678 := api.NewStringOrInt("678")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344"
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization": {
									"id": 678
								}
							}
						`),
						json.RawMessage(`
							{
								"id": "12346"
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id     = "some-uuid"
			visualization_id = "doesnt-matter/678"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))
}

func TestWidgetCreateWithText(t *testing.T) {
	sText := "widget text"

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID: "some-uuid",
					Text:        &sText,
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"text": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"text": "` + sText + `"
							}
						`),
						json.RawMessage(`
							{
								"id": "12346",
								"text": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id = "some-uuid"
			text         = "` + sText + `"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, sText, d.Get("text"))
}

func TestWidgetCreateWithParamValue(t *testing.T) {
	i678 := api.NewStringOrInt("678")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
					Options: api.WidgetOptions{
						ParameterMapping: map[string]api.WidgetParameterMapping{
							"p1": {
								Name:  "p1",
								Type:  "dashboard-level",
								Value: "v1",
							},
							"p2": {
								Name:  "p2",
								Type:  "dashboard-level",
								Value: []string{"v2_0", "v2_1"},
							},
						},
					},
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": 678,
								"options": {
									"parameterMappings": {
										"p2": {
											"name": "p2",
											"type": "dashboard-level",
											"value": [
												"v2_0",
												"v2_1"
											]
										},
										"p1": {
											"name": "p1",
											"type": "dashboard-level",
											"value": "v1"
										}
									}
								}
							}
						`),
						json.RawMessage(`
							{
								"id": "12346",
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id = "some-uuid"
			visualization_id = "678"

			parameter {
				name  = "p1"
				type  = "dashboard-level"
				value = "v1"
			}

			parameter {
				name  = "p2"
				type  = "dashboard-level"
				values = ["v2_0", "v2_1"]
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))

	params := d.Get("parameter").(*schema.Set)
	assert.Equal(t, 2, params.Len())

	for _, param := range params.List() {
		m := param.(map[string]any)
		switch m["name"].(string) {
		case "p1":
			// First parameter
			assert.Equal(t, "dashboard-level", m["type"])
			assert.Equal(t, "v1", m["value"])
			values := m["values"].([]any)
			assert.Equal(t, 0, len(values))
		case "p2":
			// Second parameter
			assert.Equal(t, "dashboard-level", m["type"])
			assert.Equal(t, "", m["value"])
			values := m["values"].([]any)
			assert.Equal(t, 2, len(values))
			assert.Equal(t, "v2_0", values[0].(string))
			assert.Equal(t, "v2_1", values[1].(string))
		default:
			t.Fatalf("Unexpected parameter: %v", m["name"])
		}
	}
}

func TestWidgetCreateWithPosition(t *testing.T) {
	i678 := api.NewStringOrInt("678")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
					Options: api.WidgetOptions{
						Position: &api.WidgetPosition{
							AutoHeight: false,
							SizeX:      3,
							SizeY:      4,
							PosX:       5,
							PosY:       6,
						},
					},
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": 678,
								"options": {
									"position": {
										"autoHeight": false,
										"sizeX": 3,
										"sizeY": 4,
										"col": 5,
										"row": 6
									}
								}
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id = "some-uuid"
			visualization_id = "678"

			position {
				size_x = 3
				size_y = 4
				pos_x = 5
				pos_y = 6
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))

	// The position is a nested type, which can only be modeled in Terraform
	// by a list type with a single element.
	assert.Equal(t, 3, d.Get("position.0.size_x"))
	assert.Equal(t, 4, d.Get("position.0.size_y"))
	assert.Equal(t, 5, d.Get("position.0.pos_x"))
	assert.Equal(t, 6, d.Get("position.0.pos_y"))
	assert.Equal(t, false, d.Get("position.0.auto_height"))
}

func TestWidgetCreateWithPositionAndAutoheight(t *testing.T) {
	i678 := api.NewStringOrInt("678")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets",
				ExpectedRequest: api.Widget{
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
					Options: api.WidgetOptions{
						Position: &api.WidgetPosition{
							AutoHeight: true,
							SizeX:      3,
							SizeY:      0,
							PosX:       5,
							PosY:       6,
						},
					},
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": 678,
								"options": {
									"position": {
										"autoHeight": true,
										"sizeX": 3,
										"sizeY": 0,
										"col": 5,
										"row": 6
									}
								}
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Create:   true,
		HCL: `
			dashboard_id = "some-uuid"
			visualization_id = "678"

			position {
				size_x = 3
				size_y = 0
				pos_x = 5
				pos_y = 6

				auto_height = true
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))

	// The position is a nested type, which can only be modeled in Terraform
	// by a list type with a single element.
	assert.Equal(t, 3, d.Get("position.0.size_x"))
	assert.Equal(t, 0, d.Get("position.0.size_y"))
	assert.Equal(t, 5, d.Get("position.0.pos_x"))
	assert.Equal(t, 6, d.Get("position.0.pos_y"))
	assert.Equal(t, true, d.Get("position.0.auto_height"))
}

func TestWidgetReadNotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12345",
								"text": "text"
							}
						`),
					},
				},
			},
		},
		Resource:    ResourceSqlWidget(),
		Read:        true,
		Removed:     true,
		RequiresNew: true,
		ID:          "some-uuid/12344",
		InstanceState: map[string]string{
			"dashboard_id":     "some-uuid",
			"visualization_id": "678",
		},
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "", d.Id(), "Resource ID should be empty")
}

func TestWidgetUpdate(t *testing.T) {
	sText := "new text"

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/widgets/12345",
				ExpectedRequest: api.Widget{
					DashboardID: "some-uuid",
					Text:        &sText,
				},
				Response: api.Widget{
					ID:          "12345",
					DashboardID: "some-uuid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/some-uuid",
				Response: api.Dashboard{
					ID: "some-uuid",
					Widgets: []json.RawMessage{
						json.RawMessage(`
							{
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"text": "new text"
							}
						`),
						json.RawMessage(`
							{
								"id": "12346",
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceSqlWidget(),
		Update:   true,
		ID:       "some-uuid/12345",
		InstanceState: map[string]string{
			"dashboard_id": "some-uuid",
			"text":         "previous text",
		},
		HCL: `
			dashboard_id = "some-uuid"
			text         = "new text"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))

	// Test that text is now set.
	{
		v, ok := d.GetOk("text")
		assert.True(t, ok, "text should be set")
		assert.Equal(t, "new text", v)
	}
}

func TestWidgetDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/widgets/12345",
			},
		},
		Resource: ResourceSqlWidget(),
		Delete:   true,
		ID:       "some-uuid/12345",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
}

func TestResourceWidgetCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlWidget(), qa.CornerCaseID("foo/bar"))
}

func TestWidgetParameterSorter(t *testing.T) {
	wp := sortWidgetParameter{
		WidgetParameter{Name: "foo"},
		WidgetParameter{Name: "bar"},
	}

	// Widget parameters should be sorted by their name to maintain deterministic ordering.
	// Since they are not ordered in the API payload, not ordering them means users would
	// see false state mismatches on comparison.
	sort.Sort(wp)

	assert.Equal(t, "bar", wp[0].Name)
	assert.Equal(t, "foo", wp[1].Name)
}
