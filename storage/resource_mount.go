package storage

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type mountCallback func(tpl interface{}, r *schema.Resource) func(context.Context,
	*schema.ResourceData, interface{}) diag.Diagnostics

func (cb mountCallback) preProcess(r *schema.Resource) func(
	ctx context.Context, d *schema.ResourceData,
	m interface{}) diag.Diagnostics {
	tpl := GenericMount{}
	return func(ctx context.Context, d *schema.ResourceData,
		m interface{}) diag.Diagnostics {
		var gm GenericMount
		scm := r.Schema
		common.DataToStructPointer(d, scm, &gm)
		// TODO: propagate ctx all the way down to GetAzureJwtProperty()
		err := gm.ValidateAndApplyDefaults(d, m.(*common.DatabricksClient))
		if err != nil {
			return diag.FromErr(err)
		}
		common.StructToData(gm, scm, d)
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return cb(tpl, r)(ctx, d, m)
	}
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

// ResourceMount mounts using given configuration
func ResourceMount() *schema.Resource {
	tpl := GenericMount{}
	r := commonMountResource(tpl, ResourceDatabricksMountSchema())
	r.CreateContext = mountCallback(mountCreate).preProcess(r)
	r.ReadContext = mountCallback(mountRead).preProcess(r)
	r.DeleteContext = mountCallback(mountDelete).preProcess(r)
	return r
}
