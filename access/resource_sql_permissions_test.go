package access

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableACLID(t *testing.T) {
	for id, ta := range map[string]SqlPermissions{
		"table/default.foo":   {Table: "foo"},
		"view/bar.foo":        {View: "foo", Database: "bar"},
		"database/bar":        {Database: "bar"},
		"catalog/":            {Catalog: true},
		"any file/":           {AnyFile: true},
		"anonymous function/": {AnonymousFunction: true},
	} {
		assert.Equal(t, id, ta.ID())
		ta2, err := loadTableACL(id)
		if ta.Database == "" && ta2.Database == "default" {
			ta.Database = "default"
		}
		assert.NoError(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, ta, ta2, id)
	}
}

func TestTableACLID_errors(t *testing.T) {
	for id, exp := range map[string]string{
		"table":      "ID must be two elements: table",
		"table/beep": "table must have two elements",
		"view/beep":  "view must have two elements",
		"vuew/beep":  "illegal ID type: vuew",
	} {
		_, err := loadTableACL(id)
		assert.EqualError(t, err, exp)
	}
}

type mockData map[string][][]string

func (md mockData) Execute(clusterID, language, commandStr string) common.CommandResults {
	data, ok := md[commandStr]
	if !ok {
		return common.CommandResults{
			ResultType: "error",
			Summary:    fmt.Sprintf("Query is not mocked: %s", commandStr),
		}
	}
	var x []any
	for _, a := range data {
		var y []any
		for _, b := range a {
			y = append(y, b)
		}
		x = append(x, y)
	}
	return common.CommandResults{
		ResultType: "table",
		Data:       x,
	}
}

func (md mockData) toCommandMock() func(string) common.CommandResults {
	return func(commandStr string) common.CommandResults {
		return md.Execute("_", "cobol", commandStr)
	}
}

