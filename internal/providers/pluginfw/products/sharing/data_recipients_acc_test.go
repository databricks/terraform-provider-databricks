package sharing_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkRecipientsDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_recipients.this"]
		require.True(t, ok, "data.databricks_recipients.this has to be there")
		num_recipients, _ := strconv.Atoi(ds.Primary.Attributes["recipients.#"])
		assert.GreaterOrEqual(t, num_recipients, 1)
		return nil
	}
}

func TestUcAccDataSourceRecipients(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_recipient" "this" {
			name                = "{var.RANDOM}-terraform-recipient"
			comment             = "made by terraform"
			authentication_type = "TOKEN"
			sharing_code        = "{var.RANDOM}"
			ip_access_list {
				// using private ip for acc testing
				allowed_ip_addresses = ["10.0.0.0/16"]
			}
		}

		data "databricks_recipients" "this" {
			depends_on = [databricks_recipient.this]
		}
		`,
		Check: checkRecipientsDataSourcePopulated(t),
	})
}
