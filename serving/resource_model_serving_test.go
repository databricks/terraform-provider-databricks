package serving

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/serving"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestModelServingReadWithSensitivePlaintextFields(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockServingEndpointsAPI().EXPECT().
				GetByName(mock.Anything, "test-endpoint").
				Return(&serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					EndpointUrl: "https://example.com/endpoint",
					Config: &serving.EndpointCoreConfigOutput{
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name: "gpt-4o-mini",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4o-mini",
									Provider: serving.ExternalModelProviderOpenai,
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										// API doesn't return plaintext fields
										OpenaiApiBase: "https://api.openai.com/v1",
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedEntityName:  "gpt-4o-mini",
									TrafficPercentage: 100,
								},
							},
						},
					},
				}, nil)
		},
		Resource: ResourceModelServing(),
		Read:     true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
			"config.0.served_entities.0.external_model.0.openai_config.0.openai_api_key_plaintext": "sk-test-key-12345",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"serving_endpoint_id": "test-endpoint",
		"endpoint_url":        "https://example.com/endpoint",
		// Verify the plaintext field is preserved from state
		"config.0.served_entities.0.external_model.0.openai_config.0.openai_api_key_plaintext": "sk-test-key-12345",
	})
}

func TestModelServingReadWithMultipleSensitiveFields(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockServingEndpointsAPI().EXPECT().
				GetByName(mock.Anything, "test-endpoint").
				Return(&serving.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
					EndpointUrl: "https://example.com/endpoint",
					Config: &serving.EndpointCoreConfigOutput{
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name: "gpt-4o-mini",
								ExternalModel: &serving.ExternalModel{
									Name:     "gpt-4o-mini",
									Provider: serving.ExternalModelProviderOpenai,
									Task:     "llm/v1/chat",
									OpenaiConfig: &serving.OpenAiConfig{
										// API doesn't return plaintext fields
										OpenaiApiBase:          "https://api.openai.com/v1",
										OpenaiApiType:          "azuread",
										MicrosoftEntraClientId: "client-id-123",
										MicrosoftEntraTenantId: "tenant-id-456",
									},
								},
							},
							{
								Name: "vertex-ai-model",
								ExternalModel: &serving.ExternalModel{
									Name:     "vertex-ai-model",
									Provider: serving.ExternalModelProviderGoogleCloudVertexAi,
									Task:     "llm/v1/chat",
									GoogleCloudVertexAiConfig: &serving.GoogleCloudVertexAiConfig{
										// API doesn't return plaintext fields
										ProjectId: "my-gcp-project",
										Region:    "us-central1",
									},
								},
							},
						},
						TrafficConfig: &serving.TrafficConfig{
							Routes: []serving.Route{
								{
									ServedEntityName:  "gpt-4o-mini",
									TrafficPercentage: 50,
								},
								{
									ServedEntityName:  "vertex-ai-model",
									TrafficPercentage: 50,
								},
							},
						},
					},
				}, nil)
		},
		Resource: ResourceModelServing(),
		Read:     true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
			"config.0.served_entities.0.external_model.0.openai_config.0.openai_api_key_plaintext":                "sk-test-key-12345",
			"config.0.served_entities.0.external_model.0.openai_config.0.microsoft_entra_client_secret_plaintext": "client-secret-xyz",
			"config.0.served_entities.1.external_model.0.google_cloud_vertex_ai_config.0.private_key_plaintext":   "-----BEGIN PRIVATE KEY-----\ntest-key\n-----END PRIVATE KEY-----", // gitleaks:allow
		},
	}.ApplyAndExpectData(t, map[string]any{
		"serving_endpoint_id":        "test-endpoint",
		"endpoint_url":               "https://example.com/endpoint",
		"config.0.served_entities.#": 2,
		// Verify all plaintext fields are preserved from state
		"config.0.served_entities.0.external_model.0.openai_config.0.openai_api_key_plaintext":                "sk-test-key-12345",
		"config.0.served_entities.0.external_model.0.openai_config.0.microsoft_entra_client_secret_plaintext": "client-secret-xyz",
		"config.0.served_entities.1.external_model.0.google_cloud_vertex_ai_config.0.private_key_plaintext":   "-----BEGIN PRIVATE KEY-----\ntest-key\n-----END PRIVATE KEY-----",
	})
}

