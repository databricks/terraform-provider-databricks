package exporter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/tags"

	sdk_uc "github.com/databricks/databricks-sdk-go/service/catalog"
	sdk_sharing "github.com/databricks/databricks-sdk-go/service/sharing"
	sdk_vs "github.com/databricks/databricks-sdk-go/service/vectorsearch"
	tf_uc "github.com/databricks/terraform-provider-databricks/catalog"
	tf_vs "github.com/databricks/terraform-provider-databricks/vectorsearch"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestEmitUserSpOrGroup(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("users,groups")
	emitUserSpOrGroup(ic, "user@example.com")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_user[<unknown>] (user_name: user@example.com)")

	emitUserSpOrGroup(ic, "users")
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users)")

	emitUserSpOrGroup(ic, "abcd1234-ab12-cd34-ef56-abcdef123456")
	assert.Equal(t, 3, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_service_principal[<unknown>] (application_id: abcd1234-ab12-cd34-ef56-abcdef123456)")

	emitUserSpOrGroup(ic, "users @ test.com")
	assert.Equal(t, 4, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users @ test.com)")

}

func TestTagPolicyExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.1/tag-policies?",
			Response: tags.ListTagPoliciesResponse{
				TagPolicies: []tags.TagPolicy{
					{
						TagKey:      "environment",
						Description: "Environment tag policy",
						Values: []tags.Value{
							{Name: "dev"},
							{Name: "staging"},
							{Name: "production"},
						},
					},
					{
						TagKey:      "team",
						Description: "Team tag policy",
						Values: []tags.Value{
							{Name: "engineering"},
							{Name: "data"},
						},
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/tag-policies/environment?",
			ReuseRequest: true,
			Response: tags.TagPolicy{
				TagKey:      "environment",
				Description: "Environment tag policy",
				Values: []tags.Value{
					{Name: "dev"},
					{Name: "staging"},
					{Name: "production"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/tag-policies/team?",
			ReuseRequest: true,
			Response: tags.TagPolicy{
				TagKey:      "team",
				Description: "Team tag policy",
				Values: []tags.Value{
					{Name: "engineering"},
					{Name: "data"},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("uc-tags")
		ic.enableServices("uc-tags")

		err := ic.Run()
		assert.NoError(t, err)

		content, err := os.ReadFile(tmpDir + "/uc-tags.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))
		assert.Contains(t, contentStr, `resource "databricks_tag_policy" "environment"`)
		assert.Contains(t, contentStr, `resource "databricks_tag_policy" "team"`)
		assert.Contains(t, contentStr, `tag_key = "environment"`)
		assert.Contains(t, contentStr, `description = "Environment tag policy"`)
		assert.Contains(t, contentStr, `name = "dev"`)
		assert.Contains(t, contentStr, `name = "staging"`)
		assert.Contains(t, contentStr, `name = "production"`)
	})
}

func TestListCatalogs(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/catalogs?",
			Response: sdk_uc.ListCatalogsResponse{
				Catalogs: []sdk_uc.CatalogInfo{
					{
						Name:        "cat1",
						CatalogType: "MANAGED_CATALOG",
					},
					{
						Name:        "cat2",
						CatalogType: "UNKNOWN",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_catalog"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_catalog[cat1_test_MANAGED_CATALOG] (id: cat1)"])
	})
}

func TestImportManagedCatalog(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/schemas?catalog_name=ctest",
			Response: sdk_uc.ListSchemasResponse{
				Schemas: []sdk_uc.SchemaInfo{
					{
						CatalogType: "MANAGED_CATALOG",
						Name:        "schema1",
						FullName:    "ctest.schema1",
					},
					{
						CatalogType: "MANAGED_CATALOG",
						Name:        "information_schema",
						FullName:    "ctest.schema1",
					},
					{
						CatalogType: "UNKNOWN",
						FullName:    "ctest.schema2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas")
		ic.enableListing("uc-schemas")
		ic.currentMetastore = currentMetastoreResponse
		d := tf_uc.ResourceCatalog().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		err := resourcesMap["databricks_catalog"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/ctest)"])
		assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.schema1)"])
	})
}

func TestImportForeignCatalog(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-catalogs,uc-grants,uc-connections")
	ic.currentMetastore = currentMetastoreResponse
	d := tf_uc.ResourceCatalog().ToResource().TestResourceData()
	d.SetId("fctest")
	d.Set("metastore_id", "1234")
	d.Set("connection_name", "conn")
	d.Set("name", "fctest")
	err := resourcesMap["databricks_catalog"].Import(ic, &resource{
		ID:   "fctest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/fctest)"])
	assert.True(t, ic.testEmits["databricks_connection[<unknown>] (id: 1234|conn)"])
}

func TestImportIsolatedManagedCatalog(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/schemas?catalog_name=ctest",
			Response: sdk_uc.ListSchemasResponse{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/bindings/catalog/ctest?",
			Response: sdk_uc.UpdateWorkspaceBindingsResponse{
				Bindings: []sdk_uc.WorkspaceBinding{
					{
						BindingType: "BINDING_TYPE_READ",
						WorkspaceId: 1234,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas")
		ic.enableListing("uc-schemas,uc-volumes,uc-models,uc-tables")
		ic.currentMetastore = currentMetastoreResponse
		d := tf_uc.ResourceCatalog().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		d.Set("isolation_mode", "ISOLATED")
		err := resourcesMap["databricks_catalog"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/ctest)"])
		assert.True(t, ic.testEmits["databricks_workspace_binding[catalog_ctest_ws_1234] (id: 1234|catalog|ctest)"])
	})
}

func TestImportSchema(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/models?catalog_name=ctest&schema_name=stest",
			Response: sdk_uc.ListRegisteredModelsResponse{
				RegisteredModels: []sdk_uc.RegisteredModelInfo{
					{
						Name:     "model1",
						FullName: "ctest.stest.model1",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/volumes?catalog_name=ctest&schema_name=stest",
			Response: sdk_uc.ListVolumesResponseContent{
				Volumes: []sdk_uc.VolumeInfo{
					{
						Name:     "volume1",
						FullName: "ctest.stest.volume1",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/tables?catalog_name=ctest&schema_name=stest",
			Response: sdk_uc.ListTablesResponse{
				Tables: []sdk_uc.TableInfo{
					{
						Name:      "table1",
						TableType: "MANAGED",
						FullName:  "ctest.stest.table1",
					},
					{
						Name:      "table2",
						TableType: "UNKNOWN",
						FullName:  "ctest.stest.table2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas,uc-volumes,uc-models,uc-tables")
		ic.enableListing("uc-schemas,uc-volumes,uc-models,uc-tables")
		ic.currentMetastore = currentMetastoreResponse
		d := tf_uc.ResourceSchema().ToResource().TestResourceData()
		d.SetId("ctest.stest")
		d.Set("catalog_name", "ctest")
		d.Set("name", "stest")
		err := resourcesMap["databricks_schema"].Import(ic, &resource{
			ID:   "ctest.stest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 5, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: schema/ctest.stest)"])
		assert.True(t, ic.testEmits["databricks_catalog[<unknown>] (id: ctest)"])
		assert.True(t, ic.testEmits["databricks_registered_model[<unknown>] (id: ctest.stest.model1)"])
		assert.True(t, ic.testEmits["databricks_volume[<unknown>] (id: ctest.stest.volume1)"])
		assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: ctest.stest.table1)"])
	})
}

func TestConnections(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/connections?",
			Response: sdk_uc.ListConnectionsResponse{
				Connections: []sdk_uc.ConnectionInfo{
					{
						Name:        "test",
						MetastoreId: "12345",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-connections,uc-grants")
		// Test Listing
		err := resourcesMap["databricks_connection"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_connection[<unknown>] (id: 12345|test)"])
		// Test Importing
		d := tf_uc.ResourceConnection().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		err = resourcesMap["databricks_connection"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: foreign_connection/ctest)"])
	})
}

func TestListExternalLocations(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/external-locations?",
			Response: sdk_uc.ListExternalLocationsResponse{
				ExternalLocations: []sdk_uc.ExternalLocationInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-external-locations,uc-storage-credentials,uc-grants")
		ic.currentMetastore = currentMetastoreResponse
		// Test listing
		err := resourcesMap["databricks_external_location"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_external_location[<unknown>] (id: test)"])
		// Test import
		d := tf_uc.ResourceExternalLocation().ToResource().TestResourceData()
		d.SetId("ext_loc")
		d.Set("credential_name", "stest")
		err = resourcesMap["databricks_external_location"].Import(ic, &resource{
			ID:   "ext_loc",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 3, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: external_location/ext_loc)"])
		assert.True(t, ic.testEmits["databricks_storage_credential[<unknown>] (id: stest)"])
	})
}

func TestServiceCredentials(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/credentials?",
			Response: sdk_uc.ListCredentialsResponse{
				Credentials: []sdk_uc.CredentialInfo{
					{
						Name:    "test-storage",
						Purpose: sdk_uc.CredentialPurposeStorage,
					},
					{
						Name:    "test-service",
						Purpose: sdk_uc.CredentialPurposeService,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-credentials,uc-grants")
		ic.currentMetastore = currentMetastoreResponse
		// Test listing
		err := resourcesMap["databricks_credential"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_credential[<unknown>] (id: test-service)"])
		// Test import
		err = resourcesMap["databricks_credential"].Import(ic, &resource{
			ID: "1234",
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: credential/1234)"])
	})
}

func TestStorageCredentials(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/storage-credentials?",
			Response: sdk_uc.ListStorageCredentialsResponse{
				StorageCredentials: []sdk_uc.StorageCredentialInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-storage-credentials,uc-grants")
		ic.currentMetastore = currentMetastoreResponse
		// Test listing
		err := resourcesMap["databricks_storage_credential"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_storage_credential[<unknown>] (id: test)"])
		// Test import
		err = resourcesMap["databricks_storage_credential"].Import(ic, &resource{
			ID: "1234",
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: storage_credential/1234)"])
	})
}

func TestListRecipients(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/recipients?",
			Response: sdk_sharing.ListRecipientsResponse{
				Recipients: []sdk_sharing.RecipientInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-shares")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_recipient"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_recipient[<unknown>] (id: test)"])
	})
}

func TestVolumes(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-volumes,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tf_uc.ResourceVolume().ToResource().TestResourceData()
	d.SetId("vtest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_volume"].Import(ic, &resource{
		ID:   "vtest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: volume/vtest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	r := &resource{
		Attribute: "databricks_volume",
		Value:     "dbc",
		Data:      d,
	}
	shouldOmitFunc := resourcesMap["databricks_volume"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tf_uc.ResourceVolume().Schema
	assert.False(t, shouldOmitFunc(nil, "volume_type", scm["volume_type"], d, r))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d, r))
	d.Set("volume_type", "MANAGED")
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "volume_type", scm["volume_type"], d, r))
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d, r))
}

func TestSqlTables(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-tables,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tf_uc.ResourceSqlTable().ToResource().TestResourceData()
	d.SetId("ttest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_sql_table"].Import(ic, &resource{
		ID:   "ttest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: table/ttest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	r := &resource{
		Attribute: "databricks_sql_table",
		Value:     "ttest",
		Data:      d,
	}
	shouldOmitFunc := resourcesMap["databricks_sql_table"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tf_uc.ResourceSqlTable().Schema
	assert.False(t, shouldOmitFunc(nil, "table_type", scm["table_type"], d, r))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d, r))
	d.Set("table_type", "MANAGED")
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "table_type", scm["table_type"], d, r))
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d, r))
}

func TestRegisteredModels(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-models,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tf_uc.ResourceRegisteredModel().ToResource().TestResourceData()
	d.SetId("mtest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_registered_model"].Import(ic, &resource{
		ID:   "mtest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: model/mtest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	r := &resource{
		Attribute: "databricks_registered_model",
		Value:     "mtest",
		Data:      d,
	}
	shouldOmitFunc := resourcesMap["databricks_registered_model"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tf_uc.ResourceRegisteredModel().Schema
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d, r))
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d, r))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d, r))

	ic.currentMetastore = currentMetastoreResponse
	d.Set("storage_location", "s3://abc/"+currentMetastoreResponse.MetastoreId+"/models/123456")
	assert.True(t, shouldOmitFunc(ic, "storage_location", scm["storage_location"], d, r))
}

func TestAuxUcFunctions(t *testing.T) {
	// Metastore Assignment
	d := tf_uc.ResourceMetastoreAssignment().ToResource().TestResourceData()
	d.Set("workspace_id", 123)
	assert.Equal(t, "ws_123", resourcesMap["databricks_metastore_assignment"].Name(nil, d))
	r := &resource{
		Attribute: "databricks_metastore_assignment",
		Value:     "ws_123",
		Data:      d,
	}

	shouldOmitFunc := resourcesMap["databricks_metastore_assignment"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	d.Set("default_catalog_name", "")

	scm := tf_uc.ResourceMetastoreAssignment().Schema
	assert.True(t, shouldOmitFunc(nil, "default_catalog_name", scm["default_catalog_name"], d, r))
	assert.False(t, shouldOmitFunc(nil, "metastore_id", scm["metastore_id"], d, r))

	// Metastore
	d = tf_uc.ResourceMetastore().ToResource().TestResourceData()
	d.SetId("1234")
	r = &resource{
		Attribute: "databricks_metastore",
		Value:     "1234",
		Data:      d,
	}
	assert.Equal(t, "1234", resourcesMap["databricks_metastore"].Name(nil, d))
	d.Set("name", "test")
	assert.Equal(t, "test", resourcesMap["databricks_metastore"].Name(nil, d))

	shouldOmitFunc = resourcesMap["databricks_metastore"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm = tf_uc.ResourceMetastore().Schema
	assert.True(t, shouldOmitFunc(nil, "default_data_access_config_id", scm["default_data_access_config_id"], d, r))
	assert.True(t, shouldOmitFunc(nil, "owner", scm["owner"], d, r))
	d.Set("owner", "test")
	assert.False(t, shouldOmitFunc(nil, "owner", scm["owner"], d, r))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d, r))

	// Connections
	d = tf_uc.ResourceConnection().ToResource().TestResourceData()
	d.SetId("1234")
	assert.Equal(t, "1234", resourcesMap["databricks_connection"].Name(nil, d))
	d.Set("name", "test")
	d.Set("connection_type", "db")
	assert.Equal(t, "db_test", resourcesMap["databricks_connection"].Name(nil, d))

	// Catalogs
	d = tf_uc.ResourceCatalog().ToResource().TestResourceData()
	d.SetId("test")
	r = &resource{
		Attribute: "databricks_catalog",
		Value:     "test",
		Data:      d,
	}
	shouldOmitFunc = resourcesMap["databricks_catalog"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm = tf_uc.ResourceCatalog().Schema
	d.Set("isolation_mode", "OPEN")
	assert.True(t, shouldOmitFunc(nil, "isolation_mode", scm["isolation_mode"], d, r))
	d.Set("isolation_mode", "ISOLATED")
	assert.False(t, shouldOmitFunc(nil, "isolation_mode", scm["isolation_mode"], d, r))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d, r))
}

func TestImportGrants(t *testing.T) {
	ic := importContextForTest()

	s := ic.Resources["databricks_grants"].Schema
	d := tf_uc.ResourceGrants().ToResource().TestResourceData()
	id := "metastore/1234"
	d.SetId(id)
	d.MarkNewResource()
	r := &resource{ID: id, Data: d}
	err := resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)

	var pList tf_uc.PermissionsList
	common.DataToStructPointer(r.Data, s, &pList)
	assert.Empty(t, pList.Assignments)

	// Test with a filled user name and no owner
	ic.meUserName = "user@domain.com"
	d.Set("catalog", "1234")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 0, len(pList.Assignments))

	// Test with a filled user name and permissions
	r.AddExtraData("owner", "otheruser@domain.com")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 1, len(pList.Assignments))
	assert.Equal(t, ic.meUserName, pList.Assignments[0].Principal)
	assert.Equal(t, sortStringsCopy(grantsPrivilegesToAdd["catalog"]), sortStringsCopy(pList.Assignments[0].Privileges))

	// Test with a filled user name and permissions
	d.Set("metastore", "")
	d.Set("catalog", "test")
	pList.Assignments = []tf_uc.PrivilegeAssignment{
		{Principal: ic.meUserName, Privileges: []string{"USE_CATALOG", "USE_SCHEMA"}},
	}
	common.StructToData(pList, s, d)
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 1, len(pList.Assignments))
	assert.Equal(t, ic.meUserName, pList.Assignments[0].Principal)
	assert.Equal(t, sortStringsCopy(append([]string{"USE_CATALOG", "USE_SCHEMA"}, grantsPrivilegesToAdd["catalog"]...)),
		sortStringsCopy(pList.Assignments[0].Privileges))

	// Test with a filled user name and unsupported objects
	d.Set("catalog", "")
	d.Set("model", "test")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
}

