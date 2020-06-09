package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestE2CustomerManagedKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	networksList, err := client.MWSCustomerManagedKeys().List(acctId)
	assert.NoError(t, err, err)
	t.Log(networksList)

}
