package acceptance

import (
	"context"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	assert.Equal(t, (*conf)["enableIpAccessLists"], expected)
}

func TestAccWorkspaceConfFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_workspace_conf" "this" {
			custom_config = {
				"enableIpAccessLists": true
			}
		}`,
		Check: func(s *terraform.State) error {
			assertEnableIpAccessList(t, "true")
			return nil
		},
	}, step{
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
	}, step{
		// Set invalid configuration
		Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLissss": "invalid"
				}
			}`,
		// Assert on server side error returned
		ExpectError: regexp.MustCompile(`cannot update workspace conf: Invalid keys`),
	}, step{
		// Set enableIpAccessLists to true with strange case
		Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLists": "TRue"
				}
			}`,
		Check: func(s *terraform.State) error {
			// Assert server side configuration is updated
			assertEnableIpAccessList(t, "true")

			// Assert state is persisted
			conf := s.RootModule().Resources["databricks_workspace_conf.this"]
			assert.Equal(t, "true", conf.Primary.Attributes["custom_config.enableIpAccessLists"])
			return nil
		},
	})
}
