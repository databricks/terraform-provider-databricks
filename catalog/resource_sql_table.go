package catalog

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MaxSqlExecWaitTimeout = 50

type SqlColumnInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type_text,omitempty" tf:"suppress_diff,alias:type"`
	Comment  string `json:"comment,omitempty"`
	Nullable bool   `json:"nullable,omitempty" tf:"default:true"`
}

type SqlTableInfo struct {
	Name                  string            `json:"name"`
	CatalogName           string            `json:"catalog_name" tf:"force_new"`
	SchemaName            string            `json:"schema_name" tf:"force_new"`
	TableType             string            `json:"table_type" tf:"force_new"`
	DataSourceFormat      string            `json:"data_source_format,omitempty" tf:"force_new"`
	ColumnInfos           []SqlColumnInfo   `json:"columns,omitempty" tf:"alias:column,computed,force_new"`
	Partitions            []string          `json:"partitions,omitempty" tf:"force_new"`
	ClusterKeys           []string          `json:"cluster_keys,omitempty" tf:"force_new"`
	StorageLocation       string            `json:"storage_location,omitempty" tf:"suppress_diff"`
	StorageCredentialName string            `json:"storage_credential_name,omitempty" tf:"force_new"`
	ViewDefinition        string            `json:"view_definition,omitempty"`
	Comment               string            `json:"comment,omitempty"`
	Properties            map[string]string `json:"properties,omitempty" tf:"computed"`
	Options               map[string]string `json:"options,omitempty" tf:"force_new"`
	ClusterID             string            `json:"cluster_id,omitempty" tf:"computed"`
	WarehouseID           string            `json:"warehouse_id,omitempty"`

	exec    common.CommandExecutor
	sqlExec sql.StatementExecutionInterface
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

// These properties are added automatically
// If we do not customize the diff using these then terraform will constantly try to remove them
// `properties` is essentially a "partially" computed field
// This needs to be replaced with something a bit more robust in the future
func sqlTableIsManagedProperty(key string) bool {
	managedProps := map[string]bool{
		"delta.lastCommitTimestamp":                                true,
		"delta.lastUpdateVersion":                                  true,
		"delta.minReaderVersion":                                   true,
		"delta.minWriterVersion":                                   true,
		"delta.enableDeletionVectors":                              true,
		"delta.enableRowTracking":                                  true,
		"delta.feature.deletionVectors":                            true,
		"delta.feature.domainMetadata":                             true,
		"delta.feature.liquid":                                     true,
		"delta.feature.rowTracking":                                true,
		"delta.feature.v2Checkpoint":                               true,
		"delta.liquid.clusteringColumns":                           true,
		"delta.rowTracking.materializedRowCommitVersionColumnName": true,
		"delta.rowTracking.materializedRowIdColumnName":            true,
		"delta.checkpoint.writeStatsAsJson":                        true,
		"delta.checkpoint.writeStatsAsStruct":                      true,
		"delta.checkpointPolicy":                                   true,
		"view.catalogAndNamespace.numParts":                        true,
		"view.catalogAndNamespace.part.0":                          true,
		"view.catalogAndNamespace.part.1":                          true,
		"view.query.out.col.0":                                     true,
		"view.query.out.numCols":                                   true,
		"view.referredTempFunctionsNames":                          true,
		"view.referredTempViewNames":                               true,
		"view.sqlConfig.spark.sql.hive.convertCTAS":                true,
		"view.sqlConfig.spark.sql.legacy.createHiveTableByDefault": true,
		"view.sqlConfig.spark.sql.parquet.compression.codec":       true,
		"view.sqlConfig.spark.sql.session.timeZone":                true,
		"view.sqlConfig.spark.sql.sources.commitProtocolClass":     true,
		"view.sqlConfig.spark.sql.sources.default":                 true,
		"view.sqlConfig.spark.sql.streaming.stopTimeout":           true,
	}
	return managedProps[key]
}

func (ti *SqlTableInfo) initCluster(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) (err error) {
	defaultClusterName := "terraform-sql-table"
	clustersAPI := clusters.NewClustersAPI(ctx, c)
	// if a cluster id is specified, start the cluster
	if ci, ok := d.GetOk("cluster_id"); ok {
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
		// if a warehouse id is specified, use the warehouse
	} else if wi, ok := d.GetOk("warehouse_id"); ok {
		ti.WarehouseID = wi.(string)
		// else, create a default cluster
	} else {
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
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{
		Latest: true,
	})
	nodeType := clustersAPI.GetSmallestNodeType(compute.NodeTypeRequest{LocalDisk: true})
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

func (ti *SqlTableInfo) serializeColumnInfo(col SqlColumnInfo) string {
	notNull := ""
	if !col.Nullable {
		notNull = " NOT NULL"
	}

	comment := ""
	if col.Comment != "" {
		comment = fmt.Sprintf(" COMMENT '%s'", parseComment(col.Comment))
	}
	return fmt.Sprintf("%s %s%s%s", col.Name, col.Type, notNull, comment) // id INT NOT NULL COMMENT 'something'
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
		if !sqlTableIsManagedProperty(key) {
			propsMap = append(propsMap, fmt.Sprintf("'%s'='%s'", key, value))
		}
	}
	return strings.Join(propsMap[:], ", ") // 'foo'='bar', 'this'='that'
}

func (ti *SqlTableInfo) serializeOptions() string {
	optionsMap := make([]string, 0, len(ti.Options))
	for key, value := range ti.Options {
		if !sqlTableIsManagedProperty(key) {
			optionsMap = append(optionsMap, fmt.Sprintf("'%s'='%s'", key, value))
		}
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

	externalFragment := ""
	if ti.TableType == "EXTERNAL" {
		externalFragment = "EXTERNAL "
	}

	createType := ti.getTableTypeString()

	statements = append(statements, fmt.Sprintf("CREATE %s%s %s", externalFragment, createType, ti.SQLFullName()))

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
		statements = append(statements, fmt.Sprintf("\nCLUSTER BY (%s)", strings.Join(ti.ClusterKeys, ", "))) // CLUSTER BY (university, major)
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

// Wrapping the column name with backtiks to avoid special character messing things up.
func getWrappedColumnName(ci SqlColumnInfo) string {
	return fmt.Sprintf("`%s`", ci.Name)
}

func (ti *SqlTableInfo) getStatementsForColumnDiffs(oldti *SqlTableInfo, statements []string, typestring string) {
	// TODO: take out "force_new" in `column` and add case to handle addition and removal of columns.
	for i, ci := range ti.ColumnInfos {
		oldCi := oldti.ColumnInfos[i]
		if ci.Name != oldCi.Name {
			statements = append(statements, fmt.Sprintf("ALTER %s %s RENAME COLUMN %s to %s", typestring, ti.SQLFullName(), getWrappedColumnName(oldCi), getWrappedColumnName(ci)))
		}
		if ci.Comment != oldCi.Comment {
			statements = append(statements, fmt.Sprintf("ALTER %s %s ALTER COLUMN %s COMMENT %s", typestring, ti.SQLFullName(), getWrappedColumnName(ci), parseComment(ci.Comment)))
		}
		if ci.Nullable != oldCi.Nullable {
			var keyWord string
			if ci.Nullable {
				keyWord = "SET"
			} else {
				keyWord = "DROP"
			}
			statements = append(statements, fmt.Sprintf("ALTER %s %s ALTER COLUMN %s %s NULLABLE", typestring, ti.SQLFullName(), getWrappedColumnName(ci), keyWord))
		}
	}
}

func (ti *SqlTableInfo) diff(oldti *SqlTableInfo) ([]string, error) {
	statements := make([]string, 0)
	typestring := ti.getTableTypeString()

	if ti.TableType == "VIEW" {
		// View only attributes
		if ti.ViewDefinition != oldti.ViewDefinition {
			statements = append(statements, fmt.Sprintf("ALTER VIEW %s AS %s", ti.SQLFullName(), ti.ViewDefinition))
		}
	} else {
		// Table only attributes
		if ti.StorageLocation != oldti.StorageLocation {
			statements = append(statements, fmt.Sprintf("ALTER TABLE %s SET %s", ti.SQLFullName(), ti.buildLocationStatement()))
		}
		if !reflect.DeepEqual(ti.ClusterKeys, oldti.ClusterKeys) {
			statements = append(statements, fmt.Sprintf("ALTER TABLE %s CLUSTER BY (%s)", ti.SQLFullName(), strings.Join(ti.ClusterKeys, ", ")))
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

	ti.getStatementsForColumnDiffs(oldti, statements, typestring)

	return statements, nil
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

func columnChangesCustomizeDiff(d *schema.ResourceDiff) error {
	if d.HasChange("column") {
		old, new := d.GetChange("column")
		oldCols := old.([]interface{})
		newCols := new.([]interface{})

		// Only handling same number of columns for now, will address different number of columns as a follow-up.
		if len(oldCols) == len(newCols) {
			for i, oldCol := range oldCols {
				oldColMap := oldCol.(map[string]interface{})
				newColMap := newCols[i].(map[string]interface{})

				if oldColMap["type"] != newColMap["type"] {
					return fmt.Errorf("changing the 'type' of an existing column is not supported")
				}
			}
		}
	}
	return nil
}

func ResourceSqlTable() common.Resource {
	tableSchema := common.StructToSchema(SqlTableInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["data_source_format"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if new == "" {
					return true
				}
				return strings.EqualFold(strings.ToLower(old), strings.ToLower(new))
			}
			s["storage_location"].DiffSuppressFunc = ucDirectoryPathSlashAndEmptySuppressDiff
			s["view_definition"].DiffSuppressFunc = common.SuppressDiffWhitespaceChange

			s["cluster_id"].ConflictsWith = []string{"warehouse_id"}
			s["warehouse_id"].ConflictsWith = []string{"cluster_id"}

			s["partitions"].ConflictsWith = []string{"cluster_keys"}
			s["cluster_keys"].ConflictsWith = []string{"partitions"}
			return s
		})
	return common.Resource{
		Schema: tableSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			err := columnChangesCustomizeDiff(d)
			if err != nil {
				return err
			}
			if d.HasChange("properties") {
				old, new := d.GetChange("properties")
				oldProps := old.(map[string]any)
				newProps := new.(map[string]any)
				old, _ = d.GetChange("options")
				options := old.(map[string]any)
				for key := range oldProps {
					if _, ok := newProps[key]; !ok {
						//options also gets exposed as properties
						if _, ok := options[key]; ok {
							newProps[key] = oldProps[key]
						}
						//some options are exposed as option.[...] properties
						if sqlTableIsManagedProperty(key) || strings.HasPrefix(key, "option.") {
							newProps[key] = oldProps[key]
						}
					}
				}
				d.SetNew("properties", newProps)
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
			d.SetId(ti.FullName())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ti, err := NewSqlTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ti, tableSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var newti = new(SqlTableInfo)
			common.DataToStructPointer(d, tableSchema, newti)
			if err := newti.initCluster(ctx, d, c); err != nil {
				return err
			}
			oldti, err := NewSqlTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			return newti.updateTable(&oldti)
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
