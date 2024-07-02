// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddExchangeForListingRequest struct {
	ExchangeId string `tfsdk:"exchange_id"`

	ListingId string `tfsdk:"listing_id"`
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
	Version string `tfsdk:"version"`
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company string `tfsdk:"company"`

	Email string `tfsdk:"email"`

	FirstName string `tfsdk:"first_name"`

	LastName string `tfsdk:"last_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ContactInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContactInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	FilterId string `tfsdk:"filter_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExchangeRequest struct {
	Exchange Exchange `tfsdk:"exchange"`
}

type CreateExchangeResponse struct {
	ExchangeId string `tfsdk:"exchange_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileRequest struct {
	DisplayName string `tfsdk:"display_name"`

	FileParent FileParent `tfsdk:"file_parent"`

	MarketplaceFileType MarketplaceFileType `tfsdk:"marketplace_file_type"`

	MimeType string `tfsdk:"mime_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileResponse struct {
	FileInfo *FileInfo `tfsdk:"file_info"`
	// Pre-signed POST URL to blob storage
	UploadUrl string `tfsdk:"upload_url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms *ConsumerTerms `tfsdk:"accepted_consumer_terms"`

	CatalogName string `tfsdk:"catalog_name"`

	ListingId string `tfsdk:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`
	// for git repo installations
	RepoDetail *RepoInstallation `tfsdk:"repo_detail"`

	ShareName string `tfsdk:"share_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateListingRequest struct {
	Listing Listing `tfsdk:"listing"`
}

type CreateListingResponse struct {
	ListingId string `tfsdk:"listing_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms ConsumerTerms `tfsdk:"accepted_consumer_terms"`

	Comment string `tfsdk:"comment"`

	Company string `tfsdk:"company"`

	FirstName string `tfsdk:"first_name"`

	IntendedUse string `tfsdk:"intended_use"`

	IsFromLighthouse bool `tfsdk:"is_from_lighthouse"`

	LastName string `tfsdk:"last_name"`

	ListingId string `tfsdk:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreatePersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePersonalizationRequestResponse struct {
	Id string `tfsdk:"id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateProviderRequest struct {
	Provider ProviderInfo `tfsdk:"provider"`
}

type CreateProviderResponse struct {
	Id string `tfsdk:"id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateProviderResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProviderResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Interval int64 `tfsdk:"interval"`

	Unit DataRefresh `tfsdk:"unit"`
}

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type DeleteExchangeFilterResponse struct {
}

// Delete an exchange
type DeleteExchangeRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type DeleteExchangeResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	FileId string `tfsdk:"-" url:"-"`
}

type DeleteFileResponse struct {
}

// Uninstall from a listing
type DeleteInstallationRequest struct {
	InstallationId string `tfsdk:"-" url:"-"`

	ListingId string `tfsdk:"-" url:"-"`
}

type DeleteInstallationResponse struct {
}

// Delete a listing
type DeleteListingRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type DeleteListingResponse struct {
}

