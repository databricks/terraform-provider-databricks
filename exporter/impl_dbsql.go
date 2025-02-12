package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	tf_sql "github.com/databricks/terraform-provider-databricks/sql"
	tf_sql_api "github.com/databricks/terraform-provider-databricks/sql/api"
)

func listQueries(ic *importContext) error {
	it := ic.workspaceClient.Queries.List(ic.Context, sql.ListQueriesRequest{PageSize: 100})
	i := 0
	for it.HasNext(ic.Context) {
		q, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(q.DisplayName) {
			continue
		}
		// TODO: look if we can create data based on the response, without calling Get
		ic.EmitIfUpdatedAfterIsoString(&resource{
			Resource:    "databricks_query",
			ID:          q.Id,
			Incremental: ic.incremental,
		}, q.UpdateTime, fmt.Sprintf("query '%s'", q.DisplayName))
		if i%50 == 0 {
			log.Printf("[INFO] Imported %d Queries", i)
		}
	}
	log.Printf("[INFO] Listed %d Queries", i)
	return nil
}

func importQuery(ic *importContext, r *resource) error {
	var query tf_sql.QueryStruct
	s := ic.Resources["databricks_query"].Schema
	common.DataToStructPointer(r.Data, s, &query)
	if query.WarehouseId != "" {
		ic.Emit(&resource{
			Resource: "databricks_sql_endpoint",
			ID:       query.WarehouseId,
		})
	}
	// emit queries specified as parameters
	for _, p := range query.Parameters {
		if p.QueryBackedValue != nil {
			ic.Emit(&resource{
				Resource: "databricks_query",
				ID:       p.QueryBackedValue.QueryId,
			})
		}
	}
	ic.emitUserOrServicePrincipal(query.OwnerUserName)
	ic.emitDirectoryOrRepo(query.ParentPath)
	// TODO: r.AddExtraData(ParentDirectoryExtraKey, directoryPath) ?
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/queries/%s", r.ID),
		"query_"+ic.Importables["databricks_query"].Name(ic, r.Data))
	if query.Catalog != "" && query.Schema != "" {
		ic.Emit(&resource{
			Resource: "databricks_schema",
			ID:       fmt.Sprintf("%s.%s", query.Catalog, query.Schema),
		})
	}
	return nil
}

func listSqlEndpoints(ic *importContext) error {
	it := ic.workspaceClient.Warehouses.List(ic.Context, sql.ListWarehousesRequest{})
	i := 0
	for it.HasNext(ic.Context) {
		q, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if !ic.MatchesName(q.Name) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_sql_endpoint",
			ID:       q.Id,
		})
		i++
		log.Printf("[INFO] Imported %d SQL endpoints", i)
	}
	return nil
}

func importSqlEndpoint(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/warehouses/%s", r.ID),
		"sql_endpoint_"+ic.Importables["databricks_sql_endpoint"].Name(ic, r.Data))
	if ic.meAdmin {
		ic.Emit(&resource{
			Resource: "databricks_sql_global_config",
			ID:       tf_sql.GlobalSqlConfigResourceID,
		})
	}
	return nil
}

func listRedashDashboards(ic *importContext) error {
	qs, err := dbsqlListObjects(ic, "/preview/sql/dashboards")
	if err != nil {
		return nil
	}
	for i, q := range qs {
		name := q["name"].(string)
		if !ic.MatchesName(name) {
			continue
		}
		ic.EmitIfUpdatedAfterIsoString(&resource{
			Resource:    "databricks_sql_dashboard",
			ID:          q["id"].(string),
			Incremental: ic.incremental,
		}, q["updated_at"].(string), fmt.Sprintf("dashboard '%s'", name))
		log.Printf("[INFO] Imported %d of %d SQL dashboards", i+1, len(qs))
	}
	return nil
}

