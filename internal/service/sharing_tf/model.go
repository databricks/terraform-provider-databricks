// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sharing_tf

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:""`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProvider) {
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringRead(existingState CreateProvider) {
}

func (a CreateProvider) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateProvider) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":   types.StringType,
			"comment":               types.StringType,
			"name":                  types.StringType,
			"recipient_profile_str": types.StringType,
		},
	}
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:""`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRecipient) {
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringRead(existingState CreateRecipient) {
}

func (a CreateRecipient) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
	}
}

func (a CreateRecipient) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":                types.StringType,
			"comment":                            types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"expiration_time":                    types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.ToAttrType(ctx),
			},
			"name":  types.StringType,
			"owner": types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.ToAttrType(ctx),
			},
			"sharing_code": types.StringType,
		},
	}
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:""`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateShare) {
}

func (newState *CreateShare) SyncEffectiveFieldsDuringRead(existingState CreateShare) {
}

func (a CreateShare) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateShare) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"storage_root": types.StringType,
		},
	}
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
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
			"name": types.StringType,
		},
	}
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRecipientRequest) {
}

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRecipientRequest) {
}

func (a DeleteRecipientRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRecipientRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (a DeleteResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteShareRequest) {
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringRead(existingState DeleteShareRequest) {
}

func (a DeleteShareRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteShareRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoRequest) {
}

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoRequest) {
}

func (a GetActivationUrlInfoRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetActivationUrlInfoRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type GetActivationUrlInfoResponse struct {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoResponse) {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoResponse) {
}

func (a GetActivationUrlInfoResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetActivationUrlInfoResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
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
			"name": types.StringType,
		},
	}
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientRequest) {
}

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringRead(existingState GetRecipientRequest) {
}

func (a GetRecipientRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRecipientRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share permissions for a recipient.
	PermissionsOut types.List `tfsdk:"permissions_out" tf:"optional"`
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientSharePermissionsResponse) {
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState GetRecipientSharePermissionsResponse) {
}

func (a GetRecipientSharePermissionsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions_out": reflect.TypeOf(ShareToPrivilegeAssignment{}),
	}
}

func (a GetRecipientSharePermissionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"permissions_out": basetypes.ListType{
				ElemType: ShareToPrivilegeAssignment{}.ToAttrType(ctx),
			},
		},
	}
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetShareRequest) {
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringRead(existingState GetShareRequest) {
}

func (a GetShareRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetShareRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_shared_data": types.BoolType,
			"name":                types.StringType,
		},
	}
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses types.List `tfsdk:"allowed_ip_addresses" tf:"optional"`
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessList) {
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringRead(existingState IpAccessList) {
}

func (a IpAccessList) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_ip_addresses": reflect.TypeOf(types.StringType),
	}
}

func (a IpAccessList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider shares.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderSharesResponse) {
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderSharesResponse) {
}

func (a ListProviderSharesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ProviderShare{}),
	}
}

func (a ListProviderSharesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ProviderShare{}.ToAttrType(ctx),
			},
		},
	}
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
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
			"data_provider_global_metastore_id": types.StringType,
			"max_results":                       types.Int64Type,
			"page_token":                        types.StringType,
		},
	}
}

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider information objects.
	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse) {
}

func (a ListProvidersResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

func (a ListProvidersResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsRequest) {
}

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringRead(existingState ListRecipientsRequest) {
}

func (a ListRecipientsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListRecipientsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_recipient_global_metastore_id": types.StringType,
			"max_results":                        types.Int64Type,
			"page_token":                         types.StringType,
		},
	}
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of recipient information objects.
	Recipients types.List `tfsdk:"recipients" tf:"optional"`
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsResponse) {
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringRead(existingState ListRecipientsResponse) {
}

func (a ListRecipientsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipients": reflect.TypeOf(RecipientInfo{}),
	}
}

func (a ListRecipientsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"recipients": basetypes.ListType{
				ElemType: RecipientInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// List shares by Provider
type ListSharesRequest struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the provider in which to list shares.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesRequest) {
}

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringRead(existingState ListSharesRequest) {
}

func (a ListSharesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListSharesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share information objects.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesResponse) {
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListSharesResponse) {
}

func (a ListSharesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ShareInfo{}),
	}
}

func (a ListSharesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ShareInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type Partition struct {
	// An array of partition values.
	Values types.List `tfsdk:"value" tf:"optional"`
}

func (newState *Partition) SyncEffectiveFieldsDuringCreateOrUpdate(plan Partition) {
}

func (newState *Partition) SyncEffectiveFieldsDuringRead(existingState Partition) {
}

func (a Partition) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue{}),
	}
}

func (a Partition) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue{}.ToAttrType(ctx),
			},
		},
	}
}

type PartitionValue struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op" tf:"optional"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key" tf:"optional"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionValue) {
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringRead(existingState PartitionValue) {
}

func (a PartitionValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PartitionValue) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                   types.StringType,
			"op":                     types.StringType,
			"recipient_property_key": types.StringType,
			"value":                  types.StringType,
		},
	}
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges types.List `tfsdk:"privileges" tf:"optional"`
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment) {
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment) {
}

func (a PrivilegeAssignment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(types.StringType),
	}
}

func (a PrivilegeAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id" tf:"optional"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile types.List `tfsdk:"recipient_profile" tf:"optional,object"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
}

