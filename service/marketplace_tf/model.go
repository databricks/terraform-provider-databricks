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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddExchangeForListingRequest struct {
	ExchangeId types.String `tfsdk:"exchange_id"`

	ListingId types.String `tfsdk:"listing_id"`
}

type AddExchangeForListingResponse struct {
	ExchangeForListing *ExchangeListing `tfsdk:"exchange_for_listing"`
}

type AssetType string

const AssetTypeAssetTypeDataTable AssetType = `ASSET_TYPE_DATA_TABLE`

const AssetTypeAssetTypeGitRepo AssetType = `ASSET_TYPE_GIT_REPO`

const AssetTypeAssetTypeMedia AssetType = `ASSET_TYPE_MEDIA`

const AssetTypeAssetTypeModel AssetType = `ASSET_TYPE_MODEL`

const AssetTypeAssetTypeNotebook AssetType = `ASSET_TYPE_NOTEBOOK`

const AssetTypeAssetTypeUnspecified AssetType = `ASSET_TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *AssetType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AssetType) Set(v string) error {
	switch v {
	case `ASSET_TYPE_DATA_TABLE`, `ASSET_TYPE_GIT_REPO`, `ASSET_TYPE_MEDIA`, `ASSET_TYPE_MODEL`, `ASSET_TYPE_NOTEBOOK`, `ASSET_TYPE_UNSPECIFIED`:
		*f = AssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSET_TYPE_DATA_TABLE", "ASSET_TYPE_GIT_REPO", "ASSET_TYPE_MEDIA", "ASSET_TYPE_MODEL", "ASSET_TYPE_NOTEBOOK", "ASSET_TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns AssetType to satisfy [pflag.Value] interface
func (f *AssetType) Type() string {
	return "AssetType"
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest struct {
	Ids []types.String `tfsdk:"-" url:"ids,omitempty"`
}

type BatchGetListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest struct {
	Ids []types.String `tfsdk:"-" url:"ids,omitempty"`
}

type BatchGetProvidersResponse struct {
	Providers []ProviderInfo `tfsdk:"providers"`
}

type Category string

const CategoryAdvertisingAndMarketing Category = `ADVERTISING_AND_MARKETING`

const CategoryClimateAndEnvironment Category = `CLIMATE_AND_ENVIRONMENT`

const CategoryCommerce Category = `COMMERCE`

const CategoryDemographics Category = `DEMOGRAPHICS`

const CategoryEconomics Category = `ECONOMICS`

const CategoryEducation Category = `EDUCATION`

const CategoryEnergy Category = `ENERGY`

const CategoryFinancial Category = `FINANCIAL`

const CategoryGaming Category = `GAMING`

const CategoryGeospatial Category = `GEOSPATIAL`

const CategoryHealth Category = `HEALTH`

const CategoryLookupTables Category = `LOOKUP_TABLES`

const CategoryManufacturing Category = `MANUFACTURING`

const CategoryMedia Category = `MEDIA`

const CategoryOther Category = `OTHER`

const CategoryPublicSector Category = `PUBLIC_SECTOR`

const CategoryRetail Category = `RETAIL`

const CategoryScienceAndResearch Category = `SCIENCE_AND_RESEARCH`

const CategorySecurity Category = `SECURITY`

const CategorySports Category = `SPORTS`

const CategoryTransportationAndLogistics Category = `TRANSPORTATION_AND_LOGISTICS`

const CategoryTravelAndTourism Category = `TRAVEL_AND_TOURISM`

// String representation for [fmt.Print]
func (f *Category) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Category) Set(v string) error {
	switch v {
	case `ADVERTISING_AND_MARKETING`, `CLIMATE_AND_ENVIRONMENT`, `COMMERCE`, `DEMOGRAPHICS`, `ECONOMICS`, `EDUCATION`, `ENERGY`, `FINANCIAL`, `GAMING`, `GEOSPATIAL`, `HEALTH`, `LOOKUP_TABLES`, `MANUFACTURING`, `MEDIA`, `OTHER`, `PUBLIC_SECTOR`, `RETAIL`, `SCIENCE_AND_RESEARCH`, `SECURITY`, `SPORTS`, `TRANSPORTATION_AND_LOGISTICS`, `TRAVEL_AND_TOURISM`:
		*f = Category(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADVERTISING_AND_MARKETING", "CLIMATE_AND_ENVIRONMENT", "COMMERCE", "DEMOGRAPHICS", "ECONOMICS", "EDUCATION", "ENERGY", "FINANCIAL", "GAMING", "GEOSPATIAL", "HEALTH", "LOOKUP_TABLES", "MANUFACTURING", "MEDIA", "OTHER", "PUBLIC_SECTOR", "RETAIL", "SCIENCE_AND_RESEARCH", "SECURITY", "SPORTS", "TRANSPORTATION_AND_LOGISTICS", "TRAVEL_AND_TOURISM"`, v)
	}
}

