package exporter

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	tf_dlt "github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/repos"
	tf_settings "github.com/databricks/terraform-provider-databricks/settings"
	tf_sharing "github.com/databricks/terraform-provider-databricks/sharing"
	tf_sql "github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
	"golang.org/x/exp/maps"
)

var (
	adlsGen2Regex                    = regexp.MustCompile(`^(abfss?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	adlsGen1Regex                    = regexp.MustCompile(`^(adls?)://([^.]+)\.(?:[^/]+)(/.*)?$`)
	wasbsRegex                       = regexp.MustCompile(`^(wasbs?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	s3Regex                          = regexp.MustCompile(`^(s3a?)://([^/]+)(/.*)?$`)
	gsRegex                          = regexp.MustCompile(`^gs://([^/]+)(/.*)?$`)
	globalWorkspaceConfName          = "global_workspace_conf"
	nameNormalizationRegex           = regexp.MustCompile(`\W+`)
	fileNameNormalizationRegex       = regexp.MustCompile(`[^-_\w/.@]`)
	jobClustersRegex                 = regexp.MustCompile(`^((job_cluster|task)\.\d+\.new_cluster\.\d+\.)`)
	dltClusterRegex                  = regexp.MustCompile(`^(cluster\.\d+\.)`)
	secretPathRegex                  = regexp.MustCompile(`^\{\{secrets\/([^\/]+)\/([^}]+)\}\}$`)
	sqlParentRegexp                  = regexp.MustCompile(`^folders/(\d+)$`)
	dltDefaultStorageRegex           = regexp.MustCompile(`^dbfs:/pipelines/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	ignoreIdeFolderRegex             = regexp.MustCompile(`^/Users/[^/]+/\.ide/.*$`)
	servedEntityFieldExtractionRegex = regexp.MustCompile(`^config\.[0-9]+\.served_entities\.([0-9]+)\.(.*)$`)
	uc3LevelIdRegex                  = regexp.MustCompile(`^([^.]+\.[^.]+\.[^.]+)$`)
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

func generateMountBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	mount := ic.mountMap[r.ID]

	b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
	b.SetAttributeValue("name", cty.StringVal(strings.Replace(r.ID, "/mnt/", "", 1)))
	if res := s3Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("s3", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[2]))
		if mount.InstanceProfile != "" {
			block.SetAttributeValue("instance_profile", cty.StringVal(mount.InstanceProfile))
		} else if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := gsRegex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("gs", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[1]))
		if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := adlsGen2Regex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("abfs", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_abfs", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv2 resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("client_id", ic.variable(
			"client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable(
			"tenant_id"+varName, "Tenant ID"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
		block.SetAttributeValue("initialize_file_system", cty.BoolVal(false))
	} else if res := adlsGen1Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("adl", nil).Body()
		storageResourceName := res[2]
		block.SetAttributeValue("storage_resource_name", cty.StringVal(storageResourceName))
		if res[3] != "" && res[3] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[3]))
		}
		varName := ic.regexFix("_"+storageResourceName+"_adl", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv1 resource %s://%s", res[1], storageResourceName)

		block.SetAttributeRaw("client_id", ic.variable("client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable("tenant_id"+varName, "Tenant IDs"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName, "Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName, "Key in secret scope that stores app client secret"+textStr))
	} else if res := wasbsRegex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("wasb", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}
		block.SetAttributeValue("auth_type", cty.StringVal("ACCESS_KEY"))

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_wasb", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting WASB resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("token_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("token_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
	} else {
		return fmt.Errorf("no matching handler for: %s", mount.URL)
	}
	body.AppendNewline()

	return nil
}

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
		Import: func(ic *importContext, r *resource) error {
			dbfsAPI := storage.NewDbfsAPI(ic.Context, ic.Client)
			content, err := dbfsAPI.Read(r.ID)
			if err != nil {
				return err
			}
			name := ic.Importables["databricks_dbfs_file"].Name(ic, r.Data)
			fileName, err := ic.saveFileIn("dbfs_files", name, content)
			log.Printf("Creating %s for %s", fileName, r)
			if err != nil {
				return err
			}
			r.Data.Set("source", fileName)
			return nil
		},
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
		},
	},
	"databricks_instance_pool": {
		WorkspaceLevel: true,
		Service:        "pools",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			raw, ok := d.GetOk("instance_pool_name")
			if !ok || raw.(string) == "" {
				return strings.Split(d.Id(), "-")[2]
			}
			return raw.(string)
		},
		List: func(ic *importContext) error {
			it := ic.workspaceClient.InstancePools.List(ic.Context)
			i := 0
			for it.HasNext(ic.Context) {
				pool, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				i++
				if !ic.MatchesName(pool.InstancePoolName) {
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_instance_pool",
					ID:       pool.InstancePoolId,
				})
				if i%50 == 0 {
					log.Printf("[INFO] Imported %d instance pools", i)
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/instance-pools/%s", r.ID),
				"inst_pool_"+ic.Importables["databricks_instance_pool"].Name(ic, r.Data))
			return nil
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_instance_pool", "instance_pool_name"),
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
			{Path: "library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "policy_id", Resource: "databricks_cluster_policy"},
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
		List: listClusters,
		Import: func(ic *importContext, r *resource) error {
			var c compute.ClusterSpec
			s := ic.Resources["databricks_cluster"].Schema
			common.DataToStructPointer(r.Data, s, &c)
			ic.importCluster(&c)
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/clusters/%s", r.ID),
				"cluster_"+ic.Importables["databricks_cluster"].Name(ic, r.Data))
			return ic.importClusterLibraries(r.Data)
		},
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
		List: func(ic *importContext) error {
			builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
			it := ic.workspaceClient.ClusterPolicies.List(ic.Context, compute.ListClusterPoliciesRequest{})
			i := 0
			for it.HasNext(ic.Context) {
				policy, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				i++
				family, isBuiltin := builtInClusterPolicies[policy.PolicyFamilyId]
				if policy.PolicyFamilyId != "" && isBuiltin && family.Name == policy.Name &&
					policy.PolicyFamilyDefinitionOverrides == "" {
					log.Printf("[DEBUG] Skipping builtin cluster policy '%s' without overrides", policy.Name)
					continue
				}
				if !ic.MatchesName(policy.Name) {
					log.Printf("[DEBUG] Policy %s doesn't match %s filter", policy.Name, ic.match)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_cluster_policy",
					ID:       policy.PolicyId,
				})
				if i%10 == 0 {
					log.Printf("[INFO] Scanned %d cluster policies", i)
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/cluster-policies/%s", r.ID),
				"cluster_policy_"+ic.Importables["databricks_cluster_policy"].Name(ic, r.Data))

			var clusterPolicy compute.Policy
			s := ic.Resources["databricks_cluster_policy"].Schema
			common.DataToStructPointer(r.Data, s, &clusterPolicy)

			var definition map[string]map[string]any
			err := json.Unmarshal([]byte(clusterPolicy.Definition), &definition)
			if err != nil {
				return err
			}
			for k, policy := range definition {
				value, vok := policy["value"]
				defaultValue, dok := policy["defaultValue"]
				typ := policy["type"]
				if !vok && !dok {
					log.Printf("[DEBUG] Skipping policy element as it doesn't have both value and defaultValue. k='%v', policy='%v'",
						k, policy)
					continue
				}
				if k == "aws_attributes.instance_profile_arn" {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       eitherString(value, defaultValue),
					})
				}
				if k == "instance_pool_id" || k == "driver_instance_pool_id" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       eitherString(value, defaultValue),
					})
				}
				if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
					strings.HasSuffix(k, ".dbfs.destination") {
					ic.emitIfDbfsFile(eitherString(value, defaultValue))
				}
				if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
					strings.HasSuffix(k, ".volumes.destination") {
					ic.emitIfVolumeFile(eitherString(value, defaultValue))
				}
				if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
					strings.HasSuffix(k, ".workspace.destination") {
					ic.emitWorkspaceFileOrRepo(eitherString(value, defaultValue))
				}
				if typ == "fixed" && (strings.HasPrefix(k, "spark_conf.") || strings.HasPrefix(k, "spark_env_vars.")) {
					either := eitherString(value, defaultValue)
					if res := secretPathRegex.FindStringSubmatch(either); res != nil {
						ic.Emit(&resource{
							Resource: "databricks_secret_scope",
							ID:       res[1],
						})
					}
				}
			}

			for _, lib := range clusterPolicy.Libraries {
				ic.emitIfDbfsFile(lib.Whl)
				ic.emitIfDbfsFile(lib.Jar)
				ic.emitIfDbfsFile(lib.Egg)
				ic.emitIfWsfsFile(lib.Whl)
				ic.emitIfWsfsFile(lib.Jar)
				ic.emitIfWsfsFile(lib.Egg)
				ic.emitIfVolumeFile(lib.Whl)
				ic.emitIfVolumeFile(lib.Jar)
			}

			policyFamilyId := clusterPolicy.PolicyFamilyId
			if policyFamilyId != "" && clusterPolicy.Definition != "" {
				// we need to set definition to empty value because otherwise it will be put into
				// generated HCL code for data source, and it only supports the `name` attribute
				r.Data.Set("definition", "")
				builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
				v, isBuiltin := builtInClusterPolicies[policyFamilyId]
				if isBuiltin && clusterPolicy.PolicyFamilyDefinitionOverrides == "" && v.Name == clusterPolicy.Name {
					r.Mode = "data"
				}
			}

			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "definition" {
				return d.Get("policy_family_id").(string) != ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "display_name" {
				userName := d.Get("user_name").(string)
				displayName := d.Get("display_name").(string)
				return displayName == "" || userName == displayName
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "display_name" {
				if ic.Client.IsAzure() {
					applicationID := d.Get("application_id").(string)
					displayName := d.Get("display_name").(string)
					return applicationID == displayName
				}
				return false
			}
			// application_id should be provided only on Azure
			// TODO: how to support Databricks-managed SPs?
			if pathString == "application_id" {
				return !ic.Client.IsAzure()
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
			if backendType != "AZURE_KEYVAULT" {
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
		List: func(ic *importContext) error {
			if !ic.mounts {
				return nil
			}
			if err := ic.refreshMounts(); err != nil {
				return err
			}
			for mountName, source := range ic.mountMap {
				if !ic.MatchesName(mountName) {
					continue
				}
				if strings.HasPrefix(source.URL, "s3a://") {
					log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
					if source.InstanceProfile != "" {
						ic.Emit(&resource{
							Resource: "databricks_instance_profile",
							ID:       source.InstanceProfile,
						})
					} else if source.ClusterID != "" {
						ic.Emit(&resource{
							Resource: "databricks_cluster",
							ID:       source.ClusterID,
						})
					}
				} else if strings.HasPrefix(source.URL, "gs://") {
					if source.ClusterID != "" {
						ic.Emit(&resource{
							Resource: "databricks_cluster",
							ID:       source.ClusterID,
						})
					}
				} else if res := adlsGen2Regex.FindStringSubmatch(source.URL); res != nil {
				} else if res := adlsGen1Regex.FindStringSubmatch(source.URL); res != nil {
				} else if res := wasbsRegex.FindStringSubmatch(source.URL); res != nil {
				} else {
					log.Printf("[INFO] No matching handler for: %s", source.URL)
					continue
				}
				log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
				ic.Emit(&resource{
					Resource: "databricks_mount",
					ID:       mountName,
					Data: ic.Resources["databricks_mount"].Data(
						&terraform.InstanceState{
							ID:         mountName,
							Attributes: map[string]string{},
						}),
				})

			}
			return nil
		},
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
			fileName, err := ic.saveFileIn("global_init_scripts", fmt.Sprintf("%s.sh", ic.ResourceName(r)), content)
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
		Search: func(ic *importContext, r *resource) error {
			repoDir, err := ic.workspaceClient.Workspace.GetStatusByPath(ic.Context, r.Value)
			if err != nil {
				return err
			}
			if repoDir.ObjectType != sdk_workspace.ObjectTypeRepo {
				return fmt.Errorf("object %s is not a repo", r.Value)
			}
			if repoDir.ResourceId != "" {
				r.ID = repoDir.ResourceId
			} else {
				r.ID = strconv.FormatInt(repoDir.ObjectId, 10)
			}
			return nil
		},
		List: func(ic *importContext) error {
			it := ic.workspaceClient.Repos.List(ic.Context, sdk_workspace.ListReposRequest{PathPrefix: "/Workspace"})
			i := 1
			for it.HasNext(ic.Context) {
				repo, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				if repo.Url != "" {
					ic.Emit(&resource{
						Resource: "databricks_repo",
						ID:       strconv.FormatInt(repo.Id, 10),
					})
				} else {
					log.Printf("[WARN] ignoring databricks_repo without Git provider. Path: %s", repo.Path)
					ic.addIgnoredResource(fmt.Sprintf("databricks_repo. path=%s", repo.Path))
				}
				if i%50 == 0 {
					log.Printf("[INFO] Scanned %d repos", i)
				}
				i++
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			path := maybeStripWorkspacePrefix(r.Data.Get("path").(string))
			if strings.HasPrefix(path, "/Repos") {
				ic.emitUserOrServicePrincipalForPath(path, "/Repos")
			} else if strings.HasPrefix(path, "/Users") {
				ic.emitUserOrServicePrincipalForPath(path, "/Users")
			}
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/repos/%s", r.ID),
				"repo_"+ic.Importables["databricks_repo"].Name(ic, r.Data))
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			resp, err := ic.workspaceClient.Workspace.Export(ic.Context, sdk_workspace.ExportRequest{
				Path:   r.ID,
				Format: sdk_workspace.ExportFormat(ic.notebooksFormat),
			})
			if err != nil {
				if apierr.IsMissing(err) {
					ic.addIgnoredResource(fmt.Sprintf("databricks_notebook. path=%s", r.ID))
				}
				return err
			}
			var fileExtension string
			if ic.notebooksFormat == "SOURCE" {
				language := r.Data.Get("language").(string)
				fileExtension = fileExtensionLanguageMapping[language]
				r.Data.Set("language", "")
			} else {
				fileExtension = fileExtensionFormatMapping[ic.notebooksFormat]
			}
			r.Data.Set("format", ic.notebooksFormat)
			objectId := r.Data.Get("object_id").(int)
			name := fileNameNormalizationRegex.ReplaceAllString(r.ID[1:], "_") + "_" + strconv.Itoa(objectId) + fileExtension
			content, _ := base64.StdEncoding.DecodeString(resp.Content)
			fileName, err := ic.saveFileIn("notebooks", name, []byte(content))
			if err != nil {
				return err
			}
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/notebooks/%d", objectId),
				"notebook_"+ic.Importables["databricks_notebook"].Name(ic, r.Data))
			// TODO: it's not completely correct condition - we need to make emit smarter -
			// emit only if permissions are different from their parent's permission.
			ic.emitWorkspaceObjectParentDirectory(r)
			return r.Data.Set("source", fileName)
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			switch pathString {
			case "language":
				return d.Get("language") == ""
			case "format":
				return d.Get("format") == "SOURCE"
			}
			return shouldOmitMd5Field(ic, pathString, as, d)
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
		WorkspaceLevel: true,
		Service:        "wsfiles",
		Name:           workspaceObjectResouceName,
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			resp, err := ic.workspaceClient.Workspace.Export(ic.Context, sdk_workspace.ExportRequest{
				Path:   r.ID,
				Format: sdk_workspace.ExportFormatAuto,
			})
			if err != nil {
				if apierr.IsMissing(err) {
					ic.addIgnoredResource(fmt.Sprintf("databricks_workspace_file. path=%s", r.ID))
				}
				return err
			}
			objectId := r.Data.Get("object_id").(int)
			parts := strings.Split(r.ID, "/")
			plen := len(parts)
			if idx := strings.Index(parts[plen-1], "."); idx != -1 {
				parts[plen-1] = parts[plen-1][:idx] + "_" + strconv.Itoa(objectId) + parts[plen-1][idx:]
			} else {
				parts[plen-1] = parts[plen-1] + "_" + strconv.Itoa(objectId)
			}
			name := fileNameNormalizationRegex.ReplaceAllString(strings.Join(parts, "/")[1:], "_")
			content, _ := base64.StdEncoding.DecodeString(resp.Content)
			fileName, err := ic.saveFileIn("workspace_files", name, []byte(content))
			if err != nil {
				return err
			}

			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/files/%d", objectId),
				"ws_file_"+ic.Importables["databricks_workspace_file"].Name(ic, r.Data))

			// TODO: it's not completely correct condition - we need to make emit smarter -
			// emit only if permissions are different from their parent's permission.
			ic.emitWorkspaceObjectParentDirectory(r)
			log.Printf("[TRACE] Creating %s for %s", fileName, r)
			return r.Data.Set("source", fileName)
		},
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		ShouldGenerateField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "enable_serverless_compute" {
				return false
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
		List: func(ic *importContext) error {
			it := ic.workspaceClient.Pipelines.ListPipelines(ic.Context, pipelines.ListPipelinesRequest{
				MaxResults: 100,
			})
			i := 0
			for it.HasNext(ic.Context) {
				q, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				i++
				if !ic.MatchesName(q.Name) {
					continue
				}
				var modifiedAt int64
				if ic.incremental {
					pipeline, err := ic.workspaceClient.Pipelines.Get(ic.Context, pipelines.GetPipelineRequest{
						PipelineId: q.PipelineId,
					})
					if err != nil {
						return err
					}
					modifiedAt = pipeline.LastModified
				}
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_pipeline",
					ID:       q.PipelineId,
				}, modifiedAt, fmt.Sprintf("DLT Pipeline '%s'", q.Name))
				if i%100 == 0 {
					log.Printf("[INFO] Imported %d DLT Pipelines", i)
				}
			}
			log.Printf("[INFO] Listed %d DLT pipelines", i)
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var pipeline tf_dlt.Pipeline
			s := ic.Resources["databricks_pipeline"].Schema
			common.DataToStructPointer(r.Data, s, &pipeline)
			if pipeline.Deployment != nil && pipeline.Deployment.Kind == "BUNDLE" {
				log.Printf("[INFO] Skipping processing of DLT Pipeline with ID %s (%s) as deployed with DABs",
					r.ID, pipeline.Name)
				return nil
			}
			if pipeline.Catalog != "" {
				var schemaName string
				if pipeline.Target != "" {
					schemaName = pipeline.Target
				} else if pipeline.Schema != "" {
					schemaName = pipeline.Schema
				}
				if schemaName != "" {
					ic.Emit(&resource{
						Resource: "databricks_schema",
						ID:       pipeline.Catalog + "." + pipeline.Target,
					})
				}
			}
			if pipeline.Deployment == nil || pipeline.Deployment.Kind != "BUNDLE" {
				for _, lib := range pipeline.Libraries {
					if lib.Notebook != nil {
						ic.emitNotebookOrRepo(lib.Notebook.Path)
					}
					if lib.File != nil {
						ic.emitNotebookOrRepo(lib.File.Path)
					}
					ic.emitIfDbfsFile(lib.Jar)
					ic.emitIfDbfsFile(lib.Whl)
				}
			}
			// Emit clusters
			for _, cluster := range pipeline.Clusters {
				if cluster.AwsAttributes != nil && cluster.AwsAttributes.InstanceProfileArn != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       cluster.AwsAttributes.InstanceProfileArn,
					})
				}
				if cluster.InstancePoolId != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       cluster.InstancePoolId,
					})
				}
				if cluster.DriverInstancePoolId != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       cluster.DriverInstancePoolId,
					})
				}
				if cluster.PolicyId != "" {
					ic.Emit(&resource{
						Resource: "databricks_cluster_policy",
						ID:       cluster.PolicyId,
					})
				}
				ic.emitInitScripts(cluster.InitScripts)
				ic.emitSecretsFromSecretsPathMap(cluster.SparkConf)
				ic.emitSecretsFromSecretsPathMap(cluster.SparkEnvVars)
			}
			ic.emitFilesFromMap(pipeline.Configuration)
			ic.emitSecretsFromSecretsPathMap(pipeline.Configuration)
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/pipelines/%s", r.ID),
				"pipeline_"+ic.Importables["databricks_pipeline"].Name(ic, r.Data))
			if pipeline.Notifications != nil {
				for _, n := range pipeline.Notifications {
					ic.emitListOfUsers(n.EmailRecipients)
				}
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if res := dltClusterRegex.FindStringSubmatch(pathString); res != nil { // analyze DLT clusters
				return makeShouldOmitFieldForCluster(dltClusterRegex)(ic, pathString, as, d)
			}
			switch pathString {
			case "storage":
				return dltDefaultStorageRegex.FindStringSubmatch(d.Get("storage").(string)) != nil
			case "edition":
				return d.Get("edition").(string) == ""
			case "creator_user_name":
				return true
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
			{Path: "notification.email_recipients", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
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
			if mse.Config.ServedEntities != nil {
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
					}
				}
			}
			if mse.Config.AutoCaptureConfig != nil && mse.Config.AutoCaptureConfig.CatalogName != "" &&
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "config" {
				return d.Get("config.#").(int) == 0
			}
			if pathString == "config.0.traffic_config" || pathString == "config.0.auto_capture_config.0.enabled" ||
				(pathString == "config.0.auto_capture_config.0.table_name_prefix" && d.Get(pathString).(string) != "") {
				return false
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
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		ShouldGenerateField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
			{Path: "config.auto_capture_config.schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
			{Path: "ai_gateway.inference_table_config.catalog_name", Resource: "databricks_catalog"},
			{Path: "ai_gateway.inference_table_config.schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog_name", "schema_name"),
				SkipDirectLookup:     true},
		},
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "storage_location" {
				return d.Get("volume_type").(string) == "MANAGED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			switch pathString {
			case "storage_location":
				return d.Get("table_type").(string) == "MANAGED"
			case "enable_predictive_optimization":
				epo := d.Get(pathString).(string)
				return epo == "" || epo == "INHERIT"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if (pathString == "url" || pathString == "credential_name") && d.Get("name").(string) == dbManagedExternalLocation {
				return true
			}
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATION_MODE_ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		WorkspaceLevel: true,
		Service:        "uc-shares",
		List: func(ic *importContext) error {
			it := ic.workspaceClient.Shares.List(ic.Context, sharing.ListSharesRequest{})
			for it.HasNext(ic.Context) {
				share, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
					Resource: "databricks_share",
					ID:       share.Name,
				}, share.Name, share.UpdatedAt, fmt.Sprintf("share '%s'", share.Name))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var share tf_sharing.ShareInfo
			s := ic.Resources["databricks_share"].Schema
			common.DataToStructPointer(r.Data, s, &share)
			// TODO: how to link recipients to share?
			ic.emitUCGrantsWithOwner("share/"+r.ID, r)
			for _, obj := range share.Objects {
				switch obj.DataObjectType {
				case "TABLE":
					ic.Emit(&resource{
						Resource: "databricks_sql_table",
						ID:       obj.Name,
					})
				case "VOLUME":
					ic.Emit(&resource{
						Resource: "databricks_volume",
						ID:       obj.Name,
					})
				case "MODEL":
					ic.Emit(&resource{
						Resource: "databricks_registered_model",
						ID:       obj.Name,
					})
				default:
					log.Printf("[INFO] Object type '%s' (name: '%s') isn't supported in share '%s'",
						obj.DataObjectType, obj.Name, r.ID)
				}
			}

			return nil
		},
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
		List: func(ic *importContext) error {
			it := ic.workspaceClient.Recipients.List(ic.Context, sharing.ListRecipientsRequest{})
			for it.HasNext(ic.Context) {
				rec, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
					Resource: "databricks_recipient",
					ID:       rec.Name,
				}, rec.Name, rec.UpdatedAt, fmt.Sprintf("recipient '%s'", rec.Name))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			owner := r.Data.Get("owner").(string)
			if owner != "" {
				emitUserSpOrGroup(ic, owner)
			}
			return nil
		},
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "storage_location" {
				location := d.Get(pathString).(string)
				if ic != nil && ic.currentMetastore != nil { // don't generate location it if it's managed...
					return strings.Contains(location, "/"+ic.currentMetastore.MetastoreId+"/models/")
				}
				return location == ""
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "default_data_access_config_id" || pathString == "storage_root_credential_id" {
				// technically, both should be marked as `computed`
				return true
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "default_catalog_name" {
				return d.Get(pathString).(string) == ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Depends: []reference{
			{Path: "metastore_id", Resource: "databricks_metastore"},
		},
	},
	"databricks_workspace_binding": {
		WorkspaceLevel: true,
		Service:        "uc-catalogs",
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "securable_name" {
				return d.Get(pathString).(string) == ""
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			switch pathString {
			case "md5", "remote_file_modified", "modification_time", "file_size":
				return true
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_volume", Match: "volume_path", MatchType: MatchLongestPrefix},
		},
	},
	"databricks_mws_permission_assignment": {
		AccountLevel: true,
		Service:      "idfed",
		List: func(ic *importContext) error {
			workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
			if err != nil {
				return err
			}
			for _, ws := range workspaces {
				// list only specific workspaces if ic.match is set
				if !ic.MatchesName(strconv.FormatInt(ws.WorkspaceId, 10)) {
					log.Printf("[DEBUG] Skipping workspace %d because it doesn't match to the filter", ws.WorkspaceId)
					continue
				}
				pas, err := ic.accountClient.WorkspaceAssignment.ListByWorkspaceId(ic.Context, ws.WorkspaceId)
				if err != nil {
					log.Printf("[ERROR] listing workspace permission assignments for workspace %d: %s",
						ws.WorkspaceId, err.Error())
					continue
				}
				log.Printf("[DEBUG] Emitting permission assignments for workspace %d", ws.WorkspaceId)
				for _, pa := range pas.PermissionAssignments {
					perm := "unknown"
					if len(pa.Permissions) > 0 {
						perm = pa.Permissions[0].String()
					}
					nm := fmt.Sprintf("mws_pa_%d_%s_%s_%d", ws.WorkspaceId, pa.Principal.DisplayName,
						perm, pa.Principal.PrincipalId)
					// We  generate Data directly to avoid calling APIs
					data := mws.ResourceMwsPermissionAssignment().ToResource().TestResourceData()
					paId := fmt.Sprintf("%d|%d", ws.WorkspaceId, pa.Principal.PrincipalId)
					data = ic.generateNewData(data, "databricks_mws_permission_assignment", paId, pa)
					data.Set("workspace_id", ws.WorkspaceId)
					data.Set("principal_id", pa.Principal.PrincipalId)
					ic.Emit(&resource{
						Resource: "databricks_mws_permission_assignment",
						ID:       paId,
						Name:     nameNormalizationRegex.ReplaceAllString(nm, "_"),
						Data:     data,
					})
					// Emit principals
					strPrincipalId := strconv.FormatInt(pa.Principal.PrincipalId, 10)
					if pa.Principal.ServicePrincipalName != "" {
						ic.Emit(&resource{
							Resource: "databricks_service_principal",
							ID:       strPrincipalId,
						})
					} else if pa.Principal.UserName != "" {
						ic.Emit(&resource{
							Resource: "databricks_user",
							ID:       strPrincipalId,
						})
					} else if pa.Principal.GroupName != "" {
						ic.Emit(&resource{
							Resource: "databricks_group",
							ID:       strPrincipalId,
						})
					}
				}
			}
			return nil
		},
		Depends: []reference{
			{Resource: "databricks_service_principal", Path: "principal_id"},
			{Resource: "databricks_user", Path: "principal_id"},
			{Resource: "databricks_group", Path: "principal_id"},
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
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			return pathString == "dashboard_change_detected" || shouldOmitMd5Field(ic, pathString, as, d)
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
		List: func(ic *importContext) error {
			if !ic.meAdmin {
				return fmt.Errorf("notifications can be imported only by admin")
			}
			it := ic.workspaceClient.NotificationDestinations.List(ic.Context, settings.ListNotificationDestinationsRequest{})
			for it.HasNext(ic.Context) {
				n, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				ic.Emit(&resource{
					Resource: "databricks_notification_destination",
					ID:       n.Id,
				})
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var notificationDestination tf_settings.NDStruct
			s := ic.Resources["databricks_notification_destination"].Schema
			common.DataToStructPointer(r.Data, s, &notificationDestination)
			if notificationDestination.DestinationType == "EMAIL" && notificationDestination.Config != nil &&
				notificationDestination.Config.Email != nil {
				for _, email := range notificationDestination.Config.Email.Addresses {
					ic.emitUserOrServicePrincipal(email)
				}
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			var notificationDestination tf_settings.NDStruct
			s := ic.Resources["databricks_notification_destination"].Schema
			common.DataToStructPointer(d, s, &notificationDestination)
			if notificationDestination.Config != nil {
				switch notificationDestination.DestinationType {
				case "WEBHOOK":
					if notificationDestination.Config.GenericWebhook != nil {
						switch pathString {
						case "config.0.generic_webhook.0.url":
							return !notificationDestination.Config.GenericWebhook.UrlSet
						case "config.0.generic_webhook.0.username":
							return !notificationDestination.Config.GenericWebhook.UsernameSet
						case "config.0.generic_webhook.0.password":
							return !notificationDestination.Config.GenericWebhook.PasswordSet
						}
					}
				case "SLACK":
					if notificationDestination.Config.Slack != nil && pathString == "config.0.slack.0.url" {
						return !notificationDestination.Config.Slack.UrlSet
					}
				case "PAGERDUTY":
					if notificationDestination.Config.Pagerduty != nil && pathString == "config.0.pagerduty.0.integration_key" {
						return !notificationDestination.Config.Pagerduty.IntegrationKeySet
					}
				case "MICROSOFT_TEAMS":
					if notificationDestination.Config.MicrosoftTeams != nil && pathString == "config.0.microsoft_teams.0.url" {
						return !notificationDestination.Config.MicrosoftTeams.UrlSet
					}
				}
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Depends: []reference{
			{Path: "config.email.addresses", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "config.microsoft_teams.url", Variable: true},
			{Path: "config.pagerduty.integration_key", Variable: true},
			{Path: "config.generic_webhook.url", Variable: true},
			{Path: "config.generic_webhook.username", Variable: true},
			{Path: "config.generic_webhook.password", Variable: true},
			{Path: "config.slack.url", Variable: true},
		},
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
		List: func(ic *importContext) error {
			endpoints, err := ic.workspaceClient.VectorSearchEndpoints.ListEndpointsAll(ic.Context, vectorsearch.ListEndpointsRequest{})
			if err != nil {
				log.Printf("[ERROR] listing vector search endpoints: %s", err.Error())
				return err
			}
			for _, ep := range endpoints {
				ic.EmitIfUpdatedAfterMillis(&resource{
					Resource: "databricks_vector_search_endpoint",
					ID:       ep.Name,
				}, ep.LastUpdatedTimestamp, fmt.Sprintf("vector search endpoint '%s'", ep.Name))

			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			indexes, err := ic.workspaceClient.VectorSearchIndexes.ListIndexesAll(ic.Context, vectorsearch.ListIndexesRequest{
				EndpointName: r.ID,
			})
			if err != nil {
				log.Printf("[ERROR] listing vector search indexes for endpoint %s: %s", r.ID, err.Error())
				return err
			}
			for _, idx := range indexes {
				ic.Emit(&resource{
					Resource: "databricks_vector_search_index",
					ID:       idx.Name,
				})
			}
			return nil
		},
	},
	"databricks_vector_search_index": {
		WorkspaceLevel: true,
		Service:        "vector-search",
		Import: func(ic *importContext, r *resource) error {
			ic.emitUCGrantsWithOwner("table/"+r.ID, r)
			s := ic.Resources["databricks_vector_search_index"].Schema
			var vsi vectorsearch.VectorIndex
			common.DataToStructPointer(r.Data, s, &vsi)
			if vsi.EndpointName != "" {
				ic.Emit(&resource{
					Resource: "databricks_vector_search_endpoint",
					ID:       vsi.EndpointName,
				})
			}
			if vsi.DeltaSyncIndexSpec != nil {
				ic.Emit(&resource{
					Resource: "databricks_sql_table",
					ID:       vsi.DeltaSyncIndexSpec.SourceTable,
				})
				if vsi.DeltaSyncIndexSpec.EmbeddingWritebackTable != "" {
					ic.Emit(&resource{
						Resource: "databricks_sql_table",
						ID:       vsi.DeltaSyncIndexSpec.EmbeddingWritebackTable,
					})
				}
				for _, col := range vsi.DeltaSyncIndexSpec.EmbeddingSourceColumns {
					if col.EmbeddingModelEndpointName != "" {
						ic.Emit(&resource{
							Resource: "databricks_model_serving",
							ID:       col.EmbeddingModelEndpointName,
						})
					}
				}
			}
			if vsi.DirectAccessIndexSpec != nil {
				for _, col := range vsi.DirectAccessIndexSpec.EmbeddingSourceColumns {
					if col.EmbeddingModelEndpointName != "" {
						ic.Emit(&resource{
							Resource: "databricks_model_serving",
							ID:       col.EmbeddingModelEndpointName,
						})
					}
				}
			}
			return nil
		},
		Depends: []reference{
			{Path: "delta_sync_index_spec.source_table", Resource: "databricks_sql_table"},
			{Path: "endpoint_name", Resource: "databricks_vector_search_endpoint"},
			{Path: "delta_sync_index_spec.embedding_source_columns.embedding_model_endpoint_name", Resource: "databricks_model_serving"},
			{Path: "direct_access_index_spec.embedding_source_columns.embedding_model_endpoint_name", Resource: "databricks_model_serving"},
		},
	},
	"databricks_mws_network_connectivity_config": {
		AccountLevel: true,
		Service:      "nccs",
		List: func(ic *importContext) error {
			updatedSinceMs := ic.getUpdatedSinceMs()
			it := ic.accountClient.NetworkConnectivity.ListNetworkConnectivityConfigurations(ic.Context,
				settings.ListNetworkConnectivityConfigurationsRequest{})
			for it.HasNext(ic.Context) {
				nc, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				if ic.incremental && nc.UpdatedTime < updatedSinceMs {
					log.Printf("[DEBUG] skipping %s that was modified at %d (last active=%d)",
						fmt.Sprintf("network connectivity config '%s'", nc.Name), nc.UpdatedTime, updatedSinceMs)
					continue
				}
				// TODO: technically we can create data directly from the API response
				ic.Emit(&resource{
					Resource: "databricks_mws_network_connectivity_config",
					ID:       nc.AccountId + "/" + nc.NetworkConnectivityConfigId,
					Name:     nc.Name,
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
							Name:     nc.Name + "_" + resourceId + "_" + rule.GroupId.String(),
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
}
