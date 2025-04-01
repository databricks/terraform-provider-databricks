package jobs

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceJobUpdate_WebhookNotifications(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.2/jobs/reset",
				ExpectedRequest: UpdateJobRequest{
					JobID: 789,
					NewSettings: &JobSettings{
						Name: "Webhook test",
						Tasks: []JobTaskSettings{
							{
								TaskKey:           "task1",
								ExistingClusterID: "abc",
							},
						},
						WebhookNotifications: &jobs.WebhookNotifications{
							OnSuccess: []jobs.Webhook{
								{Id: "id1"},
							},
						},
						MaxConcurrentRuns: 1,
					},
				},
				Response: Job{
					JobID: 789,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=789",
				Response: Job{
					Settings: &JobSettings{
						Name: "Webhook test",
						Tasks: []JobTaskSettings{
							{
								TaskKey:           "task1",
								ExistingClusterID: "abc",
							},
						},
						WebhookNotifications: &jobs.WebhookNotifications{
							OnSuccess: []jobs.Webhook{
								{Id: "id1"},
							},
						},
						MaxConcurrentRuns: 1,
					},
				},
			},
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
	}.ApplyNoError(t)
}
