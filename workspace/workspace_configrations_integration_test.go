package workspace

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := common.NewClientFromEnvironment()

	wsConfMap := map[string]string{
		"enableIpAccessLists": "true",
	}
	err := NewWorkspaceConfAPI(client).Update(wsConfMap)
	assert.NoError(t, err, err)
	resp, err := NewWorkspaceConfAPI(client).Read("enableIpAccessLists")
	t.Log(resp["enableIpAccessLists"])
	assert.NoError(t, err, err)
}
