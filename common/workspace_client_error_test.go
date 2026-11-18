package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const meNotFoundMessage = "failed to validate workspace_id: failed to get the workspace_id: " +
	"User not found"

// newMeNotFoundClient returns a workspace-level client whose SCIM /Me answers 404.
func newMeNotFoundClient(t *testing.T) *DatabricksClient {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/2.0/preview/scim/v2/Me" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		_, _ = rw.Write([]byte(`{"error_code": "NOT_FOUND", "message": "User not found"}`))
	}))
	t.Cleanup(server.Close)
	sdkClient, err := client.New(&config.Config{
		Host:  server.URL,
		Token: "dapi-test",
		HostMetadataResolver: func(ctx context.Context, host string) (*config.HostMetadata, error) {
			return nil, nil
		},
	})
	require.NoError(t, err)
	return &DatabricksClient{DatabricksClient: sdkClient}
}

func TestGetWorkspaceClientForUnifiedProvider_MeNotFoundIsNotMissing(t *testing.T) {
	c := newMeNotFoundClient(t)

	_, err := c.GetWorkspaceClientForUnifiedProvider(context.Background(), "12345")

	assert.EqualError(t, err, meNotFoundMessage)
	assert.False(t, apierr.IsMissing(err))
}

func TestMeNotFoundKeepsResourceOnReadAndDelete(t *testing.T) {
	buildWorkspaceClient := func(ctx context.Context, d *schema.ResourceData,
		c *DatabricksClient) error {
		_, err := c.WorkspaceClientUnifiedProvider(ctx, d)
		return err
	}
	r := Resource{
		Read:   buildWorkspaceClient,
		Delete: buildWorkspaceClient,
		Schema: AddNamespaceInSchema(map[string]*schema.Schema{}),
	}.ToResource()
	c := newMeNotFoundClient(t)
	ctx := context.Background()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"provider_config": []any{map[string]any{"workspace_id": "12345"}},
	})
	d.SetId("a")

	diags := r.ReadContext(ctx, d, c)
	require.True(t, diags.HasError())
	assert.Equal(t, meNotFoundMessage, diags[0].Summary)
	assert.Equal(t, "a", d.Id())

	diags = r.DeleteContext(ctx, d, c)
	require.True(t, diags.HasError())
	assert.Equal(t, meNotFoundMessage, diags[0].Summary)
	assert.Equal(t, "a", d.Id())
}
