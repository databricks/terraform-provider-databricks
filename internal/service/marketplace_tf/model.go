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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddExchangeForListingRequest struct {
	ExchangeId types.String `tfsdk:"exchange_id"`

	ListingId types.String `tfsdk:"listing_id"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingRequest
// only implements ToObjectValue() and Type().
func (o AddExchangeForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"listing_id":  o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddExchangeForListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"listing_id":  types.StringType,
		},
	}
}

type AddExchangeForListingResponse struct {
	ExchangeForListing types.Object `tfsdk:"exchange_for_listing"`
}

func (newState *AddExchangeForListingResponse) SyncFieldsDuringCreateOrUpdate(plan AddExchangeForListingResponse) {
}

func (newState *AddExchangeForListingResponse) SyncFieldsDuringRead(existingState AddExchangeForListingResponse) {
}

func (c AddExchangeForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_for_listing"] = attrs["exchange_for_listing"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingResponse
// only implements ToObjectValue() and Type().
func (o AddExchangeForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_for_listing": o.ExchangeForListing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddExchangeForListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_for_listing": ExchangeListing{}.Type(ctx),
		},
	}
}

// GetExchangeForListing returns the value of the ExchangeForListing field in AddExchangeForListingResponse as
// a ExchangeListing value.
// If the field is unknown or null, the boolean return value is false.
func (o *AddExchangeForListingResponse) GetExchangeForListing(ctx context.Context) (ExchangeListing, bool) {
	var e ExchangeListing
	if o.ExchangeForListing.IsNull() || o.ExchangeForListing.IsUnknown() {
		return e, false
	}
	var v []ExchangeListing
	d := o.ExchangeForListing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchangeForListing sets the value of the ExchangeForListing field in AddExchangeForListingResponse.
func (o *AddExchangeForListingResponse) SetExchangeForListing(ctx context.Context, v ExchangeListing) {
	vs := v.ToObjectValue(ctx)
	o.ExchangeForListing = vs
}

type BatchGetListingsRequest struct {
	Ids types.List `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsRequest
// only implements ToObjectValue() and Type().
func (o BatchGetListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": o.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetListingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in BatchGetListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetListingsRequest) GetIds(ctx context.Context) ([]types.String, bool) {
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

// SetIds sets the value of the Ids field in BatchGetListingsRequest.
func (o *BatchGetListingsRequest) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Ids = types.ListValueMust(t, vs)
}

type BatchGetListingsResponse struct {
	Listings types.List `tfsdk:"listings"`
}

func (newState *BatchGetListingsResponse) SyncFieldsDuringCreateOrUpdate(plan BatchGetListingsResponse) {
}

func (newState *BatchGetListingsResponse) SyncFieldsDuringRead(existingState BatchGetListingsResponse) {
}

func (c BatchGetListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listings"] = attrs["listings"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsResponse
// only implements ToObjectValue() and Type().
func (o BatchGetListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings": o.Listings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetListingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
		},
	}
}

// GetListings returns the value of the Listings field in BatchGetListingsResponse as
// a slice of Listing values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in BatchGetListingsResponse.
func (o *BatchGetListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type BatchGetProvidersRequest struct {
	Ids types.List `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersRequest
// only implements ToObjectValue() and Type().
func (o BatchGetProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": o.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetProvidersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in BatchGetProvidersRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetProvidersRequest) GetIds(ctx context.Context) ([]types.String, bool) {
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

// SetIds sets the value of the Ids field in BatchGetProvidersRequest.
func (o *BatchGetProvidersRequest) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Ids = types.ListValueMust(t, vs)
}

type BatchGetProvidersResponse struct {
	Providers types.List `tfsdk:"providers"`
}

func (newState *BatchGetProvidersResponse) SyncFieldsDuringCreateOrUpdate(plan BatchGetProvidersResponse) {
}

func (newState *BatchGetProvidersResponse) SyncFieldsDuringRead(existingState BatchGetProvidersResponse) {
}

func (c BatchGetProvidersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["providers"] = attrs["providers"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersResponse
// only implements ToObjectValue() and Type().
func (o BatchGetProvidersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"providers": o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BatchGetProvidersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in BatchGetProvidersResponse as
// a slice of ProviderInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *BatchGetProvidersResponse) GetProviders(ctx context.Context) ([]ProviderInfo, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in BatchGetProvidersResponse.
func (o *BatchGetProvidersResponse) SetProviders(ctx context.Context, v []ProviderInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

type ConsumerTerms struct {
	Version types.String `tfsdk:"version"`
}

func (newState *ConsumerTerms) SyncFieldsDuringCreateOrUpdate(plan ConsumerTerms) {
}

func (newState *ConsumerTerms) SyncFieldsDuringRead(existingState ConsumerTerms) {
}

func (c ConsumerTerms) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConsumerTerms
// only implements ToObjectValue() and Type().
func (o ConsumerTerms) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConsumerTerms) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.StringType,
		},
	}
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company types.String `tfsdk:"company"`

	Email types.String `tfsdk:"email"`

	FirstName types.String `tfsdk:"first_name"`

	LastName types.String `tfsdk:"last_name"`
}

func (newState *ContactInfo) SyncFieldsDuringCreateOrUpdate(plan ContactInfo) {
}

func (newState *ContactInfo) SyncFieldsDuringRead(existingState ContactInfo) {
}

func (c ContactInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["company"] = attrs["company"].SetOptional()
	attrs["email"] = attrs["email"].SetOptional()
	attrs["first_name"] = attrs["first_name"].SetOptional()
	attrs["last_name"] = attrs["last_name"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContactInfo
// only implements ToObjectValue() and Type().
func (o ContactInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ContactInfo) Type(ctx context.Context) attr.Type {
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
	Filter types.Object `tfsdk:"filter"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (o CreateExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeFilterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": ExchangeFilter{}.Type(ctx),
		},
	}
}

// GetFilter returns the value of the Filter field in CreateExchangeFilterRequest as
// a ExchangeFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExchangeFilterRequest) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in CreateExchangeFilterRequest.
func (o *CreateExchangeFilterRequest) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type CreateExchangeFilterResponse struct {
	FilterId types.String `tfsdk:"filter_id"`
}

func (newState *CreateExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(plan CreateExchangeFilterResponse) {
}

func (newState *CreateExchangeFilterResponse) SyncFieldsDuringRead(existingState CreateExchangeFilterResponse) {
}

func (c CreateExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter_id"] = attrs["filter_id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (o CreateExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_id": o.FilterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_id": types.StringType,
		},
	}
}

type CreateExchangeRequest struct {
	Exchange types.Object `tfsdk:"exchange"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeRequest
// only implements ToObjectValue() and Type().
func (o CreateExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in CreateExchangeRequest as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExchangeRequest) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange
	d := o.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in CreateExchangeRequest.
func (o *CreateExchangeRequest) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	o.Exchange = vs
}

type CreateExchangeResponse struct {
	ExchangeId types.String `tfsdk:"exchange_id"`
}

func (newState *CreateExchangeResponse) SyncFieldsDuringCreateOrUpdate(plan CreateExchangeResponse) {
}

func (newState *CreateExchangeResponse) SyncFieldsDuringRead(existingState CreateExchangeResponse) {
}

func (c CreateExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_id"] = attrs["exchange_id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeResponse
// only implements ToObjectValue() and Type().
func (o CreateExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
		},
	}
}

type CreateFileRequest struct {
	DisplayName types.String `tfsdk:"display_name"`

	FileParent types.Object `tfsdk:"file_parent"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileRequest
// only implements ToObjectValue() and Type().
func (o CreateFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":          types.StringType,
			"file_parent":           FileParent{}.Type(ctx),
			"marketplace_file_type": types.StringType,
			"mime_type":             types.StringType,
		},
	}
}

// GetFileParent returns the value of the FileParent field in CreateFileRequest as
// a FileParent value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFileRequest) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent
	d := o.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in CreateFileRequest.
func (o *CreateFileRequest) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	o.FileParent = vs
}

type CreateFileResponse struct {
	FileInfo types.Object `tfsdk:"file_info"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url"`
}

func (newState *CreateFileResponse) SyncFieldsDuringCreateOrUpdate(plan CreateFileResponse) {
}

func (newState *CreateFileResponse) SyncFieldsDuringRead(existingState CreateFileResponse) {
}

func (c CreateFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_info"] = attrs["file_info"].SetOptional()
	attrs["upload_url"] = attrs["upload_url"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileResponse
// only implements ToObjectValue() and Type().
func (o CreateFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info":  o.FileInfo,
			"upload_url": o.UploadUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFileResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info":  FileInfo{}.Type(ctx),
			"upload_url": types.StringType,
		},
	}
}

// GetFileInfo returns the value of the FileInfo field in CreateFileResponse as
// a FileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFileResponse) GetFileInfo(ctx context.Context) (FileInfo, bool) {
	var e FileInfo
	if o.FileInfo.IsNull() || o.FileInfo.IsUnknown() {
		return e, false
	}
	var v []FileInfo
	d := o.FileInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileInfo sets the value of the FileInfo field in CreateFileResponse.
func (o *CreateFileResponse) SetFileInfo(ctx context.Context, v FileInfo) {
	vs := v.ToObjectValue(ctx)
	o.FileInfo = vs
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms types.Object `tfsdk:"accepted_consumer_terms"`

	CatalogName types.String `tfsdk:"catalog_name"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type"`
	// for git repo installations
	RepoDetail types.Object `tfsdk:"repo_detail"`

	ShareName types.String `tfsdk:"share_name"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstallationRequest
// only implements ToObjectValue() and Type().
func (o CreateInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateInstallationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": ConsumerTerms{}.Type(ctx),
			"catalog_name":            types.StringType,
			"listing_id":              types.StringType,
			"recipient_type":          types.StringType,
			"repo_detail":             RepoInstallation{}.Type(ctx),
			"share_name":              types.StringType,
		},
	}
}

// GetAcceptedConsumerTerms returns the value of the AcceptedConsumerTerms field in CreateInstallationRequest as
// a ConsumerTerms value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstallationRequest) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms, bool) {
	var e ConsumerTerms
	if o.AcceptedConsumerTerms.IsNull() || o.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v []ConsumerTerms
	d := o.AcceptedConsumerTerms.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreateInstallationRequest.
func (o *CreateInstallationRequest) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms) {
	vs := v.ToObjectValue(ctx)
	o.AcceptedConsumerTerms = vs
}

// GetRepoDetail returns the value of the RepoDetail field in CreateInstallationRequest as
// a RepoInstallation value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstallationRequest) GetRepoDetail(ctx context.Context) (RepoInstallation, bool) {
	var e RepoInstallation
	if o.RepoDetail.IsNull() || o.RepoDetail.IsUnknown() {
		return e, false
	}
	var v []RepoInstallation
	d := o.RepoDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRepoDetail sets the value of the RepoDetail field in CreateInstallationRequest.
func (o *CreateInstallationRequest) SetRepoDetail(ctx context.Context, v RepoInstallation) {
	vs := v.ToObjectValue(ctx)
	o.RepoDetail = vs
}

type CreateListingRequest struct {
	Listing types.Object `tfsdk:"listing"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingRequest
// only implements ToObjectValue() and Type().
func (o CreateListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in CreateListingRequest as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateListingRequest) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing
	d := o.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in CreateListingRequest.
func (o *CreateListingRequest) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	o.Listing = vs
}

