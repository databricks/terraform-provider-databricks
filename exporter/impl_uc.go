package exporter

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	tf_uc "github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"golang.org/x/exp/slices"
)

func listUcCatalogs(ic *importContext) error {
	if ic.currentMetastore == nil {
		return fmt.Errorf("there is no UC metastore information")
	}
	it := ic.workspaceClient.Catalogs.List(ic.Context, catalog.ListCatalogsRequest{})
	for it.HasNext(ic.Context) {
		v, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		switch v.CatalogType {
		case "MANAGED_CATALOG", "FOREIGN_CATALOG", "DELTASHARING_CATALOG":
			{
				name := fmt.Sprintf("%s_%s_%s", v.Name, ic.currentMetastore.Name, v.CatalogType)
				ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
					Resource: "databricks_catalog",
					ID:       v.Name,
					Name:     nameNormalizationRegex.ReplaceAllString(name, "_"),
				}, v.Name, v.UpdatedAt, fmt.Sprintf("catalog '%s'", v.Name))
			}
		default:
			// TODO: remove this skipping - we need to emit grants for all catalogs, but
			// we need to convert these catalog types to data sources
			log.Printf("[INFO] Skipping catalog %s of type %s", v.Name, v.CatalogType)
		}
	}
	return nil
}

func importUcCatalog(ic *importContext, r *resource) error {
	var cat tf_uc.CatalogInfo
	s := ic.Resources["databricks_catalog"].Schema
	common.DataToStructPointer(r.Data, s, &cat)

	// TODO: convert `main` catalog into the data source as it's automatically created?
	// Emit: UC Connection, List schemas, Catalog grants, ...
	owner, catalogGrantsResource := ic.emitUCGrantsWithOwner("catalog/"+cat.Name, r)
	dependsOn := []*resource{}
	if owner != "" && owner != ic.meUserName {
		dependsOn = append(dependsOn, catalogGrantsResource)
	}
	if cat.ConnectionName != "" {
		ic.Emit(&resource{
			Resource: "databricks_connection",
			ID:       cat.MetastoreID + "|" + cat.ConnectionName,
		})
	} else if cat.ShareName == "" {
		// TODO: We need to be careful here if we add more catalog types... Really we need to have CatalogType in resource
		if ic.isServiceInListing("uc-schemas") {
			ignoredSchemas := []string{"information_schema"}
			it := ic.workspaceClient.Schemas.List(ic.Context, catalog.ListSchemasRequest{CatalogName: r.ID})
			for it.HasNext(ic.Context) {
				schema, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				if schema.CatalogType != "MANAGED_CATALOG" || slices.Contains(ignoredSchemas, schema.Name) {
					continue
				}
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource:  "databricks_schema",
					ID:        schema.FullName,
					DependsOn: dependsOn,
				}, schema.UpdatedAt, fmt.Sprintf("schema '%s'", schema.FullName))
			}
		}
	}
	if cat.IsolationMode == "ISOLATED" {
		ic.emitWorkspaceBindings("catalog", cat.Name)
	}
	return nil
}

