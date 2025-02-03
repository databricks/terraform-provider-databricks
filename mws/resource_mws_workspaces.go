package mws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/tokens"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// DefaultProvisionTimeout is the amount of minutes terraform will wait
// for workspace to be provisioned and DNS entry to be available. Increasing
// this may help with local DNS cache issues.
const DefaultProvisionTimeout = 20 * time.Minute

// NewWorkspacesAPI creates MWSWorkspacesAPI instance from provider meta
func NewWorkspacesAPI(ctx context.Context, m any) WorkspacesAPI {
	return WorkspacesAPI{m.(*common.DatabricksClient), ctx}
}

// WorkspacesAPI exposes the mws workspaces API
type WorkspacesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// List of workspace statuses for provisioning the workspace
const (
	WorkspaceStatusNotProvisioned = "NOT_PROVISIONED"
	WorkspaceStatusProvisioning   = "PROVISIONING"
	WorkspaceStatusRunning        = "RUNNING"
	WorkspaceStatusFailed         = "FAILED"
	WorkspaceStatusCanceled       = "CANCELLED"
)

// WorkspaceStatusesNonRunnable is a list of statuses in which the workspace is not runnable
var WorkspaceStatusesNonRunnable = []string{WorkspaceStatusCanceled, WorkspaceStatusFailed}

type GCP struct {
	ProjectID string `json:"project_id"`
}

type CloudResourceContainer struct {
	GCP *GCP `json:"gcp"`
}

type GCPManagedNetworkConfig struct {
	SubnetCIDR               string `json:"subnet_cidr" tf:"force_new"`
	GKEClusterPodIPRange     string `json:"gke_cluster_pod_ip_range" tf:"force_new"`
	GKEClusterServiceIPRange string `json:"gke_cluster_service_ip_range" tf:"force_new"`
}

type GkeConfig struct {
	ConnectivityType string `json:"connectivity_type" tf:"force_new"`
	MasterIPRange    string `json:"master_ip_range" tf:"force_new"`
}

type externalCustomerInfo struct {
	CustomerName              string `json:"customer_name"`
	AuthoritativeUserEmail    string `json:"authoritative_user_email"`
	AuthoritativeUserFullName string `json:"authoritative_user_full_name"`
}

// Workspace is the object that contains all the information for deploying a workspace
type Workspace struct {
	AccountID                           string                   `json:"account_id"`
	WorkspaceName                       string                   `json:"workspace_name"`
	DeploymentName                      string                   `json:"deployment_name,omitempty"`
	AwsRegion                           string                   `json:"aws_region,omitempty"`               // required for AWS, not allowed for GCP
	CredentialsID                       string                   `json:"credentials_id,omitempty"`           // required for AWS, not allowed for GCP
	CustomerManagedKeyID                string                   `json:"customer_managed_key_id,omitempty"`  // just for compatibility, will be removed
	StorageConfigurationID              string                   `json:"storage_configuration_id,omitempty"` // required for AWS, not allowed for GCP
	ManagedServicesCustomerManagedKeyID string                   `json:"managed_services_customer_managed_key_id,omitempty"`
	StorageCustomerManagedKeyID         string                   `json:"storage_customer_managed_key_id,omitempty"`
	PricingTier                         string                   `json:"pricing_tier,omitempty" tf:"computed"`
	PrivateAccessSettingsID             string                   `json:"private_access_settings_id,omitempty"`
	NetworkID                           string                   `json:"network_id,omitempty" tf:"suppress_diff"`
	IsNoPublicIPEnabled                 bool                     `json:"is_no_public_ip_enabled" tf:"optional,default:true"`
	WorkspaceID                         int64                    `json:"workspace_id,omitempty" tf:"computed"`
	WorkspaceURL                        string                   `json:"workspace_url,omitempty" tf:"computed"`
	WorkspaceStatus                     string                   `json:"workspace_status,omitempty" tf:"computed"`
	WorkspaceStatusMessage              string                   `json:"workspace_status_message,omitempty" tf:"computed"`
	CreationTime                        int64                    `json:"creation_time,omitempty" tf:"computed"`
	ExternalCustomerInfo                *externalCustomerInfo    `json:"external_customer_info,omitempty"`
	CloudResourceBucket                 *CloudResourceContainer  `json:"cloud_resource_container,omitempty"`
	GCPManagedNetworkConfig             *GCPManagedNetworkConfig `json:"gcp_managed_network_config,omitempty" tf:"suppress_diff"`
	GkeConfig                           *GkeConfig               `json:"gke_config,omitempty" tf:"suppress_diff"`
	Cloud                               string                   `json:"cloud,omitempty" tf:"computed"`
	Location                            string                   `json:"location,omitempty"`
	CustomTags                          map[string]string        `json:"custom_tags,omitempty"` // Optional for AWS, not allowed for GCP
}

