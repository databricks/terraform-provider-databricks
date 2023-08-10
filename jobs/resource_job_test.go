package jobs

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
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
					Queue:                  &Queue{},
					RunAs: &JobRunAs{
						UserName: "user@mail.com",
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
					JobID:         789,
					RunAsUserName: "user@mail.com",
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
						Queue: &Queue{},
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
		}
		queue {}
		run_as {
			user_name = "user@mail.com"
		}`,
	}.Apply(t)
	assert.NoError(t, err)
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
							Health: &JobHealth{
								Rules: []JobHealthRule{
									{
										Metric:    "RUN_DURATION_SECONDS",
										Operation: "GREATER_THAN",
										Value:     3600,
									},
								},
							},
						},
						{
							TaskKey: "b",
							DependsOn: []jobs.TaskDependency{
								{
									TaskKey: "a",
								},
							},
							RunIf: "ALL_DONE",
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
					Health: &JobHealth{
						Rules: []JobHealthRule{
							{
								Metric:    "RUN_DURATION_SECONDS",
								Operation: "GREATER_THAN",
								Value:     3600,
							},
						},
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

		health {
			rules {
				metric = "RUN_DURATION_SECONDS"
				op     = "GREATER_THAN"
				value  = 3600						  
			}
		}

		task {
			task_key = "a"

			existing_cluster_id = "abc"

			spark_jar_task {
				main_class_name = "com.labs.BarMain"
			}

			library {
				jar = "dbfs://aa/bb/cc.jar"
			}

			health {
				rules {
					metric = "RUN_DURATION_SECONDS"
					op     = "GREATER_THAN"
					value  = 3600						  
				}
			}
	
		}

		task {
			task_key = "b"

			depends_on {
				task_key = "a"
			}

			run_if = "ALL_DONE"

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
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_JobParameters(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "JobParameterTesting",
					Tasks: []JobTaskSettings{
						{
							TaskKey: "a",
						},
						{
							TaskKey: "b",
						},
					},
					MaxConcurrentRuns: 1,
					Parameters: []JobParameterDefinition{
						{
							Name:    "hello",
							Default: "world",
						},
						{
							Name:    "key",
							Default: "value_default",
						},
					},
				},
				Response: Job{
					JobID: 231,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=231",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "a",
							},
							{
								TaskKey: "b",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "hello",
								Default: "world",
							},
							{
								Name:    "key",
								Default: "value_default",
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "JobParameterTesting"

		parameter {
				name = "hello"
				default = "world"
		}
		parameter {
				name = "key"
				default = "value_default"
		}
	
		task {
			task_key = "a"
		}

		task {
			task_key = "b"
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "231", d.Id())
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
	assert.NoError(t, err)
	assert.Equal(t, "17", d.Id())
}

func TestResourceJobCreate_JobCompute(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "JobComputed",
					Tasks: []JobTaskSettings{
						{
							TaskKey:    "b",
							ComputeKey: "j",
							NotebookTask: &NotebookTask{
								NotebookPath: "/Stuff",
							},
						},
					},
					MaxConcurrentRuns: 1,
					Compute: []JobCompute{
						{
							ComputeKey: "j",
							ComputeSpec: &compute.ComputeSpec{
								Kind: "t",
							},
						},
					},
				},
				Response: Job{
					JobID: 18,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=18",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "JobComputed"

		compute {
			compute_key = "j"
			spec {
			  kind   	= "t"
			}
		}

		task {
			task_key = "b"

			compute_key = "j"

			notebook_task {
				notebook_path = "/Stuff"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "18", d.Id())
}

func TestResourceJobCreate_SqlSubscriptions(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name:              "TF SQL task subscriptions",
					MaxConcurrentRuns: 1,
					Tasks: []JobTaskSettings{
						{
							TaskKey: "a",
							SqlTask: &SqlTask{
								WarehouseID: "dca3a0ba199040eb",
								Alert: &SqlAlertTask{
									AlertID: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
									Subscriptions: []SqlSubscription{
										{UserName: "user@domain.com"},
										{DestinationID: "Test"},
									},
									PauseSubscriptions: true,
								},
							},
						},
						{
							TaskKey: "d",
							SqlTask: &SqlTask{
								WarehouseID: "dca3a0ba199040eb",
								Dashboard: &SqlDashboardTask{
									DashboardID: "d81a7760-7fd2-443e-bf41-95a60c2f4c7c",
									Subscriptions: []SqlSubscription{
										{UserName: "user@domain.com"},
										{DestinationID: "Test"},
									},
									CustomSubject: "test",
								},
							},
						},
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
						Name: "TF SQL task subscriptions",
						Tasks: []JobTaskSettings{
							{
								TaskKey: "a",
								SqlTask: &SqlTask{
									WarehouseID: "dca3a0ba199040eb",
									Alert: &SqlAlertTask{
										AlertID: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
										Subscriptions: []SqlSubscription{
											{UserName: "user@domain.com"},
											{DestinationID: "Test"},
										},
										PauseSubscriptions: true,
									},
								},
							},
							{
								TaskKey: "d",
								SqlTask: &SqlTask{
									WarehouseID: "dca3a0ba199040eb",
									Dashboard: &SqlDashboardTask{
										DashboardID: "d81a7760-7fd2-443e-bf41-95a60c2f4c7c",
										Subscriptions: []SqlSubscription{
											{UserName: "user@domain.com"},
											{DestinationID: "Test"},
										},
										CustomSubject: "test",
									},
								},
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `name = "TF SQL task subscriptions"

		task {
		  task_key = "a"
	  
		  sql_task {
			warehouse_id = "dca3a0ba199040eb"
			alert {
			  subscriptions {
				user_name = "user@domain.com"
			  }
			  subscriptions {
				destination_id = "Test"
			  }
			  pause_subscriptions = true
			  alert_id = "3cf91a42-6217-4f3c-a6f0-345d489051b9"
			}
		  }
		}
	  
		task {
		  task_key = "d"
	  
		  sql_task {
			warehouse_id = "dca3a0ba199040eb"
			dashboard {
			  subscriptions {
				user_name = "user@domain.com"
			  }
			  subscriptions {
				destination_id = "Test"
			  }
			  pause_subscriptions = false
			  dashboard_id = "d81a7760-7fd2-443e-bf41-95a60c2f4c7c"
			  custom_subject = "test"
			}
		  }
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_RunJobTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/create",
				ExpectedRequest: JobSettings{
					Name:              "TF RunJobTask Main Job",
					MaxConcurrentRuns: 1,
					Tasks: []JobTaskSettings{
						{
							TaskKey: "runJobTask",
							RunJobTask: &RunJobTask{
								JobID: "123",
							},
						},
					},
				},
				Response: Job{
					JobID: 123,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=123",
				Response: Job{
					JobID: 123,
					Settings: &JobSettings{
						Name: "TF SQL task subscriptions",
						Tasks: []JobTaskSettings{
							{
								TaskKey: "childJobTaskKey",
								NotebookTask: &NotebookTask{
									NotebookPath: "/Stuff",
								},
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `name = "TF RunJobTask Main Job"

		task {
		  task_key = "runJobTask"
	  
		  run_job_task {
				job_id = "123"
		  }
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "123", d.Id())
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
	assert.NoError(t, err)
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

func TestResourceJobCreate_ControlRunState_AlwaysRunningConflict(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `control_run_state = true
		always_running = true
		continuous {
			pause_status = "UNPAUSED"
		}`,
	}.ExpectError(t, "invalid config supplied. [always_running] Conflicting configuration arguments. [control_run_state] Conflicting configuration arguments")
}

func TestResourceJobCreate_ControlRunState_NoContinuous(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL:      `control_run_state = true`,
	}.ExpectError(t, "`control_run_state` must be specified only with `continuous`")
}

func TestResourceJobCreate_ControlRunState_ContinuousCreate(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/create",
				ExpectedRequest: JobSettings{
					MaxConcurrentRuns: 1,
					Name:              "Test",
					Continuous: &ContinuousConf{
						PauseStatus: "UNPAUSED",
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
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Continuous: &ContinuousConf{
							PauseStatus: "UNPAUSED",
						},
					},
				},
			},
		},
		HCL: `
		continuous {
			pause_status = "UNPAUSED"
		}
		control_run_state = true
		max_concurrent_runs = 1
		name = "Test"
		`,
	}.Apply(t)
}

func TestResourceJobCreate_ControlRunState_ContinuousUpdateRunNow(t *testing.T) {
	qa.ResourceFixture{
		Update:   true,
		ID:       "789",
		Resource: ResourceJob(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Continuous: &ContinuousConf{
							PauseStatus: "UNPAUSED",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Continuous: &ContinuousConf{
							PauseStatus: "UNPAUSED",
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/run-now",
				ExpectedRequest: RunParameters{
					JobID: 789,
				},
				Response: JobRun{},
			},
		},
		HCL: `
		continuous {
			pause_status = "UNPAUSED"
		}
		control_run_state = true
		max_concurrent_runs = 1
		name = "Test"
		`,
	}.Apply(t)
}

func TestResourceJobCreate_ControlRunState_ContinuousUpdateCancel(t *testing.T) {
	qa.ResourceFixture{
		Update:   true,
		ID:       "789",
		Resource: ResourceJob(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Continuous: &ContinuousConf{
							PauseStatus: "UNPAUSED",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Continuous: &ContinuousConf{
							PauseStatus: "UNPAUSED",
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/jobs/run-now",
				ExpectedRequest: RunParameters{
					JobID: 789,
				},
				Response: apierr.APIError{StatusCode: 404},
				Status:   404,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=789",
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
				Resource: "/api/2.0/jobs/runs/get?run_id=567",
				Response: JobRun{
					RunID: 567,
					State: RunState{
						LifeCycleState: "TERMINATED",
					},
				},
			},
		},
		HCL: `
		continuous {
			pause_status = "UNPAUSED"
		}
		control_run_state = true
		max_concurrent_runs = 1
		name = "Test"
		`,
	}.Apply(t)
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
	assert.NoError(t, err)
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
	assert.NoError(t, err)
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
					NotificationSettings: &jobs.JobNotificationSettings{
						NoAlertForSkippedRuns:  true,
						NoAlertForCanceledRuns: true,
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
						NotificationSettings: &jobs.JobNotificationSettings{
							NoAlertForSkippedRuns:  true,
							NoAlertForCanceledRuns: true,
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
		}
		notification_settings {
			no_alert_for_skipped_runs = true
			no_alert_for_canceled_runs = true
		  }
	`,
	}.Apply(t)
	assert.NoError(t, err)
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
						JobSource: &jobs.JobSource{
							JobConfigPath:       "a/b/c/databricks.yml",
							ImportFromGitBranch: "main",
							DirtyState:          "NOT_SYNCED",
						},
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
			job_source {
				job_config_path = "a/b/c/databricks.yml"
				import_from_git_branch = "main"
				dirty_state = "NOT_SYNCED"
			}
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
	assert.Error(t, err)
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
	}.ExpectError(t, "git source is not empty but Git Provider is not specified and cannot be guessed by url &{Url:https://custom.git.hosting.com/databricks/terraform-provider-databricks Provider: Branch: Tag:0.4.8 Commit: JobSource:<nil>}")
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
	assert.Error(t, err)
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
	assert.NoError(t, err)

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
				Response: apierr.APIErrorBody{
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
				Response: apierr.APIErrorBody{
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
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Featurizer New", d.Get("name"))
}

func TestResourceJobUpdate_NodeTypeToInstancePool(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						NewCluster: &clusters.Cluster{
							InstancePoolID:       "instance-pool-worker",
							DriverInstancePoolID: "instance-pool-driver",
							SparkVersion:         "spark-1",
							NumWorkers:           1,
						},
						Tasks: []JobTaskSettings{
							{
								NewCluster: &clusters.Cluster{
									InstancePoolID:       "instance-pool-worker-task",
									DriverInstancePoolID: "instance-pool-driver-task",
									SparkVersion:         "spark-2",
									NumWorkers:           2,
								},
							},
						},
						JobClusters: []JobCluster{
							{
								NewCluster: &clusters.Cluster{
									InstancePoolID:       "instance-pool-worker-job",
									DriverInstancePoolID: "instance-pool-driver-job",
									SparkVersion:         "spark-3",
									NumWorkers:           3,
								},
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
				Resource: "/api/2.1/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						NewCluster: &clusters.Cluster{
							NodeTypeID:       "node-type-id",
							DriverNodeTypeID: "driver-node-type-id",
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
		InstanceState: map[string]string{
			"new_cluster.0.node_type_id":                      "node-type-id-worker",
			"new_cluster.0.driver_node_type_id":               "node-type-id-driver",
			"task.0.new_cluster.0.node_type_id":               "node-type-id-worker-task",
			"task.0.new_cluster.0.driver_node_type_id":        "node-type-id-driver-task",
			"job_cluster.0.new_cluster.0.node_type_id":        "node-type-id-worker-job",
			"job_cluster.0.new_cluster.0.driver_node_type_id": "node-type-id-driver-job",
		},
		HCL: `
		new_cluster = {
			instance_pool_id = "instance-pool-worker"
			driver_instance_pool_id = "instance-pool-driver"
			spark_version = "spark-1"
			num_workers = 1
		}
		task = {
			new_cluster = {
				instance_pool_id = "instance-pool-worker-task"
				driver_instance_pool_id = "instance-pool-driver-task"
				spark_version = "spark-2"
				num_workers = 2
			}
		}
		job_cluster = {
			new_cluster = {
				instance_pool_id = "instance-pool-worker-job"
				driver_instance_pool_id = "instance-pool-driver-job"
				spark_version = "spark-3"
				num_workers = 3
			}
		}
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer New"
		retry_on_timeout = true`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Featurizer New", d.Get("name"))
}

func TestResourceJobUpdate_InstancePoolToNodeType(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						NewCluster: &clusters.Cluster{
							NodeTypeID:   "node-type-id-1",
							SparkVersion: "spark-1",
							NumWorkers:   1,
						},
						Tasks: []JobTaskSettings{
							{
								NewCluster: &clusters.Cluster{
									NodeTypeID:   "node-type-id-2",
									SparkVersion: "spark-2",
									NumWorkers:   2,
								},
							},
						},
						JobClusters: []JobCluster{
							{
								NewCluster: &clusters.Cluster{
									NodeTypeID:   "node-type-id-3",
									SparkVersion: "spark-3",
									NumWorkers:   3,
								},
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
				Resource: "/api/2.1/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						NewCluster: &clusters.Cluster{
							NodeTypeID:           "node-type-id",
							DriverNodeTypeID:     "driver-node-type-id",
							InstancePoolID:       "instance-pool-id-worker",
							DriverInstancePoolID: "instance-pool-id-driver",
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
		InstanceState: map[string]string{
			"new_cluster.0.instance_pool_id":           "instance-pool-id-worker",
			"new_cluster.0.driver_instance_pool_id":    "instance-pool-id-driver",
			"new_cluster.0.node_type_id":               "node-type-id-worker",
			"task.0.new_cluster.0.node_type_id":        "node-type-id-worker-task",
			"task.0.instance_pool_id":                  "instance-pool-id-worker",
			"task.0.driver_instance_pool_id":           "instance-pool-id-driver",
			"job_cluster.0.new_cluster.0.node_type_id": "node-type-id-worker-job",
			"job_cluster.0.instance_pool_id":           "instance-pool-id-worker",
			"job_cluster.0.driver_instance_pool_id":    "instance-pool-id-driver",
		},
		HCL: `
		new_cluster = {
			node_type_id = "node-type-id-1"
			spark_version = "spark-1"
			num_workers = 1
		}
		task = {
			new_cluster = {
				node_type_id = "node-type-id-2"
				spark_version = "spark-2"
				num_workers = 2
			}
		}
		job_cluster = {
			new_cluster = {
				node_type_id = "node-type-id-3"
				spark_version = "spark-3"
				num_workers = 3
			}
		}
		max_concurrent_runs = 1
		max_retries = 3
		min_retry_interval_millis = 5000
		name = "Featurizer New"
		retry_on_timeout = true`,
	}.Apply(t)
	assert.NoError(t, err)
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
	assert.NoError(t, err)
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
			Response: apierr.APIError{
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
			Response: apierr.APIError{
				Message: "nope",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?active_only=true&job_id=222",
			Status:   400,
			Response: apierr.APIError{
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

		testRestart := func(jobID int64, stopErr, startErr string) {
			err := ja.StopActiveRun(jobID, timeout)
			if stopErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, stopErr)
			}
			if stopErr == "" {
				err = ja.Start(jobID, timeout)
				if startErr == "" {
					assert.NoError(t, err)
				} else {
					assert.EqualError(t, err, startErr)
				}
			}
		}

		// no active runs for the first time
		testRestart(123, "", "")

		// one active run for the second time
		testRestart(123, "", "")

		testRestart(111, "cannot cancel run 567: nope", "")

		testRestart(222, "nope", "")

		testRestart(678, "`always_running` must be specified only "+
			"with `max_concurrent_runs = 1`. There are 2 active runs", "")
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
	assert.NoError(t, err)
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
	assert.Error(t, err)
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
