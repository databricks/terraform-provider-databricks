package service

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	wsConfMapBuilder := model.WorkspaceConfRequestMapBuilder{}
	wsConfMap := wsConfMapBuilder.Build(
		model.WithLoginLogo("https://databricks.com/wp-content/uploads/2019/09/delta-lake-logo.png"),
	)
	err := client.WorkspaceConfigurations().Update(wsConfMap)
	assert.NoError(t, err, err)
	resp, err := client.WorkspaceConfigurations().Read(model.AllWorkspaceConfKeys)
	t.Log(resp.LoginLogo)
	assert.NoError(t, err, err)
}
