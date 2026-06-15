package common

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestUserAgentAccumulatesWhenContextReused reproduces the runaway User-Agent
// observed on the heatseeker probe (extremely long headers like
// "... sdk/sdkv2 resource/directory sdk/sdkv2 resource/directory ...", repeated
// hundreds of times, which trip context-integrity checks).
//
// addContext appends "sdk/sdkv2" and "resource/<name>" to the context's
// user-agent data (via useragent.InContext, which appends without dedup). The
// normal Terraform flow hands each CRUD a fresh per-RPC context, so exactly one
// pair is added. But when a caller REUSES the enriched context across operations
// (a long-lived process threading one context), the pairs accumulate linearly.
func TestUserAgentAccumulatesWhenContextReused(t *testing.T) {
	var seen context.Context
	inner := op(func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		seen = ctx // the context a resource's CRUD would use for its SDK calls
		return nil
	})
	// Exactly how addContextToResource wraps a resource's ReadContext.
	wrapped := inner.addContext(ResourceName, "directory").addContext(Sdk, sdkName)

	const want = "sdk/sdkv2 resource/directory"

	// Fresh context per operation (correct usage): stays bounded at one pair.
	freshBase := context.Background()
	wrapped(freshBase, nil, nil)
	if got := strings.Count(useragent.FromContext(seen), want); got != 1 {
		t.Fatalf("fresh-context-per-op should be bounded at 1 pair, got %d", got)
	}

	// Reused context (the bug): the pair accumulates once per operation.
	ctx := context.Background()
	const n = 50
	for i := 0; i < n; i++ {
		wrapped(ctx, nil, nil)
		ctx = seen // reuse the enriched context for the next operation
	}
	ua := useragent.FromContext(ctx)
	got := strings.Count(ua, want)
	t.Logf("reproduced runaway User-Agent: %q repeated %d times (uaLen=%d)", want, got, len(ua))
	if got != n {
		t.Fatalf("expected %d accumulated pairs, got %d:\n%s", n, got, ua)
	}
}
