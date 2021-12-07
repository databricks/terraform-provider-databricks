package storage

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func preprocessResourceData(ctx context.Context, d *schema.ResourceData, scm map[string]*schema.Schema, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, scm, &gm); err != nil {
		return err
	}
	// TODO: propagate ctx all the way down to GetAzureJwtProperty()
	if err := gm.ValidateAndApplyDefaults(d, m.(*common.DatabricksClient)); err != nil {
		return err
	}
	if err := common.StructToData(gm, scm, d); err != nil {
		return err
	}
	if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
		return err
	}
	if err := preprocessGsMount(ctx, scm, d, m); err != nil {
		return err
	}
	return nil
}

func ResourceDatabricksMountSchema() map[string]*schema.Schema {
	tpl := GenericMount{}
	scm := common.StructToSchema(tpl, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["source"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}

		s["uri"].ConflictsWith = []string{"abfs", "wasb", "s3", "adl", "gs"}
		s["extra_configs"].ConflictsWith = []string{"abfs", "wasb", "s3", "adl", "gs"}
		s["abfs"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "adl", "gs"}
		s["wasb"].ConflictsWith = []string{"uri", "extra_configs", "abfs", "s3", "adl", "gs"}
		s["s3"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "abfs", "adl", "gs"}
		s["adl"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "abfs", "gs"}
		s["gs"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "abfs", "adl"}
		// TODO: We need to have a validation function that will check that source isn't empty if other blocks aren't specified

		return s
	})
	return scm
}

// ResourceDatabricksMount mounts using given configuration
func ResourceDatabricksMount() *schema.Resource {
	tpl := GenericMount{}
	scm := ResourceDatabricksMountSchema()

	r := commonMountResource(tpl, scm)
	r.CreateContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessResourceData(ctx, d, scm, m); err != nil {
			return diag.FromErr(err)
		}
		return mountCreate(tpl, r)(ctx, d, m)
	}
	r.ReadContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessResourceData(ctx, d, scm, m); err != nil {
			return diag.FromErr(err)
		}
		return mountRead(tpl, r)(ctx, d, m)
	}
	r.DeleteContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessResourceData(ctx, d, scm, m); err != nil {
			return diag.FromErr(err)
		}
		return mountDelete(tpl, r)(ctx, d, m)
	}
	return r
}