type CreateListingResponse struct {
	ListingId types.String `tfsdk:"listing_id"`
}

func (newState *CreateListingResponse) SyncFieldsDuringCreateOrUpdate(plan CreateListingResponse) {
}

func (newState *CreateListingResponse) SyncFieldsDuringRead(existingState CreateListingResponse) {
}

func (c CreateListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingResponse
// only implements ToObjectValue() and Type().
func (o CreateListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
		},
	}
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms types.Object `tfsdk:"accepted_consumer_terms"`

	Comment types.String `tfsdk:"comment"`

	Company types.String `tfsdk:"company"`

	FirstName types.String `tfsdk:"first_name"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	LastName types.String `tfsdk:"last_name"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequest
// only implements ToObjectValue() and Type().
func (o CreatePersonalizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreatePersonalizationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"accepted_consumer_terms": ConsumerTerms{}.Type(ctx),
			"comment":                 types.StringType,
			"company":                 types.StringType,
			"first_name":              types.StringType,
			"intended_use":            types.StringType,
			"is_from_lighthouse":      types.BoolType,
			"last_name":               types.StringType,
			"listing_id":              types.StringType,
			"recipient_type":          types.StringType,
		},
	}
}

// GetAcceptedConsumerTerms returns the value of the AcceptedConsumerTerms field in CreatePersonalizationRequest as
// a ConsumerTerms value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePersonalizationRequest) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms, bool) {
	var e ConsumerTerms
	if o.AcceptedConsumerTerms.IsNull() || o.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v []ConsumerTerms
	d := o.AcceptedConsumerTerms.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreatePersonalizationRequest.
func (o *CreatePersonalizationRequest) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms) {
	vs := v.ToObjectValue(ctx)
	o.AcceptedConsumerTerms = vs
}

type CreatePersonalizationRequestResponse struct {
	Id types.String `tfsdk:"id"`
}

