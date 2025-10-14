package permissions

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	TestingUser      = "ben"
	TestingAdminUser = "admin"
	TestingOwner     = "testOwner"
)

func TestResourcePermissionsRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/clusters/abc", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

// https://github.com/databricks/terraform-provider-databricks/issues/1227
func TestResourcePermissionsRead_RemovedCluster(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(nil, &apierr.APIError{
				StatusCode: 400,
				ErrorCode:  "INVALID_STATE",
				Message:    "Cannot access cluster X that was terminated or unpinned more than Y days ago.",
			})
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		Removed:  true,
		ID:       "/clusters/abc",
	}.ApplyNoError(t)
}

func TestResourcePermissionsRead_Mlflow_Model(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/registered-models/fakeuuid123",
				ObjectType: "registered-model",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/registered-models/fakeuuid123",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/registered-models/fakeuuid123", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_Mlflow_Model(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/registered-models/fakeuuid123",
				ObjectType: "registered-model",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"registered_model_id": "fakeuuid123",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsUpdate_Mlflow_Model(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/registered-models/fakeuuid123",
				ObjectType: "registered-model",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		InstanceState: map[string]string{
			"registered_model_id": "fakeuuid123",
		},
		HCL: `
		registered_model_id = "fakeuuid123"

		access_control {
			user_name = "ben"
			permission_level = "CAN_READ"
		}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		// Removed:  true,
		ID: "/registered-models/fakeuuid123",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/registered-models/fakeuuid123", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsDelete_Mlflow_Model(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/registered-models/fakeuuid123",
				ObjectType: "registered-model",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "fakeuuid123",
				RequestObjectType: "registered-models",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/registered-models/fakeuuid123",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/registered-models/fakeuuid123", d.Id())
}

func TestResourcePermissionsRead_SQLA_Asset(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "dbsql-dashboards",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "dashboards/abc",
				ObjectType: "dashboard",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/sql/dashboards/abc",
		HCL: `
		sql_dashboard_id = "abc"
		access_control {
			user_name = "ben"
			permission_level = "CAN_VIEW"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/sql/dashboards/abc", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_VIEW", firstElem["permission_level"])
}

func TestResourcePermissionsRead_Dashboard(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "dashboards",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "dashboards/abc",
				ObjectType: "dashboard",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/dashboards/abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/dashboards/abc", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	assert.Equal(t, "abc", d.Get("dashboard_id").(string))
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(nil, &apierr.APIError{
				StatusCode: 404,
				ErrorCode:  "NOT_FOUND",
				Message:    "Cluster does not exist",
			})
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		Removed:  true,
		ID:       "/clusters/abc",
	}.ApplyNoError(t)
}

func TestResourcePermissionsRead_some_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(nil, &apierr.APIError{
				StatusCode: 400,
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
			})
		},
		Resource: ResourcePermissions(),
		Read:     true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.Error(t, err)
}

func TestResourcePermissionsCustomizeDiff_ErrorOnCreate(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourcePermissions(),
		Create:   true,
		HCL: `
		cluster_id = "abc"
		access_control {
			permission_level = "WHATEVER"
		}`,
	}.ExpectError(t, "permission_level WHATEVER is not supported with cluster_id objects; allowed levels: CAN_ATTACH_TO, CAN_MANAGE, CAN_RESTART")
}

func TestResourcePermissionsRead_ErrorOnScimMe(t *testing.T) {
	mock := func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(nil, &apierr.APIError{
			ErrorCode: "INVALID_REQUEST",
			Message:   "Internal error happened",
		})
	}
	qa.MockWorkspaceApply(t, mock, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourcePermissions().ToResource()
		d := r.TestResourceData()
		d.SetId("/clusters/abc")
		diags := r.ReadContext(ctx, d, client)
		assert.True(t, diags.HasError())
		assert.Equal(t, "Internal error happened", diags[0].Summary)
	})
}

func TestResourcePermissionsRead_ToPermissionsEntity_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectType: "teapot",
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/clusters/abc",
	}.ExpectError(t, "expected object type cluster, got teapot")
}

func TestResourcePermissionsRead_EmptyListResultsInRemoval(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		Removed:  true,
		InstanceState: map[string]string{
			"cluster_id": "abc",
		},
		ID: "/clusters/abc",
	}.ApplyNoError(t)
}

