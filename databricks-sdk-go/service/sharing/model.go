// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
)

// The delta sharing authentication type.
type AuthenticationType string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeOauthClientCredentials AuthenticationType = `OAUTH_CLIENT_CREDENTIALS`

const AuthenticationTypeOidcFederation AuthenticationType = `OIDC_FEDERATION`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

// String representation for [fmt.Print]
func (f *AuthenticationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `OAUTH_CLIENT_CREDENTIALS`, `OIDC_FEDERATION`, `TOKEN`:
		*f = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "OAUTH_CLIENT_CREDENTIALS", "OIDC_FEDERATION", "TOKEN"`, v)
	}
}

// Values returns all possible values for AuthenticationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeDatabricks,
		AuthenticationTypeOauthClientCredentials,
		AuthenticationTypeOidcFederation,
		AuthenticationTypeToken,
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (f *AuthenticationType) Type() string {
	return "AuthenticationType"
}

// UC supported column types Copied from
// https://src.dev.databricks.com/databricks/universe@23a85902bb58695ab9293adc9f327b0714b55e72/-/blob/managed-catalog/api/messages/table.proto?L68
type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeName = `VARIANT`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE", "VARIANT"`, v)
	}
}

// Values returns all possible values for ColumnTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnTypeName) Values() []ColumnTypeName {
	return []ColumnTypeName{
		ColumnTypeNameArray,
		ColumnTypeNameBinary,
		ColumnTypeNameBoolean,
		ColumnTypeNameByte,
		ColumnTypeNameChar,
		ColumnTypeNameDate,
		ColumnTypeNameDecimal,
		ColumnTypeNameDouble,
		ColumnTypeNameFloat,
		ColumnTypeNameInt,
		ColumnTypeNameInterval,
		ColumnTypeNameLong,
		ColumnTypeNameMap,
		ColumnTypeNameNull,
		ColumnTypeNameShort,
		ColumnTypeNameString,
		ColumnTypeNameStruct,
		ColumnTypeNameTableType,
		ColumnTypeNameTimestamp,
		ColumnTypeNameTimestampNtz,
		ColumnTypeNameUserDefinedType,
		ColumnTypeNameVariant,
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

type CreateFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be created.
	Policy FederationPolicy `json:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being created.
	RecipientName string `json:"-" url:"-"`
}

func (s *CreateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateProvider struct {
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRecipient struct {
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// [Create,Update:IGN] common - id of the recipient
	Id string `json:"id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	Name string `json:"name"`
	// Cloud storage path, within an external location, under which the share
	// stores the notebook files added to it. When unset, notebook files are
	// stored under the metastore root storage location.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be deleted.
	Name string `json:"-" url:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being deleted.
	RecipientName string `json:"-" url:"-"`
}

func (s *DeleteFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" url:"-"`
}

func (s *DeleteProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" url:"-"`
}

