package common

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestCommandMock(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("DATABRICKS_TOKEN", ".")
	os.Setenv("DATABRICKS_HOST", ".")
	c := DatabricksClient{}
	c.Configure()

	called := false
	c.WithCommandMock(func(commandStr string) (string, error) {
		called = true
		assert.Equal(t, "print 1", commandStr)
		return "done", nil
	})
	res, err := c.CommandExecutor().Execute("irrelevant", "python", "print 1")

	assert.Equal(t, true, called)
	assert.Equal(t, "done", res)
	assert.NilError(t, err)
}
