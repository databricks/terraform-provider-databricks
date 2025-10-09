package jobs

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestResourceJobUpdate_WebhookNotifications(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockJobsAPI().EXPECT()
			e.Reset(mock.Anything, jobs.ResetJob{
				JobId: 789,
				NewSettings: jobs.JobSettings{
					Name: "Webhook test",
					Tasks: []jobs.Task{
						{
							TaskKey:           "task1",
							ExistingClusterId: "abc",
						},
					},
					WebhookNotifications: &jobs.WebhookNotifications{
						OnSuccess: []jobs.Webhook{
							{Id: "id1"},
						},
					},
					MaxConcurrentRuns: 1,
					Queue: &jobs.QueueSettings{
						Enabled: false,
					},
				},
			}).Return(nil)
			e.Get(mock.Anything, jobs.GetJobRequest{
				JobId: 789,
			}).Return(&jobs.Job{
				JobId: 789,
				Settings: &jobs.JobSettings{
					Name: "Webhook test",
					Tasks: []jobs.Task{
						{
							TaskKey:           "task1",
							ExistingClusterId: "abc",
						},
					},
					WebhookNotifications: &jobs.WebhookNotifications{
						OnSuccess: []jobs.Webhook{
							{Id: "id1"},
						},
					},
					MaxConcurrentRuns: 1,
				},
			}, nil)
		},
		ID:       "789",
		Update:   true,
		Resource: ResourceJob(),
		InstanceState: map[string]string{
			"webhook_notifications.#": "1",
			"webhook_notifications.0.on_duration_warning_threshold_exceeded.#":    "1",
			"webhook_notifications.0.on_duration_warning_threshold_exceeded.0.id": "id1",
			"webhook_notifications.0.on_failure.#":                                "1",
			"webhook_notifications.0.on_failure.0.id":                             "id1",
			"webhook_notifications.0.on_start.#":                                  "1",
			"webhook_notifications.0.on_start.0.id":                               "id1",
			"webhook_notifications.0.on_success.#":                                "1",
			"webhook_notifications.0.on_success.0.id":                             "id1",
		},
		HCL: `
		name = "Webhook test"
		task {
			task_key = "task1"
			existing_cluster_id = "abc"
		}
		webhook_notifications {
			// Remove everything but "on_success"
			on_success {
				id = "id1"
			}
		}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":   "789",
		"name": "Webhook test",
	})
}
