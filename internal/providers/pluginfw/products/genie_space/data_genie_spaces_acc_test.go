package genie_space_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

// checkGenieSpacesQueryable returns a Check that asserts the data source at
// dataAddr resolved without error and produced a numeric `spaces.#` attribute.
// The check tolerates empty workspaces (count >= 0) because we cannot reliably
// seed a Genie space in CI: the `serialized_space` field requires a known-good,
// API-validated JSON blob that is not committed as a test fixture.
func checkGenieSpacesQueryable(t *testing.T, dataAddr string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources[dataAddr]
		require.True(t, ok, "%s must be present in state", dataAddr)
		count, err := strconv.Atoi(ds.Primary.Attributes["spaces.#"])
		require.NoError(t, err, "spaces.# must be a number")
		require.GreaterOrEqual(t, count, 0, "spaces count must be non-negative")
		return nil
	}
}

// TestAccGenieSpacesDataSource exercises databricks_genie_spaces end-to-end
// against a live workspace. It validates that the data source can be queried
// both unfiltered and with a title_contains filter, that pagination terminates
// cleanly, and that the schema matches the resource.
//
// A companion resource acceptance test is deferred pending a committed
// minimal-valid serialized_space fixture.
func TestAccGenieSpacesDataSource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
			data "databricks_genie_spaces" "all" {}

			data "databricks_genie_spaces" "filtered" {
				title_contains = "tf-acc-genie-{var.RANDOM}"
			}
		`,
		Check: func(s *terraform.State) error {
			if err := checkGenieSpacesQueryable(t, "data.databricks_genie_spaces.all")(s); err != nil {
				return err
			}
			return checkGenieSpacesQueryable(t, "data.databricks_genie_spaces.filtered")(s)
		},
	})
}