// Delete provider
type DeleteProviderRequest struct {
	Id string `tfsdk:"-" url:"-"`
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
	Comment string `tfsdk:"comment"`

	CreatedAt int64 `tfsdk:"created_at"`

	CreatedBy string `tfsdk:"created_by"`

	Filters []ExchangeFilter `tfsdk:"filters"`

	Id string `tfsdk:"id"`

	LinkedListings []ExchangeListing `tfsdk:"linked_listings"`

	Name string `tfsdk:"name"`

	UpdatedAt int64 `tfsdk:"updated_at"`

	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Exchange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Exchange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExchangeFilter struct {
	CreatedAt int64 `tfsdk:"created_at"`

	CreatedBy string `tfsdk:"created_by"`

	ExchangeId string `tfsdk:"exchange_id"`

	FilterType ExchangeFilterType `tfsdk:"filter_type"`

	FilterValue string `tfsdk:"filter_value"`

	Id string `tfsdk:"id"`

	Name string `tfsdk:"name"`

	UpdatedAt int64 `tfsdk:"updated_at"`

	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExchangeFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	CreatedAt int64 `tfsdk:"created_at"`

	CreatedBy string `tfsdk:"created_by"`

	ExchangeId string `tfsdk:"exchange_id"`

	ExchangeName string `tfsdk:"exchange_name"`

	Id string `tfsdk:"id"`

	ListingId string `tfsdk:"listing_id"`

	ListingName string `tfsdk:"listing_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExchangeListing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeListing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileInfo struct {
	CreatedAt int64 `tfsdk:"created_at"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName string `tfsdk:"display_name"`

	DownloadLink string `tfsdk:"download_link"`

	FileParent *FileParent `tfsdk:"file_parent"`

	Id string `tfsdk:"id"`

	MarketplaceFileType MarketplaceFileType `tfsdk:"marketplace_file_type"`

	MimeType string `tfsdk:"mime_type"`

	Status FileStatus `tfsdk:"status"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage string `tfsdk:"status_message"`

	UpdatedAt int64 `tfsdk:"updated_at"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileParent struct {
	FileParentType FileParentType `tfsdk:"file_parent_type"`
	// TODO make the following fields required
	ParentId string `tfsdk:"parent_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FileParent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileParent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Id string `tfsdk:"-" url:"-"`
}

type GetExchangeResponse struct {
	Exchange *Exchange `tfsdk:"exchange"`
}

// Get a file
type GetFileRequest struct {
	FileId string `tfsdk:"-" url:"-"`
}

type GetFileResponse struct {
	FileInfo *FileInfo `tfsdk:"file_info"`
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version int64 `tfsdk:"version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetLatestVersionProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLatestVersionProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get listing content metadata
type GetListingContentMetadataRequest struct {
	ListingId string `tfsdk:"-" url:"-"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetListingContentMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingContentMetadataResponse struct {
	NextPageToken string `tfsdk:"next_page_token"`

	SharedDataObjects []SharedDataObject `tfsdk:"shared_data_objects"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetListingContentMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get listing
type GetListingRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type GetListingResponse struct {
	Listing *Listing `tfsdk:"listing"`
}

// List listings
type GetListingsRequest struct {
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {
	ListingId string `tfsdk:"-" url:"-"`
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests"`
}

// Get a provider
type GetProviderRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type GetProviderResponse struct {
	Provider *ProviderInfo `tfsdk:"provider"`
}

type Installation struct {
	Installation *InstallationDetail `tfsdk:"installation"`
}

type InstallationDetail struct {
	CatalogName string `tfsdk:"catalog_name"`

	ErrorMessage string `tfsdk:"error_message"`

	Id string `tfsdk:"id"`

	InstalledOn int64 `tfsdk:"installed_on"`

	ListingId string `tfsdk:"listing_id"`

	ListingName string `tfsdk:"listing_name"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	RepoName string `tfsdk:"repo_name"`

	RepoPath string `tfsdk:"repo_path"`

	ShareName string `tfsdk:"share_name"`

	Status InstallationStatus `tfsdk:"status"`

	TokenDetail *TokenDetail `tfsdk:"token_detail"`

	Tokens []TokenInfo `tfsdk:"tokens"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *InstallationDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstallationDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAllInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAllInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAllPersonalizationRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken string `tfsdk:"next_page_token"`

	PersonalizationRequests []PersonalizationRequest `tfsdk:"personalization_requests"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAllPersonalizationRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchange filters
type ListExchangeFiltersRequest struct {
	ExchangeId string `tfsdk:"-" url:"exchange_id"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangeFiltersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangeFiltersResponse struct {
	Filters []ExchangeFilter `tfsdk:"filters"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangeFiltersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchanges for listing
type ListExchangesForListingRequest struct {
	ListingId string `tfsdk:"-" url:"listing_id"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangesForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesForListingResponse struct {
	ExchangeListing []ExchangeListing `tfsdk:"exchange_listing"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangesForListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchanges
type ListExchangesRequest struct {
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesResponse struct {
	Exchanges []Exchange `tfsdk:"exchanges"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExchangesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List files
type ListFilesRequest struct {
	FileParent FileParent `tfsdk:"-" url:"file_parent"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFilesResponse struct {
	FileInfos []FileInfo `tfsdk:"file_infos"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all listing fulfillments
type ListFulfillmentsRequest struct {
	ListingId string `tfsdk:"-" url:"-"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFulfillmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFulfillmentsResponse struct {
	Fulfillments []ListingFulfillment `tfsdk:"fulfillments"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFulfillmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List installations for a listing
type ListInstallationsRequest struct {
	ListingId string `tfsdk:"-" url:"-"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListInstallationsResponse struct {
	Installations []InstallationDetail `tfsdk:"installations"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List listings for exchange
type ListListingsForExchangeRequest struct {
	ExchangeId string `tfsdk:"-" url:"exchange_id"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListListingsForExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsForExchangeResponse struct {
	ExchangeListings []ExchangeListing `tfsdk:"exchange_listings"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListListingsForExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List listings
type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `tfsdk:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `tfsdk:"-" url:"categories,omitempty"`

	IsAscending bool `tfsdk:"-" url:"is_ascending,omitempty"`
	// Filters each listing based on if it is free.
	IsFree bool `tfsdk:"-" url:"is_free,omitempty"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange bool `tfsdk:"-" url:"is_private_exchange,omitempty"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick bool `tfsdk:"-" url:"is_staff_pick,omitempty"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `tfsdk:"-" url:"provider_ids,omitempty"`
	// Criteria for sorting the resulting set of listings.
	SortBy SortBy `tfsdk:"-" url:"sort_by,omitempty"`
	// Matches any of the following tags
	Tags []ListingTag `tfsdk:"-" url:"tags,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId string `tfsdk:"dashboard_id"`

	Id string `tfsdk:"id"`

	Version int64 `tfsdk:"version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List providers
type ListProvidersRequest struct {
	IsFeatured bool `tfsdk:"-" url:"is_featured,omitempty"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersResponse struct {
	NextPageToken string `tfsdk:"next_page_token"`

	Providers []ProviderInfo `tfsdk:"providers"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Listing struct {
	Detail *ListingDetail `tfsdk:"detail"`

	Id string `tfsdk:"id"`
	// Next Number: 26
	Summary ListingSummary `tfsdk:"summary"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Listing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Listing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets []AssetType `tfsdk:"assets"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd int64 `tfsdk:"collection_date_end"`
	// The starting date timestamp for when the data spans
	CollectionDateStart int64 `tfsdk:"collection_date_start"`
	// Smallest unit of time in the dataset
	CollectionGranularity *DataRefreshInfo `tfsdk:"collection_granularity"`
	// Whether the dataset is free or paid
	Cost Cost `tfsdk:"cost"`
	// Where/how the data is sourced
	DataSource string `tfsdk:"data_source"`

	Description string `tfsdk:"description"`

	DocumentationLink string `tfsdk:"documentation_link"`

	EmbeddedNotebookFileInfos []FileInfo `tfsdk:"embedded_notebook_file_infos"`

	FileIds []string `tfsdk:"file_ids"`
	// Which geo region the listing data is collected from
	GeographicalCoverage string `tfsdk:"geographical_coverage"`
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	License string `tfsdk:"license"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	PricingModel string `tfsdk:"pricing_model"`

	PrivacyPolicyLink string `tfsdk:"privacy_policy_link"`
	// size of the dataset in GB
	Size float64 `tfsdk:"size"`

	SupportLink string `tfsdk:"support_link"`
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	Tags []ListingTag `tfsdk:"tags"`

	TermsOfService string `tfsdk:"terms_of_service"`
	// How often data is updated
	UpdateFrequency *DataRefreshInfo `tfsdk:"update_frequency"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListingDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingFulfillment struct {
	FulfillmentType FulfillmentType `tfsdk:"fulfillment_type"`

	ListingId string `tfsdk:"listing_id"`

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

	CreatedAt int64 `tfsdk:"created_at"`

	CreatedBy string `tfsdk:"created_by"`

	CreatedById int64 `tfsdk:"created_by_id"`

	ExchangeIds []string `tfsdk:"exchange_ids"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo *RepoInfo `tfsdk:"git_repo"`

	ListingType ListingType `tfsdk:"listingType"`

	MetastoreId string `tfsdk:"metastore_id"`

	Name string `tfsdk:"name"`

	ProviderId string `tfsdk:"provider_id"`

	ProviderRegion *RegionInfo `tfsdk:"provider_region"`

	PublishedAt int64 `tfsdk:"published_at"`

	PublishedBy string `tfsdk:"published_by"`

	Setting *ListingSetting `tfsdk:"setting"`

	Share *ShareInfo `tfsdk:"share"`
	// Enums
	Status ListingStatus `tfsdk:"status"`

	Subtitle string `tfsdk:"subtitle"`

	UpdatedAt int64 `tfsdk:"updated_at"`

	UpdatedBy string `tfsdk:"updated_by"`

	UpdatedById int64 `tfsdk:"updated_by_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListingSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingTag struct {
	// Tag name (enum)
	TagName ListingTagType `tfsdk:"tag_name"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []string `tfsdk:"tag_values"`
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
	Comment string `tfsdk:"comment"`

	ConsumerRegion RegionInfo `tfsdk:"consumer_region"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo *ContactInfo `tfsdk:"contact_info"`

	CreatedAt int64 `tfsdk:"created_at"`

	Id string `tfsdk:"id"`

	IntendedUse string `tfsdk:"intended_use"`

	IsFromLighthouse bool `tfsdk:"is_from_lighthouse"`

	ListingId string `tfsdk:"listing_id"`

	ListingName string `tfsdk:"listing_name"`

	MetastoreId string `tfsdk:"metastore_id"`

	ProviderId string `tfsdk:"provider_id"`

	RecipientType DeltaSharingRecipientType `tfsdk:"recipient_type"`

	Share *ShareInfo `tfsdk:"share"`

	Status PersonalizationRequestStatus `tfsdk:"status"`

	StatusMessage string `tfsdk:"status_message"`

	UpdatedAt int64 `tfsdk:"updated_at"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Id string `tfsdk:"id"`
}

type ProviderInfo struct {
	BusinessContactEmail string `tfsdk:"business_contact_email"`

	CompanyWebsiteLink string `tfsdk:"company_website_link"`

	DarkModeIconFileId string `tfsdk:"dark_mode_icon_file_id"`

	DarkModeIconFilePath string `tfsdk:"dark_mode_icon_file_path"`

	Description string `tfsdk:"description"`

	IconFileId string `tfsdk:"icon_file_id"`

	IconFilePath string `tfsdk:"icon_file_path"`

	Id string `tfsdk:"id"`
	// is_featured is accessible by consumers only
	IsFeatured bool `tfsdk:"is_featured"`

	Name string `tfsdk:"name"`

	PrivacyPolicyLink string `tfsdk:"privacy_policy_link"`
	// published_by is only applicable to data aggregators (e.g. Crux)
	PublishedBy string `tfsdk:"published_by"`

	SupportContactEmail string `tfsdk:"support_contact_email"`

	TermOfServiceLink string `tfsdk:"term_of_service_link"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegionInfo struct {
	Cloud string `tfsdk:"cloud"`

	Region string `tfsdk:"region"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RegionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {
	Id string `tfsdk:"-" url:"-"`
}

type RemoveExchangeForListingResponse struct {
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl string `tfsdk:"git_repo_url"`
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName string `tfsdk:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath string `tfsdk:"repo_path"`
}

// Search listings
type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `tfsdk:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `tfsdk:"-" url:"categories,omitempty"`

	IsAscending bool `tfsdk:"-" url:"is_ascending,omitempty"`

	IsFree bool `tfsdk:"-" url:"is_free,omitempty"`

	IsPrivateExchange bool `tfsdk:"-" url:"is_private_exchange,omitempty"`

	PageSize int `tfsdk:"-" url:"page_size,omitempty"`

	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `tfsdk:"-" url:"provider_ids,omitempty"`
	// Fuzzy matches query
	Query string `tfsdk:"-" url:"query"`

	SortBy SortBy `tfsdk:"-" url:"sort_by,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SearchListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchListingsResponse struct {
	Listings []Listing `tfsdk:"listings"`

	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SearchListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ShareInfo struct {
	Name string `tfsdk:"name"`

	Type ListingShareType `tfsdk:"type"`
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType string `tfsdk:"data_object_type"`
	// Name of the shared object
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	BearerToken string `tfsdk:"bearerToken"`

	Endpoint string `tfsdk:"endpoint"`

	ExpirationTime string `tfsdk:"expirationTime"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion int `tfsdk:"shareCredentialsVersion"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `tfsdk:"activation_url"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of Recipient Token creator.
	CreatedBy string `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `tfsdk:"expiration_time"`
	// Unique id of the Recipient Token.
	Id string `tfsdk:"id"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of Recipient Token updater.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExchangeFilterRequest struct {
	Filter ExchangeFilter `tfsdk:"filter"`

	Id string `tfsdk:"-" url:"-"`
}

type UpdateExchangeFilterResponse struct {
	Filter *ExchangeFilter `tfsdk:"filter"`
}

type UpdateExchangeRequest struct {
	Exchange Exchange `tfsdk:"exchange"`

	Id string `tfsdk:"-" url:"-"`
}

type UpdateExchangeResponse struct {
	Exchange *Exchange `tfsdk:"exchange"`
}

type UpdateInstallationRequest struct {
	Installation InstallationDetail `tfsdk:"installation"`

	InstallationId string `tfsdk:"-" url:"-"`

	ListingId string `tfsdk:"-" url:"-"`

	RotateToken bool `tfsdk:"rotate_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateInstallationResponse struct {
	Installation *InstallationDetail `tfsdk:"installation"`
}

type UpdateListingRequest struct {
	Id string `tfsdk:"-" url:"-"`

	Listing Listing `tfsdk:"listing"`
}

type UpdateListingResponse struct {
	Listing *Listing `tfsdk:"listing"`
}

type UpdatePersonalizationRequestRequest struct {
	ListingId string `tfsdk:"-" url:"-"`

	Reason string `tfsdk:"reason"`

	RequestId string `tfsdk:"-" url:"-"`

	Share *ShareInfo `tfsdk:"share"`

	Status PersonalizationRequestStatus `tfsdk:"status"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdatePersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePersonalizationRequestResponse struct {
	Request *PersonalizationRequest `tfsdk:"request"`
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id string `tfsdk:"-" url:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version int64 `tfsdk:"version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateProviderAnalyticsDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId string `tfsdk:"dashboard_id"`
	// id & version should be the same as the request
	Id string `tfsdk:"id"`

	Version int64 `tfsdk:"version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderRequest struct {
	Id string `tfsdk:"-" url:"-"`

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

	FilterValue string `tfsdk:"filterValue"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *VisibilityFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VisibilityFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
