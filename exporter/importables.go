package exporter

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/secrets"
	tfsql "github.com/databricks/terraform-provider-databricks/sql"
	sql_api "github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
	"golang.org/x/exp/slices"
)

var (
	adlsGen2Regex                = regexp.MustCompile(`^(abfss?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	adlsGen1Regex                = regexp.MustCompile(`^(adls?)://([^.]+)\.(?:[^/]+)(/.*)?$`)
	wasbsRegex                   = regexp.MustCompile(`^(wasbs?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	s3Regex                      = regexp.MustCompile(`^(s3a?)://([^/]+)(/.*)?$`)
	gsRegex                      = regexp.MustCompile(`^gs://([^/]+)(/.*)?$`)
	globalWorkspaceConfName      = "global_workspace_conf"
	nameNormalizationRegex       = regexp.MustCompile(`\W+`)
	fileNameNormalizationRegex   = regexp.MustCompile(`[^-_\w/.@]`)
	jobClustersRegex             = regexp.MustCompile(`^((job_cluster|task)\.\d+\.new_cluster\.\d+\.)`)
	dltClusterRegex              = regexp.MustCompile(`^(cluster\.\d+\.)`)
	userDirRegex                 = regexp.MustCompile(`^(/Users/[^/]+)(/.*)?$`)
	secretPathRegex              = regexp.MustCompile(`^\{\{secrets\/([^\/]+)\/([^}]+)\}\}$`)
	sqlParentRegexp              = regexp.MustCompile(`^folders/(\d+)$`)
	dltDefaultStorageRegex       = regexp.MustCompile(`^dbfs:/pipelines/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	ignoreIdeFolderRegex         = regexp.MustCompile(`^/Users/[^/]+/\.ide/.*$`)
	fileExtensionLanguageMapping = map[string]string{
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
			fileName, err := ic.createFile(name, content)
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
			pools, err := ic.workspaceClient.InstancePools.ListAll(ic.Context)
			if err != nil {
				return err
			}
			for i, pool := range pools {
				if !ic.MatchesName(pool.InstancePoolName) {
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_instance_pool",
					ID:       pool.InstancePoolId,
				})
				log.Printf("[INFO] Imported %d of %d instance pools", i+1, len(pools))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/instance-pools/%s", r.ID),
					Name:     "inst_pool_" + ic.Importables["databricks_instance_pool"].Name(ic, r.Data),
				})
			}
			return nil
		},
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
				return strings.Split(d.Id(), "-")[2]
			}
			return fmt.Sprintf("%s_%s", name, d.Id())
		},
		Depends: []reference{
			{Path: "aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.jar", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.whl", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.egg", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "policy_id", Resource: "databricks_cluster_policy"},
			{Path: "single_user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "single_user_name", Resource: "databricks_service_principal", Match: "application_id"},
		},
		List: func(ic *importContext) error {
			clusters, err := clusters.NewClustersAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			lastActiveMs := ic.getLastActiveMs()
			nonInteractiveClusters := []string{"JOB", "PIPELINE_MAINTENANCE", "PIPELINE", "SQL"}
			for offset, c := range clusters {
				log.Printf("[DEBUG] Cluster %s, source: %s", c.ClusterID, c.ClusterSource)
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
			var c clusters.Cluster
			s := ic.Resources["databricks_cluster"].Schema
			common.DataToStructPointer(r.Data, s, &c)
			ic.importCluster(&c)
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/clusters/%s", r.ID),
					Name:     "cluster_" + ic.Importables["databricks_cluster"].Name(ic, r.Data),
				})
			}
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
			{Path: "email_notifications.on_duration_warning_threshold_exceeded", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_failure", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_start", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "email_notifications.on_success", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "job_cluster.new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "job_cluster.new_cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "job_cluster.new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "job_cluster.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "job_cluster.new_cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "job_cluster.new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "job_cluster.new_cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "run_as.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
			{Path: "run_as.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "task.dbt_task.warehouse_id", Resource: "databricks_sql_endpoint"},
			{Path: "task.existing_cluster_id", Resource: "databricks_cluster"},
			{Path: "task.library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.library.whl", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.library.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "task.new_cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "task.new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "task.new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "task.new_cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.notebook_task.base_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_notebook"},
			{Path: "task.notebook_task.notebook_path", Resource: "databricks_repo", Match: "path", MatchType: MatchPrefix},
			{Path: "task.pipeline_task.pipeline_id", Resource: "databricks_pipeline"},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.python_wheel_task.named_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.python_wheel_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.run_job_task.job_id", Resource: "databricks_job"},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.run_job_task.job_parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.spark_jar_task.jar_uri", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.spark_jar_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.spark_python_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_repo", Match: "path", MatchType: MatchPrefix},
			{Path: "task.spark_python_task.python_file", Resource: "databricks_workspace_file", Match: "path"},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "task.spark_submit_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "task.sql_task.alert.alert_id", Resource: "databricks_sql_alert"},
			{Path: "task.sql_task.dashboard.dashboard_id", Resource: "databricks_sql_dashboard"},
			{Path: "task.sql_task.query.query_id", Resource: "databricks_sql_query"},
			{Path: "task.sql_task.warehouse_id", Resource: "databricks_sql_endpoint"},
		},
		Import: func(ic *importContext, r *resource) error {
			var job jobs.JobSettings
			s := ic.Resources["databricks_job"].Schema
			common.DataToStructPointer(r.Data, s, &job)
			ic.importCluster(job.NewCluster)
			ic.Emit(&resource{
				Resource: "databricks_cluster",
				ID:       job.ExistingClusterID,
			})
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/jobs/%s", r.ID),
					Name:     "job_" + ic.Importables["databricks_job"].Name(ic, r.Data),
				})
			}
			// Support for multitask jobs
			for _, task := range job.Tasks {
				if task.NotebookTask != nil {
					if task.NotebookTask.Source != "GIT" {
						ic.emitNotebookOrRepo(task.NotebookTask.NotebookPath)
					}
					ic.emitFilesFromMap(task.NotebookTask.BaseParameters)
				}
				if task.PipelineTask != nil {
					ic.Emit(&resource{
						Resource: "databricks_pipeline",
						ID:       task.PipelineTask.PipelineID,
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
							Resource: "databricks_sql_query",
							ID:       task.SqlTask.Query.QueryID,
						})
					}
					if task.SqlTask.Dashboard != nil {
						ic.Emit(&resource{
							Resource: "databricks_sql_dashboard",
							ID:       task.SqlTask.Dashboard.DashboardID,
						})
					}
					if task.SqlTask.Alert != nil {
						ic.Emit(&resource{
							Resource: "databricks_sql_alert",
							ID:       task.SqlTask.Alert.AlertID,
						})
					}
					if task.SqlTask.WarehouseID != "" {
						ic.Emit(&resource{
							Resource: "databricks_sql_endpoint",
							ID:       task.SqlTask.WarehouseID,
						})
					}
				}
				if task.DbtTask != nil {
					if task.DbtTask.WarehouseId != "" {
						ic.Emit(&resource{
							Resource: "databricks_sql_endpoint",
							ID:       task.DbtTask.WarehouseId,
						})
					}
				}
				if task.RunJobTask != nil && task.RunJobTask.JobID != 0 {
					ic.Emit(&resource{
						Resource: "databricks_job",
						ID:       strconv.FormatInt(task.RunJobTask.JobID, 10),
					})
					ic.emitFilesFromMap(task.RunJobTask.JobParameters)
				}
				ic.importCluster(task.NewCluster)
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       task.ExistingClusterID,
				})
				ic.emitLibraries(task.Libraries)
			}
			for _, jc := range job.JobClusters {
				ic.importCluster(jc.NewCluster)
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
			var js jobs.JobSettings
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
			policiesList, err := w.ClusterPolicies.ListAll(ic.Context, compute.ListClusterPoliciesRequest{})
			if err != nil {
				return err
			}

			builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
			for offset, policy := range policiesList {
				log.Printf("[TRACE] Scanning %d:  %v", offset+1, policy)
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
				if offset%10 == 0 {
					log.Printf("[INFO] Scanned %d of %d cluster policies", offset+1, len(policiesList))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/cluster-policies/%s", r.ID),
					Name:     "cluster_policy_" + ic.Importables["databricks_cluster_policy"].Name(ic, r.Data),
				})
			}

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
			}

			policyFamilyId := clusterPolicy.PolicyFamilyId
			if policyFamilyId != "" && clusterPolicy.Definition != "" {
				// we need to set definition to empty value because otherwise it will be put into
				// generated HCL code for data source, and it only supports the `name` attribute
				r.Data.Set("definition", "")
				builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
				_, isBuiltin := builtInClusterPolicies[policyFamilyId]
				if isBuiltin && clusterPolicy.PolicyFamilyDefinitionOverrides == "" {
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
			{Path: "libraries.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.jar", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "libraries.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "libraries.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.whl", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "libraries.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "libraries.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "libraries.egg", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
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
						ic.Emit(&resource{
							Resource: "databricks_group_member",
							ID:       fmt.Sprintf("%s|%s", parent.Value, g.ID),
							Name:     fmt.Sprintf("%s_%s_%s", parent.Display, parent.Value, g.DisplayName),
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
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       fmt.Sprintf("%s|%s", g.ID, x.Value),
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
							})
						}
					}
					if strings.HasPrefix(x.Ref, "ServicePrincipals/") {
						ic.Emit(&resource{
							Resource: "databricks_service_principal",
							ID:       x.Value,
						})
						if !builtInUserGroup {
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       fmt.Sprintf("%s|%s", g.ID, x.Value),
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
							})
						}
					}
					if strings.HasPrefix(x.Ref, "Groups/") {
						ic.Emit(&resource{
							Resource: "databricks_group",
							ID:       x.Value,
						})
						if !builtInUserGroup {
							ic.Emit(&resource{
								Resource: "databricks_group_member",
								ID:       fmt.Sprintf("%s|%s", g.ID, x.Value),
								Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
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
		// TODO: we need to add List operation here as well
		Search: func(ic *importContext, r *resource) error {
			u, err := ic.findUserByName(r.Value)
			if err != nil {
				return err
			}
			r.ID = u.ID
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			username := r.Data.Get("user_name").(string)
			r.Data.Set("force", true)
			u, err := ic.findUserByName(username)
			if err != nil {
				return err
			}
			ic.emitGroups(u)
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
		// TODO: we need to add List operation here as well
		Search: func(ic *importContext, r *resource) error {
			u, err := ic.findSpnByAppID(r.Value)
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
			u, err := ic.findSpnByAppID(applicationID)
			if err != nil {
				return err
			}
			ic.emitGroups(u)
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
			{Path: "sql_query_id", Resource: "databricks_sql_query"},
			{Path: "sql_alert_id", Resource: "databricks_sql_alert"},
			{Path: "sql_dashboard_id", Resource: "databricks_sql_dashboard"},
			{Path: "sql_endpoint_id", Resource: "databricks_sql_endpoint"},
			{Path: "registered_model_id", Resource: "databricks_mlflow_model"},
			{Path: "experiment_id", Resource: "databricks_mlflow_experiment"},
			{Path: "repo_id", Resource: "databricks_repo"},
			{Path: "directory_id", Resource: "databricks_directory", Match: "object_id"},
			{Path: "notebook_id", Resource: "databricks_notebook", Match: "object_id"},
			{Path: "access_control.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "access_control.group_name", Resource: "databricks_group", Match: "display_name"},
			{Path: "access_control.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
		},
		Ignore: func(ic *importContext, r *resource) bool {
			var permissions permissions.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			common.DataToStructPointer(r.Data, s, &permissions)
			return (len(permissions.AccessControlList) == 0)
		},
		Import: func(ic *importContext, r *resource) error {
			var permissions permissions.PermissionsEntity
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
			ssAPI := secrets.NewSecretScopesAPI(ic.Context, ic.Client)
			if scopes, err := ssAPI.List(); err == nil {
				for i, scope := range scopes {
					if !ic.MatchesName(scope.Name) {
						log.Printf("[INFO] Secret scope %s doesn't match %s filter",
							scope.Name, ic.match)
						continue
					}
					ic.Emit(&resource{
						Resource: "databricks_secret_scope",
						ID:       scope.Name,
					})
					log.Printf("[INFO] Imported %d of %d secret scopes", i+1, len(scopes))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			backendType, _ := r.Data.GetOk("backend_type")
			if backendType != "AZURE_KEYVAULT" {
				if l, err := secrets.NewSecretsAPI(ic.Context, ic.Client).List(r.ID); err == nil {
					for _, secret := range l {
						ic.Emit(&resource{
							Resource: "databricks_secret",
							ID:       fmt.Sprintf("%s|||%s", r.ID, secret.Key),
						})
					}
				}
			}
			if l, err := secrets.NewSecretAclsAPI(ic.Context, ic.Client).List(r.ID); err == nil {
				for _, acl := range l {
					ic.Emit(&resource{
						Resource: "databricks_secret_acl",
						ID:       fmt.Sprintf("%s|||%s", r.ID, acl.Principal),
					})
				}
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
		Depends: []reference{
			{Path: "string_value", Variable: true},
			{Path: "scope", Resource: "databricks_secret_scope"},
			{Path: "string_value", Resource: "vault_generic_secret", Match: "data"},
			{Path: "string_value", Resource: "aws_kms_secrets", Match: "plaintext"},
			{Path: "string_value", Resource: "azurerm_key_vault_secret", Match: "value"},
			{Path: "string_value", Resource: "aws_secretsmanager_secret_version", Match: "secret_string"},
		},
		Name: func(ic *importContext, d *schema.ResourceData) string {
			name := fmt.Sprintf("%s_%s", d.Get("scope"), d.Get("key"))
			return name + "_" + generateUniqueID(name)
		},
	},
	"databricks_secret_acl": {
		WorkspaceLevel: true,
		Service:        "secrets",
		Depends: []reference{
			{Path: "scope", Resource: "databricks_secret_scope"},
			{Path: "principal", Resource: "databricks_group", Match: "display_name"},
			{Path: "principal", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
			{Path: "principal", Resource: "databricks_service_principal", Match: "application_id"},
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
			globalInitScripts, err := workspace.NewGlobalInitScriptsAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			updatedSinceMs := ic.getUpdatedSinceMs()
			for offset, gis := range globalInitScripts {
				modifiedAt := gis.UpdatedAt
				if ic.incremental && modifiedAt < updatedSinceMs {
					log.Printf("[DEBUG] skipping global init script '%s' that was modified at %d (last active=%d)",
						gis.Name, modifiedAt, updatedSinceMs)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_global_init_script",
					ID:       gis.ScriptID,
				})
				log.Printf("[INFO] Scanned %d of %d global init scripts", offset+1, len(globalInitScripts))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			gis, err := workspace.NewGlobalInitScriptsAPI(ic.Context, ic.Client).Get(r.ID)
			if err != nil {
				return err
			}
			content, err := base64.StdEncoding.DecodeString(gis.ContentBase64)
			if err != nil {
				return err
			}
			fileName, err := ic.createFile(fmt.Sprintf("%s.sh", r.Name), content)
			log.Printf("Creating %s for %s", fileName, r)
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
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/repos/%s", r.ID),
					Name:     "repo_" + ic.Importables["databricks_repo"].Name(ic, r.Data),
				})
			}
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
			{Path: "path", Resource: "databricks_user", Match: "repos", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_service_principal", Match: "repos", MatchType: MatchPrefix},
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
			loaded := map[string]any{}
			keyNames := []string{}
			for k := range ic.workspaceConfKeys {
				keyNames = append(keyNames, k)
			}
			sort.Strings(keyNames)
			conf, err := ic.workspaceClient.WorkspaceConf.GetStatus(ic.Context, settings.GetStatusRequest{
				Keys: strings.Join(keyNames, ","),
			})
			if err != nil {
				return err
			}
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
			ipListsResp, err := ic.workspaceClient.IpAccessLists.Impl().List(ic.Context)

			if err != nil {
				return err
			}
			ipLists := ipListsResp.IpAccessLists
			updatedSinceMs := ic.getUpdatedSinceMs()
			for offset, ipList := range ipLists {
				modifiedAt := ipList.UpdatedAt
				if ic.incremental && modifiedAt < updatedSinceMs {
					log.Printf("[DEBUG] skipping IP access list '%s' that was modified at %d (last active=%d)",
						ipList.Label, modifiedAt, updatedSinceMs)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_ip_access_list",
					ID:       ipList.ListId,
				})
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
		List:           createListWorkspaceObjectsFunc(workspace.Notebook, "databricks_notebook", "notebook"),
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			contentB64, err := notebooksAPI.Export(r.ID, ic.notebooksFormat)
			if err != nil {
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
			fileName, err := ic.createFileIn("notebooks", name, []byte(content))
			if err != nil {
				return err
			}
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/notebooks/%d", objectId),
					Name:     "notebook_" + ic.Importables["databricks_notebook"].Name(ic, r.Data),
				})
			}

			// TODO: it's not completely correct condition - we need to make emit smarter -
			// emit only if permissions are different from their parent's permission.
			if ic.meAdmin {
				directorySplits := strings.Split(r.ID, "/")
				directorySplits = directorySplits[:len(directorySplits)-1]
				directoryPath := strings.Join(directorySplits, "/")

				ic.Emit(&resource{
					Resource: "databricks_directory",
					ID:       directoryPath,
				})
			}

			return r.Data.Set("source", fileName)
		},
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_directory", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_user", Match: "home", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_service_principal", Match: "home", MatchType: MatchPrefix},
		},
	},
	"databricks_workspace_file": {
		WorkspaceLevel: true,
		Service:        "notebooks",
		Name:           workspaceObjectResouceName,
		List:           createListWorkspaceObjectsFunc(workspace.File, "databricks_workspace_file", "workspace_file"),
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			contentB64, err := notebooksAPI.Export(r.ID, "AUTO")
			if err != nil {
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
			fileName, err := ic.createFileIn("workspace_files", name, []byte(content))
			if err != nil {
				return err
			}

			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/files/%d", objectId),
					Name:     "ws_file_" + ic.Importables["databricks_workspace_file"].Name(ic, r.Data),
				})
			}

			// TODO: it's not completely correct condition - we need to make emit smarter -
			// emit only if permissions are different from their parent's permission.
			if ic.meAdmin {
				directorySplits := strings.Split(r.ID, "/")
				directorySplits = directorySplits[:len(directorySplits)-1]
				directoryPath := strings.Join(directorySplits, "/")

				ic.Emit(&resource{
					Resource: "databricks_directory",
					ID:       directoryPath,
				})
			}
			log.Printf("Creating %s for %s", fileName, r)
			return r.Data.Set("source", fileName)
		},
		ShouldOmitField: shouldOmitMd5Field,
		Depends: []reference{
			{Path: "source", File: true},
			{Path: "path", Resource: "databricks_directory", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_user", Match: "home", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_service_principal", Match: "home", MatchType: MatchPrefix},
		},
	},
	"databricks_sql_query": {
		WorkspaceLevel: true,
		Service:        "sql-queries",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
			qs, err := dbsqlListObjects(ic, "/preview/sql/queries")
			if err != nil {
				return nil
			}
			updatedSinceStr := ic.getUpdatedSinceStr()
			for i, q := range qs {
				name := q["name"].(string)
				if !ic.MatchesName(name) {
					continue
				}
				updatedAt := q["updated_at"].(string)
				if ic.incremental && updatedAt < updatedSinceStr {
					log.Printf("[DEBUG] skipping query '%s' that was modified at %s (updatedSince=%s)", name,
						updatedAt, updatedSinceStr)
					continue
				}
				log.Printf("[DEBUG] emitting query '%s' that was modified at %s (updatedSince=%s)", name,
					updatedAt, updatedSinceStr)
				ic.Emit(&resource{
					Resource:    "databricks_sql_query",
					ID:          q["id"].(string),
					Incremental: ic.incremental,
				})
				log.Printf("[INFO] Imported %d of %d SQL queries", i+1, len(qs))
			}

			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var query tfsql.QueryEntity
			s := ic.Resources["databricks_sql_query"].Schema
			common.DataToStructPointer(r.Data, s, &query)
			sqlEndpointID, err := ic.getSqlEndpoint(query.DataSourceID)
			if err == nil {
				ic.Emit(&resource{
					Resource: "databricks_sql_endpoint",
					ID:       sqlEndpointID,
				})
			} else {
				log.Printf("[WARN] Can't find SQL endpoint for data source '%s'", query.DataSourceID)
			}
			// emit queries specified as parameters
			for _, p := range query.Parameter {
				if p.Query != nil {
					ic.Emit(&resource{
						Resource: "databricks_sql_query",
						ID:       p.Query.QueryID,
					})
				}
			}
			ic.emitSqlParentDirectory(query.Parent)
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/sql/queries/%s", r.ID),
					Name:     "sql_query_" + ic.Importables["databricks_sql_query"].Name(ic, r.Data),
				})
			}
			return nil
		},
		Depends: []reference{
			{Path: "data_source_id", Resource: "databricks_sql_endpoint", Match: "data_source_id"},
			{Path: "parameter.query.query_id", Resource: "databricks_sql_query", Match: "id"},
			{Path: "parent", Resource: "databricks_directory", Match: "object_id", MatchType: MatchRegexp,
				Regexp: sqlParentRegexp},
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
			endpointsList, err := ic.workspaceClient.Warehouses.ListAll(ic.Context, sql.ListWarehousesRequest{})
			if err != nil {
				return err
			}
			for i, q := range endpointsList {
				if !ic.MatchesName(q.Name) {
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_sql_endpoint",
					ID:       q.Id,
				})
				log.Printf("[INFO] Imported %d of %d SQL endpoints", i+1, len(endpointsList))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/sql/warehouses/%s", r.ID),
					Name:     "sql_endpoint_" + ic.Importables["databricks_sql_endpoint"].Name(ic, r.Data),
				})
				ic.Emit(&resource{
					Resource: "databricks_sql_global_config",
					ID:       tfsql.GlobalSqlConfigResourceID,
				})
			}
			return nil
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
			updatedSinceStr := ic.getUpdatedSinceStr()
			for i, q := range qs {
				name := q["name"].(string)
				if !ic.MatchesName(name) {
					continue
				}
				updatedAt := q["updated_at"].(string)
				if ic.incremental && updatedAt < updatedSinceStr {
					log.Printf("[DEBUG] skipping dashboard '%s' that was modified at %s (updatedSince=%s)", name,
						updatedAt, updatedSinceStr)
					continue
				}
				log.Printf("[DEBUG] emitting dashboard '%s' that was modified at %s (updatedSince=%s)", name,
					updatedAt, updatedSinceStr)
				ic.Emit(&resource{
					Resource:    "databricks_sql_dashboard",
					ID:          q["id"].(string),
					Incremental: ic.incremental,
				})
				log.Printf("[INFO] Imported %d of %d SQL dashboards", i+1, len(qs))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/sql/dashboards/%s", r.ID),
					Name:     "sql_dashboard_" + ic.Importables["databricks_sql_dashboard"].Name(ic, r.Data),
				})
			}
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
							Resource: "databricks_sql_query",
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
			{Path: "dashboard_id", Resource: "databricks_sql_dashboard", Match: "id"},
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
			{Path: "query_id", Resource: "databricks_sql_query", Match: "id"},
		},
	},
	"databricks_sql_alert": {
		WorkspaceLevel: true,
		Service:        "sql-alerts",
		Name: func(ic *importContext, d *schema.ResourceData) string {
			return d.Get("name").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
			updatedSinceStr := ic.getUpdatedSinceStr()
			alerts, err := ic.workspaceClient.Alerts.List(ic.Context)
			if err != nil {
				return err
			}
			for i, alert := range alerts {
				name := alert.Name
				if !ic.MatchesName(name) {
					continue
				}
				if ic.incremental && alert.UpdatedAt < updatedSinceStr {
					log.Printf("[DEBUG] skipping alert '%s' that was modified at %s (last active=%s)", name,
						alert.UpdatedAt, updatedSinceStr)
					continue
				}
				log.Printf("[DEBUG] emitting alert '%s' that was modified at %s (last active=%s)", name,
					alert.UpdatedAt, updatedSinceStr)
				ic.Emit(&resource{
					Resource:    "databricks_sql_alert",
					ID:          alert.Id,
					Incremental: ic.incremental,
				})
				log.Printf("[INFO] Imported %d of %d SQL alerts", i+1, len(alerts))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var alert tfsql.AlertEntity
			s := ic.Resources["databricks_sql_alert"].Schema
			common.DataToStructPointer(r.Data, s, &alert)
			if alert.QueryId != "" {
				ic.Emit(&resource{Resource: "databricks_sql_query", ID: alert.QueryId})
			}
			ic.emitSqlParentDirectory(alert.Parent)
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/sql/alerts/%s", r.ID),
					Name:     "sql_alert_" + ic.Importables["databricks_sql_alert"].Name(ic, r.Data)})
			}
			return nil
		},
		Depends: []reference{
			{Path: "query_id", Resource: "databricks_sql_query", Match: "id"},
			{Path: "parent", Resource: "databricks_directory", Match: "object_id",
				MatchType: MatchRegexp, Regexp: sqlParentRegexp},
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
			api := pipelines.NewPipelinesAPI(ic.Context, ic.Client)
			pipelinesList, err := api.List(50, "")
			if err != nil {
				return err
			}
			updatedSinceMs := ic.getUpdatedSinceMs()
			for i, q := range pipelinesList {
				if !ic.MatchesName(q.Name) {
					continue
				}
				if ic.incremental {
					pipeline, err := api.Read(q.PipelineID)
					if err != nil {
						return err
					}
					modifiedAt := pipeline.LastModified
					if modifiedAt < updatedSinceMs {
						log.Printf("[DEBUG] skipping DLT Pipeline '%s' that was modified at %d (last active=%d)",
							pipeline.Name, modifiedAt, updatedSinceMs)
						continue
					}
				}
				ic.Emit(&resource{
					Resource: "databricks_pipeline",
					ID:       q.PipelineID,
				})
				log.Printf("[INFO] Imported %d of %d DLT Pipelines", i+1, len(pipelinesList))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var pipeline pipelines.PipelineSpec
			s := ic.Resources["databricks_pipeline"].Schema
			common.DataToStructPointer(r.Data, s, &pipeline)
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
			// Emit clusters
			for _, cluster := range pipeline.Clusters {
				if cluster.AwsAttributes != nil && cluster.AwsAttributes.InstanceProfileArn != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       cluster.AwsAttributes.InstanceProfileArn,
					})
				}
				if cluster.InstancePoolID != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       cluster.InstancePoolID,
					})
				}
				if cluster.DriverInstancePoolID != "" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       cluster.DriverInstancePoolID,
					})
				}
				if cluster.PolicyID != "" {
					ic.Emit(&resource{
						Resource: "databricks_cluster_policy",
						ID:       cluster.PolicyID,
					})
				}
				ic.emitInitScripts(cluster.InitScripts)
				ic.emitSecretsFromSecretsPath(cluster.SparkConf)
				ic.emitSecretsFromSecretsPath(cluster.SparkEnvVars)
			}
			ic.emitFilesFromMap(pipeline.Configuration)
			ic.emitSecretsFromSecretsPath(pipeline.Configuration)

			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/pipelines/%s", r.ID),
					Name:     "pipeline_" + ic.Importables["databricks_pipeline"].Name(ic, r.Data),
				})
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if res := dltClusterRegex.FindStringSubmatch(pathString); res != nil { // analyze DLT clusters
				return makeShouldOmitFieldForCluster(dltClusterRegex)(ic, pathString, as, d)
			}
			if pathString == "storage" {
				return dltDefaultStorageRegex.FindStringSubmatch(d.Get("storage").(string)) != nil
			}
			return pathString == "creator_user_name" || defaultShouldOmitFieldFunc(ic, pathString, as, d)
		},
		Ignore: func(ic *importContext, r *resource) bool {
			numLibraries := r.Data.Get("library.#").(int)
			if numLibraries == 0 {
				log.Printf("[WARN] Ignoring DLT Pipeline with ID %s", r.ID)
				ic.addIgnoredResource(fmt.Sprintf("databricks_pipeline. id=%s", r.ID))
			}
			return numLibraries == 0
		},
		Depends: []reference{
			{Path: "cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
			{Path: "cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster.policy_id", Resource: "databricks_cluster_policy"},
			{Path: "configuration", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "configuration", Resource: "databricks_repo", Match: "workspace_path", MatchType: MatchPrefix},
			{Path: "configuration", Resource: "databricks_workspace_file", Match: "workspace_path"},
			{Path: "library.notebook.path", Resource: "databricks_notebook"},
			{Path: "library.notebook.path", Resource: "databricks_repo", Match: "path", MatchType: MatchPrefix},
			{Path: "library.file.path", Resource: "databricks_workspace_file"},
			{Path: "library.file.path", Resource: "databricks_repo", Match: "path", MatchType: MatchPrefix},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
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
		// TODO: think if we really need this, we need directories only for permissions,
		// and only when they are different from parents & notebooks
		List: func(ic *importContext) error {
			directoryList := ic.getAllDirectories()
			for offset, directory := range directoryList {
				if strings.HasPrefix(directory.Path, "/Repos") {
					continue
				}
				if res := ignoreIdeFolderRegex.FindStringSubmatch(directory.Path); res != nil {
					continue
				}
				ic.maybeEmitWorkspaceObject("databricks_directory", directory.Path)

				if offset%50 == 0 {
					log.Printf("[INFO] Scanned %d of %d directories", offset+1, len(directoryList))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
			// Existing permissions API doesn't allow to set permissions for
			if ic.meAdmin && r.ID != "/Shared" {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/directories/%d", r.Data.Get("object_id").(int)),
					Name:     "directory_" + ic.Importables["databricks_directory"].Name(ic, r.Data),
				})
			}

			if r.ID == "/Shared" || r.ID == "/Users" || ic.IsUserOrServicePrincipalDirectory(r.ID, "/Users") {
				r.Mode = "data"
			}

			return nil

		},
		Body: resourceOrDataBlockBody,
		Depends: []reference{
			// TODO: it should try to find longest reference to another directory object that it not itself...
			{Path: "path", Resource: "databricks_user", Match: "home", MatchType: MatchPrefix},
			{Path: "path", Resource: "databricks_service_principal", Match: "home", MatchType: MatchPrefix},
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
			endpointsList, err := ic.workspaceClient.ServingEndpoints.ListAll(ic.Context)
			if err != nil {
				return err
			}

			updatedSinceMs := ic.getUpdatedSinceMs()
			for offset, endpoint := range endpointsList {
				modifiedAt := endpoint.LastUpdatedTimestamp
				if ic.incremental && modifiedAt < updatedSinceMs {
					log.Printf("[DEBUG] skipping serving endpoint '%s' that was modified at %d (last active=%d)",
						endpoint.Name, modifiedAt, updatedSinceMs)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_model_serving",
					ID:       endpoint.Name,
				})
				if offset%50 == 0 {
					log.Printf("[INFO] Scanned %d of %d Serving Endpoints", offset+1, len(endpointsList))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				log.Printf("[DEBUG] Emitting permissions of endpoint '%s' id='%s'", r.ID, r.Data.Get("serving_endpoint_id").(string))
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/serving-endpoints/%s", r.Data.Get("serving_endpoint_id").(string)),
					Name:     "serving_endpoint_" + ic.Importables["databricks_model_serving"].Name(ic, r.Data),
				})
			}
			return nil
		},
		ShouldOmitField: func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
			if pathString == "config.0.traffic_config" ||
				(strings.HasPrefix(pathString, "config.0.served_models.") &&
					strings.HasSuffix(pathString, ".scale_to_zero_enabled")) {
				return false
			}
			return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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

			updatedSinceMs := ic.getUpdatedSinceMs()
			for offset, webhook := range webhooks {
				modifiedAt := webhook.LastUpdatedTimestamp
				if ic.incremental && modifiedAt < updatedSinceMs {
					log.Printf("[DEBUG] skipping MLflow webhook '%s' that was modified at %d (last active=%d)",
						webhook.Id, modifiedAt, updatedSinceMs)
					continue
				}
				log.Printf("[DEBUG] emitting MLflow webhook '%s' that was modified at %d (last active=%d)",
					webhook.Id, modifiedAt, updatedSinceMs)
				ic.Emit(&resource{
					Resource: "databricks_mlflow_webhook",
					ID:       webhook.Id,
				})
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
		// TODO: add Depends & Import to emit corresponding UC Volumes when support for them is added
	},
}
