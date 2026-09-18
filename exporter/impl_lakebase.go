package exporter

import (
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	database_instance_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_instance"
	postgres_branch_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_branch"
	postgres_catalog_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_catalog"
	postgres_cdf_config_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_cdf_config"
	postgres_data_api_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_data_api"
	postgres_database_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_database"
	postgres_endpoint_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_endpoint"
	postgres_project_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_project"
	postgres_role_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_role"
	postgres_synced_table_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/postgres_synced_table"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func listDatabaseInstances(ic *importContext) error {
	instances, err := ic.workspaceClient.Database.ListDatabaseInstancesAll(ic.Context, database.ListDatabaseInstancesRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, instance := range instances {
		if !ic.MatchesName(instance.Name) {
			log.Printf("[INFO] Skipping database instance %s because it doesn't match %s", instance.Name, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_database_instance",
			ID:       instance.Name,
		}, 0, fmt.Sprintf("database instance '%s'", instance.Name))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Database Instances", i)
	}
	return nil
}

func importDatabaseInstance(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex) including custom_tags!
	copyEffectiveFieldsToInputFieldsWithConverters[database_instance_resource.DatabaseInstance](
		ic, r, database.DatabaseInstance{})

	// Emit permissions for the database instance
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/database-instances/%s", r.ID),
		"database_instance_"+r.Name)
	return nil
}

func postgresResourceNameFromPath(fullName string) string {
	parts := strings.Split(fullName, "/")
	resourceNameParts := []string{}
	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case "projects", "branches", "endpoints", "databases", "roles", "cdf-configs", "cdf-statuses", "catalogs", "synced_tables":
			if i+1 < len(parts) {
				resourceNameParts = append(resourceNameParts, parts[i+1])
				i++
			}
		case "data-api":
			resourceNameParts = append(resourceNameParts, "data_api")
		}
	}
	name := strings.Join(resourceNameParts, "_")
	if name == "" {
		name = fullName
	}
	return strings.NewReplacer("-", "_", ".", "_", "/", "_").Replace(name)
}

func postgresPathPart(fullName, marker string) string {
	parts := strings.Split(fullName, "/")
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] == marker {
			return parts[i+1]
		}
	}
	return ""
}

func postgresPathParentBefore(fullName, marker string) string {
	parts := strings.Split(fullName, "/")
	for i, part := range parts {
		if part == marker {
			return strings.Join(parts[:i], "/")
		}
	}
	return ""
}

func postgresBranchPathFromName(fullName string) string {
	projectId := postgresPathPart(fullName, "projects")
	branchId := postgresPathPart(fullName, "branches")
	if projectId == "" || branchId == "" {
		return ""
	}
	return fmt.Sprintf("projects/%s/branches/%s", projectId, branchId)
}

func writePostgresTfStruct(ic *importContext, r *resource, tfStruct any) error {
	pfWrapper, ok := r.DataWrapper.(*PluginFrameworkResourceData)
	if !ok {
		log.Printf("[WARN] Unable to write postgres import state: wrapper is not PluginFrameworkResourceData for %s", r.ID)
		return nil
	}
	diags := pfWrapper.state.Set(ic.Context, tfStruct)
	if diags.HasError() {
		return fmt.Errorf("failed to write postgres import state for %s: %v", r.ID, diags)
	}
	return nil
}

func normalizePostgresProjectState(ic *importContext, r *resource) error {
	var tfProject postgres_project_resource.Project
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfProject); err != nil {
		return err
	}
	var project postgres.Project
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfProject, &project)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres project state for %s: %v", r.ID, diags)
	}

	if project.ProjectId == "" {
		project.ProjectId = postgresPathPart(project.Name, "projects")
	}
	if project.ProjectId == "" && project.Status != nil {
		project.ProjectId = project.Status.ProjectId
	}
	if project.Spec == nil && project.Status != nil {
		project.Spec = &postgres.ProjectSpec{
			BudgetPolicyId:           project.Status.BudgetPolicyId,
			CustomTags:               project.Status.CustomTags,
			DisplayName:              project.Status.DisplayName,
			EnablePgNativeLogin:      project.Status.EnablePgNativeLogin,
			HistoryRetentionDuration: project.Status.HistoryRetentionDuration,
			PgVersion:                project.Status.PgVersion,
		}
	}

	var tfProjectOut postgres_project_resource.Project
	diags = converters.GoSdkToTfSdkStruct(ic.Context, project, &tfProjectOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres project state back for %s: %v", r.ID, diags)
	}
	tfProjectOut.PurgeOnDelete = tfProject.PurgeOnDelete
	tfProjectOut.ProviderConfig = tfProject.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfProjectOut)
}

