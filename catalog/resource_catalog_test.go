package catalog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCatalogCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalog())
}

func TestCatalogCreateAlsoDeletesDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateCatalog{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}, nil)
			w.GetMockSchemasAPI().EXPECT().DeleteByFullName(mock.Anything, "a.default").Return(nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}, nil)
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

func TestCatalogCreateWithForeignCatalogDoesNotDeleteDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateCatalog{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
				ConnectionName: "g", // this indicates a foreign catalog
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
				ConnectionName: "g",
			}, nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
				ConnectionName: "g",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateCatalog{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}, nil)
			w.GetMockSchemasAPI().EXPECT().DeleteByFullName(mock.Anything, "a.default").Return(nil)
			e.Update(mock.Anything, catalog.UpdateCatalog{
				Name:                         "a",
				Owner:                        "administrators",
				Comment:                      "b",
				EnablePredictiveOptimization: "ENABLE",
				Properties: map[string]string{
					"c": "d",
				},
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
				Owner:                        "administrators",
				EnablePredictiveOptimization: "ENABLE",
			}, nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
				Owner:                        "administrators",
				EnablePredictiveOptimization: "ENABLE",
			}, nil)
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
		enable_predictive_optimization = "ENABLE"
		`,
	}.ApplyNoError(t)
}

func TestCatalogCreateCannotDeleteDefaultSchema(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateCatalog{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "b",
				Properties: map[string]string{
					"c": "d",
				},
			}, nil)
			w.GetMockSchemasAPI().EXPECT().DeleteByFullName(mock.Anything, "a.default").Return(errors.New("Something"))
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "d",
			}, nil)
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateCatalog{
				Name:  "a",
				Owner: "administrators",
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "c",
				Owner:   "administrators",
			}, nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "c",
				Owner:   "administrators",
			}, nil)
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

func TestUpdateCatalogOwnerOnly(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCatalogsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateCatalog{
				Name:  "a",
				Owner: "updatedOwner",
			}).Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "c",
				Owner:   "updatedOwner",
			}, nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:    "a",
				Comment: "c",
				Owner:   "updatedOwner",
			}, nil)
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":    "a",
			"comment": "c",
			"owner":   "administrators",
		},
		HCL: `
		name = "a"
		comment = "c"
		owner = "updatedOwner"
		`,
	}.ApplyNoError(t)
}

func TestFailIfMetastoreIdIsWrong(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "old_id",
			}, nil)
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"metastore_id": "old_id",
			"name":         "a",
			"comment":      "c",
		},
		HCL: `
		metastore_id = "new_id"
		name = "a"
		comment = "c"
		owner = "administrators"
		`,
	}.ExpectError(t, "metastore_id must be empty or equal to the metastore id assigned to the workspace: old_id. "+
		"If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set")
}

func TestUpdateCatalogIfMetastoreIdChanges(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "correct_id",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Name:    "a",
					Comment: "c",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "correct_id",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "correct_id",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"metastore_id": "wrong_id",
			"name":         "a",
			"comment":      "c",
		},
		HCL: `
		metastore_id = "correct_id"
		name = "a"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"metastore_id": "correct_id",
		"name":         "a",
		"comment":      "c",
	})
}

func TestUpdateCatalogIfMetastoreIdNotExplicitelySet(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "correct_id",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Name:    "a",
					Comment: "c",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "correct_id",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "correct_id",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"metastore_id": "correct_id",
			"name":         "a",
			"comment":      "c",
		},
		HCL: `
		name = "a"
		comment = "c"
		owner = "updatedOwner"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateCatalogOwnerAndOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "updatedOwner",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Comment: "e",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "e",
					Owner:       "updatedOwner",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/a?",
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "e",
					Owner:       "updatedOwner",
				},
			},
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":    "a",
			"comment": "c",
			"owner":   "administrators",
		},
		HCL: `
		name = "a"
		comment = "e"
		owner = "updatedOwner"
		`,
	}.ApplyNoError(t)
}

func TestUpdateCatalogUpdateRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "updatedOwner",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Comment: "e",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "administrators",
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
			"name":    "a",
			"comment": "c",
			"owner":   "administrators",
		},
		HCL: `
		name = "a"
		comment = "e"
		owner = "updatedOwner"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateCatalogUpdateRollbackError(t *testing.T) {
	serverErrMessage := "Something unexpected happened"
	rollbackErrMessage := "Internal error happened"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "updatedOwner",
				},
				Response: catalog.CatalogInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Comment: "e",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   serverErrMessage,
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/catalogs/a",
				ExpectedRequest: catalog.UpdateCatalog{
					Owner: "administrators",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   rollbackErrMessage,
				},
				Status: 400,
			},
		},
		Resource: ResourceCatalog(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":    "a",
			"comment": "c",
			"owner":   "administrators",
		},
		HCL: `
		name = "a"
		comment = "e"
		owner = "updatedOwner"
		`,
	}.Apply(t)
	errOccurred := fmt.Sprintf("%s. Owner rollback also failed: %s", serverErrMessage, rollbackErrMessage)
	qa.AssertErrorStartsWith(t, err, errOccurred)
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/catalogs",
				ExpectedRequest: catalog.CatalogInfo{
					Name:    "foreign_catalog",
					Comment: "b",
					Options: map[string]string{
						"database": "abcd",
					},
					ConnectionName: "foo",
				},
				Response: catalog.CatalogInfo{
					Name:    "foreign_catalog",
					Comment: "b",
					Options: map[string]string{
						"database": "abcd",
					},
					ConnectionName: "foo",
					MetastoreId:    "e",
					Owner:          "f",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/foreign_catalog.default?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs/foreign_catalog?",
				Response: catalog.CatalogInfo{
					Name:    "foreign_catalog",
					Comment: "b",
					Options: map[string]string{
						"database": "abcd",
					},
					ConnectionName: "foo",
					MetastoreId:    "e",
					Owner:          "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "foreign_catalog"
		comment = "b"
		options = {
			database = "abcd"
		}
		connection_name = "foo"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "foreign_catalog", d.Get("name"))
	assert.Equal(t, "foo", d.Get("connection_name"))
	assert.Equal(t, "b", d.Get("comment"))
	assert.Equal(t, map[string]interface{}{"database": "abcd"}, d.Get("options"))
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
				Resource: "/api/2.1/unity-catalog/bindings/catalog/a",
				ExpectedRequest: catalog.UpdateWorkspaceBindingsParameters{
					SecurableName: "a",
					SecurableType: "catalog",
					Add: []catalog.WorkspaceBinding{
						{
							WorkspaceId: int64(123456789101112),
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
						},
					},
				},
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							WorkspaceId: int64(123456789101112),
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
						},
					},
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

func TestCatalogSuppressCaseSensitivity(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCatalog(),
		ID:       "a",
		InstanceState: map[string]string{
			"metastore_id": "d",
			"name":         "a",
			"comment":      "c",
		},
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"force_destroy":  {Old: "", New: "false", NewComputed: false, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"isolation_mode": {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"owner":          {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
		},
		HCL: `
		name = "A"
		comment = "c"
		`,
	}.ApplyNoError(t)
}
