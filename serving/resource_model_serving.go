package serving

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultProvisionTimeout = 45 * time.Minute
const deleteCallTimeout = 10 * time.Second

func ResourceModelServing() common.Resource {
	s := common.StructToSchema(
		serving.CreateServingEndpoint{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
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

			common.MustSchemaPath(m, "config", "served_entities", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "workload_size").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "workload_type").Computed = true

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
				endpoint.Config.ServedModels = nil
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
			e.Config.Name = e.Name
			_, err = w.ServingEndpoints.UpdateConfigAndWait(ctx, e.Config, retries.Timeout[serving.ServingEndpointDetailed](d.Timeout(schema.TimeoutUpdate)))
			return err
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
