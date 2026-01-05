package exporter

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/common"
	account_setting_v2_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_setting_v2"
	workspace_setting_v2_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/workspace_setting_v2"
	tf_settings "github.com/databricks/terraform-provider-databricks/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func importWorkspaceSettingV2(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex)
	copyEffectiveFieldsToInputFieldsWithConverters[workspace_setting_v2_resource.Setting](
		ic, r, settingsv2.Setting{})

	return nil
}

func listWorkspaceSettingsV2(ic *importContext) error {
	settings, err := ic.workspaceClient.WorkspaceSettingsV2.ListWorkspaceSettingsMetadataAll(ic.Context, settingsv2.ListWorkspaceSettingsMetadataRequest{})
	if err != nil {
		return err
	}
	for _, setting := range settings {
		ic.Emit(&resource{
			Resource: "databricks_workspace_setting_v2",
			ID:       setting.Name,
			Name:     setting.Name,
		})
	}
	return nil
}

func importAccountSettingV2(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex)
	copyEffectiveFieldsToInputFieldsWithConverters[account_setting_v2_resource.Setting](
		ic, r, settingsv2.Setting{})

	return nil
}

func listAccountSettingsV2(ic *importContext) error {
	settings, err := ic.accountClient.SettingsV2.ListAccountSettingsMetadataAll(ic.Context, settingsv2.ListAccountSettingsMetadataRequest{})
	if err != nil {
		return err
	}
	for _, setting := range settings {
		ic.Emit(&resource{
			Resource: "databricks_account_setting_v2",
			ID:       setting.Name,
			Name:     setting.Name,
		})
	}
	return nil
}

func listNotificationDestinations(ic *importContext) error {
	if !ic.meAdmin {
		return fmt.Errorf("notifications can be imported only by admin")
	}
	it := ic.workspaceClient.NotificationDestinations.List(ic.Context, settings.ListNotificationDestinationsRequest{})
	for it.HasNext(ic.Context) {
		n, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.Emit(&resource{
			Resource: "databricks_notification_destination",
			ID:       n.Id,
		})
	}
	return nil
}

func importNotificationDestination(ic *importContext, r *resource) error {
	var notificationDestination tf_settings.NDStruct
	s := ic.Resources["databricks_notification_destination"].Schema
	common.DataToStructPointer(r.Data, s, &notificationDestination)
	if notificationDestination.DestinationType == "EMAIL" && notificationDestination.Config != nil &&
		notificationDestination.Config.Email != nil {
		for _, email := range notificationDestination.Config.Email.Addresses {
			ic.emitUserOrServicePrincipal(email)
		}
	}
	return nil
}

func shouldOmitForNotificationDestination(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
	var notificationDestination tf_settings.NDStruct
	s := ic.Resources["databricks_notification_destination"].Schema
	common.DataToStructPointer(d, s, &notificationDestination)
	if notificationDestination.Config != nil {
		switch notificationDestination.DestinationType {
		case "WEBHOOK":
			if notificationDestination.Config.GenericWebhook != nil {
				switch pathString {
				case "config.0.generic_webhook.0.url":
					return !notificationDestination.Config.GenericWebhook.UrlSet
				case "config.0.generic_webhook.0.username":
					return !notificationDestination.Config.GenericWebhook.UsernameSet
				case "config.0.generic_webhook.0.password":
					return !notificationDestination.Config.GenericWebhook.PasswordSet
				}
			}
		case "SLACK":
			if notificationDestination.Config.Slack != nil {
				switch pathString {
				case "config.0.slack.0.url":
					return !notificationDestination.Config.Slack.UrlSet
				case "config.0.slack.0.channel_id":
					return !notificationDestination.Config.Slack.ChannelIdSet
				case "config.0.slack.0.oauth_token":
					return !notificationDestination.Config.Slack.OauthTokenSet
				}
			}
		case "PAGERDUTY":
			if notificationDestination.Config.Pagerduty != nil && pathString == "config.0.pagerduty.0.integration_key" {
				return !notificationDestination.Config.Pagerduty.IntegrationKeySet
			}
		case "MICROSOFT_TEAMS":
			if notificationDestination.Config.MicrosoftTeams != nil {
				switch pathString {
				case "config.0.microsoft_teams.0.url":
					return !notificationDestination.Config.MicrosoftTeams.UrlSet
				case "config.0.microsoft_teams.0.channel_url":
					return !notificationDestination.Config.MicrosoftTeams.ChannelUrlSet
				case "config.0.microsoft_teams.0.auth_secret":
					return !notificationDestination.Config.MicrosoftTeams.AuthSecretSet
				case "config.0.microsoft_teams.0.tenant_id":
					return !notificationDestination.Config.MicrosoftTeams.TenantIdSet
				case "config.0.microsoft_teams.0.app_id":
					return !notificationDestination.Config.MicrosoftTeams.AppIdSet
				}
			}
		}
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
}
