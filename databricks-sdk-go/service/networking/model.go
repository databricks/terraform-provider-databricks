// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networking

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsVpcEndpointInfo struct {
	// The AWS account ID in which this VPC endpoint lives.
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks VPC endpoint service that this endpoint connects
	// to.
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the underlying VPC endpoint in AWS. Provided by the customer
	// when registering an existing AWS VPC endpoint.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzurePrivateEndpointInfo struct {
	// The name of the Private Endpoint in the Azure subscription.
	PrivateEndpointName string `json:"private_endpoint_name"`
	// The GUID of the Private Endpoint resource in the Azure subscription. This
	// is assigned by Azure when the user sets up the Private Endpoint.
	PrivateEndpointResourceGuid string `json:"private_endpoint_resource_guid"`
	// The full resource ID of the Private Endpoint.
	PrivateEndpointResourceId string `json:"private_endpoint_resource_id,omitempty"`
	// The resource ID of the Databricks Private Link Service that this Private
	// Endpoint connects to.
	PrivateLinkServiceId string `json:"private_link_service_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzurePrivateEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzurePrivateEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpointRequest struct {
	Endpoint Endpoint `json:"endpoint"`
	// The parent resource name of the account under which the endpoint is
	// created. Format: `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *CreateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreatePrivateNetworkGatewayRequest struct {
	// The network connectivity configuration that will contain the gateway.
	Parent string `json:"-" url:"-"`
	// The gateway to create.
	PrivateNetworkGateway PrivateNetworkGateway `json:"private_network_gateway"`
	// A unique identifier for this request. The request is idempotent when this
	// is provided.
	RequestId string `json:"-" url:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePrivateNetworkGatewayRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePrivateNetworkGatewayRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteEndpointRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway.
	Name string `json:"-" url:"-"`
}

func (s *DeletePrivateNetworkGatewayRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Endpoint represents a cloud networking resource in a user's cloud account and
// binds it to the Databricks account.
type Endpoint struct {
	// The Databricks Account in which the endpoint object exists.
	AccountId string `json:"account_id,omitempty"`
	// Info for an AWS VPC endpoint.
	AwsVpcEndpointInfo *AwsVpcEndpointInfo `json:"aws_vpc_endpoint_info,omitempty"`
	// Info for an Azure private endpoint.
	AzurePrivateEndpointInfo *AzurePrivateEndpointInfo `json:"azure_private_endpoint_info,omitempty"`
	// The timestamp when the endpoint was created. The timestamp is in RFC 3339
	// format in UTC timezone.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The human-readable display name of this endpoint. The input should
	// conform to RFC-1034, which restricts to letters, numbers, and hyphens,
	// with the first character a letter, the last a letter or a number, and a
	// 63 character maximum.
	DisplayName string `json:"display_name"`
	// The unique identifier for this endpoint under the account. This field is
	// a UUID generated by Databricks.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Info for a GCP Private Service Connect endpoint.
	GcpPscEndpointInfo *GcpPscEndpointInfo `json:"gcp_psc_endpoint_info,omitempty"`
	// The resource name of the endpoint, which uniquely identifies the
	// endpoint.
	Name string `json:"name,omitempty"`
	// The cloud provider region where this endpoint is located.
	Region string `json:"region"`
	// The state of the endpoint. The endpoint can only be used if the state is
	// `APPROVED`.
	State EndpointState `json:"state,omitempty"`
	// The use case that determines the type of network connectivity this
	// endpoint provides. This field is automatically determined based on the
	// endpoint configuration and cloud-specific settings.
	UseCase EndpointUseCase `json:"use_case,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Endpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Endpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointState string

const EndpointStateApproved EndpointState = `APPROVED`

const EndpointStateDisconnected EndpointState = `DISCONNECTED`

const EndpointStateFailed EndpointState = `FAILED`

const EndpointStatePending EndpointState = `PENDING`

// String representation for [fmt.Print]
func (f *EndpointState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointState) Set(v string) error {
	switch v {
	case `APPROVED`, `DISCONNECTED`, `FAILED`, `PENDING`:
		*f = EndpointState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVED", "DISCONNECTED", "FAILED", "PENDING"`, v)
	}
}

// Values returns all possible values for EndpointState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointState) Values() []EndpointState {
	return []EndpointState{
		EndpointStateApproved,
		EndpointStateDisconnected,
		EndpointStateFailed,
		EndpointStatePending,
	}
}

// Type always returns EndpointState to satisfy [pflag.Value] interface
func (f *EndpointState) Type() string {
	return "EndpointState"
}

type EndpointUseCase string

const EndpointUseCaseServiceDirect EndpointUseCase = `SERVICE_DIRECT`

// String representation for [fmt.Print]
func (f *EndpointUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointUseCase) Set(v string) error {
	switch v {
	case `SERVICE_DIRECT`:
		*f = EndpointUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SERVICE_DIRECT"`, v)
	}
}

// Values returns all possible values for EndpointUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointUseCase) Values() []EndpointUseCase {
	return []EndpointUseCase{
		EndpointUseCaseServiceDirect,
	}
}

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