func TestImportUcOnlineTable(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("uc-tables,uc-grants")
		ic.currentMetastore = currentMetastoreResponse

		otTableName := "main.tmp.tbl_ot"
		d := tf_uc.ResourceOnlineTable().ToResource().TestResourceData()
		ot := sdk_uc.OnlineTable{
			Name: otTableName,
			Spec: &sdk_uc.OnlineTableSpec{
				SourceTableFullName: "main.tmp.tbl",
				PrimaryKeyColumns:   []string{"id"},
			},
		}
		d.SetId(otTableName)
		d.MarkNewResource()
		scm := tf_uc.ResourceOnlineTable().Schema
		err := common.StructToData(ot, scm, d)
		require.NoError(t, err)

		err = resourcesMap["databricks_online_table"].Import(ic, &resource{
			ID:   otTableName,
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: main.tmp.tbl)"])
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: table/main.tmp.tbl_ot)"])
	})
}

func TestListSystemSchemasSuccess(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		currentMetastoreSuccess,
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.1/unity-catalog/metastores/%s/systemschemas?", currentMetastoreResponse.MetastoreId),
			Response: sdk_uc.ListSystemSchemasResponse{
				Schemas: []sdk_uc.SystemSchemaInfo{
					{
						Schema: "access",
						State:  "ENABLE_COMPLETED",
					},
					{
						Schema: "information_schema",
						State:  "ENABLE_COMPLETED",
					},
					{
						Schema: "marketplace",
						State:  "available",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-system-schemas")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_system_schema"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, len(ic.testEmits), 1)
	})
}

