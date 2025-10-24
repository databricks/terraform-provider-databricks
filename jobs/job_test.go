package jobs_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestAccJobTasks(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_current_user" "me" {}
		data "databricks_spark_version" "latest" {}
		data "databricks_node_type" "smallest" {
			local_disk = true
		}

		resource "databricks_notebook" "this" {
			path     = "${data.databricks_current_user.me.home}/Terraform{var.RANDOM}"
			language = "PYTHON"
			content_base64 = base64encode(<<-EOT
				# created from ${abspath(path.module)}
				display(spark.range(10))
				EOT
			)
		}

		resource "databricks_job" "this" {
			name = "{var.RANDOM}"

			job_cluster {
				job_cluster_key = "j"
				new_cluster {
					num_workers   = 0  // Setting it to zero intentionally to cover edge case.
					spark_version = data.databricks_spark_version.latest.id
					node_type_id  = data.databricks_node_type.smallest.id
					custom_tags = {
						"ResourceClass" = "SingleNode"
					}
					spark_conf = {
						"spark.databricks.cluster.profile" : "singleNode"
						"spark.master" : "local[*,4]"
					}
				}
			}

			task {
				task_key = "a"

				new_cluster {
					num_workers   = 1
					spark_version = data.databricks_spark_version.latest.id
					node_type_id  = data.databricks_node_type.smallest.id
				}

				notebook_task {
					notebook_path = databricks_notebook.this.path
				}
			}

			task {
				task_key = "b"

				depends_on {
					task_key = "a"
				}

				new_cluster {
					num_workers   = 8
					spark_version = data.databricks_spark_version.latest.id
					node_type_id  = data.databricks_node_type.smallest.id
				}

				notebook_task {
					notebook_path = databricks_notebook.this.path
				}
			}

			task {
				task_key = "c"

				job_cluster_key = "j"

				depends_on {
					task_key = "b"
				}

				run_if = "ALL_DONE"

				notebook_task {
					notebook_path = databricks_notebook.this.path
					base_parameters = {
						"param_0" = "{{job.parameters.empty_default}}"
						"param_1" = "{{job.parameters.non_empty_default}}"
					}
				}
			}

			parameter {
				name = "empty_default"
				default = ""
			}

			parameter {
				name = "non_empty_default"
				default = "non_empty"
			}
		}`,
	})
}

func TestAccForEachTask(t *testing.T) {
	t.Skip("Skipping this test because feature not enabled in Prod")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_current_user" "me" {}
		data "databricks_spark_version" "latest" {}
		data "databricks_node_type" "smallest" {
			local_disk = true
		}

		resource "databricks_notebook" "this" {
			path     = "${data.databricks_current_user.me.home}/Terraform{var.RANDOM}"
			language = "PYTHON"
			content_base64 = base64encode(<<-EOT
				# created from ${abspath(path.module)}
				display(spark.range(10))
				EOT
			)
		}

		resource "databricks_job" "this" {
			name = "{var.RANDOM}"

			job_cluster {
				job_cluster_key = "j"
				new_cluster {
					num_workers   = 20
					spark_version = data.databricks_spark_version.latest.id
					node_type_id  = data.databricks_node_type.smallest.id
				}
			}

			task {
				task_key = "for_each_task_key"
				for_each_task {
					concurrency = 1
					inputs = "[1, 2, 3, 4, 5, 6]"
					task {
						task_key        = "nested_task_key"
						job_cluster_key = "j"

						notebook_task {
							notebook_path = databricks_notebook.this.path
						}
					}
				}
			}

			parameter {
				name = "empty_default"
				default = ""
			}

			parameter {
				name = "non_empty_default"
				default = "non_empty"
			}
		}`,
	})
}

