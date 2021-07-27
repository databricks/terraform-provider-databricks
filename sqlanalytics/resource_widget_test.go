package sqlanalytics

import (
	"encoding/json"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/stretchr/testify/assert"
)

func TestWidgetCreateWithVisualization(t *testing.T) {
	i678 := 678

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
								"id": "12344",
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": "12345",
								"visualization_id": 678
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
		Resource: ResourceWidget(),
		Create:   true,
		HCL: `
			dashboard_id     = "some-uuid"
			visualization_id = "678"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))
}

func TestWidgetCreateWithVisualizationByResourceID(t *testing.T) {
	i678 := 678

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
		Resource: ResourceWidget(),
		Create:   true,
		HCL: `
			dashboard_id     = "some-uuid"
			visualization_id = "doesnt-matter/678"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
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
		Resource: ResourceWidget(),
		Create:   true,
		HCL: `
			dashboard_id = "some-uuid"
			text         = "` + sText + `"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, sText, d.Get("text"))
}

func TestWidgetCreateWithParamValue(t *testing.T) {
	i678 := 678

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
		Resource: ResourceWidget(),
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

	assert.NoError(t, err, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))
	assert.Equal(t, "678", d.Get("visualization_id"))

	// First parameter
	assert.Equal(t, "p1", d.Get("parameter.0.name"))
	assert.Equal(t, "dashboard-level", d.Get("parameter.0.type"))
	assert.Equal(t, "v1", d.Get("parameter.0.value"))
	{
		_, ok := d.GetOk("parameter.0.values")
		assert.False(t, ok, "Expected `values` to not be set")
	}

	// Second parameter
	assert.Equal(t, "p2", d.Get("parameter.1.name"))
	assert.Equal(t, "dashboard-level", d.Get("parameter.1.type"))
	assert.Equal(t, 2, d.Get("parameter.1.values.#"))
	assert.Equal(t, "v2_0", d.Get("parameter.1.values.0"))
	assert.Equal(t, "v2_1", d.Get("parameter.1.values.1"))
	{
		_, ok := d.GetOk("parameter.1.value")
		assert.False(t, ok, "Expected `value` to not be set")
	}
}

func TestWidgetCreateWithPosition(t *testing.T) {
	i678 := 678

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
		Resource: ResourceWidget(),
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

	assert.NoError(t, err, err)
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
	i678 := 678

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
		Resource: ResourceWidget(),
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

	assert.NoError(t, err, err)
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
		Resource:    ResourceWidget(),
		Read:        true,
		Removed:     true,
		RequiresNew: true,
		ID:          "some-uuid/12344",
		InstanceState: map[string]string{
			"dashboard_id":     "some-uuid",
			"visualization_id": "678",
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Resource ID should be empty")
}

func TestWidgetUpdate(t *testing.T) {
	sText := "text"

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
								"text": "text"
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
		Resource: ResourceWidget(),
		Update:   true,
		ID:       "some-uuid/12345",
		InstanceState: map[string]string{
			"dashboard_id":     "some-uuid",
			"visualization_id": "678",
		},
		HCL: `
			dashboard_id = "some-uuid"
			text         = "text"
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "some-uuid", d.Get("dashboard_id"))
	assert.Equal(t, "12345", d.Get("widget_id"))

	// Test that visualization_id is now unset (see instance state).
	{
		_, ok := d.GetOk("visualization_id")
		assert.False(t, ok, "visualization_id should not be set")
	}

	// Test that text is now set.
	{
		v, ok := d.GetOk("text")
		assert.True(t, ok, "text should be set")
		assert.Equal(t, "text", v)
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
		Resource: ResourceWidget(),
		Delete:   true,
		ID:       "some-uuid/12345",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "some-uuid/12345", d.Id(), "Resource ID should not be empty")
}

func TestResourceWidgetCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceWidget(), "foo/bar")
}
