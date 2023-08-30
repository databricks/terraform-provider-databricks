package mlflow

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/databricks-sdk-go/service/ml"
)

var (
	testWhJob = ml.CreateRegistryWebhook{
		Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
		Description: "Job webhook trigger",
		Status:      "ACTIVE",
		JobSpec: &ml.JobSpec{
			JobId:        "1234",
			WorkspaceUrl: "https://test.cloud.databricks.com",
			AccessToken:  "dapi1234",
		},
	}
	testWhID            = "12345"
	testWhJobWithIDResp = ml.RegistryWebhook{
		Id:          testWhID,
		Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
		Description: "Job webhook trigger",
		Status:      "ACTIVE",
		JobSpec: &ml.JobSpecWithoutSecret{
			JobId:        "1234",
			WorkspaceUrl: "https://test.cloud.databricks.com",
		},
	}
	testWhJobWithID = ml.UpdateRegistryWebhook{
		Id:          testWhID,
		Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
		Description: "Job webhook trigger",
		Status:      "ACTIVE",
		JobSpec: &ml.JobSpec{
			JobId:        "1234",
			WorkspaceUrl: "https://test.cloud.databricks.com",
			AccessToken:  "dapi1234",
		},
	}
	testWhHCL = `
	events = ["TRANSITION_REQUEST_CREATED"]
	description = "Job webhook trigger"
	status = "ACTIVE"
	job_spec {
		job_id = "1234"
		workspace_url = "https://test.cloud.databricks.com"
		access_token = "dapi1234"
	}
	`
	listFixture = qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/mlflow/registry-webhooks/list?",
		Response: ml.ListRegistryWebhooks{
			Webhooks: []ml.RegistryWebhook{testWhJobWithIDResp},
		},
	}
)

func TestWebookCreateJobSpec(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: testWhJob,
				Response: ml.CreateWebhookResponse{
					Webhook: &testWhJobWithIDResp,
				},
			},
			listFixture,
		},
		Resource: ResourceMlflowWebhook(),
		Create:   true,
		HCL:      testWhHCL,
	}.ApplyAndExpectData(t, map[string]any{"id": testWhID, "status": "ACTIVE", "job_spec.0.access_token": "dapi1234"})
}

func TestWebookCreateUrlSpec(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: ml.CreateRegistryWebhook{
					Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
					Description: "Url webhook trigger",
					Status:      "ACTIVE",
					HttpUrlSpec: &ml.HttpUrlSpec{
						Url:                   "https://my_cool_host/webhook",
						EnableSslVerification: true,
						Authorization:         "Bearer dapi...",
					},
				},
				Response: ml.CreateWebhookResponse{
					Webhook: &ml.RegistryWebhook{
						Id:          testWhID,
						Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
						Description: "Url webhook trigger",
						Status:      "ACTIVE",
						HttpUrlSpec: &ml.HttpUrlSpecWithoutSecret{
							Url:                   "https://my_cool_host/webhook",
							EnableSslVerification: true,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list?",
				Response: ml.ListRegistryWebhooks{
					Webhooks: []ml.RegistryWebhook{
						{
							Id:          testWhID,
							Events:      []ml.RegistryWebhookEvent{"TRANSITION_REQUEST_CREATED"},
							Description: "Url webhook trigger",
							Status:      "ACTIVE",
							HttpUrlSpec: &ml.HttpUrlSpecWithoutSecret{
								Url:                   "https://my_cool_host/webhook",
								EnableSslVerification: true,
							},
						},
					},
				},
			},
		},
		Resource: ResourceMlflowWebhook(),
		Create:   true,
		HCL: `
		events = ["TRANSITION_REQUEST_CREATED"]
		description = "Url webhook trigger"
		status = "ACTIVE"
		http_url_spec {
			url = "https://my_cool_host/webhook"
			authorization = "Bearer dapi..."
		}
		`,
	}.ApplyAndExpectData(t, map[string]any{"id": testWhID, "status": "ACTIVE", "http_url_spec.0.authorization": "Bearer dapi..."})
}

func TestWebookCreateError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: testWhJob,
				Response: ml.ListRegistryWebhooks{
					Webhooks: []ml.RegistryWebhook{testWhJobWithIDResp},
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowWebhook(),
		Create:   true,
		HCL:      testWhHCL,
	}.ExpectError(t, "error creating a webhook: ")
}

func TestWebookCreateErrorNoUrlOrJobSpec(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMlflowWebhook(),
		Create:   true,
		HCL: `
		events = ["TRANSITION_REQUEST_CREATED"]
		description = "Job webhook trigger"
		status = "ACTIVE"
		`,
	}.ExpectError(t, "at least one of http_url_spec or job_spec should be specified")
}

func TestWebookRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			listFixture,
		},
		Resource: ResourceMlflowWebhook(),
		Read:     true,
		ID:       testWhID,
	}.ApplyAndExpectData(t, map[string]any{"id": testWhID})
}

func TestWebookReadError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			listFixture,
		},
		Resource: ResourceMlflowWebhook(),
		Read:     true,
		ID:       "123456",
	}.ExpectError(t, "webhook with ID 123456 isn't found")
}

func TestWebookReadErrorBadResponse(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list?",
				Response: ml.ListRegistryWebhooks{
					Webhooks: []ml.RegistryWebhook{testWhJobWithIDResp},
				},
				Status: 400,
			},
		},
		Resource: ResourceMlflowWebhook(),
		Read:     true,
		ID:       "123456",
	}.ExpectError(t, "error reading list of webhooks: ")
}

func TestWebookUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/mlflow/registry-webhooks/update",
				ExpectedRequest: testWhJobWithID,
				Response:        testWhJobWithIDResp,
			},
			listFixture,
		},
		Resource:    ResourceMlflowWebhook(),
		Update:      true,
		RequiresNew: false,
		ID:          testWhID,
		State: map[string]any{
			"status":      "TEST_MODE",
			"events":      []string{"TRANSITION_REQUEST_CREATED"},
			"description": "Job webhook trigger",
			"job_spec": map[string]any{
				"job_id":        "1234",
				"workspace_url": "https://test.cloud.databricks.com",
				"access_token":  "dapi1234",
			},
		},
		HCL: testWhHCL,
	}.ApplyAndExpectData(t, map[string]any{"id": testWhID, "status": "ACTIVE"})
}

func TestWebookDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: fmt.Sprintf("/api/2.0/mlflow/registry-webhooks/delete?id=%s", testWhID),
			},
		},
		Resource: ResourceMlflowWebhook(),
		Delete:   true,
		ID:       testWhID,
		HCL:      testWhHCL,
	}.ApplyAndExpectData(t, map[string]any{"id": testWhID})
}

func TestWebookDeleteError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: fmt.Sprintf("/api/2.0/mlflow/registry-webhooks/delete?id=%s", testWhID),
				Status:   400,
			},
		},
		Resource: ResourceMlflowWebhook(),
		Delete:   true,
		ID:       testWhID,
		HCL:      testWhHCL,
	}.ExpectError(t, "Response from server (400 Bad Request) : unexpected end of JSON input")
}
