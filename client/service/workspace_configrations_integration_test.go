package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	wsConfMap := map[string]string{
		"enableIpAccessLists": "true",
	}
	err := client.WorkspaceConfigurations().Update(wsConfMap)
	assert.NoError(t, err, err)
	resp, err := client.WorkspaceConfigurations().Read("enableIpAccessLists")
	t.Log(resp["enableIpAccessLists"])
	assert.NoError(t, err, err)
}
