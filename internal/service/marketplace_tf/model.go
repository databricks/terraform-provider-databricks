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
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddExchangeForListingRequest struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:""`

	ListingId types.String `tfsdk:"listing_id" tf:""`
}

func (newState *AddExchangeForListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddExchangeForListingRequest) {
}

func (newState *AddExchangeForListingRequest) SyncEffectiveFieldsDuringRead(existingState AddExchangeForListingRequest) {
}

func (a AddExchangeForListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a AddExchangeForListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeId": types.StringType,
			"ListingId":  types.StringType,
		},
	}
}

type AddExchangeForListingResponse struct {
	ExchangeForListing types.List `tfsdk:"exchange_for_listing" tf:"optional,object"`
}

func (newState *AddExchangeForListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddExchangeForListingResponse) {
}

func (newState *AddExchangeForListingResponse) SyncEffectiveFieldsDuringRead(existingState AddExchangeForListingResponse) {
}

func (a AddExchangeForListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ExchangeForListing": reflect.TypeOf(ExchangeListing{}),
	}
}

func (a AddExchangeForListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeForListing": ExchangeListing{}.ToAttrType(ctx),
		},
	}
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest struct {
	Ids types.List `tfsdk:"-"`
}

func (newState *BatchGetListingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetListingsRequest) {
}

func (newState *BatchGetListingsRequest) SyncEffectiveFieldsDuringRead(existingState BatchGetListingsRequest) {
}

func (a BatchGetListingsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Ids": reflect.TypeOf(types.StringType),
	}
}

func (a BatchGetListingsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type BatchGetListingsResponse struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`
}

func (newState *BatchGetListingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetListingsResponse) {
}

func (newState *BatchGetListingsResponse) SyncEffectiveFieldsDuringRead(existingState BatchGetListingsResponse) {
}

func (a BatchGetListingsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listings": reflect.TypeOf(Listing{}),
	}
}

func (a BatchGetListingsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listings": basetypes.ListType{
				ElemType: Listing{}.ToAttrType(ctx),
			},
		},
	}
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest struct {
	Ids types.List `tfsdk:"-"`
}

func (newState *BatchGetProvidersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetProvidersRequest) {
}

func (newState *BatchGetProvidersRequest) SyncEffectiveFieldsDuringRead(existingState BatchGetProvidersRequest) {
}

func (a BatchGetProvidersRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Ids": reflect.TypeOf(types.StringType),
	}
}

func (a BatchGetProvidersRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type BatchGetProvidersResponse struct {
	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *BatchGetProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetProvidersResponse) {
}

func (newState *BatchGetProvidersResponse) SyncEffectiveFieldsDuringRead(existingState BatchGetProvidersResponse) {
}

func (a BatchGetProvidersResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Providers": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a BatchGetProvidersResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Providers": basetypes.ListType{
				ElemType: ProviderInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ConsumerTerms struct {
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *ConsumerTerms) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConsumerTerms) {
}

func (newState *ConsumerTerms) SyncEffectiveFieldsDuringRead(existingState ConsumerTerms) {
}

func (a ConsumerTerms) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ConsumerTerms) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Version": types.StringType,
		},
	}
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company types.String `tfsdk:"company" tf:"optional"`

	Email types.String `tfsdk:"email" tf:"optional"`

	FirstName types.String `tfsdk:"first_name" tf:"optional"`

	LastName types.String `tfsdk:"last_name" tf:"optional"`
}

func (newState *ContactInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContactInfo) {
}

func (newState *ContactInfo) SyncEffectiveFieldsDuringRead(existingState ContactInfo) {
}

func (a ContactInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ContactInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Company":   types.StringType,
			"Email":     types.StringType,
			"FirstName": types.StringType,
			"LastName":  types.StringType,
		},
	}
}

type CreateExchangeFilterRequest struct {
	Filter types.List `tfsdk:"filter" tf:"object"`
}

func (newState *CreateExchangeFilterRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeFilterRequest) {
}

func (newState *CreateExchangeFilterRequest) SyncEffectiveFieldsDuringRead(existingState CreateExchangeFilterRequest) {
}

func (a CreateExchangeFilterRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

func (a CreateExchangeFilterRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter": ExchangeFilter{}.ToAttrType(ctx),
		},
	}
}

type CreateExchangeFilterResponse struct {
	FilterId types.String `tfsdk:"filter_id" tf:"optional"`
}

func (newState *CreateExchangeFilterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeFilterResponse) {
}

func (newState *CreateExchangeFilterResponse) SyncEffectiveFieldsDuringRead(existingState CreateExchangeFilterResponse) {
}

func (a CreateExchangeFilterResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateExchangeFilterResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FilterId": types.StringType,
		},
	}
}

type CreateExchangeRequest struct {
	Exchange types.List `tfsdk:"exchange" tf:"object"`
}

func (newState *CreateExchangeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeRequest) {
}

func (newState *CreateExchangeRequest) SyncEffectiveFieldsDuringRead(existingState CreateExchangeRequest) {
}

func (a CreateExchangeRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Exchange": reflect.TypeOf(Exchange{}),
	}
}

func (a CreateExchangeRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Exchange": Exchange{}.ToAttrType(ctx),
		},
	}
}

type CreateExchangeResponse struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:"optional"`
}

func (newState *CreateExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeResponse) {
}

func (newState *CreateExchangeResponse) SyncEffectiveFieldsDuringRead(existingState CreateExchangeResponse) {
}

func (a CreateExchangeResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateExchangeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeId": types.StringType,
		},
	}
}

type CreateFileRequest struct {
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	FileParent types.List `tfsdk:"file_parent" tf:"object"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type" tf:""`

	MimeType types.String `tfsdk:"mime_type" tf:""`
}

func (newState *CreateFileRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFileRequest) {
}

func (newState *CreateFileRequest) SyncEffectiveFieldsDuringRead(existingState CreateFileRequest) {
}

func (a CreateFileRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileParent": reflect.TypeOf(FileParent{}),
	}
}

func (a CreateFileRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DisplayName":         types.StringType,
			"FileParent":          FileParent{}.ToAttrType(ctx),
			"MarketplaceFileType": types.StringType,
			"MimeType":            types.StringType,
		},
	}
}

type CreateFileResponse struct {
	FileInfo types.List `tfsdk:"file_info" tf:"optional,object"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url" tf:"optional"`
}

func (newState *CreateFileResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFileResponse) {
}

func (newState *CreateFileResponse) SyncEffectiveFieldsDuringRead(existingState CreateFileResponse) {
}

func (a CreateFileResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileInfo": reflect.TypeOf(FileInfo{}),
	}
}

