package finops

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

const (
	shared  = billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared
	perUser = billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypePerUser
)

func email(id string) billing.ActionConfiguration {
	return billing.ActionConfiguration{
		ActionConfigurationId: id,
		ActionType:            billing.ActionConfigurationTypeEmailNotification,
		Target:                "me@example.com",
	}
}

func block(id string) billing.ActionConfiguration {
	return billing.ActionConfiguration{
		ActionConfigurationId: id,
		ActionType:            billing.ActionConfigurationTypeBlockUsage,
	}
}

// alert returns a monthly alert on the list price of usage.
func alert(id string, scope billing.AlertConfigurationScopeType, threshold string,
	actions ...billing.ActionConfiguration) billing.AlertConfiguration {
	a := billing.AlertConfiguration{
		AlertConfigurationId: id,
		ActionConfigurations: actions,
		QuantityThreshold:    threshold,
		ScopeType:            scope,
	}
	a.QuantityType = billing.AlertConfigurationQuantityTypeListPriceDollarsUsd
	a.TimePeriod = billing.AlertConfigurationTimePeriodMonth
	a.TriggerType = billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded
	return a
}

func withOverride(a billing.AlertConfiguration,
	threshold string) billing.AlertConfiguration {
	a.PrincipalOverrides = []billing.PrincipalOverride{
		{PrincipalId: 123, OverrideThreshold: threshold},
	}
	return a
}

// padded returns a threshold as the API returns it.
func padded(threshold string) string {
	return threshold + ".000000000000000000"
}

// configuredAlerts are two shared alerts and a per-user alert with an override.
func configuredAlerts() []billing.AlertConfiguration {
	return []billing.AlertConfiguration{
		alert("", shared, "800", email("")),
		alert("", shared, "1500", block(""), email("")),
		withOverride(alert("", perUser, "200", block("")), "250"),
	}
}

// returnedAlerts are configuredAlerts as the API returns them: reordered, with IDs, and
// with padded decimals.
func returnedAlerts() []billing.AlertConfiguration {
	return []billing.AlertConfiguration{
		alert("block", shared, padded("1500"),
			email("block-email"), block("block-block")),
		withOverride(alert("pu", perUser, padded("200"), block("pu-block")), "250.00"),
		alert("email", shared, padded("800"), email("email-email")),
	}
}

// storedAlerts are returnedAlerts in the configured order.
func storedAlerts() []billing.AlertConfiguration {
	return []billing.AlertConfiguration{
		alert("email", shared, padded("800"), email("email-email")),
		alert("block", shared, padded("1500"),
			block("block-block"), email("block-email")),
		withOverride(alert("pu", perUser, padded("200"), block("pu-block")), "250.00"),
	}
}

// storedIDs are the IDs of storedAlerts, each followed by the IDs of its actions.
var storedIDs = [][]string{
	{"email", "email-email"}, {"block", "block-block", "block-email"}, {"pu", "pu-block"},
}

// alertIDs returns the IDs of the alerts in d, each followed by the IDs of its actions.
func alertIDs(d *schema.ResourceData) (ids [][]string) {
	for _, a := range alertsOf(d, ResourceBudget().Schema) {
		item := []string{a.AlertConfigurationId}
		for _, action := range a.ActionConfigurations {
			item = append(item, action.ActionConfigurationId)
		}
		ids = append(ids, item)
	}
	return ids
}

func budgetWith(alerts []billing.AlertConfiguration) *billing.BudgetConfiguration {
	return &billing.BudgetConfiguration{
		AccountId:             "account_id",
		AlertConfigurations:   slices.Clone(alerts),
		BudgetConfigurationId: "budget_configuration_id",
		DisplayName:           "genie",
	}
}

func budgetState(t *testing.T, budget *billing.BudgetConfiguration) map[string]string {
	r := ResourceBudget()
	d := schema.TestResourceDataRaw(t, r.Schema, nil)
	d.MarkNewResource()
	require.NoError(t, common.StructToData(budget, r.Schema, d))
	d.SetId("account_id|budget_configuration_id")
	return d.State().Attributes
}

// budgetHCL renders a budget with the given alerts, leaving out empty fields.
func budgetHCL(alerts ...billing.AlertConfiguration) string {
	var b strings.Builder
	b.WriteString("display_name = \"genie\"\n")
	for _, a := range alerts {
		fmt.Fprintf(&b, "alert_configurations {\ntime_period = %q\ntrigger_type = %q\n"+
			"quantity_type = %q\nquantity_threshold = %q\n",
			a.TimePeriod, a.TriggerType, a.QuantityType, a.QuantityThreshold)
		if a.ScopeType != "" {
			fmt.Fprintf(&b, "scope_type = %q\n", a.ScopeType)
		}
		for _, o := range a.PrincipalOverrides {
			fmt.Fprintf(&b, "principal_overrides {\nprincipal_id = %d\n"+
				"override_threshold = %q\n}\n", o.PrincipalId, o.OverrideThreshold)
		}
		for _, action := range a.ActionConfigurations {
			fmt.Fprintf(&b, "action_configurations {\naction_type = %q\n",
				action.ActionType)
			if action.Target != "" {
				fmt.Fprintf(&b, "target = %q\n", action.Target)
			}
			b.WriteString("}\n")
		}
		b.WriteString("}\n")
	}
	return b.String()
}