func (a ProviderInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipient_profile": reflect.TypeOf(RecipientProfile{}),
	}
}

func (a ProviderInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":               types.StringType,
			"cloud":                             types.StringType,
			"comment":                           types.StringType,
			"created_at":                        types.Int64Type,
			"created_by":                        types.StringType,
			"data_provider_global_metastore_id": types.StringType,
			"metastore_id":                      types.StringType,
			"name":                              types.StringType,
			"owner":                             types.StringType,
			"recipient_profile": basetypes.ListType{
				ElemType: RecipientProfile{}.ToAttrType(ctx),
			},
			"recipient_profile_str": types.StringType,
			"region":                types.StringType,
			"updated_at":            types.Int64Type,
			"updated_by":            types.StringType,
		},
	}
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderShare) {
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringRead(existingState ProviderShare) {
}

func (a ProviderShare) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ProviderShare) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated" tf:"optional"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens types.List `tfsdk:"tokens" tf:"optional"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo) {
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringRead(existingState RecipientInfo) {
}

func (a RecipientInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
		"tokens":             reflect.TypeOf(RecipientTokenInfo{}),
	}
}

func (a RecipientInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activated":                          types.BoolType,
			"activation_url":                     types.StringType,
			"authentication_type":                types.StringType,
			"cloud":                              types.StringType,
			"comment":                            types.StringType,
			"created_at":                         types.Int64Type,
			"created_by":                         types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.ToAttrType(ctx),
			},
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.ToAttrType(ctx),
			},
			"region":       types.StringType,
			"sharing_code": types.StringType,
			"tokens": basetypes.ListType{
				ElemType: RecipientTokenInfo{}.ToAttrType(ctx),
			},
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version" tf:"optional"`
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientProfile) {
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringRead(existingState RecipientProfile) {
}

func (a RecipientProfile) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RecipientProfile) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"share_credentials_version": types.Int64Type,
		},
	}
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo) {
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo) {
}

func (a RecipientTokenInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RecipientTokenInfo) ToAttrType(ctx context.Context) types.ObjectType {
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

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenRequest) {
}

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenRequest) {
}

func (a RetrieveTokenRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RetrieveTokenRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenResponse) {
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenResponse) {
}

func (a RetrieveTokenResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RetrieveTokenResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearerToken":             types.StringType,
			"endpoint":                types.StringType,
			"expirationTime":          types.StringType,
			"shareCredentialsVersion": types.Int64Type,
		},
	}
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds" tf:""`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan RotateRecipientToken) {
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringRead(existingState RotateRecipientToken) {
}

func (a RotateRecipientToken) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RotateRecipientToken) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"existing_token_expire_in_seconds": types.Int64Type,
			"name":                             types.StringType,
		},
	}
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:""`
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecurablePropertiesKvPairs) {
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState SecurablePropertiesKvPairs) {
}

func (a SecurablePropertiesKvPairs) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.StringType),
	}
}

func (a SecurablePropertiesKvPairs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type ShareInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"computed,optional"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"computed,optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A list of shared data objects within the share.
	Objects types.List `tfsdk:"object" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed,optional"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"computed,optional"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"computed,optional"`
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
}

func (a ShareInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(SharedDataObject{}),
	}
}

