package model

import "fmt"

// NotebookTask contains the information for notebook jobs
type NotebookTask struct {
	NotebookPath   string            `json:"notebook_path"`
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

// SparkPythonTask contains the information for python jobs
type SparkPythonTask struct {
	PythonFile string   `json:"python_file"`
	Parameters []string `json:"parameters,omitempty"`
}

// SparkJarTask contains the information for jar jobs
type SparkJarTask struct {
	JarURI        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
}

// SparkSubmitTask contains the information for spark submit jobs
type SparkSubmitTask struct {
	Parameters []string `json:"parameters,omitempty"`
}

// JobEmailNotifications contains the information for email notifications after job completion
type JobEmailNotifications struct {
	OnStart               []string `json:"on_start,omitempty"`
	OnSuccess             []string `json:"on_success,omitempty"`
	OnFailure             []string `json:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs,omitempty"`
}

// CronSchedule contains the information for the quartz cron expression
type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression"`
	TimezoneID           string `json:"timezone_id"`
	PauseStatus          string `json:"pause_status,omitempty"`
}

// JobSettings contains the information for configuring a job on databricks
type JobSettings struct {
	Name string `json:"name,omitempty" tf:"default:Untitled"`

	ExistingClusterID string   `json:"existing_cluster_id,omitempty" tf:"group:cluster_type"`
	NewCluster        *Cluster `json:"new_cluster,omitempty" tf:"group:cluster_type"`

	NotebookTask    *NotebookTask    `json:"notebook_task,omitempty" tf:"group:task_type"`
	SparkJarTask    *SparkJarTask    `json:"spark_jar_task,omitempty" tf:"group:task_type"`
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty" tf:"group:task_type"`
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty" tf:"group:task_type"`

	Libraries              []Library     `json:"libraries,omitempty" tf:"slice_set,alias:library"`
	TimeoutSeconds         int32         `json:"timeout_seconds,omitempty"`
	MaxRetries             int32         `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32         `json:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool          `json:"retry_on_timeout,omitempty"`
	Schedule               *CronSchedule `json:"schedule,omitempty"`
	MaxConcurrentRuns      int32         `json:"max_concurrent_runs,omitempty"`

	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
}

// Job contains the information when using a GET request from the Databricks Jobs api
type Job struct {
	JobID           int64        `json:"job_id,omitempty"`
	CreatorUserName string       `json:"creator_user_name,omitempty"`
	Settings        *JobSettings `json:"settings,omitempty"`
	CreatedTime     int64        `json:"created_time,omitempty"`
}

// ID returns job id as string
func (j Job) ID() string {
	return fmt.Sprintf("%d", j.JobID)
}

// UpdateJobRequest ...
type UpdateJobRequest struct {
	JobID       int64        `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings *JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
}
