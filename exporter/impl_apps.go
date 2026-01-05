package exporter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/apps"
	app_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/app"
)

func listAppsSettingsCustomTemplates(ic *importContext) error {
	templates, err := ic.workspaceClient.AppsSettings.ListCustomTemplatesAll(ic.Context, apps.ListCustomTemplatesRequest{})
	if err != nil {
		return err
	}
	for _, template := range templates {
		if !ic.MatchesName(template.Name) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_apps_settings_custom_template",
			ID:       template.Name,
		})
	}
	return nil
}

func listApps(ic *importContext) error {
	i := 0
	it := ic.workspaceClient.Apps.List(ic.Context, apps.ListAppsRequest{})
	for it.HasNext(ic.Context) {
		app, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if i%50 == 0 {
			log.Printf("[INFO] Scanned %d apps", i)
		}
		if !ic.MatchesName(app.Name) {
			log.Printf("[INFO] App name %s doesn't match selection %s", app.Name, ic.match)
			continue
		}
		var updateTimeMillis int64
		if app.UpdateTime != "" {
			updateTime, err := time.Parse(time.RFC3339, app.UpdateTime)
			if err != nil {
				return err
			}
			updateTimeMillis = updateTime.UnixMilli()
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_app",
			ID:       app.Name,
		}, updateTimeMillis, fmt.Sprintf("app '%s'", app.Name))
	}
	log.Printf("[INFO] Total %d apps are going to be exported", i)
	return nil
}

func importApp(ic *importContext, r *resource) error {
	// Convert Plugin Framework state to Go SDK struct
	var app apps.App
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper, app_resource.AppResource{}, &app); err != nil {
		return err
	}

	// Emit dependencies for each resource in the app
	for _, res := range app.Resources {
		// SQL Warehouse
		if res.SqlWarehouse != nil && res.SqlWarehouse.Id != "" {
			ic.Emit(&resource{Resource: "databricks_sql_endpoint", ID: res.SqlWarehouse.Id})
		}
		// Serving Endpoint
		if res.ServingEndpoint != nil && res.ServingEndpoint.Name != "" {
			ic.Emit(&resource{Resource: "databricks_model_serving", ID: res.ServingEndpoint.Name})
		}
		// Job
		if res.Job != nil && res.Job.Id != "" {
			ic.Emit(&resource{Resource: "databricks_job", ID: res.Job.Id})
		}
		// Secret - emit both scope and secret
		if res.Secret != nil && res.Secret.Scope != "" && res.Secret.Key != "" {
			// If secrets are enabled, importing of secret scope will emit secrets as well
			ic.Emit(&resource{
				Resource: "databricks_secret_scope",
				ID:       res.Secret.Scope,
			})
		}
		// UC Securable - emit volume if type is VOLUME
		if res.UcSecurable != nil && res.UcSecurable.SecurableType == "VOLUME" && res.UcSecurable.SecurableFullName != "" {
			ic.Emit(&resource{
				Resource: "databricks_volume",
				ID:       res.UcSecurable.SecurableFullName,
			})
		}
		// Database Instance (Lakebase)
		if res.Database != nil && res.Database.InstanceName != "" {
			ic.Emit(&resource{
				Resource: "databricks_database_instance",
				ID:       res.Database.InstanceName,
			})
		}
	}

	// Budget Policy
	if app.BudgetPolicyId != "" {
		ic.Emit(&resource{
			Resource: "databricks_budget_policy",
			ID:       app.BudgetPolicyId,
		})
	}

	// Emit permissions
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/apps/%s", app.Name), "app_"+r.Name)
	return nil
}

// createIsMatchingScopeAndKey creates a validation function that matches secret scope and key together
// to ensure we reference the correct secret when the same key exists in multiple scopes.
// This works with both SDKv2 and Plugin Framework resources by using DataWrapper when available.
func createIsMatchingScopeAndKey(scope_name_attr, key_name_attr string) func(ic *importContext, res *resource,
	ra *resourceApproximation, origPath string) bool {
	return func(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
		// Get scope and key from the source resource (databricks_app)
		// We need to adjust the path if we're matching on the key attribute
		new_scope_name_attr := scope_name_attr
		if strings.HasSuffix(origPath, "."+key_name_attr) {
			new_scope_name_attr = strings.TrimSuffix(origPath, key_name_attr) + scope_name_attr
		}
		log.Printf("[DEBUG] Matching scope and key for resource %s, scope_name_attr=%s, key_name_attr=%s, origPath=%s, new_scope_name_attr=%s", res.Resource, new_scope_name_attr, key_name_attr, origPath, new_scope_name_attr)
		// Use DataWrapper if available (works for both SDKv2 and Plugin Framework),
		// otherwise fall back to res.Data (SDKv2 only)
		var res_scope_value, res_key_value interface{}
		if res.DataWrapper != nil {
			res_scope_value = res.DataWrapper.Get(new_scope_name_attr)
			res_key_value = res.DataWrapper.Get(origPath)
		} else if res.Data != nil {
			res_scope_value = res.Data.Get(new_scope_name_attr)
			res_key_value = res.Data.Get(origPath)
		} else {
			log.Printf("[WARN] Neither DataWrapper nor Data available for resource %s", res.Resource)
			return false
		}

		// Convert to string, handling nil values
		res_scope_name, ok := res_scope_value.(string)
		if !ok || res_scope_name == "" {
			log.Printf("[DEBUG] Scope attribute '%s' not found or empty for resource %s", new_scope_name_attr, res.Resource)
			return false
		}

		res_key_name, ok := res_key_value.(string)
		if !ok || res_key_name == "" {
			log.Printf("[DEBUG] Key attribute '%s' not found or empty for resource %s", origPath, res.Resource)
			return false
		}

		// Get scope and key from the target resource approximation (databricks_secret)
		ra_scope_name, scope_found := ra.Get("scope")
		ra_key_name, key_found := ra.Get("key")
		if !scope_found || !key_found {
			log.Printf("[WARN] Can't find attributes in approximation: %s %s, scope='%v' (found? %v) key='%v' (found? %v). Resource: %s, scope='%s', key='%s'",
				ra.Type, ra.Name, ra_scope_name, scope_found, ra_key_name, key_found, res.Resource, res_scope_name, res_key_name)
			return false
		}

		// Match only if both scope and key match
		result := ra_scope_name.(string) == res_scope_name && ra_key_name.(string) == res_key_name
		log.Printf("[DEBUG] Matching scope and key for resource %s: scope='%s' (match? %v), key='%s' (match? %v), result=%v",
			res.Resource, res_scope_name, ra_scope_name.(string) == res_scope_name, res_key_name, ra_key_name.(string) == res_key_name, result)
		return result
	}
}
