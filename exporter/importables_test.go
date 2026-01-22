package exporter

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func importContextForTest() *importContext {
	p := sdkv2.DatabricksProvider()
	supportedResources := maps.Keys(resourcesMap)

	// Initialize Plugin Framework provider for tests
	ctx := context.Background()
	pfProvider, pfResources, pfSchemas := initializePluginFrameworkProvider(ctx)

	return &importContext{
		Importables:               resourcesMap,
		Resources:                 p.ResourcesMap,
		PluginFrameworkProvider:   pfProvider,
		PluginFrameworkResources:  pfResources,
		PluginFrameworkSchemas:    pfSchemas,
		testEmits:                 map[string]bool{},
		nameFixes:                 nameFixes,
		waitGroup:                 &sync.WaitGroup{},
		allUsers:                  map[string]scim.User{},
		allSps:                    map[string]scim.User{},
		channels:                  makeResourcesChannels(),
		oldWorkspaceObjectMapping: map[int64]string{},
		gitInfoCache:              map[string]gitInfoCacheEntry{},
		exportDeletedUsersAssets:  false,
		ignoredResources:          map[string]struct{}{},
		deletedResources:          map[string]struct{}{},
		State:                     newStateApproximation(supportedResources),
		emittedUsers:              map[string]struct{}{},
		userOrSpDirectories:       map[string]bool{},
		defaultChannel:            make(resourceChannel, defaultChannelSize),
		services:                  map[string]struct{}{},
		listing:                   map[string]struct{}{},
		tfvars:                    map[string]string{},
		noFormat:                  true,
	}
}

func importContextForTestWithClient(ctx context.Context, client *common.DatabricksClient) *importContext {
	ic := importContextForTest()
	ic.Client = client
	ic.Context = ctx
	if client.Config.HostType() == config.AccountHost {
		ic.accountClient, _ = client.AccountClient()
	} else {
		ic.workspaceClient, _ = client.WorkspaceClient()
	}
	return ic
}

// importContextForAccountTestWithClient creates an import context configured for account-level testing
func importContextForAccountTestWithClient(ctx context.Context, client *common.DatabricksClient, services string) *importContext {
	client.Config.AccountID = testAccountID
	client.Config.WithTesting()
	ic := importContextForTestWithClient(ctx, client)
	ic.enableServices(services)
	return ic
}

// Helper function to create an iterator from a slice
func createIteratorFromSlice[T any](items []T) listing.Iterator[T] {
	request := struct{}{}
	return listing.NewIterator(
		&request,
		func(ctx context.Context, req struct{}) ([]T, error) {
			return items, nil
		},
		func(resp []T) []T {
			return resp
		},
		func(resp []T) *struct{} {
			return nil // No pagination
		},
	)
}

func TestPermissions(t *testing.T) {
	p := permissions.ResourcePermissions()
	d := p.ToResource().TestResourceData()
	d.SetId("abc")
	ic := importContextForTest()
	ic.enableServices("access,users,groups")
	name := ic.Importables["databricks_permissions"].Name(ic, d)
	assert.Equal(t, "abc", name)

	d.MarkNewResource()
	err := common.StructToData(entity.PermissionsEntity{
		AccessControlList: []iam.AccessControlRequest{
			{
				UserName: "a",
			},
			{
				GroupName: "b",
			},
			{
				ServicePrincipalName: "123",
			},
		},
	}, p.Schema, d)
	assert.NoError(t, err)

	err = ic.Importables["databricks_permissions"].Import(ic, &resource{
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 3)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: a)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (display_name: b)"])
	assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (application_id: 123)"])
}

func TestSecretScope(t *testing.T) {
	d := secrets.ResourceSecretScope().ToResource().TestResourceData()
	d.Set("name", "abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_secret_scope"].Name(ic, d)
	assert.Equal(t, "abc_a9993e3647", name)
}

func TestSecretScopeListNoNameMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/scopes/list",

			Response: sdk_workspace.ListScopesResponse{
				Scopes: []sdk_workspace.SecretScope{
					{
						Name: "abc",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_secret_scope"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestGlobalInitScriptNameFromId(t *testing.T) {
	ic := importContextForTest()
	d := workspace.ResourceGlobalInitScript().ToResource().TestResourceData()
	d.SetId("abc")
	assert.Equal(t, "abc", resourcesMap["databricks_global_init_script"].Name(ic, d))
}

func TestGlobalInitScriptsErrors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_global_init_script"].List(ic)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "abc",
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestGlobalInitScriptsBodyErrors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/global-init-scripts/sad-emoji?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				Name:   "x.sh",
				Script: "🥺",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/global-init-scripts/second?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				Name:   "x.sh",
				Script: "YWJj",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "sad-emoji",
		})
		assert.EqualError(t, err, "illegal base64 data at input byte 0")

		err = resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "second",
		})
		assert.NotNil(t, err) // no exact match because of OS diffs
	})
}

