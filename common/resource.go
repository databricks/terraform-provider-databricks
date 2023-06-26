package common

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource aims to simplify things like error & deleted entities handling
type Resource struct {
	Create         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Read           func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Update         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Delete         func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	CustomizeDiff  func(ctx context.Context, d *schema.ResourceDiff) error
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

func (r Resource) saferCustomizeDiff() schema.CustomizeDiffFunc {
	if r.CustomizeDiff == nil {
		return nil
	}
	return func(ctx context.Context, rd *schema.ResourceDiff, _ any) (err error) {
		defer func() {
			// this is deliberate decision to convert a panic into error,
			// so that any unforeseen bug would we visible to end-user
			// as an error and not a provider crash, which is way less
			// of pleasant experience.
			if panic := recover(); panic != nil {
				err = nicerError(ctx, fmt.Errorf("panic: %v", panic),
					"customize diff for")
			}
		}()
		// we don't propagate instance of SDK client to the diff function, because
		// authentication is not deterministic at this stage with the recent Terraform
		// versions. Diff customization must be limited to hermetic checks only anyway.
		err = r.CustomizeDiff(ctx, rd)
		if err != nil {
			err = nicerError(ctx, err, "customize diff for")
		}
		return
	}
}

// ToResource converts to Terraform resource definition
func (r Resource) ToResource() *schema.Resource {
	var update func(ctx context.Context, d *schema.ResourceData,
		m any) diag.Diagnostics
	if r.Update != nil {
		update = func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
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
	generateReadFunc := func(ignoreMissing bool) func(ctx context.Context, d *schema.ResourceData,
		m any) diag.Diagnostics {
		return func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
			err := recoverable(r.Read)(ctx, d, m.(*DatabricksClient))
			// TODO: https://github.com/databricks/terraform-provider-databricks/issues/2021
			if ignoreMissing && apierr.IsMissing(err) {
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
	}
	return &schema.Resource{
		Schema:         r.Schema,
		SchemaVersion:  r.SchemaVersion,
		StateUpgraders: r.StateUpgraders,
		CustomizeDiff:  r.saferCustomizeDiff(),
		CreateContext: func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
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
		ReadContext:   generateReadFunc(true),
		UpdateContext: update,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
			err := recoverable(r.Delete)(ctx, d, m.(*DatabricksClient))
			if apierr.IsMissing(err) {
				// TODO: https://github.com/databricks/terraform-provider-databricks/issues/2021
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
				m any) (data []*schema.ResourceData, e error) {
				d.MarkNewResource()
				diags := generateReadFunc(false)(ctx, d, m)
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

// Deprecated: migrate to WorkspaceData
func DataResource(sc any, read func(context.Context, any, *DatabricksClient) error) *schema.Resource {
	// TODO: migrate to go1.18 and get schema from second function argument?..
	s := StructToSchema(sc, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
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
			// check if the resource schema has the `id` attribute (marked with `json:"id"` in the provided structure).
			// and if yes, then use it as resource ID. If not, then use default value for resource ID (`_`)
			if _, ok := s["id"]; ok {
				d.SetId(d.Get("id").(string))
			} else {
				d.SetId("_")
			}
			return
		},
	}
}

// WorkspaceData is a generic way to define data resources in Terraform provider.
//
// Example usage:
//
//	type catalogsData struct {
//		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
//	}
//	return common.WorkspaceData(func(ctx context.Context, data *catalogsData, w *databricks.WorkspaceClient) error {
//		catalogs, err := w.Catalogs.ListAll(ctx)
//		...
//	})
func WorkspaceData[T any](read func(context.Context, *T, *databricks.WorkspaceClient) error) *schema.Resource {
	var dummy T
	s := StructToSchema(dummy, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		// `id` attribute must be marked as computed, otherwise it's not set!
		if v, ok := m["id"]; ok {
			v.Computed = true
			v.Required = false
		}
		return m
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
			defer func() {
				// using recoverable() would cause more complex rewrapping of DataToStructPointer & StructToData
				if panic := recover(); panic != nil {
					diags = diag.Errorf("panic: %v", panic)
				}
			}()
			ptr := reflect.New(reflect.ValueOf(dummy).Type())
			DataToReflectValue(d, &schema.Resource{Schema: s}, ptr.Elem())
			client := m.(*DatabricksClient)
			w, err := client.WorkspaceClient()
			if err != nil {
				err = nicerError(ctx, err, "read data")
				return diag.FromErr(err)
			}
			err = read(ctx, ptr.Interface().(*T), w)
			if err != nil {
				err = nicerError(ctx, err, "read data")
				diags = diag.FromErr(err)
			}
			StructToData(ptr.Elem().Interface(), s, d)
			// check if the resource schema has the `id` attribute (marked with `json:"id"` in the provided structure).
			// and if yes, then use it as resource ID. If not, then use default value for resource ID (`_`)
			if _, ok := s["id"]; ok {
				d.SetId(d.Get("id").(string))
			} else {
				d.SetId("_")
			}
			return
		},
	}
}

func AccountData[T any](read func(context.Context, *T, *databricks.AccountClient) error) *schema.Resource {
	var dummy T
	s := StructToSchema(dummy, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		// `id` attribute must be marked as computed, otherwise it's not set!
		if v, ok := m["id"]; ok {
			v.Computed = true
			v.Required = false
		}
		return m
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
			defer func() {
				// using recoverable() would cause more complex rewrapping of DataToStructPointer & StructToData
				if panic := recover(); panic != nil {
					diags = diag.Errorf("panic: %v", panic)
				}
			}()
			ptr := reflect.New(reflect.ValueOf(dummy).Type())
			DataToReflectValue(d, &schema.Resource{Schema: s}, ptr.Elem())
			client := m.(*DatabricksClient)
			acc, err := client.AccountClient()
			if err != nil {
				err = nicerError(ctx, err, "read data")
				return diag.FromErr(err)
			}
			err = read(ctx, ptr.Interface().(*T), acc)
			if err != nil {
				err = nicerError(ctx, err, "read data")
				diags = diag.FromErr(err)
			}
			StructToData(ptr.Elem().Interface(), s, d)
			// check if the resource schema has the `id` attribute (marked with `json:"id"` in the provided structure).
			// and if yes, then use it as resource ID. If not, then use default value for resource ID (`_`)
			if _, ok := s["id"]; ok {
				d.SetId(d.Get("id").(string))
			} else {
				d.SetId("_")
			}
			return
		},
	}
}

func EqualFoldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if strings.EqualFold(old, new) {
		log.Printf("[INFO] Suppressing diff on %s", k)
		return true
	}
	return false
}
