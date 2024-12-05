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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddExchangeForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddExchangeForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of AddExchangeForListingRequest in the Terraform plugin framework type
// system.
func (a AddExchangeForListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"listing_id":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddExchangeForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddExchangeForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_for_listing": reflect.TypeOf(ExchangeListing{}),
	}
}

// ToObjectType returns the representation of AddExchangeForListingResponse in the Terraform plugin framework type
// system.
func (a AddExchangeForListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_for_listing": basetypes.ListType{
				ElemType: ExchangeListing{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of BatchGetListingsRequest in the Terraform plugin framework type
// system.
func (a BatchGetListingsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of BatchGetListingsResponse in the Terraform plugin framework type
// system.
func (a BatchGetListingsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of BatchGetProvidersRequest in the Terraform plugin framework type
// system.
func (a BatchGetProvidersRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetProvidersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of BatchGetProvidersResponse in the Terraform plugin framework type
// system.
func (a BatchGetProvidersResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConsumerTerms.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConsumerTerms) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ConsumerTerms in the Terraform plugin framework type
// system.
func (a ConsumerTerms) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContactInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ContactInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ContactInfo in the Terraform plugin framework type
// system.
func (a ContactInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"company":    types.StringType,
			"email":      types.StringType,
			"first_name": types.StringType,
			"last_name":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// ToObjectType returns the representation of CreateExchangeFilterRequest in the Terraform plugin framework type
// system.
func (a CreateExchangeFilterRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreateExchangeFilterResponse in the Terraform plugin framework type
// system.
func (a CreateExchangeFilterResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// ToObjectType returns the representation of CreateExchangeRequest in the Terraform plugin framework type
// system.
func (a CreateExchangeRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreateExchangeResponse in the Terraform plugin framework type
// system.
func (a CreateExchangeResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// ToObjectType returns the representation of CreateFileRequest in the Terraform plugin framework type
// system.
func (a CreateFileRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name": types.StringType,
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.ToObjectType(ctx),
			},
			"marketplace_file_type": types.StringType,
			"mime_type":             types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo{}),
	}
}

// ToObjectType returns the representation of CreateFileResponse in the Terraform plugin framework type
// system.
func (a CreateFileResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": basetypes.ListType{
				ElemType: FileInfo{}.ToObjectType(ctx),
			},
			"upload_url": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms{}),
		"repo_detail":             reflect.TypeOf(RepoInstallation{}),
	}
}

// ToObjectType returns the representation of CreateInstallationRequest in the Terraform plugin framework type
// system.
func (a CreateInstallationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": basetypes.ListType{
				ElemType: ConsumerTerms{}.ToObjectType(ctx),
			},
			"catalog_name":   types.StringType,
			"listing_id":     types.StringType,
			"recipient_type": types.StringType,
			"repo_detail": basetypes.ListType{
				ElemType: RepoInstallation{}.ToObjectType(ctx),
			},
			"share_name": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of CreateListingRequest in the Terraform plugin framework type
// system.
func (a CreateListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreateListingResponse in the Terraform plugin framework type
// system.
func (a CreateListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePersonalizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePersonalizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms{}),
	}
}

// ToObjectType returns the representation of CreatePersonalizationRequest in the Terraform plugin framework type
// system.
func (a CreatePersonalizationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": basetypes.ListType{
				ElemType: ConsumerTerms{}.ToObjectType(ctx),
			},
			"comment":            types.StringType,
			"company":            types.StringType,
			"first_name":         types.StringType,
			"intended_use":       types.StringType,
			"is_from_lighthouse": types.BoolType,
			"last_name":          types.StringType,
			"listing_id":         types.StringType,
			"recipient_type":     types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreatePersonalizationRequestResponse in the Terraform plugin framework type
// system.
func (a CreatePersonalizationRequestResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of CreateProviderRequest in the Terraform plugin framework type
// system.
func (a CreateProviderRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreateProviderResponse in the Terraform plugin framework type
// system.
func (a CreateProviderResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataRefreshInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataRefreshInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DataRefreshInfo in the Terraform plugin framework type
// system.
func (a DataRefreshInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"interval": types.Int64Type,
			"unit":     types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteExchangeFilterRequest in the Terraform plugin framework type
// system.
func (a DeleteExchangeFilterRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeFilterResponse struct {
}

func (newState *DeleteExchangeFilterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterResponse) {
}

func (newState *DeleteExchangeFilterResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeFilterResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteExchangeFilterResponse in the Terraform plugin framework type
// system.
func (a DeleteExchangeFilterResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteExchangeRequest in the Terraform plugin framework type
// system.
func (a DeleteExchangeRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeResponse struct {
}

func (newState *DeleteExchangeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeResponse) {
}

func (newState *DeleteExchangeResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteExchangeResponse in the Terraform plugin framework type
// system.
func (a DeleteExchangeResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteFileRequest in the Terraform plugin framework type
// system.
func (a DeleteFileRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type DeleteFileResponse struct {
}

func (newState *DeleteFileResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileResponse) {
}

func (newState *DeleteFileResponse) SyncEffectiveFieldsDuringRead(existingState DeleteFileResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteFileResponse in the Terraform plugin framework type
// system.
func (a DeleteFileResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteInstallationRequest in the Terraform plugin framework type
// system.
func (a DeleteInstallationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
		},
	}
}

type DeleteInstallationResponse struct {
}

func (newState *DeleteInstallationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstallationResponse) {
}

func (newState *DeleteInstallationResponse) SyncEffectiveFieldsDuringRead(existingState DeleteInstallationResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstallationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstallationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteInstallationResponse in the Terraform plugin framework type
// system.
func (a DeleteInstallationResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteListingRequest in the Terraform plugin framework type
// system.
func (a DeleteListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteListingResponse struct {
}

func (newState *DeleteListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteListingResponse) {
}

func (newState *DeleteListingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteListingResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteListingResponse in the Terraform plugin framework type
// system.
func (a DeleteListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteProviderRequest in the Terraform plugin framework type
// system.
func (a DeleteProviderRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteProviderResponse struct {
}

func (newState *DeleteProviderResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderResponse) {
}

func (newState *DeleteProviderResponse) SyncEffectiveFieldsDuringRead(existingState DeleteProviderResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteProviderResponse in the Terraform plugin framework type
// system.
func (a DeleteProviderResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Exchange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Exchange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters":         reflect.TypeOf(ExchangeFilter{}),
		"linked_listings": reflect.TypeOf(ExchangeListing{}),
	}
}

// ToObjectType returns the representation of Exchange in the Terraform plugin framework type
// system.
func (a Exchange) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToObjectType(ctx),
			},
			"id": types.StringType,
			"linked_listings": basetypes.ListType{
				ElemType: ExchangeListing{}.ToObjectType(ctx),
			},
			"name":       types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ExchangeFilter in the Terraform plugin framework type
// system.
func (a ExchangeFilter) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"exchange_id":  types.StringType,
			"filter_type":  types.StringType,
			"filter_value": types.StringType,
			"id":           types.StringType,
			"name":         types.StringType,
			"updated_at":   types.Int64Type,
			"updated_by":   types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeListing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeListing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ExchangeListing in the Terraform plugin framework type
// system.
func (a ExchangeListing) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":    types.Int64Type,
			"created_by":    types.StringType,
			"exchange_id":   types.StringType,
			"exchange_name": types.StringType,
			"id":            types.StringType,
			"listing_id":    types.StringType,
			"listing_name":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// ToObjectType returns the representation of FileInfo in the Terraform plugin framework type
// system.
func (a FileInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":    types.Int64Type,
			"display_name":  types.StringType,
			"download_link": types.StringType,
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.ToObjectType(ctx),
			},
			"id":                    types.StringType,
			"marketplace_file_type": types.StringType,
			"mime_type":             types.StringType,
			"status":                types.StringType,
			"status_message":        types.StringType,
			"updated_at":            types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileParent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileParent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of FileParent in the Terraform plugin framework type
// system.
func (a FileParent) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent_type": types.StringType,
			"parent_id":        types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetExchangeRequest in the Terraform plugin framework type
// system.
func (a GetExchangeRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// ToObjectType returns the representation of GetExchangeResponse in the Terraform plugin framework type
// system.
func (a GetExchangeResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetFileRequest in the Terraform plugin framework type
// system.
func (a GetFileRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo{}),
	}
}

// ToObjectType returns the representation of GetFileResponse in the Terraform plugin framework type
// system.
func (a GetFileResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": basetypes.ListType{
				ElemType: FileInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLatestVersionProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetLatestVersionProviderAnalyticsDashboardResponse in the Terraform plugin framework type
// system.
func (a GetLatestVersionProviderAnalyticsDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingContentMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingContentMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetListingContentMetadataRequest in the Terraform plugin framework type
// system.
func (a GetListingContentMetadataRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingContentMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingContentMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shared_data_objects": reflect.TypeOf(SharedDataObject{}),
	}
}

// ToObjectType returns the representation of GetListingContentMetadataResponse in the Terraform plugin framework type
// system.
func (a GetListingContentMetadataResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shared_data_objects": basetypes.ListType{
				ElemType: SharedDataObject{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetListingRequest in the Terraform plugin framework type
// system.
func (a GetListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of GetListingResponse in the Terraform plugin framework type
// system.
func (a GetListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetListingsRequest in the Terraform plugin framework type
// system.
func (a GetListingsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of GetListingsResponse in the Terraform plugin framework type
// system.
func (a GetListingsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalizationRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalizationRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetPersonalizationRequestRequest in the Terraform plugin framework type
// system.
func (a GetPersonalizationRequestRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// ToObjectType returns the representation of GetPersonalizationRequestResponse in the Terraform plugin framework type
// system.
func (a GetPersonalizationRequestResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetProviderRequest in the Terraform plugin framework type
// system.
func (a GetProviderRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of GetProviderResponse in the Terraform plugin framework type
// system.
func (a GetProviderResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Installation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Installation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// ToObjectType returns the representation of Installation in the Terraform plugin framework type
// system.
func (a Installation) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallationDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstallationDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_detail": reflect.TypeOf(TokenDetail{}),
		"tokens":       reflect.TypeOf(TokenInfo{}),
	}
}

// ToObjectType returns the representation of InstallationDetail in the Terraform plugin framework type
// system.
func (a InstallationDetail) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"error_message":  types.StringType,
			"id":             types.StringType,
			"installed_on":   types.Int64Type,
			"listing_id":     types.StringType,
			"listing_name":   types.StringType,
			"recipient_type": types.StringType,
			"repo_name":      types.StringType,
			"repo_path":      types.StringType,
			"share_name":     types.StringType,
			"status":         types.StringType,
			"token_detail": basetypes.ListType{
				ElemType: TokenDetail{}.ToObjectType(ctx),
			},
			"tokens": basetypes.ListType{
				ElemType: TokenInfo{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllInstallationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllInstallationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListAllInstallationsRequest in the Terraform plugin framework type
// system.
func (a ListAllInstallationsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllInstallationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllInstallationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail{}),
	}
}

// ToObjectType returns the representation of ListAllInstallationsResponse in the Terraform plugin framework type
// system.
func (a ListAllInstallationsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllPersonalizationRequestsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllPersonalizationRequestsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListAllPersonalizationRequestsRequest in the Terraform plugin framework type
// system.
func (a ListAllPersonalizationRequestsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllPersonalizationRequestsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllPersonalizationRequestsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// ToObjectType returns the representation of ListAllPersonalizationRequestsResponse in the Terraform plugin framework type
// system.
func (a ListAllPersonalizationRequestsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangeFiltersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangeFiltersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListExchangeFiltersRequest in the Terraform plugin framework type
// system.
func (a ListExchangeFiltersRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangeFiltersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangeFiltersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters": reflect.TypeOf(ExchangeFilter{}),
	}
}

// ToObjectType returns the representation of ListExchangeFiltersResponse in the Terraform plugin framework type
// system.
func (a ListExchangeFiltersResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListExchangesForListingRequest in the Terraform plugin framework type
// system.
func (a ListExchangesForListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listing": reflect.TypeOf(ExchangeListing{}),
	}
}

// ToObjectType returns the representation of ListExchangesForListingResponse in the Terraform plugin framework type
// system.
func (a ListExchangesForListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listing": basetypes.ListType{
				ElemType: ExchangeListing{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListExchangesRequest in the Terraform plugin framework type
// system.
func (a ListExchangesRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchanges": reflect.TypeOf(Exchange{}),
	}
}

// ToObjectType returns the representation of ListExchangesResponse in the Terraform plugin framework type
// system.
func (a ListExchangesResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchanges": basetypes.ListType{
				ElemType: Exchange{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// ToObjectType returns the representation of ListFilesRequest in the Terraform plugin framework type
// system.
func (a ListFilesRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.ToObjectType(ctx),
			},
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_infos": reflect.TypeOf(FileInfo{}),
	}
}

// ToObjectType returns the representation of ListFilesResponse in the Terraform plugin framework type
// system.
func (a ListFilesResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_infos": basetypes.ListType{
				ElemType: FileInfo{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFulfillmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFulfillmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListFulfillmentsRequest in the Terraform plugin framework type
// system.
func (a ListFulfillmentsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFulfillmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFulfillmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fulfillments": reflect.TypeOf(ListingFulfillment{}),
	}
}

// ToObjectType returns the representation of ListFulfillmentsResponse in the Terraform plugin framework type
// system.
func (a ListFulfillmentsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillments": basetypes.ListType{
				ElemType: ListingFulfillment{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstallationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstallationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListInstallationsRequest in the Terraform plugin framework type
// system.
func (a ListInstallationsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstallationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstallationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail{}),
	}
}

// ToObjectType returns the representation of ListInstallationsResponse in the Terraform plugin framework type
// system.
func (a ListInstallationsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsForExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsForExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListListingsForExchangeRequest in the Terraform plugin framework type
// system.
func (a ListListingsForExchangeRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsForExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsForExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listings": reflect.TypeOf(ExchangeListing{}),
	}
}

// ToObjectType returns the representation of ListListingsForExchangeResponse in the Terraform plugin framework type
// system.
func (a ListListingsForExchangeResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listings": basetypes.ListType{
				ElemType: ExchangeListing{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":       reflect.TypeOf(types.String{}),
		"categories":   reflect.TypeOf(types.String{}),
		"provider_ids": reflect.TypeOf(types.String{}),
		"tags":         reflect.TypeOf(ListingTag{}),
	}
}

// ToObjectType returns the representation of ListListingsRequest in the Terraform plugin framework type
// system.
func (a ListListingsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_free":             types.BoolType,
			"is_private_exchange": types.BoolType,
			"is_staff_pick":       types.BoolType,
			"page_size":           types.Int64Type,
			"page_token":          types.StringType,
			"provider_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tags": basetypes.ListType{
				ElemType: ListingTag{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of ListListingsResponse in the Terraform plugin framework type
// system.
func (a ListListingsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListProviderAnalyticsDashboardResponse in the Terraform plugin framework type
// system.
func (a ListProviderAnalyticsDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListProvidersRequest in the Terraform plugin framework type
// system.
func (a ListProvidersRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_featured": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of ListProvidersResponse in the Terraform plugin framework type
// system.
func (a ListProvidersResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Listing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Listing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"detail":  reflect.TypeOf(ListingDetail{}),
		"summary": reflect.TypeOf(ListingSummary{}),
	}
}

// ToObjectType returns the representation of Listing in the Terraform plugin framework type
// system.
func (a Listing) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"detail": basetypes.ListType{
				ElemType: ListingDetail{}.ToObjectType(ctx),
			},
			"id": types.StringType,
			"summary": basetypes.ListType{
				ElemType: ListingSummary{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":                       reflect.TypeOf(types.String{}),
		"collection_granularity":       reflect.TypeOf(DataRefreshInfo{}),
		"embedded_notebook_file_infos": reflect.TypeOf(FileInfo{}),
		"file_ids":                     reflect.TypeOf(types.String{}),
		"tags":                         reflect.TypeOf(ListingTag{}),
		"update_frequency":             reflect.TypeOf(DataRefreshInfo{}),
	}
}

// ToObjectType returns the representation of ListingDetail in the Terraform plugin framework type
// system.
func (a ListingDetail) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"collection_date_end":   types.Int64Type,
			"collection_date_start": types.Int64Type,
			"collection_granularity": basetypes.ListType{
				ElemType: DataRefreshInfo{}.ToObjectType(ctx),
			},
			"cost":               types.StringType,
			"data_source":        types.StringType,
			"description":        types.StringType,
			"documentation_link": types.StringType,
			"embedded_notebook_file_infos": basetypes.ListType{
				ElemType: FileInfo{}.ToObjectType(ctx),
			},
			"file_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"geographical_coverage": types.StringType,
			"license":               types.StringType,
			"pricing_model":         types.StringType,
			"privacy_policy_link":   types.StringType,
			"size":                  types.Float64Type,
			"support_link":          types.StringType,
			"tags": basetypes.ListType{
				ElemType: ListingTag{}.ToObjectType(ctx),
			},
			"terms_of_service": types.StringType,
			"update_frequency": basetypes.ListType{
				ElemType: DataRefreshInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingFulfillment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingFulfillment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repo_info":  reflect.TypeOf(RepoInfo{}),
		"share_info": reflect.TypeOf(ShareInfo{}),
	}
}

// ToObjectType returns the representation of ListingFulfillment in the Terraform plugin framework type
// system.
func (a ListingFulfillment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillment_type": types.StringType,
			"listing_id":       types.StringType,
			"recipient_type":   types.StringType,
			"repo_info": basetypes.ListType{
				ElemType: RepoInfo{}.ToObjectType(ctx),
			},
			"share_info": basetypes.ListType{
				ElemType: ShareInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListingSetting in the Terraform plugin framework type
// system.
func (a ListingSetting) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"visibility": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingSummary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"categories":      reflect.TypeOf(types.String{}),
		"exchange_ids":    reflect.TypeOf(types.String{}),
		"git_repo":        reflect.TypeOf(RepoInfo{}),
		"provider_region": reflect.TypeOf(RegionInfo{}),
		"setting":         reflect.TypeOf(ListingSetting{}),
		"share":           reflect.TypeOf(ShareInfo{}),
	}
}

// ToObjectType returns the representation of ListingSummary in the Terraform plugin framework type
// system.
func (a ListingSummary) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"created_at":    types.Int64Type,
			"created_by":    types.StringType,
			"created_by_id": types.Int64Type,
			"exchange_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"git_repo": basetypes.ListType{
				ElemType: RepoInfo{}.ToObjectType(ctx),
			},
			"listingType": types.StringType,
			"name":        types.StringType,
			"provider_id": types.StringType,
			"provider_region": basetypes.ListType{
				ElemType: RegionInfo{}.ToObjectType(ctx),
			},
			"published_at": types.Int64Type,
			"published_by": types.StringType,
			"setting": basetypes.ListType{
				ElemType: ListingSetting{}.ToObjectType(ctx),
			},
			"share": basetypes.ListType{
				ElemType: ShareInfo{}.ToObjectType(ctx),
			},
			"status":        types.StringType,
			"subtitle":      types.StringType,
			"updated_at":    types.Int64Type,
			"updated_by":    types.StringType,
			"updated_by_id": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_values": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ListingTag in the Terraform plugin framework type
// system.
func (a ListingTag) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_name": types.StringType,
			"tag_values": basetypes.ListType{
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PersonalizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"consumer_region": reflect.TypeOf(RegionInfo{}),
		"contact_info":    reflect.TypeOf(ContactInfo{}),
		"share":           reflect.TypeOf(ShareInfo{}),
	}
}

// ToObjectType returns the representation of PersonalizationRequest in the Terraform plugin framework type
// system.
func (a PersonalizationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"consumer_region": basetypes.ListType{
				ElemType: RegionInfo{}.ToObjectType(ctx),
			},
			"contact_info": basetypes.ListType{
				ElemType: ContactInfo{}.ToObjectType(ctx),
			},
			"created_at":         types.Int64Type,
			"id":                 types.StringType,
			"intended_use":       types.StringType,
			"is_from_lighthouse": types.BoolType,
			"listing_id":         types.StringType,
			"listing_name":       types.StringType,
			"metastore_id":       types.StringType,
			"provider_id":        types.StringType,
			"recipient_type":     types.StringType,
			"share": basetypes.ListType{
				ElemType: ShareInfo{}.ToObjectType(ctx),
			},
			"status":         types.StringType,
			"status_message": types.StringType,
			"updated_at":     types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderAnalyticsDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderAnalyticsDashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ProviderAnalyticsDashboard in the Terraform plugin framework type
// system.
func (a ProviderAnalyticsDashboard) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ProviderInfo in the Terraform plugin framework type
// system.
func (a ProviderInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"business_contact_email":   types.StringType,
			"company_website_link":     types.StringType,
			"dark_mode_icon_file_id":   types.StringType,
			"dark_mode_icon_file_path": types.StringType,
			"description":              types.StringType,
			"icon_file_id":             types.StringType,
			"icon_file_path":           types.StringType,
			"id":                       types.StringType,
			"is_featured":              types.BoolType,
			"name":                     types.StringType,
			"privacy_policy_link":      types.StringType,
			"published_by":             types.StringType,
			"support_contact_email":    types.StringType,
			"term_of_service_link":     types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegionInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RegionInfo in the Terraform plugin framework type
// system.
func (a RegionInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud":  types.StringType,
			"region": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveExchangeForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveExchangeForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RemoveExchangeForListingRequest in the Terraform plugin framework type
// system.
func (a RemoveExchangeForListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type RemoveExchangeForListingResponse struct {
}

func (newState *RemoveExchangeForListingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingResponse) {
}

func (newState *RemoveExchangeForListingResponse) SyncEffectiveFieldsDuringRead(existingState RemoveExchangeForListingResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveExchangeForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveExchangeForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RemoveExchangeForListingResponse in the Terraform plugin framework type
// system.
func (a RemoveExchangeForListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RepoInfo in the Terraform plugin framework type
// system.
func (a RepoInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"git_repo_url": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoInstallation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoInstallation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RepoInstallation in the Terraform plugin framework type
// system.
func (a RepoInstallation) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_name": types.StringType,
			"repo_path": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":       reflect.TypeOf(types.String{}),
		"categories":   reflect.TypeOf(types.String{}),
		"provider_ids": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of SearchListingsRequest in the Terraform plugin framework type
// system.
func (a SearchListingsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"categories": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_free":             types.BoolType,
			"is_private_exchange": types.BoolType,
			"page_size":           types.Int64Type,
			"page_token":          types.StringType,
			"provider_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"query": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of SearchListingsResponse in the Terraform plugin framework type
// system.
func (a SearchListingsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ShareInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ShareInfo in the Terraform plugin framework type
// system.
func (a ShareInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"type": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharedDataObject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SharedDataObject in the Terraform plugin framework type
// system.
func (a SharedDataObject) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_object_type": types.StringType,
			"name":             types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TokenDetail in the Terraform plugin framework type
// system.
func (a TokenDetail) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearerToken":             types.StringType,
			"endpoint":                types.StringType,
			"expirationTime":          types.StringType,
			"shareCredentialsVersion": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TokenInfo in the Terraform plugin framework type
// system.
func (a TokenInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url":  types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"expiration_time": types.Int64Type,
			"id":              types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// ToObjectType returns the representation of UpdateExchangeFilterRequest in the Terraform plugin framework type
// system.
func (a UpdateExchangeFilterRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToObjectType(ctx),
			},
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// ToObjectType returns the representation of UpdateExchangeFilterResponse in the Terraform plugin framework type
// system.
func (a UpdateExchangeFilterResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// ToObjectType returns the representation of UpdateExchangeRequest in the Terraform plugin framework type
// system.
func (a UpdateExchangeRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.ToObjectType(ctx),
			},
			"id": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// ToObjectType returns the representation of UpdateExchangeResponse in the Terraform plugin framework type
// system.
func (a UpdateExchangeResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// ToObjectType returns the representation of UpdateInstallationRequest in the Terraform plugin framework type
// system.
func (a UpdateInstallationRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.ToObjectType(ctx),
			},
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
			"rotate_token":    types.BoolType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInstallationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInstallationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// ToObjectType returns the representation of UpdateInstallationResponse in the Terraform plugin framework type
// system.
func (a UpdateInstallationResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of UpdateListingRequest in the Terraform plugin framework type
// system.
func (a UpdateListingRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"listing": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// ToObjectType returns the representation of UpdateListingResponse in the Terraform plugin framework type
// system.
func (a UpdateListingResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalizationRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalizationRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"share": reflect.TypeOf(ShareInfo{}),
	}
}

// ToObjectType returns the representation of UpdatePersonalizationRequestRequest in the Terraform plugin framework type
// system.
func (a UpdatePersonalizationRequestRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"reason":     types.StringType,
			"request_id": types.StringType,
			"share": basetypes.ListType{
				ElemType: ShareInfo{}.ToObjectType(ctx),
			},
			"status": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// ToObjectType returns the representation of UpdatePersonalizationRequestResponse in the Terraform plugin framework type
// system.
func (a UpdatePersonalizationRequestResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": basetypes.ListType{
				ElemType: PersonalizationRequest{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UpdateProviderAnalyticsDashboardRequest in the Terraform plugin framework type
// system.
func (a UpdateProviderAnalyticsDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":      types.StringType,
			"version": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UpdateProviderAnalyticsDashboardResponse in the Terraform plugin framework type
// system.
func (a UpdateProviderAnalyticsDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of UpdateProviderRequest in the Terraform plugin framework type
// system.
func (a UpdateProviderRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// ToObjectType returns the representation of UpdateProviderResponse in the Terraform plugin framework type
// system.
func (a UpdateProviderResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.ToObjectType(ctx),
			},
		},
	}
}
