package catalog

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	clustersApi "github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SqlColumnInfo struct {
	Name             string `json:"name"`
	TypeText         string `json:"type_text"`
	TypeJson         string `json:"type_json,omitempty"`
	TypeName         string `json:"type_name"`
	TypePrecision    int32  `json:"type_precision,omitempty"`
	TypeScale        int32  `json:"type_scale,omitempty"`
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	Position         int32  `json:"position"`
	Comment          string `json:"comment,omitempty"`
	Nullable         bool   `json:"nullable,omitempty" tf:"default:true"`
	PartitionIndex   int32  `json:"partition_index,omitempty"`
}

type SqlTableInfo struct {
	Name                  string            `json:"name"`
	CatalogName           string            `json:"catalog_name" tf:"force_new"`
	SchemaName            string            `json:"schema_name" tf:"force_new"`
	TableType             string            `json:"table_type" tf:"force_new"`
	DataSourceFormat      string            `json:"data_source_format"`
	ColumnInfos           []SqlColumnInfo   `json:"columns,omitempty" tf:"alias:column,computed"`
	StorageLocation       string            `json:"storage_location,omitempty" tf:"suppress_diff"`
	StorageCredentialName string            `json:"storage_credential_name,omitempty" tf:"force_new"`
	ViewDefinition        string            `json:"view_definition,omitempty"`
	Owner                 string            `json:"owner,omitempty" tf:"computed"`
	Comment               string            `json:"comment,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
	ClusterID             string            `json:"cluster_id,omitempty" tf:"computed"`

	exec common.CommandExecutor
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

func (a SqlTablesAPI) deleteTable(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/tables/"+name, nil)
}

func (ti *SqlTableInfo) FullName() string {
	return fmt.Sprintf("%s.%s.%s", ti.CatalogName, ti.SchemaName, ti.Name)
}

type SqlTables struct {
	Tables []SqlTableInfo `json:"tables"`
}

func (ti *SqlTableInfo) initCluster(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) (err error) {
	clustersAPI := clusters.NewClustersAPI(ctx, c)
	if ci, ok := d.GetOk("cluster_id"); ok {
		ti.ClusterID = ci.(string)
	} else {
		ti.ClusterID, err = ti.getOrCreateCluster(clustersAPI)
		if err != nil {
			return
		}
	}
	_, err = clustersAPI.StartAndGetInfo(ti.ClusterID)
	if apierr.IsMissing(err) {
		// cluster that was previously in a tfstate was deleted
		ti.ClusterID, err = ti.getOrCreateCluster(clustersAPI)
		if err != nil {
			return
		}
		_, err = clustersAPI.StartAndGetInfo(ti.ClusterID)
	}
	if err != nil {
		return
	}
	ti.exec = c.CommandExecutor(ctx)
	return nil
}

func (ti *SqlTableInfo) getOrCreateCluster(clustersAPI clusters.ClustersAPI) (string, error) {
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{
		Latest: true,
	})
	nodeType := clustersAPI.GetSmallestNodeType(clustersApi.NodeTypeRequest{LocalDisk: true})
	aclCluster, err := clustersAPI.GetOrCreateRunningCluster(
		"terraform-table", clusters.Cluster{
			ClusterName:            "terraform-sql-table",
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
		comment = fmt.Sprintf(" COMMENT %s", col.Comment)
	}
	return fmt.Sprintf("%s %s%s%s", col.Name, col.TypeText, notNull, comment) // id INT NOT NULL COMMENT something
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

func (ti *SqlTableInfo) buildTableCreateStatement() string {
	statements := make([]string, 0, 10)

	isView := ti.TableType == "VIEW"

	externalFragment := ""
	if ti.TableType == "EXTERNAL" {
		externalFragment = "EXTERNAL "
	}

	createType := "TABLE"
	if isView {
		createType = "VIEW"
	}

	statements = append(statements, fmt.Sprintf("CREATE %s%s %s", externalFragment, createType, ti.FullName()))

	if len(ti.ColumnInfos) > 0 {
		statements = append(statements, fmt.Sprintf(" (%s)", ti.serializeColumnInfos()))
	}

	if ti.Comment != "" {
		statements = append(statements, fmt.Sprintf("\nCOMMENT '%s'", ti.Comment)) // COMMENT 'this is a comment'
	}

	if len(ti.Properties) > 0 {
		statements = append(statements, fmt.Sprintf("\nTBLPROPERTIES (%s)", ti.serializeProperties())) // TBLPROPERTIES ('foo'='bar')
	}

	if !isView {
		if ti.DataSourceFormat != "" {
			statements = append(statements, fmt.Sprintf("\nUSING %s", ti.DataSourceFormat)) // USING CSV
		}

		if ti.StorageLocation != "" {
			statements = append(statements, fmt.Sprintf("\nLOCATION '%s'", ti.StorageLocation)) // LOCATION '/mnt/csv_files'

			if ti.StorageCredentialName != "" {
				statements = append(statements, fmt.Sprintf(" WITH (CREDENTIAL `%s`)", ti.StorageCredentialName))
			}
		}
	} else {
		statements = append(statements, fmt.Sprintf("\nAS %s", ti.ViewDefinition))
	}

	statements = append(statements, ";")

	return strings.Join(statements, "")
}

func (ti *SqlTableInfo) createTable(ctx context.Context, d *schema.ResourceData, s map[string]*schema.Schema, c *common.DatabricksClient) error {

	sqlQuery := ti.buildTableCreateStatement()

	log.Printf("[INFO] Executing Sql: %s", sqlQuery)
	r := ti.exec.Execute(ti.ClusterID, "sql", sqlQuery)

	if !r.Failed() {
		return nil
	}
	return fmt.Errorf("cannot execute %s: %s", sqlQuery, r.Error())
}

func ResourceSqlTable() *schema.Resource {
	tableSchema := common.StructToSchema(SqlTableInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	update := updateFunctionFactory("/unity-catalog/tables", []string{
		"owner", "name", "data_source_format", "columns", "storage_location",
		"view_definition", "comment", "properties"})
	return common.Resource{
		Schema: tableSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if d.Get("table_type") != "EXTERNAL" {
				return nil
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti = new(SqlTableInfo)
			common.DataToStructPointer(d, tableSchema, ti)
			if err := ti.initCluster(ctx, d, c); err != nil {
				return err
			}
			if err := ti.createTable(ctx, d, tableSchema, c); err != nil {
				return err
			}
			d.SetId(ti.FullName())
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ti, err := NewSqlTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ti, tableSchema, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSqlTablesAPI(ctx, c).deleteTable(d.Id())
		},
	}.ToResource()
}