func importUcSchema(ic *importContext, r *resource) error {
	schemaFullName := r.ID
	catalogName := r.Data.Get("catalog_name").(string)
	schemaName := r.Data.Get("name").(string)
	owner, schemaGrantResource := ic.emitUCGrantsWithOwner("schema/"+schemaFullName, r)
	dependsOn := []*resource{}
	if owner != "" && owner != ic.meUserName {
		dependsOn = append(dependsOn, schemaGrantResource)
	}
	// TODO: think if we need to emit upstream dependencies in case if we're going bottom-up
	ic.Emit(&resource{
		Resource: "databricks_catalog",
		ID:       catalogName,
	})
	if ic.isServiceInListing("uc-models") {
		it := ic.workspaceClient.RegisteredModels.List(ic.Context,
			catalog.ListRegisteredModelsRequest{
				CatalogName: catalogName,
				SchemaName:  schemaName,
			})
		for it.HasNext(ic.Context) {
			model, err := it.Next(ic.Context)
			if err != nil {
				return err // TODO: should we continue?
			}
			ic.EmitIfUpdatedAfterMillis(&resource{
				Resource:  "databricks_registered_model",
				ID:        model.FullName,
				DependsOn: dependsOn,
			}, model.UpdatedAt, fmt.Sprintf("registered model '%s'", model.FullName))
		}
	}
	if ic.isServiceInListing("uc-volumes") {
		// list volumes
		it := ic.workspaceClient.Volumes.List(ic.Context,
			catalog.ListVolumesRequest{
				CatalogName: catalogName,
				SchemaName:  schemaName,
			})
		for it.HasNext(ic.Context) {
			volume, err := it.Next(ic.Context)
			if err != nil {
				return err // TODO: should we continue?
			}
			ic.EmitIfUpdatedAfterMillis(&resource{
				Resource:  "databricks_volume",
				ID:        volume.FullName,
				DependsOn: dependsOn,
			}, volume.UpdatedAt, fmt.Sprintf("volume '%s'", volume.FullName))
		}
	}
	isTablesListingEnabled := ic.isServiceInListing("uc-tables")
	isOnlineTablesListingEnabled := ic.isServiceInListing("uc-online-tables")
	isVectorSearchListingEnabled := ic.isServiceInListing("vector-search")
	if isTablesListingEnabled || isOnlineTablesListingEnabled || isVectorSearchListingEnabled {
		it := ic.workspaceClient.Tables.List(ic.Context, catalog.ListTablesRequest{
			CatalogName: catalogName,
			SchemaName:  schemaName,
		})
		for it.HasNext(ic.Context) {
			table, err := it.Next(ic.Context)
			if err != nil {
				return err // TODO: should we continue?
			}
			switch table.TableType {
			case "MANAGED", "EXTERNAL", "VIEW":
				if isTablesListingEnabled {
					ic.EmitIfUpdatedAfterMillis(&resource{
						Resource:  "databricks_sql_table",
						ID:        table.FullName,
						DependsOn: dependsOn,
					}, table.UpdatedAt, fmt.Sprintf("table '%s'", table.FullName))
				}
			case "FOREIGN":
				// TODO: it's better to use SecurableKind if it will be added to the Go SDK
				switch table.DataSourceFormat {
				case "VECTOR_INDEX_FORMAT":
					if isVectorSearchListingEnabled {
						ic.Emit(&resource{
							Resource: "databricks_vector_search_index",
							ID:       table.FullName,
						})
					}
				case "MYSQL_FORMAT":
					if isOnlineTablesListingEnabled {
						ic.EmitIfUpdatedAfterMillis(&resource{
							Resource:  "databricks_online_table",
							ID:        table.FullName,
							DependsOn: dependsOn,
						}, table.UpdatedAt, fmt.Sprintf("table '%s'", table.FullName))
					}
				default:
					log.Printf("[DEBUG] Skipping foreign table %s of format %s", table.FullName, table.DataSourceFormat)
				}
			default:
				log.Printf("[DEBUG] Skipping table %s of type %s", table.FullName, table.TableType)
			}
		}
	}
	return nil
}

func importUcVolume(ic *importContext, r *resource) error {
	volumeFullName := r.ID
	ic.emitUCGrantsWithOwner("volume/"+volumeFullName, r)

	schemaFullName := r.Data.Get("catalog_name").(string) + "." + r.Data.Get("schema_name").(string)
	ic.Emit(&resource{
		Resource: "databricks_schema",
		ID:       schemaFullName,
	})
	return nil
}

// TODO: Should we try to make name unique?
// TODO: do we need to emit principals? Maybe only on account level? See comment for the owner...
func importUcGrants(ic *importContext, r *resource) error {
	if ic.meUserName == "" {
		return nil
	}
	// https://docs.databricks.com/en/data-governance/unity-catalog/manage-privileges/privileges.html#privilege-types-by-securable-object-in-unity-catalog
	var newPrivileges []string
	for k, v := range grantsPrivilegesToAdd {
		if r.Data.Get(k).(string) != "" {
			newPrivileges = append(newPrivileges, v...)
			break
		}
	}
	if len(newPrivileges) == 0 {
		return nil
	}

	owner, found := r.GetExtraData("owner")
	if !found || owner == "" || owner == ic.meUserName {
		// We don't need to change permissions if owner isn't set, or it's the same user
		return nil
	}

	var pList tf_uc.PermissionsList
	s := ic.Resources["databricks_grants"].Schema
	common.DataToStructPointer(r.Data, s, &pList)
	foundExisting := false
	for i, v := range pList.Assignments {
		if v.Principal == ic.meUserName {
			pList.Assignments[i].Privileges = append(pList.Assignments[i].Privileges, newPrivileges...)
			slices.Sort(pList.Assignments[i].Privileges)
			pList.Assignments[i].Privileges = slices.Compact(pList.Assignments[i].Privileges)
			foundExisting = true
			break
		}
	}
	if !foundExisting {
		pList.Assignments = append(pList.Assignments, tf_uc.PrivilegeAssignment{
			Principal:  ic.meUserName,
			Privileges: newPrivileges,
		})
	}
	return common.StructToData(pList, s, r.Data)
}

