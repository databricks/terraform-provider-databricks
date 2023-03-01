package serving_endpoints

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/endpoints"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestServingEndpointCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceServingEndpoint())
}

func TestServingEndpointCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				ExpectedRequest: endpoints.CreateServingEndpoint{
					Name: "test-endpoint",
					Config: endpoints.EndpointCoreConfigInput{
						ServedModels: []endpoints.ServedModelInput{
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
						TrafficConfig: &endpoints.TrafficConfig{
							Routes: []endpoints.Route{
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
				Response: endpoints.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				ExpectedRequest: endpoints.GetServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: endpoints.ServingEndpointDetailed{
					Name: "test-endpoint",
					State: &endpoints.EndpointState{
						ConfigUpdate: endpoints.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				ExpectedRequest: endpoints.GetServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: endpoints.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &endpoints.EndpointState{
						ConfigUpdate: endpoints.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &endpoints.EndpointCoreConfigOutput{
						ServedModels: []endpoints.ServedModelOutput{
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
						TrafficConfig: &endpoints.TrafficConfig{
							Routes: []endpoints.Route{
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
		Resource: ResourceServingEndpoint(),
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

func TestServingEndpointCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/serving-endpoints",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServingEndpoint(),
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestServingEndpointRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				ExpectedRequest: endpoints.GetServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: endpoints.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &endpoints.EndpointState{
						ConfigUpdate: endpoints.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &endpoints.EndpointCoreConfigOutput{
						ServedModels: []endpoints.ServedModelOutput{
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
						TrafficConfig: &endpoints.TrafficConfig{
							Routes: []endpoints.Route{
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
		Resource: ResourceServingEndpoint(),
		Read:     true,
		ID:       "test-endpoint",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "test-endpoint", d.Id())
}

func TestServingEndpointRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServingEndpoint(),
		Read:     true,
		ID:       "test-endpoint",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "test-endpoint", d.Id())
}
func TestServingEndpointUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/config",
				ExpectedRequest: endpoints.EndpointCoreConfigInput{
					ServedModels: []endpoints.ServedModelInput{
						{
							Name:               "prod_model",
							ModelName:          "ads1",
							ModelVersion:       "2",
							WorkloadSize:       "Small",
							ScaleToZeroEnabled: true,
						},
					},
					TrafficConfig: &endpoints.TrafficConfig{
						Routes: []endpoints.Route{
							{
								ServedModelName:   "prod_model",
								TrafficPercentage: 100,
							},
						},
					},
				},
				Response: endpoints.ServingEndpointDetailed{
					Name: "test-endpoint",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				ExpectedRequest: endpoints.GetServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: endpoints.ServingEndpointDetailed{
					Name: "test-endpoint",
					State: &endpoints.EndpointState{
						ConfigUpdate: endpoints.EndpointStateConfigUpdateNotUpdating,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/serving-endpoints/test-endpoint?",
				ExpectedRequest: endpoints.GetServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: endpoints.ServingEndpointDetailed{
					Id:   "test-endpoint",
					Name: "test-endpoint",
					State: &endpoints.EndpointState{
						ConfigUpdate: endpoints.EndpointStateConfigUpdateNotUpdating,
					},
					Config: &endpoints.EndpointCoreConfigOutput{
						ServedModels: []endpoints.ServedModelOutput{
							{
								Name:               "prod_model",
								ModelName:          "ads1",
								ModelVersion:       "2",
								ScaleToZeroEnabled: true,
							},
						},
						TrafficConfig: &endpoints.TrafficConfig{
							Routes: []endpoints.Route{
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
		Resource: ResourceServingEndpoint(),
		Update:   true,
		ID:       "test-endpoint",
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
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "test-endpoint", d.Id())
}

func TestServingEndpointUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/serving-endpoints/test-endpoint/config",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServingEndpoint(),
		Update:   true,
		ID:       "test-endpoint",
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
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "test-endpoint", d.Id())
}

func TestServingEndpointDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint",
				ExpectedRequest: endpoints.DeleteServingEndpointRequest{
					Name: "test-endpoint",
				},
				Response: "",
			},
		},
		Resource: ResourceServingEndpoint(),
		Delete:   true,
		ID:       "test-endpoint",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "test-endpoint", d.Id())
}

func TestServingEndpointDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/serving-endpoints/test-endpoint",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServingEndpoint(),
		Delete:   true,
		ID:       "test-endpoint",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "test-endpoint", d.Id())
}
