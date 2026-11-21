package exporter

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdktime "github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	lakewatch_datasource_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/lakewatch_datasource"
	lakewatch_rule_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/lakewatch_rule"
)

// lakewatchConfigSuffixes maps the per-catalog Lakewatch singleton resources to the last
// segment of their resource name (`catalogs/{catalog}/{suffix}`).
var lakewatchConfigSuffixes = map[string]string{
	"databricks_lakewatch_alert_config":      "alertConfig",
	"databricks_lakewatch_common_config":     "commonConfig",
	"databricks_lakewatch_datasource_config": "datasourceConfig",
	"databricks_lakewatch_rule_config":       "ruleConfig",
}

// logLakewatchError logs an error returned by the Lakewatch API. Lakewatch is in Beta and
// may not be enabled in a workspace, or the catalog may not be used by Lakewatch, so such
// errors are expected and shouldn't fail the export.
func logLakewatchError(what, catalogName string, err error) {
	if apierr.IsMissing(err) || errors.Is(err, apierr.ErrNotImplemented) ||
		errors.Is(err, apierr.ErrPermissionDenied) || errors.Is(err, apierr.ErrBadRequest) {
		log.Printf("[DEBUG] Lakewatch %s isn't available for catalog %s: %v", what, catalogName, err)
		return
	}
	log.Printf("[WARN] Can't get Lakewatch %s for catalog %s: %v", what, catalogName, err)
}

// emitLakewatchResources emits Lakewatch rules, datasources, and per-catalog configurations
// bound to the given catalog. It's called from the Import function of `databricks_catalog`.
func (ic *importContext) emitLakewatchResources(catalogName string) {
	if !ic.isServiceEnabled("lakewatch") {
		return
	}
	ic.emitLakewatchConfigs(catalogName)
	ic.emitLakewatchRules(catalogName)
	ic.emitLakewatchDatasources(catalogName)
}

func (ic *importContext) emitLakewatchConfigs(catalogName string) {
	for resourceType, suffix := range lakewatchConfigSuffixes {
		name := fmt.Sprintf("catalogs/%s/%s", catalogName, suffix)
		configured, err := ic.isLakewatchConfigSet(resourceType, name)
		if err != nil {
			logLakewatchError(suffix, catalogName, err)
			continue
		}
		if !configured {
			log.Printf("[DEBUG] Lakewatch %s isn't configured for catalog %s", suffix, catalogName)
			continue
		}
		ic.Emit(&resource{
			Resource: resourceType,
			ID:       name,
		})
	}
}

// isLakewatchConfigSet checks if a per-catalog Lakewatch configuration was ever updated.
// Singletons are readable even when not configured, so we use version/update time to
// avoid generating resources with only default values.
func (ic *importContext) isLakewatchConfigSet(resourceType, name string) (bool, error) {
	w := ic.workspaceClient.Lakewatch
	switch resourceType {
	case "databricks_lakewatch_alert_config":
		c, err := w.GetAlertConfig(ic.Context, lakewatch.GetAlertConfigRequest{Name: name})
		if err != nil {
			return false, err
		}
		return c.Version > 0 || c.UpdateTime != nil, nil
	case "databricks_lakewatch_common_config":
		c, err := w.GetCommonConfig(ic.Context, lakewatch.GetCommonConfigRequest{Name: name})
		if err != nil {
			return false, err
		}
		return c.Version > 0 || c.UpdateTime != nil, nil
	case "databricks_lakewatch_datasource_config":
		c, err := w.GetDatasourceConfig(ic.Context, lakewatch.GetDatasourceConfigRequest{Name: name})
		if err != nil {
			return false, err
		}
		return c.Version > 0 || c.UpdateTime != nil, nil
	case "databricks_lakewatch_rule_config":
		c, err := w.GetRuleConfig(ic.Context, lakewatch.GetRuleConfigRequest{Name: name})
		if err != nil {
			return false, err
		}
		return c.Version > 0 || c.UpdateTime != nil, nil
	}
	return false, fmt.Errorf("unsupported Lakewatch configuration resource: %s", resourceType)
}

