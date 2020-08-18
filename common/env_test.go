package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonEnvironmentClient(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("DATABRICKS_TOKEN", ".")
	os.Setenv("DATABRICKS_HOST", ".")
	c := CommonEnvironmentClient()
	c2 := CommonEnvironmentClient()
	assert.Equal(t, UserAgent(), c.userAgent)
	assert.Equal(t, c2.userAgent, c.userAgent)
}