// this type alias hack is required for Marshaller to work without an infinite loop
type aWorkspace Workspace

// MarshalJSON is required to overcome the limitations of `omitempty` usage with reflect_resource.go
// for workspace creation in Accounts API for AWS and GCP. It exits early on AWS and picks only
// the relevant fields for GCP.
func (w *Workspace) MarshalJSON() ([]byte, error) {
	if w.Cloud != "gcp" {
		return json.Marshal(aWorkspace(*w))
	}
	workspaceCreationRequest := map[string]any{
		"account_id":               w.AccountID,
		"cloud":                    w.Cloud,
		"cloud_resource_container": w.CloudResourceBucket,
		"location":                 w.Location,
		"workspace_name":           w.WorkspaceName,
	}
	if w.NetworkID != "" {
		workspaceCreationRequest["network_id"] = w.NetworkID
	}
	if w.PrivateAccessSettingsID != "" {
		workspaceCreationRequest["private_access_settings_id"] = w.PrivateAccessSettingsID
	}
	if w.GkeConfig != nil {
		workspaceCreationRequest["gke_config"] = w.GkeConfig
	}
	if w.GCPManagedNetworkConfig != nil {
		workspaceCreationRequest["gcp_managed_network_config"] = w.GCPManagedNetworkConfig
	}
	if w.ManagedServicesCustomerManagedKeyID != "" {
		workspaceCreationRequest["managed_services_customer_managed_key_id"] = w.ManagedServicesCustomerManagedKeyID
	}
	if w.StorageCustomerManagedKeyID != "" {
		workspaceCreationRequest["storage_customer_managed_key_id"] = w.StorageCustomerManagedKeyID
	}
	return json.Marshal(workspaceCreationRequest)
}

// Create deploys the workspace and waits till it's properly running.
// In case of error, it removes the failed deployment and returns the message
func (a WorkspacesAPI) Create(ws *Workspace, timeout time.Duration) error {
	if a.client.IsGcp() {
		ws.Cloud = "gcp"
	}
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", ws.AccountID)
	err := a.client.Post(a.context, workspacesAPIPath, ws, &ws)
	if err != nil {
		return err
	}
	if err = a.WaitForRunning(*ws, timeout); err != nil {
		log.Printf("[ERROR] Deleting failed workspace: %s", err)
		if derr := a.Delete(ws.AccountID, fmt.Sprintf("%d", ws.WorkspaceID)); derr != nil {
			return fmt.Errorf("%s - %s", err, derr)
		}
		return err
	}
	if ws.WorkspaceURL == "" {
		// WorkspaceURL is computed, yet very important field
		host := generateWorkspaceHostname(a.client, *ws)
		ws.WorkspaceURL = fmt.Sprintf("https://%s", host)
	}
	return nil
}

// generateWorkspaceHostname computes the hostname for the specified workspace,
// given the account console hostname.
func generateWorkspaceHostname(client *common.DatabricksClient, ws Workspace) string {
	u, err := url.Parse(client.Config.Host)
	if err != nil {
		// Fallback.
		log.Printf("[WARN] Unable to parse URL from client host: %v", err)
		return ws.DeploymentName + ".cloud.databricks.com"
	}

	// We expect the account console hostname to be of the form `accounts.foo[.bar]...`
	// The workspace hostname can be generated by replacing `accounts` with the deployment name.
	// If the hostname is an IP address, we're in testing mode and do fallback.
	chunks := strings.Split(u.Hostname(), ".")
	if len(chunks) == 0 || net.ParseIP(u.Hostname()) != nil {
		// Fallback.
		log.Printf("[WARN] Unable to split client host: %v", u.Hostname())
		return ws.DeploymentName + ".cloud.databricks.com"
	}
	chunks[0] = ws.DeploymentName
	return strings.Join(chunks, ".")
}

