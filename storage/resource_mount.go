package storage

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type mountCallback func(tpl any, r *schema.Resource) func(context.Context,
	*schema.ResourceData, any) diag.Diagnostics

func (cb mountCallback) preProcess(r *schema.Resource) func(
	ctx context.Context, d *schema.ResourceData,
	m any) diag.Diagnostics {
	tpl := GenericMount{}
	return func(ctx context.Context, d *schema.ResourceData,
		m any) diag.Diagnostics {
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

// If a cluster_id is specified for the mount, this function validates that the
// cluster actually exists. If the cluster does not exist, we assume (incorrectly)
// that the mount also does not exist.
//
// In cases of update requests this would lead to an attempted recreation of the mount.
//
// In cases of delete, this could orphan the mount, ie that the mount still exists
// but is not tracked by the terraform state.
//
// This is needed to resolve issue: https://github.com/databricks/terraform-provider-databricks/issues/1864
// because without a running cluster we do not have a way to do CRUD on mounts.
// The argument being that possibly orphaning mounts is a better experience than
// requiring manual edits to tfstate.
//
// Why do we need a valid cluster_id to be set here?
// Answer: Downstream read and delete code relies on a valid cluster_id being set
// in the terraform state. If the cluster_id is invalid then there is no way to update
// it using native terraform apply workflows. The only workaround is to manually edit the
// tfstate file replacing the existing invalid cluster_id with a new valid one.
func validateCluster(cb func(context.Context, *schema.ResourceData, any) diag.Diagnostics) func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		clusterId := d.Get("cluster_id").(string)
		if clusterId != "" {
			clustersAPI := clusters.NewClustersAPI(ctx, m)
			_, err := clustersAPI.Get(clusterId)
			if apierr.IsMissing(err) {
				// Return with empty ID to indicate the resource no longer exists.
				d.SetId("")
				return nil
			}
		}
		return cb(ctx, d, m)
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
	r.ReadContext = validateCluster(mountCallback(mountRead).preProcess(r))
	r.DeleteContext = validateCluster(mountCallback(mountDelete).preProcess(r))
	r.Importer = nil
	r.Timeouts = &schema.ResourceTimeout{
		Default: schema.DefaultTimeout(20 * time.Minute),
	}
	return r
}
