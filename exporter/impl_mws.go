package exporter

import (
	"log"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/mws"
)

func listMwsCredentials(ic *importContext) error {
	if !ic.accountClient.Config.IsAws() {
		return nil
	}
	updatedSinceMs := ic.getUpdatedSinceMs()
	creds, err := ic.accountClient.Credentials.List(ic.Context)
	if err != nil {
		return err
	}
	for _, cred := range creds {
		if !ic.MatchesName(cred.CredentialsName) {
			log.Printf("[INFO] Skipping mws_credentials %s because it doesn't match %s", cred.CredentialsName, ic.match)
			continue
		}
		if ic.incremental && cred.CreationTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping mws_credentials '%s' that was created at %d (last active=%d)",
				cred.CredentialsName, cred.CreationTime, updatedSinceMs)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_credentials",
			ID:       ic.accountClient.Config.AccountID + "/" + cred.CredentialsId,
			Name:     cred.CredentialsName,
		})
	}
	return nil
}

func listMwsStorageConfigurations(ic *importContext) error {
	if !ic.accountClient.Config.IsAws() {
		return nil
	}
	updatedSinceMs := ic.getUpdatedSinceMs()
	scs, err := ic.accountClient.Storage.List(ic.Context)
	if err != nil {
		return err
	}
	for _, sc := range scs {
		if !ic.MatchesName(sc.StorageConfigurationName) {
			log.Printf("[INFO] Skipping mws_storage_configurations %s because it doesn't match %s", sc.StorageConfigurationName, ic.match)
			continue
		}
		if ic.incremental && sc.CreationTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping mws_storage_configurations '%s' that was created at %d (last active=%d)",
				sc.StorageConfigurationName, sc.CreationTime, updatedSinceMs)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_storage_configurations",
			ID:       ic.accountClient.Config.AccountID + "/" + sc.StorageConfigurationId,
			Name:     sc.StorageConfigurationName,
		})
	}
	return nil
}

func listMwsVpcEndpoints(ic *importContext) error {
	if ic.accountClient.Config.IsAzure() {
		return nil
	}
	eps, err := ic.accountClient.VpcEndpoints.List(ic.Context)
	if err != nil {
		return err
	}
	for _, ep := range eps {
		if !ic.MatchesName(ep.VpcEndpointName) {
			log.Printf("[INFO] Skipping mws_vpc_endpoint %s because it doesn't match %s", ep.VpcEndpointName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_vpc_endpoint",
			ID:       ic.accountClient.Config.AccountID + "/" + ep.VpcEndpointId,
		})
	}
	return nil

}

func listMwsPrivateAccessSettings(ic *importContext) error {
	if ic.accountClient.Config.IsAzure() {
		return nil
	}
	pss, err := ic.accountClient.PrivateAccess.List(ic.Context)
	if err != nil {
		return err
	}
	for _, ps := range pss {
		if !ic.MatchesName(ps.PrivateAccessSettingsName) {
			log.Printf("[INFO] Skipping mws_private_access_settings %s because it doesn't match %s", ps.PrivateAccessSettingsName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_private_access_settings",
			ID:       ic.accountClient.Config.AccountID + "/" + ps.PrivateAccessSettingsId,
			Name:     ps.PrivateAccessSettingsName,
		})
	}
	return nil
}

func importMwsPrivateAccessSettings(ic *importContext, r *resource) error {
	var pas provisioning.PrivateAccessSettings
	s := ic.Resources["databricks_mws_private_access_settings"].Schema
	common.DataToStructPointer(r.Data, s, &pas)
	for _, ep := range pas.AllowedVpcEndpointIds {
		ic.Emit(&resource{
			Resource: "databricks_mws_vpc_endpoint",
			ID:       ic.accountClient.Config.AccountID + "/" + ep,
		})
	}
	return nil
}

func listMwsCustomerManagedKeys(ic *importContext) error {
	if ic.accountClient.Config.IsAzure() {
		return nil
	}
	kms, err := ic.accountClient.EncryptionKeys.List(ic.Context)
	if err != nil {
		return err
	}
	updatedSinceMs := ic.getUpdatedSinceMs()
	for _, kms := range kms {
		if ic.incremental && kms.CreationTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping mws_customer_managed_keys '%s' that was created at %d (last active=%d)",
				kms.CustomerManagedKeyId, kms.CreationTime, updatedSinceMs)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_customer_managed_keys",
			ID:       ic.accountClient.Config.AccountID + "/" + kms.CustomerManagedKeyId,
			Name:     kms.CustomerManagedKeyId,
		})
	}
	return nil
}

