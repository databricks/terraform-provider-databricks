package access

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/stretchr/testify/assert"
)

var (
	TestingUser      = "ben"
	TestingAdminUser = "admin"
)

func TestResourcePermissionsRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: ObjectACL{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []*AccessControl{
						{
							UserName: &TestingUser,
							AllPermissions: []*Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: &TestingAdminUser,
							AllPermissions: []*Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: identity.User{
					UserName: TestingAdminUser,
				},
			},
		},
		Resource: ResourcePermissions(),
		Read:     true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/clusters/abc", d.Id())
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_READ", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestResourcePermissionsRead_some_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: common.APIErrorBody{
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

func TestResourcePermissionsRead_ErrorOnScimMe(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: ObjectACL{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []*AccessControl{
						{
							UserName: &TestingUser,
							AllPermissions: []*Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: &TestingAdminUser,
							AllPermissions: []*Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: common.APIErrorBody{
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

func TestResourcePermissionsDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.0/preview/permissions/clusters/abc",
				ExpectedRequest: ObjectACL{},
			},
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/clusters/abc", d.Id())
}

func TestResourcePermissionsDelete_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.0/preview/permissions/clusters/abc",
				ExpectedRequest: ObjectACL{},
				Response: common.APIErrorBody{
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
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourcePermissions(),
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "At least one type of resource identifiers must be set")
}

func TestResourcePermissionsCreate_no_access_control(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourcePermissions(),
		Create:   true,
		State: map[string]interface{}{
			"cluster_id": "abc",
		},
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Invalid config supplied. Required attribute is not set")
}

func TestResourcePermissionsCreate_conflicting_fields(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourcePermissions(),
		Create:   true,
		State: map[string]interface{}{
			"cluster_id":    "abc",
			"notebook_path": "/Init",
			"access_control": []interface{}{
				map[string]interface{}{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Invalid config supplied. cluster_id: conflicts with notebook_path. notebook_path: conflicts with cluster_id")
}

func TestResourcePermissionsCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				ExpectedRequest: AccessControlChangeList{
					AccessControlList: []*AccessControlChange{
						{
							UserName:        &TestingUser,
							PermissionLevel: "CAN_READ",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: ObjectACL{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []*AccessControl{
						{
							UserName: &TestingUser,
							AllPermissions: []*Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: &TestingAdminUser,
							AllPermissions: []*Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: identity.User{
					UserName: TestingAdminUser,
				},
			},
		},
		Resource: ResourcePermissions(),
		State: map[string]interface{}{
			"cluster_id": "abc",
			"access_control": []interface{}{
				map[string]interface{}{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_READ", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestResourcePermissionsCreate_NotebookPath_NotExists(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePermissions(),
		State: map[string]interface{}{
			"notebook_path": "/Development/Init",
			"access_control": []interface{}{
				map[string]interface{}{
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
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
				Response: workspace.WorkspaceObjectStatus{
					ObjectID:   988765,
					ObjectType: "NOTEBOOK",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/permissions/notebooks/988765",
				ExpectedRequest: AccessControlChangeList{
					AccessControlList: []*AccessControlChange{
						{
							UserName:        &TestingUser,
							PermissionLevel: "CAN_USE",
						},
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/notebooks/988765",
				Response: ObjectACL{
					ObjectID:   "/notebooks/988765",
					ObjectType: "notebooks",
					AccessControlList: []*AccessControl{
						{
							UserName: &TestingUser,
							AllPermissions: []*Permission{
								{
									PermissionLevel: "CAN_USE",
									Inherited:       false,
								},
							},
						},
						{
							UserName: &TestingAdminUser,
							AllPermissions: []*Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: identity.User{
					UserName: TestingAdminUser,
				},
			},
		},
		Resource: ResourcePermissions(),
		State: map[string]interface{}{
			"notebook_path": "/Development/Init",
			"access_control": []interface{}{
				map[string]interface{}{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_USE", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestResourcePermissionsCreate_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePermissions(),
		State: map[string]interface{}{
			"cluster_id": "abc",
			"access_control": []interface{}{
				map[string]interface{}{
					"user_name":        TestingUser,
					"permission_level": "CAN_USE",
				},
			},
		},
		Create: true,
	}.Apply(t)
	if assert.Error(t, err) {
		if e, ok := err.(common.APIError); ok {
			assert.Equal(t, "INVALID_REQUEST", e.ErrorCode)
		}
	}
}
