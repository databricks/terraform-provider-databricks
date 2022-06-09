package common

import (
	"context"
	"fmt"
	"log"
	"reflect"
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

func nicerError(ctx context.Context, err error, action string) error {
	name := ResourceName.GetOrUnknown(ctx)
	if name == "unknown" {
		return err
	}
	return fmt.Errorf("cannot %s %s: %w", action,
		strings.ReplaceAll(name, "_", " "), err)
}

func recoverable(cb func(
	ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error) func(
	ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
	return func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) (err error) {
		defer func() {
			// this is deliberate decision to convert a panic into error,
			// so that any unforeseen bug would we visible to end-user
			// as an error and not a provider crash, which is way less
			// of pleasant experience.
			if panic := recover(); panic != nil {
				err = fmt.Errorf("panic: %v", panic)
			}
		}()
		err = cb(ctx, d, c)
		return
	}
}

// ToResource converts to Terraform resource definition
func (r Resource) ToResource() *schema.Resource {
	var update func(ctx context.Context, d *schema.ResourceData,
		m interface{}) diag.Diagnostics
	if r.Update != nil {
		update = func(ctx context.Context, d *schema.ResourceData,
			m interface{}) diag.Diagnostics {
			c := m.(*DatabricksClient)
			if err := recoverable(r.Update)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "update")
				return diag.FromErr(err)
			}
			if err := recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		}
	} else {
		// set ForceNew to all attributes with CRD
		queue := []*schema.Resource{
			{Schema: r.Schema},
		}
		for {
			head := queue[0]
			queue = queue[1:]
			for _, v := range head.Schema {
				if v.Computed {
					continue
				}
				if nested, ok := v.Elem.(*schema.Resource); ok {
					queue = append(queue, nested)
				}
				v.ForceNew = true
			}
			if len(queue) == 0 {
				break
			}
		}
	}
	read := func(ctx context.Context, d *schema.ResourceData,
		m interface{}) diag.Diagnostics {
		err := recoverable(r.Read)(ctx, d, m.(*DatabricksClient))
		if IsMissing(err) {
			log.Printf("[INFO] %s[id=%s] is removed on backend",
				ResourceName.GetOrUnknown(ctx), d.Id())
			d.SetId("")
			return nil
		}
		if err != nil {
			err = nicerError(ctx, err, "read")
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:         r.Schema,
		SchemaVersion:  r.SchemaVersion,
		StateUpgraders: r.StateUpgraders,
		CustomizeDiff:  r.CustomizeDiff,
		CreateContext: func(ctx context.Context, d *schema.ResourceData,
			m interface{}) diag.Diagnostics {
			c := m.(*DatabricksClient)
			err := recoverable(r.Create)(ctx, d, c)
			if err != nil {
				err = nicerError(ctx, err, "create")
				return diag.FromErr(err)
			}
			if err = recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		},
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData,
			m interface{}) diag.Diagnostics {
			err := recoverable(r.Delete)(ctx, d, m.(*DatabricksClient))
			if IsMissing(err) {
				log.Printf("[INFO] %s[id=%s] is removed on backend",
					ResourceName.GetOrUnknown(ctx), d.Id())
				d.SetId("")
				return nil
			}
			if err != nil {
				err = nicerError(ctx, err, "delete")
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData,
				m interface{}) (data []*schema.ResourceData, e error) {
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
	re := MustCompileKeyRE(name)
	return func(k, old, new string, d *schema.ResourceData) bool {
		if re.Match([]byte(name)) && old == "1" && new == "0" {
			log.Printf("[DEBUG] Suppressing diff for name=%s k=%#v platform=%#v config=%#v", name, k, old, new)
			return true
		}
		return false
	}
}

func DataResource(sc interface{}, read func(context.Context, interface{}, *DatabricksClient) error) *schema.Resource {
	// TODO: migrate to go1.18 and get schema from second function argument?..
	s := StructToSchema(sc, func(m map[string]*schema.Schema) map[string]*schema.Schema { return m })
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
			defer func() {
				// using recoverable() would cause more complex rewrapping of DataToStructPointer & StructToData
				if panic := recover(); panic != nil {
					diags = diag.Errorf("panic: %v", panic)
				}
			}()
			ptr := reflect.New(reflect.ValueOf(sc).Type())
			DataToReflectValue(d, &schema.Resource{Schema: s}, ptr.Elem())
			err := read(ctx, ptr.Interface(), m.(*DatabricksClient))
			if err != nil {
				err = nicerError(ctx, err, "read data")
				diags = diag.FromErr(err)
			}
			StructToData(ptr.Elem().Interface(), s, d)
			d.SetId("_")
			return
		},
	}
}
