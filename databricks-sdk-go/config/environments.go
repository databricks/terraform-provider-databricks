package config

import (
	"strings"

	"github.com/databricks/databricks-sdk-go/common/environment"
)

// Deprecated: Use the Cloud field and cloud-specific helper methods (IsAws, IsAzure, IsGcp)
// instead. Environment() returns environment metadata including cloud type and Azure-specific
// endpoints.
func (c *Config) Environment() environment.DatabricksEnvironment {
	// Use the provided environment if specified. Tests may configure the client with different hostnames,
	// like localhost, which are not resolvable to a known environment, while needing to mock a specific environment.
	if c.DatabricksEnvironment != nil {
		return *c.DatabricksEnvironment
	}
	envs := environment.AllEnvironments()
	if c.Host == "" && c.AzureResourceID != "" {
		// azure resource ID can also be used in lieu of host by some
		// of the clients, like Terraform. However, in this case, the workspace
		// is assumed to be a production workspace.
		azureEnv := strings.ToUpper(c.AzureEnvironment)
		if azureEnv == "" {
			azureEnv = "PUBLIC"
		}
		for _, v := range envs {
			if v.Cloud != environment.CloudAzure {
				continue
			}
			if v.AzureEnvironment.Name != azureEnv {
				continue
			}
			if strings.HasPrefix(v.DnsZone, ".dev") || strings.HasPrefix(v.DnsZone, ".staging") {
				// we can't support host-less Azure CLI auth for dev & staging environments, as users will get errors like
				// ... `The user or administrator has not consented to use the application with ID '...' named
				// 'Microsoft Azure CLI'.`.
				continue
			}
			return v
		}
	}
	hostname := c.CanonicalHostName()
	return environment.GetEnvironmentForHostname(hostname)
}