// Type always returns Category to satisfy [pflag.Value] interface
func (f *Category) Type() string {
	return "Category"
}

type ConsumerTerms struct {
	Version types.String `tfsdk:"version"`
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company types.String `tfsdk:"company"`

	Email types.String `tfsdk:"email"`

	FirstName types.String `tfsdk:"first_name"`

	LastName types.String `tfsdk:"last_name"`
}

type Cost string

const CostFree Cost = `FREE`

const CostPaid Cost = `PAID`

// String representation for [fmt.Print]
func (f *Cost) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Cost) Set(v string) error {
	switch v {
	case `FREE`, `PAID`:
		*f = Cost(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FREE", "PAID"`, v)
	}
}

// Type always returns Cost to satisfy [pflag.Value] interface
func (f *Cost) Type() string {
	return "Cost"
}

type CreateExchangeFilterRequest struct {
	Filter ExchangeFilter `tfsdk:"filter"`
}

type CreateExchangeFilterResponse struct {
	FilterId types.String `tfsdk:"filter_id"`
}

type CreateExchangeRequest struct {
	Exchange Exchange `tfsdk:"exchange"`
}

type CreateExchangeResponse struct {
	ExchangeId types.String `tfsdk:"exchange_id"`
}

type CreateFileRequest struct {
	DisplayName types.String `tfsdk:"display_name"`

	FileParent FileParent `tfsdk:"file_parent"`

	MarketplaceFileType MarketplaceFileType `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`
}

type CreateFileResponse struct {
	FileInfo *FileInfo `tfsdk:"file_info"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url"`
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms *ConsumerTerms `tfsdk:"accepted_consumer_terms"`

	CatalogName types.String `tfsdk:"catalog_name"`

	ListingId types.String `tfsdk:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`
	// for git repo installations
	RepoDetail *RepoInstallation `tfsdk:"repo_detail"`

	ShareName types.String `tfsdk:"share_name"`
}

type CreateListingRequest struct {
	Listing Listing `tfsdk:"listing"`
}

type CreateListingResponse struct {
	ListingId types.String `tfsdk:"listing_id"`
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms ConsumerTerms `tfsdk:"accepted_consumer_terms"`

	Comment types.String `tfsdk:"comment"`

	Company types.String `tfsdk:"company"`

	FirstName types.String `tfsdk:"first_name"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	LastName types.String `tfsdk:"last_name"`

	ListingId types.String `tfsdk:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`
}

type CreatePersonalizationRequestResponse struct {
	Id types.String `tfsdk:"id"`
}

type CreateProviderRequest struct {
	Provider ProviderInfo `tfsdk:"provider"`
}

type CreateProviderResponse struct {
	Id types.String `tfsdk:"id"`
}

type DataRefresh string

const DataRefreshDaily DataRefresh = `DAILY`

const DataRefreshHourly DataRefresh = `HOURLY`

const DataRefreshMinute DataRefresh = `MINUTE`

const DataRefreshMonthly DataRefresh = `MONTHLY`

const DataRefreshNone DataRefresh = `NONE`

const DataRefreshQuarterly DataRefresh = `QUARTERLY`

const DataRefreshSecond DataRefresh = `SECOND`

const DataRefreshWeekly DataRefresh = `WEEKLY`

const DataRefreshYearly DataRefresh = `YEARLY`

// String representation for [fmt.Print]
func (f *DataRefresh) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataRefresh) Set(v string) error {
	switch v {
	case `DAILY`, `HOURLY`, `MINUTE`, `MONTHLY`, `NONE`, `QUARTERLY`, `SECOND`, `WEEKLY`, `YEARLY`:
		*f = DataRefresh(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAILY", "HOURLY", "MINUTE", "MONTHLY", "NONE", "QUARTERLY", "SECOND", "WEEKLY", "YEARLY"`, v)
	}
}

// Type always returns DataRefresh to satisfy [pflag.Value] interface
func (f *DataRefresh) Type() string {
	return "DataRefresh"
}

type DataRefreshInfo struct {
	Interval types.Int64 `tfsdk:"interval"`

	Unit DataRefresh `tfsdk:"unit"`
}

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteExchangeFilterResponse struct {
}

