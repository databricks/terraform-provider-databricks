package util

import (
	"context"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CommonResource aims to simplify things like error & deleted entities handling
type CommonResource struct {
	Create         func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Read           func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Update         func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	Delete         func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error
	StateUpgraders []schema.StateUpgrader
	Schema         map[string]*schema.Schema
	SchemaVersion  int
	Timeouts       *schema.ResourceTimeout
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
		// set ForceNew to all attributes with CRD
		for _, v := range r.Schema {
			if v.Computed {
				continue
			}
			v.ForceNew = true
		}
	}
	read := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		err := r.Read(ctx, d, m.(*common.DatabricksClient))
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] %s[id=%s] is removed on backend",
				common.ResourceName.GetOrUnknown(ctx), d.Id())
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:         r.Schema,
		SchemaVersion:  r.SchemaVersion,
		StateUpgraders: r.StateUpgraders,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			c := m.(*common.DatabricksClient)
			err := r.Create(ctx, d, c)
			if e, ok := err.(common.APIError); ok && e.IsMissing() {
				log.Printf("[INFO] %s[id=%s] is removed on backend",
					common.ResourceName.GetOrUnknown(ctx), d.Id())
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
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			if err := r.Delete(ctx, d, m.(*common.DatabricksClient)); err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) (data []*schema.ResourceData, e error) {
				d.MarkNewResource()
				diags := read(ctx, d, m)
				var err error
				if diags.HasError() {
					err = diags[0].Validate()
				}
				return []*schema.ResourceData{d}, err
			},
		},
		Timeouts: r.Timeouts,
	}
}
