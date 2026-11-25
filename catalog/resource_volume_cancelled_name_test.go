package catalog

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A call whose URL can't be built, for a name with a "%" that starts no valid escape,
// fails in the rate limiter when its context is already cancelled, because the client
// waits there before it builds the request. Nothing is sent.
func TestVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled(t *testing.T) {
	forEachVolumeClient(t, testVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled)
}

func testVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled(t *testing.T,
	sdkGo bool) {
	const limiterErr = "failed in rate limiter: context canceled"
	tests := []struct {
		name       string
		run        func(context.Context, *schema.Resource, any) diag.Diagnostics
		diagnostic string
	}{
		{
			name: "refresh",
			run: func(ctx context.Context, r *schema.Resource, c any) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(percentVolumeID)
				return r.ReadContext(ctx, d, c)
			},
			diagnostic: "cannot read volume: " + limiterErr,
		},
		{
			name: "update",
			run: func(ctx context.Context, r *schema.Resource, c any) diag.Diagnostics {
				state := &terraform.InstanceState{ID: percentVolumeID,
					Attributes: map[string]string{
						"id":           percentVolumeID,
						"name":         "50%_raw",
						"catalog_name": "main",
						"schema_name":  "replay",
						"volume_type":  "MANAGED",
					}}
				diff := &terraform.InstanceDiff{
					Attributes: map[string]*terraform.ResourceAttrDiff{
						"comment": {Old: "", New: "percent"},
					}}
				d, err := schema.InternalMap(r.Schema).Data(state, diff)
				require.NoError(t, err)
				return r.UpdateContext(ctx, d, c)
			},
			diagnostic: "cannot update volume: " + limiterErr,
		},
		{
			name: "delete",
			run: func(ctx context.Context, r *schema.Resource, c any) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(percentVolumeID)
				return r.DeleteContext(ctx, d, c)
			},
			diagnostic: "cannot delete volume: " + limiterErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, server := newScriptedVolumeServer(t)
			c := newVolumeClient(t, &config.Config{Host: server.URL})
			ctx, cancel := context.WithCancel(volumeContext(context.Background()))
			cancel()

			diags := tt.run(ctx, volumeResource(sdkGo), c)

			assert.Equal(t, tt.diagnostic, diagnosticSummary(t, diags))
			assertVolumeRequests(t, nil, api)
		})
	}
}
