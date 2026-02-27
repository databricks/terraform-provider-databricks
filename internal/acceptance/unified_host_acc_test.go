package acceptance

import (
	"os"
	"testing"
)

func initUnifiedHostEnv(t *testing.T) {
	initTest(t, "account")
	unifiedHost := os.Getenv("UNIFIED_HOST")
	if unifiedHost == "" {
		Skipf(t)("UNIFIED_HOST environment variable is missing")
	}
	// Override provider env vars so the Terraform provider uses unified host.
	os.Setenv("DATABRICKS_HOST", unifiedHost)
	os.Setenv("DATABRICKS_EXPERIMENTAL_IS_UNIFIED_HOST", "true")
}

func TestAccUnifiedHostCreateJobsAWS(t *testing.T) {
	initUnifiedHostEnv(t)
	if !IsAws(t) {
		Skipf(t)("This test is only running on AWS")
	}
	run(t, []Step{
		{
			Template: `
			resource "databricks_mws_workspaces" "ws1" {
				account_id     = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name = "tf-unified-{var.RANDOM}-1"
				aws_region     = "{env.AWS_REGION}"
				compute_mode   = "SERVERLESS"
			}

			resource "databricks_mws_workspaces" "ws2" {
				account_id     = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name = "tf-unified-{var.RANDOM}-2"
				aws_region     = "{env.AWS_REGION}"
				compute_mode   = "SERVERLESS"
			}

			resource "databricks_job" "ws1" {
				name = "tf-unified-{var.RANDOM}-job-1"
				provider_config {
					workspace_id = databricks_mws_workspaces.ws1.workspace_id
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

			resource "databricks_job" "ws2" {
				name = "tf-unified-{var.RANDOM}-job-2"
				provider_config {
					workspace_id = databricks_mws_workspaces.ws2.workspace_id
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
		},
	})
}
