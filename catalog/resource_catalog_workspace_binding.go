package catalog

import (
	"github.com/databricks/terraform-provider-databricks/common"
)

func ResourceCatalogWorkspaceBinding() common.Resource {
	r := ResourceWorkspaceBinding()
	r.DeprecationMessage = "Use `databricks_workspace_binding` instead."
	return r
}
