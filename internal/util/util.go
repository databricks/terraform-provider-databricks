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
}

// NewPairID creates new ID pair
func NewPairID(left, right string) *Pair {
	return &Pair{left, right}
}

// Schema of paired fields
func (p *Pair) Schema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{}
	s[p.Left] = &schema.Schema{Type: schema.TypeString, ForceNew: true, Required: true}
	s[p.Right] = &schema.Schema{Type: schema.TypeString, ForceNew: true, Required: true}
	return s
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
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return p.ReadContext(d, func(left, right string) error {
			return pr.ReadContext(ctx, left, right, m.(*common.DatabricksClient))
		})
	}
	return &schema.Resource{
		Schema:      p.Schema(),
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			left := d.Get(p.Left).(string)
			if left == "" {
				return diag.Errorf("%s cannot be empty", p.Left)
			}
			right := d.Get(p.Right).(string)
			if right == "" {
				return diag.Errorf("%s cannot be empty", p.Right)
			}
			err := pr.CreateContext(ctx, left, right, m.(*common.DatabricksClient))
			if err != nil {
				return diag.FromErr(err)
			}
			p.Pack(d)
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			left, right, err := p.Unpack(d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = pr.DeleteContext(ctx, left, right, m.(*common.DatabricksClient))
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
