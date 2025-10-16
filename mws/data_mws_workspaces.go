package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

type WorkspaceData struct {
	AccountID                           string `json:"account_id,omitempty" tf:"computed"`
	WorkspaceName                       string `json:"workspace_name,omitempty" tf:"computed"`
	DeploymentName                      string `json:"deployment_name,omitempty" tf:"computed"`
	AwsRegion                           string `json:"aws_region,omitempty" tf:"computed"`
	CredentialsID                       string `json:"credentials_id,omitempty" tf:"computed"`
	CustomerManagedKeyID                string `json:"customer_managed_key_id,omitempty" tf:"computed"`
	StorageConfigurationID              string `json:"storage_configuration_id,omitempty" tf:"computed"`
	ManagedServicesCustomerManagedKeyID string `json:"managed_services_customer_managed_key_id,omitempty" tf:"computed"`
	StorageCustomerManagedKeyID         string `json:"storage_customer_managed_key_id,omitempty" tf:"computed"`
	PricingTier                         string `json:"pricing_tier,omitempty" tf:"computed"`
	PrivateAccessSettingsID             string `json:"private_access_settings_id,omitempty" tf:"computed"`
	NetworkID                           string `json:"network_id,omitempty" tf:"computed"`
	IsNoPublicIPEnabled                 bool   `json:"is_no_public_ip_enabled,omitempty" tf:"computed"`
	WorkspaceID                         int64  `json:"workspace_id,omitempty" tf:"computed"`
	WorkspaceURL                        string `json:"workspace_url,omitempty" tf:"computed"`
	WorkspaceStatus                     string `json:"workspace_status,omitempty" tf:"computed"`
	WorkspaceStatusMessage              string `json:"workspace_status_message,omitempty" tf:"computed"`
	ExpectedWorkspaceStatus             string `json:"expected_workspace_status,omitempty" tf:"computed"`
	CreationTime                        int64  `json:"creation_time,omitempty" tf:"computed"`
	//ExternalCustomerInfo                *externalCustomerInfo    `json:"external_customer_info,omitempty" tf:"computed"`
	//CloudResourceBucket                 *CloudResourceContainer  `json:"cloud_resource_container,omitempty" tf:"computed"`
	//GCPManagedNetworkConfig             *GCPManagedNetworkConfig `json:"gcp_managed_network_config,omitempty" tf:"computed"`
	//GkeConfig                           *GkeConfig               `json:"gke_config,omitempty" tf:"computed"`
	Cloud                string            `json:"cloud,omitempty" tf:"computed"`
	Location             string            `json:"location,omitempty" tf:"computed"`
	CustomTags           map[string]string `json:"custom_tags,omitempty" tf:"computed"`
	ComputeMode          string            `json:"compute_mode,omitempty" tf:"computed"`
	EffectiveComputeMode string            `json:"effective_compute_mode,omitempty" tf:"computed"`
}

func NewWorkspaceDataFromWorkspace(workspace Workspace) WorkspaceData {
	return WorkspaceData{
		AccountID:                           workspace.AccountID,
		WorkspaceName:                       workspace.WorkspaceName,
		DeploymentName:                      workspace.DeploymentName,
		AwsRegion:                           workspace.AwsRegion,
		CredentialsID:                       workspace.CredentialsID,
		CustomerManagedKeyID:                workspace.CustomerManagedKeyID,
		StorageConfigurationID:              workspace.StorageConfigurationID,
		ManagedServicesCustomerManagedKeyID: workspace.ManagedServicesCustomerManagedKeyID,
		StorageCustomerManagedKeyID:         workspace.StorageCustomerManagedKeyID,
		PricingTier:                         workspace.PricingTier,
		PrivateAccessSettingsID:             workspace.PrivateAccessSettingsID,
		NetworkID:                           workspace.NetworkID,
		IsNoPublicIPEnabled:                 workspace.IsNoPublicIPEnabled,
		WorkspaceID:                         workspace.WorkspaceID,
		WorkspaceURL:                        workspace.WorkspaceURL,
		WorkspaceStatus:                     workspace.WorkspaceStatus,
		WorkspaceStatusMessage:              workspace.WorkspaceStatusMessage,
		ExpectedWorkspaceStatus:             workspace.ExpectedWorkspaceStatus,
		CreationTime:                        workspace.CreationTime,
		Cloud:                               workspace.Cloud,
		Location:                            workspace.Location,
		CustomTags:                          workspace.CustomTags,
		ComputeMode:                         workspace.ComputeMode,
		EffectiveComputeMode:                workspace.EffectiveComputeMode,
	}
}

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids           map[string]int64 `json:"ids" tf:"computed"`
		MwsWorkspaces []WorkspaceData  `json:"mws_workspaces" tf:"computed"`
	}
	return common.DataResource(mwsWorkspacesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsWorkspacesData)
		if c.Config.AccountID == "" {
			return fmt.Errorf("provider block is missing `account_id` property")
		}
		workspaces, err := NewWorkspacesAPI(ctx, c).List(c.Config.AccountID)
		if err != nil {
			return err
		}
		data.Ids = map[string]int64{}
		for _, workspace := range workspaces {
			data.Ids[workspace.WorkspaceName] = workspace.WorkspaceID
			data.MwsWorkspaces = append(data.MwsWorkspaces, NewWorkspaceDataFromWorkspace(workspace))
		}
		return nil
	})
}
