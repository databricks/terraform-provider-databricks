// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
)

// These APIs manage account federation policies.
//
// Account federation policies allow users and service principals in your
// Databricks account to securely access Databricks APIs using tokens from your
// trusted identity providers (IdPs).
//
// With token federation, your users and service principals can exchange tokens
// from your IdP for Databricks OAuth tokens, which can be used to access
// Databricks APIs. Token federation eliminates the need to manage Databricks
// secrets, and allows you to centralize management of token issuance policies
// in your IdP. Databricks token federation is typically used in combination
// with [SCIM], so users in your IdP are synchronized into your Databricks
// account.
//
// Token federation is configured in your Databricks account using an account
// federation policy. An account federation policy specifies: * which IdP, or
// issuer, your Databricks account should accept tokens from * how to determine
// which Databricks user, or subject, a token is issued for
//
// To configure a federation policy, you provide the following: * The required
// token __issuer__, as specified in the “iss” claim of your tokens. The
// issuer is an https URL that identifies your IdP. * The allowed token
// __audiences__, as specified in the “aud” claim of your tokens. This
// identifier is intended to represent the recipient of the token. As long as
// the audience in the token matches at least one audience in the policy, the
// token is considered a match. If unspecified, the default value is your
// Databricks account id. * The __subject claim__, which indicates which token
// claim contains the Databricks username of the user the token was issued for.
// If unspecified, the default value is “sub”. * Optionally, the public keys
// used to validate the signature of your tokens, in JWKS format. If unspecified
// (recommended), Databricks automatically fetches the public keys from your
// issuer’s well known endpoint. Databricks strongly recommends relying on
// your issuer’s well known endpoint for discovering public keys.
//
// An example federation policy is: ``` issuer: "https://idp.mycompany.com/oidc"
// audiences: ["databricks"] subject_claim: "sub" ```
//
// An example JWT token body that matches this policy and could be used to
// authenticate to Databricks as user `username@mycompany.com` is: ``` { "iss":
// "https://idp.mycompany.com/oidc", "aud": "databricks", "sub":
// "username@mycompany.com" } ```
//
// You may also need to configure your IdP to generate tokens for your users to
// exchange with Databricks, if your users do not already have the ability to
// generate tokens that are compatible with your federation policy.
//
// You do not need to configure an OAuth application in Databricks to use token
// federation.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [SCIM]: https://docs.databricks.com/admin/users-groups/scim/index.html
type AccountFederationPolicyService interface {

	// Create account federation policy.
	Create(ctx context.Context, request CreateAccountFederationPolicyRequest) (*FederationPolicy, error)

	// Delete account federation policy.
	Delete(ctx context.Context, request DeleteAccountFederationPolicyRequest) error

	// Get account federation policy.
	Get(ctx context.Context, request GetAccountFederationPolicyRequest) (*FederationPolicy, error)

	// List account federation policies.
	List(ctx context.Context, request ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error)

	// Update account federation policy.
	Update(ctx context.Context, request UpdateAccountFederationPolicyRequest) (*FederationPolicy, error)
}

// These APIs enable administrators to manage custom OAuth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CustomAppIntegrationService interface {

	// Create Custom OAuth App Integration.
	//
	// You can retrieve the custom OAuth app integration via
	// [CustomAppIntegration/get].
	//
	// [CustomAppIntegration/get]: https://docs.databricks.com/api/account/customappintegration/get
	Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error)

	// Delete an existing Custom OAuth App Integration. You can retrieve the
	// custom OAuth app integration via [CustomAppIntegration/get].
	//
	// [CustomAppIntegration/get]: https://docs.databricks.com/api/account/customappintegration/get
	Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error

	// Gets the Custom OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error)

	// Get the list of custom OAuth app integrations for the specified
	// Databricks account.
	List(ctx context.Context, request ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error)

	// Updates an existing custom OAuth App Integration. You can retrieve the
	// custom OAuth app integration via [CustomAppIntegration/get].
	//
	// [CustomAppIntegration/get]: https://docs.databricks.com/api/account/customappintegration/get
	Update(ctx context.Context, request UpdateCustomAppIntegration) error
}

// These APIs enable administrators to view all the available published OAuth
// applications in Databricks. Administrators can add the published OAuth
// applications to their account through the OAuth Published App Integration
// APIs.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type OAuthPublishedAppsService interface {

	// Get all the available published OAuth apps in Databricks.
	List(ctx context.Context, request ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error)
}

// These APIs enable administrators to manage published OAuth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PublishedAppIntegrationService interface {

	// Create Published OAuth App Integration.
	//
	// You can retrieve the published OAuth app integration via
	// [PublishedAppIntegration/get].
	//
	// [PublishedAppIntegration/get]: https://docs.databricks.com/api/account/publishedappintegration/get
	Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error)

	// Delete an existing Published OAuth App Integration. The app can no longer
	// be used to sign in to the account until it is enabled again with
	// [PublishedAppIntegration/create].
	//
	// [PublishedAppIntegration/create]: https://docs.databricks.com/api/account/publishedappintegration/create
	Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error

	// Gets the Published OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error)

	// Get the list of published OAuth app integrations for the specified
	// Databricks account.
	List(ctx context.Context, request ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error)

	// Updates an existing published OAuth App Integration. You can retrieve the
	// published OAuth app integration via [PublishedAppIntegration/get].
	//
	// [PublishedAppIntegration/get]: https://docs.databricks.com/api/account/publishedappintegration/get
	Update(ctx context.Context, request UpdatePublishedAppIntegration) error
}

