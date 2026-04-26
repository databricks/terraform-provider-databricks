// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import "github.com/databricks/databricks-sdk-go/config"

// HostTypeForTerraform is the TF-local bridge over the deprecated SDK
// Config.HostType() — single chokepoint for future TF-specific host-type logic.
// Once SDK removes HostType(), this is where TF-local inference lives.
func (c *DatabricksClient) HostTypeForTerraform() config.HostType {
	return c.Config.HostType()
}
