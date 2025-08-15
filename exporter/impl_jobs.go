package exporter

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
	tf_jobs "github.com/databricks/terraform-provider-databricks/jobs"
	tf_workspace "github.com/databricks/terraform-provider-databricks/workspace"
)

func importTask(ic *importContext, task sdk_jobs.Task, jobName, rID string) {
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
			for _, subscription := range task.SqlTask.Alert.Subscriptions {
				if subscription.UserName != "" {
					ic.Emit(&resource{
						Resource:  "databricks_user",
						Attribute: "user_name",
						Value:     subscription.UserName,
					})
				}
				if subscription.DestinationId != "" {
					ic.Emit(&resource{
						Resource: "databricks_notification_destination",
						ID:       subscription.DestinationId,
					})
				}
			}
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
	if task.DashboardTask != nil {
		ic.Emit(&resource{
			Resource: "databricks_dashboard",
			ID:       task.DashboardTask.DashboardId,
		})
		if task.DashboardTask.WarehouseId != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_endpoint",
				ID:       task.DashboardTask.WarehouseId,
			})
		}
		if task.DashboardTask.Subscription != nil {
			for _, subscriber := range task.DashboardTask.Subscription.Subscribers {
				if subscriber.DestinationId != "" {
					ic.Emit(&resource{
						Resource: "databricks_notification_destination",
						ID:       subscriber.DestinationId,
					})
				}
				if subscriber.UserName != "" {
					ic.Emit(&resource{
						Resource:  "databricks_user",
						Attribute: "user_name",
						Value:     subscriber.UserName,
					})
				}
			}
		}
	}
	if task.PowerBiTask != nil {
		if task.PowerBiTask.WarehouseId != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_endpoint",
				ID:       task.PowerBiTask.WarehouseId,
			})
		}
		if task.PowerBiTask.ConnectionResourceName != "" && ic.currentMetastore != nil {
			ic.Emit(&resource{
				Resource: "databricks_connection",
				ID:       ic.currentMetastore.MetastoreId + "|" + task.PowerBiTask.ConnectionResourceName,
			})
		}
		for _, table := range task.PowerBiTask.Tables {
			if table.Catalog != "" && table.Schema != "" && table.Name != "" {
				ic.Emit(&resource{
					Resource: "databricks_sql_table",
					ID:       table.Catalog + "." + table.Schema + "." + table.Name,
				})
			}
		}

	}
	if task.DbtTask != nil {
		if task.DbtTask.WarehouseId != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_endpoint",
				ID:       task.DbtTask.WarehouseId,
			})
		}
		if task.DbtTask.Catalog != "" && task.DbtTask.Schema != "" {
			ic.Emit(&resource{
				Resource: "databricks_catalog",
				ID:       task.DbtTask.Catalog,
			})
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       task.DbtTask.Catalog + "." + task.DbtTask.Schema,
			})
		}
		if task.DbtTask.Source == "WORKSPACE" {
			directory := task.DbtTask.ProjectDirectory
			if ic.isInRepoOrGitFolder(directory, true) {
				ic.emitRepoOrGitFolder(directory, true)
			} else {
				// Traverse the dbt project directory and emit all objects found in it
				nbAPI := tf_workspace.NewNotebooksAPI(ic.Context, ic.Client)
				objects, err := nbAPI.List(directory, true, true)
				if err == nil {
					for _, object := range objects {
						if object.ObjectType != tf_workspace.File {
							continue
						}
						ic.maybeEmitWorkspaceObject("databricks_workspace_file", object.Path, &object)
					}
				} else {
					log.Printf("[WARN] Can't list directory %s for DBT task in job %s (id: %s)", directory, jobName, rID)
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
	if task.ForEachTask != nil {
		importTask(ic, task.ForEachTask.Task, jobName, rID)
	}
	ic.importCluster(task.NewCluster)
	if task.ExistingClusterId != "" {
		ic.Emit(&resource{
			Resource: "databricks_cluster",
			ID:       task.ExistingClusterId,
		})
	}
	ic.emitLibraries(task.Libraries)

	emitWebhookNotifications(ic, task.WebhookNotifications)
	if task.EmailNotifications != nil {
		ic.emitListOfUsers(task.EmailNotifications.OnDurationWarningThresholdExceeded)
		ic.emitListOfUsers(task.EmailNotifications.OnFailure)
		ic.emitListOfUsers(task.EmailNotifications.OnStart)
		ic.emitListOfUsers(task.EmailNotifications.OnSuccess)
		ic.emitListOfUsers(task.EmailNotifications.OnStreamingBacklogExceeded)
	}
}

func importJob(ic *importContext, r *resource) error {
	var job tf_jobs.JobSettingsResource
	s := ic.Resources["databricks_job"].Schema
	common.DataToStructPointer(r.Data, s, &job)
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/jobs/%s", r.ID),
		"job_"+ic.Importables["databricks_job"].Name(ic, r.Data))
	for _, task := range job.Tasks {
		importTask(ic, task, job.Name, r.ID)
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
	if job.Trigger != nil && job.Trigger.TableUpdate != nil {
		for _, table := range job.Trigger.TableUpdate.TableNames {
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       table,
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
	emitWebhookNotifications(ic, job.WebhookNotifications)
	for _, param := range job.Parameters {
		ic.emitIfWsfsFile(param.Default)
		ic.emitIfVolumeFile(param.Default)
	}
	for _, env := range job.Environments {
		if env.Spec != nil {
			for _, dep := range env.Spec.Dependencies {
				emitEnvironmentDependency(ic, dep)
			}
		}
	}

	return ic.importLibraries(r.Data, s)
}

func emitEnvironmentDependency(ic *importContext, dep string) {
	v := dep
	if res := requirementsFileRegexp.FindStringSubmatch(v); res != nil {
		v = res[1]
	}
	ic.emitIfWsfsFile(v)
	ic.emitIfVolumeFile(v)
}

func emitWebhookNotifications(ic *importContext, notifications *sdk_jobs.WebhookNotifications) {
	if notifications != nil {
		ic.emitJobsDestinationNotifications(notifications.OnFailure)
		ic.emitJobsDestinationNotifications(notifications.OnSuccess)
		ic.emitJobsDestinationNotifications(notifications.OnDurationWarningThresholdExceeded)
		ic.emitJobsDestinationNotifications(notifications.OnStart)
		ic.emitJobsDestinationNotifications(notifications.OnStreamingBacklogExceeded)
	}
}

func listJobs(ic *importContext) error {
	i := 0
	it := ic.workspaceClient.Jobs.List(ic.Context, sdk_jobs.ListJobsRequest{ExpandTasks: false, Limit: 100})
	for it.HasNext(ic.Context) {
		job, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if i%50 == 0 {
			log.Printf("[INFO] Scanned %d jobs", i)
		}
		if !ic.MatchesName(job.Settings.Name) {
			log.Printf("[INFO] Job name %s doesn't match selection %s", job.Settings.Name, ic.match)
			continue
		}
		if job.Settings.Deployment != nil && job.Settings.Deployment.Kind == "BUNDLE" &&
			job.Settings.EditMode == "UI_LOCKED" {
			log.Printf("[INFO] Skipping job '%s' because it's deployed by DABs", job.Settings.Name)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_job",
			ID:       strconv.FormatInt(job.JobId, 10),
		})
	}
	log.Printf("[INFO] Total %d jobs are going to be exported", i)
	return nil
}

func shouldOmitFieldInJob(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, r *resource) bool {
	switch pathString {
	case "url", "format":
		return true
	}
	var js tf_jobs.JobSettingsResource
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
		return makeShouldOmitFieldForCluster(jobClustersRegex)(ic, pathString, as, d, r)
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d, r)
}

func shouldIgnoreJob(ic *importContext, r *resource) bool {
	numTasks := r.Data.Get("task.#").(int)
	if numTasks == 0 {
		log.Printf("[WARN] Ignoring job with ID %s", r.ID)
		ic.addIgnoredResource(fmt.Sprintf("databricks_job. id=%s", r.ID))
	}
	return numTasks == 0
}

func (ic *importContext) emitJobsDestinationNotifications(notifications []sdk_jobs.Webhook) {
	for _, notification := range notifications {
		ic.Emit(&resource{
			Resource: "databricks_notification_destination",
			ID:       notification.Id,
		})
	}
}

var (
	// This is the list of dependencies that are needed for a job. It doesn't include the dependencies for for_each_task
	// as it will be added by createJobDependencies function.
	baseJobDependencies = []reference{
		{Path: "job_cluster.new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
		{Path: "job_cluster.new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		{Path: "job_cluster.new_cluster.init_scripts.volumes.destination", Resource: "databricks_file"},
		{Path: "job_cluster.new_cluster.init_scripts.workspace.destination", Resource: "databricks_workspace_file"},
		{Path: "job_cluster.new_cluster.driver_instance_pool_id", Resource: "databricks_instance_pool"},
		{Path: "job_cluster.new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
		{Path: "job_cluster.new_cluster.policy_id", Resource: "databricks_cluster_policy"},
		{Path: "run_as.service_principal_name", Resource: "databricks_service_principal", Match: "application_id"},
		{Path: "task.dbt_task.warehouse_id", Resource: "databricks_sql_endpoint"},
		{Path: "task.dbt_task.catalog", Resource: "databricks_catalog"},
		{Path: "task.dbt_task.schema", Resource: "databricks_schema", Match: "name",
			IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema"),
		},
		{Path: "task.dashboard_task.dashboard_id", Resource: "databricks_dashboard"},
		{Path: "task.dashboard_task.warehouse_id", Resource: "databricks_sql_endpoint"},
		{Path: "task.power_bi_task.warehouse_id", Resource: "databricks_sql_endpoint"},
		{Path: "task.power_bi_task.connection_resource_name", Resource: "databricks_connection", Match: "name"},
		{Path: "task.power_bi_task.tables.catalog", Resource: "databricks_catalog"},
		{Path: "task.power_bi_task.tables.schema", Resource: "databricks_schema", Match: "name",
			IsValidApproximation: createIsMatchingCatalogAndSchema("catalog", "schema"),
		},
		{Path: "task.power_bi_task.tables.name", Resource: "databricks_sql_table", Match: "name",
			IsValidApproximation: createIsMatchingCatalogAndSchemaAndTable("catalog", "schema", "name")},
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
		{Path: "task.spark_python_task.python_file", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "task.spark_submit_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		{Path: "task.spark_submit_task.parameters", Resource: "databricks_file"},
		{Path: "task.spark_submit_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "task.sql_task.file.path", Resource: "databricks_workspace_file", Match: "path"},
		{Path: "task.sql_task.file.path", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "task.dbt_task.project_directory", Resource: "databricks_directory", Match: "path"},
		{Path: "task.dbt_task.project_directory", Resource: "databricks_directory", Match: "workspace_path"},
		{Path: "task.sql_task.alert.alert_id", Resource: "databricks_alert"},
		{Path: "task.sql_task.alert.subscriptions.destination_id", Resource: "databricks_notification_destination"},
		{Path: "task.sql_task.dashboard.dashboard_id", Resource: "databricks_sql_dashboard"},
		{Path: "task.sql_task.query.query_id", Resource: "databricks_query"},
		{Path: "task.sql_task.warehouse_id", Resource: "databricks_sql_endpoint"},
		{Path: "task.webhook_notifications.on_duration_warning_threshold_exceeded.id",
			Resource: "databricks_notification_destination"},
		{Path: "task.webhook_notifications.on_failure.id", Resource: "databricks_notification_destination"},
		{Path: "task.webhook_notifications.on_start.id", Resource: "databricks_notification_destination"},
		{Path: "task.webhook_notifications.on_success.id", Resource: "databricks_notification_destination"},
		{Path: "task.webhook_notifications.on_streaming_backlog_exceeded.id", Resource: "databricks_notification_destination"},
		{Path: "parameter.default", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "parameter.default", Resource: "databricks_workspace_file", Match: "path"},
		{Path: "parameter.default", Resource: "databricks_file", Match: "path"},
		{Path: "environments.spec.dependencies", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "environments.spec.dependencies", Resource: "databricks_file"},
		{Path: "environments.spec.dependencies", Resource: "databricks_workspace_file", Match: "workspace_path",
			MatchType: MatchRegexp, Regexp: requirementsFileRegexp},
		{Path: "environments.spec.dependencies", Resource: "databricks_file", MatchType: MatchRegexp,
			Regexp: requirementsFileRegexp},
		{Path: "webhook_notifications.on_duration_warning_threshold_exceeded.id",
			Resource: "databricks_notification_destination"},
		{Path: "webhook_notifications.on_failure.id", Resource: "databricks_notification_destination"},
		{Path: "webhook_notifications.on_start.id", Resource: "databricks_notification_destination"},
		{Path: "webhook_notifications.on_success.id", Resource: "databricks_notification_destination"},
		{Path: "webhook_notifications.on_streaming_backlog_exceeded.id",
			Resource: "databricks_notification_destination"},
		{Path: "trigger.table_update.table_names", Resource: "databricks_sql_table"},
		{Path: "trigger.file_arrival.url", Resource: "databricks_external_location",
			Match: "url", MatchType: MatchLongestPrefix},
		{Path: "task.sql_task.alert.subscriptions.user_name",
			Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "task.email_notifications.on_duration_warning_threshold_exceeded", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "task.email_notifications.on_failure", Resource: "databricks_user", Match: "user_name",
			MatchType: MatchCaseInsensitive},
		{Path: "task.email_notifications.on_start", Resource: "databricks_user", Match: "user_name",
			MatchType: MatchCaseInsensitive},
		{Path: "task.email_notifications.on_success", Resource: "databricks_user", Match: "user_name",
			MatchType: MatchCaseInsensitive},
		{Path: "task.email_notifications.on_streaming_backlog_exceeded", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "run_as.user_name", Resource: "databricks_user", Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "email_notifications.on_duration_warning_threshold_exceeded", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "email_notifications.on_failure", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "email_notifications.on_start", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "email_notifications.on_success", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "email_notifications.on_streaming_backlog_exceeded", Resource: "databricks_user",
			Match: "user_name", MatchType: MatchCaseInsensitive},
		{Path: "task.library.whl", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		{Path: "task.new_cluster.init_scripts.workspace.destination", Resource: "databricks_repo", Match: "path",
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
		{Path: "job_cluster.new_cluster.init_scripts.workspace.destination",
			Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		{Path: "job_cluster.new_cluster.init_scripts.workspace.destination",
			Resource: "databricks_repo", Match: "path"},
		{Path: "parameter.default", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
	}
)

func createJobDependencies() []reference {
	dependencies := make([]reference, 0, 2*len(baseJobDependencies))
	for _, dep := range baseJobDependencies {
		dependencies = append(dependencies, dep)
		if strings.HasPrefix(dep.Path, "task.") {
			new_dep := dep
			new_dep.Path = "task.for_each_task." + dep.Path
			dependencies = append(dependencies, new_dep)
		}
	}
	return dependencies
}
