package catalog

import (
	"github.com/databricks/terraform-provider-databricks/common"
)

func ResourceLakehouseMonitor() common.Resource {
	r := ResourceQualityMonitor()
	r.DeprecationMessage = "Use `databricks_quality_monitor` instead."
	return r
}
