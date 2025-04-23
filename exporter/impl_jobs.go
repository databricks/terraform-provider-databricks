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
	if task.DbtTask != nil {
		if task.DbtTask.WarehouseId != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_endpoint",
				ID:       task.DbtTask.WarehouseId,
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

	return ic.importLibraries(r.Data, s)
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

func shouldOmitFieldInJob(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
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
		return makeShouldOmitFieldForCluster(jobClustersRegex)(ic, pathString, as, d)
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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
