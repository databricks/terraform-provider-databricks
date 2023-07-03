package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccModelServing(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "aws", "azure":
	default:
		t.Skipf("not available on %s", cloudEnv)
	}

	clusterID := os.Getenv("TEST_DEFAULT_CLUSTER_ID")
	if clusterID == "" {
		t.Skipf("default cluster not available")
	}
	// data "databricks_spark_version" "latest" {
	// }
	// resource "databricks_cluster" "this" {
	// 	cluster_name = "singlenode-{var.RANDOM}"
	// 	spark_version = data.databricks_spark_version.latest.id
	// 	instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
	// 	num_workers = 0
	// 	autotermination_minutes = 10
	// 	spark_conf = {
	// 		"spark.databricks.cluster.profile" = "singleNode"
	// 		"spark.master" = "local[*]"
	// 	}
	// 	custom_tags = {
	// 		"ResourceClass" = "SingleNode"
	// 	}
	// 	library {
	// 		pypi {
	// 			package = "mlflow"
	// 		}
	// 	}
	// }

	name := fmt.Sprintf("terraform-test-model-serving-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_mlflow_experiment" "exp" {
			name = "/Shared/%[1]s-exp"
		}
		resource "databricks_mlflow_model" "model" {
			name = "%[1]s-model"
		}
		`, name),
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())
			ctx := context.Background()
			executor, err := w.CommandExecution.Start(ctx, clusterID, compute.LanguagePython)
			if err != nil {
				return err
			}
			defer executor.Destroy(ctx)
			results, err := executor.Execute(ctx, fmt.Sprintf(`
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
			if err != nil {
				return err
			}
			return results.Err()
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

			resource "databricks_permissions" "ml_serving_usage" {
				serving_endpoint_id = databricks_model_serving.endpoint.serving_endpoint_id
			  
				access_control {
				  group_name       = "users"
				  permission_level = "CAN_VIEW"
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
