package databricks

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAzureJobResource(t *testing.T) {
	//var secretScope model.Secre
	var job model.Job
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAzureJobResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureJobResourceNewCluster(),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureJobResourceExists("databricks_job.my_job", &job),
					// verify remote values
					testAzureJobValuesNewCluster(t, &job),
				),
				Destroy: false,
			},
		},
	})
}

func testAzureJobResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_job" {
			continue
		}
		idInt, err := strconv.ParseInt(rs.Primary.ID, 10, 32)
		if err != nil {
			return err
		}
		_, err = client.Jobs().Read(idInt)
		if err != nil {
			return nil
		}
		return errors.New("resource job is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAzureJobResourceExists(n string, job *model.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
		idInt, err := strconv.ParseInt(rs.Primary.ID, 10, 32)
		if err != nil {
			return err
		}
		resp, err := conn.Jobs().Read(idInt)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*job = resp
		return nil
	}
}

// Assertions are based off of the resource definition defined in function: testAzureJobResourceNewCluster
func testAzureJobValuesNewCluster(t *testing.T, job *model.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.NotNil(t, job.Settings)
		assert.NotNil(t, job.Settings.NewCluster)
		assert.NotNil(t, job.Settings.NewCluster.Autoscale)
		assert.NotNil(t, job.Settings.NotebookTask)
		assert.Equal(t, 2, int(job.Settings.NewCluster.Autoscale.MinWorkers))
		assert.Equal(t, 3, int(job.Settings.NewCluster.Autoscale.MaxWorkers))
		assert.Equal(t, "6.4.x-scala2.11", job.Settings.NewCluster.SparkVersion)
		assert.Equal(t, "Standard_DS3_v2", job.Settings.NewCluster.NodeTypeID)
		assert.Equal(t, "/Users/jane.doe@databricks.com/my-demo-notebook", job.Settings.NotebookTask.NotebookPath)
		assert.Equal(t, "my-demo-notebook", job.Settings.Name)
		assert.Equal(t, 3600, int(job.Settings.TimeoutSeconds))
		assert.Equal(t, 1, int(job.Settings.MaxRetries))
		assert.Equal(t, 1, int(job.Settings.MaxConcurrentRuns))
		return nil
	}
}

func testAzureJobResourceNewCluster() string {
	return `
	resource "databricks_job" "my_job" {
	  new_cluster  {
		autoscale  {
		  min_workers = 2
		  max_workers = 3
		}
		spark_version = "6.4.x-scala2.11"
		node_type_id = "Standard_DS3_v2"
	  }
	  notebook_path = "/Users/jane.doe@databricks.com/my-demo-notebook"
	  name = "my-demo-notebook"
	  timeout_seconds = 3600
	  max_retries = 1
	  max_concurrent_runs = 1
	}
	`
}
