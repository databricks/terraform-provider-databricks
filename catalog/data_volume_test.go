package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataSourceVolume_ReadByFullName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes/catalog.schema.name?",
				Response: catalog.VolumeInfo{
					FullName:    "catalog.schema.name",
					CatalogName: "catalog",
					SchemaName:  "schema",
					Name:        "name",
				},
			},
		},
		Resource: DataSourceVolume(),
		HCL: `
		full_name = "catalog.schema.name"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestDataSourceVolume_InvalidConfigExclusivceInputs(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: DataSourceVolume(),
		HCL: `
		full_name = "abc"
		catalog_name = "a"
		schema_name = "b"
		name = "c"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	assert.Error(t, err)
}

func TestDataSourceVolume_InvalidConfigOnlyCatalogName(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource:    DataSourceVolume(),
		HCL:         `catalog_name = "a"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	assert.Error(t, err)
}

func TestDataSourceVolume_InvalidConfigOnlySchemaName(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource:    DataSourceVolume(),
		HCL:         `schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	assert.Error(t, err)
}

func TestDataSourceVolume_InvalidConfigOnlyName(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource:    DataSourceVolume(),
		HCL:         `name = "c"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	assert.Error(t, err)
}

func TestDataSourceVolume_ReadByCatalogSchemaName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes/catalog.schema.name?",
				Response: catalog.VolumeInfo{
					FullName:    "catalog.schema.name",
					CatalogName: "catalog",
					SchemaName:  "schema",
					Name:        "name",
				},
			},
		},
		Resource: DataSourceVolume(),
		HCL: `
		catalog_name = "catalog"
		schema_name = "schema"
		name = "name"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestDataSourceVolume_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceVolume(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
