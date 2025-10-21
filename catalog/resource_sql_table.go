package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MaxSqlExecWaitTimeout = 50
var optionPrefixes = []string{"option.", "spark.sql.dataSourceOptions."}

type SqlColumnInfo struct {
	Name     string         `json:"name"`
	Type     string         `json:"type_text,omitempty" tf:"alias:type,computed"`
	Identity IdentityColumn `json:"identity,omitempty"`
	Comment  string         `json:"comment,omitempty"`
	Nullable bool           `json:"nullable,omitempty" tf:"default:true"`
	TypeJson string         `json:"type_json,omitempty" tf:"computed"`
}

type TypeJson struct {
	Metadata map[string]any `json:"metadata,omitempty"`
}

type IdentityColumn string

const IdentityColumnNone IdentityColumn = ""
const IdentityColumnAlways IdentityColumn = "always"
const IdentityColumnDefault IdentityColumn = "default"

type SqlTableInfo struct {
	Name                  string            `json:"name"`
	CatalogName           string            `json:"catalog_name" tf:"force_new"`
	SchemaName            string            `json:"schema_name" tf:"force_new"`
	TableType             string            `json:"table_type" tf:"force_new"`
	DataSourceFormat      string            `json:"data_source_format,omitempty" tf:"force_new"`
	ColumnInfos           []SqlColumnInfo   `json:"columns,omitempty" tf:"alias:column,computed"`
	Partitions            []string          `json:"partitions,omitempty" tf:"force_new,computed"`
	ClusterKeys           []string          `json:"cluster_keys,omitempty"`
	StorageLocation       string            `json:"storage_location,omitempty" tf:"suppress_diff"`
	StorageCredentialName string            `json:"storage_credential_name,omitempty" tf:"force_new"`
	ViewDefinition        string            `json:"view_definition,omitempty"`
	Comment               string            `json:"comment,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
	Options               map[string]string `json:"options,omitempty" tf:"force_new"`
	// EffectiveProperties includes both properties and options. Options are prefixed with `option.`.
	EffectiveProperties map[string]string `json:"effective_properties" tf:"computed"`
	ClusterID           string            `json:"cluster_id,omitempty" tf:"computed"`
	WarehouseID         string            `json:"warehouse_id,omitempty"`
	Owner               string            `json:"owner,omitempty" tf:"computed"`
	TableID             string            `json:"table_id" tf:"computed"`

	exec    common.CommandExecutor
	sqlExec sql.StatementExecutionInterface
}

func (ti SqlTableInfo) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	caseInsensitiveFields := []string{"name", "catalog_name", "schema_name"}
	for _, field := range caseInsensitiveFields {
		s.SchemaPath(field).SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	}
	s.SchemaPath("data_source_format").SetCustomSuppressDiff(func(k, old, new string, d *schema.ResourceData) bool {
		if new == "" {
			return true
		}
		return strings.EqualFold(strings.ToLower(old), strings.ToLower(new))
	})
	s.SchemaPath("storage_location").SetCustomSuppressDiff(ucDirectoryPathSlashAndEmptySuppressDiff)
	s.SchemaPath("view_definition").SetCustomSuppressDiff(common.SuppressDiffWhitespaceChange)

	s.SchemaPath("cluster_id").SetConflictsWith([]string{"warehouse_id"})
	s.SchemaPath("warehouse_id").SetConflictsWith([]string{"cluster_id"})

	s.SchemaPath("partitions").SetConflictsWith([]string{"cluster_keys"})
	s.SchemaPath("cluster_keys").SetConflictsWith([]string{"partitions"})
	s.SchemaPath("column", "type").SetCustomSuppressDiff(func(k, old, new string, d *schema.ResourceData) bool {
		return getColumnType(old) == getColumnType(new)
	})
	return s
}

type SqlTablesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewSqlTablesAPI(ctx context.Context, m any) SqlTablesAPI {
	return SqlTablesAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

func (a SqlTablesAPI) getTable(name string) (ti SqlTableInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/tables/"+name, nil, &ti)
	// Copy returned properties & options to read-only attributes
	ti.EffectiveProperties = ti.Properties
	ti.Properties = nil
	return
}

func (ti *SqlTableInfo) FullName() string {
	return fmt.Sprintf("%s.%s.%s", ti.CatalogName, ti.SchemaName, ti.Name)
}

func (ti *SqlTableInfo) SQLFullName() string {
	return fmt.Sprintf("`%s`.`%s`.`%s`", ti.CatalogName, ti.SchemaName, ti.Name)
}

func parseComment(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, `\'`, `'`), `'`, `\'`)
}

