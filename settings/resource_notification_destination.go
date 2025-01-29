package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setStruct(s *settings.NotificationDestination, readND *settings.NotificationDestination) {
	if readND.Config != nil && s.Config != nil {
		switch readND.DestinationType {
		case settings.DestinationTypeSlack:
			if readND.Config.Slack != nil && s.Config.Slack != nil {
				readND.Config.Slack.Url = s.Config.Slack.Url
			}
		case settings.DestinationTypePagerduty:
			if readND.Config.Pagerduty != nil && s.Config.Pagerduty != nil {
				readND.Config.Pagerduty.IntegrationKey = s.Config.Pagerduty.IntegrationKey
			}
		case settings.DestinationTypeMicrosoftTeams:
			if readND.Config.MicrosoftTeams != nil && s.Config.MicrosoftTeams != nil {
				readND.Config.MicrosoftTeams.Url = s.Config.MicrosoftTeams.Url
			}
		case settings.DestinationTypeWebhook:
			if readND.Config.GenericWebhook != nil && s.Config.GenericWebhook != nil {
				if readND.Config.GenericWebhook.UrlSet {
					readND.Config.GenericWebhook.Url = s.Config.GenericWebhook.Url
				}
				if readND.Config.GenericWebhook.PasswordSet {
					readND.Config.GenericWebhook.Password = s.Config.GenericWebhook.Password
				}
				if readND.Config.GenericWebhook.UsernameSet {
					readND.Config.GenericWebhook.Username = s.Config.GenericWebhook.Username
				}
			}
		}
	}
}

func Create(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient) error {
	var newNDrequest settings.CreateNotificationDestinationRequest
	common.DataToStructPointer(d, ndSchema, &newNDrequest)
	createdND, err := w.NotificationDestinations.Create(ctx, newNDrequest)
	if err != nil {
		return err
	}
	d.SetId(createdND.Id)
	return nil
}

func Read(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient) error {
	var tempND settings.NotificationDestination
	common.DataToStructPointer(d, ndSchema, &tempND)

	readND, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
		Id: d.Id(),
	})
	if err != nil {
		return err
	}
	setStruct(&tempND, readND)
	return common.StructToData(readND, ndSchema, d)
}

func Update(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient) error {
	var updateNDRequest settings.UpdateNotificationDestinationRequest
	common.DataToStructPointer(d, ndSchema, &updateNDRequest)
	updateNDRequest.Id = d.Id()
	_, err := w.NotificationDestinations.Update(ctx, updateNDRequest)
	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient) error {
	return w.NotificationDestinations.Delete(ctx, settings.DeleteNotificationDestinationRequest{
		Id: d.Id(),
	})
}

type NDStruct struct {
	settings.NotificationDestination
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

	// ForceNew fields
	s.SchemaPath("config", "slack").SetForceNew()
	s.SchemaPath("config", "pagerduty").SetForceNew()
	s.SchemaPath("config", "microsoft_teams").SetForceNew()
	s.SchemaPath("config", "generic_webhook").SetForceNew()
	s.SchemaPath("config", "email").SetForceNew()

	// ConflictsWith fields
	config_eoo := []string{"config.0.slack", "config.0.pagerduty", "config.0.microsoft_teams", "config.0.generic_webhook", "config.0.email"}
	s.SchemaPath("config", "slack").SetExactlyOneOf(config_eoo)

	// RequiredWith fields
	s.SchemaPath("config", "slack").SetRequiredWith([]string{"config.0.slack.0.url"})
	s.SchemaPath("config", "pagerduty").SetRequiredWith([]string{"config.0.pagerduty.0.integration_key"})
	s.SchemaPath("config", "microsoft_teams").SetRequiredWith([]string{"config.0.microsoft_teams.0.url"})
	s.SchemaPath("config", "generic_webhook").SetRequiredWith([]string{"config.0.generic_webhook.0.url"})
	s.SchemaPath("config", "email").SetRequiredWith([]string{"config.0.email.0.addresses"})

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
			return Create(ctx, d, w)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Read(ctx, d, w)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Update(ctx, d, w)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Delete(ctx, d, w)
		},
	}
}