func (a ShareInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"name":       types.StringType,
			"object": basetypes.ListType{
				ElemType: SharedDataObject{}.ToAttrType(ctx),
			},
			"owner":            types.StringType,
			"storage_location": types.StringType,
			"storage_root":     types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
		},
	}
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the Recipient.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharePermissionsRequest) {
}

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState SharePermissionsRequest) {
}

func (a SharePermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SharePermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
	// The share name.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareToPrivilegeAssignment) {
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState ShareToPrivilegeAssignment) {
}

func (a ShareToPrivilegeAssignment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment{}),
	}
}

func (a ShareToPrivilegeAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.ToAttrType(ctx),
			},
			"share_name": types.StringType,
		},
	}
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"computed,optional"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by" tf:"computed,optional"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled types.Bool `tfsdk:"cdf_enabled" tf:"computed,optional"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The type of the data object.
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus types.String `tfsdk:"history_data_sharing_status" tf:"computed,optional"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name types.String `tfsdk:"name" tf:""`
	// Array of partitions for the shared data.
	Partitions types.List `tfsdk:"partition" tf:"optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs types.String `tfsdk:"shared_as" tf:"computed,optional"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion types.Int64 `tfsdk:"start_version" tf:"computed,optional"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status" tf:"computed,optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as" tf:"optional"`
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject) {
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringRead(existingState SharedDataObject) {
}

func (a SharedDataObject) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition": reflect.TypeOf(Partition{}),
	}
}

func (a SharedDataObject) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":                    types.Int64Type,
			"added_by":                    types.StringType,
			"cdf_enabled":                 types.BoolType,
			"comment":                     types.StringType,
			"content":                     types.StringType,
			"data_object_type":            types.StringType,
			"history_data_sharing_status": types.StringType,
			"name":                        types.StringType,
			"partition": basetypes.ListType{
				ElemType: Partition{}.ToAttrType(ctx),
			},
			"shared_as":        types.StringType,
			"start_version":    types.Int64Type,
			"status":           types.StringType,
			"string_shared_as": types.StringType,
		},
	}
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action" tf:"optional"`
	// The data object that is being added, removed, or updated.
	DataObject types.List `tfsdk:"data_object" tf:"optional,object"`
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObjectUpdate) {
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringRead(existingState SharedDataObjectUpdate) {
}

func (a SharedDataObjectUpdate) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_object": reflect.TypeOf(SharedDataObject{}),
	}
}

func (a SharedDataObjectUpdate) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action": types.StringType,
			"data_object": basetypes.ListType{
				ElemType: SharedDataObject{}.ToAttrType(ctx),
			},
		},
	}
}

type UpdatePermissionsResponse struct {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissionsResponse) {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState UpdatePermissionsResponse) {
}

func (a UpdatePermissionsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdatePermissionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateProvider struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProvider) {
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringRead(existingState UpdateProvider) {
}

func (a UpdateProvider) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateProvider) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":               types.StringType,
			"name":                  types.StringType,
			"new_name":              types.StringType,
			"owner":                 types.StringType,
			"recipient_profile_str": types.StringType,
		},
	}
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRecipient) {
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringRead(existingState UpdateRecipient) {
}

func (a UpdateRecipient) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
	}
}

func (a UpdateRecipient) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"expiration_time": types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.ToAttrType(ctx),
			},
			"name":     types.StringType,
			"new_name": types.StringType,
			"owner":    types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.ToAttrType(ctx),
			},
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

func (a UpdateResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed,optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Array of shared data object updates.
	Updates types.List `tfsdk:"updates" tf:"optional"`
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare) {
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringRead(existingState UpdateShare) {
}

func (a UpdateShare) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(SharedDataObjectUpdate{}),
	}
}

func (a UpdateShare) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"new_name":     types.StringType,
			"owner":        types.StringType,
			"storage_root": types.StringType,
			"updates": basetypes.ListType{
				ElemType: SharedDataObjectUpdate{}.ToAttrType(ctx),
			},
		},
	}
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes types.List `tfsdk:"changes" tf:"optional"`
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions) {
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions) {
}

func (a UpdateSharePermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(catalog.PermissionsChange{}),
	}
}

func (a UpdateSharePermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: catalog_tf.PermissionsChange{}.ToAttrType(ctx),
			},
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}
