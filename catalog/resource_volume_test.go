package catalog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestVolumesCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestVolumesCreateWithoutInitialOwner(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "InitialOwner",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "InitialOwner",
				},
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "InitialOwner", d.Get("owner"))
	assert.Equal(t, "MANAGED", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
}

func TestVolumesCreateWithInitialOwner(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "initialOwner",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwner",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testName",
					Comment: "This is a test comment.",
					Owner:   "testOwner",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwner",
				},
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name = "testName"
		owner = "testOwner"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "testOwner", d.Get("owner"))
	assert.Equal(t, "MANAGED", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
}

func TestVolumesCreateWithoutInitialOwner_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name = "testName"
		owner = "testOwner"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestVolumesCreateWithInitialOwner_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "initialOwner",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwner",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name = "testName"
		owner = "testOwner"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestVolumesRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
				},
			},
		},
		Resource: ResourceVolume(),
		Read:     true,
		ID:       "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "MANAGED", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
	assert.Equal(t, "/Volumes/testCatalogName/testSchemaName/testName", d.Get("volume_path"))
}

func TestResourceVolumeRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceVolume(),
		Read:     true,
		ID:       "testCatalogName.testSchemaName.testName",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "testCatalogName.testSchemaName.testName", d.Id(), "Id should not be empty for error reads")
}

func TestVolumesUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					NewName: "testNameNew",
					Comment: "This is a new test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwnerNew",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testNameNew",
					Owner:       "testOwnerNew",
				},
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
		},
		ID: "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testNameNew"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testNameNew", d.Get("name"))
	assert.Equal(t, "MANAGED", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a new test comment.", d.Get("comment"))
	assert.Equal(t, "/Volumes/testCatalogName/testSchemaName/testNameNew", d.Get("volume_path"))
}

func TestVolumesUpdateCommentOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Comment:         "",
					ForceSendFields: []string{"Comment"},
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "",
					FullName:    "testCatalogName.testSchemaName.testName",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "",
					FullName:    "testCatalogName.testSchemaName.testNameNew",
				},
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		InstanceState: map[string]string{
			"name":         "testName",
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
			"comment":      "this is a comment",
		},
		ID: "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = ""
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":         "testName",
		"volume_type":  "MANAGED",
		"catalog_name": "testCatalogName",
		"schema_name":  "testSchemaName",
		"comment":      "",
		"volume_path":  "/Volumes/testCatalogName/testSchemaName/testNameNew",
	})
}

func TestVolumesUpdateForceNewOnCatalog(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogNameNew.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogNameNew",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testNameNew",
					Owner:       "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					NewName: "testNameNew",
					Comment: "This is a new test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogNameNew",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogNameNew.testSchemaName.testName",
					Owner:       "testOwnerNew",
				},
			},
		},
		Resource:    ResourceVolume(),
		RequiresNew: true,
		Update:      true,
		ID:          "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testNameNew"
		volume_type = "MANAGED"
		catalog_name = "testCatalogNameNew"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testNameNew", d.Get("name"))
	assert.Equal(t, "MANAGED", d.Get("volume_type"))
	assert.Equal(t, "testCatalogNameNew", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a new test comment.", d.Get("comment"))
}

func TestVolumesValidateOnVolumesType(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Resource:    ResourceVolume(),
		RequiresNew: true,
		Update:      true,
		ID:          "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testName"
		volume_type = "unknown"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."		
		`,
	}.Apply(t)
	assert.ErrorContains(t, err, "value unknown is not one of EXTERNAL, MANAGED")
}

func TestVolumesUpdateForceNewOnVolumeType(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("EXTERNAL"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testNameNew",
					Owner:       "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testName",
					Comment: "This is a new test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("EXTERNAL"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwnerNew",
				},
			},
		},
		Resource:    ResourceVolume(),
		RequiresNew: true,
		Update:      true,
		ID:          "testCatalogName.testSchemaName.testName",
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
		},
		HCL: `
		name = "testName"
		volume_type = "EXTERNAL"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testNameNew", d.Get("name"))
	assert.Equal(t, "EXTERNAL", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a new test comment.", d.Get("comment"))
}

func TestVolumesUpdateWithOwner(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testNameNew",
					Owner:       "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testName",
					Comment: "This is a new test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("MANAGED"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a new test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
					Owner:       "testOwnerNew",
				},
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		ID:       "testCatalogName.testSchemaName.testName",
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
			"owner":        "testOwnerOld",
		},
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testNameNew", d.Get("name"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "testOwnerNew", d.Get("owner"))
	assert.Equal(t, "This is a new test comment.", d.Get("comment"))
}

func TestVolumesUpdateRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testName",
					Comment: "This is a new test comment.",
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerOld",
				},
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		ID:       "testCatalogName.testSchemaName.testName",
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
			"owner":        "testOwnerOld",
		},
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestVolumesUpdateRollback_Error(t *testing.T) {
	serverErrMessage := "Something unexpected happened"
	rollbackErrMessage := "Internal error happened"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testName",
					Comment: "This is a new test comment.",
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   serverErrMessage,
				},
				Status: 500,
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerOld",
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   rollbackErrMessage,
				},
				Status: 400,
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		ID:       "testCatalogName.testSchemaName.testName",
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
			"owner":        "testOwnerOld",
		},
		HCL: `
		name = "testName"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	errOccurred := fmt.Sprintf("%s. Owner rollback also failed: %s", serverErrMessage, rollbackErrMessage)
	qa.AssertErrorStartsWith(t, err, errOccurred)
}

func TestVolumeUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Owner: "testOwnerNew",
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		InstanceState: map[string]string{
			"catalog_name": "testCatalogName",
			"schema_name":  "testSchemaName",
			"volume_type":  "MANAGED",
		},
		ID: "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testNameNew"
		volume_type = "MANAGED"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestVolumeDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
			},
		},
		Resource: ResourceVolume(),
		Delete:   true,
		ID:       "testCatalogName.testSchemaName.testName",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testCatalogName.testSchemaName.testName", d.Id())
}

func TestVolumeDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceVolume(),
		Delete:   true,
		Removed:  true,
		ID:       "testCatalogName.testSchemaName.testName",
	}.ExpectError(t, "Something went wrong")
}
