package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMWSNetworks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	networksList, err := client.MWSNetworks().List(acctId)
	assert.NoError(t, err, err)
	t.Log(networksList)

	myNetwork, err := client.MWSNetworks().Create(acctId, "sri-mws-terraform-automation-network",
		"vpc-0abcdef1234567890", []string{"subnet-0123456789abcdef0", "subnet-0fedcba9876543210"}, []string{"sg-0a1b2c3d4e5f6a7b8"})
	assert.NoError(t, err, err)
	defer func() {
		err = client.MWSNetworks().Delete(acctId, myNetwork.NetworkID)
		assert.NoError(t, err, err)
	}()

	myNetworkFull, err := client.MWSNetworks().Read(acctId, myNetwork.NetworkID)
	assert.NoError(t, err, err)
	t.Log(myNetworkFull)
}
