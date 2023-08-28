package catalog

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatalogCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalog())
}

func TestCatalogCreateAlsoDeletesDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CreateCatalog{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId: "e",
					Owner:       "f",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/a.default?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId: "e",
					Owner:       "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateWithForiegnCatalogDoesNotDeleteDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CreateCatalog{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					ConnectionName: "g", // this indicates a foriegn catalog
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId:    "e",
					Owner:          "f",
					ConnectionName: "g",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId:    "e",
					Owner:          "f",
					ConnectionName: "g",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		connection_name = "g"
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateWithOwnerAlsoDeletesDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CreateCatalog{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					Owner: "testers",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/a.default?",
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					Owner: "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId: "e",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateCannotDeleteDefaultSchema(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					Owner: "testers",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/a.default?",
				Status:   400,
				Response: apierr.APIErrorBody{
					Message: "Something",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		`,
	}.Apply(t)
	require.Error(t, err)
	assert.Equal(t, "cannot remove new catalog default schema: Something", fmt.Sprint(err))

}

func TestUpdateCatalog(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Name:    "a",
					Comment: "c",
					Owner:   "administrators",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"metastore_id": "d",
			"name":         "a",
			"comment":      "c",
		},
		HCL: `
		name = "a"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestForceDeleteCatalog(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas?catalog_name=b",
				Response: []catalog.SchemaInfo{
					{
						Name:     "a",
						FullName: "b.a",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/?catalog_name=b&schema_name=a",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "c",
							TableType:   "MANAGED",
						},
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "d",
							TableType:   "VIEW",
						},
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/tables/b.a.c",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/tables/b.a.d",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/catalogs/b?force=true",
			},
		},
		Resource: ResourceCatalog(),
		Delete:   true,
		ID:       "b",
		HCL: `
		name = "b"
		comment = "c"
		owner = "administrators"
		force_destroy = true
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateDeltaSharing(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					ProviderName: "foo",
					ShareName:    "bar",
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					ProviderName: "foo",
					ShareName:    "bar",
					MetastoreId:  "e",
					Owner:        "f",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					ProviderName: "foo",
					ShareName:    "bar",
					MetastoreId:  "e",
					Owner:        "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		provider_name = "foo"
		share_name = "bar"
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateForeign(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Options: map[string]string{
						"a": "b",
					},
					ConnectionName: "foo",
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Options: map[string]string{
						"a": "b",
					},
					ConnectionName: "foo",
					MetastoreId:  "e",
					Owner:        "f",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/a.default?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Options: map[string]string{
						"a": "b",
					},
					ConnectionName: "foo",
					MetastoreId:  "e",
					Owner:        "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		options = {
			a = "b"
		}
		connection_name = "foo"
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateIsolated(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreId: "e",
					Owner:       "f",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/a.default?",
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					IsolationMode: "ISOLATED",
				},
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					IsolationMode: "ISOLATED",
					MetastoreId:   "e",
					Owner:         "f",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",

				Response: catalog.MetastoreAssignment{
					MetastoreId: "e",
					WorkspaceId: 123456789101112,
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/workspace-bindings/catalogs/a",
				ExpectedRequest: catalog.UpdateWorkspaceBindings{
					Name:             "a",
					AssignWorkspaces: []int64{123456789101112},
				},

				Response: catalog.CurrentWorkspaceBindings{
					Workspaces: []int64{123456789101112},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					IsolationMode: "ISOLATED",
					MetastoreId:   "e",
					Owner:         "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		isolation_mode = "ISOLATED"
		`,
	}.ApplyNoError(t)
}

func TestUcDirectoryPathSuppressDiff(t *testing.T) {
	assert.True(t, ucDirectoryPathSlashOnlySuppressDiff("", "abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH",
		"abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH/", nil))
	assert.True(t, ucDirectoryPathSlashOnlySuppressDiff("", "abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH/",
		"abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH", nil))
	assert.False(t, ucDirectoryPathSlashOnlySuppressDiff("", "abfss://test@test.dfs.core.windows.net/new_dir",
		"abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH/", nil))
	//
	assert.True(t, ucDirectoryPathSlashAndEmptySuppressDiff("", "abfss://test@test.dfs.core.windows.net/TF_DIR_WITH_SLASH/",
		"", nil))
	assert.False(t, ucDirectoryPathSlashOnlySuppressDiff("", "abfss://test@test.dfs.core.windows.net/new_dir",
		"abfss://test@test.dfs.core.windows.net/OTHER/", nil))
}