func (a CreateFileResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileInfo":  FileInfo{}.ToAttrType(ctx),
			"UploadUrl": types.StringType,
		},
	}
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms types.List `tfsdk:"accepted_consumer_terms" tf:"optional,object"`

	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`
	// for git repo installations
	RepoDetail types.List `tfsdk:"repo_detail" tf:"optional,object"`

	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *CreateInstallationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateInstallationRequest) {
}

func (newState *CreateInstallationRequest) SyncEffectiveFieldsDuringRead(existingState CreateInstallationRequest) {
}

func (a CreateInstallationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AcceptedConsumerTerms": reflect.TypeOf(ConsumerTerms{}),
		"RepoDetail":            reflect.TypeOf(RepoInstallation{}),
	}
}

func (a CreateInstallationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AcceptedConsumerTerms": ConsumerTerms{}.ToAttrType(ctx),
			"CatalogName":           types.StringType,
			"ListingId":             types.StringType,
			"RecipientType":         types.StringType,
			"RepoDetail":            RepoInstallation{}.ToAttrType(ctx),
			"ShareName":             types.StringType,
		},
	}
}

type CreateListingRequest struct {
	Listing types.List `tfsdk:"listing" tf:"object"`
}

func (newState *CreateListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateListingRequest) {
}

func (newState *CreateListingRequest) SyncEffectiveFieldsDuringRead(existingState CreateListingRequest) {
}

func (a CreateListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listing": reflect.TypeOf(Listing{}),
	}
}

func (a CreateListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listing": Listing{}.ToAttrType(ctx),
		},
	}
}

type CreateListingResponse struct {
	ListingId types.String `tfsdk:"listing_id" tf:"optional"`
}

func (newState *CreateListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateListingResponse) {
}

func (newState *CreateListingResponse) SyncEffectiveFieldsDuringRead(existingState CreateListingResponse) {
}

func (a CreateListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
		},
	}
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms types.List `tfsdk:"accepted_consumer_terms" tf:"object"`

	Comment types.String `tfsdk:"comment" tf:"optional"`

	Company types.String `tfsdk:"company" tf:"optional"`

	FirstName types.String `tfsdk:"first_name" tf:"optional"`

	IntendedUse types.String `tfsdk:"intended_use" tf:""`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse" tf:"optional"`

	LastName types.String `tfsdk:"last_name" tf:"optional"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`
}

func (newState *CreatePersonalizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePersonalizationRequest) {
}

func (newState *CreatePersonalizationRequest) SyncEffectiveFieldsDuringRead(existingState CreatePersonalizationRequest) {
}

func (a CreatePersonalizationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AcceptedConsumerTerms": reflect.TypeOf(ConsumerTerms{}),
	}
}

func (a CreatePersonalizationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AcceptedConsumerTerms": ConsumerTerms{}.ToAttrType(ctx),
			"Comment":               types.StringType,
			"Company":               types.StringType,
			"FirstName":             types.StringType,
			"IntendedUse":           types.StringType,
			"IsFromLighthouse":      types.BoolType,
			"LastName":              types.StringType,
			"ListingId":             types.StringType,
			"RecipientType":         types.StringType,
		},
	}
}

type CreatePersonalizationRequestResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreatePersonalizationRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePersonalizationRequestResponse) {
}

func (newState *CreatePersonalizationRequestResponse) SyncEffectiveFieldsDuringRead(existingState CreatePersonalizationRequestResponse) {
}

func (a CreatePersonalizationRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreatePersonalizationRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type CreateProviderRequest struct {
	Provider types.List `tfsdk:"provider" tf:"object"`
}

func (newState *CreateProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProviderRequest) {
}

func (newState *CreateProviderRequest) SyncEffectiveFieldsDuringRead(existingState CreateProviderRequest) {
}

func (a CreateProviderRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Provider": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a CreateProviderRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Provider": ProviderInfo{}.ToAttrType(ctx),
		},
	}
}

type CreateProviderResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreateProviderResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProviderResponse) {
}

func (newState *CreateProviderResponse) SyncEffectiveFieldsDuringRead(existingState CreateProviderResponse) {
}

func (a CreateProviderResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateProviderResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DataRefreshInfo struct {
	Interval types.Int64 `tfsdk:"interval" tf:""`

	Unit types.String `tfsdk:"unit" tf:""`
}

func (newState *DataRefreshInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataRefreshInfo) {
}

func (newState *DataRefreshInfo) SyncEffectiveFieldsDuringRead(existingState DataRefreshInfo) {
}

func (a DataRefreshInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DataRefreshInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Interval": types.Int64Type,
			"Unit":     types.StringType,
		},
	}
}

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteExchangeFilterRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterRequest) {
}

func (newState *DeleteExchangeFilterRequest) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeFilterRequest) {
}

func (a DeleteExchangeFilterRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExchangeFilterRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteExchangeFilterResponse struct {
}

func (newState *DeleteExchangeFilterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterResponse) {
}

func (newState *DeleteExchangeFilterResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeFilterResponse) {
}

func (a DeleteExchangeFilterResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExchangeFilterResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an exchange
type DeleteExchangeRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteExchangeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeRequest) {
}

func (newState *DeleteExchangeRequest) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeRequest) {
}

func (a DeleteExchangeRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExchangeRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteExchangeResponse struct {
}

func (newState *DeleteExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeResponse) {
}

func (newState *DeleteExchangeResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeResponse) {
}

func (a DeleteExchangeResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExchangeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a file
type DeleteFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

func (newState *DeleteFileRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileRequest) {
}

func (newState *DeleteFileRequest) SyncEffectiveFieldsDuringRead(existingState DeleteFileRequest) {
}

func (a DeleteFileRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteFileRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileId": types.StringType,
		},
	}
}

type DeleteFileResponse struct {
}

func (newState *DeleteFileResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileResponse) {
}

func (newState *DeleteFileResponse) SyncEffectiveFieldsDuringRead(existingState DeleteFileResponse) {
}

func (a DeleteFileResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteFileResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Uninstall from a listing
type DeleteInstallationRequest struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
}

func (newState *DeleteInstallationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstallationRequest) {
}

func (newState *DeleteInstallationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteInstallationRequest) {
}

func (a DeleteInstallationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteInstallationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"InstallationId": types.StringType,
			"ListingId":      types.StringType,
		},
	}
}

type DeleteInstallationResponse struct {
}

func (newState *DeleteInstallationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstallationResponse) {
}

func (newState *DeleteInstallationResponse) SyncEffectiveFieldsDuringRead(existingState DeleteInstallationResponse) {
}

func (a DeleteInstallationResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteInstallationResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a listing
type DeleteListingRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteListingRequest) {
}

func (newState *DeleteListingRequest) SyncEffectiveFieldsDuringRead(existingState DeleteListingRequest) {
}

