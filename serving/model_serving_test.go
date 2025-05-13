package serving_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccModelServing(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	acceptance.WorkspaceLevel(t, acceptance.Step{
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

			data "databricks_serving_endpoints" "all" {}

			resource "databricks_permissions" "ml_serving_usage" {
				serving_endpoint_id = databricks_model_serving.endpoint.serving_endpoint_id
			  
				access_control {
				  group_name       = "users"
				  permission_level = "CAN_VIEW"
				}
			}
		`, name),
	},
		acceptance.Step{
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
			data "databricks_serving_endpoints" "all" {}
		`, name),
		},
	)
}

func TestUcAccModelServingProvisionedThroughput(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-pt-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities{
						name = "pt_model"
						entity_name = "system.ai.mistral_7b_instruct_v0_2"
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
	}, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities{
						name = "pt_model"
						entity_name = "system.ai.mistral_7b_instruct_v0_2"
						entity_version = "1"
						min_provisioned_throughput = 970
						max_provisioned_throughput = 1940
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
	}, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities{
						name = "pt_model"
						entity_name = "system.ai.mistral_7b_instruct_v0_2"
						entity_version = "1"
						min_provisioned_throughput = 0
						max_provisioned_throughput = 1940
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

func TestAccModelServingExternalModel(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-em-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	scope_name := fmt.Sprintf("terraform-test-secret-scope-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_secret_scope" "scope" {
				name = "%s"
			}

			resource "databricks_secret" "key" {
				key          = "api_key"
				string_value = "fake-secret"
				scope        = databricks_secret_scope.scope.id
			}

			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities {
						name = "prod_model"
						external_model {
							provider = "anthropic"
							name = "claude-2.0"
							task = "llm/v1/chat"
							anthropic_config {
								anthropic_api_key = databricks_secret.key.config_reference
							}
						}
					}
					traffic_config {
						routes {
							served_model_name = "prod_model"
							traffic_percentage = 100
						}
					}
				}
			}
		`, scope_name, name),
	},
		acceptance.Step{
			Template: fmt.Sprintf(`
			resource "databricks_secret_scope" "scope" {
				name = "%s"
			}

			resource "databricks_secret" "key" {
				key          = "api_key"
				string_value = "fake-secret"
				scope        = databricks_secret_scope.scope.id
			}

			resource "databricks_model_serving" "endpoint" {
				name = "%s"
				config {
					served_entities {
						name = "prod_model"
						external_model {
							provider = "openai"
							name = "gpt-4o"
							task = "llm/v1/chat"
							openai_config {
								openai_api_key = databricks_secret.key.config_reference
							}
						}
					}
					traffic_config {
						routes {
							served_model_name = "prod_model"
							traffic_percentage = 100
						}
					}
				}
			}
		`, scope_name, name),
		},
	)
}

func TestUcAccModelServingProvisionedThroughputResource(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-model-serving-pt-%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_model_serving_provisioned_throughput" "endpoint" {
				name = "%s"
				config {
					served_entities {
						name = "prod_model"
						entity_name = "system.ai.llama-4-maverick"
						entity_version = "1"
						provisioned_model_units = 54
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
		acceptance.Step{
			Template: fmt.Sprintf(`
			resource "databricks_model_serving_provisioned_throughput" "endpoint" {
				name = "%s"
				config {
					served_entities {
						name = "prod_model"
						entity_name = "system.ai.llama-4-maverick"
						entity_version = "1"
						provisioned_model_units = 100
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
