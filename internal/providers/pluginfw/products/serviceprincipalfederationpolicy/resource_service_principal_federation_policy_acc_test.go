package serviceprincipalfederationpolicy_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

const baseResources = `
  resource "databricks_service_principal" "this" {
    display_name = "test"
  }

  resource "databricks_service_principal_federation_policy" "this" {
   service_principal_id = databricks_service_principal.this.id
   oidc_policy = {
	 issuer    = "https://accounts.google.com"
	 audiences = ["%s"]
	 subject   = "subjcect"
   }
  }
`

const updateServicePrincipalResources = `
  resource "databricks_service_principal" "this" {
    display_name = "test"
  }

resource "databricks_service_principal" "this2" {
    display_name = "test"
  }

  resource "databricks_service_principal_federation_policy" "this" {
   service_principal_id = databricks_service_principal.this2.id
   oidc_policy = {
	 issuer    = "https://accounts.google.com"
	 audiences = ["%s"]
	 subject   = "subjcect"
   }
  }
`

func TestMwsAccDatabricksServicePrincipalFederationPolicyResourceNewServicePrincipal(t *testing.T) {
	var updateTime string
	var createTime string
	var servicePrincipalId string
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(baseResources, "databricks"),
		Check: func(s *terraform.State) error {
			updateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["update_time"]
			createTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["create_time"]
			servicePrincipalId = s.RootModule().Resources["databricks_service_principal.this"].Primary.Attributes["id"]
			assert.NotNil(t, updateTime)
			return nil
		},
	},
		acceptance.Step{
			Template: updateServicePrincipalResources,
			Check: func(s *terraform.State) error {
				var newUpdateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["update_time"]
				var newServicePrincipalId = s.RootModule().Resources["databricks_service_principal.this2"].Primary.Attributes["id"]
				var newCreateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["create_time"]
				var spfpspid = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["service_principal_id"]
				assert.NotEqual(t, updateTime, newUpdateTime)
				assert.NotEqual(t, newCreateTime, createTime)
				assert.NotEqual(t, newServicePrincipalId, servicePrincipalId)
				assert.Equal(t, newServicePrincipalId, spfpspid)
				return nil
			},
		})
}

func TestMwsAccDatabricksServicePrincipalFederationPolicyResourceNewAudiences(t *testing.T) {
	var updateTime string
	var createTime string
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(baseResources, "databricks"),
		Check: func(s *terraform.State) error {
			updateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["update_time"]
			createTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["create_time"]
			assert.NotNil(t, updateTime)
			return nil
		},
	},
		acceptance.Step{
			Template: fmt.Sprintf(baseResources, "notdatabricks"),
			Check: func(s *terraform.State) error {
				var newUpdateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["update_time"]
				var newCreateTime = s.RootModule().Resources["databricks_service_principal_federation_policy.this"].Primary.Attributes["create_time"]
				assert.NotEqual(t, updateTime, newUpdateTime)
				assert.Equal(t, newCreateTime, createTime)
				return nil
			},
		})
}