// Delete an exchange
type DeleteExchangeRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteExchangeResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	FileId types.String `tfsdk:"-" url:"-"`
}

type DeleteFileResponse struct {
}

// Uninstall from a listing
type DeleteInstallationRequest struct {
	InstallationId types.String `tfsdk:"-" url:"-"`

	ListingId types.String `tfsdk:"-" url:"-"`
}

type DeleteInstallationResponse struct {
}

// Delete a listing
type DeleteListingRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteListingResponse struct {
}

// Delete provider
type DeleteProviderRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteProviderResponse struct {
}

type DeltaSharingRecipientType string

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeDatabricks DeltaSharingRecipientType = `DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS`

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeOpen DeltaSharingRecipientType = `DELTA_SHARING_RECIPIENT_TYPE_OPEN`

// String representation for [fmt.Print]
func (f *DeltaSharingRecipientType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeltaSharingRecipientType) Set(v string) error {
	switch v {
	case `DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS`, `DELTA_SHARING_RECIPIENT_TYPE_OPEN`:
		*f = DeltaSharingRecipientType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS", "DELTA_SHARING_RECIPIENT_TYPE_OPEN"`, v)
	}
}

// Type always returns DeltaSharingRecipientType to satisfy [pflag.Value] interface
func (f *DeltaSharingRecipientType) Type() string {
	return "DeltaSharingRecipientType"
}

type Exchange struct {
	Comment types.String `tfsdk:"comment"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	Filters []ExchangeFilter `tfsdk:"filters"`

	Id types.String `tfsdk:"id"`

	LinkedListings []ExchangeListing `tfsdk:"linked_listings"`

	Name types.String `tfsdk:"name"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`
}

type ExchangeFilter struct {
	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	ExchangeId types.String `tfsdk:"exchange_id"`

	FilterType ExchangeFilterType `tfsdk:"filter_type"`

	FilterValue types.String `tfsdk:"filter_value"`

	Id types.String `tfsdk:"id"`

	Name types.String `tfsdk:"name"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`
}

type ExchangeFilterType string

const ExchangeFilterTypeGlobalMetastoreId ExchangeFilterType = `GLOBAL_METASTORE_ID`

// String representation for [fmt.Print]
func (f *ExchangeFilterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExchangeFilterType) Set(v string) error {
	switch v {
	case `GLOBAL_METASTORE_ID`:
		*f = ExchangeFilterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GLOBAL_METASTORE_ID"`, v)
	}
}

// Type always returns ExchangeFilterType to satisfy [pflag.Value] interface
func (f *ExchangeFilterType) Type() string {
	return "ExchangeFilterType"
}

type ExchangeListing struct {
	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	ExchangeId types.String `tfsdk:"exchange_id"`

	ExchangeName types.String `tfsdk:"exchange_name"`

	Id types.String `tfsdk:"id"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`
}

type FileInfo struct {
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName types.String `tfsdk:"display_name"`

	DownloadLink types.String `tfsdk:"download_link"`

	FileParent *FileParent `tfsdk:"file_parent"`

	Id types.String `tfsdk:"id"`

	MarketplaceFileType MarketplaceFileType `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`

	Status FileStatus `tfsdk:"status"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

type FileParent struct {
	FileParentType FileParentType `tfsdk:"file_parent_type"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id"`
}

type FileParentType string

const FileParentTypeListing FileParentType = `LISTING`

const FileParentTypeProvider FileParentType = `PROVIDER`

// String representation for [fmt.Print]
func (f *FileParentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileParentType) Set(v string) error {
	switch v {
	case `LISTING`, `PROVIDER`:
		*f = FileParentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING", "PROVIDER"`, v)
	}
}

// Type always returns FileParentType to satisfy [pflag.Value] interface
func (f *FileParentType) Type() string {
	return "FileParentType"
}

type FileStatus string

const FileStatusFileStatusPublished FileStatus = `FILE_STATUS_PUBLISHED`

const FileStatusFileStatusSanitizationFailed FileStatus = `FILE_STATUS_SANITIZATION_FAILED`

const FileStatusFileStatusSanitizing FileStatus = `FILE_STATUS_SANITIZING`

const FileStatusFileStatusStaging FileStatus = `FILE_STATUS_STAGING`

