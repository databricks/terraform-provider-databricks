package oauth_test

import (
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"testing"
)

func TestAccResourceServicePrincipalFederationPolicy(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `

		resource "databricks_service_principal" "sp" {
		  display_name = "Service Principal"
		}
		resource "databricks_service_principal_federation_policy" "dspfp" {
		  service_principal_id = databricks_service_principal.sp.id
		  oidc_policy  {
			issuer    = "https://accounts.google.com"
			audiences = ["databricks"]
			subject   = "subject"
		  }     
		}
		`,
	})
}
