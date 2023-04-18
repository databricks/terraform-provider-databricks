package catalog

import (
	"testing"

	clustersApi "github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceSqlTableCreateStatement_External(t *testing.T) {
	ti := &SqlTableInfo{
		Name:                  "bar",
		CatalogName:           "main",
		SchemaName:            "foo",
		TableType:             "EXTERNAL",
		DataSourceFormat:      "DELTA",
		StorageLocation:       "s3://ext-main/foo/bar1",
		StorageCredentialName: "somecred",
		Comment:               "terraform managed",
	}
	stmt := ti.buildTableCreateStatement()
	assert.Contains(t, stmt, "CREATE EXTERNAL TABLE main.foo.bar")
	assert.Contains(t, stmt, "USING DELTA")
	assert.Contains(t, stmt, "LOCATION 's3://ext-main/foo/bar1' WITH (CREDENTIAL `somecred`)")
	assert.Contains(t, stmt, "COMMENT 'terraform managed'")
}

func TestResourceSqlTableCreateStatement_View(t *testing.T) {
	ti := &SqlTableInfo{
		Name:                  "bar",
		CatalogName:           "main",
		SchemaName:            "foo",
		TableType:             "VIEW",
		DataSourceFormat:      "DELTA",
		StorageLocation:       "s3://ext-main/foo/bar1",
		StorageCredentialName: "somecred",
		Comment:               "terraform managed",
		Properties: map[string]string{
			"one":   "two",
			"three": "four",
		},
	}
	stmt := ti.buildTableCreateStatement()
	assert.Contains(t, stmt, "CREATE VIEW main.foo.bar")
	assert.NotContains(t, stmt, "USING DELTA")
	assert.NotContains(t, stmt, "LOCATION 's3://ext-main/foo/bar1' WITH CREDENTIAL somecred")
	assert.Contains(t, stmt, "COMMENT 'terraform managed'")
	assert.Contains(t, stmt, "'one'='two'")
}

func TestResourceSqlTableSerializeProperties(t *testing.T) {
	ti := &SqlTableInfo{
		Properties: map[string]string{
			"one":   "two",
			"three": "four",
		},
	}
	assert.Contains(t, ti.serializeProperties(), "'one'='two'")
	assert.Contains(t, ti.serializeProperties(), "'three'='four'")
}

func TestResourceSqlTableCreateTable(t *testing.T) {
	_, err := qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "",
				Data:       nil,
			}
		},
		HCL: `
		name               = "bar"
		catalog_name       = "main"
		schema_name        = "foo"
		table_type         = "MANAGED"
		data_source_format = "DELTA"
		storage_location   = "abfss:container@account/somepath"
	  
		column {
		  name      = "id"
		  type_text = "int"
		}
		column {
		  name      = "name"
		  type_text = "string"
		  comment   = "name of thing"
		}
		comment = "this table is managed by terraform"
		`,
		Fixtures: append([]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: SqlTableInfo{
					Name:                  "bar",
					CatalogName:           "main",
					SchemaName:            "foo",
					TableType:             "EXTERNAL",
					DataSourceFormat:      "DELTA",
					StorageLocation:       "s3://ext-main/foo/bar1",
					StorageCredentialName: "somecred",
					Comment:               "terraform managed",
					Properties: map[string]string{
						"one":   "two",
						"three": "four",
					},
				},
			},
		}, useExistingClusterForSql...),
		Create:   true,
		Resource: ResourceSqlTable(),
	}.Apply(t)
	assert.NoError(t, err)
}

