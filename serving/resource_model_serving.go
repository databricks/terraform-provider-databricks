package serving

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultProvisionTimeout = 45 * time.Minute
const deleteCallTimeout = 10 * time.Second

// updateConfig updates the configuration of the provided serving endpoint to the provided config.
func updateConfig(ctx context.Context, w *databricks.WorkspaceClient, name string, e *serving.EndpointCoreConfigInput, d *schema.ResourceData) error {
	e.Name = name
	waiter, err := w.ServingEndpoints.UpdateConfig(ctx, *e)
	if err != nil {
		return err
	}
	_, err = waiter.GetWithTimeout(d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}

// updateTags updates the tags of the provided serving endpoint to the given tags. Any tags not present on the existing
// endpoint will be removed, any tags absent on the endpoint will be added, existing tags will be updated, and unchanged
// tags will remain as-is.
func updateTags(ctx context.Context, w *databricks.WorkspaceClient, name string, newTags []serving.EndpointTag, d *schema.ResourceData) error {
	currentEndpoint, err := w.ServingEndpoints.Get(ctx, serving.GetServingEndpointRequest{
		Name: name,
	})
	oldTags := currentEndpoint.Tags
	if err != nil {
		return err
	}
	req := serving.PatchServingEndpointTags{
		Name: name,
	}
	for _, newTag := range newTags {
		found := false
		for _, oldTag := range oldTags {
			if oldTag.Key == newTag.Key && oldTag.Value == newTag.Value {
				found = true
				break
			}
		}
		if !found {
			req.AddTags = append(req.AddTags, newTag)
		}
	}
	for _, oldTag := range oldTags {
		found := false
		for _, newTag := range newTags {
			if oldTag.Key == newTag.Key {
				found = true
				break
			}
		}
		if !found {
			req.DeleteTags = append(req.DeleteTags, oldTag.Key)
		}
	}
	if _, err := w.ServingEndpoints.Patch(ctx, req); err != nil {
		return err
	}
	return nil
}

// Update the AI Gateway configuration for a model serving endpoint.
func updateAiGateway(ctx context.Context, w *databricks.WorkspaceClient, name string, newAiGateway serving.AiGatewayConfig, d *schema.ResourceData) error {
	_, err := w.ServingEndpoints.PutAiGateway(ctx, serving.PutAiGatewayRequest{
		Name:                 name,
		Guardrails:           newAiGateway.Guardrails,
		InferenceTableConfig: newAiGateway.InferenceTableConfig,
		RateLimits:           newAiGateway.RateLimits,
		UsageTrackingConfig:  newAiGateway.UsageTrackingConfig,
	})
	return err
}

func ResourceModelServing() common.Resource {
	s := common.StructToSchema(
		serving.CreateServingEndpoint{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
			// It is allowed for users to create a serving endpoint with or without a config. Removing a config
			// from an existing model serving endpoint is a no-op (i.e. the config will remain in the state and
			// the model serving endpoint will not be changed).
			common.MustSchemaPath(m, "config").Computed = true
			common.MustSchemaPath(m, "config", "served_models").ConflictsWith = []string{"config.served_entities"}
			common.MustSchemaPath(m, "config", "served_entities").ConflictsWith = []string{"config.served_models"}

			common.MustSchemaPath(m, "config", "traffic_config").Computed = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "table_name_prefix").Computed = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "enabled").Computed = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "catalog_name").ForceNew = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "schema_name").ForceNew = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "table_name_prefix").ForceNew = true

			common.MustSchemaPath(m, "config", "served_models", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_models", "workload_type").Computed = true
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Required = false
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Optional = true
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Default = true
			common.MustSchemaPath(m, "config", "served_models").Deprecated = "Please use 'config.served_entities' instead of 'config.served_models'."
			common.MustSchemaPath(m, "rate_limits").Deprecated = "Please use AI Gateway to manage rate limits."

			common.MustSchemaPath(m, "config", "served_entities", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "workload_size").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "workload_type").Computed = true

			// route_optimized cannot be updated.
			common.MustSchemaPath(m, "route_optimized").ForceNew = true

			m["serving_endpoint_id"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e serving.CreateServingEndpoint
			common.DataToStructPointer(d, s, &e)
			wait, err := w.ServingEndpoints.Create(ctx, e)
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
			if sOrig.Config == nil {
				// If it is a new resource, then we only return ServedEntities
				if endpoint.Config != nil {
					endpoint.Config.ServedModels = nil
				}
			} else {
				// If it is an existing resource, then have to set one of the responses to nil
				if sOrig.Config.ServedModels == nil {
					endpoint.Config.ServedModels = nil
				} else if sOrig.Config.ServedEntities == nil {
					endpoint.Config.ServedEntities = nil
				}
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
			var e serving.CreateServingEndpoint
			common.DataToStructPointer(d, s, &e)
			if d.HasChange("config") {
				if err := updateConfig(ctx, w, e.Name, e.Config, d); err != nil {
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
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}
}
