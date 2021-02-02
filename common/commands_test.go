package common

import (
	"context"
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
	c.WithCommandMock(func(commandStr string) (string, error) {
		called = true
		assert.Equal(t, "print 1", commandStr)
		return "done", nil
	})
	ctx := context.Background()
	res, err := c.CommandExecutor(ctx).Execute("irrelevant", "python", "print 1")

	assert.Equal(t, true, called)
	assert.Equal(t, "done", res)
	assert.NoError(t, err, err)
}
