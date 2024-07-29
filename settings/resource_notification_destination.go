package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setState(d *schema.ResourceData) {
	d.Set("slack_url", d.Get("config.0.slack.0.url"))
	d.Set("microsoft_teams_url", d.Get("config.0.microsoft_teams.0.url"))
	d.Set("pagerduty_integration_key", d.Get("config.0.pagerduty.0.integration_key"))
	d.Set("generic_webhook_url", d.Get("config.0.generic_webhook.0.url"))
	d.Set("generic_webhook_password", d.Get("config.0.generic_webhook.0.password"))
	d.Set("generic_webhook_username", d.Get("config.0.generic_webhook.0.username"))
}

type NDStruct struct {
	SlackUrl                string `json:"slack_url,omitempty"`
	MicrosoftTeamsUrl       string `json:"microsoft_teams_url,omitempty"`
	PagerDutyIntegrationKey string `json:"pagerduty_integration_key,omitempty"`
	GenericWebhookUrl       string `json:"generic_webhook_url,omitempty"`
	GenericWebhookPassword  string `json:"generic_webhook_password,omitempty"`
	GenericWebhookUsername  string `json:"generic_webhook_username,omitempty"`
	settings.NotificationDestination
}

func customDiffSlackUrl(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("slack_url") == new
}

func customDiffMicrosoftTeamsUrl(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("microsoft_teams_url") == new
}

func customDiffPagerdutyIntegrationKey(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("pagerduty_integration_key") == new
}

func customDiffGenericWebhookUrl(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("generic_webhook_url") == new
}

func customDiffGenericWebhookUsername(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("generic_webhook_username") == new
}

func customDiffGenericWebhookPassword(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("generic_webhook_password") == new
}

func (NDStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields
	s.SchemaPath("display_name").SetRequired()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("destination_type").SetComputed()
	s.SchemaPath("config", "slack", "url_set").SetComputed()
	s.SchemaPath("config", "pagerduty", "integration_key_set").SetComputed()
	s.SchemaPath("config", "microsoft_teams", "url_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "url_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "password_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "username_set").SetComputed()

	s.SchemaPath("slack_url").SetComputed()
	s.SchemaPath("microsoft_teams_url").SetComputed()
	s.SchemaPath("pagerduty_integration_key").SetComputed()
	s.SchemaPath("generic_webhook_url").SetComputed()
	s.SchemaPath("generic_webhook_password").SetComputed()
	s.SchemaPath("generic_webhook_username").SetComputed()

	// ForceNew fields
	s.SchemaPath("destination_type").SetForceNew()

	// ConflictsWith fields
	config_eoo := []string{"config.0.slack", "config.0.pagerduty", "config.0.microsoft_teams", "config.0.generic_webhook", "config.0.email"}
	s.SchemaPath("config", "slack").SetExactlyOneOf(config_eoo)
	// s.SchemaPath("config", "pagerduty").SetExactlyOneOf(config_eoo)
	// s.SchemaPath("config", "microsoft_teams").SetExactlyOneOf(config_eoo)
	// s.SchemaPath("config", "generic_webhook").SetExactlyOneOf(config_eoo)
	// s.SchemaPath("config", "email").SetExactlyOneOf(config_eoo)

	// RequiredWith fields
	s.SchemaPath("config", "slack").SetRequiredWith([]string{"config.0.slack.0.url"})
	s.SchemaPath("config", "pagerduty").SetRequiredWith([]string{"config.0.pagerduty.0.integration_key"})
	s.SchemaPath("config", "microsoft_teams").SetRequiredWith([]string{"config.0.microsoft_teams.0.url"})
	s.SchemaPath("config", "generic_webhook").SetRequiredWith([]string{"config.0.generic_webhook.0.url"})
	s.SchemaPath("config", "email").SetRequiredWith([]string{"config.0.email.0.addresses"})

	// CustomSuppressDiff fields
	s.SchemaPath("config", "slack", "url").SetCustomSuppressDiff(customDiffSlackUrl)
	s.SchemaPath("config", "microsoft_teams", "url").SetCustomSuppressDiff(customDiffMicrosoftTeamsUrl)
	s.SchemaPath("config", "pagerduty", "integration_key").SetCustomSuppressDiff(customDiffPagerdutyIntegrationKey)
	s.SchemaPath("config", "generic_webhook", "url").SetCustomSuppressDiff(customDiffGenericWebhookUrl)
	s.SchemaPath("config", "generic_webhook", "username").SetCustomSuppressDiff(customDiffGenericWebhookUsername)
	s.SchemaPath("config", "generic_webhook", "password").SetCustomSuppressDiff(customDiffGenericWebhookPassword)

	// Sensitive fields
	s.SchemaPath("config", "generic_webhook", "password").SetSensitive()
	s.SchemaPath("config", "generic_webhook", "username").SetSensitive()
	s.SchemaPath("config", "generic_webhook", "url").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "url").SetSensitive()
	s.SchemaPath("config", "pagerduty", "integration_key").SetSensitive()
	s.SchemaPath("config", "slack", "url").SetSensitive()

	return s
}

var ndSchema = common.StructToSchema(NDStruct{}, nil)

func ResourceNotificationDestination() common.Resource {
	return common.Resource{
		Schema: ndSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {

			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var newNDrequest settings.CreateNotificationDestinationRequest
			common.DataToStructPointer(d, ndSchema, &newNDrequest)
			createdND, err := w.NotificationDestinations.Create(ctx, newNDrequest)
			if err != nil {
				return err
			}
			d.SetId(createdND.Id)
			setState(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			readND, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: d.Id(),
			})
			if err != nil {
				return err
			}

			return common.StructToData(readND, ndSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateNDRequest settings.UpdateNotificationDestinationRequest
			common.DataToStructPointer(d, ndSchema, &updateNDRequest)
			updateNDRequest.Id = d.Id()
			_, err = w.NotificationDestinations.Update(ctx, updateNDRequest)
			if err != nil {
				return err
			}
			setState(d)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.NotificationDestinations.Delete(ctx, settings.DeleteNotificationDestinationRequest{
				Id: d.Id(),
			})
		},
	}
}
