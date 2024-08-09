package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestSqlAlertReadStringValue(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/alerts/xyz?",
				Response: sql.LegacyAlert{
					CreatedAt: "2020-01-01T00:00:00.000Z",
					Id:        "xyz",
					Parent:    "abc",
					Name:      "Alert name",
					Query: &sql.AlertQuery{
						Id: "abc",
					},
					Options: &sql.AlertOptions{
						Column: "col1",
						Op:     "==",
						Value:  "10",
					},
				},
			},
		},
		Resource: ResourceSqlAlert(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestSqlAlertReadNumberValue(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/alerts/xyz?",
				Response: sql.LegacyAlert{
					CreatedAt: "2020-01-01T00:00:00.000Z",
					Id:        "xyz",
					Parent:    "abc",
					Name:      "Alert name",
					Query: &sql.AlertQuery{
						Id: "abc",
					},
					Options: &sql.AlertOptions{
						Column: "col1",
						Op:     ">",
						Value:  10,
					},
				},
			},
		},
		Resource: ResourceSqlAlert(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestSqlAlertReadBoolValue(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/alerts/xyz?",
				Response: sql.LegacyAlert{
					CreatedAt: "2020-01-01T00:00:00.000Z",
					Id:        "xyz",
					Parent:    "abc",
					Name:      "Alert name",
					Query: &sql.AlertQuery{
						Id: "abc",
					},
					Options: &sql.AlertOptions{
						Column: "col1",
						Op:     "==",
						Value:  true,
					},
				},
			},
		},
		Resource: ResourceSqlAlert(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}
