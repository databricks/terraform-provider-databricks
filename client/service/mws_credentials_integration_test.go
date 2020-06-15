package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMWSCreds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	credsList, err := client.MWSCredentials().List(acctId)
	assert.NoError(t, err, err)
	t.Log(credsList)

	myCreds, err := client.MWSCredentials().Create(acctId, "sri-mws-terraform-automation-role", "arn:aws:iam::997819999999:role/sri-e2-terraform-automation-role")
	assert.NoError(t, err, err)

	myCredsFull, err := client.MWSCredentials().Read(acctId, myCreds.CredentialsID)
	assert.NoError(t, err, err)
	t.Log(myCredsFull.AwsCredentials.StsRole.ExternalID)

	defer func() {
		err = client.MWSCredentials().Delete(acctId, myCreds.CredentialsID)
		assert.NoError(t, err, err)
	}()
}
