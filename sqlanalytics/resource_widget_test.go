package sqlanalytics

import (
	"encoding/json"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/stretchr/testify/assert"
)

func TestWidgetCreate(t *testing.T) {
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
					ID:              12345,
					DashboardID:     "some-uuid",
					VisualizationID: &i678,
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
								"id": 12344,
								"visualization_id": null
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"visualization_id": 678
							}
						`),
						json.RawMessage(`
							{
								"id": 12345,
								"visualization_id": null
							}
						`),
					},
				},
			},
		},
		Resource: ResourceWidget(),
		Create:   true,
		State: map[string]interface{}{
			"dashboard_id":     "some-uuid",
			"visualization_id": "678",
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "12345", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "678", d.Get("visualization_id"))
}
