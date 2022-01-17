package mlflow

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type HttpUrlSpec struct {
	URL                   string `json:"url"`
	EnableSslVerification bool   `json:"enable_ssl_verification" tf:"default:true"`
	Secret                string `json:"string,omitempty"`
	Authorization         string `json:"authorization,omitempty" tf:"sensitive"`
}

type JobSpec struct {
	JobID        string `json:"job_id"`
	AccessToken  string `json:"access_token" tf:"sensitive"`
	WorkspaceURL string `json:"workspace_url,omitempty"`
}

type Webhook struct {
	ID          string       `json:"id" tf:"computed"`
	Events      []string     `json:"events"`
	Status      string       `json:"status,omitempty" tf:"default:ACTIVE"`
	Description string       `json:"description,omitempty"`
	ModelName   string       `json:"model_name,omitempty" tf:"force_new"`
	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"` // TODO: add diff suppressor?
	JobSpec     *JobSpec     `json:"job_spec,omitempty"`      // TODO: add diff suppressor?
}

type webhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type webhookApiResponse struct {
	Webhook Webhook `json:"webhook"`
}

type WebhooksAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewWebhooksAPI(ctx context.Context, m interface{}) WebhooksAPI {
	return WebhooksAPI{m.(*common.DatabricksClient), ctx}
}

func (a WebhooksAPI) Create(m Webhook) (string, error) {
	var r webhookApiResponse
	err := a.client.Post(a.context, "/mlflow/registry-webhooks/create", &m, &r)
	return r.Webhook.ID, err
}

func (a WebhooksAPI) Read(ID string) (Webhook, error) {
	var m webhookListResponse
	err := a.client.Get(a.context, "/mlflow/registry-webhooks/list", nil, &m)
	if err != nil {
		return Webhook{}, fmt.Errorf("error reading list of webhooks: %w", err)
	}
	for _, wh := range m.Webhooks {
		if wh.ID == ID {
			return wh, nil
		}
	}
	return Webhook{}, fmt.Errorf("Webhook with ID %s isn't found", ID)
}

// Update the webhook entity
func (a WebhooksAPI) Update(id string, m Webhook) error {
	// Update API doesn't allow to change Model name field after creation.
	m.ModelName = ""
	m.ID = id
	return a.client.Patch(a.context, "/mlflow/registry-webhooks/update", m)
}

// Delete removes the webhook by its ID
func (a WebhooksAPI) Delete(ID string) error {
	return a.client.Delete(a.context, "/mlflow/registry-webhooks/delete", map[string]string{
		"id": ID,
	})
}

func ResourceMLFlowWebhook() *schema.Resource {
	s := common.StructToSchema(
		Webhook{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["status"].ValidateFunc = validation.StringInSlice([]string{"ACTIVE", "TEST_MODE", "DISABLED"}, true)
			// TODO: do we need a validation for Events?
			delete(m, "id")
			if p, err := common.SchemaPath(m, "http_url_spec", "enable_ssl_verification"); err == nil {
				p.Required = false
				p.Optional = true
			}
			if p, err := common.SchemaPath(m, "http_url_spec", "url"); err == nil {
				p.ValidateFunc = validation.IsURLWithHTTPS
			}
			m["http_url_spec"].ConflictsWith = []string{"job_spec"}
			m["job_spec"].ConflictsWith = []string{"http_url_spec"}

			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var m Webhook
			common.DataToStructPointer(d, s, &m)
			if m.HttpUrlSpec == nil && m.JobSpec == nil {
				return fmt.Errorf("at least one of http_url_spec or job_spec should be specified")
			}

			hookID, err := NewWebhooksAPI(ctx, c).Create(m)
			if err != nil {
				return fmt.Errorf("error creating a webhook: %w", err)
			}
			d.SetId(hookID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var mOrig Webhook
			common.DataToStructPointer(d, s, &mOrig)
			m, err := NewWebhooksAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			// We need to preserve original values as API doesn't return sensitive values
			if mOrig.JobSpec != nil && m.JobSpec != nil {
				m.JobSpec.AccessToken = mOrig.JobSpec.AccessToken
			}
			if mOrig.HttpUrlSpec != nil && m.HttpUrlSpec != nil {
				m.HttpUrlSpec.Authorization = mOrig.HttpUrlSpec.Authorization
			}
			return common.StructToData(m, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var m Webhook
			common.DataToStructPointer(d, s, &m)
			return NewWebhooksAPI(ctx, c).Update(d.Id(), m)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewWebhooksAPI(ctx, c).Delete(d.Id())
		},
		Schema: s,
	}.ToResource()
}
