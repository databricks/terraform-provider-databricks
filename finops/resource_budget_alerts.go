package finops

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

// The Budgets API returns alerts and actions in an order unrelated to the request, while
// Terraform pairs list elements with the state by position. The functions below pair
// alerts, actions, and principal overrides by content or ID instead.

// match reports whether a desired item can be paired with an existing item.
type match[T any] func(desired, existing T) bool

// pair returns the index of the existing item paired with each desired item, or -1. Each
// existing item is paired at most once, and earlier matches take precedence.
func pair[T any](desired, existing []T, matches ...match[T]) []int {
	pairs := make([]int, len(desired))
	for i := range pairs {
		pairs[i] = -1
	}
	used := make([]bool, len(existing))
	for _, m := range matches {
		for i := range desired {
			for j := 0; pairs[i] == -1 && j < len(existing); j++ {
				if !used[j] && m(desired[i], existing[j]) {
					pairs[i], used[j] = j, true
				}
			}
		}
	}
	return pairs
}

// pairedAll reports whether pairs pair every one of n existing items.
func pairedAll(pairs []int, n int) bool {
	return len(pairs) == n && !slices.Contains(pairs, -1)
}

// isPermutation reports whether desired and existing have the same items in any order.
func isPermutation[T any](desired, existing []T, same match[T]) bool {
	return pairedAll(pair(desired, existing, same), len(existing))
}

// reorder returns the paired existing items in the order of the desired items, followed
// by the unpaired ones.
func reorder[T any](existing []T, pairs []int) []T {
	if len(existing) == 0 {
		return existing
	}
	ordered := make([]T, 0, len(existing))
	used := make([]bool, len(existing))
	for _, j := range pairs {
		if j != -1 {
			ordered = append(ordered, existing[j])
			used[j] = true
		}
	}
	for j, item := range existing {
		if !used[j] {
			ordered = append(ordered, item)
		}
	}
	return ordered
}

// sameDecimal reports whether the decimals are equal, ignoring the trailing zeros that
// the API adds.
func sameDecimal(a, b string) bool {
	normalize := func(v string) string {
		if strings.Contains(v, ".") {
			v = strings.TrimRight(v, "0")
			v = strings.TrimSuffix(v, ".")
		}
		return v
	}
	return normalize(a) == normalize(b)
}

func suppressDecimalDiff(_, old, new string, _ *schema.ResourceData) bool {
	return sameDecimal(old, new)
}

func sameActionType(desired, existing billing.ActionConfiguration) bool {
	return desired.ActionType == existing.ActionType
}

func sameAction(desired, existing billing.ActionConfiguration) bool {
	return sameActionType(desired, existing) && desired.Target == existing.Target
}

// pairActions pairs identical actions first, then actions with the same ID, then actions
// of the same type, as the API rejects changes to the type of an action.
func pairActions(desired, existing []billing.ActionConfiguration) []int {
	sameID := func(d, e billing.ActionConfiguration) bool {
		return d.ActionConfigurationId != "" &&
			d.ActionConfigurationId == e.ActionConfigurationId && sameActionType(d, e)
	}
	return pair(desired, existing, sameAction, sameID, sameActionType)
}

func samePrincipal(desired, existing billing.PrincipalOverride) bool {
	return desired.PrincipalId == existing.PrincipalId
}

func sameOverride(desired, existing billing.PrincipalOverride) bool {
	return samePrincipal(desired, existing) &&
		sameDecimal(desired.OverrideThreshold, existing.OverrideThreshold)
}

// compatibleAlerts reports whether existing can be updated to desired, as the API rejects
// changes to these fields. An empty scope type matches any scope type.
func compatibleAlerts(desired, existing billing.AlertConfiguration) bool {
	return (desired.ScopeType == "" || desired.ScopeType == existing.ScopeType) &&
		desired.TimePeriod == existing.TimePeriod &&
		desired.TriggerType == existing.TriggerType &&
		desired.QuantityType == existing.QuantityType
}

func sameAlertID(desired, existing billing.AlertConfiguration) bool {
	return desired.AlertConfigurationId != "" &&
		desired.AlertConfigurationId == existing.AlertConfigurationId &&
		compatibleAlerts(desired, existing)
}

func sameAlert(desired, existing billing.AlertConfiguration) bool {
	actions, overrides := existing.ActionConfigurations, existing.PrincipalOverrides
	return compatibleAlerts(desired, existing) &&
		sameDecimal(desired.QuantityThreshold, existing.QuantityThreshold) &&
		isPermutation(desired.ActionConfigurations, actions, sameAction) &&
		isPermutation(desired.PrincipalOverrides, overrides, sameOverride)
}

