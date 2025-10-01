package serving

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	defaultPtProvisionTimeout = 10 * time.Minute
)

// ModelServingProvisionedThroughputStruct embeds SDK type with ProviderConfig
type ModelServingProvisionedThroughputStruct struct {
	serving.CreatePtEndpointRequest
	common.ProviderConfig
}

func ResourceModelServingProvisionedThroughput() common.Resource {
	s := common.StructToSchema(
		ModelServingProvisionedThroughputStruct{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m, "name").SetForceNew()
			common.CustomizeSchemaPath(m, "config", "traffic_config").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_entities", "name").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_entities", "provisioned_model_units").SetRequired()
			common.CustomizeSchemaPath(m, "config", "served_entities", "entity_name").SetRequired()
			common.CustomizeSchemaPath(m, "config", "served_entities", "entity_version").SetRequired()

			common.CustomizeSchemaPath(m, "ai_gateway", "usage_tracking_config").SetOptional().SetComputed()
			common.CustomizeSchemaPath(m, "ai_gateway", "usage_tracking_config", "enabled").SetOptional().SetComputed()
			common.CustomizeSchemaPath(m, "ai_gateway", "guardrails", "input", "pii").SetOptional().SetComputed()
			common.CustomizeSchemaPath(m, "ai_gateway", "guardrails", "input", "pii", "behavior").SetOptional().SetComputed()

			m["serving_endpoint_id"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}

			// Add provider_config customizations
			common.CustomizeSchemaPath(m, "provider_config").SetOptional()
			common.CustomizeSchemaPath(m, "provider_config", "workspace_id").SetRequired()

			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e ModelServingProvisionedThroughputStruct
			common.DataToStructPointer(d, s, &e)
			wait, err := w.ServingEndpoints.CreateProvisionedThroughputEndpoint(ctx, e)
			if err != nil {
				return err
			}
			endpoint, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for serving endpoint to be created: %s", err.Error())
				nestedErr := w.ServingEndpoints.DeleteByName(ctx, e.Name)
				if nestedErr != nil {
					log.Printf("[ERROR] Error cleaning up serving endpoint: %s", nestedErr.Error())
				}
				return err
			}
			d.SetId(endpoint.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			var sOrig serving.ServingEndpointDetailed
			common.DataToStructPointer(d, s, &sOrig)
			if err != nil {
				return err
			}
			endpoint, err := w.ServingEndpoints.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(*endpoint, s, d)
			if err != nil {
				return err
			}
			d.Set("serving_endpoint_id", endpoint.Id)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e serving.CreatePtEndpointRequest
			common.DataToStructPointer(d, s, &e)
			if d.HasChange("config") {
				var updateRequest serving.UpdateProvisionedThroughputEndpointConfigRequest
				updateRequest.Name = e.Name
				updateRequest.Config = e.Config

				waiter, err := w.ServingEndpoints.UpdateProvisionedThroughputEndpointConfig(ctx, updateRequest)
				if err != nil {
					return err
				}
				_, err = waiter.GetWithTimeout(d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return err
				}
			}
			if d.HasChange("tags") {
				if err := updateTags(ctx, w, e.Name, e.Tags, d); err != nil {
					return err
				}
			}
			if d.HasChange("ai_gateway") {
				if err := updateAiGateway(ctx, w, e.Name, *e.AiGateway, d); err != nil {
					return err
				}
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.ServingEndpoints.DeleteByName(ctx, d.Id())
		},
		Schema: s,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultPtProvisionTimeout),
			Update: schema.DefaultTimeout(defaultPtProvisionTimeout),
		},
	}
}
