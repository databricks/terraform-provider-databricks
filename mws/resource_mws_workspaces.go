package mws

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/docs"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DefaultProvisionTimeout is the amount of minutes terraform will wait
// for workspace to be provisioned and DNS entry to be available. Increasing
// this may help with local DNS cache issues.
const DefaultProvisionTimeout = 20 * time.Minute

// Workspace is the object that contains all the information for deploying a workspace
type Workspace struct {
	provisioning.Workspace
	CustomerManagedKeyID string `json:"customer_managed_key_id,omitempty"` // just for compatibility, will be removed
	GcpWorkspaceSa       string `json:"gcp_workspace_sa" tf:"computed"`
	Token                *Token `json:"token,omitempty"`
	WorkspaceURL         string `json:"workspace_url,omitempty" tf:"computed"`
}

// wait for DNS caches to refresh, as sometimes we cannot make
// API calls to new workspaces immediately after it's created
func verifyWorkspaceReachable(ctx context.Context, w *databricks.WorkspaceClient) error {
	return retries.New[struct{}](retries.WithRetryFunc(func(err error) bool {
		var dnsError *net.DNSError
		if errors.As(err, &dnsError) {
			log.Printf("[INFO] workspace is not yet reachable: %s", dnsError)
			// expected to retry on: dial tcp: lookup XXX: no such host
			return true
		}
		return false
	})).Wait(ctx, func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		_, err := w.CurrentUser.Me(ctx)
		return err
	})
}

func explainWorkspaceFailure(ctx context.Context, a *databricks.AccountClient, workspace *provisioning.Workspace) error {
	errorBase := fmt.Sprintf("workspace status message: %s", workspace.WorkspaceStatusMessage)
	if workspace.NetworkId == "" {
		return errors.New(errorBase)
	}
	network, err := a.Networks.Get(ctx, provisioning.GetNetworkRequest{NetworkId: workspace.NetworkId})
	if err != nil {
		return fmt.Errorf("%s; network error message: cannot read network: %w", errorBase, err)
	}
	var strBuffer bytes.Buffer
	for _, networkHealth := range network.ErrorMessages {
		strBuffer.WriteString(fmt.Sprintf("error: %s;error_msg: %s;", networkHealth.ErrorType, networkHealth.ErrorMessage))
	}
	return fmt.Errorf("%s, network error message: %s", errorBase, strBuffer.String())
}