func importUcStorageCredential(ic *importContext, r *resource) error {
	if r.ID == "__databricks_managed_storage_credential" {
		// it's created by default and can't be imported
		// TODO: add check for "securable_kind":"STORAGE_CREDENTIAL_DB_AWS_IAM" when we get it in the credential
		r.Mode = "data"
		data := tf_uc.ResourceStorageCredential().ToResource().TestResourceData()
		obj := tf_uc.StorageCredentialInfo{Name: r.ID}
		r.Data = ic.generateNewData(data, "databricks_storage_credential", r.ID, obj)
	}
	ic.emitUCGrantsWithOwner("storage_credential/"+r.ID, r)
	if r.Data != nil {
		isolationMode := r.Data.Get("isolation_mode").(string)
		if isolationMode == "ISOLATION_MODE_ISOLATED" {
			ic.emitWorkspaceBindings("storage_credential", r.ID)
		}
	}
	return nil
}

func listUcStorageCredentials(ic *importContext) error {
	it := ic.workspaceClient.StorageCredentials.List(ic.Context, catalog.ListStorageCredentialsRequest{})
	for it.HasNext(ic.Context) {
		v, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_storage_credential",
			ID:       v.Name,
		}, v.Name, v.UpdatedAt, fmt.Sprintf("storage credential %s", v.Name))
	}
	return nil
}

func importUcCredential(ic *importContext, r *resource) error {
	ic.emitUCGrantsWithOwner("credential/"+r.ID, r)
	if r.Data != nil {
		isolationMode := r.Data.Get("isolation_mode").(string)
		if isolationMode == "ISOLATION_MODE_ISOLATED" {
			purpose := r.Data.Get("purpose").(string)
			if purpose == "SERVICE" {
				ic.emitWorkspaceBindings("credential", r.ID)
			} else if purpose == "STORAGE" {
				ic.emitWorkspaceBindings("storage_credential", r.ID)
			}
		}
	}
	return nil
}

func listUcCredentials(ic *importContext) error {
	it := ic.workspaceClient.Credentials.ListCredentials(ic.Context, catalog.ListCredentialsRequest{})
	for it.HasNext(ic.Context) {
		v, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if v.Purpose == catalog.CredentialPurposeStorage {
			continue // we're handling storage credentials separately
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_credential",
			ID:       v.Name,
		}, v.Name, v.UpdatedAt, fmt.Sprintf("credential %s", v.Name))
	}
	return nil
}

func (ic *importContext) emitUCGrantsWithOwner(id string, parentResource *resource) (string, *resource) {
	gr := &resource{
		Resource: "databricks_grants",
		ID:       id,
	}
	var owner string
	if parentResource.Data != nil {
		ignoreFunc := ic.Importables[parentResource.Resource].Ignore
		if ignoreFunc != nil && ignoreFunc(ic, parentResource) {
			return "", nil
		}
		ownerRaw, ok := parentResource.Data.GetOk("owner")
		if ok {
			gr.AddExtraData("owner", ownerRaw)
			owner = ownerRaw.(string)
			emitUserSpOrGroup(ic, owner)
		}
	}
	ic.Emit(gr)
	return owner, gr
}

var (
	emailDomainRegex = regexp.MustCompile(`^.*@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+\.?$`)
)

