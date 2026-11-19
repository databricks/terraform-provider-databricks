package exporter

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	app_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/app"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportAppEmitsUnityCatalogDependencies(t *testing.T) {
	ctx := context.Background()
	ic := importContextForTest()
	ic.Context = ctx
	services, _ := ic.allServicesAndListing()
	ic.enableServices(services)

	app := apps.App{
		Name: "test-app",
		Resources: []apps.AppResource{
			{Name: "table", UcSecurable: &apps.AppResourceUcSecurable{SecurableType: "TABLE", SecurableFullName: "main.default.events"}},
			{Name: "connection", UcSecurable: &apps.AppResourceUcSecurable{SecurableType: "CONNECTION", SecurableFullName: "salesforce"}},
			{Name: "volume", UcSecurable: &apps.AppResourceUcSecurable{SecurableType: "VOLUME", SecurableFullName: "main.default.checkpoints"}},
		},
		TelemetryExportDestinations: []apps.TelemetryExportDestination{{
			UnityCatalog: &apps.UnityCatalog{
				LogsTable: "main.telemetry.logs", MetricsTable: "main.telemetry.metrics", TracesTable: "main.telemetry.traces",
			},
		}},
	}
	var tfApp app_resource.AppResource
	diags := converters.GoSdkToTfSdkStruct(ctx, app, &tfApp)
	require.False(t, diags.HasError(), diags.Errors())
	wrapper, state := newPluginFrameworkTestState(t, ic, ctx, "databricks_app", app.Name)
	require.Empty(t, state.SetAttribute(ctx, path.Root("name"), tfApp.Name).Errors())
	require.Empty(t, state.SetAttribute(ctx, path.Root("resources"), tfApp.Resources).Errors())
	require.Empty(t, state.SetAttribute(ctx, path.Root("telemetry_export_destinations"), tfApp.TelemetryExportDestinations).Errors())

	err := importApp(ic, &resource{Resource: "databricks_app", ID: app.Name, Name: "test_app", DataWrapper: wrapper})
	require.NoError(t, err)
	for _, expected := range []string{
		"databricks_sql_table[<unknown>] (id: main.default.events)",
		"databricks_connection[<unknown>] (id: salesforce)",
		"databricks_volume[<unknown>] (id: main.default.checkpoints)",
		"databricks_sql_table[<unknown>] (id: main.telemetry.logs)",
		"databricks_sql_table[<unknown>] (id: main.telemetry.metrics)",
		"databricks_sql_table[<unknown>] (id: main.telemetry.traces)",
	} {
		assert.Contains(t, ic.testEmits, expected)
	}
}

func TestAppExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyAppsSettingsCustomTemplates,
		emptyDatabaseInstances,
		{
			Method:   "GET",
			Resource: "/api/2.0/apps?",
			Response: apps.ListAppsResponse{
				Apps: []apps.App{
					{
						Name:        "test-app",
						Description: "Test app",
						Resources: []apps.AppResource{
							{
								Name: "sql-warehouse",
								SqlWarehouse: &apps.AppResourceSqlWarehouse{
									Id:         "warehouse-123",
									Permission: "CAN_MANAGE",
								},
							},
							{
								Name: "serving-endpoint",
								ServingEndpoint: &apps.AppResourceServingEndpoint{
									Name:       "endpoint-abc",
									Permission: "CAN_QUERY",
								},
							},
							{
								Name: "job",
								Job: &apps.AppResourceJob{
									Id:         "job-456",
									Permission: "CAN_VIEW",
								},
							},
							{
								Name: "secret",
								Secret: &apps.AppResourceSecret{
									Scope:      "my-scope",
									Key:        "my-key",
									Permission: "READ",
								},
							},
							{
								Name: "uc-volume",
								UcSecurable: &apps.AppResourceUcSecurable{
									SecurableType:     "VOLUME",
									SecurableFullName: "catalog.schema.my_volume",
									Permission:        "READ_VOLUME",
								},
							},
							{
								Name: "database",
								Database: &apps.AppResourceDatabase{
									InstanceName: "my-db-instance",
									DatabaseName: "my_database",
									Permission:   "CAN_CONNECT_AND_CREATE",
								},
							},
						},
						BudgetPolicyId: "budget-789",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/apps/test-app?",
			Response: apps.App{
				Name:        "test-app",
				Description: "Test app",
				Resources: []apps.AppResource{
					{
						Name: "sql-warehouse",
						SqlWarehouse: &apps.AppResourceSqlWarehouse{
							Id:         "warehouse-123",
							Permission: "CAN_MANAGE",
						},
					},
					{
						Name: "serving-endpoint",
						ServingEndpoint: &apps.AppResourceServingEndpoint{
							Name:       "endpoint-abc",
							Permission: "CAN_QUERY",
						},
					},
					{
						Name: "job",
						Job: &apps.AppResourceJob{
							Id:         "job-456",
							Permission: "CAN_VIEW",
						},
					},
					{
						Name: "secret",
						Secret: &apps.AppResourceSecret{
							Scope:      "my-scope",
							Key:        "my-key",
							Permission: "READ",
						},
					},
					{
						Name: "uc-volume",
						UcSecurable: &apps.AppResourceUcSecurable{
							SecurableType:     "VOLUME",
							SecurableFullName: "catalog.schema.my_volume",
							Permission:        "READ_VOLUME",
						},
					},
					{
						Name: "database",
						Database: &apps.AppResourceDatabase{
							InstanceName: "my-db-instance",
							DatabaseName: "my_database",
							Permission:   "CAN_CONNECT_AND_CREATE",
						},
					},
				},
				BudgetPolicyId: "budget-789",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances/my-db-instance?",
			Response: database.DatabaseInstance{
				Name:     "my-db-instance",
				Capacity: "CU_2",
				State:    "AVAILABLE",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/permissions/database-instances/my-db-instance?",
			ReuseRequest: true,
			Response: entity.PermissionsEntity{
				ObjectType:        "database-instances",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/apps/test-app",
			Response: entity.PermissionsEntity{
				ObjectType:        "apps",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.Config.WorkspaceID = testProviderWorkspaceID
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.enableServices("apps,lakebase")
		ic.enableListing("apps")
		ic.Directory = tmpDir
		ic.noFormat = true

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the app and its dependencies were generated in the Terraform code
		content, err := os.ReadFile(tmpDir + "/apps.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the app resource is generated
		assert.Contains(t, contentStr, `resource "databricks_app" "test_app"`)
		assert.Contains(t, contentStr, `name = "test-app"`)
		assert.Contains(t, contentStr, `description = "Test app"`)

		// Check that database instance reference is generated
		assert.Contains(t, contentStr, `instance_name = databricks_database_instance.my_db_instance.name`)

		// Check that the database instance resource itself is generated (in lakebase.tf)
		lakebaseContent, err := os.ReadFile(tmpDir + "/lakebase.tf")
		assert.NoError(t, err)
		lakebaseStr := normalizeWhitespace(string(lakebaseContent))
		assert.Contains(t, lakebaseStr, `resource "databricks_database_instance" "my_db_instance"`)
		assert.Contains(t, lakebaseStr, `name = "my-db-instance"`)
	})
}

func TestAppExportWithEmptyDescription(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyAppsSettingsCustomTemplates,
		emptyDatabaseInstances,
		{
			Method:   "GET",
			Resource: "/api/2.0/apps?",
			Response: apps.ListAppsResponse{
				Apps: []apps.App{
					{
						Name:           "data-intake",
						Description:    "", // Empty description should be omitted
						BudgetPolicyId: "4635ae18-e8d8-4528-98d3-05805c7e6308",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/apps/data-intake?",
			Response: apps.App{
				Name:           "data-intake",
				Description:    "", // Empty description should be omitted
				BudgetPolicyId: "4635ae18-e8d8-4528-98d3-05805c7e6308",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/apps/data-intake",
			Response: entity.PermissionsEntity{
				ObjectType:        "apps",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.Config.WorkspaceID = testProviderWorkspaceID
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.enableServices("apps")
		ic.enableListing("apps")
		ic.Directory = tmpDir
		ic.noFormat = true

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the app was generated without the empty description field
		content, err := os.ReadFile(tmpDir + "/apps.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the app resource is generated
		assert.Contains(t, contentStr, `resource "databricks_app" "data_intake"`)
		assert.Contains(t, contentStr, `name = "data-intake"`)
		assert.Contains(t, contentStr, `budget_policy_id = "4635ae18-e8d8-4528-98d3-05805c7e6308"`)

		// The empty description should NOT be present
		assert.NotContains(t, contentStr, `description = ""`)
	})
}

func TestAppsSettingsCustomTemplateExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.0/apps?",
			Response: apps.ListAppsResponse{
				Apps: []apps.App{},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/apps-settings/templates?",
			Response: apps.ListCustomTemplatesResponse{
				Templates: []apps.CustomTemplate{
					{
						Name:        "my-custom-template",
						Description: "Test template",
						GitRepo:     "https://github.com/example/repo.git",
						GitProvider: "github",
						Path:        "templates/app",
						Creator:     "user@example.com",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/apps-settings/templates/my-custom-template?",
			Response: apps.CustomTemplate{
				Name:        "my-custom-template",
				Description: "Test template",
				GitRepo:     "https://github.com/example/repo.git",
				GitProvider: "github",
				Path:        "templates/app",
				Creator:     "user@example.com",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/apps/templates/my-custom-template",
			Response: entity.PermissionsEntity{
				ObjectType:        "apps/templates",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.Config.WorkspaceID = testProviderWorkspaceID
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.enableServices("apps")
		ic.enableListing("apps")
		ic.Directory = tmpDir
		ic.noFormat = true

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the custom template was generated in the Terraform code
		content, err := os.ReadFile(tmpDir + "/apps.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the custom template resource is generated
		assert.Contains(t, contentStr, `resource "databricks_apps_settings_custom_template" "my_custom_template"`)
		assert.Contains(t, contentStr, `name = "my-custom-template"`)
		assert.Contains(t, contentStr, `description = "Test template"`)
		assert.Contains(t, contentStr, `git_repo = "https://github.com/example/repo.git"`)
		assert.Contains(t, contentStr, `git_provider = "github"`)
		assert.Contains(t, contentStr, `path = "templates/app"`)
	})
}

// TODO: Add test for databricks_custom_app_integration export
// The test needs proper account-level authentication mocking which requires
// more investigation into how the OAuth2 API authenticates in test mode
