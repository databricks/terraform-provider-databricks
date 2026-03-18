package app

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestBuildUpdateMask_GitRepository(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	// State with no git_repository
	state := AppResource{
		App: apps_tf.App{
			Name:        types.StringValue("test-app"),
			Description: types.StringValue("Original description"),
		},
	}

	// Plan with git_repository added
	plan := AppResource{
		App: apps_tf.App{
			Name:        types.StringValue("test-app"),
			Description: types.StringValue("Updated description"),
			GitRepository: types.ObjectValueMust(
				map[string]attr.Type{
					"provider": types.StringType,
					"url":      types.StringType,
				},
				map[string]attr.Value{
					"provider": types.StringValue("gitHub"),
					"url":      types.StringValue("https://github.com/user/repo"),
				},
			),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Contains(t, updateMask, "description")
	assert.Contains(t, updateMask, "git_repository")
}

func TestBuildUpdateMask_MultipleFields(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	// State with some values
	state := AppResource{
		App: apps_tf.App{
			Name:        types.StringValue("test-app"),
			Description: types.StringValue("Original description"),
			ComputeSize: types.StringValue("SMALL"),
		},
	}

	// Plan with multiple fields changed
	plan := AppResource{
		App: apps_tf.App{
			Name:           types.StringValue("test-app"),
			Description:    types.StringValue("Updated description"),
			ComputeSize:    types.StringValue("MEDIUM"),
			BudgetPolicyId: types.StringValue("budget-123"),
			UsagePolicyId:  types.StringValue("usage-456"),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Contains(t, updateMask, "description")
	assert.Contains(t, updateMask, "compute_size")
	assert.Contains(t, updateMask, "budget_policy_id")
	assert.Contains(t, updateMask, "usage_policy_id")
}

func TestBuildUpdateMask_NoChanges(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	// State and plan are identical (with all nullable fields explicitly set to null)
	state := AppResource{
		App: apps_tf.App{
			Name:           types.StringValue("test-app"),
			Description:    types.StringValue("Same description"),
			GitRepository:  types.ObjectNull(map[string]attr.Type{}),
			Resources:      types.ListNull(types.ObjectType{}),
			BudgetPolicyId: types.StringNull(),
			ComputeSize:    types.StringNull(),
			UsagePolicyId:  types.StringNull(),
			UserApiScopes:  types.ListNull(types.StringType),
		},
	}

	plan := AppResource{
		App: apps_tf.App{
			Name:           types.StringValue("test-app"),
			Description:    types.StringValue("Same description"),
			GitRepository:  types.ObjectNull(map[string]attr.Type{}),
			Resources:      types.ListNull(types.ObjectType{}),
			BudgetPolicyId: types.StringNull(),
			ComputeSize:    types.StringNull(),
			UsagePolicyId:  types.StringNull(),
			UserApiScopes:  types.ListNull(types.StringType),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Equal(t, "", updateMask)
}

func TestBuildUpdateMask_DescriptionOnly(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	state := AppResource{
		App: apps_tf.App{
			Name:           types.StringValue("test-app"),
			Description:    types.StringValue("Old description"),
			GitRepository:  types.ObjectNull(map[string]attr.Type{}),
			Resources:      types.ListNull(types.ObjectType{}),
			BudgetPolicyId: types.StringNull(),
			ComputeSize:    types.StringNull(),
			UsagePolicyId:  types.StringNull(),
			UserApiScopes:  types.ListNull(types.StringType),
		},
	}

	plan := AppResource{
		App: apps_tf.App{
			Name:           types.StringValue("test-app"),
			Description:    types.StringValue("New description"),
			GitRepository:  types.ObjectNull(map[string]attr.Type{}),
			Resources:      types.ListNull(types.ObjectType{}),
			BudgetPolicyId: types.StringNull(),
			ComputeSize:    types.StringNull(),
			UsagePolicyId:  types.StringNull(),
			UserApiScopes:  types.ListNull(types.StringType),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Equal(t, "description", updateMask)
}

func TestBuildUpdateMask_Resources(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	// Create state with empty resources list
	state := AppResource{
		App: apps_tf.App{
			Name:      types.StringValue("test-app"),
			Resources: types.ListNull(types.ObjectType{}),
		},
	}

	// Create plan with resources list
	plan := AppResource{
		App: apps_tf.App{
			Name: types.StringValue("test-app"),
			Resources: types.ListValueMust(
				types.ObjectType{},
				[]attr.Value{},
			),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Contains(t, updateMask, "resources")
}

func TestBuildUpdateMask_UserApiScopes(t *testing.T) {
	ctx := context.Background()
	resource := &resourceApp{}

	state := AppResource{
		App: apps_tf.App{
			Name:          types.StringValue("test-app"),
			UserApiScopes: types.ListNull(types.StringType),
		},
	}

	plan := AppResource{
		App: apps_tf.App{
			Name: types.StringValue("test-app"),
			UserApiScopes: types.ListValueMust(
				types.StringType,
				[]attr.Value{
					types.StringValue("read"),
					types.StringValue("write"),
				},
			),
		},
	}

	updateMask := resource.buildUpdateMask(ctx, plan, state)
	assert.Contains(t, updateMask, "user_api_scopes")
}
