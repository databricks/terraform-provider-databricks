package databricks

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/google/go-cmp/cmp"
)

func TestWorkspaceHost(t *testing.T) {
	awsEnv := environment.DatabricksEnvironment{
		Cloud:   environment.CloudAWS,
		DnsZone: ".cloud.databricks.com",
	}
	azureEnv := environment.DatabricksEnvironment{
		Cloud:   environment.CloudAzure,
		DnsZone: ".azuredatabricks.net",
	}

	testCases := []struct {
		name           string
		env            environment.DatabricksEnvironment
		accountHost    string
		deploymentName string
		want           string
	}{
		{
			name:           "standard AWS account host constructs workspace URL from deployment name",
			env:            awsEnv,
			accountHost:    "https://accounts.cloud.databricks.com",
			deploymentName: "myworkspace",
			want:           "https://myworkspace.cloud.databricks.com",
		},
		{
			name:           "standard AWS account host without scheme",
			env:            awsEnv,
			accountHost:    "accounts.cloud.databricks.com",
			deploymentName: "myworkspace",
			want:           "https://myworkspace.cloud.databricks.com",
		},
		{
			name:           "standard Azure account host constructs workspace URL from deployment name",
			env:            azureEnv,
			accountHost:    "https://accounts.azuredatabricks.net",
			deploymentName: "myworkspace",
			want:           "https://myworkspace.azuredatabricks.net",
		},
		{
			name:           "SPOG host without scheme returns account host as-is",
			env:            awsEnv,
			accountHost:    "myspog.custom.example.com",
			deploymentName: "myworkspace",
			want:           "myspog.custom.example.com",
		},
		{
			name:           "empty dns zone returns account host as-is",
			env:            environment.DatabricksEnvironment{},
			accountHost:    "https://accounts.cloud.databricks.com",
			deploymentName: "myworkspace",
			want:           "https://accounts.cloud.databricks.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := workspaceHost(tc.env, tc.accountHost, tc.deploymentName)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