func listMwsNetworks(ic *importContext) error {
	if ic.accountClient.Config.IsAzure() {
		return nil
	}
	networks, err := ic.accountClient.Networks.List(ic.Context)
	if err != nil {
		return err
	}
	for _, network := range networks {
		if !ic.MatchesName(network.NetworkName) {
			log.Printf("[INFO] Skipping mws_networks %s because it doesn't match %s", network.NetworkName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_mws_networks",
			ID:       ic.accountClient.Config.AccountID + "/" + network.NetworkId,
			Name:     network.NetworkName,
		})
	}
	return nil
}

func importMwsNetworks(ic *importContext, r *resource) error {
	var network mws.Network
	s := ic.Resources["databricks_mws_networks"].Schema
	common.DataToStructPointer(r.Data, s, &network)
	if network.VPCEndpoints != nil {
		for _, vpce := range network.VPCEndpoints.DataplaneRelayAPI {
			log.Printf("[DEBUG] emitting dataplane relay vpc endpoint %s", vpce)
			ic.Emit(&resource{
				Resource: "databricks_mws_vpc_endpoint",
				ID:       ic.accountClient.Config.AccountID + "/" + vpce,
			})
		}
		for _, vpce := range network.VPCEndpoints.RestAPI {
			log.Printf("[DEBUG] emitting rest api vpc endpoint %s", vpce)
			ic.Emit(&resource{
				Resource: "databricks_mws_vpc_endpoint",
				ID:       ic.accountClient.Config.AccountID + "/" + vpce,
			})
		}
	}
	return nil
}

func listMwsWorkspaces(ic *importContext) error {
	if ic.accountClient.Config.IsAzure() {
		// TODO: use listing on Azure just to emit the NCC bindings
		return nil
	}
	workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
	if err != nil {
		return err
	}
	updatedSinceMs := ic.getUpdatedSinceMs()
	for _, workspace := range workspaces {
		if !ic.MatchesName(workspace.WorkspaceName) {
			log.Printf("[INFO] Skipping mws_workspaces %s because it doesn't match %s", workspace.WorkspaceName, ic.match)
			continue
		}
		if ic.incremental && workspace.CreationTime < updatedSinceMs {
			log.Printf("[DEBUG] skipping mws_workspaces '%s' that was created at %d (last active=%d)",
				workspace.WorkspaceName, workspace.CreationTime, updatedSinceMs)
			continue
		}
		if workspace.WorkspaceStatus != "RUNNING" {
			log.Printf("[DEBUG] skipping mws_workspaces '%s' that is not running", workspace.WorkspaceName)
			continue
		}
		wsIdString := strconv.FormatInt(workspace.WorkspaceId, 10)
		ic.Emit(&resource{
			Resource: "databricks_mws_workspaces",
			ID:       ic.accountClient.Config.AccountID + "/" + wsIdString,
			Name:     workspace.WorkspaceName + "_" + wsIdString,
		})
	}
	return nil
}

func importMwsWorkspaces(ic *importContext, r *resource) error {
	var workspace mws.Workspace
	s := ic.Resources["databricks_mws_workspaces"].Schema
	common.DataToStructPointer(r.Data, s, &workspace)
	if workspace.NetworkID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_networks",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.NetworkID,
		})
	}
	if workspace.PrivateAccessSettingsID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_private_access_settings",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.PrivateAccessSettingsID,
		})
	}
	if workspace.StorageConfigurationID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_storage_configurations",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.StorageConfigurationID,
		})
	}
	if workspace.StorageCustomerManagedKeyID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_customer_managed_keys",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.CustomerManagedKeyID,
		})
	}
	if workspace.ManagedServicesCustomerManagedKeyID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_customer_managed_keys",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.ManagedServicesCustomerManagedKeyID,
		})
	}
	if workspace.CredentialsID != "" {
		ic.Emit(&resource{
			Resource: "databricks_mws_credentials",
			ID:       ic.accountClient.Config.AccountID + "/" + workspace.CredentialsID,
		})
	}
	if workspace.NetworkConnectivityConfigID != "" {
		ic.emitNccBindingAndNcc(workspace.WorkspaceID, workspace.NetworkConnectivityConfigID)
	}
	if ic.isServiceEnabled("idfed") {
		err := emitIdfedAndUsersSpsGroups(ic, workspace.WorkspaceID)
		if err != nil {
			log.Printf("[ERROR] listing workspace permission assignments for workspace %d: %s",
				workspace.WorkspaceID, err.Error())
		}
	}
	if ic.isServiceEnabled("seg") {
		ic.Emit(&resource{
			Resource: "databricks_workspace_network_option",
			ID:       strconv.FormatInt(workspace.WorkspaceID, 10),
		})
	}
	return nil
}
