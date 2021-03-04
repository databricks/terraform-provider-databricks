package access

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
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

type mockData struct {
	t    *testing.T
	sql  string
	data [][]string
}

func (md mockData) Execute(clusterID, language, commandStr string) common.CommandResults {
	assert.Equal(md.t, md.sql, commandStr)
	var x []interface{}
	for _, a := range md.data {
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

func TestTableACLGrants(t *testing.T) {
	ta := TableACL{Table: "foo"}
	// principal, actionType, objType, objectKey
	err := ta.read(mockData{t,
		"SHOW GRANT ON TABLE `default`.`foo`",
		[][]string{
			{"users", "SELECT", "database", "foo"},
			{"users", "SELECT", "table", "`default`.`foo`"},
			{"users", "READ", "table", "`default`.`foo`"},
			{"users", "SELECT", "database", "default"},
			{"interns", "DENIED_SELECT", "table", "`default`.`foo`"},
		}})
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

func TestTableACL_NotFound(t *testing.T) {
	ta := TableACL{Table: "foo"}
	err := ta.read(failedCommand("Table does not exist"))
	assert.EqualError(t, err, "Table does not exist")
}

func TestTableACL_OtherError(t *testing.T) {
	ta := TableACL{Table: "foo"}
	err := ta.read(failedCommand("Some error"))
	assert.EqualError(t, err, "Some error")
}
