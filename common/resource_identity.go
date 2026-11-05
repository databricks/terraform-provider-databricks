// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Terraform passes only the import ID (or a resource identity) to the
// provider during import, never the resource configuration. A workspace-level
// resource imported through an account-level or unified-host provider
// therefore needs its workspace ID in the identity.
// See https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/identity.
const (
	identityIDKey          = "id"
	identityWorkspaceIDKey = "workspace_id"
)

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
					Description:       "The ID of the workspace that contains the resource.",
				},
			}
		},
	}
}

// importerWithIdentity wraps an importer so that the import identity is copied
// into the resource data before the wrapped importer runs. The wrapped importer
// then parses d.Id() as usual.
func importerWithIdentity(inner *schema.ResourceImporter, isDual bool) *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: func(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
			c := m.(*DatabricksClient)
			if err := applyImportIdentity(d); err != nil {
				return nil, err
			}
			if !(isDual && IsAccountLevel(d, c)) {
				if err := requireImportWorkspaceID(d, c); err != nil {
					return nil, err
				}
			}
			if inner.StateContext != nil {
				return inner.StateContext(ctx, d, m)
			}
			if inner.State != nil {
				return inner.State(d, m)
			}
			return []*schema.ResourceData{d}, nil
		},
	}
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

// requireImportWorkspaceID fails the import when the provider cannot resolve a
// workspace by itself and the import did not supply one. Without this check,
// Read would fall back to the account client and read from the wrong API.
func requireImportWorkspaceID(d *schema.ResourceData, c *DatabricksClient) error {
	if workspaceID, _ := d.Get(workspaceIDSchemaKey).(string); workspaceID != "" {
		return nil
	}
	if c.DatabricksClient == nil || c.Config == nil {
		return nil
	}
	if c.Config.WorkspaceID != "" || c.HostTypeForTerraform() == config.WorkspaceHost {
		return nil
	}
	return fmt.Errorf("cannot import the resource with ID %q: the provider is configured for an account-level "+
		"or unified host, so the import needs the ID of the workspace that contains the resource. "+
		"Terraform does not pass provider_config to the provider during import. Use an import block "+
		"with identity = { id = %q, workspace_id = \"<workspace ID>\" } instead of id (requires Terraform 1.12 or later)",
		d.Id(), d.Id())
}

// setWorkspaceResourceIdentity sets the identity from the resource ID and
// provider_config.workspace_id. It mutates d.
func setWorkspaceResourceIdentity(d *schema.ResourceData) error {
	identity, err := d.Identity()
	if err != nil {
		// No identity schema on d, as in applyImportIdentity. In Terraform runs the
		// SDK rejects a missing identity, so this cannot hide a real failure.
		return nil
	}
	if err := identity.Set(identityIDKey, d.Id()); err != nil {
		return err
	}
	workspaceID, _ := d.Get(workspaceIDSchemaKey).(string)
	return identity.Set(identityWorkspaceIDKey, workspaceID)
}
