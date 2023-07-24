package acceptance

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestAccJobTasks(t *testing.T) {
	workspaceLevel(t, step{
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

				notebook_task {
					notebook_path = databricks_notebook.this.path
				}
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
	randomName1 := RandomName("notebook-")
	randomName2 := RandomName("updated-notebook-")
	workspaceLevel(t, step{
		// A new continuous job with empty block should be started automatically
		Template: getJobTemplate(randomName1, ``),
		Check:    resourceCheck("databricks_job.this", waitForRunToStart),
	}, step{
		// Updating the notebook should cancel the existing run
		Template: getJobTemplate(randomName2, ``),
		Check:    resourceCheck("databricks_job.this", waitForRunToStart),
	}, step{
		// Marking the job as paused should cancel existing run and not start a new one
		Template: getJobTemplate(randomName2, `pause_status = "PAUSED"`),
		Check:    resourceCheck("databricks_job.this", waitForAllRunsToEnd),
	}, step{
		// No pause status should be the equivalent of unpaused
		Template: getJobTemplate(randomName2, `pause_status = "UNPAUSED"`),
		Check:    resourceCheck("databricks_job.this", waitForRunToStart),
	})
}
