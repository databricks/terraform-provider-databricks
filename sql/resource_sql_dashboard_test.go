package sql

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/sql/api"
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
		Resource: ResourceSqlDashboard(),
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
		Resource: ResourceSqlDashboard(),
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
		Resource: ResourceSqlDashboard(),
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
		Resource: ResourceSqlDashboard(),
		Delete:   true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestResourceDashboardCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlDashboard())
}
