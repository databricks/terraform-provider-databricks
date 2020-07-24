package databricks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccJobResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: fmt.Sprintf(`resource "databricks_job" "my_job" {
					new_cluster  {
					  autoscale  {
						min_workers = 2
						max_workers = 3
					  }
					  instance_pool_id = "%s"
					  spark_version = "%s"
					}
					notebook_path = "/Users/jane.doe@databricks.com/my-demo-notebook"
					name = "%s"
					timeout_seconds = 3600
					max_retries = 1
					max_concurrent_runs = 1
				  }`, service.CommonInstancePoolID(), service.CommonRuntimeVersion(),
					epoch.RandomLongName()),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					epoch.ResourceCheck("databricks_job.my_job",
						func(client *service.DatabricksClient, id string) error {
							idInt, err := strconv.ParseInt(id, 10, 32)
							assert.NoError(t, err)
							job, err := client.Jobs().Read(idInt)
							assert.NoError(t, err)
							assert.NotNil(t, job.Settings)
							assert.NotNil(t, job.Settings.NewCluster)
							assert.NotNil(t, job.Settings.NewCluster.Autoscale)
							assert.NotNil(t, job.Settings.NotebookTask)
							assert.Equal(t, 2, int(job.Settings.NewCluster.Autoscale.MinWorkers))
							assert.Equal(t, 3, int(job.Settings.NewCluster.Autoscale.MaxWorkers))
							assert.Equal(t, service.CommonRuntimeVersion(), job.Settings.NewCluster.SparkVersion)
							assert.Equal(t, "/Users/jane.doe@databricks.com/my-demo-notebook", job.Settings.NotebookTask.NotebookPath)
							assert.Equal(t, 3600, int(job.Settings.TimeoutSeconds))
							assert.Equal(t, 1, int(job.Settings.MaxRetries))
							assert.Equal(t, 1, int(job.Settings.MaxConcurrentRuns))
							return nil
						}),
				),
				Destroy: false,
			},
		},
		CheckDestroy: epoch.ResourceCheck("databricks_job.my_job",
			func(client *service.DatabricksClient, id string) error {
				idInt, err := strconv.ParseInt(id, 10, 32)
				assert.NoError(t, err)
				_, err = client.Jobs().Read(idInt)
				if err != nil {
					return nil
				}
				return errors.New("resource job is not cleaned up")
			}),
	})
}
