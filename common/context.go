package common

import (
	"context"
	"log"
	"runtime/debug"
	"strings"
	"sync/atomic"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/internal/providers/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const sdkName = "sdkv2"

// DEBUG ONLY (do not merge): tripwire to locate runaway User-Agent growth.
//
// addContext only ever branches the context and passes the enriched copy down to
// the inner operation; it never threads the enriched copy back out. So a freshly
// dispatched CRUD/data-source wrap should see an incoming context that carries
// zero "resource/" user-agent dimensions. If the incoming context already carries
// one or more, some caller upstream is feeding an already-enriched context back
// into the wrap, which is the runaway. Dump the stack of the first such calls to
// identify that caller. Fires are capped to avoid flooding the logs.
var uaTripwireFires int32

// AddContextToAllResources ...
func AddContextToAllResources(p *schema.Provider, prefix string) {
	for k, r := range p.DataSourcesMap {
		name := strings.ReplaceAll(k, prefix+"_", "")
		wrap := op(r.ReadContext).addContext(ResourceName, name).addContext(IsData, "yes").addContext(Sdk, sdkName)
		r.ReadContext = schema.ReadContextFunc(wrap)
	}
	for k, r := range p.ResourcesMap {
		k = strings.ReplaceAll(k, prefix+"_", "")
		addContextToResource(k, r)
	}
}

// any of TF resource CRUD operation, that may need context enhancement
type op func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics

// wrap operation invokations with additional context
func (f op) addContext(k contextKey, v string) op {
	return func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		// DEBUG ONLY (do not merge): see uaTripwireFires above.
		// The -json test reporter drops passing-test output, so a plain log line
		// never reaches the CI job log. Embed the stack in a diag error and FAIL
		// the operation: failing-test output IS surfaced by the reporter. We fire
		// the first few times the incoming context already carries a resource/
		// dimension (the onset of accumulation), which pinpoints the caller that
		// re-applied the wrap to an already-enriched context.
		if ua := useragent.FromContext(ctx); strings.Count(ua, "resource/") >= 1 {
			if n := atomic.AddInt32(&uaTripwireFires, 1); n <= 6 {
				stack := string(debug.Stack())
				log.Printf("[WARN] ua-tripwire fire #%d: incoming %d resource/ dim(s) uaLen=%d k=%v v=%s\n%s",
					n, strings.Count(ua, "resource/"), len(ua), k, v, stack)
				return diag.Errorf("UA-RUNAWAY-TRIPWIRE fire #%d: incoming ctx already carries %d resource/ dim(s) (uaLen=%d) before adding k=%v v=%s; STACK:\n%s",
					n, strings.Count(ua, "resource/"), len(ua), k, v, stack)
			}
		}
		switch k {
		case ResourceName:
			ctx = useragent.InContext(ctx, "resource", v)
		case IsData:
			ctx = useragent.InContext(ctx, "data", v)
		case Sdk:
			ctx = common.SetSDKInContext(ctx, v)
		}
		ctx = context.WithValue(ctx, k, v)
		return f(ctx, d, m)
	}
}

func addContextToResource(name string, r *schema.Resource) {
	addName := func(a op) func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		return a.addContext(ResourceName, name).addContext(Sdk, sdkName)
	}
	if r.CreateContext != nil {
		r.CreateContext = addName(op(r.CreateContext))
	}
	if r.ReadContext != nil {
		r.ReadContext = addName(op(r.ReadContext))
	}
	if r.UpdateContext != nil {
		r.UpdateContext = addName(op(r.UpdateContext))
	}
	if r.DeleteContext != nil {
		r.DeleteContext = addName(op(r.DeleteContext))
	}
}
