package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccModelServing(t *testing.T) {
	name := fmt.Sprintf("terraform-test-model-serving-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
		data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "singlenode-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
			num_workers = 0
			autotermination_minutes = 10
			spark_conf = {
				"spark.databricks.cluster.profile" = "singleNode"
				"spark.master" = "local[*]"
			}
			custom_tags = {
				"ResourceClass" = "SingleNode"
			}
			library {
				pypi {
					package = "mlflow"
				}
			}
		}
		resource "databricks_mlflow_experiment" "exp" {
			name = "/Shared/%[1]s-exp"
		}
		resource "databricks_mlflow_model" "model" {
			name = "%[1]s-model"
		}
		`, name),
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())
			id := s.RootModule().Resources["databricks_cluster.this"].Primary.ID
			w.CommandExecutor.Execute(context.Background(), id, "python", fmt.Sprintf(`
				import time
				import mlflow
				import mlflow.pyfunc
				from mlflow.tracking.artifact_utils import get_artifact_uri
				from mlflow.tracking.client import MlflowClient

				mlflow.set_experiment("/Shared/%[1]s-exp")

				class SampleModel(mlflow.pyfunc.PythonModel):
					def predict(self, ctx, input_df):
						return 7
				artifact_path = 'sample_model'
				
				with mlflow.start_run() as new_run:
					mlflow.pyfunc.log_model(python_model=SampleModel(), artifact_path=artifact_path)
					run1_id = new_run.info.run_id
					source = get_artifact_uri(run_id=run1_id, artifact_path=artifact_path)

				client = MlflowClient()
				client.create_model_version(name="%[1]s-model", source=source, run_id=run1_id)
				client.create_model_version(name="%[1]s-model", source=source, run_id=run1_id)
				while client.get_model_version(name="%[1]s-model", version="1").getStatus() != ModelRegistry.ModelVersionStatus.READY:
					time.sleep(10)
				while client.get_model_version(name="%[1]s-model", version="2").getStatus() != ModelRegistry.ModelVersionStatus.READY:
					time.sleep(10)
			`, name))
			return nil
		},
	},
		step{
			Template: fmt.Sprintf(`
			resource "databricks_mlflow_experiment" "exp" {
				name = "/Shared/%[1]s-exp"
			}
			resource "databricks_mlflow_model" "model" {
				name = "%[1]s-model"
			}
			resource "databricks_model_serving" "endpoint" {
				name = "%[1]s"
				config {
					served_models {
						name = "prod_model"
						model_name = "%[1]s-model"
						model_version = "1"
						workload_size = "Small"
						scale_to_zero_enabled = true
					}
					served_models {
						name = "candidate_model"
						model_name = "%[1]s-model"
						model_version = "2"
						workload_size = "Small"
						scale_to_zero_enabled = false
					}
					traffic_config {
						routes {
							served_model_name = "prod_model"
							traffic_percentage = 90
						}
						routes {
							served_model_name = "candidate_model"
							traffic_percentage = 10
						}
					}
				}
			}
		`, name),
		},
		step{
			Template: fmt.Sprintf(`
			resource "databricks_mlflow_experiment" "exp" {
				name = "/Shared/%[1]s-exp"
			}
			resource "databricks_mlflow_model" "model" {
				name = "%[1]s-model"
			}
			resource "databricks_model_serving" "endpoint" {
				name = "%[1]s"
				config {
					served_models {
						name = "prod_model"
						model_name = "%[1]s-model"
						model_version = "1"
						workload_size = "Small"
						scale_to_zero_enabled = true
					}
					traffic_config {
						routes {
							served_model_name = "prod_model"
							traffic_percentage = 100
						}
					}
				}
			}
		`, name),
		},
	)
}