// TestCopySensitiveFields tests the reflection-based sensitive field copying logic
func TestCopySensitiveFields(t *testing.T) {
	// Test case 1: Simple plaintext field copy
	t.Run("SimpleOpenAIPlaintextCopy", func(t *testing.T) {
		src := &serving.ServingEndpointDetailed{
			Name: "test-endpoint",
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						Name: "gpt-4o-mini",
						ExternalModel: &serving.ExternalModel{
							Name:     "gpt-4o-mini",
							Provider: serving.ExternalModelProviderOpenai,
							Task:     "llm/v1/chat",
							OpenaiConfig: &serving.OpenAiConfig{
								OpenaiApiKeyPlaintext: "sk-test-key-12345",
								OpenaiApiBase:         "https://api.openai.com/v1",
							},
						},
					},
				},
			},
		}

		dst := &serving.ServingEndpointDetailed{
			Name: "test-endpoint",
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						Name: "gpt-4o-mini",
						ExternalModel: &serving.ExternalModel{
							Name:     "gpt-4o-mini",
							Provider: serving.ExternalModelProviderOpenai,
							Task:     "llm/v1/chat",
							OpenaiConfig: &serving.OpenAiConfig{
								// API doesn't return plaintext field
								OpenaiApiBase: "https://api.openai.com/v1",
							},
						},
					},
				},
			},
		}

		copySensitiveExternalModelFields(src, dst)

		assert.Equal(t, "sk-test-key-12345", dst.Config.ServedEntities[0].ExternalModel.OpenaiConfig.OpenaiApiKeyPlaintext,
			"OpenAI plaintext API key should be copied from source")
	})

	// Test case 2: Multiple plaintext fields
	t.Run("MultipleProviderPlaintextCopy", func(t *testing.T) {
		src := &serving.ServingEndpointDetailed{
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						ExternalModel: &serving.ExternalModel{
							OpenaiConfig: &serving.OpenAiConfig{
								OpenaiApiKeyPlaintext:               "sk-test-key",
								MicrosoftEntraClientSecretPlaintext: "client-secret",
							},
						},
					},
					{
						ExternalModel: &serving.ExternalModel{
							GoogleCloudVertexAiConfig: &serving.GoogleCloudVertexAiConfig{
								PrivateKeyPlaintext: "private-key-content",
							},
						},
					},
				},
			},
		}

		dst := &serving.ServingEndpointDetailed{
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						ExternalModel: &serving.ExternalModel{
							OpenaiConfig: &serving.OpenAiConfig{},
						},
					},
					{
						ExternalModel: &serving.ExternalModel{
							GoogleCloudVertexAiConfig: &serving.GoogleCloudVertexAiConfig{},
						},
					},
				},
			},
		}

		copySensitiveExternalModelFields(src, dst)

		assert.Equal(t, "sk-test-key", dst.Config.ServedEntities[0].ExternalModel.OpenaiConfig.OpenaiApiKeyPlaintext)
		assert.Equal(t, "client-secret", dst.Config.ServedEntities[0].ExternalModel.OpenaiConfig.MicrosoftEntraClientSecretPlaintext)
		assert.Equal(t, "private-key-content", dst.Config.ServedEntities[1].ExternalModel.GoogleCloudVertexAiConfig.PrivateKeyPlaintext)
	})

	// Test case 3: Nil safety
	t.Run("NilSafety", func(t *testing.T) {
		// Should not panic
		copySensitiveExternalModelFields(nil, nil)

		src := &serving.ServingEndpointDetailed{}
		copySensitiveExternalModelFields(src, nil)
		copySensitiveExternalModelFields(nil, src)

		dst := &serving.ServingEndpointDetailed{}
		copySensitiveExternalModelFields(src, dst) // Both empty, should not panic
	})

	// Test case 4: Don't overwrite existing values
	t.Run("DontOverwriteExisting", func(t *testing.T) {
		src := &serving.ServingEndpointDetailed{
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						ExternalModel: &serving.ExternalModel{
							OpenaiConfig: &serving.OpenAiConfig{
								OpenaiApiKeyPlaintext: "src-key",
							},
						},
					},
				},
			},
		}

		dst := &serving.ServingEndpointDetailed{
			Config: &serving.EndpointCoreConfigOutput{
				ServedEntities: []serving.ServedEntityOutput{
					{
						ExternalModel: &serving.ExternalModel{
							OpenaiConfig: &serving.OpenAiConfig{
								OpenaiApiKeyPlaintext: "dst-key-already-set",
							},
						},
					},
				},
			},
		}

		copySensitiveExternalModelFields(src, dst)

		// Should not overwrite existing value
		assert.Equal(t, "dst-key-already-set", dst.Config.ServedEntities[0].ExternalModel.OpenaiConfig.OpenaiApiKeyPlaintext)
	})

	// Test case 5: Test with reflection directly
	t.Run("ReflectionBasedCopy", func(t *testing.T) {
		type TestStruct struct {
			RegularField       string
			SensitivePlaintext string
		}

		src := TestStruct{
			RegularField:       "regular",
			SensitivePlaintext: "sensitive-value",
		}

		dst := TestStruct{
			RegularField: "regular",
		}

		srcVal := reflect.ValueOf(&src)
		dstVal := reflect.ValueOf(&dst)

		copySensitiveFields(srcVal, dstVal)

		assert.Equal(t, "sensitive-value", dst.SensitivePlaintext)
	})
}