func reconstructIdentity(c *SqlColumnInfo) (IdentityColumn, error) {
	if c.TypeJson == "" {
		return IdentityColumnNone, nil
	}
	var typeJson TypeJson
	err := json.Unmarshal([]byte(c.TypeJson), &typeJson)
	if err != nil {
		return IdentityColumnNone, err
	}
	if _, ok := typeJson.Metadata["delta.identity.start"]; !ok {
		return IdentityColumnNone, nil
	}
	explicit, ok := typeJson.Metadata["delta.identity.allowExplicitInsert"]
	if !ok {
		return IdentityColumnNone, nil
	}
	if explicit.(bool) {
		return IdentityColumnDefault, nil
	}
	return IdentityColumnAlways, nil
}

func (ti *SqlTableInfo) initCluster(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) (err error) {
	defaultClusterName := "terraform-sql-table"
	clustersAPI := clusters.NewClustersAPI(ctx, c)
	// if a warehouse id is specified, use the warehouse
	if wi, ok := d.GetOk("warehouse_id"); ok {
		ti.WarehouseID = wi.(string)
	} else if ci, ok := d.GetOk("cluster_id"); ok {
		// if a cluster id is specified, start the cluster
		ti.ClusterID = ci.(string)
		_, err = clustersAPI.StartAndGetInfo(ti.ClusterID)
		if apierr.IsMissing(err) {
			// cluster that was previously in a tfstate was deleted
			ti.ClusterID, err = ti.getOrCreateCluster(defaultClusterName, clustersAPI)
			if err != nil {
				return
			}
			_, err = clustersAPI.StartAndGetInfo(ti.ClusterID)
		}
		if err != nil {
			return
		}
	} else {
		// else, create a default cluster
		ti.ClusterID, err = ti.getOrCreateCluster(defaultClusterName, clustersAPI)
		if err != nil {
			return
		}
	}
	ti.exec = c.CommandExecutor(ctx)
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	ti.sqlExec = w.StatementExecution
	return nil
}

func (ti *SqlTableInfo) getOrCreateCluster(clusterName string, clustersAPI clusters.ClustersAPI) (string, error) {
	sparkVersion := clusters.LatestSparkVersionOrDefault(clustersAPI.Context(), clustersAPI.WorkspaceClient(), compute.SparkVersionRequest{
		Latest: true,
	})
	nodeType := clustersAPI.GetSmallestNodeType(clusters.NodeTypeRequest{NodeTypeRequest: compute.NodeTypeRequest{LocalDisk: true}})
	aclCluster, err := clustersAPI.GetOrCreateRunningCluster(
		clusterName, clusters.Cluster{
			ClusterName:            clusterName,
			SparkVersion:           sparkVersion,
			NodeTypeID:             nodeType,
			AutoterminationMinutes: 10,
			DataSecurityMode:       "SINGLE_USER",
			SparkConf: map[string]string{
				"spark.databricks.cluster.profile": "singleNode",
				"spark.master":                     "local[*]",
			},
			CustomTags: map[string]string{
				"ResourceClass": "SingleNode",
			},
		})
	if err != nil {
		return "", err
	}
	return aclCluster.ClusterID, nil
}

func (ci *SqlColumnInfo) getColumnType() string {
	var colType string
	switch ci.Identity {
	case IdentityColumnAlways:
		colType = fmt.Sprintf("%s GENERATED ALWAYS AS IDENTITY", ci.Type)
	case IdentityColumnDefault:
		colType = fmt.Sprintf("%s GENERATED BY DEFAULT AS IDENTITY", ci.Type)
	default:
		colType = ci.Type
	}
	return colType
}

