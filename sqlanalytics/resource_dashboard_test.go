package sqlanalytics

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/stretchr/testify/assert"
)

func TestDashboardCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/dashboards",
				ExpectedRequest: api.Dashboard{
					Name: "Dashboard name",
					Tags: []string{"t1", "t2"},
				},
				Response: api.Dashboard{
					ID:   "xyz",
					Name: "Dashboard name",
					Tags: []string{"t1", "t2"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/xyz",
				Response: api.Dashboard{
					ID:   "xyz",
					Name: "Dashboard name",
					Tags: []string{"t1", "t2"},
				},
			},
		},
		Resource: ResourceDashboard(),
		Create:   true,
		State: map[string]interface{}{
			"name": "Dashboard name",
			"tags": []interface{}{"t1", "t2"},
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "Dashboard name", d.Get("name"))
}

func TestDashboardRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/xyz",
				Response: api.Dashboard{
					ID:   "xyz",
					Name: "Dashboard name",
					Tags: []string{"t1", "t2"},
				},
			},
		},
		Resource: ResourceDashboard(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestDashboardUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/dashboards/xyz",
				Response: api.Dashboard{
					ID:   "xyz",
					Name: "Dashboard renamed",
					Tags: []string{"t2", "t3"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/dashboards/xyz",
				Response: api.Dashboard{
					ID:   "xyz",
					Name: "Dashboard renamed",
					Tags: []string{"t2", "t3"},
				},
			},
		},
		Resource: ResourceDashboard(),
		Update:   true,
		ID:       "xyz",
		State: map[string]interface{}{
			"name": "Dashboard renamed",
			"tags": []interface{}{"t2", "t3"},
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestDashboardDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/dashboards/xyz",
			},
		},
		Resource: ResourceDashboard(),
		Delete:   true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}
