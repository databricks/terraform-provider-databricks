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

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddExchangeForListingRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:""`

	ListingId types.String `tfsdk:"listing_id" tf:""`
}

func (newState *AddExchangeForListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddExchangeForListingRequest_SdkV2) {
}

func (newState *AddExchangeForListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState AddExchangeForListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddExchangeForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddExchangeForListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AddExchangeForListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"listing_id":  o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddExchangeForListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"listing_id":  types.StringType,
		},
	}
}

type AddExchangeForListingResponse_SdkV2 struct {
	ExchangeForListing types.List `tfsdk:"exchange_for_listing" tf:"optional,object"`
}

func (newState *AddExchangeForListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddExchangeForListingResponse_SdkV2) {
}

func (newState *AddExchangeForListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AddExchangeForListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddExchangeForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddExchangeForListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_for_listing": reflect.TypeOf(ExchangeListing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AddExchangeForListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_for_listing": o.ExchangeForListing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddExchangeForListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_for_listing": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
		},
	}
}

// GetExchangeForListing returns the value of the ExchangeForListing field in AddExchangeForListingResponse_SdkV2 as
// a ExchangeListing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AddExchangeForListingResponse_SdkV2) GetExchangeForListing(ctx context.Context) (ExchangeListing_SdkV2, bool) {
	var e ExchangeListing_SdkV2
	if o.ExchangeForListing.IsNull() || o.ExchangeForListing.IsUnknown() {
		return e, false
	}
	var v []ExchangeListing_SdkV2
	d := o.ExchangeForListing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchangeForListing sets the value of the ExchangeForListing field in AddExchangeForListingResponse_SdkV2.
func (o *AddExchangeForListingResponse_SdkV2) SetExchangeForListing(ctx context.Context, v ExchangeListing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_for_listing"]
	o.ExchangeForListing = types.ListValueMust(t, vs)
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest_SdkV2 struct {
	Ids types.List `tfsdk:"-"`
}

func (newState *BatchGetListingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetListingsRequest_SdkV2) {
}

func (newState *BatchGetListingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState BatchGetListingsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetListingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o BatchGetListingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": o.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetListingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in BatchGetListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetListingsRequest_SdkV2) GetIds(ctx context.Context) ([]types.String, bool) {
	if o.Ids.IsNull() || o.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in BatchGetListingsRequest_SdkV2.
func (o *BatchGetListingsRequest_SdkV2) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Ids = types.ListValueMust(t, vs)
}

type BatchGetListingsResponse_SdkV2 struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`
}

func (newState *BatchGetListingsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetListingsResponse_SdkV2) {
}

func (newState *BatchGetListingsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState BatchGetListingsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetListingsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o BatchGetListingsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings": o.Listings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetListingsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListings returns the value of the Listings field in BatchGetListingsResponse_SdkV2 as
// a slice of Listing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetListingsResponse_SdkV2) GetListings(ctx context.Context) ([]Listing_SdkV2, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing_SdkV2
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in BatchGetListingsResponse_SdkV2.
func (o *BatchGetListingsResponse_SdkV2) SetListings(ctx context.Context, v []Listing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest_SdkV2 struct {
	Ids types.List `tfsdk:"-"`
}

func (newState *BatchGetProvidersRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetProvidersRequest_SdkV2) {
}

func (newState *BatchGetProvidersRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState BatchGetProvidersRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetProvidersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o BatchGetProvidersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": o.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in BatchGetProvidersRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetProvidersRequest_SdkV2) GetIds(ctx context.Context) ([]types.String, bool) {
	if o.Ids.IsNull() || o.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in BatchGetProvidersRequest_SdkV2.
func (o *BatchGetProvidersRequest_SdkV2) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Ids = types.ListValueMust(t, vs)
}

type BatchGetProvidersResponse_SdkV2 struct {
	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *BatchGetProvidersResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BatchGetProvidersResponse_SdkV2) {
}

func (newState *BatchGetProvidersResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState BatchGetProvidersResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BatchGetProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BatchGetProvidersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o BatchGetProvidersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"providers": o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetProvidersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in BatchGetProvidersResponse_SdkV2 as
// a slice of ProviderInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetProvidersResponse_SdkV2) GetProviders(ctx context.Context) ([]ProviderInfo_SdkV2, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in BatchGetProvidersResponse_SdkV2.
func (o *BatchGetProvidersResponse_SdkV2) SetProviders(ctx context.Context, v []ProviderInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

type ConsumerTerms_SdkV2 struct {
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *ConsumerTerms_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConsumerTerms_SdkV2) {
}

func (newState *ConsumerTerms_SdkV2) SyncEffectiveFieldsDuringRead(existingState ConsumerTerms_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConsumerTerms.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConsumerTerms_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConsumerTerms_SdkV2
// only implements ToObjectValue() and Type().
func (o ConsumerTerms_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConsumerTerms_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.StringType,
		},
	}
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo_SdkV2 struct {
	Company types.String `tfsdk:"company" tf:"optional"`

	Email types.String `tfsdk:"email" tf:"optional"`

	FirstName types.String `tfsdk:"first_name" tf:"optional"`

	LastName types.String `tfsdk:"last_name" tf:"optional"`
}

func (newState *ContactInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContactInfo_SdkV2) {
}

func (newState *ContactInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ContactInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContactInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ContactInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContactInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ContactInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"company":    o.Company,
			"email":      o.Email,
			"first_name": o.FirstName,
			"last_name":  o.LastName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContactInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"company":    types.StringType,
			"email":      types.StringType,
			"first_name": types.StringType,
			"last_name":  types.StringType,
		},
	}
}

type CreateExchangeFilterRequest_SdkV2 struct {
	Filter types.List `tfsdk:"filter" tf:"object"`
}

func (newState *CreateExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeFilterRequest_SdkV2) {
}

func (newState *CreateExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateExchangeFilterRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExchangeFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
		},
	}
}

// GetFilter returns the value of the Filter field in CreateExchangeFilterRequest_SdkV2 as
// a ExchangeFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExchangeFilterRequest_SdkV2) GetFilter(ctx context.Context) (ExchangeFilter_SdkV2, bool) {
	var e ExchangeFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in CreateExchangeFilterRequest_SdkV2.
func (o *CreateExchangeFilterRequest_SdkV2) SetFilter(ctx context.Context, v ExchangeFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type CreateExchangeFilterResponse_SdkV2 struct {
	FilterId types.String `tfsdk:"filter_id" tf:"optional"`
}

func (newState *CreateExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeFilterResponse_SdkV2) {
}

func (newState *CreateExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateExchangeFilterResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeFilterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExchangeFilterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_id": o.FilterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeFilterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_id": types.StringType,
		},
	}
}

type CreateExchangeRequest_SdkV2 struct {
	Exchange types.List `tfsdk:"exchange" tf:"object"`
}

func (newState *CreateExchangeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeRequest_SdkV2) {
}

func (newState *CreateExchangeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateExchangeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExchangeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
		},
	}
}

// GetExchange returns the value of the Exchange field in CreateExchangeRequest_SdkV2 as
// a Exchange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExchangeRequest_SdkV2) GetExchange(ctx context.Context) (Exchange_SdkV2, bool) {
	var e Exchange_SdkV2
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange_SdkV2
	d := o.Exchange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in CreateExchangeRequest_SdkV2.
func (o *CreateExchangeRequest_SdkV2) SetExchange(ctx context.Context, v Exchange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange"]
	o.Exchange = types.ListValueMust(t, vs)
}

type CreateExchangeResponse_SdkV2 struct {
	ExchangeId types.String `tfsdk:"exchange_id" tf:"optional"`
}

func (newState *CreateExchangeResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExchangeResponse_SdkV2) {
}

func (newState *CreateExchangeResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateExchangeResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExchangeResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExchangeResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
		},
	}
}

type CreateFileRequest_SdkV2 struct {
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`

	FileParent types.List `tfsdk:"file_parent" tf:"object"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type" tf:""`

	MimeType types.String `tfsdk:"mime_type" tf:""`
}

func (newState *CreateFileRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFileRequest_SdkV2) {
}

func (newState *CreateFileRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateFileRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFileRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateFileRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          o.DisplayName,
			"file_parent":           o.FileParent,
			"marketplace_file_type": o.MarketplaceFileType,
			"mime_type":             o.MimeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFileRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name": types.StringType,
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.Type(ctx),
			},
			"marketplace_file_type": types.StringType,
			"mime_type":             types.StringType,
		},
	}
}

// GetFileParent returns the value of the FileParent field in CreateFileRequest_SdkV2 as
// a FileParent_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFileRequest_SdkV2) GetFileParent(ctx context.Context) (FileParent_SdkV2, bool) {
	var e FileParent_SdkV2
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent_SdkV2
	d := o.FileParent.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in CreateFileRequest_SdkV2.
func (o *CreateFileRequest_SdkV2) SetFileParent(ctx context.Context, v FileParent_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_parent"]
	o.FileParent = types.ListValueMust(t, vs)
}

type CreateFileResponse_SdkV2 struct {
	FileInfo types.List `tfsdk:"file_info" tf:"optional,object"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url" tf:"optional"`
}

func (newState *CreateFileResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFileResponse_SdkV2) {
}

func (newState *CreateFileResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateFileResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFileResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateFileResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info":  o.FileInfo,
			"upload_url": o.UploadUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFileResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
			},
			"upload_url": types.StringType,
		},
	}
}

// GetFileInfo returns the value of the FileInfo field in CreateFileResponse_SdkV2 as
// a FileInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFileResponse_SdkV2) GetFileInfo(ctx context.Context) (FileInfo_SdkV2, bool) {
	var e FileInfo_SdkV2
	if o.FileInfo.IsNull() || o.FileInfo.IsUnknown() {
		return e, false
	}
	var v []FileInfo_SdkV2
	d := o.FileInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileInfo sets the value of the FileInfo field in CreateFileResponse_SdkV2.
func (o *CreateFileResponse_SdkV2) SetFileInfo(ctx context.Context, v FileInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_info"]
	o.FileInfo = types.ListValueMust(t, vs)
}

