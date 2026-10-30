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

// The Budgets API returns alerts and actions in any order and gives an alert new IDs
// when its threshold changes, while Terraform compares list elements by position. Alerts
// are therefore matched by a key built from their content, as if the list were a map:
// orderAlerts stores alerts in the configured order, assignAlertIDs sends each alert with
// the IDs of the existing alert with the same key, and suppressReorderedAlerts ignores
// diffs that only reorder alerts.

// normalizeDecimal removes the trailing zeros that the API adds to decimals.
func normalizeDecimal(v string) string {
	if strings.Contains(v, ".") {
		v = strings.TrimRight(v, "0")
		v = strings.TrimSuffix(v, ".")
	}
	return v
}

func suppressDecimalDiff(_, old, new string, _ *schema.ResourceData) bool {
	return normalizeDecimal(old) == normalizeDecimal(new)
}

// jsonKey returns the JSON of v, which has every field that v sets.
func jsonKey(v any) string {
	raw, _ := json.Marshal(v)
	return string(raw)
}

func actionKey(a billing.ActionConfiguration) string {
	a.ActionConfigurationId = ""
	a.ForceSendFields = nil
	return jsonKey(a)
}

func overrideKey(o billing.PrincipalOverride) string {
	o.OverrideThreshold = normalizeDecimal(o.OverrideThreshold)
	o.ForceSendFields = nil
	return jsonKey(o)
}

// alertKey returns the content of an alert without its IDs. Every other field is part
// of the key, so fields that the API adds later count too. The order of actions and
// principal overrides doesn't matter, and an unset scope type is shared, the API default.
func alertKey(a billing.AlertConfiguration) string {
	actions := sortedKeys(a.ActionConfigurations, actionKey)
	overrides := sortedKeys(a.PrincipalOverrides, overrideKey)
	a.AlertConfigurationId = ""
	a.ActionConfigurations, a.PrincipalOverrides = nil, nil
	a.QuantityThreshold = normalizeDecimal(a.QuantityThreshold)
	if a.ScopeType == "" {
		a.ScopeType = billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared
	}
	a.ForceSendFields = nil
	return jsonKey([]any{a, actions, overrides})
}

func sortedKeys[T any](items []T, key func(T) string) []string {
	keys := make([]string, len(items))
	for i, item := range items {
		keys[i] = key(item)
	}
	slices.Sort(keys)
	return keys
}

// byKey groups items by key, keeping their order.
func byKey[T any](items []T, key func(T) string) map[string][]T {
	groups := map[string][]T{}
	for _, item := range items {
		groups[key(item)] = append(groups[key(item)], item)
	}
	return groups
}

// take removes and returns the first item with the key, if there is one.
func take[T any](groups map[string][]T, key string) (item T, ok bool) {
	if len(groups[key]) == 0 {
		return item, false
	}
	item, groups[key] = groups[key][0], groups[key][1:]
	return item, true
}

// orderLike returns the items in the order of the reference items with the same key,
// followed by the other items.
func orderLike[T any](items, reference []T, key func(T) string) []T {
	if len(items) == 0 {
		return items
	}
	groups := byKey(items, key)
	ordered := make([]T, 0, len(items))
	for _, ref := range reference {
		if item, ok := take(groups, key(ref)); ok {
			ordered = append(ordered, item)
		}
	}
	for _, item := range items {
		if rest, ok := take(groups, key(item)); ok {
			ordered = append(ordered, rest)
		}
	}
	return ordered
}

// orderAlerts orders the alerts of a budget returned by the API, and their actions and
// principal overrides, like the reference alerts with the same content.
func orderAlerts(budget *billing.BudgetConfiguration,
	reference []billing.AlertConfiguration) {
	refs := byKey(reference, alertKey)
	for i := range budget.AlertConfigurations {
		alert := &budget.AlertConfigurations[i]
		if ref, ok := take(refs, alertKey(*alert)); ok {
			alert.ActionConfigurations = orderLike(alert.ActionConfigurations,
				ref.ActionConfigurations, actionKey)
			alert.PrincipalOverrides = orderLike(alert.PrincipalOverrides,
				ref.PrincipalOverrides, overrideKey)
		}
	}
	alerts := orderLike(budget.AlertConfigurations, reference, alertKey)
	budget.AlertConfigurations = alerts
}

// assignAlertIDs gives each desired alert, and its actions, the IDs of the existing alert
// with the same content. The API creates the desired alerts without IDs and deletes the
// existing alerts that are left out.
func assignAlertIDs(desired, existing []billing.AlertConfiguration) {
	current := byKey(existing, alertKey)
	for i := range desired {
		alert := &desired[i]
		match, _ := take(current, alertKey(*alert))
		alert.AlertConfigurationId = match.AlertConfigurationId
		if alert.ScopeType == "" {
			alert.ScopeType = match.ScopeType
		}
		actions := byKey(match.ActionConfigurations, actionKey)
		for j := range alert.ActionConfigurations {
			action := &alert.ActionConfigurations[j]
			found, _ := take(actions, actionKey(*action))
			action.ActionConfigurationId = found.ActionConfigurationId
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
	return slices.Equal(sortedKeys(desired, alertKey), sortedKeys(existing, alertKey))
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