func testGenerate(t *testing.T, fixtures []qa.HTTPFixture, services string, asAdmin bool, cb func(*importContext)) {
	qa.HTTPFixturesApply(t, fixtures, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.Directory = fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		os.MkdirAll(ic.Directory, 0755)
		defer os.RemoveAll(ic.Directory)
		ic.testEmits = nil
		ic.meAdmin = asAdmin
		ic.importing = map[string]bool{}
		ic.variables = map[string]string{}
		ic.enableServices(services)
		ic.startImportChannels()
		cb(ic)
	})
}

func getGeneratedFile(ic *importContext, service string) string {
	fileName := fmt.Sprintf("%s/%s.tf", ic.Directory, service)
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[ERROR] can't read file %s", fileName)
		return ""
	}
	return string(content)
}

func TestGlobalInitScriptGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/global-init-scripts/a?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				ScriptId: "a",
				Name:     "New: Importing ^ Things",
				Enabled:  true,
				Script:   "YWJj",
			},
		},
	}, "wsconf", false, func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_global_init_script",
			ID:       "a",
		})

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_global_init_script" "new_importing_things" {
		  source  = "${path.module}/global_init_scripts/new_importing_things.sh"
		  name    = "New: Importing ^ Things"
		  enabled = true
		}`), getGeneratedFile(ic, "wsconf"))
	})
}

func TestSecretGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=a",
			Response: sdk_workspace.ListSecretsResponse{
				Secrets: []sdk_workspace.SecretMetadata{
					{
						Key: "b",
					},
				},
			},
		},
	}, "secrets", false, func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_secret",
			ID:       "a|||b",
		})
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_secret" "a_b_eb2980a5a2" {
		  string_value = var.string_value_a_b_eb2980a5a2
		  scope        = "a"
		  key          = "b"
		}`), getGeneratedFile(ic, "secrets"))
	})
}

// TODO: remove it completely after we remove support for legacy dashboards
func TestSqlListObjects(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries?page_size=100",
			Response: dbsqlListResponse{PageSize: 1, Page: 1, TotalCount: 2,
				Results: []map[string]any{{"key1": "value1"}}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries?page=2&page_size=100",
			Response: dbsqlListResponse{PageSize: 1, Page: 2, TotalCount: 2,
				Results: []map[string]any{{"key2": "value2"}}},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		answer, err := dbsqlListObjects(ic, "/preview/sql/queries")
		assert.NoError(t, err)
		assert.Len(t, answer, 2)
	})
}

func TestEmitSqlParent(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("directories")
	ic.emitSqlParentDirectory("")
	assert.Equal(t, 0, len(ic.testEmits))
	ic.emitSqlParentDirectory("folders/12345")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_directory[<unknown>] (object_id: 12345)")
}

func TestEmitFilesFromSlice(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FShared%2Ftest.txt&return_git_info=true",
			Response: workspace.ObjectStatus{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FShared%2Fgit%2Ftest.txt&return_git_info=true",
			Response: workspace.ObjectStatus{
				GitInfo: &sdk_workspace.RepoInfo{
					Id: 1234,
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("storage,notebooks,wsfiles,repos")
		ic.emitFilesFromSlice([]string{
			"dbfs:/FileStore/test.txt",
			"/Workspace/Shared/test.txt",
			"/Workspace/Shared/git/test.txt",
			"nothing",
		})
		assert.Equal(t, 3, len(ic.testEmits))
		assert.Contains(t, ic.testEmits, "databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/test.txt)")
		assert.Contains(t, ic.testEmits, "databricks_workspace_file[<unknown>] (id: /Shared/test.txt)")
		assert.Contains(t, ic.testEmits, "databricks_repo[<unknown>] (id: 1234)")
	})
}

func TestImportUcVolumeFile(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/fs/files/Volumes/main/default/wheels/some.whl?",
			Response:     "test",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("storage")
		ic.currentMetastore = currentMetastoreResponse

		file_path := "/Volumes/main/default/wheels/some.whl"
		d := storage.ResourceFile().ToResource().TestResourceData()
		d.SetId(file_path)
		err := resourcesMap["databricks_file"].Import(ic, &resource{
			ID:   file_path,
			Data: d,
		})
		assert.NoError(t, err)
		assert.Equal(t, file_path, d.Get("path"))
		assert.Equal(t, "uc_files/main/default/wheels/some.whl", d.Get("source"))
		// Testing auxiliary functions
		r := &resource{
			Attribute: "databricks_file",
			Value:     file_path,
			Data:      d,
		}
		shouldOmitFunc := resourcesMap["databricks_file"].ShouldOmitField
		require.NotNil(t, shouldOmitFunc)
		scm := storage.ResourceFile().Schema
		assert.True(t, shouldOmitFunc(ic, "md5", scm["md5"], d, r))
		assert.False(t, shouldOmitFunc(ic, "path", scm["path"], d, r))

		assert.Equal(t, "main/default/wheels/some.whl_f27badf8", resourcesMap["databricks_file"].Name(nil, d))
	})
}

func sortStringsCopy(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	sort.Strings(c)
	return c
}
