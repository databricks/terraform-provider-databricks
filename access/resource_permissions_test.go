package access

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
					ObjectType: "cluster",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
		New:      true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/clusters/abc", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]interface{})
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
}

func TestResourcePermissionsRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Cluster does not exist",
				},
				Status: 404,
			},
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/clusters/abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
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
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: ObjectACL{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				Response: ObjectACL{
					ObjectID:   "/clusters/abc",
					ObjectType: "clusters",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
	qa.AssertErrorStartsWith(t, err, "Invalid config supplied. [access_control] Required attribute is not set")
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
				Method:   http.MethodPut,
				Resource: "/api/2.0/preview/permissions/clusters/abc",
				ExpectedRequest: AccessControlChangeList{
					AccessControlList: []AccessControlChange{
						{
							UserName:        TestingUser,
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
					ObjectType: "cluster",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_READ",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]interface{})
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_READ", firstElem["permission_level"])
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
				Response: workspace.ObjectStatus{
					ObjectID:   988765,
					ObjectType: "NOTEBOOK",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/preview/permissions/notebooks/988765",
				ExpectedRequest: AccessControlChangeList{
					AccessControlList: []AccessControlChange{
						{
							UserName:        TestingUser,
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
					ObjectType: "notebook",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_USE",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]interface{})
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_USE", firstElem["permission_level"])
}

func TestResourcePermissionsCreate_error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
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

func TestResourcePermissionsUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/permissions/jobs/9",
				Response: ObjectACL{
					ObjectID:   "/jobs/9",
					ObjectType: "job",
					AccessControlList: []AccessControl{
						{
							UserName: TestingUser,
							AllPermissions: []Permission{
								{
									PermissionLevel: "CAN_RUN",
									Inherited:       false,
								},
							},
						},
						{
							UserName: TestingAdminUser,
							AllPermissions: []Permission{
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
				Method:       http.MethodGet,
				ReuseRequest: true,
				Resource:     "/api/2.0/preview/scim/v2/Me",
				Response: identity.ScimUser{
					UserName: TestingAdminUser,
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/preview/permissions/jobs/9",
				ExpectedRequest: AccessControlChangeList{
					AccessControlList: []AccessControlChange{
						{
							UserName:        TestingUser,
							PermissionLevel: "CAN_RUN",
						},
						{
							UserName:        TestingAdminUser,
							PermissionLevel: "IS_OWNER",
						},
					},
				},
			},
		},
		HCL: `
		job_id = 9

		access_control {
			user_name = "ben"
			permission_level = "CAN_RUN"
		}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		ID:       "/jobs/9",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/jobs/9", d.Id())
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	firstElem := ac.List()[0].(map[string]interface{})
	assert.Equal(t, TestingUser, firstElem["user_name"])
	assert.Equal(t, "CAN_RUN", firstElem["permission_level"])
}

func permissionsTestHelper(t *testing.T,
	cb func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity)) {
	if "" == os.Getenv("CLOUD_ENV") {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	client := common.NewClientFromEnvironment()

	ctx := context.Background()
	usersAPI := identity.NewUsersAPI(ctx, client)
	me, err := usersAPI.Me()
	require.NoError(t, err)

	user, err := usersAPI.Create(identity.UserEntity{
		UserName: fmt.Sprintf("tf-%s@example.com", randomName),
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, usersAPI.Delete(user.ID))
	}()

	groupsAPI := identity.NewGroupsAPI(ctx, client)
	group, err := groupsAPI.Create(fmt.Sprintf("tf-%s", randomName), []string{user.ID}, nil, nil)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, groupsAPI.Delete(group.ID))
	}()

	permissionsAPI := NewPermissionsAPI(ctx, client)
	cb(permissionsAPI, user.UserName, group.DisplayName, func(id string) PermissionsEntity {
		d := ResourcePermissions().TestResourceData()
		objectACL, err := permissionsAPI.Read(id)
		require.NoError(t, err)
		entity, err := objectACL.ToPermissionsEntity(context.Background(), d, me.UserName)
		require.NoError(t, err)
		return entity
	})
}

func TestAccPermissionsClusterPolicy(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		policy := compute.ClusterPolicy{
			Name:       group,
			Definition: "{}",
		}
		policiesAPI := compute.NewClusterPoliciesAPI(permissionsAPI.client)
		require.NoError(t, policiesAPI.Create(&policy))
		defer func() {
			assert.NoError(t, policiesAPI.Delete(policy.PolicyID))
		}()

		objectID := fmt.Sprintf("/cluster-policies/%s", policy.PolicyID)
		require.NoError(t, permissionsAPI.Update(objectID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "CAN_USE",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_USE",
				},
			},
		}))
		entity := ef(objectID)
		assert.Equal(t, "cluster-policy", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(objectID))
		entity = ef(objectID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}

func TestAccPermissionsInstancePool(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		poolsAPI := compute.NewInstancePoolsAPI(permissionsAPI.client)
		ctx := context.Background()
		ips, err := poolsAPI.Create(compute.InstancePool{
			InstancePoolName: group,
			NodeTypeID: compute.NewClustersAPI(
				ctx, permissionsAPI.client).GetSmallestNodeType(
				compute.NodeTypeRequest{
					LocalDisk: true,
				}),
		})
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, poolsAPI.Delete(ips.InstancePoolID))
		}()

		objectID := fmt.Sprintf("/instance-pools/%s", ips.InstancePoolID)
		require.NoError(t, permissionsAPI.Update(objectID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "CAN_MANAGE",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_ATTACH_TO",
				},
			},
		}))
		entity := ef(objectID)
		assert.Equal(t, "instance-pool", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(objectID))
		entity = ef(objectID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}

func TestAccPermissionsClusters(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		ctx := context.Background()
		clustersAPI := compute.NewClustersAPI(ctx, permissionsAPI.client)
		clusterInfo, err := compute.NewTinyClusterInCommonPool()
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, clustersAPI.PermanentDelete(clusterInfo.ClusterID))
		}()

		objectID := fmt.Sprintf("/clusters/%s", clusterInfo.ClusterID)
		require.NoError(t, permissionsAPI.Update(objectID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "CAN_RESTART",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_ATTACH_TO",
				},
			},
		}))
		entity := ef(objectID)
		assert.Equal(t, "cluster", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(objectID))
		entity = ef(objectID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}

func TestAccPermissionsTokens(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		objectID := "/authorization/tokens"
		require.NoError(t, permissionsAPI.Update(objectID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "CAN_USE",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_USE",
				},
			},
		}))
		entity := ef(objectID)
		assert.Equal(t, "tokens", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(objectID))
		entity = ef(objectID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}

func TestAccPermissionsJobs(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		ctx := context.Background()
		jobsAPI := compute.NewJobsAPI(permissionsAPI.client)
		job, err := jobsAPI.Create(compute.JobSettings{
			NewCluster: &compute.Cluster{
				NumWorkers:   2,
				SparkVersion: "6.4.x-scala2.11",
				NodeTypeID: compute.NewClustersAPI(
					ctx, permissionsAPI.client).GetSmallestNodeType(
					compute.NodeTypeRequest{
						LocalDisk: true,
					}),
			},
			NotebookTask: &compute.NotebookTask{
				NotebookPath: "/Production/Featurize",
			},
			Name: group,
		})
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, jobsAPI.Delete(job.ID()))
		}()

		objectID := fmt.Sprintf("/jobs/%s", job.ID())
		require.NoError(t, permissionsAPI.Update(objectID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "IS_OWNER",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_MANAGE_RUN",
				},
			},
		}))
		entity := ef(objectID)
		assert.Equal(t, "job", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(objectID))
		entity = ef(objectID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}

func TestAccPermissionsNotebooks(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI PermissionsAPI, user, group string,
		ef func(string) PermissionsEntity) {
		workspaceAPI := workspace.NewNotebooksAPI(context.Background(), permissionsAPI.client)

		notebookDir := fmt.Sprintf("/Testing/%s/something", group)
		err := workspaceAPI.Mkdirs(notebookDir)
		require.NoError(t, err)

		notebookPath := fmt.Sprintf("%s/Dummy", notebookDir)

		err = workspaceAPI.Create(workspace.ImportRequest{
			Path:      notebookPath,
			Content:   "MSsx",
			Format:    "SOURCE",
			Language:  "PYTHON",
			Overwrite: true,
		})
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, workspaceAPI.Delete(notebookDir, true))
		}()

		folder, err := workspaceAPI.Read(fmt.Sprintf("/Testing/%s", group))
		require.NoError(t, err)

		directoryID := fmt.Sprintf("/directories/%d", folder.ObjectID)
		require.NoError(t, permissionsAPI.Update(directoryID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					GroupName:       "users",
					PermissionLevel: "CAN_READ",
				},
			},
		}))
		entity := ef(directoryID)
		assert.Equal(t, "directory", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 1)

		notebook, err := workspaceAPI.Read(notebookPath)
		require.NoError(t, err)
		notebookID := fmt.Sprintf("/notebooks/%d", notebook.ObjectID)
		require.NoError(t, permissionsAPI.Update(notebookID, AccessControlChangeList{
			AccessControlList: []AccessControlChange{
				{
					UserName:        user,
					PermissionLevel: "CAN_MANAGE",
				},
				{
					GroupName:       group,
					PermissionLevel: "CAN_EDIT",
				},
			},
		}))

		entity = ef(notebookID)
		assert.Equal(t, "notebook", entity.ObjectType)
		assert.Len(t, entity.AccessControlList, 2)

		require.NoError(t, permissionsAPI.Delete(directoryID))
		entity = ef(directoryID)
		assert.Len(t, entity.AccessControlList, 0)
	})
}
