package common

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource aims to simplify things like error & deleted entities handling
type Resource struct {
	Create         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Read           func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Update         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Delete         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	CustomizeDiff  func(ctx context.Context, d *schema.ResourceDiff, c interface{}) error
	StateUpgraders []schema.StateUpgrader
	Schema         map[string]*schema.Schema
	SchemaVersion  int
	Timeouts       *schema.ResourceTimeout
}

// ToResource converts to Terraform resource definition
func (r Resource) ToResource() *schema.Resource {
	var update func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
	if r.Update != nil {
		update = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			c := m.(*DatabricksClient)
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
		err := r.Read(ctx, d, m.(*DatabricksClient))
		if IsMissing(err) {
			log.Printf("[INFO] %s[id=%s] is removed on backend",
				ResourceName.GetOrUnknown(ctx), d.Id())
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
		CustomizeDiff:  r.CustomizeDiff,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			c := m.(*DatabricksClient)
			err := r.Create(ctx, d, c)
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
			if err := r.Delete(ctx, d, m.(*DatabricksClient)); err != nil {
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

func MustCompileKeyRE(name string) *regexp.Regexp {
	regexFromName := strings.ReplaceAll(name, ".", "\\.")
	regexFromName = strings.ReplaceAll(regexFromName, ".0", ".\\d+")
	return regexp.MustCompile(regexFromName)
}

func makeEmptyBlockSuppressFunc(name string) func(k, old, new string, d *schema.ResourceData) bool {
	// TODO: move to TF field tag
	re := MustCompileKeyRE(name)
	return func(k, old, new string, d *schema.ResourceData) bool {
		log.Printf("[DEBUG] name=%s k='%v', old='%v', new='%v'", name, k, old, new)
		if re.Match([]byte(name)) && old == "1" && new == "0" {
			log.Printf("[DEBUG] Disable removal of empty block")
			return true
		}
		return false
	}
}