func (a DeleteListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteListingResponse struct {
}

func (newState *DeleteListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteListingResponse) {
}

func (newState *DeleteListingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteListingResponse) {
}

func (a DeleteListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete provider
type DeleteProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderRequest) {
}

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringRead(existingState DeleteProviderRequest) {
}

func (a DeleteProviderRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteProviderRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteProviderResponse struct {
}

func (newState *DeleteProviderResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderResponse) {
}

func (newState *DeleteProviderResponse) SyncEffectiveFieldsDuringRead(existingState DeleteProviderResponse) {
}

func (a DeleteProviderResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteProviderResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Exchange struct {
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	Filters types.List `tfsdk:"filters" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	LinkedListings types.List `tfsdk:"linked_listings" tf:"optional"`

	Name types.String `tfsdk:"name" tf:""`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`

	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *Exchange) SyncEffectiveFieldsDuringCreateOrUpdate(plan Exchange) {
}

func (newState *Exchange) SyncEffectiveFieldsDuringRead(existingState Exchange) {
}

func (a Exchange) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Filters":        reflect.TypeOf(ExchangeFilter{}),
		"LinkedListings": reflect.TypeOf(ExchangeListing{}),
	}
}

func (a Exchange) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment":   types.StringType,
			"CreatedAt": types.Int64Type,
			"CreatedBy": types.StringType,
			"Filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToAttrType(ctx),
			},
			"Id": types.StringType,
			"LinkedListings": basetypes.ListType{
				ElemType: ExchangeListing{}.ToAttrType(ctx),
			},
			"Name":      types.StringType,
			"UpdatedAt": types.Int64Type,
			"UpdatedBy": types.StringType,
		},
	}
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

func (newState *ExchangeFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeFilter) {
}

func (newState *ExchangeFilter) SyncEffectiveFieldsDuringRead(existingState ExchangeFilter) {
}

func (a ExchangeFilter) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExchangeFilter) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreatedAt":   types.Int64Type,
			"CreatedBy":   types.StringType,
			"ExchangeId":  types.StringType,
			"FilterType":  types.StringType,
			"FilterValue": types.StringType,
			"Id":          types.StringType,
			"Name":        types.StringType,
			"UpdatedAt":   types.Int64Type,
			"UpdatedBy":   types.StringType,
		},
	}
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

func (newState *ExchangeListing) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeListing) {
}

func (newState *ExchangeListing) SyncEffectiveFieldsDuringRead(existingState ExchangeListing) {
}

func (a ExchangeListing) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExchangeListing) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreatedAt":    types.Int64Type,
			"CreatedBy":    types.StringType,
			"ExchangeId":   types.StringType,
			"ExchangeName": types.StringType,
			"Id":           types.StringType,
			"ListingId":    types.StringType,
			"ListingName":  types.StringType,
		},
	}
}

type FileInfo struct {
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	DownloadLink types.String `tfsdk:"download_link" tf:"optional"`

	FileParent types.List `tfsdk:"file_parent" tf:"optional,object"`

	Id types.String `tfsdk:"id" tf:"optional"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type" tf:"optional"`

	MimeType types.String `tfsdk:"mime_type" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

func (newState *FileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncEffectiveFieldsDuringRead(existingState FileInfo) {
}

func (a FileInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileParent": reflect.TypeOf(FileParent{}),
	}
}

func (a FileInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreatedAt":           types.Int64Type,
			"DisplayName":         types.StringType,
			"DownloadLink":        types.StringType,
			"FileParent":          FileParent{}.ToAttrType(ctx),
			"Id":                  types.StringType,
			"MarketplaceFileType": types.StringType,
			"MimeType":            types.StringType,
			"Status":              types.StringType,
			"StatusMessage":       types.StringType,
			"UpdatedAt":           types.Int64Type,
		},
	}
}

type FileParent struct {
	FileParentType types.String `tfsdk:"file_parent_type" tf:"optional"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id" tf:"optional"`
}

func (newState *FileParent) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileParent) {
}

func (newState *FileParent) SyncEffectiveFieldsDuringRead(existingState FileParent) {
}

func (a FileParent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a FileParent) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileParentType": types.StringType,
			"ParentId":       types.StringType,
		},
	}
}

// Get an exchange
type GetExchangeRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetExchangeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExchangeRequest) {
}

func (newState *GetExchangeRequest) SyncEffectiveFieldsDuringRead(existingState GetExchangeRequest) {
}

func (a GetExchangeRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetExchangeRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type GetExchangeResponse struct {
	Exchange types.List `tfsdk:"exchange" tf:"optional,object"`
}

func (newState *GetExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExchangeResponse) {
}

func (newState *GetExchangeResponse) SyncEffectiveFieldsDuringRead(existingState GetExchangeResponse) {
}

func (a GetExchangeResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Exchange": reflect.TypeOf(Exchange{}),
	}
}

func (a GetExchangeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Exchange": Exchange{}.ToAttrType(ctx),
		},
	}
}

// Get a file
type GetFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

func (newState *GetFileRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFileRequest) {
}

func (newState *GetFileRequest) SyncEffectiveFieldsDuringRead(existingState GetFileRequest) {
}

func (a GetFileRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetFileRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileId": types.StringType,
		},
	}
}

type GetFileResponse struct {
	FileInfo types.List `tfsdk:"file_info" tf:"optional,object"`
}

func (newState *GetFileResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFileResponse) {
}

func (newState *GetFileResponse) SyncEffectiveFieldsDuringRead(existingState GetFileResponse) {
}

func (a GetFileResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileInfo": reflect.TypeOf(FileInfo{}),
	}
}

func (a GetFileResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileInfo": FileInfo{}.ToAttrType(ctx),
		},
	}
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (a GetLatestVersionProviderAnalyticsDashboardResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetLatestVersionProviderAnalyticsDashboardResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Version": types.Int64Type,
		},
	}
}

// Get listing content metadata
type GetListingContentMetadataRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *GetListingContentMetadataRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingContentMetadataRequest) {
}

func (newState *GetListingContentMetadataRequest) SyncEffectiveFieldsDuringRead(existingState GetListingContentMetadataRequest) {
}

func (a GetListingContentMetadataRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetListingContentMetadataRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type GetListingContentMetadataResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	SharedDataObjects types.List `tfsdk:"shared_data_objects" tf:"optional"`
}

func (newState *GetListingContentMetadataResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingContentMetadataResponse) {
}

func (newState *GetListingContentMetadataResponse) SyncEffectiveFieldsDuringRead(existingState GetListingContentMetadataResponse) {
}

func (a GetListingContentMetadataResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"SharedDataObjects": reflect.TypeOf(SharedDataObject{}),
	}
}

func (a GetListingContentMetadataResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"SharedDataObjects": basetypes.ListType{
				ElemType: SharedDataObject{}.ToAttrType(ctx),
			},
		},
	}
}

// Get listing
type GetListingRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingRequest) {
}

func (newState *GetListingRequest) SyncEffectiveFieldsDuringRead(existingState GetListingRequest) {
}

func (a GetListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type GetListingResponse struct {
	Listing types.List `tfsdk:"listing" tf:"optional,object"`
}

func (newState *GetListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingResponse) {
}

func (newState *GetListingResponse) SyncEffectiveFieldsDuringRead(existingState GetListingResponse) {
}

func (a GetListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listing": reflect.TypeOf(Listing{}),
	}
}

func (a GetListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listing": Listing{}.ToAttrType(ctx),
		},
	}
}

// List listings
type GetListingsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *GetListingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingsRequest) {
}

func (newState *GetListingsRequest) SyncEffectiveFieldsDuringRead(existingState GetListingsRequest) {
}

func (a GetListingsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetListingsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type GetListingsResponse struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetListingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingsResponse) {
}

func (newState *GetListingsResponse) SyncEffectiveFieldsDuringRead(existingState GetListingsResponse) {
}