type GcpPscEndpointInfo struct {
	// The GCP region of the PSC connection endpoint. Provided by the customer
	// when registering an existing PSC endpoint. GCP supports only same-region
	// PSC, so this must match the workspace region.
	EndpointRegion string `json:"endpoint_region"`
	// The GCP consumer project ID in which this PSC endpoint is created.
	// Provided by the customer when registering an existing PSC endpoint.
	ProjectId string `json:"project_id"`
	// The ID of the underlying Private Service Connect connection in the GCP
	// consumer project, assigned by GCP when the PSC connection is created.
	PscConnectionId string `json:"psc_connection_id,omitempty"`
	// The name of this PSC connection in the GCP consumer project. Provided by
	// the customer when registering an existing PSC endpoint.
	PscEndpoint string `json:"psc_endpoint"`
	// The ID of the Databricks service attachment this PSC endpoint connects
	// to.
	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpPscEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpPscEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEndpointRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway.
	Name string `json:"-" url:"-"`
}

func (s *GetPrivateNetworkGatewayRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListEndpointsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent resource name of the account to list endpoints for. Format:
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	Items []Endpoint `json:"items,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPrivateNetworkGatewaysRequest struct {
	// An opaque token returned by a previous list request.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The network connectivity configuration containing the gateways.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivateNetworkGatewaysRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateNetworkGatewaysRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPrivateNetworkGatewaysResponse struct {
	// An opaque token for the next page, or empty when there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`

	PrivateNetworkGateways []PrivateNetworkGateway `json:"private_network_gateways,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivateNetworkGatewaysResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateNetworkGatewaysResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A private network gateway connects serverless compute to destinations in a
// customer-managed VPC or VNet.
type PrivateNetworkGateway struct {
	// The AWS connection used by the gateway.
	AwsCloudConnection *PrivateNetworkGatewayAwsCloudConnection `json:"aws_cloud_connection,omitempty"`
	// The Azure connection used by the gateway.
	AzureCloudConnection *PrivateNetworkGatewayAzureCloudConnection `json:"azure_cloud_connection,omitempty"`
	// The provisioned bandwidth tier for an Azure gateway, in gigabits per
	// second. Required when creating an Azure gateway.
	BandwidthTierGigabitsPerSecond int `json:"bandwidth_tier_gigabits_per_second,omitempty"`
	// The time when the gateway was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The destinations routed through this gateway.
	Destinations []PrivateNetworkGatewayDestination `json:"destinations,omitempty"`
	// The human-readable name of the gateway.
	DisplayName string `json:"display_name"`
	// The failure reason when the gateway is in the FAILED state.
	ErrorMessage string `json:"error_message,omitempty"`
	// The canonical resource name of the gateway, in the form
	// `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`.
	Name string `json:"name,omitempty"`
	// The DNS resolvers used for private name resolution.
	PrivateDnsResolvers []PrivateNetworkGatewayPrivateDnsResolver `json:"private_dns_resolvers,omitempty"`
	// The current lifecycle state of the gateway.
	State PrivateNetworkGatewayGatewayState `json:"state,omitempty"`
	// The traffic routed through this gateway.
	TrafficMode PrivateNetworkGatewayTrafficMode `json:"traffic_mode"`
	// The time when the gateway was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivateNetworkGateway) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateNetworkGateway) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// AWS connection configuration.
type PrivateNetworkGatewayAwsCloudConnection struct {
	// The IAM role that Databricks assumes to manage gateway resources.
	CrossAccountRole PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole `json:"cross_account_role"`
	// The subnets where the gateway establishes connectivity.
	GatewaySubnets []PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet `json:"gateway_subnets"`
	// The security groups attached to the gateway network interface.
	SecurityGroupIds []string `json:"security_group_ids"`
}

func (s *PrivateNetworkGatewayAwsCloudConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An AWS subnet used by the gateway.
type PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet struct {
	// The AWS subnet ID.
	SubnetId string `json:"subnet_id"`
}

func (s *PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A cross-account IAM role used to manage gateway resources.
type PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole struct {
	// The ARN of the IAM role.
	RoleArn string `json:"role_arn"`
}

func (s *PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Azure connection configuration.
type PrivateNetworkGatewayAzureCloudConnection struct {
	// The subnet where the gateway establishes connectivity.
	GatewaySubnet PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet `json:"gateway_subnet"`
}

func (s *PrivateNetworkGatewayAzureCloudConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An Azure subnet used by the gateway.
type PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet struct {
	// The full Azure resource ID of the subnet.
	ResourceId string `json:"resource_id"`
}

func (s *PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A destination routed through the gateway.
type PrivateNetworkGatewayDestination struct {
	// The destination type.
	DestinationType PrivateNetworkGatewayDestinationDestinationType `json:"destination_type"`
	// The destination value.
	Value string `json:"value"`
}

func (s *PrivateNetworkGatewayDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The supported destination types.
type PrivateNetworkGatewayDestinationDestinationType string

const PrivateNetworkGatewayDestinationDestinationTypeDnsName PrivateNetworkGatewayDestinationDestinationType = `DNS_NAME`

// String representation for [fmt.Print]
func (f *PrivateNetworkGatewayDestinationDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateNetworkGatewayDestinationDestinationType) Set(v string) error {
	switch v {
	case `DNS_NAME`:
		*f = PrivateNetworkGatewayDestinationDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DNS_NAME"`, v)
	}
}

