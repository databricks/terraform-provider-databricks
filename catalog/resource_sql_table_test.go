package catalog

import (
	"testing"

	clustersApi "github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestSqlTableCreateStatement(t *testing.T) {
	ti := &SqlTableInfo{
		Name:             "bar",
		CatalogName:      "main",
		SchemaName:       "foo",
		TableType:        "EXTERNAL",
		DataSourceFormat: "DELTA",
		//ColumnInfos:           []SqlColumnInfo,
		StorageLocation:       "s3://ext-main/foo/bar1",
		StorageCredentialName: "somecred",
		//ViewDefinition        string            `json:"view_definition,omitempty"`
		Comment: "terraform managed",
		//Properties            map[string]string `json:"properties,omitempty"`
	}
	stmt := ti.buildTableCreateStatement()
	assert.Contains(t, stmt, "CREATE EXTERNAL TABLE bar")
	assert.Contains(t, stmt, "USING DELTA")
	assert.Contains(t, stmt, "LOCATION 's3://ext-main/foo/bar1' WITH CREDENTIAL somecred")
	assert.Contains(t, stmt, "COMMENT 'terraform managed'")
}

func TestSqlTableCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		HCL: `
		name               = "bar"
		catalog_name       = "main"
		schema_name        = "foo"
		table_type         = "MANAGED"
		data_source_format = "DELTA"
		storage_location   = "s3://ext-main/foo/bar1"
		column {
			name      = "id"
			position  = 0
			type_name = "INT"
			type_text = "int"
		}
		column {
			name      = "name"
			position  = 1
			type_name = "STRING"
			type_text = "varchar(64)"
		}
		comment = "this table is managed by terraform"
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlTable(),
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, 2, d.Get("column.#"))
	assert.Equal(t, "INT", d.Get("column.0.type_name"))
	assert.Equal(t, "varchar(64)", d.Get("column.1.type_text"))
	assert.Equal(t, "DELTA", d.Get("data_source_format"))
}

func TestSqlTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceTable())
}

func TestSqlTableUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]interface{}{
					"columns": map[string]interface{}{
						"comment":            "",
						"name":               "id",
						"nullable":           true,
						"partition_index":    0,
						"position":           0,
						"type_interval_type": "",
						"type_json":          "",
						"type_name":          "string",
						"type_precision":     0,
						"type_scale":         0,
						"type_text":          "string",
					},
					"storage_location": "s3://ext-main/foo/bar1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "EXTERNAL",
					DataSourceFormat: "JSON",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "string",
							TypeText: "string",
						},
					},
				},
			},
		},
		Resource: ResourceTable(),
		Update:   true,
		ID:       "main.foo.bar",
		InstanceState: map[string]string{
			"catalog_name":       "main",
			"schema_name":        "foo",
			"name":               "bar",
			"table_type":         "EXTERNAL",
			"data_source_format": "JSON",
			"storage_location":   "s3://ext-main/foo/bar",
			"column": `[{
				\"comment\": \"\",
				\"name\": \"id\",
				\"nullable\": true,
				\"partition_index\": 0,
				\"position\": 0,
				\"type_interval_type\": \"\",
				\"type_json\": \"\",
				\"type_name\": \"int\",
				\"type_precision\": 0,
				\"type_scale\": 0,
				\"type_text\": \"int\"
			}]`,
		},
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		table_type = "EXTERNAL"
		data_source_format = "JSON"
		storage_location = "s3://ext-main/foo/bar1"
		
		column {
			name = "id"
			type_text = "string"
			type_name = "string"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}

func TestManagedSqlTableUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]interface{}{
					"columns": map[string]interface{}{
						"comment":            "",
						"name":               "id",
						"nullable":           true,
						"partition_index":    0,
						"position":           0,
						"type_interval_type": "",
						"type_json":          "",
						"type_name":          "string",
						"type_precision":     0,
						"type_scale":         0,
						"type_text":          "string",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "MANAGED",
					DataSourceFormat: "JSON",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "string",
							TypeText: "string",
						},
					},
				},
			},
		},
		Resource: ResourceTable(),
		Update:   true,
		ID:       "main.foo.bar",
		InstanceState: map[string]string{
			"catalog_name":       "main",
			"schema_name":        "foo",
			"name":               "bar",
			"table_type":         "MANAGED",
			"data_source_format": "JSON",
			"storage_location":   "s3://ext-main/foo/bar",
			"column": `[{
				\"comment\": \"\",
				\"name\": \"id\",
				\"nullable\": true,
				\"partition_index\": 0,
				\"position\": 0,
				\"type_interval_type\": \"\",
				\"type_json\": \"\",
				\"type_name\": \"int\",
				\"type_precision\": 0,
				\"type_scale\": 0,
				\"type_text\": \"int\"
			}]`,
		},
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		table_type = "MANAGED"
		data_source_format = "JSON"
		
		column {
			name = "id"
			type_text = "string"
			type_name = "string"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}

var createHighConcurrencyCluster = []qa.HTTPFixture{
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/list",
		Response:     map[string]any{},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/spark-versions",
		Response: clusters.SparkVersionsList{
			SparkVersions: []clusters.SparkVersion{
				{
					Version:     "7.1.x-cpu-ml-scala2.12",
					Description: "7.1 ML (includes Apache Spark 3.0.0, Scala 2.12)",
				},
			},
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/list-node-types",
		Response: clustersApi.ListNodeTypesResponse{
			NodeTypes: []clustersApi.NodeType{
				{
					NodeTypeId:     "Standard_F4s",
					InstanceTypeId: "Standard_F4s",
					MemoryMb:       8192,
					NumCores:       4,
					NodeInstanceType: &clustersApi.NodeInstanceType{
						LocalDisks:      1,
						InstanceTypeId:  "Standard_F4s",
						LocalDiskSizeGb: 16,
					},
				},
			},
		},
	},
	{
		Method:       "POST",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/create",
		ExpectedRequest: clusters.Cluster{
			AutoterminationMinutes: 10,
			ClusterName:            "terraform-table",
			NodeTypeID:             "Standard_F4s",
			SparkVersion:           "7.3.x-scala2.12",
			CustomTags: map[string]string{
				"ResourceClass": "SingleNode",
			},
			SparkConf: map[string]string{
				"spark.databricks.repl.allowedLanguages": "python,sql",
				"spark.databricks.cluster.profile":       "singleNode",
				"spark.master":                           "local[*]",
			},
		},
		Response: clusters.ClusterID{
			ClusterID: "bcd",
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/get?cluster_id=bcd",
		Response: clusters.ClusterInfo{
			ClusterID: "bcd",
			State:     "RUNNING",
			SparkConf: map[string]string{
				"spark.databricks.cluster.profile": "singleNode",
			},
		},
	},
}
