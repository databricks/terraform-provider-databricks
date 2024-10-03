package settings

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceNotificationDestinations() common.Resource {
	// dest
	type notificationDestination struct {
		Id              string `json:"id,omitempty" tf:"computed"`
		DisplayName     string `json:"display_name,omitempty" tf:"computed"`
		DestinationType string `json:"destination_type,omitempty" tf:"computed"`
	}

	type notificationDestinationsData struct {
		DisplayNameContains      string                    `json:"display_name_contains,omitempty"`
		Type                     string                    `json:"type,omitempty"`
		NotificationDestinations []notificationDestination `json:"notification_destinations,omitempty" tf:"computed"`
	}

	return common.WorkspaceData(func(ctx context.Context, data *notificationDestinationsData, w *databricks.WorkspaceClient) error {

		if data.Type != "" {
			switch data.Type {
			case
				string(settings.DestinationTypeEmail),
				string(settings.DestinationTypeMicrosoftTeams),
				string(settings.DestinationTypePagerduty),
				string(settings.DestinationTypeSlack),
				string(settings.DestinationTypeWebhook):
			default:
				return fmt.Errorf("invalid type '%s'; valid types are EMAIL, MICROSOFT_TEAMS, PAGERDUTY, SLACK, WEBHOOK", data.Type)
			}
		}

		notificationRequest := settings.ListNotificationDestinationsRequest{}

		notificationDestinations, err := w.NotificationDestinations.ListAll(ctx, notificationRequest)

		if err != nil {
			return err
		}

		var filteredDestinations []notificationDestination

		for _, nd := range notificationDestinations {
			matches := true

			if data.DisplayNameContains != "" {
				if !strings.Contains(nd.DisplayName, data.DisplayNameContains) {
					matches = false
				}
			}

			if data.Type != "" {
				if string(nd.DestinationType) != data.Type {
					matches = false
				}
			}

			if matches {
				filteredDestinations = append(filteredDestinations, notificationDestination{
					Id:              nd.Id,
					DisplayName:     nd.DisplayName,
					DestinationType: string(nd.DestinationType),
				})
			}

		}

		if len(filteredDestinations) == 0 {
			return fmt.Errorf("could not find any notification destinations with the specified criteria")
		}

		data.NotificationDestinations = filteredDestinations

		return nil
	})
}
