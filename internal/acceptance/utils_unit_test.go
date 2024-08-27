package acceptance

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/terraform-provider-databricks/common"
)

type ResourceFixturePluginFramework struct {
	Steps []Step
	// Mocks the Databricks Command Execution API. This mock is loaded at the client
	// level so any command execution API requests are not sent to the server.
	CommandMock common.CommandMock
	Azure       bool
	AzureSPN    bool
	Gcp         bool
	AccountID   string
}

func (f ResourceFixturePluginFramework) RunUnitTest(t *testing.T) error {
	setDebugLogger()
	return runWithFixtureServer(t, f)
}

func (f ResourceFixturePluginFramework) setDatabricksEnvironmentForTest(client *common.DatabricksClient, host string) {
	if f.Azure || f.AzureSPN {
		client.Config.DatabricksEnvironment = &environment.DatabricksEnvironment{
			Cloud:              environment.CloudAzure,
			DnsZone:            host,
			AzureApplicationID: "azure-login-application-id",
			AzureEnvironment:   &environment.AzurePublicCloud,
		}
	} else if f.Gcp {
		client.Config.DatabricksEnvironment = &environment.DatabricksEnvironment{
			Cloud:   environment.CloudGCP,
			DnsZone: host,
		}
	} else {
		client.Config.DatabricksEnvironment = &environment.DatabricksEnvironment{
			Cloud:   environment.CloudAWS,
			DnsZone: host,
		}
	}
}