type CreateInstallationRequest_SdkV2 struct {
	AcceptedConsumerTerms types.List `tfsdk:"accepted_consumer_terms" tf:"optional,object"`

	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`
	// for git repo installations
	RepoDetail types.List `tfsdk:"repo_detail" tf:"optional,object"`

	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *CreateInstallationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateInstallationRequest_SdkV2) {
}

func (newState *CreateInstallationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateInstallationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateInstallationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms_SdkV2{}),
		"repo_detail":             reflect.TypeOf(RepoInstallation_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstallationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateInstallationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"accepted_consumer_terms": o.AcceptedConsumerTerms,
			"catalog_name":            o.CatalogName,
			"listing_id":              o.ListingId,
			"recipient_type":          o.RecipientType,
			"repo_detail":             o.RepoDetail,
			"share_name":              o.ShareName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstallationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": basetypes.ListType{
				ElemType: ConsumerTerms{}.Type(ctx),
			},
			"catalog_name":   types.StringType,
			"listing_id":     types.StringType,
			"recipient_type": types.StringType,
			"repo_detail": basetypes.ListType{
				ElemType: RepoInstallation{}.Type(ctx),
			},
			"share_name": types.StringType,
		},
	}
}

// GetAcceptedConsumerTerms returns the value of the AcceptedConsumerTerms field in CreateInstallationRequest_SdkV2 as
// a ConsumerTerms_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstallationRequest_SdkV2) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms_SdkV2, bool) {
	var e ConsumerTerms_SdkV2
	if o.AcceptedConsumerTerms.IsNull() || o.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v []ConsumerTerms_SdkV2
	d := o.AcceptedConsumerTerms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreateInstallationRequest_SdkV2.
func (o *CreateInstallationRequest_SdkV2) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["accepted_consumer_terms"]
	o.AcceptedConsumerTerms = types.ListValueMust(t, vs)
}

// GetRepoDetail returns the value of the RepoDetail field in CreateInstallationRequest_SdkV2 as
// a RepoInstallation_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstallationRequest_SdkV2) GetRepoDetail(ctx context.Context) (RepoInstallation_SdkV2, bool) {
	var e RepoInstallation_SdkV2
	if o.RepoDetail.IsNull() || o.RepoDetail.IsUnknown() {
		return e, false
	}
	var v []RepoInstallation_SdkV2
	d := o.RepoDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRepoDetail sets the value of the RepoDetail field in CreateInstallationRequest_SdkV2.
func (o *CreateInstallationRequest_SdkV2) SetRepoDetail(ctx context.Context, v RepoInstallation_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repo_detail"]
	o.RepoDetail = types.ListValueMust(t, vs)
}

type CreateListingRequest_SdkV2 struct {
	Listing types.List `tfsdk:"listing" tf:"object"`
}

func (newState *CreateListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateListingRequest_SdkV2) {
}

func (newState *CreateListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListing returns the value of the Listing field in CreateListingRequest_SdkV2 as
// a Listing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateListingRequest_SdkV2) GetListing(ctx context.Context) (Listing_SdkV2, bool) {
	var e Listing_SdkV2
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing_SdkV2
	d := o.Listing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in CreateListingRequest_SdkV2.
func (o *CreateListingRequest_SdkV2) SetListing(ctx context.Context, v Listing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listing"]
	o.Listing = types.ListValueMust(t, vs)
}

type CreateListingResponse_SdkV2 struct {
	ListingId types.String `tfsdk:"listing_id" tf:"optional"`
}

func (newState *CreateListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateListingResponse_SdkV2) {
}

func (newState *CreateListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
		},
	}
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest_SdkV2 struct {
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

func (newState *CreatePersonalizationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePersonalizationRequest_SdkV2) {
}

func (newState *CreatePersonalizationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreatePersonalizationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePersonalizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePersonalizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePersonalizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"accepted_consumer_terms": o.AcceptedConsumerTerms,
			"comment":                 o.Comment,
			"company":                 o.Company,
			"first_name":              o.FirstName,
			"intended_use":            o.IntendedUse,
			"is_from_lighthouse":      o.IsFromLighthouse,
			"last_name":               o.LastName,
			"listing_id":              o.ListingId,
			"recipient_type":          o.RecipientType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePersonalizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": basetypes.ListType{
				ElemType: ConsumerTerms{}.Type(ctx),
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

// GetAcceptedConsumerTerms returns the value of the AcceptedConsumerTerms field in CreatePersonalizationRequest_SdkV2 as
// a ConsumerTerms_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePersonalizationRequest_SdkV2) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms_SdkV2, bool) {
	var e ConsumerTerms_SdkV2
	if o.AcceptedConsumerTerms.IsNull() || o.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v []ConsumerTerms_SdkV2
	d := o.AcceptedConsumerTerms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreatePersonalizationRequest_SdkV2.
func (o *CreatePersonalizationRequest_SdkV2) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["accepted_consumer_terms"]
	o.AcceptedConsumerTerms = types.ListValueMust(t, vs)
}

type CreatePersonalizationRequestResponse_SdkV2 struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreatePersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePersonalizationRequestResponse_SdkV2) {
}

func (newState *CreatePersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreatePersonalizationRequestResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePersonalizationRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePersonalizationRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePersonalizationRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateProviderRequest_SdkV2 struct {
	Provider types.List `tfsdk:"provider" tf:"object"`
}

func (newState *CreateProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProviderRequest_SdkV2) {
}

func (newState *CreateProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProvider returns the value of the Provider field in CreateProviderRequest_SdkV2 as
// a ProviderInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateProviderRequest_SdkV2) GetProvider(ctx context.Context) (ProviderInfo_SdkV2, bool) {
	var e ProviderInfo_SdkV2
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Provider.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in CreateProviderRequest_SdkV2.
func (o *CreateProviderRequest_SdkV2) SetProvider(ctx context.Context, v ProviderInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider"]
	o.Provider = types.ListValueMust(t, vs)
}

type CreateProviderResponse_SdkV2 struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreateProviderResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProviderResponse_SdkV2) {
}

func (newState *CreateProviderResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateProviderResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateProviderResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DataRefreshInfo_SdkV2 struct {
	Interval types.Int64 `tfsdk:"interval" tf:""`

	Unit types.String `tfsdk:"unit" tf:""`
}

func (newState *DataRefreshInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataRefreshInfo_SdkV2) {
}

func (newState *DataRefreshInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState DataRefreshInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataRefreshInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataRefreshInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataRefreshInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o DataRefreshInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"interval": o.Interval,
			"unit":     o.Unit,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataRefreshInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"interval": types.Int64Type,
			"unit":     types.StringType,
		},
	}
}

// Delete an exchange filter
type DeleteExchangeFilterRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterRequest_SdkV2) {
}

func (newState *DeleteExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeFilterRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExchangeFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeFilterResponse_SdkV2 struct {
}

func (newState *DeleteExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterResponse_SdkV2) {
}

func (newState *DeleteExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeFilterResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeFilterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExchangeFilterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeFilterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an exchange
type DeleteExchangeRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteExchangeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeRequest_SdkV2) {
}

func (newState *DeleteExchangeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExchangeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeResponse_SdkV2 struct {
}

func (newState *DeleteExchangeResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExchangeResponse_SdkV2) {
}

func (newState *DeleteExchangeResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteExchangeResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExchangeResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExchangeResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a file
type DeleteFileRequest_SdkV2 struct {
	FileId types.String `tfsdk:"-"`
}

func (newState *DeleteFileRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileRequest_SdkV2) {
}

func (newState *DeleteFileRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteFileRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFileRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteFileRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": o.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFileRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type DeleteFileResponse_SdkV2 struct {
}

func (newState *DeleteFileResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileResponse_SdkV2) {
}

func (newState *DeleteFileResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteFileResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFileResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteFileResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFileResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Uninstall from a listing
type DeleteInstallationRequest_SdkV2 struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
}

func (newState *DeleteInstallationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstallationRequest_SdkV2) {
}

func (newState *DeleteInstallationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteInstallationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstallationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteInstallationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation_id": o.InstallationId,
			"listing_id":      o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstallationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
		},
	}
}

type DeleteInstallationResponse_SdkV2 struct {
}

func (newState *DeleteInstallationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstallationResponse_SdkV2) {
}

func (newState *DeleteInstallationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteInstallationResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstallationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstallationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteInstallationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstallationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a listing
type DeleteListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteListingRequest_SdkV2) {
}

func (newState *DeleteListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteListingResponse_SdkV2 struct {
}

func (newState *DeleteListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteListingResponse_SdkV2) {
}

func (newState *DeleteListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete provider
type DeleteProviderRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderRequest_SdkV2) {
}

func (newState *DeleteProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteProviderResponse_SdkV2 struct {
}

func (newState *DeleteProviderResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderResponse_SdkV2) {
}

func (newState *DeleteProviderResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteProviderResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteProviderResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Exchange_SdkV2 struct {
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

func (newState *Exchange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Exchange_SdkV2) {
}

func (newState *Exchange_SdkV2) SyncEffectiveFieldsDuringRead(existingState Exchange_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Exchange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Exchange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters":         reflect.TypeOf(ExchangeFilter_SdkV2{}),
		"linked_listings": reflect.TypeOf(ExchangeListing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Exchange_SdkV2
// only implements ToObjectValue() and Type().
func (o Exchange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":         o.Comment,
			"created_at":      o.CreatedAt,
			"created_by":      o.CreatedBy,
			"filters":         o.Filters,
			"id":              o.Id,
			"linked_listings": o.LinkedListings,
			"name":            o.Name,
			"updated_at":      o.UpdatedAt,
			"updated_by":      o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Exchange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
			"id": types.StringType,
			"linked_listings": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
			"name":       types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

// GetFilters returns the value of the Filters field in Exchange_SdkV2 as
// a slice of ExchangeFilter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Exchange_SdkV2) GetFilters(ctx context.Context) ([]ExchangeFilter_SdkV2, bool) {
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter_SdkV2
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in Exchange_SdkV2.
func (o *Exchange_SdkV2) SetFilters(ctx context.Context, v []ExchangeFilter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Filters = types.ListValueMust(t, vs)
}

// GetLinkedListings returns the value of the LinkedListings field in Exchange_SdkV2 as
// a slice of ExchangeListing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Exchange_SdkV2) GetLinkedListings(ctx context.Context) ([]ExchangeListing_SdkV2, bool) {
	if o.LinkedListings.IsNull() || o.LinkedListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing_SdkV2
	d := o.LinkedListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLinkedListings sets the value of the LinkedListings field in Exchange_SdkV2.
func (o *Exchange_SdkV2) SetLinkedListings(ctx context.Context, v []ExchangeListing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["linked_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LinkedListings = types.ListValueMust(t, vs)
}

type ExchangeFilter_SdkV2 struct {
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

func (newState *ExchangeFilter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeFilter_SdkV2) {
}

func (newState *ExchangeFilter_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExchangeFilter_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeFilter_SdkV2
// only implements ToObjectValue() and Type().
func (o ExchangeFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":   o.CreatedAt,
			"created_by":   o.CreatedBy,
			"exchange_id":  o.ExchangeId,
			"filter_type":  o.FilterType,
			"filter_value": o.FilterValue,
			"id":           o.Id,
			"name":         o.Name,
			"updated_at":   o.UpdatedAt,
			"updated_by":   o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeFilter_SdkV2) Type(ctx context.Context) attr.Type {
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

type ExchangeListing_SdkV2 struct {
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`

	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	ExchangeId types.String `tfsdk:"exchange_id" tf:"optional"`

	ExchangeName types.String `tfsdk:"exchange_name" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:"optional"`

	ListingName types.String `tfsdk:"listing_name" tf:"optional"`
}

