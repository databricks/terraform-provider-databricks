package jobs

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
					Libraries: []compute.Library{
						{
							Jar: "dbfs://aa/bb/cc.jar",
						},
						{
							Jar: "dbfs://ff/gg/hh.jar",
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
					Queue: &jobs.QueueSettings{
						Enabled: true,
					},
					RunAs: &JobRunAs{
						UserName: "user@mail.com",
					},
					Deployment: &jobs.JobDeployment{
						Kind:             "BUNDLE",
						MetadataFilePath: "/a/b/c",
					},
					EditMode: "UI_LOCKED",
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
						Libraries: []compute.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
							{
								Jar: "dbfs://ff/gg/hh.jar",
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
						Queue: &jobs.QueueSettings{
							Enabled: true,
						},
						RunAs: &JobRunAs{
							UserName: "user@mail.com",
						},
						Deployment: &jobs.JobDeployment{
							Kind:             "BUNDLE",
							MetadataFilePath: "/a/b/c",
						},
						EditMode: "UI_LOCKED",
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
		queue {
			enabled = true
		}
		run_as {
			user_name = "user@mail.com"
		}
		deployment {
			kind = "BUNDLE"
			metadata_file_path = "/a/b/c"
		}
		edit_mode = "UI_LOCKED"`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_MultiTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "Featurizer",
					Tasks: []JobTaskSettings{
						{
							TaskKey:           "a",
							ExistingClusterID: "abc",
							Libraries: []compute.Library{
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
										Value:     50000000000, // 5 * 10^10
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
					Queue: &jobs.QueueSettings{
						Enabled: false,
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
				Resource: "/api/2.2/jobs/get?job_id=789",
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
					value  = 50000000000				  
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

func TestResourceJobCreate_TaskOrder(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "Featurizer",
					Tasks: []JobTaskSettings{
						{
							TaskKey:           "a",
							ExistingClusterID: "abc",
							NotebookTask: &NotebookTask{
								NotebookPath: "/a",
							},
						},
						{
							TaskKey: "b",
							DependsOn: []jobs.TaskDependency{
								{
									TaskKey: "a",
								},
							},
							ExistingClusterID: "abc",
							NotebookTask: &NotebookTask{
								NotebookPath: "/b",
							},
						},
						{
							TaskKey: "c",
							DependsOn: []jobs.TaskDependency{
								{
									TaskKey: "a",
								},
								{
									TaskKey: "b",
								},
							},
							ExistingClusterID: "abc",
							NotebookTask: &NotebookTask{
								NotebookPath: "/c",
							},
						},
						{
							TaskKey: "d",
							DependsOn: []jobs.TaskDependency{
								{
									TaskKey: "a",
								},
								{
									TaskKey: "b",
								},
								{
									TaskKey: "c",
								},
							},
							ExistingClusterID: "abc",
							NotebookTask: &NotebookTask{
								NotebookPath: "/d",
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
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
				Resource: "/api/2.2/jobs/get?job_id=789",
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
							{
								TaskKey: "d",
							},
							{
								TaskKey: "c",
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

			notebook_task {
				notebook_path = "/a"
			}
		}

		task {
			task_key = "b"

			depends_on {
				task_key = "a"
			}

			existing_cluster_id = "abc"

			notebook_task {
				notebook_path = "/b"
			}
		}
		
		task {
			task_key = "c"

			depends_on {
				task_key = "a"
			}

			depends_on {
				task_key = "b"
			}

			existing_cluster_id = "abc"

			notebook_task {
				notebook_path = "/c"
			}
		}

		task {
			task_key = "d"

			depends_on {
				task_key = "a"
			}

			depends_on {
				task_key = "b"
			}

			depends_on {
				task_key = "c"
			}

			existing_cluster_id = "abc"

			notebook_task {
				notebook_path = "/d"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_ConditionTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "ConditionTaskTesting",
					Tasks: []JobTaskSettings{
						{
							TaskKey: "a",
							ConditionTask: &jobs.ConditionTask{
								Left:  "123",
								Op:    "EQUAL_TO",
								Right: "123",
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					MaxConcurrentRuns: 1,
				},
				Response: Job{
					JobID: 231,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=231",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "a",
								ConditionTask: &jobs.ConditionTask{
									Left:  "123",
									Op:    "EQUAL_TO",
									Right: "123",
								},
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "ConditionTaskTesting"
	
		task {
			task_key = "a"
			condition_task {
				left = "123"
				op = "EQUAL_TO"
				right = "123"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "231", d.Id())
}

func TestResourceJobCreate_ForEachTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "Foreach-task-testing",
					Tasks: []JobTaskSettings{
						{
							TaskKey: "for_each_task_key",
							ForEachTask: &ForEachTask{
								Concurrency: 1,
								Inputs:      "[1, 2, 3, 4, 5, 6]",
								Task: ForEachNestedTask{
									TaskKey:           "nested_task_key",
									ExistingClusterID: "abc",
									NotebookTask: &NotebookTask{
										NotebookPath: "/Stuff",
									},
								},
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					MaxConcurrentRuns: 1,
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "for_each_task_key",
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "Foreach-task-testing"

		task {
			task_key = "for_each_task_key"
			for_each_task {
				concurrency = 1

				inputs = "[1, 2, 3, 4, 5, 6]"

				task {

					task_key = "nested_task_key"

					existing_cluster_id = "abc"
					
						notebook_task {
							notebook_path = "/Stuff"
						}
				}
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}
func TestResourceJobCreate_PowerBiTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockJobsAPI().EXPECT()
			e.Create(mock.Anything, jobs.CreateJob{
				Name:              "power_bi_task_name",
				MaxConcurrentRuns: 1,
				Queue: &jobs.QueueSettings{
					Enabled: false,
				},
				Tasks: []jobs.Task{
					{
						TaskKey: "power_bi_task_key",
						PowerBiTask: &jobs.PowerBiTask{
							ConnectionResourceName: "test-connection",
							PowerBiModel: &jobs.PowerBiModel{
								AuthenticationMethod: jobs.AuthenticationMethodOauth,
								ModelName:            "TestModel",
								OverwriteExisting:    true,
								StorageMode:          jobs.StorageModeDirectQuery,
								WorkspaceName:        "TestWorkspace",
							},
							RefreshAfterUpdate: true,
							Tables: []jobs.PowerBiTable{
								{
									Catalog:     "TestCatalog",
									Name:        "TestTable1",
									Schema:      "TestSchema",
									StorageMode: jobs.StorageModeDirectQuery,
								},
								{
									Catalog:     "TestCatalog",
									Name:        "TestTable2",
									Schema:      "TestSchema",
									StorageMode: jobs.StorageModeDual,
								},
							},
							WarehouseId: "12345",
						},
					},
				},
			}).
				Return(&jobs.CreateResponse{
					JobId: 789,
				}, nil)
			e.GetByJobId(mock.Anything, int64(789)).Return(&jobs.Job{
				JobId: 789,
				Settings: &jobs.JobSettings{
					Name: "power_bi_task_name",
					Tasks: []jobs.Task{
						{
							TaskKey: "power_bi_task_key",
							PowerBiTask: &jobs.PowerBiTask{
								ConnectionResourceName: "test-connection",
								PowerBiModel: &jobs.PowerBiModel{
									AuthenticationMethod: jobs.AuthenticationMethodOauth,
									ModelName:            "TestModel",
									OverwriteExisting:    true,
									StorageMode:          jobs.StorageModeDirectQuery,
									WorkspaceName:        "TestWorkspace",
								},
								RefreshAfterUpdate: true,
								Tables: []jobs.PowerBiTable{
									{
										Catalog:     "TestCatalog",
										Name:        "TestTable1",
										Schema:      "TestSchema",
										StorageMode: jobs.StorageModeDirectQuery,
									},
									{
										Catalog:     "TestCatalog",
										Name:        "TestTable2",
										Schema:      "TestSchema",
										StorageMode: jobs.StorageModeDual,
									},
								},
								WarehouseId: "12345",
							},
						},
					},
				},
			}, nil)
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `name = "power_bi_task_name"
		task {
			task_key = "power_bi_task_key"
			power_bi_task {
				connection_resource_name = "test-connection"
				power_bi_model {
					authentication_method = "OAUTH"
					model_name = "TestModel"
					overwrite_existing = true
					storage_mode = "DIRECT_QUERY"
					workspace_name = "TestWorkspace"
				}
				refresh_after_update = true
				tables {
					catalog = "TestCatalog"
					name = "TestTable1"
					schema = "TestSchema"
					storage_mode = "DIRECT_QUERY"
				}
				tables {
					catalog = "TestCatalog"
					name = "TestTable2"
					schema = "TestSchema"
					storage_mode = "DUAL"
				}
				warehouse_id = "12345"
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
				Resource: "/api/2.2/jobs/create",
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
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					MaxConcurrentRuns: 1,
					Parameters: []jobs.JobParameterDefinition{
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
				Resource: "/api/2.2/jobs/get?job_id=231",
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
						Parameters: []jobs.JobParameterDefinition{
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

func TestResourceJobCreate_JobParameters_EmptyDefault(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name:              "JobParameterTesting",
					MaxConcurrentRuns: 1,
					Tasks: []JobTaskSettings{
						{
							TaskKey: "a",
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					Parameters: []jobs.JobParameterDefinition{
						{
							Name:    "key",
							Default: "",
						},
					},
				},
				Response: Job{
					JobID: 231,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=231",
				Response: Job{
					// good enough for mock
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey: "a",
							},
						},
						Parameters: []jobs.JobParameterDefinition{
							{
								Name:    "key",
								Default: "",
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
				name = "key"
				default = ""
		}

		task {
			task_key = "a"
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "231", d.Id())
}

func TestResourceJobCreate_JobParameters_DefaultIsRequired(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "JobParameterTesting"

		parameter {
				name = "key"
		}

		task {
			task_key = "a"
		}`,
	}.ExpectError(t, "invalid config supplied. [parameter.#.default] Missing required argument")
}

func TestResourceJobCreate_JobParameters_SingleTasksConflict(t *testing.T) {
	qa.ResourceFixture{
		Create:   true,
		Resource: ResourceJob(),
		HCL: `
		name = "JobParameterTesting"

		parameter {
				name = "key"
				default = ""
		}

		notebook_task {
			notebook_path = "/Shared/test"
		}`,
	}.ExpectError(t, "invalid config supplied. [notebook_task] Conflicting configuration arguments")
}

func TestResourceJobCreate_JobClusters(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
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
							DependsOn: []jobs.TaskDependency{
								{
									TaskKey: "a",
								},
							},
							NotebookTask: &NotebookTask{
								NotebookPath: "/Stuff",
							},
						},
						{
							TaskKey: "c",
							NewCluster: &clusters.Cluster{
								SparkVersion: "d",
								NodeTypeID:   "e",
								NumWorkers:   0,
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
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
								NumWorkers:   0,
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
				Resource: "/api/2.2/jobs/get?job_id=17",
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
			  num_workers   = 0
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

			depends_on {
				task_key = "a"
			}

			new_cluster {
				spark_version = "a"
				node_type_id = "b"
				num_workers = 3
			}

			notebook_task {
				notebook_path = "/Stuff"
			}
		}
		
		task {
			task_key = "c"
			new_cluster {
				spark_version = "d"
				node_type_id = "e"
				num_workers = 0
			}
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "17", d.Id())
}

func TestResourceJobCreate_JobCompute(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name: "JobEnvironments",
					Tasks: []JobTaskSettings{
						{
							TaskKey:        "b",
							EnvironmentKey: "j",
							NotebookTask: &NotebookTask{
								NotebookPath: "/Stuff",
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					MaxConcurrentRuns: 1,
					Environments: []jobs.JobEnvironment{
						{
							EnvironmentKey: "j",
							Spec: &compute.Environment{
								Client: "1",
								Dependencies: []string{
									"cowsay",
									"-r /Workspace/Users/lisa@company.com/my.whl",
								},
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
				Resource: "/api/2.2/jobs/get?job_id=18",
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
		name = "JobEnvironments"
		environment {
			environment_key = "j"
			spec {
			  client   	= "1"
			  dependencies = ["cowsay", "-r /Workspace/Users/lisa@company.com/my.whl"]
			}
		}
		task {
			task_key = "b"
			environment_key = "j"
			notebook_task {
				notebook_path = "/Stuff"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "18", d.Id())
}

func TestResourceJobCreate_DashboardTask(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: jobs.CreateJob{
					Name:              "DashboardTask",
					MaxConcurrentRuns: 1,
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
					Tasks: []jobs.Task{
						{
							TaskKey: "a",
							DashboardTask: &jobs.DashboardTask{
								Subscription: &jobs.Subscription{
									Subscribers: []jobs.SubscriptionSubscriber{
										{UserName: "user@domain.com"},
										{DestinationId: "Test"},
									},
									Paused:        true,
									CustomSubject: "\"custom subject\"",
								},
								WarehouseId: "\"dca3a0ba199040eb\"",
								DashboardId: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
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
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: jobs.Job{
					JobId: 789,
					Settings: &jobs.JobSettings{
						Name: "DashboardTask",
						Tasks: []jobs.Task{
							{
								TaskKey: "a",
								DashboardTask: &jobs.DashboardTask{
									Subscription: &jobs.Subscription{
										Subscribers: []jobs.SubscriptionSubscriber{
											{UserName: "user@domain.com"},
											{DestinationId: "Test"},
										},
										Paused:        true,
										CustomSubject: "\"custom subject\"",
									},
									WarehouseId: "\"dca3a0ba199040eb\"",
									DashboardId: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
								},
							},
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceJob(),
		HCL: `name = "DashboardTask"
		task {
		  task_key = "a"
		  dashboard_task {
			warehouse_id = "\"dca3a0ba199040eb\""
			subscription {
				subscribers {
    				user_name = "user@domain.com"
  				}
				subscribers {
					destination_id = "Test"
				}
				paused = true
				custom_subject = "\"custom subject\""
			}
			dashboard_id = "3cf91a42-6217-4f3c-a6f0-345d489051b9"
		  }
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "789", d.Id())
}

func TestResourceJobCreate_SqlSubscriptions(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: jobs.CreateJob{
					Name:              "TF SQL task subscriptions",
					MaxConcurrentRuns: 1,
					Tasks: []jobs.Task{
						{
							TaskKey: "a",
							SqlTask: &jobs.SqlTask{
								WarehouseId: "dca3a0ba199040eb",
								Alert: &jobs.SqlTaskAlert{
									AlertId: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
									Subscriptions: []jobs.SqlTaskSubscription{
										{UserName: "user@domain.com"},
										{DestinationId: "Test"},
									},
									PauseSubscriptions: true,
								},
							},
						},
						{
							TaskKey: "d",
							SqlTask: &jobs.SqlTask{
								WarehouseId: "dca3a0ba199040eb",
								Dashboard: &jobs.SqlTaskDashboard{
									DashboardId:        "d81a7760-7fd2-443e-bf41-95a60c2f4c7c",
									PauseSubscriptions: false,
									Subscriptions: []jobs.SqlTaskSubscription{
										{UserName: "user@domain.com"},
										{DestinationId: "Test"},
									},
									CustomSubject: "test",
								},
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: jobs.Job{
					JobId: 789,
					Settings: &jobs.JobSettings{
						Name: "TF SQL task subscriptions",
						Tasks: []jobs.Task{
							{
								TaskKey: "a",
								SqlTask: &jobs.SqlTask{
									WarehouseId: "dca3a0ba199040eb",
									Alert: &jobs.SqlTaskAlert{
										AlertId: "3cf91a42-6217-4f3c-a6f0-345d489051b9",
										Subscriptions: []jobs.SqlTaskSubscription{
											{UserName: "user@domain.com"},
											{DestinationId: "Test"},
										},
										PauseSubscriptions: true,
									},
								},
							},
							{
								TaskKey: "d",
								SqlTask: &jobs.SqlTask{
									WarehouseId: "dca3a0ba199040eb",
									Dashboard: &jobs.SqlTaskDashboard{
										DashboardId: "d81a7760-7fd2-443e-bf41-95a60c2f4c7c",
										Subscriptions: []jobs.SqlTaskSubscription{
											{UserName: "user@domain.com"},
											{DestinationId: "Test"},
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
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Name:              "TF RunJobTask Main Job",
					MaxConcurrentRuns: 1,
					Tasks: []JobTaskSettings{
						{
							TaskKey: "runJobTask",
							RunJobTask: &RunJobTask{
								JobID: 123,
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
				},
				Response: Job{
					JobID: 123,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=123",
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
	}.ApplyNoError(t)
}

func TestResourceJobCreate_Trigger_TableUpdateCreate(t *testing.T) {
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
					Trigger: &Trigger{
						PauseStatus: "UNPAUSED",
						TableUpdate: &TableUpdate{
							TableNames: []string{"catalog.schema.table1", "catalog.schema.table2"},
							Condition:  "ALL_UPDATED",
						},
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
						Trigger: &Trigger{
							PauseStatus: "UNPAUSED",
							TableUpdate: &TableUpdate{
								TableNames: []string{"catalog.schema.table1", "catalog.schema.table2"},
								Condition:  "ALL_UPDATED",
							},
						},
					},
				},
			},
		},
		HCL: `
		trigger {
			pause_status = "UNPAUSED"
			table_update {
				table_names = [
					"catalog.schema.table1",
					"catalog.schema.table2"
				]
				condition = "ALL_UPDATED"
			}
		}
		max_concurrent_runs = 1
		name = "Test"
		`,
	}.ApplyNoError(t)
}

func TestResourceJobCreate_Trigger_PeriodicCreate(t *testing.T) {
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
					Trigger: &Trigger{
						PauseStatus: "UNPAUSED",
						Periodic: &Periodic{
							Interval: 4,
							Unit:     "HOURS",
						},
					},
				},
				Response: Job{
					JobID: 1042,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=1042",
				Response: Job{
					JobID: 1042,
					Settings: &JobSettings{
						MaxConcurrentRuns: 1,
						Name:              "Test",
						Trigger: &Trigger{
							PauseStatus: "UNPAUSED",
							Periodic: &Periodic{
								Interval: 4,
								Unit:     "HOURS",
							},
						},
					},
				},
			},
		},
		HCL: `
		trigger {
			pause_status = "UNPAUSED"
			periodic {
				interval = 4
				unit = "HOURS"
			}
		}
		max_concurrent_runs = 1
		name = "Test"
		`,
	}.ApplyNoError(t)
}

func TestResourceJobUpdate_ControlRunState_ContinuousUpdateRunNow(t *testing.T) {
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
	}.ApplyNoError(t)
}

func TestResourceJobUpdate_ControlRunState_ContinuousUpdateRunNowFailsWith409(t *testing.T) {
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
				Status: 409,
				Response: common.APIErrorBody{
					ErrorCode: "CONFLICT",
					Message:   "A concurrent request to run the continuous job is already in progress. Please wait for it to complete before issuing a new request.",
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
	}.ApplyNoError(t)
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
	}.ApplyNoError(t)
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

func TestResourceJobUpdateNoQueue(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					JobID:    789,
					Settings: &JobSettings{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
						},
						Name:              "Untitled",
						MaxConcurrentRuns: 1,
					},
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						ExistingClusterID: "abc",
						Tasks: []JobTaskSettings{
							{
								TaskKey: "b",
							},
						},
						MaxConcurrentRuns: 1,
					},
				},
			},
		},
		Update:   true,
		Resource: ResourceJob(),
		ID:       "789",
		HCL: `
		max_concurrent_runs = 1

		task {
			task_key = "b"
		}`,
	}.ApplyNoError(t)
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
					Libraries: []compute.Library{
						{
							Jar: "dbfs://aa/bb/cc.jar",
						},
					},
					Name: "Featurizer",
					WebhookNotifications: &jobs.WebhookNotifications{
						OnStart:   []jobs.Webhook{{Id: "id1"}, {Id: "id2"}, {Id: "id3"}},
						OnSuccess: []jobs.Webhook{{Id: "id2"}},
						OnFailure: []jobs.Webhook{{Id: "id3"}},
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
						Libraries: []compute.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
						},
						Name: "Featurizer",
						WebhookNotifications: &jobs.WebhookNotifications{
							OnStart:   []jobs.Webhook{{Id: "id1"}, {Id: "id2"}, {Id: "id3"}},
							OnSuccess: []jobs.Webhook{{Id: "id2"}},
							OnFailure: []jobs.Webhook{{Id: "id3"}},
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
				Resource: "/api/2.2/jobs/create",
				ExpectedRequest: JobSettings{
					Tasks: []JobTaskSettings{
						{
							TaskKey: "b",
							NotebookTask: &NotebookTask{
								NotebookPath: "/GitSourcedNotebook",
							},
						},
					},
					Queue: &jobs.QueueSettings{
						Enabled: false,
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
				Resource: "/api/2.2/jobs/get?job_id=789",
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
		HCL: `
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
	}.ExpectError(t, "git source is not empty but Git Provider is not specified and cannot be guessed by url &{GitBranch: GitCommit: GitProvider: GitSnapshot:<nil> GitTag:0.4.8 GitUrl:https://custom.git.hosting.com/databricks/terraform-provider-databricks JobSource:<nil> ForceSendFields:[]}")
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
						Libraries: []compute.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
							{
								Jar: "dbfs://ff/gg/hh.jar",
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
	libraries := d.Get("library").([]interface{})
	assert.Len(t, libraries, 2)
	allDbfsLibs := []string{}
	for _, lib := range libraries {
		allDbfsLibs = append(allDbfsLibs, lib.(map[string]any)["jar"].(string))
	}
	assert.Contains(t, allDbfsLibs, "dbfs://ff/gg/hh.jar")
	assert.Contains(t, allDbfsLibs, "dbfs://aa/bb/cc.jar")

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
						Libraries: []compute.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
							{
								Jar: "dbfs://ff/gg/hh.jar",
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
						Libraries: []compute.Library{
							{
								Jar: "dbfs://aa/bb/cc.jar",
							},
							{
								Jar: "dbfs://ff/gg/hh.jar",
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

func TestResourceJobUpdate_RunIfSuppressesDiffIfAllSuccess(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						MaxConcurrentRuns: 1,
						Tasks: []JobTaskSettings{
							{
								TaskKey: "task1",
								NotebookTask: &NotebookTask{
									NotebookPath: "/foo/bar",
								},
								// The diff is suppressed here. The API payload
								// contains the "run_if" value from the terraform
								// state.
								RunIf: "ALL_SUCCESS",
							},
							{
								TaskKey: "task2",
								ForEachTask: &ForEachTask{
									Inputs: "abc",
									Task: ForEachNestedTask{
										TaskKey:      "task3",
										NotebookTask: &NotebookTask{NotebookPath: "/bar/foo"},
										// The diff is suppressed here. Value is from
										// the terraform state.
										RunIf: "ALL_SUCCESS",
									},
								},
							},
						},
						Name: "My job",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						Name: "My job",
						Tasks: []JobTaskSettings{
							{
								TaskKey:      "task1",
								NotebookTask: &NotebookTask{NotebookPath: "/foo/bar"},
								RunIf:        "ALL_SUCCESS",
							},
							{
								TaskKey: "task2",
								ForEachTask: &ForEachTask{
									Inputs: "abc",
									Task: ForEachNestedTask{
										TaskKey:      "task3",
										NotebookTask: &NotebookTask{NotebookPath: "/bar/foo"},
										RunIf:        "ALL_SUCCESS",
									},
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
		InstanceState: map[string]string{
			"task.0.run_if":                        "ALL_SUCCESS",
			"task.1.for_each_task.0.task.0.run_if": "ALL_SUCCESS",
		},
		HCL: `
		name = "My job"
		task {
			task_key = "task1"
			notebook_task {
				notebook_path = "/foo/bar"
			}
		}
		task {
			task_key = "task2"
			for_each_task {
				inputs = "abc"
				task {
					task_key = "task3"
					notebook_task {
						notebook_path = "/bar/foo"
					}
				}
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
}

func TestResourceJobUpdate_RunIfDoesNotSuppressIfNotAllSuccess(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						MaxConcurrentRuns: 1,
						Tasks: []JobTaskSettings{
							{
								TaskKey: "task1",
								NotebookTask: &NotebookTask{
									NotebookPath: "/foo/bar",
								},
								// The diff is not suppressed here. Thus the API payload
								// explicitly does not set run_if here, to unset it in the
								// job definition.
								// RunIf is not set, as implied from the HCL config.
							},
							{
								TaskKey: "task2",
								ForEachTask: &ForEachTask{
									Inputs: "abc",
									Task: ForEachNestedTask{
										TaskKey:      "task3",
										NotebookTask: &NotebookTask{NotebookPath: "/bar/foo"},
										// The diff is not suppressed. RunIf is
										// not set, as implied from the HCL config.
									},
								},
							},
						},
						Name: "My job",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					JobID: 789,
					Settings: &JobSettings{
						Name: "My job",
						Tasks: []JobTaskSettings{
							{
								TaskKey:      "task1",
								NotebookTask: &NotebookTask{NotebookPath: "/foo/bar"},
								RunIf:        "AT_LEAST_ONE_FAILED",
							},
							{
								TaskKey: "task2",
								ForEachTask: &ForEachTask{
									Task: ForEachNestedTask{
										TaskKey: "task3",
										RunIf:   "AT_LEAST_ONE_FAILED",
									},
								},
							},
						},
					},
				},
			},
		},
		ID:     "789",
		Update: true,
		InstanceState: map[string]string{
			"task.0.run_if":                        "AT_LEAST_ONE_FAILED",
			"task.1.for_each_task.0.task.0.run_if": "AT_LEAST_ONE_FAILED",
		},
		Resource: ResourceJob(),
		HCL: `
		name = "My job"
		task {
			task_key = "task1"
			notebook_task {
				notebook_path = "/foo/bar"
			}
		}
		task {
			task_key = "task2"
			for_each_task {
				inputs = "abc"
				task {
					task_key = "task3"
					notebook_task {
						notebook_path = "/bar/foo"
					}
				}
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
}

func TestResourceJobUpdate_NodeTypeToInstancePool(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: jobs.ResetJob{
					JobId: 789,
					NewSettings: jobs.JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						JobClusters: []jobs.JobCluster{
							{
								JobClusterKey: "job_cluster_1",
								NewCluster: compute.ClusterSpec{
									InstancePoolId:       "instance-pool-worker-job",
									DriverInstancePoolId: "instance-pool-driver-job",
									SparkVersion:         "spark-3",
									NumWorkers:           3,
								},
							},
						},
						Tasks: []jobs.Task{
							{
								TaskKey: "task1",
								NewCluster: &compute.ClusterSpec{
									InstancePoolId:       "instance-pool-worker-task",
									DriverInstancePoolId: "instance-pool-driver-task",
									SparkVersion:         "spark-2",
									NumWorkers:           2,
								},
							},
						},
						Name:              "Featurizer New",
						MaxConcurrentRuns: 1,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: jobs.Job{
					JobId: 789,
					Settings: &jobs.JobSettings{
						Name:              "Featurizer New",
						MaxConcurrentRuns: 1,
					},
				},
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		InstanceState: map[string]string{
			"task.0.new_cluster.0.node_type_id":               "node-type-id-worker-task",
			"task.0.new_cluster.0.driver_node_type_id":        "node-type-id-driver-task",
			"job_cluster.0.new_cluster.0.node_type_id":        "node-type-id-worker-job",
			"job_cluster.0.new_cluster.0.driver_node_type_id": "node-type-id-driver-job",
		},
		HCL: `
		task = {
			task_key = "task1"
			new_cluster = {
				instance_pool_id = "instance-pool-worker-task"
				driver_instance_pool_id = "instance-pool-driver-task"
				spark_version = "spark-2"
				num_workers = 2
			}
		}
		job_cluster = {
			job_cluster_key = "job_cluster_1"
			new_cluster = {
				instance_pool_id = "instance-pool-worker-job"
				driver_instance_pool_id = "instance-pool-driver-job"
				spark_version = "spark-3"
				num_workers = 3
			}
		}
		max_concurrent_runs = 1
		name = "Featurizer New"`,
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
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: jobs.UpdateJob{
					JobId: 789,
					NewSettings: &jobs.JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						Tasks: []jobs.Task{
							{
								TaskKey: "task1",
								NewCluster: &compute.ClusterSpec{
									NodeTypeId:   "node-type-id-2",
									SparkVersion: "spark-2",
									NumWorkers:   2,
								},
							},
						},
						JobClusters: []jobs.JobCluster{
							{
								JobClusterKey: "job_cluster_1",
								NewCluster: compute.ClusterSpec{
									NodeTypeId:   "node-type-id-3",
									SparkVersion: "spark-3",
									NumWorkers:   3,
								},
							},
						},
						Name:              "Featurizer New",
						MaxConcurrentRuns: 1,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: jobs.Job{
					JobId: 789,
					Settings: &jobs.JobSettings{
						Name:              "Featurizer New",
						MaxConcurrentRuns: 1,
					},
				},
			},
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		InstanceState: map[string]string{
			"task.0.new_cluster.0.node_type_id":        "node-type-id-worker-task",
			"task.0.instance_pool_id":                  "instance-pool-id-worker",
			"task.0.driver_instance_pool_id":           "instance-pool-id-driver",
			"job_cluster.0.new_cluster.0.node_type_id": "node-type-id-worker-job",
			"job_cluster.0.instance_pool_id":           "instance-pool-id-worker",
			"job_cluster.0.driver_instance_pool_id":    "instance-pool-id-driver",
		},
		HCL: `
		task = {
			task_key = "task1"
			new_cluster = {
				node_type_id = "node-type-id-2"
				spark_version = "spark-2"
				num_workers = 2
			}
		}
		job_cluster = {
			job_cluster_key = "job_cluster_1"
			new_cluster = {
				node_type_id = "node-type-id-3"
				spark_version = "spark-3"
				num_workers = 3
			}
		}
		max_concurrent_runs = 1
		name = "Featurizer New"`,
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
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Queue: &jobs.QueueSettings{
							Enabled: false,
						},
						Name: "Featurizer New",
						Tasks: []JobTaskSettings{
							{
								TaskKey:           "task1",
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
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []JobTaskSettings{
							{
								TaskKey:           "task1",
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
			task_key = "task1"
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
				Resource: "/api/2.2/jobs/delete",
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

func TestJobsAPIList(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25",
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
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 1,
					},
				},
				HasMore:       true,
				NextPageToken: "aaaa",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&page_token=aaaa",
			Response: JobListResponse{
				Jobs: []Job{
					{
						JobID: 2,
					},
				},
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
			Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&name=test",
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
