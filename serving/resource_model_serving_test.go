package serving

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
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

func TestModelServingCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				Response: common.APIErrorBody{
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
	}.ApplyNoError(t)
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
				Response: common.APIErrorBody{
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
				Response: common.APIErrorBody{
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
				Response: common.APIErrorBody{
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