// String representation for [fmt.Print]
func (f *FileStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileStatus) Set(v string) error {
	switch v {
	case `FILE_STATUS_PUBLISHED`, `FILE_STATUS_SANITIZATION_FAILED`, `FILE_STATUS_SANITIZING`, `FILE_STATUS_STAGING`:
		*f = FileStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_STATUS_PUBLISHED", "FILE_STATUS_SANITIZATION_FAILED", "FILE_STATUS_SANITIZING", "FILE_STATUS_STAGING"`, v)
	}
}

// Type always returns FileStatus to satisfy [pflag.Value] interface
func (f *FileStatus) Type() string {
	return "FileStatus"
}

type FilterType string

const FilterTypeMetastore FilterType = `METASTORE`

// String representation for [fmt.Print]
func (f *FilterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FilterType) Set(v string) error {
	switch v {
	case `METASTORE`:
		*f = FilterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "METASTORE"`, v)
	}
}

// Type always returns FilterType to satisfy [pflag.Value] interface
func (f *FilterType) Type() string {
	return "FilterType"
}

type FulfillmentType string

const FulfillmentTypeInstall FulfillmentType = `INSTALL`

const FulfillmentTypeRequestAccess FulfillmentType = `REQUEST_ACCESS`

// String representation for [fmt.Print]
func (f *FulfillmentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FulfillmentType) Set(v string) error {
	switch v {
	case `INSTALL`, `REQUEST_ACCESS`:
		*f = FulfillmentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INSTALL", "REQUEST_ACCESS"`, v)
	}
}

// Type always returns FulfillmentType to satisfy [pflag.Value] interface
func (f *FulfillmentType) Type() string {
	return "FulfillmentType"
}

// Get an exchange
type GetExchangeRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type GetExchangeResponse struct {
	Exchange *Exchange `tfsdk:"exchange"`
}

// Get a file
type GetFileRequest struct {
	FileId types.String `tfsdk:"-" url:"-"`
}

type GetFileResponse struct {
	FileInfo *FileInfo `tfsdk:"file_info"`
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

// Get listing content metadata
type GetListingContentMetadataRequest struct {
	ListingId types.String `tfsdk:"-" url:"-"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type GetListingContentMetadataResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	SharedDataObjects []SharedDataObject `tfsdk:"shared_data_objects"`
}

// Get listing
type GetListingRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type GetListingResponse struct {
	Listing *Listing `tfsdk:"listing"`
}

// List listings
type GetListingsRequest struct {
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type GetListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-" url:"-"`
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests"`
}

// Get a provider
type GetProviderRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type GetProviderResponse struct {
	Provider *ProviderInfo `tfsdk:"provider"`
}

type Installation struct {
	Installation *InstallationDetail `tfsdk:"installation"`
}

type InstallationDetail struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ErrorMessage types.String `tfsdk:"error_message"`

	Id types.String `tfsdk:"id"`

	InstalledOn types.Int64 `tfsdk:"installed_on"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	RepoName types.String `tfsdk:"repo_name"`

	RepoPath types.String `tfsdk:"repo_path"`

	ShareName types.String `tfsdk:"share_name"`

	Status InstallationStatus `tfsdk:"status"`

	TokenDetail *TokenDetail `tfsdk:"token_detail"`

	Tokens []TokenInfo `tfsdk:"tokens"`
}

type InstallationStatus string

const InstallationStatusFailed InstallationStatus = `FAILED`

const InstallationStatusInstalled InstallationStatus = `INSTALLED`

// String representation for [fmt.Print]
func (f *InstallationStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstallationStatus) Set(v string) error {
	switch v {
	case `FAILED`, `INSTALLED`:
		*f = InstallationStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "INSTALLED"`, v)
	}
}

// Type always returns InstallationStatus to satisfy [pflag.Value] interface
func (f *InstallationStatus) Type() string {
	return "InstallationStatus"
}

// List all installations
type ListAllInstallationsRequest struct {
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListAllInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests"`
}