func (a GetListingsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listings": reflect.TypeOf(Listing{}),
	}
}

func (a GetListingsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listings": basetypes.ListType{
				ElemType: Listing{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`
}

func (newState *GetPersonalizationRequestRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPersonalizationRequestRequest) {
}

func (newState *GetPersonalizationRequestRequest) SyncEffectiveFieldsDuringRead(existingState GetPersonalizationRequestRequest) {
}

func (a GetPersonalizationRequestRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPersonalizationRequestRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
		},
	}
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests types.List `tfsdk:"personalization_requests" tf:"optional"`
}

func (newState *GetPersonalizationRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPersonalizationRequestResponse) {
}

func (newState *GetPersonalizationRequestResponse) SyncEffectiveFieldsDuringRead(existingState GetPersonalizationRequestResponse) {
}

func (a GetPersonalizationRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PersonalizationRequests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

func (a GetPersonalizationRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PersonalizationRequests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.ToAttrType(ctx),
			},
		},
	}
}

// Get a provider
type GetProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderRequest) {
}

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringRead(existingState GetProviderRequest) {
}

func (a GetProviderRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetProviderRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type GetProviderResponse struct {
	Provider types.List `tfsdk:"provider" tf:"optional,object"`
}

func (newState *GetProviderResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderResponse) {
}

func (newState *GetProviderResponse) SyncEffectiveFieldsDuringRead(existingState GetProviderResponse) {
}

func (a GetProviderResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Provider": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a GetProviderResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Provider": ProviderInfo{}.ToAttrType(ctx),
		},
	}
}

type Installation struct {
	Installation types.List `tfsdk:"installation" tf:"optional,object"`
}

func (newState *Installation) SyncEffectiveFieldsDuringCreateOrUpdate(plan Installation) {
}

func (newState *Installation) SyncEffectiveFieldsDuringRead(existingState Installation) {
}

func (a Installation) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Installation": reflect.TypeOf(InstallationDetail{}),
	}
}

func (a Installation) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Installation": InstallationDetail{}.ToAttrType(ctx),
		},
	}
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

	TokenDetail types.List `tfsdk:"token_detail" tf:"optional,object"`

	Tokens types.List `tfsdk:"tokens" tf:"optional"`
}

func (newState *InstallationDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstallationDetail) {
}

func (newState *InstallationDetail) SyncEffectiveFieldsDuringRead(existingState InstallationDetail) {
}

func (a InstallationDetail) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TokenDetail": reflect.TypeOf(TokenDetail{}),
		"Tokens":      reflect.TypeOf(TokenInfo{}),
	}
}

func (a InstallationDetail) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CatalogName":   types.StringType,
			"ErrorMessage":  types.StringType,
			"Id":            types.StringType,
			"InstalledOn":   types.Int64Type,
			"ListingId":     types.StringType,
			"ListingName":   types.StringType,
			"RecipientType": types.StringType,
			"RepoName":      types.StringType,
			"RepoPath":      types.StringType,
			"ShareName":     types.StringType,
			"Status":        types.StringType,
			"TokenDetail":   TokenDetail{}.ToAttrType(ctx),
			"Tokens": basetypes.ListType{
				ElemType: TokenInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// List all installations
type ListAllInstallationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAllInstallationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllInstallationsRequest) {
}

func (newState *ListAllInstallationsRequest) SyncEffectiveFieldsDuringRead(existingState ListAllInstallationsRequest) {
}

func (a ListAllInstallationsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAllInstallationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListAllInstallationsResponse struct {
	Installations types.List `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAllInstallationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllInstallationsResponse) {
}

func (newState *ListAllInstallationsResponse) SyncEffectiveFieldsDuringRead(existingState ListAllInstallationsResponse) {
}

func (a ListAllInstallationsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Installations": reflect.TypeOf(InstallationDetail{}),
	}
}

func (a ListAllInstallationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Installations": basetypes.ListType{
				ElemType: InstallationDetail{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAllPersonalizationRequestsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllPersonalizationRequestsRequest) {
}

func (newState *ListAllPersonalizationRequestsRequest) SyncEffectiveFieldsDuringRead(existingState ListAllPersonalizationRequestsRequest) {
}

func (a ListAllPersonalizationRequestsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAllPersonalizationRequestsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	PersonalizationRequests types.List `tfsdk:"personalization_requests" tf:"optional"`
}

func (newState *ListAllPersonalizationRequestsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllPersonalizationRequestsResponse) {
}

func (newState *ListAllPersonalizationRequestsResponse) SyncEffectiveFieldsDuringRead(existingState ListAllPersonalizationRequestsResponse) {
}

func (a ListAllPersonalizationRequestsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PersonalizationRequests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

func (a ListAllPersonalizationRequestsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"PersonalizationRequests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.ToAttrType(ctx),
			},
		},
	}
}

// List exchange filters
type ListExchangeFiltersRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangeFiltersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangeFiltersRequest) {
}

func (newState *ListExchangeFiltersRequest) SyncEffectiveFieldsDuringRead(existingState ListExchangeFiltersRequest) {
}

func (a ListExchangeFiltersRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListExchangeFiltersRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeId": types.StringType,
			"PageSize":   types.Int64Type,
			"PageToken":  types.StringType,
		},
	}
}

type ListExchangeFiltersResponse struct {
	Filters types.List `tfsdk:"filters" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangeFiltersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangeFiltersResponse) {
}

func (newState *ListExchangeFiltersResponse) SyncEffectiveFieldsDuringRead(existingState ListExchangeFiltersResponse) {
}

func (a ListExchangeFiltersResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Filters": reflect.TypeOf(ExchangeFilter{}),
	}
}

func (a ListExchangeFiltersResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List exchanges for listing
type ListExchangesForListingRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangesForListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesForListingRequest) {
}

func (newState *ListExchangesForListingRequest) SyncEffectiveFieldsDuringRead(existingState ListExchangesForListingRequest) {
}

func (a ListExchangesForListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListExchangesForListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListExchangesForListingResponse struct {
	ExchangeListing types.List `tfsdk:"exchange_listing" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangesForListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesForListingResponse) {
}

func (newState *ListExchangesForListingResponse) SyncEffectiveFieldsDuringRead(existingState ListExchangesForListingResponse) {
}

func (a ListExchangesForListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ExchangeListing": reflect.TypeOf(ExchangeListing{}),
	}
}

func (a ListExchangesForListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeListing": basetypes.ListType{
				ElemType: ExchangeListing{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List exchanges
type ListExchangesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesRequest) {
}

func (newState *ListExchangesRequest) SyncEffectiveFieldsDuringRead(existingState ListExchangesRequest) {
}

func (a ListExchangesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListExchangesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListExchangesResponse struct {
	Exchanges types.List `tfsdk:"exchanges" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesResponse) {
}

func (newState *ListExchangesResponse) SyncEffectiveFieldsDuringRead(existingState ListExchangesResponse) {
}

func (a ListExchangesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Exchanges": reflect.TypeOf(Exchange{}),
	}
}

func (a ListExchangesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Exchanges": basetypes.ListType{
				ElemType: Exchange{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List files
type ListFilesRequest struct {
	FileParent types.List `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListFilesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFilesRequest) {
}

func (newState *ListFilesRequest) SyncEffectiveFieldsDuringRead(existingState ListFilesRequest) {
}

func (a ListFilesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileParent": reflect.TypeOf(FileParent{}),
	}
}

func (a ListFilesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileParent": FileParent{}.ToAttrType(ctx),
			"PageSize":   types.Int64Type,
			"PageToken":  types.StringType,
		},
	}
}

