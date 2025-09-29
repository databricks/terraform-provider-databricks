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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddExchangeForListingRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"exchange_id"`

	ListingId types.String `tfsdk:"listing_id"`
}

func (to *AddExchangeForListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddExchangeForListingRequest_SdkV2) {
}

func (to *AddExchangeForListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AddExchangeForListingRequest_SdkV2) {
}

func (c AddExchangeForListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_id"] = attrs["exchange_id"].SetRequired()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()

	return attrs
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
	ExchangeForListing types.List `tfsdk:"exchange_for_listing"`
}

func (to *AddExchangeForListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddExchangeForListingResponse_SdkV2) {
	if !from.ExchangeForListing.IsNull() && !from.ExchangeForListing.IsUnknown() {
		if toExchangeForListing, ok := to.GetExchangeForListing(ctx); ok {
			if fromExchangeForListing, ok := from.GetExchangeForListing(ctx); ok {
				// Recursively sync the fields of ExchangeForListing
				toExchangeForListing.SyncFieldsDuringCreateOrUpdate(ctx, fromExchangeForListing)
				to.SetExchangeForListing(ctx, toExchangeForListing)
			}
		}
	}
}

func (to *AddExchangeForListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AddExchangeForListingResponse_SdkV2) {
	if !from.ExchangeForListing.IsNull() && !from.ExchangeForListing.IsUnknown() {
		if toExchangeForListing, ok := to.GetExchangeForListing(ctx); ok {
			if fromExchangeForListing, ok := from.GetExchangeForListing(ctx); ok {
				toExchangeForListing.SyncFieldsDuringRead(ctx, fromExchangeForListing)
				to.SetExchangeForListing(ctx, toExchangeForListing)
			}
		}
	}
}

func (c AddExchangeForListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_for_listing"] = attrs["exchange_for_listing"].SetOptional()
	attrs["exchange_for_listing"] = attrs["exchange_for_listing"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ExchangeListing_SdkV2{}.Type(ctx),
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

type BatchGetListingsRequest_SdkV2 struct {
	Ids types.List `tfsdk:"-"`
}

func (to *BatchGetListingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetListingsRequest_SdkV2) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (to *BatchGetListingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BatchGetListingsRequest_SdkV2) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (c BatchGetListingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ids"] = attrs["ids"].SetOptional()

	return attrs
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
	Listings types.List `tfsdk:"listings"`
}

func (to *BatchGetListingsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *BatchGetListingsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BatchGetListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (c BatchGetListingsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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

type BatchGetProvidersRequest_SdkV2 struct {
	Ids types.List `tfsdk:"-"`
}

func (to *BatchGetProvidersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetProvidersRequest_SdkV2) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (to *BatchGetProvidersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BatchGetProvidersRequest_SdkV2) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (c BatchGetProvidersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ids"] = attrs["ids"].SetOptional()

	return attrs
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
	Providers types.List `tfsdk:"providers"`
}

func (to *BatchGetProvidersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (to *BatchGetProvidersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BatchGetProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (c BatchGetProvidersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
	Version types.String `tfsdk:"version"`
}

func (to *ConsumerTerms_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConsumerTerms_SdkV2) {
}

func (to *ConsumerTerms_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ConsumerTerms_SdkV2) {
}

func (c ConsumerTerms_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Company types.String `tfsdk:"company"`

	Email types.String `tfsdk:"email"`

	FirstName types.String `tfsdk:"first_name"`

	LastName types.String `tfsdk:"last_name"`
}

func (to *ContactInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ContactInfo_SdkV2) {
}

func (to *ContactInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ContactInfo_SdkV2) {
}

func (c ContactInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Filter types.List `tfsdk:"filter"`
}

func (to *CreateExchangeFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeFilterRequest_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *CreateExchangeFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeFilterRequest_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (c CreateExchangeFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetRequired()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ExchangeFilter_SdkV2{}.Type(ctx),
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
	FilterId types.String `tfsdk:"filter_id"`
}

func (to *CreateExchangeFilterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeFilterResponse_SdkV2) {
}

func (to *CreateExchangeFilterResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeFilterResponse_SdkV2) {
}

func (c CreateExchangeFilterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Exchange types.List `tfsdk:"exchange"`
}

func (to *CreateExchangeRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeRequest_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				// Recursively sync the fields of Exchange
				toExchange.SyncFieldsDuringCreateOrUpdate(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (to *CreateExchangeRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeRequest_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (c CreateExchangeRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetRequired()
	attrs["exchange"] = attrs["exchange"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Exchange_SdkV2{}.Type(ctx),
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
	ExchangeId types.String `tfsdk:"exchange_id"`
}

func (to *CreateExchangeResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeResponse_SdkV2) {
}

func (to *CreateExchangeResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeResponse_SdkV2) {
}

func (c CreateExchangeResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	DisplayName types.String `tfsdk:"display_name"`

	FileParent types.List `tfsdk:"file_parent"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`
}

func (to *CreateFileRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFileRequest_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				// Recursively sync the fields of FileParent
				toFileParent.SyncFieldsDuringCreateOrUpdate(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (to *CreateFileRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateFileRequest_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (c CreateFileRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["file_parent"] = attrs["file_parent"].SetRequired()
	attrs["file_parent"] = attrs["file_parent"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["marketplace_file_type"] = attrs["marketplace_file_type"].SetRequired()
	attrs["mime_type"] = attrs["mime_type"].SetRequired()

	return attrs
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
				ElemType: FileParent_SdkV2{}.Type(ctx),
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
	FileInfo types.List `tfsdk:"file_info"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url"`
}

