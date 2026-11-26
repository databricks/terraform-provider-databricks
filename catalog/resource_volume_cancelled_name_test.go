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

// unparsableNameCall is a refresh, update or delete of the volume percentVolumeID, whose
// name has a "%" that starts no valid escape, and the URL the client fails to parse.
type unparsableNameCall struct {
	action string
	url    string
	run    func(context.Context, *schema.Resource, any) diag.Diagnostics
}

func unparsableNameCalls(t *testing.T) []unparsableNameCall {
	return []unparsableNameCall{
		{
			action: "read",
			url:    volumesAPIPath + "/" + percentVolumeID + "?",
			run: func(ctx context.Context, r *schema.Resource, c any) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(percentVolumeID)
				return r.ReadContext(ctx, d, c)
			},
		},
		{
			action: "update",
			url:    volumesAPIPath + "/" + percentVolumeID,
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
		},
		{
			action: "delete",
			url:    volumesAPIPath + "/" + percentVolumeID + "?",
			run: func(ctx context.Context, r *schema.Resource, c any) diag.Diagnostics {
				d := r.TestResourceData()
				d.SetId(percentVolumeID)
				return r.DeleteContext(ctx, d, c)
			},
		},
	}
}

// Building the request fails before anything is sent, after the rate limiter.
func TestVolumeCallsWithUnparsableNameFailToBuildTheRequest(t *testing.T) {
	forEachVolumeClient(t, testVolumeCallsWithUnparsableNameFailToBuildTheRequest)
}

func testVolumeCallsWithUnparsableNameFailToBuildTheRequest(t *testing.T, sdkGo bool) {
	for _, call := range unparsableNameCalls(t) {
		t.Run(call.action, func(t *testing.T) {
			api, server := newScriptedVolumeServer(t)
			c := newVolumeClient(t, &config.Config{Host: server.URL})

			diags := call.run(volumeContext(context.Background()), volumeResource(sdkGo), c)

			assert.Equal(t, "cannot "+call.action+" volume: failed creating new request: "+
				`parse "`+call.url+`": invalid URL escape "%_r"`,
				diagnosticSummary(t, diags))
			assertVolumeRequests(t, nil, api)
		})
	}
}

// With its context already cancelled, the same call fails in the rate limiter, because
// the client waits there before it builds the request.
func TestVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled(t *testing.T) {
	forEachVolumeClient(t, testVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled)
}

func testVolumeCallsWithUnparsableNameFailInRateLimiterWhenCancelled(t *testing.T,
	sdkGo bool) {
	for _, call := range unparsableNameCalls(t) {
		t.Run(call.action, func(t *testing.T) {
			api, server := newScriptedVolumeServer(t)
			c := newVolumeClient(t, &config.Config{Host: server.URL})
			ctx, cancel := context.WithCancel(volumeContext(context.Background()))
			cancel()

			diags := call.run(ctx, volumeResource(sdkGo), c)

			assert.Equal(t, "cannot "+call.action+
				" volume: failed in rate limiter: context canceled",
				diagnosticSummary(t, diags))
			assertVolumeRequests(t, nil, api)
		})
	}
}
