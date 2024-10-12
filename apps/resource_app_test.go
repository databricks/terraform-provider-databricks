package apps

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func getTestApp(warehouseName, endpointName string) *apps.App {
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
		Url: "my-custom-app-123.cloud.databricksapps.com",
		ActiveDeployment: &apps.AppDeployment{
			DeploymentId:   "01ef0bda89f21f08a8351f41e4a9b948",
			SourceCodePath: "/Workspace/user@test.com/my_custom_app",
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
		},
		CreateTime: "2019-08-24T14:15:22Z",
		Creator:    "user@test.com",
		UpdateTime: "2019-08-24T14:15:22Z",
		Updater:    "user@test.com",
		PendingDeployment: &apps.AppDeployment{
			DeploymentId:   "01ef0bda89f21f08a8351f41e4a9b948",
			SourceCodePath: "/Workspace/user@test.com/my_custom_app",
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
		},
		Resources: []apps.AppResource{
			{
				Name:        "api-key",
				Description: "API key for external service.",
				Secret: &apps.AppResourceSecret{
					Scope:      "my-scope",
					Key:        "my-key",
					Permission: "READ",
				},
				SqlWarehouse: &apps.AppResourceSqlWarehouse{
					Id:         warehouseName,
					Permission: "CAN_MANAGE",
				},
				ServingEndpoint: &apps.AppResourceServingEndpoint{
					Name:       endpointName,
					Permission: "CAN_MANAGE",
				},
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
		"resources": []any{map[string]any{
			"name":        "api-key",
			"description": "API key for external service.",
			"secret": []any{map[string]any{
				"scope":      "my-scope",
				"key":        "my-key",
				"permission": "READ",
			}},
			"sql_warehouse": []any{map[string]any{
				"id":         warehouseName,
				"permission": "CAN_MANAGE",
			}},
			"serving_endpoint": []any{map[string]any{
				"name":       endpointName,
				"permission": "CAN_MANAGE",
			}},
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
				Name:        "my-custom-app",
				Description: "My app description.",
				Resources: []apps.AppResource{
					{
						Name:        "api-key",
						Description: "API key for external service.",
						Secret: &apps.AppResourceSecret{
							Scope:      "my-scope",
							Key:        "my-key",
							Permission: "READ",
						},
						SqlWarehouse: &apps.AppResourceSqlWarehouse{
							Id:         "e9ca293f79a74b5c",
							Permission: "CAN_MANAGE",
						},
						ServingEndpoint: &apps.AppResourceServingEndpoint{
							Name:       "databricks-meta-llama-3-1-70b-instruct",
							Permission: "CAN_MANAGE",
						},
						Job: &apps.AppResourceJob{
							Id:         "1234",
							Permission: "CAN_MANAGE",
						},
					},
				},
			}).Return(&apps.WaitGetAppActive[apps.App]{Poll: poll.Simple(*getTestApp("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"))}, nil)
			api.GetByName(mock.Anything, "my-custom-app").Return(
				getTestApp("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
		},
		Create: true,
		HCL: `
		name = "my-custom-app"
		description = "My app description."
		resources {
			name = "api-key"
			description = "API key for external service."
			secret {
				scope = "my-scope"
				key = "my-key"
				permission = "READ"
			}
			sql_warehouse {
				id = "e9ca293f79a74b5c"
				permission = "CAN_MANAGE"
			}
			serving_endpoint {
				name = "databricks-meta-llama-3-1-70b-instruct"
				permission = "CAN_MANAGE"
			}
			job {
				id = "1234"
				permission = "CAN_MANAGE"
			}
		}
		`,
		Resource: ResourceApp(),
	}.ApplyAndExpectData(t, getTestAppData("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"))
}

func TestResourceAppsRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			a.GetMockAppsAPI().EXPECT().
				GetByName(mock.Anything, "my-custom-app").
				Return(getTestApp("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
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
				Name:        "my-custom-app",
				Description: "My app description.",
				Resources: []apps.AppResource{
					{
						Name:        "api-key",
						Description: "API key for external service.",
						Secret: &apps.AppResourceSecret{
							Scope:      "my-scope",
							Key:        "my-key",
							Permission: "READ",
						},
						SqlWarehouse: &apps.AppResourceSqlWarehouse{
							Id:         "new_warehouse",
							Permission: "CAN_MANAGE",
						},
						ServingEndpoint: &apps.AppResourceServingEndpoint{
							Name:       "new_endpoint",
							Permission: "CAN_MANAGE",
						},
						Job: &apps.AppResourceJob{
							Id:         "1234",
							Permission: "CAN_MANAGE",
						},
					},
				},
			}).Return(getTestApp("new_warehouse", "new_endpoint"), nil)
			api.GetByName(mock.Anything, "my-custom-app").
				Return(getTestApp("new_warehouse", "new_endpoint"), nil)
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
			resources {
				name = "api-key"
				description = "API key for external service."
				secret {
					scope = "my-scope"
					key = "my-key"
					permission = "READ"
				}
				sql_warehouse {
					id = "new_warehouse"
					permission = "CAN_MANAGE"
				}
				serving_endpoint {
					name = "new_endpoint"
					permission = "CAN_MANAGE"
				}
				job {
					id = "1234"
					permission = "CAN_MANAGE"
				}
			}
		`,
		ID: "my-custom-app",
	}.ApplyAndExpectData(t, getTestAppData("new_warehouse", "new_endpoint"))
}

func TestResourceAppsDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			a.GetMockAppsAPI().EXPECT().DeleteByName(mock.Anything, "my-custom-app").Return(getTestApp("e9ca293f79a74b5c", "databricks-meta-llama-3-1-70b-instruct"), nil)
		},
		Resource: ResourceApp(),
		Delete:   true,
		ID:       "my-custom-app",
	}.ApplyAndExpectData(t, nil)
}
