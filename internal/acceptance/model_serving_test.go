package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccModelServing(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_models {
						name = "prod_model"
						model_name = "experiment-fixture-model"
						model_version = "1"
						workload_size = "Small"
						scale_to_zero_enabled = true
					}
					served_models {
						name = "candidate_model"
						model_name = "experiment-fixture-model"
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
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_models {
						name = "prod_model"
						model_name = "experiment-fixture-model"
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

func TestUcAccModelServingProvisionedThroughput(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-pt-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	unityWorkspaceLevel(t, step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities{
						name = "pt_model"
						entity_name = "system.ai.mistral_7b_instruct_v0_1"
						entity_version = "1"
						min_provisioned_throughput = 0
						max_provisioned_throughput = 970
					}
					traffic_config {
						routes {
							served_model_name = "pt_model"
							traffic_percentage = 100
						}
					}
				}
			}
		`, name),
	},
	)
}
