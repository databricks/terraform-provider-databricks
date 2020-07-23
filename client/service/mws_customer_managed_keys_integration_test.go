package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMWSCustomerManagedKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	networksList, err := client.MWSCustomerManagedKeys().List(acctID)
	assert.NoError(t, err, err)
	t.Log(networksList)
}