type ListFilesResponse struct {
	FileInfos types.List `tfsdk:"file_infos" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFilesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFilesResponse) {
}

func (newState *ListFilesResponse) SyncEffectiveFieldsDuringRead(existingState ListFilesResponse) {
}

func (a ListFilesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FileInfos": reflect.TypeOf(FileInfo{}),
	}
}

func (a ListFilesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileInfos": basetypes.ListType{
				ElemType: FileInfo{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List all listing fulfillments
type ListFulfillmentsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListFulfillmentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFulfillmentsRequest) {
}

func (newState *ListFulfillmentsRequest) SyncEffectiveFieldsDuringRead(existingState ListFulfillmentsRequest) {
}

func (a ListFulfillmentsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListFulfillmentsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListFulfillmentsResponse struct {
	Fulfillments types.List `tfsdk:"fulfillments" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFulfillmentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFulfillmentsResponse) {
}

func (newState *ListFulfillmentsResponse) SyncEffectiveFieldsDuringRead(existingState ListFulfillmentsResponse) {
}

func (a ListFulfillmentsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Fulfillments": reflect.TypeOf(ListingFulfillment{}),
	}
}

func (a ListFulfillmentsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Fulfillments": basetypes.ListType{
				ElemType: ListingFulfillment{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List installations for a listing
type ListInstallationsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListInstallationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstallationsRequest) {
}

func (newState *ListInstallationsRequest) SyncEffectiveFieldsDuringRead(existingState ListInstallationsRequest) {
}

func (a ListInstallationsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListInstallationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
			"PageSize":  types.Int64Type,
			"PageToken": types.StringType,
		},
	}
}

type ListInstallationsResponse struct {
	Installations types.List `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListInstallationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstallationsResponse) {
}

func (newState *ListInstallationsResponse) SyncEffectiveFieldsDuringRead(existingState ListInstallationsResponse) {
}

func (a ListInstallationsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Installations": reflect.TypeOf(InstallationDetail{}),
	}
}

func (a ListInstallationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Installations": basetypes.ListType{
				ElemType: InstallationDetail{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List listings for exchange
type ListListingsForExchangeRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListListingsForExchangeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsForExchangeRequest) {
}

func (newState *ListListingsForExchangeRequest) SyncEffectiveFieldsDuringRead(existingState ListListingsForExchangeRequest) {
}

func (a ListListingsForExchangeRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListListingsForExchangeRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeId": types.StringType,
			"PageSize":   types.Int64Type,
			"PageToken":  types.StringType,
		},
	}
}

type ListListingsForExchangeResponse struct {
	ExchangeListings types.List `tfsdk:"exchange_listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListListingsForExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsForExchangeResponse) {
}

func (newState *ListListingsForExchangeResponse) SyncEffectiveFieldsDuringRead(existingState ListListingsForExchangeResponse) {
}

func (a ListListingsForExchangeResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ExchangeListings": reflect.TypeOf(ExchangeListing{}),
	}
}

func (a ListListingsForExchangeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExchangeListings": basetypes.ListType{
				ElemType: ExchangeListing{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List listings
type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets types.List `tfsdk:"-"`
	// Matches any of the following categories
	Categories types.List `tfsdk:"-"`
	// Filters each listing based on if it is free.
	IsFree types.Bool `tfsdk:"-"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange types.Bool `tfsdk:"-"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Matches any of the following provider ids
	ProviderIds types.List `tfsdk:"-"`
	// Matches any of the following tags
	Tags types.List `tfsdk:"-"`
}

func (newState *ListListingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsRequest) {
}

func (newState *ListListingsRequest) SyncEffectiveFieldsDuringRead(existingState ListListingsRequest) {
}

func (a ListListingsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Assets":      reflect.TypeOf(types.StringType),
		"Categories":  reflect.TypeOf(types.StringType),
		"ProviderIds": reflect.TypeOf(types.StringType),
		"Tags":        reflect.TypeOf(ListingTag{}),
	}
}

func (a ListListingsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"IsFree":            types.BoolType,
			"IsPrivateExchange": types.BoolType,
			"IsStaffPick":       types.BoolType,
			"PageSize":          types.Int64Type,
			"PageToken":         types.StringType,
			"ProviderIds": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Tags": basetypes.ListType{
				ElemType: ListingTag{}.ToAttrType(ctx),
			},
		},
	}
}

type ListListingsResponse struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListListingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsResponse) {
}

func (newState *ListListingsResponse) SyncEffectiveFieldsDuringRead(existingState ListListingsResponse) {
}

func (a ListListingsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listings": reflect.TypeOf(Listing{}),
	}
}

func (a ListListingsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listings": basetypes.ListType{
				ElemType: Listing{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`

	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *ListProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderAnalyticsDashboardResponse) {
}

func (newState *ListProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderAnalyticsDashboardResponse) {
}

func (a ListProviderAnalyticsDashboardResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListProviderAnalyticsDashboardResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DashboardId": types.StringType,
			"Id":          types.StringType,
			"Version":     types.Int64Type,
		},
	}
}

// List providers
type ListProvidersRequest struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersRequest) {
}

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringRead(existingState ListProvidersRequest) {
}

func (a ListProvidersRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListProvidersRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IsFeatured": types.BoolType,
			"PageSize":   types.Int64Type,
			"PageToken":  types.StringType,
		},
	}
}

type ListProvidersResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse) {
}

func (a ListProvidersResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Providers": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a ListProvidersResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"Providers": basetypes.ListType{
				ElemType: ProviderInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type Listing struct {
	Detail types.List `tfsdk:"detail" tf:"optional,object"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// Next Number: 26
	Summary types.List `tfsdk:"summary" tf:"object"`
}

func (newState *Listing) SyncEffectiveFieldsDuringCreateOrUpdate(plan Listing) {
}

func (newState *Listing) SyncEffectiveFieldsDuringRead(existingState Listing) {
}

func (a Listing) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Detail":  reflect.TypeOf(ListingDetail{}),
		"Summary": reflect.TypeOf(ListingSummary{}),
	}
}