func normalizePostgresBranchState(ic *importContext, r *resource) error {
	var tfBranch postgres_branch_resource.Branch
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfBranch); err != nil {
		return err
	}
	var branch postgres.Branch
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfBranch, &branch)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres branch state for %s: %v", r.ID, diags)
	}

	if branch.BranchId == "" {
		branch.BranchId = postgresPathPart(branch.Name, "branches")
	}
	if branch.BranchId == "" && branch.Status != nil {
		branch.BranchId = branch.Status.BranchId
	}
	if branch.Parent == "" {
		projectId := postgresPathPart(branch.Name, "projects")
		if projectId != "" {
			branch.Parent = "projects/" + projectId
		}
	}
	if branch.Spec == nil && branch.Status != nil {
		branch.Spec = &postgres.BranchSpec{
			ExpireTime:       branch.Status.ExpireTime,
			IsProtected:      branch.Status.IsProtected,
			SourceBranch:     branch.Status.SourceBranch,
			SourceBranchLsn:  branch.Status.SourceBranchLsn,
			SourceBranchTime: branch.Status.SourceBranchTime,
		}
		if branch.Status.ExpireTime == nil {
			branch.Spec.NoExpiry = true
		}
	}

	var tfBranchOut postgres_branch_resource.Branch
	diags = converters.GoSdkToTfSdkStruct(ic.Context, branch, &tfBranchOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres branch state back for %s: %v", r.ID, diags)
	}
	tfBranchOut.PurgeOnDelete = tfBranch.PurgeOnDelete
	tfBranchOut.ReplaceExisting = tfBranch.ReplaceExisting
	tfBranchOut.ProviderConfig = tfBranch.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfBranchOut)
}

func normalizePostgresEndpointState(ic *importContext, r *resource) error {
	var tfEndpoint postgres_endpoint_resource.Endpoint
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfEndpoint); err != nil {
		return err
	}
	var endpoint postgres.Endpoint
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfEndpoint, &endpoint)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres endpoint state for %s: %v", r.ID, diags)
	}

	if endpoint.EndpointId == "" {
		endpoint.EndpointId = postgresPathPart(endpoint.Name, "endpoints")
	}
	if endpoint.EndpointId == "" && endpoint.Status != nil {
		endpoint.EndpointId = endpoint.Status.EndpointId
	}
	if endpoint.Parent == "" {
		projectId := postgresPathPart(endpoint.Name, "projects")
		branchId := postgresPathPart(endpoint.Name, "branches")
		if projectId != "" && branchId != "" {
			endpoint.Parent = fmt.Sprintf("projects/%s/branches/%s", projectId, branchId)
		}
	}
	if endpoint.Spec == nil && endpoint.Status != nil {
		endpoint.Spec = &postgres.EndpointSpec{
			AutoscalingLimitMaxCu:  endpoint.Status.AutoscalingLimitMaxCu,
			AutoscalingLimitMinCu:  endpoint.Status.AutoscalingLimitMinCu,
			Disabled:               endpoint.Status.Disabled,
			EndpointType:           endpoint.Status.EndpointType,
			Settings:               endpoint.Status.Settings,
			SuspendTimeoutDuration: endpoint.Status.SuspendTimeoutDuration,
		}
		if endpoint.Status.Group != nil {
			endpoint.Spec.Group = &postgres.EndpointGroupSpec{
				EnableReadableSecondaries: endpoint.Status.Group.EnableReadableSecondaries,
				Max:                       endpoint.Status.Group.Max,
				Min:                       endpoint.Status.Group.Min,
			}
		}
	}

	var tfEndpointOut postgres_endpoint_resource.Endpoint
	diags = converters.GoSdkToTfSdkStruct(ic.Context, endpoint, &tfEndpointOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres endpoint state back for %s: %v", r.ID, diags)
	}
	tfEndpointOut.ReplaceExisting = tfEndpoint.ReplaceExisting
	tfEndpointOut.ProviderConfig = tfEndpoint.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfEndpointOut)
}

