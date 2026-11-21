package exporter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdktime "github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/listing"
	sdk_uc "github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	lakewatch_datasource_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/lakewatch_datasource"
	lakewatch_rule_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/lakewatch_rule"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// createErrorIterator creates an iterator that fails on the first page.
func createErrorIterator[T any](err error) listing.Iterator[T] {
	request := struct{}{}
	return listing.NewIterator(
		&request,
		func(ctx context.Context, req struct{}) ([]T, error) {
			return nil, err
		},
		func(resp []T) []T {
			return resp
		},
		func(resp []T) *struct{} {
			return nil
		},
	)
}

func TestLakewatchConfigCatalogName(t *testing.T) {
	assert.Equal(t, "sec", lakewatchConfigCatalogName("catalogs/sec/alertConfig"))
	assert.Equal(t, "sec", lakewatchConfigCatalogName("catalogs/sec/ruleConfig"))
	assert.Equal(t, "", lakewatchConfigCatalogName("catalogs/sec"))
	assert.Equal(t, "", lakewatchConfigCatalogName("catalogs//alertConfig"))
	assert.Equal(t, "", lakewatchConfigCatalogName("rules/abc"))
}

func TestUpdateTimeToMillis(t *testing.T) {
	assert.Equal(t, int64(0), updateTimeToMillis(nil))
	ts := time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, ts.UnixMilli(), updateTimeToMillis(sdktime.New(ts)))
}

func TestEmitLakewatchResources(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		api := mw.GetMockLakewatchAPI().EXPECT()
		api.GetAlertConfig(mock.Anything, lakewatch.GetAlertConfigRequest{Name: "catalogs/sec/alertConfig"}).
			Return(&lakewatch.AlertConfig{Name: "catalogs/sec/alertConfig", Version: 2}, nil)
		// Never updated - should be skipped
		api.GetCommonConfig(mock.Anything, lakewatch.GetCommonConfigRequest{Name: "catalogs/sec/commonConfig"}).
			Return(&lakewatch.CommonConfig{Name: "catalogs/sec/commonConfig"}, nil)
		api.GetDatasourceConfig(mock.Anything, lakewatch.GetDatasourceConfigRequest{Name: "catalogs/sec/datasourceConfig"}).
			Return(nil, fmt.Errorf("not configured: %w", apierr.ErrResourceDoesNotExist))
		api.GetRuleConfig(mock.Anything, lakewatch.GetRuleConfigRequest{Name: "catalogs/sec/ruleConfig"}).
			Return(&lakewatch.RuleConfig{Name: "catalogs/sec/ruleConfig",
				UpdateTime: sdktime.New(time.Now())}, nil)
		api.ListRules(mock.Anything, lakewatch.ListRulesRequest{CatalogName: "sec"}).
			Return(createIteratorFromSlice([]lakewatch.Rule{
				{Name: "rules/r1", CatalogName: "sec", SchemaName: "detections", RuleName: "rule1"},
			}))
		api.ListDatasources(mock.Anything, lakewatch.ListDatasourcesRequest{CatalogName: "sec"}).
			Return(createIteratorFromSlice([]lakewatch.Datasource{
				{Name: "datasources/d1", CatalogName: "sec", SchemaName: "raw", DatasourceName: "ds1"},
			}))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakewatch")

		ic.emitLakewatchResources("sec")

		assert.Equal(t, map[string]bool{
			"databricks_lakewatch_alert_config[<unknown>] (id: catalogs/sec/alertConfig)": true,
			"databricks_lakewatch_rule_config[<unknown>] (id: catalogs/sec/ruleConfig)":   true,
			"databricks_lakewatch_rule[<unknown>] (id: rules/r1)":                         true,
			"databricks_lakewatch_datasource[<unknown>] (id: datasources/d1)":             true,
		}, ic.testEmits)
	})
}

func TestEmitLakewatchResourcesNotAvailable(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		api := mw.GetMockLakewatchAPI().EXPECT()
		notEnabled := fmt.Errorf("lakewatch isn't enabled: %w", apierr.ErrNotImplemented)
		api.GetAlertConfig(mock.Anything, mock.Anything).Return(nil, notEnabled)
		api.GetCommonConfig(mock.Anything, mock.Anything).Return(nil, notEnabled)
		api.GetDatasourceConfig(mock.Anything, mock.Anything).Return(nil, notEnabled)
		api.GetRuleConfig(mock.Anything, mock.Anything).Return(nil, errors.New("unexpected error"))
		api.ListRules(mock.Anything, lakewatch.ListRulesRequest{CatalogName: "main"}).
			Return(createErrorIterator[lakewatch.Rule](notEnabled))
		api.ListDatasources(mock.Anything, lakewatch.ListDatasourcesRequest{CatalogName: "main"}).
			Return(createErrorIterator[lakewatch.Datasource](
				fmt.Errorf("no access: %w", apierr.ErrPermissionDenied)))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakewatch")

		// Errors must not fail the export, and nothing should be emitted
		ic.emitLakewatchResources("main")

		assert.Empty(t, ic.testEmits)
	})
}

