package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandMock(t *testing.T) {
	defer CleanupEnvironment()()
	err := os.Setenv("DATABRICKS_TOKEN", ".")
	assert.NoError(t, err)
	err = os.Setenv("DATABRICKS_HOST", ".")
	assert.NoError(t, err)
	c := DatabricksClient{}
	err = c.Configure()
	assert.NoError(t, err)

	called := false
	c.WithCommandMock(func(commandStr string) (string, error) {
		called = true
		assert.Equal(t, "print 1", commandStr)
		return "done", nil
	})
	res, err := c.CommandExecutor().Execute("irrelevant", "python", "print 1")

	assert.Equal(t, true, called)
	assert.Equal(t, "done", res)
	assert.NoError(t, err, err)
}
