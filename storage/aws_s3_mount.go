package storage

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	}
	// nolint should be a bigger context-aware refactor
	r.Create = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return err
		}
		return mountCreate(tpl, r)(d, m)
	}

	r.Read = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return err
		}
		return mountRead(tpl, r)(d, m)
	}
	r.Delete = func(d *schema.ResourceData, m interface{}) error {
		err := preprocessS3Mount(d, m)
		if err != nil {
			return err
		}
		return mountDelete(tpl, r)(d, m)
	}
	return r
}

func preprocessS3Mount(d *schema.ResourceData, m interface{}) error {
	clusterID := d.Get("cluster_id").(string)
	instanceProfile := d.Get("instance_profile").(string)
	if clusterID == "" && instanceProfile == "" {
		return fmt.Errorf("Either cluster_id or instance_profile must be specified")
	}
	clustersAPI := compute.NewClustersAPI(m)
	if clusterID != "" {
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
	if instanceProfile != "" {
		cluster, err := getOrCreateMountingClusterWithInstanceProfile(clustersAPI, instanceProfile)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}

func getOrCreateMountingClusterWithInstanceProfile(clustersAPI compute.ClustersAPI, instanceProfile string) (i compute.ClusterInfo, err error) {
	ia, err := arn.Parse(instanceProfile)
	if err != nil {
		return i, err
	}
	instanceProfileParts := strings.Split(ia.Resource, "/")
	if len(instanceProfileParts) != 2 {
		return i, fmt.Errorf("Should have gotten two parts: %v", instanceProfileParts)
	}
	clusterName := fmt.Sprintf("terraform-mount-%s", instanceProfileParts[1])
	return clustersAPI.GetOrCreateRunningCluster(clusterName, compute.Cluster{
		NumWorkers:   1,
		ClusterName:  clusterName,
		SparkVersion: compute.CommonRuntimeVersion(),
		NodeTypeID: clustersAPI.GetSmallestNodeType(compute.NodeTypeRequest{
			LocalDisk: true,
		}),
		AutoterminationMinutes: 10,
		AwsAttributes: &compute.AwsAttributes{
			InstanceProfileArn: instanceProfile,
			Availability:       "SPOT",
		},
	})
}