func (newState *ExchangeListing_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeListing_SdkV2) {
}

func (newState *ExchangeListing_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExchangeListing_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeListing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeListing_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeListing_SdkV2
// only implements ToObjectValue() and Type().
func (o ExchangeListing_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":    o.CreatedAt,
			"created_by":    o.CreatedBy,
			"exchange_id":   o.ExchangeId,
			"exchange_name": o.ExchangeName,
			"id":            o.Id,
			"listing_id":    o.ListingId,
			"listing_name":  o.ListingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeListing_SdkV2) Type(ctx context.Context) attr.Type {
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

type FileInfo_SdkV2 struct {
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

func (newState *FileInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo_SdkV2) {
}

func (newState *FileInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState FileInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o FileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":            o.CreatedAt,
			"display_name":          o.DisplayName,
			"download_link":         o.DownloadLink,
			"file_parent":           o.FileParent,
			"id":                    o.Id,
			"marketplace_file_type": o.MarketplaceFileType,
			"mime_type":             o.MimeType,
			"status":                o.Status,
			"status_message":        o.StatusMessage,
			"updated_at":            o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":    types.Int64Type,
			"display_name":  types.StringType,
			"download_link": types.StringType,
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.Type(ctx),
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

// GetFileParent returns the value of the FileParent field in FileInfo_SdkV2 as
// a FileParent_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileInfo_SdkV2) GetFileParent(ctx context.Context) (FileParent_SdkV2, bool) {
	var e FileParent_SdkV2
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent_SdkV2
	d := o.FileParent.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in FileInfo_SdkV2.
func (o *FileInfo_SdkV2) SetFileParent(ctx context.Context, v FileParent_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_parent"]
	o.FileParent = types.ListValueMust(t, vs)
}

type FileParent_SdkV2 struct {
	FileParentType types.String `tfsdk:"file_parent_type" tf:"optional"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id" tf:"optional"`
}

func (newState *FileParent_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileParent_SdkV2) {
}

func (newState *FileParent_SdkV2) SyncEffectiveFieldsDuringRead(existingState FileParent_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileParent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileParent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileParent_SdkV2
// only implements ToObjectValue() and Type().
func (o FileParent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent_type": o.FileParentType,
			"parent_id":        o.ParentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileParent_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent_type": types.StringType,
			"parent_id":        types.StringType,
		},
	}
}

// Get an exchange
type GetExchangeRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetExchangeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExchangeRequest_SdkV2) {
}

func (newState *GetExchangeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetExchangeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExchangeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExchangeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExchangeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetExchangeResponse_SdkV2 struct {
	Exchange types.List `tfsdk:"exchange" tf:"optional,object"`
}

func (newState *GetExchangeResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExchangeResponse_SdkV2) {
}

func (newState *GetExchangeResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetExchangeResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExchangeResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExchangeResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExchangeResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
		},
	}
}

// GetExchange returns the value of the Exchange field in GetExchangeResponse_SdkV2 as
// a Exchange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetExchangeResponse_SdkV2) GetExchange(ctx context.Context) (Exchange_SdkV2, bool) {
	var e Exchange_SdkV2
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange_SdkV2
	d := o.Exchange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in GetExchangeResponse_SdkV2.
func (o *GetExchangeResponse_SdkV2) SetExchange(ctx context.Context, v Exchange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange"]
	o.Exchange = types.ListValueMust(t, vs)
}

// Get a file
type GetFileRequest_SdkV2 struct {
	FileId types.String `tfsdk:"-"`
}

func (newState *GetFileRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFileRequest_SdkV2) {
}

func (newState *GetFileRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetFileRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFileRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetFileRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": o.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFileRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type GetFileResponse_SdkV2 struct {
	FileInfo types.List `tfsdk:"file_info" tf:"optional,object"`
}

func (newState *GetFileResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFileResponse_SdkV2) {
}

func (newState *GetFileResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetFileResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFileResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetFileResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info": o.FileInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFileResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
			},
		},
	}
}

// GetFileInfo returns the value of the FileInfo field in GetFileResponse_SdkV2 as
// a FileInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetFileResponse_SdkV2) GetFileInfo(ctx context.Context) (FileInfo_SdkV2, bool) {
	var e FileInfo_SdkV2
	if o.FileInfo.IsNull() || o.FileInfo.IsUnknown() {
		return e, false
	}
	var v []FileInfo_SdkV2
	d := o.FileInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileInfo sets the value of the FileInfo field in GetFileResponse_SdkV2.
func (o *GetFileResponse_SdkV2) SetFileInfo(ctx context.Context, v FileInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_info"]
	o.FileInfo = types.ListValueMust(t, vs)
}

type GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2 struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) {
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.Int64Type,
		},
	}
}

// Get listing content metadata
type GetListingContentMetadataRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *GetListingContentMetadataRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingContentMetadataRequest_SdkV2) {
}

func (newState *GetListingContentMetadataRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingContentMetadataRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingContentMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingContentMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingContentMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingContentMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GetListingContentMetadataResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	SharedDataObjects types.List `tfsdk:"shared_data_objects" tf:"optional"`
}

func (newState *GetListingContentMetadataResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingContentMetadataResponse_SdkV2) {
}

func (newState *GetListingContentMetadataResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingContentMetadataResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingContentMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingContentMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shared_data_objects": reflect.TypeOf(SharedDataObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingContentMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     o.NextPageToken,
			"shared_data_objects": o.SharedDataObjects,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingContentMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shared_data_objects": basetypes.ListType{
				ElemType: SharedDataObject{}.Type(ctx),
			},
		},
	}
}

// GetSharedDataObjects returns the value of the SharedDataObjects field in GetListingContentMetadataResponse_SdkV2 as
// a slice of SharedDataObject_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingContentMetadataResponse_SdkV2) GetSharedDataObjects(ctx context.Context) ([]SharedDataObject_SdkV2, bool) {
	if o.SharedDataObjects.IsNull() || o.SharedDataObjects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject_SdkV2
	d := o.SharedDataObjects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSharedDataObjects sets the value of the SharedDataObjects field in GetListingContentMetadataResponse_SdkV2.
func (o *GetListingContentMetadataResponse_SdkV2) SetSharedDataObjects(ctx context.Context, v []SharedDataObject_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shared_data_objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SharedDataObjects = types.ListValueMust(t, vs)
}

// Get listing
type GetListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingRequest_SdkV2) {
}

func (newState *GetListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetListingResponse_SdkV2 struct {
	Listing types.List `tfsdk:"listing" tf:"optional,object"`
}

func (newState *GetListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingResponse_SdkV2) {
}

func (newState *GetListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListing returns the value of the Listing field in GetListingResponse_SdkV2 as
// a Listing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingResponse_SdkV2) GetListing(ctx context.Context) (Listing_SdkV2, bool) {
	var e Listing_SdkV2
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing_SdkV2
	d := o.Listing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in GetListingResponse_SdkV2.
func (o *GetListingResponse_SdkV2) SetListing(ctx context.Context, v Listing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listing"]
	o.Listing = types.ListValueMust(t, vs)
}

// List listings
type GetListingsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *GetListingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingsRequest_SdkV2) {
}

func (newState *GetListingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GetListingsResponse_SdkV2 struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetListingsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetListingsResponse_SdkV2) {
}

func (newState *GetListingsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetListingsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetListingsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetListingsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in GetListingsResponse_SdkV2 as
// a slice of Listing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingsResponse_SdkV2) GetListings(ctx context.Context) ([]Listing_SdkV2, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing_SdkV2
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in GetListingsResponse_SdkV2.
func (o *GetListingsResponse_SdkV2) SetListings(ctx context.Context, v []Listing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`
}

func (newState *GetPersonalizationRequestRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPersonalizationRequestRequest_SdkV2) {
}

func (newState *GetPersonalizationRequestRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPersonalizationRequestRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalizationRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalizationRequestRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPersonalizationRequestRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalizationRequestRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
		},
	}
}

type GetPersonalizationRequestResponse_SdkV2 struct {
	PersonalizationRequests types.List `tfsdk:"personalization_requests" tf:"optional"`
}

func (newState *GetPersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPersonalizationRequestResponse_SdkV2) {
}

func (newState *GetPersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPersonalizationRequestResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalizationRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPersonalizationRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"personalization_requests": o.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalizationRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.Type(ctx),
			},
		},
	}
}

// GetPersonalizationRequests returns the value of the PersonalizationRequests field in GetPersonalizationRequestResponse_SdkV2 as
// a slice of PersonalizationRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPersonalizationRequestResponse_SdkV2) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest_SdkV2, bool) {
	if o.PersonalizationRequests.IsNull() || o.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest_SdkV2
	d := o.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in GetPersonalizationRequestResponse_SdkV2.
func (o *GetPersonalizationRequestResponse_SdkV2) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PersonalizationRequests = types.ListValueMust(t, vs)
}

// Get a provider
type GetProviderRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderRequest_SdkV2) {
}

func (newState *GetProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetProviderResponse_SdkV2 struct {
	Provider types.List `tfsdk:"provider" tf:"optional,object"`
}

func (newState *GetProviderResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderResponse_SdkV2) {
}

func (newState *GetProviderResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetProviderResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetProviderResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProvider returns the value of the Provider field in GetProviderResponse_SdkV2 as
// a ProviderInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetProviderResponse_SdkV2) GetProvider(ctx context.Context) (ProviderInfo_SdkV2, bool) {
	var e ProviderInfo_SdkV2
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Provider.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in GetProviderResponse_SdkV2.
func (o *GetProviderResponse_SdkV2) SetProvider(ctx context.Context, v ProviderInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider"]
	o.Provider = types.ListValueMust(t, vs)
}

type Installation_SdkV2 struct {
	Installation types.List `tfsdk:"installation" tf:"optional,object"`
}

func (newState *Installation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Installation_SdkV2) {
}

func (newState *Installation_SdkV2) SyncEffectiveFieldsDuringRead(existingState Installation_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Installation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Installation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Installation_SdkV2
// only implements ToObjectValue() and Type().
func (o Installation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": o.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Installation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
		},
	}
}

// GetInstallation returns the value of the Installation field in Installation_SdkV2 as
// a InstallationDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Installation_SdkV2) GetInstallation(ctx context.Context) (InstallationDetail_SdkV2, bool) {
	var e InstallationDetail_SdkV2
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail_SdkV2
	d := o.Installation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in Installation_SdkV2.
func (o *Installation_SdkV2) SetInstallation(ctx context.Context, v InstallationDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installation"]
	o.Installation = types.ListValueMust(t, vs)
}

type InstallationDetail_SdkV2 struct {
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

func (newState *InstallationDetail_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstallationDetail_SdkV2) {
}

func (newState *InstallationDetail_SdkV2) SyncEffectiveFieldsDuringRead(existingState InstallationDetail_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallationDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstallationDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_detail": reflect.TypeOf(TokenDetail_SdkV2{}),
		"tokens":       reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallationDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o InstallationDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   o.CatalogName,
			"error_message":  o.ErrorMessage,
			"id":             o.Id,
			"installed_on":   o.InstalledOn,
			"listing_id":     o.ListingId,
			"listing_name":   o.ListingName,
			"recipient_type": o.RecipientType,
			"repo_name":      o.RepoName,
			"repo_path":      o.RepoPath,
			"share_name":     o.ShareName,
			"status":         o.Status,
			"token_detail":   o.TokenDetail,
			"tokens":         o.Tokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstallationDetail_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: TokenDetail{}.Type(ctx),
			},
			"tokens": basetypes.ListType{
				ElemType: TokenInfo{}.Type(ctx),
			},
		},
	}
}

// GetTokenDetail returns the value of the TokenDetail field in InstallationDetail_SdkV2 as
// a TokenDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstallationDetail_SdkV2) GetTokenDetail(ctx context.Context) (TokenDetail_SdkV2, bool) {
	var e TokenDetail_SdkV2
	if o.TokenDetail.IsNull() || o.TokenDetail.IsUnknown() {
		return e, false
	}
	var v []TokenDetail_SdkV2
	d := o.TokenDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenDetail sets the value of the TokenDetail field in InstallationDetail_SdkV2.
func (o *InstallationDetail_SdkV2) SetTokenDetail(ctx context.Context, v TokenDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_detail"]
	o.TokenDetail = types.ListValueMust(t, vs)
}

// GetTokens returns the value of the Tokens field in InstallationDetail_SdkV2 as
// a slice of TokenInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstallationDetail_SdkV2) GetTokens(ctx context.Context) ([]TokenInfo_SdkV2, bool) {
	if o.Tokens.IsNull() || o.Tokens.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo_SdkV2
	d := o.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in InstallationDetail_SdkV2.
func (o *InstallationDetail_SdkV2) SetTokens(ctx context.Context, v []TokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tokens = types.ListValueMust(t, vs)
}

// List all installations
type ListAllInstallationsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAllInstallationsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllInstallationsRequest_SdkV2) {
}

func (newState *ListAllInstallationsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAllInstallationsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllInstallationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllInstallationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllInstallationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllInstallationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAllInstallationsResponse_SdkV2 struct {
	Installations types.List `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListAllInstallationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllInstallationsResponse_SdkV2) {
}

func (newState *ListAllInstallationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAllInstallationsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllInstallationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllInstallationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllInstallationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   o.Installations,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllInstallationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetInstallations returns the value of the Installations field in ListAllInstallationsResponse_SdkV2 as
// a slice of InstallationDetail_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAllInstallationsResponse_SdkV2) GetInstallations(ctx context.Context) ([]InstallationDetail_SdkV2, bool) {
	if o.Installations.IsNull() || o.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail_SdkV2
	d := o.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListAllInstallationsResponse_SdkV2.
func (o *ListAllInstallationsResponse_SdkV2) SetInstallations(ctx context.Context, v []InstallationDetail_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Installations = types.ListValueMust(t, vs)
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAllPersonalizationRequestsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllPersonalizationRequestsRequest_SdkV2) {
}

func (newState *ListAllPersonalizationRequestsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAllPersonalizationRequestsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllPersonalizationRequestsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllPersonalizationRequestsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllPersonalizationRequestsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllPersonalizationRequestsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAllPersonalizationRequestsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	PersonalizationRequests types.List `tfsdk:"personalization_requests" tf:"optional"`
}

func (newState *ListAllPersonalizationRequestsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllPersonalizationRequestsResponse_SdkV2) {
}

func (newState *ListAllPersonalizationRequestsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAllPersonalizationRequestsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllPersonalizationRequestsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllPersonalizationRequestsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllPersonalizationRequestsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          o.NextPageToken,
			"personalization_requests": o.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllPersonalizationRequestsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.Type(ctx),
			},
		},
	}
}

