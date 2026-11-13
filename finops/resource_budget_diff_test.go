package finops

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

func teamTagClause() billing.BudgetConfigurationFilterTagClause {
	return billing.BudgetConfigurationFilterTagClause{
		Key: "Team", Value: &billing.BudgetConfigurationFilterClause{
			Operator: billing.BudgetConfigurationFilterOperatorIn, Values: []string{"FinOps"},
		},
	}
}

// Exercise SDKv2 planning against existing state without invoking CRUD or an API.
func TestResourceBudgetDiff(t *testing.T) {
	for _, tc := range []struct {
		name     string
		before   func(*billing.BudgetConfiguration)
		change   func(*billing.BudgetConfiguration)
		wantNew  bool
		wantDiff bool
	}{
		{name: "workspace values", change: func(b *billing.BudgetConfiguration) {
			b.Filter.WorkspaceId.Values = []int64{987654321}
		}, wantNew: true, wantDiff: true},
		{name: "remove workspace clause", change: func(b *billing.BudgetConfiguration) {
			b.Filter.WorkspaceId = nil
		}, wantNew: true, wantDiff: true},
		{name: "tag values", change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags[0].Value.Values = []string{"Production"}
		}, wantNew: true, wantDiff: true},
		{name: "tag key", change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags[0].Key = "Team"
		}, wantNew: true, wantDiff: true},
		{name: "add tag clause", change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = append(b.Filter.Tags, billing.BudgetConfigurationFilterTagClause{
				Key: "Team", Value: &billing.BudgetConfigurationFilterClause{
					Operator: billing.BudgetConfigurationFilterOperatorIn, Values: []string{"FinOps"},
				},
			})
		}, wantNew: true, wantDiff: true},
		{name: "remove tag clause", change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = nil
		}, wantNew: true, wantDiff: true},
		{name: "add workspace clause", before: func(b *billing.BudgetConfiguration) {
			b.Filter.WorkspaceId = nil
		}, wantNew: true, wantDiff: true},
		{name: "add filter", before: func(b *billing.BudgetConfiguration) {
			b.Filter = nil
		}, wantNew: true, wantDiff: true},
		{name: "remove filter", change: func(b *billing.BudgetConfiguration) {
			b.Filter = nil
		}, wantNew: true, wantDiff: true},
		{name: "unchanged"},
		{name: "equivalent empty tag clauses", before: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = nil
		}, change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = []billing.BudgetConfigurationFilterTagClause{}
		}},
		{name: "add workspace value", change: func(b *billing.BudgetConfiguration) {
			b.Filter.WorkspaceId.Values = append(b.Filter.WorkspaceId.Values, 987654321)
		}, wantNew: true, wantDiff: true},
		{name: "remove workspace value", before: func(b *billing.BudgetConfiguration) {
			b.Filter.WorkspaceId.Values = append(b.Filter.WorkspaceId.Values, 987654321)
		}, wantNew: true, wantDiff: true},
		{name: "tag order", before: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = append(b.Filter.Tags, teamTagClause())
		}, change: func(b *billing.BudgetConfiguration) {
			b.Filter.Tags = []billing.BudgetConfigurationFilterTagClause{teamTagClause(), b.Filter.Tags[0]}
		}, wantNew: true, wantDiff: true},
		{name: "display name", change: func(b *billing.BudgetConfiguration) {
			b.DisplayName = "renamed"
		}, wantDiff: true},
		{name: "alert threshold", change: func(b *billing.BudgetConfiguration) {
			b.AlertConfigurations[0].QuantityThreshold = "940"
		}, wantDiff: true},
		{name: "alert target", change: func(b *billing.BudgetConfiguration) {
			b.AlertConfigurations[0].ActionConfigurations[0].Target = "other@databricks.com"
		}, wantDiff: true},
		{name: "equivalent threshold", change: func(b *billing.BudgetConfiguration) {
			b.AlertConfigurations[0].QuantityThreshold = "840.84"
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := ResourceBudget().ToResource()
			old, next := getTestBudget(), getTestBudget()
			if tc.before != nil {
				tc.before(old)
			}
			if tc.change != nil {
				tc.change(next)
			}
			stateData := schema.TestResourceDataRaw(t, r.Schema, nil)
			stateData.MarkNewResource()
			require.NoError(t, common.StructToData(old, r.Schema, stateData))
			stateData.SetId("account_id|budget_configuration_id")
			configData := schema.TestResourceDataRaw(t, r.Schema, nil)
			configData.MarkNewResource()
			require.NoError(t, common.StructToData(next, r.Schema, configData))
			config := terraform.NewResourceConfigRaw(map[string]any{
				"display_name":         configData.Get("display_name"),
				"filter":               configData.Get("filter"),
				"alert_configurations": configData.Get("alert_configurations"),
			})
			require.False(t, r.Validate(config).HasError(), "%v", r.Validate(config))
			diff, err := r.Diff(context.Background(), stateData.State(), config, nil)
			require.NoError(t, err)
			assert.Equal(t, tc.wantDiff, !diff.Empty(), "unexpected planned changes")
			assert.Equal(t, tc.wantNew, diff.RequiresNew(), "unexpected replacement decision")
			if diff != nil {
				for key, attr := range diff.Attributes {
					if attr.Old != attr.New {
						t.Logf("%s: %q -> %q (requires_new=%t)", key, attr.Old, attr.New, attr.RequiresNew)
					}
				}
			}
		})
	}
}