func TestResourceBudgetCreate_KeepsConfiguredOrder(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			created, read := budgetWith(returnedAlerts()), budgetWith(returnedAlerts())
			api := a.GetMockBudgetsAPI().EXPECT()
			api.Create(mock.Anything, mock.Anything).Return(
				&billing.CreateBudgetConfigurationResponse{Budget: created}, nil)
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: read}, nil)
		},
		Resource:  ResourceBudget(),
		Create:    true,
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		HCL:       budgetHCL(configuredAlerts()...),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, storedIDs, alertIDs(d))
}

func TestResourceBudgetRead_KeepsStateOrder(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			read := budgetWith(returnedAlerts())
			api := a.GetMockBudgetsAPI().EXPECT()
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: read}, nil)
		},
		Resource:      ResourceBudget(),
		Read:          true,
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		ID:            "account_id|budget_configuration_id",
		InstanceState: budgetState(t, budgetWith(storedAlerts())),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, storedIDs, alertIDs(d))
}

func TestResourceBudgetPlan_IgnoresAlertOrder(t *testing.T) {
	// The state of budgets created by earlier versions has the alerts in the API order.
	qa.ResourceFixture{
		Resource:      ResourceBudget(),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		InstanceState: budgetState(t, budgetWith(returnedAlerts())),
		HCL:           budgetHCL(configuredAlerts()...),
		ExpectedDiff:  map[string]*terraform.ResourceAttrDiff{},
	}.ApplyNoError(t)
}

func TestResourceBudgetPlan_ShowsAlertChange(t *testing.T) {
	alerts := configuredAlerts()
	alerts[2].QuantityThreshold = "220"
	qa.ResourceFixture{
		Resource:      ResourceBudget(),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		InstanceState: budgetState(t, budgetWith(storedAlerts())),
		HCL:           budgetHCL(alerts...),
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"alert_configurations.2.quantity_threshold": {Old: padded("200"), New: "220"},
		},
	}.ApplyNoError(t)
}

func TestResourceBudgetPlan_IgnoresEmptyFilter(t *testing.T) {
	budget := budgetWith(storedAlerts())
	budget.Filter = &billing.BudgetConfigurationFilter{}
	qa.ResourceFixture{
		Resource:      ResourceBudget(),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		InstanceState: budgetState(t, budget),
		HCL:           budgetHCL(configuredAlerts()...),
		ExpectedDiff:  map[string]*terraform.ResourceAttrDiff{},
	}.ApplyNoError(t)
}

func TestResourceBudgetPlan_ShowsFilterRemoval(t *testing.T) {
	budget := budgetWith(storedAlerts())
	budget.Filter = &billing.BudgetConfigurationFilter{
		WorkspaceId: &billing.BudgetConfigurationFilterWorkspaceIdClause{
			Operator: billing.BudgetConfigurationFilterOperatorIn,
			Values:   []int64{123},
		},
	}
	qa.ResourceFixture{
		Resource:      ResourceBudget(),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		InstanceState: budgetState(t, budget),
		HCL:           budgetHCL(configuredAlerts()...),
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"filter.#":                         {Old: "1", New: "0"},
			"filter.0.workspace_id.#":          {Old: "1", New: "0"},
			"filter.0.workspace_id.0.operator": {Old: "IN", New: "", NewRemoved: true},
			"filter.0.workspace_id.0.values.#": {Old: "1", New: "0"},
			"filter.0.workspace_id.0.values.0": {Old: "123", New: "0", NewRemoved: true},
		},
	}.ApplyNoError(t)
}

// expectUpdate expects a read of the budget with the before alerts, an update with the
// request alerts, and reads with the after alerts.
func expectUpdate(before, request,
	after []billing.AlertConfiguration) func(*mocks.MockAccountClient) {
	return func(a *mocks.MockAccountClient) {
		current, updated, read := budgetWith(before), budgetWith(after), budgetWith(after)
		api := a.GetMockBudgetsAPI().EXPECT()
		api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
			&billing.GetBudgetConfigurationResponse{Budget: current}, nil).Once()
		api.Update(mock.Anything, billing.UpdateBudgetConfigurationRequest{
			Budget: billing.UpdateBudgetConfigurationBudget{
				AccountId:             "account_id",
				AlertConfigurations:   request,
				BudgetConfigurationId: "budget_configuration_id",
				DisplayName:           "genie",
			},
			BudgetId: "budget_configuration_id",
		}).Return(&billing.UpdateBudgetConfigurationResponse{Budget: updated}, nil)
		api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
			&billing.GetBudgetConfigurationResponse{Budget: read}, nil)
	}
}

