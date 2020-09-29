package access

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
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
				Response: identity.ScimUser{
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

func TestResourcePermissionsDelete_Job(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.0/preview/permissions/jobs/101",
				ExpectedRequest: ObjectACL{},
			},
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/jobs/101",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/jobs/101", d.Id())
}

func TestJobResourcePermissionsDelete_Job_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.0/preview/permissions/jobs/101",
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
		ID:       "/jobs/101",
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
				Response: identity.ScimUser{
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
				Response: identity.ScimUser{
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

func TestAccAddOrModifyDeleteJobPermissions(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	randomName := qa.RandomName()

	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()

	jobSettings := compute.JobSettings{
		NewCluster: &compute.Cluster{
			NumWorkers:   2,
			SparkVersion: "6.4.x-scala2.11",
			SparkConf:    nil,
			NodeTypeID:   "Standard_DS3_v2",
		},
		Name: "Job " + randomName,
	}
	jobCreate, err := compute.NewJobsAPI(client).Create(jobSettings)

	assert.NoError(t, err)

	//job, err := compute.NewJobsAPI(client).Read(fmt.Sprint(jobCreate.ID()))

	jobID := fmt.Sprint(jobCreate.ID())

	groupName := "Group" + randomName
	_, groupsAPIErr := identity.NewGroupsAPI(client).Create(groupName, nil, nil, nil)

	assert.NoError(t, groupsAPIErr)

	userName := randomName + "@" + randomName + ".com"
	_, usersAPIErr := identity.NewUsersAPI(client).Create(userName, randomName, nil, nil)

	assert.NoError(t, usersAPIErr)

	ownerName := jobCreate.CreatorUserName

	groupACL := AccessControlChange{
		GroupName:       &groupName,
		PermissionLevel: "CAN_MANAGE",
	}

	ownerACL := AccessControlChange{
		UserName:        &ownerName,
		PermissionLevel: "CAN_MANAGE",
	}

	userACL := AccessControlChange{
		UserName:        &userName,
		PermissionLevel: "IS_OWNER",
	}

	accessControlChange := []*AccessControlChange{&groupACL, &userACL, &ownerACL}

	jobACL := AccessControlChangeList{
		accessControlChange,
	}

	param := &jobACL

	permissionsAPIErr := NewPermissionsAPI(client).AddOrModify(fmt.Sprintf("/jobs/%s/", jobID), param)

	assert.NoError(t, permissionsAPIErr)

	ownerNewACL := AccessControlChange{
		UserName:        &ownerName,
		PermissionLevel: "IS_OWNER",
	}

	ownerAccessControlChange := []*AccessControlChange{&ownerNewACL}

	newAcl := AccessControlChangeList{
		ownerAccessControlChange,
	}
	permissionsAPIUpdateErr := NewPermissionsAPI(client).SetOrDelete(fmt.Sprintf("/jobs/%s/", jobID), &newAcl)

	assert.NoError(t, permissionsAPIUpdateErr)

	jobsAPIErr := compute.NewJobsAPI(client).Delete(fmt.Sprint(jobCreate.JobID))

	assert.NoError(t, jobsAPIErr)

}
