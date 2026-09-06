package exporter

import (
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/common"
	tf_settings "github.com/databricks/terraform-provider-databricks/settings"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func importWorkspaceSettingV2(ic *importContext, r *resource) error {
	copyEffectiveSettingV2Fields(ic, r)
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
	copyEffectiveSettingV2Fields(ic, r)
	return nil
}

func copyEffectiveSettingV2Fields(ic *importContext, r *resource) {
	wrapper, ok := r.DataWrapper.(*PluginFrameworkResourceData)
	if !ok {
		return
	}

	schema := wrapper.GetSchema()
	copiedFields := []string{}
	for _, fieldName := range schema.GetFields() {
		if !strings.HasPrefix(fieldName, "effective_") {
			continue
		}
		inputFieldName := strings.TrimPrefix(fieldName, "effective_")
		if schema.GetField(inputFieldName) == nil {
			continue
		}
		if _, ok := wrapper.GetOk(inputFieldName); ok {
			continue
		}

		var value attr.Value
		diags := wrapper.state.GetAttribute(ic.Context, path.Root(fieldName), &value)
		if diags.HasError() || value == nil || value.IsNull() || value.IsUnknown() {
			continue
		}
		diags = wrapper.state.SetAttribute(ic.Context, path.Root(inputFieldName), value)
		if diags.HasError() {
			log.Printf("[WARN] Failed to copy %s to %s for %s: %v", fieldName, inputFieldName, r.ID, diags)
			continue
		}
		copiedFields = append(copiedFields, fmt.Sprintf("%s->%s", fieldName, inputFieldName))
	}
	if len(copiedFields) > 0 {
		log.Printf("[TRACE] Copied effective setting fields for %s: %s", r.ID, strings.Join(copiedFields, ", "))
	}
}

// shouldGenerateForSettingV2 forces emission of the inner `value` field inside
// the simple `boolean_val`, `string_val`, and `integer_val` blocks. The
// underlying Setting API marks `value` with `omitempty`, so a zero value (false,
// "", 0) is lost during the Go-SDK -> TF-SDK round-trip. Without this override
// an explicitly configured `value = false` would be silently dropped from the
// generated HCL, leaving an empty block (or the whole resource without it).
func shouldGenerateForSettingV2(ic *importContext, pathString string, fieldSchema FieldSchema,
	wrapper ResourceDataWrapper, r *resource) bool {
	switch pathString {
	case "boolean_val.value", "string_val.value", "integer_val.value":
		return true
	}
	return false
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