func TestResourceBudgetUpdate_PairsAlertsByContent(t *testing.T) {
	// Terraform pairs the configured alerts with the state, which has them in the API
	// order, by position. The update must still keep the scope type of every alert.
	alerts := configuredAlerts()
	alerts[0].QuantityThreshold = "850"
	request := []billing.AlertConfiguration{
		alert("email", shared, "850", email("email-email")),
		alert("block", shared, "1500", block("block-block"), email("block-email")),
		withOverride(alert("pu", perUser, "200", block("pu-block")), "250"),
	}
	// The API gives an alert new IDs when its threshold changes.
	after := returnedAlerts()
	after[2] = alert("email-v2", shared, padded("850"), email("email-v2-email"))
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: expectUpdate(returnedAlerts(), request, after),
		Resource:              ResourceBudget(),
		Update:                true,
		AccountID:             "account_id",
		Host:                  "https://accounts.cloud.databricks.com",
		ID:                    "account_id|budget_configuration_id",
		InstanceState:         budgetState(t, budgetWith(returnedAlerts())),
		HCL:                   budgetHCL(alerts...),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, [][]string{{"email-v2", "email-v2-email"},
		{"block", "block-block", "block-email"}, {"pu", "pu-block"}}, alertIDs(d))
}

func TestResourceBudgetUpdate_RemovesFirstAlert(t *testing.T) {
	request := []billing.AlertConfiguration{
		alert("block", shared, "1500", block("block-block"), email("block-email")),
		withOverride(alert("pu", perUser, "200", block("pu-block")), "250"),
	}
	after := returnedAlerts()[:2]
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: expectUpdate(returnedAlerts(), request, after),
		Resource:              ResourceBudget(),
		Update:                true,
		AccountID:             "account_id",
		Host:                  "https://accounts.cloud.databricks.com",
		ID:                    "account_id|budget_configuration_id",
		InstanceState:         budgetState(t, budgetWith(storedAlerts())),
		HCL:                   budgetHCL(configuredAlerts()[1:]...),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, storedIDs[1:], alertIDs(d))
}

func TestResourceBudgetUpdate_UnsetScopeType(t *testing.T) {
	// SDKv2 copies the scope type of the removed per-user alert to the shared alert that
	// takes its position, which must not turn the shared alert into a per-user alert.
	stored := []billing.AlertConfiguration{storedAlerts()[2], storedAlerts()[0]}
	request := []billing.AlertConfiguration{
		alert("email", shared, "800", email("email-email")),
	}
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: expectUpdate(stored, request, storedAlerts()[:1]),
		Resource:              ResourceBudget(),
		Update:                true,
		AccountID:             "account_id",
		Host:                  "https://accounts.cloud.databricks.com",
		ID:                    "account_id|budget_configuration_id",
		InstanceState:         budgetState(t, budgetWith(stored)),
		HCL:                   budgetHCL(alert("", "", "800", email(""))),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, storedIDs[:1], alertIDs(d))
	assert.Equal(t, string(shared), d.Get("alert_configurations.0.scope_type"))
}

func TestResourceBudgetUpdate_RecreatesAlertWithNewScopeType(t *testing.T) {
	// The API rejects changes to the scope type of an alert, so the alert is replaced.
	// The threshold diff is suppressed, so the request keeps the stored decimal.
	request := []billing.AlertConfiguration{
		alert("", perUser, padded("800"), email("")),
	}
	after := []billing.AlertConfiguration{
		alert("new", perUser, padded("800"), email("new-email")),
	}
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: expectUpdate(storedAlerts()[:1], request, after),
		Resource:              ResourceBudget(),
		Update:                true,
		AccountID:             "account_id",
		Host:                  "https://accounts.cloud.databricks.com",
		ID:                    "account_id|budget_configuration_id",
		InstanceState:         budgetState(t, budgetWith(storedAlerts()[:1])),
		HCL:                   budgetHCL(alert("", perUser, "800", email(""))),
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, [][]string{{"new", "new-email"}}, alertIDs(d))
}

func TestPairAlerts_PairsScopedAlertsFirst(t *testing.T) {
	// An alert without a scope type matches both existing alerts, so the per-user alert
	// is paired first to leave the shared alert for it.
	desired := []billing.AlertConfiguration{
		alert("", "", "100", block("")),
		alert("", perUser, "100", block("")),
	}
	existing := []billing.AlertConfiguration{
		alert("u", perUser, "100", block("ub")),
		alert("s", shared, "100", block("sb")),
	}
	assert.Equal(t, []int{1, 0}, pairAlerts(desired, existing, sameAlert))
}