// Values returns all possible values for PrivateNetworkGatewayDestinationDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateNetworkGatewayDestinationDestinationType) Values() []PrivateNetworkGatewayDestinationDestinationType {
	return []PrivateNetworkGatewayDestinationDestinationType{
		PrivateNetworkGatewayDestinationDestinationTypeDnsName,
	}
}

// Type always returns PrivateNetworkGatewayDestinationDestinationType to satisfy [pflag.Value] interface
func (f *PrivateNetworkGatewayDestinationDestinationType) Type() string {
	return "PrivateNetworkGatewayDestinationDestinationType"
}

// The lifecycle state of the gateway.
type PrivateNetworkGatewayGatewayState string

const PrivateNetworkGatewayGatewayStateCreating PrivateNetworkGatewayGatewayState = `CREATING`

const PrivateNetworkGatewayGatewayStateDeleting PrivateNetworkGatewayGatewayState = `DELETING`

const PrivateNetworkGatewayGatewayStateEstablished PrivateNetworkGatewayGatewayState = `ESTABLISHED`

const PrivateNetworkGatewayGatewayStateFailed PrivateNetworkGatewayGatewayState = `FAILED`

// String representation for [fmt.Print]
func (f *PrivateNetworkGatewayGatewayState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateNetworkGatewayGatewayState) Set(v string) error {
	switch v {
	case `CREATING`, `DELETING`, `ESTABLISHED`, `FAILED`:
		*f = PrivateNetworkGatewayGatewayState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATING", "DELETING", "ESTABLISHED", "FAILED"`, v)
	}
}

// Values returns all possible values for PrivateNetworkGatewayGatewayState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateNetworkGatewayGatewayState) Values() []PrivateNetworkGatewayGatewayState {
	return []PrivateNetworkGatewayGatewayState{
		PrivateNetworkGatewayGatewayStateCreating,
		PrivateNetworkGatewayGatewayStateDeleting,
		PrivateNetworkGatewayGatewayStateEstablished,
		PrivateNetworkGatewayGatewayStateFailed,
	}
}

// Type always returns PrivateNetworkGatewayGatewayState to satisfy [pflag.Value] interface
func (f *PrivateNetworkGatewayGatewayState) Type() string {
	return "PrivateNetworkGatewayGatewayState"
}

type PrivateNetworkGatewayOperationMetadata struct {
	// The operation performed on the gateway.
	OperationType PrivateNetworkGatewayOperationMetadataOperationType `json:"operation_type,omitempty"`
}

func (s *PrivateNetworkGatewayOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PrivateNetworkGatewayOperationMetadataOperationType string

const PrivateNetworkGatewayOperationMetadataOperationTypeCreate PrivateNetworkGatewayOperationMetadataOperationType = `CREATE`

// String representation for [fmt.Print]
func (f *PrivateNetworkGatewayOperationMetadataOperationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateNetworkGatewayOperationMetadataOperationType) Set(v string) error {
	switch v {
	case `CREATE`:
		*f = PrivateNetworkGatewayOperationMetadataOperationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATE"`, v)
	}
}

// Values returns all possible values for PrivateNetworkGatewayOperationMetadataOperationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateNetworkGatewayOperationMetadataOperationType) Values() []PrivateNetworkGatewayOperationMetadataOperationType {
	return []PrivateNetworkGatewayOperationMetadataOperationType{
		PrivateNetworkGatewayOperationMetadataOperationTypeCreate,
	}
}

// Type always returns PrivateNetworkGatewayOperationMetadataOperationType to satisfy [pflag.Value] interface
func (f *PrivateNetworkGatewayOperationMetadataOperationType) Type() string {
	return "PrivateNetworkGatewayOperationMetadataOperationType"
}

// A private DNS resolver used by the gateway.
type PrivateNetworkGatewayPrivateDnsResolver struct {
	// The resolver type.
	ResolverType PrivateNetworkGatewayPrivateDnsResolverResolverType `json:"resolver_type"`
	// The resolver value.
	Value string `json:"value"`
}

func (s *PrivateNetworkGatewayPrivateDnsResolver) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The supported resolver types.
type PrivateNetworkGatewayPrivateDnsResolverResolverType string

const PrivateNetworkGatewayPrivateDnsResolverResolverTypeIpAddress PrivateNetworkGatewayPrivateDnsResolverResolverType = `IP_ADDRESS`

// String representation for [fmt.Print]
func (f *PrivateNetworkGatewayPrivateDnsResolverResolverType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateNetworkGatewayPrivateDnsResolverResolverType) Set(v string) error {
	switch v {
	case `IP_ADDRESS`:
		*f = PrivateNetworkGatewayPrivateDnsResolverResolverType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IP_ADDRESS"`, v)
	}
}

// Values returns all possible values for PrivateNetworkGatewayPrivateDnsResolverResolverType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateNetworkGatewayPrivateDnsResolverResolverType) Values() []PrivateNetworkGatewayPrivateDnsResolverResolverType {
	return []PrivateNetworkGatewayPrivateDnsResolverResolverType{
		PrivateNetworkGatewayPrivateDnsResolverResolverTypeIpAddress,
	}
}

// Type always returns PrivateNetworkGatewayPrivateDnsResolverResolverType to satisfy [pflag.Value] interface
func (f *PrivateNetworkGatewayPrivateDnsResolverResolverType) Type() string {
	return "PrivateNetworkGatewayPrivateDnsResolverResolverType"
}

// The traffic routed through the gateway.
type PrivateNetworkGatewayTrafficMode string

const PrivateNetworkGatewayTrafficModeAllTraffic PrivateNetworkGatewayTrafficMode = `ALL_TRAFFIC`

const PrivateNetworkGatewayTrafficModeSpecificDestinations PrivateNetworkGatewayTrafficMode = `SPECIFIC_DESTINATIONS`

// String representation for [fmt.Print]
func (f *PrivateNetworkGatewayTrafficMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateNetworkGatewayTrafficMode) Set(v string) error {
	switch v {
	case `ALL_TRAFFIC`, `SPECIFIC_DESTINATIONS`:
		*f = PrivateNetworkGatewayTrafficMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_TRAFFIC", "SPECIFIC_DESTINATIONS"`, v)
	}
}

// Values returns all possible values for PrivateNetworkGatewayTrafficMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateNetworkGatewayTrafficMode) Values() []PrivateNetworkGatewayTrafficMode {
	return []PrivateNetworkGatewayTrafficMode{
		PrivateNetworkGatewayTrafficModeAllTraffic,
		PrivateNetworkGatewayTrafficModeSpecificDestinations,
	}
}

// Type always returns PrivateNetworkGatewayTrafficMode to satisfy [pflag.Value] interface
func (f *PrivateNetworkGatewayTrafficMode) Type() string {
	return "PrivateNetworkGatewayTrafficMode"
}

type UpdatePrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway, in the form
	// `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`.
	Name string `json:"-" url:"-"`
	// The gateway containing the desired mutable field values.
	PrivateNetworkGateway PrivateNetworkGateway `json:"private_network_gateway"`
	// The fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdatePrivateNetworkGatewayRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
