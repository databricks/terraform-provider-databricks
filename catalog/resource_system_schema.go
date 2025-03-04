package catalog

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSystemSchema() common.Resource {
	systemSchema := common.StructToSchema(catalog.SystemSchemaInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["metastore_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		m["full_name"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		m["state"].Computed = true
		return m
	})
	pi := common.NewPairID("metastore_id", "schema").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return systemSchema
		})
	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		o, n := d.GetChange("schema")
		old, okOld := o.(string)
		new, okNew := n.(string)
		if !okNew || !okOld {
			return fmt.Errorf("internal type casting error")
		}
		log.Printf("[DEBUG] Old system schema: %s, new: %s", old, new)
		w, err := c.WorkspaceClient()
		if err != nil {
			return err
		}
		metastoreSummary, err := w.Metastores.Summary(ctx)
		if err != nil {
			return err
		}
		//enable new schema
		err = w.SystemSchemas.Enable(ctx, catalog.EnableRequest{
			MetastoreId: metastoreSummary.MetastoreId,
			SchemaName:  new,
		})
		//ignore "schema <schema-name> already exists" error
		//ignore "<schema-name> system schema can only be enabled by Databricks" error
		if err != nil && !strings.Contains(err.Error(), "already exists") && !strings.Contains(err.Error(), "can only be enabled by Databricks") {
			return err
		}
		//disable old schemas if needed
		if old != "" {
			err = w.SystemSchemas.Disable(ctx, catalog.DisableRequest{
				MetastoreId: metastoreSummary.MetastoreId,
				SchemaName:  old,
			})
			if err != nil {
				return err
			}
		}
		d.Set("metastore_id", metastoreSummary.MetastoreId)
		pi.Pack(d)
		return nil
	}
	return common.Resource{
		Schema: systemSchema,
		Create: createOrUpdate,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, schemaName, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			metastoreSummary, err := w.Metastores.Summary(ctx)
			if err != nil {
				return err
			}
			systemSchemaInfo, err := w.SystemSchemas.ListByMetastoreId(ctx, metastoreSummary.MetastoreId)
			if err != nil {
				return err
			}
			for _, schema := range systemSchemaInfo.Schemas {
				if schema.Schema == schemaName {
					err = common.StructToData(schema, systemSchema, d)
					if err != nil {
						return err
					}
					// only track enabled/legacy schemas
					if schema.State != catalog.SystemSchemaInfoStateEnableCompleted &&
						schema.State != catalog.SystemSchemaInfoStateEnableInitialized &&
						schema.State != catalog.SystemSchemaInfoStateUnavailable {
						log.Printf("[WARN] %s is not enabled, ignoring it", schemaName)
						d.SetId("")
						return nil
					}

					d.Set("full_name", fmt.Sprintf("system.%s", schemaName))
					return nil
				}
			}
			log.Printf("[WARN] %s does not exist, ignoring it", schemaName)
			d.SetId("")
			return nil
		},
		Update: createOrUpdate,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, schemaName, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			metastoreSummary, err := w.Metastores.Summary(ctx)
			if err != nil {
				return err
			}
			return w.SystemSchemas.Disable(ctx, catalog.DisableRequest{
				MetastoreId: metastoreSummary.MetastoreId,
				SchemaName:  schemaName,
			})
		},
	}
}
