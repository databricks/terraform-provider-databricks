package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTablesData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables?catalog_name=a&schema_name=b",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							FullName: "a.b.c",
							Name:     "c",
						},
						{
							FullName: "a.b.d",
							Name:     "d",
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

// https://github.com/databricks/terraform-provider-databricks/issues/1264
func TestTablesDataIssue1264(t *testing.T) {
	r := DataSourceTables()
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables?catalog_name=a&schema_name=b",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							Name:     "a",
							FullName: "a.b.a",
						},
						{
							Name:     "b",
							FullName: "a.b.b",
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
	assert.True(t, s.Contains("a.b.a"))

	d, err = qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables?catalog_name=a&schema_name=b",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							Name:     "c",
							FullName: "a.b.c",
						},
						{
							Name:     "d",
							FullName: "a.b.d",
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
	assert.True(t, s.Contains("a.b.c"))
}

func TestTablesData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceTables(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
