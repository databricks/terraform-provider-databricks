package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccJobResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "databricks_job" "this" {
					new_cluster  {
					  autoscale  {
						min_workers = 2
						max_workers = 3
					  }
					  instance_pool_id = "%s"
					  spark_version = "%s"
					}
					notebook_task {
						notebook_path = "/Production/MakeFeatures"
					}
					email_notifications {
						no_alert_for_skipped_runs = true
					}
					name = "%s"
					timeout_seconds = 3600
					max_retries = 1
					max_concurrent_runs = 1
				  }`, CommonInstancePoolID(), CommonRuntimeVersion(),
					qa.RandomLongName()),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					acceptance.ResourceCheck("databricks_job.this",
						func(client *common.DatabricksClient, id string) error {
							job, err := NewJobsAPI(client).Read(id)
							assert.NoError(t, err)
							assert.NotNil(t, job.Settings)
							assert.NotNil(t, job.Settings.NewCluster)
							assert.NotNil(t, job.Settings.NewCluster.Autoscale)
							assert.NotNil(t, job.Settings.NotebookTask)
							assert.Equal(t, 2, int(job.Settings.NewCluster.Autoscale.MinWorkers))
							assert.Equal(t, 3, int(job.Settings.NewCluster.Autoscale.MaxWorkers))
							assert.Equal(t, CommonRuntimeVersion(), job.Settings.NewCluster.SparkVersion)
							assert.Equal(t, "/Production/MakeFeatures", job.Settings.NotebookTask.NotebookPath)
							assert.Equal(t, 3600, int(job.Settings.TimeoutSeconds))
							assert.Equal(t, 1, int(job.Settings.MaxRetries))
							assert.Equal(t, 1, int(job.Settings.MaxConcurrentRuns))
							return nil
						}),
				),
			},
		},
		CheckDestroy: acceptance.ResourceCheck("databricks_job.this",
			func(client *common.DatabricksClient, id string) error {
				_, err := NewJobsAPI(client).Read(id)
				if err != nil {
					return nil
				}
				return errors.New("resource job is not cleaned up")
			}),
	})
}
