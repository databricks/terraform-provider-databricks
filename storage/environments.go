package storage

import (
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
)

// Reference Azure environment.
type azureEnvironment struct {
	Name                      string
	ServiceManagementEndpoint string
	ResourceManagerEndpoint   string
	ActiveDirectoryEndpoint   string
}

var AzurePublicCloud = azureEnvironment{
	Name:                      "PUBLIC",
	ServiceManagementEndpoint: "https://management.core.windows.net/",
	ResourceManagerEndpoint:   "https://management.azure.com/",
	ActiveDirectoryEndpoint:   "https://login.microsoftonline.com/",
}

var AzureUsGovernmentCloud = azureEnvironment{
	Name:                      "USGOVERNMENT",
	ServiceManagementEndpoint: "https://management.core.usgovcloudapi.net/",
	ResourceManagerEndpoint:   "https://management.usgovcloudapi.net/",
	ActiveDirectoryEndpoint:   "https://login.microsoftonline.us/",
}

var AzureChinaCloud = azureEnvironment{
	Name:                      "CHINA",
	ServiceManagementEndpoint: "https://management.core.chinacloudapi.cn/",
	ResourceManagerEndpoint:   "https://management.chinacloudapi.cn/",
	ActiveDirectoryEndpoint:   "https://login.chinacloudapi.cn/",
}

func environment(cfg *config.Config) (azureEnvironment, error) {
	switch strings.ToUpper(cfg.AzureEnvironment) {
	case "PUBLIC", "":
		return AzurePublicCloud, nil
	case "USGOVERNMENT":
		return AzureUsGovernmentCloud, nil
	case "CHINA":
		return AzureChinaCloud, nil
	}

	// If the environment is not specified, infer the environment from
	// the host.
	switch {
	case strings.HasSuffix(cfg.Host, ".dev.azuredatabricks.net"):
		return AzurePublicCloud, nil
	case strings.HasSuffix(cfg.Host, ".staging.azuredatabricks.net"):
		return AzurePublicCloud, nil
	case strings.HasSuffix(cfg.Host, ".azuredatabricks.net"):
		return AzurePublicCloud, nil
	case strings.HasSuffix(cfg.Host, ".databricks.azure.us"):
		return AzureUsGovernmentCloud, nil
	case strings.HasSuffix(cfg.Host, ".databricks.azure.cn"):
		return AzureChinaCloud, nil
	}

	return azureEnvironment{}, fmt.Errorf("unable to infer Azure environment")
}

func azureActiveDirectoryEndpoint(cfg *config.Config) (string, error) {
	env, err := environment(cfg)
	if err != nil {
		return "", err
	}
	return env.ActiveDirectoryEndpoint, nil
}

func azureDomain(cfg *config.Config) (string, error) {
	env, err := environment(cfg)
	if err != nil {
		return "", err
	}

	switch env.Name {
	case "PUBLIC":
		return "core.windows.net", nil
	case "USGOVERNMENT":
		return "core.usgovcloudapi.net", nil
	case "CHINA":
		return "core.chinacloudapi.cn", nil
	default:
		return "", fmt.Errorf("unknown Azure domain for environment: %q", env.Name)
	}
}
