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
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegration) {
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegration) {
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

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegrationOutput) {
}

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegrationOutput) {
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId types.String `tfsdk:"app_id" tf:"optional"`
	// Token access policy
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegration) {
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegration) {
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegrationOutput) {
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegrationOutput) {
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *CreateServicePrincipalSecretRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServicePrincipalSecretRequest) {
}

func (newState *CreateServicePrincipalSecretRequest) SyncEffectiveFieldsDuringRead(existingState CreateServicePrincipalSecretRequest) {
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

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServicePrincipalSecretResponse) {
}

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringRead(existingState CreateServicePrincipalSecretResponse) {
}

type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails types.String `tfsdk:"authorization_details" tf:"optional"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url" tf:"optional"`
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneInfo) {
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringRead(existingState DataPlaneInfo) {
}

type DeleteCustomAppIntegrationOutput struct {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCustomAppIntegrationOutput) {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState DeleteCustomAppIntegrationOutput) {
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *DeleteCustomAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCustomAppIntegrationRequest) {
}

func (newState *DeleteCustomAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCustomAppIntegrationRequest) {
}

type DeletePublishedAppIntegrationOutput struct {
}

func (newState *DeletePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePublishedAppIntegrationOutput) {
}

func (newState *DeletePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState DeletePublishedAppIntegrationOutput) {
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *DeletePublishedAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePublishedAppIntegrationRequest) {
}

func (newState *DeletePublishedAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState DeletePublishedAppIntegrationRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteServicePrincipalSecretRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServicePrincipalSecretRequest) {
}

func (newState *DeleteServicePrincipalSecretRequest) SyncEffectiveFieldsDuringRead(existingState DeleteServicePrincipalSecretRequest) {
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
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationOutput) {
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationOutput) {
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *GetCustomAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationRequest) {
}

func (newState *GetCustomAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationRequest) {
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationsOutput) {
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationsOutput) {
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
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationOutput) {
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationOutput) {
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *GetPublishedAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationRequest) {
}

func (newState *GetPublishedAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationRequest) {
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationsOutput) {
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationsOutput) {
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps []PublishedAppOutput `tfsdk:"apps" tf:"optional"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppsOutput) {
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppsOutput) {
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCustomAppIntegrationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCustomAppIntegrationsRequest) {
}

func (newState *ListCustomAppIntegrationsRequest) SyncEffectiveFieldsDuringRead(existingState ListCustomAppIntegrationsRequest) {
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListOAuthPublishedAppsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListOAuthPublishedAppsRequest) {
}

func (newState *ListOAuthPublishedAppsRequest) SyncEffectiveFieldsDuringRead(existingState ListOAuthPublishedAppsRequest) {
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListPublishedAppIntegrationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPublishedAppIntegrationsRequest) {
}

func (newState *ListPublishedAppIntegrationsRequest) SyncEffectiveFieldsDuringRead(existingState ListPublishedAppIntegrationsRequest) {
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the secrets for this service principal.
	// Provide this token to retrieve the next page of secret entries. When
	// providing a `page_token`, all other parameters provided to the request
	// must match the previous request. To list all of the secrets for a service
	// principal, it is necessary to continue requesting pages of entries until
	// the response contains no `next_page_token`. Note that the number of
	// entries returned must not be used to determine when the listing is
	// complete.
	PageToken types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *ListServicePrincipalSecretsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalSecretsRequest) {
}

func (newState *ListServicePrincipalSecretsRequest) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalSecretsRequest) {
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of the secrets
	Secrets []SecretInfo `tfsdk:"secrets" tf:"optional"`
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalSecretsResponse) {
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalSecretsResponse) {
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

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedAppOutput) {
}

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringRead(existingState PublishedAppOutput) {
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

func (newState *SecretInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretInfo) {
}

func (newState *SecretInfo) SyncEffectiveFieldsDuringRead(existingState SecretInfo) {
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes types.Int64 `tfsdk:"access_token_ttl_in_minutes" tf:"optional"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes types.Int64 `tfsdk:"refresh_token_ttl_in_minutes" tf:"optional"`
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessPolicy) {
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringRead(existingState TokenAccessPolicy) {
}

type UpdateCustomAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []types.String `tfsdk:"redirect_urls" tf:"optional"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegration) {
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegration) {
}

type UpdateCustomAppIntegrationOutput struct {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegrationOutput) {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegrationOutput) {
}

type UpdatePublishedAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy []TokenAccessPolicy `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegration) {
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegration) {
}

type UpdatePublishedAppIntegrationOutput struct {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegrationOutput) {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegrationOutput) {
}
