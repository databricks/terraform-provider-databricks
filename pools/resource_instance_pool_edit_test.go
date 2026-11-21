package pools

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

func TestResourceInstancePoolAvailabilityValidation(t *testing.T) {
	for _, tc := range []struct {
		block string
		field string
		value string
		valid bool
	}{
		{"aws_attributes", "availability", "SPOT_WITH_FALLBACK", true},
		{"azure_attributes", "availability", "SPOT_WITH_FALLBACK_AZURE", true},
		{"gcp_attributes", "gcp_availability", "PREEMPTIBLE_WITH_FALLBACK_GCP", true},
		{"aws_attributes", "availability", "SPOT_WITH_FALLBACK_AZURE", false},
		{"azure_attributes", "availability", "SPOT_WITH_FALLBACK", false},
	} {
		t.Run(tc.block+"/"+tc.value, func(t *testing.T) {
			diagnostics := ResourceInstancePool().ToResource().Validate(terraform.NewResourceConfigRaw(map[string]any{
				"instance_pool_name":                    "test-pool",
				"node_type_id":                          "test-node",
				"idle_instance_autotermination_minutes": 10,
				tc.block:                                []any{map[string]any{tc.field: tc.value}},
			}))
			assert.Equal(t, !tc.valid, diagnostics.HasError(), "%v", diagnostics)
		})
	}
}

func TestResourceInstancePoolCloudAttributePlan(t *testing.T) {
	for _, tc := range []struct {
		block       string
		field       string
		oldValue    any
		newValue    any
		requiresNew bool
	}{
		{"aws_attributes", "zone_id", "us-east-1a", "auto", false},
		{"aws_attributes", "zone_id", "auto", "us-east-1b", false},
		{"aws_attributes", "availability", "ON_DEMAND", "SPOT_WITH_FALLBACK", false},
		{"aws_attributes", "spot_bid_price_percent", 100, 75, false},
		{"azure_attributes", "availability", "ON_DEMAND_AZURE", "SPOT_WITH_FALLBACK_AZURE", false},
		{"azure_attributes", "spot_bid_max_price", -1.0, 0.0, false},
		{"azure_attributes", "spot_bid_max_price", 0.0, 0.25, false},
		{"gcp_attributes", "zone_id", "us-central1-a", "auto", false},
		{"gcp_attributes", "zone_id", "auto", "us-central1-b", false},
		{"gcp_attributes", "gcp_availability", "ON_DEMAND_GCP", "PREEMPTIBLE_WITH_FALLBACK_GCP", false},
		{"aws_attributes", "instance_profile_arn", "arn:aws:iam::123456789012:instance-profile/old", "arn:aws:iam::123456789012:instance-profile/new", true},
		{"gcp_attributes", "local_ssd_count", 1, 2, true},
	} {
		t.Run(fmt.Sprintf("%s/%s/%v", tc.block, tc.field, tc.newValue), func(t *testing.T) {
			resource := ResourceInstancePool().ToResource()
			oldConfig := map[string]any{
				"instance_pool_name":                    "test-pool",
				"node_type_id":                          "test-node",
				"idle_instance_autotermination_minutes": 10,
				tc.block:                                []any{map[string]any{tc.field: tc.oldValue}},
			}
			data := schema.TestResourceDataRaw(t, resource.Schema, oldConfig)
			data.SetId("pool-1")
			newConfig := map[string]any{
				"instance_pool_name":                    "test-pool",
				"node_type_id":                          "test-node",
				"idle_instance_autotermination_minutes": 10,
				tc.block:                                []any{map[string]any{tc.field: tc.newValue}},
			}
			c := &common.DatabricksClient{
				DatabricksClient: &client.DatabricksClient{Config: &config.Config{}},
			}
			diff, err := resource.Diff(context.Background(), data.State(), terraform.NewResourceConfigRaw(newConfig), c)
			require.NoError(t, err)
			require.NotNil(t, diff)
			attributeDiff := diff.Attributes[tc.block+".0."+tc.field]
			require.NotNil(t, attributeDiff)
			assert.Equal(t, fmt.Sprint(tc.newValue), attributeDiff.New)
			assert.Equal(t, tc.requiresNew, diff.RequiresNew())
		})
	}
}
