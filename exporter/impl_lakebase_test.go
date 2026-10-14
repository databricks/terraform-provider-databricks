package exporter

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	sdkduration "github.com/databricks/databricks-sdk-go/common/types/duration"
	sdktime "github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// emptyPostgresProjectsFixture is a reusable fixture for listing Postgres projects (empty list).
// Use it in tests that enable lakebase listing so listPostgresProjects does not hit an unmocked endpoint.
var emptyPostgresProjectsFixture = qa.HTTPFixture{
	Method:   "GET",
	Resource: "/api/2.0/postgres/projects?",
	Response: postgres.ListProjectsResponse{
		Projects: []postgres.Project{},
	},
}

var emptyPostgresCatalogsFixture = qa.HTTPFixture{
	Method:   "GET",
	Resource: "/api/2.1/unity-catalog/catalogs?max_results=0",
	Response: catalog.ListCatalogsResponse{
		Catalogs: []catalog.CatalogInfo{},
	},
}

func postgresCatalogsFixture(catalogs []catalog.CatalogInfo) qa.HTTPFixture {
	return qa.HTTPFixture{
		Method:       "GET",
		Resource:     "/api/2.1/unity-catalog/catalogs?max_results=0",
		ReuseRequest: true,
		Response: catalog.ListCatalogsResponse{
			Catalogs: catalogs,
		},
	}
}

func emptyPostgresDatabasesFixture(branchName string) qa.HTTPFixture {
	return qa.HTTPFixture{
		Method:   "GET",
		Resource: fmt.Sprintf("/api/2.0/postgres/%s/databases?", branchName),
		Response: postgres.ListDatabasesResponse{
			Databases: []postgres.Database{},
		},
	}
}

func emptyPostgresRolesFixture(branchName string) qa.HTTPFixture {
	return qa.HTTPFixture{
		Method:   "GET",
		Resource: fmt.Sprintf("/api/2.0/postgres/%s/roles?", branchName),
		Response: postgres.ListRolesResponse{
			Roles: []postgres.Role{},
		},
	}
}

// postgresNameWrapper is a minimal ResourceDataWrapper for testing Postgres name generation.
type postgresNameWrapper struct {
	name string
	id   string
}

func (w *postgresNameWrapper) GetOk(key string) (interface{}, bool) {
	if key == "name" {
		return w.name, w.name != ""
	}
	return nil, false
}

func (w *postgresNameWrapper) Get(key string) interface{} {
	if key == "name" {
		return w.name
	}
	return nil
}

func (w *postgresNameWrapper) Id() string { return w.id }

func (w *postgresNameWrapper) SetId(string) {}

func (w *postgresNameWrapper) Set(string, interface{}) error { return nil }

func (w *postgresNameWrapper) GetSchema() SchemaWrapper { return nil }

func (w *postgresNameWrapper) IsPluginFramework() bool { return true }

func (w *postgresNameWrapper) GetTypedStruct(context.Context, interface{}) error { return nil }

func TestPostgresResourceNameFromPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"projects/my-app", "my_app"},
		{"projects/my_project", "my_project"},
		{"projects/my-app/branches/main", "my_app_main"},
		{"projects/demo/branches/dev", "demo_dev"},
		{"projects/my-app/branches/main/endpoints/primary", "my_app_main_primary"},
		{"projects/p1/branches/b2/endpoints/e3", "p1_b2_e3"},
		{"short", "short"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := postgresResourceNameFromPath(tt.path)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPostgresProjectNameUnified(t *testing.T) {
	ic := importContextForTest()
	wrapper := &postgresNameWrapper{name: "projects/my-app", id: "projects/my-app"}
	name := resourcesMap["databricks_postgres_project"].NameUnified(ic, wrapper)
	assert.Equal(t, "my_app", name)
}

func TestPostgresBranchNameUnified(t *testing.T) {
	ic := importContextForTest()
	wrapper := &postgresNameWrapper{name: "projects/my-app/branches/main", id: "projects/my-app/branches/main"}
	name := resourcesMap["databricks_postgres_branch"].NameUnified(ic, wrapper)
	assert.Equal(t, "my_app_main", name)
}

func TestPostgresEndpointNameUnified(t *testing.T) {
	ic := importContextForTest()
	wrapper := &postgresNameWrapper{name: "projects/my-app/branches/main/endpoints/primary", id: "projects/my-app/branches/main/endpoints/primary"}
	name := resourcesMap["databricks_postgres_endpoint"].NameUnified(ic, wrapper)
	assert.Equal(t, "my_app_main_primary", name)
}

func TestPostgresRemainingResourceNamesUnified(t *testing.T) {
	ic := importContextForTest()
	tests := []struct {
		resource string
		path     string
		want     string
	}{
		{"databricks_postgres_catalog", "catalogs/app_catalog", "app_catalog"},
		{"databricks_postgres_database", "projects/my-app/branches/main/databases/app-db", "my_app_main_app_db"},
		{"databricks_postgres_role", "projects/my-app/branches/main/roles/app-owner", "my_app_main_app_owner"},
		{"databricks_postgres_data_api", "projects/my-app/branches/main/databases/app-db/data-api", "my_app_main_app_db_data_api"},
		{"databricks_postgres_cdf_config", "projects/my-app/branches/main/databases/app-db/cdf-configs/public", "my_app_main_app_db_public"},
		{"databricks_postgres_synced_table", "synced_tables/app_catalog.default.users_synced", "app_catalog_default_users_synced"},
	}
	for _, tt := range tests {
		t.Run(tt.resource, func(t *testing.T) {
			wrapper := &postgresNameWrapper{name: tt.path, id: tt.path}
			name := resourcesMap[tt.resource].NameUnified(ic, wrapper)
			assert.Equal(t, tt.want, name)
		})
	}
}

func TestPostgresProjectImport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches?",
			Response: postgres.ListBranchesResponse{
				Branches: []postgres.Branch{
					{Name: "projects/my-app/branches/main"},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakebase,access")
		ic.meAdmin = true
		ic.testEmits = make(map[string]bool)

		r := &resource{
			Resource:    "databricks_postgres_project",
			ID:          "projects/my-app",
			Name:        "my_app",
			DataWrapper: &postgresNameWrapper{name: "projects/my-app", id: "projects/my-app"},
		}
		err := resourcesMap["databricks_postgres_project"].Import(ic, r)
		assert.NoError(t, err)
		assert.True(t, ic.testEmits["databricks_postgres_branch[<unknown>] (id: projects/my-app/branches/main)"])
		assert.True(t, ic.testEmits["databricks_permissions[postgres_project_my_app] (id: /database-projects/my-app)"])
	})
}

func TestPostgresBranchImport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/endpoints?",
			Response: postgres.ListEndpointsResponse{
				Endpoints: []postgres.Endpoint{
					{Name: "projects/my-app/branches/main/endpoints/primary"},
				},
			},
		},
		emptyPostgresDatabasesFixture("projects/my-app/branches/main"),
		emptyPostgresRolesFixture("projects/my-app/branches/main"),
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakebase")
		ic.testEmits = make(map[string]bool)

		r := &resource{
			Resource:    "databricks_postgres_branch",
			ID:          "projects/my-app/branches/main",
			DataWrapper: &postgresNameWrapper{name: "projects/my-app/branches/main", id: "projects/my-app/branches/main"},
		}
		err := resourcesMap["databricks_postgres_branch"].Import(ic, r)
		assert.NoError(t, err)
		assert.True(t, ic.testEmits["databricks_postgres_endpoint[<unknown>] (id: projects/my-app/branches/main/endpoints/primary)"])
	})
}

