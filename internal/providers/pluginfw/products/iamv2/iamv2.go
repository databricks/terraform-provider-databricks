package iamv2

import (
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resolveApiLevel decides whether a resolve-by-external-id data source should
// call the account API or the workspace (proxy) API. It mirrors
// common.validateApiLevelForUnifiedHost (common/unified_provider.go) for the
// plugin framework: an explicit `api` argument always wins; otherwise the
// choice is inferred from the provider's configured host, except on a
// UnifiedHost, where the host alone is ambiguous and `api` must be set.
func resolveApiLevel(client *common.DatabricksClient, api types.String) (isAccount bool, diags diag.Diagnostics) {
	if !api.IsNull() && !api.IsUnknown() {
		apiLevel := api.ValueString()
		if apiLevel != "account" && apiLevel != "workspace" {
			diags.AddError("Invalid api value", "api must be either \"account\" or \"workspace\"")
			return false, diags
		}
		return apiLevel == "account", diags
	}
	if client.HostTypeForTerraform() == config.UnifiedHost {
		diags.AddError("Missing api value", "please set api to account or workspace")
		return false, diags
	}
	return client.HostTypeForTerraform() == config.AccountHost, diags
}
