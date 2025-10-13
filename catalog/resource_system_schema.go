package catalog

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SystemSchemaInfo struct {
	catalog.SystemSchemaInfo
	common.Namespace
}

func ResourceSystemSchema() common.Resource {
	systemSchema := common.StructToSchema(SystemSchemaInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["metastore_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		m["full_name"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		m["auto_enabled"] = &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		}
		m["state"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		common.NamespaceCustomizeSchemaMap(m)
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
		w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
		if err != nil {
			//ignore "<schema-name> system schema can only be enabled by Databricks" error, also mark it to make delete no-op
			d.Set("auto_enabled", strings.Contains(err.Error(), "can only be enabled by Databricks"))
			//ignore "schema <schema-name> already exists" error
			if !strings.Contains(err.Error(), "already exists") && !strings.Contains(err.Error(), "can only be enabled by Databricks") {
				return err
			}
		}
		//disable old schemas if needed
		if old != "" {
			if d.Get("auto_enabled").(bool) {
				log.Printf("[WARN] %s is auto enabled, ignoring it", old)
				return nil
			}
			err = safeDisable(ctx, w, metastoreSummary.MetastoreId, old)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
					if schema.State != string(SystemSchemaInfoStateEnableCompleted) &&
						schema.State != string(SystemSchemaInfoStateEnableInitialized) &&
						schema.State != string(SystemSchemaInfoStateUnavailable) {
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
			if d.Get("auto_enabled").(bool) {
				log.Printf("[WARN] %s is auto enabled, ignoring it", schemaName)
				return nil
			}
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			metastoreSummary, err := w.Metastores.Summary(ctx)
			if err != nil {
				return err
			}
			return safeDisable(ctx, w, metastoreSummary.MetastoreId, schemaName)
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
	}
}

func safeDisable(ctx context.Context, w *databricks.WorkspaceClient, metastoreId, schemaName string) error {
	err := w.SystemSchemas.Disable(ctx, catalog.DisableRequest{
		MetastoreId: metastoreId,
		SchemaName:  schemaName,
	})
	if err != nil {
		//ignore "<schema-name> system schema can only be disabled by Databricks" error
		if !strings.Contains(err.Error(), "can only be disabled by Databricks") {
			return err
		}
		log.Printf("[WARN] %s can be disabled only by Databricks, ignoring it", schemaName)
	}
	return nil
}