func TestListSystemSchemasErrorGetMetastore(t *testing.T) {
	ic := importContextForTest()
	err := resourcesMap["databricks_system_schema"].List(ic)
	assert.EqualError(t, err, "there is no UC metastore information")
}

func TestListSystemSchemasErrorListing(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.1/unity-catalog/metastores/%s/systemschemas?", currentMetastoreResponse.MetastoreId),
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_system_schema"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func TestListUcAllowListError(t *testing.T) {
	ic := importContextForTest()
	err := resourcesMap["databricks_artifact_allowlist"].List(ic)
	assert.EqualError(t, err, "there is no UC metastore information")
}

func TestListUcAllowListSuccess(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-artifact-allowlist")
	ic.currentMetastore = currentMetastoreResponse
	err := resourcesMap["databricks_artifact_allowlist"].List(ic)
	assert.NoError(t, err)
	assert.Equal(t, len(ic.testEmits), 3)
	// Test ignore function
	d := tf_uc.ResourceArtifactAllowlist().ToResource().TestResourceData()
	d.MarkNewResource()
	d.Set("id", "abc")
	res := ic.Importables["databricks_artifact_allowlist"].Ignore(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.True(t, res)
	assert.Contains(t, ic.ignoredResources, "databricks_artifact_allowlist. id=abc")
	// Test ignore function, with blocks
	err = common.StructToData(
		tf_uc.ArtifactAllowlistInfo{
			ArtifactType: "INIT_SCRIPT",
			ArtifactMatchers: []sdk_uc.ArtifactMatcher{
				{
					Artifact:  "/Volumes/inits",
					MatchType: "PREFIX_MATCH",
				},
			},
		},
		tf_uc.ResourceArtifactAllowlist().Schema, d)
	assert.NoError(t, err)
	res = ic.Importables["databricks_artifact_allowlist"].Ignore(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.False(t, res)
}

func TestStorageCredentialListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/storage-credentials?",
			Status:   200,
			Response: &sdk_uc.ListStorageCredentialsResponse{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_storage_credential"].List(ic)
		assert.NoError(t, err)
	})
}

