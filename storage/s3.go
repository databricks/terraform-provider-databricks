package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// S3IamMount describes the object for a aws mount using iam role
type S3IamMount struct {
	BucketName      string `json:"bucket_name" tf:"force_new"`
	InstanceProfile string `json:"instance_profile,omitempty" tf:"force_new"`
}

// Source ...
func (m S3IamMount) Source() string {
	return fmt.Sprintf("s3a://%s", m.BucketName)
}

// Name ...
func (m S3IamMount) Name() string {
	return m.BucketName
}

// Config ...
func (m S3IamMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func (m S3IamMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	nm := d.Get("name").(string)
	if nm != "" {
		return nil
	}
	nm = m.Name()
	if nm != "" {
		d.Set("name", nm)
		return nil
	}
	return fmt.Errorf("'name' is not detected & it's impossible to infer it")
}

func preprocessS3MountGeneric(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	common.DataToStructPointer(d, s, &gm)
	// TODO: move into Validate function
	if !(strings.HasPrefix(gm.URI, "s3://") || strings.HasPrefix(gm.URI, "s3a://") || gm.S3 != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	instanceProfile := ""
	if gm.S3 != nil {
		instanceProfile = gm.S3.InstanceProfile
	}
	if clusterID == "" && instanceProfile == "" {
		return fmt.Errorf("either cluster_id or instance_profile must be specified to mount S3 bucket")
	}
	clustersAPI := clusters.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil || len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("cluster %s must have EC2 instance profile attached", clusterID)
		}
	}
	if instanceProfile != "" {
		cluster, err := GetOrCreateMountingClusterWithInstanceProfile(clustersAPI, instanceProfile)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}
