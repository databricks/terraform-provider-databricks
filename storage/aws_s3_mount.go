package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// AWSIamMount describes the object for a aws mount using iam role
type AWSIamMount struct {
	S3BucketName string `json:"s3_bucket_name"`
}

// Source ...
func (m AWSIamMount) Source() string {
	return fmt.Sprintf("s3a://%s", m.S3BucketName)
}

// Config ...
func (m AWSIamMount) Config() map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func ResourceAWSS3Mount() *schema.Resource {
	tpl := AWSIamMount{}
	r := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"instance_profile_arn"},
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
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster_id"},
			},
		},
		SchemaVersion: 2,
	}
	r.Create = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return nil
		}
		return mountCreate(tpl, r)(d, m)
	}

	r.Read = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return nil
		}
		return mountRead(tpl, r)(d, m)
	}
	r.Delete = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return nil
		}
		return mountDelete(tpl, r)(d, m)
	}
	return r
}

func preprocessS3Mount(d *schema.ResourceData, m interface{}) error {
	clustersAPI := compute.NewClustersAPI(m)
	if clusterID, ok := d.Get("cluster_id").(string); ok && clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil {
			return fmt.Errorf("Cluster %s must have AWS attributes", clusterID)
		}
		if len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("Cluster %s must have EC2 instance profile attached", clusterID)
		}
	}
	if instanceProfile, ok := d.Get("instance_profile").(string); ok && instanceProfile != "" {
		ia, err := arn.Parse(instanceProfile)
		if err != nil {
			return err
		}
		clusterName := fmt.Sprintf("terraform-mount-%s", ia.Resource)
		cluster, err := clustersAPI.GetOrCreateRunningCluster(clusterName, compute.Cluster{
			NumWorkers:             1,
			ClusterName:            clusterName,
			SparkVersion:           compute.CommonRuntimeVersion(),
			NodeTypeID:             clustersAPI.GetSmallestNodeTypeWithStorage(),
			AutoterminationMinutes: 10,
			AwsAttributes: &compute.AwsAttributes{
				InstanceProfileArn: instanceProfile,
				Availability:       "SPOT",
			},
		})
		if err != nil {
			return err
		}
		d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}
