---
subcategory: "MLflow"
---
# databricks_mlflow_webhook Resource

This resource allows you to create [MLflow Model Registry Webhooks](https://docs.databricks.com/applications/mlflow/model-registry-webhooks.html) in Databricks.  Webhooks enable you to listen for Model Registry events so your integrations can automatically trigger actions. You can use webhooks to automate and integrate your machine learning pipeline with existing CI/CD tools and workflows. Webhooks allow trigger execution of a Databricks job or call a web service on specific event(s) that is generated in the MLflow Registry - stage transitioning, creation of registered model, creation of transition request, etc.

## Example Usage

### Triggering Databricks job

```hcl
data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/MLFlowWebhook"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    import json
 
    event_message = dbutils.widgets.get("event_message")
    event_message_dict = json.loads(event_message)
    print(f"event data={event_message_dict}")
    EOT
  )
}

resource "databricks_job" "this" {
  name = "Terraform MLflowWebhook Demo (${data.databricks_current_user.me.alphanumeric})"

  task {
    task_key = "task1"

    new_cluster {
      num_workers   = 1
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }

    notebook_task {
      notebook_path = databricks_notebook.this.path
    }
  }
}

resource "databricks_token" "pat_for_webhook" {
  comment          = "MLflow Webhook"
  lifetime_seconds = 86400000
}

resource "databricks_mlflow_webhook" "job" {
  events      = ["TRANSITION_REQUEST_CREATED"]
  description = "Databricks Job webhook trigger"
  status      = "ACTIVE"
  job_spec {
    job_id        = databricks_job.this.id
    workspace_url = data.databricks_current_user.me.workspace_url
    access_token  = databricks_token.pat_for_webhook.token_value
  }
}
```

### POSTing to URL

```hcl
resource "databricks_mlflow_webhook" "url" {
  events      = ["TRANSITION_REQUEST_CREATED"]
  description = "URL webhook trigger"
  http_url_spec {
    url = "https://my_cool_host/webhook"
  }
}
```

## Argument Reference

The following arguments are supported:

* `model_name` - (Optional) Name of MLflow model for which webhook will be created. If the model name is not specified, a registry-wide webhook is created that listens for the specified events across all versions of all registered models.
* `description` - Optional description of the MLflow webhook.
* `status` - Optional status of webhook. Possible values are `ACTIVE`, `TEST_MODE`, `DISABLED`. Default is `ACTIVE`.
* `events` - (Required) The list of events that will trigger execution of Databricks job or POSTing to an URL, for example, `MODEL_VERSION_CREATED`, `MODEL_VERSION_TRANSITIONED_STAGE`, `TRANSITION_REQUEST_CREATED`, etc.  Refer to the [Webhooks API documentation](https://docs.databricks.com/dev-tools/api/latest/mlflow.html#operation/create-registry-webhook) for a full list of supported events.

Configuration must include one of `http_url_spec` or `job_spec` blocks, but not both.

### job_spec

* `access_token` - (Required, Sensitive) The personal access token used to authorize webhook's job runs.
* `job_id` - (Required) ID of the Databricks job that the webhook runs.
* `workspace_url` - (Optional) URL of the workspace containing the job that this webhook runs. If not specified, the job’s workspace URL is assumed to be the same as the workspace where the webhook is created.

### http_url_spec

* `url` - (Required) External HTTPS URL called on event trigger (by using a POST request). Structure of payload depends on the event type, refer to [documentation](https://docs.databricks.com/applications/mlflow/model-registry-webhooks.html) for more details.
* `authorization` - (Optional) Value of the authorization header that should be sent in the request sent by the wehbook.  It should be of the form `<auth type> <credentials>`, e.g. `Bearer <access_token>`. If set to an empty string, no authorization header will be included in the request.
* `enable_ssl_verification` - (Optional) Enable/disable SSL certificate validation. Default is `true`. For self-signed certificates, this field must be `false` AND the destination server must disable certificate validation as well. For security purposes, it is encouraged to perform secret validation with the HMAC-encoded portion of the payload and acknowledge the risk associated with disabling hostname validation whereby it becomes more likely that requests can be maliciously routed to an unintended host.
* `secret` - (Optional, Sensitive) Shared secret required for HMAC encoding payload. The HMAC-encoded payload will be sent in the header as `X-Databricks-Signature: encoded_payload`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Unique ID of the MLflow Webhook.

## Access Control

* MLflow webhooks could be configured only by workspace admins.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_experiment](mlflow_experiment.md) to manage [MLflow experiments](https://docs.databricks.com/data/data-sources/mlflow-experiment.html) in Databricks.
* [databricks_mlflow_model](mlflow_model.md) to create [MLflow models](https://docs.databricks.com/applications/mlflow/models.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
