package acceptance

import (
	"context"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertEnableIpAccessList(t *testing.T, expected string) {
	w, err := databricks.NewWorkspaceClient(&databricks.Config{})
	require.NoError(t, err)
	conf, err := w.WorkspaceConf.GetStatus(context.Background(), settings.GetStatusRequest{
		Keys: "enableIpAccessLists",
	})
	require.NoError(t, err)
	assert.Len(t, *conf, 1)
	assert.Equal(t, expected, (*conf)["enableIpAccessLists"])
}

func TestAccWorkspaceConfFullLifecycle(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_workspace_conf" "this" {
			custom_config = {
				"enableIpAccessLists": true
			}
		}`,
		Check: func(s *terraform.State) error {
			assertEnableIpAccessList(t, "true")
			return nil
		},
	}, Step{
		// Set enableIpAccessLists to false
		Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLists": "false"
				}
			}`,
		Check: func(s *terraform.State) error {
			// Assert server side configuration is updated
			assertEnableIpAccessList(t, "false")

			// Assert state is persisted
			conf := s.RootModule().Resources["databricks_workspace_conf.this"]
			assert.Equal(t, "false", conf.Primary.Attributes["custom_config.enableIpAccessLists"])
			return nil
		},
	}, Step{
		// Set invalid configuration
		Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLissss": "invalid"
				}
			}`,
		// Assert on server side error returned
		ExpectError: regexp.MustCompile(`cannot update workspace conf: Invalid keys`),
	}, Step{
		// Set enableIpAccessLists to true with strange case and maxTokenLifetimeDays to verify
		// failed deletion case
		Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLists": "TRue",
					"maxTokenLifetimeDays": 90
				}
			}`,
		Check: func(s *terraform.State) error {
			// Assert server side configuration is updated
			assertEnableIpAccessList(t, "TRue")

			// Assert state is persisted
			conf := s.RootModule().Resources["databricks_workspace_conf.this"]
			assert.Equal(t, "TRue", conf.Primary.Attributes["custom_config.enableIpAccessLists"])
			return nil
		},
	})
}