func TestImportStorageCredentialGrants(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.1/unity-catalog/permissions/storage_credential/abc",
			Response: sdk_uc.GetPermissionsResponse{
				PrivilegeAssignments: []sdk_uc.PrivilegeAssignment{
					{
						Principal:  "principal",
						Privileges: []sdk_uc.Privilege{"CREATE EXTERNAL LOCATION"},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := tf_uc.ResourceStorageCredential().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_storage_credential"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

func TestExternalLocationListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/external-locations?",
			Status:   200,
			Response: &sdk_uc.ListExternalLocationsResponse{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_external_location"].List(ic)
		assert.NoError(t, err)
	})
}

func TestImportExternalLocationGrants(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.1/unity-catalog/permissions/external-locations/abc",
			Response: sdk_uc.GetPermissionsResponse{
				PrivilegeAssignments: []sdk_uc.PrivilegeAssignment{
					{
						Principal:  "principal",
						Privileges: []sdk_uc.Privilege{"ALL PRIVILEGES"},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := tf_uc.ResourceExternalLocation().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_external_location"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

func TestListMetastores(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		metastores := []sdk_uc.MetastoreInfo{
			{
				Name:        "test",
				MetastoreId: "1234",
			},
		}
		ma.GetMockAccountMetastoresAPI().EXPECT().
			List(mock.Anything).
			Return(createIteratorFromSlice(metastores))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "uc-metastores")
		err := resourcesMap["databricks_metastore"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_metastore[<unknown>] (id: 1234)"])
	})
}

func TestImportVectorSearchEndpointList(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/vector-search/endpoints?",
			Response: sdk_vs.ListEndpointResponse{
				Endpoints: []sdk_vs.EndpointInfo{
					{
						Name:                 "test",
						LastUpdatedTimestamp: 1234567890,
					},
					{
						Name:                 "test2",
						LastUpdatedTimestamp: 2234567890,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("vector-search")
		ic.currentMetastore = currentMetastoreResponse

		err := resourcesMap["databricks_vector_search_endpoint"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_vector_search_endpoint[<unknown>] (id: test)"])
		assert.True(t, ic.testEmits["databricks_vector_search_endpoint[<unknown>] (id: test2)"])
	})
}

func TestImportVectorSearchEndpoint(t *testing.T) {
	vseName := "test"
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/vector-search/indexes?endpoint_name=test",
			Response: sdk_vs.ListVectorIndexesResponse{
				VectorIndexes: []sdk_vs.MiniVectorIndex{
					{
						Name: "idx1",
					},
					{
						Name: "idx2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("vector-search")
		ic.currentMetastore = currentMetastoreResponse

		d := tf_vs.ResourceVectorSearchEndpoint().ToResource().TestResourceData()
		vse := sdk_vs.EndpointInfo{
			Name: vseName,
		}
		d.SetId(vseName)
		d.MarkNewResource()
		scm := tf_vs.ResourceVectorSearchEndpoint().Schema
		err := common.StructToData(vse, scm, d)
		require.NoError(t, err)

		err = resourcesMap["databricks_vector_search_endpoint"].Import(ic, &resource{
			ID:   vseName,
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_vector_search_index[<unknown>] (id: idx1)"])
		assert.True(t, ic.testEmits["databricks_vector_search_index[<unknown>] (id: idx2)"])
	})
}

func TestImportVectorSearchIndex(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("vector-search,uc-tables,uc-grants,model-serving")
		ic.currentMetastore = currentMetastoreResponse

		vsiName := "main.tmp.vsi"
		d := tf_vs.ResourceVectorSearchIndex().ToResource().TestResourceData()
		ot := sdk_vs.VectorIndex{
			Name:         vsiName,
			PrimaryKey:   "id",
			EndpointName: "vs",
			DeltaSyncIndexSpec: &sdk_vs.DeltaSyncVectorIndexSpecResponse{
				SourceTable: "main.tmp.tbl",
				EmbeddingSourceColumns: []sdk_vs.EmbeddingSourceColumn{
					{
						Name:                       "col1",
						EmbeddingModelEndpointName: "test",
					},
				},
			},
			DirectAccessIndexSpec: &sdk_vs.DirectAccessVectorIndexSpec{
				EmbeddingSourceColumns: []sdk_vs.EmbeddingSourceColumn{
					{
						Name:                       "col1",
						EmbeddingModelEndpointName: "test",
					},
				},
			},
		}
		d.SetId(vsiName)
		d.MarkNewResource()
		scm := tf_vs.ResourceVectorSearchIndex().Schema
		err := common.StructToData(ot, scm, d)
		require.NoError(t, err)

		err = resourcesMap["databricks_vector_search_index"].Import(ic, &resource{
			ID:   vsiName,
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 4, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: table/main.tmp.vsi)"])
		assert.True(t, ic.testEmits["databricks_vector_search_endpoint[<unknown>] (id: vs)"])
		assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: main.tmp.tbl)"])
		assert.True(t, ic.testEmits["databricks_model_serving[<unknown>] (id: test)"])
	})
}

func TestImportVectorSearchIndexWithDirectAccess(t *testing.T) {
	d := tf_vs.ResourceVectorSearchIndex().ToResource().TestResourceData()
	d.SetId("test-index")
	d.Set("endpoint_name", "test-endpoint")

	// Set up direct access index spec
	directAccess := map[string]any{
		"embedding_source_columns": []map[string]any{
			{
				"name":                          "embedding_column",
				"embedding_model_endpoint_name": "embedding-model",
			},
		},
	}
	d.Set("direct_access_index_spec", []map[string]any{directAccess})

	ic := importContextForTest()
	ic.enableServices("vector-search,uc-tables,uc-grants,model-serving")

	r := &resource{
		Resource: "databricks_vector_search_index",
		ID:       "test-index",
		Data:     d,
	}

	err := resourcesMap["databricks_vector_search_index"].Import(ic, r)
	assert.NoError(t, err)

	// Check that the right resources were emitted
	assert.True(t, ic.testEmits["databricks_vector_search_endpoint[<unknown>] (id: test-endpoint)"])
	assert.True(t, ic.testEmits["databricks_model_serving[<unknown>] (id: embedding-model)"])
}

func TestRfaAccessRequestDestinationsName(t *testing.T) {
	ic := importContextForTest()
	d := ic.PluginFrameworkResources["databricks_rfa_access_request_destinations"]
	assert.NotNil(t, d, "databricks_rfa_access_request_destinations should be registered")

	// Test name generation from ID
	testData := map[string]string{
		"CATALOG,main":                "rfa_catalog_main",
		"SCHEMA,main.default":         "rfa_schema_main_default",
		"TABLE,main.default.my_table": "rfa_table_main_default_my_table",
	}

	for id, expectedName := range testData {
		// Create a mock resource data with the ID
		// Since this is a Plugin Framework resource, we need to use a different approach
		// For now, just test the Name function logic directly
		parts := strings.Split(id, ",")
		if len(parts) == 2 {
			securableType := strings.ToLower(parts[0])
			fullName := parts[1]
			safeName := nameNormalizationRegex.ReplaceAllString(fullName, "_")
			generatedName := fmt.Sprintf("rfa_%s_%s", securableType, safeName)
			assert.Equal(t, expectedName, generatedName, "Name generation for ID %s", id)
		}
	}
}

func TestEmitRfaAccessRequestDestinations_WithDestinations(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockRfaAPI().EXPECT().GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "CATALOG",
			FullName:      "main",
		}).Return(&sdk_uc.AccessRequestDestinations{
			Destinations: []sdk_uc.NotificationDestination{
				{
					DestinationId:   "admin@example.com",
					DestinationType: sdk_uc.DestinationTypeEmail,
				},
				{
					DestinationId:   "slack-dest-123",
					DestinationType: sdk_uc.DestinationTypeSlack,
				},
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-rfa")

		ic.emitRfaAccessRequestDestinations("CATALOG", "main")

		// Should emit the RFA resource
		assert.True(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: CATALOG,main)"],
			"Should emit RFA destinations when they exist")
	})
}

func TestEmitRfaAccessRequestDestinations_NoDestinations(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockRfaAPI().EXPECT().GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "CATALOG",
			FullName:      "test_catalog",
		}).Return(&sdk_uc.AccessRequestDestinations{
			Destinations: []sdk_uc.NotificationDestination{},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-rfa")

		ic.emitRfaAccessRequestDestinations("CATALOG", "test_catalog")

		// Should NOT emit when no destinations configured
		assert.False(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: CATALOG,test_catalog)"],
			"Should not emit RFA destinations when none exist")
	})
}

func TestEmitRfaAccessRequestDestinations_ServiceNotEnabled(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		// No RFA API calls should be made when service is not enabled
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		// Note: NOT enabling the 'uc-rfa' service

		ic.emitRfaAccessRequestDestinations("CATALOG", "main")

		// Should NOT emit when service not enabled
		assert.False(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: CATALOG,main)"],
			"Should not emit RFA destinations when service not enabled")
	})
}

