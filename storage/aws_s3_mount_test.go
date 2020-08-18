package storage

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

// Test interface compliance via compile time error
var _ Mount = (*AWSIamMount)(nil)

func TestAwsS3Mount_mount(t *testing.T) {
	randomBucketName := qa.RandomName()
	awsMount := AWSIamMount{
		S3BucketName: randomBucketName,
	}
	expSource := fmt.Sprintf("s3a://%s", randomBucketName)
	testMountPointCreateHelper(t, awsMount, expSource, "{}")
}
