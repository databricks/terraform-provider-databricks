package provider

import (
	"context"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func addContextToAllResources(p *schema.Provider) {
	v := p.TerraformVersion
	for k, r := range p.DataSourcesMap {
		addContextToResource(v, k, r)
	}
	for k, r := range p.ResourcesMap {
		addContextToResource(v, k, r)
	}
}

func addContextToResource(version, name string, r *schema.Resource) {
	name = strings.ReplaceAll(name, "databricks_", "")
	if r.CreateContext != nil {
		r.CreateContext = addContextToStage(version, name, r.CreateContext)
	}
	if r.ReadContext != nil {
		r.ReadContext = addContextToStage(version, name, r.ReadContext)
	}
	if r.UpdateContext != nil {
		r.UpdateContext = addContextToStage(version, name, r.UpdateContext)
	}
	if r.DeleteContext != nil {
		r.DeleteContext = addContextToStage(version, name, r.DeleteContext)
	}
}

func addContextToStage(version, name string,
	f func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics) func(
	ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		ctx = context.WithValue(ctx, common.ResourceName, name)
		ctx = context.WithValue(ctx, common.TerraformVersion, version)
		return f(ctx, d, m)
	}
}
