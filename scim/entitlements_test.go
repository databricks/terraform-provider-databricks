package scim

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

type entitlementScenario struct {
	name          string
	create        bool
	initialValues map[string]bool
	values        map[string]bool
}

type entitlementEntity struct {
	name     string
	idField  string
	idPrefix string
	endpoint string
}

var entitlementScenarios = []entitlementScenario{
	{
		name:   "add_to_empty",
		values: allEntitlementValues(true),
	},
	{
		name:          "remove_existing",
		initialValues: allEntitlementValues(true),
		values:        allEntitlementValues(false),
	},
	{
		name:   "set_explicitly_to_false",
		create: true,
		values: allEntitlementValues(false),
	},
	{
		name:   "some_true_some_false",
		create: true,
		values: map[string]bool{
			"allow_cluster_create":       false,
			"allow_instance_pool_create": false,
			"databricks_sql_access":      true,
			"workspace_access":           true,
		},
	},
}

var entitlementEntities = []entitlementEntity{
	{
		name:     "group",
		idField:  "group_id",
		idPrefix: "group",
		endpoint: "/api/2.0/preview/scim/v2/Groups/abc",
	},
	{
		name:     "user",
		idField:  "user_id",
		idPrefix: "user",
		endpoint: "/api/2.0/preview/scim/v2/Users/abc",
	},
	{
		name:     "service_principal",
		idField:  "service_principal_id",
		idPrefix: "spn",
		endpoint: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
	},
}

func allEntitlementValues(value bool) map[string]bool {
	return map[string]bool{
		"allow_cluster_create":       value,
		"allow_instance_pool_create": value,
		"databricks_sql_access":      value,
		"workspace_access":           value,
	}
}

func entitlementComplexValues(values map[string]bool) []ComplexValue {
	result := make([]ComplexValue, 0, len(values))
	for _, entitlement := range possibleEntitlements {
		if values[entitlementMapping[entitlement]] {
			result = append(result, ComplexValue{Value: entitlement})
		}
	}
	if len(result) == 0 {
		return []ComplexValue{{Value: ""}}
	}
	return result
}

func entitlementHCL(entity entitlementEntity, values map[string]bool) string {
	var hcl strings.Builder
	fmt.Fprintf(&hcl, "%s = \"abc\"\n", entity.idField)
	for _, entitlement := range possibleEntitlements {
		field := entitlementMapping[entitlement]
		if value, ok := values[field]; ok {
			fmt.Fprintf(&hcl, "%s = %t\n", field, value)
		}
	}
	return hcl.String()
}

func entitlementState(entity entitlementEntity, values map[string]bool) map[string]string {
	state := map[string]string{entity.idField: "abc"}
	for field, value := range values {
		state[field] = fmt.Sprintf("%t", value)
	}
	return state
}

func TestResourceEntitlementsLifecycleScenarios(t *testing.T) {
	for _, scenario := range entitlementScenarios {
		for _, entity := range entitlementEntities {
			t.Run(scenario.name+"/"+entity.name, func(t *testing.T) {
				remoteEntitlements := entitlementComplexValues(scenario.values)
				if len(remoteEntitlements) == 1 && remoteEntitlements[0].Value == "" {
					remoteEntitlements = nil
				}
				fixture := qa.ResourceFixture{
					Fixtures: []qa.HTTPFixture{
						{
							Method:   "PATCH",
							Resource: entity.endpoint,
							ExpectedRequest: PatchRequestComplexValue([]patchOperation{{
								"replace", "entitlements", entitlementComplexValues(scenario.values),
							}}),
							Response: map[string]any{"id": "abc"},
						},
						{
							Method:   "GET",
							Resource: entity.endpoint + "?attributes=entitlements",
							Response: map[string]any{
								"id":           "abc",
								"entitlements": remoteEntitlements,
							},
						},
					},
					Resource:      ResourceEntitlements(),
					HCL:           entitlementHCL(entity, scenario.values),
					Create:        scenario.create,
					Update:        !scenario.create,
					InstanceState: entitlementState(entity, scenario.initialValues),
				}
				if !scenario.create {
					fixture.ID = entity.idPrefix + "/abc"
				}
				expected := map[string]any{"id": entity.idPrefix + "/abc"}
				for field := range allEntitlementValues(false) {
					expected[field] = scenario.values[field]
				}
				fixture.ApplyAndExpectData(t, expected)
			})
		}
	}
}
