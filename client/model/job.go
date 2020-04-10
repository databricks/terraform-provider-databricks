package model

//go:generate easytags $GOFILE

type NotebookTask struct {
	NotebookPath   string            `json:"notebook_path,omitempty"`
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

type SparkPythonTask struct {
	PythonFile string   `json:"python_file,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
}

type SparkJarTask struct {
	JarUri        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
}

type SparkSubmitTask struct {
	Parameters []string `json:"parameters,omitempty"`
}

type JobEmailNotifications struct {
	OnStart               []string `json:"on_start,omitempty"`
	OnSuccess             []string `json:"on_success,omitempty"`
	OnFailure             []string `json:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs,omitempty"`
}

type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression,omitempty"`
	TimezoneId           string `json:"timezone_id,omitempty"`
}

type JobSettings struct {
	ExistingClusterId      string                 `json:"existing_cluster_id,omitempty"`
	NewCluster             *Cluster               `json:"new_cluster,omitempty"`
	NotebookTask           *NotebookTask          `json:"notebook_task,omitempty"`
	SparkJarTask           *SparkJarTask          `json:"spark_jar_task,omitempty"`
	SparkPythonTask        *SparkPythonTask       `json:"spark_python_task,omitempty"`
	SparkSubmitTask        *SparkSubmitTask       `json:"spark_submit_task,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Libraries              []Library              `json:"libraries,omitempty"`
	EmailNotifications     *JobEmailNotifications `json:"email_notifications,omitempty"`
	TimeoutSeconds         int32                  `json:"timeout_seconds,omitempty"`
	MaxRetries             int32                  `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32                  `json:"max_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool                   `json:"retry_on_timeout,omitempty"`
	Schedule               *CronSchedule          `json:"schedule,omitempty"`
	MaxConcurrentRuns      int32                  `json:"max_concurrent_runs,omitempty"`
}

type Job struct {
	JobId           int64        `json:"job_id,omitempty"`
	CreatorUserName string       `json:"creator_user_name,omitempty"`
	Settings        *JobSettings `json:"settings,omitempty"`
	CreatedTime     int64        `json:"created_time,omitempty"`
}