func TestResourceSqlTableUpdateTable(t *testing.T) {
	d, err := qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {

			return common.CommandResults{
				ResultType: "",
				Data:       nil,
			}
		},
		HCL: `
		name               = "bar"
		catalog_name       = "main"
		schema_name        = "foo"
		table_type         = "EXTERNAL"
		data_source_format = "DELTA"
		storage_location   = "abfss:container@account/somepath"
		comment 		   = "this table is managed by terraform"
		cluster_id         = "gone"
		properties	       = {
			"one" = "two"
		}
		column {
			name      = "one"
			type_text = "string"
			comment   = "managed comment"
			nullable  = false
		}
		column {
			name      = "two"
			type_text = "string"
		}
		`,
		InstanceState: map[string]string{
			"name":               "bar",
			"catalog_name":       "main",
			"schema_name":        "foo",
			"table_type":         "EXTERNAL",
			"data_source_format": "DELTA",
			"storage_location":   "s3://ext-main/foo/bar1",
			"comment":            "terraform managed",
			"column.#":           "2",
			"column.0.name":      "one",
			"column.0.type_text": "string",
			"column.0.comment":   "old comment",
			"column.0.nullable":  "false",
			"column.1.name":      "two",
			"column.1.type_text": "string",
		},
		Fixtures: append([]qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.1/unity-catalog/tables/main.foo.bar",
				ReuseRequest: true,
				Response: SqlTableInfo{
					Name:                  "bar",
					CatalogName:           "main",
					SchemaName:            "foo",
					TableType:             "EXTERNAL",
					DataSourceFormat:      "DELTA",
					StorageLocation:       "s3://ext-main/foo/bar1",
					StorageCredentialName: "somecred",
					Comment:               "terraform managed",
					Properties: map[string]string{
						"delta.lastCommitTimestamp": "87698768",
						"delta.minWriterVersion":    "1",
					},
					ColumnInfos: []SqlColumnInfo{
						{
							Name:     "one",
							TypeText: "string",
							Comment:  "managed comment",
							Nullable: false,
						},
						{
							Name:     "two",
							TypeText: "string",
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/start",
				ExpectedRequest: clusters.ClusterID{
					ClusterID: "gone",
				},
				Status: 404,
			},
		}, createClusterForSql...),
		Resource: ResourceSqlTable(),
		ID:       "main.foo.bar",
		Update:   true,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "bar", d.Get("name"))
}

func TestResourceSqlTableDeleteTable(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			assert.Equal(t, "DROP TABLE main.foo.bar", commandStr)
			return common.CommandResults{
				ResultType: "",
				Data:       nil,
			}
		},
		Resource: ResourceSqlTable(),
		State: map[string]any{
			"name":               "bar",
			"catalog_name":       "main",
			"schema_name":        "foo",
			"table_type":         "EXTERNAL",
			"data_source_format": "DELTA",
			"storage_location":   "abfss:container@account/somepath",
			"comment":            "this table is managed by terraform",
		},
		Fixtures: append([]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: SqlTableInfo{
					Name:                  "bar",
					CatalogName:           "main",
					SchemaName:            "foo",
					TableType:             "EXTERNAL",
					DataSourceFormat:      "DELTA",
					StorageLocation:       "s3://ext-main/foo/bar1",
					StorageCredentialName: "somecred",
					Comment:               "terraform managed",
					Properties: map[string]string{
						"delta.lastCommitTimestamp": "87698768",
						"delta.minWriterVersion":    "1",
					},
				},
			},
		}, createClusterForSql...),
		Delete: true,
		ID:     "main.foo.bar",
	}.ApplyNoError(t)
}

