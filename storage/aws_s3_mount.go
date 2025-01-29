package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AWSIamMount describes the object for a aws mount using iam role
type AWSIamMount struct {
	S3BucketName string `json:"s3_bucket_name"`
}

// Source ...
func (m AWSIamMount) Source(_ *common.DatabricksClient) string {
	return fmt.Sprintf("s3a://%s", m.S3BucketName)
}

// Name ...
func (m AWSIamMount) Name() string {
	return m.S3BucketName
}

func (m AWSIamMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

// Config ...
func (m AWSIamMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

// ResourceAWSS3Mount ...
func ResourceAWSS3Mount() common.Resource {
	tpl := AWSIamMount{}
	r := common.Resource{
		DeprecationMessage: "Resource is deprecated and will be removed in further versions. " +
			"Please rewrite configuration using `databricks_mount` resource. More info at " +
			"https://registry.terraform.io/providers/databricks/databricks/latest/docs/" +
			"resources/mount#migration-from-other-mount-resources",
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mount_name": {
				// TODO: have it by default as storage_resource_name
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"s3_bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_profile": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
		SchemaVersion: 2,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
	r.Create = func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
		if err := preprocessS3Mount(ctx, d, m); err != nil {
			return err
		}
		return mountCreate(tpl, r)(ctx, d, m)
	}
	r.Read = func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
		if err := preprocessS3Mount(ctx, d, m); err != nil {
			return err
		}
		return mountRead(tpl, r)(ctx, d, m)
	}
	r.Delete = func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
		if err := preprocessS3Mount(ctx, d, m); err != nil {
			return err
		}
		return mountDelete(tpl, r)(ctx, d, m)
	}
	return r
}

func preprocessS3Mount(ctx context.Context, d *schema.ResourceData, m any) error {
	clusterID := d.Get("cluster_id").(string)
	instanceProfile := d.Get("instance_profile").(string)
	if clusterID == "" && instanceProfile == "" {
		return fmt.Errorf("either cluster_id or instance_profile must be specified")
	}
	clustersAPI := clusters.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil {
			return fmt.Errorf("cluster %s must have AWS attributes", clusterID)
		}
		if len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("cluster %s must have EC2 instance profile attached", clusterID)
		}
	} else if instanceProfile != "" {
		return mountS3ViaProfileAndSetClusterID(clustersAPI, instanceProfile, d)
	}
	return nil
}

// GetOrCreateMountingClusterWithInstanceProfile ...
func GetOrCreateMountingClusterWithInstanceProfile(
	clustersAPI clusters.ClustersAPI, instanceProfile string,
) (i clusters.ClusterInfo, err error) {
	arnSections := strings.SplitN(instanceProfile, ":", 6)
	if len(arnSections) != 6 {
		err = fmt.Errorf("invalid arn: %s", instanceProfile)
		return
	}
	instanceProfileParts := strings.Split(arnSections[5], "/")
	clusterName := fmt.Sprintf("terraform-mount-%s", strings.Join(instanceProfileParts[1:], "-"))
	cluster := getCommonClusterObject(clustersAPI, clusterName)
	cluster.AwsAttributes = &clusters.AwsAttributes{
		InstanceProfileArn: instanceProfile,
		Availability:       "SPOT",
		ZoneID:             "auto",
	}
	return clustersAPI.GetOrCreateRunningCluster(clusterName, cluster)
}
