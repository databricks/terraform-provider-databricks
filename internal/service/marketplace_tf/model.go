// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package marketplace_tf

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddExchangeForListingRequest struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:""`

	ListingId types.String `tfsdk:"listing_id" tf:""`
}

type AddExchangeForListingResponse struct {
	ExchangeForListing []ExchangeListing `tfsdk:"exchange_for_listing" tf:"optional,object"`
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest struct {
	Ids []types.String `tfsdk:"-"`
}

type BatchGetListingsResponse struct {
	Listings []Listing `tfsdk:"listings" tf:"optional"`
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest struct {
	Ids []types.String `tfsdk:"-"`
}

type BatchGetProvidersResponse struct {
	Providers []ProviderInfo `tfsdk:"providers" tf:"optional"`
}

type ConsumerTerms struct {
	Version types.String `tfsdk:"version" tf:""`
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company types.String `tfsdk:"company" tf:"optional"`

	Email types.String `tfsdk:"email" tf:"optional"`

	FirstName types.String `tfsdk:"first_name" tf:"optional"`

	LastName types.String `tfsdk:"last_name" tf:"optional"`
}

type CreateExchangeFilterRequest struct {
	Filter []ExchangeFilter `tfsdk:"filter" tf:"object"`
}

type CreateExchangeFilterResponse struct {
	FilterId types.String `tfsdk:"filter_id" tf:"optional"`
}

type CreateExchangeRequest struct {
	Exchange []Exchange `tfsdk:"exchange" tf:"object"`
}

type CreateExchangeResponse struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:"optional"`
}

type CreateFileRequest struct {
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	FileParent []FileParent `tfsdk:"file_parent" tf:"object"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type" tf:""`

	MimeType types.String `tfsdk:"mime_type" tf:""`
}

type CreateFileResponse struct {
	FileInfo []FileInfo `tfsdk:"file_info" tf:"optional,object"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url" tf:"optional"`
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms []ConsumerTerms `tfsdk:"accepted_consumer_terms" tf:"optional,object"`

	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`
	// for git repo installations
	RepoDetail []RepoInstallation `tfsdk:"repo_detail" tf:"optional,object"`

	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

type CreateListingRequest struct {
	Listing []Listing `tfsdk:"listing" tf:"object"`
}

type CreateListingResponse struct {
	ListingId types.String `tfsdk:"listing_id" tf:"optional"`
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms []ConsumerTerms `tfsdk:"accepted_consumer_terms" tf:"object"`

	Comment types.String `tfsdk:"comment" tf:"optional"`

	Company types.String `tfsdk:"company" tf:"optional"`

	FirstName types.String `tfsdk:"first_name" tf:"optional"`

	IntendedUse types.String `tfsdk:"intended_use" tf:""`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse" tf:"optional"`

	LastName types.String `tfsdk:"last_name" tf:"optional"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`
}

type CreatePersonalizationRequestResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

type CreateProviderRequest struct {
	Provider []ProviderInfo `tfsdk:"provider" tf:"object"`
}

type CreateProviderResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

type DataRefreshInfo struct {
	Interval types.Int64 `tfsdk:"interval" tf:""`

	Unit types.String `tfsdk:"unit" tf:""`
}

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {
	Id types.String `tfsdk:"-"`
}

type DeleteExchangeFilterResponse struct {
}

// Delete an exchange
type DeleteExchangeRequest struct {
	Id types.String `tfsdk:"-"`
}

type DeleteExchangeResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

type DeleteFileResponse struct {
}

// Uninstall from a listing
type DeleteInstallationRequest struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
}

type DeleteInstallationResponse struct {
}

// Delete a listing
type DeleteListingRequest struct {
	Id types.String `tfsdk:"-"`
}

type DeleteListingResponse struct {
}

// Delete provider
type DeleteProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

type DeleteProviderResponse struct {
}

