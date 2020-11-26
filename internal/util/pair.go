package util

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Pair defines an ID pair
type Pair struct {
	Left, Right string
	schema      map[string]*schema.Schema
}

// NewPairID creates new ID pair
func NewPairID(left, right string) *Pair {
	return &Pair{
		Left:  left,
		Right: right,
		schema: map[string]*schema.Schema{
			left:  {Type: schema.TypeString, ForceNew: true, Required: true},
			right: {Type: schema.TypeString, ForceNew: true, Required: true},
		},
	}
}

// Schema sets custom schema
func (p *Pair) Schema(do func(map[string]*schema.Schema) map[string]*schema.Schema) *Pair {
	p.schema = do(p.schema)
	return p
}

// Unpack ID into two strings and set data
func (p *Pair) Unpack(d *schema.ResourceData) (string, string, error) {
	id := d.Id()
	parts := strings.SplitN(id, "|", 2)
	if len(parts) != 2 {
		d.SetId("")
		return "", "", fmt.Errorf("Invalid ID: %s", id)
	}
	if parts[0] == "" {
		d.SetId("")
		return "", "", fmt.Errorf("%s cannot be empty", p.Left)
	}
	if parts[1] == "" {
		d.SetId("")
		return "", "", fmt.Errorf("%s cannot be empty", p.Right)
	}
	err := d.Set(p.Left, parts[0])
	if err != nil {
		return "", "", err
	}
	err = d.Set(p.Right, parts[1])
	if err != nil {
		return "", "", err
	}
	return parts[0], parts[1], nil
}

// ReadContext helper function
func (p *Pair) ReadContext(d *schema.ResourceData, do func(left, right string) error) diag.Diagnostics {
	left, right, err := p.Unpack(d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = do(left, right)
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// Pack data attributes to ID
func (p *Pair) Pack(d *schema.ResourceData) {
	d.SetId(fmt.Sprintf("%s|%s", d.Get(p.Left), d.Get(p.Right)))
}

// BindResource defines resource with simplified functions
type BindResource struct {
	ReadContext   func(ctx context.Context, left, right string, c *common.DatabricksClient) error
	CreateContext func(ctx context.Context, left, right string, c *common.DatabricksClient) error
	DeleteContext func(ctx context.Context, left, right string, c *common.DatabricksClient) error
}

// BindResource creates resource that relies on binding ID pair with simple schema & importer
func (p *Pair) BindResource(pr BindResource) *schema.Resource {
	return CommonResource{
		Schema: p.schema,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			left, right, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return pr.ReadContext(ctx, left, right, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			left := d.Get(p.Left).(string)
			if left == "" {
				return fmt.Errorf("%s cannot be empty", p.Left)
			}
			right := d.Get(p.Right).(string)
			if right == "" {
				return fmt.Errorf("%s cannot be empty", p.Right)
			}
			err := pr.CreateContext(ctx, left, right, c)
			if err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			left, right, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return pr.DeleteContext(ctx, left, right, c)
		},
	}.ToResource()
}
