package exporter

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportModelServingEmitsBedrockAndBudgetDependencies(t *testing.T) {
	ic := importContextForTest()
	services, _ := ic.allServicesAndListing()
	ic.enableServices(services)
	d := schema.TestResourceDataRaw(t, ic.Resources["databricks_model_serving"].Schema, map[string]any{
		"name":                "bedrock-endpoint",
		"serving_endpoint_id": "endpoint-id",
		"budget_policy_id":    "budget-policy-id",
		"config": []any{map[string]any{
			"served_entities": []any{map[string]any{
				"external_model": []any{map[string]any{
					"name":     "claude",
					"provider": "amazon-bedrock",
					"task":     "llm/v1/chat",
					"amazon_bedrock_config": []any{map[string]any{
						"aws_region":                 "us-west-2",
						"bedrock_provider":           "anthropic",
						"instance_profile_arn":       "arn:aws:iam::123456789012:instance-profile/bedrock",
						"uc_service_credential_name": "bedrock-credential",
					}},
				}},
			}},
		}},
	})
	d.SetId("bedrock-endpoint")

	err := ic.Importables["databricks_model_serving"].Import(ic, &resource{
		Resource: "databricks_model_serving", ID: d.Id(), Name: "bedrock_endpoint", Data: d,
	})
	require.NoError(t, err)
	for _, expected := range []string{
		"databricks_credential[<unknown>] (id: bedrock-credential)",
		"databricks_instance_profile[<unknown>] (id: arn:aws:iam::123456789012:instance-profile/bedrock)",
		"databricks_budget_policy[<unknown>] (id: budget-policy-id)",
	} {
		assert.Contains(t, ic.testEmits, expected)
	}
}