func (ti *SqlTableInfo) serializeColumnInfo(col SqlColumnInfo) string {
	var colType = col.getColumnType()

	notNull := ""
	if !col.Nullable {
		notNull = " NOT NULL"
	}

	comment := ""
	if col.Comment != "" {
		comment = fmt.Sprintf(" COMMENT '%s'", parseComment(col.Comment))
	}
	return fmt.Sprintf("%s %s%s%s", col.getWrappedColumnName(), colType, notNull, comment) // id INT NOT NULL COMMENT 'something'
}

func (ti *SqlTableInfo) serializeColumnInfos() string {
	columnFragments := make([]string, len(ti.ColumnInfos))
	for i, col := range ti.ColumnInfos {
		columnFragments[i] = ti.serializeColumnInfo(col)
	}
	return strings.Join(columnFragments[:], ", ") // id INT NOT NULL, name STRING, age INT
}

func (ti *SqlTableInfo) serializeProperties() string {
	propsMap := make([]string, 0, len(ti.Properties))
	for key, value := range ti.Properties {
		propsMap = append(propsMap, fmt.Sprintf("'%s'='%s'", key, value))
	}
	return strings.Join(propsMap[:], ", ") // 'foo'='bar', 'this'='that'
}

func (ti *SqlTableInfo) serializeOptions() string {
	optionsMap := make([]string, 0, len(ti.Options))
	for key, value := range ti.Options {
		optionsMap = append(optionsMap, fmt.Sprintf("'%s'='%s'", key, value))
	}
	return strings.Join(optionsMap[:], ", ") // 'foo'='bar', 'this'='that'
}

func (ti *SqlTableInfo) buildLocationStatement() string {
	statements := make([]string, 0, 10)
	statements = append(statements, fmt.Sprintf("LOCATION '%s'", ti.StorageLocation)) // LOCATION '/mnt/csv_files'

	if ti.StorageCredentialName != "" {
		statements = append(statements, fmt.Sprintf(" WITH (CREDENTIAL `%s`)", ti.StorageCredentialName))
	}
	return strings.Join(statements, "")
}

func (ti *SqlTableInfo) getTableTypeString() string {
	if ti.TableType == "VIEW" {
		return "VIEW"
	}
	return "TABLE"
}

func (ti *SqlTableInfo) buildTableCreateStatement() string {
	statements := make([]string, 0, 10)

	isView := ti.TableType == "VIEW"

	createType := ti.getTableTypeString()

	// Use CREATE OR REPLACE for managed Delta tables
	if ti.DataSourceFormat == "DELTA" && ti.TableType == "MANAGED" {
		statements = append(statements, fmt.Sprintf("CREATE OR REPLACE %s %s", createType, ti.SQLFullName()))
	} else {
		statements = append(statements, fmt.Sprintf("CREATE %s %s", createType, ti.SQLFullName()))
	}

	if len(ti.ColumnInfos) > 0 {
		statements = append(statements, fmt.Sprintf(" (%s)", ti.serializeColumnInfos()))
	}

	if !isView {
		if ti.DataSourceFormat != "" {
			statements = append(statements, fmt.Sprintf("\nUSING %s", ti.DataSourceFormat)) // USING CSV
		}
	}

	if len(ti.Partitions) > 0 {
		statements = append(statements, fmt.Sprintf("\nPARTITIONED BY (%s)", strings.Join(ti.Partitions, ", "))) // PARTITIONED BY (university, major)
	}

	if len(ti.ClusterKeys) > 0 {
		statements = append(statements, fmt.Sprintf("\nCLUSTER BY %s", ti.getWrappedClusterKeys())) // CLUSTER BY (`university`, `major`)
	}

	if ti.Comment != "" {
		statements = append(statements, fmt.Sprintf("\nCOMMENT '%s'", parseComment(ti.Comment))) // COMMENT 'this is a comment'
	}

	if len(ti.Properties) > 0 {
		statements = append(statements, fmt.Sprintf("\nTBLPROPERTIES (%s)", ti.serializeProperties())) // TBLPROPERTIES ('foo'='bar')
	}

	if len(ti.Options) > 0 {
		statements = append(statements, fmt.Sprintf("\nOPTIONS (%s)", ti.serializeOptions())) // OPTIONS ('foo'='bar')
	}

	if !isView {
		if ti.StorageLocation != "" {
			statements = append(statements, "\n"+ti.buildLocationStatement())
		}
	} else {
		statements = append(statements, fmt.Sprintf("\nAS %s", ti.ViewDefinition))
	}

	statements = append(statements, ";")

	return strings.Join(statements, "")
}

