package settings

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceNotificationDestinations_DisplayNameContains(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI()
			e.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return([]settings.ListNotificationDestinationsResult{
				{
					Id:              "1",
					DisplayName:     "test destination",
					DestinationType: settings.DestinationTypeSlack,
				},
				{
					Id:              "2",
					DisplayName:     "another destination",
					DestinationType: settings.DestinationTypeEmail,
				},
			}, nil)
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		display_name_contains = "test"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"notification_destinations.#":                  1,
		"notification_destinations.0.id":               "1",
		"notification_destinations.0.display_name":     "test destination",
		"notification_destinations.0.destination_type": "SLACK",
	})
}

func TestDataSourceNotificationDestinations_Type(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI()
			e.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return([]settings.ListNotificationDestinationsResult{
				{
					Id:              "1",
					DisplayName:     "test destination",
					DestinationType: settings.DestinationTypeSlack,
				},
				{
					Id:              "2",
					DisplayName:     "another destination",
					DestinationType: settings.DestinationTypeEmail,
				},
			}, nil)
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		type = "SLACK"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"notification_destinations.#":                  1,
		"notification_destinations.0.id":               "1",
		"notification_destinations.0.display_name":     "test destination",
		"notification_destinations.0.destination_type": "SLACK",
	})
}

func TestDataSourceNotificationDestinations_DisplayNameAndType(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI()
			e.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return([]settings.ListNotificationDestinationsResult{
				{
					Id:              "1",
					DisplayName:     "test destination",
					DestinationType: settings.DestinationTypeSlack,
				},
				{
					Id:              "2",
					DisplayName:     "another test destination",
					DestinationType: settings.DestinationTypeEmail,
				},
				{
					Id:              "3",
					DisplayName:     "third destination",
					DestinationType: settings.DestinationTypeSlack,
				},
			}, nil)
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
			display_name_contains = "test"
			type = "SLACK"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"notification_destinations.#":                  1,
		"notification_destinations.0.id":               "1",
		"notification_destinations.0.display_name":     "test destination",
		"notification_destinations.0.destination_type": "SLACK",
	})
}

func TestDataSourceNotificationDestinations_InvalidType(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
			type = "INVALID"
		`,
	}.ExpectError(t, "invalid type 'INVALID'; valid types are EMAIL, MICROSOFT_TEAMS, PAGERDUTY, SLACK, WEBHOOK")
}

func TestDataSourceNotificationDestinations_NoMatches(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI()
			e.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return([]settings.ListNotificationDestinationsResult{
				{
					Id:              "1",
					DisplayName:     "test destination",
					DestinationType: settings.DestinationTypeSlack,
				},
				{
					Id:              "2",
					DisplayName:     "another destination",
					DestinationType: settings.DestinationTypeEmail,
				},
			}, nil)
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
			display_name_contains = "invalid"
		`,
	}.ExpectError(t, "could not find any notification destinations with the specified criteria")
}

func TestDataSourceNotificationDestinations_APIError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI()
			e.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return(nil, fmt.Errorf("api error"))
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "api error")
}
