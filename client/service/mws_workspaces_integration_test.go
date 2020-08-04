package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMwsAccWorkspace(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := os.Getenv("DATABRICKS_ACCOUNT_ID")
	client := CommonEnvironmentClient()
	workspaceList, err := client.MWSWorkspaces().List(acctID)
	assert.NoError(t, err, err)
	t.Log(workspaceList)
}