// Wrapping the column name with backticks to avoid special character messing things up.
func (ci SqlColumnInfo) getWrappedColumnName() string {
	return fmt.Sprintf("`%s`", ci.Name)
}

// Wrapping column name with backticks to avoid special character messing things up.
func (ti *SqlTableInfo) getWrappedClusterKeys() string {
	if len(ti.ClusterKeys) == 1 {
		clusterKey := strings.ToUpper(ti.ClusterKeys[0])
		// If the cluster key is AUTO or NONE, we don't need to wrap it with backticks.
		if slices.Contains([]string{"AUTO", "NONE"}, clusterKey) {
			return clusterKey
		}
	}
	return "(`" + strings.Join(ti.ClusterKeys, "`,`") + "`)"
}

func (ti *SqlTableInfo) getStatementsForColumnDiffs(oldti *SqlTableInfo, statements []string, typestring string) []string {
	if len(ti.ColumnInfos) != len(oldti.ColumnInfos) {
		statements = ti.addOrRemoveColumnStatements(oldti, statements, typestring)
	} else {
		statements = ti.alterExistingColumnStatements(oldti, statements, typestring)
	}
	return statements
}

func (ti *SqlTableInfo) addOrRemoveColumnStatements(oldti *SqlTableInfo, statements []string, typestring string) []string {
	nameToOldColumn := make(map[string]SqlColumnInfo)
	nameToNewColumn := make(map[string]SqlColumnInfo)
	for _, ci := range oldti.ColumnInfos {
		nameToOldColumn[ci.Name] = ci
	}
	for _, newCi := range ti.ColumnInfos {
		nameToNewColumn[newCi.Name] = newCi
	}

	removeColumnStatements := make([]string, 0)

	for name, oldCi := range nameToOldColumn {
		if _, exists := nameToNewColumn[name]; !exists {
			// Remove old column if old column is no longer found in the config.
			removeColumnStatements = append(removeColumnStatements, oldCi.getWrappedColumnName())
		}
	}
	if len(removeColumnStatements) > 0 {
		removeColumnStatementsStr := strings.Join(removeColumnStatements, ", ")
		statements = append(statements, fmt.Sprintf("ALTER %s %s DROP COLUMN IF EXISTS (%s)", typestring, ti.SQLFullName(), removeColumnStatementsStr))
	}

	for i, newCi := range ti.ColumnInfos {
		if _, exists := nameToOldColumn[newCi.Name]; !exists {
			// Add new column if new column is detected.
			newCiStatement := ti.serializeColumnInfo(newCi)
			if i == 0 {
				// If this is the first column, add column with `FIRST` keyword
				statements = append(statements, fmt.Sprintf("ALTER %s %s ADD COLUMN %s FIRST", typestring, ti.SQLFullName(), newCiStatement))
			} else {
				// Find out the name of the column before this column and add after the previous one.
				statements = append(statements, fmt.Sprintf("ALTER %s %s ADD COLUMN %s AFTER %s", typestring, ti.SQLFullName(), newCiStatement, ti.ColumnInfos[i-1].Name))
			}
		}
	}

	return statements
}

func (ti *SqlTableInfo) alterExistingColumnStatements(oldti *SqlTableInfo, statements []string, typestring string) []string {
	for i, ci := range ti.ColumnInfos {
		oldCi := oldti.ColumnInfos[i]
		if ci.Name != oldCi.Name {
			statements = append(statements, fmt.Sprintf("ALTER %s %s RENAME COLUMN %s to %s", typestring, ti.SQLFullName(), oldCi.getWrappedColumnName(), ci.getWrappedColumnName()))
		}
		if ci.Comment != oldCi.Comment {
			statements = append(statements, fmt.Sprintf("ALTER %s %s ALTER COLUMN %s COMMENT '%s'", typestring, ti.SQLFullName(), ci.getWrappedColumnName(), parseComment(ci.Comment)))
		}
		if ci.Nullable != oldCi.Nullable {
			var keyWord string
			if ci.Nullable {
				keyWord = "DROP"
			} else {
				keyWord = "SET"
			}
			statements = append(statements, fmt.Sprintf("ALTER %s %s ALTER COLUMN %s %s NOT NULL", typestring, ti.SQLFullName(), ci.getWrappedColumnName(), keyWord))
		}
	}
	return statements
}

