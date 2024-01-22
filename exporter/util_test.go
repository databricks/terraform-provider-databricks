package exporter

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/stretchr/testify/assert"
)

func TestImportClusterEmitsInitScripts(t *testing.T) {
	ic := importContextForTest()
	ic.importCluster(&clusters.Cluster{
		InitScripts: []clusters.InitScriptStorageInfo{
			{
				Dbfs: &clusters.DbfsStorageInfo{
					Destination: "/mnt/abc/test.sh",
				},
			},
		},
	})
	assert.Equal(t, 1, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_dbfs_file[<unknown>] (id: /mnt/abc/test.sh)"])
}

func TestAddAwsMounts(t *testing.T) {
	ic := importContextForTest()
	ic.mountMap = map[string]mount{}
	ic.addAwsMounts("abc", map[string]string{
		"foo": "bar",
		"baz": "🙄",
	})
	assert.Equal(t, 2, len(ic.mountMap))
}

var (
	userListIdUsernameFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/Users?attributes=id%2CuserName&count=100&startIndex=1",
		Response: iam.ListUsersResponse{
			Resources: []iam.User{
				{
					Id:       "id",
					UserName: "user@domain.com",
				},
			},
			TotalResults: 1,
			StartIndex:   1,
		},
		ReuseRequest: true,
	}
	userListIdUsernameFixture2 = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/Users?attributes=id%2CuserName&count=100&startIndex=2",
		Response: iam.ListUsersResponse{
			Resources:    []iam.User{},
			TotalResults: 1,
			StartIndex:   2,
		},
		ReuseRequest: true,
	}
	userListFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/Users?attributes=userName%2Cid&startIndex=1",
		Response: scim.User{
			ID:       "id",
			UserName: "user@domain.com",
		},
	}
	userReadFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/Users/id?attributes=id,userName,displayName,active,externalId,entitlements,groups,roles",
		Response: iam.User{
			Id:       "id",
			UserName: "user@domain.com",
		},
	}
	spListIdUsernameFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?attributes=id%2CuserName&count=100&startIndex=1",
		Response: iam.ListServicePrincipalResponse{
			Resources: []iam.ServicePrincipal{
				{
					Id:            "id",
					ApplicationId: "21aab5a7-ee70-4385-34d4-a77278be5cb6",
				},
			},
			TotalResults: 1,
			StartIndex:   1,
		},
	}
	spListIdUsernameFixture2 = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?attributes=id%2CuserName&count=100&startIndex=2",
		Response: iam.ListServicePrincipalResponse{
			Resources:    []iam.ServicePrincipal{},
			TotalResults: 1,
			StartIndex:   2,
		},
	}
	spListFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?attributes=id%2CuserName&startIndex=1",
		Response: scim.UserList{
			Resources: []scim.User{
				{
					ID:            "id",
					ApplicationID: "21aab5a7-ee70-4385-34d4-a77278be5cb6",
				},
			},
		},
	}
	spReadFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/id?attributes=userName,displayName,active,externalId,entitlements,groups,roles",
		Response: iam.ServicePrincipal{
			Id:            "id",
			ApplicationId: "21aab5a7-ee70-4385-34d4-a77278be5cb6",
		},
	}
)

func TestEmitUser(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		assert.True(t, len(ic.testEmits) == 0)
		ic.emitUserOrServicePrincipal("user@domain.com")
		assert.True(t, len(ic.testEmits) == 1)
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (id: id)"])
	})
}

func TestEmitServicePrincipal(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		spListIdUsernameFixture,
		spListIdUsernameFixture2,
		spListFixture,
		spReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.emitUserOrServicePrincipal("21aab5a7-ee70-4385-34d4-a77278be5cb6")
		assert.True(t, len(ic.testEmits) == 1)
		assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (id: id)"])
	})
}

func TestEmitUserError(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?attributes=id%2CuserName&count=100&startIndex=1",
			Response: iam.ListUsersResponse{
				Resources: []iam.User{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.emitUserOrServicePrincipal("abc")
		assert.True(t, len(ic.testEmits) == 0)
	})
}