func normalizePostgresDatabaseState(ic *importContext, r *resource) error {
	var tfDatabase postgres_database_resource.Database
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfDatabase); err != nil {
		return err
	}
	var database postgres.Database
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfDatabase, &database)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres database state for %s: %v", r.ID, diags)
	}

	if database.DatabaseId == "" {
		database.DatabaseId = postgresPathPart(database.Name, "databases")
	}
	if database.DatabaseId == "" && database.Status != nil {
		database.DatabaseId = database.Status.DatabaseId
	}
	if database.Parent == "" {
		database.Parent = postgresBranchPathFromName(database.Name)
	}
	if database.Spec == nil && database.Status != nil {
		database.Spec = &postgres.DatabaseDatabaseSpec{
			PostgresDatabase: database.Status.PostgresDatabase,
			Role:             database.Status.Role,
		}
	}

	var tfDatabaseOut postgres_database_resource.Database
	diags = converters.GoSdkToTfSdkStruct(ic.Context, database, &tfDatabaseOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres database state back for %s: %v", r.ID, diags)
	}
	tfDatabaseOut.ReplaceExisting = tfDatabase.ReplaceExisting
	tfDatabaseOut.ProviderConfig = tfDatabase.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfDatabaseOut)
}

func normalizePostgresRoleState(ic *importContext, r *resource) error {
	var tfRole postgres_role_resource.Role
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfRole); err != nil {
		return err
	}
	var role postgres.Role
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfRole, &role)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres role state for %s: %v", r.ID, diags)
	}

	if role.RoleId == "" {
		role.RoleId = postgresPathPart(role.Name, "roles")
	}
	if role.RoleId == "" && role.Status != nil {
		role.RoleId = role.Status.RoleId
	}
	if role.Parent == "" {
		role.Parent = postgresBranchPathFromName(role.Name)
	}
	if role.Spec == nil && role.Status != nil {
		role.Spec = &postgres.RoleRoleSpec{
			Attributes:      role.Status.Attributes,
			AuthMethod:      role.Status.AuthMethod,
			IdentityType:    role.Status.IdentityType,
			MembershipRoles: role.Status.MembershipRoles,
			PostgresRole:    role.Status.PostgresRole,
		}
	}

	var tfRoleOut postgres_role_resource.Role
	diags = converters.GoSdkToTfSdkStruct(ic.Context, role, &tfRoleOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres role state back for %s: %v", r.ID, diags)
	}
	tfRoleOut.ReplaceExisting = tfRole.ReplaceExisting
	tfRoleOut.ProviderConfig = tfRole.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfRoleOut)
}

func normalizePostgresDataApiState(ic *importContext, r *resource) error {
	var tfDataApi postgres_data_api_resource.DataApi
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfDataApi); err != nil {
		return err
	}
	var dataApi postgres.DataApi
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfDataApi, &dataApi)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres data api state for %s: %v", r.ID, diags)
	}

	if dataApi.Parent == "" {
		dataApi.Parent = postgresPathParentBefore(dataApi.Name, "data-api")
	}
	if dataApi.Spec == nil && dataApi.Status != nil {
		dataApi.Spec = &postgres.DataApiDataApiSpec{
			DbAggregatesEnabled:      dataApi.Status.DbAggregatesEnabled,
			DbExtraSearchPath:        dataApi.Status.DbExtraSearchPath,
			DbMaxRows:                dataApi.Status.DbMaxRows,
			DbSchemas:                dataApi.Status.DbSchemas,
			JwtCacheMaxLifetime:      dataApi.Status.JwtCacheMaxLifetime,
			JwtRoleClaimKey:          dataApi.Status.JwtRoleClaimKey,
			OpenapiMode:              dataApi.Status.OpenapiMode,
			ServerCorsAllowedOrigins: dataApi.Status.ServerCorsAllowedOrigins,
			ServerTimingEnabled:      dataApi.Status.ServerTimingEnabled,
		}
	}

	var tfDataApiOut postgres_data_api_resource.DataApi
	diags = converters.GoSdkToTfSdkStruct(ic.Context, dataApi, &tfDataApiOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres data api state back for %s: %v", r.ID, diags)
	}
	tfDataApiOut.ProviderConfig = tfDataApi.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfDataApiOut)
}

func normalizePostgresCatalogState(ic *importContext, r *resource) error {
	var tfCatalog postgres_catalog_resource.Catalog
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfCatalog); err != nil {
		return err
	}
	var catalog postgres.Catalog
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfCatalog, &catalog)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres catalog state for %s: %v", r.ID, diags)
	}

	if catalog.CatalogId == "" {
		catalog.CatalogId = postgresPathPart(catalog.Name, "catalogs")
	}
	if catalog.Spec == nil && catalog.Status != nil {
		catalog.Spec = &postgres.CatalogCatalogSpec{
			Branch:           catalog.Status.Branch,
			PostgresDatabase: catalog.Status.PostgresDatabase,
		}
	}

	var tfCatalogOut postgres_catalog_resource.Catalog
	diags = converters.GoSdkToTfSdkStruct(ic.Context, catalog, &tfCatalogOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres catalog state back for %s: %v", r.ID, diags)
	}
	tfCatalogOut.ProviderConfig = tfCatalog.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfCatalogOut)
}

