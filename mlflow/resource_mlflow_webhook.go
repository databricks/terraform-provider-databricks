package mlflow

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type CreateRegistryWebhook struct {
	ml.CreateRegistryWebhook
	common.Namespace
}

func readWebHook(w *databricks.WorkspaceClient, ctx context.Context, ID string) (ml.RegistryWebhook, error) {
	m, err := w.ModelRegistry.ListWebhooksAll(ctx, ml.ListWebhooksRequest{})
	if err != nil {
		return ml.RegistryWebhook{}, fmt.Errorf("error reading list of webhooks: %w", err)
	}

	for _, wh := range m {
		if wh.Id == ID {
			return wh, nil
		}
	}
	return ml.RegistryWebhook{}, fmt.Errorf("webhook with ID %s isn't found", ID)
}

func ResourceMlflowWebhook() common.Resource {
	s := common.StructToSchema(
		CreateRegistryWebhook{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["status"].ValidateFunc = validation.StringInSlice([]string{"ACTIVE", "TEST_MODE", "DISABLED"}, true)
			if p, err := common.SchemaPath(m, "http_url_spec", "url"); err == nil {
				p.ValidateFunc = validation.IsURLWithHTTPS
			}
			m["http_url_spec"].ConflictsWith = []string{"job_spec"}
			m["job_spec"].ConflictsWith = []string{"http_url_spec"}
			common.MustSchemaPath(m, "http_url_spec", "enable_ssl_verification").Default = true
			common.MustSchemaPath(m, "http_url_spec", "secret").Sensitive = true
			common.MustSchemaPath(m, "job_spec", "access_token").Sensitive = true
			common.NamespaceCustomizeSchemaMap(m)
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var m ml.CreateRegistryWebhook
			common.DataToStructPointer(d, s, &m)
			if m.HttpUrlSpec == nil && m.JobSpec == nil {
				return fmt.Errorf("at least one of http_url_spec or job_spec should be specified")
			}

			resp, err := w.ModelRegistry.CreateWebhook(ctx, m)
			if err != nil {
				return fmt.Errorf("error creating a webhook: %w", err)
			}
			d.SetId(resp.Webhook.Id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var mOrig ml.CreateRegistryWebhook
			common.DataToStructPointer(d, s, &mOrig)
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			m, err := readWebHook(w, ctx, d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(m, s, d)
			if err != nil {
				return err
			}
			// We need to preserve original values as API doesn't return sensitive values
			if mOrig.JobSpec != nil && m.JobSpec != nil {
				d.Set("job_spec.0.access_token", mOrig.JobSpec.AccessToken)
			}
			if mOrig.HttpUrlSpec != nil && m.HttpUrlSpec != nil {
				d.Set("http_url_spec.0.authorization", mOrig.HttpUrlSpec.Authorization)
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var m ml.UpdateRegistryWebhook
			common.DataToStructPointer(d, s, &m)
			m.Id = d.Id()
			_, err = w.ModelRegistry.UpdateWebhook(ctx, m)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.ModelRegistry.DeleteWebhook(ctx, ml.DeleteWebhookRequest{Id: d.Id()})
		},
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
	}
}
