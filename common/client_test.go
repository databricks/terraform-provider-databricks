package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func configureAndAuthenticate(dc *DatabricksClient) (*DatabricksClient, error) {
	// Skip host-metadata resolution to avoid DNS/connection hangs on non-real hosts.
	if dc.Config.HostMetadataResolver == nil {
		dc.Config.HostMetadataResolver = func(ctx context.Context, host string) (*config.HostMetadata, error) {
			return nil, nil
		}
	}
	req, err := http.NewRequest("GET", dc.Config.Host, nil)
	if err != nil {
		return dc, err
	}
	return dc, dc.Config.Authenticate(req)
}

func failsToAuthenticateWith(t *testing.T, dc *DatabricksClient, message string) {
	_, err := configureAndAuthenticate(dc)
	if dc.Config.AuthType != "" {
		log.Printf("[INFO] Auth is: %s", dc.Config.AuthType)
	}
	if assert.NotNil(t, err, "expected to have error: %s", message) {
		assert.True(t, strings.HasPrefix(err.Error(), message), "Expected to have '%s' error, but got '%s'", message, err.Error())
	}
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	t.Setenv("PATH", "testdata:/bin")
	t.Setenv("HOME", t.TempDir())
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Username: "foo",
				Password: "bar",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:     "https://localhost:443",
				Username: "foo",
				Password: "bar",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "basic", dc.Config.AuthType)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "localhost:443",
				Token: "...",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
	assert.Equal(t, "...", dc.Config.Token)
	assert.Equal(t, "https://localhost:443", dc.Config.Host)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token: "dapi345678",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:       "foo",
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
}

func TestDatabricksClientConfigure_BasicAuthDoesNotTakePrecedence(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:       "foo",
				Token:      "configured",
				Username:   "foo",
				Password:   "bar",
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	}, "validate: more than one authorization method configured: basic and pat.")
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Config.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "nohost",
			},
		},
	}, NoAuth+
		". Config: token=***, profile=nohost, config_file=testdata/.databrickscfg")
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "invalidhost",
			},
		},
	}, "resolve: testdata/.databrickscfg has no invalidhost profile configured. "+
		"Config: token=***, profile=invalidhost, config_file=testdata/.databrickscfg")
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.invalid file",
				Profile:    "invalidhost",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/az",
				Profile:    "invalidhost",
			},
		},
	}, `resolve: cannot parse config file`)
}

func TestDatabricksClient_FormatURL(t *testing.T) {
	client := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://some.host",
			},
		},
	}
	assert.Equal(t, "https://some.host/#job/123", client.FormatURL("#job/123"))
}

func TestDatabricksIsGcp(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://demo.gcp.databricks.com/",
				Token: "dapi123",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, true, dc.IsGcp())
}

func TestIsAzure_Error(t *testing.T) {
	dc := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "notoken",
			},
		},
	}
	assert.Equal(t, false, dc.IsAzure())
}

func TestClientForHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:     "https://accounts.cloud.databricks.com/",
				Username: "abc",
				Password: "bcd",
			},
		},
	})
	assert.NoError(t, err)
	assert.True(t, dc.IsAws())
	cc, err := dc.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	assert.NoError(t, err)
	assert.Equal(t, dc.Config.Username, cc.Config.Username)
	assert.Equal(t, dc.Config.Password, cc.Config.Password)
	assert.NotEqual(t, dc.Config.Host, cc.Config.Host)
}

func TestClientForHostAuthError(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "notoken",
			},
		},
	}
	_, err := c.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_NonsenseAuth(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AuthType: "nonsense",
			},
		},
	}, "default auth: auth type \"nonsense\" not found")
}

func TestGetJWTProperty_AzureCLI_SP(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AzureClientID:     "a",
				AzureClientSecret: "b",
				AzureTenantID:     "c",
				Host:              "https://adb-1232.azuredatabricks.net",
			},
		},
	}
	tid, err := aa.GetAzureJwtProperty("tid")
	assert.NoError(t, err)
	assert.Equal(t, "c", tid)
}

