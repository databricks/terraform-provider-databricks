package mws

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccCustomerManagedKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID, acctIDset := os.LookupEnv("DATABRICKS_ACCOUNT_ID")
	if !acctIDset {
		t.Skip("MWS tests skipped unless env 'DATABRICKS_ACCOUNT_ID' is set")
	}
	client := common.CommonEnvironmentClient()
	networksList, err := NewMWSCustomerManagedKeysAPI(client).List(acctID)
	assert.NoError(t, err, err)
	t.Log(networksList)
}
