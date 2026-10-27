package workspaceconf_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/qa/replay/clienttest"
	"github.com/databricks/terraform-provider-databricks/workspace"
)

func TestWorkspaceConfHelpers(t *testing.T) {
	for _, test := range []struct {
		name          string
		keys          []string
		expectedError string
	}{
		{"TestAccWorkspaceConf_GetValidKey", []string{"enableIpAccessLists"}, ""},
		{"TestAccWorkspaceConf_GetInvalidKey", []string{"invalidKey", "enableIpAccessLists"}, ""},
		{"TestAccWorkspaceConf_GetOnlyInvalidKeys", []string{"invalidKey"}, "failed to get workspace conf because all keys are invalid: invalidKey"},
	} {
		t.Run(test.name, func(t *testing.T) {
			client := clienttest.WorkspaceClient(t, "testdata/client/"+test.name+".yaml")
			conf, err := workspace.SafeGetStatus(context.Background(), client, test.keys)
			if test.expectedError != "" {
				assert.ErrorContains(t, err, test.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Contains(t, conf, "enableIpAccessLists")
		})
	}
	for _, test := range []struct {
		name          string
		removed       map[string]struct{}
		value         string
		expectedError string
	}{
		{"TestAccWorkspaceConf_SetInvalidKey", map[string]struct{}{}, "invalidValue", "failed to set workspace conf because some new keys are invalid: invalidKey"},
		{"TestAccWorkspaceConf_DeleteInvalidKey", map[string]struct{}{"invalidKey": {}}, "", ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			client := clienttest.WorkspaceClient(t, "testdata/client/"+test.name+".yaml")
			err := workspace.SafeSetStatus(context.Background(), client, test.removed, map[string]string{"invalidKey": test.value})
			if test.expectedError != "" {
				assert.ErrorContains(t, err, test.expectedError)
				return
			}
			assert.NoError(t, err)
		})
	}
}
