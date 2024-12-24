package exporter

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"

	"github.com/databricks/terraform-provider-databricks/common"
	tf_jobs "github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/workspace"
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
			if ic.isInRepoOrGitFolder(directory, true) {
				ic.emitRepoOrGitFolder(directory, true)
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
	if job.WebhookNotifications != nil {
		ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnFailure)
		ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnSuccess)
		ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnDurationWarningThresholdExceeded)
		ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnStart)
		ic.emitJobsDestinationNotifications(job.WebhookNotifications.OnStreamingBacklogExceeded)
	}
	for _, param := range job.Parameters {
		ic.emitIfWsfsFile(param.Default)
		ic.emitIfVolumeFile(param.Default)
	}

	return ic.importLibraries(r.Data, s)
}
