package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/jobs"
	"github.com/databrickslabs/terraform-provider-databricks/permissions"
	"github.com/databrickslabs/terraform-provider-databricks/policies"
	"github.com/databrickslabs/terraform-provider-databricks/pools"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
	"github.com/databrickslabs/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func permissionsTestHelper(t *testing.T,
	cb func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity)) {
	if os.Getenv("CLOUD_ENV") == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	client := common.NewClientFromEnvironment()

	ctx := context.Background()
	usersAPI := scim.NewUsersAPI(ctx, client)
	me, err := usersAPI.Me()
	require.NoError(t, err)

	user, err := usersAPI.Create(scim.User{
		UserName: fmt.Sprintf("tf-%s@example.com", randomName),
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, usersAPI.Delete(user.ID))
	}()

	groupsAPI := scim.NewGroupsAPI(ctx, client)
	group, err := groupsAPI.Create(scim.Group{
		DisplayName: fmt.Sprintf("tf-%s", randomName),
		Members: []scim.ComplexValue{
			{
				Value: user.ID,
			},
		},
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, groupsAPI.Delete(group.ID))
	}()

	permissionsAPI := permissions.NewPermissionsAPI(ctx, client)
	cb(permissionsAPI, user.UserName, group.DisplayName, func(id string) permissions.PermissionsEntity {
		d := permissions.ResourcePermissions().TestResourceData()
		objectACL, err := permissionsAPI.Read(id)
		require.NoError(t, err)
		entity, err := objectACL.ToPermissionsEntity(d, me.UserName)
		require.NoError(t, err)
		return entity
	})
}

func TestAccPermissionsClusterPolicy(t *testing.T) {
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		policy := policies.ClusterPolicy{
			Name:       group,
			Definition: "{}",
		}
		ctx := context.Background()
		client := common.NewClientFromEnvironment()
		policiesAPI := policies.NewClusterPoliciesAPI(ctx, client)
		require.NoError(t, policiesAPI.Create(&policy))
		defer func() {
			assert.NoError(t, policiesAPI.Delete(policy.PolicyID))
		}()

		objectID := fmt.Sprintf("/cluster-policies/%s", policy.PolicyID)
		require.NoError(t, permissionsAPI.Update(objectID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		client := common.NewClientFromEnvironment()
		poolsAPI := pools.NewInstancePoolsAPI(context.Background(), client)
		ctx := context.Background()
		ips, err := poolsAPI.Create(pools.InstancePool{
			InstancePoolName: group,
			NodeTypeID: clusters.NewClustersAPI(ctx, client).GetSmallestNodeType(
				clusters.NodeTypeRequest{
					LocalDisk: true,
				}),
		})
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, poolsAPI.Delete(ips.InstancePoolID))
		}()

		objectID := fmt.Sprintf("/instance-pools/%s", ips.InstancePoolID)
		require.NoError(t, permissionsAPI.Update(objectID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		ctx := context.Background()
		client := common.NewClientFromEnvironment()
		clustersAPI := clusters.NewClustersAPI(ctx, client)
		clusterInfo, err := compute.NewTinyClusterInCommonPool()
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, clustersAPI.PermanentDelete(clusterInfo.ClusterID))
		}()

		objectID := fmt.Sprintf("/clusters/%s", clusterInfo.ClusterID)
		require.NoError(t, permissionsAPI.Update(objectID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		objectID := "/authorization/tokens"
		require.NoError(t, permissionsAPI.Update(objectID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		ctx := context.Background()
		client := common.NewClientFromEnvironment()
		jobsAPI := jobs.NewJobsAPI(ctx, client)
		job, err := jobsAPI.Create(jobs.JobSettings{
			NewCluster: &clusters.Cluster{
				NumWorkers:   2,
				SparkVersion: "6.4.x-scala2.11",
				NodeTypeID: clusters.NewClustersAPI(ctx, client).GetSmallestNodeType(
					clusters.NodeTypeRequest{
						LocalDisk: true,
					}),
			},
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: "/Production/Featurize",
			},
			Name: group,
		})
		require.NoError(t, err)
		defer func() {
			assert.NoError(t, jobsAPI.Delete(job.ID()))
		}()

		objectID := fmt.Sprintf("/jobs/%s", job.ID())
		require.NoError(t, permissionsAPI.Update(objectID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
	permissionsTestHelper(t, func(permissionsAPI permissions.PermissionsAPI, user, group string,
		ef func(string) permissions.PermissionsEntity) {
		client := common.NewClientFromEnvironment()
		workspaceAPI := workspace.NewNotebooksAPI(context.Background(), client)

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
		require.NoError(t, permissionsAPI.Update(directoryID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
		require.NoError(t, permissionsAPI.Update(notebookID, permissions.AccessControlChangeList{
			AccessControlList: []permissions.AccessControlChange{
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
