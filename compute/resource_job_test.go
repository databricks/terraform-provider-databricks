package compute

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestAwsAccJobsCreate(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	jobsAPI := NewJobsAPI(context.Background(), client)

	jobSettings := JobSettings{
		NewCluster: &Cluster{
			NumWorkers:   2,
			SparkVersion: "6.4.x-scala2.11",
			SparkConf:    nil,
			AwsAttributes: &AwsAttributes{
				Availability: "ON_DEMAND",
			},
			NodeTypeID: "r3.xlarge",
		},
		NotebookTask: &NotebookTask{
			NotebookPath: "/Users/sri.tikkireddy@databricks.com/demo-terraform/demo-notebook",
		},
		Name: "1-sri-test-job",
		Libraries: []Library{
			{
				Maven: &Maven{
					Coordinates: "org.jsoup:jsoup:1.7.2",
				},
			},
		},
		EmailNotifications: &JobEmailNotifications{
			OnStart:   []string{},
			OnSuccess: []string{},
			OnFailure: []string{},
		},
		TimeoutSeconds: 3600,
		MaxRetries:     1,
		Schedule: &CronSchedule{
			QuartzCronExpression: "0 15 22 ? * *",
			TimezoneID:           "America/Los_Angeles",
		},
		MaxConcurrentRuns: 1,
	}

	job, err := jobsAPI.Create(jobSettings)
	assert.NoError(t, err, err)
	id := job.ID()
	defer func() {
		err := jobsAPI.Delete(id)
		assert.NoError(t, err, err)
	}()
	t.Log(id)
	job, err = jobsAPI.Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == "6.4.x-scala2.11",
		"Something is wrong with spark version")

	jobSettings.NewCluster.SparkVersion = "6.1.x-scala2.11"

	err = jobsAPI.Update(id, jobSettings)
	assert.NoError(t, err, err)

	job, err = jobsAPI.Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == "6.1.x-scala2.11",
		"Something is wrong with spark version")
}

func TestResourceJobCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				ExpectedRequest: JobSettings{
					ExistingClusterID: "abc",
					SparkJarTask: &SparkJarTask{
						MainClassName: "com.labs.BarMain",
					},
					Libraries: []Library{
						{
							Jar: "dbfs://ff/gg/hh.jar",
						},
						{
							Jar: "dbfs://aa/bb/cc.jar",
						},
					},
					Name:                   "Featurizer",
					MaxRetries:             3,
					MinRetryIntervalMillis: 5000,
					RetryOnTimeout:         true,
					MaxConcurrentRuns:      1,
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						ExistingClusterID: "abc",
						SparkJarTask: &SparkJarTask{
							MainClassName: "com.labs.BarMain",
						},
						Libraries: []Library{
							{
								Jar: "dbfs://ff/gg/hh.jar",
							},
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
						},
						Name:                   "Featurizer",
						MaxRetries:             3,
						MinRetryIntervalMillis: 5000,
						RetryOnTimeout:         true,
						MaxConcurrentRuns:      1,
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}
		library {
			jar = "dbfs://aa/bb/cc.jar"
		}
		library {
			jar = "dbfs://ff/gg/hh.jar"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}
		library {
			jar = "dbfs://aa/bb/cc.jar"
		}
		library {
			jar = "dbfs://ff/gg/hh.jar"
		}`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceJobRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						ExistingClusterID: "abc",
						SparkJarTask: &SparkJarTask{
							MainClassName: "com.labs.BarMain",
							Parameters:    []string{"--cleanup", "full"},
						},
						Libraries: []Library{
							{
								Jar: "dbfs://ff/gg/hh.jar",
							},
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
						},
						Name:                   "Featurizer",
						MaxRetries:             3,
						MinRetryIntervalMillis: 5000,
						RetryOnTimeout:         true,
						MaxConcurrentRuns:      1,
					},
				},
			},
		},
		Resource: ResourceJob(),
		Read:     true,
		New:      true,
		ID:       "789",
	}.Apply(t)
	assert.NoError(t, err, err)

	assert.Equal(t, "Featurizer", d.Get("name"))
	assert.Equal(t, 2, d.Get("library.#"))
	assert.Equal(t, "dbfs://ff/gg/hh.jar", d.Get("library.1850263921.jar"))
	assert.Equal(t, "dbfs://aa/bb/cc.jar", d.Get("library.587400796.jar"))

	assert.Equal(t, 2, d.Get("spark_jar_task.0.parameters.#"))
	assert.Equal(t, "com.labs.BarMain", d.Get("spark_jar_task.0.main_class_name"))
	assert.Equal(t, "--cleanup", d.Get("spark_jar_task.0.parameters.0"))
	assert.Equal(t, "full", d.Get("spark_jar_task.0.parameters.1"))

	assert.Equal(t, 5000, d.Get("min_retry_interval_millis"))
	assert.Equal(t, 3, d.Get("max_retries"))
	assert.Equal(t, 1, d.Get("max_concurrent_runs"))
	assert.Equal(t, 1, d.Get("spark_jar_task.#"))
	assert.Equal(t, true, d.Get("retry_on_timeout"))
	assert.Equal(t, "abc", d.Get("existing_cluster_id"))
}

func TestResourceJobRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceJob(),
		Read:     true,
		New:      true,
		Removed:  true,
		ID:       "789",
	}.ApplyNoError(t)
}

func TestResourceJobRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceJob(),
		Read:     true,
		New:      true,
		ID:       "789",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id(), "Id should not be empty for error reads")
}

func TestResourceJobUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						ExistingClusterID: "abc",
						SparkJarTask: &SparkJarTask{
							MainClassName: "com.labs.BarMain",
							Parameters:    []string{"--cleanup", "full"},
						},
						Libraries: []Library{
							{
								Jar: "dbfs://ff/gg/hh.jar",
							},
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
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
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						ExistingClusterID: "abc",
						SparkJarTask: &SparkJarTask{
							MainClassName: "com.labs.BarMain",
							Parameters:    []string{"--cleanup", "full"},
						},
						Libraries: []Library{
							{
								Jar: "dbfs://ff/gg/hh.jar",
							},
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
						},
						Name:                   "Featurizer New",
						MaxRetries:             3,
						MinRetryIntervalMillis: 5000,
						RetryOnTimeout:         true,
						MaxConcurrentRuns:      1,
					},
				},
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer New"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
			parameters = ["--cleanup", "full"]
		}
		library {
			jar = "dbfs://aa/bb/cc.jar"
		}
		library {
			jar = "dbfs://ff/gg/hh.jar"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Featurizer New", d.Get("name"))
}

func TestResourceJobUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/reset",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer New"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
			parameters = ["--cleanup", "full"]
		}
		library {
			jar = "dbfs://aa/bb/cc.jar"
		}
		library {
			jar = "dbfs://ff/gg/hh.jar"
		}`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/delete",
				ExpectedRequest: map[string]int{
					"job_id": 789,
				},
			},
		},
		ID:       "789",
		Delete:   true,
		Resource: ResourceJob(),
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		ID:       "789",
		Delete:   true,
		Resource: ResourceJob(),
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "789", d.Id())
}
