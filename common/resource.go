// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource aims to simplify things like error & deleted entities handling
type Resource struct {
	Create                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Read                            func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Update                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Delete                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	CustomizeDiff                   func(ctx context.Context, d *schema.ResourceDiff) error
	StateUpgraders                  []schema.StateUpgrader
	Schema                          map[string]*schema.Schema
	SchemaVersion                   int
	Timeouts                        *schema.ResourceTimeout
	DeprecationMessage              string
	Importer                        *schema.ResourceImporter
	CanSkipReadAfterCreateAndUpdate func(d *schema.ResourceData) bool
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
			if r.CanSkipReadAfterCreateAndUpdate != nil && r.CanSkipReadAfterCreateAndUpdate(d) {
				return nil
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
	// Ignore missing for read for resources, but not for data sources.
	ignoreMissingForRead := (r.Create != nil || r.Update != nil || r.Delete != nil)
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
	resource := &schema.Resource{
		Schema:             r.Schema,
		SchemaVersion:      r.SchemaVersion,
		StateUpgraders:     r.StateUpgraders,
		CustomizeDiff:      r.saferCustomizeDiff(),
		ReadContext:        generateReadFunc(ignoreMissingForRead),
		UpdateContext:      update,
		Importer:           r.Importer,
		Timeouts:           r.Timeouts,
		DeprecationMessage: r.DeprecationMessage,
	}
	if r.Create != nil {
		resource.CreateContext = func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			c := m.(*DatabricksClient)
			err := recoverable(r.Create)(ctx, d, c)
			if err != nil {
				err = nicerError(ctx, err, "create")
				return diag.FromErr(err)
			}
			if r.CanSkipReadAfterCreateAndUpdate != nil && r.CanSkipReadAfterCreateAndUpdate(d) {
				return nil
			}
			if err = recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		}
	}
	if r.Delete != nil {
		resource.DeleteContext = func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			err := recoverable(r.Delete)(ctx, d, m.(*DatabricksClient))
			if apierr.IsMissing(err) {
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
		}
	}
	if resource.Importer == nil {
		resource.Importer = &schema.ResourceImporter{
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
		}
	}
	return resource
}

func MustCompileKeyRE(name string) *regexp.Regexp {
	regexFromName := strings.ReplaceAll(name, ".", "\\.")
	regexFromName = strings.ReplaceAll(regexFromName, ".0", ".\\d+")
	return regexp.MustCompile(regexFromName)
}

// WorkspacePathPrefixDiffSuppress suppresses diffs for workspace paths where both sides
// may or may not include the `/Workspace` prefix.
//
// This is the case for dashboards, alerts and queries where at create time, the user may include the `/Workspace`
// prefix for the `parent_path` field, but the read response will not include the prefix.
func WorkspacePathPrefixDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	const prefix = "/Workspace"
	return strings.TrimPrefix(old, prefix) == strings.TrimPrefix(new, prefix)
}

// WorkspaceOrEmptyPathPrefixDiffSuppress is similar WorkspacePathPrefixDiffSuppress but also suppresses diffs
// when the new value is empty (not specified by user).
func WorkspaceOrEmptyPathPrefixDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	const prefix = "/Workspace"
	return (old != "" && new == "") || strings.TrimPrefix(old, prefix) == strings.TrimPrefix(new, prefix)
}

func EqualFoldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if strings.EqualFold(old, new) {
		log.Printf("[INFO] Suppressing diff on %s", k)
		return true
	}
	return false
}

func NoCustomize(m map[string]*schema.Schema) map[string]*schema.Schema {
	return m
}

var NoAuth string = "default auth: cannot configure default credentials, " +
	"please check https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication " +
	"to configure credentials for your preferred authentication method"

func AddAccountIdField(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["account_id"] = &schema.Schema{
		Type:       schema.TypeString,
		Computed:   true,
		Optional:   true,
		Deprecated: "Configuring `account_id` at the resource-level is deprecated; please specify it in the `provider {}` configuration block instead",
	}
	return s
}
