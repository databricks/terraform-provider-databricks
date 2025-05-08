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

func ResourceModelServingProvisionedThroughput() common.Resource {
	s := common.StructToSchema(
		serving.CreatePtEndpointRequest{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
			common.MustSchemaPath(m, "config", "traffic_config").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "provisioned_model_units").Required = true

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
			Create: schema.DefaultTimeout(defaultPtProvisionTimeout),
			Update: schema.DefaultTimeout(defaultPtProvisionTimeout),
		},
	}
}
