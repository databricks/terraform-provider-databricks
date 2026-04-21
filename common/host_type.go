package common

import "github.com/databricks/databricks-sdk-go/config"

// HostTypeForTerraform is the TF-local bridge over the deprecated SDK
// Config.HostType() — single chokepoint for future TF-specific host-type logic.
// Once SDK removes HostType(), this is where TF-local inference lives.
func (c *DatabricksClient) HostTypeForTerraform() config.HostType {
	if c.Config.Experimental_IsUnifiedHost {
		return config.UnifiedHost
	}
	return c.Config.HostType()
}
