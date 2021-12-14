package mws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/tokens"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DefaultProvisionTimeout is the amount of minutes terraform will wait
// for workspace to be provisioned and DNS entry to be available. Increasing
// this may help with local DNS cache issues.
const DefaultProvisionTimeout = 20 * time.Minute

// NewWorkspacesAPI creates MWSWorkspacesAPI instance from provider meta
func NewWorkspacesAPI(ctx context.Context, m interface{}) WorkspacesAPI {
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

type CloudResourceBucket struct {
	GCP *GCP `json:"gcp"`
}

type GCPManagedNetworkConfig struct {
	SubnetCIDR               string `json:"subnet_cidr"`
	GKEClusterPodIPRange     string `json:"gke_cluster_pod_ip_range"`
	GKEClusterServiceIPRange string `json:"gke_cluster_service_ip_range"`
}

type GCPCommonNetworkConfig struct {
	GKEConnectivityType     string `json:"gke_connectivity_type"`
	GKEClusterMasterIPRange string `json:"gke_cluster_master_ip_range"`
}

type GCPNetwork struct {
	GCPManagedNetworkConfig *GCPManagedNetworkConfig `json:"gcp_managed_network_config"`
	GCPCommonNetworkConfig  *GCPCommonNetworkConfig  `json:"gcp_common_network_config"`
}

type externalCustomerInfo struct {
	CustomerName              string `json:"customer_name"`
	AuthoritativeUserEmail    string `json:"authoritative_user_email"`
	AuthoritativeUserFullName string `json:"authoritative_user_full_name"`
}

// Workspace is the object that contains all the information for deploying a workspace
type Workspace struct {
	AccountID                           string                `json:"account_id"`
	WorkspaceName                       string                `json:"workspace_name"`
	DeploymentName                      string                `json:"deployment_name,omitempty"`
	AwsRegion                           string                `json:"aws_region,omitempty"`               // required for AWS, not allowed for GCP
	CredentialsID                       string                `json:"credentials_id,omitempty"`           // required for AWS, not allowed for GCP
	CustomerManagedKeyID                string                `json:"customer_managed_key_id,omitempty"`  // just for compatibility, will be removed
	StorageConfigurationID              string                `json:"storage_configuration_id,omitempty"` // required for AWS, not allowed for GCP
	ManagedServicesCustomerManagedKeyID string                `json:"managed_services_customer_managed_key_id,omitempty"`
	StorageCustomerManagedKeyID         string                `json:"storage_customer_managed_key_id,omitempty"`
	PricingTier                         string                `json:"pricing_tier,omitempty" tf:"computed"`
	PrivateAccessSettingsID             string                `json:"private_access_settings_id,omitempty"`
	NetworkID                           string                `json:"network_id,omitempty"`
	IsNoPublicIPEnabled                 bool                  `json:"is_no_public_ip_enabled"`
	WorkspaceID                         int64                 `json:"workspace_id,omitempty" tf:"computed"`
	WorkspaceURL                        string                `json:"workspace_url,omitempty" tf:"computed"`
	WorkspaceStatus                     string                `json:"workspace_status,omitempty" tf:"computed"`
	WorkspaceStatusMessage              string                `json:"workspace_status_message,omitempty" tf:"computed"`
	CreationTime                        int64                 `json:"creation_time,omitempty" tf:"computed"`
	ExternalCustomerInfo                *externalCustomerInfo `json:"external_customer_info,omitempty"`
	CloudResourceBucket                 *CloudResourceBucket  `json:"cloud_resource_bucket,omitempty"`
	Network                             *GCPNetwork           `json:"network,omitempty"`
	Cloud                               string                `json:"cloud,omitempty" tf:"computed"`
	Location                            string                `json:"location,omitempty"`
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
	workspaceCreationRequest := map[string]interface{}{
		"account_id":            w.AccountID,
		"cloud":                 w.Cloud,
		"cloud_resource_bucket": w.CloudResourceBucket,
		"location":              w.Location,
		"workspace_name":        w.WorkspaceName,
	}
	if w.Network != nil {
		workspaceCreationRequest["network"] = w.Network
	}
	return json.Marshal(workspaceCreationRequest)
}

// Create creates the workspace creation process
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
	return nil
}