func TestEmitUserOrServicePrincipalForPath(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.emitUserOrServicePrincipalForPath("/Users/user@domain.com/abc", "/Users")
		assert.True(t, len(ic.testEmits) == 1)
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (id: id)"])
	})
}

func TestEmitUserOrServicePrincipalForPath_NoEmit(t *testing.T) {
	// Negative cases
	ic := importContextForTest()
	ic.emitUserOrServicePrincipalForPath("/Shared/abc", "/Users")
	assert.True(t, len(ic.testEmits) == 0)

	ic = importContextForTest()
	ic.emitUserOrServicePrincipalForPath("/Users/", "/Users")
	assert.True(t, len(ic.testEmits) == 0)
}

func TestEmitNotebookOrRepo(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.emitNotebookOrRepo("/Users/user@domain.com/abc")
		assert.True(t, len(ic.testEmits) == 1)
		assert.True(t, ic.testEmits["databricks_notebook[<unknown>] (id: /Users/user@domain.com/abc)"])
	})

	// test for repository
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.emitNotebookOrRepo("/Repos/user@domain.com/repo/abc")
		assert.True(t, len(ic.testEmits) == 1)
		assert.True(t, ic.testEmits["databricks_repo[<unknown>] (path: /Repos/user@domain.com/repo)"])
	})
}

func TestIsUserOrServicePrincipalDirectory(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		result_false_partslength_more_than_3 := ic.IsUserOrServicePrincipalDirectory("/Users/user@domain.com/abc", "/Users")
		assert.False(t, result_false_partslength_more_than_3)
	})

	ic := importContextForTest()
	result_false_partslength_less_than_3 := ic.IsUserOrServicePrincipalDirectory("/Users", "/Users")
	assert.False(t, result_false_partslength_less_than_3)

	ic = importContextForTest()
	result_false_part2_empty := ic.IsUserOrServicePrincipalDirectory("/Users/", "/Users")
	assert.False(t, result_false_part2_empty)

	ic = importContextForTest()
	result_false_notprefix_with_user := ic.IsUserOrServicePrincipalDirectory("/Shared", "/Users")
	assert.False(t, result_false_notprefix_with_user)

	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		result_true_user_directory := ic.IsUserOrServicePrincipalDirectory("/Users/user@domain.com", "/Users")
		assert.True(t, result_true_user_directory)
	})

	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userListIdUsernameFixture,
		userListIdUsernameFixture2,
		userListFixture,
		userReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		result_true_user_directory := ic.IsUserOrServicePrincipalDirectory("/Users/user@domain.com/", "/Users")
		assert.True(t, result_true_user_directory)
	})

	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		spListIdUsernameFixture,
		spListIdUsernameFixture2,
		spListFixture,
		spReadFixture,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		result_true_sp_directory := ic.IsUserOrServicePrincipalDirectory("/Users/21aab5a7-ee70-4385-34d4-a77278be5cb6", "/Users")
		assert.True(t, result_true_sp_directory)
	})
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("a", "10")
	assert.Equal(t, 10, getEnvAsInt("a", 1))
	//
	os.Setenv("a", "abc")
	assert.Equal(t, 1, getEnvAsInt("a", 1))
	//
	assert.Equal(t, 1, getEnvAsInt("b", 1))
}

func TestExcludeAuxiliaryDirectories(t *testing.T) {
	assert.True(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{Path: "", ObjectType: workspace.Directory}))
	assert.True(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{ObjectType: workspace.File}))
	assert.True(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{Path: "/Users/user@domain.com/abc",
		ObjectType: workspace.Directory}))
	// should be ignored
	assert.False(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{Path: "/Users/user@domain.com/.ide",
		ObjectType: workspace.Directory}))
	assert.False(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{Path: "/Shared/.bundle",
		ObjectType: workspace.Directory}))
	assert.False(t, excludeAuxiliaryDirectories(workspace.ObjectStatus{Path: "/Users/user@domain.com/abc/__pycache__",
		ObjectType: workspace.Directory}))
}
