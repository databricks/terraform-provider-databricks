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

func (to *AddExchangeForListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddExchangeForListingRequest) {
}

func (to *AddExchangeForListingRequest) SyncFieldsDuringRead(ctx context.Context, from AddExchangeForListingRequest) {
}

func (m AddExchangeForListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AddExchangeForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingRequest
// only implements ToObjectValue() and Type().
func (m AddExchangeForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": m.ExchangeId,
			"listing_id":  m.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddExchangeForListingRequest) Type(ctx context.Context) attr.Type {
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

func (to *AddExchangeForListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddExchangeForListingResponse) {
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

func (to *AddExchangeForListingResponse) SyncFieldsDuringRead(ctx context.Context, from AddExchangeForListingResponse) {
	if !from.ExchangeForListing.IsNull() && !from.ExchangeForListing.IsUnknown() {
		if toExchangeForListing, ok := to.GetExchangeForListing(ctx); ok {
			if fromExchangeForListing, ok := from.GetExchangeForListing(ctx); ok {
				toExchangeForListing.SyncFieldsDuringRead(ctx, fromExchangeForListing)
				to.SetExchangeForListing(ctx, toExchangeForListing)
			}
		}
	}
}

func (m AddExchangeForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AddExchangeForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_for_listing": reflect.TypeOf(ExchangeListing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddExchangeForListingResponse
// only implements ToObjectValue() and Type().
func (m AddExchangeForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_for_listing": m.ExchangeForListing,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddExchangeForListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange_for_listing": ExchangeListing{}.Type(ctx),
		},
	}
}

// GetExchangeForListing returns the value of the ExchangeForListing field in AddExchangeForListingResponse as
// a ExchangeListing value.
// If the field is unknown or null, the boolean return value is false.
func (m *AddExchangeForListingResponse) GetExchangeForListing(ctx context.Context) (ExchangeListing, bool) {
	var e ExchangeListing
	if m.ExchangeForListing.IsNull() || m.ExchangeForListing.IsUnknown() {
		return e, false
	}
	var v ExchangeListing
	d := m.ExchangeForListing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeForListing sets the value of the ExchangeForListing field in AddExchangeForListingResponse.
func (m *AddExchangeForListingResponse) SetExchangeForListing(ctx context.Context, v ExchangeListing) {
	vs := v.ToObjectValue(ctx)
	m.ExchangeForListing = vs
}

type BatchGetListingsRequest struct {
	Ids types.List `tfsdk:"-"`
}

func (to *BatchGetListingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetListingsRequest) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (to *BatchGetListingsRequest) SyncFieldsDuringRead(ctx context.Context, from BatchGetListingsRequest) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (m BatchGetListingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BatchGetListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsRequest
// only implements ToObjectValue() and Type().
func (m BatchGetListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": m.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BatchGetListingsRequest) Type(ctx context.Context) attr.Type {
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
func (m *BatchGetListingsRequest) GetIds(ctx context.Context) ([]types.String, bool) {
	if m.Ids.IsNull() || m.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in BatchGetListingsRequest.
func (m *BatchGetListingsRequest) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Ids = types.ListValueMust(t, vs)
}

type BatchGetListingsResponse struct {
	Listings types.List `tfsdk:"listings"`
}

func (to *BatchGetListingsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *BatchGetListingsResponse) SyncFieldsDuringRead(ctx context.Context, from BatchGetListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (m BatchGetListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BatchGetListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetListingsResponse
// only implements ToObjectValue() and Type().
func (m BatchGetListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings": m.Listings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BatchGetListingsResponse) Type(ctx context.Context) attr.Type {
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
func (m *BatchGetListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if m.Listings.IsNull() || m.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := m.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in BatchGetListingsResponse.
func (m *BatchGetListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Listings = types.ListValueMust(t, vs)
}

type BatchGetProvidersRequest struct {
	Ids types.List `tfsdk:"-"`
}

func (to *BatchGetProvidersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetProvidersRequest) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (to *BatchGetProvidersRequest) SyncFieldsDuringRead(ctx context.Context, from BatchGetProvidersRequest) {
	if !from.Ids.IsNull() && !from.Ids.IsUnknown() && to.Ids.IsNull() && len(from.Ids.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Ids, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Ids = from.Ids
	}
}

func (m BatchGetProvidersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BatchGetProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersRequest
// only implements ToObjectValue() and Type().
func (m BatchGetProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": m.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BatchGetProvidersRequest) Type(ctx context.Context) attr.Type {
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
func (m *BatchGetProvidersRequest) GetIds(ctx context.Context) ([]types.String, bool) {
	if m.Ids.IsNull() || m.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in BatchGetProvidersRequest.
func (m *BatchGetProvidersRequest) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Ids = types.ListValueMust(t, vs)
}

type BatchGetProvidersResponse struct {
	Providers types.List `tfsdk:"providers"`
}

func (to *BatchGetProvidersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BatchGetProvidersResponse) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (to *BatchGetProvidersResponse) SyncFieldsDuringRead(ctx context.Context, from BatchGetProvidersResponse) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (m BatchGetProvidersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BatchGetProvidersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BatchGetProvidersResponse
// only implements ToObjectValue() and Type().
func (m BatchGetProvidersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"providers": m.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BatchGetProvidersResponse) Type(ctx context.Context) attr.Type {
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
func (m *BatchGetProvidersResponse) GetProviders(ctx context.Context) ([]ProviderInfo, bool) {
	if m.Providers.IsNull() || m.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo
	d := m.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in BatchGetProvidersResponse.
func (m *BatchGetProvidersResponse) SetProviders(ctx context.Context, v []ProviderInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Providers = types.ListValueMust(t, vs)
}

type ConsumerTerms struct {
	Version types.String `tfsdk:"version"`
}

func (to *ConsumerTerms) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConsumerTerms) {
}

func (to *ConsumerTerms) SyncFieldsDuringRead(ctx context.Context, from ConsumerTerms) {
}

func (m ConsumerTerms) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ConsumerTerms) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConsumerTerms
// only implements ToObjectValue() and Type().
func (m ConsumerTerms) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConsumerTerms) Type(ctx context.Context) attr.Type {
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

func (to *ContactInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ContactInfo) {
}

func (to *ContactInfo) SyncFieldsDuringRead(ctx context.Context, from ContactInfo) {
}

func (m ContactInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ContactInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContactInfo
// only implements ToObjectValue() and Type().
func (m ContactInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"company":    m.Company,
			"email":      m.Email,
			"first_name": m.FirstName,
			"last_name":  m.LastName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ContactInfo) Type(ctx context.Context) attr.Type {
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

func (to *CreateExchangeFilterRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeFilterRequest) {
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

func (to *CreateExchangeFilterRequest) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeFilterRequest) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m CreateExchangeFilterRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (m CreateExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": m.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExchangeFilterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": ExchangeFilter{}.Type(ctx),
		},
	}
}

// GetFilter returns the value of the Filter field in CreateExchangeFilterRequest as
// a ExchangeFilter value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExchangeFilterRequest) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v ExchangeFilter
	d := m.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in CreateExchangeFilterRequest.
func (m *CreateExchangeFilterRequest) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	m.Filter = vs
}

type CreateExchangeFilterResponse struct {
	FilterId types.String `tfsdk:"filter_id"`
}

func (to *CreateExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeFilterResponse) {
}

func (to *CreateExchangeFilterResponse) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeFilterResponse) {
}

func (m CreateExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (m CreateExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_id": m.FilterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_id": types.StringType,
		},
	}
}

type CreateExchangeRequest struct {
	Exchange types.Object `tfsdk:"exchange"`
}

func (to *CreateExchangeRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeRequest) {
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

func (to *CreateExchangeRequest) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeRequest) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (m CreateExchangeRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExchangeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeRequest
// only implements ToObjectValue() and Type().
func (m CreateExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": m.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in CreateExchangeRequest as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExchangeRequest) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if m.Exchange.IsNull() || m.Exchange.IsUnknown() {
		return e, false
	}
	var v Exchange
	d := m.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchange sets the value of the Exchange field in CreateExchangeRequest.
func (m *CreateExchangeRequest) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	m.Exchange = vs
}

type CreateExchangeResponse struct {
	ExchangeId types.String `tfsdk:"exchange_id"`
}

func (to *CreateExchangeResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExchangeResponse) {
}

func (to *CreateExchangeResponse) SyncFieldsDuringRead(ctx context.Context, from CreateExchangeResponse) {
}

func (m CreateExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExchangeResponse
// only implements ToObjectValue() and Type().
func (m CreateExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": m.ExchangeId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExchangeResponse) Type(ctx context.Context) attr.Type {
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

func (to *CreateFileRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFileRequest) {
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

func (to *CreateFileRequest) SyncFieldsDuringRead(ctx context.Context, from CreateFileRequest) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (m CreateFileRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["file_parent"] = attrs["file_parent"].SetRequired()
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
func (m CreateFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileRequest
// only implements ToObjectValue() and Type().
func (m CreateFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          m.DisplayName,
			"file_parent":           m.FileParent,
			"marketplace_file_type": m.MarketplaceFileType,
			"mime_type":             m.MimeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFileRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateFileRequest) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if m.FileParent.IsNull() || m.FileParent.IsUnknown() {
		return e, false
	}
	var v FileParent
	d := m.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileParent sets the value of the FileParent field in CreateFileRequest.
func (m *CreateFileRequest) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	m.FileParent = vs
}

type CreateFileResponse struct {
	FileInfo types.Object `tfsdk:"file_info"`
	// Pre-signed POST URL to blob storage
	UploadUrl types.String `tfsdk:"upload_url"`
}

func (to *CreateFileResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFileResponse) {
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

func (to *CreateFileResponse) SyncFieldsDuringRead(ctx context.Context, from CreateFileResponse) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				toFileInfo.SyncFieldsDuringRead(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (m CreateFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFileResponse
// only implements ToObjectValue() and Type().
func (m CreateFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info":  m.FileInfo,
			"upload_url": m.UploadUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFileResponse) Type(ctx context.Context) attr.Type {
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
func (m *CreateFileResponse) GetFileInfo(ctx context.Context) (FileInfo, bool) {
	var e FileInfo
	if m.FileInfo.IsNull() || m.FileInfo.IsUnknown() {
		return e, false
	}
	var v FileInfo
	d := m.FileInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileInfo sets the value of the FileInfo field in CreateFileResponse.
func (m *CreateFileResponse) SetFileInfo(ctx context.Context, v FileInfo) {
	vs := v.ToObjectValue(ctx)
	m.FileInfo = vs
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

func (to *CreateInstallationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateInstallationRequest) {
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

func (to *CreateInstallationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateInstallationRequest) {
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

func (m CreateInstallationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["recipient_type"] = attrs["recipient_type"].SetOptional()
	attrs["repo_detail"] = attrs["repo_detail"].SetOptional()
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
func (m CreateInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms{}),
		"repo_detail":             reflect.TypeOf(RepoInstallation{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstallationRequest
// only implements ToObjectValue() and Type().
func (m CreateInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"accepted_consumer_terms": m.AcceptedConsumerTerms,
			"catalog_name":            m.CatalogName,
			"listing_id":              m.ListingId,
			"recipient_type":          m.RecipientType,
			"repo_detail":             m.RepoDetail,
			"share_name":              m.ShareName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateInstallationRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateInstallationRequest) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms, bool) {
	var e ConsumerTerms
	if m.AcceptedConsumerTerms.IsNull() || m.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v ConsumerTerms
	d := m.AcceptedConsumerTerms.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreateInstallationRequest.
func (m *CreateInstallationRequest) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms) {
	vs := v.ToObjectValue(ctx)
	m.AcceptedConsumerTerms = vs
}

// GetRepoDetail returns the value of the RepoDetail field in CreateInstallationRequest as
// a RepoInstallation value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstallationRequest) GetRepoDetail(ctx context.Context) (RepoInstallation, bool) {
	var e RepoInstallation
	if m.RepoDetail.IsNull() || m.RepoDetail.IsUnknown() {
		return e, false
	}
	var v RepoInstallation
	d := m.RepoDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepoDetail sets the value of the RepoDetail field in CreateInstallationRequest.
func (m *CreateInstallationRequest) SetRepoDetail(ctx context.Context, v RepoInstallation) {
	vs := v.ToObjectValue(ctx)
	m.RepoDetail = vs
}

type CreateListingRequest struct {
	Listing types.Object `tfsdk:"listing"`
}

func (to *CreateListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateListingRequest) {
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

func (to *CreateListingRequest) SyncFieldsDuringRead(ctx context.Context, from CreateListingRequest) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (m CreateListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateListingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingRequest
// only implements ToObjectValue() and Type().
func (m CreateListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": m.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in CreateListingRequest as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateListingRequest) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if m.Listing.IsNull() || m.Listing.IsUnknown() {
		return e, false
	}
	var v Listing
	d := m.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListing sets the value of the Listing field in CreateListingRequest.
func (m *CreateListingRequest) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	m.Listing = vs
}

type CreateListingResponse struct {
	ListingId types.String `tfsdk:"listing_id"`
}

func (to *CreateListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateListingResponse) {
}

func (to *CreateListingResponse) SyncFieldsDuringRead(ctx context.Context, from CreateListingResponse) {
}

func (m CreateListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateListingResponse
// only implements ToObjectValue() and Type().
func (m CreateListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateListingResponse) Type(ctx context.Context) attr.Type {
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

func (to *CreatePersonalizationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePersonalizationRequest) {
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

func (to *CreatePersonalizationRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePersonalizationRequest) {
	if !from.AcceptedConsumerTerms.IsNull() && !from.AcceptedConsumerTerms.IsUnknown() {
		if toAcceptedConsumerTerms, ok := to.GetAcceptedConsumerTerms(ctx); ok {
			if fromAcceptedConsumerTerms, ok := from.GetAcceptedConsumerTerms(ctx); ok {
				toAcceptedConsumerTerms.SyncFieldsDuringRead(ctx, fromAcceptedConsumerTerms)
				to.SetAcceptedConsumerTerms(ctx, toAcceptedConsumerTerms)
			}
		}
	}
}

func (m CreatePersonalizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["accepted_consumer_terms"] = attrs["accepted_consumer_terms"].SetRequired()
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
func (m CreatePersonalizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"accepted_consumer_terms": reflect.TypeOf(ConsumerTerms{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequest
// only implements ToObjectValue() and Type().
func (m CreatePersonalizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"accepted_consumer_terms": m.AcceptedConsumerTerms,
			"comment":                 m.Comment,
			"company":                 m.Company,
			"first_name":              m.FirstName,
			"intended_use":            m.IntendedUse,
			"is_from_lighthouse":      m.IsFromLighthouse,
			"last_name":               m.LastName,
			"listing_id":              m.ListingId,
			"recipient_type":          m.RecipientType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePersonalizationRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreatePersonalizationRequest) GetAcceptedConsumerTerms(ctx context.Context) (ConsumerTerms, bool) {
	var e ConsumerTerms
	if m.AcceptedConsumerTerms.IsNull() || m.AcceptedConsumerTerms.IsUnknown() {
		return e, false
	}
	var v ConsumerTerms
	d := m.AcceptedConsumerTerms.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAcceptedConsumerTerms sets the value of the AcceptedConsumerTerms field in CreatePersonalizationRequest.
func (m *CreatePersonalizationRequest) SetAcceptedConsumerTerms(ctx context.Context, v ConsumerTerms) {
	vs := v.ToObjectValue(ctx)
	m.AcceptedConsumerTerms = vs
}

type CreatePersonalizationRequestResponse struct {
	Id types.String `tfsdk:"id"`
}

func (to *CreatePersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePersonalizationRequestResponse) {
}

func (to *CreatePersonalizationRequestResponse) SyncFieldsDuringRead(ctx context.Context, from CreatePersonalizationRequestResponse) {
}

func (m CreatePersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (m CreatePersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateProviderAnalyticsDashboardRequest struct {
}

func (to *CreateProviderAnalyticsDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderAnalyticsDashboardRequest) {
}

func (to *CreateProviderAnalyticsDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from CreateProviderAnalyticsDashboardRequest) {
}

func (m CreateProviderAnalyticsDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (m CreateProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateProviderRequest struct {
	Provider types.Object `tfsdk:"provider"`
}

func (to *CreateProviderRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderRequest) {
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

func (to *CreateProviderRequest) SyncFieldsDuringRead(ctx context.Context, from CreateProviderRequest) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (m CreateProviderRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderRequest
// only implements ToObjectValue() and Type().
func (m CreateProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": m.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in CreateProviderRequest as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProviderRequest) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if m.Provider.IsNull() || m.Provider.IsUnknown() {
		return e, false
	}
	var v ProviderInfo
	d := m.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProvider sets the value of the Provider field in CreateProviderRequest.
func (m *CreateProviderRequest) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	m.Provider = vs
}

type CreateProviderResponse struct {
	Id types.String `tfsdk:"id"`
}

func (to *CreateProviderResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProviderResponse) {
}

func (to *CreateProviderResponse) SyncFieldsDuringRead(ctx context.Context, from CreateProviderResponse) {
}

func (m CreateProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProviderResponse
// only implements ToObjectValue() and Type().
func (m CreateProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProviderResponse) Type(ctx context.Context) attr.Type {
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

func (to *DataRefreshInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataRefreshInfo) {
}

func (to *DataRefreshInfo) SyncFieldsDuringRead(ctx context.Context, from DataRefreshInfo) {
}

func (m DataRefreshInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataRefreshInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataRefreshInfo
// only implements ToObjectValue() and Type().
func (m DataRefreshInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"interval": m.Interval,
			"unit":     m.Unit,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataRefreshInfo) Type(ctx context.Context) attr.Type {
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

func (to *DeleteExchangeFilterRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeFilterRequest) {
}

func (to *DeleteExchangeFilterRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeFilterRequest) {
}

func (m DeleteExchangeFilterRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (m DeleteExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExchangeFilterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeFilterResponse struct {
}

func (to *DeleteExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeFilterResponse) {
}

func (to *DeleteExchangeFilterResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeFilterResponse) {
}

func (m DeleteExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeFilterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (m DeleteExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExchangeRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteExchangeRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeRequest) {
}

func (to *DeleteExchangeRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeRequest) {
}

func (m DeleteExchangeRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeRequest
// only implements ToObjectValue() and Type().
func (m DeleteExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteExchangeResponse struct {
}

func (to *DeleteExchangeResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExchangeResponse) {
}

func (to *DeleteExchangeResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteExchangeResponse) {
}

func (m DeleteExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExchangeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExchangeResponse
// only implements ToObjectValue() and Type().
func (m DeleteExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

func (to *DeleteFileRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFileRequest) {
}

func (to *DeleteFileRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteFileRequest) {
}

func (m DeleteFileRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileRequest
// only implements ToObjectValue() and Type().
func (m DeleteFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": m.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type DeleteFileResponse struct {
}

func (to *DeleteFileResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFileResponse) {
}

func (to *DeleteFileResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteFileResponse) {
}

func (m DeleteFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileResponse
// only implements ToObjectValue() and Type().
func (m DeleteFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFileResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteInstallationRequest struct {
	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`
}

func (to *DeleteInstallationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstallationRequest) {
}

func (to *DeleteInstallationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteInstallationRequest) {
}

func (m DeleteInstallationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationRequest
// only implements ToObjectValue() and Type().
func (m DeleteInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation_id": m.InstallationId,
			"listing_id":      m.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteInstallationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation_id": types.StringType,
			"listing_id":      types.StringType,
		},
	}
}

type DeleteInstallationResponse struct {
}

func (to *DeleteInstallationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstallationResponse) {
}

func (to *DeleteInstallationResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteInstallationResponse) {
}

func (m DeleteInstallationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstallationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteInstallationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstallationResponse
// only implements ToObjectValue() and Type().
func (m DeleteInstallationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteInstallationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteListingRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteListingRequest) {
}

func (to *DeleteListingRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteListingRequest) {
}

func (m DeleteListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingRequest
// only implements ToObjectValue() and Type().
func (m DeleteListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteListingResponse struct {
}

func (to *DeleteListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteListingResponse) {
}

func (to *DeleteListingResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteListingResponse) {
}

func (m DeleteListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteListingResponse
// only implements ToObjectValue() and Type().
func (m DeleteListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteProviderRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProviderRequest) {
}

func (to *DeleteProviderRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteProviderRequest) {
}

func (m DeleteProviderRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest
// only implements ToObjectValue() and Type().
func (m DeleteProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteProviderResponse struct {
}

func (to *DeleteProviderResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProviderResponse) {
}

func (to *DeleteProviderResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteProviderResponse) {
}

func (m DeleteProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderResponse
// only implements ToObjectValue() and Type().
func (m DeleteProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProviderResponse) Type(ctx context.Context) attr.Type {
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

func (to *Exchange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Exchange) {
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

func (to *Exchange) SyncFieldsDuringRead(ctx context.Context, from Exchange) {
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

func (m Exchange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Exchange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters":         reflect.TypeOf(ExchangeFilter{}),
		"linked_listings": reflect.TypeOf(ExchangeListing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Exchange
// only implements ToObjectValue() and Type().
func (m Exchange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":         m.Comment,
			"created_at":      m.CreatedAt,
			"created_by":      m.CreatedBy,
			"filters":         m.Filters,
			"id":              m.Id,
			"linked_listings": m.LinkedListings,
			"name":            m.Name,
			"updated_at":      m.UpdatedAt,
			"updated_by":      m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Exchange) Type(ctx context.Context) attr.Type {
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
func (m *Exchange) GetFilters(ctx context.Context) ([]ExchangeFilter, bool) {
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in Exchange.
func (m *Exchange) SetFilters(ctx context.Context, v []ExchangeFilter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Filters = types.ListValueMust(t, vs)
}

// GetLinkedListings returns the value of the LinkedListings field in Exchange as
// a slice of ExchangeListing values.
// If the field is unknown or null, the boolean return value is false.
func (m *Exchange) GetLinkedListings(ctx context.Context) ([]ExchangeListing, bool) {
	if m.LinkedListings.IsNull() || m.LinkedListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := m.LinkedListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLinkedListings sets the value of the LinkedListings field in Exchange.
func (m *Exchange) SetLinkedListings(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["linked_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LinkedListings = types.ListValueMust(t, vs)
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

func (to *ExchangeFilter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeFilter) {
}

func (to *ExchangeFilter) SyncFieldsDuringRead(ctx context.Context, from ExchangeFilter) {
}

func (m ExchangeFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExchangeFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeFilter
// only implements ToObjectValue() and Type().
func (m ExchangeFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":   m.CreatedAt,
			"created_by":   m.CreatedBy,
			"exchange_id":  m.ExchangeId,
			"filter_type":  m.FilterType,
			"filter_value": m.FilterValue,
			"id":           m.Id,
			"name":         m.Name,
			"updated_at":   m.UpdatedAt,
			"updated_by":   m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeFilter) Type(ctx context.Context) attr.Type {
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

func (to *ExchangeListing) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeListing) {
}

func (to *ExchangeListing) SyncFieldsDuringRead(ctx context.Context, from ExchangeListing) {
}

func (m ExchangeListing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExchangeListing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeListing
// only implements ToObjectValue() and Type().
func (m ExchangeListing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":    m.CreatedAt,
			"created_by":    m.CreatedBy,
			"exchange_id":   m.ExchangeId,
			"exchange_name": m.ExchangeName,
			"id":            m.Id,
			"listing_id":    m.ListingId,
			"listing_name":  m.ListingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeListing) Type(ctx context.Context) attr.Type {
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

func (to *FileInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileInfo) {
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

func (to *FileInfo) SyncFieldsDuringRead(ctx context.Context, from FileInfo) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (m FileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo
// only implements ToObjectValue() and Type().
func (m FileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":            m.CreatedAt,
			"display_name":          m.DisplayName,
			"download_link":         m.DownloadLink,
			"file_parent":           m.FileParent,
			"id":                    m.Id,
			"marketplace_file_type": m.MarketplaceFileType,
			"mime_type":             m.MimeType,
			"status":                m.Status,
			"status_message":        m.StatusMessage,
			"updated_at":            m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileInfo) Type(ctx context.Context) attr.Type {
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
func (m *FileInfo) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if m.FileParent.IsNull() || m.FileParent.IsUnknown() {
		return e, false
	}
	var v FileParent
	d := m.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileParent sets the value of the FileParent field in FileInfo.
func (m *FileInfo) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	m.FileParent = vs
}

type FileParent struct {
	FileParentType types.String `tfsdk:"file_parent_type"`
	// TODO make the following fields required
	ParentId types.String `tfsdk:"parent_id"`
}

func (to *FileParent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileParent) {
}

func (to *FileParent) SyncFieldsDuringRead(ctx context.Context, from FileParent) {
}

func (m FileParent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FileParent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileParent
// only implements ToObjectValue() and Type().
func (m FileParent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent_type": m.FileParentType,
			"parent_id":        m.ParentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileParent) Type(ctx context.Context) attr.Type {
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

func (to *GetExchangeRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExchangeRequest) {
}

func (to *GetExchangeRequest) SyncFieldsDuringRead(ctx context.Context, from GetExchangeRequest) {
}

func (m GetExchangeRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeRequest
// only implements ToObjectValue() and Type().
func (m GetExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExchangeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetExchangeResponse struct {
	Exchange types.Object `tfsdk:"exchange"`
}

func (to *GetExchangeResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExchangeResponse) {
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

func (to *GetExchangeResponse) SyncFieldsDuringRead(ctx context.Context, from GetExchangeResponse) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (m GetExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExchangeResponse
// only implements ToObjectValue() and Type().
func (m GetExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": m.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in GetExchangeResponse as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExchangeResponse) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if m.Exchange.IsNull() || m.Exchange.IsUnknown() {
		return e, false
	}
	var v Exchange
	d := m.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchange sets the value of the Exchange field in GetExchangeResponse.
func (m *GetExchangeResponse) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	m.Exchange = vs
}

type GetFileRequest struct {
	FileId types.String `tfsdk:"-"`
}

func (to *GetFileRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFileRequest) {
}

func (to *GetFileRequest) SyncFieldsDuringRead(ctx context.Context, from GetFileRequest) {
}

func (m GetFileRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileRequest
// only implements ToObjectValue() and Type().
func (m GetFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_id": m.FileId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_id": types.StringType,
		},
	}
}

type GetFileResponse struct {
	FileInfo types.Object `tfsdk:"file_info"`
}

func (to *GetFileResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFileResponse) {
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

func (to *GetFileResponse) SyncFieldsDuringRead(ctx context.Context, from GetFileResponse) {
	if !from.FileInfo.IsNull() && !from.FileInfo.IsUnknown() {
		if toFileInfo, ok := to.GetFileInfo(ctx); ok {
			if fromFileInfo, ok := from.GetFileInfo(ctx); ok {
				toFileInfo.SyncFieldsDuringRead(ctx, fromFileInfo)
				to.SetFileInfo(ctx, toFileInfo)
			}
		}
	}
}

func (m GetFileResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetFileResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_info": reflect.TypeOf(FileInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFileResponse
// only implements ToObjectValue() and Type().
func (m GetFileResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_info": m.FileInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFileResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_info": FileInfo{}.Type(ctx),
		},
	}
}

// GetFileInfo returns the value of the FileInfo field in GetFileResponse as
// a FileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetFileResponse) GetFileInfo(ctx context.Context) (FileInfo, bool) {
	var e FileInfo
	if m.FileInfo.IsNull() || m.FileInfo.IsUnknown() {
		return e, false
	}
	var v FileInfo
	d := m.FileInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileInfo sets the value of the FileInfo field in GetFileResponse.
func (m *GetFileResponse) SetFileInfo(ctx context.Context, v FileInfo) {
	vs := v.ToObjectValue(ctx)
	m.FileInfo = vs
}

type GetLatestVersionProviderAnalyticsDashboardRequest struct {
}

func (to *GetLatestVersionProviderAnalyticsDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardRequest) {
}

func (to *GetLatestVersionProviderAnalyticsDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardRequest) {
}

func (m GetLatestVersionProviderAnalyticsDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLatestVersionProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (m GetLatestVersionProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

func (to *GetLatestVersionProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (to *GetLatestVersionProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionProviderAnalyticsDashboardResponse) {
}

func (m GetLatestVersionProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLatestVersionProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (m GetLatestVersionProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
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

func (to *GetListingContentMetadataRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingContentMetadataRequest) {
}

func (to *GetListingContentMetadataRequest) SyncFieldsDuringRead(ctx context.Context, from GetListingContentMetadataRequest) {
}

func (m GetListingContentMetadataRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingContentMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataRequest
// only implements ToObjectValue() and Type().
func (m GetListingContentMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingContentMetadataRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetListingContentMetadataResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingContentMetadataResponse) {
	if !from.SharedDataObjects.IsNull() && !from.SharedDataObjects.IsUnknown() && to.SharedDataObjects.IsNull() && len(from.SharedDataObjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SharedDataObjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SharedDataObjects = from.SharedDataObjects
	}
}

func (to *GetListingContentMetadataResponse) SyncFieldsDuringRead(ctx context.Context, from GetListingContentMetadataResponse) {
	if !from.SharedDataObjects.IsNull() && !from.SharedDataObjects.IsUnknown() && to.SharedDataObjects.IsNull() && len(from.SharedDataObjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SharedDataObjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SharedDataObjects = from.SharedDataObjects
	}
}

func (m GetListingContentMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingContentMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shared_data_objects": reflect.TypeOf(SharedDataObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingContentMetadataResponse
// only implements ToObjectValue() and Type().
func (m GetListingContentMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     m.NextPageToken,
			"shared_data_objects": m.SharedDataObjects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingContentMetadataResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetListingContentMetadataResponse) GetSharedDataObjects(ctx context.Context) ([]SharedDataObject, bool) {
	if m.SharedDataObjects.IsNull() || m.SharedDataObjects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject
	d := m.SharedDataObjects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSharedDataObjects sets the value of the SharedDataObjects field in GetListingContentMetadataResponse.
func (m *GetListingContentMetadataResponse) SetSharedDataObjects(ctx context.Context, v []SharedDataObject) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["shared_data_objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SharedDataObjects = types.ListValueMust(t, vs)
}

type GetListingRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingRequest) {
}

func (to *GetListingRequest) SyncFieldsDuringRead(ctx context.Context, from GetListingRequest) {
}

func (m GetListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingRequest
// only implements ToObjectValue() and Type().
func (m GetListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetListingResponse struct {
	Listing types.Object `tfsdk:"listing"`
}

func (to *GetListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingResponse) {
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

func (to *GetListingResponse) SyncFieldsDuringRead(ctx context.Context, from GetListingResponse) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (m GetListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingResponse
// only implements ToObjectValue() and Type().
func (m GetListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": m.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in GetListingResponse as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetListingResponse) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if m.Listing.IsNull() || m.Listing.IsUnknown() {
		return e, false
	}
	var v Listing
	d := m.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListing sets the value of the Listing field in GetListingResponse.
func (m *GetListingResponse) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	m.Listing = vs
}

type GetListingsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *GetListingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingsRequest) {
}

func (to *GetListingsRequest) SyncFieldsDuringRead(ctx context.Context, from GetListingsRequest) {
}

func (m GetListingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsRequest
// only implements ToObjectValue() and Type().
func (m GetListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingsRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetListingsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *GetListingsResponse) SyncFieldsDuringRead(ctx context.Context, from GetListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (m GetListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetListingsResponse
// only implements ToObjectValue() and Type().
func (m GetListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        m.Listings,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetListingsResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if m.Listings.IsNull() || m.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := m.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in GetListingsResponse.
func (m *GetListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Listings = types.ListValueMust(t, vs)
}

type GetPersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`
}

func (to *GetPersonalizationRequestRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalizationRequestRequest) {
}

func (to *GetPersonalizationRequestRequest) SyncFieldsDuringRead(ctx context.Context, from GetPersonalizationRequestRequest) {
}

func (m GetPersonalizationRequestRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPersonalizationRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestRequest
// only implements ToObjectValue() and Type().
func (m GetPersonalizationRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPersonalizationRequestRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing_id": types.StringType,
		},
	}
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests types.List `tfsdk:"personalization_requests"`
}

func (to *GetPersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalizationRequestResponse) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (to *GetPersonalizationRequestResponse) SyncFieldsDuringRead(ctx context.Context, from GetPersonalizationRequestResponse) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (m GetPersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (m GetPersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"personalization_requests": m.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetPersonalizationRequestResponse) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest, bool) {
	if m.PersonalizationRequests.IsNull() || m.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest
	d := m.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in GetPersonalizationRequestResponse.
func (m *GetPersonalizationRequestResponse) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PersonalizationRequests = types.ListValueMust(t, vs)
}

type GetProviderRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetProviderRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProviderRequest) {
}

func (to *GetProviderRequest) SyncFieldsDuringRead(ctx context.Context, from GetProviderRequest) {
}

func (m GetProviderRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest
// only implements ToObjectValue() and Type().
func (m GetProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetProviderResponse struct {
	Provider types.Object `tfsdk:"provider"`
}

func (to *GetProviderResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProviderResponse) {
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

func (to *GetProviderResponse) SyncFieldsDuringRead(ctx context.Context, from GetProviderResponse) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (m GetProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderResponse
// only implements ToObjectValue() and Type().
func (m GetProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": m.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in GetProviderResponse as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetProviderResponse) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if m.Provider.IsNull() || m.Provider.IsUnknown() {
		return e, false
	}
	var v ProviderInfo
	d := m.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProvider sets the value of the Provider field in GetProviderResponse.
func (m *GetProviderResponse) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	m.Provider = vs
}

type Installation struct {
	Installation types.Object `tfsdk:"installation"`
}

func (to *Installation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Installation) {
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

func (to *Installation) SyncFieldsDuringRead(ctx context.Context, from Installation) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (m Installation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Installation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Installation
// only implements ToObjectValue() and Type().
func (m Installation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": m.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Installation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": InstallationDetail{}.Type(ctx),
		},
	}
}

// GetInstallation returns the value of the Installation field in Installation as
// a InstallationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *Installation) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if m.Installation.IsNull() || m.Installation.IsUnknown() {
		return e, false
	}
	var v InstallationDetail
	d := m.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallation sets the value of the Installation field in Installation.
func (m *Installation) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	m.Installation = vs
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

func (to *InstallationDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstallationDetail) {
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

func (to *InstallationDetail) SyncFieldsDuringRead(ctx context.Context, from InstallationDetail) {
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

func (m InstallationDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m InstallationDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_detail": reflect.TypeOf(TokenDetail{}),
		"tokens":       reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallationDetail
// only implements ToObjectValue() and Type().
func (m InstallationDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   m.CatalogName,
			"error_message":  m.ErrorMessage,
			"id":             m.Id,
			"installed_on":   m.InstalledOn,
			"listing_id":     m.ListingId,
			"listing_name":   m.ListingName,
			"recipient_type": m.RecipientType,
			"repo_name":      m.RepoName,
			"repo_path":      m.RepoPath,
			"share_name":     m.ShareName,
			"status":         m.Status,
			"token_detail":   m.TokenDetail,
			"tokens":         m.Tokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstallationDetail) Type(ctx context.Context) attr.Type {
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
func (m *InstallationDetail) GetTokenDetail(ctx context.Context) (TokenDetail, bool) {
	var e TokenDetail
	if m.TokenDetail.IsNull() || m.TokenDetail.IsUnknown() {
		return e, false
	}
	var v TokenDetail
	d := m.TokenDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenDetail sets the value of the TokenDetail field in InstallationDetail.
func (m *InstallationDetail) SetTokenDetail(ctx context.Context, v TokenDetail) {
	vs := v.ToObjectValue(ctx)
	m.TokenDetail = vs
}

// GetTokens returns the value of the Tokens field in InstallationDetail as
// a slice of TokenInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstallationDetail) GetTokens(ctx context.Context) ([]TokenInfo, bool) {
	if m.Tokens.IsNull() || m.Tokens.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo
	d := m.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in InstallationDetail.
func (m *InstallationDetail) SetTokens(ctx context.Context, v []TokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tokens = types.ListValueMust(t, vs)
}

type ListAllInstallationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAllInstallationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllInstallationsRequest) {
}

func (to *ListAllInstallationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAllInstallationsRequest) {
}

func (m ListAllInstallationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAllInstallationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsRequest
// only implements ToObjectValue() and Type().
func (m ListAllInstallationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllInstallationsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListAllInstallationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllInstallationsResponse) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (to *ListAllInstallationsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAllInstallationsResponse) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (m ListAllInstallationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAllInstallationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllInstallationsResponse
// only implements ToObjectValue() and Type().
func (m ListAllInstallationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   m.Installations,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllInstallationsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListAllInstallationsResponse) GetInstallations(ctx context.Context) ([]InstallationDetail, bool) {
	if m.Installations.IsNull() || m.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail
	d := m.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListAllInstallationsResponse.
func (m *ListAllInstallationsResponse) SetInstallations(ctx context.Context, v []InstallationDetail) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Installations = types.ListValueMust(t, vs)
}

type ListAllPersonalizationRequestsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAllPersonalizationRequestsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllPersonalizationRequestsRequest) {
}

func (to *ListAllPersonalizationRequestsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAllPersonalizationRequestsRequest) {
}

func (m ListAllPersonalizationRequestsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAllPersonalizationRequestsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsRequest
// only implements ToObjectValue() and Type().
func (m ListAllPersonalizationRequestsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllPersonalizationRequestsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListAllPersonalizationRequestsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllPersonalizationRequestsResponse) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (to *ListAllPersonalizationRequestsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAllPersonalizationRequestsResponse) {
	if !from.PersonalizationRequests.IsNull() && !from.PersonalizationRequests.IsUnknown() && to.PersonalizationRequests.IsNull() && len(from.PersonalizationRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PersonalizationRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PersonalizationRequests = from.PersonalizationRequests
	}
}

func (m ListAllPersonalizationRequestsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAllPersonalizationRequestsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personalization_requests": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllPersonalizationRequestsResponse
// only implements ToObjectValue() and Type().
func (m ListAllPersonalizationRequestsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"personalization_requests": m.PersonalizationRequests,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllPersonalizationRequestsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListAllPersonalizationRequestsResponse) GetPersonalizationRequests(ctx context.Context) ([]PersonalizationRequest, bool) {
	if m.PersonalizationRequests.IsNull() || m.PersonalizationRequests.IsUnknown() {
		return nil, false
	}
	var v []PersonalizationRequest
	d := m.PersonalizationRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalizationRequests sets the value of the PersonalizationRequests field in ListAllPersonalizationRequestsResponse.
func (m *ListAllPersonalizationRequestsResponse) SetPersonalizationRequests(ctx context.Context, v []PersonalizationRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["personalization_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PersonalizationRequests = types.ListValueMust(t, vs)
}

type ListConsumerProvidersRequest struct {
	IsFeatured types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListConsumerProvidersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListConsumerProvidersRequest) {
}

func (to *ListConsumerProvidersRequest) SyncFieldsDuringRead(ctx context.Context, from ListConsumerProvidersRequest) {
}

func (m ListConsumerProvidersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListConsumerProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConsumerProvidersRequest
// only implements ToObjectValue() and Type().
func (m ListConsumerProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_featured": m.IsFeatured,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListConsumerProvidersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_featured": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListExchangeFiltersRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangeFiltersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangeFiltersRequest) {
}

func (to *ListExchangeFiltersRequest) SyncFieldsDuringRead(ctx context.Context, from ListExchangeFiltersRequest) {
}

func (m ListExchangeFiltersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangeFiltersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersRequest
// only implements ToObjectValue() and Type().
func (m ListExchangeFiltersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": m.ExchangeId,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangeFiltersRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListExchangeFiltersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangeFiltersResponse) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
}

func (to *ListExchangeFiltersResponse) SyncFieldsDuringRead(ctx context.Context, from ListExchangeFiltersResponse) {
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
}

func (m ListExchangeFiltersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangeFiltersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filters": reflect.TypeOf(ExchangeFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangeFiltersResponse
// only implements ToObjectValue() and Type().
func (m ListExchangeFiltersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filters":         m.Filters,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangeFiltersResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListExchangeFiltersResponse) GetFilters(ctx context.Context) ([]ExchangeFilter, bool) {
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return nil, false
	}
	var v []ExchangeFilter
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in ListExchangeFiltersResponse.
func (m *ListExchangeFiltersResponse) SetFilters(ctx context.Context, v []ExchangeFilter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Filters = types.ListValueMust(t, vs)
}

type ListExchangesForListingRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangesForListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesForListingRequest) {
}

func (to *ListExchangesForListingRequest) SyncFieldsDuringRead(ctx context.Context, from ListExchangesForListingRequest) {
}

func (m ListExchangesForListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangesForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingRequest
// only implements ToObjectValue() and Type().
func (m ListExchangesForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangesForListingRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListExchangesForListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesForListingResponse) {
	if !from.ExchangeListing.IsNull() && !from.ExchangeListing.IsUnknown() && to.ExchangeListing.IsNull() && len(from.ExchangeListing.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListing, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListing = from.ExchangeListing
	}
}

func (to *ListExchangesForListingResponse) SyncFieldsDuringRead(ctx context.Context, from ListExchangesForListingResponse) {
	if !from.ExchangeListing.IsNull() && !from.ExchangeListing.IsUnknown() && to.ExchangeListing.IsNull() && len(from.ExchangeListing.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListing, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListing = from.ExchangeListing
	}
}

func (m ListExchangesForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangesForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listing": reflect.TypeOf(ExchangeListing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesForListingResponse
// only implements ToObjectValue() and Type().
func (m ListExchangesForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listing": m.ExchangeListing,
			"next_page_token":  m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangesForListingResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListExchangesForListingResponse) GetExchangeListing(ctx context.Context) ([]ExchangeListing, bool) {
	if m.ExchangeListing.IsNull() || m.ExchangeListing.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := m.ExchangeListing.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListing sets the value of the ExchangeListing field in ListExchangesForListingResponse.
func (m *ListExchangesForListingResponse) SetExchangeListing(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listing"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExchangeListing = types.ListValueMust(t, vs)
}

type ListExchangesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListExchangesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesRequest) {
}

func (to *ListExchangesRequest) SyncFieldsDuringRead(ctx context.Context, from ListExchangesRequest) {
}

func (m ListExchangesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesRequest
// only implements ToObjectValue() and Type().
func (m ListExchangesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangesRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListExchangesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExchangesResponse) {
	if !from.Exchanges.IsNull() && !from.Exchanges.IsUnknown() && to.Exchanges.IsNull() && len(from.Exchanges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exchanges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exchanges = from.Exchanges
	}
}

func (to *ListExchangesResponse) SyncFieldsDuringRead(ctx context.Context, from ListExchangesResponse) {
	if !from.Exchanges.IsNull() && !from.Exchanges.IsUnknown() && to.Exchanges.IsNull() && len(from.Exchanges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exchanges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exchanges = from.Exchanges
	}
}

func (m ListExchangesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExchangesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchanges": reflect.TypeOf(Exchange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExchangesResponse
// only implements ToObjectValue() and Type().
func (m ListExchangesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchanges":       m.Exchanges,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExchangesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListExchangesResponse) GetExchanges(ctx context.Context) ([]Exchange, bool) {
	if m.Exchanges.IsNull() || m.Exchanges.IsUnknown() {
		return nil, false
	}
	var v []Exchange
	d := m.Exchanges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchanges sets the value of the Exchanges field in ListExchangesResponse.
func (m *ListExchangesResponse) SetExchanges(ctx context.Context, v []Exchange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exchanges"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exchanges = types.ListValueMust(t, vs)
}

type ListFilesRequest struct {
	FileParent types.Object `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFilesRequest) {
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

func (to *ListFilesRequest) SyncFieldsDuringRead(ctx context.Context, from ListFilesRequest) {
	if !from.FileParent.IsNull() && !from.FileParent.IsUnknown() {
		if toFileParent, ok := to.GetFileParent(ctx); ok {
			if fromFileParent, ok := from.GetFileParent(ctx); ok {
				toFileParent.SyncFieldsDuringRead(ctx, fromFileParent)
				to.SetFileParent(ctx, toFileParent)
			}
		}
	}
}

func (m ListFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_parent"] = attrs["file_parent"].SetRequired()
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
func (m ListFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_parent": reflect.TypeOf(FileParent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesRequest
// only implements ToObjectValue() and Type().
func (m ListFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_parent": m.FileParent,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFilesRequest) Type(ctx context.Context) attr.Type {
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
func (m *ListFilesRequest) GetFileParent(ctx context.Context) (FileParent, bool) {
	var e FileParent
	if m.FileParent.IsNull() || m.FileParent.IsUnknown() {
		return e, false
	}
	var v FileParent
	d := m.FileParent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileParent sets the value of the FileParent field in ListFilesRequest.
func (m *ListFilesRequest) SetFileParent(ctx context.Context, v FileParent) {
	vs := v.ToObjectValue(ctx)
	m.FileParent = vs
}

type ListFilesResponse struct {
	FileInfos types.List `tfsdk:"file_infos"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFilesResponse) {
	if !from.FileInfos.IsNull() && !from.FileInfos.IsUnknown() && to.FileInfos.IsNull() && len(from.FileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileInfos = from.FileInfos
	}
}

func (to *ListFilesResponse) SyncFieldsDuringRead(ctx context.Context, from ListFilesResponse) {
	if !from.FileInfos.IsNull() && !from.FileInfos.IsUnknown() && to.FileInfos.IsNull() && len(from.FileInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileInfos = from.FileInfos
	}
}

func (m ListFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_infos": reflect.TypeOf(FileInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFilesResponse
// only implements ToObjectValue() and Type().
func (m ListFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_infos":      m.FileInfos,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFilesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListFilesResponse) GetFileInfos(ctx context.Context) ([]FileInfo, bool) {
	if m.FileInfos.IsNull() || m.FileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := m.FileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileInfos sets the value of the FileInfos field in ListFilesResponse.
func (m *ListFilesResponse) SetFileInfos(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FileInfos = types.ListValueMust(t, vs)
}

type ListFulfillmentsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListFulfillmentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFulfillmentsRequest) {
}

func (to *ListFulfillmentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListFulfillmentsRequest) {
}

func (m ListFulfillmentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFulfillmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsRequest
// only implements ToObjectValue() and Type().
func (m ListFulfillmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFulfillmentsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListFulfillmentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFulfillmentsResponse) {
	if !from.Fulfillments.IsNull() && !from.Fulfillments.IsUnknown() && to.Fulfillments.IsNull() && len(from.Fulfillments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fulfillments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fulfillments = from.Fulfillments
	}
}

func (to *ListFulfillmentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListFulfillmentsResponse) {
	if !from.Fulfillments.IsNull() && !from.Fulfillments.IsUnknown() && to.Fulfillments.IsNull() && len(from.Fulfillments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fulfillments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fulfillments = from.Fulfillments
	}
}

func (m ListFulfillmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFulfillmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fulfillments": reflect.TypeOf(ListingFulfillment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFulfillmentsResponse
// only implements ToObjectValue() and Type().
func (m ListFulfillmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fulfillments":    m.Fulfillments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFulfillmentsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListFulfillmentsResponse) GetFulfillments(ctx context.Context) ([]ListingFulfillment, bool) {
	if m.Fulfillments.IsNull() || m.Fulfillments.IsUnknown() {
		return nil, false
	}
	var v []ListingFulfillment
	d := m.Fulfillments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFulfillments sets the value of the Fulfillments field in ListFulfillmentsResponse.
func (m *ListFulfillmentsResponse) SetFulfillments(ctx context.Context, v []ListingFulfillment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fulfillments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Fulfillments = types.ListValueMust(t, vs)
}

type ListInstallationsRequest struct {
	ListingId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListInstallationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstallationsRequest) {
}

func (to *ListInstallationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListInstallationsRequest) {
}

func (m ListInstallationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListInstallationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsRequest
// only implements ToObjectValue() and Type().
func (m ListInstallationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstallationsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListInstallationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstallationsResponse) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (to *ListInstallationsResponse) SyncFieldsDuringRead(ctx context.Context, from ListInstallationsResponse) {
	if !from.Installations.IsNull() && !from.Installations.IsUnknown() && to.Installations.IsNull() && len(from.Installations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Installations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Installations = from.Installations
	}
}

func (m ListInstallationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListInstallationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installations": reflect.TypeOf(InstallationDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstallationsResponse
// only implements ToObjectValue() and Type().
func (m ListInstallationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installations":   m.Installations,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstallationsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListInstallationsResponse) GetInstallations(ctx context.Context) ([]InstallationDetail, bool) {
	if m.Installations.IsNull() || m.Installations.IsUnknown() {
		return nil, false
	}
	var v []InstallationDetail
	d := m.Installations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallations sets the value of the Installations field in ListInstallationsResponse.
func (m *ListInstallationsResponse) SetInstallations(ctx context.Context, v []InstallationDetail) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["installations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Installations = types.ListValueMust(t, vs)
}

type ListListingsForExchangeRequest struct {
	ExchangeId types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListListingsForExchangeRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsForExchangeRequest) {
}

func (to *ListListingsForExchangeRequest) SyncFieldsDuringRead(ctx context.Context, from ListListingsForExchangeRequest) {
}

func (m ListListingsForExchangeRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListListingsForExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeRequest
// only implements ToObjectValue() and Type().
func (m ListListingsForExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_id": m.ExchangeId,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListListingsForExchangeRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListListingsForExchangeResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsForExchangeResponse) {
	if !from.ExchangeListings.IsNull() && !from.ExchangeListings.IsUnknown() && to.ExchangeListings.IsNull() && len(from.ExchangeListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListings = from.ExchangeListings
	}
}

func (to *ListListingsForExchangeResponse) SyncFieldsDuringRead(ctx context.Context, from ListListingsForExchangeResponse) {
	if !from.ExchangeListings.IsNull() && !from.ExchangeListings.IsUnknown() && to.ExchangeListings.IsNull() && len(from.ExchangeListings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExchangeListings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExchangeListings = from.ExchangeListings
	}
}

func (m ListListingsForExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListListingsForExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange_listings": reflect.TypeOf(ExchangeListing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsForExchangeResponse
// only implements ToObjectValue() and Type().
func (m ListListingsForExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange_listings": m.ExchangeListings,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListListingsForExchangeResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListListingsForExchangeResponse) GetExchangeListings(ctx context.Context) ([]ExchangeListing, bool) {
	if m.ExchangeListings.IsNull() || m.ExchangeListings.IsUnknown() {
		return nil, false
	}
	var v []ExchangeListing
	d := m.ExchangeListings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeListings sets the value of the ExchangeListings field in ListListingsForExchangeResponse.
func (m *ListListingsForExchangeResponse) SetExchangeListings(ctx context.Context, v []ExchangeListing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExchangeListings = types.ListValueMust(t, vs)
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

func (to *ListListingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsRequest) {
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

func (to *ListListingsRequest) SyncFieldsDuringRead(ctx context.Context, from ListListingsRequest) {
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

func (m ListListingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m ListListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":              m.Assets,
			"categories":          m.Categories,
			"is_free":             m.IsFree,
			"is_private_exchange": m.IsPrivateExchange,
			"is_staff_pick":       m.IsStaffPick,
			"page_size":           m.PageSize,
			"page_token":          m.PageToken,
			"provider_ids":        m.ProviderIds,
			"tags":                m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListListingsRequest) Type(ctx context.Context) attr.Type {
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
func (m *ListListingsRequest) GetAssets(ctx context.Context) ([]types.String, bool) {
	if m.Assets.IsNull() || m.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListListingsRequest.
func (m *ListListingsRequest) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in ListListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListListingsRequest) GetCategories(ctx context.Context) ([]types.String, bool) {
	if m.Categories.IsNull() || m.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in ListListingsRequest.
func (m *ListListingsRequest) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in ListListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListListingsRequest) GetProviderIds(ctx context.Context) ([]types.String, bool) {
	if m.ProviderIds.IsNull() || m.ProviderIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ProviderIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviderIds sets the value of the ProviderIds field in ListListingsRequest.
func (m *ListListingsRequest) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ProviderIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListListingsRequest as
// a slice of ListingTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListListingsRequest) GetTags(ctx context.Context) ([]ListingTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListListingsRequest.
func (m *ListListingsRequest) SetTags(ctx context.Context, v []ListingTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ListListingsResponse struct {
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListListingsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *ListListingsResponse) SyncFieldsDuringRead(ctx context.Context, from ListListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (m ListListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListListingsResponse
// only implements ToObjectValue() and Type().
func (m ListListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        m.Listings,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListListingsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if m.Listings.IsNull() || m.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := m.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in ListListingsResponse.
func (m *ListListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Listings = types.ListValueMust(t, vs)
}

type ListProviderAnalyticsDashboardRequest struct {
}

func (to *ListProviderAnalyticsDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderAnalyticsDashboardRequest) {
}

func (to *ListProviderAnalyticsDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from ListProviderAnalyticsDashboardRequest) {
}

func (m ListProviderAnalyticsDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderAnalyticsDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (m ListProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderAnalyticsDashboardResponse) {
}

func (to *ListProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(ctx context.Context, from ListProviderAnalyticsDashboardResponse) {
}

func (m ListProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (m ListProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
			"id":           m.Id,
			"version":      m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"version":      types.Int64Type,
		},
	}
}

type ListProvidersRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListProvidersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersRequest) {
}

func (to *ListProvidersRequest) SyncFieldsDuringRead(ctx context.Context, from ListProvidersRequest) {
}

func (m ListProvidersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest
// only implements ToObjectValue() and Type().
func (m ListProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProvidersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProvidersResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Providers types.List `tfsdk:"providers"`
}

func (to *ListProvidersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersResponse) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (to *ListProvidersResponse) SyncFieldsDuringRead(ctx context.Context, from ListProvidersResponse) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (m ListProvidersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProvidersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse
// only implements ToObjectValue() and Type().
func (m ListProvidersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"providers":       m.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProvidersResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListProvidersResponse) GetProviders(ctx context.Context) ([]ProviderInfo, bool) {
	if m.Providers.IsNull() || m.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo
	d := m.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse.
func (m *ListProvidersResponse) SetProviders(ctx context.Context, v []ProviderInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Providers = types.ListValueMust(t, vs)
}

type Listing struct {
	Detail types.Object `tfsdk:"detail"`

	Id types.String `tfsdk:"id"`

	Summary types.Object `tfsdk:"summary"`
}

func (to *Listing) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Listing) {
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

func (to *Listing) SyncFieldsDuringRead(ctx context.Context, from Listing) {
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

func (m Listing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Listing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"detail":  reflect.TypeOf(ListingDetail{}),
		"summary": reflect.TypeOf(ListingSummary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Listing
// only implements ToObjectValue() and Type().
func (m Listing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"detail":  m.Detail,
			"id":      m.Id,
			"summary": m.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Listing) Type(ctx context.Context) attr.Type {
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
func (m *Listing) GetDetail(ctx context.Context) (ListingDetail, bool) {
	var e ListingDetail
	if m.Detail.IsNull() || m.Detail.IsUnknown() {
		return e, false
	}
	var v ListingDetail
	d := m.Detail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetail sets the value of the Detail field in Listing.
func (m *Listing) SetDetail(ctx context.Context, v ListingDetail) {
	vs := v.ToObjectValue(ctx)
	m.Detail = vs
}

// GetSummary returns the value of the Summary field in Listing as
// a ListingSummary value.
// If the field is unknown or null, the boolean return value is false.
func (m *Listing) GetSummary(ctx context.Context) (ListingSummary, bool) {
	var e ListingSummary
	if m.Summary.IsNull() || m.Summary.IsUnknown() {
		return e, false
	}
	var v ListingSummary
	d := m.Summary.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSummary sets the value of the Summary field in Listing.
func (m *Listing) SetSummary(ctx context.Context, v ListingSummary) {
	vs := v.ToObjectValue(ctx)
	m.Summary = vs
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

func (to *ListingDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingDetail) {
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

func (to *ListingDetail) SyncFieldsDuringRead(ctx context.Context, from ListingDetail) {
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

func (m ListingDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListingDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m ListingDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":                       m.Assets,
			"collection_date_end":          m.CollectionDateEnd,
			"collection_date_start":        m.CollectionDateStart,
			"collection_granularity":       m.CollectionGranularity,
			"cost":                         m.Cost,
			"data_source":                  m.DataSource,
			"description":                  m.Description,
			"documentation_link":           m.DocumentationLink,
			"embedded_notebook_file_infos": m.EmbeddedNotebookFileInfos,
			"file_ids":                     m.FileIds,
			"geographical_coverage":        m.GeographicalCoverage,
			"license":                      m.License,
			"pricing_model":                m.PricingModel,
			"privacy_policy_link":          m.PrivacyPolicyLink,
			"size":                         m.Size,
			"support_link":                 m.SupportLink,
			"tags":                         m.Tags,
			"terms_of_service":             m.TermsOfService,
			"update_frequency":             m.UpdateFrequency,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListingDetail) Type(ctx context.Context) attr.Type {
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
func (m *ListingDetail) GetAssets(ctx context.Context) ([]types.String, bool) {
	if m.Assets.IsNull() || m.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListingDetail.
func (m *ListingDetail) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Assets = types.ListValueMust(t, vs)
}

// GetCollectionGranularity returns the value of the CollectionGranularity field in ListingDetail as
// a DataRefreshInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingDetail) GetCollectionGranularity(ctx context.Context) (DataRefreshInfo, bool) {
	var e DataRefreshInfo
	if m.CollectionGranularity.IsNull() || m.CollectionGranularity.IsUnknown() {
		return e, false
	}
	var v DataRefreshInfo
	d := m.CollectionGranularity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollectionGranularity sets the value of the CollectionGranularity field in ListingDetail.
func (m *ListingDetail) SetCollectionGranularity(ctx context.Context, v DataRefreshInfo) {
	vs := v.ToObjectValue(ctx)
	m.CollectionGranularity = vs
}

// GetEmbeddedNotebookFileInfos returns the value of the EmbeddedNotebookFileInfos field in ListingDetail as
// a slice of FileInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingDetail) GetEmbeddedNotebookFileInfos(ctx context.Context) ([]FileInfo, bool) {
	if m.EmbeddedNotebookFileInfos.IsNull() || m.EmbeddedNotebookFileInfos.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := m.EmbeddedNotebookFileInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddedNotebookFileInfos sets the value of the EmbeddedNotebookFileInfos field in ListingDetail.
func (m *ListingDetail) SetEmbeddedNotebookFileInfos(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedded_notebook_file_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddedNotebookFileInfos = types.ListValueMust(t, vs)
}

// GetFileIds returns the value of the FileIds field in ListingDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingDetail) GetFileIds(ctx context.Context) ([]types.String, bool) {
	if m.FileIds.IsNull() || m.FileIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FileIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileIds sets the value of the FileIds field in ListingDetail.
func (m *ListingDetail) SetFileIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FileIds = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListingDetail as
// a slice of ListingTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingDetail) GetTags(ctx context.Context) ([]ListingTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ListingTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListingDetail.
func (m *ListingDetail) SetTags(ctx context.Context, v []ListingTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// GetUpdateFrequency returns the value of the UpdateFrequency field in ListingDetail as
// a DataRefreshInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingDetail) GetUpdateFrequency(ctx context.Context) (DataRefreshInfo, bool) {
	var e DataRefreshInfo
	if m.UpdateFrequency.IsNull() || m.UpdateFrequency.IsUnknown() {
		return e, false
	}
	var v DataRefreshInfo
	d := m.UpdateFrequency.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdateFrequency sets the value of the UpdateFrequency field in ListingDetail.
func (m *ListingDetail) SetUpdateFrequency(ctx context.Context, v DataRefreshInfo) {
	vs := v.ToObjectValue(ctx)
	m.UpdateFrequency = vs
}

type ListingFulfillment struct {
	FulfillmentType types.String `tfsdk:"fulfillment_type"`

	ListingId types.String `tfsdk:"listing_id"`

	RecipientType types.String `tfsdk:"recipient_type"`

	RepoInfo types.Object `tfsdk:"repo_info"`

	ShareInfo types.Object `tfsdk:"share_info"`
}

func (to *ListingFulfillment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingFulfillment) {
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

func (to *ListingFulfillment) SyncFieldsDuringRead(ctx context.Context, from ListingFulfillment) {
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

func (m ListingFulfillment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListingFulfillment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repo_info":  reflect.TypeOf(RepoInfo{}),
		"share_info": reflect.TypeOf(ShareInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingFulfillment
// only implements ToObjectValue() and Type().
func (m ListingFulfillment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fulfillment_type": m.FulfillmentType,
			"listing_id":       m.ListingId,
			"recipient_type":   m.RecipientType,
			"repo_info":        m.RepoInfo,
			"share_info":       m.ShareInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListingFulfillment) Type(ctx context.Context) attr.Type {
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
func (m *ListingFulfillment) GetRepoInfo(ctx context.Context) (RepoInfo, bool) {
	var e RepoInfo
	if m.RepoInfo.IsNull() || m.RepoInfo.IsUnknown() {
		return e, false
	}
	var v RepoInfo
	d := m.RepoInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepoInfo sets the value of the RepoInfo field in ListingFulfillment.
func (m *ListingFulfillment) SetRepoInfo(ctx context.Context, v RepoInfo) {
	vs := v.ToObjectValue(ctx)
	m.RepoInfo = vs
}

// GetShareInfo returns the value of the ShareInfo field in ListingFulfillment as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingFulfillment) GetShareInfo(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if m.ShareInfo.IsNull() || m.ShareInfo.IsUnknown() {
		return e, false
	}
	var v ShareInfo
	d := m.ShareInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShareInfo sets the value of the ShareInfo field in ListingFulfillment.
func (m *ListingFulfillment) SetShareInfo(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	m.ShareInfo = vs
}

type ListingSetting struct {
	Visibility types.String `tfsdk:"visibility"`
}

func (to *ListingSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingSetting) {
}

func (to *ListingSetting) SyncFieldsDuringRead(ctx context.Context, from ListingSetting) {
}

func (m ListingSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListingSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingSetting
// only implements ToObjectValue() and Type().
func (m ListingSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visibility": m.Visibility,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListingSetting) Type(ctx context.Context) attr.Type {
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

	ListingType types.String `tfsdk:"listing_type"`

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

func (to *ListingSummary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingSummary) {
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

func (to *ListingSummary) SyncFieldsDuringRead(ctx context.Context, from ListingSummary) {
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

func (m ListingSummary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["categories"] = attrs["categories"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["created_by_id"] = attrs["created_by_id"].SetOptional()
	attrs["exchange_ids"] = attrs["exchange_ids"].SetOptional()
	attrs["git_repo"] = attrs["git_repo"].SetOptional()
	attrs["listing_type"] = attrs["listing_type"].SetRequired()
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
func (m ListingSummary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m ListingSummary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"categories":      m.Categories,
			"created_at":      m.CreatedAt,
			"created_by":      m.CreatedBy,
			"created_by_id":   m.CreatedById,
			"exchange_ids":    m.ExchangeIds,
			"git_repo":        m.GitRepo,
			"listing_type":    m.ListingType,
			"name":            m.Name,
			"provider_id":     m.ProviderId,
			"provider_region": m.ProviderRegion,
			"published_at":    m.PublishedAt,
			"published_by":    m.PublishedBy,
			"setting":         m.Setting,
			"share":           m.Share,
			"status":          m.Status,
			"subtitle":        m.Subtitle,
			"updated_at":      m.UpdatedAt,
			"updated_by":      m.UpdatedBy,
			"updated_by_id":   m.UpdatedById,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListingSummary) Type(ctx context.Context) attr.Type {
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
			"listing_type":    types.StringType,
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
func (m *ListingSummary) GetCategories(ctx context.Context) ([]types.String, bool) {
	if m.Categories.IsNull() || m.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in ListingSummary.
func (m *ListingSummary) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Categories = types.ListValueMust(t, vs)
}

// GetExchangeIds returns the value of the ExchangeIds field in ListingSummary as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingSummary) GetExchangeIds(ctx context.Context) ([]types.String, bool) {
	if m.ExchangeIds.IsNull() || m.ExchangeIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ExchangeIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchangeIds sets the value of the ExchangeIds field in ListingSummary.
func (m *ListingSummary) SetExchangeIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exchange_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExchangeIds = types.ListValueMust(t, vs)
}

// GetGitRepo returns the value of the GitRepo field in ListingSummary as
// a RepoInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingSummary) GetGitRepo(ctx context.Context) (RepoInfo, bool) {
	var e RepoInfo
	if m.GitRepo.IsNull() || m.GitRepo.IsUnknown() {
		return e, false
	}
	var v RepoInfo
	d := m.GitRepo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGitRepo sets the value of the GitRepo field in ListingSummary.
func (m *ListingSummary) SetGitRepo(ctx context.Context, v RepoInfo) {
	vs := v.ToObjectValue(ctx)
	m.GitRepo = vs
}

// GetProviderRegion returns the value of the ProviderRegion field in ListingSummary as
// a RegionInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingSummary) GetProviderRegion(ctx context.Context) (RegionInfo, bool) {
	var e RegionInfo
	if m.ProviderRegion.IsNull() || m.ProviderRegion.IsUnknown() {
		return e, false
	}
	var v RegionInfo
	d := m.ProviderRegion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviderRegion sets the value of the ProviderRegion field in ListingSummary.
func (m *ListingSummary) SetProviderRegion(ctx context.Context, v RegionInfo) {
	vs := v.ToObjectValue(ctx)
	m.ProviderRegion = vs
}

// GetSetting returns the value of the Setting field in ListingSummary as
// a ListingSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingSummary) GetSetting(ctx context.Context) (ListingSetting, bool) {
	var e ListingSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v ListingSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in ListingSummary.
func (m *ListingSummary) SetSetting(ctx context.Context, v ListingSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// GetShare returns the value of the Share field in ListingSummary as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListingSummary) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if m.Share.IsNull() || m.Share.IsUnknown() {
		return e, false
	}
	var v ShareInfo
	d := m.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShare sets the value of the Share field in ListingSummary.
func (m *ListingSummary) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	m.Share = vs
}

type ListingTag struct {
	// Tag name (enum)
	TagName types.String `tfsdk:"tag_name"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues types.List `tfsdk:"tag_values"`
}

func (to *ListingTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListingTag) {
	if !from.TagValues.IsNull() && !from.TagValues.IsUnknown() && to.TagValues.IsNull() && len(from.TagValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagValues = from.TagValues
	}
}

func (to *ListingTag) SyncFieldsDuringRead(ctx context.Context, from ListingTag) {
	if !from.TagValues.IsNull() && !from.TagValues.IsUnknown() && to.TagValues.IsNull() && len(from.TagValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagValues = from.TagValues
	}
}

func (m ListingTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListingTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_values": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListingTag
// only implements ToObjectValue() and Type().
func (m ListingTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_name":   m.TagName,
			"tag_values": m.TagValues,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListingTag) Type(ctx context.Context) attr.Type {
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
func (m *ListingTag) GetTagValues(ctx context.Context) ([]types.String, bool) {
	if m.TagValues.IsNull() || m.TagValues.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.TagValues.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagValues sets the value of the TagValues field in ListingTag.
func (m *ListingTag) SetTagValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TagValues = types.ListValueMust(t, vs)
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

func (to *PersonalizationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalizationRequest) {
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

func (to *PersonalizationRequest) SyncFieldsDuringRead(ctx context.Context, from PersonalizationRequest) {
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

func (m PersonalizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PersonalizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"consumer_region": reflect.TypeOf(RegionInfo{}),
		"contact_info":    reflect.TypeOf(ContactInfo{}),
		"share":           reflect.TypeOf(ShareInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalizationRequest
// only implements ToObjectValue() and Type().
func (m PersonalizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            m.Comment,
			"consumer_region":    m.ConsumerRegion,
			"contact_info":       m.ContactInfo,
			"created_at":         m.CreatedAt,
			"id":                 m.Id,
			"intended_use":       m.IntendedUse,
			"is_from_lighthouse": m.IsFromLighthouse,
			"listing_id":         m.ListingId,
			"listing_name":       m.ListingName,
			"metastore_id":       m.MetastoreId,
			"provider_id":        m.ProviderId,
			"recipient_type":     m.RecipientType,
			"share":              m.Share,
			"status":             m.Status,
			"status_message":     m.StatusMessage,
			"updated_at":         m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PersonalizationRequest) Type(ctx context.Context) attr.Type {
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
func (m *PersonalizationRequest) GetConsumerRegion(ctx context.Context) (RegionInfo, bool) {
	var e RegionInfo
	if m.ConsumerRegion.IsNull() || m.ConsumerRegion.IsUnknown() {
		return e, false
	}
	var v RegionInfo
	d := m.ConsumerRegion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConsumerRegion sets the value of the ConsumerRegion field in PersonalizationRequest.
func (m *PersonalizationRequest) SetConsumerRegion(ctx context.Context, v RegionInfo) {
	vs := v.ToObjectValue(ctx)
	m.ConsumerRegion = vs
}

// GetContactInfo returns the value of the ContactInfo field in PersonalizationRequest as
// a ContactInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *PersonalizationRequest) GetContactInfo(ctx context.Context) (ContactInfo, bool) {
	var e ContactInfo
	if m.ContactInfo.IsNull() || m.ContactInfo.IsUnknown() {
		return e, false
	}
	var v ContactInfo
	d := m.ContactInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetContactInfo sets the value of the ContactInfo field in PersonalizationRequest.
func (m *PersonalizationRequest) SetContactInfo(ctx context.Context, v ContactInfo) {
	vs := v.ToObjectValue(ctx)
	m.ContactInfo = vs
}

// GetShare returns the value of the Share field in PersonalizationRequest as
// a ShareInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *PersonalizationRequest) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if m.Share.IsNull() || m.Share.IsUnknown() {
		return e, false
	}
	var v ShareInfo
	d := m.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShare sets the value of the Share field in PersonalizationRequest.
func (m *PersonalizationRequest) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	m.Share = vs
}

type ProviderAnalyticsDashboard struct {
	Id types.String `tfsdk:"id"`
}

func (to *ProviderAnalyticsDashboard) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderAnalyticsDashboard) {
}

func (to *ProviderAnalyticsDashboard) SyncFieldsDuringRead(ctx context.Context, from ProviderAnalyticsDashboard) {
}

func (m ProviderAnalyticsDashboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ProviderAnalyticsDashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderAnalyticsDashboard
// only implements ToObjectValue() and Type().
func (m ProviderAnalyticsDashboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProviderAnalyticsDashboard) Type(ctx context.Context) attr.Type {
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

func (to *ProviderInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderInfo) {
}

func (to *ProviderInfo) SyncFieldsDuringRead(ctx context.Context, from ProviderInfo) {
}

func (m ProviderInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ProviderInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo
// only implements ToObjectValue() and Type().
func (m ProviderInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"business_contact_email":   m.BusinessContactEmail,
			"company_website_link":     m.CompanyWebsiteLink,
			"dark_mode_icon_file_id":   m.DarkModeIconFileId,
			"dark_mode_icon_file_path": m.DarkModeIconFilePath,
			"description":              m.Description,
			"icon_file_id":             m.IconFileId,
			"icon_file_path":           m.IconFilePath,
			"id":                       m.Id,
			"is_featured":              m.IsFeatured,
			"name":                     m.Name,
			"privacy_policy_link":      m.PrivacyPolicyLink,
			"published_by":             m.PublishedBy,
			"support_contact_email":    m.SupportContactEmail,
			"term_of_service_link":     m.TermOfServiceLink,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProviderInfo) Type(ctx context.Context) attr.Type {
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

func (to *RegionInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegionInfo) {
}

func (to *RegionInfo) SyncFieldsDuringRead(ctx context.Context, from RegionInfo) {
}

func (m RegionInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegionInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegionInfo
// only implements ToObjectValue() and Type().
func (m RegionInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud":  m.Cloud,
			"region": m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegionInfo) Type(ctx context.Context) attr.Type {
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

func (to *RemoveExchangeForListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveExchangeForListingRequest) {
}

func (to *RemoveExchangeForListingRequest) SyncFieldsDuringRead(ctx context.Context, from RemoveExchangeForListingRequest) {
}

func (m RemoveExchangeForListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RemoveExchangeForListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingRequest
// only implements ToObjectValue() and Type().
func (m RemoveExchangeForListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveExchangeForListingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type RemoveExchangeForListingResponse struct {
}

func (to *RemoveExchangeForListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveExchangeForListingResponse) {
}

func (to *RemoveExchangeForListingResponse) SyncFieldsDuringRead(ctx context.Context, from RemoveExchangeForListingResponse) {
}

func (m RemoveExchangeForListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveExchangeForListingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveExchangeForListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveExchangeForListingResponse
// only implements ToObjectValue() and Type().
func (m RemoveExchangeForListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveExchangeForListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl types.String `tfsdk:"git_repo_url"`
}

func (to *RepoInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInfo) {
}

func (to *RepoInfo) SyncFieldsDuringRead(ctx context.Context, from RepoInfo) {
}

func (m RepoInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo
// only implements ToObjectValue() and Type().
func (m RepoInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_repo_url": m.GitRepoUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoInfo) Type(ctx context.Context) attr.Type {
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

func (to *RepoInstallation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInstallation) {
}

func (to *RepoInstallation) SyncFieldsDuringRead(ctx context.Context, from RepoInstallation) {
}

func (m RepoInstallation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoInstallation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInstallation
// only implements ToObjectValue() and Type().
func (m RepoInstallation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_name": m.RepoName,
			"repo_path": m.RepoPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoInstallation) Type(ctx context.Context) attr.Type {
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

func (to *SearchListingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchListingsRequest) {
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

func (to *SearchListingsRequest) SyncFieldsDuringRead(ctx context.Context, from SearchListingsRequest) {
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

func (m SearchListingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchListingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets":       reflect.TypeOf(types.String{}),
		"categories":   reflect.TypeOf(types.String{}),
		"provider_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsRequest
// only implements ToObjectValue() and Type().
func (m SearchListingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":              m.Assets,
			"categories":          m.Categories,
			"is_free":             m.IsFree,
			"is_private_exchange": m.IsPrivateExchange,
			"page_size":           m.PageSize,
			"page_token":          m.PageToken,
			"provider_ids":        m.ProviderIds,
			"query":               m.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchListingsRequest) Type(ctx context.Context) attr.Type {
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
func (m *SearchListingsRequest) GetAssets(ctx context.Context) ([]types.String, bool) {
	if m.Assets.IsNull() || m.Assets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in SearchListingsRequest.
func (m *SearchListingsRequest) SetAssets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Assets = types.ListValueMust(t, vs)
}

// GetCategories returns the value of the Categories field in SearchListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchListingsRequest) GetCategories(ctx context.Context) ([]types.String, bool) {
	if m.Categories.IsNull() || m.Categories.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Categories.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCategories sets the value of the Categories field in SearchListingsRequest.
func (m *SearchListingsRequest) SetCategories(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["categories"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Categories = types.ListValueMust(t, vs)
}

// GetProviderIds returns the value of the ProviderIds field in SearchListingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchListingsRequest) GetProviderIds(ctx context.Context) ([]types.String, bool) {
	if m.ProviderIds.IsNull() || m.ProviderIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ProviderIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviderIds sets the value of the ProviderIds field in SearchListingsRequest.
func (m *SearchListingsRequest) SetProviderIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["provider_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ProviderIds = types.ListValueMust(t, vs)
}

type SearchListingsResponse struct {
	Listings types.List `tfsdk:"listings"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchListingsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (to *SearchListingsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchListingsResponse) {
	if !from.Listings.IsNull() && !from.Listings.IsUnknown() && to.Listings.IsNull() && len(from.Listings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Listings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Listings = from.Listings
	}
}

func (m SearchListingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchListingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listings": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchListingsResponse
// only implements ToObjectValue() and Type().
func (m SearchListingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listings":        m.Listings,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchListingsResponse) Type(ctx context.Context) attr.Type {
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
func (m *SearchListingsResponse) GetListings(ctx context.Context) ([]Listing, bool) {
	if m.Listings.IsNull() || m.Listings.IsUnknown() {
		return nil, false
	}
	var v []Listing
	d := m.Listings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListings sets the value of the Listings field in SearchListingsResponse.
func (m *SearchListingsResponse) SetListings(ctx context.Context, v []Listing) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["listings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Listings = types.ListValueMust(t, vs)
}

type ShareInfo struct {
	Name types.String `tfsdk:"name"`

	Type_ types.String `tfsdk:"type"`
}

func (to *ShareInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ShareInfo) {
}

func (to *ShareInfo) SyncFieldsDuringRead(ctx context.Context, from ShareInfo) {
}

func (m ShareInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ShareInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo
// only implements ToObjectValue() and Type().
func (m ShareInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
			"type": m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ShareInfo) Type(ctx context.Context) attr.Type {
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

func (to *SharedDataObject) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharedDataObject) {
}

func (to *SharedDataObject) SyncFieldsDuringRead(ctx context.Context, from SharedDataObject) {
}

func (m SharedDataObject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SharedDataObject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject
// only implements ToObjectValue() and Type().
func (m SharedDataObject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_object_type": m.DataObjectType,
			"name":             m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SharedDataObject) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_object_type": types.StringType,
			"name":             types.StringType,
		},
	}
}

type TokenDetail struct {
	BearerToken types.String `tfsdk:"bearer_token"`

	Endpoint types.String `tfsdk:"endpoint"`

	ExpirationTime types.String `tfsdk:"expiration_time"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

func (to *TokenDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenDetail) {
}

func (to *TokenDetail) SyncFieldsDuringRead(ctx context.Context, from TokenDetail) {
}

func (m TokenDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenDetail
// only implements ToObjectValue() and Type().
func (m TokenDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearer_token":              m.BearerToken,
			"endpoint":                  m.Endpoint,
			"expiration_time":           m.ExpirationTime,
			"share_credentials_version": m.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"expiration_time":           types.StringType,
			"share_credentials_version": types.Int64Type,
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

func (to *TokenInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenInfo) {
}

func (to *TokenInfo) SyncFieldsDuringRead(ctx context.Context, from TokenInfo) {
}

func (m TokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo
// only implements ToObjectValue() and Type().
func (m TokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url":  m.ActivationUrl,
			"created_at":      m.CreatedAt,
			"created_by":      m.CreatedBy,
			"expiration_time": m.ExpirationTime,
			"id":              m.Id,
			"updated_at":      m.UpdatedAt,
			"updated_by":      m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenInfo) Type(ctx context.Context) attr.Type {
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

func (to *UpdateExchangeFilterRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeFilterRequest) {
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

func (to *UpdateExchangeFilterRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeFilterRequest) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m UpdateExchangeFilterRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetRequired()
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
func (m UpdateExchangeFilterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterRequest
// only implements ToObjectValue() and Type().
func (m UpdateExchangeFilterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": m.Filter,
			"id":     m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExchangeFilterRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateExchangeFilterRequest) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v ExchangeFilter
	d := m.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterRequest.
func (m *UpdateExchangeFilterRequest) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	m.Filter = vs
}

type UpdateExchangeFilterResponse struct {
	Filter types.Object `tfsdk:"filter"`
}

func (to *UpdateExchangeFilterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeFilterResponse) {
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

func (to *UpdateExchangeFilterResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeFilterResponse) {
	if !from.Filter.IsNull() && !from.Filter.IsUnknown() {
		if toFilter, ok := to.GetFilter(ctx); ok {
			if fromFilter, ok := from.GetFilter(ctx); ok {
				toFilter.SyncFieldsDuringRead(ctx, fromFilter)
				to.SetFilter(ctx, toFilter)
			}
		}
	}
}

func (m UpdateExchangeFilterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateExchangeFilterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter": reflect.TypeOf(ExchangeFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeFilterResponse
// only implements ToObjectValue() and Type().
func (m UpdateExchangeFilterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter": m.Filter,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExchangeFilterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter": ExchangeFilter{}.Type(ctx),
		},
	}
}

// GetFilter returns the value of the Filter field in UpdateExchangeFilterResponse as
// a ExchangeFilter value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateExchangeFilterResponse) GetFilter(ctx context.Context) (ExchangeFilter, bool) {
	var e ExchangeFilter
	if m.Filter.IsNull() || m.Filter.IsUnknown() {
		return e, false
	}
	var v ExchangeFilter
	d := m.Filter.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilter sets the value of the Filter field in UpdateExchangeFilterResponse.
func (m *UpdateExchangeFilterResponse) SetFilter(ctx context.Context, v ExchangeFilter) {
	vs := v.ToObjectValue(ctx)
	m.Filter = vs
}

type UpdateExchangeRequest struct {
	Exchange types.Object `tfsdk:"exchange"`

	Id types.String `tfsdk:"-"`
}

func (to *UpdateExchangeRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeRequest) {
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

func (to *UpdateExchangeRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeRequest) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (m UpdateExchangeRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exchange"] = attrs["exchange"].SetRequired()
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
func (m UpdateExchangeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeRequest
// only implements ToObjectValue() and Type().
func (m UpdateExchangeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": m.Exchange,
			"id":       m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExchangeRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateExchangeRequest) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if m.Exchange.IsNull() || m.Exchange.IsUnknown() {
		return e, false
	}
	var v Exchange
	d := m.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeRequest.
func (m *UpdateExchangeRequest) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	m.Exchange = vs
}

type UpdateExchangeResponse struct {
	Exchange types.Object `tfsdk:"exchange"`
}

func (to *UpdateExchangeResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExchangeResponse) {
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

func (to *UpdateExchangeResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateExchangeResponse) {
	if !from.Exchange.IsNull() && !from.Exchange.IsUnknown() {
		if toExchange, ok := to.GetExchange(ctx); ok {
			if fromExchange, ok := from.GetExchange(ctx); ok {
				toExchange.SyncFieldsDuringRead(ctx, fromExchange)
				to.SetExchange(ctx, toExchange)
			}
		}
	}
}

func (m UpdateExchangeResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateExchangeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exchange": reflect.TypeOf(Exchange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExchangeResponse
// only implements ToObjectValue() and Type().
func (m UpdateExchangeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exchange": m.Exchange,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExchangeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exchange": Exchange{}.Type(ctx),
		},
	}
}

// GetExchange returns the value of the Exchange field in UpdateExchangeResponse as
// a Exchange value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateExchangeResponse) GetExchange(ctx context.Context) (Exchange, bool) {
	var e Exchange
	if m.Exchange.IsNull() || m.Exchange.IsUnknown() {
		return e, false
	}
	var v Exchange
	d := m.Exchange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExchange sets the value of the Exchange field in UpdateExchangeResponse.
func (m *UpdateExchangeResponse) SetExchange(ctx context.Context, v Exchange) {
	vs := v.ToObjectValue(ctx)
	m.Exchange = vs
}

type UpdateInstallationRequest struct {
	Installation types.Object `tfsdk:"installation"`

	InstallationId types.String `tfsdk:"-"`

	ListingId types.String `tfsdk:"-"`

	RotateToken types.Bool `tfsdk:"rotate_token"`
}

func (to *UpdateInstallationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInstallationRequest) {
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

func (to *UpdateInstallationRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateInstallationRequest) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (m UpdateInstallationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["installation"] = attrs["installation"].SetRequired()
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
func (m UpdateInstallationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationRequest
// only implements ToObjectValue() and Type().
func (m UpdateInstallationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation":    m.Installation,
			"installation_id": m.InstallationId,
			"listing_id":      m.ListingId,
			"rotate_token":    m.RotateToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateInstallationRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateInstallationRequest) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if m.Installation.IsNull() || m.Installation.IsUnknown() {
		return e, false
	}
	var v InstallationDetail
	d := m.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationRequest.
func (m *UpdateInstallationRequest) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	m.Installation = vs
}

type UpdateInstallationResponse struct {
	Installation types.Object `tfsdk:"installation"`
}

func (to *UpdateInstallationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInstallationResponse) {
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

func (to *UpdateInstallationResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateInstallationResponse) {
	if !from.Installation.IsNull() && !from.Installation.IsUnknown() {
		if toInstallation, ok := to.GetInstallation(ctx); ok {
			if fromInstallation, ok := from.GetInstallation(ctx); ok {
				toInstallation.SyncFieldsDuringRead(ctx, fromInstallation)
				to.SetInstallation(ctx, toInstallation)
			}
		}
	}
}

func (m UpdateInstallationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateInstallationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"installation": reflect.TypeOf(InstallationDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInstallationResponse
// only implements ToObjectValue() and Type().
func (m UpdateInstallationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"installation": m.Installation,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateInstallationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"installation": InstallationDetail{}.Type(ctx),
		},
	}
}

// GetInstallation returns the value of the Installation field in UpdateInstallationResponse as
// a InstallationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInstallationResponse) GetInstallation(ctx context.Context) (InstallationDetail, bool) {
	var e InstallationDetail
	if m.Installation.IsNull() || m.Installation.IsUnknown() {
		return e, false
	}
	var v InstallationDetail
	d := m.Installation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstallation sets the value of the Installation field in UpdateInstallationResponse.
func (m *UpdateInstallationResponse) SetInstallation(ctx context.Context, v InstallationDetail) {
	vs := v.ToObjectValue(ctx)
	m.Installation = vs
}

type UpdateListingRequest struct {
	Id types.String `tfsdk:"-"`

	Listing types.Object `tfsdk:"listing"`
}

func (to *UpdateListingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateListingRequest) {
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

func (to *UpdateListingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateListingRequest) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (m UpdateListingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["listing"] = attrs["listing"].SetRequired()
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
func (m UpdateListingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingRequest
// only implements ToObjectValue() and Type().
func (m UpdateListingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      m.Id,
			"listing": m.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateListingRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateListingRequest) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if m.Listing.IsNull() || m.Listing.IsUnknown() {
		return e, false
	}
	var v Listing
	d := m.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListing sets the value of the Listing field in UpdateListingRequest.
func (m *UpdateListingRequest) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	m.Listing = vs
}

type UpdateListingResponse struct {
	Listing types.Object `tfsdk:"listing"`
}

func (to *UpdateListingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateListingResponse) {
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

func (to *UpdateListingResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateListingResponse) {
	if !from.Listing.IsNull() && !from.Listing.IsUnknown() {
		if toListing, ok := to.GetListing(ctx); ok {
			if fromListing, ok := from.GetListing(ctx); ok {
				toListing.SyncFieldsDuringRead(ctx, fromListing)
				to.SetListing(ctx, toListing)
			}
		}
	}
}

func (m UpdateListingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateListingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"listing": reflect.TypeOf(Listing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateListingResponse
// only implements ToObjectValue() and Type().
func (m UpdateListingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing": m.Listing,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateListingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"listing": Listing{}.Type(ctx),
		},
	}
}

// GetListing returns the value of the Listing field in UpdateListingResponse as
// a Listing value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateListingResponse) GetListing(ctx context.Context) (Listing, bool) {
	var e Listing
	if m.Listing.IsNull() || m.Listing.IsUnknown() {
		return e, false
	}
	var v Listing
	d := m.Listing.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListing sets the value of the Listing field in UpdateListingResponse.
func (m *UpdateListingResponse) SetListing(ctx context.Context, v Listing) {
	vs := v.ToObjectValue(ctx)
	m.Listing = vs
}

type UpdatePersonalizationRequestRequest struct {
	ListingId types.String `tfsdk:"-"`

	Reason types.String `tfsdk:"reason"`

	RequestId types.String `tfsdk:"-"`

	Share types.Object `tfsdk:"share"`

	Status types.String `tfsdk:"status"`
}

func (to *UpdatePersonalizationRequestRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalizationRequestRequest) {
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

func (to *UpdatePersonalizationRequestRequest) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalizationRequestRequest) {
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				toShare.SyncFieldsDuringRead(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
}

func (m UpdatePersonalizationRequestRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["reason"] = attrs["reason"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
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
func (m UpdatePersonalizationRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"share": reflect.TypeOf(ShareInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestRequest
// only implements ToObjectValue() and Type().
func (m UpdatePersonalizationRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"listing_id": m.ListingId,
			"reason":     m.Reason,
			"request_id": m.RequestId,
			"share":      m.Share,
			"status":     m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePersonalizationRequestRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdatePersonalizationRequestRequest) GetShare(ctx context.Context) (ShareInfo, bool) {
	var e ShareInfo
	if m.Share.IsNull() || m.Share.IsUnknown() {
		return e, false
	}
	var v ShareInfo
	d := m.Share.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShare sets the value of the Share field in UpdatePersonalizationRequestRequest.
func (m *UpdatePersonalizationRequestRequest) SetShare(ctx context.Context, v ShareInfo) {
	vs := v.ToObjectValue(ctx)
	m.Share = vs
}

type UpdatePersonalizationRequestResponse struct {
	Request types.Object `tfsdk:"request"`
}

func (to *UpdatePersonalizationRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalizationRequestResponse) {
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

func (to *UpdatePersonalizationRequestResponse) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalizationRequestResponse) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				toRequest.SyncFieldsDuringRead(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (m UpdatePersonalizationRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdatePersonalizationRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(PersonalizationRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalizationRequestResponse
// only implements ToObjectValue() and Type().
func (m UpdatePersonalizationRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": m.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePersonalizationRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": PersonalizationRequest{}.Type(ctx),
		},
	}
}

// GetRequest returns the value of the Request field in UpdatePersonalizationRequestResponse as
// a PersonalizationRequest value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePersonalizationRequestResponse) GetRequest(ctx context.Context) (PersonalizationRequest, bool) {
	var e PersonalizationRequest
	if m.Request.IsNull() || m.Request.IsUnknown() {
		return e, false
	}
	var v PersonalizationRequest
	d := m.Request.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequest sets the value of the Request field in UpdatePersonalizationRequestResponse.
func (m *UpdatePersonalizationRequestResponse) SetRequest(ctx context.Context, v PersonalizationRequest) {
	vs := v.ToObjectValue(ctx)
	m.Request = vs
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id types.String `tfsdk:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version types.Int64 `tfsdk:"version"`
}

func (to *UpdateProviderAnalyticsDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderAnalyticsDashboardRequest) {
}

func (to *UpdateProviderAnalyticsDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderAnalyticsDashboardRequest) {
}

func (m UpdateProviderAnalyticsDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateProviderAnalyticsDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardRequest
// only implements ToObjectValue() and Type().
func (m UpdateProviderAnalyticsDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      m.Id,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProviderAnalyticsDashboardRequest) Type(ctx context.Context) attr.Type {
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

func (to *UpdateProviderAnalyticsDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderAnalyticsDashboardResponse) {
}

func (to *UpdateProviderAnalyticsDashboardResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderAnalyticsDashboardResponse) {
}

func (m UpdateProviderAnalyticsDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateProviderAnalyticsDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderAnalyticsDashboardResponse
// only implements ToObjectValue() and Type().
func (m UpdateProviderAnalyticsDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
			"id":           m.Id,
			"version":      m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProviderAnalyticsDashboardResponse) Type(ctx context.Context) attr.Type {
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

func (to *UpdateProviderRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderRequest) {
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

func (to *UpdateProviderRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderRequest) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (m UpdateProviderRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetRequired()
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
func (m UpdateProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderRequest
// only implements ToObjectValue() and Type().
func (m UpdateProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":       m.Id,
			"provider": m.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProviderRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateProviderRequest) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if m.Provider.IsNull() || m.Provider.IsUnknown() {
		return e, false
	}
	var v ProviderInfo
	d := m.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProvider sets the value of the Provider field in UpdateProviderRequest.
func (m *UpdateProviderRequest) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	m.Provider = vs
}

type UpdateProviderResponse struct {
	Provider types.Object `tfsdk:"provider"`
}

func (to *UpdateProviderResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProviderResponse) {
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

func (to *UpdateProviderResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateProviderResponse) {
	if !from.Provider.IsNull() && !from.Provider.IsUnknown() {
		if toProvider, ok := to.GetProvider(ctx); ok {
			if fromProvider, ok := from.GetProvider(ctx); ok {
				toProvider.SyncFieldsDuringRead(ctx, fromProvider)
				to.SetProvider(ctx, toProvider)
			}
		}
	}
}

func (m UpdateProviderResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateProviderResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProviderResponse
// only implements ToObjectValue() and Type().
func (m UpdateProviderResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": m.Provider,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProviderResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": ProviderInfo{}.Type(ctx),
		},
	}
}

// GetProvider returns the value of the Provider field in UpdateProviderResponse as
// a ProviderInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProviderResponse) GetProvider(ctx context.Context) (ProviderInfo, bool) {
	var e ProviderInfo
	if m.Provider.IsNull() || m.Provider.IsUnknown() {
		return e, false
	}
	var v ProviderInfo
	d := m.Provider.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProvider sets the value of the Provider field in UpdateProviderResponse.
func (m *UpdateProviderResponse) SetProvider(ctx context.Context, v ProviderInfo) {
	vs := v.ToObjectValue(ctx)
	m.Provider = vs
}