func (ti *SqlTableInfo) diff(oldti *SqlTableInfo) ([]string, error) {
	statements := make([]string, 0)
	typestring := ti.getTableTypeString()

	if ti.TableType == "VIEW" {
		// View only attributes
		ti.formatViewDefinition()
		if ti.ViewDefinition != oldti.ViewDefinition {
			statements = append(statements, fmt.Sprintf("ALTER VIEW %s AS %s", ti.SQLFullName(), ti.ViewDefinition))
		}
	} else {
		// Table only attributes
		if ti.StorageLocation != oldti.StorageLocation {
			statements = append(statements, fmt.Sprintf("ALTER TABLE %s SET %s", ti.SQLFullName(), ti.buildLocationStatement()))
		}
		equal := slices.Equal(ti.ClusterKeys, oldti.ClusterKeys)
		if !equal {
			statements = append(statements, fmt.Sprintf("ALTER TABLE %s CLUSTER BY %s", ti.SQLFullName(), ti.getWrappedClusterKeys()))
		}
	}

	// Attributes common to both views and tables
	if ti.Comment != oldti.Comment {
		statements = append(statements, fmt.Sprintf("COMMENT ON %s %s IS '%s'", typestring, ti.SQLFullName(), parseComment(ti.Comment)))
	}

	if !reflect.DeepEqual(ti.Properties, oldti.Properties) {
		// First handle removal of properties
		removeProps := make([]string, 0)
		for key := range oldti.Properties {
			if _, ok := ti.Properties[key]; !ok {
				removeProps = append(removeProps, key)
			}
		}
		if len(removeProps) > 0 {
			statements = append(statements, fmt.Sprintf("ALTER %s %s UNSET TBLPROPERTIES IF EXISTS (%s)", typestring, ti.SQLFullName(), strings.Join(removeProps, ",")))
		}
		// Next handle property changes and additions
		statements = append(statements, fmt.Sprintf("ALTER %s %s SET TBLPROPERTIES (%s)", typestring, ti.SQLFullName(), ti.serializeProperties()))
	}

	statements = ti.getStatementsForColumnDiffs(oldti, statements, typestring)

	return statements, nil
}

// formatViewDefinition removes empty lines and changes tabs to 4 spaces
// in order to compare view definitions correctly
func (ti *SqlTableInfo) formatViewDefinition() {

	// remove empty lines
	// 1\n\n\n2 => 1\n2
	ti.ViewDefinition = regexp.MustCompile(`[\r\n]+`).ReplaceAllString(ti.ViewDefinition, "\n")

	// change tab to 4 spaces
	// 1\t2 => 1    2
	ti.ViewDefinition = strings.ReplaceAll(ti.ViewDefinition, "\t", "    ")
}

