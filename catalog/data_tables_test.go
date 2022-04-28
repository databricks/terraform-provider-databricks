package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTablesData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/?catalog_name=a&schema_name=b",
				Response: Tables{
					Tables: []TableInfo{
						{
							Name: "a.b.c",
						},
						{
							Name: "a.b.d",
						},
					},
				},
			},
		},
		Resource: DataSourceTables(),
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

// https://github.com/databrickslabs/terraform-provider-databricks/issues/1264
func TestTablesDataIssue1264(t *testing.T) {
	r := DataSourceTables()
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/?catalog_name=a&schema_name=b",
				Response: Tables{
					Tables: []TableInfo{
						{
							Name: "a",
						},
						{
							Name: "b",
						},
					},
				},
			},
		},
		Resource: r,
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	s := d.Get("ids").(*schema.Set)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("..a"))

	d, err = qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/?catalog_name=a&schema_name=b",
				Response: Tables{
					Tables: []TableInfo{
						{
							Name: "c",
						},
						{
							Name: "d",
						},
					},
				},
			},
		},
		Resource: r,
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	s = d.Get("ids").(*schema.Set)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("..c"))
}

func TestTablesData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceTables(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
