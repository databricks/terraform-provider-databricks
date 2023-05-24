package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

var (
	TestingName             = "testName"
	TestingVolumeType       = catalog.VolumeType("testVolumeType")
	TestingVolumeTypeString = "testVolumeType"
	TestingCatalogName      = "testCatalogName"
	TestingSchemaName       = "testSchemaName"
	TestingFullName         = "testCatalogName.testSchemaName.testName"
	TestingComment          = "This is a test comment."
)

func TestVolumesCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateVolumes(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					Name:        TestingName,
					VolumeType:  TestingVolumeType,
					CatalogName: TestingCatalogName,
					SchemaName:  TestingSchemaName,
					Comment:     TestingComment,
				},
				Response: catalog.VolumeInfo{
					Name:        TestingName,
					VolumeType:  TestingVolumeType,
					CatalogName: TestingCatalogName,
					SchemaName:  TestingSchemaName,
					Comment:     TestingComment,
					FullName:    TestingFullName,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/" + TestingFullName + "?",
				Response: catalog.VolumeInfo{
					Name:        TestingName,
					VolumeType:  TestingVolumeType,
					CatalogName: TestingCatalogName,
					SchemaName:  TestingSchemaName,
					Comment:     TestingComment,
					FullName:    TestingFullName,
				},
			},
		},
		Resource: ResourceVolumes(),
		State: map[string]any{
			"name":         TestingName,
			"volume_type":  TestingVolumeType,
			"catalog_name": TestingCatalogName,
			"schema_name":  TestingSchemaName,
			"comment":      TestingComment,
		},
		Create: true,
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
