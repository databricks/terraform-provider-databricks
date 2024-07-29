package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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

	// Sensitive fields
	// s.SchemaPath("config", "generic_webhook", "password").SetSensitive()
	// s.SchemaPath("config", "generic_webhook", "username").SetSensitive()
	// s.SchemaPath("config", "generic_webhook", "url").SetSensitive()
	// s.SchemaPath("config", "microsoft_teams", "url").SetSensitive()
	// s.SchemaPath("config", "pagerduty", "integration_key").SetSensitive()
	// s.SchemaPath("config", "slack", "url").SetSensitive()

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
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var s settings.NotificationDestination
			common.DataToStructPointer(d, ndSchema, &s)

			readND, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: d.Id(),
			})
			if err != nil {
				return err
			}

			err = common.StructToData(readND, ndSchema, d)
			if err != nil {
				return err
			}
			d.Set("config.0.slack.0.url", s.Config.Slack.Url)
			// fmt.Println(d.Get("config.0.slack.0.url"))
			return nil
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
