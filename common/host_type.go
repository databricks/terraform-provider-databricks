package common

import (
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
)

// HostTypeForTerraform determines the host type without relying on SDK-internal
// fields like resolvedHostType or isTesting. This is the Terraform provider's
// own implementation that uses only the explicit Experimental_IsUnifiedHost flag
// and URL-prefix matching.
func (c *DatabricksClient) HostTypeForTerraform() config.HostType {
	if c.Config.Experimental_IsUnifiedHost {
		return config.UnifiedHost
	}

	// Normalize the host to ensure the scheme is present before checking
	// prefixes. Profiles saved without "https://" (e.g. from user input)
	// would otherwise fail the prefix check and be misclassified as
	// workspace hosts.
	host := c.Config.Host
	if host != "" && !strings.Contains(host, "://") {
		host = "https://" + host
	}
	accountsPrefixes := []string{
		"https://accounts.",
		"https://accounts-dod.",
	}
	for _, prefix := range accountsPrefixes {
		if strings.HasPrefix(host, prefix) {
			return config.AccountHost
		}
	}

	return config.WorkspaceHost
}