// These APIs manage service principal federation policies.
//
// Service principal federation, also known as Workload Identity Federation,
// allows your automated workloads running outside of Databricks to securely
// access Databricks APIs without the need for Databricks secrets. With Workload
// Identity Federation, your application (or workload) authenticates to
// Databricks as a Databricks service principal, using tokens provided by the
// workload runtime.
//
// Databricks strongly recommends using Workload Identity Federation to
// authenticate to Databricks from automated workloads, over alternatives such
// as OAuth client secrets or Personal Access Tokens, whenever possible.
// Workload Identity Federation is supported by many popular services, including
// Github Actions, Azure DevOps, GitLab, Terraform Cloud, and Kubernetes
// clusters, among others.
//
// Workload identity federation is configured in your Databricks account using a
// service principal federation policy. A service principal federation policy
// specifies: * which IdP, or issuer, the service principal is allowed to
// authenticate from * which workload identity, or subject, is allowed to
// authenticate as the Databricks service principal
//
// To configure a federation policy, you provide the following: * The required
// token __issuer__, as specified in the “iss” claim of workload identity
// tokens. The issuer is an https URL that identifies the workload identity
// provider. * The required token __subject__, as specified in the “sub”
// claim of workload identity tokens. The subject uniquely identifies the
// workload in the workload runtime environment. * The allowed token
// __audiences__, as specified in the “aud” claim of workload identity
// tokens. The audience is intended to represent the recipient of the token. As
// long as the audience in the token matches at least one audience in the
// policy, the token is considered a match. If unspecified, the default value is
// your Databricks account id. * Optionally, the public keys used to validate
// the signature of the workload identity tokens, in JWKS format. If unspecified
// (recommended), Databricks automatically fetches the public keys from the
// issuer’s well known endpoint. Databricks strongly recommends relying on the
// issuer’s well known endpoint for discovering public keys.
//
// An example service principal federation policy, for a Github Actions
// workload, is: ``` issuer: "https://token.actions.githubusercontent.com"
// audiences: ["https://github.com/my-github-org"] subject:
// "repo:my-github-org/my-repo:environment:prod" ```
//
// An example JWT token body that matches this policy and could be used to
// authenticate to Databricks is: ``` { "iss":
// "https://token.actions.githubusercontent.com", "aud":
// "https://github.com/my-github-org", "sub":
// "repo:my-github-org/my-repo:environment:prod" } ```
//
// You may also need to configure the workload runtime to generate tokens for
// your workloads.
//
// You do not need to configure an OAuth application in Databricks to use token
// federation.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ServicePrincipalFederationPolicyService interface {

	// Create account federation policy.
	Create(ctx context.Context, request CreateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error)

	// Delete account federation policy.
	Delete(ctx context.Context, request DeleteServicePrincipalFederationPolicyRequest) error

	// Get account federation policy.
	Get(ctx context.Context, request GetServicePrincipalFederationPolicyRequest) (*FederationPolicy, error)

	// List account federation policies.
	List(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error)

	// Update account federation policy.
	Update(ctx context.Context, request UpdateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error)
}

// These APIs enable administrators to manage service principal secrets.
//
// You can use the generated secrets to obtain OAuth access tokens for a service
// principal, which can then be used to access Databricks Accounts and Workspace
// APIs. For more information, see [Authentication using OAuth tokens for
// service principals].
//
// In addition, the generated secrets can be used to configure the Databricks
// Terraform Provider to authenticate with the service principal. For more
// information, see [Databricks Terraform Provider].
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Authentication using OAuth tokens for service principals]: https://docs.databricks.com/dev-tools/authentication-oauth.html
// [Databricks Terraform Provider]: https://github.com/databricks/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-service-principal
type ServicePrincipalSecretsService interface {

	// Create a secret for the given service principal.
	Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error)

	// Delete a secret from the given service principal.
	Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error

	// List all secrets associated with the given service principal. This
	// operation only returns information about the secrets themselves and does
	// not include the secret values.
	List(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error)
}

// These APIs enable administrators to manage service principal secrets at the
// workspace level. To use these APIs, the service principal must be first added
// to the current workspace.
//
// You can use the generated secrets to obtain OAuth access tokens for a service
// principal, which can then be used to access Databricks Accounts and Workspace
// APIs. For more information, see [Authentication using OAuth tokens for
// service principals].
//
// In addition, the generated secrets can be used to configure the Databricks
// Terraform Providerto authenticate with the service principal. For more
// information, see [Databricks Terraform Provider].
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Authentication using OAuth tokens for service principals]: https://docs.databricks.com/dev-tools/authentication-oauth.html
// [Databricks Terraform Provider]: https://github.com/databricks/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-service-principal
type ServicePrincipalSecretsProxyService interface {

	// Create a secret for the given service principal.
	Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error)

	// Delete a secret from the given service principal.
	Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error

	// List all secrets associated with the given service principal. This
	// operation only returns information about the secrets themselves and does
	// not include the secret values.
	List(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error)
}
