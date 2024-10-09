// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package oauth2_tf

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential" tf:"optional"`
	// Name of the custom OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of OAuth redirect urls
	RedirectUrls []types.String `tfsdk:"redirect_urls" tf:"optional"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes []types.String `tfsdk:"scopes" tf:"optional"`
	// Token access policy
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret types.String `tfsdk:"client_secret" tf:"optional"`
	// Unique integration id for the custom OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId types.String `tfsdk:"app_id" tf:"optional"`
	// Token access policy
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// ID of the secret
	Id types.String `tfsdk:"id" tf:"optional"`
	// Secret Value
	Secret types.String `tfsdk:"secret" tf:"optional"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash" tf:"optional"`
	// Status of the secret
	Status types.String `tfsdk:"status" tf:"optional"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails types.String `tfsdk:"authorization_details" tf:"optional"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url" tf:"optional"`
}

type DeleteCustomAppIntegrationOutput struct {
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

type DeletePublishedAppIntegrationOutput struct {
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

type DeleteResponse struct {
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential" tf:"optional"`

	CreateTime types.String `tfsdk:"create_time" tf:"optional"`

	CreatedBy types.Int64 `tfsdk:"created_by" tf:"optional"`

	CreatorUsername types.String `tfsdk:"creator_username" tf:"optional"`
	// ID of this custom app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
	// The display name of the custom OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of OAuth redirect urls
	RedirectUrls []types.String `tfsdk:"redirect_urls" tf:"optional"`

	Scopes []types.String `tfsdk:"scopes" tf:"optional"`
	// Token access policy
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	AppId types.String `tfsdk:"app_id" tf:"optional"`

	CreateTime types.String `tfsdk:"create_time" tf:"optional"`

	CreatedBy types.Int64 `tfsdk:"created_by" tf:"optional"`
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
	// Display name of the published OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// Token access policy
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps []PublishedAppOutput `tfsdk:"apps" tf:"optional"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

type ListServicePrincipalSecretsResponse struct {
	// List of the secrets
	Secrets []SecretInfo `tfsdk:"secrets" tf:"optional"`
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId types.String `tfsdk:"app_id" tf:"optional"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// Description of the published OAuth app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient types.Bool `tfsdk:"is_confidential_client" tf:"optional"`
	// The display name of the published OAuth app.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls []types.String `tfsdk:"redirect_urls" tf:"optional"`
	// Required scopes for the published OAuth app.
	Scopes []types.String `tfsdk:"scopes" tf:"optional"`
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// ID of the secret
	Id types.String `tfsdk:"id" tf:"optional"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash" tf:"optional"`
	// Status of the secret
	Status types.String `tfsdk:"status" tf:"optional"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes types.Int64 `tfsdk:"access_token_ttl_in_minutes" tf:"optional"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes types.Int64 `tfsdk:"refresh_token_ttl_in_minutes" tf:"optional"`
}

type UpdateCustomAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []types.String `tfsdk:"redirect_urls" tf:"optional"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

type UpdateCustomAppIntegrationOutput struct {
}

type UpdatePublishedAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional"`
}

type UpdatePublishedAppIntegrationOutput struct {
}