// GetPersonalizationRequests returns the value of the PersonalizationRequests field in ListAllPersonalizationRequestsResponse_SdkV2 as
// a slice of PersonalizationRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAllPersonalizationRequestsResponse_SdkV2) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest_SdkV2, bool) {
	if o.PersonalizationRequests.IsNull() || o.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest_SdkV2
	d := o.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in ListAllPersonalizationRequestsResponse_SdkV2.
func (o *ListAllPersonalizationRequestsResponse_SdkV2) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PersonalizationRequests = types.ListValueMust(t, vs)
}

// List exchange filters
type ListExchangeFiltersRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangeFiltersRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangeFiltersRequest_SdkV2) {
}

func (newState *ListExchangeFiltersRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangeFiltersRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangeFiltersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangeFiltersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangeFiltersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangeFiltersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListExchangeFiltersResponse_SdkV2 struct {
	Filters types.List `tfsdk:"filters" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangeFiltersResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangeFiltersResponse_SdkV2) {
}

func (newState *ListExchangeFiltersResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangeFiltersResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangeFiltersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangeFiltersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters": reflect.TypeOf(ExchangeFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangeFiltersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filters":         o.Filters,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangeFiltersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFilters returns the value of the Filters field in ListExchangeFiltersResponse_SdkV2 as
// a slice of ExchangeFilter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangeFiltersResponse_SdkV2) GetFilters(ctx context.Context) ([]ExchangeFilter_SdkV2, bool) {
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter_SdkV2
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in ListExchangeFiltersResponse_SdkV2.
func (o *ListExchangeFiltersResponse_SdkV2) SetFilters(ctx context.Context, v []ExchangeFilter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Filters = types.ListValueMust(t, vs)
}

// List exchanges for listing
type ListExchangesForListingRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangesForListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesForListingRequest_SdkV2) {
}

func (newState *ListExchangesForListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangesForListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesForListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangesForListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesForListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListExchangesForListingResponse_SdkV2 struct {
	ExchangeListing types.List `tfsdk:"exchange_listing" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangesForListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesForListingResponse_SdkV2) {
}

func (newState *ListExchangesForListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangesForListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesForListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listing": reflect.TypeOf(ExchangeListing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangesForListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listing": o.ExchangeListing,
			"next_page_token":  o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesForListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listing": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchangeListing returns the value of the ExchangeListing field in ListExchangesForListingResponse_SdkV2 as
// a slice of ExchangeListing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangesForListingResponse_SdkV2) GetExchangeListing(ctx context.Context) ([]ExchangeListing_SdkV2, bool) {
	if o.ExchangeListing.IsNull() || o.ExchangeListing.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing_SdkV2
	d := o.ExchangeListing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListing sets the value of the ExchangeListing field in ListExchangesForListingResponse_SdkV2.
func (o *ListExchangesForListingResponse_SdkV2) SetExchangeListing(ctx context.Context, v []ExchangeListing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listing"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeListing = types.ListValueMust(t, vs)
}

// List exchanges
type ListExchangesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExchangesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesRequest_SdkV2) {
}

func (newState *ListExchangesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListExchangesResponse_SdkV2 struct {
	Exchanges types.List `tfsdk:"exchanges" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExchangesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExchangesResponse_SdkV2) {
}

func (newState *ListExchangesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExchangesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExchangesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExchangesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchanges": reflect.TypeOf(Exchange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExchangesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchanges":       o.Exchanges,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchanges": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchanges returns the value of the Exchanges field in ListExchangesResponse_SdkV2 as
// a slice of Exchange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangesResponse_SdkV2) GetExchanges(ctx context.Context) ([]Exchange_SdkV2, bool) {
	if o.Exchanges.IsNull() || o.Exchanges.IsUnknown() {
		return nil, false
	}
	var v []Exchange_SdkV2
	d := o.Exchanges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchanges sets the value of the Exchanges field in ListExchangesResponse_SdkV2.
func (o *ListExchangesResponse_SdkV2) SetExchanges(ctx context.Context, v []Exchange_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchanges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exchanges = types.ListValueMust(t, vs)
}

// List files
type ListFilesRequest_SdkV2 struct {
	FileParent types.List `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListFilesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFilesRequest_SdkV2) {
}

func (newState *ListFilesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFilesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFilesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFilesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent": o.FileParent,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFilesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent": basetypes.ListType{
				ElemType: FileParent{}.Type(ctx),
			},
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// GetFileParent returns the value of the FileParent field in ListFilesRequest_SdkV2 as
// a FileParent_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFilesRequest_SdkV2) GetFileParent(ctx context.Context) (FileParent_SdkV2, bool) {
	var e FileParent_SdkV2
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent_SdkV2
	d := o.FileParent.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in ListFilesRequest_SdkV2.
func (o *ListFilesRequest_SdkV2) SetFileParent(ctx context.Context, v FileParent_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_parent"]
	o.FileParent = types.ListValueMust(t, vs)
}

type ListFilesResponse_SdkV2 struct {
	FileInfos types.List `tfsdk:"file_infos" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFilesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFilesResponse_SdkV2) {
}

func (newState *ListFilesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFilesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFilesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_infos": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFilesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_infos":      o.FileInfos,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFilesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_infos": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFileInfos returns the value of the FileInfos field in ListFilesResponse_SdkV2 as
// a slice of FileInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFilesResponse_SdkV2) GetFileInfos(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if o.FileInfos.IsNull() || o.FileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := o.FileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileInfos sets the value of the FileInfos field in ListFilesResponse_SdkV2.
func (o *ListFilesResponse_SdkV2) SetFileInfos(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileInfos = types.ListValueMust(t, vs)
}

// List all listing fulfillments
type ListFulfillmentsRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListFulfillmentsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFulfillmentsRequest_SdkV2) {
}

func (newState *ListFulfillmentsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFulfillmentsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFulfillmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFulfillmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFulfillmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFulfillmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListFulfillmentsResponse_SdkV2 struct {
	Fulfillments types.List `tfsdk:"fulfillments" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFulfillmentsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFulfillmentsResponse_SdkV2) {
}

func (newState *ListFulfillmentsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFulfillmentsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFulfillmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFulfillmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fulfillments": reflect.TypeOf(ListingFulfillment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFulfillmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fulfillments":    o.Fulfillments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFulfillmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillments": basetypes.ListType{
				ElemType: ListingFulfillment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFulfillments returns the value of the Fulfillments field in ListFulfillmentsResponse_SdkV2 as
// a slice of ListingFulfillment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFulfillmentsResponse_SdkV2) GetFulfillments(ctx context.Context) ([]ListingFulfillment_SdkV2, bool) {
	if o.Fulfillments.IsNull() || o.Fulfillments.IsUnknown() {
		return nil, false
	}
	var v []ListingFulfillment_SdkV2
	d := o.Fulfillments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFulfillments sets the value of the Fulfillments field in ListFulfillmentsResponse_SdkV2.
func (o *ListFulfillmentsResponse_SdkV2) SetFulfillments(ctx context.Context, v []ListingFulfillment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fulfillments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Fulfillments = types.ListValueMust(t, vs)
}

// List installations for a listing
type ListInstallationsRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListInstallationsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstallationsRequest_SdkV2) {
}

func (newState *ListInstallationsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListInstallationsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstallationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstallationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstallationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstallationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListInstallationsResponse_SdkV2 struct {
	Installations types.List `tfsdk:"installations" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListInstallationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstallationsResponse_SdkV2) {
}

func (newState *ListInstallationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListInstallationsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstallationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstallationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstallationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   o.Installations,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstallationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetInstallations returns the value of the Installations field in ListInstallationsResponse_SdkV2 as
// a slice of InstallationDetail_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListInstallationsResponse_SdkV2) GetInstallations(ctx context.Context) ([]InstallationDetail_SdkV2, bool) {
	if o.Installations.IsNull() || o.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail_SdkV2
	d := o.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListInstallationsResponse_SdkV2.
func (o *ListInstallationsResponse_SdkV2) SetInstallations(ctx context.Context, v []InstallationDetail_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Installations = types.ListValueMust(t, vs)
}

// List listings for exchange
type ListListingsForExchangeRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListListingsForExchangeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsForExchangeRequest_SdkV2) {
}

func (newState *ListListingsForExchangeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListListingsForExchangeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsForExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsForExchangeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListListingsForExchangeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsForExchangeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListListingsForExchangeResponse_SdkV2 struct {
	ExchangeListings types.List `tfsdk:"exchange_listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListListingsForExchangeResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsForExchangeResponse_SdkV2) {
}

func (newState *ListListingsForExchangeResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListListingsForExchangeResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsForExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsForExchangeResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listings": reflect.TypeOf(ExchangeListing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListListingsForExchangeResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listings": o.ExchangeListings,
			"next_page_token":   o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsForExchangeResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listings": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchangeListings returns the value of the ExchangeListings field in ListListingsForExchangeResponse_SdkV2 as
// a slice of ExchangeListing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsForExchangeResponse_SdkV2) GetExchangeListings(ctx context.Context) ([]ExchangeListing_SdkV2, bool) {
	if o.ExchangeListings.IsNull() || o.ExchangeListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing_SdkV2
	d := o.ExchangeListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListings sets the value of the ExchangeListings field in ListListingsForExchangeResponse_SdkV2.
func (o *ListListingsForExchangeResponse_SdkV2) SetExchangeListings(ctx context.Context, v []ExchangeListing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeListings = types.ListValueMust(t, vs)
}

// List listings
type ListListingsRequest_SdkV2 struct {
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

func (newState *ListListingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsRequest_SdkV2) {
}

func (newState *ListListingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListListingsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":       reflect.TypeOf(types.String{}),
		"categories":   reflect.TypeOf(types.String{}),
		"provider_ids": reflect.TypeOf(types.String{}),
		"tags":         reflect.TypeOf(ListingTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListListingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":              o.Assets,
			"categories":          o.Categories,
			"is_free":             o.IsFree,
			"is_private_exchange": o.IsPrivateExchange,
			"is_staff_pick":       o.IsStaffPick,
			"page_size":           o.PageSize,
			"page_token":          o.PageToken,
			"provider_ids":        o.ProviderIds,
			"tags":                o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: ListingTag{}.Type(ctx),
			},
		},
	}
}

// GetAssets returns the value of the Assets field in ListListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest_SdkV2) GetAssets(ctx context.Context) ([]types.String, bool) {
	if o.Assets.IsNull() || o.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListListingsRequest_SdkV2.
func (o *ListListingsRequest_SdkV2) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in ListListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest_SdkV2) GetCategories(ctx context.Context) ([]types.String, bool) {
	if o.Categories.IsNull() || o.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in ListListingsRequest_SdkV2.
func (o *ListListingsRequest_SdkV2) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in ListListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest_SdkV2) GetProviderIds(ctx context.Context) ([]types.String, bool) {
	if o.ProviderIds.IsNull() || o.ProviderIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ProviderIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviderIds sets the value of the ProviderIds field in ListListingsRequest_SdkV2.
func (o *ListListingsRequest_SdkV2) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ProviderIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListListingsRequest_SdkV2 as
// a slice of ListingTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest_SdkV2) GetTags(ctx context.Context) ([]ListingTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListListingsRequest_SdkV2.
func (o *ListListingsRequest_SdkV2) SetTags(ctx context.Context, v []ListingTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ListListingsResponse_SdkV2 struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListListingsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListListingsResponse_SdkV2) {
}

func (newState *ListListingsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListListingsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListListingsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListListingsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in ListListingsResponse_SdkV2 as
// a slice of Listing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsResponse_SdkV2) GetListings(ctx context.Context) ([]Listing_SdkV2, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing_SdkV2
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in ListListingsResponse_SdkV2.
func (o *ListListingsResponse_SdkV2) SetListings(ctx context.Context, v []Listing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type ListProviderAnalyticsDashboardResponse_SdkV2 struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`

	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *ListProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderAnalyticsDashboardResponse_SdkV2) {
}

func (newState *ListProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProviderAnalyticsDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderAnalyticsDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProviderAnalyticsDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"id":           o.Id,
			"version":      o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderAnalyticsDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
		},
	}
}

// List providers
type ListProvidersRequest_SdkV2 struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListProvidersRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersRequest_SdkV2) {
}

func (newState *ListProvidersRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProvidersRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProvidersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_featured": o.IsFeatured,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_featured": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListProvidersResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse_SdkV2) {
}

func (newState *ListProvidersResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProvidersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"providers":       o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in ListProvidersResponse_SdkV2 as
// a slice of ProviderInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProvidersResponse_SdkV2) GetProviders(ctx context.Context) ([]ProviderInfo_SdkV2, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse_SdkV2.
func (o *ListProvidersResponse_SdkV2) SetProviders(ctx context.Context, v []ProviderInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

type Listing_SdkV2 struct {
	Detail types.List `tfsdk:"detail" tf:"optional,object"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// Next Number: 26
	Summary types.List `tfsdk:"summary" tf:"object"`
}

func (newState *Listing_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Listing_SdkV2) {
}

func (newState *Listing_SdkV2) SyncEffectiveFieldsDuringRead(existingState Listing_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Listing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Listing_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"detail":  reflect.TypeOf(ListingDetail_SdkV2{}),
		"summary": reflect.TypeOf(ListingSummary_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Listing_SdkV2
// only implements ToObjectValue() and Type().
func (o Listing_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"detail":  o.Detail,
			"id":      o.Id,
			"summary": o.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Listing_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"detail": basetypes.ListType{
				ElemType: ListingDetail{}.Type(ctx),
			},
			"id": types.StringType,
			"summary": basetypes.ListType{
				ElemType: ListingSummary{}.Type(ctx),
			},
		},
	}
}

// GetDetail returns the value of the Detail field in Listing_SdkV2 as
// a ListingDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Listing_SdkV2) GetDetail(ctx context.Context) (ListingDetail_SdkV2, bool) {
	var e ListingDetail_SdkV2
	if o.Detail.IsNull() || o.Detail.IsUnknown() {
		return e, false
	}
	var v []ListingDetail_SdkV2
	d := o.Detail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDetail sets the value of the Detail field in Listing_SdkV2.
func (o *Listing_SdkV2) SetDetail(ctx context.Context, v ListingDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["detail"]
	o.Detail = types.ListValueMust(t, vs)
}

// GetSummary returns the value of the Summary field in Listing_SdkV2 as
// a ListingSummary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Listing_SdkV2) GetSummary(ctx context.Context) (ListingSummary_SdkV2, bool) {
	var e ListingSummary_SdkV2
	if o.Summary.IsNull() || o.Summary.IsUnknown() {
		return e, false
	}
	var v []ListingSummary_SdkV2
	d := o.Summary.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSummary sets the value of the Summary field in Listing_SdkV2.
func (o *Listing_SdkV2) SetSummary(ctx context.Context, v ListingSummary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["summary"]
	o.Summary = types.ListValueMust(t, vs)
}

type ListingDetail_SdkV2 struct {
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

func (newState *ListingDetail_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingDetail_SdkV2) {
}

func (newState *ListingDetail_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListingDetail_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":                       reflect.TypeOf(types.String{}),
		"collection_granularity":       reflect.TypeOf(DataRefreshInfo_SdkV2{}),
		"embedded_notebook_file_infos": reflect.TypeOf(FileInfo_SdkV2{}),
		"file_ids":                     reflect.TypeOf(types.String{}),
		"tags":                         reflect.TypeOf(ListingTag_SdkV2{}),
		"update_frequency":             reflect.TypeOf(DataRefreshInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o ListingDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":                       o.Assets,
			"collection_date_end":          o.CollectionDateEnd,
			"collection_date_start":        o.CollectionDateStart,
			"collection_granularity":       o.CollectionGranularity,
			"cost":                         o.Cost,
			"data_source":                  o.DataSource,
			"description":                  o.Description,
			"documentation_link":           o.DocumentationLink,
			"embedded_notebook_file_infos": o.EmbeddedNotebookFileInfos,
			"file_ids":                     o.FileIds,
			"geographical_coverage":        o.GeographicalCoverage,
			"license":                      o.License,
			"pricing_model":                o.PricingModel,
			"privacy_policy_link":          o.PrivacyPolicyLink,
			"size":                         o.Size,
			"support_link":                 o.SupportLink,
			"tags":                         o.Tags,
			"terms_of_service":             o.TermsOfService,
			"update_frequency":             o.UpdateFrequency,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"collection_date_end":   types.Int64Type,
			"collection_date_start": types.Int64Type,
			"collection_granularity": basetypes.ListType{
				ElemType: DataRefreshInfo{}.Type(ctx),
			},
			"cost":               types.StringType,
			"data_source":        types.StringType,
			"description":        types.StringType,
			"documentation_link": types.StringType,
			"embedded_notebook_file_infos": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
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
				ElemType: ListingTag{}.Type(ctx),
			},
			"terms_of_service": types.StringType,
			"update_frequency": basetypes.ListType{
				ElemType: DataRefreshInfo{}.Type(ctx),
			},
		},
	}
}

// GetAssets returns the value of the Assets field in ListingDetail_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetAssets(ctx context.Context) ([]types.String, bool) {
	if o.Assets.IsNull() || o.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCollectionGranularity returns the value of the CollectionGranularity field in ListingDetail_SdkV2 as
// a DataRefreshInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetCollectionGranularity(ctx context.Context) (DataRefreshInfo_SdkV2, bool) {
	var e DataRefreshInfo_SdkV2
	if o.CollectionGranularity.IsNull() || o.CollectionGranularity.IsUnknown() {
		return e, false
	}
	var v []DataRefreshInfo_SdkV2
	d := o.CollectionGranularity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCollectionGranularity sets the value of the CollectionGranularity field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetCollectionGranularity(ctx context.Context, v DataRefreshInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["collection_granularity"]
	o.CollectionGranularity = types.ListValueMust(t, vs)
}

// GetEmbeddedNotebookFileInfos returns the value of the EmbeddedNotebookFileInfos field in ListingDetail_SdkV2 as
// a slice of FileInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetEmbeddedNotebookFileInfos(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if o.EmbeddedNotebookFileInfos.IsNull() || o.EmbeddedNotebookFileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := o.EmbeddedNotebookFileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddedNotebookFileInfos sets the value of the EmbeddedNotebookFileInfos field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetEmbeddedNotebookFileInfos(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedded_notebook_file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddedNotebookFileInfos = types.ListValueMust(t, vs)
}

// GetFileIds returns the value of the FileIds field in ListingDetail_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetFileIds(ctx context.Context) ([]types.String, bool) {
	if o.FileIds.IsNull() || o.FileIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FileIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileIds sets the value of the FileIds field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetFileIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListingDetail_SdkV2 as
// a slice of ListingTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetTags(ctx context.Context) ([]ListingTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetTags(ctx context.Context, v []ListingTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetUpdateFrequency returns the value of the UpdateFrequency field in ListingDetail_SdkV2 as
// a DataRefreshInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail_SdkV2) GetUpdateFrequency(ctx context.Context) (DataRefreshInfo_SdkV2, bool) {
	var e DataRefreshInfo_SdkV2
	if o.UpdateFrequency.IsNull() || o.UpdateFrequency.IsUnknown() {
		return e, false
	}
	var v []DataRefreshInfo_SdkV2
	d := o.UpdateFrequency.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdateFrequency sets the value of the UpdateFrequency field in ListingDetail_SdkV2.
func (o *ListingDetail_SdkV2) SetUpdateFrequency(ctx context.Context, v DataRefreshInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["update_frequency"]
	o.UpdateFrequency = types.ListValueMust(t, vs)
}

type ListingFulfillment_SdkV2 struct {
	FulfillmentType types.String `tfsdk:"fulfillment_type" tf:"optional"`

	ListingId types.String `tfsdk:"listing_id" tf:""`

	RecipientType types.String `tfsdk:"recipient_type" tf:"optional"`

	RepoInfo types.List `tfsdk:"repo_info" tf:"optional,object"`

	ShareInfo types.List `tfsdk:"share_info" tf:"optional,object"`
}

func (newState *ListingFulfillment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingFulfillment_SdkV2) {
}

func (newState *ListingFulfillment_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListingFulfillment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingFulfillment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingFulfillment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repo_info":  reflect.TypeOf(RepoInfo_SdkV2{}),
		"share_info": reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingFulfillment_SdkV2
// only implements ToObjectValue() and Type().
func (o ListingFulfillment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fulfillment_type": o.FulfillmentType,
			"listing_id":       o.ListingId,
			"recipient_type":   o.RecipientType,
			"repo_info":        o.RepoInfo,
			"share_info":       o.ShareInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingFulfillment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillment_type": types.StringType,
			"listing_id":       types.StringType,
			"recipient_type":   types.StringType,
			"repo_info": basetypes.ListType{
				ElemType: RepoInfo{}.Type(ctx),
			},
			"share_info": basetypes.ListType{
				ElemType: ShareInfo{}.Type(ctx),
			},
		},
	}
}

// GetRepoInfo returns the value of the RepoInfo field in ListingFulfillment_SdkV2 as
// a RepoInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingFulfillment_SdkV2) GetRepoInfo(ctx context.Context) (RepoInfo_SdkV2, bool) {
	var e RepoInfo_SdkV2
	if o.RepoInfo.IsNull() || o.RepoInfo.IsUnknown() {
		return e, false
	}
	var v []RepoInfo_SdkV2
	d := o.RepoInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRepoInfo sets the value of the RepoInfo field in ListingFulfillment_SdkV2.
func (o *ListingFulfillment_SdkV2) SetRepoInfo(ctx context.Context, v RepoInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repo_info"]
	o.RepoInfo = types.ListValueMust(t, vs)
}

// GetShareInfo returns the value of the ShareInfo field in ListingFulfillment_SdkV2 as
// a ShareInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingFulfillment_SdkV2) GetShareInfo(ctx context.Context) (ShareInfo_SdkV2, bool) {
	var e ShareInfo_SdkV2
	if o.ShareInfo.IsNull() || o.ShareInfo.IsUnknown() {
		return e, false
	}
	var v []ShareInfo_SdkV2
	d := o.ShareInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShareInfo sets the value of the ShareInfo field in ListingFulfillment_SdkV2.
func (o *ListingFulfillment_SdkV2) SetShareInfo(ctx context.Context, v ShareInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["share_info"]
	o.ShareInfo = types.ListValueMust(t, vs)
}

type ListingSetting_SdkV2 struct {
	Visibility types.String `tfsdk:"visibility" tf:"optional"`
}

func (newState *ListingSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingSetting_SdkV2) {
}

func (newState *ListingSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListingSetting_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o ListingSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visibility": o.Visibility,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"visibility": types.StringType,
		},
	}
}

