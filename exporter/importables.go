package exporter

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/databricks-sdk-go/service/iam"
	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
	tfcatalog "github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	tfpipelines "github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/repos"
	tfsettings "github.com/databricks/terraform-provider-databricks/settings"
	tfsharing "github.com/databricks/terraform-provider-databricks/sharing"
	tfsql "github.com/databricks/terraform-provider-databricks/sql"
	sql_api "github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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
	ParentDirectoryExtraKey = "parent_directory"
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
				log.Printf("[INFO] Imported %d instance pools", i)
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
		Service: "access",
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
			{Path: "single_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "library.jar", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.whl", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "library.egg", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
		List: func(ic *importContext) error {
			clusters, err := clusters.NewClustersAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			lastActiveMs := ic.getLastActiveMs()
			nonInteractiveClusters := []string{"JOB", "MODELS", "PIPELINE_MAINTENANCE", "PIPELINE", "SQL"}
			for offset, c := range clusters {
				if slices.Contains(nonInteractiveClusters, string(c.ClusterSource)) {
					// TODO: Should we check cluster name as well?
					// jobRunClusterNameRegex = regexp.MustCompile(`^job-\d+-run-\d+$`)
					// jobRunClusterNameRegex.MatchString(c.ClusterName)
					log.Printf("[INFO] Skipping non-interactive cluster %s", c.ClusterID)
					continue
				}
				if strings.HasPrefix(c.ClusterName, "terraform-") {
					log.Printf("[INFO] Skipping terraform-specific cluster %s", c.ClusterName)
					continue
				}
				if !ic.MatchesName(c.ClusterName) {
					log.Printf("[INFO] Skipping %s because it doesn't match %s", c.ClusterName, ic.match)
					continue
				}
				if c.LastActivityTime > 0 && c.LastActivityTime < lastActiveMs {
					log.Printf("[INFO] Older inactive cluster %s", c.ClusterName)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       c.ClusterID,
				})
				log.Printf("[INFO] Scanned %d of %d clusters", offset+1, len(clusters))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var c compute.ClusterSpec
			s := ic.Resources["databricks_cluster"].Schema
			common.DataToStructPointer(r.Data, s, &c)
			ic.importCluster(&c)
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/clusters/%s", r.ID),
				"cluster_"+ic.Importables["databricks_cluster"].Name(ic, r.Data))
			return ic.importClusterLibraries(r.Data, s)
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
		Depends: []reference{
			{Path: "job_cluster.new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "job_cluster.new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "job_cluster.new_cluster.init_scripts.volumes.destination", Resource: "databricks_file"},
			{Path: "job_cluster.new_cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "job_cluster.new_cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "job_cluster.new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "job_cluster.new_cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "run_as.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "task.dbt_task.warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "task.existing_cluster_id", Resource: "databricks_cluster"},
			{Path: "task.library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.jar", Resource: "databricks_file"},
			{Path: "task.library.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.whl", Resource: "databricks_file"},
			{Path: "task.library.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.library.requirements", Resource: "databricks_file"},
			{Path: "task.library.requirements", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "task.new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.new_cluster.init_scripts.volumes.destination", Resource: "databricks_file"},
			{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "task.new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "task.new_cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "task.new_cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_file"},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_notebook"},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_notebook", Match: "workspace_path"},
			{Path: "task.notebook_task.warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "task.pipeline_task.pipeline_id", Resource: "databricks_pipeline"},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_file"},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.run_job_task.job_id", Resource: "databricks_job"},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.spark_jar_task.jar_uri", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_file"},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.spark_python_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_workspace_file", Match: "path"},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_file"},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.sql_task.file.path", Resource: "databricks_workspace_file", Match: "path"},
			{Path: "task.dbt_task.project_directory", Resource: "databricks_directory", Match: "path"},
			{Path: "task.sql_task.alert.alert_id", Resource: "databricks_alert"},
			{Path: "task.sql_task.dashboard.dashboard_id", Resource: "databricks_sql_dashboard"},
			{Path: "task.sql_task.query.query_id", Resource: "databricks_query"},
			{Path: "task.sql_task.warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "task.webhook_notifications.on_duration_warning_threshold_exceeded.id", Resource: "databricks_notification_destination"},
			{Path: "task.webhook_notifications.on_failure.id", Resource: "databricks_notification_destination"},
			{Path: "task.webhook_notifications.on_start.id", Resource: "databricks_notification_destination"},
			{Path: "task.webhook_notifications.on_success.id", Resource: "databricks_notification_destination"},
			{Path: "task.webhook_notifications.on_streaming_backlog_exceeded.id", Resource: "databricks_notification_destination"},
			{Path: "task.email_notifications.on_duration_warning_threshold_exceeded", Resource: "databricks_user",
				Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.email_notifications.on_failure", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.email_notifications.on_start", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.email_notifications.on_success", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.email_notifications.on_streaming_backlog_exceeded", Resource: "databricks_user",
				Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "run_as.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "webhook_notifications.on_duration_warning_threshold_exceeded.id", Resource: "databricks_notification_destination"},
			{Path: "webhook_notifications.on_failure.id", Resource: "databricks_notification_destination"},
			{Path: "webhook_notifications.on_start.id", Resource: "databricks_notification_destination"},
			{Path: "webhook_notifications.on_success.id", Resource: "databricks_notification_destination"},
			{Path: "webhook_notifications.on_streaming_backlog_exceeded.id", Resource: "databricks_notification_destination"},
			{Path: "email_notifications.on_duration_warning_threshold_exceeded", Resource: "databricks_user",
				Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_failure", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_start", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_success", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_streaming_backlog_exceeded", Resource: "databricks_user",
				Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.library.whl", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_repo", Match: "path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
			{Path: "job_cluster.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
				MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		},
		Import: func(ic *importContext, r *resource) error {
			var job jobs.JobSettingsResource
			s := ic.Resources["databricks_job"].Schema
			common.DataToStructPointer(r.Data, s, &job)
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/jobs/%s", r.ID),
				"job_"+ic.Importables["databricks_job"].Name(ic, r.Data))
			for _, task := range job.Tasks {
				if task.NotebookTask != nil {
					if task.NotebookTask.Source != "GIT" {
						ic.emitNotebookOrRepo(task.NotebookTask.NotebookPath)
					}
					ic.emitFilesFromMap(task.NotebookTask.BaseParameters)
					if task.NotebookTask.WarehouseId != "" {
						ic.Emit(&resource{
							Resource: "databricks_sql_endpoint",
							ID:       task.NotebookTask.WarehouseId,
						})
					}
				}
				if task.PipelineTask != nil {
					ic.Emit(&resource{
						Resource: "databricks_pipeline",
						ID:       task.PipelineTask.PipelineId,
					})
				}
				if task.SparkPythonTask != nil {
					if task.SparkPythonTask.Source != "GIT" {
						if strings.HasPrefix(task.SparkPythonTask.PythonFile, "dbfs:") {
							ic.Emit(&resource{
								Resource: "databricks_dbfs_file",
								ID:       task.SparkPythonTask.PythonFile,
							})
						} else {
							ic.emitWorkspaceFileOrRepo(task.SparkPythonTask.PythonFile)
						}
					}
					ic.emitFilesFromSlice(task.SparkPythonTask.Parameters)
				}
				if task.PythonWheelTask != nil {
					ic.emitFilesFromSlice(task.PythonWheelTask.Parameters)
					ic.emitFilesFromMap(task.PythonWheelTask.NamedParameters)
				}
				if task.SparkJarTask != nil {
					ic.emitFilesFromSlice(task.SparkJarTask.Parameters)
				}
				if task.SparkSubmitTask != nil {
					ic.emitFilesFromSlice(task.SparkSubmitTask.Parameters)
				}
				if task.SqlTask != nil {
					if task.SqlTask.Query != nil {
						ic.Emit(&resource{
							Resource: "databricks_query",
							ID:       task.SqlTask.Query.QueryId,
						})
					}
					if task.SqlTask.Dashboard != nil {
						ic.Emit(&resource{
							Resource: "databricks_sql_dashboard",
							ID:       task.SqlTask.Dashboard.DashboardId,
						})
					}
					if task.SqlTask.Alert != nil {
						ic.Emit(&resource{
							Resource: "databricks_alert",
							ID:       task.SqlTask.Alert.AlertId,
						})
					}
					if task.SqlTask.WarehouseId != "" {
						ic.Emit(&resource{
							Resource: "databricks_sql_endpoint",
							ID:       task.SqlTask.WarehouseId,
						})
					}
					if task.SqlTask.File != nil && task.SqlTask.File.Source == "WORKSPACE" {
						ic.emitWorkspaceFileOrRepo(task.SqlTask.File.Path)
					}
				}
				if task.DbtTask != nil {
					if task.DbtTask.WarehouseId != "" {
						ic.Emit(&resource{
							Resource: "databricks_sql_endpoint",
							ID:       task.DbtTask.WarehouseId,
						})
					}
					if task.DbtTask.Source == "WORKSPACE" {
						directory := task.DbtTask.ProjectDirectory
						if strings.HasPrefix(directory, "/Repos") {
							ic.emitRepoByPath(directory)
						} else {
							// Traverse the dbt project directory and emit all objects found in it
							nbAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
							objects, err := nbAPI.List(directory, true, true)
							if err == nil {
								for _, object := range objects {
									if object.ObjectType != workspace.File {
										continue
									}
									ic.maybeEmitWorkspaceObject("databricks_workspace_file", object.Path, &object)
								}
							} else {
								log.Printf("[WARN] Can't list directory %s for DBT task in job %s (id: %s)", directory, job.Name, r.ID)
							}
						}
					}
				}
				if task.RunJobTask != nil && task.RunJobTask.JobId != 0 {
					ic.Emit(&resource{
						Resource: "databricks_job",
						ID:       strconv.FormatInt(task.RunJobTask.JobId, 10),
					})
					ic.emitFilesFromMap(task.RunJobTask.JobParameters)
				}
				ic.importCluster(task.NewCluster)
				if task.ExistingClusterId != "" {
					ic.Emit(&resource{
						Resource: "databricks_cluster",
						ID:       task.ExistingClusterId,
					})
				}
				ic.emitLibraries(task.Libraries)

				if task.WebhookNotifications != nil {
					ic.emitJobsDestinationNotifications(task.WebhookNotifications.OnFailure)
					ic.emitJobsDestinationNotifications(task.WebhookNotifications.OnSuccess)
					ic.emitJobsDestinationNotifications(task.WebhookNotifications.OnDurationWarningThresholdExceeded)
					ic.emitJobsDestinationNotifications(task.WebhookNotifications.OnStart)
					ic.emitJobsDestinationNotifications(task.WebhookNotifications.OnStreamingBacklogExceeded)
				}
				if task.EmailNotifications != nil {
					ic.emitListOfUsers(task.EmailNotifications.OnDurationWarningThresholdExceeded)
					ic.emitListOfUsers(task.EmailNotifications.OnFailure)
					ic.emitListOfUsers(task.EmailNotifications.OnStart)
					ic.emitListOfUsers(task.EmailNotifications.OnSuccess)
					ic.emitListOfUsers(task.EmailNotifications.OnStreamingBacklogExceeded)
				}
			}
			for _, jc := range job.JobClusters {
				ic.importCluster(&jc.NewCluster)
			}
			if job.RunAs != nil {
				if job.RunAs.UserName != "" {
					ic.Emit(&resource{
						Resource:  "databricks_user",
						Attribute: "user_name",
						Value:     job.RunAs.UserName,
					})
				}
				if job.RunAs.ServicePrincipalName != "" {
					ic.Emit(&resource{
						Resource:  "databricks_service_principal",
						Attribute: "application_id",
						Value:     job.RunAs.ServicePrincipalName,
					})
				}
			}
			if job.EmailNotifications != nil {
				ic.emitListOfUsers(job.EmailNotifications.OnDurationWarningThresholdExceeded)
				ic.emitListOfUsers(job.EmailNotifications.OnFailure)
				ic.emitListOfUsers(job.EmailNotifications.OnStart)
				ic.emitListOfUsers(job.EmailNotifications.OnSuccess)
				ic.emitListOfUsers(job.EmailNotifications.OnStreamingBacklogExceeded)
			}
			if job.WebhookNotifications != nil {
				ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnFailure)
				ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnSuccess)
				ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnDurationWarningThresholdExceeded)
				ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnStart)
				ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnStreamingBacklogExceeded)
			}

			return ic.importLibraries(r.Data, s)
		},
		List: func(ic *importContext) error {
			if l, err := jobs.NewJobsAPI(ic.Context, ic.Client).List(); err == nil {
				ic.importJobs(l)
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			switch pathString {
			case "url", "format":
				return true
			}
			var js jobs.JobSettingsResource
			common.DataToStructPointer(d, ic.Resources["databricks_job"].Schema, &js)
			switch pathString {
			case "email_notifications":
				if js.EmailNotifications != nil {
					return reflect.DeepEqual(*js.EmailNotifications, sdk_jobs.JobEmailNotifications{})
				}
			case "webhook_notifications":
				if js.WebhookNotifications != nil {
					return reflect.DeepEqual(*js.WebhookNotifications, sdk_jobs.WebhookNotifications{})
				}
			case "notification_settings":
				if js.NotificationSettings != nil {
					return reflect.DeepEqual(*js.NotificationSettings, sdk_jobs.JobNotificationSettings{})
				}
			case "run_as":
				if js.RunAs != nil && (js.RunAs.UserName != "" || js.RunAs.ServicePrincipalName != "") {
					var user string
					if js.RunAs.UserName != "" {
						user = js.RunAs.UserName
					} else {
						user = js.RunAs.ServicePrincipalName
					}
					return user == ic.meUserName
				}
				return true
			}
			if strings.HasPrefix(pathString, "task.") {
				parts := strings.Split(pathString, ".")
				if len(parts) > 2 {
					taskIndex, err := strconv.Atoi(parts[1])
					if err == nil && taskIndex >= 0 && taskIndex < len(js.Tasks) {
						blockName := parts[len(parts)-1]
						switch blockName {
						case "notification_settings":
							if js.Tasks[taskIndex].NotificationSettings != nil {
								return reflect.DeepEqual(*js.Tasks[taskIndex].NotificationSettings,
									sdk_jobs.TaskNotificationSettings{})
							}
						case "email_notifications":
							if js.Tasks[taskIndex].EmailNotifications != nil {
								return reflect.DeepEqual(*js.Tasks[taskIndex].EmailNotifications,
									sdk_jobs.TaskEmailNotifications{})
							}
						case "webhook_notifications":
							if js.Tasks[taskIndex].WebhookNotifications != nil {
								return reflect.DeepEqual(*js.Tasks[taskIndex].WebhookNotifications,
									sdk_jobs.WebhookNotifications{})
							}
						}
					}
				}
				if strings.HasSuffix(pathString, ".notebook_task.0.source") && js.GitSource == nil && d.Get(pathString).(string) == "WORKSPACE" {
					return true
				}
				// TODO: add should omit for new cluster in the task?
				// TODO: double check it
			}
			if res := jobClustersRegex.FindStringSubmatch(pathString); res != nil { // analyze job clusters
				return makeShouldOmitFieldForCluster(jobClustersRegex)(ic, pathString, as, d)
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			numTasks := r.Data.Get("task.#").(int)
			if numTasks == 0 {
				log.Printf("[WARN] Ignoring job with ID %s", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_job. id=%s", r.ID))
			}
			return numTasks == 0
		},
	},
	"databricks_cluster_policy": {
		WorkspaceLevel: true,
		Service:        "policies",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		List: func(ic *importContext) error {
			w, err := ic.Client.WorkspaceClient()
			if err != nil {
				return err
			}
			builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
			it := w.ClusterPolicies.List(ic.Context, compute.ListClusterPoliciesRequest{})
			i := 0
			for it.HasNext(ic.Context) {
				policy, err := it.Next(ic.Context)
				if err != nil {
					return err
				}
				i++
				log.Printf("[TRACE] Scanning %d:  %v", i, policy)
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
		// TODO: special formatting required, where JSON is written line by line
		// so that we're able to do the references
		Body: resourceOrDataBlockBody,
	},
	"databricks_group": {
		Service:        "groups",
		WorkspaceLevel: true,
		AccountLevel:   true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("display_name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for offset, g := range ic.allGroups {
				if !ic.MatchesName(g.DisplayName) {
					log.Printf("[INFO] Group %s doesn't match %s filter", g.DisplayName, ic.match)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       g.ID,
				})
				log.Printf("[INFO] Scanned %d of %d groups", offset+1, len(ic.allGroups))
			}
			return nil
		},
		Search: func(ic *importContext, r *resource) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				if g.DisplayName == r.Value && r.Attribute == "display_name" {
					r.ID = g.ID
					return nil
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			groupName := r.Data.Get("display_name").(string)
			if (!ic.accountLevel && (groupName == "admins" || groupName == "users")) ||
				(ic.accountLevel && groupName == "account users") {
				// Workspace admins & users or Account users are to be imported through "data block"
				r.Mode = "data"
				r.Data.Set("workspace_access", false)
				r.Data.Set("databricks_sql_access", false)
				r.Data.Set("allow_instance_pool_create", false)
				r.Data.Set("allow_cluster_create", false)
				r.Data.State().Set(&terraform.InstanceState{
					ID: r.ID,
					Attributes: map[string]string{
						"display_name": r.Name,
					},
				})
			} else if r.Data != nil {
				r.Data.Set("force", true)
			}
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				if r.ID != g.ID {
					continue
				}
				ic.emitRoles("group", g.ID, g.Roles)
				builtInUserGroup := (ic.accountLevel && g.DisplayName == "account users") || (!ic.accountLevel && g.DisplayName == "users")
				if builtInUserGroup && !ic.importAllUsers {
					log.Printf("[INFO] Skipping import of entire user directory ...")
					continue
				}
				if len(g.Members) > 10 {
					log.Printf("[INFO] Importing %d members of %s",
						len(g.Members), g.DisplayName)
				}
				for _, parent := range g.Groups {
					ic.Emit(&resource{
						Resource: "databricks_group",
						ID:       parent.Value,
					})
					if parent.Type == "direct" {
						id := fmt.Sprintf("%s|%s", parent.Value, g.ID)
						ic.Emit(&resource{
							Resource: "databricks_group_member",
							ID:       id,
							Name:     fmt.Sprintf("%s_%s_%s", parent.Display, parent.Value, g.DisplayName),
							Data:     ic.makeGroupMemberData(id, parent.Value, g.ID),
						})
					}
				}
				for i, x := range g.Members {
					if strings.HasPrefix(x.Ref, "Users/") {
						ic.Emit(&resource{
							Resource: "databricks_user",
							ID:       x.Value,
						})
						if !builtInUserGroup {
							id := fmt.Sprintf("%s|%s", g.ID, x.Value)
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       id,
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
								Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
							})
						}
					}
					if strings.HasPrefix(x.Ref, "ServicePrincipals/") {
						ic.Emit(&resource{
							Resource: "databricks_service_principal",
							ID:       x.Value,
						})
						if !builtInUserGroup {
							id := fmt.Sprintf("%s|%s", g.ID, x.Value)
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       id,
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
								Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
							})
						}
					}
					if strings.HasPrefix(x.Ref, "Groups/") {
						ic.Emit(&resource{
							Resource: "databricks_group",
							ID:       x.Value,
						})
						if !builtInUserGroup {
							id := fmt.Sprintf("%s|%s", g.ID, x.Value)
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       id,
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
								Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
							})
						}
					}
					if len(g.Members) > 10 {
						log.Printf("[INFO] Imported %d of %d members of %s", i+1, len(g.Members), g.DisplayName)
					}
				}
			}
			if ic.accountLevel {
				ic.Emit(&resource{
					Resource: "databricks_access_control_rule_set",
					ID: fmt.Sprintf("accounts/%s/groups/%s/ruleSets/default",
						ic.Client.Config.AccountID, r.ID),
				})
			}

			return nil
		},
		Body: resourceOrDataBlockBody,
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
		List: func(ic *importContext) error {
			ic.getUsersMapping()
			ic.allUsersMutex.RLocker().Lock()
			userMapping := maps.Clone(ic.allUsersMapping)
			ic.allUsersMutex.RLocker().Unlock()
			for userName, userScimId := range userMapping {
				log.Printf("[TRACE] Emitting user %s, SCIM id=%s", userName, userScimId)
				ic.Emit(&resource{
					Resource: "databricks_user",
					ID:       userScimId,
				})
			}
			return nil
		},
		Search: func(ic *importContext, r *resource) error {
			u, err := ic.findUserByName(r.Value, false)
			if err != nil {
				return err
			}
			r.ID = u.ID
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			username := r.Data.Get("user_name").(string)
			r.Data.Set("force", true)
			u, err := ic.findUserByName(username, false)
			if err != nil {
				return err
			}
			ic.emitGroups(*u)
			ic.emitRoles("user", u.ID, u.Roles)
			return nil
		},
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
		List: func(ic *importContext) error {
			ic.getSpsMapping()
			ic.spsMutex.RLock()
			spsMapping := maps.Clone(ic.allSpsMapping)
			ic.spsMutex.RLocker().Unlock()
			for applicationId, appScimId := range spsMapping {
				log.Printf("[TRACE] Emitting service principal %s, SCIM id=%s", applicationId, appScimId)
				ic.Emit(&resource{
					Resource: "databricks_service_principal",
					ID:       appScimId,
				})
			}
			return nil
		},
		Search: func(ic *importContext, r *resource) error {
			u, err := ic.findSpnByAppID(r.Value, false)
			if err != nil {
				return err
			}
			r.ID = u.ID
			return nil
		},
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
			if pathString == "application_id" {
				return !ic.Client.IsAzure()
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Import: func(ic *importContext, r *resource) error {
			applicationID := r.Data.Get("application_id").(string)
			r.Data.Set("force", true)
			u, err := ic.findSpnByAppID(applicationID, false)
			if err != nil {
				return err
			}
			ic.emitGroups(*u)
			ic.emitRoles("service_principal", u.ID, u.Roles)
			if ic.accountLevel {
				ic.Emit(&resource{
					Resource: "databricks_access_control_rule_set",
					ID: fmt.Sprintf("accounts/%s/servicePrincipals/%s/ruleSets/default",
						ic.Client.Config.AccountID, applicationID),
				})
			}
			return nil
		},
	},
	"databricks_permissions": {
		Service:        "access",
		WorkspaceLevel: true,
		Name: func(ic *importContext, d *schema.ResourceData) string {
			s := strings.Split(d.Id(), "/")
			return s[len(s)-1]
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
			{Path: "s3_bucket_name", Resource: "aws_s3_bucket", Match: "bucket"}, // this should be changed somehow & avoid clashes with GCS bucket_name
			{Path: "instance_profile", Resource: "databricks_instance_profile"},
			{Path: "cluster_id", Resource: "databricks_cluster"},
			{Path: "storage_account_name", Resource: "azurerm_storage_account", Match: "name"}, // similarly for WASBS vs ABFSS
			{Path: "container_name", Resource: "azurerm_storage_container", Match: "name"},
			{Path: "storage_resource_name", Resource: "azurerm_data_lake_store", Match: "name"},
		},
	},
	"databricks_global_init_script": {
		WorkspaceLevel: true,
		Service:        "workspace",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				return d.Id()
			}
			return name
		},
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
			reposAPI := repos.NewReposAPI(ic.Context, ic.Client)
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			repoDir, err := notebooksAPI.Read(r.Value)
			if err != nil {
				return err
			}
			repo, err := reposAPI.Read(fmt.Sprintf("%d", repoDir.ObjectID))
			if err != nil {
				return err
			}
			r.ID = fmt.Sprintf("%d", repo.ID)
			return nil
		},
		List: func(ic *importContext) error {
			objList, err := repos.NewReposAPI(ic.Context, ic.Client).ListAll()
			if err != nil {
				return err
			}
			for offset, repo := range objList {
				if repo.Url != "" {
					ic.Emit(&resource{
						Resource: "databricks_repo",
						ID:       fmt.Sprintf("%d", repo.ID),
					})
				} else {
					log.Printf("[WARN] ignoring databricks_repo without Git provider. Path: %s", repo.Path)
					ic.addIgnoredResource(fmt.Sprintf("databricks_repo. path=%s", repo.Path))
				}
				log.Printf("[INFO] Scanned %d of %d repos", offset+1, len(objList))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.Data.Get("path").(string), "/Repos")
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
		},
	},
	"databricks_workspace_conf": {
		WorkspaceLevel: true,
		Service:        "workspace",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return globalWorkspaceConfName
		},
		List: func(ic *importContext) error {
			_, err := ic.workspaceClient.WorkspaceConf.GetStatus(ic.Context, settings.GetStatusRequest{
				Keys: "zDummyKey",
			})
			/* this is done to pass the TestImportingNoResourcesError test in exporter_test.go
			Commonly, some of the keys in a workspace conf are Nil
			In the simulated server all are returned with a value.
			We have a zDummyKey - which will always return an error in a real workspace but a value in the simulated workspace
			if no keys have nil values, we should not emit this object
			*/
			if err != nil {
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
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			contentB64, err := notebooksAPI.Export(r.ID, ic.notebooksFormat)
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
			content, _ := base64.StdEncoding.DecodeString(contentB64)
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
	"databricks_workspace_file": {
		WorkspaceLevel: true,
		Service:        "wsfiles",
		Name:           workspaceObjectResouceName,
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			contentB64, err := notebooksAPI.Export(r.ID, "AUTO")
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
			content, _ := base64.StdEncoding.DecodeString(contentB64)
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
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("display_name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
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
		},
		Import: func(ic *importContext, r *resource) error {
			var query tfsql.QueryStruct
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
		},
		// TODO: exclude owner if it's the current user?
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_query", "display_name"),
		Depends: []reference{
			{Path: "warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "parameter.query_backed_value.query_id", Resource: "databricks_query", Match: "id"},
			{Path: "owner_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "owner_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "catalog", Resource: "databricks_catalog"},
			{Path: "schema", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema"),
				SkipDirectLookup:     true},
			// TODO: add match like for workspace files?
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
			// TODO: add support for Repos?
		},
	},
	"databricks_sql_endpoint": {
		WorkspaceLevel: true,
		Service:        "sql-endpoints",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				name = d.Id()
			}
			return name
		},
		List: func(ic *importContext) error {
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
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/warehouses/%s", r.ID),
				"sql_endpoint_"+ic.Importables["databricks_sql_endpoint"].Name(ic, r.Data))
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_sql_global_config",
					ID:       tfsql.GlobalSqlConfigResourceID,
				})
			}
			return nil
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_sql_endpoint", "name"),
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
		Service:        "sql-endpoints",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return "sql_global_config"
		},
		List: func(ic *importContext) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_sql_global_config",
					ID:       tfsql.GlobalSqlConfigResourceID,
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
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
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
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/sql/dashboards/%s", r.ID),
				"sql_dashboard_"+ic.Importables["databricks_sql_dashboard"].Name(ic, r.Data))
			dashboardID := r.ID
			dashboardAPI := tfsql.NewDashboardAPI(ic.Context, ic.Client)
			dashboard, err := dashboardAPI.Read(dashboardID)
			if err != nil {
				return err
			}

			ic.emitSqlParentDirectory(dashboard.Parent)
			for _, rv := range dashboard.Widgets {
				var widget sql_api.Widget
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
					var visualization sql_api.Visualization
					err = json.Unmarshal(widget.Visualization, &visualization)
					if err != nil {
						log.Printf("[WARN] Problems decoding visualization for widget with ID: %s", widget.ID.String())
						continue
					}
					if len(visualization.Query) > 0 {
						var query sql_api.Query
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
		},
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
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string) + "_" + d.Id()
			return name
		},
		Depends: []reference{
			{Path: "query_id", Resource: "databricks_query"},
		},
	},
	"databricks_alert": {
		WorkspaceLevel: true,
		Service:        "alerts",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("display_name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
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
		},
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
			{Path: "owner_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "owner_user_name", Resource: "databricks_service_principal", Match: "application_id"},
			// TODO: add match like for workspace files?
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_directory", Match: "workspace_path"},
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
			var pipeline tfpipelines.Pipeline
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
			var pipeline tfpipelines.Pipeline
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
		Body: resourceOrDataBlockBody,
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
			return pathString == "config.0.auto_capture_config.0.enabled"
		},
		Depends: []reference{
			{Path: "config.served_entities.entity_name", Resource: "databricks_registered_model"},
			{Path: "config.auto_capture_config.catalog_name", Resource: "databricks_catalog"},
			{Path: "config.auto_capture_config.schema_name", Resource: "databricks_schema", Match: "name",
				IsValidApproximation: createIsMatchingCatalogAndSchema("config.0.auto_capture_config.0.catalog_name", "config.0.auto_capture_config.0.schema_name"),
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
		List: func(ic *importContext) error {
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
				if v.State == catalog.SystemSchemaInfoStateEnableCompleted || v.State == catalog.SystemSchemaInfoStateEnableInitialized {
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
		},
	},
	"databricks_artifact_allowlist": {
		WorkspaceLevel: true,
		Service:        "uc-artifact-allowlist",
		List: func(ic *importContext) error {
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
		},
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
		List: func(ic *importContext) error {
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
					log.Printf("[INFO] Skipping catalog %s of type %s", v.Name, v.CatalogType)
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var cat tfcatalog.CatalogInfo
			s := ic.Resources["databricks_catalog"].Schema
			common.DataToStructPointer(r.Data, s, &cat)

			// Emit: UC Connection, List schemas, Catalog grants, ...
			owner, catalogGrantsResource := ic.emitUCGrantsWithOwner("catalog/"+cat.Name, r)
			dependsOn := []*resource{}
			if owner != "" && owner != ic.meUserName {
				dependsOn = append(dependsOn, catalogGrantsResource)
			}
			// TODO: emit owner?  Should we do this? Because it's a account-level identity... Create a separate function for that...
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
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_catalog", "name"),
		Depends: []reference{
			{Path: "connection_name", Resource: "databricks_connection", Match: "name"},
			{Path: "storage_root", Resource: "databricks_external_location", Match: "url", MatchType: MatchLongestPrefix},
		},
		// TODO: convert `main` catalog into the data source as it's automatically created?
		//   This will require addition of the databricks_catalog data source
	},
	"databricks_schema": {
		WorkspaceLevel: true,
		Service:        "uc-schemas",
		Import: func(ic *importContext, r *resource) error {
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
			// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "catalog/" + catalogName})

			// TODO: somehow add depends on catalog's grant...
			// TODO: emit owner? See comment in catalog resource
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
		},
		ShouldOmitField: shouldOmitForUnityCatalog,
		Ignore:          generateIgnoreObjectWithEmptyAttributeValue("databricks_schema", "name"),
		Depends: []reference{
			{Path: "catalog_name", Resource: "databricks_catalog"},
			{Path: "storage_root", Resource: "databricks_external_location", Match: "url", MatchType: MatchLongestPrefix},
		},
	},
	"databricks_volume": {
		WorkspaceLevel: true,
		Service:        "uc-volumes",
		Import: func(ic *importContext, r *resource) error {
			volumeFullName := r.ID
			ic.emitUCGrantsWithOwner("volume/"+volumeFullName, r)

			schemaFullName := r.Data.Get("catalog_name").(string) + "." + r.Data.Get("schema_name").(string)
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       schemaFullName,
			})
			// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "schema/" + schemaFullName})
			// TODO: emit owner? See comment in catalog resource
			return nil
		},
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
		},
	},
	"databricks_sql_table": {
		WorkspaceLevel: true,
		Service:        "uc-tables",
		Import: func(ic *importContext, r *resource) error {
			tableFullName := r.ID
			ic.emitUCGrantsWithOwner("table/"+tableFullName, r)
			schemaFullName := r.Data.Get("catalog_name").(string) + "." + r.Data.Get("schema_name").(string)
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       schemaFullName,
			})
			// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "schema/" + schemaFullName})
			// TODO: emit owner? See comment in catalog resource
			return nil
		},
		Ignore: generateIgnoreObjectWithEmptyAttributeValue("databricks_sql_table", "name"),
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
		},
	},
	"databricks_grants": {
		WorkspaceLevel: true,
		Service:        "uc-grants",
		// TODO: Should we try to make name unique?
		// TODO: do we need to emit principals? Maybe only on account level? See comment for the owner...
		Import: func(ic *importContext, r *resource) error {
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

			var pList tfcatalog.PermissionsList
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
				pList.Assignments = append(pList.Assignments, tfcatalog.PrivilegeAssignment{
					Principal:  ic.meUserName,
					Privileges: newPrivileges,
				})
			}
			return common.StructToData(pList, s, r.Data)
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
			// TODO: add similar matchers for users/groups/SPs on account level...
			{Path: "grant.principal", Resource: "databricks_recipient", IsValidApproximation: isMatchingShareRecipient},
			//	{Path: "", Resource: ""},
			//	{Path: "", Resource: ""},
		},
	},
	"databricks_storage_credential": {
		WorkspaceLevel: true,
		Service:        "uc-storage-credentials",
		Import: func(ic *importContext, r *resource) error {
			ic.emitUCGrantsWithOwner("storage_credential/"+r.ID, r)
			if r.Data != nil {
				isolationMode := r.Data.Get("isolation_mode").(string)
				if isolationMode == "ISOLATION_MODE_ISOLATED" {
					ic.emitWorkspaceBindings("storage_credential", r.ID)
				}
			}
			return nil
		},
		List: func(ic *importContext) error {
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
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "isolation_mode" {
				return d.Get(pathString).(string) != "ISOLATION_MODE_ISOLATED"
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
		},
		Depends: []reference{
			{Path: "azure_service_principal.client_secret", Variable: true},
		},
	},
	"databricks_external_location": {
		WorkspaceLevel: true,
		Service:        "uc-external-locations",
		Import: func(ic *importContext, r *resource) error {
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
		},
		List: func(ic *importContext) error {
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
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
		List: func(ic *importContext) error {
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
		},
		// TODO: think what to do with the sensitive fields in the `options`?
		Import: func(ic *importContext, r *resource) error {
			// TODO: do we need to emit the owner See comment for the owner...
			connectionName := r.Data.Get("name").(string)
			ic.emitUCGrantsWithOwner("foreign_connection/"+connectionName, r)
			return nil
		},
		ShouldOmitField: shouldOmitForUnityCatalog,
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
			// TODO: do we need to emit the owner See comment for the owner...
			var share tfsharing.ShareInfo
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
		// TODO: do we need to emit the owner See comment for the owner...
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
			// r.AddDependsOn(&resource{Resource: "databricks_grants", ID: "schema/" + schemaFullName})
			// TODO: emit owner? See comment in catalog resource
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
		},
	},
	"databricks_metastore": {
		AccountLevel: true,
		Service:      "uc-metastores",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				return d.Id()
			}
			return name
		},
		List: func(ic *importContext) error {
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
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitUCGrantsWithOwner("metastore/"+r.ID, r)
			// TODO: emit owner? See comment in catalog resource
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
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "default_data_access_config_id" || pathString == "storage_root_credential_id" {
				// technically, both should be marked as `computed`
				return true
			}
			return shouldOmitForUnityCatalog(ic, pathString, as, d)
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
		Import: func(ic *importContext, r *resource) error {
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
		},
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
		Service:      "access",
		List: func(ic *importContext) error {
			workspaces, err := ic.accountClient.Workspaces.List(ic.Context)
			if err != nil {
				return err
			}
			for _, ws := range workspaces {
				pas, err := ic.accountClient.WorkspaceAssignment.ListByWorkspaceId(ic.Context, ws.WorkspaceId)
				if err != nil {
					log.Printf("[ERROR] listing workspace permission assignments for workspace %d: %s", ws.WorkspaceId, err.Error())
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
					scm := ic.Resources["databricks_mws_permission_assignment"].Schema
					data.MarkNewResource()
					paId := fmt.Sprintf("%d|%d", ws.WorkspaceId, pa.Principal.PrincipalId)
					data.SetId(paId)
					common.StructToData(pa, scm, data)
					data.Set("workspace_id", ws.WorkspaceId)
					data.Set("principal_id", pa.Principal.PrincipalId)
					ic.Emit(&resource{
						Resource: "databricks_mws_permission_assignment",
						ID:       paId,
						Name:     nameNormalizationRegex.ReplaceAllString(nm, "_"),
						Data:     data,
					})
					// Emit principals
					strPrincipalId := fmt.Sprintf("%d", pa.Principal.PrincipalId)
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
		List: func(ic *importContext) error {
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
		},
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
		Import: func(ic *importContext, r *resource) error {
			path := r.Data.Get("path").(string)
			if strings.HasPrefix(path, "/Repos") {
				ic.emitRepoByPath(path)
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
				ic.Emit(&resource{
					Resource: "databricks_directory",
					ID:       parentPath,
				})
			}
			warehouseId := r.Data.Get("warehouse_id").(string)
			if warehouseId != "" {
				ic.Emit(&resource{
					Resource: "databricks_sql_endpoint",
					ID:       warehouseId,
				})
			}

			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			return pathString == "dashboard_change_detected" || shouldOmitMd5Field(ic, pathString, as, d)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			return strings.HasPrefix(r.Data.Get("path").(string), "/Repos") || strings.HasPrefix(r.Data.Get("parent_path").(string), "/Repos")
		},
		Depends: []reference{
			{Path: "file_path", File: true},
			{Path: "warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "parent_path", Resource: "databricks_directory"},
			{Path: "parent_path", Resource: "databricks_user", Match: "home"},
			{Path: "parent_path", Resource: "databricks_service_principal"},
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
			var notificationDestination tfsettings.NDStruct
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
			var notificationDestination tfsettings.NDStruct
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
			// TODO: emit owner? See comment in catalog resource
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
}
