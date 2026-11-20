// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Terraform passes only the import ID (or a resource identity) to the
// provider during import, never the resource configuration. A workspace-level
// resource imported through an account-level or unified-host provider
// therefore needs its workspace ID in the import identity.
// See https://developer.hashicorp.com/terraform/plugin/framework/resources/identity.

// workspaceResourceIdentity is the identity of a workspace-level resource. ID
// holds the same string as the import ID, so ImportState parses it like one.
// WorkspaceID is an import input only. The stored identity leaves it null, so
// provider_config.workspace_id is the only copy of the workspace ID in state.
type workspaceResourceIdentity struct {
	ID          types.String `tfsdk:"id"`
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// WorkspaceResourceIdentitySchema returns the identity schema shared by
// workspace-level resources.
func WorkspaceResourceIdentitySchema() identityschema.Schema {
	return identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
				Description:       "The import ID of the resource.",
			},
			"workspace_id": identityschema.StringAttribute{
				OptionalForImport: true,
				Description: "The ID of the workspace that contains the resource. Used only on import, " +
					"to set provider_config.workspace_id.",
			},
		},
	}
}

// ResolveImportIdentity returns the import ID and the workspace ID for an
// import. When importing by identity, the import ID is the identity id.
func ResolveImportIdentity(ctx context.Context, req resource.ImportStateRequest) (string, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	importID := req.ID
	var workspaceID string
	if req.Identity != nil && !req.Identity.Raw.IsNull() {
		var identity workspaceResourceIdentity
		diags.Append(req.Identity.Get(ctx, &identity)...)
		if diags.HasError() {
			return "", "", diags
		}
		if importID == "" {
			importID = identity.ID.ValueString()
		}
		workspaceID = identity.WorkspaceID.ValueString()
	}
	if importID == "" {
		diags.AddError("Missing import ID", `Import identity must set "id" to the import ID of the resource.`)
		return "", "", diags
	}
	return importID, workspaceID, diags
}

// SetImportWorkspaceID sets provider_config.workspace_id in the imported state
// when the import supplied one.
func SetImportWorkspaceID(ctx context.Context, state *tfsdk.State, workspaceID string) diag.Diagnostics {
	if workspaceID == "" {
		return nil
	}
	return state.SetAttribute(ctx, path.Root("provider_config"),
		ProviderConfig{WorkspaceID: types.StringValue(workspaceID)}.ToObjectValue(ctx))
}

// SetImportWorkspaceID_SdkV2 is SetImportWorkspaceID for resources migrated
// from SDKv2, whose provider_config is a list block.
func SetImportWorkspaceID_SdkV2(ctx context.Context, state *tfsdk.State, workspaceID string) diag.Diagnostics {
	if workspaceID == "" {
		return nil
	}
	providerConfig, diags := types.ListValue(ProviderConfig{}.Type(ctx),
		[]attr.Value{ProviderConfig{WorkspaceID: types.StringValue(workspaceID)}.ToObjectValue(ctx)})
	if diags.HasError() {
		return diags
	}
	return state.SetAttribute(ctx, path.Root("provider_config"), providerConfig)
}

// SetResourceIdentity sets the identity id to the import ID and leaves the
// import-only workspace_id null.
func SetResourceIdentity(ctx context.Context, identity *tfsdk.ResourceIdentity, importID string) diag.Diagnostics {
	if identity == nil {
		return nil
	}
	return identity.Set(ctx, workspaceResourceIdentity{
		ID:          types.StringValue(importID),
		WorkspaceID: types.StringNull(),
	})
}