func TestEmitRfaAccessRequestDestinations_ApiError(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockRfaAPI().EXPECT().GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "CATALOG",
			FullName:      "missing",
		}).Return(nil, errors.New("RESOURCE_DOES_NOT_EXIST: Catalog does not exist"))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-rfa")

		// Should not panic or error, just log and return
		ic.emitRfaAccessRequestDestinations("CATALOG", "missing")

		// Should NOT emit when API returns error
		assert.False(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: CATALOG,missing)"],
			"Should not emit RFA destinations when API returns error")
	})
}

func TestEmitRfaAccessRequestDestinations_MultipleSecurableTypes(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		rfaAPI := mw.GetMockRfaAPI().EXPECT()
		rfaAPI.GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "CATALOG",
			FullName:      "main",
		}).Return(&sdk_uc.AccessRequestDestinations{
			Destinations: []sdk_uc.NotificationDestination{
				{DestinationId: "admin@example.com", DestinationType: sdk_uc.DestinationTypeEmail},
			},
		}, nil)
		rfaAPI.GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "SCHEMA",
			FullName:      "main.default",
		}).Return(&sdk_uc.AccessRequestDestinations{
			Destinations: []sdk_uc.NotificationDestination{
				{DestinationId: "data-team@example.com", DestinationType: sdk_uc.DestinationTypeEmail},
			},
		}, nil)
		rfaAPI.GetAccessRequestDestinations(mock.Anything, sdk_uc.GetAccessRequestDestinationsRequest{
			SecurableType: "TABLE",
			FullName:      "main.default.users",
		}).Return(&sdk_uc.AccessRequestDestinations{
			Destinations: []sdk_uc.NotificationDestination{
				{DestinationId: "table-owner@example.com", DestinationType: sdk_uc.DestinationTypeEmail},
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-rfa")

		ic.emitRfaAccessRequestDestinations("CATALOG", "main")
		ic.emitRfaAccessRequestDestinations("SCHEMA", "main.default")
		ic.emitRfaAccessRequestDestinations("TABLE", "main.default.users")

		// Should emit all three RFA resources
		assert.True(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: CATALOG,main)"])
		assert.True(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: SCHEMA,main.default)"])
		assert.True(t, ic.testEmits["databricks_rfa_access_request_destinations[<unknown>] (id: TABLE,main.default.users)"])
	})
}

