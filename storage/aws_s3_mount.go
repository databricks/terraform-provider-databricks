package storage

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
func (m AWSIamMount) Config() map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func ResourceAWSS3Mount() *schema.Resource {
	return commonMountResource(AWSIamMount{}, map[string]*schema.Schema{
		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
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
	})
}
