package catalog

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "InitialOwner", d.Get("owner"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "testOwner", d.Get("owner"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
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
				Response: apierr.APIErrorBody{
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
		volume_type = "testVolumeType"
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
				},
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
				Response: apierr.APIErrorBody{
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
		volume_type = "testVolumeType"
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
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
}

func TestResourceVolumeRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: apierr.APIErrorBody{
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
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName?",
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
					Name:    "testNameNew",
					Comment: "This is a new test comment.",
					Owner:   "testOwnerNew",
				},
				Response: catalog.VolumeInfo{
					Name:        "testNameNew",
					VolumeType:  catalog.VolumeType("testVolumeType"),
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
		HCL: `
		name = "testNameNew"
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a new test comment."
		owner = "testOwnerNew"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testNameNew", d.Get("name"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a new test comment.", d.Get("comment"))
}

func TestVolumeUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/testCatalogName.testSchemaName.testName",
				ExpectedRequest: catalog.UpdateVolumeRequestContent{
					Name:    "testNameNew",
					Comment: "This is a new test comment.",
					Owner:   "testOwnerNew",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceVolume(),
		Update:   true,
		ID:       "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testNameNew"
		volume_type = "testVolumeType"
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
				Response: apierr.APIErrorBody{
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

func TestVolumesList(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/unity-catalog/volumes?catalog_name=&schema_name=",
			Response: map[string]any{},
		},
	})
	require.NoError(t, err)

	w, err := client.WorkspaceClient()
	require.NoError(t, err)

	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	vLists, err := w.Volumes.Impl().List(ctx, catalog.ListVolumesRequest{})

	require.NoError(t, err)
	assert.Equal(t, 0, len(vLists.Volumes))
}