func TestCreateIsMatchingSecurableType(t *testing.T) {
	// Test the IsValidApproximation function for different securable types
	testCases := []struct {
		name         string
		expectedType string
		actualType   string
		shouldMatch  bool
	}{
		{
			name:         "CATALOG matches CATALOG",
			expectedType: "CATALOG",
			actualType:   "CATALOG",
			shouldMatch:  true,
		},
		{
			name:         "CATALOG does not match SCHEMA",
			expectedType: "CATALOG",
			actualType:   "SCHEMA",
			shouldMatch:  false,
		},
		{
			name:         "TABLE matches TABLE",
			expectedType: "TABLE",
			actualType:   "TABLE",
			shouldMatch:  true,
		},
		{
			name:         "VOLUME does not match TABLE",
			expectedType: "VOLUME",
			actualType:   "TABLE",
			shouldMatch:  false,
		},
		{
			name:         "CREDENTIAL matches CREDENTIAL",
			expectedType: "CREDENTIAL",
			actualType:   "CREDENTIAL",
			shouldMatch:  true,
		},
		{
			name:         "EXTERNAL_LOCATION matches EXTERNAL_LOCATION",
			expectedType: "EXTERNAL_LOCATION",
			actualType:   "EXTERNAL_LOCATION",
			shouldMatch:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ic := importContextForTest()

			// Create a mock resource with DataWrapper that returns the securable type
			r := &resource{
				Resource: "databricks_rfa_access_request_destinations",
				ID:       "TEST,test_object",
				DataWrapper: &testDataWrapper{
					securableType: tc.actualType,
				},
			}

			// Create the validation function
			validationFunc := createIsMatchingSecurableType(tc.expectedType)

			// Test the validation
			result := validationFunc(ic, r, nil, "securable.full_name")

			assert.Equal(t, tc.shouldMatch, result,
				"Expected match=%v for %s vs %s", tc.shouldMatch, tc.expectedType, tc.actualType)
		})
	}
}

