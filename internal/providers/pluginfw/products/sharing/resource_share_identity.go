package sharing

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// The identity lets a share be imported through an account-level or
// unified-host provider by passing the workspace ID. It matches the identity of
// the SDKv2 databricks_share, so state can move between the two implementations.
var _ resource.ResourceWithIdentity = &ShareResource{}

func (r *ShareResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = tfschema.WorkspaceResourceIdentitySchema()
}
