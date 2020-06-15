package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMWSWorkspace(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	workspaceList, err := client.MWSWorkspaces().List(acctId)
	assert.NoError(t, err, err)
	t.Log(workspaceList)
}