func (ic *importContext) emitLakewatchRules(catalogName string) {
	it := ic.workspaceClient.Lakewatch.ListRules(ic.Context, lakewatch.ListRulesRequest{CatalogName: catalogName})
	i := 0
	for it.HasNext(ic.Context) {
		rule, err := it.Next(ic.Context)
		if err != nil {
			logLakewatchError("rules", catalogName, err)
			break
		}
		fullName := fmt.Sprintf("%s.%s.%s", rule.CatalogName, rule.SchemaName, rule.RuleName)
		if !ic.MatchesName(fullName) {
			log.Printf("[INFO] Skipping Lakewatch rule %s because it doesn't match %s", fullName, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_lakewatch_rule",
			ID:       rule.Name,
		}, updateTimeToMillis(rule.UpdateTime), fmt.Sprintf("Lakewatch rule '%s'", fullName))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Lakewatch rules in catalog %s", i, catalogName)
	}
}

func (ic *importContext) emitLakewatchDatasources(catalogName string) {
	it := ic.workspaceClient.Lakewatch.ListDatasources(ic.Context,
		lakewatch.ListDatasourcesRequest{CatalogName: catalogName})
	i := 0
	for it.HasNext(ic.Context) {
		ds, err := it.Next(ic.Context)
		if err != nil {
			logLakewatchError("datasources", catalogName, err)
			break
		}
		fullName := fmt.Sprintf("%s.%s.%s", ds.CatalogName, ds.SchemaName, ds.DatasourceName)
		if !ic.MatchesName(fullName) {
			log.Printf("[INFO] Skipping Lakewatch datasource %s because it doesn't match %s", fullName, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_lakewatch_datasource",
			ID:       ds.Name,
		}, updateTimeToMillis(ds.UpdateTime), fmt.Sprintf("Lakewatch datasource '%s'", fullName))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Lakewatch datasources in catalog %s", i, catalogName)
	}
}

// updateTimeToMillis converts an optional update time into milliseconds, returning 0 when unset.
func updateTimeToMillis(t *sdktime.Time) int64 {
	if t == nil {
		return 0
	}
	return t.AsTime().UnixMilli()
}

// lakewatchConfigCatalogName extracts the catalog name from a Lakewatch configuration
// resource name (`catalogs/{catalog}/{suffix}`).
func lakewatchConfigCatalogName(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) == 3 && parts[0] == "catalogs" && parts[1] != "" {
		return parts[1]
	}
	return ""
}

func lakewatchConfigNameUnified(suffix string) func(ic *importContext, wrapper ResourceDataWrapper) string {
	return func(ic *importContext, wrapper ResourceDataWrapper) string {
		if catalogName := lakewatchConfigCatalogName(wrapper.Id()); catalogName != "" {
			return catalogName + "_" + suffix
		}
		return generateUniqueID(wrapper.Id())
	}
}

func lakewatchObjectNameUnified(nameAttr string) func(ic *importContext, wrapper ResourceDataWrapper) string {
	return func(ic *importContext, wrapper ResourceDataWrapper) string {
		catalogName, _ := wrapper.Get("catalog_name").(string)
		schemaName, _ := wrapper.Get("schema_name").(string)
		objName, _ := wrapper.Get(nameAttr).(string)
		if catalogName != "" && schemaName != "" && objName != "" {
			return catalogName + "_" + schemaName + "_" + objName
		}
		return generateUniqueID(wrapper.Id())
	}
}

// emitLakewatchSchema emits a schema if both catalog and schema names are known.
func (ic *importContext) emitLakewatchSchema(catalogName, schemaName string) {
	if catalogName != "" && schemaName != "" {
		ic.Emit(&resource{
			Resource: "databricks_schema",
			ID:       catalogName + "." + schemaName,
		})
	}
}

// emitLakewatchTableSchema emits the schema of a table that Lakewatch writes to. Catalog and
// schema are optional, and default to the ones of the datasource.
func (ic *importContext) emitLakewatchTableSchema(table *lakewatch.TableName, ds *lakewatch.Datasource) {
	if table == nil {
		return
	}
	catalogName, schemaName := table.CatalogName, table.SchemaName
	if catalogName == "" {
		catalogName = ds.CatalogName
	}
	if schemaName == "" {
		schemaName = ds.SchemaName
	}
	ic.emitLakewatchSchema(catalogName, schemaName)
}

// emitLakewatchCommonDependencies emits UC objects and the run-as identity used by rules
// and datasources.
func (ic *importContext) emitLakewatchCommonDependencies(catalogName, schemaName, runAs, warehouseId string) {
	ic.emitLakewatchSchema(catalogName, schemaName)
	if warehouseId != "" {
		ic.Emit(&resource{
			Resource: "databricks_sql_endpoint",
			ID:       warehouseId,
		})
	}
	ic.emitUserOrServicePrincipal(runAs)
}