func (a Listing) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Detail":  ListingDetail{}.ToAttrType(ctx),
			"Id":      types.StringType,
			"Summary": ListingSummary{}.ToAttrType(ctx),
		},
	}
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets types.List `tfsdk:"assets" tf:"optional"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd types.Int64 `tfsdk:"collection_date_end" tf:"optional"`
	// The starting date timestamp for when the data spans
	CollectionDateStart types.Int64 `tfsdk:"collection_date_start" tf:"optional"`
	// Smallest unit of time in the dataset
	CollectionGranularity types.List `tfsdk:"collection_granularity" tf:"optional,object"`
	// Whether the dataset is free or paid
	Cost types.String `tfsdk:"cost" tf:"optional"`
	// Where/how the data is sourced
	DataSource types.String `tfsdk:"data_source" tf:"optional"`

	Description types.String `tfsdk:"description" tf:"optional"`

	DocumentationLink types.String `tfsdk:"documentation_link" tf:"optional"`

	EmbeddedNotebookFileInfos types.List `tfsdk:"embedded_notebook_file_infos" tf:"optional"`

	FileIds types.List `tfsdk:"file_ids" tf:"optional"`
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
	Tags types.List `tfsdk:"tags" tf:"optional"`

	TermsOfService types.String `tfsdk:"terms_of_service" tf:"optional"`
	// How often data is updated
	UpdateFrequency types.List `tfsdk:"update_frequency" tf:"optional,object"`
}

func (newState *ListingDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingDetail) {
}

func (newState *ListingDetail) SyncEffectiveFieldsDuringRead(existingState ListingDetail) {
}

func (a ListingDetail) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Assets":                    reflect.TypeOf(types.StringType),
		"CollectionGranularity":     reflect.TypeOf(DataRefreshInfo{}),
		"EmbeddedNotebookFileInfos": reflect.TypeOf(FileInfo{}),
		"FileIds":                   reflect.TypeOf(types.StringType),
		"Tags":                      reflect.TypeOf(ListingTag{}),
		"UpdateFrequency":           reflect.TypeOf(DataRefreshInfo{}),
	}
}

func (a ListingDetail) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"CollectionDateEnd":     types.Int64Type,
			"CollectionDateStart":   types.Int64Type,
			"CollectionGranularity": DataRefreshInfo{}.ToAttrType(ctx),
			"Cost":                  types.StringType,
			"DataSource":            types.StringType,
			"Description":           types.StringType,
			"DocumentationLink":     types.StringType,
			"EmbeddedNotebookFileInfos": basetypes.ListType{
				ElemType: FileInfo{}.ToAttrType(ctx),
			},
			"FileIds": basetypes.ListType{
				ElemType: types.StringType,
			},
			"GeographicalCoverage": types.StringType,
			"License":              types.StringType,
			"PricingModel":         types.StringType,
			"PrivacyPolicyLink":    types.StringType,
			"Size":                 types.Float64Type,
			"SupportLink":          types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ListingTag{}.ToAttrType(ctx),
			},
			"TermsOfService":  types.StringType,
			"UpdateFrequency": DataRefreshInfo{}.ToAttrType(ctx),
		},
	}
}

type ListingFulfillment struct {
	FulfillmentType types.String `tfsdk:"fulfillment_type" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:""`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	RepoInfo types.List `tfsdk:"repo_info" tf:"optional,object"`

	ShareInfo types.List `tfsdk:"share_info" tf:"optional,object"`
}

func (newState *ListingFulfillment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingFulfillment) {
}

func (newState *ListingFulfillment) SyncEffectiveFieldsDuringRead(existingState ListingFulfillment) {
}

func (a ListingFulfillment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RepoInfo":  reflect.TypeOf(RepoInfo{}),
		"ShareInfo": reflect.TypeOf(ShareInfo{}),
	}
}

func (a ListingFulfillment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FulfillmentType": types.StringType,
			"ListingId":       types.StringType,
			"RecipientType":   types.StringType,
			"RepoInfo":        RepoInfo{}.ToAttrType(ctx),
			"ShareInfo":       ShareInfo{}.ToAttrType(ctx),
		},
	}
}

type ListingSetting struct {
	Visibility types.String `tfsdk:"visibility" tf:"optional"`
}

func (newState *ListingSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingSetting) {
}

func (newState *ListingSetting) SyncEffectiveFieldsDuringRead(existingState ListingSetting) {
}

func (a ListingSetting) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListingSetting) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Visibility": types.StringType,
		},
	}
}

// Next Number: 26
type ListingSummary struct {
	Categories types.List `tfsdk:"categories" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	CreatedById types.Int64 `tfsdk:"created_by_id" tf:"optional"`

	ExchangeIds types.List `tfsdk:"exchange_ids" tf:"optional"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo types.List `tfsdk:"git_repo" tf:"optional,object"`

	ListingType types.String `tfsdk:"listingType" tf:""`

	Name types.String `tfsdk:"name" tf:""`

	ProviderId types.String `tfsdk:"provider_id" tf:"optional"`

	ProviderRegion types.List `tfsdk:"provider_region" tf:"optional,object"`

	PublishedAt types.Int64 `tfsdk:"published_at" tf:"optional"`

	PublishedBy types.String `tfsdk:"published_by" tf:"optional"`

	Setting types.List `tfsdk:"setting" tf:"optional,object"`

	Share types.List `tfsdk:"share" tf:"optional,object"`
	// Enums
	Status types.String `tfsdk:"status" tf:"optional"`

	Subtitle types.String `tfsdk:"subtitle" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`

	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`

	UpdatedById types.Int64 `tfsdk:"updated_by_id" tf:"optional"`
}

func (newState *ListingSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingSummary) {
}

func (newState *ListingSummary) SyncEffectiveFieldsDuringRead(existingState ListingSummary) {
}

func (a ListingSummary) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Categories":     reflect.TypeOf(types.StringType),
		"ExchangeIds":    reflect.TypeOf(types.StringType),
		"GitRepo":        reflect.TypeOf(RepoInfo{}),
		"ProviderRegion": reflect.TypeOf(RegionInfo{}),
		"Setting":        reflect.TypeOf(ListingSetting{}),
		"Share":          reflect.TypeOf(ShareInfo{}),
	}
}

func (a ListingSummary) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"CreatedAt":   types.Int64Type,
			"CreatedBy":   types.StringType,
			"CreatedById": types.Int64Type,
			"ExchangeIds": basetypes.ListType{
				ElemType: types.StringType,
			},
			"GitRepo":        RepoInfo{}.ToAttrType(ctx),
			"ListingType":    types.StringType,
			"Name":           types.StringType,
			"ProviderId":     types.StringType,
			"ProviderRegion": RegionInfo{}.ToAttrType(ctx),
			"PublishedAt":    types.Int64Type,
			"PublishedBy":    types.StringType,
			"Setting":        ListingSetting{}.ToAttrType(ctx),
			"Share":          ShareInfo{}.ToAttrType(ctx),
			"Status":         types.StringType,
			"Subtitle":       types.StringType,
			"UpdatedAt":      types.Int64Type,
			"UpdatedBy":      types.StringType,
			"UpdatedById":    types.Int64Type,
		},
	}
}