func TestGetJWTProperty_NonAzure(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://abc.cloud.databricks.com",
				Token: "abc",
			},
		},
	}
	_, err := aa.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't get Azure JWT token in non-Azure environment")
}

func TestGetJWTProperty_Authenticate_Fail(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")
	t.Setenv("FAIL", "yes")
	t.Setenv("ARM_TENANT_ID", "tenant-id")

	client := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://adb-1232.azuredatabricks.net",
			},
		},
	}
	_, err := client.GetAzureJwtProperty("tid")
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(),
		"default auth: cannot configure default credentials"))
}

type mockInternalUserService struct {
	count int
}

func (m *mockInternalUserService) Me(ctx context.Context) (user *iam.User, err error) {
	m.count++
	return &iam.User{
		UserName: "test",
	}, nil
}

func TestCachedMe_Me_MakesSingleRequest(t *testing.T) {
	mock := &mockInternalUserService{}
	cm := newCachedMe(mock)
	cm.Me(context.Background())
	cm.Me(context.Background())
	assert.Equal(t, 1, mock.count)
}

func TestWorkspaceClientForWorkspace_AccountAPIFails_FallsBackToDirect(t *testing.T) {
	mockAcc := mocks.NewMockAccountClient(t)
	mockWorkspacesAPI := mockAcc.GetMockWorkspacesAPI()

	// Setup the mock to return an error (e.g. no account-level access)
	mockWorkspacesAPI.EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
		WorkspaceId: 12345,
	}).Return(nil, fmt.Errorf("workspace not found"))

	// Create a DatabricksClient with the mock account client
	dc := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token: "dapi123",
			},
		},
	}
	dc.SetAccountClient(mockAcc.AccountClient)

	// When account API fails, fallback creates a direct workspace client
	workspaceClient, err := dc.WorkspaceClientForWorkspace(context.Background(), 12345)
	assert.NoError(t, err)
	assert.NotNil(t, workspaceClient)

	// Verify the client is cached
	dc.mu.Lock()
	cachedClient, exists := dc.cachedWorkspaceClients[12345]
	dc.mu.Unlock()
	assert.True(t, exists)
	assert.Equal(t, workspaceClient, cachedClient)
}

func TestWorkspaceClientForWorkspace_WorkspaceExistsNotInCache(t *testing.T) {
	mockAcc := mocks.NewMockAccountClient(t)
	mockAcc.AccountClient.Config = &config.Config{
		Host:  "https://accounts.cloud.databricks.com",
		Token: "dapi123", // Instantiating WorkspaceClient attempts authentication, this allows Configure() to complete quickly.
	}
	mockWorkspacesAPI := mockAcc.GetMockWorkspacesAPI()

	// Create a mock workspace
	mockWorkspace := &provisioning.Workspace{
		WorkspaceId:    12345,
		WorkspaceName:  "test-workspace",
		DeploymentName: "test-deployment",
	}

	// Setup the mock to return the workspace
	mockWorkspacesAPI.EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
		WorkspaceId: 12345,
	}).Return(mockWorkspace, nil)

	// Create a DatabricksClient with the mock account client
	dc := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}
	dc.SetAccountClient(mockAcc.AccountClient)

	// Call the method with the workspace ID
	workspaceClient, err := dc.WorkspaceClientForWorkspace(context.Background(), 12345)

	// Verify no error and client is returned
	assert.NoError(t, err)
	assert.NotNil(t, workspaceClient)

	// Verify the workspace client is configured correctly
	assert.Equal(t, fmt.Sprintf("https://%s.cloud.databricks.com", mockWorkspace.DeploymentName), workspaceClient.Config.Host)

	// Verify the client is cached
	dc.mu.Lock()
	cachedClient, exists := dc.cachedWorkspaceClients[12345]
	dc.mu.Unlock()

	assert.True(t, exists)
	assert.Equal(t, workspaceClient, cachedClient)
}