// List exchange filters
type ListExchangeFiltersRequest struct {
	ExchangeId types.String `tfsdk:"-" url:"exchange_id"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListExchangeFiltersResponse struct {
	Filters []ExchangeFilter `tfsdk:"filters"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List exchanges for listing
type ListExchangesForListingRequest struct {
	ListingId types.String `tfsdk:"-" url:"listing_id"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListExchangesForListingResponse struct {
	ExchangeListing []ExchangeListing `tfsdk:"exchange_listing"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List exchanges
type ListExchangesRequest struct {
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListExchangesResponse struct {
	Exchanges []Exchange `tfsdk:"exchanges"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List files
type ListFilesRequest struct {
	FileParent FileParent `tfsdk:"-" url:"file_parent"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListFilesResponse struct {
	FileInfos []FileInfo `tfsdk:"file_infos"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List all listing fulfillments
type ListFulfillmentsRequest struct {
	ListingId types.String `tfsdk:"-" url:"-"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListFulfillmentsResponse struct {
	Fulfillments []ListingFulfillment `tfsdk:"fulfillments"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List installations for a listing
type ListInstallationsRequest struct {
	ListingId types.String `tfsdk:"-" url:"-"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List listings for exchange
type ListListingsForExchangeRequest struct {
	ExchangeId types.String `tfsdk:"-" url:"exchange_id"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListListingsForExchangeResponse struct {
	ExchangeListings []ExchangeListing `tfsdk:"exchange_listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List listings
type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `tfsdk:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `tfsdk:"-" url:"categories,omitempty"`

	IsAscending types.Bool `tfsdk:"-" url:"is_ascending,omitempty"`
	// Filters each listing based on if it is free.
	IsFree types.Bool `tfsdk:"-" url:"is_free,omitempty"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange types.Bool `tfsdk:"-" url:"is_private_exchange,omitempty"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick types.Bool `tfsdk:"-" url:"is_staff_pick,omitempty"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []types.String `tfsdk:"-" url:"provider_ids,omitempty"`
	// Criteria for sorting the resulting set of listings.
	SortBy SortBy `tfsdk:"-" url:"sort_by,omitempty"`
	// Matches any of the following tags
	Tags []ListingTag `tfsdk:"-" url:"tags,omitempty"`
}

type ListListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`

	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

// List providers
type ListProvidersRequest struct {
	IsFeatured types.Bool `tfsdk:"-" url:"is_featured,omitempty"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListProvidersResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Providers []ProviderInfo `tfsdk:"providers"`
}

type Listing struct {
	Detail *ListingDetail `tfsdk:"detail"`

	Id types.String `tfsdk:"id"`
	// we can not use just ProviderListingSummary since we already have same
	// name on entity side of the state
	ProviderSummary *ProviderListingSummaryInfo `tfsdk:"provider_summary"`
	// Next Number: 26
	Summary ListingSummary `tfsdk:"summary"`
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets []AssetType `tfsdk:"assets"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd types.Int64 `tfsdk:"collection_date_end"`
	// The starting date timestamp for when the data spans
	CollectionDateStart types.Int64 `tfsdk:"collection_date_start"`
	// Smallest unit of time in the dataset
	CollectionGranularity *DataRefreshInfo `tfsdk:"collection_granularity"`
	// Whether the dataset is free or paid
	Cost Cost `tfsdk:"cost"`
	// Where/how the data is sourced
	DataSource types.String `tfsdk:"data_source"`

	Description types.String `tfsdk:"description"`

	DocumentationLink types.String `tfsdk:"documentation_link"`

	EmbeddedNotebookFileInfos []FileInfo `tfsdk:"embedded_notebook_file_infos"`

	FileIds []types.String `tfsdk:"file_ids"`
	// Which geo region the listing data is collected from
	GeographicalCoverage types.String `tfsdk:"geographical_coverage"`
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	License types.String `tfsdk:"license"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	PricingModel types.String `tfsdk:"pricing_model"`

	PrivacyPolicyLink types.String `tfsdk:"privacy_policy_link"`
	// size of the dataset in GB
	Size types.Float64 `tfsdk:"size"`

	SupportLink types.String `tfsdk:"support_link"`
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	Tags []ListingTag `tfsdk:"tags"`

	TermsOfService types.String `tfsdk:"terms_of_service"`
	// How often data is updated
	UpdateFrequency *DataRefreshInfo `tfsdk:"update_frequency"`
}

type ListingFulfillment struct {
	FulfillmentType FulfillmentType `tfsdk:"fulfillment_type"`

	ListingId types.String `tfsdk:"listing_id"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	RepoInfo *RepoInfo `tfsdk:"repo_info"`

	ShareInfo *ShareInfo `tfsdk:"share_info"`
}

type ListingSetting struct {
	// filters are joined with `or` conjunction.
	Filters []VisibilityFilter `tfsdk:"filters"`

	Visibility Visibility `tfsdk:"visibility"`
}

type ListingShareType string

const ListingShareTypeFull ListingShareType = `FULL`

const ListingShareTypeSample ListingShareType = `SAMPLE`

// String representation for [fmt.Print]
func (f *ListingShareType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingShareType) Set(v string) error {
	switch v {
	case `FULL`, `SAMPLE`:
		*f = ListingShareType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL", "SAMPLE"`, v)
	}
}

// Type always returns ListingShareType to satisfy [pflag.Value] interface
func (f *ListingShareType) Type() string {
	return "ListingShareType"
}

// Enums
type ListingStatus string

const ListingStatusDraft ListingStatus = `DRAFT`

const ListingStatusPending ListingStatus = `PENDING`

const ListingStatusPublished ListingStatus = `PUBLISHED`

const ListingStatusSuspended ListingStatus = `SUSPENDED`

// String representation for [fmt.Print]
func (f *ListingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingStatus) Set(v string) error {
	switch v {
	case `DRAFT`, `PENDING`, `PUBLISHED`, `SUSPENDED`:
		*f = ListingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DRAFT", "PENDING", "PUBLISHED", "SUSPENDED"`, v)
	}
}

// Type always returns ListingStatus to satisfy [pflag.Value] interface
func (f *ListingStatus) Type() string {
	return "ListingStatus"
}

// Next Number: 26
type ListingSummary struct {
	Categories []Category `tfsdk:"categories"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	CreatedById types.Int64 `tfsdk:"created_by_id"`

	ExchangeIds []types.String `tfsdk:"exchange_ids"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo *RepoInfo `tfsdk:"git_repo"`

	ListingType ListingType `tfsdk:"listingType"`

	MetastoreId types.String `tfsdk:"metastore_id"`

	Name types.String `tfsdk:"name"`

	ProviderId types.String `tfsdk:"provider_id"`

	ProviderRegion *RegionInfo `tfsdk:"provider_region"`

	PublishedAt types.Int64 `tfsdk:"published_at"`

	PublishedBy types.String `tfsdk:"published_by"`

	Setting *ListingSetting `tfsdk:"setting"`

	Share *ShareInfo `tfsdk:"share"`
	// Enums
	Status ListingStatus `tfsdk:"status"`

	Subtitle types.String `tfsdk:"subtitle"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`

	UpdatedById types.Int64 `tfsdk:"updated_by_id"`
}

type ListingTag struct {
	// Tag name (enum)
	TagName ListingTagType `tfsdk:"tag_name"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []types.String `tfsdk:"tag_values"`
}

type ListingTagType string

const ListingTagTypeListingTagTypeLanguage ListingTagType = `LISTING_TAG_TYPE_LANGUAGE`

const ListingTagTypeListingTagTypeTask ListingTagType = `LISTING_TAG_TYPE_TASK`

const ListingTagTypeListingTagTypeUnspecified ListingTagType = `LISTING_TAG_TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ListingTagType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingTagType) Set(v string) error {
	switch v {
	case `LISTING_TAG_TYPE_LANGUAGE`, `LISTING_TAG_TYPE_TASK`, `LISTING_TAG_TYPE_UNSPECIFIED`:
		*f = ListingTagType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING_TAG_TYPE_LANGUAGE", "LISTING_TAG_TYPE_TASK", "LISTING_TAG_TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns ListingTagType to satisfy [pflag.Value] interface
func (f *ListingTagType) Type() string {
	return "ListingTagType"
}

type ListingType string

const ListingTypePersonalized ListingType = `PERSONALIZED`

const ListingTypeStandard ListingType = `STANDARD`

// String representation for [fmt.Print]
func (f *ListingType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingType) Set(v string) error {
	switch v {
	case `PERSONALIZED`, `STANDARD`:
		*f = ListingType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PERSONALIZED", "STANDARD"`, v)
	}
}

// Type always returns ListingType to satisfy [pflag.Value] interface
func (f *ListingType) Type() string {
	return "ListingType"
}

type MarketplaceFileType string

const MarketplaceFileTypeEmbeddedNotebook MarketplaceFileType = `EMBEDDED_NOTEBOOK`

const MarketplaceFileTypeProviderIcon MarketplaceFileType = `PROVIDER_ICON`

// String representation for [fmt.Print]
func (f *MarketplaceFileType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MarketplaceFileType) Set(v string) error {
	switch v {
	case `EMBEDDED_NOTEBOOK`, `PROVIDER_ICON`:
		*f = MarketplaceFileType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMBEDDED_NOTEBOOK", "PROVIDER_ICON"`, v)
	}
}

// Type always returns MarketplaceFileType to satisfy [pflag.Value] interface
func (f *MarketplaceFileType) Type() string {
	return "MarketplaceFileType"
}

type PersonalizationRequest struct {
	Comment types.String `tfsdk:"comment"`

	ConsumerRegion RegionInfo `tfsdk:"consumer_region"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo *ContactInfo `tfsdk:"contact_info"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	Id types.String `tfsdk:"id"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`

	MetastoreId types.String `tfsdk:"metastore_id"`

	ProviderId types.String `tfsdk:"provider_id"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	Share *ShareInfo `tfsdk:"share"`

	Status PersonalizationRequestStatus `tfsdk:"status"`

	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

type PersonalizationRequestStatus string

const PersonalizationRequestStatusDenied PersonalizationRequestStatus = `DENIED`

const PersonalizationRequestStatusFulfilled PersonalizationRequestStatus = `FULFILLED`

const PersonalizationRequestStatusNew PersonalizationRequestStatus = `NEW`

const PersonalizationRequestStatusRequestPending PersonalizationRequestStatus = `REQUEST_PENDING`

// String representation for [fmt.Print]
func (f *PersonalizationRequestStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalizationRequestStatus) Set(v string) error {
	switch v {
	case `DENIED`, `FULFILLED`, `NEW`, `REQUEST_PENDING`:
		*f = PersonalizationRequestStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DENIED", "FULFILLED", "NEW", "REQUEST_PENDING"`, v)
	}
}

// Type always returns PersonalizationRequestStatus to satisfy [pflag.Value] interface
func (f *PersonalizationRequestStatus) Type() string {
	return "PersonalizationRequestStatus"
}

type ProviderAnalyticsDashboard struct {
	Id types.String `tfsdk:"id"`
}

type ProviderIconFile struct {
	IconFileId types.String `tfsdk:"icon_file_id"`

	IconFilePath types.String `tfsdk:"icon_file_path"`

	IconType ProviderIconType `tfsdk:"icon_type"`
}

type ProviderIconType string

const ProviderIconTypeDark ProviderIconType = `DARK`

const ProviderIconTypePrimary ProviderIconType = `PRIMARY`

const ProviderIconTypeProviderIconTypeUnspecified ProviderIconType = `PROVIDER_ICON_TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ProviderIconType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProviderIconType) Set(v string) error {
	switch v {
	case `DARK`, `PRIMARY`, `PROVIDER_ICON_TYPE_UNSPECIFIED`:
		*f = ProviderIconType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DARK", "PRIMARY", "PROVIDER_ICON_TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns ProviderIconType to satisfy [pflag.Value] interface
func (f *ProviderIconType) Type() string {
	return "ProviderIconType"
}

type ProviderInfo struct {
	BusinessContactEmail types.String `tfsdk:"business_contact_email"`

	CompanyWebsiteLink types.String `tfsdk:"company_website_link"`

	DarkModeIconFileId types.String `tfsdk:"dark_mode_icon_file_id"`

	DarkModeIconFilePath types.String `tfsdk:"dark_mode_icon_file_path"`

	Description types.String `tfsdk:"description"`

	IconFileId types.String `tfsdk:"icon_file_id"`

	IconFilePath types.String `tfsdk:"icon_file_path"`

	Id types.String `tfsdk:"id"`
	// is_featured is accessible by consumers only
	IsFeatured types.Bool `tfsdk:"is_featured"`

	Name types.String `tfsdk:"name"`

	PrivacyPolicyLink types.String `tfsdk:"privacy_policy_link"`
	// published_by is only applicable to data aggregators (e.g. Crux)
	PublishedBy types.String `tfsdk:"published_by"`

	SupportContactEmail types.String `tfsdk:"support_contact_email"`

	TermOfServiceLink types.String `tfsdk:"term_of_service_link"`
}

// we can not use just ProviderListingSummary since we already have same name on
// entity side of the state
type ProviderListingSummaryInfo struct {
	Description types.String `tfsdk:"description"`

	IconFiles []ProviderIconFile `tfsdk:"icon_files"`

	Name types.String `tfsdk:"name"`
}

type RegionInfo struct {
	Cloud types.String `tfsdk:"cloud"`

	Region types.String `tfsdk:"region"`
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {
	Id types.String `tfsdk:"-" url:"-"`
}

type RemoveExchangeForListingResponse struct {
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url"`
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName types.String `tfsdk:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path"`
}

// Search listings
type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `tfsdk:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `tfsdk:"-" url:"categories,omitempty"`

	IsAscending types.Bool `tfsdk:"-" url:"is_ascending,omitempty"`

	IsFree types.Bool `tfsdk:"-" url:"is_free,omitempty"`

	IsPrivateExchange types.Bool `tfsdk:"-" url:"is_private_exchange,omitempty"`

	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`

	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []types.String `tfsdk:"-" url:"provider_ids,omitempty"`
	// Fuzzy matches query
	Query types.String `tfsdk:"-" url:"query"`

	SortBy SortBy `tfsdk:"-" url:"sort_by,omitempty"`
}

type SearchListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

type ShareInfo struct {
	Name types.String `tfsdk:"name"`

	Type ListingShareType `tfsdk:"type"`
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType types.String `tfsdk:"data_object_type"`
	// Name of the shared object
	Name types.String `tfsdk:"name"`
}

type SortBy string

const SortBySortByDate SortBy = `SORT_BY_DATE`

const SortBySortByRelevance SortBy = `SORT_BY_RELEVANCE`

const SortBySortByTitle SortBy = `SORT_BY_TITLE`

const SortBySortByUnspecified SortBy = `SORT_BY_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *SortBy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SortBy) Set(v string) error {
	switch v {
	case `SORT_BY_DATE`, `SORT_BY_RELEVANCE`, `SORT_BY_TITLE`, `SORT_BY_UNSPECIFIED`:
		*f = SortBy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SORT_BY_DATE", "SORT_BY_RELEVANCE", "SORT_BY_TITLE", "SORT_BY_UNSPECIFIED"`, v)
	}
}

