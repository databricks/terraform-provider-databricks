package access

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableACLID(t *testing.T) {
	for id, ta := range map[string]TableACL{
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
	var x []interface{}
	for _, a := range data {
		var y []interface{}
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
	ta := TableACL{Table: "foo", exec: mockData{
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
	assert.Len(t, ta.Grants, 1)
	assert.Len(t, ta.Denies, 1)
	assert.Len(t, ta.Grants[0].Privileges, 2)
	assert.Len(t, ta.Denies[0].Privileges, 1)
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
	ta := TableACL{Table: "foo", exec: failedCommand("Table does not exist")}
	err := ta.read()
	assert.EqualError(t, err, "Table does not exist")
}

func TestTableACL_OtherError(t *testing.T) {
	ta := TableACL{Table: "foo", exec: failedCommand("Some error")}
	err := ta.read()
	assert.EqualError(t, err, "Some error")
}

func TestTableACL_Revoke(t *testing.T) {
	ta := TableACL{Table: "foo", exec: mockData{
		"SHOW GRANT ON TABLE `default`.`foo`": {
			{"users", "SELECT", "database", "foo"},
			{"users", "SELECT", "table", "`default`.`foo`"},
			{"users", "READ", "table", "`default`.`foo`"},
			{"users", "SELECT", "database", "default"},
			{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
		},
		"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `users`":   {},
		"REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM `interns`": {},
	}}
	err := ta.revoke()
	require.NoError(t, err)
}

func TestTableACL_Enforce(t *testing.T) {
	ta := TableACL{
		Table: "foo",
		Grants: []TablePermissions{
			{"engineers", []string{"MODIFY", "SELECT", "READ"}},
			{"support", []string{"SELECT"}},
		},
		Denies: []TablePermissions{
			{"foo@example.com", []string{"SELECT"}},
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
			"DENY SELECT ON TABLE `default`.`foo` TO `foo@example.com`":          {},
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
		Response:     map[string]interface{}{},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/spark-versions",
		Response: compute.SparkVersionsList{
			SparkVersions: []compute.SparkVersion{
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
		Response: compute.NodeTypeList{
			NodeTypes: []compute.NodeType{
				{
					NodeTypeID:     "Standard_F4s",
					InstanceTypeID: "Standard_F4s",
					MemoryMB:       8192,
					NumCores:       4,
					NodeInstanceType: &compute.NodeInstanceType{
						LocalDisks:      1,
						InstanceTypeID:  "Standard_F4s",
						LocalDiskSizeGB: 16,
						LocalNVMeDisks:  0,
					},
				},
			},
		},
	},
	{
		Method:       "POST",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/create",
		ExpectedRequest: compute.Cluster{
			AutoterminationMinutes: 10,
			ClusterName:            "terrraform-table-acl",
			NodeTypeID:             "Standard_F4s",
			SparkVersion:           "7.3.x-scala2.12",
			CustomTags: map[string]string{
				"ResourceClass": "SingleNode",
			},
			SparkConf: map[string]string{
				"spark.databricks.acl.dfAclsEnabled": "true",
				"spark.master":                       "local[*]",
			},
		},
		Response: compute.ClusterID{
			ClusterID: "bcd",
		},
	},
	{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/get?cluster_id=bcd",
		Response: compute.ClusterInfo{
			ClusterID: "bcd",
			State:     "RUNNING",
			SparkConf: map[string]string{
				"spark.databricks.acl.dfAclsEnabled": "true",
			},
		},
	},
}

func TestResourceTableACL_Read(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: mockData{
			"SHOW GRANT ON TABLE `default`.`foo`": {
				{"users", "SELECT", "database", "foo"},
				{"users", "SELECT", "table", "`default`.`foo`"},
				{"users", "READ", "table", "`default`.`foo`"},
				{"users", "SELECT", "database", "default"},
				{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
			},
		}.toCommandMock(),
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceTableACL(),
		Read:     true,
		New:      true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceTableACL_Read_Error(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceTableACL(),
		Read:     true,
		New:      true,
		ID:       "something",
	}.ExpectError(t, "ID must be two elements: something")
}

func TestResourceTableACL_Read_ErrorCommand(t *testing.T) {
	qa.ResourceFixture{
		CommandMock: failedCommand("does not compute").toCommandMock(),
		Fixtures:    createHighConcurrencyCluster,
		Resource:    ResourceTableACL(),
		ID:          "database/foo",
		Read:        true,
		New:         true,
	}.ExpectError(t, "does not compute")
}

func TestResourceTableACL_Create(t *testing.T) {
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
			"DENY SELECT ON TABLE `default`.`foo` TO `users`":                            {},
		}.toCommandMock(),
		HCL: `
		table = "foo"
		grant {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		deny {
			principal = "users"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceTableACL(),
		Create:   true,
	}.ApplyNoError(t)
}

func TestResourceTableACL_Create_Error(t *testing.T) {
	qa.ResourceFixture{
		HCL: `table = "foo"
		grant {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		deny {
			principal = "users"
			privileges = ["SELECT"]
		}`,
		CommandMock: failedCommand("Some error").toCommandMock(),
		Fixtures:    createHighConcurrencyCluster,
		Resource:    ResourceTableACL(),
		Create:      true,
	}.ExpectError(t, "Some error")
}

func TestResourceTableACL_Update(t *testing.T) {
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
			"DENY SELECT ON TABLE `default`.`foo` TO `users`":                            {},
		}.toCommandMock(),
		HCL: `
		table = "foo"
		grant {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		deny {
			principal = "users"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceTableACL(),
		Update:   true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceTableACL_Delete(t *testing.T) {
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
		grant {
			principal = "serge@example.com"
			privileges = ["SELECT", "READ", "MODIFY"]
		}
		deny {
			principal = "users"
			privileges = ["SELECT"]
		}
		`,
		Fixtures: createHighConcurrencyCluster,
		Resource: ResourceTableACL(),
		Delete:   true,
		ID:       "table/default.foo",
	}.ApplyNoError(t)
}

func TestResourceTableACL_CornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceTableACL(), "database/foo")
}