func TestPostgresProjectExport(t *testing.T) {
	sourceBranchTime := sdktime.New(time.Date(2026, 8, 22, 11, 35, 52, 0, time.UTC))
	historyRetentionDuration := sdkduration.New(168 * time.Hour)
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		postgresCatalogsFixture([]catalog.CatalogInfo{
			{
				Name:        "app_catalog",
				CatalogType: catalog.CatalogTypeManagedOnlineCatalog,
			},
		}),
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances?",
			Response: database.ListDatabaseInstancesResponse{
				DatabaseInstances: []database.DatabaseInstance{},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects?",
			Response: postgres.ListProjectsResponse{
				Projects: []postgres.Project{
					{Name: "projects/my-app"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app?",
			Response: postgres.Project{
				Name: "projects/my-app",
				Status: &postgres.ProjectStatus{
					DisplayName:              "My Project",
					HistoryRetentionDuration: historyRetentionDuration,
					PgVersion:                17,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches?",
			Response: postgres.ListBranchesResponse{
				Branches: []postgres.Branch{
					{Name: "projects/my-app/branches/main", Parent: "projects/my-app"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main?",
			Response: postgres.Branch{
				Name:   "projects/my-app/branches/main",
				Parent: "projects/my-app",
				Status: &postgres.BranchStatus{
					BranchId:         "main",
					IsProtected:      true,
					SourceBranchTime: sourceBranchTime,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/endpoints?",
			Response: postgres.ListEndpointsResponse{
				Endpoints: []postgres.Endpoint{
					{Name: "projects/my-app/branches/main/endpoints/primary", Parent: "projects/my-app/branches/main"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/endpoints/primary?",
			Response: postgres.Endpoint{
				Name:   "projects/my-app/branches/main/endpoints/primary",
				Parent: "projects/my-app/branches/main",
				Status: &postgres.EndpointStatus{
					AutoscalingLimitMaxCu: 4,
					AutoscalingLimitMinCu: 1,
					EndpointId:            "primary",
					EndpointType:          postgres.EndpointTypeEndpointTypeReadWrite,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/databases?",
			Response: postgres.ListDatabasesResponse{
				Databases: []postgres.Database{
					{
						Name:   "projects/my-app/branches/main/databases/app-db",
						Parent: "projects/my-app/branches/main",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/databases/app-db?",
			Response: postgres.Database{
				Name:   "projects/my-app/branches/main/databases/app-db",
				Parent: "projects/my-app/branches/main",
				Status: &postgres.DatabaseDatabaseStatus{
					DatabaseId:       "app-db",
					PostgresDatabase: "app_db",
					Role:             "projects/my-app/branches/main/roles/app-owner",
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/databases/app-db/cdf-configs?",
			Response: postgres.ListCdfConfigsResponse{
				CdfConfigs: []postgres.CdfConfig{
					{
						Name:           "projects/my-app/branches/main/databases/app-db/cdf-configs/public",
						Catalog:        "main",
						Schema:         "app_replicated",
						PostgresSchema: "public",
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/postgres/projects/my-app/branches/main/databases/app-db/data-api?",
			ReuseRequest: true,
			Response: postgres.DataApi{
				Name:   "projects/my-app/branches/main/databases/app-db/data-api",
				Parent: "projects/my-app/branches/main/databases/app-db",
				Status: &postgres.DataApiDataApiStatus{
					DbSchemas:                []string{"public"},
					DbMaxRows:                1000,
					ServerCorsAllowedOrigins: []string{"https://app.example.com"},
					ServerTimingEnabled:      true,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/databases/app-db/cdf-configs/public?",
			Response: postgres.CdfConfig{
				Name:           "projects/my-app/branches/main/databases/app-db/cdf-configs/public",
				CdfConfigId:    "public",
				Catalog:        "main",
				Schema:         "app_replicated",
				PostgresSchema: "public",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/roles?",
			Response: postgres.ListRolesResponse{
				Roles: []postgres.Role{
					{
						Name:   "projects/my-app/branches/main/roles/app-owner",
						Parent: "projects/my-app/branches/main",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/projects/my-app/branches/main/roles/app-owner?",
			Response: postgres.Role{
				Name:   "projects/my-app/branches/main/roles/app-owner",
				Parent: "projects/my-app/branches/main",
				Status: &postgres.RoleRoleStatus{
					RoleId:          "app-owner",
					PostgresRole:    "app_owner",
					MembershipRoles: []postgres.RoleMembershipRole{postgres.RoleMembershipRoleDatabricksSuperuser},
					Attributes: &postgres.RoleAttributes{
						Createdb: true,
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/postgres/catalogs/app_catalog?",
			Response: postgres.Catalog{
				Name:      "catalogs/app_catalog",
				CatalogId: "app_catalog",
				Status: &postgres.CatalogCatalogStatus{
					Branch:           "projects/my-app/branches/main",
					PostgresDatabase: "app_db",
					Project:          "projects/my-app",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.Context = ctx
		ic.workspaceClient, _ = client.WorkspaceClient()
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("lakebase")
		ic.enableServices("lakebase")

		err := ic.Run()
		assert.NoError(t, err)

		content, err := os.ReadFile(tmpDir + "/lakebase.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		assert.Contains(t, contentStr, `resource "databricks_postgres_project" "my_app"`)
		assert.Contains(t, contentStr, `project_id = "my-app"`)
		assert.Contains(t, contentStr, `display_name = "My Project"`)
		assert.Contains(t, contentStr, `history_retention_duration = "168h0m0s"`)
		assert.NotContains(t, contentStr, `history_retention_duration = "\"168h0m0s\""`)
		assert.Contains(t, contentStr, `pg_version = 17`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_branch" "my_app_main"`)
		assert.Contains(t, contentStr, `branch_id = "main"`)
		assert.Contains(t, contentStr, `source_branch_time = "2026-08-22T11:35:52Z"`)
		assert.NotContains(t, contentStr, `source_branch_time = "\"2026-08-22T11:35:52Z\""`)
		assert.Contains(t, contentStr, `parent = databricks_postgres_project.my_app.name`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_endpoint" "my_app_main_primary"`)
		assert.Contains(t, contentStr, `endpoint_id = "primary"`)
		assert.Contains(t, contentStr, `autoscaling_limit_max_cu = 4`)
		assert.Contains(t, contentStr, `parent = databricks_postgres_branch.my_app_main.name`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_database" "my_app_main_app_db"`)
		assert.Contains(t, contentStr, `database_id = "app-db"`)
		assert.Contains(t, contentStr, `postgres_database = "app_db"`)
		assert.Contains(t, contentStr, `role = databricks_postgres_role.my_app_main_app_owner.name`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_role" "my_app_main_app_owner"`)
		assert.Contains(t, contentStr, `role_id = "app-owner"`)
		assert.Contains(t, contentStr, `postgres_role = "app_owner"`)
		assert.Contains(t, contentStr, `membership_roles = ["DATABRICKS_SUPERUSER"]`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_cdf_config" "my_app_main_app_db_public"`)
		assert.Contains(t, contentStr, `parent = databricks_postgres_database.my_app_main_app_db.name`)
		assert.Contains(t, contentStr, `postgres_schema = "public"`)
		assert.Contains(t, contentStr, `schema = "app_replicated"`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_data_api" "my_app_main_app_db_data_api"`)
		assert.Contains(t, contentStr, `db_schemas = ["public"]`)
		assert.Contains(t, contentStr, `db_max_rows = 1000`)
		assert.Contains(t, contentStr, `server_cors_allowed_origins = ["https://app.example.com"]`)
		assert.Contains(t, contentStr, `resource "databricks_postgres_catalog" "app_catalog"`)
		assert.Contains(t, contentStr, `catalog_id = "app_catalog"`)
		assert.Contains(t, contentStr, `branch = databricks_postgres_branch.my_app_main.name`)
	})
}

func TestDatabaseInstanceName(t *testing.T) {
	ic := importContextForTest()
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}, map[string]any{
		"id":   "test-instance",
		"name": "test-instance",
	})
	d.SetId("test-instance")

	// Create wrapper for the resource data
	wrapper := &SDKv2ResourceData{
		data:   d,
		schema: &schema.Resource{Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}}},
	}

	name := resourcesMap["databricks_database_instance"].NameUnified(ic, wrapper)
	assert.Equal(t, "test-instance", name)
}

func TestDatabaseInstanceImport(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("lakebase,access")
	ic.meAdmin = true
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}, map[string]any{
		"id":   "test-instance",
		"name": "test-instance",
	})
	d.SetId("test-instance")
	r := &resource{
		ID:   "test-instance",
		Name: "test-instance",
		Data: d,
	}
	err := resourcesMap["databricks_database_instance"].Import(ic, r)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 1)
	assert.True(t, ic.testEmits["databricks_permissions[database_instance_test-instance] (id: /database-instances/test-instance)"])
}

func TestDatabaseInstanceIgnore(t *testing.T) {
	ic := importContextForTest()

	// Test with empty name - should be ignored
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]any{})
	d.SetId("test-instance")
	r := &resource{
		ID:   "test-instance",
		Data: d,
	}
	ignore := resourcesMap["databricks_database_instance"].Ignore(ic, r)
	assert.True(t, ignore)

	// Test with valid name - should not be ignored
	d2 := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]any{
		"name": "test-instance",
	})
	d2.SetId("test-instance")
	r2 := &resource{
		ID:   "test-instance",
		Data: d2,
	}
	ignore2 := resourcesMap["databricks_database_instance"].Ignore(ic, r2)
	assert.False(t, ignore2)
}

func TestDatabaseInstanceExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyPostgresProjectsFixture,
		emptyPostgresCatalogsFixture,
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances?",
			Response: database.ListDatabaseInstancesResponse{
				DatabaseInstances: []database.DatabaseInstance{
					{
						Name:                      "prod-instance",
						Capacity:                  "CU_2",
						State:                     "AVAILABLE",
						NodeCount:                 2,
						EnableReadableSecondaries: true,
						UsagePolicyId:             "policy-123",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances/prod-instance?",
			Response: database.DatabaseInstance{
				Name:                               "prod-instance",
				Capacity:                           "CU_2",
				EffectiveCapacity:                  "CU_2",
				State:                              "AVAILABLE",
				NodeCount:                          2,
				EffectiveNodeCount:                 2,
				EnableReadableSecondaries:          true,
				EffectiveEnableReadableSecondaries: true,
				UsagePolicyId:                      "policy-123",
				EffectiveUsagePolicyId:             "policy-123",
				EffectiveCustomTags: []database.CustomTag{
					{
						Key:   "Environment",
						Value: "Production",
					},
					{
						Key:   "Team",
						Value: "DataPlatform",
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/permissions/database-instances/prod-instance?",
			ReuseRequest: true,
			Response: entity.PermissionsEntity{
				ObjectType:        "database-instances",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.Config.WorkspaceID = testProviderWorkspaceID
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("lakebase")
		ic.enableServices("lakebase")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the database instance was generated in the Terraform code
		content, err := os.ReadFile(tmpDir + "/lakebase.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with expected fields
		assert.Contains(t, contentStr, `resource "databricks_database_instance" "prod_instance"`)
		assert.Contains(t, contentStr, `name = "prod-instance"`)
		assert.Contains(t, contentStr, `capacity = "CU_2"`)
		// These simple-type fields are automatically exported from their effective_* counterparts
		assert.Contains(t, contentStr, `node_count = 2`)
		assert.Contains(t, contentStr, `enable_readable_secondaries = true`)
		assert.Contains(t, contentStr, `usage_policy_id = "policy-123"`)
		// Note: Complex types like custom_tags require deeper Plugin Framework integration
		// The conversion logic exists but wrapper.Set() validation is blocking it
		// This is documented in COMPLEX_TYPES_HANDLING.md as a known limitation
	})
}
