package jobs

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/qa"
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
					Libraries: []libraries.Library{
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
						Libraries: []libraries.Library{
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

func TestResourceJobCreate_MultiTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "Featurizer",
					Tasks: []JobTaskSettings{
						{
							TaskKey:           "a",
							ExistingClusterID: "abc",
							Libraries: []libraries.Library{
								{
									Jar: "dbfs://aa/bb/cc.jar",
								},
							},
							SparkJarTask: &SparkJarTask{
								MainClassName: "com.labs.BarMain",
							},
						},
						{
							TaskKey: "b",
							NewCluster: &clusters.Cluster{
								SparkVersion: "a",
								NodeTypeID:   "b",
								NumWorkers:   1,
								AzureAttributes: &clusters.AzureAttributes{
									SpotBidMaxPrice: 0.99,
								},
							},
							NotebookTask: &NotebookTask{
								NotebookPath: "/Stuff",
							},
						},
					},
					MaxConcurrentRuns: 1,
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=789",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
							{
								TaskKey: "a",
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "Featurizer"
		
		task {
			task_key = "a"

			existing_cluster_id = "abc"

			spark_jar_task {
				main_class_name = "com.labs.BarMain"
			}

			library {
				jar = "dbfs://aa/bb/cc.jar"
			}
		}

		task {
			task_key = "b"

			new_cluster {
				spark_version = "a"
				node_type_id = "b"
				num_workers = 1
				azure_attributes {
					spot_bid_max_price = 0.99
				}
			}

			notebook_task {
				notebook_path = "/Stuff"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_JobClusters(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "JobClustered",
					Tasks: []JobTaskSettings{
						{
							TaskKey:       "a",
							JobClusterKey: "j",
						},
						{
							TaskKey: "b",
							NewCluster: &clusters.Cluster{
								SparkVersion: "a",
								NodeTypeID:   "b",
								NumWorkers:   3,
							},
							NotebookTask: &NotebookTask{
								NotebookPath: "/Stuff",
							},
						},
					},
					MaxConcurrentRuns: 1,
					JobClusters: []JobCluster{
						{
							JobClusterKey: "j",
							NewCluster: &clusters.Cluster{
								SparkVersion: "b",
								NodeTypeID:   "c",
								NumWorkers:   7,
							},
						},
						{
							JobClusterKey: "k",
							NewCluster: &clusters.Cluster{
								SparkVersion: "x",
								NodeTypeID:   "y",
								NumWorkers:   9,
							},
						},
					},
				},
				Response: Job{
					JobID: 17,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=17",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
							{
								TaskKey: "a",
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "JobClustered"

		job_cluster {
			job_cluster_key = "j"
			new_cluster {
			  num_workers   = 7
			  spark_version = "b"
			  node_type_id  = "c"
			}
		}
		
		job_cluster {
			job_cluster_key = "k"
			new_cluster {
			  num_workers   = 9
			  spark_version = "x"
			  node_type_id  = "y"
			}
		}
		
		task {
			task_key = "a"
			job_cluster_key = "j"
		}

		task {
			task_key = "b"

			new_cluster {
				spark_version = "a"
				node_type_id = "b"
				num_workers = 3
			}

			notebook_task {
				notebook_path = "/Stuff"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "17", d.Id())
}

func TestResourceJobCreate_AlwaysRunning(t *testing.T) {
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
					Name:              "Featurizer",
					MaxRetries:        3,
					MaxConcurrentRuns: 1,
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
						Name:       "Featurizer",
						MaxRetries: 3,
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/run-now",
				ExpectedRequest: RunParameters{
					JobID: 789,
				},
				Response: JobRun{
					RunID: 890,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/get?run_id=890",
				Response: JobRun{
					State: RunState{
						LifeCycleState: "RUNNING",
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_retries = 3
		name = "Featurizer"
		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}
		always_running = true
		`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_AlwaysRunning_Conflict(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		always_running = true
		max_concurrent_runs = 2
		`,
	}.ExpectError(t, "`always_running` must be specified only with `max_concurrent_runs = 1`")
}

func TestResourceJobCreateSingleNode(t *testing.T) {
	cluster := clusters.Cluster{
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
	cluster := clusters.Cluster{
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

func TestResourceJobCreateWithWebhooks(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				ExpectedRequest: JobSettings{
					ExistingClusterID: "abc",
					MaxConcurrentRuns: 1,
					SparkJarTask: &SparkJarTask{
						MainClassName: "com.labs.BarMain",
					},
					Libraries: []libraries.Library{
						{
							Jar: "dbfs://aa/bb/cc.jar",
						},
					},
					Name: "Featurizer",
					WebhookNotifications: &WebhookNotifications{
						OnStart:   []Webhook{{ID: "id1"}, {ID: "id2"}, {ID: "id3"}},
						OnSuccess: []Webhook{{ID: "id2"}},
						OnFailure: []Webhook{{ID: "id3"}},
					},
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
						MaxConcurrentRuns: 1,
						SparkJarTask: &SparkJarTask{
							MainClassName: "com.labs.BarMain",
						},
						Libraries: []libraries.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
						},
						Name: "Featurizer",
						WebhookNotifications: &WebhookNotifications{
							OnStart:   []Webhook{{ID: "id1"}, {ID: "id2"}, {ID: "id3"}},
							OnSuccess: []Webhook{{ID: "id2"}},
							OnFailure: []Webhook{{ID: "id3"}},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		name = "Featurizer"
		max_concurrent_runs = 1
		spark_jar_task {
			main_class_name = "com.labs.BarMain"
		}
		library {
			jar = "dbfs://aa/bb/cc.jar"
		}
		webhook_notifications {
			on_start {
				id = "id3" 
			}
			on_start {
				id = "id1" 
			}
			on_start {
				id = "id2" 
			}
			on_success {
				id = "id2" 
			}
			on_failure {
				id = "id3" 
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreateFromGitSource(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					ExistingClusterID: "abc",
					Tasks: []JobTaskSettings{
						{
							TaskKey: "b",
							NotebookTask: &NotebookTask{
								NotebookPath: "/GitSourcedNotebook",
							},
						},
					},
					Name:              "GitSourceJob",
					MaxConcurrentRuns: 1,
					GitSource: &GitSource{
						Url:      "https://github.com/databricks/terraform-provider-databricks",
						Tag:      "0.4.8",
						Provider: "gitHub",
					},
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						// good enough for mock
						ExistingClusterID: "abc",
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
						},
						Name:              "GitSourceJob",
						MaxConcurrentRuns: 1,
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		name = "GitSourceJob"

		git_source {
			url = "https://github.com/databricks/terraform-provider-databricks"
			tag = "0.4.8"
		}

		task {
			task_key = "b"

			notebook_task {
				notebook_path = "/GitSourcedNotebook"
			}
		}`,
	}.ApplyNoError(t)
}

func resourceJobCreateFromGitSourceConflict(t *testing.T, conflictingArgs []string, gitSource string) {
	var hclTemplate = `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		name = "GitSourceJob"
		
		%s
		
		task {
			task_key = "b"

			notebook_task {
				notebook_path = "/GitSourcedNotebook"
			}
		}
	`
	var hcl = fmt.Sprintf(hclTemplate, gitSource)
	_, err := qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL:      hcl,
	}.Apply(t)
	assert.Error(t, err, err)
	var found = false
	for _, fieldName := range conflictingArgs {
		require.Equal(t, true, strings.Contains(err.Error(), fieldName))
		found = true
	}
	require.Equal(t, true, found)
}

func TestResourceJobCreateFromGitSourceTagAndBranchConflict(t *testing.T) {
	var gitSource = `git_source {
		url = "https://github.com/databricks/terraform-provider-databricks"
		tag = "0.4.8"
		branch = "main"
	}`
	resourceJobCreateFromGitSourceConflict(t, []string{"branch", "tag"}, gitSource)
}
func TestResourceJobCreateFromGitSourceTagAndCommitConflict(t *testing.T) {
	var gitSource = `git_source {
		url = "https://github.com/databricks/terraform-provider-databricks"
		tag = "0.4.8"
		commit = "a26bf6"
	}`
	resourceJobCreateFromGitSourceConflict(t, []string{"commit", "tag"}, gitSource)
}

func TestResourceJobCreateFromGitSourceBranchAndCommitConflict(t *testing.T) {
	var gitSource = `git_source {
		url = "https://github.com/databricks/terraform-provider-databricks"
		branch = "main"
		commit = "a26bf6"
	}`
	resourceJobCreateFromGitSourceConflict(t, []string{"branch", "commit"}, gitSource)
}

func TestResourceJobCreateFromGitSourceWithoutProviderFail(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `existing_cluster_id = "abc"
		max_concurrent_runs = 1
		name = "GitSourceJob"

		git_source {
			url = "https://custom.git.hosting.com/databricks/terraform-provider-databricks"
			tag = "0.4.8"
		}

		task {
			task_key = "b"

			notebook_task {
				notebook_path = "/GitSourcedNotebook"
			}
		}
	`,
	}.ExpectError(t, "git source is not empty but Git Provider is not specified and cannot be guessed by url &{Url:https://custom.git.hosting.com/databricks/terraform-provider-databricks Provider: Branch: Tag:0.4.8 Commit:}")
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
						Libraries: []libraries.Library{
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
						Libraries: []libraries.Library{
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
						Libraries: []libraries.Library{
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

func TestResourceJobUpdate_Tasks(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Name: "Featurizer New",
						Tasks: []JobTaskSettings{
							{
								ExistingClusterID: "abc",
								SparkJarTask: &SparkJarTask{
									MainClassName: "com.labs.BarMain",
								},
							},
						},
						MaxConcurrentRuns: 1,
					},
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=789",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								ExistingClusterID: "abc",
								SparkJarTask: &SparkJarTask{
									MainClassName: "com.labs.BarMain",
								},
							},
						},
					},
				},
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "Featurizer New"
		task {
			existing_cluster_id = "abc"
			spark_jar_task {
				main_class_name = "com.labs.BarMain"
			}
		}`,
	}.ApplyNoError(t)
}

func TestResourceJobUpdate_Restart(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Name:              "Featurizer New",
						MaxConcurrentRuns: 1,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						Name: "Featurizer New",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=789",
				Response: JobRunsList{},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/run-now",
				ExpectedRequest: RunParameters{
					JobID: 789,
				},
				Response: JobRun{
					RunID: 890,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/get?run_id=890",
				Response: JobRun{
					State: RunState{
						LifeCycleState: "RUNNING",
					},
				},
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "Featurizer New"
		always_running = true
		`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "789", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Featurizer New", d.Get("name"))
}

func TestJobRestarts(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "POST",
			Resource:     "/api/2.0/jobs/run-now",
			ReuseRequest: true,
			ExpectedRequest: RunParameters{
				JobID: 123,
			},
			Response: JobRun{
				RunID: 234,
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/jobs/runs/get?run_id=234",
			ReuseRequest: true,
			Response: JobRun{
				State: RunState{
					LifeCycleState: "RUNNING",
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/get?run_id=345",
			Status:   400,
			Response: common.APIError{
				Message: "nope",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/get?run_id=456",
			Response: JobRun{
				State: RunState{
					LifeCycleState: "INTERNAL_ERROR",
					StateMessage:   "Quota exceeded",
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/jobs/runs/get?run_id=890",
			ReuseRequest: true,
			Response: JobRun{
				State: RunState{
					LifeCycleState: "SOMETHING",
					StateMessage:   "Checking...",
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=123",
			Response: JobRunsList{
				Runs: []JobRun{},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=123",
			Response: JobRunsList{
				Runs: []JobRun{
					{
						RunID: 567,
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/runs/cancel",
			ExpectedRequest: map[string]any{
				"run_id": 567,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=111",
			Response: JobRunsList{
				Runs: []JobRun{
					{
						RunID: 567,
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/jobs/runs/cancel",
			Status:   400,
			Response: common.APIError{
				Message: "nope",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=222",
			Status:   400,
			Response: common.APIError{
				Message: "nope",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/jobs/runs/get?run_id=567",
			ReuseRequest: true,
			Response: JobRun{
				State: RunState{
					LifeCycleState: "TERMINATED",
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=678",
			Response: JobRunsList{
				Runs: []JobRun{
					{
						RunID: 789,
					},
					{
						RunID: 890,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ja := NewJobsAPI(ctx, client)
		timeout := 500 * time.Millisecond

		err := ja.Start(123, timeout)
		assert.NoError(t, err)

		err = ja.waitForRunState(345, "RUNNING", timeout)
		assert.EqualError(t, err, "cannot get job RUNNING: nope")

		err = ja.waitForRunState(456, "TERMINATED", timeout)
		assert.EqualError(t, err, "cannot get job TERMINATED: Quota exceeded")

		err = ja.waitForRunState(890, "RUNNING", timeout)
		assert.EqualError(t, err, "run is SOMETHING: Checking...")

		// no active runs for the first time
		err = ja.Restart("123", timeout)
		assert.NoError(t, err)

		// one active run for the second time
		err = ja.Restart("123", timeout)
		assert.NoError(t, err)

		err = ja.Restart("111", timeout)
		assert.EqualError(t, err, "cannot cancel run 567: nope")

		err = ja.Restart("a", timeout)
		assert.EqualError(t, err, "strconv.ParseInt: parsing \"a\": invalid syntax")

		err = ja.Restart("222", timeout)
		assert.EqualError(t, err, "nope")

		err = ja.Restart("678", timeout)
		assert.EqualError(t, err, "`always_running` must be specified only "+
			"with `max_concurrent_runs = 1`. There are 2 active runs")
	})
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

func TestJobsAPIList(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 1,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewJobsAPI(ctx, client)
		l, err := a.List()
		require.NoError(t, err)
		assert.Len(t, l, 1)
	})
}

func TestJobsAPIListMultiplePages(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 1,
					},
				},
				HasMore: true,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=1",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 2,
					},
				},
				HasMore: false,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewJobsAPI(ctx, client)
		l, err := a.List()
		require.NoError(t, err)
		assert.Len(t, l, 2)
	})
}

func TestJobsAPIListByName(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&name=test&offset=0",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 1,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewJobsAPI(ctx, client)
		l, err := a.ListByName("test", false)
		require.NoError(t, err)
		assert.Len(t, l, 1)
	})
}

func TestJobsAPIRunsList(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
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
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewJobsAPI(ctx, client)
		l, err := a.RunsList(JobRunsListRequest{
			JobID:         234,
			CompletedOnly: true,
			Limit:         1,
			Offset:        0,
		})
		require.NoError(t, err)
		assert.Len(t, l.Runs, 1)
	})
}

func TestJobResourceCornerCases_HTTP(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceJob(), qa.CornerCaseID("10"))
}

func TestJobResourceCornerCases_WrongID(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceJob(),
		qa.CornerCaseID("x"),
		qa.CornerCaseSkipCRUD("create"),
		qa.CornerCaseExpectError(`strconv.ParseInt: parsing "x": invalid syntax`))
}

func TestJobResource_SparkConfDiffSuppress(t *testing.T) {
	jr := ResourceJob()
	scs := common.MustSchemaPath(jr.Schema, "new_cluster", "spark_conf")
	assert.True(t, scs.DiffSuppressFunc("new_cluster.0.spark_conf.%", "1", "0", nil))
	assert.False(t, scs.DiffSuppressFunc("new_cluster.0.spark_conf.%", "1", "1", nil))
}