func emitUserSpOrGroup(ic *importContext, userOrSOrGroupPName string) {
	if common.StringIsUUID(userOrSOrGroupPName) {
		ic.Emit(&resource{
			Resource:  "databricks_service_principal",
			Attribute: "application_id",
			Value:     userOrSOrGroupPName,
		})
	} else if emailDomainRegex.MatchString(userOrSOrGroupPName) {
		ic.Emit(&resource{
			Resource:  "databricks_user",
			Attribute: "user_name",
			Value:     userOrSOrGroupPName,
		})
	} else {
		ic.Emit(&resource{
			Resource:  "databricks_group",
			Attribute: "display_name",
			Value:     userOrSOrGroupPName,
		})
	}
}

func shouldOmitForUnityCatalog(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if pathString == "owner" {
		return d.Get(pathString).(string) == ""
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d)
}

func shouldOmitWithIsolationMode(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if pathString == "isolation_mode" {
		return d.Get(pathString).(string) != "ISOLATION_MODE_ISOLATED"
	}
	return shouldOmitForUnityCatalog(ic, pathString, as, d)
}

func createIsMatchingCatalogAndSchema(catalog_name_attr, schema_name_attr string) func(ic *importContext, res *resource,
	ra *resourceApproximation, origPath string) bool {
	return func(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
		// catalog and schema names for the source resource. We need to copy the original catalog_name_attr
		// to a new variable because we're going to modify it
		new_catalog_name_attr := catalog_name_attr
		if strings.HasSuffix(origPath, "."+schema_name_attr) {
			new_catalog_name_attr = strings.TrimSuffix(origPath, schema_name_attr) + catalog_name_attr
		}
		res_catalog_name := res.Data.Get(new_catalog_name_attr).(string)
		res_schema_name := res.Data.Get(origPath).(string)
		// In some cases catalog or schema name could be empty, like, in non-UC DLT pipelines, so we need to skip it
		if res_catalog_name == "" || res_schema_name == "" {
			return false
		}
		// catalog and schema names for target resource approximation
		ra_catalog_name, cat_found := ra.Get("catalog_name")
		ra_schema_name, schema_found := ra.Get("name")
		if !cat_found || !schema_found {
			log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v). Resource: %s, catalog='%s', schema='%s'",
				ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, res.Resource, res_catalog_name, res_schema_name)
			return false
		}
		result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name
		return result

	}
}

func createIsMatchingCatalogAndSchemaAndTable(catalog_name_attr, schema_name_attr, table_name_attr string) func(ic *importContext, res *resource,
	ra *resourceApproximation, origPath string) bool {
	return func(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
		// catalog and schema names for the source resource. We need to copy the original catalog_name_attr
		// to a new variable because we're going to modify it
		new_catalog_name_attr := catalog_name_attr
		new_schema_name_attr := schema_name_attr
		if strings.HasSuffix(origPath, "."+table_name_attr) {
			prefix := strings.TrimSuffix(origPath, table_name_attr)
			new_catalog_name_attr = prefix + catalog_name_attr
			new_schema_name_attr = prefix + schema_name_attr
		}
		res_catalog_name := res.Data.Get(new_catalog_name_attr).(string)
		res_schema_name := res.Data.Get(new_schema_name_attr).(string)
		res_table_name := res.Data.Get(origPath).(string)
		// In some cases catalog or schema name could be empty, like, in non-UC DLT pipelines, so we need to skip it
		if res_catalog_name == "" || res_schema_name == "" || res_table_name == "" {
			return false
		}
		// catalog and schema names for target resource approximation
		ra_catalog_name, cat_found := ra.Get("catalog_name")
		ra_schema_name, schema_found := ra.Get("schema_name")
		ra_table_name, table_found := ra.Get("name")
		if !cat_found || !schema_found || !table_found {
			log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v) table='%v' (found? %v). Resource: %s, catalog='%s', schema='%s', table='%s'",
				ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, ra_table_name,
				table_found, res.Resource, res_catalog_name, res_schema_name, res_table_name)
			return false
		}
		result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name && ra_table_name.(string) == res_table_name
		return result

	}
}

