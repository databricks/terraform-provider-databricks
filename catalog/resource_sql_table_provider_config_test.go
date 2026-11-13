package catalog

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sqlTableFullName = "main.foo.bar"

const sqlTableInfoJSON = `{
	"name": "bar",
	"catalog_name": "main",
	"schema_name": "foo",
	"table_type": "EXTERNAL",
	"data_source_format": "DELTA",
	"storage_location": "s3://bucket/path",
	"comment": "managed by terraform",
	"owner": "someone",
	"table_id": "tbl-1",
	"properties": {"k": "v"},
	"columns": [{"name": "one", "type_text": "string", "type_name": "STRING", "position": 0, "nullable": true}]
}`

// accountHostRoundTripper plays the account console of an account-level
// provider: it answers every request with an HTML page and counts the
// unity-catalog table lookups that reach it.
type accountHostRoundTripper struct {
	tableCalls atomic.Int32
}

func (rt *accountHostRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if strings.Contains(r.URL.Path, "/unity-catalog/tables/") {
		rt.tableCalls.Add(1)
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"text/html"}},
		Body:       io.NopCloser(strings.NewReader("<!DOCTYPE html><html><body>console</body></html>")),
		Request:    r,
	}, nil
}

// newAccountLevelClientWithWorkspace returns a client configured like an
// account-level provider, with a workspace client for workspace "123"
// registered, and the account host transport used to count stray requests.
func newAccountLevelClientWithWorkspace(t *testing.T) (*common.DatabricksClient, *accountHostRoundTripper) {
	workspace := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.1/unity-catalog/tables/"+sqlTableFullName {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(sqlTableInfoJSON))
			return
		}
		http.NotFound(w, r)
	}))
	t.Cleanup(workspace.Close)

	rt := &accountHostRoundTripper{}
	cfg := (&config.Config{
		Host:          "https://accounts.cloud.databricks.com",
		AccountID:     "00000000-0000-0000-0000-000000000000",
		Token:         "token",
		HTTPTransport: rt,
	}).WithTesting()
	dc, err := client.New(cfg)
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: dc}
	w, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host:  workspace.URL,
		Token: "token",
	})
	require.NoError(t, err)
	c.SetWorkspaceClientForWorkspace("123", w)
	return c, rt
}

func TestResourceSqlTableRead_ProviderConfigWithAccountLevelProvider(t *testing.T) {
	c, accountHost := newAccountLevelClientWithWorkspace(t)

	r := ResourceSqlTable().ToResource()
	d := r.TestResourceData()
	d.SetId(sqlTableFullName)
	require.NoError(t, d.Set("comment", "stale"))
	require.NoError(t, d.Set("provider_config", []any{map[string]any{"workspace_id": "123"}}))

	diags := r.ReadContext(context.Background(), d, c)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, int32(0), accountHost.tableCalls.Load(),
		"the table lookup must go to the workspace, not the account host")
	assert.Equal(t, "managed by terraform", d.Get("comment"))
	assert.Equal(t, "someone", d.Get("owner"))
	assert.Equal(t, "tbl-1", d.Get("table_id"))
	assert.Equal(t, "v", d.Get("effective_properties.k"))
	assert.Equal(t, "one", d.Get("column.0.name"))
}