func (s *DeleteRecipientRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteShareRequest struct {
	// The name of the share.
	Name string `json:"-" url:"-"`
}

func (s *DeleteShareRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a UC dependency.
type DeltaSharingDependency struct {
	Function *DeltaSharingFunctionDependency `json:"function,omitempty"`

	Table *DeltaSharingTableDependency `json:"table,omitempty"`
}

func (s *DeltaSharingDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a list of dependencies.
type DeltaSharingDependencyList struct {
	// An array of Dependency.
	Dependencies []DeltaSharingDependency `json:"dependencies,omitempty"`
}

func (s *DeltaSharingDependencyList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeltaSharingFunction struct {
	// The aliass of registered model.
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// The comment of the function.
	Comment string `json:"comment,omitempty"`
	// The data type of the function.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// The dependency list of the function.
	DependencyList *DeltaSharingDependencyList `json:"dependency_list,omitempty"`
	// The full data type of the function.
	FullDataType string `json:"full_data_type,omitempty"`
	// The id of the function.
	Id string `json:"id,omitempty"`
	// The function parameter information.
	InputParams *FunctionParameterInfos `json:"input_params,omitempty"`
	// The name of the function.
	Name string `json:"name,omitempty"`
	// The properties of the function.
	Properties string `json:"properties,omitempty"`
	// The routine definition of the function.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// The name of the schema that the function belongs to.
	Schema string `json:"schema,omitempty"`
	// The securable kind of the function.
	SecurableKind SharedSecurableKind `json:"securable_kind,omitempty"`
	// The name of the share that the function belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the function belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The storage location of the function.
	StorageLocation string `json:"storage_location,omitempty"`
	// The tags of the function.
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSharingFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {
	FunctionName string `json:"function_name,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSharingFunctionDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingFunctionDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {
	SchemaName string `json:"schema_name,omitempty"`

	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSharingTableDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingTableDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FederationPolicy struct {
	// Description of the policy. This is a user-provided description.
	Comment string `json:"comment,omitempty"`
	// System-generated timestamp indicating when the policy was created.
	CreateTime string `json:"create_time,omitempty"`
	// Unique, immutable system-generated identifier for the federation policy.
	Id string `json:"id,omitempty"`
	// Name of the federation policy. A recipient can have multiple policies
	// with different names. The name must contain only lowercase alphanumeric
	// characters, numbers, and hyphens.
	Name string `json:"name,omitempty"`
	// Specifies the policy to use for validating OIDC claims in the federated
	// tokens.
	OidcPolicy *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	// System-generated timestamp indicating when the policy was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FederationPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FederationPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo struct {
	// The comment of the parameter.
	Comment string `json:"comment,omitempty"`
	// The name of the parameter.
	Name string `json:"name,omitempty"`
	// The default value of the parameter.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of the function parameter.
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// The position of the parameter.
	Position int `json:"position,omitempty"`
	// The interval type of the parameter type.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The type of the parameter in JSON format.
	TypeJson string `json:"type_json,omitempty"`
	// The type of the parameter in Enum format.
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// The precision of the parameter type.
	TypePrecision int `json:"type_precision,omitempty"`
	// The scale of the parameter type.
	TypeScale int `json:"type_scale,omitempty"`
	// The type of the parameter in text format.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func (s *FunctionParameterInfos) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FunctionParameterMode string

const FunctionParameterModeIn FunctionParameterMode = `IN`

const FunctionParameterModeInout FunctionParameterMode = `INOUT`

const FunctionParameterModeOut FunctionParameterMode = `OUT`

// String representation for [fmt.Print]
func (f *FunctionParameterMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`, `INOUT`, `OUT`:
		*f = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN", "INOUT", "OUT"`, v)
	}
}

// Values returns all possible values for FunctionParameterMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterMode) Values() []FunctionParameterMode {
	return []FunctionParameterMode{
		FunctionParameterModeIn,
		FunctionParameterModeInout,
		FunctionParameterModeOut,
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

type FunctionParameterType string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (f *FunctionParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*f = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Values returns all possible values for FunctionParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterType) Values() []FunctionParameterType {
	return []FunctionParameterType{
		FunctionParameterTypeColumn,
		FunctionParameterTypeParam,
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" url:"-"`
}

func (s *GetActivationUrlInfoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be retrieved.
	Name string `json:"-" url:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being retrieved.
	RecipientName string `json:"-" url:"-"`
}

func (s *GetFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" url:"-"`
}

func (s *GetProviderRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" url:"-"`
}

func (s *GetRecipientRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRecipientSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRecipientSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetShareRequest struct {
	// Whether to include the share's data objects in the response. When false
	// or omitted, the response omits `objects`.
	IncludeSharedData bool `json:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetShareRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetShareRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

func (s *IpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListFederationPoliciesRequest struct {
	MaxResults int `json:"-" url:"max_results,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of the recipient. This is the name of the recipient for which the
	// policies are being listed.
	RecipientName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFederationPoliciesResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []FederationPolicy `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProviderShareAssetsRequest struct {
	// Maximum number of functions to return.
	FunctionMaxResults int `json:"-" url:"function_max_results,omitempty"`
	// Maximum number of notebooks to return.
	NotebookMaxResults int `json:"-" url:"notebook_max_results,omitempty"`
	// The name of the provider who owns the share.
	ProviderName string `json:"-" url:"-"`
	// The name of the share.
	ShareName string `json:"-" url:"-"`
	// Maximum number of skills to return.
	SkillMaxResults int `json:"-" url:"skill_max_results,omitempty"`
	// Maximum number of tables to return.
	TableMaxResults int `json:"-" url:"table_max_results,omitempty"`
	// Maximum number of volumes to return.
	VolumeMaxResults int `json:"-" url:"volume_max_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProviderShareAssetsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderShareAssetsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to ListProviderShareAssets, which contains the list of assets of a
// share.
type ListProviderShareAssetsResponse struct {
	// The list of functions in the share.
	Functions []DeltaSharingFunction `json:"functions,omitempty"`
	// The list of notebooks in the share.
	Notebooks []NotebookFile `json:"notebooks,omitempty"`
	// The metadata of the share.
	Share *Share `json:"share,omitempty"`
	// The list of skills in the share.
	Skills []Skill `json:"skills,omitempty"`
	// The list of tables in the share.
	Tables []Table `json:"tables,omitempty"`
	// The list of volumes in the share.
	Volumes []Volume `json:"volumes,omitempty"`
}

func (s *ListProviderShareAssetsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider shares.
	Shares []ProviderShare `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProviderSharesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderSharesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId string `json:"-" url:"data_provider_global_metastore_id,omitempty"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
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
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider information objects.
	Providers []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId string `json:"-" url:"data_recipient_global_metastore_id,omitempty"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRecipientsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRecipientsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of recipient information objects.
	Recipients []RecipientInfo `json:"recipients,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRecipientsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRecipientsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Name of the provider in which to list shares.
	Name string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSharesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSharesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share information objects.
	Shares []ShareInfo `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSharesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSharesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotebookFile struct {
	// The comment of the notebook file.
	Comment string `json:"comment,omitempty"`
	// The id of the notebook file.
	Id string `json:"id,omitempty"`
	// Name of the notebook file.
	Name string `json:"name,omitempty"`
	// The name of the share that the notebook file belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the notebook file belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The tags of the notebook file.
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookFile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookFile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specifies the policy to use for validating OIDC claims in your federated
// tokens from Delta Sharing Clients. Refer to
// https://docs.databricks.com/en/delta-sharing/create-recipient-oidc-fed for
// more details.
type OidcFederationPolicy struct {
	// The allowed token audiences, as specified in the 'aud' claim of federated
	// tokens. The audience identifier is intended to represent the recipient of
	// the token. Can be any non-empty string value. As long as the audience in
	// the token matches at least one audience in the policy,
	Audiences []string `json:"audiences,omitempty"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	Issuer string `json:"issuer"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. The subject claim identifies the identity of the user
	// or machine accessing the resource. Examples for Entra ID (AAD): - U2M
	// flow (group access): If the subject claim is `groups`, this must be the
	// Object ID of the group in Entra ID. - U2M flow (user access): If the
	// subject claim is `oid`, this must be the Object ID of the user in Entra
	// ID. - M2M flow (OAuth App access): If the subject claim is `azp`, this
	// must be the client ID of the OAuth app registered in Entra ID.
	Subject string `json:"subject"`
	// The claim that contains the subject of the token. Depending on the
	// identity provider and the use case (U2M or M2M), this can vary: - For
	// Entra ID (AAD): * U2M flow (group access): Use `groups`. * U2M flow (user
	// access): Use `oid`. * M2M flow (OAuth App access): Use `azp`. - For other
	// IdPs, refer to the specific IdP documentation.
	//
	// Supported `subject_claim` values are: - `oid`: Object ID of the user. -
	// `azp`: Client ID of the OAuth app. - `groups`: Object ID of the group. -
	// `sub`: Subject identifier for other use cases.
	SubjectClaim string `json:"subject_claim"`
}

func (s *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Partition struct {
	// An array of partition values. A row belongs to the partition when it
	// matches every value in the array, and each `name` must appear at most
	// once in the array.
	Values []PartitionValue `json:"values,omitempty"`
}

func (s *Partition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PartitionValue struct {
	// The name of the partition column.
	Name string `json:"name,omitempty"`
	// The operator to apply for the value. **EQUAL** matches the partition
	// value exactly; **LIKE** matches it against a SQL `LIKE` pattern.
	Op PartitionValueOp `json:"op,omitempty"`
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PartitionValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PartitionValueOp string

const PartitionValueOpEqual PartitionValueOp = `EQUAL`

const PartitionValueOpLike PartitionValueOp = `LIKE`

// String representation for [fmt.Print]
func (f *PartitionValueOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PartitionValueOp) Set(v string) error {
	switch v {
	case `EQUAL`, `LIKE`:
		*f = PartitionValueOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "LIKE"`, v)
	}
}

// Values returns all possible values for PartitionValueOp.
//
// There is no guarantee on the order of the values in the slice.
func (f *PartitionValueOp) Values() []PartitionValueOp {
	return []PartitionValueOp{
		PartitionValueOpEqual,
		PartitionValueOpLike,
	}
}

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (f *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []string `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []string `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Privilege string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable Privilege = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel Privilege = `CREATE_MODEL`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential Privilege = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeCreateVolume Privilege = `CREATE_VOLUME`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeManage Privilege = `MANAGE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeReadVolume Privilege = `READ_VOLUME`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseConnection Privilege = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets Privilege = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume Privilege = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE`, `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FOREIGN_SECURABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `MANAGE`, `MANAGE_ALLOWLIST`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "CREATE", "CREATE_CATALOG", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FOREIGN_SECURABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "MANAGE", "MANAGE_ALLOWLIST", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for Privilege.
//
// There is no guarantee on the order of the values in the slice.
func (f *Privilege) Values() []Privilege {
	return []Privilege{
		PrivilegeAccess,
		PrivilegeAllPrivileges,
		PrivilegeApplyTag,
		PrivilegeCreate,
		PrivilegeCreateCatalog,
		PrivilegeCreateConnection,
		PrivilegeCreateExternalLocation,
		PrivilegeCreateExternalTable,
		PrivilegeCreateExternalVolume,
		PrivilegeCreateForeignCatalog,
		PrivilegeCreateForeignSecurable,
		PrivilegeCreateFunction,
		PrivilegeCreateManagedStorage,
		PrivilegeCreateMaterializedView,
		PrivilegeCreateModel,
		PrivilegeCreateProvider,
		PrivilegeCreateRecipient,
		PrivilegeCreateSchema,
		PrivilegeCreateServiceCredential,
		PrivilegeCreateShare,
		PrivilegeCreateStorageCredential,
		PrivilegeCreateTable,
		PrivilegeCreateView,
		PrivilegeCreateVolume,
		PrivilegeExecute,
		PrivilegeManage,
		PrivilegeManageAllowlist,
		PrivilegeModify,
		PrivilegeReadFiles,
		PrivilegeReadPrivateFiles,
		PrivilegeReadVolume,
		PrivilegeRefresh,
		PrivilegeSelect,
		PrivilegeSetSharePermission,
		PrivilegeUsage,
		PrivilegeUseCatalog,
		PrivilegeUseConnection,
		PrivilegeUseMarketplaceAssets,
		PrivilegeUseProvider,
		PrivilegeUseRecipient,
		PrivilegeUseSchema,
		PrivilegeUseShare,
		PrivilegeWriteFiles,
		PrivilegeWritePrivateFiles,
		PrivilegeWriteVolume,
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProviderInfo struct {
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud string `json:"cloud,omitempty"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Provider.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProviderShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`

	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Cloud string `json:"cloud,omitempty"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// [Create,Update:IGN] common - id of the recipient
	Id string `json:"id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RecipientInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearer_token,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RecipientProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this recipient token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique ID of the recipient token.
	Id string `json:"id,omitempty"`
	// Time at which this recipient token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient token updater.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RecipientTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelAlias struct {
	// Name of the alias.
	AliasName string `json:"alias_name,omitempty"`
	// Numeric model version that alias will reference.
	VersionNum int64 `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" url:"-"`
}

func (s *RetrieveTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearerToken,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RetrieveTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds"`
	// The name of the Recipient.
	Name string `json:"-" url:"-"`
}

func (s *RotateRecipientToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties"`
}

func (s *SecurablePropertiesKvPairs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Share struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Share) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Share) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A share is a named collection of data and AI assets from a single Unity
// Catalog metastore that a provider makes available to recipients through Delta
// Sharing. A share is created empty; tables, views, volumes, models, notebook
// files, or whole schemas are added to it with **update**, and recipients are
// granted access to it through share permissions.
type ShareInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Name of the share.
	Name string `json:"name,omitempty"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `json:"objects,omitempty"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Full storage path, derived from `storage_root`, where the share's
	// notebook files are stored.
	StorageLocation string `json:"storage_location,omitempty"`
	// Cloud storage path, within an external location, under which the share
	// stores the notebook files added to it. When unset, notebook files are
	// stored under the metastore root storage location.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of share updater.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ShareInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// The name of the Recipient.
	Name string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharePermissionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	// The share name.
	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ShareToPrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareToPrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt int64 `json:"added_at,omitempty"`
	// Username of the sharer.
	AddedBy string `json:"added_by,omitempty"`
	// Whether to share the table's change data feed (CDF) with recipients.
	// Requires change data feed to be enabled on the source table, and
	// implicitly enables history sharing. Only applies to tables.
	CdfEnabled bool `json:"cdf_enabled,omitempty"`
	// A user-provided comment when adding the data object to the share.
	Comment string `json:"comment,omitempty"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content string `json:"content,omitempty"`
	// The type of the data object. One of **TABLE**, **VIEW**,
	// **MATERIALIZED_VIEW**, **STREAMING_TABLE**, **FOREIGN_TABLE**,
	// **SCHEMA**, **VOLUME**, **MODEL**, **FUNCTION**, **FEATURE_SPEC**, or
	// **NOTEBOOK_FILE**. Defaults to **TABLE** when omitted.
	DataObjectType SharedDataObjectDataObjectType `json:"data_object_type,omitempty"`
	// Whether to share the table's history, which lets recipients run time
	// travel queries and streaming reads against it. When a **TABLE** is added
	// without this field or `cdf_enabled`, history sharing is **ENABLED** by
	// default; otherwise the default is **DISABLED**.
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `json:"history_data_sharing_status,omitempty"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`, a schema's is `<catalog>.<schema>`, a
	// volume's, model's, or function's is `<catalog>.<schema>.<object_name>`,
	// and a notebook file's is its workspace path.
	Name string `json:"name"`
	// Array of partitions for the shared data. Only rows that belong to at
	// least one of the listed partitions are shared. Only applies to tables.
	Partitions []Partition `json:"partitions,omitempty"`
	// A user-provided alias name for table-like data objects within the share.
	//
	// Use this field for table-like objects (for example: TABLE, VIEW,
	// MATERIALIZED_VIEW, STREAMING_TABLE, FOREIGN_TABLE). For non-table objects
	// (for example: VOLUME, MODEL, NOTEBOOK_FILE, FUNCTION), use
	// `string_shared_as` instead.
	//
	// Important: For non-table objects, this field must be omitted entirely.
	//
	// Format: Must be a 2-part name `<schema_name>.<table_name>` (e.g.,
	// "sales_schema.orders_table") - Both schema and table names must contain
	// only alphanumeric characters and underscores - No periods, spaces,
	// forward slashes, or control characters are allowed within each part - Do
	// not include the catalog name (use 2 parts, not 3)
	//
	// Behavior: - If not provided, the service automatically generates the
	// alias as `<schema>.<table>` from the object's original name - If you
	// don't want to specify this field, omit it entirely from the request (do
	// not pass an empty string) - The `shared_as` name must be unique within
	// the share
	//
	// Examples: - Valid: "analytics_schema.customer_view" - Invalid:
	// "catalog.analytics_schema.customer_view" (3 parts not allowed) - Invalid:
	// "analytics-schema.customer-view" (hyphens not allowed)
	SharedAs string `json:"shared_as,omitempty"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion int64 `json:"start_version,omitempty"`
	// Sharing status of the data object. **ACTIVE** means the object is being
	// shared with recipients without issues; **PERMISSION_DENIED** means the
	// share owner has lost access to the object, so it is not being shared with
	// recipients.
	Status SharedDataObjectStatus `json:"status,omitempty"`
	// A user-provided alias name for non-table data objects within the share.
	//
	// Use this field for non-table objects (for example: VOLUME, MODEL,
	// NOTEBOOK_FILE, FUNCTION).
	//
	// For table-like objects (for example: TABLE, VIEW, MATERIALIZED_VIEW,
	// STREAMING_TABLE, FOREIGN_TABLE), use `shared_as` instead.
	//
	// Important: For table-like objects, this field must be omitted entirely.
	//
	// Format: - For VOLUME: Must be a 2-part name `<schema_name>.<volume_name>`
	// (e.g., "data_schema.ml_models") - For FUNCTION: Must be a 2-part name
	// `<schema_name>.<function_name>` (e.g., "udf_schema.calculate_tax") - For
	// MODEL: Must be a 2-part name `<schema_name>.<model_name>` (e.g.,
	// "models.prediction_model") - For NOTEBOOK_FILE: Should be the notebook
	// file name (e.g., "analysis_notebook.py")
	//
	// - All names must contain only alphanumeric characters and underscores -
	// No periods, spaces, forward slashes, or control characters are allowed
	// within each part
	//
	// Behavior: - If not provided, the service automatically generates the
	// alias from the object's original name - If you don't want to specify this
	// field, omit it entirely from the request (do not pass an empty string) -
	// The `string_shared_as` name must be unique for objects of the same type
	// within the share
	//
	// Examples: - Valid for VOLUME: "data_schema.training_data" - Valid for
	// FUNCTION: "analytics.calculate_revenue"
	//
	// - Invalid: "catalog.data_schema.training_data" (3 parts not allowed for
	// volumes) - Invalid: "data-schema.training-data" (hyphens not allowed)
	StringSharedAs string `json:"string_shared_as,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SharedDataObjectDataObjectType string

const SharedDataObjectDataObjectTypeFeatureSpec SharedDataObjectDataObjectType = `FEATURE_SPEC`

const SharedDataObjectDataObjectTypeForeignTable SharedDataObjectDataObjectType = `FOREIGN_TABLE`

const SharedDataObjectDataObjectTypeFunction SharedDataObjectDataObjectType = `FUNCTION`

const SharedDataObjectDataObjectTypeMaterializedView SharedDataObjectDataObjectType = `MATERIALIZED_VIEW`

const SharedDataObjectDataObjectTypeModel SharedDataObjectDataObjectType = `MODEL`

const SharedDataObjectDataObjectTypeNotebookFile SharedDataObjectDataObjectType = `NOTEBOOK_FILE`

const SharedDataObjectDataObjectTypeSchema SharedDataObjectDataObjectType = `SCHEMA`

const SharedDataObjectDataObjectTypeSkill SharedDataObjectDataObjectType = `SKILL`

const SharedDataObjectDataObjectTypeStreamingTable SharedDataObjectDataObjectType = `STREAMING_TABLE`

const SharedDataObjectDataObjectTypeTable SharedDataObjectDataObjectType = `TABLE`

const SharedDataObjectDataObjectTypeView SharedDataObjectDataObjectType = `VIEW`

const SharedDataObjectDataObjectTypeVolume SharedDataObjectDataObjectType = `VOLUME`

// String representation for [fmt.Print]
func (f *SharedDataObjectDataObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectDataObjectType) Set(v string) error {
	switch v {
	case `FEATURE_SPEC`, `FOREIGN_TABLE`, `FUNCTION`, `MATERIALIZED_VIEW`, `MODEL`, `NOTEBOOK_FILE`, `SCHEMA`, `SKILL`, `STREAMING_TABLE`, `TABLE`, `VIEW`, `VOLUME`:
		*f = SharedDataObjectDataObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FEATURE_SPEC", "FOREIGN_TABLE", "FUNCTION", "MATERIALIZED_VIEW", "MODEL", "NOTEBOOK_FILE", "SCHEMA", "SKILL", "STREAMING_TABLE", "TABLE", "VIEW", "VOLUME"`, v)
	}
}

// Values returns all possible values for SharedDataObjectDataObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectDataObjectType) Values() []SharedDataObjectDataObjectType {
	return []SharedDataObjectDataObjectType{
		SharedDataObjectDataObjectTypeFeatureSpec,
		SharedDataObjectDataObjectTypeForeignTable,
		SharedDataObjectDataObjectTypeFunction,
		SharedDataObjectDataObjectTypeMaterializedView,
		SharedDataObjectDataObjectTypeModel,
		SharedDataObjectDataObjectTypeNotebookFile,
		SharedDataObjectDataObjectTypeSchema,
		SharedDataObjectDataObjectTypeSkill,
		SharedDataObjectDataObjectTypeStreamingTable,
		SharedDataObjectDataObjectTypeTable,
		SharedDataObjectDataObjectTypeView,
		SharedDataObjectDataObjectTypeVolume,
	}
}

// Type always returns SharedDataObjectDataObjectType to satisfy [pflag.Value] interface
func (f *SharedDataObjectDataObjectType) Type() string {
	return "SharedDataObjectDataObjectType"
}

type SharedDataObjectHistoryDataSharingStatus string

const SharedDataObjectHistoryDataSharingStatusDisabled SharedDataObjectHistoryDataSharingStatus = `DISABLED`

const SharedDataObjectHistoryDataSharingStatusEnabled SharedDataObjectHistoryDataSharingStatus = `ENABLED`

// String representation for [fmt.Print]
func (f *SharedDataObjectHistoryDataSharingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectHistoryDataSharingStatus) Set(v string) error {
	switch v {
	case `DISABLED`, `ENABLED`:
		*f = SharedDataObjectHistoryDataSharingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLED", "ENABLED"`, v)
	}
}

// Values returns all possible values for SharedDataObjectHistoryDataSharingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectHistoryDataSharingStatus) Values() []SharedDataObjectHistoryDataSharingStatus {
	return []SharedDataObjectHistoryDataSharingStatus{
		SharedDataObjectHistoryDataSharingStatusDisabled,
		SharedDataObjectHistoryDataSharingStatusEnabled,
	}
}