// An integration test which creates a continuous job with control_run_state = true, verifying
// that a job run was triggered within 5 minutes of the job creation. Then, the test updates the
// job, verifying that the existing run was cancelled within 5 minutes of the update.
func TestAccJobControlRunState(t *testing.T) {
	getJobTemplate := func(name string, continuousBlock string) string {
		return `
		data "databricks_current_user" "me" {}
		data "databricks_spark_version" "latest" {}
		data "databricks_node_type" "smallest" {
			local_disk = true
		}

		resource "databricks_notebook" "this" {
			path     = "${data.databricks_current_user.me.home}/Terraform` + name + `"
			language = "PYTHON"
			content_base64 = base64encode(<<-EOT
				# created from ${abspath(path.module)}
				import time

				display(spark.range(10))
				time.sleep(3600)
				EOT
			)
		}

		resource "databricks_job" "this" {
			name = "{var.RANDOM}"

			task {
				task_key = "a"

				new_cluster {
					num_workers   = 1
					spark_version = data.databricks_spark_version.latest.id
					node_type_id  = data.databricks_node_type.smallest.id
				}

				notebook_task {
					notebook_path = databricks_notebook.this.path
				}
			}

			continuous {
				` + continuousBlock + `
			}

			control_run_state = true
		}`
	}
	previousRunIds := make([]int64, 0)
	checkIfRunHasStarted := func(ctx context.Context, w *databricks.WorkspaceClient, jobID int64) (bool, error) {
		runs, err := w.Jobs.ListRunsAll(ctx, jobs.ListRunsRequest{JobId: jobID})
		assert.NoError(t, err)
		runIdsMap := make(map[int64]bool)
		for _, id := range previousRunIds {
			runIdsMap[id] = true
		}

		for _, run := range runs {
			if _, ok := runIdsMap[run.RunId]; !ok && run.State.LifeCycleState == "RUNNING" {
				previousRunIds = append(previousRunIds, run.RunId)
				return true, nil
			}
		}
		return false, nil
	}
	checkIfAllRunsHaveEnded := func(ctx context.Context, w *databricks.WorkspaceClient, jobID int64) (bool, error) {
		runs, err := w.Jobs.ListRunsAll(ctx, jobs.ListRunsRequest{JobId: jobID})
		assert.NoError(t, err)
		for _, run := range runs {
			if run.State.LifeCycleState == "RUNNING" {
				return false, nil
			}
		}
		return true, nil
	}
	retryFor := func(ctx context.Context, client *common.DatabricksClient, id string, lastErr error, f func(context.Context, *databricks.WorkspaceClient, int64) (bool, error)) error {
		ctx = context.WithValue(ctx, common.Api, common.API_2_1)
		w, err := client.WorkspaceClient()
		assert.NoError(t, err)
		jobID, err := strconv.ParseInt(id, 10, 64)
		assert.NoError(t, err)
		for i := 0; i < 100; i++ {
			success, err := f(ctx, w, jobID)
			if err != nil {
				return err
			}
			if success {
				return nil
			}
			// sleep for 5 seconds
			time.Sleep(5 * time.Second)
		}
		return lastErr
	}
	waitForRunToStart := func(ctx context.Context, client *common.DatabricksClient, id string) error {
		return retryFor(ctx, client, id, errors.New("timed out waiting for job run to start"), checkIfRunHasStarted)
	}
	waitForAllRunsToEnd := func(ctx context.Context, client *common.DatabricksClient, id string) error {
		return retryFor(ctx, client, id, errors.New("timed out waiting for job run to end"), checkIfAllRunsHaveEnded)
	}
	randomName1 := acceptance.RandomName("notebook-")
	randomName2 := acceptance.RandomName("updated-notebook-")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		// A new continuous job with empty block should be started automatically
		Template: getJobTemplate(randomName1, ``),
		Check:    acceptance.ResourceCheck("databricks_job.this", waitForRunToStart),
	}, acceptance.Step{
		// Updating the notebook should cancel the existing run
		Template: getJobTemplate(randomName2, ``),
		Check:    acceptance.ResourceCheck("databricks_job.this", waitForRunToStart),
	}, acceptance.Step{
		// Marking the job as paused should cancel existing run and not start a new one
		Template: getJobTemplate(randomName2, `pause_status = "PAUSED"`),
		Check:    acceptance.ResourceCheck("databricks_job.this", waitForAllRunsToEnd),
	}, acceptance.Step{
		// No pause status should be the equivalent of unpaused
		Template: getJobTemplate(randomName2, `pause_status = "UNPAUSED"`),
		Check:    acceptance.ResourceCheck("databricks_job.this", waitForRunToStart),
	})
}

func runAsTemplate(runAs string) string {
	return `
	data "databricks_current_user" "me" {}
	data "databricks_spark_version" "latest" {}
	data "databricks_node_type" "smallest" {
		local_disk = true
	}

	resource "databricks_notebook" "this" {
		path     = "${data.databricks_current_user.me.home}/Terraform{var.RANDOM}"
		language = "PYTHON"
		content_base64 = base64encode(<<-EOT
			# created from ${abspath(path.module)}
			display(spark.range(10))
			EOT
		)
	}

	resource "databricks_job" "this" {
		name = "{var.RANDOM}"

		job_cluster {
			job_cluster_key = "j"
			new_cluster {
				num_workers   = 20
				spark_version = data.databricks_spark_version.latest.id
				node_type_id  = data.databricks_node_type.smallest.id
			}
		}

		task {
			task_key = "c"
			job_cluster_key = "j"
			notebook_task {
				notebook_path = databricks_notebook.this.path
			}
		}

		run_as {
			` + runAs + `
		}
	}`
}

