package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandMock(t *testing.T) {
	c := DatabricksClient{
		Host:  ".",
		Token: ".",
	}
	err := c.Configure()
	assert.NoError(t, err)

	called := false
	c.WithCommandMock(func(commandStr string) CommandResults {
		called = true
		assert.Equal(t, "print 1", commandStr)
		return CommandResults{
			ResultType: "text",
			Data:       "done",
		}
	})
	cr := c.CommandExecutor().Execute("irrelevant", "python", "print 1")

	assert.Equal(t, true, called)
	assert.Equal(t, false, cr.Failed())
	assert.Equal(t, "done", cr.Text())
}

func TestCommandResultsScan(t *testing.T) {
	// TODO: fix it
	cr := CommandResults{
		ResultType: "table",
		Data: []interface{}{
			[]interface{}{"DENIED_SELECT", 1, false},
			[]interface{}{"DENIED_MODIFY", 2, true},
		},
	}
	var a string
	var b int
	var c bool
	for cr.Scan(&a, &b, &c) {
		fmt.Printf("%v / %v / %v\n", a, b, c)
	}
}