func TestWorkspaceClientForWorkspace_WorkspaceExistsInCache(t *testing.T) {
	// Create a mock workspace client
	mockWorkspaceClient := &databricks.WorkspaceClient{}

	// Create a DatabricksClient with the mock workspace client in cache
	dc := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}

	// Set the workspace client in cache
	dc.SetWorkspaceClientForWorkspace(12345, mockWorkspaceClient)

	// Call the method with the workspace ID
	workspaceClient, err := dc.WorkspaceClientForWorkspace(context.Background(), 12345)

	// Verify no error and the cached client is returned
	assert.NoError(t, err)
	assert.Equal(t, mockWorkspaceClient, workspaceClient)
}

func TestAddApiField_ValidValues(t *testing.T) {
	s := AddApiField(map[string]*schema.Schema{})
	apiSchema := s["api"]
	require.NotNil(t, apiSchema)
	require.NotNil(t, apiSchema.ValidateFunc)

	// "account" is valid
	_, errs := apiSchema.ValidateFunc("account", "api")
	assert.Empty(t, errs)

	// "workspace" is valid
	_, errs = apiSchema.ValidateFunc("workspace", "api")
	assert.Empty(t, errs)
}

func TestAddApiField_InvalidValues(t *testing.T) {
	s := AddApiField(map[string]*schema.Schema{})
	validateFunc := s["api"].ValidateFunc

	for _, invalid := range []string{"", "foo", "ACCOUNT", "Workspace", "acct", "ws"} {
		_, errs := validateFunc(invalid, "api")
		assert.NotEmpty(t, errs, "expected error for value %q", invalid)
	}
}

func testSchemaWithApiField() map[string]*schema.Schema {
	return AddApiField(map[string]*schema.Schema{})
}

func TestGetApiLevel_ReturnsValueWhenSet(t *testing.T) {
	d := schema.TestResourceDataRaw(t, testSchemaWithApiField(), map[string]any{
		"api": "account",
	})
	assert.Equal(t, "account", GetApiLevel(d))
}

func TestGetApiLevel_ReturnsWorkspaceWhenSet(t *testing.T) {
	d := schema.TestResourceDataRaw(t, testSchemaWithApiField(), map[string]any{
		"api": "workspace",
	})
	assert.Equal(t, "workspace", GetApiLevel(d))
}

func TestGetApiLevel_ReturnsEmptyWhenNotSet(t *testing.T) {
	d := schema.TestResourceDataRaw(t, testSchemaWithApiField(), map[string]any{})
	assert.Equal(t, "", GetApiLevel(d))
}

// testGetApiLevelFromDiff exercises GetApiLevelFromDiff by running a full schema diff
// with a CustomizeDiff that captures the result. ResourceDiff has no simple test
// constructor, so we use r.Diff() to create one through the SDK's diff pipeline.
func testGetApiLevelFromDiff(t *testing.T, apiValue string) string {
	t.Helper()
	testSchema := map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Required: true},
	}
	AddApiField(testSchema)

	var captured string
	r := &schema.Resource{
		Schema: testSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
			captured = GetApiLevelFromDiff(d)
			return nil
		},
	}

	rawConfig := map[string]interface{}{"name": "test"}
	if apiValue != "" {
		rawConfig["api"] = apiValue
	}
	rc := terraform.NewResourceConfigRaw(rawConfig)
	_, err := r.Diff(context.Background(), nil, rc, nil)
	require.NoError(t, err)
	return captured
}

func TestGetApiLevelFromDiff_ReturnsAccount(t *testing.T) {
	assert.Equal(t, "account", testGetApiLevelFromDiff(t, "account"))
}

func TestGetApiLevelFromDiff_ReturnsWorkspace(t *testing.T) {
	assert.Equal(t, "workspace", testGetApiLevelFromDiff(t, "workspace"))
}

func TestGetApiLevelFromDiff_ReturnsEmptyWhenNotSet(t *testing.T) {
	assert.Equal(t, "", testGetApiLevelFromDiff(t, ""))
}

func TestCurrentWorkspaceID_ReturnsCachedValue(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
		// No cachedWorkspaceClient — any attempt to call WorkspaceClient()
		// and make an API call would fail, proving the cache works.
	}
	c.SetCachedWorkspaceID(12345)

	id, err := c.CurrentWorkspaceID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(12345), id)

	// Second call also returns cached value.
	id2, err := c.CurrentWorkspaceID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(12345), id2)
}
