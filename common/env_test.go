package common

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonEnvironmentClient(t *testing.T) {
	ResetCommonEnvironmentClient()
	defer CleanupEnvironment()()
	os.Setenv("DATABRICKS_TOKEN", ".")
	os.Setenv("DATABRICKS_HOST", ".")
	os.Setenv("DATABRICKS_DEBUG_HEADERS", "true")
	os.Setenv("DATABRICKS_DEBUG_TRUNCATE_BYTES", "1024")
	c := CommonEnvironmentClient()
	c2 := CommonEnvironmentClient()
	ctx := context.Background()
	assert.Equal(t, c2.userAgent(ctx), c.userAgent(ctx))
	assert.Equal(t, "databricks-tf-provider/"+version+" (+unknown) terraform/unknown", c.userAgent(ctx))

	ctx = context.WithValue(ctx, ResourceName, "cluster")
	c.Provider = &schema.Provider{
		TerraformVersion: "0.12",
	}
	assert.Equal(t, "databricks-tf-provider/"+version+" (+cluster) terraform/0.12", c.userAgent(ctx))

	defer func() {
		paniced := recover()
		log.Printf("[INFO] paniced with: %s", paniced)
		require.NotNil(t, paniced, "Must have paniced!")
	}()
	commonClient = nil
	onceClient = sync.Once{}
	os.Setenv("ARM_ENVIRONMENT", "ANY")
	// and this one will panic
	CommonEnvironmentClient()
}