// Type always returns SharedDataObjectHistoryDataSharingStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectHistoryDataSharingStatus) Type() string {
	return "SharedDataObjectHistoryDataSharingStatus"
}

type SharedDataObjectStatus string

const SharedDataObjectStatusActive SharedDataObjectStatus = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatus = `PERMISSION_DENIED`

// String representation for [fmt.Print]
func (f *SharedDataObjectStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `PERMISSION_DENIED`:
		*f = SharedDataObjectStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PERMISSION_DENIED"`, v)
	}
}

// Values returns all possible values for SharedDataObjectStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectStatus) Values() []SharedDataObjectStatus {
	return []SharedDataObjectStatus{
		SharedDataObjectStatusActive,
		SharedDataObjectStatusPermissionDenied,
	}
}

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectStatus) Type() string {
	return "SharedDataObjectStatus"
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**. **ADD** adds the data object to
	// the share, **REMOVE** removes it, and **UPDATE** changes the properties
	// of a data object that is already in the share.
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being added, removed, or updated.
	DataObject *SharedDataObject `json:"data_object,omitempty"`
}

func (s *SharedDataObjectUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SharedDataObjectUpdateAction string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

const SharedDataObjectUpdateActionUpdate SharedDataObjectUpdateAction = `UPDATE`

// String representation for [fmt.Print]
func (f *SharedDataObjectUpdateAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectUpdateAction) Set(v string) error {
	switch v {
	case `ADD`, `REMOVE`, `UPDATE`:
		*f = SharedDataObjectUpdateAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD", "REMOVE", "UPDATE"`, v)
	}
}

// Values returns all possible values for SharedDataObjectUpdateAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectUpdateAction) Values() []SharedDataObjectUpdateAction {
	return []SharedDataObjectUpdateAction{
		SharedDataObjectUpdateActionAdd,
		SharedDataObjectUpdateActionRemove,
		SharedDataObjectUpdateActionUpdate,
	}
}

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (f *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

// The SecurableKind of a delta-shared object.
type SharedSecurableKind string

const SharedSecurableKindFunctionFeatureSpec SharedSecurableKind = `FUNCTION_FEATURE_SPEC`

const SharedSecurableKindFunctionRegisteredModel SharedSecurableKind = `FUNCTION_REGISTERED_MODEL`

const SharedSecurableKindFunctionStandard SharedSecurableKind = `FUNCTION_STANDARD`

// String representation for [fmt.Print]
func (f *SharedSecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedSecurableKind) Set(v string) error {
	switch v {
	case `FUNCTION_FEATURE_SPEC`, `FUNCTION_REGISTERED_MODEL`, `FUNCTION_STANDARD`:
		*f = SharedSecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FUNCTION_FEATURE_SPEC", "FUNCTION_REGISTERED_MODEL", "FUNCTION_STANDARD"`, v)
	}
}

// Values returns all possible values for SharedSecurableKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedSecurableKind) Values() []SharedSecurableKind {
	return []SharedSecurableKind{
		SharedSecurableKindFunctionFeatureSpec,
		SharedSecurableKindFunctionRegisteredModel,
		SharedSecurableKindFunctionStandard,
	}
}

