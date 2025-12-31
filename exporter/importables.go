package exporter

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sql"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"

	"github.com/databricks/terraform-provider-databricks/common"
	alert_v2_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/alert_v2"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	tf_dlt "github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/repos"
	tf_sql "github.com/databricks/terraform-provider-databricks/sql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"golang.org/x/exp/maps"
)

var (
	globalWorkspaceConfName          = "global_workspace_conf"
	nameNormalizationRegex           = regexp.MustCompile(`\W+`)
	fileNameNormalizationRegex       = regexp.MustCompile(`[^-_\w/.@]`)
	jobClustersRegex                 = regexp.MustCompile(`^((job_cluster|task)\.\d+\.new_cluster\.\d+\.)`)
	dltClusterRegex                  = regexp.MustCompile(`^(cluster\.\d+\.)`)
	secretPathRegex                  = regexp.MustCompile(`^\{\{secrets\/([^\/]+)\/([^}]+)\}\}$`)
	secretScopePathRegex             = regexp.MustCompile(`^\{\{secrets\/([^\/]+)\/[^}]+\}\}$`)
	sqlParentRegexp                  = regexp.MustCompile(`^folders/(\d+)$`)
	requirementsFileRegexp           = regexp.MustCompile(`-r\s+(/.*)$`)
	dltDefaultStorageRegex           = regexp.MustCompile(`^dbfs:/pipelines/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	ignoreIdeFolderRegex             = regexp.MustCompile(`^/Users/[^/]+/\.ide/.*$`)
	servedEntityFieldExtractionRegex = regexp.MustCompile(`^config\.[0-9]+\.served_entities\.([0-9]+)\.(.*)$`)
	uc3LevelIdRegex                  = regexp.MustCompile(`^([^.]+\.[^.]+\.[^.]+)$`)
	globIncludeDirectoryRegex        = regexp.MustCompile(`^(/.+)/\*\*$`)
	fileExtensionLanguageMapping     = map[string]string{
		"SCALA":  ".scala",
		"PYTHON": ".py",
		"SQL":    ".sql",
		"R":      ".r",
	}
	fileExtensionFormatMapping = map[string]string{
		"HTML":       ".html",
		"JUPYTER":    ".ipynb",
		"DBC":        ".dbc",
		"R_MARKDOWN": ".Rmd",
	}
	grantsPrivilegesToAdd = map[string][]string{
		// TODO: think how to handle these if the TF user isn't owner of the metastore...
		// "metastore": {`CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_PROVIDER`,
		// 	`CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_STORAGE_CREDENTIAL`, `SET_SHARE_PERMISSION`},
		"catalog":            {`CREATE_SCHEMA`},
		"schema":             {`CREATE_FUNCTION`, `CREATE_TABLE`, `CREATE_MODEL`, `CREATE_VOLUME`},
		"volume":             {`WRITE_VOLUME`},
		"external_location":  {`CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_MANAGED_STORAGE`},
		"storage_credential": {`CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`},
		"foreign_connection": {`CREATE_FOREIGN_CATALOG`},
	}
	ParentDirectoryExtraKey   = "parent_directory"
	dbManagedExternalLocation = "__databricks_managed_storage_location"
)

var resourcesMap map[string]importable = map[string]importable{
	"databricks_dbfs_file": {
		WorkspaceLevel: true,
		Service:        "storage",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			fileNameMd5 := fmt.Sprintf("%x", md5.Sum([]byte(d.Id())))
			s := strings.Split(d.Id(), "/")
			name := "_" + fileNameMd5 + "_" + s[len(s)-1]
			return name
		},
		Import:          importDbfsFile,
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
		},
	},
	"databricks_instance_pool": {
		WorkspaceLevel: true,
		Service:        "pools",
		Name:           instancePoolName,
		List:           listInstancePools,
		Import:         importInstancePool,
		Ignore:         generateIgnoreObjectWithEmptyAttributeValue("databricks_instance_pool", "instance_pool_name"),
	},
	"databricks_instance_profile": {
		Service:        "access",
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			arn := d.Get("instance_profile_arn").(string)
			splits := strings.Split(arn, "/")
			return splits[len(splits)-1]
		},
	},
	"databricks_group_role": {
		Service:        "access",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Depends: []reference{
			{Path: "group_id", Resource: "databricks_group"},
			{Path: "role", Resource: "databricks_instance_profile", Match: "instance_profile_arn"},
		},
	},
	"databricks_user_role": {
		Service:        "access",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Depends: []reference{
			{Path: "user_id", Resource: "databricks_user"},
			{Path: "role", Resource: "databricks_instance_profile", Match: "instance_profile_arn"},
		},
	},
	"databricks_service_principal_role": {
		Service:        "access",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Depends: []reference{
			{Path: "service_principal_id", Resource: "databricks_service_principal"},
			{Path: "role", Resource: "databricks_instance_profile", Match: "instance_profile_arn"},
		},
	},
	"databricks_cluster": {
		WorkspaceLevel: true,
		Service:        "compute",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("cluster_name").(string)
			if name == "" {
				parts := strings.Split(d.Id(), "-")
				if len(parts) > 2 {
					return parts[2]
				}
				return d.Id()
			}
			return fmt.Sprintf("%s_%s", name, d.Id())
		},
		Depends: []reference{
			{Path: "aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "init_scripts.volumes.destination", Resource: "databricks_file"},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.jar", Resource: "databricks_file"},
			{Path: "library.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_file"},
			{Path: "library.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.whl", Resource: "databricks_workspace_file"},
			{Path: "library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "policy_id", Resource: "databricks_cluster_policy"},
			{Path: "docker_image.basic_auth.password", Resource: "databricks_secret", Match: "config_reference"},
			{Path: "docker_image.basic_auth.username", Resource: "databricks_secret", Match: "config_reference"},
			{Path: "docker_image.basic_auth.password", Resource: "databricks_secret_scope",
				MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "docker_image.basic_auth.username", Resource: "databricks_secret_scope",
				MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "spark_conf", Resource: "databricks_secret", Match: "config_reference"},
			{Path: "spark_env_vars", Resource: "databricks_secret", Match: "config_reference"},
			{Path: "spark_conf", Resource: "databricks_secret_scope",
				MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "spark_env_vars", Resource: "databricks_secret_scope",
				MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "single_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "single_user_name", Resource: "databricks_group", Match: "display_name"},
			{Path: "single_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "library.jar", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.whl", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.egg", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
		List:            listClusters,
		Import:          importCluster,
		ShouldOmitField: makeShouldOmitFieldForCluster(nil),
	},
	"databricks_job": {
		ApiVersion:     common.API_2_1,
		WorkspaceLevel: true,
		Service:        "jobs",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				name = "job"
			}
			return nameNormalizationRegex.ReplaceAllString(
				fmt.Sprintf("%s_%s", name, d.Id()), "_")
		},
		Depends:         createJobDependencies(),
		Import:          importJob,
		List:            listJobs,
		ShouldOmitField: shouldOmitFieldInJob,
		Ignore:          shouldIgnoreJob,
	},
	"databricks_cluster_policy": {
		WorkspaceLevel: true,
		Service:        "policies",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		List:   listClusterPolicies,
		Import: importClusterPolicy,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "definition" {
				return d.Get("policy_family_id").(string) != ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "libraries.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "libraries.jar", Resource: "databricks_file"},
			{Path: "libraries.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "libraries.whl", Resource: "databricks_file"},
			{Path: "libraries.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "libraries.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.whl", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "libraries.egg", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "libraries.jar", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
		// TODO: implement a custom Body that will write with special formatting, where
		// JSON is written line by line so that we're able to do the references
	},
	"databricks_group": {
		Service:        "groups",
		WorkspaceLevel: true,
		AccountLevel:   true,
		Name:           makeNamePlusIdFunc("display_name"),
		List:           listGroups,
		Search:         searchGroup,
		Import:         importGroup,
	},
	"databricks_group_member": {
		Service:        "groups",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Depends: []reference{
			{Path: "group_id", Resource: "databricks_group"},
			{Path: "member_id", Resource: "databricks_user"},
			{Path: "member_id", Resource: "databricks_group"},
			{Path: "member_id", Resource: "databricks_service_principal"},
		},
	},
	"databricks_user": {
		Service:        "users",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			s := d.Get("user_name").(string)
			// if CLI argument includeUserDomains is set then it includes domain portion as well
			if ic.includeUserDomains {
				return nameNormalizationRegex.ReplaceAllString(s, "_") + "_" + d.Id()
			}
			return nameNormalizationRegex.ReplaceAllString(strings.Split(s, "@")[0], "_") + "_" + d.Id()
		},
		List:   listUsers,
		Search: searchUser,
		Import: importUser,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if r.Mode == "data" {
				return pathString != "user_name"
			}
			if pathString == "display_name" {
				userName := d.Get("user_name").(string)
				displayName := d.Get("display_name").(string)
				return displayName == "" || userName == displayName
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
	},
	"databricks_service_principal": {
		Service:        "users",
		AccountLevel:   true,
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("display_name").(string)
			if name == "" {
				name = d.Get("application_id").(string)
				if len(name) > 8 {
					name = name[0:8]
				}
			}
			return name + "_" + d.Id()
		},
		List:   listServicePrincipals,
		Search: searchServicePrincipal,
		Import: importServicePrincipal,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if r.Mode == "data" {
				return pathString != "application_id"
			}
			if pathString == "display_name" {
				if ic.Client.IsAzure() {
					if ic.targetCloud != "" {
						return false
					}
					applicationID := d.Get("application_id").(string)
					displayName := d.Get("display_name").(string)
					externalID := d.Get("external_id").(string)
					return applicationID == displayName && externalID != ""
				}
				return false
			}
			// application_id should be provided only on Azure and only for Azure-managed SPs that
			// have external_id set
			if pathString == "application_id" {
				return (ic.Client.IsAzure() && ic.targetCloud != "azure" && ic.targetCloud != "") ||
					!ic.Client.IsAzure() || (d.Get("external_id").(string) == "")

			}
			if pathString == "external_id" {
				return ic.targetCloud != "" || d.Get("external_id").(string) == ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
	},
	"databricks_permissions": {
		Service:        "access",
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			s := strings.Split(d.Id(), "/")
			return s[len(s)-1]
		},
		List: func(ic *importContext) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       "/authorization/tokens",
					Name:     "tokens_usage",
				})
			}
			return nil
		},
		Depends: []reference{
			{Path: "job_id", Resource: "databricks_job"},
			{Path: "pipeline_id", Resource: "databricks_pipeline"},
			{Path: "cluster_id", Resource: "databricks_cluster"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster_policy_id", Resource: "databricks_cluster_policy"},
			{Path: "sql_query_id", Resource: "databricks_query"},
			{Path: "sql_alert_id", Resource: "databricks_alert"},
			{Path: "sql_dashboard_id", Resource: "databricks_sql_dashboard"},
			{Path: "sql_endpoint_id", Resource: "databricks_sql_endpoint"},
			{Path: "dashboard_id", Resource: "databricks_dashboard"},
			{Path: "registered_model_id", Resource: "databricks_mlflow_model"},
			{Path: "experiment_id", Resource: "databricks_mlflow_experiment"},
			{Path: "repo_id", Resource: "databricks_repo"},
			{Path: "vector_search_endpoint_id", Resource: "databricks_vector_search_endpoint", Match: "endpoint_id"},
			{Path: "serving_endpoint_id", Resource: "databricks_model_serving", Match: "serving_endpoint_id"},
			{Path: "database_instance_name", Resource: "databricks_database_instance", Match: "name"},
			{Path: "app_name", Resource: "databricks_app", Match: "name"},
			// TODO: can we fill _path component for it, and then match on user/SP home instead?
			{Path: "directory_id", Resource: "databricks_directory", Match: "object_id"},
			{Path: "notebook_id", Resource: "databricks_notebook", Match: "object_id"},
			{Path: "workspace_file_id", Resource: "databricks_workspace_file", Match: "object_id"},
			{Path: "access_control.group_name", Resource: "databricks_group", Match: "display_name"},
			{Path: "access_control.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "access_control.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
		Ignore: func(ic *importContext, r *resource) bool {
			return (r.Data.Get("access_control.#").(int) == 0)
		},
		Import: func(ic *importContext, r *resource) error {
			var permissions entity.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			common.DataToStructPointer(r.Data, s, &permissions)
			for _, ac := range permissions.AccessControlList {
				ic.Emit(&resource{
					Resource:  "databricks_user",
					Attribute: "user_name",
					Value:     ac.UserName,
				})
				ic.Emit(&resource{
					Resource:  "databricks_group",
					Attribute: "display_name",
					Value:     ac.GroupName,
				})
				ic.Emit(&resource{
					Resource:  "databricks_service_principal",
					Attribute: "application_id",
					Value:     ac.ServicePrincipalName,
				})
			}
			return nil
		},
	},
	"databricks_permission_assignment": {
		Service:        "idfed",
		WorkspaceLevel: true,
		List:           listWorkspacePermissionAssignments,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			switch pathString {
			case "principal_id":
				return true
			case "user_name", "service_principal_name", "group_name":
				return d.Get(pathString).(string) == ""
			default:
				return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
			}
		},
		// Note: We don't need dependencies here as we assign permissions by name, not by ID
	},
	"databricks_secret_scope": {
		Service:        "secrets",
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			return name + "_" + generateUniqueID(name)
		},
		List: func(ic *importContext) error {
			scopes := ic.workspaceClient.Secrets.ListScopes(ic.Context)
			for scopes.HasNext(ic.Context) {
				scope, err := scopes.Next(ic.Context)
				if err != nil {
					return err
				}
				if !ic.MatchesName(scope.Name) {
					log.Printf("[INFO] Secret scope %s doesn't match %s filter", scope.Name, ic.match)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_secret_scope",
					ID:       scope.Name,
				})
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			backendType, _ := r.Data.GetOk("backend_type")
			if backendType != "AZURE_KEYVAULT" || ic.targetCloud != "" {
				secrets := ic.workspaceClient.Secrets.ListSecrets(ic.Context, sdk_workspace.ListSecretsRequest{
					Scope: r.ID,
				})
				for secrets.HasNext(ic.Context) {
					secret, err := secrets.Next(ic.Context)
					if err != nil {
						return err
					}
					ic.Emit(&resource{
						Resource: "databricks_secret",
						ID:       fmt.Sprintf("%s|||%s", r.ID, secret.Key),
					})
				}
			}
			if backendType == "AZURE_KEYVAULT" || ic.targetCloud != "" {
				r.Data.Set("backend_type ", "DATABRICKS")
				r.Data.Set("keyvault_metadata", nil)
			}
			acls, err := ic.workspaceClient.Secrets.ListAclsByScope(ic.Context, r.ID)
			if err != nil {
				return err
			}
			for _, acl := range acls.Items {
				ic.Emit(&resource{
					Resource: "databricks_secret_acl",
					ID:       fmt.Sprintf("%s|||%s", r.ID, acl.Principal),
				})
			}
			return nil
		},
		Ignore: func(ic *importContext, r *resource) bool {
			return r.Data.Get("name").(string) == ""
		},
	},
	"databricks_secret": {
		WorkspaceLevel: true,
		Service:        "secrets",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := fmt.Sprintf("%s_%s", d.Get("scope"), d.Get("key"))
			return name + "_" + generateUniqueID(name)
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.exportSecrets {
				resp, err := ic.workspaceClient.Secrets.GetSecret(ic.Context, sdk_workspace.GetSecretRequest{
					Scope: r.Data.Get("scope").(string),
					Key:   r.Data.Get("key").(string),
				})
				if err != nil {
					return err
				}
				secret, err := base64.StdEncoding.DecodeString(resp.Value)
				if err != nil {
					return err
				}
				varName := ic.generateVariableName("string_value", ic.ResourceName(r))
				ic.addTfVar(varName, string(secret))
			}
			return nil
		},
		Depends: []reference{
			{Path: "string_value", Variable: true},
			{Path: "scope", Resource: "databricks_secret_scope"},
		},
	},
	"databricks_secret_acl": {
		WorkspaceLevel: true,
		Service:        "secrets",
		Depends: []reference{
			{Path: "scope", Resource: "databricks_secret_scope"},
			{Path: "principal", Resource: "databricks_group", Match: "display_name"},
			{Path: "principal", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "principal", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_mount": {
		WorkspaceLevel: true,
		Service:        "mounts",
		Body:           generateMountBody,
		List:           listMounts,
		Depends: []reference{
			{Path: "instance_profile", Resource: "databricks_instance_profile"},
			{Path: "cluster_id", Resource: "databricks_cluster"},
		},
	},
	"databricks_global_init_script": {
		WorkspaceLevel: true,
		Service:        "wsconf",
		Name:           makeNameOrIdFunc("name"),
		List: func(ic *importContext) error {
			globalInitScripts, err := ic.workspaceClient.GlobalInitScripts.ListAll(ic.Context)
			if err != nil {
				return err
			}
			for offset, gis := range globalInitScripts {
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_global_init_script",
					ID:       gis.ScriptId,
				}, int64(gis.UpdatedAt), fmt.Sprintf("global init script '%s'", gis.Name))
				log.Printf("[INFO] Scanned %d of %d global init scripts", offset+1, len(globalInitScripts))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			gis, err := ic.workspaceClient.GlobalInitScripts.GetByScriptId(ic.Context, r.ID)
			if err != nil {
				return err
			}
			content, err := base64.StdEncoding.DecodeString(gis.Script)
			if err != nil {
				return err
			}
			fileName, err := ic.saveContentIn("global_init_scripts", fmt.Sprintf("%s.sh", ic.ResourceName(r)), content)
			if err != nil {
				return err
			}
			return r.Data.Set("source", fileName)
		},
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
		},
	},
	"databricks_repo": {
		WorkspaceLevel: true,
		Service:        "repos",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("path").(string)
			if name == "" {
				return "repo_" + d.Id()
			}
			return nameNormalizationRegex.ReplaceAllString(name[7:], "_") + "_" + d.Id()
		},
		Search: searchRepoByPath,
		List:   listRepos,
		Import: importRepo,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			switch pathString {
			case "path":
				return false
			case "branch":
				return d.Get("branch").(string) == ""
			case "tag":
				return d.Get("tag").(string) == ""
			case "git_provider":
				url := d.Get("url").(string)
				provider := repos.GetGitProviderFromUrl(url)
				return provider != "" // omit git_provider only for well-known URLs
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			shouldIgnore := r.Data.Get("url").(string) == ""
			if shouldIgnore {
				path := r.Data.Get("path").(string)
				log.Printf("[WARN] ignoring databricks_repo without Git provider. Path: %s", path)
				ic.addIgnoredResource(fmt.Sprintf("databricks_repo. path=%s", r.Data.Get("path").(string)))
			}
			return shouldIgnore
		},
		Depends: []reference{
			{Path: "path", Resource: "databricks_user", Match: "repos",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_service_principal", Match: "repos",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_user", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_service_principal", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
	},
	"databricks_workspace_conf": {
		WorkspaceLevel: true,
		Service:        "wsconf",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return globalWorkspaceConfName
		},
		List: func(ic *importContext) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_workspace_conf",
					ID:       globalWorkspaceConfName,
				})
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			keyNames := maps.Keys(ic.workspaceConfKeys)
			sort.Strings(keyNames)
			conf, err := ic.workspaceClient.WorkspaceConf.GetStatus(ic.Context, settings.GetStatusRequest{
				Keys: strings.Join(keyNames, ","),
			})
			if err != nil {
				log.Printf("[WARN] Error getting workspace conf: %s", err)
				return err
			}
			loaded := map[string]any{}
			for k, v := range *conf {
				if v == "" {
					continue
				}
				loaded[k] = v
			}
			r.Data.Set("custom_config", loaded)
			return nil
		},
	},
	"databricks_ip_access_list": {
		WorkspaceLevel: true,
		Service:        "access",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("list_type").(string) + "_" + d.Get("label").(string)
		},
		List: func(ic *importContext) error {
			ipLists, err := ic.workspaceClient.IpAccessLists.ListAll(ic.Context)

			if err != nil {
				return err
			}
			for offset, ipList := range ipLists {
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_ip_access_list",
					ID:       ipList.ListId,
				}, ipList.UpdatedAt, fmt.Sprintf("IP access list '%s'", ipList.Label))
				log.Printf("[INFO] Scanned %d of %d IP Access Lists", offset+1, len(ipLists))
			}
			if len(ipLists) > 0 {
				ic.Emit(&resource{
					Resource: "databricks_workspace_conf",
					ID:       globalWorkspaceConfName,
					Data: ic.Resources["databricks_workspace_conf"].Data(
						&terraform.InstanceState{
							ID:         globalWorkspaceConfName,
							Attributes: map[string]string{},
						}),
				})
			}
			return nil
		},
	},
	"databricks_notebook": {
		WorkspaceLevel: true,
		Service:        "notebooks",
		Name:           workspaceObjectResouceName,
		Import:         ImportNotebook,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			switch pathString {
			case "language":
				return d.Get("language") == ""
			case "format":
				return d.Get("format") == "SOURCE"
			}
			return shouldOmitMd5Field(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_directory", MatchType: MatchLongestPrefix,
				SearchValueTransformFunc: appendEndingSlashToDirName, ExtraLookupKey: ParentDirectoryExtraKey},
			{Path: "path", Resource: "databricks_user", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_service_principal", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
	},
	"databricks_workspace_file": {
		WorkspaceLevel:  true,
		Service:         "wsfiles",
		Name:            workspaceObjectResouceName,
		Import:          importWorkspaceFile,
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_directory", MatchType: MatchLongestPrefix,
				SearchValueTransformFunc: appendEndingSlashToDirName, ExtraLookupKey: ParentDirectoryExtraKey},
			{Path: "path", Resource: "databricks_user", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_service_principal", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
	},
	"databricks_query": {
		WorkspaceLevel: true,
		Service:        "queries",
		Name:           makeNamePlusIdFunc("display_name"),
		List:           listQueries,
		Import:         importQuery,
		// TODO: exclude owner if it's the current user?
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_query", "display_name"),
		Depends: []reference{
			{Path: "warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "parameter.query_backed_value.query_id", Resource: "databricks_query", Match: "id"},
			{Path: "catalog", Resource: "databricks_catalog"},
			{Path: "schema", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema"),
				SkipDirectLookup:     true},
			// TODO: add match like for workspace files?
			{Path: "parent_path", Resource: "databricks_user", Match: "home"},
			{Path: "parent_path", Resource: "databricks_service_principal", Match: "home"},
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
			{Path: "owner_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			// TODO: add support for Repos?
		},
	},
	"databricks_sql_endpoint": {
		WorkspaceLevel: true,
		Service:        "sql-endpoints",
		Name:           makeNameOrIdFunc("name"),
		List:           listSqlEndpoints,
		Import:         importSqlEndpoint,
		Ignore:         generateIgnoreObjectWithEmptyAttributeValue("databricks_sql_endpoint", "name"),
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			switch pathString {
			case "enable_serverless_compute":
				return false
			case "tags":
				return d.Get("tags.0.custom_tags.#").(int) == 0
			case "channel.0.name":
				channelName := d.Get(pathString).(string)
				return channelName == "" || channelName == "CHANNEL_NAME_CURRENT"
			case "channel":
				channelName := d.Get(pathString + ".0.name").(string)
				return channelName == "" || channelName == "CHANNEL_NAME_CURRENT"
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		ShouldGenerateField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			// We need to generate it even if it's false...
			return pathString == "enable_serverless_compute"
		},
	},
	"databricks_sql_global_config": {
		WorkspaceLevel: true,
		Service:        "wsconf",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return "sql_global_config"
		},
		List: func(ic *importContext) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_sql_global_config",
					ID:       tf_sql.GlobalSqlConfigResourceID,
				})
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			arn := r.Data.Get("instance_profile_arn").(string)
			if arn != "" {
				ic.Emit(&resource{
					Resource: "databricks_instance_profile",
					ID:       arn,
				})
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "enable_serverless_compute" {
				return false
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "instance_profile_arn", Resource: "databricks_instance_profile"},
		},
	},
	"databricks_sql_dashboard": {
		WorkspaceLevel: true,
		Service:        "sql-dashboards",
		Name:           makeNamePlusIdFunc("name"),
		List:           listRedashDashboards,
		Import:         importRedashDashboard,
		Depends: []reference{
			{Path: "parent", Resource: "databricks_directory", Match: "object_id", MatchType: MatchRegexp,
				Regexp: sqlParentRegexp},
		},
	},
	"databricks_sql_widget": {
		WorkspaceLevel: true,
		Service:        "sql-dashboards",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Id()
		},
		Depends: []reference{
			{Path: "visualization_id", Resource: "databricks_sql_visualization", Match: "visualization_id"},
			{Path: "dashboard_id", Resource: "databricks_sql_dashboard"},
		},
	},
	"databricks_sql_visualization": {
		WorkspaceLevel: true,
		Service:        "sql-dashboards",
		Name:           makeNamePlusIdFunc("name"),
		Depends: []reference{
			{Path: "query_id", Resource: "databricks_query"},
		},
	},
	"databricks_alert": {
		WorkspaceLevel: true,
		Service:        "alerts",
		Name:           makeNamePlusIdFunc("display_name"),
		List:           listAlerts,
		Import: func(ic *importContext, r *resource) error {
			var alert sql.Alert
			s := ic.Resources["databricks_alert"].Schema
			common.DataToStructPointer(r.Data, s, &alert)
			if alert.QueryId != "" {
				ic.Emit(&resource{Resource: "databricks_query", ID: alert.QueryId})
			}
			ic.emitDirectoryOrRepo(alert.ParentPath)
			ic.emitUserOrServicePrincipal(alert.OwnerUserName)
			// TODO: r.AddExtraData(ParentDirectoryExtraKey, directoryPath) ?
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/alerts/%s", r.ID),
				"alert_"+ic.Importables["databricks_alert"].Name(ic, r.Data))
			return nil
		},
		// TODO: exclude owner if it's the current user?
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_alert", "display_name"),
		Depends: []reference{
			{Path: "query_id", Resource: "databricks_query"},
			// TODO: add match like for workspace files?
			{Path: "parent_path", Resource: "databricks_user", Match: "home"},
			{Path: "parent_path", Resource: "databricks_service_principal", Match: "home"},
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
			{Path: "owner_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_alert_v2": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "alerts",
		Name:            makeNamePlusIdFunc("display_name"),
		List:            listAlertsV2,
		// Body function removed - using generic HCL generation for Plugin Framework resources
		Import: func(ic *importContext, r *resource) error {
			// Convert Plugin Framework state to Go SDK struct
			var alert sql.AlertV2
			if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper, alert_v2_resource.AlertV2{}, &alert); err != nil {
				return err
			}

			// Emit dependencies - now using plain Go strings!
			if alert.WarehouseId != "" {
				ic.Emit(&resource{Resource: "databricks_sql_endpoint", ID: alert.WarehouseId})
			}

			if alert.ParentPath != "" {
				ic.emitDirectoryOrRepo(alert.ParentPath)
			}

			if alert.OwnerUserName != "" {
				ic.emitUserOrServicePrincipal(alert.OwnerUserName)
			}

			// Handle evaluation.notification.subscriptions
			if alert.Evaluation.Notification != nil {
				for _, sub := range alert.Evaluation.Notification.Subscriptions {
					if sub.DestinationId != "" {
						ic.Emit(&resource{Resource: "databricks_notification_destination", ID: sub.DestinationId})
					}
					// user_email is only for users (email addresses), not service principals (UUIDs)
					// emitUserOrServicePrincipal will automatically handle this correctly
					if sub.UserEmail != "" {
						ic.emitUserOrServicePrincipal(sub.UserEmail)
					}
				}
			}

			// For Plugin Framework resources, we can't use r.Data directly, use the wrapper ID
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/alerts/%s", r.ID),
				"alert_v2_"+r.Name)
			return nil
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_alert_v2", "display_name"),
		Depends: []reference{
			{Path: "warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "parent_path", Resource: "databricks_user", Match: "home"},
			{Path: "parent_path", Resource: "databricks_service_principal", Match: "home"},
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
			{Path: "owner_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "evaluation.notification.subscriptions.destination_id", Resource: "databricks_notification_destination"},
			{Path: "evaluation.notification.subscriptions.user_email", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_apps_settings_custom_template": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "apps",
		Name:            func(ic *importContext, d *schema.ResourceData) string { return d.Id() },
		List:            listAppsSettingsCustomTemplates,
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_apps_settings_custom_template", "name"),
	},

	"databricks_custom_app_integration": {
		AccountLevel: true,
		Service:      "oauth",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				return "custom_app_" + d.Id()
			}
			return name + "_" + d.Id()[:8]
		},
		List:   listCustomAppIntegrations,
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_custom_app_integration", "name"),
	},
	"databricks_account_federation_policy": {
		AccountLevel:    true,
		PluginFramework: true,
		Service:         "oauth",
		List:            listAccountFederationPolicies,
	},
	"databricks_service_principal_federation_policy": {
		AccountLevel:    true,
		PluginFramework: true,
		Service:         "oauth",
		List:            listServicePrincipalFederationPolicies,
		Depends: []reference{
			{Path: "service_principal_id", Resource: "databricks_service_principal"},
		},
	},
	"databricks_app": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "apps",
		Name:            func(ic *importContext, d *schema.ResourceData) string { return d.Id() },
		List:            listApps,
		Import:          importApp,
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_app", "name"),
		Depends: []reference{
			{Path: "resources.sql_warehouse.id", Resource: "databricks_sql_endpoint"},
			{Path: "resources.serving_endpoint.name", Resource: "databricks_model_serving"},
			{Path: "resources.job.id", Resource: "databricks_job"},
			{Path: "resources.secret.scope", Resource: "databricks_secret_scope"},
			{Path: "resources.secret.key", Resource: "databricks_secret", Match: "key",
				IsValidApproximation: createIsMatchingScopeAndKey("scope", "key")},
			{Path: "resources.uc_securable.securable_full_name", Resource: "databricks_volume"},
			{Path: "resources.database.instance_name", Resource: "databricks_database_instance", Match: "name"},
			{Path: "budget_policy_id", Resource: "databricks_budget_policy", Match: "policy_id"},
		},
	},
	"databricks_pipeline": {
		WorkspaceLevel: true,
		Service:        "dlt",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				return d.Id()
			}
			return name + "_" + d.Id()
		},
		List:   listPipelines,
		Import: importPipeline,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if res := dltClusterRegex.FindStringSubmatch(pathString); res != nil { // analyze DLT clusters
				return makeShouldOmitFieldForCluster(dltClusterRegex)(ic, pathString, as, d, r)
			}
			switch pathString {
			case "storage":
				return dltDefaultStorageRegex.FindStringSubmatch(d.Get("storage").(string)) != nil
			case "edition":
				return d.Get("edition").(string) == ""
			case "creator_user_name":
				return true
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			var pipeline tf_dlt.Pipeline
			s := ic.Resources["databricks_pipeline"].Schema
			common.DataToStructPointer(r.Data, s, &pipeline)
			if pipeline.Deployment != nil && pipeline.Deployment.Kind == "BUNDLE" {
				log.Printf("[WARN] Ignoring DLT Pipeline with ID %s as deployed with DABs", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_pipeline. id=%s", r.ID))
				return true
			}
			numLibraries := len(pipeline.Libraries)
			if numLibraries == 0 {
				log.Printf("[WARN] Ignoring DLT Pipeline with ID %s due to the lack of libraries", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_pipeline. id=%s", r.ID))
			}
			return numLibraries == 0
		},
		Depends: []reference{
			{Path: "catalog", Resource: "databricks_catalog"},
			{Path: "target", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "target"),
				SkipDirectLookup:     true},
			{Path: "schema", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema"),
				SkipDirectLookup:     true},
			{Path: "cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "cluster.init_scripts.volumes.destination", Resource: "databricks_file"},
			{Path: "cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "configuration", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "configuration", Resource: "databricks_file"},
			{Path: "configuration", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.notebook.path", Resource: "databricks_notebook"},
			{Path: "library.notebook.path", Resource: "databricks_notebook", Match: "workspace_path"},
			{Path: "library.file.path", Resource: "databricks_workspace_file"},
			{Path: "library.file.path", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "root_path", Resource: "databricks_directory"},
			{Path: "root_path", Resource: "databricks_directory", Match: "workspace_path"},
			{Path: "environment.dependencies", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "environment.dependencies", Resource: "databricks_file"},
			{Path: "environment.dependencies", Resource: "databricks_workspace_file", Match: "workspace_path",
				MatchType: MatchRegexp, Regexp: requirementsFileRegexp},
			{Path: "environment.dependencies", Resource: "databricks_file", MatchType: MatchRegexp,
				Regexp: requirementsFileRegexp},
			{Path: "notification.email_recipients", Resource: "databricks_user",
				Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "event_log.catalog", Resource: "databricks_catalog"},
			{Path: "event_log.schema", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema")},
			{Path: "event_log.name", Resource: "databricks_sql_table", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchemaAndTable("catalog", "schema", "name")},
			{Path: "run_as.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "run_as.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "configuration", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.notebook.path", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.notebook.path", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.file.path", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.file.path", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_directory", MatchType: MatchRegexp,
				Regexp: globIncludeDirectoryRegex},
			{Path: "library.glob.include", Resource: "databricks_directory", Match: "workspace_path",
				MatchType: MatchRegexp, Regexp: globIncludeDirectoryRegex},
			{Path: "library.glob.include", Resource: "databricks_notebook", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_workspace_file", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_notebook", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_workspace_file", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.glob.include", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "root_path", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
	},
	"databricks_directory": {
		WorkspaceLevel: true,
		Service:        "directories",
		Name:           workspaceObjectResouceName,
		Search: func(ic *importContext, r *resource) error {
			objId, err := strconv.ParseInt(r.Value, 10, 64)
			if err != nil {
				return err
			}
			directoryList := ic.getAllDirectories()
			for _, directory := range directoryList {
				if directory.ObjectID == objId {
					r.ID = directory.Path
					return nil
				}
			}
			return fmt.Errorf("can't find directory with object_id: %s", r.Value)
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			// Existing permissions API doesn't allow to set permissions for
			if r.ID != "/Shared" {
				ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/directories/%d", r.Data.Get("object_id").(int)),
					"directory_"+ic.Importables["databricks_directory"].Name(ic, r.Data))
			}

			if r.ID == "/Shared" || r.ID == "/Users" || ic.IsUserOrServicePrincipalDirectory(r.ID, "/Users", true) {
				r.Mode = "data"
			}
			return nil
		},
		Depends: []reference{
			{Path: "path", Resource: "databricks_user", Match: "home"},
			{Path: "path", Resource: "databricks_service_principal", Match: "home"},
			// TODO: it should try to find longest reference to another directory object that it not itself...
			{Path: "path", Resource: "databricks_user", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "path", Resource: "databricks_service_principal", Match: "home",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
	},
	"databricks_model_serving": {
		WorkspaceLevel: true,
		Service:        "model-serving",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			nameMd5 := fmt.Sprintf("%x", md5.Sum([]byte(d.Id())))
			return strings.ToLower(d.Id()) + "_" + nameMd5[:8]
		},
		List: func(ic *importContext) error {
			it := ic.workspaceClient.ServingEndpoints.List(ic.Context)
			i := 0
			for it.HasNext(ic.Context) {
				endpoint, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				if !ic.MatchesName(endpoint.Name) {
					log.Printf("[INFO] Skipping serving endpoint %s because it doesn't match %s", endpoint.Name, ic.match)
					continue
				}
				if endpoint.Config != nil && endpoint.Config.ServedEntities != nil && len(endpoint.Config.ServedEntities) > 0 {
					if endpoint.Config.ServedEntities[0].FoundationModel != nil {
						log.Printf("[INFO] skipping endpoint %s that is foundation model", endpoint.Name)
						continue
					}
				}
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_model_serving",
					ID:       endpoint.Name,
				}, endpoint.LastUpdatedTimestamp, fmt.Sprintf("serving endpoint '%s'", endpoint.Name))
				i++
				if i%50 == 0 {
					log.Printf("[INFO] Scanned %d Serving Endpoints", i)
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/serving-endpoints/%s", r.Data.Get("serving_endpoint_id").(string)),
				"serving_endpoint_"+ic.Importables["databricks_model_serving"].Name(ic, r.Data))
			s := ic.Resources["databricks_model_serving"].Schema
			var mse serving.CreateServingEndpoint
			common.DataToStructPointer(r.Data, s, &mse)
			if mse.Config != nil {
				for _, se := range mse.Config.ServedEntities {
					if se.EntityName != "" {
						if se.EntityVersion != "" { // we have an UC model or model from model registry
							if uc3LevelIdRegex.MatchString(se.EntityName) {
								ic.Emit(&resource{
									Resource: "databricks_registered_model",
									ID:       se.EntityName,
								})
							}
							// TODO: add else branch to emit databricks_model when we have support for it
						}
						// TODO: add else branch to emit UC function when we add support for them...
					}
					if se.InstanceProfileArn != "" {
						ic.Emit(&resource{
							Resource: "databricks_instance_profile",
							ID:       se.InstanceProfileArn,
						})
					}
					ic.emitSecretsFromSecretsPathMap(se.EnvironmentVars)
					if se.ExternalModel != nil {
						if se.ExternalModel.DatabricksModelServingConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.DatabricksModelServingConfig.DatabricksApiToken)
						}
						if se.ExternalModel.Ai21labsConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.Ai21labsConfig.Ai21labsApiKey)
						}
						if se.ExternalModel.AnthropicConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.AnthropicConfig.AnthropicApiKey)
						}
						if se.ExternalModel.AmazonBedrockConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.AmazonBedrockConfig.AwsAccessKeyId)
							ic.emitSecretsFromSecretPathString(se.ExternalModel.AmazonBedrockConfig.AwsSecretAccessKey)
						}
						if se.ExternalModel.CohereConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.CohereConfig.CohereApiKey)
						}
						if se.ExternalModel.OpenaiConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.OpenaiConfig.OpenaiApiKey)
							ic.emitSecretsFromSecretPathString(se.ExternalModel.OpenaiConfig.MicrosoftEntraClientSecret)
						}
						if se.ExternalModel.PalmConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.PalmConfig.PalmApiKey)
						}
						if se.ExternalModel.GoogleCloudVertexAiConfig != nil {
							ic.emitSecretsFromSecretPathString(se.ExternalModel.GoogleCloudVertexAiConfig.PrivateKey)
						}
						if se.ExternalModel.CustomProviderConfig != nil {
							if se.ExternalModel.CustomProviderConfig.ApiKeyAuth != nil {
								ic.emitSecretsFromSecretPathString(se.ExternalModel.CustomProviderConfig.ApiKeyAuth.Value)
							}
							if se.ExternalModel.CustomProviderConfig.BearerTokenAuth != nil {
								ic.emitSecretsFromSecretPathString(se.ExternalModel.CustomProviderConfig.BearerTokenAuth.Token)
							}
						}
					}
				}
			}
			if mse.Config != nil && mse.Config.AutoCaptureConfig != nil && mse.Config.AutoCaptureConfig.CatalogName != "" &&
				mse.Config.AutoCaptureConfig.SchemaName != "" {
				ic.Emit(&resource{
					Resource: "databricks_schema",
					ID:       mse.Config.AutoCaptureConfig.CatalogName + "." + mse.Config.AutoCaptureConfig.SchemaName,
				})
			}
			// Auto-capture for AI Gateway
			if mse.AiGateway != nil && mse.AiGateway.InferenceTableConfig != nil &&
				mse.AiGateway.InferenceTableConfig.CatalogName != "" &&
				mse.AiGateway.InferenceTableConfig.SchemaName != "" {
				ic.Emit(&resource{
					Resource: "databricks_schema",
					ID:       mse.AiGateway.InferenceTableConfig.CatalogName + "." + mse.AiGateway.InferenceTableConfig.SchemaName,
				})
			}
			if mse.AiGateway != nil && mse.AiGateway.RateLimits != nil {
				for _, rl := range mse.AiGateway.RateLimits {
					// principal could be a group, user, or service principal
					var isGroupFound bool
					if !common.StringIsUUID(rl.Principal) {
						err := ic.cacheGroups()
						if err == nil {
							for _, g := range ic.allGroups {
								if g.DisplayName == rl.Principal {
									ic.Emit(&resource{
										Resource:  "databricks_group",
										Attribute: "display_name",
										Value:     rl.Principal,
									})
									isGroupFound = true
									break
								}
							}
						}
					}
					if !isGroupFound {
						ic.emitUserOrServicePrincipal(rl.Principal)
					}
				}
			}
			if mse.EmailNotifications != nil {
				ic.emitListOfUsers(mse.EmailNotifications.OnUpdateFailure)
				ic.emitListOfUsers(mse.EmailNotifications.OnUpdateSuccess)
			}
			return nil
		},
		Ignore: func(ic *importContext, r *resource) bool {
			// We need to ignore endpoints that are foundation models
			endpoint, err := ic.workspaceClient.ServingEndpoints.GetByName(ic.Context, r.ID)
			if err != nil {
				log.Printf("[WARN] Can't get endpoint by name %s: %s", r.ID, err)
				return false
			}
			res := (endpoint.Config != nil && endpoint.Config.ServedEntities != nil &&
				len(endpoint.Config.ServedEntities) > 0 && endpoint.Config.ServedEntities[0].FoundationModel != nil)
			log.Printf("[DEBUG] Ignore serving endpoint %s? %t", r.ID, res)
			return res
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "config" {
				return d.Get("config.#").(int) == 0
			}
			if pathString == "config.0.traffic_config" || pathString == "config.0.auto_capture_config.0.enabled" ||
				(pathString == "config.0.auto_capture_config.0.table_name_prefix" && d.Get(pathString).(string) != "") {
				return false
			}
			if strings.HasPrefix(pathString, "config.0.traffic_config.0.routes") && strings.HasSuffix(pathString, ".served_model_name") {
				return true
			}
			if res := servedEntityFieldExtractionRegex.FindStringSubmatch(pathString); res != nil {
				field := res[2]
				log.Printf("[DEBUG] ShouldOmitField: extracted field from %s: '%s'", pathString, field)
				switch field {
				case "scale_to_zero_enabled", "name":
					return false
				case "workload_size", "workload_type":
					return d.Get(pathString).(string) == ""
				}
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		ShouldGenerateField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			// We need to generate some fields even if they have zero value...
			if strings.HasSuffix(pathString, ".scale_to_zero_enabled") {
				extModelBlockCoordinate := strings.Replace(pathString, ".scale_to_zero_enabled", ".external_model", 1)
				return d.Get(extModelBlockCoordinate+".#").(int) == 0
			}
			return pathString == "config.0.auto_capture_config.0.enabled" || pathString == "ai_gateway.0.inference_table_config.0.enabled"
		},
		Depends: []reference{
			{Path: "config.served_entities.entity_name", Resource: "databricks_registered_model"},
			{Path: "config.served_entities.instance_profile_arn", Resource: "databricks_instance_profile",
				Match: "instance_profile_arn"},
			{Path: "config.auto_capture_config.catalog_name", Resource: "databricks_catalog"},
			{Path: "config.served_entities.external_model.databricks_model_serving_config.databricks_api_token",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.ai21labs_config.ai21labs_api_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.anthropic_config.anthropic_api_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.amazon_bedrock_config.aws_access_key_id",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.amazon_bedrock_config.aws_secret_access_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.cohere_config.cohere_api_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.openai_config.openai_api_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.openai_config.microsoft_entra_client_secret",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.palm_config.palm_api_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.google_cloud_vertex_ai_config.private_key",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.custom_provider_config.api_key_auth.value",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.custom_provider_config.bearer_token_auth.token",
				Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.external_model.databricks_model_serving_config.databricks_api_token",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.ai21labs_config.ai21labs_api_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.anthropic_config.anthropic_api_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.amazon_bedrock_config.aws_access_key_id",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.amazon_bedrock_config.aws_secret_access_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.cohere_config.cohere_api_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.openai_config.openai_api_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.openai_config.microsoft_entra_client_secret",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.palm_config.palm_api_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.google_cloud_vertex_ai_config.private_key",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.custom_provider_config.api_key_auth.value",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.external_model.custom_provider_config.bearer_token_auth.token",
				Resource: "databricks_secret_scope", MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.served_entities.environment_vars", Resource: "databricks_secret", Match: "config_reference"},
			{Path: "config.served_entities.environment_vars", Resource: "databricks_secret_scope",
				MatchType: MatchRegexp, Regexp: secretScopePathRegex},
			{Path: "config.auto_capture_config.schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "ai_gateway.inference_table_config.catalog_name", Resource: "databricks_catalog"},
			{Path: "ai_gateway.inference_table_config.schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "ai_gateway.rate_limits.principal", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "ai_gateway.rate_limits.principal", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "ai_gateway.rate_limits.principal", Resource: "databricks_group", Match: "display_name"},
			{Path: "email_notifications.on_update_failure", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_update_success", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_database_instance": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "lakebase",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Id()
		},
		List:                   listDatabaseInstances,
		Import:                 importDatabaseInstance,
		ShouldOmitFieldUnified: shouldOmitWithEffectiveFields,
		Ignore:                 generateIgnoreObjectWithEmptyAttributeValue("databricks_database_instance", "name"),
	},
	"databricks_mlflow_webhook": {
		WorkspaceLevel: true,
		Service:        "mlflow-webhooks",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return "webhook_" + d.Id()
		},
		List: func(ic *importContext) error {
			webhooks, err := ic.workspaceClient.ModelRegistry.ListWebhooksAll(ic.Context, ml.ListWebhooksRequest{})
			if err != nil {
				return err
			}
			for offset, webhook := range webhooks {
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_mlflow_webhook",
					ID:       webhook.Id,
				}, webhook.LastUpdatedTimestamp, fmt.Sprintf("webhook '%s'", webhook.Id))
				if webhook.JobSpec != nil && webhook.JobSpec.JobId != "" {
					ic.Emit(&resource{
						Resource: "databricks_job",
						ID:       webhook.JobSpec.JobId,
					})
				}
				if offset%50 == 0 {
					log.Printf("[INFO] Scanned %d of %d MLflow webhooks", offset+1, len(webhooks))
				}
			}
			return nil
		},
		Depends: []reference{
			{Path: "job_spec.job_id", Resource: "databricks_job"},
			{Path: "job_spec.access_token", Variable: true},
			// We can enable it, but we don't know if authorization is set or not because API doesn't return it
			// {Path: "http_url_spec.authorization", Variable: true},
		},
	},
	"databricks_access_control_rule_set": {
		AccountLevel: true,
		Service:      "access",
		List: func(ic *importContext) error {
			accountId := ic.Client.Config.AccountID
			// emit default ruleset
			ic.Emit(&resource{
				Resource: "databricks_access_control_rule_set",
				ID:       fmt.Sprintf("accounts/%s/ruleSets/default", accountId),
			})
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var rule iam.RuleSetResponse
			s := ic.Resources["databricks_access_control_rule_set"].Schema
			common.DataToStructPointer(r.Data, s, &rule)

			for _, grant := range rule.GrantRules {
				for _, principal := range grant.Principals {
					parts := strings.Split(principal, "/")
					if len(parts) != 2 {
						log.Printf("[WARN] Incorrect principal: '%s'", principal)
						continue
					}
					switch parts[0] {
					case "users":
						ic.Emit(&resource{
							Resource:  "databricks_user",
							Attribute: "user_name",
							Value:     parts[1],
						})
					case "servicePrincipals":
						ic.Emit(&resource{
							Resource:  "databricks_service_principal",
							Attribute: "application_id",
							Value:     parts[1],
						})
					case "groups":
						ic.Emit(&resource{
							Resource:  "databricks_group",
							Attribute: "display_name",
							Value:     parts[1],
						})
					default:
						log.Printf("[WARN] Unknown principal type: '%s'", parts[0])
					}
				}
			}

			return nil
		},
		Depends: []reference{
			{Path: "grant_rules.principals", Resource: "databricks_user", Match: "acl_principal_id"},
			{Path: "grant_rules.principals", Resource: "databricks_group", Match: "acl_principal_id"},
			{Path: "grant_rules.principals", Resource: "databricks_service_principal", Match: "acl_principal_id"},
			{Path: "name", Resource: "databricks_service_principal", Match: "application_id", MatchType: MatchRegexp,
				Regexp: regexp.MustCompile("^accounts/[^/]+/servicePrincipals/([^/]+)/ruleSets/default$")},
			{Path: "name", Resource: "databricks_group", MatchType: MatchRegexp,
				Regexp: regexp.MustCompile("^accounts/[^/]+/groups/([^/]+)/ruleSets/default$")},
			{Path: "name", Resource: "databricks_budget_policy", Match: "policy_id", MatchType: MatchRegexp,
				Regexp: regexp.MustCompile(`^accounts/[^/]+/budgetPolicies/([^/]+)/ruleSets/default$`)},
		},
		Ignore: func(ic *importContext, r *resource) bool {
			// We're ignoring ACLs without grant rules because we don't know about that at time of emitting from groups/service principals
			var rule iam.RuleSetResponse
			s := ic.Resources["databricks_access_control_rule_set"].Schema
			common.DataToStructPointer(r.Data, s, &rule)
			shouldIgnore := len(rule.GrantRules) == 0
			if shouldIgnore {
				log.Printf("[WARN] ignoring databricks_access_control_rule_set without grant rules. ID: %s", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_access_control_rule_set. ID=%s", r.ID))
			}
			return shouldIgnore
		},
	},
	"databricks_system_schema": {
		WorkspaceLevel: true,
		Service:        "uc-system-schemas",
		List:           listSystemSchemas,
	},
	"databricks_artifact_allowlist": {
		WorkspaceLevel: true,
		Service:        "uc-artifact-allowlist",
		List:           listArtifactAllowLists,
		Ignore: func(ic *importContext, r *resource) bool {
			numBlocks := r.Data.Get("artifact_matcher.#").(int)
			if numBlocks == 0 {
				log.Printf("[WARN] Ignoring artifcat allowlist with ID %s", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_artifact_allowlist. id=%s", r.ID))
			}
			return numBlocks == 0
		},
		Depends: []reference{
			{Path: "artifact_matcher.artifact", Resource: "databricks_volume", Match: "volume_path",
				IsValidApproximation: isMatchingAllowListArtifact},
			{Path: "artifact_matcher.artifact", Resource: "databricks_external_location", Match: "url",
				IsValidApproximation: isMatchingAllowListArtifact},
			{Path: "artifact_matcher.artifact", Resource: "databricks_volume", Match: "volume_path",
				MatchType: MatchLongestPrefix, IsValidApproximation: isMatchingAllowListArtifact},
			{Path: "artifact_matcher.artifact", Resource: "databricks_external_location", Match: "url",
				MatchType: MatchLongestPrefix, IsValidApproximation: isMatchingAllowListArtifact},
		},
	},
	"databricks_catalog": {
		WorkspaceLevel: true,
		Service:        "uc-catalogs",
		List:           listUcCatalogs,
		Import:         importUcCatalog,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			res := (r.Data != nil && (r.Data.Get("name").(string) == "" || r.Data.Get("name").(string) == "system"))
			if res {
				ic.addIgnoredResource(fmt.Sprintf("databricks_catalog. id=%s", r.ID))
			}
			return res
		},
		Depends: []reference{
			{Path: "connection_name", Resource: "databricks_connection", Match: "name"},
			{Path: "storage_root", Resource: "databricks_external_location", Match: "url", MatchType: MatchLongestPrefix},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_schema": {
		WorkspaceLevel:  true,
		Service:         "uc-schemas",
		Import:          importUcSchema,
		ShouldOmitField: shouldOmitForUnityCatalog,
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_schema", "name"),
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "storage_root", Resource: "databricks_external_location", Match: "url", MatchType: MatchLongestPrefix},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_volume": {
		WorkspaceLevel: true,
		Service:        "uc-volumes",
		Import:         importUcVolume,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "storage_location" {
				return d.Get("volume_type").(string) == "MANAGED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_volume", "name"),
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "storage_location", Resource: "databricks_external_location",
				Match: "url", MatchType: MatchLongestPrefix},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_sql_table": {
		WorkspaceLevel: true,
		Service:        "uc-tables",
		Import:         importSqlTable,
		Ignore:         generateIgnoreObjectWithEmptyAttributeValue("databricks_sql_table", "name"),
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			log.Printf("[INFO] ShouldOmitField: %s", pathString)
			switch pathString {
			case "storage_location":
				return d.Get("table_type").(string) == "MANAGED"
			case "enable_predictive_optimization":
				epo := d.Get(pathString).(string)
				return epo == "" || epo == "INHERIT"
			case "column", "partitions":
				return d.Get("table_type").(string) == "VIEW"
			}
			if strings.HasPrefix(pathString, "column.") {
				if d.Get("table_type").(string) == "VIEW" {
					return true
				}
				if strings.HasSuffix(pathString, ".nullable") {
					return d.Get(pathString).(bool)
				}
				if strings.HasSuffix(pathString, ".type") {
					return false
				}
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "storage_location", Resource: "databricks_external_location",
				Match: "url", MatchType: MatchLongestPrefix},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_data_quality_monitor": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "dq",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			// ID format is "object_type,object_id" (e.g., "table,abc-123-def")
			id := d.Id()
			parts := strings.Split(id, ",")
			if len(parts) == 2 {
				objectType := parts[0]
				objectId := parts[1]
				// Create name like "table_monitor_abc12345"
				if len(objectId) > 8 {
					return fmt.Sprintf("%s_monitor_%s", objectType, objectId[:8])
				}
				return fmt.Sprintf("%s_monitor_%s", objectType, objectId)
			}
			return "monitor_" + generateUniqueID(id)
		},
		Import: importDataQualityMonitor,
		List:   listDataQualityMonitors,
		// Monitors should be also emitted as dependencies from tables/schemas (TODO: we need to add it)
		Depends: []reference{
			// object_id matches either table_id or schema_id depending on object_type
			{Path: "object_id", Resource: "databricks_sql_table", Match: "table_id"},
			{Path: "object_id", Resource: "databricks_schema", Match: "schema_id"},
			// Full names match resource.id directly
			{Path: "data_profiling_config.monitored_table_name", Resource: "databricks_sql_table"},
			{Path: "data_profiling_config.baseline_table_name", Resource: "databricks_sql_table"},
			{Path: "data_profiling_config.warehouse_id", Resource: "databricks_sql_endpoint"},
			// Email addresses match user_name field
			{Path: "data_profiling_config.notification_settings.on_failure.email_addresses",
				Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_quality_monitor_v2": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "dq",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			// ID format is "object_type,object_id" (e.g., "schema,abc-123-def")
			id := d.Id()
			parts := strings.Split(id, ",")
			if len(parts) == 2 {
				objectType := parts[0]
				objectId := parts[1]
				// Create name like "schema_monitor_v2_abc12345"
				if len(objectId) > 8 {
					return fmt.Sprintf("%s_monitor_v2_%s", objectType, objectId[:8])
				}
				return fmt.Sprintf("%s_monitor_v2_%s", objectType, objectId)
			}
			return "monitor_v2_" + generateUniqueID(id)
		},
		Import: importQualityMonitorV2,
		List:   listQualityMonitorsV2,
		// Monitors should be also emitted as dependencies from tables/schemas (TODO: we need to add it)
		Depends: []reference{
			// object_id matches schema_id for schema-level monitors
			{Path: "object_id", Resource: "databricks_schema", Match: "schema_id"},
		},
	},

	"databricks_grants": {
		WorkspaceLevel: true,
		Service:        "uc-grants",
		Import:         importUcGrants,
		Ignore: func(ic *importContext, r *resource) bool {
			return (r.Data.Get("grant.#").(int) == 0)
		},
		Depends: []reference{
			{Path: "catalog", Resource: "databricks_catalog"},
			{Path: "schema", Resource: "databricks_schema"},
			{Path: "volume", Resource: "databricks_volume"},
			{Path: "share", Resource: "databricks_share"},
			{Path: "table", Resource: "databricks_sql_table"},
			{Path: "foreign_connection", Resource: "databricks_connection", Match: "name"},
			{Path: "metastore", Resource: "databricks_metastore"},
			{Path: "model", Resource: "databricks_registered_model"},
			{Path: "external_location", Resource: "databricks_external_location", Match: "name"},
			{Path: "storage_credential", Resource: "databricks_storage_credential"},
			{Path: "credential", Resource: "databricks_credential"},
			// TODO: add similar matchers for users/groups/SPs on account level...
			{Path: "grant.principal", Resource: "databricks_recipient", IsValidApproximation: isMatchingShareRecipient},
			//	{Path: "", Resource: ""},
			//	{Path: "", Resource: ""},
		},
	},
	"databricks_storage_credential": {
		WorkspaceLevel:  true,
		Service:         "uc-storage-credentials",
		Import:          importUcStorageCredential,
		List:            listUcStorageCredentials,
		ShouldOmitField: shouldOmitWithIsolationMode,
		Depends: []reference{
			{Path: "azure_service_principal.client_secret", Variable: true},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_credential": {
		WorkspaceLevel:  true,
		Service:         "uc-credentials",
		Import:          importUcCredential,
		List:            listUcCredentials,
		ShouldOmitField: shouldOmitWithIsolationMode,
		Depends: []reference{
			{Path: "azure_service_principal.client_secret", Variable: true},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_external_location": {
		WorkspaceLevel: true,
		Service:        "uc-external-locations",
		Import:         importUcExternalLocation,
		List:           listUcExternalLocations,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if (pathString == "url" || pathString == "credential_name") && d.Get("name").(string) == dbManagedExternalLocation {
				return true
			}
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATION_MODE_ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		// This external location is automatically created when metastore is created with the `storage_root`
		Ignore: func(ic *importContext, r *resource) bool {
			return r.ID == "metastore_default_location"
		},
		Depends: []reference{
			{Path: "credential_name", Resource: "databricks_storage_credential", Match: "name"},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_connection": {
		WorkspaceLevel: true,
		Service:        "uc-connections",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			connectionName := d.Get("name").(string)
			connectionType := d.Get("connection_type").(string)
			if connectionName == "" || connectionType == "" {
				return d.Id()
			}
			return connectionType + "_" + connectionName
		},
		List: listUcConnections,
		// TODO: think what to do with the sensitive fields in the `options`?
		Import: func(ic *importContext, r *resource) error {
			connectionName := r.Data.Get("name").(string)
			ic.emitUCGrantsWithOwner("foreign_connection/"+connectionName, r)
			return nil
		},
		Ignore: func(ic *importContext, r *resource) bool {
			res := (r.Data.Get("connection_type").(string) == "ONLINE_CATALOG" &&
				strings.HasPrefix(r.Data.Get("name").(string), "internal-") &&
				r.Data.Get("owner").(string) == "System user")
			if res {
				ic.addIgnoredResource(fmt.Sprintf("databricks_connection. id=%s", r.ID))
			}
			return res
		},
		ShouldOmitField: shouldOmitForUnityCatalog,
		Depends: []reference{
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_share": {
		WorkspaceLevel:  true,
		Service:         "uc-shares",
		List:            listUcShares,
		Import:          importUcShare,
		ShouldOmitField: shouldOmitForUnityCatalog,
		Depends: []reference{
			{Path: "object.name", Resource: "databricks_volume", IsValidApproximation: isMatchignShareObject("VOLUME")},
			{Path: "object.name", Resource: "databricks_registered_model", IsValidApproximation: isMatchignShareObject("MODEL")},
			{Path: "object.name", Resource: "databricks_schema", IsValidApproximation: isMatchignShareObject("SCHEMA")},
			{Path: "object.name", Resource: "databricks_sql_table", IsValidApproximation: isMatchignShareObject("TABLE")},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_recipient": {
		WorkspaceLevel: true,
		Service:        "uc-shares",
		List:           listUcRecipients,
		Import:         importUcRecipient,
		Depends: []reference{
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
		// TODO: emit variable for sharing_code ...
		// TODO: add depends for sharing_code?
	},
	"databricks_registered_model": {
		WorkspaceLevel: true,
		Service:        "uc-models",
		// TODO: it doesn't work right now, need a fix in the Go SDK
		// List: func(ic *importContext) error {
		// 	models, err := ic.workspaceClient.RegisteredModels.ListAll(ic.Context, catalog.ListRegisteredModelsRequest{})
		// 	if err != nil {
		// 		return err
		// 	}
		// 	for _, model := range models {
		// TODO: Add name matching...
		// 		ic.EmitIfUpdatedAfterMillis(&resource{
		// 			Resource: "databricks_registered_model",
		// 			ID:       model.FullName,
		// 		}, model.UpdatedAt, fmt.Sprintf("registered model '%s'", model.FullName))
		// 	}
		// 	return nil
		// },
		Import: func(ic *importContext, r *resource) error {
			modelFullName := r.ID
			ic.emitUCGrantsWithOwner("model/"+modelFullName, r)
			schemaFullName := r.Data.Get("catalog_name").(string) + "." + r.Data.Get("schema_name").(string)
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       schemaFullName,
			})
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "storage_location" {
				location := d.Get(pathString).(string)
				if ic != nil && ic.currentMetastore != nil { // don't generate location it if it's managed...
					return strings.Contains(location, "/"+ic.currentMetastore.MetastoreId+"/models/")
				}
				return location == ""
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_registered_model", "name"),
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "storage_root", Resource: "databricks_external_location", Match: "url", MatchType: MatchLongestPrefix},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_metastore": {
		AccountLevel: true,
		Service:      "uc-metastores",
		Name:         makeNameOrIdFunc("name"),
		List:         listUcMetastores,
		Import:       importUcMetastores,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "default_data_access_config_id" || pathString == "storage_root_credential_id" {
				// technically, both should be marked as `computed`
				return true
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_metastore_assignment": {
		AccountLevel: true,
		Service:      "uc-metastores",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return fmt.Sprintf("ws_%d", d.Get("workspace_id").(int))
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "default_catalog_name" {
				return d.Get(pathString).(string) == ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "metastore_id", Resource: "databricks_metastore"},
		},
	},
	"databricks_workspace_binding": {
		WorkspaceLevel: true,
		Service:        "uc-catalogs",
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "securable_name" {
				return d.Get(pathString).(string) == ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "securable_name", Resource: "databricks_catalog", Match: "name",
				IsValidApproximation: isMatchingSecurableTypeAndName, SkipDirectLookup: true},
			{Path: "securable_name", Resource: "databricks_storage_credential", Match: "name",
				IsValidApproximation: isMatchingSecurableTypeAndName, SkipDirectLookup: true},
			{Path: "securable_name", Resource: "databricks_external_location", Match: "name",
				IsValidApproximation: isMatchingSecurableTypeAndName, SkipDirectLookup: true},
		},
	},
	"databricks_file": {
		WorkspaceLevel: true,
		Service:        "storage",
		// TODO: can we implement incremental mode?
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := strings.TrimPrefix(d.Id(), "/Volumes/")
			fileNameMd5 := fmt.Sprintf("%x", md5.Sum([]byte(name)))
			return strings.ToLower(name) + "_" + fileNameMd5[:8]
		},
		Import: importUcVolumeFile,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			switch pathString {
			case "md5", "remote_file_modified", "modification_time", "file_size":
				return true
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_volume", Match: "volume_path", MatchType: MatchLongestPrefix},
		},
	},
	"databricks_mws_permission_assignment": {
		AccountLevel: true,
		Service:      "idfed",
		List:         listMwsPermissionAssignments,
		Depends: []reference{
			{Resource: "databricks_service_principal", Path: "principal_id"},
			{Resource: "databricks_user", Path: "principal_id"},
			{Resource: "databricks_group", Path: "principal_id"},
			{Resource: "databricks_mws_workspaces", Path: "workspace_id", Match: "workspace_id"},
		},
	},
	"databricks_dashboard": {
		WorkspaceLevel: true,
		Service:        "dashboards",
		List:           listLakeviewDashboards,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			s := d.Get("parent_path").(string)
			if s != "" {
				s = s[1:]
				if s != "" {
					s = s + "_"
				}
			}
			dname := d.Get("display_name").(string)
			if dname != "" {
				s = s + dname
			}
			s = s + "_" + d.Id()
			return nameNormalizationRegex.ReplaceAllString(s, "_")
		},
		Import: importLakeviewDashboard,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			return pathString == "dashboard_change_detected" || shouldOmitMd5Field(ic, pathString, as, d, r)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			return ic.isInRepoOrGitFolder(r.Data.Get("path").(string), false) || ic.isInRepoOrGitFolder(r.Data.Get("parent_path").(string), true)
		},
		Depends: []reference{
			{Path: "file_path", File: true},
			{Path: "warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "parent_path", Resource: "databricks_user", Match: "home"},
			{Path: "parent_path", Resource: "databricks_service_principal", Match: "home"},
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
		},
	},
	"databricks_notification_destination": {
		WorkspaceLevel: true,
		Service:        "settings",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("display_name").(string)
			if name != "" {
				name += "_"
			}
			id := d.Id()
			if len(id) >= 8 {
				id = id[:8]
			}
			return nameNormalizationRegex.ReplaceAllString(fmt.Sprintf("%s_%s", name, id), "_")
		},
		List:            listNotificationDestinations,
		Import:          importNotificationDestination,
		ShouldOmitField: shouldOmitForNotificationDestination,
		Depends: []reference{
			{Path: "config.pagerduty.integration_key", Variable: true},
			{Path: "config.generic_webhook.url", Variable: true},
			{Path: "config.generic_webhook.username", Variable: true},
			{Path: "config.generic_webhook.password", Variable: true},
			{Path: "config.slack.url", Variable: true},
			{Path: "config.slack.channel_id", Variable: true},
			{Path: "config.slack.oauth_token", Variable: true},
			{Path: "config.microsoft_teams.url", Variable: true},
			{Path: "config.microsoft_teams.channel_url", Variable: true},
			{Path: "config.microsoft_teams.auth_secret", Variable: true},
			{Path: "config.microsoft_teams.tenant_id", Variable: true},
			{Path: "config.microsoft_teams.app_id", Variable: true},
			{Path: "config.email.addresses", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_workspace_setting_v2": {
		WorkspaceLevel:         true,
		Service:                "settings",
		PluginFramework:        true,
		List:                   listWorkspaceSettingsV2,
		Import:                 importWorkspaceSettingV2,
		ShouldOmitFieldUnified: shouldOmitWithEffectiveFields,
	},
	"databricks_account_setting_v2": {
		AccountLevel:           true,
		Service:                "settings",
		PluginFramework:        true,
		List:                   listAccountSettingsV2,
		Import:                 importAccountSettingV2,
		ShouldOmitFieldUnified: shouldOmitWithEffectiveFields,
	},
	"databricks_online_table": {
		WorkspaceLevel: true,
		Service:        "uc-online-tables",
		Import: func(ic *importContext, r *resource) error {
			ic.emitUCGrantsWithOwner("table/"+r.ID, r)
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       r.Data.Get("spec.0.source_table_full_name").(string),
			})
			return nil
		},
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_online_table", "name"),
		ShouldOmitField: shouldOmitForUnityCatalog,
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "spec.source_table_full_name", Resource: "databricks_sql_table"},
			{Path: "owner", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "owner", Resource: "databricks_group", Match: "display_name"},
			{Path: "owner", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		},
	},
	"databricks_vector_search_endpoint": {
		WorkspaceLevel: true,
		Service:        "vector-search",
		List:           listVectorSearchEndpoints,
		Import:         importVectorSearchEndpoint,
	},
	"databricks_vector_search_index": {
		WorkspaceLevel: true,
		Service:        "vector-search",
		Import:         importVectorSearchIndex,
		Depends: []reference{
			{Path: "delta_sync_index_spec.source_table", Resource: "databricks_sql_table"},
			{Path: "endpoint_name", Resource: "databricks_vector_search_endpoint"},
			{Path: "delta_sync_index_spec.embedding_source_columns.embedding_model_endpoint_name", Resource: "databricks_model_serving"},
			{Path: "delta_sync_index_spec.embedding_source_columns.model_endpoint_name_for_query", Resource: "databricks_model_serving"},
			{Path: "direct_access_index_spec.embedding_source_columns.embedding_model_endpoint_name", Resource: "databricks_model_serving"},
			{Path: "direct_access_index_spec.embedding_source_columns.model_endpoint_name_for_query", Resource: "databricks_model_serving"},
		},
	},
	"databricks_mws_network_connectivity_config": {
		AccountLevel: true,
		Service:      "nccs",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		List: func(ic *importContext) error {
			updatedSinceMs := ic.getUpdatedSinceMs()
			it := ic.accountClient.NetworkConnectivity.ListNetworkConnectivityConfigurations(ic.Context,
				settings.ListNetworkConnectivityConfigurationsRequest{})
			for it.HasNext(ic.Context) {
				nc, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				if !ic.MatchesName(nc.Name) {
					log.Printf("[INFO] Skipping mws_network_connectivity_config %s because it doesn't match %s", nc.Name, ic.match)
					continue
				}
				if ic.incremental && nc.UpdatedTime < updatedSinceMs {
					log.Printf("[DEBUG] skipping mws_network_connectivity_config '%s' that was modified at %d (last active=%d)",
						nc.Name, nc.UpdatedTime, updatedSinceMs)
					continue
				}
				// TODO: technically we can create data directly from the API response
				ic.Emit(&resource{
					Resource: "databricks_mws_network_connectivity_config",
					ID:       nc.AccountId + "/" + nc.NetworkConnectivityConfigId,
				})
				if nc.EgressConfig.TargetRules != nil {
					for _, rule := range nc.EgressConfig.TargetRules.AzurePrivateEndpointRules {
						// TODO: technically we can create data directly from the API response
						resourceId := strings.ReplaceAll(rule.ResourceId, "/subscriptions/", "")
						resourceId = strings.ReplaceAll(resourceId, "/resourceGroups/", "_")
						resourceId = strings.ReplaceAll(resourceId, "/providers/Microsoft", "_")
						ic.Emit(&resource{
							Resource: "databricks_mws_ncc_private_endpoint_rule",
							ID:       nc.NetworkConnectivityConfigId + "/" + rule.RuleId,
							Name:     nc.Name + "_" + resourceId + "_" + rule.GroupId,
						})
					}
					for _, rule := range nc.EgressConfig.TargetRules.AwsPrivateEndpointRules {
						ic.Emit(&resource{
							Resource: "databricks_mws_ncc_private_endpoint_rule",
							ID:       nc.NetworkConnectivityConfigId + "/" + rule.RuleId,
							Name:     nc.Name + "_" + rule.EndpointService,
						})
					}
				}
			}
			return nil
		},
	},
	"databricks_mws_ncc_private_endpoint_rule": {
		AccountLevel: true,
		Service:      "nccs",
		Depends: []reference{
			{Path: "network_connectivity_config_id", Resource: "databricks_mws_network_connectivity_config",
				Match: "network_connectivity_config_id"},
		},
	},
	"databricks_mws_ncc_binding": {
		AccountLevel: true,
		Service:      "nccs",
		List: func(ic *importContext) error {
			workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
			if err != nil {
				return err
			}
			for _, workspace := range workspaces {
				if workspace.NetworkConnectivityConfigId != "" {
					ic.emitNccBindingAndNcc(workspace.WorkspaceId, workspace.NetworkConnectivityConfigId)
					if !ic.accountClient.Config.IsAzure() {
						wsIdString := strconv.FormatInt(workspace.WorkspaceId, 10)
						ic.Emit(&resource{
							Resource: "databricks_mws_workspaces",
							ID:       ic.accountClient.Config.AccountID + "/" + wsIdString,
							Name:     workspace.WorkspaceName + "_" + wsIdString,
						})
					}
				}
			}
			return nil
		},
		Depends: []reference{
			{Path: "network_connectivity_config_id", Resource: "databricks_mws_network_connectivity_config",
				Match: "network_connectivity_config_id"},
			{Path: "workspace_id", Resource: "databricks_mws_workspaces", Match: "workspace_id"},
		},
	},
	"databricks_mws_credentials": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsCredentials,
	},
	"databricks_mws_storage_configurations": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsStorageConfigurations,
	},
	"databricks_mws_vpc_endpoint": {
		AccountLevel: true,
		Service:      "mws",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("vpc_endpoint_name").(string)
		},
		List: listMwsVpcEndpoints,
	},
	"databricks_mws_private_access_settings": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsPrivateAccessSettings,
		Import:       importMwsPrivateAccessSettings,
		Depends: []reference{
			{Path: "allowed_vpc_endpoint_ids", Resource: "databricks_mws_vpc_endpoint", Match: "vpc_endpoint_id"},
		},
	},
	"databricks_mws_customer_managed_keys": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsCustomerManagedKeys,
	},
	"databricks_mws_networks": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsNetworks,
		Import:       importMwsNetworks,
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
			if pathString == "vpc_endpoints" && d.Get("vpc_endpoints.#") != 0 {
				return false
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
		},
		Depends: []reference{
			{Path: "vpc_endpoints.dataplane_relay", Resource: "databricks_mws_vpc_endpoint", Match: "vpc_endpoint_id"},
			{Path: "vpc_endpoints.rest_api", Resource: "databricks_mws_vpc_endpoint", Match: "vpc_endpoint_id"},
		},
	},
	"databricks_mws_workspaces": {
		AccountLevel: true,
		Service:      "mws",
		List:         listMwsWorkspaces,
		Import:       importMwsWorkspaces,
		Depends: []reference{
			{Path: "network_id", Resource: "databricks_mws_networks", Match: "network_id"},
			{Path: "private_access_settings_id", Resource: "databricks_mws_private_access_settings", Match: "private_access_settings_id"},
			{Path: "storage_configuration_id", Resource: "databricks_mws_storage_configurations", Match: "storage_configuration_id"},
			{Path: "storage_customer_managed_key_id", Resource: "databricks_mws_customer_managed_keys", Match: "customer_managed_key_id"},
			{Path: "managed_services_customer_managed_key_id", Resource: "databricks_mws_customer_managed_keys", Match: "customer_managed_key_id"},
			{Path: "credentials_id", Resource: "databricks_mws_credentials", Match: "credentials_id"},
		},
	},
	"databricks_budget_policy": {
		AccountLevel:    true,
		PluginFramework: true,
		Service:         "billing",
		Name:            func(ic *importContext, d *schema.ResourceData) string { return d.Id() },
		List:            listBudgetPolicies,
		Import:          importBudgetPolicy,
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_budget_policy", "policy_id"),
		Depends: []reference{
			{Path: "binding_workspace_ids", Resource: "databricks_mws_workspaces", Match: "workspace_id"},
		},
	},
	"databricks_budget": {
		AccountLevel: true,
		Service:      "billing",
		List:         listBudgets,
		Import:       importBudget,
		Depends: []reference{
			{Path: "filter.workspace_id.values", Resource: "databricks_mws_workspaces", Match: "workspace_id"},
			{Path: "alert_configurations.action_configurations.target", Resource: "databricks_user", Match: "user_name"},
		},
	},
	"databricks_tag_policy": {
		WorkspaceLevel:  true,
		PluginFramework: true,
		Service:         "uc-tags",
		List:            listTagPolicies,
		// TODO: add import function that will emit access control rule set for the tag policy
		// This requires knowing the account ID, so will be added later
	},
}