func TestTableACLGrants(t *testing.T) {
	ta := SqlPermissions{Table: "foo", exec: mockData{
		"SHOW GRANT ON TABLE `default`.`foo`": {
			// principal, actionType, objType, objectKey
			{"users", "SELECT", "database", "foo"},
			{"users", "SELECT", "table", "`default`.`foo`"},
			{"users", "READ", "table", "`default`.`foo`"},
			{"users", "SELECT", "database", "default"},
			{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
		},
	}}
	err := ta.read()
	assert.NoError(t, err)
	assert.Len(t, ta.PrivilegeAssignments, 1)
	assert.Len(t, ta.PrivilegeAssignments[0].Privileges, 2)
}

func TestDatabaseACLGrants(t *testing.T) {
	ta := SqlPermissions{Database: "default",
		exec: mockData{
			"SHOW GRANT ON DATABASE `default`": {
				// principal, actionType, objType, objectKey
				// Test with and without backticks
				{"users", "SELECT", "database", "default"},
				{"users", "USAGE", "database", "`default`"},
			},
		}}
	err := ta.read()
	assert.NoError(t, err)
	assert.Len(t, ta.PrivilegeAssignments, 1)
	assert.Len(t, ta.PrivilegeAssignments[0].Privileges, 2)
}

type failedCommand string

func (fc failedCommand) Execute(clusterID, language, commandStr string) common.CommandResults {
	return common.CommandResults{
		ResultType: "error",
		Summary:    string(fc),
	}
}

func (fc failedCommand) toCommandMock() func(commandStr string) common.CommandResults {
	return func(commandStr string) common.CommandResults {
		return fc.Execute("..", "sql", commandStr)
	}
}

func TestTableACL_NotFound(t *testing.T) {
	ta := SqlPermissions{Table: "foo", exec: failedCommand("Table does not exist")}
	err := ta.read()
	assert.EqualError(t, err, "Table does not exist")
}

func TestTableACL_OtherError(t *testing.T) {
	ta := SqlPermissions{Table: "foo", exec: failedCommand("Some error")}
	err := ta.read()
	assert.EqualError(t, err, "cannot read current grants: Some error")
}

func TestTableACL_Revoke(t *testing.T) {
	ta := SqlPermissions{Table: "foo", exec: mockData{
		"SHOW GRANT ON TABLE `default`.`foo`": {
			{"users", "SELECT", "database", "foo"},
			{"users", "SELECT", "table", "`default`.`foo`"},
			{"users", "READ", "table", "`default`.`foo`"},
			{"users", "SELECT", "database", "default"},
		},
		"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":   {},
		"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`": {},
	}}
	err := ta.revoke()
	require.NoError(t, err)
}

func TestTableACL_Enforce(t *testing.T) {
	ta := SqlPermissions{
		Table: "foo",
		PrivilegeAssignments: []PrivilegeAssignment{
			{"engineers", []string{"MODIFY", "SELECT", "READ"}},
			{"support", []string{"SELECT"}},
		},
		exec: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
				{"interns", "DENIED_READ", "table", "`default`.`foo`"},
			},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":        {},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`":      {},
			"GRANT MODIFY, SELECT, READ ON TABLE `default`.`foo` TO `engineers`": {},
			"GRANT SELECT ON TABLE `default`.`foo` TO `support`":                 {},
		},
	}
	err := ta.enforce()
	require.NoError(t, err)
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
		Resource:     "/api/2.1/clusters/spark-versions",
		Response: compute.GetSparkVersionsResponse{
			Versions: []compute.SparkVersion{
				{
					Key:  "15.4.x-scala2.12",
					Name: "15.4 LTS (includes Apache Spark 3.5.0, Scala 2.12)",
				},
			},
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.1/clusters/list-node-types",
		Response: compute.ListNodeTypesResponse{
			NodeTypes: []compute.NodeType{
				{
					NodeTypeId:     "Standard_F4s",
					InstanceTypeId: "Standard_F4s",
					MemoryMb:       8192,
					NumCores:       4,
					NodeInstanceType: &compute.NodeInstanceType{
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
			ClusterName:            "terraform-table-acl",
			NodeTypeID:             "Standard_F4s",
			SparkVersion:           "15.4.x-scala2.12",
			DataSecurityMode:       "LEGACY_TABLE_ACL",
			NumWorkers:             1,
			// CustomTags: map[string]string{
			// 	"ResourceClass": "SingleNode",
			// },
			// SparkConf: map[string]string{
			// 	"spark.databricks.cluster.profile": "singleNode",
			// 	"spark.master":                     "local[*]",
			// },
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
			ClusterID:        "bcd",
			State:            "RUNNING",
			DataSecurityMode: "LEGACY_TABLE_ACL",
			// SparkConf: map[string]string{
			// 	"spark.databricks.cluster.profile": "singleNode",
			// },
		},
	},
}

var createSharedCluster = []qa.HTTPFixture{
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/list",
		Response:     map[string]any{},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.1/clusters/spark-versions",
		Response: compute.GetSparkVersionsResponse{
			Versions: []compute.SparkVersion{
				{
					Key:  "15.4.x-scala2.12",
					Name: "15.4 LTS (includes Apache Spark 3.5.0, Scala 2.12)",
				},
			},
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.1/clusters/list-node-types",
		Response: compute.ListNodeTypesResponse{
			NodeTypes: []compute.NodeType{
				{
					NodeTypeId:     "Standard_F4s",
					InstanceTypeId: "Standard_F4s",
					MemoryMb:       8192,
					NumCores:       4,
					NodeInstanceType: &compute.NodeInstanceType{
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
			ClusterName:            "terraform-table-acl",
			NodeTypeID:             "Standard_F4s",
			SparkVersion:           "15.4.x-scala2.12",
			DataSecurityMode:       "LEGACY_TABLE_ACL",
			NumWorkers:             1,
			// CustomTags: map[string]string{
			// 	"ResourceClass": "SingleNode",
			// },
			// SparkConf: map[string]string{
			// 	"spark.databricks.cluster.profile": "singleNode",
			// 	"spark.master":                     "local[*]",
			// },
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
			ClusterID:        "bcd",
			State:            "RUNNING",
			DataSecurityMode: "USER_ISOLATION",
		},
	},
}

func TestResourceSqlPermissions_Read(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"bob@example.com", "OWN", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
		}.toCommandMock(),
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Read:     true,
		New:      true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_ReadSharedCluster(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"bob@example.com", "OWN", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
		}.toCommandMock(),
		Fixtures: createSharedCluster,
		Resource: ResourceSqlPermissions(),
		Read:     true,
		New:      true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_Read_Error(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceSqlPermissions(),
		Read:     true,
		New:      true,
		ID:       "something",
	}.ExpectError(t, "ID must be two elements: something")
}

func TestResourceSqlPermissions_Read_ErrorCommand(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: failedCommand("does not clusters").toCommandMock(),
		Fixtures:    createHighConcurrencyCluster,
		Resource:    ResourceSqlPermissions(),
		ID:          "database/foo",
		Read:        true,
		New:         true,
	}.ExpectError(t, "cannot read current grants: does not clusters")
}

func TestResourceSqlPermissions_Create(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				// TODO: transform mockData into a sequence,
				// as this query should return two different results,
				// based on the order of execution
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":          {},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`":        {},
			"GRANT MODIFY, SELECT ON TABLE `default`.`foo` TO `serge@example.com`": {},
		}.toCommandMock(),
		HCL: `
		table = "foo"
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT", "MODIFY"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Create:   true,
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_Create_Catalog(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			// yes, space in the end is needed
			"SHOW GRANT ON CATALOG ": {
				{"users", "SELECT", "CATALOG$", "None"},
				{"users", "MODIFY", "CATALOG$", "None"},
			},
			"REVOKE ALL PRIVILEGES ON CATALOG  FROM `users`":  {},
			"GRANT SELECT ON CATALOG  TO `serge@example.com`": {},
		}.toCommandMock(),
		HCL: `
		catalog = true
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Create:   true,
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_Create_Error(t *testing.T) {
	qa.ResourceFixture{
		HCL: `table = "foo"
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}`,
		CommandMock: failedCommand("Some error").toCommandMock(),
		Fixtures:    createHighConcurrencyCluster,
		Resource:    ResourceSqlPermissions(),
		Create:      true,
	}.ExpectError(t, "cannot read current grants: Some error")
}

func TestResourceSqlPermissions_Create_Error2(t *testing.T) {
	qa.ResourceFixture{
		HCL: `table = "foo"
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}`,
		CommandMock: func(commandStr string) common.CommandResults {
			md := mockData{
				"SHOW GRANT ON TABLE `default`.`foo`": {},
			}
			if _, ok := md[commandStr]; ok {
				return md.toCommandMock()(commandStr)
			}
			return common.CommandResults{
				ResultType: "error",
				Cause:      "com.x.y.z.d.Exceptions$SQLExecutionException: org.apache.spark.s...",
				Summary:    "Error in SQL statement: ParseException: \nAction Unknown ActionType READ cannot be granted on tab... (127 more bytes)",
			}
		},
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Create:   true,
	}.ExpectError(t, "cannot execute GRANT READ, MODIFY, SELECT ON TABLE `default`.`foo` TO `serge@example.com`: Action Unknown ActionType READ cannot be granted on tab... (127 more bytes)")
}

func TestResourceSqlPermissions_Update(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				// TODO: transform mockData into a sequence,
				// as this query should return two different results,
				// based on the order of execution
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":                {},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`":              {},
			"GRANT READ, MODIFY, SELECT ON TABLE `default`.`foo` TO `serge@example.com`": {},
		}.toCommandMock(),
		InstanceState: map[string]string{
			"table": "foo",
		},
		HCL: `
		table = "foo"
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Update:   true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_Delete(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":   {},
			"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`": {},
		}.toCommandMock(),
		HCL: `
		table = "foo"
		privilege_assignments {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Delete:   true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceSqlPermissions_CornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlPermissions(), qa.CornerCaseID("database/foo"))
}

func TestResourceSqlPermissions_NoUpdateAnyFile(t *testing.T) {
	d, err := qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON ANY FILE ": {
				{"users", "SELECT", "ANY_FILE", "None"},
			},
		}.toCommandMock(),
		HCL: `
		any_file = "true"
		privilege_assignments {
			principal = "users"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Update:   true,
		InstanceState: map[string]string{
			"any_file":                             "true",
			"privilege_assignments.#":              "1",
			"privilege_assignments.0.principal":    "users",
			"privilege_assignments.0.privileges.#": "1",
			"privilege_assignments.0.privileges.0": "SELECT",
		},
		ID: "any file/",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, d.Get("privilege_assignments.#"))
	assert.Equal(t, 1, d.Get("privilege_assignments.0.privileges.#"))
	assert.Equal(t, "users", d.Get("privilege_assignments.0.principal"))
	assert.Equal(t, "SELECT", d.Get("privilege_assignments.0.privileges.0"))
	assert.Equal(t, true, d.Get("any_file"))
}

func TestResourceSqlPermissions_NoUpdateAnonymousFunction(t *testing.T) {
	d, err := qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON ANONYMOUS FUNCTION ": {
				{"users", "SELECT", "ANONYMOUS_FUNCTION", "None"},
			},
		}.toCommandMock(),
		HCL: `
		anonymous_function = "true"
		privilege_assignments {
			principal = "users"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceSqlPermissions(),
		Update:   true,
		InstanceState: map[string]string{
			"anonymous_function":                   "true",
			"privilege_assignments.#":              "1",
			"privilege_assignments.0.principal":    "users",
			"privilege_assignments.0.privileges.#": "1",
			"privilege_assignments.0.privileges.0": "SELECT",
		},
		ID: "anonymous function/",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, d.Get("privilege_assignments.#"))
	assert.Equal(t, 1, d.Get("privilege_assignments.0.privileges.#"))
	assert.Equal(t, "users", d.Get("privilege_assignments.0.principal"))
	assert.Equal(t, "SELECT", d.Get("privilege_assignments.0.privileges.0"))
	assert.Equal(t, true, d.Get("anonymous_function"))
}
