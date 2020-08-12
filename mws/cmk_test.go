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
	acctID := os.Getenv("DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	networksList, err := NewMWSCustomerManagedKeysAPI(client).List(acctID)
	assert.NoError(t, err, err)
	t.Log(networksList)
}
