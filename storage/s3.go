package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// S3IamMount describes the object for a aws mount using iam role
type S3IamMount struct {
	BucketName      string `json:"bucket_name" tf:"force_new"`
	InstanceProfile string `json:"instance_profile,omitempty" tf:"force_new"`
}

// Source ...
func (m S3IamMount) Source(_ *common.DatabricksClient) string {
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

func preprocessS3MountGeneric(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m any) error {
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
		if apierr.IsMissing(err) {
			if instanceProfile == "" {
				return fmt.Errorf("instance profile is required to re-create mounting cluster")
			}
			return mountS3ViaProfileAndSetClusterID(clustersAPI, instanceProfile, d)
		}
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil || len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("cluster %s must have EC2 instance profile attached", clusterID)
		}
	} else if instanceProfile != "" {
		return mountS3ViaProfileAndSetClusterID(clustersAPI, instanceProfile, d)
	}
	return nil
}

func mountS3ViaProfileAndSetClusterID(clustersAPI clusters.ClustersAPI,
	instanceProfile string, d *schema.ResourceData) error {
	cluster, err := GetOrCreateMountingClusterWithInstanceProfile(clustersAPI, instanceProfile)
	if err != nil {
		return fmt.Errorf("mount via profile: %w", err)
	}
	return d.Set("cluster_id", cluster.ClusterID)
}
