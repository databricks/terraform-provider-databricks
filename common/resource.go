// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// workspaceIDSchemaKey is the key for the workspace ID in schema
const workspaceIDSchemaKey = "provider_config.0.workspace_id"

func workspaceIDFromRawConfig(d *schema.ResourceData) (string, bool) {
	path := cty.Path{
		cty.GetAttrStep{Name: "provider_config"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: "workspace_id"},
	}
	rawValue, diags := d.GetRawConfigAt(path)
	if diags.HasError() || rawValue.IsNull() || !rawValue.IsKnown() {
		return "", false
	}
	if rawValue.Type() == cty.String {
		return rawValue.AsString(), true
	}
	return "", false
}

// populateProviderConfigInState writes the effective workspace ID into
// provider_config in the resource state.
//
// During refresh reads (terraform plan), the prior state value must be preserved
// so that CustomizeDiff can compare the old effective workspace ID against the
// new one and trigger ForceNew when they differ. If this hook resolved the "new
// effective" workspace ID and wrote it during refresh, it would overwrite the old
// value before CustomizeDiff runs, making workspace-change detection impossible.
//
// Therefore this hook only resolves from provider-level sources (workspace_id,
// host) on the first time — when no workspace ID exists in state yet (after Create).
// On subsequent reads, it preserves whatever is already in state.
func populateProviderConfigInState(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
	// If provider_config.workspace_id already exists in state, preserve it.
	// During refresh reads the state value reflects the workspace the Read
	// actually targeted. Overwriting it would prevent CustomizeDiff from
	// detecting workspace changes.
	if existing := d.Get(workspaceIDSchemaKey); existing != nil {
		if existingStr, ok := existing.(string); ok && existingStr != "" {
			if err := d.Set("provider_config", []map[string]any{{"workspace_id": existingStr}}); err != nil {
				return fmt.Errorf("failed to set provider_config in state: %w", err)
			}
			return nil
		}
	}

	// No workspace ID in state yet (first time — after Create/Import).
	// Resolve from provider config to populate state for the first time:
	// 1. provider_config.workspace_id from raw config
	// 2. workspace_id from provider
	// 3. Lazy resolution via CurrentWorkspaceID API call (GET /api/2.0/preview/scim/v2/Me)
	//
	// Account hosts without workspace_id never reach here: plan-time validation
	// rejects them, and dual resources at account level are guarded by the api
	// field early return above.
	wsID, _ := workspaceIDFromRawConfig(d)
	if wsID == "" && c.DatabricksClient != nil && c.Config != nil {
		wsID = c.Config.WorkspaceID
	}
	if wsID == "" && c.DatabricksClient != nil {
		resolvedID, err := c.CurrentWorkspaceID(ctx)
		if err != nil {
			return fmt.Errorf("failed to resolve workspace_id: %w", err)
		}
		wsID = strconv.FormatInt(resolvedID, 10)
	}

	if wsID != "" {
		if err := d.Set("provider_config", []map[string]any{{"workspace_id": wsID}}); err != nil {
			return fmt.Errorf("failed to set provider_config in state: %w", err)
		}
	}
	return nil
}

// Resource aims to simplify things like error & deleted entities handling
type Resource struct {
	Create                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Read                            func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Update                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Delete                          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	CustomizeDiff                   func(ctx context.Context, d *schema.ResourceDiff, c *DatabricksClient) error
	StateUpgraders                  []schema.StateUpgrader
	Schema                          map[string]*schema.Schema
	SchemaVersion                   int
	Timeouts                        *schema.ResourceTimeout
	DeprecationMessage              string
	Importer                        *schema.ResourceImporter
	CanSkipReadAfterCreateAndUpdate func(d *schema.ResourceData) bool
	// IsDual marks this resource as a dual resource that can operate at both
	// account and workspace level (has an "api" field from AddApiField).
	// When true, workspace-tracking logic is skipped for account-level operations.
	IsDual bool
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
	return func(ctx context.Context, rd *schema.ResourceDiff, m any) (err error) {
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
		c, ok := m.(*DatabricksClient)
		if !ok {
			return nicerError(ctx, fmt.Errorf("expected *DatabricksClient, got %T", m), "customize diff for")
		}
		err = r.CustomizeDiff(ctx, rd, c)
		if err != nil {
			err = nicerError(ctx, err, "customize diff for")
		}
		return
	}
}

// ToResource converts to Terraform resource definition
func (r Resource) ToResource() *schema.Resource {
	// Check if this resource has provider_config in its schema (unified provider resource)
	_, hasProviderConfig := r.Schema["provider_config"]
	isDual := r.IsDual

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
			// No post-Read hook needed for Update: provider_config.workspace_id is
			// already in state from a previous Create, and r.Read() never touches it.
			// If the workspace ID had changed, CustomizeDiff would have triggered
			// ForceNew (destroy+create), so Update only runs when it's unchanged.
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

			c := m.(*DatabricksClient)
			err := recoverable(r.Read)(ctx, d, c)
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
			// Post-Read hook for refresh reads and imports: populate provider_config in state.
			// During normal refresh, provider_config is already in prior state and r.Read()
			// doesn't touch it, so the hook preserves the existing value (no-op).
			// During import (no prior state), the hook resolves the effective workspace ID
			// from workspace_id / cached host for the first time.
			// Dual resources at account level have no workspace to track, so skip.
			if hasProviderConfig && d.Id() != "" && !(isDual && IsAccountLevel(d, c)) {
				if hookErr := populateProviderConfigInState(ctx, d, c); hookErr != nil {
					return diag.FromErr(nicerError(ctx, hookErr, "populate provider_config for"))
				}
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
				// Resources that skip Read after Create (e.g. notebooks with SOURCE
				// format) still need provider_config populated for CustomizeDiff to
				// detect workspace changes. Populate it here since Read won't run.
				// Dual resources at account level have no workspace to track, so skip.
				if hasProviderConfig && d.Id() != "" && !(isDual && IsAccountLevel(d, c)) {
					if hookErr := populateProviderConfigInState(ctx, d, c); hookErr != nil {
						return diag.FromErr(nicerError(ctx, hookErr, "populate provider_config for"))
					}
				}
				return nil
			}
			if err = recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			// Post-Create hook: populate provider_config in state after Read.
			// Must run after Read so that Read's DatabricksClientForUnifiedProvider
			// sees an empty workspace_id and uses the original provider client
			// (which has the cached workspace ID) rather than creating a new client.
			// Dual resources at account level have no workspace to track, so skip.
			if hasProviderConfig && d.Id() != "" && !(isDual && IsAccountLevel(d, c)) {
				if hookErr := populateProviderConfigInState(ctx, d, c); hookErr != nil {
					return diag.FromErr(nicerError(ctx, hookErr, "populate provider_config for"))
				}
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
