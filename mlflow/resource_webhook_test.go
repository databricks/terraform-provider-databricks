package mlflow

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

var (
	testWhJob = Webhook{
		Events:      []string{"TRANSITION_REQUEST_CREATED"},
		Description: "Job webhook trigger",
		Status:      "ACTIVE",
		JobSpec: &JobSpec{
			JobID:        "1234",
			WorkspaceURL: "https://test.cloud.databricks.com",
			AccessToken:  "dapi1234",
		},
	}
	testWhID        = "12345"
	testWhJobWithID = Webhook{
		ID:          testWhID,
		Events:      []string{"TRANSITION_REQUEST_CREATED"},
		Description: "Job webhook trigger",
		Status:      "ACTIVE",
		JobSpec: &JobSpec{
			JobID:        "1234",
			WorkspaceURL: "https://test.cloud.databricks.com",
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
)

func TestWebookCreateJobSpec(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: testWhJob,
				Response: webhookApiResponse{
					Webhook: testWhJobWithID,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{testWhJobWithID},
				},
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Create:   true,
		HCL:      testWhHCL,
	}.ApplyAndExpectData(t, map[string]interface{}{"id": testWhID, "status": "ACTIVE"})
}

func TestWebookCreateUrlSpec(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: Webhook{
					Events:      []string{"TRANSITION_REQUEST_CREATED"},
					Description: "Url webhook trigger",
					Status:      "ACTIVE",
					HttpUrlSpec: &HttpUrlSpec{
						URL:                   "https://my_cool_host/webhook",
						EnableSslVerification: true,
						Authorization:         "Bearer dapi...",
					},
				},
				Response: webhookApiResponse{
					Webhook: Webhook{
						ID:          testWhID,
						Events:      []string{"TRANSITION_REQUEST_CREATED"},
						Description: "Url webhook trigger",
						Status:      "ACTIVE",
						HttpUrlSpec: &HttpUrlSpec{
							URL:                   "https://my_cool_host/webhook",
							EnableSslVerification: true},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{
						{
							ID:          testWhID,
							Events:      []string{"TRANSITION_REQUEST_CREATED"},
							Description: "Url webhook trigger",
							Status:      "ACTIVE",
							HttpUrlSpec: &HttpUrlSpec{
								URL:                   "https://my_cool_host/webhook",
								EnableSslVerification: true,
							},
						},
					},
				},
			},
		},
		Resource: ResourceMLFlowWebhook(),
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
	}.ApplyAndExpectData(t, map[string]interface{}{"id": testWhID, "status": "ACTIVE"})
}

func TestWebookCreateError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registry-webhooks/create",
				ExpectedRequest: testWhJob,
				Response: webhookApiResponse{
					Webhook: testWhJobWithID,
				},
				Status: 400,
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Create:   true,
		HCL:      testWhHCL,
	}.ExpectError(t, "error creating a webhook: ")
}

func TestWebookCreateErrorNoUrlOrJobSpec(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMLFlowWebhook(),
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
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{testWhJobWithID},
				},
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Read:     true,
		ID:       testWhID,
	}.ApplyAndExpectData(t, map[string]interface{}{"id": testWhID})
}

func TestWebookReadError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{testWhJobWithID},
				},
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Read:     true,
		ID:       "123456",
	}.ExpectError(t, "Webhook with ID 123456 isn't found")
}

func TestWebookReadErrorBadResponse(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{testWhJobWithID},
				},
				Status: 400,
			},
		},
		Resource: ResourceMLFlowWebhook(),
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
				Response: webhookApiResponse{
					Webhook: testWhJobWithID,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registry-webhooks/list",
				Response: webhookListResponse{
					Webhooks: []Webhook{testWhJobWithID},
				},
			},
		},
		Resource:    ResourceMLFlowWebhook(),
		Update:      true,
		RequiresNew: false,
		ID:          testWhID,
		State: map[string]interface{}{
			"status":      "TEST_MODE",
			"events":      []string{"TRANSITION_REQUEST_CREATED"},
			"description": "Job webhook trigger",
			"job_spec": map[string]interface{}{
				"job_id":        "1234",
				"workspace_url": "https://test.cloud.databricks.com",
				"access_token":  "dapi1234",
			},
		},
		HCL: testWhHCL,
	}.ApplyAndExpectData(t, map[string]interface{}{"id": testWhID, "status": "ACTIVE"})
}

func TestWebookDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registry-webhooks/delete",
				ExpectedRequest: map[string]interface{}{
					"id": testWhID,
				},
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Delete:   true,
		ID:       testWhID,
		HCL:      testWhHCL,
	}.ApplyAndExpectData(t, map[string]interface{}{"id": testWhID})
}

func TestWebookDeleteError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registry-webhooks/delete",
				ExpectedRequest: map[string]interface{}{
					"id": testWhID,
				},
				Status: 400,
			},
		},
		Resource: ResourceMLFlowWebhook(),
		Delete:   true,
		ID:       testWhID,
		HCL:      testWhHCL,
	}.ExpectError(t, "Response from server (400 Bad Request) : unexpected end of JSON input")
}
