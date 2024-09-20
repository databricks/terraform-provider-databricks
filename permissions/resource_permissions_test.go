package permissions

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	TestingUser      = "ben"
	TestingAdminUser = "admin"
	TestingOwner     = "testOwner"
	me               = qa.HTTPFixture{
		ReuseRequest: true,
		Method:       "GET",
		Resource:     "/api/2.0/preview/scim/v2/Me",
		Response: scim.User{
			UserName: TestingAdminUser,
		},
	}
)

func TestEntityAccessControlChangeString(t *testing.T) {
	assert.Equal(t, "me CAN_READ", AccessControlChangeApiRequest{
		UserName:        "me",
		PermissionLevel: "CAN_READ",
	}.String())
}

func TestEntityAccessControlString(t *testing.T) {
	assert.Equal(t, "me[CAN_READ (from [parent]) CAN_MANAGE]", AccessControlApiResponse{
		UserName: "me",
		AllPermissions: []PermissionApiResponse{
			{
				InheritedFromObject: []string{"parent"},
				PermissionLevel:     "CAN_READ",
			},
			{
				PermissionLevel: "CAN_MANAGE",
			},
		},
	}.String())
}

func TestResourcePermissionsRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "/clusters/abc",
					ObjectType: "cluster",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Status:   400,
				Response: apierr.APIError{
					ErrorCode: "INVALID_STATE",
					Message:   "Cannot access cluster X that was terminated or unpinned more than Y days ago.",
				},
			},
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
		// Pass list of API request mocks
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				Response: ObjectAclApiResponse{
					ObjectID:   "/registered-models/fakeuuid123",
					ObjectType: "registered-model",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				Response: ObjectAclApiResponse{
					ObjectID:   "/registered-models/fakeuuid123",
					ObjectType: "registered-model",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				Response: ObjectAclApiResponse{
					ObjectID:   "/registered-models/fakeuuid123",
					ObjectType: "registered-model",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				Response: ObjectAclApiResponse{
					ObjectID:   "/registered-models/fakeuuid123",
					ObjectType: "registered-model",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/registered-models/fakeuuid123",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/sql/permissions/dashboards/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/sql/dashboards/abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/sql/dashboards/abc", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]any)
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsRead_Dashboard(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/dashboards/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Cluster does not exist",
				},
				Status: 404,
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
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
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/permissions/clusters/abc",
			Response: ObjectAclApiResponse{
				ObjectID:   "/clusters/abc",
				ObjectType: "clusters",
				AccessControlList: []AccessControlApiResponse{
					{
						UserName: TestingUser,
						AllPermissions: []PermissionApiResponse{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me",
			Response: apierr.APIError{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourcePermissions().ToResource()
		d := r.TestResourceData()
		d.SetId("/clusters/abc")
		diags := r.ReadContext(ctx, d, client)
		assert.True(t, diags.HasError())
		assert.Equal(t, "Internal error happened", diags[0].Summary)
	})
}

func TestResourcePermissionsRead_EmptyListResultsInRemoval(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "/clusters/abc",
					ObjectType: "cluster",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/clusters/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/clusters/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.Error(t, err)
}

func TestResourcePermissionsCreate_invalid(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{me},
		Resource: ResourcePermissions(),
		Create:   true,
	}.ExpectError(t, "at least one type of resource identifier must be set; allowed fields: authorization, cluster_id, cluster_policy_id, dashboard_id, directory_id, directory_path, experiment_id, instance_pool_id, job_id, notebook_id, notebook_path, pipeline_id, registered_model_id, repo_id, repo_path, serving_endpoint_id, sql_alert_id, sql_dashboard_id, sql_endpoint_id, sql_query_id, workspace_file_id, workspace_file_path")
}

func TestResourcePermissionsCreate_no_access_control(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourcePermissions(),
		Create:   true,
		State: map[string]any{
			"cluster_id": "abc",
		},
	}.ExpectError(t, "invalid config supplied. [access_control] Missing required argument")
}

func TestResourcePermissionsCreate_conflicting_fields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
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
		Fixtures: []qa.HTTPFixture{me},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/clusters/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_ATTACH_TO",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "/clusters/abc",
					ObjectType: "cluster",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_ATTACH_TO",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/preview/sql/permissions/dashboards/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/sql/permissions/dashboards/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   "PUT",
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
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
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_USE",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "IS_OWNER",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   "PUT",
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
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
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_PARAMETER_VALUE",
					Message:   "PUT requests for warehouse *** with no existing owner must provide a new owner.",
				},
				Status: 400,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_USE",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_USE",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "IS_OWNER",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   "PUT",
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
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
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/sql/warehouses/abc",
				Response: ObjectAclApiResponse{
					ObjectID:   "dashboards/abc",
					ObjectType: "dashboard",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_USE",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
						{
							UserName:        TestingOwner,
							PermissionLevel: "IS_OWNER",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"notebook_path": "/Development/Init",
			"access_control": []any{
				map[string]any{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.Error(t, err)
}

func TestResourcePermissionsCreate_NotebookPath(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
				Response: workspace.ObjectStatus{
					ObjectID:   988765,
					ObjectType: "NOTEBOOK",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/notebooks/988765",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/notebooks/988765",
				Response: ObjectAclApiResponse{
					ObjectID:   "/notebooks/988765",
					ObjectType: "notebook",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
				Response: workspace.ObjectStatus{
					ObjectID:   988765,
					ObjectType: workspace.File,
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/files/988765",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/files/988765",
				Response: ObjectAclApiResponse{
					ObjectID:   "/files/988765",
					ObjectType: "file",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/clusters/abc",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
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
		Fixtures: []qa.HTTPFixture{
			me,
			qa.HTTPFailures[0],
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
		Fixtures: []qa.HTTPFixture{
			me,
			qa.HTTPFailures[0],
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/jobs/9",
				Response: ObjectAclApiResponse{
					ObjectID:   "/jobs/9",
					ObjectType: "job",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_VIEW",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/jobs/9",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_VIEW",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "IS_OWNER",
						},
					},
				},
			},
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

func TestResourcePermissionsUpdateTokensAlwaysThereForAdmins(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Me",
			Response: scim.User{
				UserName: "me",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/permissions/authorization/tokens",
			ExpectedRequest: AccessControlChangeListApiRequest{
				AccessControlList: []AccessControlChangeApiRequest{
					{
						UserName:        "me",
						PermissionLevel: "CAN_MANAGE",
					},
					{
						GroupName:       "admins",
						PermissionLevel: "CAN_MANAGE",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping, _ := getResourcePermissions("/authorization/tokens")
		err := p.Update("/authorization/tokens", []AccessControlChangeApiRequest{
			{
				UserName:        "me",
				PermissionLevel: "CAN_MANAGE",
			},
		}, mapping)
		assert.NoError(t, err)
	})
}

func TestShouldKeepAdminsOnAnythingExceptPasswordsAndAssignsOwnerForJob(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/jobs/123",
			Response: ObjectAclApiResponse{
				ObjectID:   "/jobs/123",
				ObjectType: "job",
				AccessControlList: []AccessControlApiResponse{
					{
						GroupName: "admins",
						AllPermissions: []PermissionApiResponse{
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
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/get?job_id=123",
			Response: jobs.Job{
				CreatorUserName: "creator@example.com",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/permissions/jobs/123",
			ExpectedRequest: ObjectAclApiResponse{
				AccessControlList: []AccessControlApiResponse{
					{
						GroupName:       "admins",
						PermissionLevel: "CAN_MANAGE",
					},
					{
						UserName:        "creator@example.com",
						PermissionLevel: "IS_OWNER",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping, _ := getResourcePermissions("/jobs/123")
		err := p.Delete("/jobs/123", mapping)
		assert.NoError(t, err)
	})
}

func TestShouldDeleteNonExistentJob(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/jobs/123",
			Response: ObjectAclApiResponse{
				ObjectID:   "/jobs/123",
				ObjectType: "job",
				AccessControlList: []AccessControlApiResponse{
					{
						GroupName: "admins",
						AllPermissions: []PermissionApiResponse{
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
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/jobs/get?job_id=123",
			Status:   400,
			Response: apierr.APIError{
				StatusCode: 400,
				Message:    "Job 123 does not exist.",
				ErrorCode:  "INVALID_PARAMETER_VALUE",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping, _ := getResourcePermissions("/jobs/123")
		err := p.Delete("/jobs/123", mapping)
		assert.NoError(t, err)
	})
}

func TestShouldKeepAdminsOnAnythingExceptPasswordsAndAssignsOwnerForPipeline(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/permissions/pipelines/123",
			Response: ObjectAclApiResponse{
				ObjectID:   "/pipelines/123",
				ObjectType: "pipeline",
				AccessControlList: []AccessControlApiResponse{
					{
						GroupName: "admins",
						AllPermissions: []PermissionApiResponse{
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
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines/123?",
			Response: jobs.Job{
				CreatorUserName: "creator@example.com",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/permissions/pipelines/123",
			ExpectedRequest: ObjectAclApiResponse{
				AccessControlList: []AccessControlApiResponse{
					{
						GroupName:       "admins",
						PermissionLevel: "CAN_MANAGE",
					},
					{
						UserName:        "creator@example.com",
						PermissionLevel: "IS_OWNER",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		p := NewPermissionsAPI(ctx, client)
		mapping, _ := getResourcePermissions("/pipelines/123")
		err := p.Delete("/pipelines/123", mapping)
		assert.NoError(t, err)
	})
}

func TestPathPermissionsResourceIDFields(t *testing.T) {
	var m resourcePermissions
	for _, x := range allResourcePermissions() {
		if x.field == "notebook_path" {
			m = x
		}
	}
	w, err := databricks.NewWorkspaceClient(&databricks.Config{})
	require.NoError(t, err)
	_, err = m.idRetriever(context.Background(), w, "x")
	assert.ErrorContains(t, err, "cannot load path x")
}

func TestObjectACLToPermissionsEntityCornerCases(t *testing.T) {
	_, _, err := getResourcePermissionsForObjectAcl(ObjectAclApiResponse{
		ObjectType: "bananas",
		AccessControlList: []AccessControlApiResponse{
			{
				GroupName: "admins",
			},
		},
	})
	assert.EqualError(t, err, "unknown object type bananas")
}

func TestEntityAccessControlToAccessControlChange(t *testing.T) {
	_, res := AccessControlApiResponse{}.toAccessControlChange()
	assert.False(t, res)
}

func TestCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourcePermissions(), qa.CornerCaseSkipCRUD("create"), qa.CornerCaseID("/clusters/x"))
}

func TestDeleteMissing(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny: true,
			Status:   404,
			Response: apierr.NotFound("missing"),
		},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FRepos%2FDevelopment%2FInit",
				Response: workspace.ObjectStatus{
					ObjectID:   988765,
					ObjectType: "repo",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/repos/988765",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/repos/988765",
				Response: ObjectAclApiResponse{
					ObjectID:   "/repos/988765",
					ObjectType: "repo",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_RUN",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/preview/sql/permissions/queries/id111",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/sql/permissions/queries/id111",
				Response: ObjectAclApiResponse{
					ObjectID:   "queries/id111",
					ObjectType: "query",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/preview/sql/permissions/queries/id111",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/sql/permissions/queries/id111",
				Response: ObjectAclApiResponse{
					ObjectID:   "queries/id111",
					ObjectType: "query",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FFirst",
				Response: workspace.ObjectStatus{
					ObjectID:   123456,
					ObjectType: "directory",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/directories/123456",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/directories/123456",
				Response: ObjectAclApiResponse{
					ObjectID:   "/directories/123456",
					ObjectType: "directory",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName: TestingUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_RUN",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []PermissionApiResponse{
								{
									PermissionLevel: "CAN_MANAGE",
									Inherited:       false,
								},
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/authorization/passwords",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							GroupName:       "admins",
							PermissionLevel: "CAN_USE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/authorization/passwords",
				Response: ObjectAclApiResponse{
					ObjectID:   "/authorization/passwords",
					ObjectType: "passwords",
					AccessControlList: []AccessControlApiResponse{
						{
							GroupName:       "admins",
							PermissionLevel: "CAN_USE",
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			me,
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/permissions/directories/0",
				ExpectedRequest: AccessControlChangeListApiRequest{
					AccessControlList: []AccessControlChangeApiRequest{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							GroupName:       "admins",
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/permissions/directories/0",
				Response: ObjectAclApiResponse{
					ObjectID:   "/directories/0",
					ObjectType: "directory",
					AccessControlList: []AccessControlApiResponse{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_READ",
						},
						{
							GroupName:       "admins",
							PermissionLevel: "CAN_MANAGE",
						},
					},
				},
			},
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
