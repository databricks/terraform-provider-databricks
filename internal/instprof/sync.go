package instprof

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Synchronized works around the fact, that we can register instance profile only once
func Synchronized(t *testing.T, cb func(string)) {
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	client := common.NewClientFromEnvironment()
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(client)
	resource.RetryContext(context.Background(), 10*time.Minute, func() *resource.RetryError {
		list, err := instanceProfilesAPI.List()
		if err != nil {
			return resource.NonRetryableError(err)
		}
		for _, ip := range list {
			if ip.InstanceProfileArn == arn {
				return resource.RetryableError(fmt.Errorf(
					"%s is registered, waiting to release", arn))
			}
		}
		cb(arn)
		return nil
	})
}