func TestResourcePermissionsRead_EmptyListResultsInRemovalWith504Errors(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)

			req := iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}

			// Fail 3 times with a 504 error. These should be retried
			// transparently.
			call := mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, req).Return(nil, apierr.ErrDeadlineExceeded)
			call.Repeatability = 3

			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, req).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
			}, nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		Removed:  true,
		InstanceState: map[string]string{
			"cluster_id": "abc",
		},
		ID: "/clusters/abc",
	}.ApplyNoError(t)
}

func TestResourcePermissionsDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/clusters/abc", d.Id())
}

func TestResourcePermissionsDelete_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.Error(t, err)
}

func TestResourcePermissionsCreate_invalid(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourcePermissions(),
		Create:   true,
	}.ExpectError(t, "at least one type of resource identifier must be set; allowed fields: alert_v2_id, app_name, authorization, cluster_id, cluster_policy_id, dashboard_id, database_instance_name, directory_id, directory_path, experiment_id, instance_pool_id, job_id, notebook_id, notebook_path, pipeline_id, registered_model_id, repo_id, repo_path, serving_endpoint_id, sql_alert_id, sql_dashboard_id, sql_endpoint_id, sql_query_id, vector_search_endpoint_id, workspace_file_id, workspace_file_path")
}

func TestResourcePermissionsCreate_no_access_control(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourcePermissions(),
		Create:   true,
		State: map[string]any{
			"cluster_id": "abc",
		},
	}.ExpectError(t, "invalid config supplied. [access_control] Missing required argument")
}

func TestResourcePermissionsCreate_conflicting_fields(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourcePermissions(),
		Create:   true,
		State: map[string]any{
			"cluster_id":    "abc",
			"notebook_path": "/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
	}.ExpectError(t, "invalid config supplied. [cluster_id] Conflicting configuration arguments. [notebook_path] Conflicting configuration arguments")
}

func TestResourcePermissionsCreate_AdminsThrowError(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
		},
		Resource: ResourcePermissions(),
		Create:   true,
		HCL: `
		cluster_id = "abc"
		access_control {
			group_name = "admins"
			permission_level = "CAN_MANAGE"
		}
		`,
	}.Apply(t)
	assert.EqualError(t, err, "it is not possible to modify admin permissions for cluster resources")
}

func TestResourcePermissionsCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/clusters/abc",
				ObjectType: "cluster",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_ATTACH_TO",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "clusters",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_ATTACH_TO",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"cluster_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_ATTACH_TO",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_ATTACH_TO", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_SQLA_Asset(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "dbsql-dashboards",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/dashboards/abc",
				ObjectType: "dashboard",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_RUN",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "dbsql-dashboards",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_RUN",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"sql_dashboard_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_RUN",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_RUN", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_SQLA_Endpoint(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_USE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "IS_OWNER",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "warehouses/abc",
				ObjectType: "warehouses",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanUse}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelIsOwner}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"sql_endpoint_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_USE", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_SQLA_Endpoint_WithOwnerError(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_USE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "IS_OWNER",
					},
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				Message:    "PUT requests for warehouse *** with no existing owner must provide a new owner.",
				StatusCode: 400,
			})
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_USE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "warehouses/abc",
				ObjectType: "warehouses",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanUse}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelIsOwner}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"sql_endpoint_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_USE", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_SQLA_Endpoint_WithOwner(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingOwner,
						PermissionLevel: "IS_OWNER",
					},
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_USE",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "abc",
				RequestObjectType: "sql/warehouses",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "warehouses/abc",
				ObjectType: "warehouses",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanUse}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
					{
						UserName:       TestingOwner,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelIsOwner}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"sql_endpoint_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
				map[string]any{
					"user_name":        TestingOwner,
					"permission_level": "IS_OWNER",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	accessControlList := ac.List()
	require.Equal(t, 2, len(accessControlList))
	foundTestingUser := false
	foundTestingOwner := false

	for _, entry := range accessControlList {
		entryMap, ok := entry.(map[string]any)
		if !ok {
			t.Fatalf("Expected the entry to be of type map[string]any, got %T", entry)
		}
		if userName, exists := entryMap["user_name"].(string); exists {
			switch userName {
			case TestingUser:
				foundTestingUser = true
				assert.Equal(t, "CAN_USE", entryMap["permission_level"], "Permission level for TestingUser is not CAN_USE")
			case TestingOwner:
				foundTestingOwner = true
				assert.Equal(t, "IS_OWNER", entryMap["permission_level"], "Permission level for TestingOwner is not IS_OWNER")
			}
		}
	}
	assert.True(t, foundTestingUser)
	assert.True(t, foundTestingOwner)
}