func (ic *importContext) emitWorkspaceBindings(securableType, securableName string) {
	bindings, err := ic.workspaceClient.WorkspaceBindings.GetBindingsAll(ic.Context, catalog.GetBindingsRequest{
		SecurableName: securableName,
		SecurableType: string(securableType),
	})
	if err != nil {
		log.Printf("[ERROR] listing %s bindings for %s: %s", securableType, securableName, err.Error())
		return
	}
	for _, binding := range bindings {
		id := fmt.Sprintf("%d|%s|%s", binding.WorkspaceId, securableType, securableName)
		// We were creating Data instance explicitly because of the bug in the databricks_catalog_workspace_binding
		// implementation. Technically, after the fix is merged we can remove this, but we're keeping it as-is now
		// to decrease a number of API calls.
		d := ic.Resources["databricks_workspace_binding"].Data(
			&terraform.InstanceState{
				ID: id,
				Attributes: map[string]string{
					"workspace_id":   fmt.Sprintf("%d", binding.WorkspaceId),
					"securable_type": securableType,
					"securable_name": securableName,
					"binding_type":   binding.BindingType.String(),
				},
			})
		ic.Emit(&resource{
			Resource: "databricks_workspace_binding",
			ID:       id,
			Name:     fmt.Sprintf("%s_%s_ws_%d", securableType, securableName, binding.WorkspaceId),
			Data:     d,
		})
	}
}

func isMatchingSecurableTypeAndName(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	res_securable_type := res.Data.Get("securable_type").(string)
	res_securable_name := res.Data.Get("securable_name").(string)
	ra_name, _ := ra.Get("name")
	return ra.Type == ("databricks_"+res_securable_type) && ra_name.(string) == res_securable_name
}

func isMatchingAllowListArtifact(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	objPath := strings.Replace(origPath, ".artifact", ".match_type", 1)
	matchType, ok := res.Data.GetOk(objPath)
	artifactType := res.Data.Get("artifact_type").(string)
	return ok && matchType.(string) == "PREFIX_MATCH" && (artifactType == "LIBRARY_JAR" || artifactType == "INIT_SCRIPT")
}

func importUcExternalLocation(ic *importContext, r *resource) error {
	if r.ID == dbManagedExternalLocation {
		// it's created by default and can't be imported
		// TODO: add check for "securable_kind":"EXTERNAL_LOCATION_DB_STORAGE" when we get it in the credential
		r.Mode = "data"
		data := tf_uc.ResourceExternalLocation().ToResource().TestResourceData()
		obj := tf_uc.ExternalLocationInfo{ExternalLocationInfo: catalog.ExternalLocationInfo{Name: r.ID}}
		r.Data = ic.generateNewData(data, "databricks_external_location", r.ID, obj)
	}
	ic.emitUCGrantsWithOwner("external_location/"+r.ID, r)
	credentialName := r.Data.Get("credential_name").(string)
	ic.Emit(&resource{
		Resource: "databricks_storage_credential",
		ID:       credentialName,
	})
	if r.Data != nil {
		isolationMode := r.Data.Get("isolation_mode").(string)
		if isolationMode == "ISOLATION_MODE_ISOLATED" {
			ic.emitWorkspaceBindings("external_location", r.ID)
		}
	}
	// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "storage_credential/" + credentialName})
	return nil
}

func listUcExternalLocations(ic *importContext) error {
	it := ic.workspaceClient.ExternalLocations.List(ic.Context, catalog.ListExternalLocationsRequest{})
	for it.HasNext(ic.Context) {
		v, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if v.Name != "metastore_default_location" {
			ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
				Resource: "databricks_external_location",
				ID:       v.Name,
			}, v.Name, v.UpdatedAt, fmt.Sprintf("external location %s", v.Name))
		}
	}
	return nil
}

func listUcConnections(ic *importContext) error {
	it := ic.workspaceClient.Connections.List(ic.Context, catalog.ListConnectionsRequest{})
	for it.HasNext(ic.Context) {
		conn, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_connection",
			ID:       conn.MetastoreId + "|" + conn.Name,
		}, conn.Name, conn.UpdatedAt, fmt.Sprintf("connection '%s'", conn.Name))
	}
	return nil
}

func listUcMetastores(ic *importContext) error {
	it := ic.accountClient.Metastores.List(ic.Context)
	for it.HasNext(ic.Context) {
		mstore, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_metastore",
			ID:       mstore.MetastoreId,
		}, mstore.Name, mstore.UpdatedAt, fmt.Sprintf("metastore '%s'", mstore.Name))
	}
	return nil
}

