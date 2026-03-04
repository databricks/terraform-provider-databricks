package acceptance

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
)

func accountHostCreateJobTest(t *testing.T) {
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	jobName := "tf-account-" + RandomName() + "-job-1"

	run(t, []Step{
		{
			Template: `
			resource "databricks_job" "j1" {
				name = "` + jobName + `"
				provider_config {
					workspace_id = ` + workspaceID + `
				}
				task {
					task_key = "check"
					condition_task {
						left  = "true"
						op    = "EQUAL_TO"
						right = "true"
					}
				}
			}
			`,
			Check: ResourceCheck("databricks_job.j1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
				if err != nil {
					return err
				}
				jobID, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					return err
				}
				res, err := w.Jobs.Get(ctx, jobs.GetJobRequest{JobId: jobID})
				if err != nil {
					return err
				}
				if res.Settings.Name != jobName {
					return fmt.Errorf("expected job name %q, got %q", jobName, res.Settings.Name)
				}
				return nil
			}),
		},
	})
}

func TestAccAccountHostCreateJobsAWS(t *testing.T) {
	LoadAccountEnv(t)
	if !IsAws(t) {
		Skipf(t)("This test is only running on AWS")
	}
	accountHostCreateJobTest(t)
}

func TestAccAccountHostCreateJobsGCP(t *testing.T) {
	LoadAccountEnv(t)
	if !IsGcp(t) {
		Skipf(t)("This test is only running on GCP")
	}
	accountHostCreateJobTest(t)
}

func TestAccAccountHostCreateJobsAzure(t *testing.T) {
	LoadAccountEnv(t)
	if !IsAzure(t) {
		Skipf(t)("This test is only running on Azure")
	}
	accountHostCreateJobTest(t)
}