func TestResourcePermissionsCreate_NotebookPath_NotExists(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/Development/Init").Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"notebook_path": "/Development/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.Error(t, err)
}

func TestResourcePermissionsCreate_NotebookPath(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/Development/Init").Return(&workspace.ObjectInfo{
				ObjectId:   988765,
				ObjectType: workspace.ObjectTypeNotebook,
			}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "988765",
				RequestObjectType: "notebooks",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "988765",
				RequestObjectType: "notebooks",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/notebooks/988765",
				ObjectType: "notebook",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"notebook_path": "/Development/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_WorkspaceFilePath(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/Development/Init").Return(&workspace.ObjectInfo{
				ObjectId:   988765,
				ObjectType: workspace.ObjectTypeFile,
			}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "988765",
				RequestObjectType: "files",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "988765",
				RequestObjectType: "files",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/files/988765",
				ObjectType: "file",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"workspace_file_path": "/Development/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_error(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourcePermissions(),
		State: map[string]any{
			"cluster_id": "abc",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.ExpectError(t, "permission_level CAN_USE is not supported with cluster_id objects; allowed levels: CAN_ATTACH_TO, CAN_MANAGE, CAN_RESTART")
}

func TestResourcePermissionsCreate_PathIdRetriever_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/foo/bar").Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "i'm a teapot",
				StatusCode: 418,
			})
		},
		Resource: ResourcePermissions(),
		Create:   true,
		HCL: `notebook_path = "/foo/bar"

		access_control {
			user_name = "ben"
			permission_level = "CAN_RUN"
		}`,
	}.ExpectError(t, "cannot load path /foo/bar: i'm a teapot")
}

func TestResourcePermissionsCreate_ActualUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Set(mock.Anything, mock.Anything).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "i'm a teapot",
				StatusCode: 418,
			})
		},
		Resource: ResourcePermissions(),
		Create:   true,
		HCL: `cluster_id = "abc"

		access_control {
			user_name = "ben"
			permission_level = "CAN_MANAGE"
		}`,
	}.ExpectError(t, "i'm a teapot")
}

func TestResourcePermissionsUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "admin"}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "9",
				RequestObjectType: "jobs",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/jobs/9",
				ObjectType: "job",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_VIEW",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "9",
				RequestObjectType: "jobs",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_VIEW",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "IS_OWNER",
					},
				},
			}).Return(nil, nil)
		},
		InstanceState: map[string]string{
			"job_id": "9",
		},
		HCL: `
		job_id = 9

		access_control {
			user_name = "ben"
			permission_level = "CAN_VIEW"
		}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		ID:       "/jobs/9",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/jobs/9", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_VIEW", firstElem["permission_level"])
}

func getResourcePermissions(field, objectType string) resourcePermissions {
	for _, mapping := range allResourcePermissions() {
		if mapping.field == field && mapping.objectType == objectType {
			return mapping
		}
	}
	panic(fmt.Sprintf("could not find resource permissions for field %s and object type %s", field, objectType))
}

func TestResourcePermissionsUpdateTokensAlwaysThereForAdmins(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: "me"}, nil)
		mwc.GetMockPermissionsAPI().EXPECT().Set(mock.Anything, iam.SetObjectPermissions{
			RequestObjectId:   "tokens",
			RequestObjectType: "authorization",
			AccessControlList: []iam.AccessControlRequest{
				{
					UserName:        "me",
					PermissionLevel: "CAN_MANAGE",
				},
				{
					GroupName:       "admins",
					PermissionLevel: "CAN_MANAGE",
				},
			},
		}).Return(nil, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping := getResourcePermissions("authorization", "tokens")
		err := p.Update("/authorization/tokens", entity.PermissionsEntity{
			AccessControlList: []iam.AccessControlRequest{
				{
					UserName:        "me",
					PermissionLevel: "CAN_MANAGE",
				},
			},
		}, mapping)
		assert.NoError(t, err)
	})
}

func TestShouldKeepAdminsOnAnythingExceptPasswordsAndAssignsOwnerForJob(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{
			JobId: int64(123),
		}).Return(&jobs.Job{
			CreatorUserName: "creator@example.com",
		}, nil)
		e := mwc.GetMockPermissionsAPI().EXPECT()
		e.Get(mock.Anything, iam.GetPermissionRequest{
			RequestObjectId:   "123",
			RequestObjectType: "jobs",
		}).Return(&iam.ObjectPermissions{
			ObjectId:   "/jobs/123",
			ObjectType: "job",
			AccessControlList: []iam.AccessControlResponse{
				{
					GroupName: "admins",
					AllPermissions: []iam.Permission{
						{
							PermissionLevel: "CAN_DO_EVERYTHING",
							Inherited:       true,
						},
						{
							PermissionLevel: "CAN_MANAGE",
							Inherited:       false,
						},
					},
				},
			},
		}, nil)
		e.Set(mock.Anything, iam.SetObjectPermissions{
			RequestObjectId:   "123",
			RequestObjectType: "jobs",
			AccessControlList: []iam.AccessControlRequest{
				{
					GroupName:       "admins",
					PermissionLevel: "CAN_MANAGE",
				},
				{
					UserName:        "creator@example.com",
					PermissionLevel: "IS_OWNER",
				},
			},
		}).Return(nil, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping := getResourcePermissions("job_id", "job")
		err := p.Delete(ctx, "/jobs/123", mapping)
		assert.NoError(t, err)
	})
}

func TestShouldDeleteNonExistentJob(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
			RequestObjectId:   "123",
			RequestObjectType: "jobs",
		}).Return(&iam.ObjectPermissions{
			ObjectId:   "/jobs/123",
			ObjectType: "job",
			AccessControlList: []iam.AccessControlResponse{
				{
					GroupName: "admins",
					AllPermissions: []iam.Permission{
						{
							PermissionLevel: "CAN_DO_EVERYTHING",
							Inherited:       true,
						},
						{
							PermissionLevel: "CAN_MANAGE",
							Inherited:       false,
						},
					},
				},
			},
		}, nil)
		mwc.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{
			JobId: int64(123),
		}).Return(nil, &apierr.APIError{
			StatusCode: 400,
			Message:    "Job 123 does not exist.",
			ErrorCode:  "INVALID_PARAMETER_VALUE",
		})
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping := getResourcePermissions("job_id", "job")
		err := p.Delete(ctx, "/jobs/123", mapping)
		assert.NoError(t, err)
	})
}

func TestShouldKeepAdminsOnAnythingExceptPasswordsAndAssignsOwnerForPipeline(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockPipelinesAPI().EXPECT().GetByPipelineId(mock.Anything, "123").Return(&pipelines.GetPipelineResponse{
			CreatorUserName: "creator@example.com",
		}, nil)
		e := mwc.GetMockPermissionsAPI().EXPECT()
		e.Get(mock.Anything, iam.GetPermissionRequest{
			RequestObjectId:   "123",
			RequestObjectType: "pipelines",
		}).Return(&iam.ObjectPermissions{
			ObjectId:   "/pipelines/123",
			ObjectType: "pipeline",
			AccessControlList: []iam.AccessControlResponse{
				{
					GroupName: "admins",
					AllPermissions: []iam.Permission{
						{
							PermissionLevel: "CAN_DO_EVERYTHING",
							Inherited:       true,
						},
						{
							PermissionLevel: "CAN_MANAGE",
							Inherited:       false,
						},
					},
				},
			},
		}, nil)
		e.Set(mock.Anything, iam.SetObjectPermissions{
			RequestObjectId:   "123",
			RequestObjectType: "pipelines",
			AccessControlList: []iam.AccessControlRequest{
				{
					GroupName:       "admins",
					PermissionLevel: "CAN_MANAGE",
				},
				{
					UserName:        "creator@example.com",
					PermissionLevel: "IS_OWNER",
				},
			},
		}).Return(nil, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping := getResourcePermissions("pipeline_id", "pipelines")
		err := p.Delete(ctx, "/pipelines/123", mapping)
		assert.NoError(t, err)
	})
}

func TestPathPermissionsResourceIDFields(t *testing.T) {
	m := getResourcePermissions("notebook_path", "notebook")
	w, err := databricks.NewWorkspaceClient(&databricks.Config{})
	require.NoError(t, err)
	_, err = m.idRetriever(context.Background(), w, "x")
	assert.ErrorContains(t, err, "cannot load path x")
}

func TestDeleteMissing(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
			RequestObjectId:   "x",
			RequestObjectType: "clusters",
		}).Return(nil, apierr.ErrNotFound)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := ResourcePermissions().ToResource()
		d := p.TestResourceData()
		d.SetId("/clusters/x")
		diags := p.DeleteContext(ctx, d, client)
		assert.Nil(t, diags)
	})
}

func TestResourcePermissionsCreate_RepoPath(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/Repos/Development/Init").Return(&workspace.ObjectInfo{
				ObjectId:   988765,
				ObjectType: workspace.ObjectTypeRepo,
			}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "988765",
				RequestObjectType: "repos",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/repos/988765",
				ObjectType: "repo",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_RUN",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "988765",
				RequestObjectType: "repos",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"repo_path": "/Repos/Development/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

// when caller does not specify CAN_MANAGE permission during create, it should be explictly added
func TestResourcePermissionsCreate_Sql_Queries(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "id111",
				RequestObjectType: "sql/queries",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_RUN",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "id111",
				RequestObjectType: "sql/queries",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "queries/id111",
				ObjectType: "query",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRun}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"sql_query_id": "id111",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_RUN",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_RUN", firstElem["permission_level"])
}

// when caller does not specify CAN_MANAGE permission during update, it should be explictly added
func TestResourcePermissionsUpdate_Sql_Queries(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "id111",
				RequestObjectType: "sql/queries",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_RUN",
					},
					{
						UserName:        TestingAdminUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "id111",
				RequestObjectType: "sql/queries",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "queries/id111",
				ObjectType: "query",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRun}},
					},
					{
						UserName:       TestingAdminUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
		},
		InstanceState: map[string]string{
			"sql_query_id": "id111",
		},
		HCL: `
		sql_query_id = "id111",

		access_control = {
			user_name = "ben",
			permission_level = "CAN_RUN",
			}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		ID:       "/sql/queries/id111",
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_RUN", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_DirectoryPath(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			mwc.GetMockWorkspaceAPI().EXPECT().GetStatusByPath(mock.Anything, "/First").Return(&workspace.ObjectInfo{
				ObjectId:   123456,
				ObjectType: workspace.ObjectTypeDirectory,
			}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "123456",
				RequestObjectType: "directories",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
				},
			}).Return(nil, nil)
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "123456",
				RequestObjectType: "directories",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/directories/123456",
				ObjectType: "directory",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName: TestingUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_RUN",
								Inherited:       false,
							},
						},
					},
					{
						UserName: TestingAdminUser,
						AllPermissions: []iam.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"directory_path": "/First",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.NoError(t, err, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsPasswordUsage(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "passwords",
				RequestObjectType: "authorization",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/authorization/passwords",
				ObjectType: "passwords",
				AccessControlList: []iam.AccessControlResponse{
					{
						GroupName:      "admins",
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanUse}},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "passwords",
				RequestObjectType: "authorization",
				AccessControlList: []iam.AccessControlRequest{
					{
						GroupName:       "admins",
						PermissionLevel: "CAN_USE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		HCL: `
		authorization = "passwords"
		access_control {
			group_name       = "admins"
			permission_level = "CAN_USE"
		}
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, "admins", firstElem["group_name"])
	assert.Equal(t, "CAN_USE", firstElem["permission_level"])
}

func TestResourcePermissionsRootDirectory(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   "0",
				RequestObjectType: "directories",
			}).Return(&iam.ObjectPermissions{
				ObjectId:   "/directories/0",
				ObjectType: "directory",
				AccessControlList: []iam.AccessControlResponse{
					{
						UserName:       TestingUser,
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanRead}},
					},
					{
						GroupName:      "admins",
						AllPermissions: []iam.Permission{{PermissionLevel: iam.PermissionLevelCanManage}},
					},
				},
			}, nil)
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   "0",
				RequestObjectType: "directories",
				AccessControlList: []iam.AccessControlRequest{
					{
						UserName:        TestingUser,
						PermissionLevel: "CAN_READ",
					},
					{
						GroupName:       "admins",
						PermissionLevel: "CAN_MANAGE",
					},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		HCL: `
		directory_id = "0"
		access_control {
			user_name        = "ben"
			permission_level = "CAN_READ"
		}
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}
