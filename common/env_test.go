package common

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonEnvironmentClient(t *testing.T) {
	ResetCommonEnvironmentClient()
	defer CleanupEnvironment()()
	os.Setenv("DATABRICKS_TOKEN", ".")
	os.Setenv("DATABRICKS_HOST", ".")
	c := CommonEnvironmentClient()
	c2 := CommonEnvironmentClient()
	ctx := context.Background()
	assert.Equal(t, c2.userAgent(ctx), c.userAgent(ctx))
	assert.Equal(t, "databricks-tf-provider/"+version+" (+unknown) terraform/unknown", c.userAgent(ctx))

	ctx = context.WithValue(ctx, ResourceName, "cluster")
	ctx = context.WithValue(ctx, TerraformVersion, "0.12")

	assert.Equal(t, "databricks-tf-provider/"+version+" (+cluster) terraform/0.12", c.userAgent(ctx))
}
