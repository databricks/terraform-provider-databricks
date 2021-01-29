package provider

import (
	"context"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func addContextToAllResources(p *schema.Provider) {
	for k, r := range p.DataSourcesMap {
		addContextToResource(k, r)
	}
	for k, r := range p.ResourcesMap {
		addContextToResource(k, r)
	}
}

func addContextToResource(name string, r *schema.Resource) {
	name = strings.ReplaceAll(name, "databricks_", "")
	if r.CreateContext != nil {
		r.CreateContext = addContextToStage(name, r.CreateContext)
	}
	if r.ReadContext != nil {
		r.ReadContext = addContextToStage(name, r.ReadContext)
	}
	if r.UpdateContext != nil {
		r.UpdateContext = addContextToStage(name, r.UpdateContext)
	}
	if r.DeleteContext != nil {
		r.DeleteContext = addContextToStage(name, r.DeleteContext)
	}
}

func addContextToStage(name string,
	f func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics) func(
	ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		ctx = context.WithValue(ctx, common.ResourceName, name)
		return f(ctx, d, m)
	}
}
