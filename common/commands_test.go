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
	c.WithCommandMock(func(commandStr string) CommandResults {
		called = true
		assert.Equal(t, "print 1", commandStr)
		return CommandResults{
			ResultType: "text",
			Data:       "done",
		}
	})
	ctx := context.Background()
	cr := c.CommandExecutor(ctx).Execute("irrelevant", "python", "print 1")

	assert.Equal(t, true, called)
	assert.Equal(t, false, cr.Failed())
	assert.Equal(t, "done", cr.Text())
}
