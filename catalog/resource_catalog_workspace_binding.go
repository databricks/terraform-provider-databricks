package catalog

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var getSecurableName = func(d *schema.ResourceData) string {
	securableName, ok := d.GetOk("securable_name")
	if !ok {
		securableName = d.Get("catalog_name")
	}
	return securableName.(string)
}

func ResourceCatalogWorkspaceBinding() common.Resource {
	workspaceBindingSchema := common.StructToSchema(catalog.WorkspaceBinding{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["catalog_name"] = &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"catalog_name", "securable_name"},
				Deprecated:   "Please use 'securable_name' and 'securable_type instead.",
			}
			m["securable_name"] = &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"catalog_name", "securable_name"},
			}
			m["securable_type"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "catalog",
			}
			m["binding_type"].Default = catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite
			m["binding_type"].ValidateFunc = validation.StringInSlice([]string{
				string(catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite),
				string(catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly),
			}, false)
			return m
		},
	)
	return common.Resource{
		Schema:        workspaceBindingSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 0,
				Type:    bindingSchemaV0(),
				Upgrade: bindingMigrateV0,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.WorkspaceBinding
			common.DataToStructPointer(d, workspaceBindingSchema, &update)

			securableName := getSecurableName(d)
			_, err = w.WorkspaceBindings.UpdateBindings(ctx, catalog.UpdateWorkspaceBindingsParameters{
				Add:           []catalog.WorkspaceBinding{update},
				SecurableName: securableName,
				SecurableType: d.Get("securable_type").(string),
			})
			d.SetId(fmt.Sprintf("%d|%s|%s", update.WorkspaceId, d.Get("securable_type").(string), securableName))
			return err
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// TODO: fix Read operation by splitting `id` into parts... Test with actual import. Remove not necessary code in exporter?
			workspaceId := int64(d.Get("workspace_id").(int))
			securable_name := getSecurableName(d)
			securable_type := d.Get("securable_type").(string)
			if workspaceId == 0 || securable_name == "" || securable_type == "" {
				parts := strings.Split(d.Id(), "|")
				if len(parts) != 3 {
					return fmt.Errorf("incorrect binding id: %s", d.Id())
				}
				securable_name = parts[2]
				securable_type = parts[1]
				workspaceId, err = strconv.ParseInt(parts[0], 10, 0)
				if err != nil {
					return err
				}
				d.Set("securable_name", securable_name)
				d.Set("securable_type", securable_type)
				d.Set("workspace_id", workspaceId)
			}
			bindings, err := w.WorkspaceBindings.GetBindings(ctx, catalog.GetBindingsRequest{
				SecurableName: securable_name,
				SecurableType: securable_type,
			})
			if err != nil {
				return err
			}
			for _, binding := range bindings.Bindings {
				if binding.WorkspaceId == workspaceId {
					return common.StructToData(binding, workspaceBindingSchema, d)
				}
			}
			return apierr.NotFound("Catalog has no binding to this workspace")
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.WorkspaceBinding
			common.DataToStructPointer(d, workspaceBindingSchema, &update)
			_, err = w.WorkspaceBindings.UpdateBindings(ctx, catalog.UpdateWorkspaceBindingsParameters{
				Remove:        []catalog.WorkspaceBinding{update},
				SecurableName: getSecurableName(d),
				SecurableType: d.Get("securable_type").(string),
			})
			return err
		},
	}
}

// migrate to v1 state, as catalog_name is moved to securable_name
func bindingMigrateV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	newState := map[string]any{}
	log.Printf("[INFO] Upgrade workspace binding schema")
	newState["securable_name"] = rawState["catalog_name"]
	newState["securable_type"] = "catalog"
	newState["catalog_name"] = rawState["catalog_name"]
	newState["workspace_id"] = rawState["workspace_id"]
	newState["binding_type"] = string(catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite)
	return newState, nil
}

func bindingSchemaV0() cty.Type {
	return (&schema.Resource{
		Schema: map[string]*schema.Schema{
			"catalog_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}}).CoreConfigSchema().ImpliedType()
}