// Type always returns SortBy to satisfy [pflag.Value] interface
func (f *SortBy) Type() string {
	return "SortBy"
}

type TokenDetail struct {
	BearerToken types.String `tfsdk:"bearerToken"`

	Endpoint types.String `tfsdk:"endpoint"`

	ExpirationTime types.String `tfsdk:"expirationTime"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion"`
}

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of Recipient Token creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// Unique id of the Recipient Token.
	Id types.String `tfsdk:"id"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of Recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

type UpdateExchangeFilterRequest struct {
	Filter ExchangeFilter `tfsdk:"filter"`

	Id types.String `tfsdk:"-" url:"-"`
}

type UpdateExchangeFilterResponse struct {
	Filter *ExchangeFilter `tfsdk:"filter"`
}

type UpdateExchangeRequest struct {
	Exchange Exchange `tfsdk:"exchange"`

	Id types.String `tfsdk:"-" url:"-"`
}

type UpdateExchangeResponse struct {
	Exchange *Exchange `tfsdk:"exchange"`
}

type UpdateInstallationRequest struct {
	Installation InstallationDetail `tfsdk:"installation"`

	InstallationId types.String `tfsdk:"-" url:"-"`

	ListingId types.String `tfsdk:"-" url:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token"`
}

type UpdateInstallationResponse struct {
	Installation *InstallationDetail `tfsdk:"installation"`
}