func normalizePostgresCdfConfigState(ic *importContext, r *resource) error {
	var tfCdfConfig postgres_cdf_config_resource.CdfConfig
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfCdfConfig); err != nil {
		return err
	}
	var cdfConfig postgres.CdfConfig
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfCdfConfig, &cdfConfig)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres cdf config state for %s: %v", r.ID, diags)
	}

	if cdfConfig.CdfConfigId == "" {
		cdfConfig.CdfConfigId = postgresPathPart(cdfConfig.Name, "cdf-configs")
	}
	parent := postgresPathParentBefore(cdfConfig.Name, "cdf-configs")

	var tfCdfConfigOut postgres_cdf_config_resource.CdfConfig
	diags = converters.GoSdkToTfSdkStruct(ic.Context, cdfConfig, &tfCdfConfigOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres cdf config state back for %s: %v", r.ID, diags)
	}
	if parent != "" {
		tfCdfConfigOut.Parent = types.StringValue(parent)
	}
	tfCdfConfigOut.ProviderConfig = tfCdfConfig.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfCdfConfigOut)
}

func normalizePostgresSyncedTableState(ic *importContext, r *resource) error {
	var tfSyncedTable postgres_synced_table_resource.SyncedTable
	if err := r.DataWrapper.GetTypedStruct(ic.Context, &tfSyncedTable); err != nil {
		return err
	}
	var syncedTable postgres.SyncedTable
	diags := converters.TfSdkToGoSdkStruct(ic.Context, tfSyncedTable, &syncedTable)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres synced table state for %s: %v", r.ID, diags)
	}

	if syncedTable.SyncedTableId == "" {
		syncedTable.SyncedTableId = strings.TrimPrefix(syncedTable.Name, "synced_tables/")
	}

	var tfSyncedTableOut postgres_synced_table_resource.SyncedTable
	diags = converters.GoSdkToTfSdkStruct(ic.Context, syncedTable, &tfSyncedTableOut)
	if diags.HasError() {
		return fmt.Errorf("failed to convert postgres synced table state back for %s: %v", r.ID, diags)
	}
	tfSyncedTableOut.ProviderConfig = tfSyncedTable.ProviderConfig
	return writePostgresTfStruct(ic, r, &tfSyncedTableOut)
}

func shouldOmitPostgresField(ic *importContext, pathString string, fieldSchema FieldSchema, wrapper ResourceDataWrapper, r *resource) bool {
	switch pathString {
	case "project_id", "branch_id", "endpoint_id", "database_id", "role_id", "catalog_id", "cdf_config_id", "synced_table_id":
		_, ok := wrapper.GetOk(pathString)
		return !ok
	}
	if pathString == "spec" || strings.HasPrefix(pathString, "spec.") {
		_, ok := wrapper.GetOk(pathString)
		return !ok
	}
	return DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, r)
}

func listPostgresProjects(ic *importContext) error {
	projects, err := ic.workspaceClient.Postgres.ListProjectsAll(ic.Context, postgres.ListProjectsRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, project := range projects {
		if !ic.MatchesName(project.Name) {
			log.Printf("[INFO] Skipping postgres project %s because it doesn't match %s", project.Name, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_postgres_project",
			ID:       project.Name,
		}, 0, fmt.Sprintf("postgres project '%s'", project.Name))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Postgres Projects", i)
	}
	return nil
}

func listPostgresCatalogs(ic *importContext) error {
	catalogs, err := ic.workspaceClient.Catalogs.ListAll(ic.Context, catalog.ListCatalogsRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, ucCatalog := range catalogs {
		if ucCatalog.CatalogType != catalog.CatalogTypeManagedOnlineCatalog {
			continue
		}
		if !ic.MatchesName(ucCatalog.Name) {
			log.Printf("[INFO] Skipping postgres catalog %s because it doesn't match %s", ucCatalog.Name, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_postgres_catalog",
			ID:       "catalogs/" + ucCatalog.Name,
		}, ucCatalog.UpdatedAt, fmt.Sprintf("postgres catalog '%s'", ucCatalog.Name))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Postgres Catalogs", i)
	}
	return nil
}

func postgresProjectNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	projectId, _ := wrapper.GetOk("project_id")
	if s, ok := projectId.(string); ok {
		return postgresResourceNameFromPath("projects/" + s)
	}
	return ""
}

func importPostgresProject(ic *importContext, r *resource) error {
	projectNameVal, _ := r.DataWrapper.GetOk("name")
	projectName, ok := projectNameVal.(string)
	if !ok || projectName == "" {
		return nil
	}
	if err := normalizePostgresProjectState(ic, r); err != nil {
		return err
	}
	projectId := postgresPathPart(projectName, "projects")
	if projectId == "" {
		projectId = projectName
	}
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/database-projects/%s", projectId),
		"postgres_project_"+r.Name)
	branches, err := ic.workspaceClient.Postgres.ListBranchesAll(ic.Context, postgres.ListBranchesRequest{
		Parent: projectName,
	})
	if err != nil {
		return err
	}
	for _, branch := range branches {
		ic.Emit(&resource{
			Resource: "databricks_postgres_branch",
			ID:       branch.Name,
		})
	}
	return nil
}

func postgresBranchNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresBranch(ic *importContext, r *resource) error {
	branchNameVal, _ := r.DataWrapper.GetOk("name")
	branchName, ok := branchNameVal.(string)
	if !ok || branchName == "" {
		return nil
	}
	if err := normalizePostgresBranchState(ic, r); err != nil {
		return err
	}
	endpoints, err := ic.workspaceClient.Postgres.ListEndpointsAll(ic.Context, postgres.ListEndpointsRequest{
		Parent: branchName,
	})
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints {
		ic.Emit(&resource{
			Resource: "databricks_postgres_endpoint",
			ID:       endpoint.Name,
		})
	}
	databases, err := ic.workspaceClient.Postgres.ListDatabasesAll(ic.Context, postgres.ListDatabasesRequest{
		Parent: branchName,
	})
	if err != nil {
		return err
	}
	for _, database := range databases {
		ic.Emit(&resource{
			Resource: "databricks_postgres_database",
			ID:       database.Name,
		})
	}
	roles, err := ic.workspaceClient.Postgres.ListRolesAll(ic.Context, postgres.ListRolesRequest{
		Parent: branchName,
	})
	if err != nil {
		return err
	}
	for _, role := range roles {
		ic.Emit(&resource{
			Resource: "databricks_postgres_role",
			ID:       role.Name,
		})
	}
	return nil
}

func postgresEndpointNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresEndpoint(ic *importContext, r *resource) error {
	endpointNameVal, _ := r.DataWrapper.GetOk("name")
	endpointName, ok := endpointNameVal.(string)
	if !ok || endpointName == "" {
		return nil
	}
	return normalizePostgresEndpointState(ic, r)
}

func postgresDatabaseNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresDatabase(ic *importContext, r *resource) error {
	databaseNameVal, _ := r.DataWrapper.GetOk("name")
	databaseName, ok := databaseNameVal.(string)
	if !ok || databaseName == "" {
		return nil
	}
	if err := normalizePostgresDatabaseState(ic, r); err != nil {
		return err
	}
	cdfConfigs, err := ic.workspaceClient.Postgres.ListCdfConfigsAll(ic.Context, postgres.ListCdfConfigsRequest{
		Parent: databaseName,
	})
	if err != nil {
		return err
	}
	for _, cdfConfig := range cdfConfigs {
		ic.Emit(&resource{
			Resource: "databricks_postgres_cdf_config",
			ID:       cdfConfig.Name,
		})
	}
	dataApiName := databaseName + "/data-api"
	if _, err := ic.workspaceClient.Postgres.GetDataApi(ic.Context, postgres.GetDataApiRequest{Name: dataApiName}); err == nil {
		ic.Emit(&resource{
			Resource: "databricks_postgres_data_api",
			ID:       dataApiName,
		})
	} else if !apierr.IsMissing(err) {
		return err
	}
	return nil
}

func postgresRoleNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresRole(ic *importContext, r *resource) error {
	return normalizePostgresRoleState(ic, r)
}

func postgresDataApiNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresDataApi(ic *importContext, r *resource) error {
	return normalizePostgresDataApiState(ic, r)
}

func postgresCatalogNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresCatalog(ic *importContext, r *resource) error {
	return normalizePostgresCatalogState(ic, r)
}

func postgresCdfConfigNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresCdfConfig(ic *importContext, r *resource) error {
	return normalizePostgresCdfConfigState(ic, r)
}

func postgresSyncedTableNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	name, _ := wrapper.GetOk("name")
	if s, ok := name.(string); ok {
		return postgresResourceNameFromPath(s)
	}
	return ""
}

func importPostgresSyncedTable(ic *importContext, r *resource) error {
	return normalizePostgresSyncedTableState(ic, r)
}