func importUcMetastores(ic *importContext, r *resource) error {
	ic.emitUCGrantsWithOwner("metastore/"+r.ID, r)
	if ic.accountLevel {
		// emit metastore assignments
		assignments, err := ic.accountClient.MetastoreAssignments.ListByMetastoreId(ic.Context, r.ID)
		if err == nil {
			for _, workspaceID := range assignments.WorkspaceIds {
				ic.Emit(&resource{
					Resource: "databricks_metastore_assignment",
					ID:       fmt.Sprintf("%d|%s", workspaceID, r.ID),
				})
			}
		} else {
			log.Printf("[ERROR] listing metastore assignments: %s", err.Error())
		}
		// TODO: emit storage credentials associated with specific metastores, but we'll need to solve
		// a problem of importing a resource... This will require to changing ID from name to metastore ID + name.
	}
	return nil
}

func importUcVolumeFile(ic *importContext, r *resource) error {
	parts := strings.Split(r.ID, "/")
	// Converting /Volumes/<catalog>/<schema>/<table>/<file> to <catalog>.<schema>.<table>
	if len(parts) > 5 {
		volumeId := strings.Join(parts[2:5], ".")
		ic.Emit(&resource{
			Resource: "databricks_volume",
			ID:       volumeId,
		})
		// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "volume/" + volumeId})
	}

	// download & store file
	resp, err := ic.workspaceClient.Files.DownloadByFilePath(ic.Context, r.ID)
	if err != nil {
		return err
	}
	// write file
	fileName := ic.prefix + fileNameNormalizationRegex.ReplaceAllString(strings.TrimPrefix(r.ID, "/Volumes/"), "_")
	local, relativeName, err := ic.createFileIn("uc_files", fileName)
	if err != nil {
		return err
	}
	defer local.Close()
	defer resp.Contents.Close()
	_, err = io.Copy(local, resp.Contents)
	if err != nil {
		return err
	}
	r.Data.Set("source", relativeName)
	r.Data.Set("path", r.ID)

	return nil
}

func listSystemSchemas(ic *importContext) error {
	if ic.currentMetastore == nil {
		return fmt.Errorf("there is no UC metastore information")
	}
	currentMetastore := ic.currentMetastore.MetastoreId
	systemSchemas, err := ic.workspaceClient.SystemSchemas.ListAll(ic.Context,
		catalog.ListSystemSchemasRequest{MetastoreId: currentMetastore})
	if err != nil {
		return err
	}
	for _, v := range systemSchemas {
		if v.Schema == "information_schema" || v.Schema == "__internal_logging" {
			continue
		}
		if v.State == "ENABLE_COMPLETED" || v.State == "ENABLE_INITIALIZED" {
			id := fmt.Sprintf("%s|%s", currentMetastore, v.Schema)
			data := ic.Resources["databricks_system_schema"].Data(
				&terraform.InstanceState{
					ID: id,
					Attributes: map[string]string{
						"metastore_id": currentMetastore,
						"schema":       v.Schema,
					},
				})
			ic.Emit(&resource{
				Resource: "databricks_system_schema",
				ID:       id,
				Data:     data,
				Name:     nameNormalizationRegex.ReplaceAllString(id, "_"),
			})
		} else {
			log.Printf("[DEBUG] Skipping system schema %s with state %s", v.Schema, v.State)
		}
	}
	return nil
}

func listArtifactAllowLists(ic *importContext) error {
	if ic.currentMetastore == nil {
		return fmt.Errorf("there is no UC metastore information")
	}
	artifactTypes := []string{"INIT_SCRIPT", "LIBRARY_JAR", "LIBRARY_MAVEN"}
	for _, v := range artifactTypes {
		id := fmt.Sprintf("%s|%s", ic.currentMetastore.MetastoreId, v)
		name := fmt.Sprintf("%s_%s_%s", v, ic.currentMetastore.Name, ic.currentMetastore.MetastoreId[:8])
		ic.Emit(&resource{
			Resource: "databricks_artifact_allowlist",
			ID:       id,
			Name:     nameNormalizationRegex.ReplaceAllString(name, "_"),
		})
	}
	return nil
}

func importSqlTable(ic *importContext, r *resource) error {
	tableFullName := r.ID
	ic.emitUCGrantsWithOwner("table/"+tableFullName, r)
	schemaFullName := r.Data.Get("catalog_name").(string) + "." + r.Data.Get("schema_name").(string)
	ic.Emit(&resource{
		Resource: "databricks_schema",
		ID:       schemaFullName,
	})
	return nil
}
