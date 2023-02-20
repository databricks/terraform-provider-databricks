package acceptance

import (
	"testing"
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