// Next Number: 26
type ListingSummary_SdkV2 struct {
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

func (newState *ListingSummary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingSummary_SdkV2) {
}

func (newState *ListingSummary_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListingSummary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingSummary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"categories":      reflect.TypeOf(types.String{}),
		"exchange_ids":    reflect.TypeOf(types.String{}),
		"git_repo":        reflect.TypeOf(RepoInfo_SdkV2{}),
		"provider_region": reflect.TypeOf(RegionInfo_SdkV2{}),
		"setting":         reflect.TypeOf(ListingSetting_SdkV2{}),
		"share":           reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingSummary_SdkV2
// only implements ToObjectValue() and Type().
func (o ListingSummary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"categories":      o.Categories,
			"created_at":      o.CreatedAt,
			"created_by":      o.CreatedBy,
			"created_by_id":   o.CreatedById,
			"exchange_ids":    o.ExchangeIds,
			"git_repo":        o.GitRepo,
			"listingType":     o.ListingType,
			"name":            o.Name,
			"provider_id":     o.ProviderId,
			"provider_region": o.ProviderRegion,
			"published_at":    o.PublishedAt,
			"published_by":    o.PublishedBy,
			"setting":         o.Setting,
			"share":           o.Share,
			"status":          o.Status,
			"subtitle":        o.Subtitle,
			"updated_at":      o.UpdatedAt,
			"updated_by":      o.UpdatedBy,
			"updated_by_id":   o.UpdatedById,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingSummary_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: RepoInfo{}.Type(ctx),
			},
			"listingType": types.StringType,
			"name":        types.StringType,
			"provider_id": types.StringType,
			"provider_region": basetypes.ListType{
				ElemType: RegionInfo{}.Type(ctx),
			},
			"published_at": types.Int64Type,
			"published_by": types.StringType,
			"setting": basetypes.ListType{
				ElemType: ListingSetting{}.Type(ctx),
			},
			"share": basetypes.ListType{
				ElemType: ShareInfo{}.Type(ctx),
			},
			"status":        types.StringType,
			"subtitle":      types.StringType,
			"updated_at":    types.Int64Type,
			"updated_by":    types.StringType,
			"updated_by_id": types.Int64Type,
		},
	}
}

