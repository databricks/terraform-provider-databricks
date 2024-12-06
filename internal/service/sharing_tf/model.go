// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sharing_tf

import (
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:""`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProvider) {
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringRead(existingState CreateProvider) {
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:""`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRecipient) {
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringRead(existingState CreateRecipient) {
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:""`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateShare) {
}

func (newState *CreateShare) SyncEffectiveFieldsDuringRead(existingState CreateShare) {
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderRequest) {
}

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringRead(existingState DeleteProviderRequest) {
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRecipientRequest) {
}

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRecipientRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteShareRequest) {
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringRead(existingState DeleteShareRequest) {
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoRequest) {
}

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoRequest) {
}

type GetActivationUrlInfoResponse struct {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoResponse) {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoResponse) {
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderRequest) {
}

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringRead(existingState GetProviderRequest) {
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientRequest) {
}

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringRead(existingState GetRecipientRequest) {
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `tfsdk:"permissions_out" tf:"optional"`
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientSharePermissionsResponse) {
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState GetRecipientSharePermissionsResponse) {
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetShareRequest) {
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringRead(existingState GetShareRequest) {
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []types.String `tfsdk:"allowed_ip_addresses" tf:"optional"`
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessList) {
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringRead(existingState IpAccessList) {
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider shares.
	Shares []ProviderShare `tfsdk:"shares" tf:"optional"`
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderSharesResponse) {
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderSharesResponse) {
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersRequest) {
}

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringRead(existingState ListProvidersRequest) {
}

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider information objects.
	Providers []ProviderInfo `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse) {
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsRequest) {
}

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringRead(existingState ListRecipientsRequest) {
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of recipient information objects.
	Recipients []RecipientInfo `tfsdk:"recipients" tf:"optional"`
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsResponse) {
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringRead(existingState ListRecipientsResponse) {
}

// List shares by Provider
type ListSharesRequest struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the provider in which to list shares.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesRequest) {
}

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringRead(existingState ListSharesRequest) {
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share information objects.
	Shares []ShareInfo `tfsdk:"shares" tf:"optional"`
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesResponse) {
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListSharesResponse) {
}

type Partition struct {
	// An array of partition values.
	Values []PartitionValue `tfsdk:"value" tf:"optional"`
}

func (newState *Partition) SyncEffectiveFieldsDuringCreateOrUpdate(plan Partition) {
}

func (newState *Partition) SyncEffectiveFieldsDuringRead(existingState Partition) {
}

type PartitionValue struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op" tf:"optional"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key" tf:"optional"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionValue) {
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringRead(existingState PartitionValue) {
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges []types.String `tfsdk:"privileges" tf:"optional"`
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment) {
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment) {
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id" tf:"optional"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile []RecipientProfile `tfsdk:"recipient_profile" tf:"optional,object"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderShare) {
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringRead(existingState ProviderShare) {
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated" tf:"optional"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional,object"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `tfsdk:"tokens" tf:"optional"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo) {
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringRead(existingState RecipientInfo) {
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version" tf:"optional"`
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientProfile) {
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringRead(existingState RecipientProfile) {
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo) {
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo) {
}

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenRequest) {
}

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenRequest) {
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenResponse) {
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenResponse) {
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds" tf:""`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan RotateRecipientToken) {
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringRead(existingState RotateRecipientToken) {
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:""`
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecurablePropertiesKvPairs) {
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState SecurablePropertiesKvPairs) {
}

type ShareInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"computed,optional"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"computed,optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `tfsdk:"object" tf:"optional"`
	// Username of current owner of share.
	Owner          types.String `tfsdk:"owner" tf:"optional"`
	EffectiveOwner types.String `tfsdk:"effective_owner" tf:"computed,optional"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"computed,optional"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"computed,optional"`
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
	newState.EffectiveOwner = newState.Owner
	newState.Owner = plan.Owner
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
	newState.EffectiveOwner = existingState.EffectiveOwner
	if existingState.EffectiveOwner.ValueString() == newState.Owner.ValueString() {
		newState.Owner = existingState.Owner
	}
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the Recipient.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharePermissionsRequest) {
}

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState SharePermissionsRequest) {
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments" tf:"optional"`
	// The share name.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareToPrivilegeAssignment) {
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState ShareToPrivilegeAssignment) {
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"computed,optional"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by" tf:"computed,optional"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled          types.Bool `tfsdk:"cdf_enabled" tf:"optional"`
	EffectiveCdfEnabled types.Bool `tfsdk:"effective_cdf_enabled" tf:"computed,optional"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The type of the data object.
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus          types.String `tfsdk:"history_data_sharing_status" tf:"optional"`
	EffectiveHistoryDataSharingStatus types.String `tfsdk:"effective_history_data_sharing_status" tf:"computed,optional"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name types.String `tfsdk:"name" tf:""`
	// Array of partitions for the shared data.
	Partitions []Partition `tfsdk:"partition" tf:"optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs          types.String `tfsdk:"shared_as" tf:"optional"`
	EffectiveSharedAs types.String `tfsdk:"effective_shared_as" tf:"computed,optional"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion          types.Int64 `tfsdk:"start_version" tf:"optional"`
	EffectiveStartVersion types.Int64 `tfsdk:"effective_start_version" tf:"computed,optional"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status" tf:"computed,optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as" tf:"optional"`
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject) {
	newState.EffectiveCdfEnabled = newState.CdfEnabled
	newState.CdfEnabled = plan.CdfEnabled
	newState.EffectiveHistoryDataSharingStatus = newState.HistoryDataSharingStatus
	newState.HistoryDataSharingStatus = plan.HistoryDataSharingStatus
	newState.EffectiveSharedAs = newState.SharedAs
	newState.SharedAs = plan.SharedAs
	newState.EffectiveStartVersion = newState.StartVersion
	newState.StartVersion = plan.StartVersion
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringRead(existingState SharedDataObject) {
	newState.EffectiveCdfEnabled = existingState.EffectiveCdfEnabled
	if existingState.EffectiveCdfEnabled.ValueBool() == newState.CdfEnabled.ValueBool() {
		newState.CdfEnabled = existingState.CdfEnabled
	}
	newState.EffectiveHistoryDataSharingStatus = existingState.EffectiveHistoryDataSharingStatus
	if existingState.EffectiveHistoryDataSharingStatus.ValueString() == newState.HistoryDataSharingStatus.ValueString() {
		newState.HistoryDataSharingStatus = existingState.HistoryDataSharingStatus
	}
	newState.EffectiveSharedAs = existingState.EffectiveSharedAs
	if existingState.EffectiveSharedAs.ValueString() == newState.SharedAs.ValueString() {
		newState.SharedAs = existingState.SharedAs
	}
	newState.EffectiveStartVersion = existingState.EffectiveStartVersion
	if existingState.EffectiveStartVersion.ValueInt64() == newState.StartVersion.ValueInt64() {
		newState.StartVersion = existingState.StartVersion
	}
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action" tf:"optional"`
	// The data object that is being added, removed, or updated.
	DataObject []SharedDataObject `tfsdk:"data_object" tf:"optional,object"`
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObjectUpdate) {
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringRead(existingState SharedDataObjectUpdate) {
}

type UpdatePermissionsResponse struct {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissionsResponse) {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState UpdatePermissionsResponse) {
}

type UpdateProvider struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProvider) {
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringRead(existingState UpdateProvider) {
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional,object"`
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRecipient) {
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringRead(existingState UpdateRecipient) {
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of share.
	Owner          types.String `tfsdk:"owner" tf:"optional"`
	EffectiveOwner types.String `tfsdk:"effective_owner" tf:"computed,optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `tfsdk:"updates" tf:"optional"`
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare) {
	newState.EffectiveOwner = newState.Owner
	newState.Owner = plan.Owner
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringRead(existingState UpdateShare) {
	newState.EffectiveOwner = existingState.EffectiveOwner
	if existingState.EffectiveOwner.ValueString() == newState.Owner.ValueString() {
		newState.Owner = existingState.Owner
	}
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes catalog.PermissionsChange `tfsdk:"changes" tf:"optional"`
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions) {
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions) {
}
