package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/a.b.c?",
				Response: catalog.TableInfo{
					CreatedAt:   1706294508998,
					CatalogName: "a",
					SchemaName:  "b",
					Name:        "c",
				},
			},
		},
		Resource: DataSourceTable(),
		HCL: `
		full_name = "a.b.c"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	assert := assert.New(t)
	assert.Equal(d.Get("catalog_name"), "a")
	assert.Equal(d.Get("schema_name"), "b")
	assert.Equal(d.Get("name"), "c")
}

func TestTableData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceTable(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