// GetCategories returns the value of the Categories field in ListingSummary_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetCategories(ctx context.Context) ([]types.String, bool) {
	if o.Categories.IsNull() || o.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetExchangeIds returns the value of the ExchangeIds field in ListingSummary_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetExchangeIds(ctx context.Context) ([]types.String, bool) {
	if o.ExchangeIds.IsNull() || o.ExchangeIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ExchangeIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeIds sets the value of the ExchangeIds field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetExchangeIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeIds = types.ListValueMust(t, vs)
}

// GetGitRepo returns the value of the GitRepo field in ListingSummary_SdkV2 as
// a RepoInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetGitRepo(ctx context.Context) (RepoInfo_SdkV2, bool) {
	var e RepoInfo_SdkV2
	if o.GitRepo.IsNull() || o.GitRepo.IsUnknown() {
		return e, false
	}
	var v []RepoInfo_SdkV2
	d := o.GitRepo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitRepo sets the value of the GitRepo field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetGitRepo(ctx context.Context, v RepoInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_repo"]
	o.GitRepo = types.ListValueMust(t, vs)
}

// GetProviderRegion returns the value of the ProviderRegion field in ListingSummary_SdkV2 as
// a RegionInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetProviderRegion(ctx context.Context) (RegionInfo_SdkV2, bool) {
	var e RegionInfo_SdkV2
	if o.ProviderRegion.IsNull() || o.ProviderRegion.IsUnknown() {
		return e, false
	}
	var v []RegionInfo_SdkV2
	d := o.ProviderRegion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProviderRegion sets the value of the ProviderRegion field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetProviderRegion(ctx context.Context, v RegionInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_region"]
	o.ProviderRegion = types.ListValueMust(t, vs)
}

