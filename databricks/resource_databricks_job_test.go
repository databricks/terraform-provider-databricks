package databricks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccJobResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
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

func TestResourceJobCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/create",
			ExpectedRequest: model.JobSettings{
				ExistingClusterID: "abc",
				NotebookTask:      &model.NotebookTask{},
				SparkJarTask: &model.SparkJarTask{
					JarURI:        "dbfs://a/b/c.jar",
					MainClassName: "com.labs.BarMain",
				},
				Name:                   "Featurizer",
				MaxRetries:             3,
				MinRetryIntervalMillis: 5000,
				RetryOnTimeout:         true,
				MaxConcurrentRuns:      1,
			},
			Response: model.Job{
				JobID: 789,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=789",
			Response: model.Job{
				JobID: 789,
				Settings: &model.JobSettings{
					ExistingClusterID: "abc",
					NotebookTask:      &model.NotebookTask{},
					SparkJarTask: &model.SparkJarTask{
						JarURI:        "dbfs://a/b/c.jar",
						MainClassName: "com.labs.BarMain",
					},
					Name:                   "Featurizer",
					MaxRetries:             3,
					MinRetryIntervalMillis: 5000,
					RetryOnTimeout:         true,
					MaxConcurrentRuns:      1,
				},
			},
		},
	}, resourceJob, map[string]interface{}{
		"existing_cluster_id":       "abc",
		"jar_main_class_name":       "com.labs.BarMain",
		"jar_uri":                   "dbfs://a/b/c.jar",
		"max_concurrent_runs":       1,
		"max_retries":               3,
		"min_retry_interval_millis": 5000,
		"name":                      "Featurizer",
		"retry_on_timeout":          true,
	}, resourceJobCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/create",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceJob, map[string]interface{}{
		"existing_cluster_id":       "abc",
		"jar_main_class_name":       "com.labs.BarMain",
		"jar_uri":                   "dbfs://a/b/c.jar",
		"max_concurrent_runs":       1,
		"max_retries":               3,
		"min_retry_interval_millis": 5000,
		"name":                      "Featurizer",
		"retry_on_timeout":          true,
	}, resourceJobCreate)
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceJobRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=789",
			Response: model.Job{
				JobID: 789,
				Settings: &model.JobSettings{
					ExistingClusterID: "abc",
					NotebookTask:      &model.NotebookTask{},
					SparkJarTask: &model.SparkJarTask{
						JarURI:        "dbfs://a/b/c.jar",
						MainClassName: "com.labs.BarMain",
					},
					Name:                   "Featurizer",
					MaxRetries:             3,
					MinRetryIntervalMillis: 5000,
					RetryOnTimeout:         true,
					MaxConcurrentRuns:      1,
				},
			},
		},
	}, resourceJob, nil, actionWithID("789", resourceJobRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id(), "Id should not be empty")
	assert.Equal(t, "com.labs.BarMain", d.Get("jar_main_class_name"))
	assert.Equal(t, "dbfs://a/b/c.jar", d.Get("jar_uri"))
	assert.Equal(t, 789, d.Get("job_id"))
	assert.Equal(t, 1, d.Get("max_concurrent_runs"))
	assert.Equal(t, 3, d.Get("max_retries"))
	assert.Equal(t, 5000, d.Get("min_retry_interval_millis"))
}

func TestResourceJobRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=789",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceJob, nil, actionWithID("789", resourceJobRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceJobRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=789",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceJob, nil, actionWithID("789", resourceJobRead))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id(), "Id should not be empty for error reads")
}

func TestResourceJobUpdate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/reset",
			ExpectedRequest: model.UpdateJobRequest{
				JobID: 789,
				NewSettings: &model.JobSettings{
					ExistingClusterID: "abc",
					NotebookTask:      &model.NotebookTask{},
					SparkJarTask: &model.SparkJarTask{
						JarURI:        "dbfs://a/b/c.jar",
						MainClassName: "com.labs.Progress",
					},
					Name:                   "Featurizer New",
					MaxRetries:             3,
					MinRetryIntervalMillis: 5000,
					RetryOnTimeout:         true,
					MaxConcurrentRuns:      1,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=789",
			Response: model.Job{
				JobID: 789,
				Settings: &model.JobSettings{
					ExistingClusterID: "abc",
					NotebookTask:      &model.NotebookTask{},
					SparkJarTask: &model.SparkJarTask{
						JarURI:        "dbfs://a/b/c.jar",
						MainClassName: "com.labs.Progress",
					},
					Name:                   "Featurizer New",
					MaxRetries:             3,
					MinRetryIntervalMillis: 5000,
					RetryOnTimeout:         true,
					MaxConcurrentRuns:      1,
				},
			},
		},
	}, resourceJob, map[string]interface{}{
		"existing_cluster_id":       "abc",
		"jar_main_class_name":       "com.labs.Progress",
		"jar_uri":                   "dbfs://a/b/c.jar",
		"max_concurrent_runs":       1,
		"max_retries":               3,
		"min_retry_interval_millis": 5000,
		"name":                      "Featurizer New",
		"retry_on_timeout":          true,
	}, actionWithID("789", resourceJobUpdate))
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id(), "Id should be the same as in reading")
}

func TestResourceJobUpdate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/reset",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceJob, map[string]interface{}{
		"existing_cluster_id":       "abc",
		"jar_main_class_name":       "com.labs.Progress",
		"jar_uri":                   "dbfs://a/b/c.jar",
		"max_concurrent_runs":       1,
		"max_retries":               3,
		"min_retry_interval_millis": 5000,
		"name":                      "Featurizer New",
		"retry_on_timeout":          true,
	}, actionWithID("789", resourceJobUpdate))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/delete",
			ExpectedRequest: map[string]int{
				"job_id": 789,
			},
		},
	}, resourceJob, nil, actionWithID("789", resourceJobDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceJob, nil, actionWithID("789", resourceJobDelete))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id())
}
