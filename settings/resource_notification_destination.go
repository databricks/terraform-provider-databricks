package settings

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// writeOnlyFields lists the sensitive config fields that also have a write-only
// variant, `<field>_wo`, keyed by config block. Write-only values are never stored in
// the plan or state, so they're read from the raw configuration on every create and
// update: the API replaces the whole config block on update, so leaving them out
// would clear them.
var writeOnlyFields = map[string][]string{
	"generic_webhook": {"url", "username", "password"},
	"slack":           {"url", "oauth_token"},
	"microsoft_teams": {"url", "auth_secret"},
	"pagerduty":       {"integration_key"},
}

// Config blocks whose field is required, either as the field or its write-only variant.
var requiredConfigFields = map[string]string{
	"slack":           "url",
	"microsoft_teams": "url",
	"generic_webhook": "url",
	"pagerduty":       "integration_key",
}

// getWriteOnlyConfig returns the value of `config.0.<block>.0.<field>_wo` from the raw
// configuration, if it's set.
func getWriteOnlyConfig(d *schema.ResourceData, block, field string) (string, bool) {
	path := cty.GetAttrPath("config").IndexInt(0).GetAttr(block).IndexInt(0).GetAttr(field + "_wo")
	v, diags := d.GetRawConfigAt(path)
	if diags.HasError() || !v.IsKnown() || v.IsNull() || !v.Type().Equals(cty.String) {
		return "", false
	}
	return v.AsString(), true
}

// applyWriteOnlyConfig sets the configured write-only values on the request config.
func applyWriteOnlyConfig(d *schema.ResourceData, config *settings.Config) {
	if config == nil {
		return
	}
	set := func(block, field string, target *string) {
		if v, ok := getWriteOnlyConfig(d, block, field); ok {
			*target = v
		}
	}
	if c := config.GenericWebhook; c != nil {
		set("generic_webhook", "url", &c.Url)
		set("generic_webhook", "username", &c.Username)
		set("generic_webhook", "password", &c.Password)
	}
	if c := config.Slack; c != nil {
		set("slack", "url", &c.Url)
		set("slack", "oauth_token", &c.OauthToken)
	}
	if c := config.MicrosoftTeams; c != nil {
		set("microsoft_teams", "url", &c.Url)
		set("microsoft_teams", "auth_secret", &c.AuthSecret)
	}
	if c := config.Pagerduty; c != nil {
		set("pagerduty", "integration_key", &c.IntegrationKey)
	}
}

