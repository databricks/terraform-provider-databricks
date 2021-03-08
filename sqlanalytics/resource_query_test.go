package sqlanalytics

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/stretchr/testify/assert"
)

func TestQueryCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
				},
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
				},
			},
		},
		Resource: ResourceQuery(),
		Create:   true,
		State: map[string]interface{}{
			"data_source_id": "xyz",
			"name":           "Query name",
			"description":    "Query description",
			"query":          "SELECT 1",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Query name", d.Get("name"))
	assert.Equal(t, "Query description", d.Get("description"))
	assert.Equal(t, "SELECT 1", d.Get("query"))
}

func TestQueryRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
				},
			},
		},
		Resource: ResourceQuery(),
		Read:     true,
		ID:       "foo",
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
}

func TestQueryUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Updated name",
					Description:  "Updated description",
					Query:        "SELECT 2",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Updated name",
					Description:  "Updated description",
					Query:        "SELECT 2",
				},
			},
		},
		Resource: ResourceQuery(),
		Update:   true,
		ID:       "foo",
		State: map[string]interface{}{
			"data_source_id": "xyz",
			"name":           "Updated name",
			"description":    "Updated description",
			"query":          "SELECT 2",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Updated name", d.Get("name"))
	assert.Equal(t, "Updated description", d.Get("description"))
	assert.Equal(t, "SELECT 2", d.Get("query"))
}

func TestQueryUpdateWithParams(t *testing.T) {
	body := api.Query{
		ID:           "foo",
		DataSourceID: "xyz",
		Name:         "Updated name",
		Query:        "SELECT 1, 2, 3, 4",
		Options: &api.QueryOptions{
			Parameters: []interface{}{
				api.QueryParameterText{
					QueryParameter: api.QueryParameter{
						Name:  "1",
						Title: "Title for column 1",
					},
				},
				api.QueryParameterNumber{
					QueryParameter: api.QueryParameter{
						Name:  "2",
						Title: "Title for column 2",
					},
				},
				api.QueryParameterEnum{
					QueryParameter: api.QueryParameter{
						Name:  "3",
						Title: "Title for column 3",
					},
					Options: "e1\ne2",
					Multi: &api.QueryParameterMultipleValuesOptions{
						Prefix:    "\"",
						Suffix:    "\"",
						Separator: ",",
					},
				},
				api.QueryParameterEnum{
					QueryParameter: api.QueryParameter{
						Name:  "3",
						Title: "Title for column 3 without multiple",
					},
					Options: "e1\ne2",
					Multi:   nil,
				},
				api.QueryParameterQuery{
					QueryParameter: api.QueryParameter{
						Name:  "4",
						Title: "Title for column 4",
					},
					QueryID: "abc",
					Multi: &api.QueryParameterMultipleValuesOptions{
						Prefix:    "\"",
						Suffix:    "\"",
						Separator: ",",
					},
				},
				api.QueryParameterQuery{
					QueryParameter: api.QueryParameter{
						Name:  "4",
						Title: "Title for column 4 without multiple",
					},
					QueryID: "abc",
					Multi:   nil,
				},
			},
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: body,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: body,
			},
		},
		Resource: ResourceQuery(),
		Update:   true,
		ID:       "foo",
		HCL: `
			data_source_id = "xyz"
			name = "name"
			query = "SELECT 1, 2, 3, 4"
			
			parameter {
				name = "1"
				title = "Title for column 1"
				type = "text"
			}

			parameter {
				name = "2"
				title = "Title for column 2"
				type = "number"
			}

			parameter {
				name = "3"
				title = "Title for column 3"
				enum {
					options = ["e1", "e2"]
					multiple {
						prefix = "\""
						suffix = "\""
						separator = ","
					}
				}
			}

			parameter {
				name = "3"
				title = "Title for column 3 without multiple"
				enum {
					options = ["e1", "e2"]
				}
			}

			parameter {
				name = "4"
				title = "Title for column 4"
				query {
					query_id = "abc"
					multiple {
						prefix = "\""
						suffix = "\""
						separator = ","
					}
				}
			}

			parameter {
				name = "4"
				title = "Title for column 4 without multiple"
				query {
					query_id = "abc"
					multiple {
						prefix = "\""
						suffix = "\""
						separator = ","
					}
				}
			}
		`,
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Updated name", d.Get("name"))
	assert.Equal(t, "SELECT 1, 2, 3, 4", d.Get("query"))
	assert.Len(t, d.Get("parameter").([]interface{}), 6)
}

func TestQueryDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/queries/foo",
			},
		},
		Resource: ResourceQuery(),
		Delete:   true,
		ID:       "foo",
		State: map[string]interface{}{
			"data_source_id": "xyz",
			"name":           "Updated name",
			"description":    "Updated description",
			"query":          "SELECT 2",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	// Delete doesn't touch schema.ResourceData, so the ID should survive.
	assert.Equal(t, "foo", d.Id())
}