func (a WorkspacesAPI) verifyWorkspaceReachable(ws Workspace) *resource.RetryError {
	ctx, cancel := context.WithTimeout(a.context, 10*time.Second)
	defer cancel()
	// wait for DNS caches to refresh, as sometimes we cannot make
	// API calls to new workspaces immediately after it's created
	wsClient, err := a.client.ClientForHost(a.context, ws.WorkspaceURL)
	if err != nil {
		return resource.NonRetryableError(err)
	}
	// make a request to SCIM API, just to verify there are no errors
	var response map[string]any
	err = wsClient.Get(ctx, "/preview/scim/v2/Me", nil, &response)
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		err = fmt.Errorf("workspace %s is not yet reachable: %s",
			ws.WorkspaceURL, dnsError)
		log.Printf("[INFO] %s", err)
		// expected to retry on: dial tcp: lookup XXX: no such host
		return resource.RetryableError(err)
	}
	return nil
}

func (a WorkspacesAPI) explainWorkspaceFailure(ws Workspace) error {
	if ws.NetworkID == "" {
		return fmt.Errorf(ws.WorkspaceStatusMessage)
	}
	network, nerr := NewNetworksAPI(a.context, a.client).Read(ws.AccountID, ws.NetworkID)
	if nerr != nil {
		return fmt.Errorf("failed to start workspace. Cannot read network: %s", nerr)
	}
	var strBuffer bytes.Buffer
	for _, networkHealth := range network.ErrorMessages {
		strBuffer.WriteString(fmt.Sprintf("error: %s;error_msg: %s;",
			networkHealth.ErrorType, networkHealth.ErrorMessage))
	}
	return fmt.Errorf("Workspace failed to create: %v, network error message: %v",
		ws.WorkspaceStatusMessage, strBuffer.String())
}

// WaitForRunning will wait until workspace is running, otherwise will try to explain why it failed
func (a WorkspacesAPI) WaitForRunning(ws Workspace, timeout time.Duration) error {
	return resource.RetryContext(a.context, timeout, func() *resource.RetryError {
		workspace, err := a.Read(ws.AccountID, fmt.Sprintf("%d", ws.WorkspaceID))
		if err != nil {
			return resource.NonRetryableError(err)
		}
		switch workspace.WorkspaceStatus {
		case WorkspaceStatusRunning:
			log.Printf("[INFO] Workspace is now running")
			if strings.Contains(ws.DeploymentName, "900150983cd24fb0") {
				// nobody would probably name workspace as 900150983cd24fb0,
				// so we'll use it as unit testing shim
				return nil
			}
			return a.verifyWorkspaceReachable(workspace)
		case WorkspaceStatusCanceled, WorkspaceStatusFailed:
			log.Printf("[ERROR] Cannot start workspace: %s", workspace.WorkspaceStatusMessage)
			err = a.explainWorkspaceFailure(workspace)
			return resource.NonRetryableError(err)
		default:
			log.Printf("[INFO] Workspace %s is %s: %s", workspace.DeploymentName,
				workspace.WorkspaceStatus, workspace.WorkspaceStatusMessage)
			return resource.RetryableError(fmt.Errorf(workspace.WorkspaceStatusMessage))
		}
	})
}

var workspaceRunningUpdatesAllowed = []string{"credentials_id", "network_id", "storage_customer_managed_key_id", "private_access_settings_id", "managed_services_customer_managed_key_id", "custom_tags"}

