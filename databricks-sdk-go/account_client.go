// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/disasterrecovery"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/databricks-sdk-go/service/networking"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
)

type AccountClient struct {
	Config *config.Config

	// These APIs manage access rules on resources in an account. Currently,
	// only grant rules are supported. A grant rule specifies a role assigned to
	// a set of principals. A list of rules attached to a resource is called a
	// rule set.
	AccessControl iam.AccountAccessControlInterface

	// This API allows you to download billable usage logs for the specified
	// account and date range. This feature works with all account types.
	BillableUsage billing.BillableUsageInterface

	// A service serves REST API about Budget policies
	BudgetPolicy billing.BudgetPolicyInterface

	// These APIs manage budget configurations for this account. Budgets enable
	// you to monitor usage across your account. You can set up budgets to
	// either track account-wide spending, or apply filters to track the
	// spending of specific teams, projects, or workspaces.
	Budgets billing.BudgetsInterface

	// These APIs manage credential configurations for this workspace.
	// Databricks needs access to a cross-account service IAM role in your AWS
	// account so that Databricks can deploy clusters in the appropriate VPC for
	// the new workspace. A credential configuration encapsulates this role
	// information, and its ID is used when creating a new workspace.
	Credentials provisioning.CredentialsInterface

	// These APIs enable administrators to manage custom OAuth app integrations,
	// which is required for adding/using Custom OAuth App Integration like
	// Tableau Cloud for Databricks in AWS cloud.
	CustomAppIntegration oauth2.CustomAppIntegrationInterface

	// Manage disaster recovery configurations and execute failover operations.
	DisasterRecovery disasterrecovery.DisasterRecoveryInterface

	// These APIs manage encryption key configurations for this workspace
	// (optional). A key configuration encapsulates the AWS KMS key information
	// and some information about how the key configuration can be used. There
	// are two possible uses for key configurations:
	//
	// * Managed services: A key configuration can be used to encrypt a
	// workspace's notebook and secret data in the control plane, as well as
	// Databricks SQL queries and query history. * Storage: A key configuration
	// can be used to encrypt a workspace's DBFS and EBS data in the data plane.
	//
	// In both of these cases, the key configuration's ID is used when creating
	// a new workspace. This Preview feature is available if your account is on
	// the E2 version of the platform. Updating a running workspace with
	// workspace storage encryption requires that the workspace is on the E2
	// version of the platform. If you have an older workspace, it might not be
	// on the E2 version of the platform. If you are not sure, contact your
	// Databricks representative.
	EncryptionKeys provisioning.EncryptionKeysInterface

	// These APIs manage endpoint configurations for this account.
	Endpoints networking.EndpointsInterface

	// These APIs manage account federation policies.
	//
	// Account federation policies allow users and service principals in your
	// Databricks account to securely access Databricks APIs using tokens from
	// your trusted identity providers (IdPs).
	//
	// With token federation, your users and service principals can exchange
	// tokens from your IdP for Databricks OAuth tokens, which can be used to
	// access Databricks APIs. Token federation eliminates the need to manage
	// Databricks secrets, and allows you to centralize management of token
	// issuance policies in your IdP. Databricks token federation is typically
	// used in combination with [SCIM], so users in your IdP are synchronized
	// into your Databricks account.
	//
	// Token federation is configured in your Databricks account using an
	// account federation policy. An account federation policy specifies: *
	// which IdP, or issuer, your Databricks account should accept tokens from *
	// how to determine which Databricks user, or subject, a token is issued for
	//
	// To configure a federation policy, you provide the following: * The
	// required token __issuer__, as specified in the “iss” claim of your
	// tokens. The issuer is an https URL that identifies your IdP. * The
	// allowed token __audiences__, as specified in the “aud” claim of your
	// tokens. This identifier is intended to represent the recipient of the
	// token. As long as the audience in the token matches at least one audience
	// in the policy, the token is considered a match. If unspecified, the
	// default value is your Databricks account id. * The __subject claim__,
	// which indicates which token claim contains the Databricks username of the
	// user the token was issued for. If unspecified, the default value is
	// “sub”. * Optionally, the public keys used to validate the signature
	// of your tokens, in JWKS format. If unspecified (recommended), Databricks
	// automatically fetches the public keys from your issuer’s well known
	// endpoint. Databricks strongly recommends relying on your issuer’s well
	// known endpoint for discovering public keys.
	//
	// An example federation policy is: ``` issuer:
	// "https://idp.mycompany.com/oidc" audiences: ["databricks"] subject_claim:
	// "sub" ```
	//
	// An example JWT token body that matches this policy and could be used to
	// authenticate to Databricks as user `username@mycompany.com` is: ``` {
	// "iss": "https://idp.mycompany.com/oidc", "aud": "databricks", "sub":
	// "username@mycompany.com" } ```
	//
	// You may also need to configure your IdP to generate tokens for your users
	// to exchange with Databricks, if your users do not already have the
	// ability to generate tokens that are compatible with your federation
	// policy.
	//
	// You do not need to configure an OAuth application in Databricks to use
	// token federation.
	//
	// [SCIM]: https://docs.databricks.com/admin/users-groups/scim/index.html
	FederationPolicy oauth2.AccountFederationPolicyInterface

	// Groups simplify identity management, making it easier to assign access to
	// Databricks account, data, and other securable objects.
	//
	// It is best practice to assign access to workspaces and access-control
	// policies in Unity Catalog to groups, instead of to users individually.
	// All Databricks account identities can be assigned as members of groups,
	// and members inherit permissions that are assigned to their group.
	GroupsV2 iam.AccountGroupsV2Interface

	// These APIs are used to manage identities and the workspace access of
	// these identities in Databricks.
	IamV2 iamv2.AccountIamV2Interface

	// The Accounts IP Access List API enables account admins to configure IP
	// access lists for access to the account console.
	//
	// Account IP Access Lists affect web application access and REST API access
	// to the account console and account APIs. If the feature is disabled for
	// the account, all access is allowed for this account. There is support for
	// allow lists (inclusion) and block lists (exclusion).
	//
	// When a connection is attempted: 1. **First, all block lists are
	// checked.** If the connection IP address matches any block list, the
	// connection is rejected. 2. **If the connection was not rejected by block
	// lists**, the IP address is compared with the allow lists.
	//
	// If there is at least one allow list for the account, the connection is
	// allowed only if the IP address matches an allow list. If there are no
	// allow lists for the account, all IP addresses are allowed.
	//
	// For all allow lists and block lists combined, the account supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	//
	// After changes to the account-level IP access lists, it can take a few
	// minutes for changes to take effect.
	IpAccessLists settings.AccountIpAccessListsInterface

	// These APIs manage log delivery configurations for this account. The two
	// supported log types for this API are _billable usage logs_ (AWS only) and
	// _audit logs_ (AWS and GCP). This feature is in Public Preview. This
	// feature works with all account ID types.
	//
	// Log delivery works with all account types. However, if your account is on
	// the E2 version of the platform or on a select custom plan that allows
	// multiple workspaces per account, you can optionally configure different
	// storage destinations for each workspace. Log delivery status is also
	// provided to know the latest status of log delivery attempts.
	//
	// The high-level flow of billable usage delivery (AWS only):
	//
	// 1. **Create storage**: In AWS, [create a new AWS S3 bucket] with a
	// specific bucket policy. Using Databricks APIs, call the Account API to
	// create a [storage configuration object](:method:Storage/Create) that uses
	// the bucket name.
	//
	// 2. **Create credentials**: In AWS, create the appropriate AWS IAM role.
	// For full details, including the required IAM role policies and trust
	// relationship, see [Billable usage log delivery]. Using Databricks APIs,
	// call the Account API to create a [credential configuration
	// object](:method:Credentials/Create) that uses the IAM role's ARN.
	//
	// 3. **Create log delivery configuration**: Using Databricks APIs, call the
	// Account API to [create a log delivery
	// configuration](:method:LogDelivery/Create) that uses the credential and
	// storage configuration objects from previous steps. You can specify if the
	// logs should include all events of that log type in your account (_Account
	// level_ delivery) or only events for a specific set of workspaces
	// (_workspace level_ delivery). Account level log delivery applies to all
	// current and future workspaces plus account level logs, while workspace
	// level log delivery solely delivers logs related to the specified
	// workspaces. You can create multiple types of delivery configurations per
	// account.
	//
	// For billable usage delivery (AWS only): * For more information about
	// billable usage logs, see [Billable usage log delivery]. For the CSV
	// schema, see the [Usage page]. * The delivery location is
	// `<bucket-name>/<prefix>/billable-usage/csv/`, where `<prefix>` is the
	// name of the optional delivery path prefix you set up during log delivery
	// configuration. Files are named
	// `workspaceId=<workspace-id>-usageMonth=<month>.csv`. * All billable usage
	// logs apply to specific workspaces (_workspace level_ logs). You can
	// aggregate usage for your entire account by creating an _account level_
	// delivery configuration that delivers logs for all current and future
	// workspaces in your account. * The files are delivered daily by
	// overwriting the month's CSV file for each workspace.
	//
	// For audit log delivery (AWS and GCP): * For more information about about
	// audit log delivery, see Audit log delivery [AWS] or [GCP], which includes
	// information about the used JSON schema. * The delivery location is
	// `<bucket-name>/<delivery-path-prefix>/workspaceId=<workspaceId>/date=<yyyy-mm-dd>/auditlogs_<internal-id>.json`.
	// Files may get overwritten with the same content multiple times to achieve
	// exactly-once delivery. * If the audit log delivery configuration included
	// specific workspace IDs, only _workspace-level_ audit logs for those
	// workspaces are delivered. If the log delivery configuration applies to
	// the entire account (_account level_ delivery configuration), the audit
	// log delivery includes workspace-level audit logs for all workspaces in
	// the account as well as account-level audit logs. See Audit log delivery
	// [AWS] or [GCP] for details. * Auditable events are typically available in
	// logs within 15 minutes.
	//
	// [AWS]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [Billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// [GCP]: https://docs.databricks.com/gcp/en/admin/account-settings/audit-logs
	// [Usage page]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [create a new AWS S3 bucket]: https://docs.databricks.com/administration-guide/account-api/aws-storage.html
	LogDelivery billing.LogDeliveryInterface

	// These APIs manage metastore assignments to a workspace.
	MetastoreAssignments catalog.AccountMetastoreAssignmentsInterface

	// These APIs manage Unity Catalog metastores for an account. A metastore
	// contains catalogs that can be associated with workspaces
	Metastores catalog.AccountMetastoresInterface

	// These APIs provide configurations for the network connectivity of your
	// workspaces for serverless compute resources. This API provides stable
	// subnets for your workspace so that you can configure your firewalls on
	// your Azure Storage accounts to allow access from Databricks. You can also
	// use the API to provision private endpoints for Databricks to privately
	// connect serverless compute resources to your Azure resources using Azure
	// Private Link. See [configure serverless secure connectivity].
	//
	// [configure serverless secure connectivity]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security
	NetworkConnectivity settings.NetworkConnectivityInterface

	// These APIs manage network policies for this account. Network policies
	// control which network destinations can be accessed from the Databricks
	// environment. Each Databricks account includes a default policy named
	// 'default-policy'. 'default-policy' is associated with any workspace
	// lacking an explicit network policy assignment, and is automatically
	// associated with each newly created workspace. 'default-policy' is
	// reserved and cannot be deleted, but it can be updated to customize the
	// default network access rules for your account.
	NetworkPolicies settings.NetworkPoliciesInterface

	// These APIs manage network configurations for customer-managed VPCs
	// (optional). Its ID is used when creating a new workspace if you use
	// customer-managed VPCs.
	Networks provisioning.NetworksInterface

	// These APIs enable administrators to view all the available published
	// OAuth applications in Databricks. Administrators can add the published
	// OAuth applications to their account through the OAuth Published App
	// Integration APIs.
	OAuthPublishedApps oauth2.OAuthPublishedAppsInterface

	// These APIs manage private access settings for this account.
	PrivateAccess provisioning.PrivateAccessInterface

	// These APIs manage private network gateways under network connectivity
	// configurations.
	PrivateNetworkGateways networking.PrivateNetworkGatewaysInterface

	// These APIs enable administrators to manage published OAuth app
	// integrations, which is required for adding/using Published OAuth App
	// Integration like Tableau Desktop for Databricks in AWS cloud.
	PublishedAppIntegration oauth2.PublishedAppIntegrationInterface

	// These APIs manage service principal federation policies.
	//
	// Service principal federation, also known as Workload Identity Federation,
	// allows your automated workloads running outside of Databricks to securely
	// access Databricks APIs without the need for Databricks secrets. With
	// Workload Identity Federation, your application (or workload)
	// authenticates to Databricks as a Databricks service principal, using
	// tokens provided by the workload runtime.
	//
	// Databricks strongly recommends using Workload Identity Federation to
	// authenticate to Databricks from automated workloads, over alternatives
	// such as OAuth client secrets or Personal Access Tokens, whenever
	// possible. Workload Identity Federation is supported by many popular
	// services, including Github Actions, Azure DevOps, GitLab, Terraform
	// Cloud, and Kubernetes clusters, among others.
	//
	// Workload identity federation is configured in your Databricks account
	// using a service principal federation policy. A service principal
	// federation policy specifies: * which IdP, or issuer, the service
	// principal is allowed to authenticate from * which workload identity, or
	// subject, is allowed to authenticate as the Databricks service principal
	//
	// To configure a federation policy, you provide the following: * The
	// required token __issuer__, as specified in the “iss” claim of
	// workload identity tokens. The issuer is an https URL that identifies the
	// workload identity provider. * The required token __subject__, as
	// specified in the “sub” claim of workload identity tokens. The subject
	// uniquely identifies the workload in the workload runtime environment. *
	// The allowed token __audiences__, as specified in the “aud” claim of
	// workload identity tokens. The audience is intended to represent the
	// recipient of the token. As long as the audience in the token matches at
	// least one audience in the policy, the token is considered a match. If
	// unspecified, the default value is your Databricks account id. *
	// Optionally, the public keys used to validate the signature of the
	// workload identity tokens, in JWKS format. If unspecified (recommended),
	// Databricks automatically fetches the public keys from the issuer’s well
	// known endpoint. Databricks strongly recommends relying on the issuer’s
	// well known endpoint for discovering public keys.
	//
	// An example service principal federation policy, for a Github Actions
	// workload, is: ``` issuer: "https://token.actions.githubusercontent.com"
	// audiences: ["https://github.com/my-github-org"] subject:
	// "repo:my-github-org/my-repo:environment:prod" ```
	//
	// An example JWT token body that matches this policy and could be used to
	// authenticate to Databricks is: ``` { "iss":
	// "https://token.actions.githubusercontent.com", "aud":
	// "https://github.com/my-github-org", "sub":
	// "repo:my-github-org/my-repo:environment:prod" } ```
	//
	// You may also need to configure the workload runtime to generate tokens
	// for your workloads.
	//
	// You do not need to configure an OAuth application in Databricks to use
	// token federation.
	ServicePrincipalFederationPolicy oauth2.ServicePrincipalFederationPolicyInterface

	// These APIs enable administrators to manage service principal secrets.
	//
	// You can use the generated secrets to obtain OAuth access tokens for a
	// service principal, which can then be used to access Databricks Accounts
	// and Workspace APIs. For more information, see [Authentication using OAuth
	// tokens for service principals].
	//
	// In addition, the generated secrets can be used to configure the
	// Databricks Terraform Provider to authenticate with the service principal.
	// For more information, see [Databricks Terraform Provider].
	//
	// [Authentication using OAuth tokens for service principals]: https://docs.databricks.com/dev-tools/authentication-oauth.html
	// [Databricks Terraform Provider]: https://github.com/databricks/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-service-principal
	ServicePrincipalSecrets oauth2.ServicePrincipalSecretsInterface

	// Identities for use with jobs, automated tools, and systems such as
	// scripts, apps, and CI/CD platforms. Databricks recommends creating
	// service principals to run production jobs or modify production data. If
	// all processes that act on production data run with service principals,
	// interactive users do not need any write, delete, or modify privileges in
	// production. This eliminates the risk of a user overwriting production
	// data by accident.
	ServicePrincipalsV2 iam.AccountServicePrincipalsV2Interface

	// Accounts Settings API allows users to manage settings at the account
	// level.
	Settings settings.AccountSettingsInterface

	// APIs to manage account level settings
	SettingsV2 settingsv2.AccountSettingsV2Interface

	// These APIs manage storage configurations for this workspace. A root
	// storage S3 bucket in your account is required to store objects like
	// cluster logs, notebook revisions, and job results. You can also use the
	// root storage S3 bucket for storage of non-production DBFS data. A storage
	// configuration encapsulates this bucket information, and its ID is used
	// when creating a new workspace.
	Storage provisioning.StorageInterface

	// These APIs manage storage credentials for a particular metastore.
	StorageCredentials catalog.AccountStorageCredentialsInterface

	// These APIs manage usage dashboards for this account. Usage dashboards
	// enable you to gain insights into your usage with pre-built dashboards:
	// visualize breakdowns, analyze tag attributions, and identify cost
	// drivers.
	UsageDashboards billing.UsageDashboardsInterface

	// User identities recognized by Databricks and represented by email
	// addresses.
	//
	// Databricks recommends using SCIM provisioning to sync users and groups
	// automatically from your identity provider to your Databricks account.
	// SCIM streamlines onboarding a new employee or team by using your identity
	// provider to create users and groups in Databricks account and give them
	// the proper level of access. When a user leaves your organization or no
	// longer needs access to Databricks account, admins can terminate the user
	// in your identity provider and that user’s account will also be removed
	// from Databricks account. This ensures a consistent offboarding process
	// and prevents unauthorized users from accessing sensitive data.
	UsersV2 iam.AccountUsersV2Interface

	// These APIs manage VPC endpoint configurations for this account.
	VpcEndpoints provisioning.VpcEndpointsInterface

	// The Workspace Permission Assignment API allows you to manage workspace
	// permissions for principals in your account.
	WorkspaceAssignment iam.WorkspaceAssignmentInterface

	// These APIs allow configuration of network settings for Databricks
	// workspaces by selecting which network policy to associate with the
	// workspace. Each workspace is always associated with exactly one network
	// policy that controls which network destinations can be accessed from the
	// Databricks environment. By default, workspaces are associated with the
	// 'default-policy' network policy. You cannot create or delete a
	// workspace's network option, only update it to associate the workspace
	// with a different policy
	WorkspaceNetworkConfiguration settings.WorkspaceNetworkConfigurationInterface

	// These APIs manage workspaces for this account. A Databricks workspace is
	// an environment for accessing all of your Databricks assets. The workspace
	// organizes objects (notebooks, libraries, and experiments) into folders,
	// and provides access to data and computational resources such as clusters
	// and jobs.
	//
	// These endpoints are available if your account is on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	Workspaces provisioning.WorkspacesInterface

	// Groups simplify identity management, making it easier to assign access to
	// Databricks account, data, and other securable objects.
	//
	// It is best practice to assign access to workspaces and access-control
	// policies in Unity Catalog to groups, instead of to users individually.
	// All Databricks account identities can be assigned as members of groups,
	// and members inherit permissions that are assigned to their group.
	//
	// Deprecated: Use the GroupsV2 API instead.
	Groups iam.AccountGroupsInterface

	// Identities for use with jobs, automated tools, and systems such as
	// scripts, apps, and CI/CD platforms. Databricks recommends creating
	// service principals to run production jobs or modify production data. If
	// all processes that act on production data run with service principals,
	// interactive users do not need any write, delete, or modify privileges in
	// production. This eliminates the risk of a user overwriting production
	// data by accident.
	//
	// Deprecated: Use the ServicePrincipalsV2 API instead.
	ServicePrincipals iam.AccountServicePrincipalsInterface

	// User identities recognized by Databricks and represented by email
	// addresses.
	//
	// Databricks recommends using SCIM provisioning to sync users and groups
	// automatically from your identity provider to your Databricks account.
	// SCIM streamlines onboarding a new employee or team by using your identity
	// provider to create users and groups in Databricks account and give them
	// the proper level of access. When a user leaves your organization or no
	// longer needs access to Databricks account, admins can terminate the user
	// in your identity provider and that user’s account will also be removed
	// from Databricks account. This ensures a consistent offboarding process
	// and prevents unauthorized users from accessing sensitive data.
	//
	// Deprecated: Use the UsersV2 API instead.
	Users iam.AccountUsersInterface
}