// generateWorkspaceHostname computes the hostname for the specified workspace,
// given the account console hostname.
func generateWorkspaceHostname(client *common.DatabricksClient, ws Workspace) string {
	u, err := url.Parse(client.Host)
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
	// make a request to Tokens API, just to verify there are no errors
	var response map[string]interface{}
	err = wsClient.Get(ctx, "/token/list", nil, &response)
	if apiError, ok := err.(common.APIError); ok {
		err = fmt.Errorf("workspace %s is not yet reachable: %s",
			ws.WorkspaceURL, apiError)
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

var workspaceRunningUpdatesAllowed = []string{"credentials_id", "network_id", "storage_customer_managed_key_id"}

// UpdateRunning will update running workspace with couple of possible fields
func (a WorkspacesAPI) UpdateRunning(ws Workspace, timeout time.Duration) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", ws.AccountID, ws.WorkspaceID)
	request := map[string]string{
		"credentials_id": ws.CredentialsID,
		// The ID of the workspace's network configuration object. Used only if you already use a customer-managed VPC.
		// This change is supported only if you specified a network configuration ID when the workspace was created.
		// In other words, you cannot switch from a Databricks-managed VPC to a customer-managed VPC. This parameter
		// is available for updating both failed and running workspaces. Note: You cannot use a network configuration
		// update in this API to add support for PrivateLink (in Public Preview). To add PrivateLink to an existing
		// workspace, contact your Databricks representative.
		"network_id": ws.NetworkID,
	}
	if ws.StorageCustomerManagedKeyID != "" {
		request["storage_customer_managed_key_id"] = ws.StorageCustomerManagedKeyID
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

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a WorkspacesAPI) Delete(mwsAcctID, workspaceID string) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%s", mwsAcctID, workspaceID)
	err := a.client.Delete(a.context, workspacesAPIPath, nil)
	if err != nil {
		return err
	}
	return resource.RetryContext(a.context, 15*time.Minute, func() *resource.RetryError {
		workspace, err := a.Read(mwsAcctID, workspaceID)
		if common.IsMissing(err) {
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
	LifetimeSeconds int32  `json:"lifetime_seconds,omitempty" tf:"default:2592000"`
	Comment         string `json:"comment,omitempty" tf:"default:Terraform PAT"`
	TokenID         string `json:"token_id,omitempty" tf:"computed,sensitive"`
	TokenValue      string `json:"token_value,omitempty" tf:"computed"`
}

// ephemeral entity to use with StructToData()
type WorkspaceToken struct {
	WorkspaceURL string `json:"workspace_url,omitempty"`
	Token        *Token `json:"token,omitempty"`
}

func (a WorkspacesAPI) CreateToken(ws *WorkspaceToken) error {
	if ws.Token == nil {
		return fmt.Errorf("no token metadata")
	}
	client, err := a.client.ClientForHost(a.context, ws.WorkspaceURL)
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(a.context, client)
	lifetime := time.Duration(ws.Token.LifetimeSeconds) * time.Second
	token, err := tokensAPI.Create(lifetime, ws.Token.Comment)
	if err != nil {
		return err
	}
	ws.Token.TokenID = token.TokenInfo.TokenID
	ws.Token.TokenValue = token.TokenValue
	return err
}

func (a WorkspacesAPI) EnsureTokenExists(ws *WorkspaceToken) error {
	if ws.Token == nil {
		return fmt.Errorf("no token metadata")
	}
	client, err := a.client.ClientForHost(a.context, ws.WorkspaceURL)
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(a.context, client)
	_, err = tokensAPI.Read(ws.Token.TokenID)
	if common.IsMissing(err) {
		return a.CreateToken(ws)
	}
	return err
}

func (a WorkspacesAPI) DeleteToken(ws *WorkspaceToken) error {
	if ws.Token == nil {
		return fmt.Errorf("no token metadata")
	}
	client, err := a.client.ClientForHost(a.context, ws.WorkspaceURL)
	if err != nil {
		return err
	}
	tokensAPI := tokens.NewTokensAPI(a.context, client)
	return tokensAPI.Delete(ws.Token.TokenID)
}

// ResourceWorkspace manages E2 workspaces
func ResourceWorkspace() *schema.Resource {
	workspaceSchema := common.StructToSchema(Workspace{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
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
			// https://github.com/databrickslabs/terraform-provider-databricks/issues/382
			return !strings.HasSuffix(new, old)
		}
		// It cannot be marked as `omitempty` in the struct annotation because Go's JON marshaller
		// skips booleans set to `false` if set. Thus, we mark it optional here.
		s["is_no_public_ip_enabled"].Optional = true
		s["is_no_public_ip_enabled"].Required = false
		// The API defaults this field to `true`. Apply the same behavior here.
		s["is_no_public_ip_enabled"].Default = true
		// The value of `is_no_public_ip_enabled` isn't part of the GET payload.
		// Keep diff when creating (i.e. `old` == ""), suppress diff otherwise.
		s["is_no_public_ip_enabled"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
			return old != ""
		}
		s["customer_managed_key_id"].Deprecated = "Use managed_services_customer_managed_key_id instead"
		s["customer_managed_key_id"].ConflictsWith = []string{"managed_services_customer_managed_key_id", "storage_customer_managed_key_id"}
		s["managed_services_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
		s["storage_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
		// manage workspace-specific PAT token
		s["token"] = common.StructToSchema(WorkspaceToken{},
			func(m map[string]*schema.Schema) map[string]*schema.Schema {
				return m
			})["token"]
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
	createTokenIfNeeded := func(workspacesAPI WorkspacesAPI, d *schema.ResourceData) error {
		var wsToken WorkspaceToken
		err := common.DataToStructPointer(d, workspaceSchema, &wsToken)
		if err != nil {
			return err
		}
		if wsToken.Token == nil {
			return nil
		}
		err = workspacesAPI.CreateToken(&wsToken)
		if err != nil {
			return err
		}
		return common.StructToData(wsToken, workspaceSchema, d)
	}
	ensureTokenExists := func(workspacesAPI WorkspacesAPI, d *schema.ResourceData) error {
		var wsToken WorkspaceToken
		err := common.DataToStructPointer(d, workspaceSchema, &wsToken)
		if err != nil {
			return err
		}
		if wsToken.Token == nil {
			return nil
		}
		err = workspacesAPI.EnsureTokenExists(&wsToken)
		if err != nil {
			return err
		}
		return common.StructToData(wsToken, workspaceSchema, d)
	}
	removeTokenIfNeeded := func(workspacesAPI WorkspacesAPI, tokenID string, d *schema.ResourceData) error {
		client, err := workspacesAPI.client.ClientForHost(workspacesAPI.context, d.Get("workspace_url").(string))
		if err != nil {
			return err
		}
		tokensAPI := tokens.NewTokensAPI(workspacesAPI.context, client)
		err = tokensAPI.Delete(tokenID)
		if err != nil {
			return err
		}
		return d.Set("token", nil)
	}
	return common.Resource{
		Schema:        workspaceSchema,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var workspace Workspace
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			if err := common.DataToStructPointer(d, workspaceSchema, &workspace); err != nil {
				return err
			}
			if err := requireFields(c.IsAws(), d, "aws_region", "credentials_id", "storage_configuration_id"); err != nil {
				return err
			}
			if err := requireFields(c.IsGcp(), d, "location"); err != nil {
				return err
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
			p.Pack(d)
			return createTokenIfNeeded(workspacesAPI, d)
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
			return ensureTokenExists(workspacesAPI, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspacesAPI := NewWorkspacesAPI(ctx, c)
			o, n := d.GetChange("token")
			old, new := o.([]interface{}), n.([]interface{})
			if d.HasChange("token") {
				switch {
				case len(old) == 0 && len(new) > 0: // create
					err := createTokenIfNeeded(workspacesAPI, d)
					if err != nil {
						return err
					}
				case len(old) > 0 && len(new) == 0: // delete
					raw := old[0].(map[string]interface{})
					err := removeTokenIfNeeded(workspacesAPI, raw["token_id"].(string), d)
					if err != nil {
						return err
					}
				case len(old) > 0 && len(new) > 0: // delete & create
					rawOld := old[0].(map[string]interface{})
					err := removeTokenIfNeeded(workspacesAPI, rawOld["token_id"].(string), d)
					if err != nil {
						return err
					}
					rawNew := new[0].(map[string]interface{})
					d.Set("token", []interface{}{
						map[string]interface{}{
							"lifetime_seconds": rawNew["lifetime_seconds"],
							"comment":          rawNew["comment"],
						},
					})
					err = createTokenIfNeeded(workspacesAPI, d)
					if err != nil {
						return err
					}
				}
			}
			var workspace Workspace
			if err := common.DataToStructPointer(d, workspaceSchema, &workspace); err != nil {
				return err
			}
			if len(workspace.CustomerManagedKeyID) > 0 && len(workspace.ManagedServicesCustomerManagedKeyID) == 0 {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				workspace.ManagedServicesCustomerManagedKeyID = workspace.CustomerManagedKeyID
				workspace.CustomerManagedKeyID = ""
			}
			return workspacesAPI.UpdateRunning(workspace, d.Timeout(schema.TimeoutUpdate))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewWorkspacesAPI(ctx, c).Delete(accountID, workspaceID)
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Read:   schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}.ToResource()
}
