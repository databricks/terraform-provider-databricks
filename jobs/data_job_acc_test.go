package jobs_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceJob(t *testing.T) {
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
			name = "job-datasource-acceptance-test"

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

		}

		data "databricks_job" "this" {
			job_name = databricks_job.this.name
		}`,
	})
}

func TestAccDataSourceJob_InvalidID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
				workspace_id = "invalid"
			}
		}`,
		ExpectError: regexp.MustCompile(`failed to parse workspace_id.*invalid syntax`),
	})
}

func TestAccDataSourceJob_MismatchedID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
				workspace_id = "123"
			}
		}`,
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccDataSourceJob_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
				workspace_id = ""
			}
		}`,
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
	})
}

func TestAccDataSourceJob_EmptyBlock(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
			}
		}`,
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
	})
}

func TestAccDataSourceJob_Match(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
		}`,
	}, acceptance.Step{
		Template: jobClusterTemplate("") + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
				workspace_id = "1142582526922259"
			}
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_job.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			id := r.Primary.Attributes["provider_config.0.workspace_id"]
			if id != "1142582526922259" {
				return fmt.Errorf("wrong workspace_id found: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