type Token struct {
	LifetimeSeconds int64           `json:"lifetime_seconds,omitempty" tf:"default:2592000"`
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

func createToken(ctx context.Context, w *databricks.WorkspaceClient, t *Token) error {
	token, err := w.Tokens.Create(ctx, settings.CreateTokenRequest{
		LifetimeSeconds: t.LifetimeSeconds,
		Comment:         t.Comment,
	})
	if isInvalidClient(err) {
		return fmt.Errorf("cannot create token: the principal used by Databricks (client ID %s) is not authorized to create a token in this workspace. "+
			"If this is a UC-enabled workspace, add this client to the workspace, either using databricks_mws_permission_assignment or manually (see https://docs.databricks.com/en/admin/users-groups/service-principals.html#assign-a-service-principal-to-a-workspace-using-the-account-console for instructions). "+
			"If this is not a UC-enabled workspace, remove the token block from this configuration and create a workspace-level service principal to configure resources in the workspace (see https://docs.databricks.com/en/admin/users-groups/service-principals.html#add-a-service-principal-to-a-workspace-using-the-workspace-admin-settings for instructions)", w.Config.ClientID)
	}
	if err != nil {
		return fmt.Errorf("cannot create token: %w", err)
	}
	t.TokenValue = SensitiveString(token.TokenValue)
	t.TokenID = token.TokenInfo.TokenId
	return nil
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

func ensureTokenExists(ctx context.Context, w *databricks.WorkspaceClient, token *Token) error {
	tokens := w.Tokens.List(ctx)
	for tokens.HasNext(ctx) {
		t, err := tokens.Next(ctx)
		// If we cannot authenticate to the workspace and we're using an in-house OAuth principal,
		// log a warning but do not fail. This can happen if the provider is authenticated with a
		// different principal than was used to create the workspace.
		if isInvalidClient(err) {
			tflog.Debug(ctx, fmt.Sprintf("unable to fetch token with ID %s from workspace using the provided service principal, continuing", token.TokenID))
			return nil
		}
		if err != nil {
			return fmt.Errorf("cannot read token: %w", err)
		}
		if t.TokenId == token.TokenID {
			return nil
		}
	}
	return createToken(ctx, w, token)
}

func removeToken(ctx context.Context, w *databricks.WorkspaceClient, tokenID string) error {
	err := w.Tokens.Delete(ctx, settings.RevokeTokenRequest{TokenId: tokenID})
	if isInvalidClient(err) {
		tflog.Debug(ctx, fmt.Sprintf("unable to delete token with ID %s from workspace using the provided service principal, continuing", tokenID))
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot remove token: %w", err)
	}
	return nil
}

// ResourceMwsWorkspaces manages E2 workspaces
func ResourceMwsWorkspaces() common.Resource {
	var computedFields = map[string]struct{}{
		"cloud":                    {},
		"workspace_id":             {},
		"workspace_url":            {},
		"workspace_status":         {},
		"workspace_status_message": {},
		"creation_time":            {},
		"gke_config":               {},
		"pricing_tier":             {},
	}

	var workspaceRunningUpdatesAllowed = map[string]struct{}{
		"credentials_id":                           {},
		"network_id":                               {},
		"storage_customer_managed_key_id":          {},
		"private_access_settings_id":               {},
		"managed_services_customer_managed_key_id": {},
		"custom_tags":                              {},
		"token":                                    {},
	}

	workspaceSchema := common.StructToSchema(Workspace{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			for name, fieldSchema := range s {
				if _, ok := computedFields[name]; ok {
					fieldSchema.Computed = true
				} else {
					_, ok := workspaceRunningUpdatesAllowed[name]
					fieldSchema.ForceNew = !ok
				}
			}
			s["account_id"].Sensitive = true
			common.CustomizeSchemaPath(s, "account_id").SetRequired()
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
			s["is_no_public_ip_enabled"].Default = true

			s["pricing_tier"].DiffSuppressFunc = common.EqualFoldDiffSuppress

			s["customer_managed_key_id"].Deprecated = "Use managed_services_customer_managed_key_id instead"
			s["customer_managed_key_id"].ConflictsWith = []string{"managed_services_customer_managed_key_id", "storage_customer_managed_key_id"}
			s["managed_services_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}
			s["storage_customer_managed_key_id"].ConflictsWith = []string{"customer_managed_key_id"}

			docOptions := docs.DocOptions{
				Section:  docs.Guides,
				Slug:     "gcp-workspace",
				Fragment: "creating-a-databricks-workspace",
			}
			common.CustomizeSchemaPath(s, "gke_config").SetDeprecated(getGkeDeprecationMessage("gke_config", docOptions))
			common.CustomizeSchemaPath(s, "gcp_managed_network_config", "gke_cluster_pod_ip_range").SetDeprecated(getGkeDeprecationMessage("gcp_managed_network_config.gke_cluster_pod_ip_range", docOptions))
			common.CustomizeSchemaPath(s, "gcp_managed_network_config", "gke_cluster_service_ip_range").SetDeprecated(getGkeDeprecationMessage("gcp_managed_network_config.gke_cluster_service_ip_range", docOptions))
			common.CustomizeSchemaPath(s, "gcp_managed_network_config", "subnet_cidr").SetRequired()
			common.CustomizeSchemaPath(s, "workspace_name").SetRequired()
			common.CustomizeSchemaPath(s, "cloud_resource_container", "gcp").SetMinItems(1)
			common.CustomizeSchemaPath(s, "cloud_resource_container", "gcp", "project_id").SetRequired()
			for _, field := range []string{"authoritative_user_email", "authoritative_user_full_name", "customer_name"} {
				common.CustomizeSchemaPath(s, "external_customer_info", field).SetRequired()
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
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			var createWorkspaceRequest provisioning.CreateWorkspaceRequest
			common.DataToStructPointer(d, workspaceSchema, &createWorkspaceRequest)
			var workspaceConfig Workspace
			common.DataToStructPointer(d, workspaceSchema, &workspaceConfig)

			// Validate required fields by cloud
			if err := requireFields(c.IsAws(), d, "aws_region", "credentials_id", "storage_configuration_id"); err != nil {
				return err
			}
			if err := requireFields(c.IsGcp(), d, "location"); err != nil {
				return err
			}
			if !c.IsAws() && createWorkspaceRequest.CustomTags != nil {
				return fmt.Errorf("custom_tags are only allowed for AWS workspaces")
			}
			if customerManagedKeyId, ok := d.GetOk("customer_managed_key_id"); ok && createWorkspaceRequest.ManagedServicesCustomerManagedKeyId == "" {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				createWorkspaceRequest.ManagedServicesCustomerManagedKeyId = customerManagedKeyId.(string)
			}

			// Create the workspace. If creation fails, clean it up and return.
			wait, err := a.Workspaces.Create(ctx, createWorkspaceRequest)
			if err != nil {
				return err
			}
			workspace, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate))
			if err != nil {
				workspace, getErr := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: wait.Response.WorkspaceId})
				if getErr != nil {
					err = fmt.Errorf("workspace creation failed: %w; failed to get workspace: %w", err, getErr)
				} else {
					err = fmt.Errorf("%w: %w", err, explainWorkspaceFailure(ctx, a, workspace))
				}
				log.Printf("[ERROR] Deleting failed workspace: %s", err)
				if derr := a.Workspaces.Delete(ctx, provisioning.DeleteWorkspaceRequest{WorkspaceId: wait.Response.WorkspaceId}); derr != nil {
					return fmt.Errorf("workspace creation failed: %w; failed workspace cleanup failed: %w", err, derr)
				}
				return err
			}

			// Once the workspace is running, wait for the API to be available by polling the SCIM Me endpoint.
			w, err := c.WorkspaceClientForWorkspace(ctx, workspace.WorkspaceId)
			if err != nil {
				return err
			}
			if err := verifyWorkspaceReachable(ctx, w); err != nil {
				return err
			}
			workspaceConfig.Workspace = *workspace
			workspaceConfig.WorkspaceURL = w.Config.CanonicalHostName()
			if c.IsGcp() {
				workspaceConfig.GcpWorkspaceSa = fmt.Sprintf("db-%d@prod-gcp-%s.iam.gserviceaccount.com", workspace.WorkspaceId, workspace.Location)
			}

			// Create a token if requested
			if workspaceConfig.Token != nil {
				err = createToken(ctx, w, workspaceConfig.Token)
				if err != nil {
					return err
				}
			}
			if err := common.StructToData(workspaceConfig, workspaceSchema, d); err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			var workspaceConfig Workspace
			common.DataToStructPointer(d, workspaceSchema, &workspaceConfig)
			workspace, err := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: common.MustInt64(workspaceID)})
			if err != nil {
				return err
			}
			// The gke_config, gcp_managed_network_config.0.gke_cluster_pod_ip_range, and
			// gcp_managed_network_config.0.gke_cluster_service_ip_range fields do not need
			// to be removed from the returned plan because they are marked with "suppress_diff",
			// so their diff will be removed anyways.

			// Default the value of `is_no_public_ip_enabled` because it isn't part of the GET payload.
			// The field is only used on creation and we therefore suppress all diffs.
			workspace.IsNoPublicIpEnabled = true
			workspaceConfig.Workspace = *workspace
			_, err = a.Workspaces.WaitGetWorkspaceRunning(ctx, workspace.WorkspaceId, d.Timeout(schema.TimeoutRead), nil)
			if err != nil {
				return err
			}
			w, err := c.WorkspaceClientForWorkspace(ctx, workspace.WorkspaceId)
			if err != nil {
				return err
			}
			workspaceConfig.WorkspaceURL = w.Config.CanonicalHostName()
			if workspaceConfig.Token != nil {
				if err := ensureTokenExists(ctx, w, workspaceConfig.Token); err != nil {
					return err
				}
			}
			return common.StructToData(workspaceConfig, workspaceSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			var updateWorkspaceRequest provisioning.UpdateWorkspaceRequest
			common.DataToStructPointer(d, workspaceSchema, &updateWorkspaceRequest)
			var workspaceConfig Workspace
			common.DataToStructPointer(d, workspaceSchema, &workspaceConfig)
			// WorkspaceId in UpdateWorkspaceRequest is a path parameter, thus tagged with `json:"-"`.
			// This causes it not to be set in DataToStructPointer. Instead, the workspace ID can be
			// retrieved from Workspace.
			updateWorkspaceRequest.WorkspaceId = workspaceConfig.WorkspaceId
			if customerManagedKeyId, ok := d.GetOk("customer_managed_key_id"); ok && updateWorkspaceRequest.ManagedServicesCustomerManagedKeyId == "" {
				log.Print("[INFO] Using existing customer_managed_key_id as value for new managed_services_customer_managed_key_id")
				updateWorkspaceRequest.ManagedServicesCustomerManagedKeyId = customerManagedKeyId.(string)
			}

			// If the workspace has been modified, call the update API and wait for it to be ready.
			if d.HasChangeExcept("token") {
				wait, err := a.Workspaces.Update(ctx, updateWorkspaceRequest)
				if err != nil {
					return err
				}
				workspace, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return fmt.Errorf("%w: %w", err, explainWorkspaceFailure(ctx, a, workspace))
				}
				workspaceConfig.Workspace = *workspace
			}

			// If the `token` field has been modified, update the token correspondingly.
			if d.HasChange("token") {
				w, err := c.WorkspaceClientForWorkspace(ctx, workspaceConfig.WorkspaceId)
				if err != nil {
					return err
				}

				// If there was a token present in the config before, revoke it.
				rawOld, _ := d.GetChange("token")
				oldTokens := rawOld.([]any)
				if len(oldTokens) > 0 {
					raw := oldTokens[0].(map[string]any)
					id := raw["token_id"].(string)
					if err := removeToken(ctx, w, id); err != nil {
						return err
					}
				}

				// If there is a token present in the config now, create a new one.
				if workspaceConfig.Token != nil {
					if err := createToken(ctx, w, workspaceConfig.Token); err != nil {
						return err
					}
				}
			}
			return common.StructToData(workspaceConfig, workspaceSchema, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, workspaceID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			if err := a.Workspaces.Delete(ctx, provisioning.DeleteWorkspaceRequest{WorkspaceId: common.MustInt64(workspaceID)}); err != nil && !apierr.IsMissing(err) {
				return err
			}
			// Wait for delete
			return retries.New[struct{}]().Wait(ctx, func(ctx context.Context) error {
				_, err := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: common.MustInt64(workspaceID)})
				if err != nil && apierr.IsMissing(err) {
					return nil
				}
				return fmt.Errorf("workspace %s still exists", d.Id())
			})
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			old, new := d.GetChange("private_access_settings_id")
			if old != "" && new == "" {
				return fmt.Errorf("cannot remove private access setting from workspace")
			}
			// For `gke_config`, `gcp_managed_network_config.0.gke_cluster_pod_ip_range` or
			// `gcp_managed_network_config.0.gke_cluster_service_ip_range`, users should be able to
			// remove these keys without recreating the workspace as part of the GKE deprecation process.
			//
			// Otherwise, any change for these keys will cause the workspace resource to be recreated.
			//
			// This should only run on update, thus we skip this check if the ID is not known.
			if d.Id() != "" {
				for _, key := range []string{"gke_config.#", "gcp_managed_network_config.0.gke_cluster_pod_ip_range", "gcp_managed_network_config.0.gke_cluster_service_ip_range"} {
					// These fields are all tagged with "suppress_diff". This means that removing them from
					// the config doesn't result in their disappearing from the diff. Thus, there is no change
					// in the plan when these fields are removed.
					if !d.HasChange(key) {
						continue
					}
					if err := d.ForceNew(key); err != nil {
						return err
					}
				}
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
