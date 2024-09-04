package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const userDataSourceTemplate = `
resource "databricks_user" "this" {
	user_name = "tf-{var.RANDOM}@example.com"
}
data "databricks_user" "this" {
	user_name = databricks_user.this.user_name
}`

func checkUserDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_user.this"]
		require.True(t, ok, "data.databricks_user.this has to be there")
		assert.Equal(t, s.Modules[0].Resources["databricks_user.this"].Primary.ID, r.Primary.ID)
		return nil
	}
}

func TestMwsAccUserData(t *testing.T) {
	AccountLevel(t, Step{
		Template: userDataSourceTemplate,
		Check:    checkUserDataSourcePopulated(t),
	})
}

func TestAccUserData(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: userDataSourceTemplate,
		Check:    checkUserDataSourcePopulated(t),
	})
}