type ListingTag struct {
	// Tag name (enum)
	TagName types.String `tfsdk:"tag_name" tf:"optional"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues types.List `tfsdk:"tag_values" tf:"optional"`
}

func (newState *ListingTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingTag) {
}

func (newState *ListingTag) SyncEffectiveFieldsDuringRead(existingState ListingTag) {
}

func (a ListingTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"TagValues": reflect.TypeOf(types.StringType),
	}
}

func (a ListingTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"TagName": types.StringType,
			"TagValues": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type PersonalizationRequest struct {
	Comment types.String `tfsdk:"comment" tf:"optional"`

	ConsumerRegion types.List `tfsdk:"consumer_region" tf:"object"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo types.List `tfsdk:"contact_info" tf:"optional,object"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	IntendedUse types.String `tfsdk:"intended_use" tf:"optional"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:"optional"`

	ListingName types.String `tfsdk:"listing_name" tf:"optional"`

	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`

	ProviderId types.String `tfsdk:"provider_id" tf:"optional"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	Share types.List `tfsdk:"share" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:"optional"`

	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

func (newState *PersonalizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalizationRequest) {
}

func (newState *PersonalizationRequest) SyncEffectiveFieldsDuringRead(existingState PersonalizationRequest) {
}

func (a PersonalizationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ConsumerRegion": reflect.TypeOf(RegionInfo{}),
		"ContactInfo":    reflect.TypeOf(ContactInfo{}),
		"Share":          reflect.TypeOf(ShareInfo{}),
	}
}

func (a PersonalizationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment":          types.StringType,
			"ConsumerRegion":   RegionInfo{}.ToAttrType(ctx),
			"ContactInfo":      ContactInfo{}.ToAttrType(ctx),
			"CreatedAt":        types.Int64Type,
			"Id":               types.StringType,
			"IntendedUse":      types.StringType,
			"IsFromLighthouse": types.BoolType,
			"ListingId":        types.StringType,
			"ListingName":      types.StringType,
			"MetastoreId":      types.StringType,
			"ProviderId":       types.StringType,
			"RecipientType":    types.StringType,
			"Share":            ShareInfo{}.ToAttrType(ctx),
			"Status":           types.StringType,
			"StatusMessage":    types.StringType,
			"UpdatedAt":        types.Int64Type,
		},
	}
}

type ProviderAnalyticsDashboard struct {
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *ProviderAnalyticsDashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderAnalyticsDashboard) {
}

func (newState *ProviderAnalyticsDashboard) SyncEffectiveFieldsDuringRead(existingState ProviderAnalyticsDashboard) {
}

func (a ProviderAnalyticsDashboard) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ProviderAnalyticsDashboard) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
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

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
}

func (a ProviderInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ProviderInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"BusinessContactEmail": types.StringType,
			"CompanyWebsiteLink":   types.StringType,
			"DarkModeIconFileId":   types.StringType,
			"DarkModeIconFilePath": types.StringType,
			"Description":          types.StringType,
			"IconFileId":           types.StringType,
			"IconFilePath":         types.StringType,
			"Id":                   types.StringType,
			"IsFeatured":           types.BoolType,
			"Name":                 types.StringType,
			"PrivacyPolicyLink":    types.StringType,
			"PublishedBy":          types.StringType,
			"SupportContactEmail":  types.StringType,
			"TermOfServiceLink":    types.StringType,
		},
	}
}

type RegionInfo struct {
	Cloud types.String `tfsdk:"cloud" tf:"optional"`

	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *RegionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegionInfo) {
}

func (newState *RegionInfo) SyncEffectiveFieldsDuringRead(existingState RegionInfo) {
}

func (a RegionInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RegionInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Cloud":  types.StringType,
			"Region": types.StringType,
		},
	}
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *RemoveExchangeForListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingRequest) {
}

func (newState *RemoveExchangeForListingRequest) SyncEffectiveFieldsDuringRead(existingState RemoveExchangeForListingRequest) {
}

func (a RemoveExchangeForListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RemoveExchangeForListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type RemoveExchangeForListingResponse struct {
}

func (newState *RemoveExchangeForListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingResponse) {
}

func (newState *RemoveExchangeForListingResponse) SyncEffectiveFieldsDuringRead(existingState RemoveExchangeForListingResponse) {
}

func (a RemoveExchangeForListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RemoveExchangeForListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url" tf:""`
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInfo) {
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringRead(existingState RepoInfo) {
}

func (a RepoInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RepoInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"GitRepoUrl": types.StringType,
		},
	}
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName types.String `tfsdk:"repo_name" tf:""`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path" tf:""`
}

func (newState *RepoInstallation) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInstallation) {
}

func (newState *RepoInstallation) SyncEffectiveFieldsDuringRead(existingState RepoInstallation) {
}

func (a RepoInstallation) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RepoInstallation) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RepoName": types.StringType,
			"RepoPath": types.StringType,
		},
	}
}

// Search listings
type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets types.List `tfsdk:"-"`
	// Matches any of the following categories
	Categories types.List `tfsdk:"-"`

	IsFree types.Bool `tfsdk:"-"`

	IsPrivateExchange types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Matches any of the following provider ids
	ProviderIds types.List `tfsdk:"-"`
	// Fuzzy matches query
	Query types.String `tfsdk:"-"`
}

func (newState *SearchListingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchListingsRequest) {
}

func (newState *SearchListingsRequest) SyncEffectiveFieldsDuringRead(existingState SearchListingsRequest) {
}

func (a SearchListingsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Assets":      reflect.TypeOf(types.StringType),
		"Categories":  reflect.TypeOf(types.StringType),
		"ProviderIds": reflect.TypeOf(types.StringType),
	}
}

func (a SearchListingsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"IsFree":            types.BoolType,
			"IsPrivateExchange": types.BoolType,
			"PageSize":          types.Int64Type,
			"PageToken":         types.StringType,
			"ProviderIds": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Query": types.StringType,
		},
	}
}

type SearchListingsResponse struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchListingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchListingsResponse) {
}

func (newState *SearchListingsResponse) SyncEffectiveFieldsDuringRead(existingState SearchListingsResponse) {
}

func (a SearchListingsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listings": reflect.TypeOf(Listing{}),
	}
}

func (a SearchListingsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listings": basetypes.ListType{
				ElemType: Listing{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

type ShareInfo struct {
	Name types.String `tfsdk:"name" tf:""`

	Type types.String `tfsdk:"type" tf:""`
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
}

func (a ShareInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ShareInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name": types.StringType,
			"Type": types.StringType,
		},
	}
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Name of the shared object
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject) {
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringRead(existingState SharedDataObject) {
}

func (a SharedDataObject) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SharedDataObject) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DataObjectType": types.StringType,
			"Name":           types.StringType,
		},
	}
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

func (newState *TokenDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenDetail) {
}

func (newState *TokenDetail) SyncEffectiveFieldsDuringRead(existingState TokenDetail) {
}

func (a TokenDetail) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TokenDetail) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"BearerToken":             types.StringType,
			"Endpoint":                types.StringType,
			"ExpirationTime":          types.StringType,
			"ShareCredentialsVersion": types.Int64Type,
		},
	}
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

func (newState *TokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenInfo) {
}

func (newState *TokenInfo) SyncEffectiveFieldsDuringRead(existingState TokenInfo) {
}

func (a TokenInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ActivationUrl":  types.StringType,
			"CreatedAt":      types.Int64Type,
			"CreatedBy":      types.StringType,
			"ExpirationTime": types.Int64Type,
			"Id":             types.StringType,
			"UpdatedAt":      types.Int64Type,
			"UpdatedBy":      types.StringType,
		},
	}
}

