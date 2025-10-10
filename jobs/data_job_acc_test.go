package jobs_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const jobDataSourceTemplate = `
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
`

func TestAccDataSourceJob(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: jobDataSourceTemplate + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
		}`,
	}, acceptance.Step{
		Template: jobDataSourceTemplate + `
		data "databricks_job" "this" {
			job_name = databricks_job.this.name
			provider_config {
				workspace_id = ""
			}
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_job.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			policy := r.Primary.Attributes["provider_config.0.workspace_id"]
			if policy == "" {
				return fmt.Errorf("Provider Config Workspace ID is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
