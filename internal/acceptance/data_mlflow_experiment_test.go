package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceMlflowExperiment(t *testing.T) {
	workspaceLevel(t,
		step{
			Template: `
				data "databricks_current_user" "me" {}

				resource "databricks_mlflow_experiment" "this" {
				  name              = "${data.databricks_current_user.me.home}/experiment-{var.RANDOM}"
				  artifact_location = "dbfs:/tmp/my-experiment"
				  description       = "My MLflow experiment description"
				}`,
		},
		step{
			Template: `
				data "databricks_current_user" "me" {}

				resource "databricks_mlflow_experiment" "this" {
				  name              = "${data.databricks_current_user.me.home}/experiment-{var.RANDOM}"
				  artifact_location = "dbfs:/tmp/my-experiment"
				  description       = "My MLflow experiment description"
				}

				data "databricks_mlflow_experiment" "by_name" {
				  depends_on = [databricks_mlflow_experiment.this]
				  name       = "${data.databricks_current_user.me.home}/experiment-{var.RANDOM}"
				}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mlflow_experiment.by_name"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				id := r.Primary.Attributes["id"]
				if id == "" {
					return fmt.Errorf("id is empty: %v", r.Primary.Attributes)
				}
				expect := "dbfs:/tmp/my-experiment"
				loc := r.Primary.Attributes["artifact_location"]
				if loc != expect {
					return fmt.Errorf("incorrect artifact location. expected: %v, received: %v",
						expect, loc)
				}
				return nil
			},
		},
		step{
			Template: `
				data "databricks_current_user" "me" {}

				resource "databricks_mlflow_experiment" "this" {
				  name              = "${data.databricks_current_user.me.home}/experiment-{var.RANDOM}"
				  artifact_location = "dbfs:/tmp/my-experiment"
				  description       = "My MLflow experiment description"
				}

				data "databricks_mlflow_experiment" "by_id" {
				  depends_on = [databricks_mlflow_experiment.this]
				  experiment_id = "${databricks_mlflow_experiment.this.experiment_id}"
				}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mlflow_experiment.by_id"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				id := r.Primary.Attributes["id"]
				if id == "" {
					return fmt.Errorf("id is empty: %v", r.Primary.Attributes)
				}
				expect := "dbfs:/tmp/my-experiment"
				loc := r.Primary.Attributes["artifact_location"]
				if loc != expect {
					return fmt.Errorf("incorrect artifact location. expected: %v, received: %v",
						expect, loc)
				}
				return nil
			},
		})
}
