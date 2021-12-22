---
subcategory: "Workspace"
---
# databricks_global_init_script Resource

This resource allows you to manage [Databricks global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).

## Example Usage

You can declare Terraform-managed global init script by specifying `source` attribute of corresponding local file.

```hcl
resource "databricks_global_init_script" "init1" {
  source = "${path.module}/init.sh"
  name   = "my init script"
}
```

You can also create managed global init script with inline sources through `content_base64` attribute.

```hcl
resource "databricks_global_init_script" "init2" {
  content_base64 = base64encode(<<-EOT
    #!/bin/bash
    echo "hello world"
    EOT
  )
  name = "hello script"
}
```
    
## Argument Reference

-> **Note** Global init script in the Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed global init script won't be overwritten by Terraform, if there's no local change to source.

The size of a global init script source code must not exceed 64Kb. The following arguments are supported:

* `name` (string, required) - the name of the script.  It should be unique
* `source` - Path to script's source code on local filesystem. Conflicts with `content_base64`
* `content_base64` - The base64-encoded source code global init script. Conflicts with `source`. Use of `content_base64` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances
* `enabled` (bool, optional default: `false`) specifies if the script is enabled for execution, or not
* `position` (integer, optional default: `null`) - the position of a global init script, where `0` represents the first global init script to run, `1` is the second global init script to run, and so on. When omitted, the script gets the last position.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID assigned to a global init script by API

## Access Control

Global init scripts are available only for administrators, so you can't change permissions for it.

## Import

The resource global init script can be imported using script ID:

```bash
$ terraform import databricks_global_init_script.this script_id
```
