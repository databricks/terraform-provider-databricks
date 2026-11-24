package terraformtest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

const (
	testVolumeRead       = "/api/2.1/unity-catalog/volumes/main.replay.volume"
	testClassicUserAgent = "databricks-tf-provider/1.0.0 databricks-sdk-go/0.202.0 " +
		"resource/volume"
	testSdkGoUserAgent = testClassicUserAgent + " sdk-go-core/0.1.0 sdk-go-uc-volumes/0.8.0"
)

func requestWithUserAgent(method, path, userAgent string) *http.Request {
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set("User-Agent", userAgent)
	return req
}

func TestSdkGoUserAgentCheck(t *testing.T) {
	sdkGo := sdkImplementations{"databricks_volume": sdkGoImplementation}
	classic := sdkImplementations{"databricks_volume": classicSdkImplementation}
	lacks := `GET ` + testVolumeRead + ` selects sdk_go, but its User-Agent has no sdk-go ` +
		`module: "` + testClassicUserAgent + `"`
	carries := " doesn't select sdk_go, but its User-Agent has an sdk-go module: "
	for _, tc := range []struct {
		name                    string
		selection               sdkImplementations
		method, path, userAgent string
		err                     string
	}{
		{"sdk-go read with module", sdkGo, http.MethodGet, testVolumeRead,
			testSdkGoUserAgent, ""},
		{"sdk-go read without module", sdkGo, http.MethodGet, testVolumeRead,
			testClassicUserAgent, lacks},
		{"existing-path write in sdk-go step", sdkGo, http.MethodPatch, testVolumeRead,
			testClassicUserAgent, ""},
		{"write with module in sdk-go step", sdkGo, http.MethodPatch, testVolumeRead,
			testSdkGoUserAgent, "PATCH " + testVolumeRead + carries},
		{"existing-path lookup in sdk-go step", sdkGo, http.MethodGet,
			"/api/2.0/preview/scim/v2/Me", testClassicUserAgent, ""},
		{"classic read without module", classic, http.MethodGet, testVolumeRead,
			testClassicUserAgent, ""},
		{"classic read with module", classic, http.MethodGet, testVolumeRead,
			testSdkGoUserAgent, "GET " + testVolumeRead + carries},
		{"unselected read without module", nil, http.MethodGet, testVolumeRead,
			testClassicUserAgent, ""},
		{"unselected read with module", nil, http.MethodGet, testVolumeRead,
			testSdkGoUserAgent, "GET " + testVolumeRead + carries},
		{"importer read without resource", sdkGo, http.MethodGet, testVolumeRead,
			"databricks-tf-provider/1.0.0 sdk-go-uc-volumes/0.8.0", ""},
		{"other resource with module", sdkGo, http.MethodGet, testVolumeRead,
			"databricks-tf-provider/1.0.0 resource/file sdk-go-uc-volumes/0.8.0",
			"GET " + testVolumeRead + carries},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.selection.checkUserAgent(
				requestWithUserAgent(tc.method, tc.path, tc.userAgent))
			if tc.err == "" {
				assert.NoError(t, err)
				return
			}
			assert.ErrorContains(t, err, tc.err)
		})
	}
}

func TestSdkGoRequestsCoverEveryResourceWithSdkGo(t *testing.T) {
	names := sdkv2.SdkGoResourceNames()
	for _, name := range names {
		assert.NotEmpty(t, sdkGoRequests[name], name)
	}
	for name := range sdkGoRequests {
		assert.Contains(t, names, name)
	}
}

// The check fails the step that sends a request, in both directions, and destroy
// follows the selection retained with its cleanup configuration.
func TestUserAgentCheckFailsTheStepThatSendsTheRequest(t *testing.T) {
	d := &testDefinition{Steps: []step{
		{Name: "classic", Config: `provider "databricks" {}`},
		{Name: "sdk_go", Config: `provider "databricks" {}`,
			SdkImplementations: sdkImplementations{"databricks_volume": sdkGoImplementation}},
	}}
	check := newUserAgentCheck(httpfixture.NewHandler(), d)
	read := []httpfixture.Call{{
		Request:  httpfixture.Request{Method: http.MethodGet, Path: testVolumeRead},
		Response: httpfixture.Response{Status: http.StatusOK},
	}}
	serve := func(userAgent string) {
		recorder := httptest.NewRecorder()
		check.ServeHTTP(recorder,
			requestWithUserAgent(http.MethodGet, testVolumeRead, userAgent))
		assert.Equal(t, http.StatusOK, recorder.Code, "the fixtures still answer")
	}

	require.NoError(t, check.StartStep("step 1 (classic)", read))
	serve(testSdkGoUserAgent)
	assert.ErrorContains(t, check.Verify(),
		"step 1 (classic): GET "+testVolumeRead+" doesn't select sdk_go")

	assert.ErrorContains(t, check.StartStep("step 2 (sdk_go)", read), "step 1 (classic)")
	serve(testClassicUserAgent)
	assert.ErrorContains(t, check.Verify(),
		"step 2 (sdk_go): GET "+testVolumeRead+" selects sdk_go, but its User-Agent "+
			"has no sdk-go module")

	assert.ErrorContains(t, check.StartStep("step 3 (destroy)", read), "step 2 (sdk_go)")
	serve(testSdkGoUserAgent)
	assert.NoError(t, check.Verify())
}
