# databricks_azure_blob_mount Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `directory` - (Optional) (String) 

* `auth_type` - (Required) (String) 

* `token_secret_scope` - (Required) (String) 

* `container_name` - (Required) (String) 

* `storage_account_name` - (Required) (String) 

* `token_secret_key` - (Required) (String) 

* `cluster_id` - (Optional) (String) 

* `mount_name` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the azure blob mount.

* `source` - (String) 


## Import

The resource azure blob mount can be imported using the `object`, e.g.

```bash
$ terraform import databricks_azure_blob_mount.object <azure blob mount id>
```