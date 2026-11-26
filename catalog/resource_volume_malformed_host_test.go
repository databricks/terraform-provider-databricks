package catalog

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// A host that can't be parsed passes provider configuration, and every volume call then
// fails in the classic client's request visitor, before anything is sent. Replay can't
// express it, because the runner always points the provider at its fixture server.
func TestVolumeCallsWithMalformedHostFailInRequestVisitor(t *testing.T) {
	forEachVolumeClient(t, testVolumeCallsWithMalformedHostFailInRequestVisitor)
}

func testVolumeCallsWithMalformedHostFailInRequestVisitor(t *testing.T, sdkGo bool) {
	const host = "https://bad host.example.com"
	const visitorErr = `failed during request visitor: parse "` + host +
		`": invalid character " " in host name`
	inner, err := client.New(&config.Config{Host: host, Token: "fixture-token"})
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)
	tests := []struct {
		name       string
		run        func(context.Context, *schema.Resource) diag.Diagnostics
		diagnostic string
	}{
		{
			name: "create",
			run: func(ctx context.Context, r *schema.Resource) diag.Diagnostics {
				d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
					"name":         "retry_volume",
					"catalog_name": "main",
					"schema_name":  "replay",
					"volume_type":  "MANAGED",
				})
				d.MarkNewResource()
				return r.CreateContext(ctx, d, c)
			},
			diagnostic: "cannot create volume: " + visitorErr,
		},
		{
			name: "refresh",
			run: func(ctx context.Context, r *schema.Resource) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(retryVolumeID)
				return r.ReadContext(ctx, d, c)
			},
			diagnostic: "cannot read volume: " + visitorErr,
		},
		{
			name: "update",
			run: func(ctx context.Context, r *schema.Resource) diag.Diagnostics {
				return r.UpdateContext(ctx, ownerAndCommentUpdateData(t, r), c)
			},
			diagnostic: "cannot update volume: " + visitorErr,
		},
		{
			name: "delete",
			run: func(ctx context.Context, r *schema.Resource) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(retryVolumeID)
				return r.DeleteContext(ctx, d, c)
			},
			diagnostic: "cannot delete volume: " + visitorErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := tt.run(volumeContext(context.Background()), volumeResource(sdkGo))
			assert.Equal(t, tt.diagnostic, diagnosticSummary(t, diags))
		})
	}
}