func importLakewatchRule(ic *importContext, r *resource) error {
	var rule lakewatch.Rule
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper, lakewatch_rule_resource.Rule{}, &rule); err != nil {
		return err
	}
	warehouseId := ""
	if rule.Batch != nil {
		warehouseId = rule.Batch.WarehouseId
		if rule.Batch.Sql != nil {
			ic.emitLakewatchSchema(rule.Batch.Sql.DefaultCatalogName, rule.Batch.Sql.DefaultSchemaName)
			// Only schemas from the rule's catalog are emitted to avoid exporting unrelated
			// objects. References to other schemas are resolved if they're exported anyway.
			for _, name := range findSqlSchemaNames(rule.Batch.Sql.Query) {
				if strings.EqualFold(name.Catalog, rule.CatalogName) {
					ic.emitLakewatchSchema(name.Catalog, name.Schema)
				}
			}
		}
		if rule.Batch.Custom != nil && strings.HasPrefix(rule.Batch.Custom.Notebook, "/") {
			ic.emitNotebookOrRepo(rule.Batch.Custom.Notebook)
		}
	}
	ic.emitLakewatchCommonDependencies(rule.CatalogName, rule.SchemaName, rule.RunAsPrincipal, warehouseId)
	return nil
}

func importLakewatchDatasource(ic *importContext, r *resource) error {
	var ds lakewatch.Datasource
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		lakewatch_datasource_resource.Datasource{}, &ds); err != nil {
		return err
	}
	ic.emitLakewatchCommonDependencies(ds.CatalogName, ds.SchemaName, ds.RunAsPrincipal, ds.WarehouseId)
	for _, parser := range ds.Parsers {
		ic.emitLakewatchTableSchema(parser.TargetTable, &ds)
	}
	if ds.ExistingTable != nil {
		ic.emitLakewatchSchema(ds.ExistingTable.Catalog, ds.ExistingTable.Schema)
	}
	if ds.Autoloader != nil {
		ic.emitLakewatchTableSchema(&ds.Autoloader.RawTable, &ds)
		ic.emitVolumeByPath(ds.Autoloader.Location)
		if ds.Autoloader.AwsSecretScope != "" {
			ic.Emit(&resource{
				Resource: "databricks_secret_scope",
				ID:       ds.Autoloader.AwsSecretScope,
			})
		}
		if ds.Autoloader.ServiceCredentials != "" {
			ic.Emit(&resource{
				Resource: "databricks_credential",
				ID:       ds.Autoloader.ServiceCredentials,
			})
		}
	}
	if ds.UnityCatalogConnection != nil {
		ic.emitConnectionByName(ds.UnityCatalogConnection.ConnectionName)
	}
	return nil
}

func importLakewatchCommonConfig(ic *importContext, r *resource) error {
	for _, attr := range []string{"custom_content_pack_path", "custom_presets_path"} {
		if v, ok := r.DataWrapper.Get(attr).(string); ok {
			ic.emitVolumeByPath(v)
		}
	}
	return nil
}

// importLakewatchFolderConfig emits the schema registered as the default folder for
// datasource & rule configurations.
func importLakewatchFolderConfig(ic *importContext, r *resource) error {
	if folder, ok := r.DataWrapper.Get("default_folder").(string); ok {
		ic.emitLakewatchSchema(lakewatchConfigCatalogName(r.ID), folder)
	}
	return nil
}

// isMatchingLakewatchDefaultFolder checks that the `default_folder` schema belongs to the
// catalog of the Lakewatch configuration (`parent` in the format `catalogs/{catalog}`).
func isMatchingLakewatchDefaultFolder(ic *importContext, res *resource, ra *resourceApproximation,
	origPath string) bool {
	catalogName := strings.TrimPrefix(resourceGetString(res, "parent"), "catalogs/")
	schemaName := resourceGetString(res, origPath)
	if catalogName == "" || schemaName == "" {
		return false
	}
	raCatalogName, catFound := ra.Get("catalog_name")
	raSchemaName, schemaFound := ra.Get("name")
	return catFound && schemaFound && raCatalogName.(string) == catalogName &&
		raSchemaName.(string) == schemaName
}
