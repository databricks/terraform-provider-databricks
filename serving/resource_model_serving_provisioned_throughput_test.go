package serving

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestModelServingProvisionedThroughputCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceModelServingProvisionedThroughput())
}

func TestModelServingProvisionedThroughputCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints/pt",
				ExpectedRequest: serving.CreatePtEndpointRequest{
					Name: "test-endpoint",
					Config: serving.PtEndpointCoreConfig{
						ServedEntities: []serving.PtServedModel{
							{
								Name:                  "prod_model",
								EntityName:            "ads1",
								EntityVersion:         "2",
								ProvisionedModelUnits: 50,
							},
							{
								Name:                  "candidate_model",
								EntityName:            "ads1",
								EntityVersion:         "4",
								ProvisionedModelUnits: 50,
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
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name:          "prod_model",
								EntityName:    "ads1",
								EntityVersion: "2",
							},
							{
								Name:          "candidate_model",
								EntityName:    "ads1",
								EntityVersion: "4",
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
		Resource: ResourceModelServingProvisionedThroughput(),
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
				}
				served_entities {
					name = "candidate_model"
					entity_name = "ads1"
					entity_version = "4"
					provisioned_model_units = 50
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

func TestModelServingProvisionedThroughputCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints/pt",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Create:   true,
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingProvisionedThroughputCreate_WithErrorOnWait(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints/pt",
				ExpectedRequest: serving.CreatePtEndpointRequest{
					Name: "test-endpoint",
					Config: serving.PtEndpointCoreConfig{
						ServedEntities: []serving.PtServedModel{
							{
								Name:                  "prod_model",
								EntityName:            "ads1",
								EntityVersion:         "2",
								ProvisionedModelUnits: 50,
							},
							{
								Name:                  "candidate_model",
								EntityName:            "ads1",
								EntityVersion:         "4",
								ProvisionedModelUnits: 50,
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
				Response: common.APIErrorBody{
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
		Resource: ResourceModelServingProvisionedThroughput(),
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
				}
				served_entities {
					name = "candidate_model"
					entity_name = "ads1"
					entity_version = "4"
					provisioned_model_units = 50
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

func TestModelServingProvisionedThroughputRead(t *testing.T) {
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
					Config: &serving.EndpointCoreConfigOutput{
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name:                  "prod_model",
								EntityName:            "ads1",
								EntityVersion:         "2",
								ProvisionedModelUnits: 50,
							},
							{
								Name:                  "candidate_model",
								EntityName:            "ads1",
								EntityVersion:         "4",
								ProvisionedModelUnits: 50,
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
		Resource: ResourceModelServingProvisionedThroughput(),
		Read:     true,
		ID:       "test-endpoint",
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputReadEmptyConfig(t *testing.T) {
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
		Resource: ResourceModelServingProvisionedThroughput(),
		Read:     true,
		ID:       "test-endpoint",
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Read:     true,
		ID:       "test-endpoint",
	}.ExpectError(t, "Internal error happened")
}

func TestModelServingProvisionedThroughputUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/pt/test-endpoint/config",
				ExpectedRequest: serving.UpdateProvisionedThroughputEndpointConfigRequest{
					Name: "test-endpoint",
					Config: serving.PtEndpointCoreConfig{
						ServedEntities: []serving.PtServedModel{
							{
								Name:                  "prod_model",
								EntityName:            "ads1",
								EntityVersion:         "2",
								ProvisionedModelUnits: 50,
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
						ServedEntities: []serving.ServedEntityOutput{
							{
								Name:          "prod_model",
								EntityName:    "ads1",
								EntityVersion: "2",
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
		Resource: ResourceModelServingProvisionedThroughput(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
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

func TestModelServingProvisionedThroughputUpdate_RemoveConfigIsNoOp(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceModelServingProvisionedThroughput(),
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name":                                               "test-endpoint",
			"config.#":                                           "1",
			"config.0.served_entities.#":                         "1",
			"config.0.served_entities.0.name":                    "prod_model",
			"config.0.served_entities.0.entity_name":             "ads1",
			"config.0.served_entities.0.entity_version":          "2",
			"config.0.served_entities.0.provisioned_model_units": "50",
			"serving_endpoint_id":                                "id",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
				}
			}
			`,
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{},
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/pt/test-endpoint/config",
				ExpectedRequest: serving.UpdateProvisionedThroughputEndpointConfigRequest{
					Name: "test-endpoint",
					Config: serving.PtEndpointCoreConfig{
						ServedEntities: []serving.PtServedModel{
							{
								Name:                  "prod_model",
								EntityName:            "ads1",
								EntityVersion:         "2",
								ProvisionedModelUnits: 50,
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
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name": "test-endpoint",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
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

func TestModelServingProvisionedThroughputUpdate_Tags(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/tags",
				ExpectedRequest: serving.PatchServingEndpointTags{
					Name:    "test-endpoint",
					AddTags: []serving.EndpointTag{{Key: "env", Value: "prod"}},
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
					Tags: []serving.EndpointTag{{Key: "env", Value: "prod"}},
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name":                                   "test-endpoint",
			"tags.#":                                 "1",
			"tags.0.key":                             "env",
			"tags.0.value":                           "prod",
			"config.#":                               "1",
			"config.0.served_entities.#":             "1",
			"config.0.served_entities.0.name":        "prod_model",
			"config.0.served_entities.0.entity_name": "ads1",
			"config.0.served_entities.0.entity_version":          "2",
			"config.0.served_entities.0.provisioned_model_units": "50",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
				}
			}
			tags = [{ key = "env", value = "prod" }]
		`,
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputUpdate_AiGateway(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/ai-gateway",
				ExpectedRequest: serving.PutAiGatewayRequest{
					Name: "test-endpoint",
					Guardrails: &serving.AiGatewayGuardrails{
						Input: &serving.AiGatewayGuardrailParameters{
							Safety: true,
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
					AiGateway: &serving.AiGatewayConfig{
						Guardrails: &serving.AiGatewayGuardrails{
							Input: &serving.AiGatewayGuardrailParameters{
								Safety: true,
							},
						},
					},
					State: &serving.EndpointState{
						ConfigUpdate: serving.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Update:   true,
		ID:       "test-endpoint",
		InstanceState: map[string]string{
			"name":                              "test-endpoint",
			"ai_gateway.#":                      "1",
			"ai_gateway.0.guardrails.#":         "1",
			"ai_gateway.0.guardrails.0.input.#": "1",
			"ai_gateway.0.guardrails.0.input.0.safety": "true",
			"config.#":                                           "1",
			"config.0.served_entities.#":                         "1",
			"config.0.served_entities.0.name":                    "prod_model",
			"config.0.served_entities.0.entity_name":             "ads1",
			"config.0.served_entities.0.entity_version":          "2",
			"config.0.served_entities.0.provisioned_model_units": "50",
		},
		HCL: `
			name = "test-endpoint"
			config {
				served_entities {
					name = "prod_model"
					entity_name = "ads1"
					entity_version = "2"
					provisioned_model_units = 50
				}
			}
			ai_gateway {
				guardrails {
					input {
						safety = true
					}
				}
			}
		`,
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: "",
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Delete:   true,
		ID:       "test-endpoint",
	}.ApplyNoError(t)
}

func TestModelServingProvisionedThroughputDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceModelServingProvisionedThroughput(),
		Delete:   true,
		ID:       "test-endpoint",
	}.ExpectError(t, "Internal error happened")
}
