package serving

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultProvisionTimeout = 45 * time.Minute

func ResourceModelServing() *schema.Resource {
	s := common.StructToSchema(
		serving.CreateServingEndpoint{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
			common.MustSchemaPath(m, "config", "served_models").ConflictsWith = []string{"config.served_entities"}
			common.MustSchemaPath(m, "config", "served_entities").ConflictsWith = []string{"config.served_models"}
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Required = false
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Optional = true
			common.MustSchemaPath(m, "config", "served_models", "scale_to_zero_enabled").Default = true
			common.MustSchemaPath(m, "config", "served_models", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_models", "workload_type").Default = "CPU"
			// TODO: `config.served_models.workload_type` should be a `Optional+Computed` field. Also consider this for other similar fields.
			// In this scenario, if a workspace does not have GPU serving, specifying `workload_type` = 'CPU' will get empty response from API.
			common.MustSchemaPath(m, "config", "served_models", "workload_type").DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old == "" && new == "CPU"
			}
			common.MustSchemaPath(m, "config", "traffic_config").Computed = true
			common.MustSchemaPath(m, "config", "served_models").Deprecated = "Please use 'config.served_entities' instead of 'config.served_models'."

			common.MustSchemaPath(m, "config", "served_entities", "name").Computed = true
			common.MustSchemaPath(m, "config", "served_entities", "workload_type").Default = "CPU"
			common.MustSchemaPath(m, "config", "served_entities", "workload_type").DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old == "" && new == "CPU"
			}
			common.MustSchemaPath(m, "config", "auto_capture_config", "catalog_name").ForceNew = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "schema_name").ForceNew = true
			common.MustSchemaPath(m, "config", "auto_capture_config", "table_name_prefix").ForceNew = true

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
			endpoint, err := w.ServingEndpoints.CreateAndWait(ctx, e, retries.Timeout[serving.ServingEndpointDetailed](d.Timeout(schema.TimeoutCreate)))
			if err != nil {
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
			// WIP: There must be a way to know which resource the user is trying to create
			// otherwise it will lead to a diff for existing resources in previous providers
			if d.IsNewResource() || sOrig.Config == nil {
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
	}.ToResource()
}
