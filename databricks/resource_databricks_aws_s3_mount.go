package databricks

import (
	"fmt"

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
func (m AWSIamMount) Config() (c map[string]string) {
	return //no extra config for S3 mounts here...
}

func resourceAWSS3MountEntity() Mount {
	return new(AWSIamMount)
}

func resourceAWSS3Mount() *schema.Resource {
	resource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
	resource.Create = mountCreate(resourceAWSS3MountEntity, resource)
	resource.Read = mountRead(resourceAWSS3MountEntity, resource)
	resource.Delete = mountDelete(resourceAWSS3MountEntity, resource)
	return resource
}
