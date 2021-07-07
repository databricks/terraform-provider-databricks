package compute

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
					Schedule: &CronSchedule{
						QuartzCronExpression: "0 15 22 ? * *",
						TimezoneID:           "America/Los_Angeles",
						PauseStatus:          "PAUSED",
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
						Schedule: &CronSchedule{
							QuartzCronExpression: "0 15 22 ? * *",
							TimezoneID:           "America/Los_Angeles",
							PauseStatus:          "PAUSED",
						},
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
		schedule {
			quartz_cron_expression = "0 15 22 ? * *"
			timezone_id = "America/Los_Angeles"
			pause_status = "PAUSED"
		}
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

func TestResourceJobCreateSingleNode(t *testing.T) {
	cluster := Cluster{
		NumWorkers: 0, SparkVersion: "7.3.x-scala2.12", NodeTypeID: "Standard_DS3_v2",
		SparkConf: map[string]string{
			"spark.master":                     "local[*]",
			"spark.databricks.cluster.profile": "singleNode",
		},
		CustomTags: map[string]string{
			"ResourceClass": "SingleNode",
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				ExpectedRequest: JobSettings{
					NewCluster: &cluster,
					SparkJarTask: &SparkJarTask{
						MainClassName: "com.labs.BarMain",
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
		HCL: `new_cluster  {
			num_workers   = 0
			spark_version = "7.3.x-scala2.12"
			node_type_id  = "Standard_DS3_v2"
			spark_conf {
				"spark.master" = "local[*]"
				"spark.databricks.cluster.profile" = "singleNode"
			}
            custom_tags {
                "ResourceClass" = "SingleNode"
            }
		  }	
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreateNWorkers(t *testing.T) {
	cluster := Cluster{
		NumWorkers: 5, SparkVersion: "7.3.x-scala2.12", NodeTypeID: "Standard_DS3_v2",
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				ExpectedRequest: JobSettings{
					NewCluster: &cluster,
					SparkJarTask: &SparkJarTask{
						MainClassName: "com.labs.BarMain",
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
		HCL: `new_cluster  {
			num_workers   = 5
			spark_version = "7.3.x-scala2.12"
			node_type_id  = "Standard_DS3_v2"
		  }	
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreateSingleNode_Fail(t *testing.T) {
	_, err := qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `new_cluster  {
			num_workers   = 0
			spark_version = "7.3.x-scala2.12"
			node_type_id  = "Standard_DS3_v2"
		  }	
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
	assert.Error(t, err, err)
	require.Equal(t, true, strings.Contains(err.Error(), "NumWorkers could be 0 only for SingleNode clusters"))
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

func TestAccRestartJob(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("This test is supposed to be run only from IDE")
	}
	// TODO: remove this test after functionality is ready
	client := common.CommonEnvironmentClient()
	jobsAPI := NewJobsAPI(context.Background(), client)
	err := jobsAPI.Restart("210")
	assert.NoError(t, err)
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

func TestResourceJobUpdate_FailNumWorkersZero(t *testing.T) {
	_, err := qa.ResourceFixture{
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		HCL: `new_cluster  {
			num_workers   = 0
			spark_version = "7.3.x-scala2.12"
			node_type_id  = "Standard_DS3_v2"
		  }
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer New"
		retry_on_timeout = true

		spark_jar_task {
			main_class_name = "com.labs.BarMain"
			parameters = ["--cleanup", "full"]
		}`,
	}.Apply(t)
	assert.Error(t, err, err)
	require.Equal(t, true, strings.Contains(err.Error(), "NumWorkers could be 0 only for SingleNode clusters"))
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

func TestJobsAPIRunsList(t *testing.T) {
	c, s, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=234&limit=1",
			Response: JobRunsList{
				Runs: []JobRun{
					{
						State: RunState{
							ResultState:    "SUCCESS",
							LifeCycleState: "TERMINATED",
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	defer s.Close()

	a := NewJobsAPI(context.Background(), c)
	l, err := a.RunsList(JobRunsListRequest{
		JobID:         234,
		CompletedOnly: true,
		Limit:         1,
		Offset:        0,
	})
	require.NoError(t, err)
	assert.Len(t, l.Runs, 1)
}