func (to *CreateFileResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFileResponse_SdkV2) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				// Recursively sync the fields of FileInfo
				toFileInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (to *CreateFileResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateFileResponse_SdkV2) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				toFileInfo.SyncFieldsDuringRead(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (c CreateFileResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_info"] = attrs["file_info"].SetOptional()
	attrs["file_info"] = attrs["file_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
				ElemType: FileInfo_SdkV2{}.Type(ctx),
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
	AcceptedConsumerTerms types.List `tfsdk:"accepted_consumer_terms"`

	CatalogName types.String `tfsdk:"catalog_name"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type"`
	// for git repo installations
	RepoDetail types.List `tfsdk:"repo_detail"`

	ShareName types.String `tfsdk:"share_name"`
}

func (to *CreateInstallationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateInstallationRequest_SdkV2) {
	if !from.AcceptedConsumerTerms.IsNull() && !from.AcceptedConsumerTerms.IsUnknown() {
		if toAcceptedConsumerTerms, ok := to.GetAcceptedConsumerTerms(ctx); ok {
			if fromAcceptedConsumerTerms, ok := from.GetAcceptedConsumerTerms(ctx); ok {
				// Recursively sync the fields of AcceptedConsumerTerms
				toAcceptedConsumerTerms.SyncFieldsDuringCreateOrUpdate(ctx, fromAcceptedConsumerTerms)
				to.SetAcceptedConsumerTerms(ctx, toAcceptedConsumerTerms)
			}
		}
	}
	if !from.RepoDetail.IsNull() && !from.RepoDetail.IsUnknown() {
		if toRepoDetail, ok := to.GetRepoDetail(ctx); ok {
			if fromRepoDetail, ok := from.GetRepoDetail(ctx); ok {
				// Recursively sync the fields of RepoDetail
				toRepoDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromRepoDetail)
				to.SetRepoDetail(ctx, toRepoDetail)
			}
		}
	}
}

func (to *CreateInstallationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateInstallationRequest_SdkV2) {
	if !from.AcceptedConsumerTerms.IsNull() && !from.AcceptedConsumerTerms.IsUnknown() {
		if toAcceptedConsumerTerms, ok := to.GetAcceptedConsumerTerms(ctx); ok {
			if fromAcceptedConsumerTerms, ok := from.GetAcceptedConsumerTerms(ctx); ok {
				toAcceptedConsumerTerms.SyncFieldsDuringRead(ctx, fromAcceptedConsumerTerms)
				to.SetAcceptedConsumerTerms(ctx, toAcceptedConsumerTerms)
			}
		}
	}
	if !from.RepoDetail.IsNull() && !from.RepoDetail.IsUnknown() {
		if toRepoDetail, ok := to.GetRepoDetail(ctx); ok {
			if fromRepoDetail, ok := from.GetRepoDetail(ctx); ok {
				toRepoDetail.SyncFieldsDuringRead(ctx, fromRepoDetail)
				to.SetRepoDetail(ctx, toRepoDetail)
			}
		}
	}
}

func (c CreateInstallationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].SetOptional()
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["repo_detail"] = attrs["repo_detail"].SetOptional()
	attrs["repo_detail"] = attrs["repo_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["share_name"] = attrs["share_name"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()

	return attrs
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
				ElemType: ConsumerTerms_SdkV2{}.Type(ctx),
			},
			"catalog_name":   types.StringType,
			"listing_id":     types.StringType,
			"recipient_type": types.StringType,
			"repo_detail": basetypes.ListType{
				ElemType: RepoInstallation_SdkV2{}.Type(ctx),
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
	Listing types.List `tfsdk:"listing"`
}

func (to *CreateListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateListingRequest_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				// Recursively sync the fields of Listing
				toListing.SyncFieldsDuringCreateOrUpdate(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (to *CreateListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateListingRequest_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (c CreateListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetRequired()
	attrs["listing"] = attrs["listing"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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
	ListingId types.String `tfsdk:"listing_id"`
}

func (to *CreateListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateListingResponse_SdkV2) {
}

func (to *CreateListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateListingResponse_SdkV2) {
}

func (c CreateListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AcceptedConsumerTerms types.List `tfsdk:"accepted_consumer_terms"`

	Comment types.String `tfsdk:"comment"`

	Company types.String `tfsdk:"company"`

	FirstName types.String `tfsdk:"first_name"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	LastName types.String `tfsdk:"last_name"`

	ListingId types.String `tfsdk:"-"`

	RecipientType types.String `tfsdk:"recipient_type"`
}

func (to *CreatePersonalizationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePersonalizationRequest_SdkV2) {
	if !from.AcceptedConsumerTerms.IsNull() && !from.AcceptedConsumerTerms.IsUnknown() {
		if toAcceptedConsumerTerms, ok := to.GetAcceptedConsumerTerms(ctx); ok {
			if fromAcceptedConsumerTerms, ok := from.GetAcceptedConsumerTerms(ctx); ok {
				// Recursively sync the fields of AcceptedConsumerTerms
				toAcceptedConsumerTerms.SyncFieldsDuringCreateOrUpdate(ctx, fromAcceptedConsumerTerms)
				to.SetAcceptedConsumerTerms(ctx, toAcceptedConsumerTerms)
			}
		}
	}
}

func (to *CreatePersonalizationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePersonalizationRequest_SdkV2) {
	if !from.AcceptedConsumerTerms.IsNull() && !from.AcceptedConsumerTerms.IsUnknown() {
		if toAcceptedConsumerTerms, ok := to.GetAcceptedConsumerTerms(ctx); ok {
			if fromAcceptedConsumerTerms, ok := from.GetAcceptedConsumerTerms(ctx); ok {
				toAcceptedConsumerTerms.SyncFieldsDuringRead(ctx, fromAcceptedConsumerTerms)
				to.SetAcceptedConsumerTerms(ctx, toAcceptedConsumerTerms)
			}
		}
	}
}

func (c CreatePersonalizationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].SetRequired()
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["company"] = attrs["company"].SetOptional()
	attrs["first_name"] = attrs["first_name"].SetOptional()
	attrs["intended_use"] = attrs["intended_use"].SetRequired()
	attrs["is_from_lighthouse"] = attrs["is_from_lighthouse"].SetOptional()
	attrs["last_name"] = attrs["last_name"].SetOptional()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()

	return attrs
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
				ElemType: ConsumerTerms_SdkV2{}.Type(ctx),
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
	Id types.String `tfsdk:"id"`
}

func (to *CreatePersonalizationRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePersonalizationRequestResponse_SdkV2) {
}

func (to *CreatePersonalizationRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePersonalizationRequestResponse_SdkV2) {
}

func (c CreatePersonalizationRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type CreateProviderAnalyticsDashboardRequest_SdkV2 struct {
}

func (to *CreateProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderAnalyticsDashboardRequest_SdkV2) {
}

func (to *CreateProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProviderAnalyticsDashboardRequest_SdkV2) {
}

func (c CreateProviderAnalyticsDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProviderAnalyticsDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderAnalyticsDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateProviderAnalyticsDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProviderAnalyticsDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateProviderRequest_SdkV2 struct {
	Provider types.List `tfsdk:"provider"`
}

func (to *CreateProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderRequest_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				// Recursively sync the fields of Provider
				toProvider.SyncFieldsDuringCreateOrUpdate(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (to *CreateProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProviderRequest_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (c CreateProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["provider"] = attrs["provider"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
	Id types.String `tfsdk:"id"`
}

func (to *CreateProviderResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderResponse_SdkV2) {
}

func (to *CreateProviderResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProviderResponse_SdkV2) {
}

func (c CreateProviderResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Interval types.Int64 `tfsdk:"interval"`

	Unit types.String `tfsdk:"unit"`
}

func (to *DataRefreshInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataRefreshInfo_SdkV2) {
}

func (to *DataRefreshInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataRefreshInfo_SdkV2) {
}

func (c DataRefreshInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteExchangeFilterRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteExchangeFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeFilterRequest_SdkV2) {
}

func (to *DeleteExchangeFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeFilterRequest_SdkV2) {
}

func (c DeleteExchangeFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteExchangeFilterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeFilterResponse_SdkV2) {
}

func (to *DeleteExchangeFilterResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeFilterResponse_SdkV2) {
}

func (c DeleteExchangeFilterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteExchangeRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteExchangeRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeRequest_SdkV2) {
}

func (to *DeleteExchangeRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeRequest_SdkV2) {
}

func (c DeleteExchangeRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteExchangeResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeResponse_SdkV2) {
}

func (to *DeleteExchangeResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeResponse_SdkV2) {
}

func (c DeleteExchangeResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteFileRequest_SdkV2 struct {
	FileId types.String `tfsdk:"-"`
}

func (to *DeleteFileRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFileRequest_SdkV2) {
}

func (to *DeleteFileRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFileRequest_SdkV2) {
}

func (c DeleteFileRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_id"] = attrs["file_id"].SetRequired()

	return attrs
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

func (to *DeleteFileResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFileResponse_SdkV2) {
}

func (to *DeleteFileResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFileResponse_SdkV2) {
}

func (c DeleteFileResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteInstallationRequest_SdkV2 struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
}

func (to *DeleteInstallationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstallationRequest_SdkV2) {
}

func (to *DeleteInstallationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteInstallationRequest_SdkV2) {
}

func (c DeleteInstallationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["installation_id"] = attrs["installation_id"].SetRequired()

	return attrs
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

func (to *DeleteInstallationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstallationResponse_SdkV2) {
}

func (to *DeleteInstallationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteInstallationResponse_SdkV2) {
}

func (c DeleteInstallationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteListingRequest_SdkV2) {
}

func (to *DeleteListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteListingRequest_SdkV2) {
}

func (c DeleteListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteListingResponse_SdkV2) {
}

func (to *DeleteListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteListingResponse_SdkV2) {
}

func (c DeleteListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteProviderRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProviderRequest_SdkV2) {
}

func (to *DeleteProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteProviderRequest_SdkV2) {
}

func (c DeleteProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteProviderResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProviderResponse_SdkV2) {
}

func (to *DeleteProviderResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteProviderResponse_SdkV2) {
}

func (c DeleteProviderResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (to *Exchange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Exchange_SdkV2) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
	if !from.LinkedListings.IsNull() && !from.LinkedListings.IsUnknown() && to.LinkedListings.IsNull() && len(from.LinkedListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LinkedListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LinkedListings = from.LinkedListings
	}
}

func (to *Exchange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Exchange_SdkV2) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
	if !from.LinkedListings.IsNull() && !from.LinkedListings.IsUnknown() && to.LinkedListings.IsNull() && len(from.LinkedListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LinkedListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LinkedListings = from.LinkedListings
	}
}

func (c Exchange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ExchangeFilter_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"linked_listings": basetypes.ListType{
				ElemType: ExchangeListing_SdkV2{}.Type(ctx),
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

func (to *ExchangeFilter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeFilter_SdkV2) {
}

func (to *ExchangeFilter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExchangeFilter_SdkV2) {
}

func (c ExchangeFilter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	ExchangeId types.String `tfsdk:"exchange_id"`

	ExchangeName types.String `tfsdk:"exchange_name"`

	Id types.String `tfsdk:"id"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`
}

func (to *ExchangeListing_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeListing_SdkV2) {
}

func (to *ExchangeListing_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExchangeListing_SdkV2) {
}

func (c ExchangeListing_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName types.String `tfsdk:"display_name"`

	DownloadLink types.String `tfsdk:"download_link"`

	FileParent types.List `tfsdk:"file_parent"`

	Id types.String `tfsdk:"id"`

	MarketplaceFileType types.String `tfsdk:"marketplace_file_type"`

	MimeType types.String `tfsdk:"mime_type"`

	Status types.String `tfsdk:"status"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (to *FileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileInfo_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				// Recursively sync the fields of FileParent
				toFileParent.SyncFieldsDuringCreateOrUpdate(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (to *FileInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileInfo_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (c FileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["download_link"] = attrs["download_link"].SetOptional()
	attrs["file_parent"] = attrs["file_parent"].SetOptional()
	attrs["file_parent"] = attrs["file_parent"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
				ElemType: FileParent_SdkV2{}.Type(ctx),
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
	FileParentType types.String `tfsdk:"file_parent_type"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id"`
}

func (to *FileParent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileParent_SdkV2) {
}

func (to *FileParent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileParent_SdkV2) {
}

func (c FileParent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetExchangeRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetExchangeRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExchangeRequest_SdkV2) {
}

func (to *GetExchangeRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExchangeRequest_SdkV2) {
}

func (c GetExchangeRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	Exchange types.List `tfsdk:"exchange"`
}

func (to *GetExchangeResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExchangeResponse_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				// Recursively sync the fields of Exchange
				toExchange.SyncFieldsDuringCreateOrUpdate(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (to *GetExchangeResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExchangeResponse_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (c GetExchangeResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetOptional()
	attrs["exchange"] = attrs["exchange"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Exchange_SdkV2{}.Type(ctx),
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

type GetFileRequest_SdkV2 struct {
	FileId types.String `tfsdk:"-"`
}

func (to *GetFileRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFileRequest_SdkV2) {
}

func (to *GetFileRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFileRequest_SdkV2) {
}

func (c GetFileRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_id"] = attrs["file_id"].SetRequired()

	return attrs
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
	FileInfo types.List `tfsdk:"file_info"`
}

func (to *GetFileResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFileResponse_SdkV2) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				// Recursively sync the fields of FileInfo
				toFileInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (to *GetFileResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFileResponse_SdkV2) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				toFileInfo.SyncFieldsDuringRead(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (c GetFileResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_info"] = attrs["file_info"].SetOptional()
	attrs["file_info"] = attrs["file_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: FileInfo_SdkV2{}.Type(ctx),
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

type GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2 struct {
}

func (to *GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) {
}

func (to *GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) {
}

func (c GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionProviderAnalyticsDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2 struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

func (to *GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) {
}

func (to *GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) {
}

func (c GetLatestVersionProviderAnalyticsDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetListingContentMetadataRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *GetListingContentMetadataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingContentMetadataRequest_SdkV2) {
}

func (to *GetListingContentMetadataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingContentMetadataRequest_SdkV2) {
}

func (c GetListingContentMetadataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	SharedDataObjects types.List `tfsdk:"shared_data_objects"`
}

func (to *GetListingContentMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingContentMetadataResponse_SdkV2) {
	if !from.SharedDataObjects.IsNull() && !from.SharedDataObjects.IsUnknown() && to.SharedDataObjects.IsNull() && len(from.SharedDataObjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SharedDataObjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SharedDataObjects = from.SharedDataObjects
	}
}

func (to *GetListingContentMetadataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingContentMetadataResponse_SdkV2) {
	if !from.SharedDataObjects.IsNull() && !from.SharedDataObjects.IsUnknown() && to.SharedDataObjects.IsNull() && len(from.SharedDataObjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SharedDataObjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SharedDataObjects = from.SharedDataObjects
	}
}

func (c GetListingContentMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: SharedDataObject_SdkV2{}.Type(ctx),
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

type GetListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingRequest_SdkV2) {
}

func (to *GetListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingRequest_SdkV2) {
}

func (c GetListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	Listing types.List `tfsdk:"listing"`
}

func (to *GetListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingResponse_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				// Recursively sync the fields of Listing
				toListing.SyncFieldsDuringCreateOrUpdate(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (to *GetListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingResponse_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (c GetListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetOptional()
	attrs["listing"] = attrs["listing"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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

type GetListingsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *GetListingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingsRequest_SdkV2) {
}

func (to *GetListingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingsRequest_SdkV2) {
}

func (c GetListingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetListingsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *GetListingsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (c GetListingsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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

type GetPersonalizationRequestRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`
}

func (to *GetPersonalizationRequestRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalizationRequestRequest_SdkV2) {
}

func (to *GetPersonalizationRequestRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPersonalizationRequestRequest_SdkV2) {
}

func (c GetPersonalizationRequestRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()

	return attrs
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
	PersonalizationRequests types.List `tfsdk:"personalization_requests"`
}

func (to *GetPersonalizationRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalizationRequestResponse_SdkV2) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (to *GetPersonalizationRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPersonalizationRequestResponse_SdkV2) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (c GetPersonalizationRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: PersonalizationRequest_SdkV2{}.Type(ctx),
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

type GetProviderRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProviderRequest_SdkV2) {
}

func (to *GetProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetProviderRequest_SdkV2) {
}

func (c GetProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	Provider types.List `tfsdk:"provider"`
}

func (to *GetProviderResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProviderResponse_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				// Recursively sync the fields of Provider
				toProvider.SyncFieldsDuringCreateOrUpdate(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (to *GetProviderResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetProviderResponse_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (c GetProviderResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["provider"] = attrs["provider"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
	Installation types.List `tfsdk:"installation"`
}

func (to *Installation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Installation_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				// Recursively sync the fields of Installation
				toInstallation.SyncFieldsDuringCreateOrUpdate(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (to *Installation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Installation_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (c Installation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetOptional()
	attrs["installation"] = attrs["installation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: InstallationDetail_SdkV2{}.Type(ctx),
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

	TokenDetail types.List `tfsdk:"token_detail"`

	Tokens types.List `tfsdk:"tokens"`
}

func (to *InstallationDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstallationDetail_SdkV2) {
	if !from.TokenDetail.IsNull() && !from.TokenDetail.IsUnknown() {
		if toTokenDetail, ok := to.GetTokenDetail(ctx); ok {
			if fromTokenDetail, ok := from.GetTokenDetail(ctx); ok {
				// Recursively sync the fields of TokenDetail
				toTokenDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenDetail)
				to.SetTokenDetail(ctx, toTokenDetail)
			}
		}
	}
	if !from.Tokens.IsNull() && !from.Tokens.IsUnknown() && to.Tokens.IsNull() && len(from.Tokens.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tokens, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tokens = from.Tokens
	}
}

func (to *InstallationDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from InstallationDetail_SdkV2) {
	if !from.TokenDetail.IsNull() && !from.TokenDetail.IsUnknown() {
		if toTokenDetail, ok := to.GetTokenDetail(ctx); ok {
			if fromTokenDetail, ok := from.GetTokenDetail(ctx); ok {
				toTokenDetail.SyncFieldsDuringRead(ctx, fromTokenDetail)
				to.SetTokenDetail(ctx, toTokenDetail)
			}
		}
	}
	if !from.Tokens.IsNull() && !from.Tokens.IsUnknown() && to.Tokens.IsNull() && len(from.Tokens.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tokens, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tokens = from.Tokens
	}
}

func (c InstallationDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["token_detail"] = attrs["token_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
				ElemType: TokenDetail_SdkV2{}.Type(ctx),
			},
			"tokens": basetypes.ListType{
				ElemType: TokenInfo_SdkV2{}.Type(ctx),
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

type ListAllInstallationsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAllInstallationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllInstallationsRequest_SdkV2) {
}

func (to *ListAllInstallationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAllInstallationsRequest_SdkV2) {
}

func (c ListAllInstallationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Installations types.List `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListAllInstallationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllInstallationsResponse_SdkV2) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (to *ListAllInstallationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAllInstallationsResponse_SdkV2) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (c ListAllInstallationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: InstallationDetail_SdkV2{}.Type(ctx),
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

type ListAllPersonalizationRequestsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAllPersonalizationRequestsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllPersonalizationRequestsRequest_SdkV2) {
}

func (to *ListAllPersonalizationRequestsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAllPersonalizationRequestsRequest_SdkV2) {
}

func (c ListAllPersonalizationRequestsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	PersonalizationRequests types.List `tfsdk:"personalization_requests"`
}

func (to *ListAllPersonalizationRequestsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllPersonalizationRequestsResponse_SdkV2) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (to *ListAllPersonalizationRequestsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAllPersonalizationRequestsResponse_SdkV2) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (c ListAllPersonalizationRequestsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: PersonalizationRequest_SdkV2{}.Type(ctx),
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

type ListConsumerProvidersRequest_SdkV2 struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListConsumerProvidersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListConsumerProvidersRequest_SdkV2) {
}

func (to *ListConsumerProvidersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListConsumerProvidersRequest_SdkV2) {
}

func (c ListConsumerProvidersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["is_featured"] = attrs["is_featured"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListConsumerProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListConsumerProvidersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConsumerProvidersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListConsumerProvidersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_featured": o.IsFeatured,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListConsumerProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_featured": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListExchangeFiltersRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangeFiltersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangeFiltersRequest_SdkV2) {
}

func (to *ListExchangeFiltersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangeFiltersRequest_SdkV2) {
}

func (c ListExchangeFiltersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_id"] = attrs["exchange_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Filters types.List `tfsdk:"filters"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExchangeFiltersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangeFiltersResponse_SdkV2) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
}

func (to *ListExchangeFiltersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangeFiltersResponse_SdkV2) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
}

func (c ListExchangeFiltersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ExchangeFilter_SdkV2{}.Type(ctx),
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

type ListExchangesForListingRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangesForListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesForListingRequest_SdkV2) {
}

func (to *ListExchangesForListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangesForListingRequest_SdkV2) {
}

func (c ListExchangesForListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	ExchangeListing types.List `tfsdk:"exchange_listing"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExchangesForListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesForListingResponse_SdkV2) {
	if !from.ExchangeListing.IsNull() && !from.ExchangeListing.IsUnknown() && to.ExchangeListing.IsNull() && len(from.ExchangeListing.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListing, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListing = from.ExchangeListing
	}
}

func (to *ListExchangesForListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangesForListingResponse_SdkV2) {
	if !from.ExchangeListing.IsNull() && !from.ExchangeListing.IsUnknown() && to.ExchangeListing.IsNull() && len(from.ExchangeListing.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListing, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListing = from.ExchangeListing
	}
}

func (c ListExchangesForListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ExchangeListing_SdkV2{}.Type(ctx),
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

type ListExchangesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesRequest_SdkV2) {
}

func (to *ListExchangesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangesRequest_SdkV2) {
}

func (c ListExchangesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Exchanges types.List `tfsdk:"exchanges"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExchangesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesResponse_SdkV2) {
	if !from.Exchanges.IsNull() && !from.Exchanges.IsUnknown() && to.Exchanges.IsNull() && len(from.Exchanges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exchanges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exchanges = from.Exchanges
	}
}

func (to *ListExchangesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExchangesResponse_SdkV2) {
	if !from.Exchanges.IsNull() && !from.Exchanges.IsUnknown() && to.Exchanges.IsNull() && len(from.Exchanges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exchanges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exchanges = from.Exchanges
	}
}

func (c ListExchangesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: Exchange_SdkV2{}.Type(ctx),
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

type ListFilesRequest_SdkV2 struct {
	FileParent types.List `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListFilesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFilesRequest_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				// Recursively sync the fields of FileParent
				toFileParent.SyncFieldsDuringCreateOrUpdate(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (to *ListFilesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFilesRequest_SdkV2) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (c ListFilesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_parent"] = attrs["file_parent"].SetRequired()
	attrs["file_parent"] = attrs["file_parent"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
				ElemType: FileParent_SdkV2{}.Type(ctx),
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
	FileInfos types.List `tfsdk:"file_infos"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFilesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFilesResponse_SdkV2) {
	if !from.FileInfos.IsNull() && !from.FileInfos.IsUnknown() && to.FileInfos.IsNull() && len(from.FileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileInfos = from.FileInfos
	}
}

func (to *ListFilesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFilesResponse_SdkV2) {
	if !from.FileInfos.IsNull() && !from.FileInfos.IsUnknown() && to.FileInfos.IsNull() && len(from.FileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileInfos = from.FileInfos
	}
}

func (c ListFilesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: FileInfo_SdkV2{}.Type(ctx),
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

type ListFulfillmentsRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListFulfillmentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFulfillmentsRequest_SdkV2) {
}

func (to *ListFulfillmentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFulfillmentsRequest_SdkV2) {
}

func (c ListFulfillmentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Fulfillments types.List `tfsdk:"fulfillments"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFulfillmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFulfillmentsResponse_SdkV2) {
	if !from.Fulfillments.IsNull() && !from.Fulfillments.IsUnknown() && to.Fulfillments.IsNull() && len(from.Fulfillments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fulfillments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fulfillments = from.Fulfillments
	}
}

func (to *ListFulfillmentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFulfillmentsResponse_SdkV2) {
	if !from.Fulfillments.IsNull() && !from.Fulfillments.IsUnknown() && to.Fulfillments.IsNull() && len(from.Fulfillments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fulfillments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fulfillments = from.Fulfillments
	}
}

func (c ListFulfillmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ListingFulfillment_SdkV2{}.Type(ctx),
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

type ListInstallationsRequest_SdkV2 struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListInstallationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstallationsRequest_SdkV2) {
}

func (to *ListInstallationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListInstallationsRequest_SdkV2) {
}

func (c ListInstallationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Installations types.List `tfsdk:"installations"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListInstallationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstallationsResponse_SdkV2) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (to *ListInstallationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListInstallationsResponse_SdkV2) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (c ListInstallationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: InstallationDetail_SdkV2{}.Type(ctx),
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

type ListListingsForExchangeRequest_SdkV2 struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListListingsForExchangeRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsForExchangeRequest_SdkV2) {
}

func (to *ListListingsForExchangeRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListListingsForExchangeRequest_SdkV2) {
}

func (c ListListingsForExchangeRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange_id"] = attrs["exchange_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	ExchangeListings types.List `tfsdk:"exchange_listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListListingsForExchangeResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsForExchangeResponse_SdkV2) {
	if !from.ExchangeListings.IsNull() && !from.ExchangeListings.IsUnknown() && to.ExchangeListings.IsNull() && len(from.ExchangeListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListings = from.ExchangeListings
	}
}

func (to *ListListingsForExchangeResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListListingsForExchangeResponse_SdkV2) {
	if !from.ExchangeListings.IsNull() && !from.ExchangeListings.IsUnknown() && to.ExchangeListings.IsNull() && len(from.ExchangeListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListings = from.ExchangeListings
	}
}

func (c ListListingsForExchangeResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ExchangeListing_SdkV2{}.Type(ctx),
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

func (to *ListListingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsRequest_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ProviderIds.IsNull() && !from.ProviderIds.IsUnknown() && to.ProviderIds.IsNull() && len(from.ProviderIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProviderIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProviderIds = from.ProviderIds
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ListListingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListListingsRequest_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ProviderIds.IsNull() && !from.ProviderIds.IsUnknown() && to.ProviderIds.IsNull() && len(from.ProviderIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProviderIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProviderIds = from.ProviderIds
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c ListListingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["assets"] = attrs["assets"].SetOptional()
	attrs["categories"] = attrs["categories"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["is_free"] = attrs["is_free"].SetOptional()
	attrs["is_private_exchange"] = attrs["is_private_exchange"].SetOptional()
	attrs["is_staff_pick"] = attrs["is_staff_pick"].SetOptional()
	attrs["provider_ids"] = attrs["provider_ids"].SetOptional()

	return attrs
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
				ElemType: ListingTag_SdkV2{}.Type(ctx),
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
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListListingsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *ListListingsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (c ListListingsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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

type ListProviderAnalyticsDashboardRequest_SdkV2 struct {
}

func (to *ListProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderAnalyticsDashboardRequest_SdkV2) {
}

func (to *ListProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProviderAnalyticsDashboardRequest_SdkV2) {
}

func (c ListProviderAnalyticsDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderAnalyticsDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProviderAnalyticsDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderAnalyticsDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListProviderAnalyticsDashboardResponse_SdkV2 struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`

	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

func (to *ListProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderAnalyticsDashboardResponse_SdkV2) {
}

func (to *ListProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProviderAnalyticsDashboardResponse_SdkV2) {
}

func (c ListProviderAnalyticsDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type ListProvidersRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListProvidersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersRequest_SdkV2) {
}

func (to *ListProvidersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProvidersRequest_SdkV2) {
}

func (c ListProvidersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProvidersResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Providers types.List `tfsdk:"providers"`
}

func (to *ListProvidersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (to *ListProvidersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (c ListProvidersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
	Detail types.List `tfsdk:"detail"`

	Id types.String `tfsdk:"id"`

	Summary types.List `tfsdk:"summary"`
}

func (to *Listing_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Listing_SdkV2) {
	if !from.Detail.IsNull() && !from.Detail.IsUnknown() {
		if toDetail, ok := to.GetDetail(ctx); ok {
			if fromDetail, ok := from.GetDetail(ctx); ok {
				// Recursively sync the fields of Detail
				toDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromDetail)
				to.SetDetail(ctx, toDetail)
			}
		}
	}
	if !from.Summary.IsNull() && !from.Summary.IsUnknown() {
		if toSummary, ok := to.GetSummary(ctx); ok {
			if fromSummary, ok := from.GetSummary(ctx); ok {
				// Recursively sync the fields of Summary
				toSummary.SyncFieldsDuringCreateOrUpdate(ctx, fromSummary)
				to.SetSummary(ctx, toSummary)
			}
		}
	}
}

func (to *Listing_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Listing_SdkV2) {
	if !from.Detail.IsNull() && !from.Detail.IsUnknown() {
		if toDetail, ok := to.GetDetail(ctx); ok {
			if fromDetail, ok := from.GetDetail(ctx); ok {
				toDetail.SyncFieldsDuringRead(ctx, fromDetail)
				to.SetDetail(ctx, toDetail)
			}
		}
	}
	if !from.Summary.IsNull() && !from.Summary.IsUnknown() {
		if toSummary, ok := to.GetSummary(ctx); ok {
			if fromSummary, ok := from.GetSummary(ctx); ok {
				toSummary.SyncFieldsDuringRead(ctx, fromSummary)
				to.SetSummary(ctx, toSummary)
			}
		}
	}
}

func (c Listing_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["detail"] = attrs["detail"].SetOptional()
	attrs["detail"] = attrs["detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["summary"] = attrs["summary"].SetRequired()
	attrs["summary"] = attrs["summary"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ListingDetail_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"summary": basetypes.ListType{
				ElemType: ListingSummary_SdkV2{}.Type(ctx),
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
	Assets types.List `tfsdk:"assets"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd types.Int64 `tfsdk:"collection_date_end"`
	// The starting date timestamp for when the data spans
	CollectionDateStart types.Int64 `tfsdk:"collection_date_start"`
	// Smallest unit of time in the dataset
	CollectionGranularity types.List `tfsdk:"collection_granularity"`
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
	UpdateFrequency types.List `tfsdk:"update_frequency"`
}

func (to *ListingDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingDetail_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.CollectionGranularity.IsNull() && !from.CollectionGranularity.IsUnknown() {
		if toCollectionGranularity, ok := to.GetCollectionGranularity(ctx); ok {
			if fromCollectionGranularity, ok := from.GetCollectionGranularity(ctx); ok {
				// Recursively sync the fields of CollectionGranularity
				toCollectionGranularity.SyncFieldsDuringCreateOrUpdate(ctx, fromCollectionGranularity)
				to.SetCollectionGranularity(ctx, toCollectionGranularity)
			}
		}
	}
	if !from.EmbeddedNotebookFileInfos.IsNull() && !from.EmbeddedNotebookFileInfos.IsUnknown() && to.EmbeddedNotebookFileInfos.IsNull() && len(from.EmbeddedNotebookFileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddedNotebookFileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddedNotebookFileInfos = from.EmbeddedNotebookFileInfos
	}
	if !from.FileIds.IsNull() && !from.FileIds.IsUnknown() && to.FileIds.IsNull() && len(from.FileIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileIds = from.FileIds
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.UpdateFrequency.IsNull() && !from.UpdateFrequency.IsUnknown() {
		if toUpdateFrequency, ok := to.GetUpdateFrequency(ctx); ok {
			if fromUpdateFrequency, ok := from.GetUpdateFrequency(ctx); ok {
				// Recursively sync the fields of UpdateFrequency
				toUpdateFrequency.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdateFrequency)
				to.SetUpdateFrequency(ctx, toUpdateFrequency)
			}
		}
	}
}

func (to *ListingDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListingDetail_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.CollectionGranularity.IsNull() && !from.CollectionGranularity.IsUnknown() {
		if toCollectionGranularity, ok := to.GetCollectionGranularity(ctx); ok {
			if fromCollectionGranularity, ok := from.GetCollectionGranularity(ctx); ok {
				toCollectionGranularity.SyncFieldsDuringRead(ctx, fromCollectionGranularity)
				to.SetCollectionGranularity(ctx, toCollectionGranularity)
			}
		}
	}
	if !from.EmbeddedNotebookFileInfos.IsNull() && !from.EmbeddedNotebookFileInfos.IsUnknown() && to.EmbeddedNotebookFileInfos.IsNull() && len(from.EmbeddedNotebookFileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddedNotebookFileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddedNotebookFileInfos = from.EmbeddedNotebookFileInfos
	}
	if !from.FileIds.IsNull() && !from.FileIds.IsUnknown() && to.FileIds.IsNull() && len(from.FileIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileIds = from.FileIds
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.UpdateFrequency.IsNull() && !from.UpdateFrequency.IsUnknown() {
		if toUpdateFrequency, ok := to.GetUpdateFrequency(ctx); ok {
			if fromUpdateFrequency, ok := from.GetUpdateFrequency(ctx); ok {
				toUpdateFrequency.SyncFieldsDuringRead(ctx, fromUpdateFrequency)
				to.SetUpdateFrequency(ctx, toUpdateFrequency)
			}
		}
	}
}

func (c ListingDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets"] = attrs["assets"].SetOptional()
	attrs["collection_date_end"] = attrs["collection_date_end"].SetOptional()
	attrs["collection_date_start"] = attrs["collection_date_start"].SetOptional()
	attrs["collection_granularity"] = attrs["collection_granularity"].SetOptional()
	attrs["collection_granularity"] = attrs["collection_granularity"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["update_frequency"] = attrs["update_frequency"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: DataRefreshInfo_SdkV2{}.Type(ctx),
			},
			"cost":               types.StringType,
			"data_source":        types.StringType,
			"description":        types.StringType,
			"documentation_link": types.StringType,
			"embedded_notebook_file_infos": basetypes.ListType{
				ElemType: FileInfo_SdkV2{}.Type(ctx),
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
				ElemType: ListingTag_SdkV2{}.Type(ctx),
			},
			"terms_of_service": types.StringType,
			"update_frequency": basetypes.ListType{
				ElemType: DataRefreshInfo_SdkV2{}.Type(ctx),
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
	FulfillmentType types.String `tfsdk:"fulfillment_type"`

	ListingId types.String `tfsdk:"listing_id"`

	RecipientType types.String `tfsdk:"recipient_type"`

	RepoInfo types.List `tfsdk:"repo_info"`

	ShareInfo types.List `tfsdk:"share_info"`
}

func (to *ListingFulfillment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingFulfillment_SdkV2) {
	if !from.RepoInfo.IsNull() && !from.RepoInfo.IsUnknown() {
		if toRepoInfo, ok := to.GetRepoInfo(ctx); ok {
			if fromRepoInfo, ok := from.GetRepoInfo(ctx); ok {
				// Recursively sync the fields of RepoInfo
				toRepoInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRepoInfo)
				to.SetRepoInfo(ctx, toRepoInfo)
			}
		}
	}
	if !from.ShareInfo.IsNull() && !from.ShareInfo.IsUnknown() {
		if toShareInfo, ok := to.GetShareInfo(ctx); ok {
			if fromShareInfo, ok := from.GetShareInfo(ctx); ok {
				// Recursively sync the fields of ShareInfo
				toShareInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromShareInfo)
				to.SetShareInfo(ctx, toShareInfo)
			}
		}
	}
}

func (to *ListingFulfillment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListingFulfillment_SdkV2) {
	if !from.RepoInfo.IsNull() && !from.RepoInfo.IsUnknown() {
		if toRepoInfo, ok := to.GetRepoInfo(ctx); ok {
			if fromRepoInfo, ok := from.GetRepoInfo(ctx); ok {
				toRepoInfo.SyncFieldsDuringRead(ctx, fromRepoInfo)
				to.SetRepoInfo(ctx, toRepoInfo)
			}
		}
	}
	if !from.ShareInfo.IsNull() && !from.ShareInfo.IsUnknown() {
		if toShareInfo, ok := to.GetShareInfo(ctx); ok {
			if fromShareInfo, ok := from.GetShareInfo(ctx); ok {
				toShareInfo.SyncFieldsDuringRead(ctx, fromShareInfo)
				to.SetShareInfo(ctx, toShareInfo)
			}
		}
	}
}

func (c ListingFulfillment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fulfillment_type"] = attrs["fulfillment_type"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["repo_info"] = attrs["repo_info"].SetOptional()
	attrs["repo_info"] = attrs["repo_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["share_info"] = attrs["share_info"].SetOptional()
	attrs["share_info"] = attrs["share_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: RepoInfo_SdkV2{}.Type(ctx),
			},
			"share_info": basetypes.ListType{
				ElemType: ShareInfo_SdkV2{}.Type(ctx),
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
	Visibility types.String `tfsdk:"visibility"`
}

func (to *ListingSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingSetting_SdkV2) {
}

func (to *ListingSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListingSetting_SdkV2) {
}

func (c ListingSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type ListingSummary_SdkV2 struct {
	Categories types.List `tfsdk:"categories"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	CreatedBy types.String `tfsdk:"created_by"`

	CreatedById types.Int64 `tfsdk:"created_by_id"`

	ExchangeIds types.List `tfsdk:"exchange_ids"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo types.List `tfsdk:"git_repo"`

	ListingType types.String `tfsdk:"listing_type"`

	Name types.String `tfsdk:"name"`

	ProviderId types.String `tfsdk:"provider_id"`

	ProviderRegion types.List `tfsdk:"provider_region"`

	PublishedAt types.Int64 `tfsdk:"published_at"`

	PublishedBy types.String `tfsdk:"published_by"`

	Setting types.List `tfsdk:"setting"`

	Share types.List `tfsdk:"share"`

	Status types.String `tfsdk:"status"`

	Subtitle types.String `tfsdk:"subtitle"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`

	UpdatedBy types.String `tfsdk:"updated_by"`

	UpdatedById types.Int64 `tfsdk:"updated_by_id"`
}

func (to *ListingSummary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingSummary_SdkV2) {
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ExchangeIds.IsNull() && !from.ExchangeIds.IsUnknown() && to.ExchangeIds.IsNull() && len(from.ExchangeIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeIds = from.ExchangeIds
	}
	if !from.GitRepo.IsNull() && !from.GitRepo.IsUnknown() {
		if toGitRepo, ok := to.GetGitRepo(ctx); ok {
			if fromGitRepo, ok := from.GetGitRepo(ctx); ok {
				// Recursively sync the fields of GitRepo
				toGitRepo.SyncFieldsDuringCreateOrUpdate(ctx, fromGitRepo)
				to.SetGitRepo(ctx, toGitRepo)
			}
		}
	}
	if !from.ProviderRegion.IsNull() && !from.ProviderRegion.IsUnknown() {
		if toProviderRegion, ok := to.GetProviderRegion(ctx); ok {
			if fromProviderRegion, ok := from.GetProviderRegion(ctx); ok {
				// Recursively sync the fields of ProviderRegion
				toProviderRegion.SyncFieldsDuringCreateOrUpdate(ctx, fromProviderRegion)
				to.SetProviderRegion(ctx, toProviderRegion)
			}
		}
	}
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				// Recursively sync the fields of Share
				toShare.SyncFieldsDuringCreateOrUpdate(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (to *ListingSummary_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListingSummary_SdkV2) {
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ExchangeIds.IsNull() && !from.ExchangeIds.IsUnknown() && to.ExchangeIds.IsNull() && len(from.ExchangeIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeIds = from.ExchangeIds
	}
	if !from.GitRepo.IsNull() && !from.GitRepo.IsUnknown() {
		if toGitRepo, ok := to.GetGitRepo(ctx); ok {
			if fromGitRepo, ok := from.GetGitRepo(ctx); ok {
				toGitRepo.SyncFieldsDuringRead(ctx, fromGitRepo)
				to.SetGitRepo(ctx, toGitRepo)
			}
		}
	}
	if !from.ProviderRegion.IsNull() && !from.ProviderRegion.IsUnknown() {
		if toProviderRegion, ok := to.GetProviderRegion(ctx); ok {
			if fromProviderRegion, ok := from.GetProviderRegion(ctx); ok {
				toProviderRegion.SyncFieldsDuringRead(ctx, fromProviderRegion)
				to.SetProviderRegion(ctx, toProviderRegion)
			}
		}
	}
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				toShare.SyncFieldsDuringRead(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (c ListingSummary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["categories"] = attrs["categories"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["created_by_id"] = attrs["created_by_id"].SetOptional()
	attrs["exchange_ids"] = attrs["exchange_ids"].SetOptional()
	attrs["git_repo"] = attrs["git_repo"].SetOptional()
	attrs["git_repo"] = attrs["git_repo"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["listing_type"] = attrs["listing_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["provider_id"] = attrs["provider_id"].SetOptional()
	attrs["provider_region"] = attrs["provider_region"].SetOptional()
	attrs["provider_region"] = attrs["provider_region"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["published_at"] = attrs["published_at"].SetOptional()
	attrs["published_by"] = attrs["published_by"].SetOptional()
	attrs["setting"] = attrs["setting"].SetOptional()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share"] = attrs["share"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
			"listing_type":    o.ListingType,
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
				ElemType: RepoInfo_SdkV2{}.Type(ctx),
			},
			"listing_type": types.StringType,
			"name":         types.StringType,
			"provider_id":  types.StringType,
			"provider_region": basetypes.ListType{
				ElemType: RegionInfo_SdkV2{}.Type(ctx),
			},
			"published_at": types.Int64Type,
			"published_by": types.StringType,
			"setting": basetypes.ListType{
				ElemType: ListingSetting_SdkV2{}.Type(ctx),
			},
			"share": basetypes.ListType{
				ElemType: ShareInfo_SdkV2{}.Type(ctx),
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
	TagName types.String `tfsdk:"tag_name"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues types.List `tfsdk:"tag_values"`
}

func (to *ListingTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingTag_SdkV2) {
	if !from.TagValues.IsNull() && !from.TagValues.IsUnknown() && to.TagValues.IsNull() && len(from.TagValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagValues = from.TagValues
	}
}

func (to *ListingTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListingTag_SdkV2) {
	if !from.TagValues.IsNull() && !from.TagValues.IsUnknown() && to.TagValues.IsNull() && len(from.TagValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagValues = from.TagValues
	}
}

func (c ListingTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Comment types.String `tfsdk:"comment"`

	ConsumerRegion types.List `tfsdk:"consumer_region"`

	ContactInfo types.List `tfsdk:"contact_info"`

	CreatedAt types.Int64 `tfsdk:"created_at"`

	Id types.String `tfsdk:"id"`

	IntendedUse types.String `tfsdk:"intended_use"`

	IsFromLighthouse types.Bool `tfsdk:"is_from_lighthouse"`

	ListingId types.String `tfsdk:"listing_id"`

	ListingName types.String `tfsdk:"listing_name"`

	MetastoreId types.String `tfsdk:"metastore_id"`

	ProviderId types.String `tfsdk:"provider_id"`

	RecipientType types.String `tfsdk:"recipient_type"`

	Share types.List `tfsdk:"share"`

	Status types.String `tfsdk:"status"`

	StatusMessage types.String `tfsdk:"status_message"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (to *PersonalizationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalizationRequest_SdkV2) {
	if !from.ConsumerRegion.IsNull() && !from.ConsumerRegion.IsUnknown() {
		if toConsumerRegion, ok := to.GetConsumerRegion(ctx); ok {
			if fromConsumerRegion, ok := from.GetConsumerRegion(ctx); ok {
				// Recursively sync the fields of ConsumerRegion
				toConsumerRegion.SyncFieldsDuringCreateOrUpdate(ctx, fromConsumerRegion)
				to.SetConsumerRegion(ctx, toConsumerRegion)
			}
		}
	}
	if !from.ContactInfo.IsNull() && !from.ContactInfo.IsUnknown() {
		if toContactInfo, ok := to.GetContactInfo(ctx); ok {
			if fromContactInfo, ok := from.GetContactInfo(ctx); ok {
				// Recursively sync the fields of ContactInfo
				toContactInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromContactInfo)
				to.SetContactInfo(ctx, toContactInfo)
			}
		}
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				// Recursively sync the fields of Share
				toShare.SyncFieldsDuringCreateOrUpdate(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (to *PersonalizationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PersonalizationRequest_SdkV2) {
	if !from.ConsumerRegion.IsNull() && !from.ConsumerRegion.IsUnknown() {
		if toConsumerRegion, ok := to.GetConsumerRegion(ctx); ok {
			if fromConsumerRegion, ok := from.GetConsumerRegion(ctx); ok {
				toConsumerRegion.SyncFieldsDuringRead(ctx, fromConsumerRegion)
				to.SetConsumerRegion(ctx, toConsumerRegion)
			}
		}
	}
	if !from.ContactInfo.IsNull() && !from.ContactInfo.IsUnknown() {
		if toContactInfo, ok := to.GetContactInfo(ctx); ok {
			if fromContactInfo, ok := from.GetContactInfo(ctx); ok {
				toContactInfo.SyncFieldsDuringRead(ctx, fromContactInfo)
				to.SetContactInfo(ctx, toContactInfo)
			}
		}
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				toShare.SyncFieldsDuringRead(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (c PersonalizationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["consumer_region"] = attrs["consumer_region"].SetRequired()
	attrs["consumer_region"] = attrs["consumer_region"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["contact_info"] = attrs["contact_info"].SetOptional()
	attrs["contact_info"] = attrs["contact_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["share"] = attrs["share"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
				ElemType: RegionInfo_SdkV2{}.Type(ctx),
			},
			"contact_info": basetypes.ListType{
				ElemType: ContactInfo_SdkV2{}.Type(ctx),
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
				ElemType: ShareInfo_SdkV2{}.Type(ctx),
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
	Id types.String `tfsdk:"id"`
}

func (to *ProviderAnalyticsDashboard_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderAnalyticsDashboard_SdkV2) {
}

func (to *ProviderAnalyticsDashboard_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProviderAnalyticsDashboard_SdkV2) {
}

func (c ProviderAnalyticsDashboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *ProviderInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderInfo_SdkV2) {
}

func (to *ProviderInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProviderInfo_SdkV2) {
}

func (c ProviderInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Cloud types.String `tfsdk:"cloud"`

	Region types.String `tfsdk:"region"`
}

func (to *RegionInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegionInfo_SdkV2) {
}

func (to *RegionInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegionInfo_SdkV2) {
}

func (c RegionInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type RemoveExchangeForListingRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *RemoveExchangeForListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveExchangeForListingRequest_SdkV2) {
}

func (to *RemoveExchangeForListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RemoveExchangeForListingRequest_SdkV2) {
}

func (c RemoveExchangeForListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *RemoveExchangeForListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveExchangeForListingResponse_SdkV2) {
}

func (to *RemoveExchangeForListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RemoveExchangeForListingResponse_SdkV2) {
}

func (c RemoveExchangeForListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	GitRepoUrl types.String `tfsdk:"git_repo_url"`
}

func (to *RepoInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInfo_SdkV2) {
}

func (to *RepoInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoInfo_SdkV2) {
}

func (c RepoInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	RepoName types.String `tfsdk:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath types.String `tfsdk:"repo_path"`
}

func (to *RepoInstallation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInstallation_SdkV2) {
}

func (to *RepoInstallation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoInstallation_SdkV2) {
}

func (c RepoInstallation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *SearchListingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchListingsRequest_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ProviderIds.IsNull() && !from.ProviderIds.IsUnknown() && to.ProviderIds.IsNull() && len(from.ProviderIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProviderIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProviderIds = from.ProviderIds
	}
}

func (to *SearchListingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchListingsRequest_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
	if !from.Categories.IsNull() && !from.Categories.IsUnknown() && to.Categories.IsNull() && len(from.Categories.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Categories, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Categories = from.Categories
	}
	if !from.ProviderIds.IsNull() && !from.ProviderIds.IsUnknown() && to.ProviderIds.IsNull() && len(from.ProviderIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProviderIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProviderIds = from.ProviderIds
	}
}

func (c SearchListingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query"] = attrs["query"].SetRequired()
	attrs["is_free"] = attrs["is_free"].SetOptional()
	attrs["is_private_exchange"] = attrs["is_private_exchange"].SetOptional()
	attrs["provider_ids"] = attrs["provider_ids"].SetOptional()
	attrs["categories"] = attrs["categories"].SetOptional()
	attrs["assets"] = attrs["assets"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchListingsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *SearchListingsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchListingsResponse_SdkV2) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (c SearchListingsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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
	Name types.String `tfsdk:"name"`

	Type_ types.String `tfsdk:"type"`
}

func (to *ShareInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ShareInfo_SdkV2) {
}

func (to *ShareInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ShareInfo_SdkV2) {
}

func (c ShareInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	DataObjectType types.String `tfsdk:"data_object_type"`
	// Name of the shared object
	Name types.String `tfsdk:"name"`
}

func (to *SharedDataObject_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharedDataObject_SdkV2) {
}

func (to *SharedDataObject_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SharedDataObject_SdkV2) {
}

func (c SharedDataObject_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	BearerToken types.String `tfsdk:"bearer_token"`

	Endpoint types.String `tfsdk:"endpoint"`

	ExpirationTime types.String `tfsdk:"expiration_time"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

func (to *TokenDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenDetail_SdkV2) {
}

func (to *TokenDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenDetail_SdkV2) {
}

func (c TokenDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bearer_token"] = attrs["bearer_token"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["share_credentials_version"] = attrs["share_credentials_version"].SetOptional()

	return attrs
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
			"bearer_token":              o.BearerToken,
			"endpoint":                  o.Endpoint,
			"expiration_time":           o.ExpirationTime,
			"share_credentials_version": o.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"expiration_time":           types.StringType,
			"share_credentials_version": types.Int64Type,
		},
	}
}

type TokenInfo_SdkV2 struct {
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

func (to *TokenInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenInfo_SdkV2) {
}

func (to *TokenInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenInfo_SdkV2) {
}

func (c TokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Filter types.List `tfsdk:"filter"`

	Id types.String `tfsdk:"-"`
}

func (to *UpdateExchangeFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeFilterRequest_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *UpdateExchangeFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeFilterRequest_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (c UpdateExchangeFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetRequired()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
				ElemType: ExchangeFilter_SdkV2{}.Type(ctx),
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
	Filter types.List `tfsdk:"filter"`
}

func (to *UpdateExchangeFilterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeFilterResponse_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				// Recursively sync the fields of Filter
				toFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (to *UpdateExchangeFilterResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeFilterResponse_SdkV2) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (c UpdateExchangeFilterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["filter"] = attrs["filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ExchangeFilter_SdkV2{}.Type(ctx),
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
	Exchange types.List `tfsdk:"exchange"`

	Id types.String `tfsdk:"-"`
}

func (to *UpdateExchangeRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeRequest_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				// Recursively sync the fields of Exchange
				toExchange.SyncFieldsDuringCreateOrUpdate(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (to *UpdateExchangeRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeRequest_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (c UpdateExchangeRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetRequired()
	attrs["exchange"] = attrs["exchange"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
				ElemType: Exchange_SdkV2{}.Type(ctx),
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
	Exchange types.List `tfsdk:"exchange"`
}

func (to *UpdateExchangeResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeResponse_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				// Recursively sync the fields of Exchange
				toExchange.SyncFieldsDuringCreateOrUpdate(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (to *UpdateExchangeResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeResponse_SdkV2) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (c UpdateExchangeResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetOptional()
	attrs["exchange"] = attrs["exchange"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Exchange_SdkV2{}.Type(ctx),
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
	Installation types.List `tfsdk:"installation"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token"`
}

func (to *UpdateInstallationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInstallationRequest_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				// Recursively sync the fields of Installation
				toInstallation.SyncFieldsDuringCreateOrUpdate(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (to *UpdateInstallationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateInstallationRequest_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (c UpdateInstallationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetRequired()
	attrs["installation"] = attrs["installation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["rotate_token"] = attrs["rotate_token"].SetOptional()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["installation_id"] = attrs["installation_id"].SetRequired()

	return attrs
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
				ElemType: InstallationDetail_SdkV2{}.Type(ctx),
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
	Installation types.List `tfsdk:"installation"`
}

func (to *UpdateInstallationResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInstallationResponse_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				// Recursively sync the fields of Installation
				toInstallation.SyncFieldsDuringCreateOrUpdate(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (to *UpdateInstallationResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateInstallationResponse_SdkV2) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (c UpdateInstallationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetOptional()
	attrs["installation"] = attrs["installation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: InstallationDetail_SdkV2{}.Type(ctx),
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

	Listing types.List `tfsdk:"listing"`
}

func (to *UpdateListingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateListingRequest_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				// Recursively sync the fields of Listing
				toListing.SyncFieldsDuringCreateOrUpdate(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (to *UpdateListingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateListingRequest_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (c UpdateListingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetRequired()
	attrs["listing"] = attrs["listing"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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
	Listing types.List `tfsdk:"listing"`
}

func (to *UpdateListingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateListingResponse_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				// Recursively sync the fields of Listing
				toListing.SyncFieldsDuringCreateOrUpdate(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (to *UpdateListingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateListingResponse_SdkV2) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (c UpdateListingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetOptional()
	attrs["listing"] = attrs["listing"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: Listing_SdkV2{}.Type(ctx),
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

	Reason types.String `tfsdk:"reason"`

	RequestId types.String `tfsdk:"-"`

	Share types.List `tfsdk:"share"`

	Status types.String `tfsdk:"status"`
}

func (to *UpdatePersonalizationRequestRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalizationRequestRequest_SdkV2) {
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				// Recursively sync the fields of Share
				toShare.SyncFieldsDuringCreateOrUpdate(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (to *UpdatePersonalizationRequestRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalizationRequestRequest_SdkV2) {
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				toShare.SyncFieldsDuringRead(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (c UpdatePersonalizationRequestRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["reason"] = attrs["reason"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share"] = attrs["share"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetRequired()
	attrs["listing_id"] = attrs["listing_id"].SetRequired()
	attrs["request_id"] = attrs["request_id"].SetRequired()

	return attrs
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
				ElemType: ShareInfo_SdkV2{}.Type(ctx),
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
	Request types.List `tfsdk:"request"`
}

func (to *UpdatePersonalizationRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalizationRequestResponse_SdkV2) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				// Recursively sync the fields of Request
				toRequest.SyncFieldsDuringCreateOrUpdate(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (to *UpdatePersonalizationRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalizationRequestResponse_SdkV2) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				toRequest.SyncFieldsDuringRead(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (c UpdatePersonalizationRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["request"] = attrs["request"].SetOptional()
	attrs["request"] = attrs["request"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: PersonalizationRequest_SdkV2{}.Type(ctx),
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
	Version types.Int64 `tfsdk:"version"`
}

func (to *UpdateProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderAnalyticsDashboardRequest_SdkV2) {
}

func (to *UpdateProviderAnalyticsDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderAnalyticsDashboardRequest_SdkV2) {
}

func (c UpdateProviderAnalyticsDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["version"] = attrs["version"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	DashboardId types.String `tfsdk:"dashboard_id"`
	// id & version should be the same as the request
	Id types.String `tfsdk:"id"`

	Version types.Int64 `tfsdk:"version"`
}

func (to *UpdateProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderAnalyticsDashboardResponse_SdkV2) {
}

func (to *UpdateProviderAnalyticsDashboardResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderAnalyticsDashboardResponse_SdkV2) {
}

func (c UpdateProviderAnalyticsDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

	Provider types.List `tfsdk:"provider"`
}

func (to *UpdateProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderRequest_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				// Recursively sync the fields of Provider
				toProvider.SyncFieldsDuringCreateOrUpdate(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (to *UpdateProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderRequest_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (c UpdateProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["provider"] = attrs["provider"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
	Provider types.List `tfsdk:"provider"`
}

func (to *UpdateProviderResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderResponse_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				// Recursively sync the fields of Provider
				toProvider.SyncFieldsDuringCreateOrUpdate(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (to *UpdateProviderResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderResponse_SdkV2) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (c UpdateProviderResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["provider"] = attrs["provider"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
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