// Type always returns SharedSecurableKind to satisfy [pflag.Value] interface
func (f *SharedSecurableKind) Type() string {
	return "SharedSecurableKind"
}

type SharesListRequest struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharesListRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharesListRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An agent skill shared through a Delta Sharing share, as served to a
// recipient.
type Skill struct {
	// The name from the skill bundle frontmatter, which may differ from the
	// resource name.
	BundleName string `json:"bundle_name,omitempty"`
	// The comment of the skill.
	Comment string `json:"comment,omitempty"`
	// The bundle description parsed from the skill bundle's SKILL.md
	// frontmatter.
	Description string `json:"description,omitempty"`
	// The shared-object ID that recipients match replicas on. For a skill
	// shared through its schema, which has no shared-object row, this is
	// derived from the shared schema and the skill.
	Id string `json:"id,omitempty"`
	// The name of the skill.
	Name string `json:"name,omitempty"`
	// The name of the schema that the skill belongs to.
	Schema string `json:"schema,omitempty"`
	// The name of the share that the skill belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the skill belongs to.
	ShareId string `json:"share_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Skill) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Skill) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Table struct {
	// The comment of the table.
	Comment string `json:"comment,omitempty"`
	// The id of the table.
	Id string `json:"id,omitempty"`
	// The catalog and schema of the materialized table
	MaterializationNamespace string `json:"materialization_namespace,omitempty"`
	// The name of a materialized table.
	MaterializedTableName string `json:"materialized_table_name,omitempty"`
	// The name of the table.
	Name string `json:"name,omitempty"`
	// The name of the schema that the table belongs to.
	Schema string `json:"schema,omitempty"`
	// The name of the share that the table belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the table belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The Tags of the table.
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Table) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Table) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProvider struct {
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Name of the provider.
	Name string `json:"-" url:"-"`
	// New name for the provider.
	NewName string `json:"new_name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// [Create,Update:IGN] common - id of the recipient
	Id string `json:"id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of the recipient.
	Name string `json:"-" url:"-"`
	// New name for the recipient. .
	NewName string `json:"new_name,omitempty"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
	// New name for the share.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Cloud storage path, within an external location, under which the share
	// stores the notebook files added to it. When unset, notebook files are
	// stored under the metastore root storage location.
	StorageRoot string `json:"storage_root,omitempty"`
	// Array of shared data object updates. The server limits how many updates a
	// single request can contain.
	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSharePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
	// Optional. Whether to return the latest permissions list of the share in
	// the response.
	OmitPermissionsList bool `json:"omit_permissions_list,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateSharePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSharePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func (s *UpdateSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Volume struct {
	// The comment of the volume.
	Comment string `json:"comment,omitempty"`
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	Id string `json:"id,omitempty"`
	// The name of the volume.
	Name string `json:"name,omitempty"`
	// The name of the schema that the volume belongs to.
	Schema string `json:"schema,omitempty"`
	// The name of the share that the volume belongs to.
	Share string `json:"share,omitempty"`
	// / The id of the share that the volume belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The tags of the volume.
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Volume) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Volume) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