type Exchange struct {
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	Filters []ExchangeFilter `tfsdk:"filters" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	LinkedListings []ExchangeListing `tfsdk:"linked_listings" tf:"optional"`

	Name types.String `tfsdk:"name" tf:""`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`

	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type ExchangeFilter struct {
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	ExchangeId types.String `tfsdk:"exchange_id" tf:""`

	FilterType types.String `tfsdk:"filter_type" tf:""`

	FilterValue types.String `tfsdk:"filter_value" tf:""`

	Id types.String `tfsdk:"id" tf:"optional"`

	Name types.String `tfsdk:"name" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`

	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type ExchangeListing struct {
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	ExchangeId types.String `tfsdk:"exchange_id" tf:"optional"`

	ExchangeName types.String `tfsdk:"exchange_name" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:"optional"`

	ListingName types.String `tfsdk:"listing_name" tf:"optional"`
}

type FileInfo struct {
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	DownloadLink types.String `tfsdk:"download_link" tf:"optional"`

	FileParent []FileParent `tfsdk:"file_parent" tf:"optional,object"`

	Id types.String `tfsdk:"id" tf:"optional"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type" tf:"optional"`

	MimeType types.String `tfsdk:"mime_type" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

type FileParent struct {
	FileParentType types.String `tfsdk:"file_parent_type" tf:"optional"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id" tf:"optional"`
}

// Get an exchange
type GetExchangeRequest struct {
	Id types.String `tfsdk:"-"`
}

type GetExchangeResponse struct {
	Exchange []Exchange `tfsdk:"exchange" tf:"optional,object"`
}

// Get a file
type GetFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

type GetFileResponse struct {
	FileInfo []FileInfo `tfsdk:"file_info" tf:"optional,object"`
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

// Get listing content metadata
type GetListingContentMetadataRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type GetListingContentMetadataResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	SharedDataObjects []SharedDataObject `tfsdk:"shared_data_objects" tf:"optional"`
}

// Get listing
type GetListingRequest struct {
	Id types.String `tfsdk:"-"`
}

type GetListingResponse struct {
	Listing []Listing `tfsdk:"listing" tf:"optional,object"`
}

// List listings
type GetListingsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type GetListingsResponse struct {
	Listings []Listing `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests" tf:"optional"`
}

// Get a provider
type GetProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

type GetProviderResponse struct {
	Provider []ProviderInfo `tfsdk:"provider" tf:"optional,object"`
}

type Installation struct {
	Installation []InstallationDetail `tfsdk:"installation" tf:"optional,object"`
}

type InstallationDetail struct {
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	InstalledOn types.Int64 `tfsdk:"installed_on" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:"optional"`

	ListingName types.String `tfsdk:"listing_name" tf:"optional"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	RepoName types.String `tfsdk:"repo_name" tf:"optional"`

	RepoPath types.String `tfsdk:"repo_path" tf:"optional"`

	ShareName types.String `tfsdk:"share_name" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`

	TokenDetail []TokenDetail `tfsdk:"token_detail" tf:"optional,object"`

	Tokens []TokenInfo `tfsdk:"tokens" tf:"optional"`
}

// List all installations
type ListAllInstallationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListAllInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests" tf:"optional"`
}

// List exchange filters
type ListExchangeFiltersRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListExchangeFiltersResponse struct {
	Filters []ExchangeFilter `tfsdk:"filters" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List exchanges for listing
type ListExchangesForListingRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListExchangesForListingResponse struct {
	ExchangeListing []ExchangeListing `tfsdk:"exchange_listing" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List exchanges
type ListExchangesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListExchangesResponse struct {
	Exchanges []Exchange `tfsdk:"exchanges" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List files
type ListFilesRequest struct {
	FileParent []FileParent `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListFilesResponse struct {
	FileInfos []FileInfo `tfsdk:"file_infos" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List all listing fulfillments
type ListFulfillmentsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListFulfillmentsResponse struct {
	Fulfillments []ListingFulfillment `tfsdk:"fulfillments" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List installations for a listing
type ListInstallationsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List listings for exchange
type ListListingsForExchangeRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListListingsForExchangeResponse struct {
	ExchangeListings []ExchangeListing `tfsdk:"exchange_listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List listings
type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets []types.String `tfsdk:"-"`
	// Matches any of the following categories
	Categories []types.String `tfsdk:"-"`
	// Filters each listing based on if it is free.
	IsFree types.Bool `tfsdk:"-"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange types.Bool `tfsdk:"-"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Matches any of the following provider ids
	ProviderIds []types.String `tfsdk:"-"`
	// Matches any of the following tags
	Tags []ListingTag `tfsdk:"-"`
}

type ListListingsResponse struct {
	Listings []Listing `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`

	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

// List providers
type ListProvidersRequest struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

type ListProvidersResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Providers []ProviderInfo `tfsdk:"providers" tf:"optional"`
}

type Listing struct {
	Detail []ListingDetail `tfsdk:"detail" tf:"optional,object"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// Next Number: 26
	Summary []ListingSummary `tfsdk:"summary" tf:"object"`
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets []types.String `tfsdk:"assets" tf:"optional"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd types.Int64 `tfsdk:"collection_date_end" tf:"optional"`
	// The starting date timestamp for when the data spans
	CollectionDateStart types.Int64 `tfsdk:"collection_date_start" tf:"optional"`
	// Smallest unit of time in the dataset
	CollectionGranularity []DataRefreshInfo `tfsdk:"collection_granularity" tf:"optional,object"`
	// Whether the dataset is free or paid
	Cost types.String `tfsdk:"cost" tf:"optional"`
	// Where/how the data is sourced
	DataSource types.String `tfsdk:"data_source" tf:"optional"`

	Description types.String `tfsdk:"description" tf:"optional"`

	DocumentationLink types.String `tfsdk:"documentation_link" tf:"optional"`

	EmbeddedNotebookFileInfos []FileInfo `tfsdk:"embedded_notebook_file_infos" tf:"optional"`

	FileIds []types.String `tfsdk:"file_ids" tf:"optional"`
	// Which geo region the listing data is collected from
	GeographicalCoverage types.String `tfsdk:"geographical_coverage" tf:"optional"`
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	License types.String `tfsdk:"license" tf:"optional"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	PricingModel types.String `tfsdk:"pricing_model" tf:"optional"`

	PrivacyPolicyLink types.String `tfsdk:"privacy_policy_link" tf:"optional"`
	// size of the dataset in GB
	Size types.Float64 `tfsdk:"size" tf:"optional"`

	SupportLink types.String `tfsdk:"support_link" tf:"optional"`
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	Tags []ListingTag `tfsdk:"tags" tf:"optional"`

	TermsOfService types.String `tfsdk:"terms_of_service" tf:"optional"`
	// How often data is updated
	UpdateFrequency []DataRefreshInfo `tfsdk:"update_frequency" tf:"optional,object"`
}

type ListingFulfillment struct {
	FulfillmentType types.String `tfsdk:"fulfillment_type" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:""`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	RepoInfo []RepoInfo `tfsdk:"repo_info" tf:"optional,object"`

	ShareInfo []ShareInfo `tfsdk:"share_info" tf:"optional,object"`
}

type ListingSetting struct {
	Visibility types.String `tfsdk:"visibility" tf:"optional"`
}

// Next Number: 26
type ListingSummary struct {
	Categories []types.String `tfsdk:"categories" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	CreatedById types.Int64 `tfsdk:"created_by_id" tf:"optional"`

	ExchangeIds []types.String `tfsdk:"exchange_ids" tf:"optional"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo []RepoInfo `tfsdk:"git_repo" tf:"optional,object"`

	ListingType types.String `tfsdk:"listingType" tf:""`

	Name types.String `tfsdk:"name" tf:""`

	ProviderId types.String `tfsdk:"provider_id" tf:"optional"`

	ProviderRegion []RegionInfo `tfsdk:"provider_region" tf:"optional,object"`

	PublishedAt types.Int64 `tfsdk:"published_at" tf:"optional"`

	PublishedBy types.String `tfsdk:"published_by" tf:"optional"`

	Setting []ListingSetting `tfsdk:"setting" tf:"optional,object"`

	Share []ShareInfo `tfsdk:"share" tf:"optional,object"`
	// Enums
	Status types.String `tfsdk:"status" tf:"optional"`

	Subtitle types.String `tfsdk:"subtitle" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`

	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`

	UpdatedById types.Int64 `tfsdk:"updated_by_id" tf:"optional"`
}

type ListingTag struct {
	// Tag name (enum)
	TagName types.String `tfsdk:"tag_name" tf:"optional"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []types.String `tfsdk:"tag_values" tf:"optional"`
}

type PersonalizationRequest struct {
	Comment types.String `tfsdk:"comment" tf:"optional"`

	ConsumerRegion []RegionInfo `tfsdk:"consumer_region" tf:"object"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo []ContactInfo `tfsdk:"contact_info" tf:"optional,object"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	IntendedUse types.String `tfsdk:"intended_use" tf:"optional"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:"optional"`

	ListingName types.String `tfsdk:"listing_name" tf:"optional"`

	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`

	ProviderId types.String `tfsdk:"provider_id" tf:"optional"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	Share []ShareInfo `tfsdk:"share" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:"optional"`

	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

type ProviderAnalyticsDashboard struct {
	Id types.String `tfsdk:"id" tf:""`
}

type ProviderInfo struct {
	BusinessContactEmail types.String `tfsdk:"business_contact_email" tf:""`

	CompanyWebsiteLink types.String `tfsdk:"company_website_link" tf:"optional"`

	DarkModeIconFileId types.String `tfsdk:"dark_mode_icon_file_id" tf:"optional"`

	DarkModeIconFilePath types.String `tfsdk:"dark_mode_icon_file_path" tf:"optional"`

	Description types.String `tfsdk:"description" tf:"optional"`

	IconFileId types.String `tfsdk:"icon_file_id" tf:"optional"`

	IconFilePath types.String `tfsdk:"icon_file_path" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// is_featured is accessible by consumers only
	IsFeatured types.Bool `tfsdk:"is_featured" tf:"optional"`

	Name types.String `tfsdk:"name" tf:""`

	PrivacyPolicyLink types.String `tfsdk:"privacy_policy_link" tf:""`
	// published_by is only applicable to data aggregators (e.g. Crux)
	PublishedBy types.String `tfsdk:"published_by" tf:"optional"`

	SupportContactEmail types.String `tfsdk:"support_contact_email" tf:"optional"`

	TermOfServiceLink types.String `tfsdk:"term_of_service_link" tf:""`
}

type RegionInfo struct {
	Cloud types.String `tfsdk:"cloud" tf:"optional"`

	Region types.String `tfsdk:"region" tf:"optional"`
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {
	Id types.String `tfsdk:"-"`
}

type RemoveExchangeForListingResponse struct {
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url" tf:""`
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName types.String `tfsdk:"repo_name" tf:""`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path" tf:""`
}

// Search listings
type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets []types.String `tfsdk:"-"`
	// Matches any of the following categories
	Categories []types.String `tfsdk:"-"`

	IsFree types.Bool `tfsdk:"-"`

	IsPrivateExchange types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Matches any of the following provider ids
	ProviderIds []types.String `tfsdk:"-"`
	// Fuzzy matches query
	Query types.String `tfsdk:"-"`
}

type SearchListingsResponse struct {
	Listings []Listing `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type ShareInfo struct {
	Name types.String `tfsdk:"name" tf:""`

	Type types.String `tfsdk:"type" tf:""`
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Name of the shared object
	Name types.String `tfsdk:"name" tf:"optional"`
}

type TokenDetail struct {
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`

	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`

	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of Recipient Token creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// Unique id of the Recipient Token.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of Recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type UpdateExchangeFilterRequest struct {
	Filter []ExchangeFilter `tfsdk:"filter" tf:"object"`

	Id types.String `tfsdk:"-"`
}

type UpdateExchangeFilterResponse struct {
	Filter []ExchangeFilter `tfsdk:"filter" tf:"optional,object"`
}

type UpdateExchangeRequest struct {
	Exchange []Exchange `tfsdk:"exchange" tf:"object"`

	Id types.String `tfsdk:"-"`
}

type UpdateExchangeResponse struct {
	Exchange []Exchange `tfsdk:"exchange" tf:"optional,object"`
}

type UpdateInstallationRequest struct {
	Installation []InstallationDetail `tfsdk:"installation" tf:"object"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token" tf:"optional"`
}

type UpdateInstallationResponse struct {
	Installation []InstallationDetail `tfsdk:"installation" tf:"optional,object"`
}

type UpdateListingRequest struct {
	Id types.String `tfsdk:"-"`

	Listing []Listing `tfsdk:"listing" tf:"object"`
}

type UpdateListingResponse struct {
	Listing []Listing `tfsdk:"listing" tf:"optional,object"`
}

type UpdatePersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`

	Reason types.String `tfsdk:"reason" tf:"optional"`

	RequestId types.String `tfsdk:"-"`

	Share []ShareInfo `tfsdk:"share" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:""`
}

type UpdatePersonalizationRequestResponse struct {
	Request []PersonalizationRequest `tfsdk:"request" tf:"optional,object"`
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

type UpdateProviderRequest struct {
	Id types.String `tfsdk:"-"`

	Provider []ProviderInfo `tfsdk:"provider" tf:"object"`
}

type UpdateProviderResponse struct {
	Provider []ProviderInfo `tfsdk:"provider" tf:"optional,object"`
}
