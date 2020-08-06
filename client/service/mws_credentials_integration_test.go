package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMwsAccCreds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := os.Getenv("DATABRICKS_ACCOUNT_ID")
	client := CommonEnvironmentClient()
	credsList, err := client.MWSCredentials().List(acctID)
	assert.NoError(t, err, err)
	t.Log(credsList)

	myCreds, err := client.MWSCredentials().Create(acctID, "sri-mws-terraform-automation-role", "arn:aws:iam::997819999999:role/sri-e2-terraform-automation-role")
	assert.NoError(t, err, err)

	myCredsFull, err := client.MWSCredentials().Read(acctID, myCreds.CredentialsID)
	assert.NoError(t, err, err)
	t.Log(myCredsFull.AwsCredentials.StsRole.ExternalID)

	defer func() {
		err = client.MWSCredentials().Delete(acctID, myCreds.CredentialsID)
		assert.NoError(t, err, err)
	}()
}