func TestAccJobRunAsUser(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_user" "this" {
			user_name = "` + qa.RandomEmail() + `"
		}
	` + runAsTemplate(`user_name = databricks_user.this.user_name`),
	})
}

func TestUcAccJobRunAsServicePrincipal(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	spId := acceptance.GetEnvOrSkipTest(t, "ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID")
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: runAsTemplate(`service_principal_name = "` + spId + `"`),
	})
}

func TestUcAccJobRunAsMutations(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	spId := acceptance.GetEnvOrSkipTest(t, "ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID")
	// Note: the attribute must match the type of principal that the test is run as.
	ctx := context.Background()
	attribute := acceptance.GetRunAsAttribute(t, ctx)
	acceptance.UnityWorkspaceLevel(
		t,
		// Provision job with service principal `run_as`
		acceptance.Step{
			Template: runAsTemplate(`service_principal_name = "` + spId + `"`),
		},
		// Update job to a user `run_as`
		acceptance.Step{
			Template: runAsTemplate(attribute + ` = data.databricks_current_user.me.user_name`),
		},
		// Update job back to a service principal `run_as`
		acceptance.Step{
			Template: runAsTemplate(`service_principal_name = "` + spId + `"`),
		},
	)
}

func TestAccRemoveWebhooks(t *testing.T) {
	acceptance.Skipf(t)("There is no API to create notification destinations. Once available, add here and enable this test.")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource databricks_job test {
			webhook_notifications {
				on_success {
					id = "a90cc1be-a29e-4eb7-a7e9-e4b0d4a7e7ae"
				}
			}
		}
		`,
	}, acceptance.Step{
		Template: `
		resource databricks_job test {}
		`,
	})
}

func TestAccPeriodicTrigger(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_job" "this" {
			name = "{var.RANDOM}"

			trigger {
				pause_status = "UNPAUSED"
				periodic {
					interval = 17
					unit = "HOURS"
				}
			}
		}`,
		Check: acceptance.ResourceCheck("databricks_job.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)

			jobID, err := strconv.ParseInt(id, 10, 64)
			assert.NoError(t, err)

			res, err := w.Jobs.Get(ctx, jobs.GetJobRequest{
				JobId: jobID,
			})
			assert.NoError(t, err)

			assert.Equal(t, jobs.PauseStatus("UNPAUSED"), res.Settings.Trigger.PauseStatus)
			assert.Equal(t, 17, res.Settings.Trigger.Periodic.Interval)
			assert.Equal(t, jobs.PeriodicTriggerConfigurationTimeUnit("HOURS"), res.Settings.Trigger.Periodic.Unit)

			return nil
		}),
	})
}

func TestAccJobDashboardTask(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_dashboard" "d1" {
			display_name			= 	"Dashboard Task Test {var.RANDOM}"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			parent_path				= 	"/Shared/provider-test"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
		}

		resource "databricks_job" "this" {
			name = "Dashboard Task Test {var.RANDOM}"

			task {
				task_key = "d"

				dashboard_task {
					dashboard_id = databricks_dashboard.d1.id
					warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
				}
			}
		}`,
	})
}

func TestAccJobClusterPolicySparkVersion(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_current_user" "me" {}
		data "databricks_spark_version" "latest" {}
		data "databricks_node_type" "smallest" {
			local_disk = true
		}
		resource "databricks_notebook" "this" {
			path     = "${data.databricks_current_user.me.home}/Terraform{var.RANDOM}"
			language = "PYTHON"
			content_base64 = base64encode(<<-EOT
				# created from ${abspath(path.module)}
				display(spark.range(10))
				EOT
			)
		}
		resource "databricks_cluster_policy" "this" {
			name = "test-policy-{var.RANDOM}"
			definition = jsonencode({
				"spark_version": {
					"type": "fixed",
					"value": data.databricks_spark_version.latest.id
				}
			})
		}
		resource "databricks_job" "this" {
			name = "test-job-{var.RANDOM}"
			job_cluster {
				job_cluster_key = "test-cluster"
				new_cluster {
					num_workers = 0
					node_type_id = data.databricks_node_type.smallest.id
					custom_tags = {
						"ResourceClass" = "SingleNode"
					}
					spark_conf = {
						"spark.databricks.cluster.profile" : "singleNode"
						"spark.master" : "local[*,4]"
					}

					// Apply the cluster policy to the job cluster for the Spark version.
					policy_id = databricks_cluster_policy.this.id
					apply_policy_default_values = true
				}
			}
			task {
				task_key = "test-task"
				job_cluster_key = "test-cluster"
				notebook_task {
					notebook_path = databricks_notebook.this.path
				}
			}
		}
`,
		// The configuration uses "apply_policy_default_values = true" to set the Spark version.
		// This means permanent drift will occur for the values sourced from the policy.
		ExpectNonEmptyPlan: true,
	})
}
