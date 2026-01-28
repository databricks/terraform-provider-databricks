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

// preserveConfigOrderPt re-orders the served_entities in the API response for provisioned
// throughput endpoints to match the order specified in the HCL configuration. This prevents
// spurious diffs when the API returns items in a different order (e.g., alphabetically).
func preserveConfigOrderPt(s map[string]*schema.Schema, d *schema.ResourceData, apiResponse *serving.EndpointCoreConfigOutput) {
	var config serving.CreatePtEndpointRequest
	common.DataToStructPointer(d, s, &config)

	if apiResponse == nil {
		return
	}

	// Re-order served_entities to match config order
	if len(config.Config.ServedEntities) > 0 && len(apiResponse.ServedEntities) > 0 {
		configNames := make([]string, len(config.Config.ServedEntities))
		for i, entity := range config.Config.ServedEntities {
			configNames[i] = entity.Name
		}
		apiResponse.ServedEntities = reorderByName(configNames, apiResponse.ServedEntities,
			func(e serving.ServedEntityOutput) string { return e.Name })
	}
}

type ModelServingProvisionedThroughputSchemaStruct struct {
	serving.CreatePtEndpointRequest
	common.Namespace
}

func ResourceModelServingProvisionedThroughput() common.Resource {
	s := common.StructToSchema(
		ModelServingProvisionedThroughputSchemaStruct{},
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

			// Tags should have Set type
			m["tags"].Type = schema.TypeSet

			m["serving_endpoint_id"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}
			common.NamespaceCustomizeSchemaMap(m)
			return m
		})

	return common.Resource{
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var e serving.CreatePtEndpointRequest
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			var sOrig serving.ServingEndpointDetailed
			common.DataToStructPointer(d, s, &sOrig)
			if err != nil {
				return err
			}
			endpoint, err := w.ServingEndpoints.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			preserveConfigOrderPt(s, d, endpoint.Config)
			err = common.StructToData(*endpoint, s, d)
			if err != nil {
				return err
			}
			d.Set("serving_endpoint_id", endpoint.Id)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
