package sdkv2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

func TestParseSdkGoResources(t *testing.T) {
	for value, want := range map[string][]string{
		"":                       nil,
		" , ":                    nil,
		"databricks_volume":      {"databricks_volume"},
		" databricks_volume , ,": {"databricks_volume"},
	} {
		got, err := parseSdkGoResources(value)
		require.NoError(t, err, value)
		assert.ElementsMatch(t, want, got, value)
	}
}

func TestParseSdkGoResourcesRejectsUnknownResources(t *testing.T) {
	_, err := parseSdkGoResources("databricks_volume,databricks_user")
	assert.EqualError(t, err, `USE_SDK_GO_RESOURCES: resource "databricks_user" has no `+
		`sdk-go implementation; supported resources: databricks_volume`)
}

func TestSdkGoResourcesOptionReplacesEnvironment(t *testing.T) {
	options := &sdkV2ProviderOptions{}
	WithSdkGoResources([]string{"databricks_volume"})(options)
	got, err := options.selectedSdkGoResources()
	require.NoError(t, err)
	assert.Equal(t, []string{"databricks_volume"}, got)

	WithSdkGoResources(nil)(options)
	got, err = options.selectedSdkGoResources()
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestSdkGoResourceNamesCoverRegisteredResources(t *testing.T) {
	resources := Resources()
	for _, name := range SdkGoResourceNames() {
		assert.Contains(t, resources, name)
	}
}

func TestEnableSdkGoResourcesMarksOnlySelectedResources(t *testing.T) {
	seen := map[string]bool{}
	read := func(name string) *schema.Resource {
		return &schema.Resource{ReadContext: func(ctx context.Context,
			_ *schema.ResourceData, _ any) diag.Diagnostics {
			seen[name] = common.UsesSdkGo(ctx)
			return nil
		}}
	}
	resources := map[string]*schema.Resource{
		"selected": read("selected"),
		"other":    read("other"),
	}
	enableSdkGoResources(resources, []string{"selected", "missing"})
	for _, r := range resources {
		r.ReadContext(context.Background(), nil, nil)
	}
	assert.Equal(t, map[string]bool{"selected": true, "other": false}, seen)
}

func TestDatabricksProviderRejectsUnknownSdkGoResource(t *testing.T) {
	p := DatabricksProvider(WithSdkGoResources([]string{"databricks_user"}))
	config := terraform.NewResourceConfigRaw(map[string]any{})
	diags := p.Configure(context.Background(), config)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary,
		`resource "databricks_user" has no sdk-go implementation`)
}

// sdkGoEnvChild marks the test process that TestDatabricksProviderReadsUseSdkGoResources
// starts, and names the outcome it checks.
const sdkGoEnvChild = "DATABRICKS_TF_TEST_SDK_GO_ENV_CHILD"

// TestDatabricksProviderReadsUseSdkGoResources pins the provider reading
// USE_SDK_GO_RESOURCES when no test option selects the resources. The provider reads it
// once per process, so each value runs the test binary again.
func TestDatabricksProviderReadsUseSdkGoResources(t *testing.T) {
	if want := os.Getenv(sdkGoEnvChild); want != "" {
		checkSdkGoResourcesFromEnv(t, want)
		return
	}
	for value, want := range map[string]string{
		"":                                  "existing_client",
		"databricks_volume":                 "sdk_go",
		" databricks_volume , ,":            "sdk_go",
		"databricks_volume,databricks_user": "rejected",
	} {
		t.Run(want+"/"+value, func(t *testing.T) {
			child := exec.Command(os.Args[0], "-test.v",
				"-test.run=^TestDatabricksProviderReadsUseSdkGoResources$")
			child.Env = append(os.Environ(), useSdkGoResourcesEnv+"="+value,
				sdkGoEnvChild+"="+want)
			output, err := child.CombinedOutput()
			require.NoError(t, err, "%s", output)
			assert.Contains(t, string(output),
				"--- PASS: TestDatabricksProviderReadsUseSdkGoResources", "%s", output)
		})
	}
}

func checkSdkGoResourcesFromEnv(t *testing.T, want string) {
	p := DatabricksProvider()
	if want == "rejected" {
		diags := p.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary,
			`resource "databricks_user" has no sdk-go implementation`)
		return
	}
	api := &userAgentRecorder{}
	inner, err := client.New(&config.Config{Host: "https://volume-api.test",
		Token: "fixture-token", HTTPTransport: api, Loaders: []config.Loader{noLoader{}}})
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)
	r := p.ResourcesMap["databricks_volume"]
	d := r.TestResourceData()
	d.SetId("main.replay.env_volume")

	require.Empty(t, r.ReadContext(context.Background(), d, c))

	require.Len(t, api.userAgents, 1)
	assert.Equal(t, want == "sdk_go",
		strings.Contains(api.userAgents[0], " sdk-go-uc-volumes/"), api.userAgents[0])
}

// userAgentRecorder answers the volume GET and records its User-Agent, and answers
// anything else, such as host metadata, with 404.
type userAgentRecorder struct {
	userAgents []string
}

func (u *userAgentRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	if req.URL.Path == "/api/2.1/unity-catalog/volumes/main.replay.env_volume" {
		u.userAgents = append(u.userAgents, req.Header.Get("User-Agent"))
		recorder.Header().Set("Content-Type", "application/json")
		_, _ = recorder.WriteString(`{"catalog_name": "main", "schema_name": "replay",
			"name": "env_volume", "full_name": "main.replay.env_volume",
			"volume_type": "MANAGED"}`)
	} else {
		http.NotFound(recorder, req)
	}
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

type noLoader struct{}

func (noLoader) Name() string                   { return "none" }
func (noLoader) Configure(*config.Config) error { return nil }