type UpdateListingRequest struct {
	Id types.String `tfsdk:"-" url:"-"`

	Listing Listing `tfsdk:"listing"`
}

type UpdateListingResponse struct {
	Listing *Listing `tfsdk:"listing"`
}

type UpdatePersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-" url:"-"`

	Reason types.String `tfsdk:"reason"`

	RequestId types.String `tfsdk:"-" url:"-"`

	Share *ShareInfo `tfsdk:"share"`

	Status PersonalizationRequestStatus `tfsdk:"status"`
}

type UpdatePersonalizationRequestResponse struct {
	Request *PersonalizationRequest `tfsdk:"request"`
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-" url:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId types.String `tfsdk:"dashboard_id"`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

type UpdateProviderRequest struct {
	Id types.String `tfsdk:"-" url:"-"`

	Provider ProviderInfo `tfsdk:"provider"`
}

type UpdateProviderResponse struct {
	Provider *ProviderInfo `tfsdk:"provider"`
}

type Visibility string

const VisibilityPrivate Visibility = `PRIVATE`

const VisibilityPublic Visibility = `PUBLIC`

// String representation for [fmt.Print]
func (f *Visibility) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Visibility) Set(v string) error {
	switch v {
	case `PRIVATE`, `PUBLIC`:
		*f = Visibility(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRIVATE", "PUBLIC"`, v)
	}
}

// Type always returns Visibility to satisfy [pflag.Value] interface
func (f *Visibility) Type() string {
	return "Visibility"
}

type VisibilityFilter struct {
	FilterType FilterType `tfsdk:"filterType"`

	FilterValue types.String `tfsdk:"filterValue"`
}
