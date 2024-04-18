package serving

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
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

			common.MustSchemaPath(m, "config", "served_entities", "scale_to_zero_enabled").Required = false
			common.MustSchemaPath(m, "config", "served_entities", "scale_to_zero_enabled").Optional = true
			common.MustSchemaPath(m, "config", "served_entities", "scale_to_zero_enabled").Default = true
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
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			old, new := d.GetChange("config.0.auto_capture_config.0.enabled")
			if old != nil && old == false && new == true {
				d.ForceNew("config.0.auto_capture_config.0.enabled")
			}
			err := validateExternalModelConfig(d)
			if err != nil {
				return err
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e serving.CreateServingEndpoint
			common.DataToStructPointer(d, s, &e)
			for i := range e.Config.ServedEntities {
				e.Config.ServedEntities[i].ForceSendFields = append(e.Config.ServedEntities[i].ForceSendFields, "ScaleToZeroEnabled")
			}
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
			for i := range e.Config.ServedEntities {
				e.Config.ServedEntities[i].ForceSendFields = append(e.Config.ServedEntities[i].ForceSendFields, "ScaleToZeroEnabled")
			}
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

func validateExternalModelConfig(d *schema.ResourceDiff) error {
	_, e := d.GetOk("config.0.served_entities.0.external_model")
	provider, p := d.GetOk("config.0.served_entities.0.external_model.0.provider")

	if !e || !p {
		return nil
	}

	name := strings.ReplaceAll(provider.(string), "-", "_")
	config := d.Get(fmt.Sprintf("config.0.served_entities.0.external_model.0.%s_config", name)).([]interface{})

	if len(config) == 0 {
		return fmt.Errorf("external_model provider is set to \"%s\" but \"%s_config\" block is missing", name, name)
	}

	if configBlock, ok := d.Get("config.0.served_entities.0.external_model.0").(map[string]interface{}); ok {
		var found []string
		for key, value := range configBlock {
			if strings.HasSuffix(key, "_config") && len(value.([]interface{})) > 0 {
				found = append(found, key)
			}
		}
		slices.Sort(found)
		if len(found) > 1 {
			msg := strings.Join(found, ", ")
			return fmt.Errorf("only one external_model config block is allowed. Found: %s", msg)
		}
	}
	return nil
}
