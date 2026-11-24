package terraformtest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestStepsPreserveConfigurationAndAlwaysDestroy(t *testing.T) {
	h := httpfixture.NewHandler()
	definition := &testDefinition{Steps: []step{
		{Name: "create", Config: "first", State: &expectedState{Resources: map[string]expectedResource{"databricks_group.this": {}}}},
		{Name: "rename", Config: "second"},
	}}
	steps := testSteps(t, definition, h)
	require.Len(t, steps, 3)
	require.Equal(t, "first", steps[0].Config)
	require.Equal(t, "second", steps[1].Config)
	require.Equal(t, "second", steps[2].Config)
	require.False(t, steps[0].Destroy)
	require.True(t, steps[2].Destroy)
	require.Len(t, steps[0].ConfigStateChecks, 1)
	require.Empty(t, steps[1].ConfigStateChecks)
	require.Len(t, steps[2].ConfigStateChecks, 1)
	var resp statecheck.CheckStateResponse
	steps[2].ConfigStateChecks[0].CheckState(context.Background(), statecheck.CheckStateRequest{State: groupState()}, &resp)
	require.ErrorContains(t, resp.Error, "step 3 (destroy): expected empty state after destroy")
}

func TestStepsKeepStateChecksOptionalAndDestroyMandatory(t *testing.T) {
	for _, state := range []*expectedState{nil, {}, {Resources: map[string]expectedResource{}}} {
		steps := testSteps(t, &testDefinition{Steps: []step{{Name: "create", Config: "config", State: state}}}, httpfixture.NewHandler())
		require.Empty(t, steps[0].ConfigStateChecks)
		require.Len(t, steps[1].ConfigStateChecks, 1)
		var resp statecheck.CheckStateResponse
		steps[1].ConfigStateChecks[0].CheckState(context.Background(), statecheck.CheckStateRequest{State: groupState()}, &resp)
		require.ErrorContains(t, resp.Error, "step 2 (destroy): expected empty state after destroy")
		steps[1].ConfigStateChecks[0].CheckState(context.Background(), statecheck.CheckStateRequest{State: stateWithResources()}, &resp)
		require.NoError(t, resp.Error)
	}
}

func TestStepInstallsAndVerifiesFixtures(t *testing.T) {
	h := httpfixture.NewHandler()
	s := terraformStep(t, h, "step 1 (read)", "config", []httpfixture.Call{{
		Request:  httpfixture.Request{Method: "GET", Path: "/groups"},
		Response: httpfixture.Response{Status: 200},
	}}, nil)
	s.PreConfig()
	require.ErrorContains(t, h.Verify(), "step 1 (read): unconsumed call 1")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/groups", nil))
	require.Equal(t, 200, w.Code)
	s.PostApplyFunc()
	require.NoError(t, h.Verify())
}

func TestFixtureTransportRejectsExternalTargets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(204) }))
	defer server.Close()
	transport := fixtureTransport{host: server.URL, transport: http.DefaultTransport}
	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)
	resp, err := transport.RoundTrip(req)
	require.NoError(t, err)
	require.Equal(t, 204, resp.StatusCode)
	resp.Body.Close()
	for _, target := range []string{"https://example.invalid", "http://example.invalid"} {
		req, err := http.NewRequest("GET", target, nil)
		require.NoError(t, err)
		_, err = transport.RoundTrip(req)
		require.ErrorContains(t, err, "hermetic replay rejected request outside fixture server")
	}
}

func TestFixtureConfigIgnoresAmbientCredentials(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()
	t.Setenv("DATABRICKS_HOST", "https://example.invalid")
	t.Setenv("DATABRICKS_TOKEN", "ambient-token")
	t.Setenv("DATABRICKS_ACCOUNT_ID", "ambient-account")
	t.Setenv("DATABRICKS_CONFIG_FILE", "/nonexistent-fixture-profile")
	cfg := &config.Config{}
	require.NoError(t, fixtureConfig(server.URL)(cfg))
	require.NoError(t, cfg.EnsureResolved())
	require.Equal(t, server.URL, cfg.Host)
	require.Equal(t, "fixture-token", cfg.Token)
	require.Equal(t, "pat", cfg.AuthType)
	require.Empty(t, cfg.AccountID)
	account := &config.Config{AccountID: "fixture-account"}
	require.NoError(t, fixtureConfig(server.URL)(account))
	require.NoError(t, account.EnsureResolved())
	require.Equal(t, "fixture-account", account.AccountID)
}

func TestConfigureTerraformDisablesDownloadsAndExtraRefresh(t *testing.T) {
	// The Go test executable is enough to verify CLI selection without invoking it.
	executable, err := os.Executable()
	require.NoError(t, err)
	t.Setenv("TF_ACC_TERRAFORM_PATH", executable)
	t.Setenv("TF_ACC_TERRAFORM_VERSION", "1.0.0")
	t.Setenv("TF_ACC_REFRESH_AFTER_APPLY", "1")
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "ambient")
	t.Setenv("TF_ACC_PROVIDER_HOST", "example.invalid")
	configureTerraform(t)
	require.Equal(t, executable, os.Getenv("TF_ACC_TERRAFORM_PATH"))
	require.Empty(t, os.Getenv("TF_ACC_TERRAFORM_VERSION"))
	_, refresh := os.LookupEnv("TF_ACC_REFRESH_AFTER_APPLY")
	require.False(t, refresh)
	require.Equal(t, "1", os.Getenv("CHECKPOINT_DISABLE"))
	require.Equal(t, "databricks", os.Getenv("TF_ACC_PROVIDER_NAMESPACE"))
	require.Equal(t, "registry.terraform.io", os.Getenv("TF_ACC_PROVIDER_HOST"))
}
