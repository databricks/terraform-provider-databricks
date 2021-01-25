# databricks_caller_identity Data Source

Retrieves information about [databricks_user](../resources/group.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API. Might be useful in applying the same Terraform by different users in the shared workspace for testing purposes. 

## Example Usage

Create personalized job and notebook:

```hcl
data "databricks_caller_identity" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" { 
    local_disk = true 
}

resource "databricks_notebook" "this" {
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
    EOT
  )
  path = "${data.databricks_caller_identity.home}/Terraform"
  language = "PYTHON"
}

resource "databricks_job" "this" {
    name = "Terraform Demo (${data.databricks_caller_identity.alphanumeric})"
    
    new_cluster  {
        num_workers   = 1
        spark_version = data.databricks_spark_version.latest.id
        node_type_id  = data.databricks_node_type.smallest.id
    }
    
    notebook_task {
        notebook_path = databricks_notebook.this.path
    }
    
    email_notifications {}
}

output "notebook_url" {
    value = databricks_notebook.this.url
}
```

## Exported attributes

Data source exposes the following attributes:

* `id` -  The id of the calling user.
* `user_name` - Name of the [user](../resources/user.md), e.g. `mr.foo@example.com`.
* `home` - Home folder of the [user](../resources/user.md), e.g. `/Users/mr.foo@example.com`.
* `alphanumeric` - Alphanumeric representation of user local name. e.g. `mr_foo`.