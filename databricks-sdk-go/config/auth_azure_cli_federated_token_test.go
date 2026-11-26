package config

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestAzureCliCredentials_FederatedTokenServicePrincipal(t *testing.T) {
	// This test verifies the fix where service principals authenticated via
	// federated token (like in AKS with workload identity) use a fallback mechanism:
	// try with tenant ID first, then retry without tenant ID if it fails.

	env.CleanupEnvironment(t)
	t.Setenv("PATH", testdataPath())

	// Simulate a service principal with a client ID (not systemAssignedIdentity/userAssignedIdentity)
	// This represents the federated token scenario that occurs in AKS workload identity
	t.Setenv("AZ_USER_NAME", "5817e630-86b3-4f67-a38e-a63e6a1a401c")
	t.Setenv("AZ_USER_TYPE", "servicePrincipal")

	// This makes the mock az command fail when --tenant is passed, simulating the federated
	// token scenario where tenant ID causes authentication failure
	t.Setenv("FAIL_IF_TENANT_ID_SET", "true")

	aa := AzureCliCredentials{}
	cfg := &Config{
		Host:          "https://adb-1891644720860465.5.azuredatabricks.net/",
		AzureTenantID: "e6a2f6d5-ece9-4c0d-9464-9c493497cb8f",
	}

	// With the fallback fix, this should work: first attempt with --tenant fails,
	// then fallback without --tenant succeeds
	visitor, err := aa.Configure(context.Background(), cfg)

	assert.NoError(t, err, "Authentication should work with federated token service principals via fallback mechanism")
	assert.NotNil(t, visitor, "Should return a valid credentials provider")
}
