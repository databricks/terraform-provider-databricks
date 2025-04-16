// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Pair defines an ID pair
type Pair struct {
	left, right string
	separator   string
	schema      map[string]*schema.Schema
}

// NewPairID creates new ID pair
func NewPairID(left, right string) *Pair {
	return NewPairSeparatedID(left, right, "|")
}

// NewPairSeparatedID creates new ID pair with a custom separator
func NewPairSeparatedID(left, right, separator string) *Pair {
	return &Pair{
		left:      left,
		right:     right,
		separator: separator,
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
	parts := strings.SplitN(id, p.separator, 2)
	if len(parts) != 2 {
		d.SetId("")
		return "", "", fmt.Errorf("invalid ID: %s", id)
	}
	if parts[0] == "" {
		d.SetId("")
		return "", "", fmt.Errorf("%s cannot be empty", p.left)
	}
	if parts[1] == "" {
		d.SetId("")
		return "", "", fmt.Errorf("%s cannot be empty", p.right)
	}
	err := p.setField(d, p.left, parts[0])
	if err != nil {
		return parts[0], parts[1], err
	}
	err = p.setField(d, p.right, parts[1])
	return parts[0], parts[1], err
}

func (p *Pair) setField(d *schema.ResourceData, col, val string) error {
	if p.schema[col].Type != schema.TypeInt {
		return d.Set(col, val)
	}
	i64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	return d.Set(col, i64)
}

// Pack data attributes to ID
func (p *Pair) Pack(d *schema.ResourceData) {
	d.SetId(fmt.Sprintf("%v%s%v", d.Get(p.left), p.separator, d.Get(p.right)))
}

// BindResource defines resource with simplified functions
type BindResource struct {
	ReadContext   func(ctx context.Context, left, right string, c *DatabricksClient) error
	CreateContext func(ctx context.Context, left, right string, c *DatabricksClient) error
	DeleteContext func(ctx context.Context, left, right string, c *DatabricksClient) error
}

// BindResource creates resource that relies on binding ID pair with simple schema & importer
func (p *Pair) BindResource(pr BindResource) Resource {
	return Resource{
		Schema: p.schema,
		Read: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			left, right, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return pr.ReadContext(ctx, left, right, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			left := d.Get(p.left).(string)
			if left == "" {
				return fmt.Errorf("%s cannot be empty", p.left)
			}
			right := d.Get(p.right).(string)
			if right == "" {
				return fmt.Errorf("%s cannot be empty", p.right)
			}
			err := pr.CreateContext(ctx, left, right, c)
			if err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			left, right, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return pr.DeleteContext(ctx, left, right, c)
		},
	}
}