func TestCreateIsMatchingSecurableType_NilDataWrapper(t *testing.T) {
	ic := importContextForTest()

	r := &resource{
		Resource:    "databricks_rfa_access_request_destinations",
		ID:          "TEST,test_object",
		DataWrapper: nil,
	}

	validationFunc := createIsMatchingSecurableType("CATALOG")
	result := validationFunc(ic, r, nil, "securable.full_name")

	assert.False(t, result, "Should return false when DataWrapper is nil")
}

func TestCreateIsMatchingSecurableType_InvalidType(t *testing.T) {
	ic := importContextForTest()

	r := &resource{
		Resource: "databricks_rfa_access_request_destinations",
		ID:       "TEST,test_object",
		DataWrapper: &testDataWrapper{
			securableType: 123, // Not a string
		},
	}

	validationFunc := createIsMatchingSecurableType("CATALOG")
	result := validationFunc(ic, r, nil, "securable.full_name")

	assert.False(t, result, "Should return false when securable type is not a string")
}

// testDataWrapper is a minimal mock for testing IsValidApproximation
type testDataWrapper struct {
	securableType interface{}
}

func (t *testDataWrapper) Get(key string) interface{} {
	if key == "securable.type" {
		return t.securableType
	}
	return nil
}

func (t *testDataWrapper) GetOk(key string) (interface{}, bool) {
	if key == "securable.type" {
		return t.securableType, t.securableType != nil
	}
	return nil, false
}

func (t *testDataWrapper) Id() string {
	return "test-id"
}

func (t *testDataWrapper) SetId(id string) {}

func (t *testDataWrapper) Set(key string, value interface{}) error {
	return nil
}

func (t *testDataWrapper) GetSchema() SchemaWrapper {
	return nil
}

func (t *testDataWrapper) IsPluginFramework() bool {
	return true
}

func (t *testDataWrapper) GetTypedStruct(ctx context.Context, target interface{}) error {
	return nil
}
