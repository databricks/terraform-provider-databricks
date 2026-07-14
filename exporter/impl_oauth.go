package exporter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
)

func listCustomAppIntegrations(ic *importContext) error {
	it := ic.accountClient.CustomAppIntegration.List(ic.Context, oauth2.ListCustomAppIntegrationsRequest{})
	for it.HasNext(ic.Context) {
		integration, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if !ic.MatchesName(integration.Name) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_custom_app_integration",
			ID:       integration.IntegrationId,
		})
	}
	return nil
}

func listAccountFederationPolicies(ic *importContext) error {
	it := ic.accountClient.FederationPolicy.List(ic.Context, oauth2.ListAccountFederationPoliciesRequest{})
	for it.HasNext(ic.Context) {
		policy, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if !ic.MatchesName(policy.Name) {
			continue
		}
		if policy.ServicePrincipalId != 0 {
			continue
		}
		// Extract policy name from the full path
		// Format: accounts/<account-id>/federationPolicies/<policy-name>
		name := ""
		parts := strings.Split(policy.Name, "/")
		if len(parts) > 0 {
			name = parts[len(parts)-1]
		} else {
			name = policy.PolicyId
		}

		ic.Emit(&resource{
			Resource: "databricks_account_federation_policy",
			ID:       policy.PolicyId,
			Name:     fmt.Sprintf("acc_fed_policy_%s", name),
		})
	}
	return nil
}

func listServicePrincipalFederationPolicies(ic *importContext) error {
	// First, list all service principals
	sps, err := ic.accountClient.ServicePrincipals.ListAll(ic.Context, iam.ListAccountServicePrincipalsRequest{
		Attributes: "id",
	})
	if err != nil {
		return err
	}

	// For each service principal, list its federation policies
	for _, sp := range sps {
		ic.Emit(&resource{
			Resource: "databricks_service_principal",
			ID:       sp.Id,
		})
		err = emitServicePrincipalFederationPolicies(ic, sp.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func emitServicePrincipalFederationPolicies(ic *importContext, spIdStr string) error {
	spId, err := strconv.ParseInt(spIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse service principal ID %s: %w", spIdStr, err)
	}
	it := ic.accountClient.ServicePrincipalFederationPolicy.List(ic.Context, oauth2.ListServicePrincipalFederationPoliciesRequest{
		ServicePrincipalId: spId,
	})
	for it.HasNext(ic.Context) {
		policy, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		// Extract policy name from the full path
		// Format: accounts/<account-id>/servicePrincipals/<sp-id>/federationPolicies/<policy-name>
		name := ""
		parts := strings.Split(policy.Name, "/")
		if len(parts) > 0 {
			name = parts[len(parts)-1]
		} else {
			name = policy.PolicyId
		}
		ic.Emit(&resource{
			Resource: "databricks_service_principal_federation_policy",
			ID:       fmt.Sprintf("%d,%s", spId, policy.PolicyId),
			Name:     fmt.Sprintf("sp_fed_policy_%d_%s", spId, name),
		})
	}
	return nil
}