func (ti *SqlTableInfo) updateTable(oldti *SqlTableInfo) error {
	statements, err := ti.diff(oldti)
	if err != nil {
		return err
	}
	for _, statement := range statements {
		err = ti.applySql(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ti *SqlTableInfo) createTable() error {
	return ti.applySql(ti.buildTableCreateStatement())
}

func (ti *SqlTableInfo) deleteTable() error {
	return ti.applySql(fmt.Sprintf("DROP %s %s", ti.getTableTypeString(), ti.SQLFullName()))
}

func (ti *SqlTableInfo) applySql(sqlQuery string) error {
	log.Printf("[INFO] Executing Sql: %s", sqlQuery)
	if ti.WarehouseID != "" {
		execCtx, cancel := context.WithTimeout(context.Background(), time.Duration(MaxSqlExecWaitTimeout)*time.Second)
		defer cancel()
		sqlRes, err := ti.sqlExec.ExecuteStatement(execCtx, sql.ExecuteStatementRequest{
			Statement:     sqlQuery,
			WaitTimeout:   fmt.Sprintf("%ds", MaxSqlExecWaitTimeout), //max allowed by sql exec
			WarehouseId:   ti.WarehouseID,
			OnWaitTimeout: sql.ExecuteStatementRequestOnWaitTimeoutCancel,
		})
		if err != nil {
			return err
		}
		if sqlRes.Status.State != "SUCCEEDED" {
			return fmt.Errorf("statement failed to execute: %s", sqlRes.Status.State)
		}
		return nil
	}

	r := ti.exec.Execute(ti.ClusterID, "sql", sqlQuery)
	if r.Failed() {
		return fmt.Errorf("cannot execute %s: %s", sqlQuery, r.Error())
	}
	return nil
}

func columnChangesCustomizeDiff(d *schema.ResourceDiff, newTable *SqlTableInfo) error {
	// Using plain type casting for oldCols because DiffToStructPointer does not support old value in the diff.
	old, _ := d.GetChange("column")
	oldCols := old.([]interface{})
	newColumnInfos := newTable.ColumnInfos

	if len(oldCols) == len(newColumnInfos) {
		err := assertNoColumnTypeDiff(oldCols, newColumnInfos)
		if err != nil {
			return err
		}
	} else {
		err := assertNoColumnMembershipAndFieldValueUpdate(oldCols, newColumnInfos)
		if err != nil {
			return err
		}
	}
	return nil
}

var columnTypeAliases = map[string]string{
	"integer": "int",
	"long":    "bigint",
	"real":    "float",
	"short":   "smallint",
	"byte":    "tinyint",
	"decimal": "decimal(10,0)",
	"dec":     "decimal(10,0)",
	"numeric": "decimal(10,0)",
}

func getColumnType(columnType string) string {
	// client side normalization is necessary to match normalization
	// that is happening on the backend side
	columnTypeNoBackticks := strings.ReplaceAll(columnType, "`", "")
	caseInsensitiveColumnType := strings.ToLower(columnTypeNoBackticks)
	if alias, ok := columnTypeAliases[caseInsensitiveColumnType]; ok {
		return alias
	}
	return caseInsensitiveColumnType
}

func assertNoColumnTypeDiff(oldCols []interface{}, newColumnInfos []SqlColumnInfo) error {
	for i, oldCol := range oldCols {
		oldColMap := oldCol.(map[string]interface{})
		if getColumnType(oldColMap["type"].(string)) != getColumnType(newColumnInfos[i].Type) {
			return fmt.Errorf("changing the 'type' of an existing column is not supported")
		}
		if oldColMap["identity"].(string) != string(newColumnInfos[i].Identity) {
			return fmt.Errorf("changing the 'identity' type of an existing column is not supported")
		}
	}
	return nil
}

// This function will throw if column addition or removal is happening together with column info field values.
func assertNoColumnMembershipAndFieldValueUpdate(oldCols []interface{}, newColumnInfos []SqlColumnInfo) error {
	oldColsNameToMap := make(map[string]map[string]interface{})
	newColsNameToMap := make(map[string]SqlColumnInfo)
	for _, oldCol := range oldCols {
		oldColMap := oldCol.(map[string]interface{})
		oldColsNameToMap[oldColMap["name"].(string)] = oldColMap
	}
	for _, newCol := range newColumnInfos {
		newColsNameToMap[newCol.Name] = newCol
	}
	for name, oldColMap := range oldColsNameToMap {
		if newCol, exists := newColsNameToMap[name]; exists {
			if getColumnType(oldColMap["type"].(string)) != getColumnType(newCol.Type) || oldColMap["nullable"] != newCol.Nullable || oldColMap["comment"] != newCol.Comment {
				return fmt.Errorf("detected changes in both number of columns and existing column field values, please do not change number of columns and update column values at the same time")
			}
		}
	}
	return nil
}

func ResourceSqlTable() common.Resource {
	tableSchema := common.StructToSchema(SqlTableInfo{}, nil)
	return common.Resource{
		Schema: tableSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if d.HasChange("column") {
				var newTableStruct SqlTableInfo
				common.DiffToStructPointer(d, tableSchema, &newTableStruct)
				err := columnChangesCustomizeDiff(d, &newTableStruct)
				if err != nil {
					return err
				}
			}
			// Compute the new effective property/options.
			// If the user changed a property or option, the resource will already be considered changed.
			// If the user specified a property but the value of that property has changed, that will appear
			// as a change in the effective property/option. To cause a diff to be detected, we need to
			// reset the effective property/option to the requested value.
			userSpecifiedProperties := d.Get("properties").(map[string]any)
			userSpecifiedOptions := d.Get("options").(map[string]any)
			effectiveProperties := d.Get("effective_properties").(map[string]any)
			diff := make(map[string]any)
			for k, userSpecifiedValue := range userSpecifiedProperties {
				if effectiveValue, ok := effectiveProperties[k]; !ok || effectiveValue != userSpecifiedValue {
					diff[k] = userSpecifiedValue
				}
			}
			for userOptName, userSpecifiedValue := range userSpecifiedOptions {
				var found bool
				var effectiveValue any
				var effectOptName string
				// If the option is not found, check if the user specified the option without the prefix
				// i.e. if user specified `multiLine` for JSON, then backend returns `spark.sql.dataSourceOptions.multiLine`
				for _, prefix := range optionPrefixes {
					effectOptName = prefix + userOptName
					if v, ok := effectiveProperties[effectOptName]; ok {
						found = true
						effectiveValue = v
						break
					}
				}
				if !found || effectiveValue != userSpecifiedValue {
					diff[effectOptName] = userSpecifiedValue
				}
			}
			if len(diff) > 0 {
				for k, v := range diff {
					effectiveProperties[k] = v
				}
				d.SetNew("effective_properties", effectiveProperties)
			}
			// No support yet for changing the COMMENT on a VIEW
			// Once added this can be removed
			if d.HasChange("comment") && d.Get("table_type") == "VIEW" {
				d.ForceNew("comment")
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti = new(SqlTableInfo)
			common.DataToStructPointer(d, tableSchema, ti)
			if err := ti.initCluster(ctx, d, c); err != nil {
				return err
			}
			if err := ti.createTable(); err != nil {
				return err
			}
			if ti.Owner != "" {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				err = w.Tables.Update(ctx, catalog.UpdateTableRequest{
					FullName: ti.FullName(),
					Owner:    ti.Owner,
				})
				if err != nil {
					return err
				}
			}
			d.SetId(ti.FullName())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ti, err := NewSqlTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			partitionInfo, err := w.Tables.GetByFullName(ctx, d.Id())
			if err != nil {
				return err
			}
			partitionIndexes := map[int]string{}
			for i := range ti.ColumnInfos {
				c := &ti.ColumnInfos[i]
				c.Identity, err = reconstructIdentity(c)
				if err != nil {
					return err
				}
			}

			for i := range partitionInfo.Columns {
				c := &partitionInfo.Columns[i]
				if slices.Contains(c.ForceSendFields, "PartitionIndex") {
					partitionIndexes[c.PartitionIndex] = c.Name
				}
			}
			indexes := slices.Sorted(maps.Keys(partitionIndexes))

			partitions := []string{}
			for _, p := range indexes {
				partitions = append(partitions, partitionIndexes[p])
			}

			d.Set("partitions", partitions)
			return common.StructToData(ti, tableSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var newti = new(SqlTableInfo)
			common.DataToStructPointer(d, tableSchema, newti)
			if err := newti.initCluster(ctx, d, c); err != nil {
				return err
			}
			oldti, err := NewSqlTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			err = newti.updateTable(&oldti)
			if err != nil {
				return err
			}
			if d.HasChange("owner") {
				// if new owner is not specified, set it to the current user
				if newti.Owner == "" {
					currentUser, err := w.CurrentUser.Me(ctx)
					if err != nil {
						return err
					}
					newti.Owner = currentUser.UserName
				}
				return w.Tables.Update(ctx, catalog.UpdateTableRequest{
					FullName: newti.FullName(),
					Owner:    newti.Owner,
				})
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti = new(SqlTableInfo)
			common.DataToStructPointer(d, tableSchema, ti)
			if err := ti.initCluster(ctx, d, c); err != nil {
				return err
			}
			return ti.deleteTable()
		},
	}
}