// GetSetting returns the value of the Setting field in ListingSummary_SdkV2 as
// a ListingSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetSetting(ctx context.Context) (ListingSetting_SdkV2, bool) {
	var e ListingSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []ListingSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetSetting(ctx context.Context, v ListingSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

// GetShare returns the value of the Share field in ListingSummary_SdkV2 as
// a ShareInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary_SdkV2) GetShare(ctx context.Context) (ShareInfo_SdkV2, bool) {
	var e ShareInfo_SdkV2
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo_SdkV2
	d := o.Share.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in ListingSummary_SdkV2.
func (o *ListingSummary_SdkV2) SetShare(ctx context.Context, v ShareInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["share"]
	o.Share = types.ListValueMust(t, vs)
}

type ListingTag_SdkV2 struct {
	// Tag name (enum)
	TagName types.String `tfsdk:"tag_name" tf:"optional"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues types.List `tfsdk:"tag_values" tf:"optional"`
}

func (newState *ListingTag_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListingTag_SdkV2) {
}

func (newState *ListingTag_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListingTag_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListingTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListingTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_values": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingTag_SdkV2
// only implements ToObjectValue() and Type().
func (o ListingTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_name":   o.TagName,
			"tag_values": o.TagValues,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_name": types.StringType,
			"tag_values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTagValues returns the value of the TagValues field in ListingTag_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingTag_SdkV2) GetTagValues(ctx context.Context) ([]types.String, bool) {
	if o.TagValues.IsNull() || o.TagValues.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TagValues.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagValues sets the value of the TagValues field in ListingTag_SdkV2.
func (o *ListingTag_SdkV2) SetTagValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagValues = types.ListValueMust(t, vs)
}

type PersonalizationRequest_SdkV2 struct {
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

func (newState *PersonalizationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalizationRequest_SdkV2) {
}

func (newState *PersonalizationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PersonalizationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PersonalizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"consumer_region": reflect.TypeOf(RegionInfo_SdkV2{}),
		"contact_info":    reflect.TypeOf(ContactInfo_SdkV2{}),
		"share":           reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PersonalizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"consumer_region":    o.ConsumerRegion,
			"contact_info":       o.ContactInfo,
			"created_at":         o.CreatedAt,
			"id":                 o.Id,
			"intended_use":       o.IntendedUse,
			"is_from_lighthouse": o.IsFromLighthouse,
			"listing_id":         o.ListingId,
			"listing_name":       o.ListingName,
			"metastore_id":       o.MetastoreId,
			"provider_id":        o.ProviderId,
			"recipient_type":     o.RecipientType,
			"share":              o.Share,
			"status":             o.Status,
			"status_message":     o.StatusMessage,
			"updated_at":         o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"consumer_region": basetypes.ListType{
				ElemType: RegionInfo{}.Type(ctx),
			},
			"contact_info": basetypes.ListType{
				ElemType: ContactInfo{}.Type(ctx),
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
				ElemType: ShareInfo{}.Type(ctx),
			},
			"status":         types.StringType,
			"status_message": types.StringType,
			"updated_at":     types.Int64Type,
		},
	}
}

// GetConsumerRegion returns the value of the ConsumerRegion field in PersonalizationRequest_SdkV2 as
// a RegionInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest_SdkV2) GetConsumerRegion(ctx context.Context) (RegionInfo_SdkV2, bool) {
	var e RegionInfo_SdkV2
	if o.ConsumerRegion.IsNull() || o.ConsumerRegion.IsUnknown() {
		return e, false
	}
	var v []RegionInfo_SdkV2
	d := o.ConsumerRegion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConsumerRegion sets the value of the ConsumerRegion field in PersonalizationRequest_SdkV2.
func (o *PersonalizationRequest_SdkV2) SetConsumerRegion(ctx context.Context, v RegionInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["consumer_region"]
	o.ConsumerRegion = types.ListValueMust(t, vs)
}

// GetContactInfo returns the value of the ContactInfo field in PersonalizationRequest_SdkV2 as
// a ContactInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest_SdkV2) GetContactInfo(ctx context.Context) (ContactInfo_SdkV2, bool) {
	var e ContactInfo_SdkV2
	if o.ContactInfo.IsNull() || o.ContactInfo.IsUnknown() {
		return e, false
	}
	var v []ContactInfo_SdkV2
	d := o.ContactInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContactInfo sets the value of the ContactInfo field in PersonalizationRequest_SdkV2.
func (o *PersonalizationRequest_SdkV2) SetContactInfo(ctx context.Context, v ContactInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["contact_info"]
	o.ContactInfo = types.ListValueMust(t, vs)
}

// GetShare returns the value of the Share field in PersonalizationRequest_SdkV2 as
// a ShareInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest_SdkV2) GetShare(ctx context.Context) (ShareInfo_SdkV2, bool) {
	var e ShareInfo_SdkV2
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo_SdkV2
	d := o.Share.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in PersonalizationRequest_SdkV2.
func (o *PersonalizationRequest_SdkV2) SetShare(ctx context.Context, v ShareInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["share"]
	o.Share = types.ListValueMust(t, vs)
}

type ProviderAnalyticsDashboard_SdkV2 struct {
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *ProviderAnalyticsDashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderAnalyticsDashboard_SdkV2) {
}

func (newState *ProviderAnalyticsDashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProviderAnalyticsDashboard_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderAnalyticsDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderAnalyticsDashboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderAnalyticsDashboard_SdkV2
// only implements ToObjectValue() and Type().
func (o ProviderAnalyticsDashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderAnalyticsDashboard_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type ProviderInfo_SdkV2 struct {
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

func (newState *ProviderInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo_SdkV2) {
}

func (newState *ProviderInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProviderInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ProviderInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"business_contact_email":   o.BusinessContactEmail,
			"company_website_link":     o.CompanyWebsiteLink,
			"dark_mode_icon_file_id":   o.DarkModeIconFileId,
			"dark_mode_icon_file_path": o.DarkModeIconFilePath,
			"description":              o.Description,
			"icon_file_id":             o.IconFileId,
			"icon_file_path":           o.IconFilePath,
			"id":                       o.Id,
			"is_featured":              o.IsFeatured,
			"name":                     o.Name,
			"privacy_policy_link":      o.PrivacyPolicyLink,
			"published_by":             o.PublishedBy,
			"support_contact_email":    o.SupportContactEmail,
			"term_of_service_link":     o.TermOfServiceLink,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type RegionInfo_SdkV2 struct {
	Cloud types.String `tfsdk:"cloud" tf:"optional"`

	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *RegionInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegionInfo_SdkV2) {
}

func (newState *RegionInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RegionInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegionInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegionInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RegionInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud":  o.Cloud,
			"region": o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegionInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud":  types.StringType,
			"region": types.StringType,
		},
	}
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (newState *RemoveExchangeForListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingRequest_SdkV2) {
}

func (newState *RemoveExchangeForListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState RemoveExchangeForListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveExchangeForListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveExchangeForListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RemoveExchangeForListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveExchangeForListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type RemoveExchangeForListingResponse_SdkV2 struct {
}

func (newState *RemoveExchangeForListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingResponse_SdkV2) {
}

func (newState *RemoveExchangeForListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RemoveExchangeForListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveExchangeForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveExchangeForListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RemoveExchangeForListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveExchangeForListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoInfo_SdkV2 struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url" tf:""`
}

func (newState *RepoInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInfo_SdkV2) {
}

func (newState *RepoInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepoInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_repo_url": o.GitRepoUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"git_repo_url": types.StringType,
		},
	}
}

type RepoInstallation_SdkV2 struct {
	// the user-specified repo name for their installed git repo listing
	RepoName types.String `tfsdk:"repo_name" tf:""`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path" tf:""`
}

func (newState *RepoInstallation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInstallation_SdkV2) {
}

func (newState *RepoInstallation_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepoInstallation_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoInstallation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoInstallation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInstallation_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoInstallation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_name": o.RepoName,
			"repo_path": o.RepoPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoInstallation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_name": types.StringType,
			"repo_path": types.StringType,
		},
	}
}

// Search listings
type SearchListingsRequest_SdkV2 struct {
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

func (newState *SearchListingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchListingsRequest_SdkV2) {
}

func (newState *SearchListingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState SearchListingsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchListingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchListingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":       reflect.TypeOf(types.String{}),
		"categories":   reflect.TypeOf(types.String{}),
		"provider_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchListingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":              o.Assets,
			"categories":          o.Categories,
			"is_free":             o.IsFree,
			"is_private_exchange": o.IsPrivateExchange,
			"page_size":           o.PageSize,
			"page_token":          o.PageToken,
			"provider_ids":        o.ProviderIds,
			"query":               o.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchListingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAssets returns the value of the Assets field in SearchListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest_SdkV2) GetAssets(ctx context.Context) ([]types.String, bool) {
	if o.Assets.IsNull() || o.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in SearchListingsRequest_SdkV2.
func (o *SearchListingsRequest_SdkV2) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in SearchListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest_SdkV2) GetCategories(ctx context.Context) ([]types.String, bool) {
	if o.Categories.IsNull() || o.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in SearchListingsRequest_SdkV2.
func (o *SearchListingsRequest_SdkV2) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in SearchListingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest_SdkV2) GetProviderIds(ctx context.Context) ([]types.String, bool) {
	if o.ProviderIds.IsNull() || o.ProviderIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ProviderIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviderIds sets the value of the ProviderIds field in SearchListingsRequest_SdkV2.
func (o *SearchListingsRequest_SdkV2) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ProviderIds = types.ListValueMust(t, vs)
}

type SearchListingsResponse_SdkV2 struct {
	Listings types.List `tfsdk:"listings" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchListingsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchListingsResponse_SdkV2) {
}

func (newState *SearchListingsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState SearchListingsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchListingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchListingsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchListingsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchListingsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in SearchListingsResponse_SdkV2 as
// a slice of Listing_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsResponse_SdkV2) GetListings(ctx context.Context) ([]Listing_SdkV2, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing_SdkV2
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in SearchListingsResponse_SdkV2.
func (o *SearchListingsResponse_SdkV2) SetListings(ctx context.Context, v []Listing_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type ShareInfo_SdkV2 struct {
	Name types.String `tfsdk:"name" tf:""`

	Type_ types.String `tfsdk:"type" tf:""`
}

func (newState *ShareInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo_SdkV2) {
}

func (newState *ShareInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ShareInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ShareInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ShareInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
			"type": o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ShareInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"type": types.StringType,
		},
	}
}

type SharedDataObject_SdkV2 struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Name of the shared object
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *SharedDataObject_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject_SdkV2) {
}

func (newState *SharedDataObject_SdkV2) SyncEffectiveFieldsDuringRead(existingState SharedDataObject_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharedDataObject_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject_SdkV2
// only implements ToObjectValue() and Type().
func (o SharedDataObject_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_object_type": o.DataObjectType,
			"name":             o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObject_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_object_type": types.StringType,
			"name":             types.StringType,
		},
	}
}

type TokenDetail_SdkV2 struct {
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`

	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`

	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

func (newState *TokenDetail_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenDetail_SdkV2) {
}

func (newState *TokenDetail_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenDetail_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearerToken":             o.BearerToken,
			"endpoint":                o.Endpoint,
			"expirationTime":          o.ExpirationTime,
			"shareCredentialsVersion": o.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearerToken":             types.StringType,
			"endpoint":                types.StringType,
			"expirationTime":          types.StringType,
			"shareCredentialsVersion": types.Int64Type,
		},
	}
}

type TokenInfo_SdkV2 struct {
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

func (newState *TokenInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenInfo_SdkV2) {
}

func (newState *TokenInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url":  o.ActivationUrl,
			"created_at":      o.CreatedAt,
			"created_by":      o.CreatedBy,
			"expiration_time": o.ExpirationTime,
			"id":              o.Id,
			"updated_at":      o.UpdatedAt,
			"updated_by":      o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type UpdateExchangeFilterRequest_SdkV2 struct {
	Filter types.List `tfsdk:"filter" tf:"object"`

	Id types.String `tfsdk:"-"`
}

func (newState *UpdateExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeFilterRequest_SdkV2) {
}

func (newState *UpdateExchangeFilterRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeFilterRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExchangeFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
			"id":     o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
			"id": types.StringType,
		},
	}
}

// GetFilter returns the value of the Filter field in UpdateExchangeFilterRequest_SdkV2 as
// a ExchangeFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeFilterRequest_SdkV2) GetFilter(ctx context.Context) (ExchangeFilter_SdkV2, bool) {
	var e ExchangeFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterRequest_SdkV2.
func (o *UpdateExchangeFilterRequest_SdkV2) SetFilter(ctx context.Context, v ExchangeFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type UpdateExchangeFilterResponse_SdkV2 struct {
	Filter types.List `tfsdk:"filter" tf:"optional,object"`
}

func (newState *UpdateExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeFilterResponse_SdkV2) {
}

func (newState *UpdateExchangeFilterResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeFilterResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeFilterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExchangeFilterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeFilterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
		},
	}
}

// GetFilter returns the value of the Filter field in UpdateExchangeFilterResponse_SdkV2 as
// a ExchangeFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeFilterResponse_SdkV2) GetFilter(ctx context.Context) (ExchangeFilter_SdkV2, bool) {
	var e ExchangeFilter_SdkV2
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter_SdkV2
	d := o.Filter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterResponse_SdkV2.
func (o *UpdateExchangeFilterResponse_SdkV2) SetFilter(ctx context.Context, v ExchangeFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter"]
	o.Filter = types.ListValueMust(t, vs)
}

type UpdateExchangeRequest_SdkV2 struct {
	Exchange types.List `tfsdk:"exchange" tf:"object"`

	Id types.String `tfsdk:"-"`
}

func (newState *UpdateExchangeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeRequest_SdkV2) {
}

func (newState *UpdateExchangeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExchangeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
			"id":       o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
			"id": types.StringType,
		},
	}
}

// GetExchange returns the value of the Exchange field in UpdateExchangeRequest_SdkV2 as
// a Exchange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeRequest_SdkV2) GetExchange(ctx context.Context) (Exchange_SdkV2, bool) {
	var e Exchange_SdkV2
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange_SdkV2
	d := o.Exchange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeRequest_SdkV2.
func (o *UpdateExchangeRequest_SdkV2) SetExchange(ctx context.Context, v Exchange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange"]
	o.Exchange = types.ListValueMust(t, vs)
}

type UpdateExchangeResponse_SdkV2 struct {
	Exchange types.List `tfsdk:"exchange" tf:"optional,object"`
}

func (newState *UpdateExchangeResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExchangeResponse_SdkV2) {
}

func (newState *UpdateExchangeResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateExchangeResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExchangeResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExchangeResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
		},
	}
}

// GetExchange returns the value of the Exchange field in UpdateExchangeResponse_SdkV2 as
// a Exchange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeResponse_SdkV2) GetExchange(ctx context.Context) (Exchange_SdkV2, bool) {
	var e Exchange_SdkV2
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange_SdkV2
	d := o.Exchange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeResponse_SdkV2.
func (o *UpdateExchangeResponse_SdkV2) SetExchange(ctx context.Context, v Exchange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange"]
	o.Exchange = types.ListValueMust(t, vs)
}

type UpdateInstallationRequest_SdkV2 struct {
	Installation types.List `tfsdk:"installation" tf:"object"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token" tf:"optional"`
}

func (newState *UpdateInstallationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInstallationRequest_SdkV2) {
}

func (newState *UpdateInstallationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateInstallationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInstallationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInstallationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateInstallationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation":    o.Installation,
			"installation_id": o.InstallationId,
			"listing_id":      o.ListingId,
			"rotate_token":    o.RotateToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInstallationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
			"rotate_token":    types.BoolType,
		},
	}
}