// pairAlerts pairs desired alerts with existing ones. Alerts with a scope type are paired
// first, as alerts without one match any scope type.
func pairAlerts(desired, existing []billing.AlertConfiguration,
	matches ...match[billing.AlertConfiguration]) []int {
	scoped := make([]match[billing.AlertConfiguration], 0, 2*len(matches))
	for _, m := range matches {
		scoped = append(scoped, func(d, e billing.AlertConfiguration) bool {
			return d.ScopeType != "" && m(d, e)
		})
	}
	return pair(desired, existing, append(scoped, matches...)...)
}

// orderAlerts orders the alerts of a budget returned by the API, and their actions and
// principal overrides, like the reference alerts. Alerts without a reference go last.
func orderAlerts(budget *billing.BudgetConfiguration,
	reference []billing.AlertConfiguration) {
	alerts := budget.AlertConfigurations
	pairs := pairAlerts(reference, alerts, sameAlert, sameAlertID, compatibleAlerts)
	for i, j := range pairs {
		if j == -1 {
			continue
		}
		ref, alert := reference[i], &alerts[j]
		alert.ActionConfigurations = reorder(alert.ActionConfigurations,
			pairActions(ref.ActionConfigurations, alert.ActionConfigurations))
		alert.PrincipalOverrides = reorder(alert.PrincipalOverrides,
			pair(ref.PrincipalOverrides, alert.PrincipalOverrides, samePrincipal))
	}
	budget.AlertConfigurations = reorder(alerts, pairs)
}

// assignAlertIDs gives desired alerts and actions the IDs of the existing ones they
// update, so that an update never changes immutable fields. The API creates alerts and
// actions without an ID and deletes the existing ones that are left out.
func assignAlertIDs(desired, existing []billing.AlertConfiguration) {
	pairs := pairAlerts(desired, existing, sameAlert, sameAlertID, compatibleAlerts)
	for i, j := range pairs {
		var current billing.AlertConfiguration
		if j != -1 {
			current = existing[j]
		}
		alert := &desired[i]
		alert.AlertConfigurationId = current.AlertConfigurationId
		if alert.ScopeType == "" {
			alert.ScopeType = current.ScopeType
		}
		assignActionIDs(alert.ActionConfigurations, current.ActionConfigurations)
	}
}

func assignActionIDs(desired, existing []billing.ActionConfiguration) {
	for i, j := range pairActions(desired, existing) {
		desired[i].ActionConfigurationId = ""
		if j != -1 {
			desired[i].ActionConfigurationId = existing[j].ActionConfigurationId
		}
	}
}

// clearUnsetScopeTypes clears the scope types missing from the configuration, which
// SDKv2 copies from the alert at the same position in the state.
func clearUnsetScopeTypes(d *schema.ResourceData, alerts []billing.AlertConfiguration) {
	for i := range alerts {
		path := cty.GetAttrPath("alert_configurations").IndexInt(i).GetAttr("scope_type")
		if v, diags := d.GetRawConfigAt(path); !diags.HasError() && v.IsNull() {
			alerts[i].ScopeType = ""
		}
	}
}

// suppressReorderedAlerts suppresses the diff when the configuration has the same alerts
// as the state in a different order.
func suppressReorderedAlerts(_, _, _ string, d *schema.ResourceData) bool {
	before, after := d.GetChange("alert_configurations")
	existing, err := alertsFromList(before)
	if err != nil {
		return false
	}
	desired, err := alertsFromList(after)
	if err != nil {
		return false
	}
	return pairedAll(pairAlerts(desired, existing, sameAlert), len(existing))
}

// alertsFromList converts alert_configurations of the state or the configuration, whose
// keys are the JSON names of the SDK fields.
func alertsFromList(v any) (alerts []billing.AlertConfiguration, err error) {
	raw, err := json.Marshal(v)
	if err == nil {
		err = json.Unmarshal(raw, &alerts)
	}
	return alerts, err
}

// alertsOf returns the alerts in d, in their order.
func alertsOf(d *schema.ResourceData,
	s map[string]*schema.Schema) []billing.AlertConfiguration {
	var budget billing.BudgetConfiguration
	common.DataToStructPointer(d, s, &budget)
	return budget.AlertConfigurations
}