// UpdateRunning will update running workspace with couple of possible fields
func (a WorkspacesAPI) UpdateRunning(ws Workspace, timeout time.Duration) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", ws.AccountID, ws.WorkspaceID)
	request := map[string]any{}

	if ws.CredentialsID != "" {
		request["credentials_id"] = ws.CredentialsID
	}

	// The ID of the workspace's network configuration object. Used only if you already use a customer-managed VPC.
	// This change is supported only if you specified a network configuration ID when the workspace was created.
	// In other words, you cannot switch from a Databricks-managed VPC to a customer-managed VPC. This parameter
	// is available for updating both failed and running workspaces.
	if ws.NetworkID != "" {
		request["network_id"] = ws.NetworkID
	}

	if ws.PrivateAccessSettingsID != "" {
		request["private_access_settings_id"] = ws.PrivateAccessSettingsID
	}
	if ws.StorageCustomerManagedKeyID != "" {
		request["storage_customer_managed_key_id"] = ws.StorageCustomerManagedKeyID
	}
	if ws.CustomTags != nil {
		if !a.client.IsAws() {
			return fmt.Errorf("custom_tags are only allowed for AWS workspaces")
		}
		request["custom_tags"] = ws.CustomTags
	}

	if len(request) == 0 {
		return nil
	}

	err := a.client.Patch(a.context, workspacesAPIPath, request)
	if err != nil {
		return err
	}
	return a.WaitForRunning(ws, timeout)
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a WorkspacesAPI) Read(mwsAcctID, workspaceID string) (Workspace, error) {
	var mwsWorkspace Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%s", mwsAcctID, workspaceID)
	err := a.client.Get(a.context, workspacesAPIPath, nil, &mwsWorkspace)
	if err == nil && mwsWorkspace.WorkspaceURL == "" {
		// generate workspace URL based on client's hostname, if response contains no URL
		host := generateWorkspaceHostname(a.client, mwsWorkspace)
		mwsWorkspace.WorkspaceURL = fmt.Sprintf("https://%s", host)
	}
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id
// and wait till it's properly removed
func (a WorkspacesAPI) Delete(mwsAcctID, workspaceID string) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%s", mwsAcctID, workspaceID)
	err := a.client.Delete(a.context, workspacesAPIPath, nil)
	if err != nil {
		return err
	}
	return resource.RetryContext(a.context, 15*time.Minute, func() *resource.RetryError {
		workspace, err := a.Read(mwsAcctID, workspaceID)
		if apierr.IsMissing(err) {
			log.Printf("[INFO] Workspace %s/%s is removed.", mwsAcctID, workspaceID)
			return nil
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		msg := fmt.Errorf("Workspace %s is not removed yet. Workspace status: %s %s",
			workspace.WorkspaceName, workspace.WorkspaceStatus, workspace.WorkspaceStatusMessage)
		log.Printf("[INFO] %s", msg)
		return resource.RetryableError(msg)
	})
}

// List will list all workspaces in a given mws account
func (a WorkspacesAPI) List(mwsAcctID string) ([]Workspace, error) {
	var mwsWorkspacesList []Workspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	err := a.client.Get(a.context, workspacesAPIPath, nil, &mwsWorkspacesList)
	return mwsWorkspacesList, err
}

type Token struct {
	LifetimeSeconds int32           `json:"lifetime_seconds,omitempty" tf:"default:2592000"`
	Comment         string          `json:"comment,omitempty" tf:"default:Terraform PAT"`
	TokenID         string          `json:"token_id,omitempty" tf:"computed"`
	TokenValue      SensitiveString `json:"token_value,omitempty" tf:"computed,sensitive"`
}

type SensitiveString string

func (s SensitiveString) GoString() string {
	return "****"
}

func (s SensitiveString) String() string {
	return "****"
}

// ephemeral entity to use with StructToData()
type WorkspaceToken struct {
	WorkspaceURL string `json:"workspace_url,omitempty"`
	Token        *Token `json:"token,omitempty"`
}

func CreateTokenIfNeeded(workspacesAPI WorkspacesAPI,
	workspaceSchema map[string]*schema.Schema, d *schema.ResourceData,
) error {
	var wsToken WorkspaceToken
	common.DataToStructPointer(d, workspaceSchema, &wsToken)
	if wsToken.Token == nil {
		return nil
	}
	client, err := workspacesAPI.client.ClientForHost(workspacesAPI.context, wsToken.WorkspaceURL)
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(workspacesAPI.context, client)
	lifetime := time.Duration(wsToken.Token.LifetimeSeconds) * time.Second
	token, err := tokensAPI.Create(lifetime, wsToken.Token.Comment)
	if isInvalidClient(err) {
		return fmt.Errorf("cannot create token: the principal used by Databricks (client ID %s) is not authorized to create a token in this workspace. "+
			"If this is a UC-enabled workspace, add this client to the workspace, either using databricks_mws_permission_assignment or manually (see https://docs.databricks.com/en/admin/users-groups/service-principals.html#assign-a-service-principal-to-a-workspace-using-the-account-console for instructions). "+
			"If this is not a UC-enabled workspace, remove the token block from this configuration and create a workspace-level service principal to configure resources in the workspace (see https://docs.databricks.com/en/admin/users-groups/service-principals.html#add-a-service-principal-to-a-workspace-using-the-workspace-admin-settings for instructions)", client.Config.ClientID)
	}
	if err != nil {
		return fmt.Errorf("cannot create token: %w", err)
	}
	wsToken.Token.TokenID = token.TokenInfo.TokenID
	wsToken.Token.TokenValue = SensitiveString(token.TokenValue)
	return common.StructToData(wsToken, workspaceSchema, d)
}

// isInvalidClient checks whether the API request failed due to the client being invalid.
// This can happen if the provided client does not belong to the workspace. For UC workspaces,
// it is possible for an admin to add the user/service principal used by Terraform to the
// workspace so that Terraform is able to create a token. For non-UC workspaces, it is not
// possible to add account-level users/service principals to the workspace, so customers
// need to manually create a workspace-level service principal and use it to authenticate to
// the workspace.
func isInvalidClient(err error) bool {
	return errors.Is(err, databricks.ErrUnauthenticated)
}

func EnsureTokenExistsIfNeeded(a WorkspacesAPI,
	workspaceSchema map[string]*schema.Schema, d *schema.ResourceData,
) error {
	var wsToken WorkspaceToken
	common.DataToStructPointer(d, workspaceSchema, &wsToken)
	if wsToken.Token == nil {
		return nil
	}
	client, err := a.client.ClientForHost(a.context, wsToken.WorkspaceURL)
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(a.context, client)
	_, err = tokensAPI.Read(wsToken.Token.TokenID)
	// If we cannot authenticate to the workspace and we're using an in-house OAuth principal,
	// log a warning but do not fail. This can happen if the provider is authenticated with a
	// different principal than was used to create the workspace.
	if isInvalidClient(err) {
		tflog.Debug(a.context, fmt.Sprintf("unable to fetch token with ID %s from workspace using the provided service principal, continuing", wsToken.Token.TokenID))
		return nil
	}
	if apierr.IsMissing(err) {
		return CreateTokenIfNeeded(a, workspaceSchema, d)
	}
	if err != nil {
		return fmt.Errorf("cannot read token: %w", err)
	}
	return nil
}

func removeTokenIfNeeded(a WorkspacesAPI, tokenID string, d *schema.ResourceData) error {
	client, err := a.client.ClientForHost(a.context, d.Get("workspace_url").(string))
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(a.context, client)
	err = tokensAPI.Delete(tokenID)
	if isInvalidClient(err) {
		tflog.Debug(a.context, fmt.Sprintf("unable to delete token with ID %s from workspace using the provided service principal, continuing", tokenID))
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot remove token: %w", err)
	}
	return d.Set("token", nil)
}

func UpdateTokenIfNeeded(workspacesAPI WorkspacesAPI,
	workspaceSchema map[string]*schema.Schema, d *schema.ResourceData,
) error {
	o, n := d.GetChange("token")
	old, new := o.([]any), n.([]any)
	if d.HasChange("token") {
		switch {
		case len(old) == 0 && len(new) > 0: // create
			return CreateTokenIfNeeded(workspacesAPI, workspaceSchema, d)
		case len(old) > 0 && len(new) == 0: // delete
			raw := old[0].(map[string]any)
			id := raw["token_id"].(string)
			return removeTokenIfNeeded(workspacesAPI, id, d)
		case len(old) > 0 && len(new) > 0: // delete & create
			rawOld := old[0].(map[string]any)
			id := rawOld["token_id"].(string)
			err := removeTokenIfNeeded(workspacesAPI, id, d)
			if err != nil {
				return err
			}
			rawNew := new[0].(map[string]any)
			d.Set("token", []any{
				map[string]any{
					"lifetime_seconds": rawNew["lifetime_seconds"],
					"comment":          rawNew["comment"],
				},
			})
			return CreateTokenIfNeeded(workspacesAPI, workspaceSchema, d)
		}
	}
	return nil
}

// ResourceMwsWorkspaces manages E2 workspaces
func ResourceMwsWorkspaces() common.Resource {
	workspaceSchema := common.StructToSchema(Workspace{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			for name, fieldSchema := range s {
				if fieldSchema.Computed {
					// skip checking all changes from remote state
					continue
				}
				fieldSchema.ForceNew = true
				for _, allowed := range workspaceRunningUpdatesAllowed {
					if allowed == name {
						// allow updating only a few specific fields
						fieldSchema.ForceNew = false
						break
					}
				}
			}
			s["account_id"].Sensitive = true
			s["deployment_name"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" && new != "" {
					return false
				}
				// Most of E2 accounts require a prefix and API returns it.
				// This is certainly a hack to get things working for Terraform operating model.
				// https://github.com/databricks/terraform-provider-databricks/issues/382
				return !strings.HasSuffix(new, old)
			}
			// The value of `is_no_public_ip_enabled` isn't part of the GET payload.
			// Keep diff when creating (i.e. `old` == ""), suppress diff otherwise.
			s["is_no_public_ip_enabled"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old != ""
			}

			s["pricing_tier"].DiffSuppressFunc = common.EqualFoldDiffSuppress

			s["customer_managed_key_id"].Deprecated = "Use managed_services_customer_managed_key_id instead"
			s["customer_managed_key_id"].ConflictsWith = []string{"managed_services_customer_managed_key_id", "storage_customer_managed_key_id"}
			s["managed_services_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
			s["storage_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
			// manage workspace-specific PAT token
			s["token"] = common.StructToSchema(WorkspaceToken{},
				func(m map[string]*schema.Schema) map[string]*schema.Schema {
					return m
				})["token"]

			s["gcp_workspace_sa"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			return s
		})
	p := common.NewPairSeparatedID("account_id", "workspace_id", "/").Schema(
		func(_ map[string]*schema.Schema) map[string]*schema.Schema {
			return workspaceSchema
		})
	requireFields := func(onThisCloud bool, d *schema.ResourceData, fields ...string) error {
		if !onThisCloud {
			return nil
		}
		for _, fieldName := range fields {
			if d.Get(fieldName) == workspaceSchema[fieldName].ZeroValue() {
				return fmt.Errorf("%s is required", fieldName)
			}
		}
		return nil
	}
	return common.Resource{
		Schema:        workspaceSchema,
		SchemaVersion: 3,
		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 2,
				Type:    workspaceSchemaV2(),
				Upgrade: workspaceMigrateV2,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			common.DataToStructPointer(d, workspaceSchema, &workspace)
			if err := requireFields(c.IsAws(), d, "aws_region", "credentials_id", "storage_configuration_id"); err != nil {
				return err
			}
			if err := requireFields(c.IsGcp(), d, "location"); err != nil {
				return err
			}
			if !c.IsAws() && workspace.CustomTags != nil {
				return fmt.Errorf("custom_tags are only allowed for AWS workspaces")
			}
			if len(workspace.CustomerManagedKeyID) > 0 && len(workspace.ManagedServicesCustomerManagedKeyID) == 0 {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				workspace.ManagedServicesCustomerManagedKeyID = workspace.CustomerManagedKeyID
				workspace.CustomerManagedKeyID = ""
			}
			if err := workspacesAPI.Create(&workspace, d.Timeout(schema.TimeoutCreate)); err != nil {
				return err
			}
			d.Set("workspace_id", workspace.WorkspaceID)
			d.Set("workspace_url", workspace.WorkspaceURL)
			if workspace.Cloud == "gcp" {
				d.Set("gcp_workspace_sa", fmt.Sprintf("db-%d@prod-gcp-%s.iam.gserviceaccount.com",
					workspace.WorkspaceID, workspace.Location))
			}
			p.Pack(d)
			return CreateTokenIfNeeded(workspacesAPI, workspaceSchema, d)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			workspace, err := workspacesAPI.Read(accountID, workspaceID)
			if err != nil {
				return err
			}
			// Default the value of `is_no_public_ip_enabled` because it isn't part of the GET payload.
			// The field is only used on creation and we therefore suppress all diffs.
			workspace.IsNoPublicIPEnabled = true
			if err = common.StructToData(workspace, workspaceSchema, d); err != nil {
				return err
			}
			err = workspacesAPI.WaitForRunning(workspace, d.Timeout(schema.TimeoutRead))
			if err != nil {
				return err
			}
			return EnsureTokenExistsIfNeeded(workspacesAPI, workspaceSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			common.DataToStructPointer(d, workspaceSchema, &workspace)
			if len(workspace.CustomerManagedKeyID) > 0 && len(workspace.ManagedServicesCustomerManagedKeyID) == 0 {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				workspace.ManagedServicesCustomerManagedKeyID = workspace.CustomerManagedKeyID
				workspace.CustomerManagedKeyID = ""
			}
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if d.HasChangeExcept("token") {
				err := workspacesAPI.UpdateRunning(workspace, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return err
				}
			}
			return UpdateTokenIfNeeded(workspacesAPI, workspaceSchema, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewWorkspacesAPI(ctx, c).Delete(accountID, workspaceID)
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			old, new := d.GetChange("private_access_settings_id")
			if old != "" && new == "" {
				return fmt.Errorf("cannot remove private access setting from workspace")
			}
			return nil
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Read:   schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}
}

func workspaceMigrateV2(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "cloud_resource_bucket":
			newState["cloud_resource_container"] = v
			log.Printf("[INFO] cloud_resource_bucket is renamed to cloud_resource_container")
		case "network":
			block, ok := rawState["network"].([]any)
			if !ok {
				log.Printf("[ERROR] how can network not be a single-element list?")
				continue
			}
			if len(block) == 0 {
				log.Printf("[ERROR] network block is empty")
				continue
			}
			oldNetwork, ok := block[0].(map[string]any)
			if !ok {
				log.Printf("[ERROR] how can network not be a map?..")
				continue
			}
			networkId, ok := oldNetwork["network_id"]
			if ok {
				newState["network_id"] = networkId
			}
			unsafeCommonNetworkConfig, ok := oldNetwork["gcp_common_network_config"]
			if ok {
				blocks, ok := unsafeCommonNetworkConfig.([]any)
				if ok {
					old, ok := blocks[0].(map[string]any)
					if ok {
						newState["gke_config"] = []any{
							map[string]any{
								"master_ip_range":   old["gke_cluster_master_ip_range"],
								"connectivity_type": old["gke_connectivity_type"],
							},
						}
						log.Printf("[INFO] moved network.gcp_common_network_config to gke_config")
					}
				}
			}
			managedNetworkConfig, ok := oldNetwork["gcp_managed_network_config"]
			if ok {
				newState["gcp_managed_network_config"] = managedNetworkConfig
			}
			log.Printf("[INFO] network fields are moved to the top level")
		default:
			newState[k] = v
		}
	}
	return newState, nil
}

func workspaceSchemaV2() cty.Type {
	return (&schema.Resource{
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"aws_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"credentials_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"customer_managed_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_no_public_ip_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_services_customer_managed_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pricing_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_access_settings_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_customer_managed_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"workspace_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workspace_status_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workspace_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_resource_bucket": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gcp": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"project_id": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"external_customer_info": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authoritative_user_email": {
							Type:     schema.TypeString,
							Required: true,
						},
						"authoritative_user_full_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"customer_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"network": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"gcp_common_network_config": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gke_cluster_master_ip_range": {
										Type:     schema.TypeString,
										Required: true,
									},
									"gke_connectivity_type": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"gcp_managed_network_config": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gke_cluster_pod_ip_range": {
										Type:     schema.TypeString,
										Required: true,
									},
									"gke_cluster_service_ip_range": {
										Type:     schema.TypeString,
										Required: true,
									},
									"subnet_cidr": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"token": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lifetime_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"token_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"token_value": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
			"custom_tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}).CoreConfigSchema().ImpliedType()
}
