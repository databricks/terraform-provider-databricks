package common

import "context"

// WithCommandMock mocks all command executions for this client
func (c *DatabricksClient) WithCommandMock(mock CommandMock) {
	c.WithCommandExecutor(func(_ context.Context, _ *DatabricksClient) CommandExecutor {
		return commandExecutorMock{
			mock: mock,
		}
	})
}

// WithCommandExecutor sets command executor implementation to use
func (c *DatabricksClient) WithCommandExecutor(cef func(context.Context, *DatabricksClient) CommandExecutor) {
	c.commandFactory = cef
}

// CommandExecutor service
func (c *DatabricksClient) CommandExecutor(ctx context.Context) CommandExecutor {
	return c.commandFactory(ctx, c)
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) (string, error)

// CommandExecutorMock simplifies command testing
type commandExecutorMock struct {
	mock CommandMock
}

// Execute mock command with given mock function
func (c commandExecutorMock) Execute(clusterID, language, commandStr string) (string, error) {
	return c.mock(commandStr)
}

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(clusterID, language, commandStr string) (string, error)
}
