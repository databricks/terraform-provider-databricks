package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceDataUsers_DisplayNameContains(t *testing.T) {
	accountLevel(t, step{
		Template: `
		data "databricks_data_users" "this" {
			display_name_contains = "testuser"
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_data_users.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			ids := r.Primary.Attributes["users.#"]
			if ids == "" {
				return fmt.Errorf("users is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}

func TestAccDataSourceDataUsers_UserNameContains(t *testing.T) {
	accountLevel(t, step{
		Template: `
		data "databricks_data_users" "this" {
			user_name_contains = "example.com"
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_data_users.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			usersCount := r.Primary.Attributes["users.#"]
			if usersCount == "" || usersCount == "0" {
				return fmt.Errorf("users list is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}

func TestAccDataSourceDataUsers_NoFilters(t *testing.T) {
	accountLevel(t, step{
		Template: `
		data "databricks_data_users" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_data_users.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			usersCount := r.Primary.Attributes["users.#"]
			if usersCount == "" || usersCount == "0" {
				return fmt.Errorf("users list is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
