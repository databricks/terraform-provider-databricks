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
	acctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	workspaceList, err := client.MWSWorkspaces().List(acctID)
	assert.NoError(t, err, err)
	t.Log(workspaceList)
}
