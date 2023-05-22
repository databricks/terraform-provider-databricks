package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

var (
	TestingName             = "testName"
	TestingVolumeType       = VolumeType("testVolumeType")
	TestingVolumeTypeString = "testVolumeType"
	TestingCatalogName      = "testCatalogName"
	TestingSchemaName       = "testSchemaName"
	TestingComment          = "This is a test comment."
)

func TestVolumesCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateVolumes(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: CreateVolumeRequestContent{
					Name:        TestingName,
					VolumeType:  TestingVolumeType,
					CatalogName: TestingCatalogName,
					SchemaName:  TestingSchemaName,
					Comment:     TestingComment,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes/testName",
				Response: VolumeInfo{
					Name:        TestingName,
					VolumeType:  TestingVolumeType,
					CatalogName: TestingCatalogName,
					SchemaName:  TestingSchemaName,
					Comment:     TestingComment,
				},
			},
		},
		Resource: ResourceVolumes(),
		Create:   true,
		HCL: `
		name = "testName"
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingName, d.Get("name"))
	assert.Equal(t, TestingVolumeTypeString, d.Get("volume_type"))
	assert.Equal(t, TestingCatalogName, d.Get("catalog_name"))
	assert.Equal(t, TestingSchemaName, d.Get("schema_name"))
	assert.Equal(t, TestingComment, d.Get("comment"))
}
