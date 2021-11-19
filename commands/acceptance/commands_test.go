package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestAccContext(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
	clusterID := clusterInfo.ClusterID
	ctx := context.Background()
	c := commands.NewCommandsAPI(ctx, client)

	result := c.Execute(clusterID, "python", `print('hello world')`)
	assert.Equal(t, "hello world", result.Text())

	// exceptions are regexed away for readability
	result = c.Execute(clusterID, "python", `raise Exception("Not Found")`)
	qa.AssertErrorStartsWith(t, result.Err(), "Not Found")

	// but errors are not
	result = c.Execute(clusterID, "python", `raise KeyError("foo")`)
	qa.AssertErrorStartsWith(t, result.Err(), "KeyError: 'foo'")

	// so it is more clear to read and debug
	result = c.Execute(clusterID, "python", `return 'hello world'`)
	qa.AssertErrorStartsWith(t, result.Err(), "SyntaxError: 'return' outside function")

	result = c.Execute(clusterID, "python", `"Hello World!"`)
	assert.Equal(t, "'Hello World!'", result.Text())

	result = c.Execute(clusterID, "python", `
 		print("Hello World!")
 		dbutils.notebook.exit("success")`)
	assert.Equal(t, "success", result.Text())
}
