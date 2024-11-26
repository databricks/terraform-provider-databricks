package apps

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func getTestAppRequest(warehouseName, endpointName string) *apps.App {
	return &apps.App{
		Name:        "my-custom-app",
		Description: "My app description.",
		Resources: []apps.AppResource{
			{
				Name: "sql-warehouse",
				SqlWarehouse: &apps.AppResourceSqlWarehouse{
					Id:         warehouseName,
					Permission: "CAN_MANAGE",
				},
			},
			{
				Name: "job",
				Job: &apps.AppResourceJob{
					Id:         "1234",
					Permission: "CAN_MANAGE",
				},
			},
			{
				Name: "serving-endpoint",
				ServingEndpoint: &apps.AppResourceServingEndpoint{
					Name:       endpointName,
					Permission: "CAN_MANAGE",
				},
			},
			{
				Name:        "api-key",
				Description: "API key for external service.",
				Secret: &apps.AppResourceSecret{
					Scope:      "my-scope",
					Key:        "my-key",
					Permission: "READ",
				},
			},
		},
	}
}

func getTestAppDeployment(path string) *apps.AppDeployment {
	return &apps.AppDeployment{
		DeploymentId:   "01ef0bda89f21f08a8351f41e4a9b948",
		SourceCodePath: path,
		Mode:           "SNAPSHOT",
		DeploymentArtifacts: &apps.AppDeploymentArtifacts{
			SourceCodePath: "/Workspace/Users/9627a015-e892-43f7-9085-eec3892da408/src/01ef1a1ed75d1964b62234a35efa61fc",
		},
		Status: &apps.AppDeploymentStatus{
			State:   "SUCCEEDED",
			Message: "Deployment is in progress.",
		},
		CreateTime: "2019-08-24T14:15:22Z",
		Creator:    "user@test.com",
		UpdateTime: "2019-08-24T14:15:22Z",
	}
}

func getTestAppResponse(warehouseName, endpointName string) *apps.App {
	return &apps.App{
		Name:        "my-custom-app",
		Description: "My app description.",
		ComputeStatus: &apps.ComputeStatus{
			State:   "ERROR",
			Message: "App compute is active.",
		},
		AppStatus: &apps.ApplicationStatus{
			State:   "DEPLOYING",
			Message: "Application is running.",
		},
		Url:               "my-custom-app-123.cloud.databricksapps.com",
		ActiveDeployment:  getTestAppDeployment("/Workspace/user@test.com/my_custom_app"),
		CreateTime:        "2019-08-24T14:15:22Z",
		Creator:           "user@test.com",
		UpdateTime:        "2019-08-24T14:15:22Z",
		Updater:           "user@test.com",
		PendingDeployment: getTestAppDeployment("/Workspace/user@test.com/my_custom_app"),
		Resources: []apps.AppResource{
			{
				Name:        "api-key",
				Description: "API key for external service.",
				Secret: &apps.AppResourceSecret{
					Scope:      "my-scope",
					Key:        "my-key",
					Permission: "READ",
				},
			},
			{
				Name: "sql-warehouse",
				SqlWarehouse: &apps.AppResourceSqlWarehouse{
					Id:         warehouseName,
					Permission: "CAN_MANAGE",
				},
			},
			{
				Name: "serving-endpoint",
				ServingEndpoint: &apps.AppResourceServingEndpoint{
					Name:       endpointName,
					Permission: "CAN_MANAGE",
				},
			},
			{
				Name: "job",
				Job: &apps.AppResourceJob{
					Id:         "1234",
					Permission: "CAN_MANAGE",
				},
			},
		},
		ServicePrincipalId:    0,
		ServicePrincipalName:  "string",
		DefaultSourceCodePath: "/Workspace/user@test.com/my_custom_app",
	}
}

func getTestAppData(warehouseName, endpointName string) map[string]interface{} {
	return map[string]any{
		"name":           "my-custom-app",
		"description":    "My app description.",
		"compute_status": []any{map[string]any{"state": "ERROR", "message": "App compute is active."}},
		"app_status":     []any{map[string]any{"state": "DEPLOYING", "message": "Application is running."}},
		"url":            "my-custom-app-123.cloud.databricksapps.com",
		"active_deployment": []any{map[string]any{
			"deployment_id":        "01ef0bda89f21f08a8351f41e4a9b948",
			"source_code_path":     "/Workspace/user@test.com/my_custom_app",
			"mode":                 "SNAPSHOT",
			"deployment_artifacts": []any{map[string]any{"source_code_path": "/Workspace/Users/9627a015-e892-43f7-9085-eec3892da408/src/01ef1a1ed75d1964b62234a35efa61fc"}},
			"status":               []any{map[string]any{"state": "SUCCEEDED", "message": "Deployment is in progress."}},
			"create_time":          "2019-08-24T14:15:22Z",
			"creator":              "user@test.com",
			"update_time":          "2019-08-24T14:15:22Z",
		}},
		"resource": []any{map[string]any{
			"name":             "api-key",
			"description":      "API key for external service.",
			"job":              []any{},
			"sql_warehouse":    []any{},
			"serving_endpoint": []any{},
			"secret": []any{map[string]any{
				"scope":      "my-scope",
				"key":        "my-key",
				"permission": "READ",
			}},
		}, map[string]any{
			"name":             "sql-warehouse",
			"description":      "",
			"job":              []any{},
			"secret":           []any{},
			"serving_endpoint": []any{},
			"sql_warehouse": []any{map[string]any{
				"id":         warehouseName,
				"permission": "CAN_MANAGE",
			}},
		}, map[string]any{
			"name":          "serving-endpoint",
			"description":   "",
			"job":           []any{},
			"secret":        []any{},
			"sql_warehouse": []any{},
			"serving_endpoint": []any{map[string]any{
				"name":       endpointName,
				"permission": "CAN_MANAGE",
			}},
		}, map[string]any{
			"name":             "job",
			"description":      "",
			"sql_warehouse":    []any{},
			"secret":           []any{},
			"serving_endpoint": []any{},
			"job": []any{map[string]any{
				"id":         "1234",
				"permission": "CAN_MANAGE",
			}},
		}},
		"create_time": "2019-08-24T14:15:22Z",
		"creator":     "user@test.com",
	}
}

func TestResourceAppsCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			api := a.GetMockAppsAPI().EXPECT()
			api.Create(mock.Anything, apps.CreateAppRequest{
				App: getTestAppRequest("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"),
			}).Return(&apps.WaitGetAppActive[apps.App]{
				Name: "my-custom-app",
				Poll: poll.Simple(*getTestAppResponse("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct")),
			}, nil)
			api.GetByName(mock.Anything, "my-custom-app").Return(
				getTestAppResponse("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
		},
		Create: true,
		HCL: `
		name = "my-custom-app"
		description = "My app description."
		resource {
			name = "api-key"
			description = "API key for external service."
			secret {
				scope = "my-scope"
				key = "my-key"
				permission = "READ"
			}
		}
		resource {
			name = "sql-warehouse"
			sql_warehouse {
				id = "e9ca293f79a74b5c"
				permission = "CAN_MANAGE"
			}			
		}
		resource {
			name = "serving-endpoint"
			serving_endpoint {
				name = "databricks-meta-llama-3-1-70b-instruct"
				permission = "CAN_MANAGE"
			}
		}
		resource {
			name = "job"
			job {
				id = "1234"
				permission = "CAN_MANAGE"
			}	
		}			
		`,
		Resource: ResourceApp(),
	}.ApplyAndExpectData(t, getTestAppData("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"))
}

func TestResourceAppsCreateExactlyOnce(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {},
		Create:                  true,
		HCL: `
		name = "my-custom-app"
		description = "My app description."
		resource {
			name = "api-key"
			description = "API key for external service."
			secret {
				scope = "my-scope"
				key = "my-key"
				permission = "READ"
			}
			serving_endpoint {
				name = "databricks-meta-llama-3-1-70b-instruct"
				permission = "CAN_MANAGE"
			}
		}		
		`,
		Resource: ResourceApp(),
	}.ExpectError(t, "exactly one resource type per resource block should be provided")
}

func TestResourceAppsRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			a.GetMockAppsAPI().EXPECT().
				GetByName(mock.Anything, "my-custom-app").
				Return(getTestAppResponse("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
		},
		Resource: ResourceApp(),
		Read:     true,
		New:      true,
		ID:       "my-custom-app",
	}.ApplyAndExpectData(t, getTestAppData("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"))
}

func TestResourceAppsUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			api := a.GetMockAppsAPI().EXPECT()
			api.Update(mock.Anything, apps.UpdateAppRequest{
				Name: "my-custom-app",
				App:  getTestAppRequest("e9ca293f79a74b5c", "new_endpoint"),
			}).Return(getTestAppResponse("e9ca293f79a74b5c", "new_endpoint"), nil)
			api.GetByName(mock.Anything, "my-custom-app").
				Return(getTestAppResponse("e9ca293f79a74b5c", "new_endpoint"), nil)
		},
		Resource: ResourceApp(),
		Update:   true,
		InstanceState: map[string]string{
			"name":        "my-custom-app",
			"description": "My app description.",
		},
		HCL: `
			name = "my-custom-app"
			description = "My app description."
			resource {
				name = "api-key"
				description = "API key for external service."
				secret {
					scope = "my-scope"
					key = "my-key"
					permission = "READ"
				}
			}
			resource {
				name = "sql-warehouse"
				sql_warehouse {
					id = "e9ca293f79a74b5c"
					permission = "CAN_MANAGE"
				}			
			}
			resource {
				name = "serving-endpoint"
				serving_endpoint {
					name = "new_endpoint"
					permission = "CAN_MANAGE"
				}
			}
			resource {
				name = "job"
				job {
					id = "1234"
					permission = "CAN_MANAGE"
				}	
			}	
		`,
		ID: "my-custom-app",
	}.ApplyAndExpectData(t, getTestAppData("e9ca293f79a74b5c", "new_endpoint"))
}

func TestResourceAppsDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			a.GetMockAppsAPI().EXPECT().DeleteByName(mock.Anything, "my-custom-app").Return(getTestAppResponse("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
		},
		Resource: ResourceApp(),
		Delete:   true,
		ID:       "my-custom-app",
	}.ApplyAndExpectData(t, nil)
}