func setStruct(s *settings.NotificationDestination, readND *settings.NotificationDestination) {
	if readND.Config != nil && s.Config != nil {
		switch readND.DestinationType {
		case settings.DestinationTypeSlack:
			if readND.Config.Slack != nil && s.Config.Slack != nil {
				if readND.Config.Slack.UrlSet {
					readND.Config.Slack.Url = s.Config.Slack.Url
				}
				if readND.Config.Slack.ChannelIdSet {
					readND.Config.Slack.ChannelId = s.Config.Slack.ChannelId
				}
				if readND.Config.Slack.OauthTokenSet {
					readND.Config.Slack.OauthToken = s.Config.Slack.OauthToken
				}
			}
		case settings.DestinationTypeDatabricksAppSlack:
			if readND.Config.DatabricksAppSlack != nil && s.Config.DatabricksAppSlack != nil {
				if readND.Config.DatabricksAppSlack.ChannelIdSet {
					readND.Config.DatabricksAppSlack.ChannelId = s.Config.DatabricksAppSlack.ChannelId
				}
				if readND.Config.DatabricksAppSlack.TeamIdSet {
					readND.Config.DatabricksAppSlack.TeamId = s.Config.DatabricksAppSlack.TeamId
				}
			}
		case settings.DestinationTypeDatabricksAppTeams:
			if readND.Config.DatabricksAppTeams != nil && s.Config.DatabricksAppTeams != nil {
				if readND.Config.DatabricksAppTeams.ChannelIdSet {
					readND.Config.DatabricksAppTeams.ChannelId = s.Config.DatabricksAppTeams.ChannelId
				}
				if readND.Config.DatabricksAppTeams.TeamIdSet {
					readND.Config.DatabricksAppTeams.TeamId = s.Config.DatabricksAppTeams.TeamId
				}
				if readND.Config.DatabricksAppTeams.TenantIdSet {
					readND.Config.DatabricksAppTeams.TenantId = s.Config.DatabricksAppTeams.TenantId
				}
			}
		case settings.DestinationTypePagerduty:
			if readND.Config.Pagerduty != nil && s.Config.Pagerduty != nil {
				if readND.Config.Pagerduty.IntegrationKeySet {
					readND.Config.Pagerduty.IntegrationKey = s.Config.Pagerduty.IntegrationKey
				}
			}
		case settings.DestinationTypeMicrosoftTeams:
			if readND.Config.MicrosoftTeams != nil && s.Config.MicrosoftTeams != nil {
				if readND.Config.MicrosoftTeams.UrlSet {
					readND.Config.MicrosoftTeams.Url = s.Config.MicrosoftTeams.Url
				}
				if readND.Config.MicrosoftTeams.ChannelUrlSet {
					readND.Config.MicrosoftTeams.ChannelUrl = s.Config.MicrosoftTeams.ChannelUrl
				}
				if readND.Config.MicrosoftTeams.AuthSecretSet {
					readND.Config.MicrosoftTeams.AuthSecret = s.Config.MicrosoftTeams.AuthSecret
				}
				if readND.Config.MicrosoftTeams.TenantIdSet {
					readND.Config.MicrosoftTeams.TenantId = s.Config.MicrosoftTeams.TenantId
				}
				if readND.Config.MicrosoftTeams.AppIdSet {
					readND.Config.MicrosoftTeams.AppId = s.Config.MicrosoftTeams.AppId
				}
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
	applyWriteOnlyConfig(d, newNDrequest.Config)
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
	applyWriteOnlyConfig(d, updateNDRequest.Config)
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
	common.Namespace
}

func (NDStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields
	s.SchemaPath("display_name").SetRequired()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("destination_type").SetComputed()
	s.SchemaPath("config", "slack", "url_set").SetComputed()
	s.SchemaPath("config", "slack", "channel_id_set").SetComputed()
	s.SchemaPath("config", "slack", "oauth_token_set").SetComputed()
	s.SchemaPath("config", "databricks_app_slack", "channel_id_set").SetComputed()
	s.SchemaPath("config", "databricks_app_slack", "team_id_set").SetComputed()
	s.SchemaPath("config", "databricks_app_teams", "channel_id_set").SetComputed()
	s.SchemaPath("config", "databricks_app_teams", "team_id_set").SetComputed()
	s.SchemaPath("config", "databricks_app_teams", "tenant_id_set").SetComputed()
	s.SchemaPath("config", "pagerduty", "integration_key_set").SetComputed()
	s.SchemaPath("config", "microsoft_teams", "url_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "url_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "password_set").SetComputed()
	s.SchemaPath("config", "generic_webhook", "username_set").SetComputed()

	// ForceNew fields
	s.SchemaPath("config", "databricks_app_slack").SetForceNew()
	s.SchemaPath("config", "databricks_app_teams").SetForceNew()
	s.SchemaPath("config", "slack").SetForceNew()
	s.SchemaPath("config", "pagerduty").SetForceNew()
	s.SchemaPath("config", "microsoft_teams").SetForceNew()
	s.SchemaPath("config", "generic_webhook").SetForceNew()
	s.SchemaPath("config", "email").SetForceNew()

	// ConflictsWith fields
	config_eoo := []string{"config.0.slack", "config.0.pagerduty", "config.0.microsoft_teams", "config.0.generic_webhook", "config.0.email", "config.0.databricks_app_slack", "config.0.databricks_app_teams"}
	s.SchemaPath("config", "slack").SetExactlyOneOf(config_eoo)

	// RequiredWith fields
	s.SchemaPath("config", "email").SetRequiredWith([]string{"config.0.email.0.addresses"})

	// Write-only fields
	s.AddNewField("config_wo_version", &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
		Description:  "Increment to apply a change to the write-only `*_wo` config fields.",
	})
	for block, fields := range writeOnlyFields {
		for _, field := range fields {
			path := fmt.Sprintf("config.0.%s.0.%s", block, field)
			s.SchemaPath("config", block).AddNewField(field+"_wo", &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				WriteOnly:     true,
				ConflictsWith: []string{path},
				RequiredWith:  []string{"config_wo_version"},
			})
			s.SchemaPath("config", block, field).SetConflictsWith([]string{path + "_wo"})
		}
	}
	// Required fields, which can be set either directly or write-only
	for block, field := range requiredConfigFields {
		path := fmt.Sprintf("config.0.%s.0.%s", block, field)
		s.SchemaPath("config", block, field).SetExactlyOneOf([]string{path, path + "_wo"})
	}

	// Sensitive fields
	s.SchemaPath("config", "generic_webhook", "password").SetSensitive()
	s.SchemaPath("config", "generic_webhook", "username").SetSensitive()
	s.SchemaPath("config", "generic_webhook", "url").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "url").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "channel_url").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "auth_secret").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "tenant_id").SetSensitive()
	s.SchemaPath("config", "microsoft_teams", "app_id").SetSensitive()
	s.SchemaPath("config", "pagerduty", "integration_key").SetSensitive()
	s.SchemaPath("config", "databricks_app_slack", "channel_id").SetSensitive()
	s.SchemaPath("config", "databricks_app_slack", "team_id").SetSensitive()
	s.SchemaPath("config", "databricks_app_teams", "channel_id").SetSensitive()
	s.SchemaPath("config", "databricks_app_teams", "team_id").SetSensitive()
	s.SchemaPath("config", "databricks_app_teams", "tenant_id").SetSensitive()
	s.SchemaPath("config", "slack", "url").SetSensitive()
	s.SchemaPath("config", "slack", "oauth_token").SetSensitive()
	s.SchemaPath("config", "slack", "channel_id").SetSensitive()

	common.NamespaceCustomizeSchema(s)
	return s
}

var ndSchema = common.StructToSchema(NDStruct{}, nil)

func ResourceNotificationDestination() common.Resource {
	return common.Resource{
		Schema: ndSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {

			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return Create(ctx, d, w)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return Read(ctx, d, w)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return Update(ctx, d, w)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return Delete(ctx, d, w)
		},
	}
}