func TestEmitLakewatchResourcesServiceNotEnabled(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		// No Lakewatch API calls should be made when the service is not enabled.
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs")

		ic.emitLakewatchResources("sec")

		assert.Empty(t, ic.testEmits)
	})
}

func TestLakewatchExport(t *testing.T) {
	updateTime := sdktime.New(time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC))
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		currentMetastoreSuccess,
		{
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/catalogs?max_results=0",
			ReuseRequest: true,
			Response: sdk_uc.ListCatalogsResponse{
				Catalogs: []sdk_uc.CatalogInfo{
					{Name: "sec", CatalogType: "MANAGED_CATALOG"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/catalogs/sec?",
			ReuseRequest: true,
			Response: sdk_uc.CatalogInfo{
				Name:        "sec",
				CatalogType: "MANAGED_CATALOG",
				MetastoreId: "12345678-1234",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/catalogs/sec/alertConfig?",
			ReuseRequest: true,
			Response: lakewatch.AlertConfig{
				Name:                                "catalogs/sec/alertConfig",
				DefaultSecurityAlertCollationWindow: "2h",
				Version:                             3,
				UpdateTime:                          updateTime,
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/catalogs/sec/commonConfig?",
			ReuseRequest: true,
			Response: lakewatch.CommonConfig{
				Name:                          "catalogs/sec/commonConfig",
				PerResourcePermissionsEnabled: true,
				Version:                       1,
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/catalogs/sec/datasourceConfig?",
			ReuseRequest: true,
			Response: lakewatch.DatasourceConfig{Name: "catalogs/sec/datasourceConfig",
				DefaultFolder: "raw", Version: 1},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/catalogs/sec/ruleConfig?",
			ReuseRequest: true,
			Status:       404,
			Response:     apierr.APIError{ErrorCode: "NOT_FOUND", StatusCode: 404, Message: "not found"},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/rules?catalog_name=sec",
			ReuseRequest: true,
			Response: lakewatch.ListRulesResponse{
				Rules: []lakewatch.Rule{
					{Name: "rules/r1", CatalogName: "sec", SchemaName: "detections", RuleName: "failed_logins"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/rules/r1?",
			ReuseRequest: true,
			Response: lakewatch.Rule{
				Name:        "rules/r1",
				RuleId:      "r1",
				CatalogName: "sec",
				SchemaName:  "detections",
				RuleName:    "failed_logins",
				Fidelity:    lakewatch.RuleFidelityFidelityHigh,
				Severity:    "High",
				Output:      lakewatch.RuleOutput{Summary: "Failed logins for {user}"},
				Schedule:    lakewatch.RuleSchedule{AtLeastEvery: "1h"},
				Batch: &lakewatch.RuleInputBatch{
					WarehouseId: "wh-123",
					Sql: &lakewatch.RuleInputBatchSql{Query: "SELECT * FROM sec.detections.logins l\n" +
						"JOIN `sec`.`raw`.`events` e ON l.id = e.id -- FROM sec.raw.commented\n" +
						"WHERE l.src = 'sec.raw.literal' AND l.x IN (SELECT x FROM other.sch.tbl)"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/datasources?catalog_name=sec",
			ReuseRequest: true,
			Response: lakewatch.ListDatasourcesResponse{
				Datasources: []lakewatch.Datasource{
					{Name: "datasources/d1", CatalogName: "sec", SchemaName: "raw", DatasourceName: "trivy"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/lakewatch/datasources/d1?",
			ReuseRequest: true,
			Response: lakewatch.Datasource{
				Name:           "datasources/d1",
				CatalogName:    "sec",
				SchemaName:     "raw",
				DatasourceName: "trivy",
				WarehouseId:    "wh-123",
				Schedule:       lakewatch.Schedule{AtLeastEvery: "1h"},
				Autoloader: &lakewatch.AutoLoader{
					Format:   lakewatch.AutoLoaderFormatAutoloaderFormatJson,
					Location: "/Volumes/sec/raw/logs/trivy/",
					RawTable: lakewatch.TableName{CatalogName: "sec", SchemaName: "raw", TableName: "trivy_raw"},
				},
				Parsers: []lakewatch.ParserRef{
					{Id: "p1", Version: "2", TargetTable: &lakewatch.TableName{
						CatalogName: "sec", SchemaName: "raw", TableName: "trivy_report"}},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/schemas/sec.raw?",
			ReuseRequest: true,
			Response:     sdk_uc.SchemaInfo{CatalogName: "sec", Name: "raw", FullName: "sec.raw"},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/schemas/sec.detections?",
			ReuseRequest: true,
			Response:     sdk_uc.SchemaInfo{CatalogName: "sec", Name: "detections", FullName: "sec.detections"},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/volumes/sec.raw.logs?",
			ReuseRequest: true,
			Response: sdk_uc.VolumeInfo{CatalogName: "sec", SchemaName: "raw", Name: "logs",
				FullName: "sec.raw.logs", VolumeType: "MANAGED"},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.Config.WorkspaceID = testProviderWorkspaceID
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("uc-catalogs")
		ic.enableServices("uc-catalogs,uc-schemas,uc-volumes,lakewatch")

		err := ic.Run()
		assert.NoError(t, err)

		content, err := os.ReadFile(tmpDir + "/lakewatch.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		assert.Contains(t, contentStr, `resource "databricks_lakewatch_alert_config" "sec_alert_config"`)
		assert.Contains(t, contentStr, `parent = "catalogs/${databricks_catalog.sec_test_managed_catalog.id}"`)
		assert.Contains(t, contentStr, `default_security_alert_collation_window = "2h"`)
		assert.Contains(t, contentStr, `resource "databricks_lakewatch_common_config" "sec_common_config"`)
		assert.Contains(t, contentStr, `per_resource_permissions_enabled = true`)
		assert.Contains(t, contentStr, `resource "databricks_lakewatch_datasource_config" "sec_datasource_config"`)
		assert.Contains(t, contentStr, `default_folder = databricks_schema.sec_raw.name`)
		assert.NotContains(t, contentStr, `databricks_lakewatch_rule_config`)

		assert.Contains(t, contentStr, `resource "databricks_lakewatch_datasource" "sec_raw_trivy"`)
		assert.Contains(t, contentStr, `schema_name = databricks_schema.sec_raw.name`)
		assert.Contains(t, contentStr, `location = "${databricks_volume.sec_raw_logs.volume_path}/trivy/"`)
		assert.Contains(t, contentStr, `table_name = "trivy_raw"`)
		assert.Contains(t, contentStr, `table_name = "trivy_report"`)
		assert.NotContains(t, contentStr, `catalog_name = "sec"`)
		assert.NotContains(t, contentStr, `schema_name = "raw"`)

		assert.Contains(t, contentStr, `resource "databricks_lakewatch_rule" "sec_detections_failed_logins"`)
		assert.Contains(t, contentStr, `catalog_name = databricks_catalog.sec_test_managed_catalog.id`)
		assert.Contains(t, contentStr, `rule_name = "failed_logins"`)
		assert.Contains(t, contentStr, `fidelity = "FIDELITY_HIGH"`)
		assert.Contains(t, contentStr, "query = \"SELECT * FROM ${databricks_schema.sec_detections.id}.logins l\\n"+
			"JOIN ${databricks_schema.sec_raw.id}.`events` e ON l.id = e.id -- FROM sec.raw.commented\\n"+
			"WHERE l.src = 'sec.raw.literal' AND l.x IN (SELECT x FROM other.sch.tbl)\"")
		assert.Contains(t, contentStr, `warehouse_id = "wh-123"`)
		assert.NotContains(t, contentStr, `rule_id`)
	})
}

// newLakewatchTestWrapper creates a Plugin Framework state with the given top-level attributes set.
func newLakewatchTestWrapper(t *testing.T, ic *importContext, resourceType, id string,
	attrs map[string]any) ResourceDataWrapper {
	wrapper, state := newPluginFrameworkTestState(t, ic, ic.Context, resourceType, id)
	for name, value := range attrs {
		require.Empty(t, state.SetAttribute(ic.Context, path.Root(name), value).Errors(), name)
	}
	return wrapper
}

func TestImportLakewatchRule(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	ic.enableServices("lakewatch,uc-schemas,sql-endpoints,notebooks")
	rule := lakewatch.Rule{
		Name:        "rules/r1",
		CatalogName: "sec",
		SchemaName:  "detections",
		RuleName:    "rule1",
		Batch: &lakewatch.RuleInputBatch{
			WarehouseId: "wh-123",
			Custom:      &lakewatch.RuleInputCustomNotebook{Notebook: "/detection"},
			Sql:         &lakewatch.RuleInputBatchSql{DefaultCatalogName: "logs", DefaultSchemaName: "aws"},
		},
	}
	var tfRule lakewatch_rule_resource.Rule
	require.False(t, converters.GoSdkToTfSdkStruct(ic.Context, rule, &tfRule).HasError())
	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_rule", rule.Name, map[string]any{
		"name":         tfRule.Name,
		"catalog_name": tfRule.CatalogName,
		"schema_name":  tfRule.SchemaName,
		"rule_name":    tfRule.RuleName,
		"batch":        tfRule.Batch,
	})
	assert.Equal(t, "sec_detections_rule1",
		ic.Importables["databricks_lakewatch_rule"].NameUnified(ic, wrapper))

	err := importLakewatchRule(ic, &resource{Resource: "databricks_lakewatch_rule", ID: rule.Name,
		DataWrapper: wrapper})
	require.NoError(t, err)
	for _, expected := range []string{
		"databricks_schema[<unknown>] (id: sec.detections)",
		"databricks_sql_endpoint[<unknown>] (id: wh-123)",
		"databricks_notebook[<unknown>] (id: /detection)",
		"databricks_schema[<unknown>] (id: logs.aws)",
	} {
		assert.Contains(t, ic.testEmits, expected)
	}
}

func TestImportLakewatchDatasource(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	ic.currentMetastore = currentMetastoreResponse
	ic.enableServices("lakewatch,uc-schemas,uc-volumes,sql-endpoints,secrets,uc-credentials,uc-connections")
	ds := lakewatch.Datasource{
		Name:           "datasources/d1",
		CatalogName:    "sec",
		SchemaName:     "raw",
		DatasourceName: "ds1",
		WarehouseId:    "wh-123",
		Autoloader: &lakewatch.AutoLoader{
			AwsSecretScope:     "aws-keys",
			ServiceCredentials: "sqs-cred",
			Location:           "/Volumes/sec/landing/logs/trivy/",
			RawTable:           lakewatch.TableName{TableName: "trivy_raw"},
		},
		Parsers: []lakewatch.ParserRef{
			{Id: "p1", TargetTable: &lakewatch.TableName{SchemaName: "parsed", TableName: "trivy"}},
		},
		ExistingTable:          &lakewatch.ExistingTable{Catalog: "src", Schema: "events", Name: "audit"},
		UnityCatalogConnection: &lakewatch.UnityCatalogConnection{ConnectionName: "slack"},
	}
	var tfDs lakewatch_datasource_resource.Datasource
	require.False(t, converters.GoSdkToTfSdkStruct(ic.Context, ds, &tfDs).HasError())
	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_datasource", ds.Name, map[string]any{
		"name":                     tfDs.Name,
		"catalog_name":             tfDs.CatalogName,
		"schema_name":              tfDs.SchemaName,
		"datasource_name":          tfDs.DatasourceName,
		"warehouse_id":             tfDs.WarehouseId,
		"autoloader":               tfDs.Autoloader,
		"parsers":                  tfDs.Parsers,
		"existing_table":           tfDs.ExistingTable,
		"unity_catalog_connection": tfDs.UnityCatalogConnection,
	})
	assert.Equal(t, "sec_raw_ds1",
		ic.Importables["databricks_lakewatch_datasource"].NameUnified(ic, wrapper))

	err := importLakewatchDatasource(ic, &resource{Resource: "databricks_lakewatch_datasource", ID: ds.Name,
		DataWrapper: wrapper})
	require.NoError(t, err)
	for _, expected := range []string{
		"databricks_schema[<unknown>] (id: sec.raw)",
		"databricks_sql_endpoint[<unknown>] (id: wh-123)",
		"databricks_secret_scope[<unknown>] (id: aws-keys)",
		"databricks_credential[<unknown>] (id: sqs-cred)",
		"databricks_connection[<unknown>] (id: 12345678-1234|slack)",
		"databricks_volume[<unknown>] (id: sec.landing.logs)",
		"databricks_schema[<unknown>] (id: sec.parsed)",
		"databricks_schema[<unknown>] (id: src.events)",
	} {
		assert.Contains(t, ic.testEmits, expected)
	}
}

func TestEmitVolumeByPath(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-volumes")
	ic.emitVolumeByPath("/Volumes/main/default/vol/some/dir/")
	ic.emitVolumeByPath("/Volumes/main/default/vol2")
	ic.emitVolumeByPath("/Volumes/main/default")
	ic.emitVolumeByPath("/Workspace/main/default/vol")
	ic.emitVolumeByPath("")
	assert.Equal(t, map[string]bool{
		"databricks_volume[<unknown>] (id: main.default.vol)":  true,
		"databricks_volume[<unknown>] (id: main.default.vol2)": true,
	}, ic.testEmits)
}

func TestImportLakewatchConfigs(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	ic.enableServices("lakewatch,uc-schemas,uc-volumes")

	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_rule_config", "catalogs/sec/ruleConfig",
		map[string]any{"default_folder": basetypes.NewStringValue("rules")})
	err := importLakewatchFolderConfig(ic, &resource{Resource: "databricks_lakewatch_rule_config",
		ID: "catalogs/sec/ruleConfig", DataWrapper: wrapper})
	require.NoError(t, err)
	assert.Equal(t, "sec_rule_config",
		ic.Importables["databricks_lakewatch_rule_config"].NameUnified(ic, wrapper))

	wrapper = newLakewatchTestWrapper(t, ic, "databricks_lakewatch_common_config", "catalogs/sec/commonConfig",
		map[string]any{
			"custom_content_pack_path": basetypes.NewStringValue("/Volumes/sec/content/packs/custom"),
			"custom_presets_path":      basetypes.NewStringValue("/Volumes/sec/content/presets"),
		})
	err = importLakewatchCommonConfig(ic, &resource{Resource: "databricks_lakewatch_common_config",
		ID: "catalogs/sec/commonConfig", DataWrapper: wrapper})
	require.NoError(t, err)

	assert.Equal(t, map[string]bool{
		"databricks_schema[<unknown>] (id: sec.rules)":           true,
		"databricks_volume[<unknown>] (id: sec.content.packs)":   true,
		"databricks_volume[<unknown>] (id: sec.content.presets)": true,
	}, ic.testEmits)
}

func TestIsMatchingLakewatchDefaultFolder(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_rule_config", "catalogs/sec/ruleConfig",
		map[string]any{
			"parent":         basetypes.NewStringValue("catalogs/sec"),
			"default_folder": basetypes.NewStringValue("rules"),
		})
	res := &resource{Resource: "databricks_lakewatch_rule_config", ID: "catalogs/sec/ruleConfig",
		DataWrapper: wrapper}
	schemaApproximation := func(catalogName, name string) *resourceApproximation {
		return &resourceApproximation{
			Type: "databricks_schema",
			Instances: []instanceApproximation{{Attributes: map[string]any{
				"catalog_name": catalogName,
				"name":         name,
			}}},
		}
	}
	assert.True(t, isMatchingLakewatchDefaultFolder(ic, res, schemaApproximation("sec", "rules"), "default_folder"))
	assert.False(t, isMatchingLakewatchDefaultFolder(ic, res, schemaApproximation("other", "rules"), "default_folder"))
	assert.False(t, isMatchingLakewatchDefaultFolder(ic, res, schemaApproximation("sec", "other"), "default_folder"))
}

func TestIsMatchingCatalogAndSchemaAndTablePluginFramework(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	var tfDs lakewatch_datasource_resource.Datasource
	require.False(t, converters.GoSdkToTfSdkStruct(ic.Context, lakewatch.Datasource{
		ExistingTable: &lakewatch.ExistingTable{Catalog: "src", Schema: "events", Name: "audit"},
	}, &tfDs).HasError())
	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_datasource", "datasources/d1",
		map[string]any{"existing_table": tfDs.ExistingTable})
	res := &resource{Resource: "databricks_lakewatch_datasource", ID: "datasources/d1", DataWrapper: wrapper}
	tableApproximation := func(catalogName, schemaName, name string) *resourceApproximation {
		return &resourceApproximation{
			Type: "databricks_sql_table",
			Instances: []instanceApproximation{{Attributes: map[string]any{
				"catalog_name": catalogName,
				"schema_name":  schemaName,
				"name":         name,
			}}},
		}
	}
	matcher := createIsMatchingCatalogAndSchemaAndTable("catalog", "schema", "name")
	assert.True(t, matcher(ic, res, tableApproximation("src", "events", "audit"), "existing_table.name"))
	assert.False(t, matcher(ic, res, tableApproximation("src", "other", "audit"), "existing_table.name"))
}