func (newState *CreatePersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(plan CreatePersonalizationRequestResponse) {
}

func (newState *CreatePersonalizationRequestResponse) SyncFieldsDuringRead(existingState CreatePersonalizationRequestResponse) {
}

func (c CreatePersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (o CreatePersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateProviderAnalyticsDashboardRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (o CreateProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateProviderRequest struct {
	Provider types.Object `tfsdk:"provider"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderRequest
// only implements ToObjectValue() and Type().
func (o CreateProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in CreateProviderRequest as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateProviderRequest) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo
	d := o.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in CreateProviderRequest.
func (o *CreateProviderRequest) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	o.Provider = vs
}

type CreateProviderResponse struct {
	Id types.String `tfsdk:"id"`
}

func (newState *CreateProviderResponse) SyncFieldsDuringCreateOrUpdate(plan CreateProviderResponse) {
}

func (newState *CreateProviderResponse) SyncFieldsDuringRead(existingState CreateProviderResponse) {
}

func (c CreateProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderResponse
// only implements ToObjectValue() and Type().
func (o CreateProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DataRefreshInfo struct {
	Interval types.Int64 `tfsdk:"interval"`

	Unit types.String `tfsdk:"unit"`
}

func (newState *DataRefreshInfo) SyncFieldsDuringCreateOrUpdate(plan DataRefreshInfo) {
}

func (newState *DataRefreshInfo) SyncFieldsDuringRead(existingState DataRefreshInfo) {
}

func (c DataRefreshInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["interval"] = attrs["interval"].SetRequired()
	attrs["unit"] = attrs["unit"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataRefreshInfo
// only implements ToObjectValue() and Type().
func (o DataRefreshInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"interval": o.Interval,
			"unit":     o.Unit,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataRefreshInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"interval": types.Int64Type,
			"unit":     types.StringType,
		},
	}
}

type DeleteExchangeFilterRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (o DeleteExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeFilterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeFilterResponse struct {
}

func (newState *DeleteExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteExchangeFilterResponse) {
}

func (newState *DeleteExchangeFilterResponse) SyncFieldsDuringRead(existingState DeleteExchangeFilterResponse) {
}

func (c DeleteExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (o DeleteExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExchangeRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeRequest
// only implements ToObjectValue() and Type().
func (o DeleteExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeResponse struct {
}

func (newState *DeleteExchangeResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteExchangeResponse) {
}

func (newState *DeleteExchangeResponse) SyncFieldsDuringRead(existingState DeleteExchangeResponse) {
}

func (c DeleteExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeResponse
// only implements ToObjectValue() and Type().
func (o DeleteExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteFileRequest struct {
	FileId types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileRequest
// only implements ToObjectValue() and Type().
func (o DeleteFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": o.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type DeleteFileResponse struct {
}

func (newState *DeleteFileResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteFileResponse) {
}

func (newState *DeleteFileResponse) SyncFieldsDuringRead(existingState DeleteFileResponse) {
}

func (c DeleteFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileResponse
// only implements ToObjectValue() and Type().
func (o DeleteFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFileResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteInstallationRequest struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationRequest
// only implements ToObjectValue() and Type().
func (o DeleteInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation_id": o.InstallationId,
			"listing_id":      o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstallationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
		},
	}
}

type DeleteInstallationResponse struct {
}

func (newState *DeleteInstallationResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteInstallationResponse) {
}

func (newState *DeleteInstallationResponse) SyncFieldsDuringRead(existingState DeleteInstallationResponse) {
}

func (c DeleteInstallationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationResponse
// only implements ToObjectValue() and Type().
func (o DeleteInstallationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstallationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteListingRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingRequest
// only implements ToObjectValue() and Type().
func (o DeleteListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteListingResponse struct {
}

func (newState *DeleteListingResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteListingResponse) {
}

func (newState *DeleteListingResponse) SyncFieldsDuringRead(existingState DeleteListingResponse) {
}

func (c DeleteListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingResponse
// only implements ToObjectValue() and Type().
func (o DeleteListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteProviderRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest
// only implements ToObjectValue() and Type().
func (o DeleteProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteProviderResponse struct {
}

func (newState *DeleteProviderResponse) SyncFieldsDuringCreateOrUpdate(plan DeleteProviderResponse) {
}

func (newState *DeleteProviderResponse) SyncFieldsDuringRead(existingState DeleteProviderResponse) {
}

func (c DeleteProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderResponse
// only implements ToObjectValue() and Type().
func (o DeleteProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Exchange struct {
	Comment types.String `tfsdk:"comment"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	Filters types.List `tfsdk:"filters"`

	Id types.String `tfsdk:"id"`

	LinkedListings types.List `tfsdk:"linked_listings"`

	Name types.String `tfsdk:"name"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *Exchange) SyncFieldsDuringCreateOrUpdate(plan Exchange) {
}

func (newState *Exchange) SyncFieldsDuringRead(existingState Exchange) {
}

func (c Exchange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["linked_listings"] = attrs["linked_listings"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Exchange
// only implements ToObjectValue() and Type().
func (o Exchange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Exchange) Type(ctx context.Context) attr.Type {
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

// GetFilters returns the value of the Filters field in Exchange as
// a slice of ExchangeFilter values.
// If the field is unknown or null, the boolean return value is false.
func (o *Exchange) GetFilters(ctx context.Context) ([]ExchangeFilter, bool) {
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in Exchange.
func (o *Exchange) SetFilters(ctx context.Context, v []ExchangeFilter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Filters = types.ListValueMust(t, vs)
}

// GetLinkedListings returns the value of the LinkedListings field in Exchange as
// a slice of ExchangeListing values.
// If the field is unknown or null, the boolean return value is false.
func (o *Exchange) GetLinkedListings(ctx context.Context) ([]ExchangeListing, bool) {
	if o.LinkedListings.IsNull() || o.LinkedListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := o.LinkedListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLinkedListings sets the value of the LinkedListings field in Exchange.
func (o *Exchange) SetLinkedListings(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["linked_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LinkedListings = types.ListValueMust(t, vs)
}

type ExchangeFilter struct {
	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	ExchangeId types.String `tfsdk:"exchange_id"`

	FilterType types.String `tfsdk:"filter_type"`

	FilterValue types.String `tfsdk:"filter_value"`

	Id types.String `tfsdk:"id"`

	Name types.String `tfsdk:"name"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *ExchangeFilter) SyncFieldsDuringCreateOrUpdate(plan ExchangeFilter) {
}

func (newState *ExchangeFilter) SyncFieldsDuringRead(existingState ExchangeFilter) {
}

func (c ExchangeFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["exchange_id"] = attrs["exchange_id"].SetRequired()
	attrs["filter_type"] = attrs["filter_type"].SetRequired()
	attrs["filter_value"] = attrs["filter_value"].SetRequired()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeFilter
// only implements ToObjectValue() and Type().
func (o ExchangeFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExchangeFilter) Type(ctx context.Context) attr.Type {
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
	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	ExchangeId types.String `tfsdk:"exchange_id"`

	ExchangeName types.String `tfsdk:"exchange_name"`

	Id types.String `tfsdk:"id"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`
}

func (newState *ExchangeListing) SyncFieldsDuringCreateOrUpdate(plan ExchangeListing) {
}

func (newState *ExchangeListing) SyncFieldsDuringRead(existingState ExchangeListing) {
}

func (c ExchangeListing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["exchange_id"] = attrs["exchange_id"].SetOptional()
	attrs["exchange_name"] = attrs["exchange_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetOptional()
	attrs["listing_name"] = attrs["listing_name"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeListing
// only implements ToObjectValue() and Type().
func (o ExchangeListing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExchangeListing) Type(ctx context.Context) attr.Type {
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
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName types.String `tfsdk:"display_name"`

	DownloadLink types.String `tfsdk:"download_link"`

	FileParent types.Object `tfsdk:"file_parent"`

	Id types.String `tfsdk:"id"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`

	Status types.String `tfsdk:"status"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (newState *FileInfo) SyncFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncFieldsDuringRead(existingState FileInfo) {
}

func (c FileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["download_link"] = attrs["download_link"].SetOptional()
	attrs["file_parent"] = attrs["file_parent"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["marketplace_file_type"] = attrs["marketplace_file_type"].SetOptional()
	attrs["mime_type"] = attrs["mime_type"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo
// only implements ToObjectValue() and Type().
func (o FileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o FileInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":            types.Int64Type,
			"display_name":          types.StringType,
			"download_link":         types.StringType,
			"file_parent":           FileParent{}.Type(ctx),
			"id":                    types.StringType,
			"marketplace_file_type": types.StringType,
			"mime_type":             types.StringType,
			"status":                types.StringType,
			"status_message":        types.StringType,
			"updated_at":            types.Int64Type,
		},
	}
}

// GetFileParent returns the value of the FileParent field in FileInfo as
// a FileParent value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileInfo) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent
	d := o.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in FileInfo.
func (o *FileInfo) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	o.FileParent = vs
}

type FileParent struct {
	FileParentType types.String `tfsdk:"file_parent_type"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id"`
}

func (newState *FileParent) SyncFieldsDuringCreateOrUpdate(plan FileParent) {
}

func (newState *FileParent) SyncFieldsDuringRead(existingState FileParent) {
}

func (c FileParent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_parent_type"] = attrs["file_parent_type"].SetOptional()
	attrs["parent_id"] = attrs["parent_id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileParent
// only implements ToObjectValue() and Type().
func (o FileParent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent_type": o.FileParentType,
			"parent_id":        o.ParentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileParent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent_type": types.StringType,
			"parent_id":        types.StringType,
		},
	}
}

type GetExchangeRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeRequest
// only implements ToObjectValue() and Type().
func (o GetExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetExchangeResponse struct {
	Exchange types.Object `tfsdk:"exchange"`
}

func (newState *GetExchangeResponse) SyncFieldsDuringCreateOrUpdate(plan GetExchangeResponse) {
}

func (newState *GetExchangeResponse) SyncFieldsDuringRead(existingState GetExchangeResponse) {
}

func (c GetExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeResponse
// only implements ToObjectValue() and Type().
func (o GetExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in GetExchangeResponse as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetExchangeResponse) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange
	d := o.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in GetExchangeResponse.
func (o *GetExchangeResponse) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	o.Exchange = vs
}

type GetFileRequest struct {
	FileId types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileRequest
// only implements ToObjectValue() and Type().
func (o GetFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": o.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type GetFileResponse struct {
	FileInfo types.Object `tfsdk:"file_info"`
}

func (newState *GetFileResponse) SyncFieldsDuringCreateOrUpdate(plan GetFileResponse) {
}

func (newState *GetFileResponse) SyncFieldsDuringRead(existingState GetFileResponse) {
}

func (c GetFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_info"] = attrs["file_info"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileResponse
// only implements ToObjectValue() and Type().
func (o GetFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info": o.FileInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFileResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": FileInfo{}.Type(ctx),
		},
	}
}

// GetFileInfo returns the value of the FileInfo field in GetFileResponse as
// a FileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetFileResponse) GetFileInfo(ctx context.Context) (FileInfo, bool) {
	var e FileInfo
	if o.FileInfo.IsNull() || o.FileInfo.IsUnknown() {
		return e, false
	}
	var v []FileInfo
	d := o.FileInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileInfo sets the value of the FileInfo field in GetFileResponse.
func (o *GetFileResponse) SetFileInfo(ctx context.Context, v FileInfo) {
	vs := v.ToObjectValue(ctx)
	o.FileInfo = vs
}

type GetLatestVersionProviderAnalyticsDashboardRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLatestVersionProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (o GetLatestVersionProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(plan GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (newState *GetLatestVersionProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(existingState GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (c GetLatestVersionProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (o GetLatestVersionProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.Int64Type,
		},
	}
}

type GetListingContentMetadataRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataRequest
// only implements ToObjectValue() and Type().
func (o GetListingContentMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingContentMetadataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GetListingContentMetadataResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	SharedDataObjects types.List `tfsdk:"shared_data_objects"`
}

func (newState *GetListingContentMetadataResponse) SyncFieldsDuringCreateOrUpdate(plan GetListingContentMetadataResponse) {
}

func (newState *GetListingContentMetadataResponse) SyncFieldsDuringRead(existingState GetListingContentMetadataResponse) {
}

func (c GetListingContentMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["shared_data_objects"] = attrs["shared_data_objects"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataResponse
// only implements ToObjectValue() and Type().
func (o GetListingContentMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     o.NextPageToken,
			"shared_data_objects": o.SharedDataObjects,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingContentMetadataResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shared_data_objects": basetypes.ListType{
				ElemType: SharedDataObject{}.Type(ctx),
			},
		},
	}
}

// GetSharedDataObjects returns the value of the SharedDataObjects field in GetListingContentMetadataResponse as
// a slice of SharedDataObject values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingContentMetadataResponse) GetSharedDataObjects(ctx context.Context) ([]SharedDataObject, bool) {
	if o.SharedDataObjects.IsNull() || o.SharedDataObjects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject
	d := o.SharedDataObjects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSharedDataObjects sets the value of the SharedDataObjects field in GetListingContentMetadataResponse.
func (o *GetListingContentMetadataResponse) SetSharedDataObjects(ctx context.Context, v []SharedDataObject) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shared_data_objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SharedDataObjects = types.ListValueMust(t, vs)
}

type GetListingRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingRequest
// only implements ToObjectValue() and Type().
func (o GetListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetListingResponse struct {
	Listing types.Object `tfsdk:"listing"`
}

func (newState *GetListingResponse) SyncFieldsDuringCreateOrUpdate(plan GetListingResponse) {
}

func (newState *GetListingResponse) SyncFieldsDuringRead(existingState GetListingResponse) {
}

func (c GetListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingResponse
// only implements ToObjectValue() and Type().
func (o GetListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in GetListingResponse as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingResponse) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing
	d := o.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in GetListingResponse.
func (o *GetListingResponse) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	o.Listing = vs
}

type GetListingsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsRequest
// only implements ToObjectValue() and Type().
func (o GetListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GetListingsResponse struct {
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *GetListingsResponse) SyncFieldsDuringCreateOrUpdate(plan GetListingsResponse) {
}

func (newState *GetListingsResponse) SyncFieldsDuringRead(existingState GetListingsResponse) {
}

func (c GetListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listings"] = attrs["listings"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsResponse
// only implements ToObjectValue() and Type().
func (o GetListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetListingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in GetListingsResponse as
// a slice of Listing values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in GetListingsResponse.
func (o *GetListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type GetPersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestRequest
// only implements ToObjectValue() and Type().
func (o GetPersonalizationRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalizationRequestRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
		},
	}
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests types.List `tfsdk:"personalization_requests"`
}

func (newState *GetPersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(plan GetPersonalizationRequestResponse) {
}

func (newState *GetPersonalizationRequestResponse) SyncFieldsDuringRead(existingState GetPersonalizationRequestResponse) {
}

func (c GetPersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["personalization_requests"] = attrs["personalization_requests"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (o GetPersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"personalization_requests": o.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.Type(ctx),
			},
		},
	}
}

// GetPersonalizationRequests returns the value of the PersonalizationRequests field in GetPersonalizationRequestResponse as
// a slice of PersonalizationRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPersonalizationRequestResponse) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest, bool) {
	if o.PersonalizationRequests.IsNull() || o.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest
	d := o.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in GetPersonalizationRequestResponse.
func (o *GetPersonalizationRequestResponse) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PersonalizationRequests = types.ListValueMust(t, vs)
}

type GetProviderRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest
// only implements ToObjectValue() and Type().
func (o GetProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetProviderResponse struct {
	Provider types.Object `tfsdk:"provider"`
}

func (newState *GetProviderResponse) SyncFieldsDuringCreateOrUpdate(plan GetProviderResponse) {
}

func (newState *GetProviderResponse) SyncFieldsDuringRead(existingState GetProviderResponse) {
}

func (c GetProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderResponse
// only implements ToObjectValue() and Type().
func (o GetProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in GetProviderResponse as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetProviderResponse) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo
	d := o.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in GetProviderResponse.
func (o *GetProviderResponse) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	o.Provider = vs
}

type Installation struct {
	Installation types.Object `tfsdk:"installation"`
}

func (newState *Installation) SyncFieldsDuringCreateOrUpdate(plan Installation) {
}

func (newState *Installation) SyncFieldsDuringRead(existingState Installation) {
}

func (c Installation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Installation
// only implements ToObjectValue() and Type().
func (o Installation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": o.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Installation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": InstallationDetail{}.Type(ctx),
		},
	}
}

// GetInstallation returns the value of the Installation field in Installation as
// a InstallationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *Installation) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail
	d := o.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in Installation.
func (o *Installation) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	o.Installation = vs
}

type InstallationDetail struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ErrorMessage types.String `tfsdk:"error_message"`

	Id types.String `tfsdk:"id"`

	InstalledOn types.Int64 `tfsdk:"installed_on"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`

	RecipientType types.String `tfsdk:"recipient_type"`

	RepoName types.String `tfsdk:"repo_name"`

	RepoPath types.String `tfsdk:"repo_path"`

	ShareName types.String `tfsdk:"share_name"`

	Status types.String `tfsdk:"status"`

	TokenDetail types.Object `tfsdk:"token_detail"`

	Tokens types.List `tfsdk:"tokens"`
}

func (newState *InstallationDetail) SyncFieldsDuringCreateOrUpdate(plan InstallationDetail) {
}

func (newState *InstallationDetail) SyncFieldsDuringRead(existingState InstallationDetail) {
}

func (c InstallationDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["installed_on"] = attrs["installed_on"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetOptional()
	attrs["listing_name"] = attrs["listing_name"].SetOptional()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["repo_name"] = attrs["repo_name"].SetOptional()
	attrs["repo_path"] = attrs["repo_path"].SetOptional()
	attrs["share_name"] = attrs["share_name"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["token_detail"] = attrs["token_detail"].SetOptional()
	attrs["tokens"] = attrs["tokens"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallationDetail
// only implements ToObjectValue() and Type().
func (o InstallationDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InstallationDetail) Type(ctx context.Context) attr.Type {
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
			"token_detail":   TokenDetail{}.Type(ctx),
			"tokens": basetypes.ListType{
				ElemType: TokenInfo{}.Type(ctx),
			},
		},
	}
}

// GetTokenDetail returns the value of the TokenDetail field in InstallationDetail as
// a TokenDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstallationDetail) GetTokenDetail(ctx context.Context) (TokenDetail, bool) {
	var e TokenDetail
	if o.TokenDetail.IsNull() || o.TokenDetail.IsUnknown() {
		return e, false
	}
	var v []TokenDetail
	d := o.TokenDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenDetail sets the value of the TokenDetail field in InstallationDetail.
func (o *InstallationDetail) SetTokenDetail(ctx context.Context, v TokenDetail) {
	vs := v.ToObjectValue(ctx)
	o.TokenDetail = vs
}

// GetTokens returns the value of the Tokens field in InstallationDetail as
// a slice of TokenInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstallationDetail) GetTokens(ctx context.Context) ([]TokenInfo, bool) {
	if o.Tokens.IsNull() || o.Tokens.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo
	d := o.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in InstallationDetail.
func (o *InstallationDetail) SetTokens(ctx context.Context, v []TokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tokens = types.ListValueMust(t, vs)
}

type ListAllInstallationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsRequest
// only implements ToObjectValue() and Type().
func (o ListAllInstallationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllInstallationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAllInstallationsResponse struct {
	Installations types.List `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListAllInstallationsResponse) SyncFieldsDuringCreateOrUpdate(plan ListAllInstallationsResponse) {
}

func (newState *ListAllInstallationsResponse) SyncFieldsDuringRead(existingState ListAllInstallationsResponse) {
}

func (c ListAllInstallationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installations"] = attrs["installations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsResponse
// only implements ToObjectValue() and Type().
func (o ListAllInstallationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   o.Installations,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllInstallationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetInstallations returns the value of the Installations field in ListAllInstallationsResponse as
// a slice of InstallationDetail values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAllInstallationsResponse) GetInstallations(ctx context.Context) ([]InstallationDetail, bool) {
	if o.Installations.IsNull() || o.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail
	d := o.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListAllInstallationsResponse.
func (o *ListAllInstallationsResponse) SetInstallations(ctx context.Context, v []InstallationDetail) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Installations = types.ListValueMust(t, vs)
}

type ListAllPersonalizationRequestsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsRequest
// only implements ToObjectValue() and Type().
func (o ListAllPersonalizationRequestsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllPersonalizationRequestsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	PersonalizationRequests types.List `tfsdk:"personalization_requests"`
}

func (newState *ListAllPersonalizationRequestsResponse) SyncFieldsDuringCreateOrUpdate(plan ListAllPersonalizationRequestsResponse) {
}

func (newState *ListAllPersonalizationRequestsResponse) SyncFieldsDuringRead(existingState ListAllPersonalizationRequestsResponse) {
}

func (c ListAllPersonalizationRequestsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["personalization_requests"] = attrs["personalization_requests"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsResponse
// only implements ToObjectValue() and Type().
func (o ListAllPersonalizationRequestsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          o.NextPageToken,
			"personalization_requests": o.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllPersonalizationRequestsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"personalization_requests": basetypes.ListType{
				ElemType: PersonalizationRequest{}.Type(ctx),
			},
		},
	}
}

// GetPersonalizationRequests returns the value of the PersonalizationRequests field in ListAllPersonalizationRequestsResponse as
// a slice of PersonalizationRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAllPersonalizationRequestsResponse) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest, bool) {
	if o.PersonalizationRequests.IsNull() || o.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest
	d := o.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in ListAllPersonalizationRequestsResponse.
func (o *ListAllPersonalizationRequestsResponse) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PersonalizationRequests = types.ListValueMust(t, vs)
}

type ListExchangeFiltersRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersRequest
// only implements ToObjectValue() and Type().
func (o ListExchangeFiltersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangeFiltersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListExchangeFiltersResponse struct {
	Filters types.List `tfsdk:"filters"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListExchangeFiltersResponse) SyncFieldsDuringCreateOrUpdate(plan ListExchangeFiltersResponse) {
}

func (newState *ListExchangeFiltersResponse) SyncFieldsDuringRead(existingState ListExchangeFiltersResponse) {
}

func (c ListExchangeFiltersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersResponse
// only implements ToObjectValue() and Type().
func (o ListExchangeFiltersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filters":         o.Filters,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangeFiltersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filters": basetypes.ListType{
				ElemType: ExchangeFilter{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFilters returns the value of the Filters field in ListExchangeFiltersResponse as
// a slice of ExchangeFilter values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangeFiltersResponse) GetFilters(ctx context.Context) ([]ExchangeFilter, bool) {
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in ListExchangeFiltersResponse.
func (o *ListExchangeFiltersResponse) SetFilters(ctx context.Context, v []ExchangeFilter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Filters = types.ListValueMust(t, vs)
}

type ListExchangesForListingRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingRequest
// only implements ToObjectValue() and Type().
func (o ListExchangesForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesForListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListExchangesForListingResponse struct {
	ExchangeListing types.List `tfsdk:"exchange_listing"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListExchangesForListingResponse) SyncFieldsDuringCreateOrUpdate(plan ListExchangesForListingResponse) {
}

func (newState *ListExchangesForListingResponse) SyncFieldsDuringRead(existingState ListExchangesForListingResponse) {
}

func (c ListExchangesForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_listing"] = attrs["exchange_listing"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingResponse
// only implements ToObjectValue() and Type().
func (o ListExchangesForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listing": o.ExchangeListing,
			"next_page_token":  o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesForListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listing": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchangeListing returns the value of the ExchangeListing field in ListExchangesForListingResponse as
// a slice of ExchangeListing values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangesForListingResponse) GetExchangeListing(ctx context.Context) ([]ExchangeListing, bool) {
	if o.ExchangeListing.IsNull() || o.ExchangeListing.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := o.ExchangeListing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListing sets the value of the ExchangeListing field in ListExchangesForListingResponse.
func (o *ListExchangesForListingResponse) SetExchangeListing(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listing"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeListing = types.ListValueMust(t, vs)
}

type ListExchangesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesRequest
// only implements ToObjectValue() and Type().
func (o ListExchangesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListExchangesResponse struct {
	Exchanges types.List `tfsdk:"exchanges"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListExchangesResponse) SyncFieldsDuringCreateOrUpdate(plan ListExchangesResponse) {
}

func (newState *ListExchangesResponse) SyncFieldsDuringRead(existingState ListExchangesResponse) {
}

func (c ListExchangesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchanges"] = attrs["exchanges"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesResponse
// only implements ToObjectValue() and Type().
func (o ListExchangesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchanges":       o.Exchanges,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExchangesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchanges": basetypes.ListType{
				ElemType: Exchange{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchanges returns the value of the Exchanges field in ListExchangesResponse as
// a slice of Exchange values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExchangesResponse) GetExchanges(ctx context.Context) ([]Exchange, bool) {
	if o.Exchanges.IsNull() || o.Exchanges.IsUnknown() {
		return nil, false
	}
	var v []Exchange
	d := o.Exchanges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchanges sets the value of the Exchanges field in ListExchangesResponse.
func (o *ListExchangesResponse) SetExchanges(ctx context.Context, v []Exchange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchanges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exchanges = types.ListValueMust(t, vs)
}

type ListFilesRequest struct {
	FileParent types.Object `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesRequest
// only implements ToObjectValue() and Type().
func (o ListFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent": o.FileParent,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_parent": FileParent{}.Type(ctx),
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

// GetFileParent returns the value of the FileParent field in ListFilesRequest as
// a FileParent value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFilesRequest) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if o.FileParent.IsNull() || o.FileParent.IsUnknown() {
		return e, false
	}
	var v []FileParent
	d := o.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileParent sets the value of the FileParent field in ListFilesRequest.
func (o *ListFilesRequest) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	o.FileParent = vs
}

type ListFilesResponse struct {
	FileInfos types.List `tfsdk:"file_infos"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListFilesResponse) SyncFieldsDuringCreateOrUpdate(plan ListFilesResponse) {
}

func (newState *ListFilesResponse) SyncFieldsDuringRead(existingState ListFilesResponse) {
}

func (c ListFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_infos"] = attrs["file_infos"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesResponse
// only implements ToObjectValue() and Type().
func (o ListFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_infos":      o.FileInfos,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_infos": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFileInfos returns the value of the FileInfos field in ListFilesResponse as
// a slice of FileInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFilesResponse) GetFileInfos(ctx context.Context) ([]FileInfo, bool) {
	if o.FileInfos.IsNull() || o.FileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := o.FileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileInfos sets the value of the FileInfos field in ListFilesResponse.
func (o *ListFilesResponse) SetFileInfos(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileInfos = types.ListValueMust(t, vs)
}

type ListFulfillmentsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsRequest
// only implements ToObjectValue() and Type().
func (o ListFulfillmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFulfillmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListFulfillmentsResponse struct {
	Fulfillments types.List `tfsdk:"fulfillments"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListFulfillmentsResponse) SyncFieldsDuringCreateOrUpdate(plan ListFulfillmentsResponse) {
}

func (newState *ListFulfillmentsResponse) SyncFieldsDuringRead(existingState ListFulfillmentsResponse) {
}

func (c ListFulfillmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fulfillments"] = attrs["fulfillments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsResponse
// only implements ToObjectValue() and Type().
func (o ListFulfillmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fulfillments":    o.Fulfillments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFulfillmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillments": basetypes.ListType{
				ElemType: ListingFulfillment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFulfillments returns the value of the Fulfillments field in ListFulfillmentsResponse as
// a slice of ListingFulfillment values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFulfillmentsResponse) GetFulfillments(ctx context.Context) ([]ListingFulfillment, bool) {
	if o.Fulfillments.IsNull() || o.Fulfillments.IsUnknown() {
		return nil, false
	}
	var v []ListingFulfillment
	d := o.Fulfillments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFulfillments sets the value of the Fulfillments field in ListFulfillmentsResponse.
func (o *ListFulfillmentsResponse) SetFulfillments(ctx context.Context, v []ListingFulfillment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fulfillments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Fulfillments = types.ListValueMust(t, vs)
}

type ListInstallationsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsRequest
// only implements ToObjectValue() and Type().
func (o ListInstallationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": o.ListingId,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstallationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListInstallationsResponse struct {
	Installations types.List `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListInstallationsResponse) SyncFieldsDuringCreateOrUpdate(plan ListInstallationsResponse) {
}

func (newState *ListInstallationsResponse) SyncFieldsDuringRead(existingState ListInstallationsResponse) {
}

func (c ListInstallationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installations"] = attrs["installations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsResponse
// only implements ToObjectValue() and Type().
func (o ListInstallationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   o.Installations,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstallationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installations": basetypes.ListType{
				ElemType: InstallationDetail{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetInstallations returns the value of the Installations field in ListInstallationsResponse as
// a slice of InstallationDetail values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListInstallationsResponse) GetInstallations(ctx context.Context) ([]InstallationDetail, bool) {
	if o.Installations.IsNull() || o.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail
	d := o.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListInstallationsResponse.
func (o *ListInstallationsResponse) SetInstallations(ctx context.Context, v []InstallationDetail) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Installations = types.ListValueMust(t, vs)
}

type ListListingsForExchangeRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeRequest
// only implements ToObjectValue() and Type().
func (o ListListingsForExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": o.ExchangeId,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsForExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_id": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListListingsForExchangeResponse struct {
	ExchangeListings types.List `tfsdk:"exchange_listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListListingsForExchangeResponse) SyncFieldsDuringCreateOrUpdate(plan ListListingsForExchangeResponse) {
}

func (newState *ListListingsForExchangeResponse) SyncFieldsDuringRead(existingState ListListingsForExchangeResponse) {
}

func (c ListListingsForExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_listings"] = attrs["exchange_listings"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeResponse
// only implements ToObjectValue() and Type().
func (o ListListingsForExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listings": o.ExchangeListings,
			"next_page_token":   o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsForExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_listings": basetypes.ListType{
				ElemType: ExchangeListing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExchangeListings returns the value of the ExchangeListings field in ListListingsForExchangeResponse as
// a slice of ExchangeListing values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsForExchangeResponse) GetExchangeListings(ctx context.Context) ([]ExchangeListing, bool) {
	if o.ExchangeListings.IsNull() || o.ExchangeListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := o.ExchangeListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListings sets the value of the ExchangeListings field in ListListingsForExchangeResponse.
func (o *ListListingsForExchangeResponse) SetExchangeListings(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeListings = types.ListValueMust(t, vs)
}

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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsRequest
// only implements ToObjectValue() and Type().
func (o ListListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListListingsRequest) Type(ctx context.Context) attr.Type {
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

// GetAssets returns the value of the Assets field in ListListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest) GetAssets(ctx context.Context) ([]types.String, bool) {
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

// SetAssets sets the value of the Assets field in ListListingsRequest.
func (o *ListListingsRequest) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in ListListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest) GetCategories(ctx context.Context) ([]types.String, bool) {
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

// SetCategories sets the value of the Categories field in ListListingsRequest.
func (o *ListListingsRequest) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in ListListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest) GetProviderIds(ctx context.Context) ([]types.String, bool) {
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

// SetProviderIds sets the value of the ProviderIds field in ListListingsRequest.
func (o *ListListingsRequest) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ProviderIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListListingsRequest as
// a slice of ListingTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsRequest) GetTags(ctx context.Context) ([]ListingTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListListingsRequest.
func (o *ListListingsRequest) SetTags(ctx context.Context, v []ListingTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ListListingsResponse struct {
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListListingsResponse) SyncFieldsDuringCreateOrUpdate(plan ListListingsResponse) {
}

func (newState *ListListingsResponse) SyncFieldsDuringRead(existingState ListListingsResponse) {
}

func (c ListListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listings"] = attrs["listings"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsResponse
// only implements ToObjectValue() and Type().
func (o ListListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListListingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in ListListingsResponse as
// a slice of Listing values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in ListListingsResponse.
func (o *ListListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type ListProviderAnalyticsDashboardRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (o ListProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`

	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

func (newState *ListProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(plan ListProviderAnalyticsDashboardResponse) {
}

func (newState *ListProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(existingState ListProviderAnalyticsDashboardResponse) {
}

func (c ListProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (o ListProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"id":           o.Id,
			"version":      o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
		},
	}
}

type ListProvidersRequest struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest
// only implements ToObjectValue() and Type().
func (o ListProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_featured": o.IsFeatured,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_featured": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListProvidersResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Providers types.List `tfsdk:"providers"`
}

func (newState *ListProvidersResponse) SyncFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncFieldsDuringRead(existingState ListProvidersResponse) {
}

func (c ListProvidersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["providers"] = attrs["providers"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse
// only implements ToObjectValue() and Type().
func (o ListProvidersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"providers":       o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in ListProvidersResponse as
// a slice of ProviderInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProvidersResponse) GetProviders(ctx context.Context) ([]ProviderInfo, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse.
func (o *ListProvidersResponse) SetProviders(ctx context.Context, v []ProviderInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

type Listing struct {
	Detail types.Object `tfsdk:"detail"`

	Id types.String `tfsdk:"id"`

	Summary types.Object `tfsdk:"summary"`
}

func (newState *Listing) SyncFieldsDuringCreateOrUpdate(plan Listing) {
}

func (newState *Listing) SyncFieldsDuringRead(existingState Listing) {
}

func (c Listing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["detail"] = attrs["detail"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["summary"] = attrs["summary"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Listing
// only implements ToObjectValue() and Type().
func (o Listing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"detail":  o.Detail,
			"id":      o.Id,
			"summary": o.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Listing) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"detail":  ListingDetail{}.Type(ctx),
			"id":      types.StringType,
			"summary": ListingSummary{}.Type(ctx),
		},
	}
}

// GetDetail returns the value of the Detail field in Listing as
// a ListingDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *Listing) GetDetail(ctx context.Context) (ListingDetail, bool) {
	var e ListingDetail
	if o.Detail.IsNull() || o.Detail.IsUnknown() {
		return e, false
	}
	var v []ListingDetail
	d := o.Detail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDetail sets the value of the Detail field in Listing.
func (o *Listing) SetDetail(ctx context.Context, v ListingDetail) {
	vs := v.ToObjectValue(ctx)
	o.Detail = vs
}

// GetSummary returns the value of the Summary field in Listing as
// a ListingSummary value.
// If the field is unknown or null, the boolean return value is false.
func (o *Listing) GetSummary(ctx context.Context) (ListingSummary, bool) {
	var e ListingSummary
	if o.Summary.IsNull() || o.Summary.IsUnknown() {
		return e, false
	}
	var v []ListingSummary
	d := o.Summary.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSummary sets the value of the Summary field in Listing.
func (o *Listing) SetSummary(ctx context.Context, v ListingSummary) {
	vs := v.ToObjectValue(ctx)
	o.Summary = vs
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets types.List `tfsdk:"assets"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd types.Int64 `tfsdk:"collection_date_end"`
	// The starting date timestamp for when the data spans
	CollectionDateStart types.Int64 `tfsdk:"collection_date_start"`
	// Smallest unit of time in the dataset
	CollectionGranularity types.Object `tfsdk:"collection_granularity"`
	// Whether the dataset is free or paid
	Cost types.String `tfsdk:"cost"`
	// Where/how the data is sourced
	DataSource types.String `tfsdk:"data_source"`

	Description types.String `tfsdk:"description"`

	DocumentationLink types.String `tfsdk:"documentation_link"`

	EmbeddedNotebookFileInfos types.List `tfsdk:"embedded_notebook_file_infos"`

	FileIds types.List `tfsdk:"file_ids"`
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
	Tags types.List `tfsdk:"tags"`

	TermsOfService types.String `tfsdk:"terms_of_service"`
	// How often data is updated
	UpdateFrequency types.Object `tfsdk:"update_frequency"`
}

func (newState *ListingDetail) SyncFieldsDuringCreateOrUpdate(plan ListingDetail) {
}

func (newState *ListingDetail) SyncFieldsDuringRead(existingState ListingDetail) {
}

func (c ListingDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets"] = attrs["assets"].SetOptional()
	attrs["collection_date_end"] = attrs["collection_date_end"].SetOptional()
	attrs["collection_date_start"] = attrs["collection_date_start"].SetOptional()
	attrs["collection_granularity"] = attrs["collection_granularity"].SetOptional()
	attrs["cost"] = attrs["cost"].SetOptional()
	attrs["data_source"] = attrs["data_source"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["documentation_link"] = attrs["documentation_link"].SetOptional()
	attrs["embedded_notebook_file_infos"] = attrs["embedded_notebook_file_infos"].SetOptional()
	attrs["file_ids"] = attrs["file_ids"].SetOptional()
	attrs["geographical_coverage"] = attrs["geographical_coverage"].SetOptional()
	attrs["license"] = attrs["license"].SetOptional()
	attrs["pricing_model"] = attrs["pricing_model"].SetOptional()
	attrs["privacy_policy_link"] = attrs["privacy_policy_link"].SetOptional()
	attrs["size"] = attrs["size"].SetOptional()
	attrs["support_link"] = attrs["support_link"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["terms_of_service"] = attrs["terms_of_service"].SetOptional()
	attrs["update_frequency"] = attrs["update_frequency"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingDetail
// only implements ToObjectValue() and Type().
func (o ListingDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListingDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"collection_date_end":    types.Int64Type,
			"collection_date_start":  types.Int64Type,
			"collection_granularity": DataRefreshInfo{}.Type(ctx),
			"cost":                   types.StringType,
			"data_source":            types.StringType,
			"description":            types.StringType,
			"documentation_link":     types.StringType,
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
			"update_frequency": DataRefreshInfo{}.Type(ctx),
		},
	}
}

// GetAssets returns the value of the Assets field in ListingDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetAssets(ctx context.Context) ([]types.String, bool) {
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

// SetAssets sets the value of the Assets field in ListingDetail.
func (o *ListingDetail) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCollectionGranularity returns the value of the CollectionGranularity field in ListingDetail as
// a DataRefreshInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetCollectionGranularity(ctx context.Context) (DataRefreshInfo, bool) {
	var e DataRefreshInfo
	if o.CollectionGranularity.IsNull() || o.CollectionGranularity.IsUnknown() {
		return e, false
	}
	var v []DataRefreshInfo
	d := o.CollectionGranularity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCollectionGranularity sets the value of the CollectionGranularity field in ListingDetail.
func (o *ListingDetail) SetCollectionGranularity(ctx context.Context, v DataRefreshInfo) {
	vs := v.ToObjectValue(ctx)
	o.CollectionGranularity = vs
}

// GetEmbeddedNotebookFileInfos returns the value of the EmbeddedNotebookFileInfos field in ListingDetail as
// a slice of FileInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetEmbeddedNotebookFileInfos(ctx context.Context) ([]FileInfo, bool) {
	if o.EmbeddedNotebookFileInfos.IsNull() || o.EmbeddedNotebookFileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := o.EmbeddedNotebookFileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddedNotebookFileInfos sets the value of the EmbeddedNotebookFileInfos field in ListingDetail.
func (o *ListingDetail) SetEmbeddedNotebookFileInfos(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedded_notebook_file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddedNotebookFileInfos = types.ListValueMust(t, vs)
}

// GetFileIds returns the value of the FileIds field in ListingDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetFileIds(ctx context.Context) ([]types.String, bool) {
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

// SetFileIds sets the value of the FileIds field in ListingDetail.
func (o *ListingDetail) SetFileIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListingDetail as
// a slice of ListingTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetTags(ctx context.Context) ([]ListingTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListingDetail.
func (o *ListingDetail) SetTags(ctx context.Context, v []ListingTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetUpdateFrequency returns the value of the UpdateFrequency field in ListingDetail as
// a DataRefreshInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingDetail) GetUpdateFrequency(ctx context.Context) (DataRefreshInfo, bool) {
	var e DataRefreshInfo
	if o.UpdateFrequency.IsNull() || o.UpdateFrequency.IsUnknown() {
		return e, false
	}
	var v []DataRefreshInfo
	d := o.UpdateFrequency.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdateFrequency sets the value of the UpdateFrequency field in ListingDetail.
func (o *ListingDetail) SetUpdateFrequency(ctx context.Context, v DataRefreshInfo) {
	vs := v.ToObjectValue(ctx)
	o.UpdateFrequency = vs
}

type ListingFulfillment struct {
	FulfillmentType types.String `tfsdk:"fulfillment_type"`

	ListingId types.String `tfsdk:"listing_id"`

	RecipientType types.String `tfsdk:"recipient_type"`

	RepoInfo types.Object `tfsdk:"repo_info"`

	ShareInfo types.Object `tfsdk:"share_info"`
}

func (newState *ListingFulfillment) SyncFieldsDuringCreateOrUpdate(plan ListingFulfillment) {
}

func (newState *ListingFulfillment) SyncFieldsDuringRead(existingState ListingFulfillment) {
}

func (c ListingFulfillment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fulfillment_type"] = attrs["fulfillment_type"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["repo_info"] = attrs["repo_info"].SetOptional()
	attrs["share_info"] = attrs["share_info"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingFulfillment
// only implements ToObjectValue() and Type().
func (o ListingFulfillment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListingFulfillment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fulfillment_type": types.StringType,
			"listing_id":       types.StringType,
			"recipient_type":   types.StringType,
			"repo_info":        RepoInfo{}.Type(ctx),
			"share_info":       ShareInfo{}.Type(ctx),
		},
	}
}

// GetRepoInfo returns the value of the RepoInfo field in ListingFulfillment as
// a RepoInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingFulfillment) GetRepoInfo(ctx context.Context) (RepoInfo, bool) {
	var e RepoInfo
	if o.RepoInfo.IsNull() || o.RepoInfo.IsUnknown() {
		return e, false
	}
	var v []RepoInfo
	d := o.RepoInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRepoInfo sets the value of the RepoInfo field in ListingFulfillment.
func (o *ListingFulfillment) SetRepoInfo(ctx context.Context, v RepoInfo) {
	vs := v.ToObjectValue(ctx)
	o.RepoInfo = vs
}

// GetShareInfo returns the value of the ShareInfo field in ListingFulfillment as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingFulfillment) GetShareInfo(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if o.ShareInfo.IsNull() || o.ShareInfo.IsUnknown() {
		return e, false
	}
	var v []ShareInfo
	d := o.ShareInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShareInfo sets the value of the ShareInfo field in ListingFulfillment.
func (o *ListingFulfillment) SetShareInfo(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	o.ShareInfo = vs
}

type ListingSetting struct {
	Visibility types.String `tfsdk:"visibility"`
}

func (newState *ListingSetting) SyncFieldsDuringCreateOrUpdate(plan ListingSetting) {
}

func (newState *ListingSetting) SyncFieldsDuringRead(existingState ListingSetting) {
}

func (c ListingSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["visibility"] = attrs["visibility"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingSetting
// only implements ToObjectValue() and Type().
func (o ListingSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visibility": o.Visibility,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"visibility": types.StringType,
		},
	}
}

type ListingSummary struct {
	Categories types.List `tfsdk:"categories"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	CreatedById types.Int64 `tfsdk:"created_by_id"`

	ExchangeIds types.List `tfsdk:"exchange_ids"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo types.Object `tfsdk:"git_repo"`

	ListingType types.String `tfsdk:"listingType"`

	Name types.String `tfsdk:"name"`

	ProviderId types.String `tfsdk:"provider_id"`

	ProviderRegion types.Object `tfsdk:"provider_region"`

	PublishedAt types.Int64 `tfsdk:"published_at"`

	PublishedBy types.String `tfsdk:"published_by"`

	Setting types.Object `tfsdk:"setting"`

	Share types.Object `tfsdk:"share"`

	Status types.String `tfsdk:"status"`

	Subtitle types.String `tfsdk:"subtitle"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`

	UpdatedById types.Int64 `tfsdk:"updated_by_id"`
}

func (newState *ListingSummary) SyncFieldsDuringCreateOrUpdate(plan ListingSummary) {
}

func (newState *ListingSummary) SyncFieldsDuringRead(existingState ListingSummary) {
}

func (c ListingSummary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["categories"] = attrs["categories"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["created_by_id"] = attrs["created_by_id"].SetOptional()
	attrs["exchange_ids"] = attrs["exchange_ids"].SetOptional()
	attrs["git_repo"] = attrs["git_repo"].SetOptional()
	attrs["listingType"] = attrs["listingType"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["provider_id"] = attrs["provider_id"].SetOptional()
	attrs["provider_region"] = attrs["provider_region"].SetOptional()
	attrs["published_at"] = attrs["published_at"].SetOptional()
	attrs["published_by"] = attrs["published_by"].SetOptional()
	attrs["setting"] = attrs["setting"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["subtitle"] = attrs["subtitle"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["updated_by_id"] = attrs["updated_by_id"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingSummary
// only implements ToObjectValue() and Type().
func (o ListingSummary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListingSummary) Type(ctx context.Context) attr.Type {
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
			"git_repo":        RepoInfo{}.Type(ctx),
			"listingType":     types.StringType,
			"name":            types.StringType,
			"provider_id":     types.StringType,
			"provider_region": RegionInfo{}.Type(ctx),
			"published_at":    types.Int64Type,
			"published_by":    types.StringType,
			"setting":         ListingSetting{}.Type(ctx),
			"share":           ShareInfo{}.Type(ctx),
			"status":          types.StringType,
			"subtitle":        types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
			"updated_by_id":   types.Int64Type,
		},
	}
}

// GetCategories returns the value of the Categories field in ListingSummary as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetCategories(ctx context.Context) ([]types.String, bool) {
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

// SetCategories sets the value of the Categories field in ListingSummary.
func (o *ListingSummary) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetExchangeIds returns the value of the ExchangeIds field in ListingSummary as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetExchangeIds(ctx context.Context) ([]types.String, bool) {
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

// SetExchangeIds sets the value of the ExchangeIds field in ListingSummary.
func (o *ListingSummary) SetExchangeIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExchangeIds = types.ListValueMust(t, vs)
}

// GetGitRepo returns the value of the GitRepo field in ListingSummary as
// a RepoInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetGitRepo(ctx context.Context) (RepoInfo, bool) {
	var e RepoInfo
	if o.GitRepo.IsNull() || o.GitRepo.IsUnknown() {
		return e, false
	}
	var v []RepoInfo
	d := o.GitRepo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitRepo sets the value of the GitRepo field in ListingSummary.
func (o *ListingSummary) SetGitRepo(ctx context.Context, v RepoInfo) {
	vs := v.ToObjectValue(ctx)
	o.GitRepo = vs
}

// GetProviderRegion returns the value of the ProviderRegion field in ListingSummary as
// a RegionInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetProviderRegion(ctx context.Context) (RegionInfo, bool) {
	var e RegionInfo
	if o.ProviderRegion.IsNull() || o.ProviderRegion.IsUnknown() {
		return e, false
	}
	var v []RegionInfo
	d := o.ProviderRegion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProviderRegion sets the value of the ProviderRegion field in ListingSummary.
func (o *ListingSummary) SetProviderRegion(ctx context.Context, v RegionInfo) {
	vs := v.ToObjectValue(ctx)
	o.ProviderRegion = vs
}

// GetSetting returns the value of the Setting field in ListingSummary as
// a ListingSetting value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetSetting(ctx context.Context) (ListingSetting, bool) {
	var e ListingSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []ListingSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in ListingSummary.
func (o *ListingSummary) SetSetting(ctx context.Context, v ListingSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

// GetShare returns the value of the Share field in ListingSummary as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingSummary) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo
	d := o.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in ListingSummary.
func (o *ListingSummary) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	o.Share = vs
}

type ListingTag struct {
	// Tag name (enum)
	TagName types.String `tfsdk:"tag_name"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues types.List `tfsdk:"tag_values"`
}

func (newState *ListingTag) SyncFieldsDuringCreateOrUpdate(plan ListingTag) {
}

func (newState *ListingTag) SyncFieldsDuringRead(existingState ListingTag) {
}

func (c ListingTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_name"] = attrs["tag_name"].SetOptional()
	attrs["tag_values"] = attrs["tag_values"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingTag
// only implements ToObjectValue() and Type().
func (o ListingTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_name":   o.TagName,
			"tag_values": o.TagValues,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListingTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_name": types.StringType,
			"tag_values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTagValues returns the value of the TagValues field in ListingTag as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListingTag) GetTagValues(ctx context.Context) ([]types.String, bool) {
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

// SetTagValues sets the value of the TagValues field in ListingTag.
func (o *ListingTag) SetTagValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagValues = types.ListValueMust(t, vs)
}

type PersonalizationRequest struct {
	Comment types.String `tfsdk:"comment"`

	ConsumerRegion types.Object `tfsdk:"consumer_region"`

	ContactInfo types.Object `tfsdk:"contact_info"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	Id types.String `tfsdk:"id"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`

	MetastoreId types.String `tfsdk:"metastore_id"`

	ProviderId types.String `tfsdk:"provider_id"`

	RecipientType types.String `tfsdk:"recipient_type"`

	Share types.Object `tfsdk:"share"`

	Status types.String `tfsdk:"status"`

	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (newState *PersonalizationRequest) SyncFieldsDuringCreateOrUpdate(plan PersonalizationRequest) {
}

func (newState *PersonalizationRequest) SyncFieldsDuringRead(existingState PersonalizationRequest) {
}

func (c PersonalizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["consumer_region"] = attrs["consumer_region"].SetRequired()
	attrs["contact_info"] = attrs["contact_info"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["intended_use"] = attrs["intended_use"].SetOptional()
	attrs["is_from_lighthouse"] = attrs["is_from_lighthouse"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetOptional()
	attrs["listing_name"] = attrs["listing_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["provider_id"] = attrs["provider_id"].SetOptional()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalizationRequest
// only implements ToObjectValue() and Type().
func (o PersonalizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PersonalizationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"consumer_region":    RegionInfo{}.Type(ctx),
			"contact_info":       ContactInfo{}.Type(ctx),
			"created_at":         types.Int64Type,
			"id":                 types.StringType,
			"intended_use":       types.StringType,
			"is_from_lighthouse": types.BoolType,
			"listing_id":         types.StringType,
			"listing_name":       types.StringType,
			"metastore_id":       types.StringType,
			"provider_id":        types.StringType,
			"recipient_type":     types.StringType,
			"share":              ShareInfo{}.Type(ctx),
			"status":             types.StringType,
			"status_message":     types.StringType,
			"updated_at":         types.Int64Type,
		},
	}
}

// GetConsumerRegion returns the value of the ConsumerRegion field in PersonalizationRequest as
// a RegionInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest) GetConsumerRegion(ctx context.Context) (RegionInfo, bool) {
	var e RegionInfo
	if o.ConsumerRegion.IsNull() || o.ConsumerRegion.IsUnknown() {
		return e, false
	}
	var v []RegionInfo
	d := o.ConsumerRegion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConsumerRegion sets the value of the ConsumerRegion field in PersonalizationRequest.
func (o *PersonalizationRequest) SetConsumerRegion(ctx context.Context, v RegionInfo) {
	vs := v.ToObjectValue(ctx)
	o.ConsumerRegion = vs
}

// GetContactInfo returns the value of the ContactInfo field in PersonalizationRequest as
// a ContactInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest) GetContactInfo(ctx context.Context) (ContactInfo, bool) {
	var e ContactInfo
	if o.ContactInfo.IsNull() || o.ContactInfo.IsUnknown() {
		return e, false
	}
	var v []ContactInfo
	d := o.ContactInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContactInfo sets the value of the ContactInfo field in PersonalizationRequest.
func (o *PersonalizationRequest) SetContactInfo(ctx context.Context, v ContactInfo) {
	vs := v.ToObjectValue(ctx)
	o.ContactInfo = vs
}

// GetShare returns the value of the Share field in PersonalizationRequest as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *PersonalizationRequest) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo
	d := o.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in PersonalizationRequest.
func (o *PersonalizationRequest) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	o.Share = vs
}

type ProviderAnalyticsDashboard struct {
	Id types.String `tfsdk:"id"`
}

func (newState *ProviderAnalyticsDashboard) SyncFieldsDuringCreateOrUpdate(plan ProviderAnalyticsDashboard) {
}

func (newState *ProviderAnalyticsDashboard) SyncFieldsDuringRead(existingState ProviderAnalyticsDashboard) {
}

func (c ProviderAnalyticsDashboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderAnalyticsDashboard
// only implements ToObjectValue() and Type().
func (o ProviderAnalyticsDashboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderAnalyticsDashboard) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
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

func (newState *ProviderInfo) SyncFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncFieldsDuringRead(existingState ProviderInfo) {
}

func (c ProviderInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["business_contact_email"] = attrs["business_contact_email"].SetRequired()
	attrs["company_website_link"] = attrs["company_website_link"].SetOptional()
	attrs["dark_mode_icon_file_id"] = attrs["dark_mode_icon_file_id"].SetOptional()
	attrs["dark_mode_icon_file_path"] = attrs["dark_mode_icon_file_path"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["icon_file_id"] = attrs["icon_file_id"].SetOptional()
	attrs["icon_file_path"] = attrs["icon_file_path"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_featured"] = attrs["is_featured"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["privacy_policy_link"] = attrs["privacy_policy_link"].SetRequired()
	attrs["published_by"] = attrs["published_by"].SetOptional()
	attrs["support_contact_email"] = attrs["support_contact_email"].SetOptional()
	attrs["term_of_service_link"] = attrs["term_of_service_link"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo
// only implements ToObjectValue() and Type().
func (o ProviderInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ProviderInfo) Type(ctx context.Context) attr.Type {
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
	Cloud types.String `tfsdk:"cloud"`

	Region types.String `tfsdk:"region"`
}

func (newState *RegionInfo) SyncFieldsDuringCreateOrUpdate(plan RegionInfo) {
}

func (newState *RegionInfo) SyncFieldsDuringRead(existingState RegionInfo) {
}

func (c RegionInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegionInfo
// only implements ToObjectValue() and Type().
func (o RegionInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud":  o.Cloud,
			"region": o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegionInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud":  types.StringType,
			"region": types.StringType,
		},
	}
}

type RemoveExchangeForListingRequest struct {
	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingRequest
// only implements ToObjectValue() and Type().
func (o RemoveExchangeForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveExchangeForListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type RemoveExchangeForListingResponse struct {
}

func (newState *RemoveExchangeForListingResponse) SyncFieldsDuringCreateOrUpdate(plan RemoveExchangeForListingResponse) {
}

func (newState *RemoveExchangeForListingResponse) SyncFieldsDuringRead(existingState RemoveExchangeForListingResponse) {
}

func (c RemoveExchangeForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingResponse
// only implements ToObjectValue() and Type().
func (o RemoveExchangeForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveExchangeForListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url"`
}

func (newState *RepoInfo) SyncFieldsDuringCreateOrUpdate(plan RepoInfo) {
}

func (newState *RepoInfo) SyncFieldsDuringRead(existingState RepoInfo) {
}

func (c RepoInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["git_repo_url"] = attrs["git_repo_url"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo
// only implements ToObjectValue() and Type().
func (o RepoInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_repo_url": o.GitRepoUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"git_repo_url": types.StringType,
		},
	}
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName types.String `tfsdk:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path"`
}

func (newState *RepoInstallation) SyncFieldsDuringCreateOrUpdate(plan RepoInstallation) {
}

func (newState *RepoInstallation) SyncFieldsDuringRead(existingState RepoInstallation) {
}

func (c RepoInstallation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["repo_name"] = attrs["repo_name"].SetRequired()
	attrs["repo_path"] = attrs["repo_path"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInstallation
// only implements ToObjectValue() and Type().
func (o RepoInstallation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_name": o.RepoName,
			"repo_path": o.RepoPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoInstallation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_name": types.StringType,
			"repo_path": types.StringType,
		},
	}
}

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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsRequest
// only implements ToObjectValue() and Type().
func (o SearchListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchListingsRequest) Type(ctx context.Context) attr.Type {
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

// GetAssets returns the value of the Assets field in SearchListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest) GetAssets(ctx context.Context) ([]types.String, bool) {
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

// SetAssets sets the value of the Assets field in SearchListingsRequest.
func (o *SearchListingsRequest) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in SearchListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest) GetCategories(ctx context.Context) ([]types.String, bool) {
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

// SetCategories sets the value of the Categories field in SearchListingsRequest.
func (o *SearchListingsRequest) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in SearchListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsRequest) GetProviderIds(ctx context.Context) ([]types.String, bool) {
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

// SetProviderIds sets the value of the ProviderIds field in SearchListingsRequest.
func (o *SearchListingsRequest) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ProviderIds = types.ListValueMust(t, vs)
}

type SearchListingsResponse struct {
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *SearchListingsResponse) SyncFieldsDuringCreateOrUpdate(plan SearchListingsResponse) {
}

func (newState *SearchListingsResponse) SyncFieldsDuringRead(existingState SearchListingsResponse) {
}

func (c SearchListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listings"] = attrs["listings"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsResponse
// only implements ToObjectValue() and Type().
func (o SearchListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        o.Listings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchListingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listings": basetypes.ListType{
				ElemType: Listing{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetListings returns the value of the Listings field in SearchListingsResponse as
// a slice of Listing values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if o.Listings.IsNull() || o.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := o.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in SearchListingsResponse.
func (o *SearchListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Listings = types.ListValueMust(t, vs)
}

type ShareInfo struct {
	Name types.String `tfsdk:"name"`

	Type_ types.String `tfsdk:"type"`
}

func (newState *ShareInfo) SyncFieldsDuringCreateOrUpdate(plan ShareInfo) {
}

func (newState *ShareInfo) SyncFieldsDuringRead(existingState ShareInfo) {
}

func (c ShareInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["type"] = attrs["type"].SetRequired()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo
// only implements ToObjectValue() and Type().
func (o ShareInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
			"type": o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ShareInfo) Type(ctx context.Context) attr.Type {
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
	DataObjectType types.String `tfsdk:"data_object_type"`
	// Name of the shared object
	Name types.String `tfsdk:"name"`
}

func (newState *SharedDataObject) SyncFieldsDuringCreateOrUpdate(plan SharedDataObject) {
}

func (newState *SharedDataObject) SyncFieldsDuringRead(existingState SharedDataObject) {
}

func (c SharedDataObject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_object_type"] = attrs["data_object_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject
// only implements ToObjectValue() and Type().
func (o SharedDataObject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_object_type": o.DataObjectType,
			"name":             o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObject) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_object_type": types.StringType,
			"name":             types.StringType,
		},
	}
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

func (newState *TokenDetail) SyncFieldsDuringCreateOrUpdate(plan TokenDetail) {
}

func (newState *TokenDetail) SyncFieldsDuringRead(existingState TokenDetail) {
}

func (c TokenDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bearerToken"] = attrs["bearerToken"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetOptional()
	attrs["expirationTime"] = attrs["expirationTime"].SetOptional()
	attrs["shareCredentialsVersion"] = attrs["shareCredentialsVersion"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenDetail
// only implements ToObjectValue() and Type().
func (o TokenDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TokenDetail) Type(ctx context.Context) attr.Type {
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

func (newState *TokenInfo) SyncFieldsDuringCreateOrUpdate(plan TokenInfo) {
}

func (newState *TokenInfo) SyncFieldsDuringRead(existingState TokenInfo) {
}

func (c TokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activation_url"] = attrs["activation_url"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo
// only implements ToObjectValue() and Type().
func (o TokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TokenInfo) Type(ctx context.Context) attr.Type {
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
	Filter types.Object `tfsdk:"filter"`

	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (o UpdateExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
			"id":     o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeFilterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": ExchangeFilter{}.Type(ctx),
			"id":     types.StringType,
		},
	}
}

// GetFilter returns the value of the Filter field in UpdateExchangeFilterRequest as
// a ExchangeFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeFilterRequest) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterRequest.
func (o *UpdateExchangeFilterRequest) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type UpdateExchangeFilterResponse struct {
	Filter types.Object `tfsdk:"filter"`
}

func (newState *UpdateExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateExchangeFilterResponse) {
}

func (newState *UpdateExchangeFilterResponse) SyncFieldsDuringRead(existingState UpdateExchangeFilterResponse) {
}

func (c UpdateExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (o UpdateExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": o.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": ExchangeFilter{}.Type(ctx),
		},
	}
}

// GetFilter returns the value of the Filter field in UpdateExchangeFilterResponse as
// a ExchangeFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeFilterResponse) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if o.Filter.IsNull() || o.Filter.IsUnknown() {
		return e, false
	}
	var v []ExchangeFilter
	d := o.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterResponse.
func (o *UpdateExchangeFilterResponse) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	o.Filter = vs
}

type UpdateExchangeRequest struct {
	Exchange types.Object `tfsdk:"exchange"`

	Id types.String `tfsdk:"-"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeRequest
// only implements ToObjectValue() and Type().
func (o UpdateExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
			"id":       o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
			"id":       types.StringType,
		},
	}
}

// GetExchange returns the value of the Exchange field in UpdateExchangeRequest as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeRequest) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange
	d := o.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeRequest.
func (o *UpdateExchangeRequest) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	o.Exchange = vs
}

type UpdateExchangeResponse struct {
	Exchange types.Object `tfsdk:"exchange"`
}

func (newState *UpdateExchangeResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateExchangeResponse) {
}

func (newState *UpdateExchangeResponse) SyncFieldsDuringRead(existingState UpdateExchangeResponse) {
}

func (c UpdateExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeResponse
// only implements ToObjectValue() and Type().
func (o UpdateExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": o.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in UpdateExchangeResponse as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExchangeResponse) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if o.Exchange.IsNull() || o.Exchange.IsUnknown() {
		return e, false
	}
	var v []Exchange
	d := o.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeResponse.
func (o *UpdateExchangeResponse) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	o.Exchange = vs
}

type UpdateInstallationRequest struct {
	Installation types.Object `tfsdk:"installation"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationRequest
// only implements ToObjectValue() and Type().
func (o UpdateInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateInstallationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation":    InstallationDetail{}.Type(ctx),
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
			"rotate_token":    types.BoolType,
		},
	}
}

// GetInstallation returns the value of the Installation field in UpdateInstallationRequest as
// a InstallationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInstallationRequest) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail
	d := o.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationRequest.
func (o *UpdateInstallationRequest) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	o.Installation = vs
}

type UpdateInstallationResponse struct {
	Installation types.Object `tfsdk:"installation"`
}

func (newState *UpdateInstallationResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateInstallationResponse) {
}

func (newState *UpdateInstallationResponse) SyncFieldsDuringRead(existingState UpdateInstallationResponse) {
}

func (c UpdateInstallationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationResponse
// only implements ToObjectValue() and Type().
func (o UpdateInstallationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": o.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInstallationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": InstallationDetail{}.Type(ctx),
		},
	}
}

// GetInstallation returns the value of the Installation field in UpdateInstallationResponse as
// a InstallationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInstallationResponse) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if o.Installation.IsNull() || o.Installation.IsUnknown() {
		return e, false
	}
	var v []InstallationDetail
	d := o.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationResponse.
func (o *UpdateInstallationResponse) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	o.Installation = vs
}

type UpdateListingRequest struct {
	Id types.String `tfsdk:"-"`

	Listing types.Object `tfsdk:"listing"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingRequest
// only implements ToObjectValue() and Type().
func (o UpdateListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":      types.StringType,
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in UpdateListingRequest as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateListingRequest) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing
	d := o.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in UpdateListingRequest.
func (o *UpdateListingRequest) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	o.Listing = vs
}

type UpdateListingResponse struct {
	Listing types.Object `tfsdk:"listing"`
}

func (newState *UpdateListingResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateListingResponse) {
}

func (newState *UpdateListingResponse) SyncFieldsDuringRead(existingState UpdateListingResponse) {
}

func (c UpdateListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingResponse
// only implements ToObjectValue() and Type().
func (o UpdateListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": o.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in UpdateListingResponse as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateListingResponse) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if o.Listing.IsNull() || o.Listing.IsUnknown() {
		return e, false
	}
	var v []Listing
	d := o.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListing sets the value of the Listing field in UpdateListingResponse.
func (o *UpdateListingResponse) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	o.Listing = vs
}

type UpdatePersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`

	Reason types.String `tfsdk:"reason"`

	RequestId types.String `tfsdk:"-"`

	Share types.Object `tfsdk:"share"`

	Status types.String `tfsdk:"status"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestRequest
// only implements ToObjectValue() and Type().
func (o UpdatePersonalizationRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdatePersonalizationRequestRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
			"reason":     types.StringType,
			"request_id": types.StringType,
			"share":      ShareInfo{}.Type(ctx),
			"status":     types.StringType,
		},
	}
}

// GetShare returns the value of the Share field in UpdatePersonalizationRequestRequest as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePersonalizationRequestRequest) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if o.Share.IsNull() || o.Share.IsUnknown() {
		return e, false
	}
	var v []ShareInfo
	d := o.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in UpdatePersonalizationRequestRequest.
func (o *UpdatePersonalizationRequestRequest) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	o.Share = vs
}

type UpdatePersonalizationRequestResponse struct {
	Request types.Object `tfsdk:"request"`
}

func (newState *UpdatePersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(plan UpdatePersonalizationRequestResponse) {
}

func (newState *UpdatePersonalizationRequestResponse) SyncFieldsDuringRead(existingState UpdatePersonalizationRequestResponse) {
}

func (c UpdatePersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["request"] = attrs["request"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (o UpdatePersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": o.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": PersonalizationRequest{}.Type(ctx),
		},
	}
}

// GetRequest returns the value of the Request field in UpdatePersonalizationRequestResponse as
// a PersonalizationRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePersonalizationRequestResponse) GetRequest(ctx context.Context) (PersonalizationRequest, bool) {
	var e PersonalizationRequest
	if o.Request.IsNull() || o.Request.IsUnknown() {
		return e, false
	}
	var v []PersonalizationRequest
	d := o.Request.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRequest sets the value of the Request field in UpdatePersonalizationRequestResponse.
func (o *UpdatePersonalizationRequestResponse) SetRequest(ctx context.Context, v PersonalizationRequest) {
	vs := v.ToObjectValue(ctx)
	o.Request = vs
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (o UpdateProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":      types.StringType,
			"version": types.Int64Type,
		},
	}
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId types.String `tfsdk:"dashboard_id"`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

func (newState *UpdateProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateProviderAnalyticsDashboardResponse) {
}

func (newState *UpdateProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(existingState UpdateProviderAnalyticsDashboardResponse) {
}

func (c UpdateProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (o UpdateProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"id":           o.Id,
			"version":      o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
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

	Provider types.Object `tfsdk:"provider"`
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderRequest
// only implements ToObjectValue() and Type().
func (o UpdateProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":       o.Id,
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":       types.StringType,
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in UpdateProviderRequest as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateProviderRequest) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo
	d := o.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in UpdateProviderRequest.
func (o *UpdateProviderRequest) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	o.Provider = vs
}

type UpdateProviderResponse struct {
	Provider types.Object `tfsdk:"provider"`
}

func (newState *UpdateProviderResponse) SyncFieldsDuringCreateOrUpdate(plan UpdateProviderResponse) {
}

func (newState *UpdateProviderResponse) SyncFieldsDuringRead(existingState UpdateProviderResponse) {
}

func (c UpdateProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetOptional()

	return attrs
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderResponse
// only implements ToObjectValue() and Type().
func (o UpdateProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": o.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in UpdateProviderResponse as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateProviderResponse) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if o.Provider.IsNull() || o.Provider.IsUnknown() {
		return e, false
	}
	var v []ProviderInfo
	d := o.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvider sets the value of the Provider field in UpdateProviderResponse.
func (o *UpdateProviderResponse) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	o.Provider = vs
}
