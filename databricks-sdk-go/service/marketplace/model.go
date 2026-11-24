// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddExchangeForListingRequest struct {
	ExchangeId string `json:"exchange_id"`

	ListingId string `json:"listing_id"`
}

func (s *AddExchangeForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AddExchangeForListingResponse struct {
	ExchangeForListing *ExchangeListing `json:"exchange_for_listing,omitempty"`
}

func (s *AddExchangeForListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AssetType string

const AssetTypeAssetTypeAgentSkill AssetType = `ASSET_TYPE_AGENT_SKILL`

const AssetTypeAssetTypeApp AssetType = `ASSET_TYPE_APP`

const AssetTypeAssetTypeDataTable AssetType = `ASSET_TYPE_DATA_TABLE`

const AssetTypeAssetTypeGitRepo AssetType = `ASSET_TYPE_GIT_REPO`

const AssetTypeAssetTypeMcp AssetType = `ASSET_TYPE_MCP`

const AssetTypeAssetTypeMedia AssetType = `ASSET_TYPE_MEDIA`

const AssetTypeAssetTypeMetricView AssetType = `ASSET_TYPE_METRIC_VIEW`

const AssetTypeAssetTypeModel AssetType = `ASSET_TYPE_MODEL`

const AssetTypeAssetTypeNotebook AssetType = `ASSET_TYPE_NOTEBOOK`

const AssetTypeAssetTypePartnerIntegration AssetType = `ASSET_TYPE_PARTNER_INTEGRATION`

// String representation for [fmt.Print]
func (f *AssetType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AssetType) Set(v string) error {
	switch v {
	case `ASSET_TYPE_AGENT_SKILL`, `ASSET_TYPE_APP`, `ASSET_TYPE_DATA_TABLE`, `ASSET_TYPE_GIT_REPO`, `ASSET_TYPE_MCP`, `ASSET_TYPE_MEDIA`, `ASSET_TYPE_METRIC_VIEW`, `ASSET_TYPE_MODEL`, `ASSET_TYPE_NOTEBOOK`, `ASSET_TYPE_PARTNER_INTEGRATION`:
		*f = AssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSET_TYPE_AGENT_SKILL", "ASSET_TYPE_APP", "ASSET_TYPE_DATA_TABLE", "ASSET_TYPE_GIT_REPO", "ASSET_TYPE_MCP", "ASSET_TYPE_MEDIA", "ASSET_TYPE_METRIC_VIEW", "ASSET_TYPE_MODEL", "ASSET_TYPE_NOTEBOOK", "ASSET_TYPE_PARTNER_INTEGRATION"`, v)
	}
}

// Values returns all possible values for AssetType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AssetType) Values() []AssetType {
	return []AssetType{
		AssetTypeAssetTypeAgentSkill,
		AssetTypeAssetTypeApp,
		AssetTypeAssetTypeDataTable,
		AssetTypeAssetTypeGitRepo,
		AssetTypeAssetTypeMcp,
		AssetTypeAssetTypeMedia,
		AssetTypeAssetTypeMetricView,
		AssetTypeAssetTypeModel,
		AssetTypeAssetTypeNotebook,
		AssetTypeAssetTypePartnerIntegration,
	}
}

// Type always returns AssetType to satisfy [pflag.Value] interface
func (f *AssetType) Type() string {
	return "AssetType"
}

type BatchGetListingsRequest struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

func (s *BatchGetListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BatchGetListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`
}

func (s *BatchGetListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BatchGetProvidersRequest struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

func (s *BatchGetProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BatchGetProvidersResponse struct {
	Providers []ProviderInfo `json:"providers,omitempty"`
}

func (s *BatchGetProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for Category.
//
// There is no guarantee on the order of the values in the slice.
func (f *Category) Values() []Category {
	return []Category{
		CategoryAdvertisingAndMarketing,
		CategoryClimateAndEnvironment,
		CategoryCommerce,
		CategoryDemographics,
		CategoryEconomics,
		CategoryEducation,
		CategoryEnergy,
		CategoryFinancial,
		CategoryGaming,
		CategoryGeospatial,
		CategoryHealth,
		CategoryLookupTables,
		CategoryManufacturing,
		CategoryMedia,
		CategoryOther,
		CategoryPublicSector,
		CategoryRetail,
		CategoryScienceAndResearch,
		CategorySecurity,
		CategorySports,
		CategoryTransportationAndLogistics,
		CategoryTravelAndTourism,
	}
}

// Type always returns Category to satisfy [pflag.Value] interface
func (f *Category) Type() string {
	return "Category"
}

type ConsumerTerms struct {
	Version string `json:"version"`
}

func (s *ConsumerTerms) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	// Company of the consumer contact.
	Company string `json:"company,omitempty"`
	// Email address of the consumer contact.
	Email string `json:"email,omitempty"`
	// First name of the consumer contact.
	FirstName string `json:"first_name,omitempty"`
	// Last name of the consumer contact.
	LastName string `json:"last_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for Cost.
//
// There is no guarantee on the order of the values in the slice.
func (f *Cost) Values() []Cost {
	return []Cost{
		CostFree,
		CostPaid,
	}
}

// Type always returns Cost to satisfy [pflag.Value] interface
func (f *Cost) Type() string {
	return "Cost"
}

type CreateExchangeFilterRequest struct {
	Filter ExchangeFilter `json:"filter"`
}

func (s *CreateExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExchangeFilterResponse struct {
	FilterId string `json:"filter_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExchangeRequest struct {
	Exchange Exchange `json:"exchange"`
}

func (s *CreateExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExchangeResponse struct {
	ExchangeId string `json:"exchange_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileRequest struct {
	// Name shown to consumers for the file, for example a notebook title. At
	// most 255 characters.
	DisplayName string `json:"display_name,omitempty"`
	// Provider profile or listing to attach the file to.
	FileParent FileParent `json:"file_parent"`
	// Purpose of the file. PROVIDER_ICON requires a PROVIDER parent;
	// EMBEDDED_NOTEBOOK requires a LISTING parent.
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type"`
	// MIME type of the file content: text/html for notebooks; image/png,
	// image/jpeg or image/svg+xml for icons.
	MimeType string `json:"mime_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileResponse struct {
	// Metadata of the created file.
	FileInfo *FileInfo `json:"file_info,omitempty"`
	// Short-lived URL for uploading the file content with an HTTP PUT request.
	UploadUrl string `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms *ConsumerTerms `json:"accepted_consumer_terms,omitempty"`

	CatalogName string `json:"catalog_name,omitempty"`

	ListingId string `json:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	// for git repo installations
	RepoDetail *RepoInstallation `json:"repo_detail,omitempty"`

	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateListingRequest struct {
	// Listing to create.
	Listing Listing `json:"listing"`
}

func (s *CreateListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateListingResponse struct {
	// ID of the created listing.
	ListingId string `json:"listing_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms ConsumerTerms `json:"accepted_consumer_terms"`

	Comment string `json:"comment,omitempty"`

	Company string `json:"company,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	IntendedUse string `json:"intended_use"`

	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	LastName string `json:"last_name,omitempty"`

	ListingId string `json:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePersonalizationRequestResponse struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateProviderRequest struct {
	// Provider profile to create.
	Provider ProviderInfo `json:"provider"`
}

func (s *CreateProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateProviderResponse struct {
	// ID of the created provider profile.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for DataRefresh.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataRefresh) Values() []DataRefresh {
	return []DataRefresh{
		DataRefreshDaily,
		DataRefreshHourly,
		DataRefreshMinute,
		DataRefreshMonthly,
		DataRefreshNone,
		DataRefreshQuarterly,
		DataRefreshSecond,
		DataRefreshWeekly,
		DataRefreshYearly,
	}
}

// Type always returns DataRefresh to satisfy [pflag.Value] interface
func (f *DataRefresh) Type() string {
	return "DataRefresh"
}

type DataRefreshInfo struct {
	// Number of time units. Must be greater than zero.
	Interval int64 `json:"interval"`
	// Time unit that the interval is expressed in.
	Unit DataRefresh `json:"unit"`
}

func (s *DataRefreshInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExchangeFilterRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *DeleteExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExchangeRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *DeleteExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteFileRequest struct {
	// ID of the file to delete.
	FileId string `json:"-" url:"-"`
}

func (s *DeleteFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteInstallationRequest struct {
	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`
}

func (s *DeleteInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteListingRequest struct {
	// ID of the listing to delete.
	Id string `json:"-" url:"-"`
}

func (s *DeleteListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteProviderRequest struct {
	// ID of the provider profile to delete.
	Id string `json:"-" url:"-"`
}

func (s *DeleteProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for DeltaSharingRecipientType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeltaSharingRecipientType) Values() []DeltaSharingRecipientType {
	return []DeltaSharingRecipientType{
		DeltaSharingRecipientTypeDeltaSharingRecipientTypeDatabricks,
		DeltaSharingRecipientTypeDeltaSharingRecipientTypeOpen,
	}
}

// Type always returns DeltaSharingRecipientType to satisfy [pflag.Value] interface
func (f *DeltaSharingRecipientType) Type() string {
	return "DeltaSharingRecipientType"
}

type Exchange struct {
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	Filters []ExchangeFilter `json:"filters,omitempty"`

	Id string `json:"id,omitempty"`

	LinkedListings []ExchangeListing `json:"linked_listings,omitempty"`

	Name string `json:"name"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Exchange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Exchange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExchangeFilter struct {
	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	ExchangeId string `json:"exchange_id"`

	FilterType ExchangeFilterType `json:"filter_type"`

	FilterValue string `json:"filter_value"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for ExchangeFilterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExchangeFilterType) Values() []ExchangeFilterType {
	return []ExchangeFilterType{
		ExchangeFilterTypeGlobalMetastoreId,
	}
}

// Type always returns ExchangeFilterType to satisfy [pflag.Value] interface
func (f *ExchangeFilterType) Type() string {
	return "ExchangeFilterType"
}

type ExchangeListing struct {
	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	ExchangeId string `json:"exchange_id,omitempty"`

	ExchangeName string `json:"exchange_name,omitempty"`

	Id string `json:"id,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExchangeListing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeListing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileInfo struct {
	// Time the file was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName string `json:"display_name,omitempty"`
	// Short-lived URL for downloading the file content.
	DownloadLink string `json:"download_link,omitempty"`
	// Provider or listing that the file belongs to.
	FileParent *FileParent `json:"file_parent,omitempty"`
	// ID of the file.
	Id string `json:"id,omitempty"`
	// Purpose of the file, such as a provider icon or a notebook attached to a
	// listing.
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type,omitempty"`
	// MIME type of the file content, for example text/html for notebooks or
	// image/png for icons.
	MimeType string `json:"mime_type,omitempty"`
	// Processing status of the file. New files start in FILE_STATUS_STAGING and
	// become FILE_STATUS_PUBLISHED after Databricks has sanitized them.
	Status FileStatus `json:"status,omitempty"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage string `json:"status_message,omitempty"`
	// Time the file was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileParent struct {
	// Type of the parent: PROVIDER for provider icons, LISTING for notebooks
	// attached to a listing.
	FileParentType FileParentType `json:"file_parent_type,omitempty" url:"file_parent_type,omitempty"`
	// ID of the provider or listing that the file belongs to.
	ParentId string `json:"parent_id,omitempty" url:"parent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileParent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileParent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileParentType string

const FileParentTypeListing FileParentType = `LISTING`

const FileParentTypeListingResource FileParentType = `LISTING_RESOURCE`

const FileParentTypeProvider FileParentType = `PROVIDER`

// String representation for [fmt.Print]
func (f *FileParentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileParentType) Set(v string) error {
	switch v {
	case `LISTING`, `LISTING_RESOURCE`, `PROVIDER`:
		*f = FileParentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING", "LISTING_RESOURCE", "PROVIDER"`, v)
	}
}

// Values returns all possible values for FileParentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileParentType) Values() []FileParentType {
	return []FileParentType{
		FileParentTypeListing,
		FileParentTypeListingResource,
		FileParentTypeProvider,
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

// Values returns all possible values for FileStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileStatus) Values() []FileStatus {
	return []FileStatus{
		FileStatusFileStatusPublished,
		FileStatusFileStatusSanitizationFailed,
		FileStatusFileStatusSanitizing,
		FileStatusFileStatusStaging,
	}
}

// Type always returns FileStatus to satisfy [pflag.Value] interface
func (f *FileStatus) Type() string {
	return "FileStatus"
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

// Values returns all possible values for FulfillmentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FulfillmentType) Values() []FulfillmentType {
	return []FulfillmentType{
		FulfillmentTypeInstall,
		FulfillmentTypeRequestAccess,
	}
}

// Type always returns FulfillmentType to satisfy [pflag.Value] interface
func (f *FulfillmentType) Type() string {
	return "FulfillmentType"
}

type GetExchangeRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *GetExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExchangeResponse struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

func (s *GetExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFileRequest struct {
	// ID of the file.
	FileId string `json:"-" url:"-"`
}

func (s *GetFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFileResponse struct {
	// Metadata of the file.
	FileInfo *FileInfo `json:"file_info,omitempty"`
}

func (s *GetFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// Latest available version of the dashboard template.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLatestVersionProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLatestVersionProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingContentMetadataRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingContentMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingContentMetadataResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	SharedDataObjects []SharedDataObject `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingContentMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *GetListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetListingResponse struct {
	// The requested listing.
	Listing *Listing `json:"listing,omitempty"`
}

func (s *GetListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetListingsRequest struct {
	// Maximum number of results to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token from the next_page_token field of a previous response. Omit to get
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingsResponse struct {
	// Listings on this page.
	Listings []Listing `json:"listings,omitempty"`
	// Token for the next page. Not set when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPersonalizationRequestRequest struct {
	ListingId string `json:"-" url:"-"`
}

func (s *GetPersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`
}

func (s *GetPersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProviderRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *GetProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProviderResponse struct {
	// The requested provider profile.
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func (s *GetProviderResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Installation struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func (s *Installation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type InstallationDetail struct {
	CatalogName string `json:"catalog_name,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`

	Id string `json:"id,omitempty"`

	InstalledOn int64 `json:"installed_on,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	RepoName string `json:"repo_name,omitempty"`

	RepoPath string `json:"repo_path,omitempty"`

	ShareName string `json:"share_name,omitempty"`

	Status InstallationStatus `json:"status,omitempty"`

	TokenDetail *TokenDetail `json:"token_detail,omitempty"`

	Tokens []TokenInfo `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for InstallationStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstallationStatus) Values() []InstallationStatus {
	return []InstallationStatus{
		InstallationStatusFailed,
		InstallationStatusInstalled,
	}
}

// Type always returns InstallationStatus to satisfy [pflag.Value] interface
func (f *InstallationStatus) Type() string {
	return "InstallationStatus"
}

type ListAllInstallationsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllInstallationsResponse struct {
	Installations []InstallationDetail `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllPersonalizationRequestsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllPersonalizationRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllPersonalizationRequestsResponse struct {
	// Token for the next page. Not set when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Personalization requests on this page.
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllPersonalizationRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListConsumerProvidersRequest struct {
	IsFeatured bool `json:"-" url:"is_featured,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListConsumerProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConsumerProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangeFiltersRequest struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangeFiltersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangeFiltersResponse struct {
	Filters []ExchangeFilter `json:"filters,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangeFiltersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesForListingRequest struct {
	ListingId string `json:"-" url:"listing_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesForListingResponse struct {
	ExchangeListing []ExchangeListing `json:"exchange_listing,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesForListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesResponse struct {
	Exchanges []Exchange `json:"exchanges,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFilesRequest struct {
	// Provider profile or listing whose files to list.
	FileParent FileParent `json:"-" url:"file_parent"`
	// Maximum number of results to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token from the next_page_token field of a previous response. Omit to get
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFilesResponse struct {
	// Files on this page.
	FileInfos []FileInfo `json:"file_infos,omitempty"`
	// Token for the next page. Not set when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFulfillmentsRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFulfillmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFulfillmentsResponse struct {
	Fulfillments []ListingFulfillment `json:"fulfillments,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFulfillmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListInstallationsRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListInstallationsResponse struct {
	Installations []InstallationDetail `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsForExchangeRequest struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsForExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsForExchangeResponse struct {
	ExchangeListings []ExchangeListing `json:"exchange_listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsForExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `json:"-" url:"categories,omitempty"`
	// Filters each listing based on if it is free.
	IsFree bool `json:"-" url:"is_free,omitempty"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange bool `json:"-" url:"is_private_exchange,omitempty"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick bool `json:"-" url:"is_staff_pick,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" url:"provider_ids,omitempty"`
	// Matches listings with this tag
	Tags *ListingTag `json:"-" url:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProviderAnalyticsDashboardResponse struct {
	// ID of the AI/BI dashboard in the workspace. Use it to open the dashboard.
	DashboardId string `json:"dashboard_id"`
	// Marketplace-specific ID of the dashboard record. Use it as the `id` when
	// updating the dashboard.
	Id string `json:"id"`
	// Version of the dashboard template that the dashboard was created or last
	// updated from.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersRequest struct {
	// Maximum number of results to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token from the next_page_token field of a previous response. Omit to get
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersResponse struct {
	// Token for the next page. Not set when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Provider profiles on this page.
	Providers []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Marketplace listing: a data product that a provider offers to consumers.
// The summary holds the listing's identity, status and sharing configuration,
// and the detail holds the descriptive content that consumers see on the
// listing page.
type Listing struct {
	// Descriptive content shown to consumers on the listing page, such as the
	// description, links and data attributes.
	Detail *ListingDetail `json:"detail,omitempty"`
	// ID of the listing. Assigned by the server; not allowed in create
	// requests.
	Id string `json:"id,omitempty"`
	// Identity, status and sharing configuration of the listing.
	Summary ListingSummary `json:"summary"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Listing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Listing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingDetail struct {
	// Types of assets included in the listing. At least one is required unless
	// the listing is a DRAFT. Duplicate values are ignored.
	Assets []AssetType `json:"assets,omitempty"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd int64 `json:"collection_date_end,omitempty"`
	// The starting date timestamp for when the data spans
	CollectionDateStart int64 `json:"collection_date_start,omitempty"`
	// Smallest unit of time in the dataset
	CollectionGranularity *DataRefreshInfo `json:"collection_granularity,omitempty"`
	// Whether the dataset is free or paid
	Cost Cost `json:"cost,omitempty"`
	// Where/how the data is sourced
	DataSource string `json:"data_source,omitempty"`
	// Detailed description of the data product. Supports basic Markdown
	// formatting such as bold, italics and lists.
	Description string `json:"description,omitempty"`
	// URL of documentation that helps consumers use the data product.
	DocumentationLink string `json:"documentation_link,omitempty"`
	// Sample notebooks attached to the listing. Create each notebook with
	// createFile, using the EMBEDDED_NOTEBOOK file type and the listing as the
	// parent, and reference it here.
	EmbeddedNotebookFileInfos []FileInfo `json:"embedded_notebook_file_infos,omitempty"`

	FileIds []string `json:"file_ids,omitempty"`
	// Which geo region the listing data is collected from
	GeographicalCoverage string `json:"geographical_coverage,omitempty"`
	// License of the data asset. Required for listings with model-based assets.
	License string `json:"license,omitempty"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid
	PricingModel string `json:"pricing_model,omitempty"`
	// URL of the privacy policy for the listing. Defaults to the provider
	// profile's privacy policy link when empty.
	PrivacyPolicyLink string `json:"privacy_policy_link,omitempty"`
	// size of the dataset in GB
	Size float64 `json:"size,omitempty"`
	// URL where consumers can get support for the listing.
	SupportLink string `json:"support_link,omitempty"`
	// Tags that annotate the listing with fixed, low-cardinality values, such
	// as supported languages or tasks.
	Tags []ListingTag `json:"tags,omitempty"`
	// URL of the terms of service for the listing. Defaults to the provider
	// profile's terms of service link when empty.
	TermsOfService string `json:"terms_of_service,omitempty"`
	// How often data is updated
	UpdateFrequency *DataRefreshInfo `json:"update_frequency,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListingDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingFulfillment struct {
	FulfillmentType FulfillmentType `json:"fulfillment_type,omitempty"`

	ListingId string `json:"listing_id"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	RepoInfo *RepoInfo `json:"repo_info,omitempty"`

	ShareInfo *ShareInfo `json:"share_info,omitempty"`
}

func (s *ListingFulfillment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListingSetting struct {
	// PUBLIC listings are visible to all consumers in the Marketplace. PRIVATE
	// listings are visible only to members of the private exchanges that they
	// are shared in.
	Visibility Visibility `json:"visibility,omitempty"`
}

func (s *ListingSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for ListingShareType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingShareType) Values() []ListingShareType {
	return []ListingShareType{
		ListingShareTypeFull,
		ListingShareTypeSample,
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

// Values returns all possible values for ListingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingStatus) Values() []ListingStatus {
	return []ListingStatus{
		ListingStatusDraft,
		ListingStatusPending,
		ListingStatusPublished,
		ListingStatusSuspended,
	}
}

// Type always returns ListingStatus to satisfy [pflag.Value] interface
func (f *ListingStatus) Type() string {
	return "ListingStatus"
}

type ListingSummary struct {
	// Categories that consumers can use to filter listings. They also appear as
	// tags on the listing tile and detail page.
	Categories []Category `json:"categories,omitempty"`
	// Time the listing was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Name of the user who created the listing.
	CreatedBy string `json:"created_by,omitempty"`
	// ID of the user who created the listing.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// IDs of the private exchanges that the listing is shared in. A listing can
	// be linked to at most 10 exchanges.
	ExchangeIds []string `json:"exchange_ids,omitempty"`
	// Git repository that holds the listing's assets, used instead of a share
	// for notebook and Git repo listings. A listing cannot have both a share
	// and a Git repo.
	GitRepo *RepoInfo `json:"git_repo,omitempty"`
	// How consumers get access to the listing. STANDARD listings are instantly
	// available; PERSONALIZED listings require the provider to approve each
	// consumer request. Cannot be changed after the listing is created.
	ListingType ListingType `json:"listingType"`
	// Display name of the listing. Must start with a letter or digit and be at
	// most 100 characters.
	Name string `json:"name"`
	// ID of the provider profile that owns the listing.
	ProviderId string `json:"provider_id,omitempty"`
	// Cloud and region of the provider's Unity Catalog metastore. Not allowed
	// in create requests.
	ProviderRegion *RegionInfo `json:"provider_region,omitempty"`

	PublishedAt int64 `json:"published_at,omitempty"`

	PublishedBy string `json:"published_by,omitempty"`
	// Listing settings, such as whether the listing is visible in the public
	// Marketplace or only in private exchanges.
	Setting *ListingSetting `json:"setting,omitempty"`
	// Delta Sharing share that holds the listing's data assets. Required to
	// publish a STANDARD listing, except for app and MCP server listings;
	// notebook and Git repo listings can use git_repo instead. The caller must
	// have access to the share.
	Share *ShareInfo `json:"share,omitempty"`
	// Publication status of the listing. Only PUBLISHED listings are visible to
	// consumers; create a listing as DRAFT and set the status to PUBLISHED when
	// it is ready.
	Status ListingStatus `json:"status,omitempty"`
	// One-line subtitle shown with the listing name.
	Subtitle string `json:"subtitle,omitempty"`
	// Time the listing was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Name of the user who last updated the listing.
	UpdatedBy string `json:"updated_by,omitempty"`
	// ID of the user who last updated the listing.
	UpdatedById int64 `json:"updated_by_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListingSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingTag struct {
	// Tag name (enum)
	TagName ListingTagType `json:"tag_name,omitempty" url:"tag_name,omitempty"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []string `json:"tag_values,omitempty" url:"tag_values,omitempty"`
}

func (s *ListingTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListingTagType string

const ListingTagTypeListingTagTypeLanguage ListingTagType = `LISTING_TAG_TYPE_LANGUAGE`

const ListingTagTypeListingTagTypeTask ListingTagType = `LISTING_TAG_TYPE_TASK`

// String representation for [fmt.Print]
func (f *ListingTagType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingTagType) Set(v string) error {
	switch v {
	case `LISTING_TAG_TYPE_LANGUAGE`, `LISTING_TAG_TYPE_TASK`:
		*f = ListingTagType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING_TAG_TYPE_LANGUAGE", "LISTING_TAG_TYPE_TASK"`, v)
	}
}

// Values returns all possible values for ListingTagType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingTagType) Values() []ListingTagType {
	return []ListingTagType{
		ListingTagTypeListingTagTypeLanguage,
		ListingTagTypeListingTagTypeTask,
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

// Values returns all possible values for ListingType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingType) Values() []ListingType {
	return []ListingType{
		ListingTypePersonalized,
		ListingTypeStandard,
	}
}

// Type always returns ListingType to satisfy [pflag.Value] interface
func (f *ListingType) Type() string {
	return "ListingType"
}

type MarketplaceFileType string

const MarketplaceFileTypeApp MarketplaceFileType = `APP`

const MarketplaceFileTypeEmbeddedNotebook MarketplaceFileType = `EMBEDDED_NOTEBOOK`

const MarketplaceFileTypeProviderIcon MarketplaceFileType = `PROVIDER_ICON`

// String representation for [fmt.Print]
func (f *MarketplaceFileType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MarketplaceFileType) Set(v string) error {
	switch v {
	case `APP`, `EMBEDDED_NOTEBOOK`, `PROVIDER_ICON`:
		*f = MarketplaceFileType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APP", "EMBEDDED_NOTEBOOK", "PROVIDER_ICON"`, v)
	}
}

// Values returns all possible values for MarketplaceFileType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MarketplaceFileType) Values() []MarketplaceFileType {
	return []MarketplaceFileType{
		MarketplaceFileTypeApp,
		MarketplaceFileTypeEmbeddedNotebook,
		MarketplaceFileTypeProviderIcon,
	}
}

// Type always returns MarketplaceFileType to satisfy [pflag.Value] interface
func (f *MarketplaceFileType) Type() string {
	return "MarketplaceFileType"
}

type PersonalizationRequest struct {
	// Comment that the consumer entered with the request.
	Comment string `json:"comment,omitempty"`
	// Cloud and region of the consumer's Unity Catalog metastore.
	ConsumerRegion RegionInfo `json:"consumer_region"`
	// Contact details that the consumer entered with the request.
	ContactInfo *ContactInfo `json:"contact_info,omitempty"`
	// Time the request was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// ID of the personalization request.
	Id string `json:"id,omitempty"`
	// How the consumer intends to use the data product, as entered with the
	// request.
	IntendedUse string `json:"intended_use,omitempty"`

	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`
	// ID of the listing that the request was made for.
	ListingId string `json:"listing_id,omitempty"`
	// Name of the listing that the request was made for.
	ListingName string `json:"listing_name,omitempty"`
	// ID of the consumer's Unity Catalog metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// ID of the provider profile that owns the listing.
	ProviderId string `json:"provider_id,omitempty"`
	// How the consumer receives the shared data:
	// DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS for a consumer on a Databricks
	// workspace, DELTA_SHARING_RECIPIENT_TYPE_OPEN for a consumer on another
	// platform who uses a credential file.
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	// Share information is required for data listings but should be
	// empty/ignored for non-data listings (MCP and App).
	Share *ShareInfo `json:"share,omitempty"`
	// Current status of the request.
	Status PersonalizationRequestStatus `json:"status,omitempty"`
	// Reason that the provider gave in the most recent status update.
	StatusMessage string `json:"status_message,omitempty"`
	// Time the request was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

// Values returns all possible values for PersonalizationRequestStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalizationRequestStatus) Values() []PersonalizationRequestStatus {
	return []PersonalizationRequestStatus{
		PersonalizationRequestStatusDenied,
		PersonalizationRequestStatusFulfilled,
		PersonalizationRequestStatusNew,
		PersonalizationRequestStatusRequestPending,
	}
}

// Type always returns PersonalizationRequestStatus to satisfy [pflag.Value] interface
func (f *PersonalizationRequestStatus) Type() string {
	return "PersonalizationRequestStatus"
}

type ProviderAnalyticsDashboard struct {
	// Marketplace-specific ID of the dashboard record, distinct from the AI/BI
	// dashboard ID.
	Id string `json:"id"`
}

func (s *ProviderAnalyticsDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ProviderInfo struct {
	// Email address that Databricks uses to contact the provider. Not shown to
	// consumers.
	BusinessContactEmail string `json:"business_contact_email"`
	// URL of the provider's website. Shown on all of the provider's listings.
	CompanyWebsiteLink string `json:"company_website_link,omitempty"`
	// ID of the provider icon file used in dark mode, created with createFile
	// using the PROVIDER_ICON file type.
	DarkModeIconFileId string `json:"dark_mode_icon_file_id,omitempty"`
	// URL for downloading the dark mode provider icon. Populated by the server
	// from dark_mode_icon_file_id.
	DarkModeIconFilePath string `json:"dark_mode_icon_file_path,omitempty"`
	// Description of the provider organization, shown on the provider profile
	// and on all of its listings.
	Description string `json:"description,omitempty"`
	// ID of the provider icon file, created with createFile using the
	// PROVIDER_ICON file type.
	IconFileId string `json:"icon_file_id,omitempty"`
	// URL for downloading the provider icon. Populated by the server from
	// icon_file_id.
	IconFilePath string `json:"icon_file_path,omitempty"`
	// ID of the provider profile. Assigned by the server; not allowed in create
	// or update requests.
	Id string `json:"id,omitempty"`
	// Whether Databricks features the provider in the Marketplace. Set by
	// Databricks; not allowed in create or update requests.
	IsFeatured bool `json:"is_featured,omitempty"`
	// Display name of the provider. Consumers can filter listings by provider
	// name.
	Name string `json:"name"`
	// URL of the provider's privacy policy. Shown on all of the provider's
	// listings unless a listing sets its own privacy policy.
	PrivacyPolicyLink string `json:"privacy_policy_link"`
	// published_by is only applicable to data aggregators (e.g. Crux)
	PublishedBy string `json:"published_by,omitempty"`
	// Email address that consumers can use to request support. Shown on all of
	// the provider's listings.
	SupportContactEmail string `json:"support_contact_email,omitempty"`
	// URL of the provider's terms of service. Shown on all of the provider's
	// listings unless a listing sets its own terms of service.
	TermOfServiceLink string `json:"term_of_service_link"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegionInfo struct {
	// Cloud provider: aws, azure or gcp.
	Cloud string `json:"cloud,omitempty"`
	// Cloud region, for example us-west-2.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RemoveExchangeForListingRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *RemoveExchangeForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl string `json:"git_repo_url"`
}

func (s *RepoInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName string `json:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath string `json:"repo_path"`
}

func (s *RepoInstallation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `json:"-" url:"categories,omitempty"`

	IsFree bool `json:"-" url:"is_free,omitempty"`

	IsPrivateExchange bool `json:"-" url:"is_private_exchange,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" url:"provider_ids,omitempty"`
	// Fuzzy matches query
	Query string `json:"-" url:"query"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ShareInfo struct {
	// Name of the Delta Sharing share in the provider's Unity Catalog
	// metastore.
	Name string `json:"name"`

	Type ListingShareType `json:"type"`
}

func (s *ShareInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType string `json:"data_object_type,omitempty"`
	// Name of the shared object
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenDetail struct {
	BearerToken string `json:"bearerToken,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Recipient Token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique id of the Recipient Token.
	Id string `json:"id,omitempty"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of Recipient Token updater.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExchangeFilterRequest struct {
	Filter ExchangeFilter `json:"filter"`

	Id string `json:"-" url:"-"`
}

func (s *UpdateExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExchangeFilterResponse struct {
	Filter *ExchangeFilter `json:"filter,omitempty"`
}

func (s *UpdateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExchangeRequest struct {
	Exchange Exchange `json:"exchange"`

	Id string `json:"-" url:"-"`
}

func (s *UpdateExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExchangeResponse struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

func (s *UpdateExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateInstallationRequest struct {
	Installation InstallationDetail `json:"installation"`

	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`

	RotateToken bool `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateInstallationResponse struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func (s *UpdateInstallationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateListingRequest struct {
	// ID of the listing to update.
	Id string `json:"-" url:"-"`
	// Full definition of the listing, including the fields that are not
	// changing.
	Listing Listing `json:"listing"`
}

func (s *UpdateListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateListingResponse struct {
	// The updated listing.
	Listing *Listing `json:"listing,omitempty"`
}

func (s *UpdateListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdatePersonalizationRequestRequest struct {
	// ID of the listing that the request was made for.
	ListingId string `json:"-" url:"-"`
	// Message for the consumer that explains the status change, at most 255
	// characters.
	Reason string `json:"reason,omitempty"`
	// ID of the personalization request.
	RequestId string `json:"-" url:"-"`
	// Share to grant the consumer access to. Required when fulfilling a request
	// for a listing with data assets and not allowed with other statuses.
	// Ignored for app and MCP server listings.
	Share *ShareInfo `json:"share,omitempty"`
	// New status of the request. Allowed transitions: NEW or REQUEST_PENDING to
	// REQUEST_PENDING, FULFILLED or DENIED; FULFILLED to DENIED; DENIED to
	// REQUEST_PENDING.
	Status PersonalizationRequestStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdatePersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePersonalizationRequestResponse struct {
	// The updated personalization request.
	Request *PersonalizationRequest `json:"request,omitempty"`
}

func (s *UpdatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// Marketplace-specific ID of the dashboard record, as returned when the
	// dashboard was created or retrieved.
	Id string `json:"-" url:"-"`
	// Version of the dashboard template to update to. Must be newer than the
	// version that the dashboard currently uses; use the version returned by
	// getLatestVersionProviderAnalyticsDashboard.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateProviderAnalyticsDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// ID of the AI/BI dashboard that now contains the updated template.
	DashboardId string `json:"dashboard_id"`
	// Marketplace-specific ID of the dashboard record, the same as in the
	// request.
	Id string `json:"id"`
	// Version of the dashboard template that the dashboard now uses.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderRequest struct {
	// ID of the provider profile to update.
	Id string `json:"-" url:"-"`
	// Full definition of the provider profile, including the fields that are
	// not changing.
	Provider ProviderInfo `json:"provider"`
}

func (s *UpdateProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateProviderResponse struct {
	// The updated provider profile.
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func (s *UpdateProviderResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for Visibility.
//
// There is no guarantee on the order of the values in the slice.
func (f *Visibility) Values() []Visibility {
	return []Visibility{
		VisibilityPrivate,
		VisibilityPublic,
	}
}

// Type always returns Visibility to satisfy [pflag.Value] interface
func (f *Visibility) Type() string {
	return "Visibility"
}