type UpdateExchangeFilterRequest struct {
	Filter types.List `tfsdk:"filter" tf:"object"`

	Id types.String `tfsdk:"-"`
}

func (newState *UpdateExchangeFilterRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeFilterRequest) {
}

func (newState *UpdateExchangeFilterRequest) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeFilterRequest) {
}

func (a UpdateExchangeFilterRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

func (a UpdateExchangeFilterRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter": ExchangeFilter{}.ToAttrType(ctx),
			"Id":     types.StringType,
		},
	}
}

type UpdateExchangeFilterResponse struct {
	Filter types.List `tfsdk:"filter" tf:"optional,object"`
}

func (newState *UpdateExchangeFilterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeFilterResponse) {
}

func (newState *UpdateExchangeFilterResponse) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeFilterResponse) {
}

func (a UpdateExchangeFilterResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

func (a UpdateExchangeFilterResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter": ExchangeFilter{}.ToAttrType(ctx),
		},
	}
}

type UpdateExchangeRequest struct {
	Exchange types.List `tfsdk:"exchange" tf:"object"`

	Id types.String `tfsdk:"-"`
}

func (newState *UpdateExchangeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeRequest) {
}

func (newState *UpdateExchangeRequest) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeRequest) {
}

func (a UpdateExchangeRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Exchange": reflect.TypeOf(Exchange{}),
	}
}

func (a UpdateExchangeRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Exchange": Exchange{}.ToAttrType(ctx),
			"Id":       types.StringType,
		},
	}
}

type UpdateExchangeResponse struct {
	Exchange types.List `tfsdk:"exchange" tf:"optional,object"`
}

func (newState *UpdateExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeResponse) {
}

func (newState *UpdateExchangeResponse) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeResponse) {
}

func (a UpdateExchangeResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Exchange": reflect.TypeOf(Exchange{}),
	}
}

func (a UpdateExchangeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Exchange": Exchange{}.ToAttrType(ctx),
		},
	}
}

type UpdateInstallationRequest struct {
	Installation types.List `tfsdk:"installation" tf:"object"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token" tf:"optional"`
}

func (newState *UpdateInstallationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInstallationRequest) {
}

func (newState *UpdateInstallationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateInstallationRequest) {
}

func (a UpdateInstallationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Installation": reflect.TypeOf(InstallationDetail{}),
	}
}

func (a UpdateInstallationRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Installation":   InstallationDetail{}.ToAttrType(ctx),
			"InstallationId": types.StringType,
			"ListingId":      types.StringType,
			"RotateToken":    types.BoolType,
		},
	}
}

type UpdateInstallationResponse struct {
	Installation types.List `tfsdk:"installation" tf:"optional,object"`
}

func (newState *UpdateInstallationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInstallationResponse) {
}

func (newState *UpdateInstallationResponse) SyncEffectiveFieldsDuringRead(existingState UpdateInstallationResponse) {
}

func (a UpdateInstallationResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Installation": reflect.TypeOf(InstallationDetail{}),
	}
}

func (a UpdateInstallationResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Installation": InstallationDetail{}.ToAttrType(ctx),
		},
	}
}

type UpdateListingRequest struct {
	Id types.String `tfsdk:"-"`

	Listing types.List `tfsdk:"listing" tf:"object"`
}

func (newState *UpdateListingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateListingRequest) {
}

func (newState *UpdateListingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateListingRequest) {
}

func (a UpdateListingRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listing": reflect.TypeOf(Listing{}),
	}
}

func (a UpdateListingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id":      types.StringType,
			"Listing": Listing{}.ToAttrType(ctx),
		},
	}
}

type UpdateListingResponse struct {
	Listing types.List `tfsdk:"listing" tf:"optional,object"`
}

func (newState *UpdateListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateListingResponse) {
}

func (newState *UpdateListingResponse) SyncEffectiveFieldsDuringRead(existingState UpdateListingResponse) {
}

func (a UpdateListingResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Listing": reflect.TypeOf(Listing{}),
	}
}

func (a UpdateListingResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Listing": Listing{}.ToAttrType(ctx),
		},
	}
}

type UpdatePersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`

	Reason types.String `tfsdk:"reason" tf:"optional"`

	RequestId types.String `tfsdk:"-"`

	Share types.List `tfsdk:"share" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:""`
}

func (newState *UpdatePersonalizationRequestRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalizationRequestRequest) {
}

func (newState *UpdatePersonalizationRequestRequest) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalizationRequestRequest) {
}

func (a UpdatePersonalizationRequestRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Share": reflect.TypeOf(ShareInfo{}),
	}
}

func (a UpdatePersonalizationRequestRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ListingId": types.StringType,
			"Reason":    types.StringType,
			"RequestId": types.StringType,
			"Share":     ShareInfo{}.ToAttrType(ctx),
			"Status":    types.StringType,
		},
	}
}

type UpdatePersonalizationRequestResponse struct {
	Request types.List `tfsdk:"request" tf:"optional,object"`
}

func (newState *UpdatePersonalizationRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalizationRequestResponse) {
}

func (newState *UpdatePersonalizationRequestResponse) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalizationRequestResponse) {
}

func (a UpdatePersonalizationRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Request": reflect.TypeOf(PersonalizationRequest{}),
	}
}

func (a UpdatePersonalizationRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Request": PersonalizationRequest{}.ToAttrType(ctx),
		},
	}
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *UpdateProviderAnalyticsDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderAnalyticsDashboardRequest) {
}

func (newState *UpdateProviderAnalyticsDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UpdateProviderAnalyticsDashboardRequest) {
}

func (a UpdateProviderAnalyticsDashboardRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateProviderAnalyticsDashboardRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id":      types.StringType,
			"Version": types.Int64Type,
		},
	}
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *UpdateProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderAnalyticsDashboardResponse) {
}

func (newState *UpdateProviderAnalyticsDashboardResponse) SyncEffectiveFieldsDuringRead(existingState UpdateProviderAnalyticsDashboardResponse) {
}

func (a UpdateProviderAnalyticsDashboardResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateProviderAnalyticsDashboardResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DashboardId": types.StringType,
			"Id":          types.StringType,
			"Version":     types.Int64Type,
		},
	}
}

type UpdateProviderRequest struct {
	Id types.String `tfsdk:"-"`

	Provider types.List `tfsdk:"provider" tf:"object"`
}

func (newState *UpdateProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderRequest) {
}

func (newState *UpdateProviderRequest) SyncEffectiveFieldsDuringRead(existingState UpdateProviderRequest) {
}

func (a UpdateProviderRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Provider": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a UpdateProviderRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id":       types.StringType,
			"Provider": ProviderInfo{}.ToAttrType(ctx),
		},
	}
}

type UpdateProviderResponse struct {
	Provider types.List `tfsdk:"provider" tf:"optional,object"`
}

func (newState *UpdateProviderResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderResponse) {
}

func (newState *UpdateProviderResponse) SyncEffectiveFieldsDuringRead(existingState UpdateProviderResponse) {
}

func (a UpdateProviderResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Provider": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a UpdateProviderResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Provider": ProviderInfo{}.ToAttrType(ctx),
		},
	}
}