func importRedashDashboard(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/dashboards/%s", r.ID),
		"sql_dashboard_"+ic.Importables["databricks_sql_dashboard"].Name(ic, r.Data))
	dashboardID := r.ID
	dashboardAPI := tf_sql.NewDashboardAPI(ic.Context, ic.Client)
	dashboard, err := dashboardAPI.Read(dashboardID)
	if err != nil {
		return err
	}

	ic.emitSqlParentDirectory(dashboard.Parent)
	for _, rv := range dashboard.Widgets {
		var widget tf_sql_api.Widget
		err = json.Unmarshal(rv, &widget)
		if err != nil {
			log.Printf("[WARN] Problems decoding widget for dashboard with ID: %s", dashboardID)
			continue
		}
		widgetID := dashboardID + "/" + widget.ID.String()
		ic.Emit(&resource{
			Resource: "databricks_sql_widget",
			ID:       widgetID,
		})

		if widget.VisualizationID != nil {
			var visualization tf_sql_api.Visualization
			err = json.Unmarshal(widget.Visualization, &visualization)
			if err != nil {
				log.Printf("[WARN] Problems decoding visualization for widget with ID: %s", widget.ID.String())
				continue
			}
			if len(visualization.Query) > 0 {
				var query tf_sql_api.Query
				err = json.Unmarshal(visualization.Query, &query)
				if err != nil {
					log.Printf("[WARN] Problems decoding query for visualization with ID: %s", visualization.ID.String())
					continue
				}
				visualizationID := query.ID + "/" + visualization.ID.String()
				ic.Emit(&resource{
					Resource: "databricks_sql_visualization",
					ID:       visualizationID,
				})
				ic.Emit(&resource{
					Resource: "databricks_query",
					ID:       query.ID,
				})
				sqlEndpointID, err := ic.getSqlEndpoint(query.DataSourceID)
				if err != nil {
					log.Printf("[WARN] Can't find SQL endpoint for data source id %s", query.DataSourceID)
				} else {
					ic.Emit(&resource{
						Resource: "databricks_sql_endpoint",
						ID:       sqlEndpointID,
					})
				}
			} else {
				log.Printf("[DEBUG] Empty query in visualization %v", visualization)
			}
		}
	}
	return nil
}

func listAlerts(ic *importContext) error {
	it := ic.workspaceClient.Alerts.List(ic.Context, sql.ListAlertsRequest{PageSize: 100})
	i := 0
	for it.HasNext(ic.Context) {
		a, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(a.DisplayName) {
			continue
		}
		// TODO: look if we can create data based on the response, without calling Get
		ic.EmitIfUpdatedAfterIsoString(&resource{
			Resource:    "databricks_alert",
			ID:          a.Id,
			Incremental: ic.incremental,
		}, a.UpdateTime, fmt.Sprintf("alert '%s'", a.DisplayName))
		if i%50 == 0 {
			log.Printf("[INFO] Imported %d Alerts", i)
		}
	}
	log.Printf("[INFO] Listed %d Alerts", i)
	return nil
}

func listLakeviewDashboards(ic *importContext) error {
	it := ic.workspaceClient.Lakeview.List(ic.Context, dashboards.ListDashboardsRequest{PageSize: 1000})
	i := 0
	for it.HasNext(ic.Context) {
		d, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(d.DisplayName) {
			continue
		}
		// TODO: add emit for incremental mode. But this information isn't included into the List response
		ic.Emit(&resource{
			Resource: "databricks_dashboard",
			ID:       d.DashboardId,
		})
		if i%100 == 0 {
			log.Printf("[INFO] Processed %d dashboards", i)
		}
	}
	log.Printf("[INFO] Listed %d dashboards", i)
	return nil
}

func importLakeviewDashboard(ic *importContext, r *resource) error {
	path := r.Data.Get("path").(string)
	if ic.isInRepoOrGitFolder(path, false) {
		ic.emitRepoOrGitFolder(path, false)
		return nil
	}
	parts := strings.Split(path, "/")
	plen := len(parts)
	if idx := strings.Index(parts[plen-1], "."); idx != -1 {
		parts[plen-1] = parts[plen-1][:idx] + "_" + r.ID + parts[plen-1][idx:]
	} else {
		parts[plen-1] = parts[plen-1] + "_" + r.ID
	}
	name := fileNameNormalizationRegex.ReplaceAllString(strings.Join(parts, "/")[1:], "_")
	fileName, err := ic.saveFileIn("dashboards", name, []byte(r.Data.Get("serialized_dashboard").(string)))
	if err != nil {
		return err
	}
	r.Data.Set("file_path", fileName)
	r.Data.Set("serialized_dashboard", "")

	ic.emitPermissionsIfNotIgnored(r, "/dashboards/"+r.ID,
		"dashboard_"+ic.Importables["databricks_dashboard"].Name(ic, r.Data))
	parentPath := r.Data.Get("parent_path").(string)
	if parentPath != "" && parentPath != "/" {
		ic.emitDirectoryOrRepo(parentPath)
	}
	warehouseId := r.Data.Get("warehouse_id").(string)
	if warehouseId != "" {
		ic.Emit(&resource{
			Resource: "databricks_sql_endpoint",
			ID:       warehouseId,
		})
	}

	return nil
}
