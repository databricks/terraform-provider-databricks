package environment

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Cloud string

const (
	CloudAWS     Cloud = "AWS"
	CloudAzure   Cloud = "Azure"
	CloudGCP     Cloud = "GCP"
	CloudUnknown Cloud = ""
)

// UnmarshalJSON automatically normalizes the cloud string when parsing JSON
func (c *Cloud) UnmarshalJSON(data []byte) error {
	var rawString string

	if err := json.Unmarshal(data, &rawString); err != nil {
		return err
	}

	*c = normalizeCloud(rawString)
	return nil
}

func normalizeCloud(cloud string) Cloud {
	switch strings.ToUpper(cloud) {
	case "AWS":
		return CloudAWS
	case "AZURE":
		return CloudAzure
	case "GCP":
		return CloudGCP
	case "":
		return CloudUnknown
	// For forward compatibility with new cloud providers that are not yet supported by the SDK.
	default:
		return Cloud(cloud)
	}
}

type DatabricksEnvironment struct {
	Cloud              Cloud
	DnsZone            string
	AzureApplicationID string
	AzureEnvironment   *azureEnvironment
}

func (de DatabricksEnvironment) DeploymentURL(name string) string {
	return fmt.Sprintf("https://%s%s", name, de.DnsZone)
}

func (de DatabricksEnvironment) AzureServiceManagementEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ServiceManagementEndpoint
}

func (de DatabricksEnvironment) AzureResourceManagerEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ResourceManagerEndpoint
}

func (de DatabricksEnvironment) AzureActiveDirectoryEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ActiveDirectoryEndpoint
}

// we default to AWS Prod environment since this case will be a hit for PVC
func DefaultEnvironment() DatabricksEnvironment {
	return DatabricksEnvironment{
		Cloud:   CloudAWS,
		DnsZone: ".cloud.databricks.com",
	}
}

var envs = []DatabricksEnvironment{
	{Cloud: CloudAWS, DnsZone: ".dev.databricks.com"},
	{Cloud: CloudAWS, DnsZone: ".staging.cloud.databricks.com"},
	{Cloud: CloudAWS, DnsZone: ".cloud.databricks.us"},
	DefaultEnvironment(),

	{Cloud: CloudAzure, DnsZone: ".dev.azuredatabricks.net", AzureApplicationID: "62a912ac-b58e-4c1d-89ea-b2dbfc7358fc", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".staging.azuredatabricks.net", AzureApplicationID: "4a67d088-db5c-48f1-9ff2-0aace800ae68", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".azuredatabricks.net", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".databricks.azure.us", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzureUsGovernmentCloud},
	{Cloud: CloudAzure, DnsZone: ".databricks.azure.cn", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzureChinaCloud},

	{Cloud: CloudGCP, DnsZone: ".dev.gcp.databricks.com"},
	{Cloud: CloudGCP, DnsZone: ".staging.gcp.databricks.com"},
	{Cloud: CloudGCP, DnsZone: ".gcp.databricks.com"},
}

func AllEnvironments() []DatabricksEnvironment {
	return append([]DatabricksEnvironment{}, envs...)
}

func GetEnvironmentForHostname(hostname string) DatabricksEnvironment {
	for _, e := range envs {
		if strings.HasSuffix(hostname, e.DnsZone) {
			return e
		}
	}
	return DefaultEnvironment()
}
