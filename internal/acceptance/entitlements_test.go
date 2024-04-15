package acceptance

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

type entitlement struct {
	name  string
	value bool
}

func (e entitlement) String() string {
	return fmt.Sprintf("%s = %t", e.name, e.value)
}

func entitlementsStepBuilder(t *testing.T, c **httpclient.ApiClient, r entitlementResource) func(entitlements []entitlement) step {
	return func(entitlements []entitlement) step {
		entitlementsBuf := strings.Builder{}
		for _, entitlement := range entitlements {
			entitlementsBuf.WriteString(fmt.Sprintf("%s\n", entitlement.String()))
		}
		return step{
			Template: fmt.Sprintf(`
			%s
			resource "databricks_entitlements" "entitlements_users" {
				group_id = %s
				%s
			}
		`, r.dataSourceTemplate(), r.tfReference(), entitlementsBuf.String()),
			Check: func(s *terraform.State) error {
				groupId := s.RootModule().Resources["data.databricks_group.example"].Primary.ID
				var res iam.Group
				ctx := context.Background()
				err := (*c).Do(ctx, "GET", fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%s?attributes=entitlements", groupId),
					httpclient.WithResponseUnmarshal(&res))
				assert.NoError(t, err)
				receivedEntitlements := make([]string, 0, len(res.Entitlements))
				for _, entitlement := range res.Entitlements {
					receivedEntitlements = append(receivedEntitlements, entitlement.Value)
				}
				expectedEntitlements := make([]string, 0, len(entitlements))
				for _, entitlement := range entitlements {
					if entitlement.value {
						expectedEntitlements = append(expectedEntitlements, strings.ReplaceAll(entitlement.name, "_", "-"))
					}
				}
				assert.ElementsMatch(t, expectedEntitlements, receivedEntitlements)
				return nil
			},
		}

	}
}

func makeEntitlementsSteps(t *testing.T, r entitlementResource, entitlementsSteps [][]entitlement) []step {
	r.setDisplayName(RandomName("entitlements-"))
	var c *httpclient.ApiClient
	makeEntitlementsStep := entitlementsStepBuilder(t, &c, r)
	steps := make([]step, len(entitlementsSteps))
	for i, entitlements := range entitlementsSteps {
		steps[i] = makeEntitlementsStep(entitlements)
	}
	steps[0].PreConfig = makePreconfig(t, &c, r)
	return steps
}

func makePreconfig(t *testing.T, c **httpclient.ApiClient, r entitlementResource) func() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
	return func() {
		w := databricks.Must(databricks.NewWorkspaceClient())
		r.setWorkspaceClient(w)
		var err error
		*c, err = w.Config.NewApiClient()
		assert.NoError(t, err)
		ctx := context.Background()
		err = r.create(ctx)
		assert.NoError(t, err)
		t.Cleanup(func() {
			r.cleanUp(ctx)
		})
	}
}

func entitlementsTest(t *testing.T, f func(*testing.T, entitlementResource)) {
	loadWorkspaceEnv(t)
	sp := &servicePrincipalResource{}
	if isAzure(t) {
		sp.applicationId = GetEnvOrSkipTest(t, "ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID")
	}
	resources := []entitlementResource{
		// &groupResource{},
		&userResource{},
		// sp,
	}
	for _, r := range resources {
		t.Run(fmt.Sprintf("%T", r), func(t *testing.T) {
			f(t, r)
		})
	}
}

func TestAccEntitlementsAddToEmpty(t *testing.T) {
	entitlementsTest(t, func(t *testing.T, r entitlementResource) {
		steps := makeEntitlementsSteps(t, r, [][]entitlement{
			{},
			{
				{"allow_cluster_create", true},
				{"allow_instance_pool_create", true},
				{"workspace_access", true},
				{"databricks_sql_access", true},
			},
		})
		workspaceLevel(t, steps...)
	})
}

func TestAccEntitlementsSetExplicitlyToFalse(t *testing.T) {
	entitlementsTest(t, func(t *testing.T, r entitlementResource) {
		steps := makeEntitlementsSteps(t, r, [][]entitlement{
			{
				{"allow_cluster_create", false},
				{"allow_instance_pool_create", false},
				{"workspace_access", false},
				{"databricks_sql_access", false},
			},
			{},
			{
				{"allow_cluster_create", false},
				{"allow_instance_pool_create", false},
				{"workspace_access", false},
				{"databricks_sql_access", false},
			},
		})
		workspaceLevel(t, steps...)
	})
}

func TestAccEntitlementsRemoveExisting(t *testing.T) {
	entitlementsTest(t, func(t *testing.T, r entitlementResource) {
		steps := makeEntitlementsSteps(t, r, [][]entitlement{
			{
				{"allow_cluster_create", true},
				{"allow_instance_pool_create", true},
				{"workspace_access", true},
				{"databricks_sql_access", true},
			},
			{},
		})
		workspaceLevel(t, steps...)
	})
}

func TestAccEntitlementsSomeTrueSomeFalse(t *testing.T) {
	entitlementsTest(t, func(t *testing.T, r entitlementResource) {
		steps := makeEntitlementsSteps(t, r, [][]entitlement{
			{
				{"allow_cluster_create", false},
				{"allow_instance_pool_create", false},
				{"workspace_access", true},
				{"databricks_sql_access", true},
			},
		})
		workspaceLevel(t, steps...)
	})
}
