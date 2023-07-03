package sql

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWarehouseData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:        "foo",
					ClusterSize: "Small",
					ID:          "abc",
					State:       "RUNNING",
				},
			},
			dataSourceListHTTPFixture,
		},
		Resource:    DataSourceWarehouse(),
		HCL:         `id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "foo", d.Get("name"))
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestWarehouseData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `id = "abc"`,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}

func TestWarehouseDataByName_ListError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}

func TestWarehouseDataByName_NotFoundError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses",

				Response: EndpointList{
					Endpoints: []SQLEndpoint{
						{
							Name:        "foo",
							ClusterSize: "Small",
							ID:          "abc",
							State:       "RUNNING",
						},
					},
				},
			},
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.ExpectError(t, "can't find SQL warehouse with the name 'abc'")
}

func TestWarehouseDataByName_DuplicatesError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses",

				Response: EndpointList{
					Endpoints: []SQLEndpoint{
						{
							Name:        "abc",
							ClusterSize: "Small",
							ID:          "abc",
							State:       "RUNNING",
						},
						{
							Name:        "Abc",
							ClusterSize: "Small",
							ID:          "abc2",
							State:       "RUNNING",
						},
					},
				},
			},
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.ExpectError(t, "there are multiple SQL warehouses with the name 'abc'")
}

func TestWarehouseDataByName(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses",

				Response: EndpointList{
					Endpoints: []SQLEndpoint{
						{
							Name:        "foo",
							ClusterSize: "Small",
							ID:          "abc2",
							State:       "RUNNING",
						},
						{
							Name:        "test",
							ClusterSize: "Small",
							ID:          "abc",
							State:       "RUNNING",
						},
					},
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:        "test",
					ClusterSize: "Small",
					ID:          "abc",
					State:       "RUNNING",
				},
			},
			dataSourceListHTTPFixture,
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "test"`,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}
