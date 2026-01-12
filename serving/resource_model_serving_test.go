package serving

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/databricks/databricks-sdk-go/service/serving"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestModelServingCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceModelServing())
}

func TestModelServingCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: serving.CreateServingEndpoint{
					Name: "test-endpoint",
					Config: &serving.EndpointCoreConfigInput{
						ServedModels: []serving.ServedModelInput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								WorkloadSize:       "Small",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								WorkloadSize:       "Small",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &serving.EndpointCoreConfigOutput{
						ServedModels: []serving.ServedModelOutput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-endpoint"
			config {
				served_models {
					name = "prod_model"
					model_name = "ads1"
					model_version = "2"
					workload_size = "Small"
					scale_to_zero_enabled = true
				}
				served_models {
					name = "candidate_model"
					model_name = "ads1"
					model_version = "4"
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
			`,
		Create: true,
	}.ApplyNoError(t)
}

func TestModelServingCreateGPU(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: serving.CreateServingEndpoint{
					Name: "test-endpoint",
					Config: &serving.EndpointCoreConfigInput{
						ServedModels: []serving.ServedModelInput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								WorkloadSize:       "Small",
								WorkloadType:       "GPU_MEDIUM",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								WorkloadSize:       "Small",
								WorkloadType:       "GPU_MEDIUM",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &serving.EndpointCoreConfigOutput{
						ServedModels: []serving.ServedModelOutput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-endpoint"
			config {
				served_models {
					name = "prod_model"
					model_name = "ads1"
					model_version = "2"
					workload_size = "Small"
					workload_type = "GPU_MEDIUM"
					scale_to_zero_enabled = true
				}
				served_models {
					name = "candidate_model"
					model_name = "ads1"
					model_version = "4"
					workload_size = "Small"
					workload_type = "GPU_MEDIUM"
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
			`,
		Create: true,
	}.ApplyNoError(t)
}

func TestModelServingCreateAzureOpenAI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: serving.CreateServingEndpoint{
					Name: "test-azure-endpoint",
					Config: &serving.EndpointCoreConfigInput{
						ServedEntities: []serving.ServedEntityInput{
							{
								Name: "azure-model",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4",
									Provider: "openai",
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										OpenaiApiBase:        "https://example.openai.azure.com",
										OpenaiApiVersion:     "2023-05-15",
										OpenaiDeploymentName: "deployment-name",
										OpenaiApiKey:         "secret-key",
										OpenaiApiType:        "azure",
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "azure-model",
									TrafficPercentage: 100,
								},
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-azure-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-azure-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Name: "test-azure-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-azure-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-azure-endpoint",
					Name: "test-azure-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &serving.EndpointCoreConfigOutput{
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name: "azure-model",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4",
									Provider: "openai",
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										OpenaiApiBase:        "https://example.openai.azure.com",
										OpenaiApiVersion:     "2023-05-15",
										OpenaiDeploymentName: "deployment-name",
										OpenaiApiType:        "azure",
										// API typically doesn't return the key
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "azure-model",
									TrafficPercentage: 100,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
						azure_openai_config {
							openai_api_base = "https://example.openai.azure.com"
							openai_api_version = "2023-05-15"
							openai_deployment_name = "deployment-name"
							openai_api_key = "secret-key"
						}
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ApplyNoError(t)
}

func TestModelServingCreateAzureOpenAI_MissingConfig(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ExpectError(t, "azure_openai_config must be specified when using provider 'azure-openai'")
}

func TestModelServingCreateAzureOpenAI_MissingAuthentication(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
						azure_openai_config {
							openai_api_base = "https://example.openai.azure.com"
							openai_api_version = "2023-05-15"
							openai_deployment_name = "deployment-name"
						}
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ExpectError(t, "azure_openai_config requires either API key authentication (openai_api_key or openai_api_key_plaintext) or Microsoft Entra authentication (microsoft_entra_client_id, microsoft_entra_tenant_id, and microsoft_entra_client_secret/microsoft_entra_client_secret_plaintext)")
}

func TestModelServingCreateAzureOpenAI_ConflictingAuthMethods(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
						azure_openai_config {
							openai_api_base = "https://example.openai.azure.com"
							openai_api_version = "2023-05-15"
							openai_deployment_name = "deployment-name"
							openai_api_key = "secret-key"
							microsoft_entra_client_id = "client-id"
							microsoft_entra_tenant_id = "tenant-id"
							microsoft_entra_client_secret = "client-secret"
						}
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ExpectError(t, "cannot specify both API key authentication (openai_api_key/openai_api_key_plaintext) and Microsoft Entra authentication (microsoft_entra_*) fields")
}

func TestModelServingCreateAzureOpenAI_EntraAuth(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: serving.CreateServingEndpoint{
					Name: "test-azure-endpoint",
					Config: &serving.EndpointCoreConfigInput{
						ServedEntities: []serving.ServedEntityInput{
							{
								Name: "azure-model",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4",
									Provider: "openai",
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										OpenaiApiBase:              "https://example.openai.azure.com",
										OpenaiApiVersion:           "2023-05-15",
										OpenaiDeploymentName:       "deployment-name",
										OpenaiApiType:              "azuread",
										MicrosoftEntraClientId:     "client-id",
										MicrosoftEntraTenantId:     "tenant-id",
										MicrosoftEntraClientSecret: "{{secrets/scope/client-secret}}",
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "azure-model",
									TrafficPercentage: 100,
								},
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-azure-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-azure-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Name: "test-azure-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-azure-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-azure-endpoint",
					Name: "test-azure-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &serving.EndpointCoreConfigOutput{
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name: "azure-model",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4",
									Provider: "openai",
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										OpenaiApiBase:              "https://example.openai.azure.com",
										OpenaiApiVersion:           "2023-05-15",
										OpenaiDeploymentName:       "deployment-name",
										OpenaiApiType:              "azuread",
										MicrosoftEntraClientId:     "client-id",
										MicrosoftEntraTenantId:     "tenant-id",
										MicrosoftEntraClientSecret: "{{secrets/scope/client-secret}}",
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "azure-model",
									TrafficPercentage: 100,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
						azure_openai_config {
							openai_api_base = "https://example.openai.azure.com"
							openai_api_version = "2023-05-15"
							openai_deployment_name = "deployment-name"
							microsoft_entra_client_id = "client-id"
							microsoft_entra_tenant_id = "tenant-id"
							microsoft_entra_client_secret = "{{secrets/scope/client-secret}}"
						}
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ApplyNoError(t)
}

func TestModelServingCreateAzureOpenAI_EntraMissingClientId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-azure-endpoint"
			config {
				served_entities {
					name = "azure-model"
					external_model {
						name = "gpt-4"
						provider = "azure-openai"
						task = "llm/v1/chat"
						azure_openai_config {
							openai_api_base = "https://example.openai.azure.com"
							openai_api_version = "2023-05-15"
							openai_deployment_name = "deployment-name"
							microsoft_entra_tenant_id = "tenant-id"
							microsoft_entra_client_secret = "client-secret"
						}
					}
				}
				traffic_config {
					routes {
						served_model_name = "azure-model"
						traffic_percentage = 100
					}
				}
			}
			`,
		Create: true,
	}.ExpectError(t, "microsoft_entra_client_id is required when using Microsoft Entra authentication")
}

func TestModelServingCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServing(),
		Create:   true,
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingCreate_WithErrorOnWait(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: serving.CreateServingEndpoint{
					Name: "test-endpoint",
					Config: &serving.EndpointCoreConfigInput{
						ServedModels: []serving.ServedModelInput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								WorkloadSize:       "Small",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								WorkloadSize:       "Small",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
			},
		},
		Resource: ResourceModelServing(),
		HCL: `
			name = "test-endpoint"
			config {
				served_models {
					name = "prod_model"
					model_name = "ads1"
					model_version = "2"
					workload_size = "Small"
					scale_to_zero_enabled = true
				}
				served_models {
					name = "candidate_model"
					model_name = "ads1"
					model_version = "4"
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
			`,
		Create: true,
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					EndpointUrl: "https://example.com/endpoint",
					Config: &serving.EndpointCoreConfigOutput{
						ServedModels: []serving.ServedModelOutput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								ModelName:          "ads1",
								ModelVersion:       "4",
								ScaleToZeroEnabled: false,
							},
						},
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name:               "prod_model",
								EntityName:         "ads1",
								EntityVersion:      "2",
								ScaleToZeroEnabled: true,
							},
							{
								Name:               "candidate_model",
								EntityName:         "ads1",
								EntityVersion:      "4",
								ScaleToZeroEnabled: false,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									ServedEntityName:  "prod_model", // Server returns both fields with same value
									TrafficPercentage: 90,
								},
								{
									ServedModelName:   "candidate_model",
									ServedEntityName:  "candidate_model", // Server returns both fields with same value
									TrafficPercentage: 10,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		Read:     true,
		ID:       "test-endpoint",
	}.ApplyAndExpectData(t, map[string]any{
		"serving_endpoint_id":                                   "test-endpoint",
		"endpoint_url":                                          "https://example.com/endpoint",
		"config.0.served_entities.#":                            2,
		"config.0.served_entities.0.name":                       "prod_model",
		"config.0.served_entities.1.name":                       "candidate_model",
		"config.0.traffic_config.#":                             1,
		"config.0.traffic_config.0.routes.#":                    2,
		"config.0.traffic_config.0.routes.0.served_model_name":  "prod_model",
		"config.0.traffic_config.0.routes.1.served_model_name":  "candidate_model",
		"config.0.traffic_config.0.routes.0.served_entity_name": "prod_model",
		"config.0.traffic_config.0.routes.1.served_entity_name": "candidate_model",
		"config.0.traffic_config.0.routes.0.traffic_percentage": 90,
		"config.0.traffic_config.0.routes.1.traffic_percentage": 10,
	})
}

func TestModelServingReadEmptyConfig(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: map[string]any{
					"creation_timestamp":           1743085336000,
					"creator":                      "b76b6808-9e10-43b3-be20-6b6d19ed1af0",
					"creator_display_name":         "DECO-TF-AWS-PROD-IS-SPN",
					"creator_kind":                 "ServicePrincipal",
					"id":                           "84f4b90597b94fb1846a96cb505772f1",
					"last_updated_timestamp":       1743085336000,
					"name":                         "test-endpoint-462f54a7-fefd-4d48-bdc2-2659a5439d94",
					"permission_level":             "CAN_MANAGE",
					"resource_credential_strategy": "EMBEDDED_CREDENTIALS",
					"route_optimized":              false,
					"state": map[string]any{
						"config_update": "NOT_UPDATING",
						"ready":         "NOT_READY",
						"suspend":       "NOT_SUSPENDED",
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		Read:     true,
		ID:       "test-endpoint",
	}.ApplyNoError(t)
}

func TestModelServingRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServing(),
		Read:     true,
		ID:       "test-endpoint",
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/config",
				ExpectedRequest: serving.EndpointCoreConfigInput{
					Name: "test-endpoint",
					ServedModels: []serving.ServedModelInput{
						{
							Name:               "prod_model",
							ModelName:          "ads1",
							ModelVersion:       "2",
							WorkloadSize:       "Small",
							ScaleToZeroEnabled: true,
						},
					},
					TrafficConfig: &serving.TrafficConfig{
						Routes: []serving.Route{
							{
								ServedModelName:   "prod_model",
								TrafficPercentage: 100,
							},
						},
					},
				},
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &serving.EndpointCoreConfigOutput{
						ServedModels: []serving.ServedModelOutput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								ScaleToZeroEnabled: true,
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedModelName:   "prod_model",
									TrafficPercentage: 100,
								},
							},
						},
					},
				},
			},
		},
		Resource: ResourceModelServing(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_models {
					name = "prod_model"
					model_name = "ads1"
					model_version = "2"
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
			`,
	}.ApplyNoError(t)
}

func TestModelServingUpdate_RemoveConfigIsNoOp(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceModelServing(),
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name":                          "test-endpoint",
			"config.#":                      "1",
			"config.0.served_models.#":      "1",
			"config.0.served_models.0.name": "prod_model",
			"serving_endpoint_id":           "id",
			"endpoint_url":                  "https://example.com/endpoint",
		},
		HCL: `
			name = "test-endpoint"
			`,
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{},
	}.ApplyNoError(t)
}

func TestModelServingUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/config",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServing(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_models {
					name = "prod_model"
					model_name = "ads1"
					model_version = "2"
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
			`,
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: "",
			},
		},
		Resource: ResourceModelServing(),
		Delete:   true,
		ID:       "test-endpoint",
	}.ApplyNoError(t)
}

func TestModelServingDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServing(),
		Delete:   true,
		ID:       "test-endpoint",
	}.ExpectError(t, "Internal error happened")
}
