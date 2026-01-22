package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
)

type azureEnvironment struct {
	Name                    string
	Domain                  string
	ActiveDirectoryEndpoint string
}

var AzurePublicCloud = azureEnvironment{
	Name:                    "PUBLIC",
	Domain:                  "core.windows.net",
	ActiveDirectoryEndpoint: "https://login.microsoftonline.com/",
}

var AzureUsGovernmentCloud = azureEnvironment{
	Name:                    "USGOVERNMENT",
	Domain:                  "core.usgovcloudapi.net",
	ActiveDirectoryEndpoint: "https://login.microsoftonline.us/",
}

var AzureChinaCloud = azureEnvironment{
	Name:                    "CHINA",
	Domain:                  "core.chinacloudapi.cn",
	ActiveDirectoryEndpoint: "https://login.chinacloudapi.cn/",
}

func azureActiveDirectoryEndpoint(cfg *config.Config) string {
	env, err := environment(cfg)
	if err != nil {
		// TODO: The error is swallowed for backward compatibility. We should
		// consider returning it to the caller.
		log.Printf("[DEBUG] Failed to get Azure Active Directory endpoint: %s", err)
		return ""
	}
	return env.ActiveDirectoryEndpoint
}

func azureDomain(cfg *config.Config) string {
	env, err := environment(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to get Azure domain: %s", err))
	}
	return env.Domain
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
