package common

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/internal/providers/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const sdkName = "sdkv2"

// AddContextToAllResources ...
func AddContextToAllResources(p *schema.Provider, prefix string) {
	for k, r := range p.DataSourcesMap {
		name := strings.ReplaceAll(k, prefix+"_", "")
		wrap := op(r.ReadContext).addContext(ResourceName, name).addContext(IsData, "yes")
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
		switch k {
		case ResourceName:
			ctx = useragent.InContext(ctx, "resource", v)
		case IsData:
			ctx = useragent.InContext(ctx, "data", v)
		}
		ctx = common.SetSDKInContext(ctx, sdkName)
		ctx = context.WithValue(ctx, k, v)
		return f(ctx, d, m)
	}
}

func addContextToResource(name string, r *schema.Resource) {
	addName := func(a op) func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		return a.addContext(ResourceName, name)
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
