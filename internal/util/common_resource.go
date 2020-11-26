package util

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CommonResource aims to simplify things like error & deleted entities handling
type CommonResource struct {
	Create func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Read   func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Update func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Delete func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Schema map[string]*schema.Schema
}

// ToResource converts to Terraform resource definition
func (r CommonResource) ToResource() *schema.Resource {
	var update func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
	if r.Update != nil {
		update = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			c := m.(*common.DatabricksClient)
			if err := r.Update(ctx, d, c); err != nil {
				return diag.FromErr(err)
			}
			if err := r.Read(ctx, d, c); err != nil {
				return diag.FromErr(err)
			}
			return nil
		}
	} else {
		// set ForceNew to all non-optional attributes
		for _, v := range r.Schema {
			if !v.Optional {
				v.ForceNew = true
			}
		}
	}
	return &schema.Resource{
		Schema: r.Schema,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			c := m.(*common.DatabricksClient)
			err := r.Create(ctx, d, c)
			if e, ok := err.(common.APIError); ok && e.IsMissing() {
				// removing missing resource
				d.SetId("")
				return nil
			}
			if err != nil {
				return diag.FromErr(err)
			}
			if err = r.Read(ctx, d, c); err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := r.Read(ctx, d, m.(*common.DatabricksClient))
			if e, ok := err.(common.APIError); ok && e.IsMissing() {
				// removing missing resource
				d.SetId("")
				return nil
			}
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		UpdateContext: update,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			if err := r.Delete(ctx, d, m.(*common.DatabricksClient)); err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
