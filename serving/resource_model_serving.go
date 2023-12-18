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

type CreateServingEndpoint struct {
	// The core config of the serving endpoint.
	Config EndpointCoreConfigInput `json:"config"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: only external
	// and foundation model endpoints are supported as of now.
	RateLimits []serving.RateLimit `json:"rate_limits,omitempty"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []serving.EndpointTag `json:"tags,omitempty"`
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *serving.AutoCaptureConfigInput `json:"auto_capture_config,omitempty"`
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
	// A list of served models for the endpoint to serve.
	// A serving endpoint can have up to 10 served models.
	ServedModels []serving.ServedModelInput `json:"served_models,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *serving.TrafficConfig `json:"traffic_config,omitempty"`
}

func ResourceModelServing() *schema.Resource {
	s := common.StructToSchema(
		CreateServingEndpoint{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
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
