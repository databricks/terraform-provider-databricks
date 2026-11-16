// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Terraform passes only the import ID (or a resource identity) to the
// provider during import, never the resource configuration. A workspace-level
// resource imported through an account-level or unified-host provider
// therefore needs its workspace ID in the import identity.
// See https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/identity.
const (
	identityIDKey          = "id"
	identityWorkspaceIDKey = "workspace_id"
)

// workspace_id is an import input only. The stored identity leaves it null,
// so provider_config.workspace_id is the only copy of the workspace ID in state.
func workspaceResourceIdentity() *schema.ResourceIdentity {
	return &schema.ResourceIdentity{
		SchemaFunc: func() map[string]*schema.Schema {
			return map[string]*schema.Schema{
				identityIDKey: {
					Type:              schema.TypeString,
					RequiredForImport: true,
					Description:       "The import ID of the resource.",
				},
				identityWorkspaceIDKey: {
					Type:              schema.TypeString,
					OptionalForImport: true,
					Description: "The ID of the workspace that contains the resource. Used only on import, " +
						"to set provider_config.workspace_id.",
				},
			}
		},
	}
}

// importerWithIdentity wraps an importer so that the import identity is copied
// into the resource data before the wrapped importer runs. The wrapped importer
// then parses d.Id() as usual.
func importerWithIdentity(inner *schema.ResourceImporter) *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: func(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
			if err := applyImportIdentity(d); err != nil {
				return nil, err
			}
			imported, err := runImporter(ctx, inner, d, m)
			if err != nil {
				return nil, err
			}
			for _, r := range imported {
				if err := setResourceIdentity(r); err != nil {
					return nil, err
				}
			}
			return imported, nil
		},
	}
}

func runImporter(ctx context.Context, inner *schema.ResourceImporter, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	if inner.StateContext != nil {
		return inner.StateContext(ctx, d, m)
	}
	if inner.State != nil {
		return inner.State(d, m)
	}
	return []*schema.ResourceData{d}, nil
}

// applyImportIdentity sets the resource ID from the identity when importing by
// identity, and sets provider_config.workspace_id from the identity when given.
// It mutates d.
func applyImportIdentity(d *schema.ResourceData) error {
	identity, err := d.Identity()
	if err != nil {
		// d has no identity schema. Terraform runs always attach one, but unit-test
		// fixtures build ResourceData without it.
		return nil
	}
	if d.Id() == "" {
		id, _ := identity.Get(identityIDKey).(string)
		if id == "" {
			return fmt.Errorf("import identity must set %q to the import ID of the resource", identityIDKey)
		}
		d.SetId(id)
	}
	workspaceID, _ := identity.Get(identityWorkspaceIDKey).(string)
	if workspaceID == "" {
		return nil
	}
	return d.Set("provider_config", []map[string]any{{"workspace_id": workspaceID}})
}

// setResourceIdentity sets the identity id to the resource ID and clears the
// import-only workspace_id. It mutates d.
func setResourceIdentity(d *schema.ResourceData) error {
	identity, err := d.Identity()
	if err != nil {
		// No identity schema on d, as in applyImportIdentity. In Terraform runs the
		// SDK rejects a missing identity, so this cannot hide a real failure.
		return nil
	}
	if err := identity.Set(identityIDKey, d.Id()); err != nil {
		return err
	}
	return identity.Set(identityWorkspaceIDKey, nil)
}