// GetInstallation returns the value of the Installation field in UpdateInstallationRequest_SdkV2 as
// a InstallationDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInstallationRequest_SdkV2) GetInstallation(ctx context.Context) (InstallationDetail_SdkV2, bool) {
	var e InstallationDetail_SdkV2
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail_SdkV2
	d := o.Installation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationRequest_SdkV2.
func (o *UpdateInstallationRequest_SdkV2) SetInstallation(ctx context.Context, v InstallationDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installation"]
	o.Installation = types.ListValueMust(t, vs)
}

type UpdateInstallationResponse_SdkV2 struct {
	Installation types.List `tfsdk:"installation" tf:"optional,object"`
}

func (newState *UpdateInstallationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInstallationResponse_SdkV2) {
}

func (newState *UpdateInstallationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateInstallationResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInstallationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInstallationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateInstallationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": o.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInstallationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
		},
	}
}

// GetInstallation returns the value of the Installation field in UpdateInstallationResponse_SdkV2 as
// a InstallationDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInstallationResponse_SdkV2) GetInstallation(ctx context.Context) (InstallationDetail_SdkV2, bool) {
	var e InstallationDetail_SdkV2
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail_SdkV2
	d := o.Installation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationResponse_SdkV2.
func (o *UpdateInstallationResponse_SdkV2) SetInstallation(ctx context.Context, v InstallationDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installation"]
	o.Installation = types.ListValueMust(t, vs)
}

type UpdateListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`

	Listing types.List `tfsdk:"listing" tf:"object"`
}

func (newState *UpdateListingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateListingRequest_SdkV2) {
}

func (newState *UpdateListingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateListingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateListingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateListingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateListingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"listing": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListing returns the value of the Listing field in UpdateListingRequest_SdkV2 as
// a Listing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateListingRequest_SdkV2) GetListing(ctx context.Context) (Listing_SdkV2, bool) {
	var e Listing_SdkV2
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing_SdkV2
	d := o.Listing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in UpdateListingRequest_SdkV2.
func (o *UpdateListingRequest_SdkV2) SetListing(ctx context.Context, v Listing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listing"]
	o.Listing = types.ListValueMust(t, vs)
}

type UpdateListingResponse_SdkV2 struct {
	Listing types.List `tfsdk:"listing" tf:"optional,object"`
}

func (newState *UpdateListingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateListingResponse_SdkV2) {
}

func (newState *UpdateListingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateListingResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateListingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateListingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateListingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListing returns the value of the Listing field in UpdateListingResponse_SdkV2 as
// a Listing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateListingResponse_SdkV2) GetListing(ctx context.Context) (Listing_SdkV2, bool) {
	var e Listing_SdkV2
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing_SdkV2
	d := o.Listing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in UpdateListingResponse_SdkV2.
func (o *UpdateListingResponse_SdkV2) SetListing(ctx context.Context, v Listing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listing"]
	o.Listing = types.ListValueMust(t, vs)
}

type UpdatePersonalizationRequestRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	Reason types.String `tfsdk:"reason" tf:"optional"`

	RequestId types.String `tfsdk:"-"`

	Share types.List `tfsdk:"share" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:""`
}

func (newState *UpdatePersonalizationRequestRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalizationRequestRequest_SdkV2) {
}

func (newState *UpdatePersonalizationRequestRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalizationRequestRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalizationRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalizationRequestRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"share": reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdatePersonalizationRequestRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"reason":     o.Reason,
			"request_id": o.RequestId,
			"share":      o.Share,
			"status":     o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePersonalizationRequestRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"reason":     types.StringType,
			"request_id": types.StringType,
			"share": basetypes.ListType{
				ElemType: ShareInfo{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetShare returns the value of the Share field in UpdatePersonalizationRequestRequest_SdkV2 as
// a ShareInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePersonalizationRequestRequest_SdkV2) GetShare(ctx context.Context) (ShareInfo_SdkV2, bool) {
	var e ShareInfo_SdkV2
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo_SdkV2
	d := o.Share.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in UpdatePersonalizationRequestRequest_SdkV2.
func (o *UpdatePersonalizationRequestRequest_SdkV2) SetShare(ctx context.Context, v ShareInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["share"]
	o.Share = types.ListValueMust(t, vs)
}

type UpdatePersonalizationRequestResponse_SdkV2 struct {
	Request types.List `tfsdk:"request" tf:"optional,object"`
}

func (newState *UpdatePersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalizationRequestResponse_SdkV2) {
}

func (newState *UpdatePersonalizationRequestResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalizationRequestResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalizationRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalizationRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(PersonalizationRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdatePersonalizationRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": o.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePersonalizationRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": basetypes.ListType{
				ElemType: PersonalizationRequest{}.Type(ctx),
			},
		},
	}
}

// GetRequest returns the value of the Request field in UpdatePersonalizationRequestResponse_SdkV2 as
// a PersonalizationRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePersonalizationRequestResponse_SdkV2) GetRequest(ctx context.Context) (PersonalizationRequest_SdkV2, bool) {
	var e PersonalizationRequest_SdkV2
	if o.Request.IsNull() || o.Request.IsUnknown() {
		return e, false
	}
	var v []PersonalizationRequest_SdkV2
	d := o.Request.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRequest sets the value of the Request field in UpdatePersonalizationRequestResponse_SdkV2.
func (o *UpdatePersonalizationRequestResponse_SdkV2) SetRequest(ctx context.Context, v PersonalizationRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["request"]
	o.Request = types.ListValueMust(t, vs)
}

type UpdateProviderAnalyticsDashboardRequest_SdkV2 struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *UpdateProviderAnalyticsDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderAnalyticsDashboardRequest_SdkV2) {
}

func (newState *UpdateProviderAnalyticsDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateProviderAnalyticsDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderAnalyticsDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateProviderAnalyticsDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderAnalyticsDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":      types.StringType,
			"version": types.Int64Type,
		},
	}
}

type UpdateProviderAnalyticsDashboardResponse_SdkV2 struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id" tf:""`

	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *UpdateProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderAnalyticsDashboardResponse_SdkV2) {
}

func (newState *UpdateProviderAnalyticsDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateProviderAnalyticsDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderAnalyticsDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderAnalyticsDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateProviderAnalyticsDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"id":           o.Id,
			"version":      o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderAnalyticsDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
		},
	}
}

type UpdateProviderRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`

	Provider types.List `tfsdk:"provider" tf:"object"`
}

func (newState *UpdateProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderRequest_SdkV2) {
}

func (newState *UpdateProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":       o.Id,
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProvider returns the value of the Provider field in UpdateProviderRequest_SdkV2 as
// a ProviderInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateProviderRequest_SdkV2) GetProvider(ctx context.Context) (ProviderInfo_SdkV2, bool) {
	var e ProviderInfo_SdkV2
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Provider.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in UpdateProviderRequest_SdkV2.
func (o *UpdateProviderRequest_SdkV2) SetProvider(ctx context.Context, v ProviderInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider"]
	o.Provider = types.ListValueMust(t, vs)
}

type UpdateProviderResponse_SdkV2 struct {
	Provider types.List `tfsdk:"provider" tf:"optional,object"`
}

func (newState *UpdateProviderResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProviderResponse_SdkV2) {
}

func (newState *UpdateProviderResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateProviderResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProviderResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateProviderResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProvider returns the value of the Provider field in UpdateProviderResponse_SdkV2 as
// a ProviderInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateProviderResponse_SdkV2) GetProvider(ctx context.Context) (ProviderInfo_SdkV2, bool) {
	var e ProviderInfo_SdkV2
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Provider.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in UpdateProviderResponse_SdkV2.
func (o *UpdateProviderResponse_SdkV2) SetProvider(ctx context.Context, v ProviderInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider"]
	o.Provider = types.ListValueMust(t, vs)
}
