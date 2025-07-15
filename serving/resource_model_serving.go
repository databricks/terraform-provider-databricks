package serving

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	DefaultProvisionTimeout = 45 * time.Minute
	deleteCallTimeout       = 10 * time.Second
)

// suppressRouteModelEntityNameDiff suppresses diff for served_model_name and served_entity_name
// in traffic config routes since they are semantically equivalent. The server returns both fields
// with the same value even when only one is explicitly set in the HCL configuration.
func suppressRouteModelEntityNameDiff(k, old, new string, d *schema.ResourceData) bool {
	// Check if this is a served_model_name or served_entity_name field in routes
	if !(strings.Contains(k, "served_model_name") || strings.Contains(k, "served_entity_name")) {
		return false
	}

	// Extract the base path (e.g., "config.0.traffic_config.0.routes.0")
	// so we can check the corresponding field
	basePath := k
	var otherField string

	if strings.Contains(k, "served_model_name") {
		otherField = "served_entity_name"
		basePath = strings.Replace(k, "served_model_name", "", 1)
	} else {
		otherField = "served_model_name"
		basePath = strings.Replace(k, "served_entity_name", "", 1)
	}

	// Get the value of the other field
	otherFieldKey := basePath + otherField
	otherFieldValue := d.Get(otherFieldKey).(string)

	// If the new value is empty (not set in config) and the old value equals the other field's value,
	// then suppress the diff (server is returning both fields with same value)
	if new == "" && old != "" && old == otherFieldValue {
		log.Printf("[DEBUG] Suppressing diff for %v: old=%#v new=%#v (server returned both served_model_name and served_entity_name with same value)", k, old, new)
		return true
	}

	// If the old value is empty (not previously set) and the new value equals the other field's value,
	// then suppress the diff (server is returning both fields with same value)
	if old == "" && new != "" && new == otherFieldValue {
		log.Printf("[DEBUG] Suppressing diff for %v: old=%#v new=%#v (server returned both served_model_name and served_entity_name with same value)", k, old, new)
		return true
	}

	return false
}

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
		FallbackConfig:       newAiGateway.FallbackConfig,
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
			// Use the newer CustomizeSchemaPath approach for better maintainability
			common.CustomizeSchemaPath(m, "name").SetForceNew()

			// It is allowed for users to create a serving endpoint with or without a config. Removing a config
			// from an existing model serving endpoint is a no-op (i.e. the config will remain in the state and
			// the model serving endpoint will not be changed).
			common.CustomizeSchemaPath(m, "config").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_models").SetConflictsWith([]string{"config.served_entities"})
			common.CustomizeSchemaPath(m, "config", "served_entities").SetConflictsWith([]string{"config.served_models"})

			common.CustomizeSchemaPath(m, "config", "traffic_config").SetComputed()
			common.CustomizeSchemaPath(m, "config", "auto_capture_config", "table_name_prefix").SetComputed()
			common.CustomizeSchemaPath(m, "config", "auto_capture_config", "enabled").SetComputed()
			common.CustomizeSchemaPath(m, "config", "auto_capture_config", "catalog_name").SetForceNew()
			common.CustomizeSchemaPath(m, "config", "auto_capture_config", "schema_name").SetForceNew()
			common.CustomizeSchemaPath(m, "config", "auto_capture_config", "table_name_prefix").SetForceNew()

			common.CustomizeSchemaPath(m, "config", "served_models", "name").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_models", "workload_type").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").SetOptional().SetDefault(true)
			common.CustomizeSchemaPath(m, "config", "served_models").SetDeprecated("Please use 'config.served_entities' instead of 'config.served_models'.")
			common.CustomizeSchemaPath(m, "rate_limits").SetDeprecated("Please use AI Gateway to manage rate limits.")

			common.CustomizeSchemaPath(m, "config", "served_entities", "name").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_entities", "workload_size").SetComputed()
			common.CustomizeSchemaPath(m, "config", "served_entities", "workload_type").SetComputed()

			// Apply custom suppress diff to traffic config routes for served_model_name and served_entity_name
			common.CustomizeSchemaPath(m, "config", "traffic_config", "routes", "served_model_name").SetCustomSuppressDiff(suppressRouteModelEntityNameDiff)
			common.CustomizeSchemaPath(m, "config", "traffic_config", "routes", "served_entity_name").SetCustomSuppressDiff(suppressRouteModelEntityNameDiff)

			common.MustSchemaPath(m, "ai_gateway", "guardrails", "input", "invalid_keywords").Deprecated = "Please use 'pii' and 'safety' instead."
			common.MustSchemaPath(m, "ai_gateway", "guardrails", "input", "valid_topics").Deprecated = "Please use 'pii' and 'safety' instead."
			common.MustSchemaPath(m, "ai_gateway", "guardrails", "output", "invalid_keywords").Deprecated = "Please use 'pii' and 'safety' instead."
			common.MustSchemaPath(m, "ai_gateway", "guardrails", "output", "valid_topics").Deprecated = "Please use 'pii' and 'safety' instead."

			// route_optimized cannot be updated.
			common.CustomizeSchemaPath(m, "route_optimized").SetForceNew()

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