var ErrNotAccountClient = errors.New("invalid Databricks Account configuration - host incorrect or account_id missing")
var ErrWorkspaceIDInAccountClient = errors.New("WorkspaceID must not be set when using AccountClient")

// NewAccountClient creates new Databricks SDK client for Accounts or returns
// error in case configuration is wrong
func NewAccountClient(c ...*Config) (*AccountClient, error) {
	var cfg *config.Config
	if len(c) == 1 {
		// first config
		cfg = (*config.Config)(c[0])
	} else {
		// default config
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountClient{
		Config: cfg,

		AccessControl:                    iam.NewAccountAccessControl(apiClient),
		BillableUsage:                    billing.NewBillableUsage(apiClient),
		BudgetPolicy:                     billing.NewBudgetPolicy(apiClient),
		Budgets:                          billing.NewBudgets(apiClient),
		Credentials:                      provisioning.NewCredentials(apiClient),
		CustomAppIntegration:             oauth2.NewCustomAppIntegration(apiClient),
		DisasterRecovery:                 disasterrecovery.NewDisasterRecovery(apiClient),
		EncryptionKeys:                   provisioning.NewEncryptionKeys(apiClient),
		Endpoints:                        networking.NewEndpoints(apiClient),
		FederationPolicy:                 oauth2.NewAccountFederationPolicy(apiClient),
		GroupsV2:                         iam.NewAccountGroupsV2(apiClient),
		IamV2:                            iamv2.NewAccountIamV2(apiClient),
		IpAccessLists:                    settings.NewAccountIpAccessLists(apiClient),
		LogDelivery:                      billing.NewLogDelivery(apiClient),
		MetastoreAssignments:             catalog.NewAccountMetastoreAssignments(apiClient),
		Metastores:                       catalog.NewAccountMetastores(apiClient),
		NetworkConnectivity:              settings.NewNetworkConnectivity(apiClient),
		NetworkPolicies:                  settings.NewNetworkPolicies(apiClient),
		Networks:                         provisioning.NewNetworks(apiClient),
		OAuthPublishedApps:               oauth2.NewOAuthPublishedApps(apiClient),
		PrivateAccess:                    provisioning.NewPrivateAccess(apiClient),
		PrivateNetworkGateways:           networking.NewPrivateNetworkGateways(apiClient),
		PublishedAppIntegration:          oauth2.NewPublishedAppIntegration(apiClient),
		ServicePrincipalFederationPolicy: oauth2.NewServicePrincipalFederationPolicy(apiClient),
		ServicePrincipalSecrets:          oauth2.NewServicePrincipalSecrets(apiClient),
		ServicePrincipalsV2:              iam.NewAccountServicePrincipalsV2(apiClient),
		Settings:                         settings.NewAccountSettings(apiClient),
		SettingsV2:                       settingsv2.NewAccountSettingsV2(apiClient),
		Storage:                          provisioning.NewStorage(apiClient),
		StorageCredentials:               catalog.NewAccountStorageCredentials(apiClient),
		UsageDashboards:                  billing.NewUsageDashboards(apiClient),
		UsersV2:                          iam.NewAccountUsersV2(apiClient),
		VpcEndpoints:                     provisioning.NewVpcEndpoints(apiClient),
		WorkspaceAssignment:              iam.NewWorkspaceAssignment(apiClient),
		WorkspaceNetworkConfiguration:    settings.NewWorkspaceNetworkConfiguration(apiClient),
		Workspaces:                       provisioning.NewWorkspaces(apiClient),
		Groups:                           iam.NewAccountGroups(apiClient),
		ServicePrincipals:                iam.NewAccountServicePrincipals(apiClient),
		Users:                            iam.NewAccountUsers(apiClient),
	}, nil
}
