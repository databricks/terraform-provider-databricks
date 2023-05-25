package catalog

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVolumesCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestVolumesCreate(t *testing.T) {
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
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/" + "testCatalogName.testSchemaName.testName" + "?",
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
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
}

func TestVolumesRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/" + "testCatalogName.testSchemaName.testName" + "?",
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
		Resource: ResourceVolumes(),
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

func TestVolumesUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/volumes/" + "testCatalogName.testSchemaName.testName" + "?",
				Response: catalog.VolumeInfo{
					Name:        "testName",
					VolumeType:  catalog.VolumeType("testVolumeType"),
					CatalogName: "testCatalogName",
					SchemaName:  "testSchemaName",
					Comment:     "This is a test comment.",
					FullName:    "testCatalogName.testSchemaName.testName",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/volumes/" + "testCatalogName.testSchemaName.testName",
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
				},
			},
		},
		Resource: ResourceVolumes(),
		Update:   true,
		ID:       "testCatalogName.testSchemaName.testName",
		HCL: `
		name = "testName"
		volume_type = "testVolumeType"
		catalog_name = "testCatalogName"
		schema_name = "testSchemaName"
		comment = "This is a test comment."
		owner = "testOwner"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "testVolumeType", d.Get("volume_type"))
	assert.Equal(t, "testCatalogName", d.Get("catalog_name"))
	assert.Equal(t, "testSchemaName", d.Get("schema_name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
}

func TestVolumeDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/volumes/" + "testCatalogName.testSchemaName.testName" + "?",
			},
		},
		Resource: ResourceVolumes(),
		Delete:   true,
		ID:       "testCatalogName.testSchemaName.testName",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testCatalogName.testSchemaName.testName", d.Id())
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