func TestResourceSqlTableUpdateView_Definition(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			assert.NotContains(t, commandStr, "COMMENT ON VIEW", "Changing comments on a VIEW requires new. Resource should not attempt to alter comments via SQL command")
			return common.CommandResults{
				ResultType: "",
				Data:       nil,
			}
		},
		HCL: `
		name               = "barview"
		catalog_name       = "main"
		schema_name        = "foo"
		table_type         = "VIEW"
		cluster_id         = "existingcluster"
		view_definition    = "SELECT * FROM main.foo.bar2"
		`,
		InstanceState: map[string]string{
			"name":            "barview",
			"catalog_name":    "main",
			"schema_name":     "foo",
			"table_type":      "VIEW",
			"view_definition": "SELECT * FROM main.foo.bar",
		},
		Fixtures: append([]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.barview",
				Response: SqlTableInfo{
					Name:           "barview",
					CatalogName:    "main",
					SchemaName:     "foo",
					TableType:      "VIEW",
					ViewDefinition: "SELECT * FROM main.foo.bar",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.barview",
				Response: SqlTableInfo{
					Name:           "barview",
					CatalogName:    "main",
					SchemaName:     "foo",
					TableType:      "VIEW",
					ViewDefinition: "SELECT * FROM main.foo.bar2",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/start",
				ExpectedRequest: clusters.ClusterID{
					ClusterID: "existingcluster",
				},
				Status: 404,
			},
		}, createClusterForSql...),
		Resource: ResourceSqlTable(),
		Update:   true,
		ID:       "main.foo.barview",
	}.ApplyNoError(t)
}

func TestResourceSqlTableUpdateView_Comments(t *testing.T) {
	d, err := qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			assert.NotContains(t, commandStr, "COMMENT ON VIEW", "Changing comments on a VIEW requires new. Resource should not attempt to alter comments via SQL command")
			return common.CommandResults{
				ResultType: "",
				Data:       nil,
			}
		},
		HCL: `
		name               = "barview"
		catalog_name       = "main"
		schema_name        = "foo"
		table_type         = "VIEW"
		comment 		   = "this table is managed by terraform"
		cluster_id         = "existingcluster"
		view_definition    = "SELECT * FROM main.foo.bar"

		`,
		InstanceState: map[string]string{
			"name":            "barview",
			"catalog_name":    "main",
			"schema_name":     "foo",
			"table_type":      "VIEW",
			"comment":         "terraform managed",
			"view_definition": "SELECT * FROM main.foo.bar",
		},
		Fixtures: append([]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.barview",
				Response: SqlTableInfo{
					Name:           "barview",
					CatalogName:    "main",
					SchemaName:     "foo",
					TableType:      "VIEW",
					ViewDefinition: "SELECT * FROM main.foo.bar",
					Comment:        "to be changed (requires new)",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/start",
				ExpectedRequest: clusters.ClusterID{
					ClusterID: "existingcluster",
				},
				Status: 404,
			},
		}, createClusterForSql...),
		Resource:    ResourceSqlTable(),
		Create:      true,
		RequiresNew: true,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "barview", d.Get("name"))
}

var baseClusterFixture = []qa.HTTPFixture{
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
				{
					Version:     "7.3.x-scala2.12",
					Description: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
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
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/get?cluster_id=existingcluster",
		Response: clusters.ClusterInfo{
			ClusterID:   "existingcluster",
			ClusterName: "terraform-sql-table",
			State:       "RUNNING",
			SparkConf: map[string]string{
				"spark.databricks.acl.dfAclsEnabled": "true",
				"spark.databricks.cluster.profile":   "singleNode",
			},
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/get?cluster_id=gone",
		Status:       404,
	},
	{
		Method:       "POST",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/start",
		ExpectedRequest: clusters.ClusterID{
			ClusterID: "existingcluster",
		},
	},
}

var useExistingClusterForSql = append([]qa.HTTPFixture{
	{
		Method:   "GET",
		Resource: "/api/2.0/clusters/list?",
		Response: clusters.ClusterList{
			Clusters: []clusters.ClusterInfo{
				{
					ClusterID:   "existingcluster",
					ClusterName: "terraform-sql-table",
					State:       "RUNNING",
					SparkConf: map[string]string{
						"spark.databricks.acl.dfAclsEnabled": "true",
						"spark.databricks.cluster.profile":   "singleNode",
					},
				},
			},
		},
	},
}, baseClusterFixture...)

var createClusterForSql = append([]qa.HTTPFixture{
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/list?",
		Response:     map[string]any{},
	},
	{
		Method:       "POST",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/create",
		Response: clusters.ClusterID{
			ClusterID: "existingcluster",
		},
	},
}, baseClusterFixture...)

func TestResourceSqlTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlTable())
}
