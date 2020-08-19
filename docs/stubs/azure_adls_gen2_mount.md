# databricks_azure_adls_gen2_mount Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `cluster_id` - (Optional) (String) 

* `container_name` - (Required) (String) 

* `storage_account_name` - (Required) (String) 

* `client_secret_key` - (Required) (String) 

* `mount_name` - (Required) (String) 

* `directory` - (Computed) (String) 

* `tenant_id` - (Required) (String) 

* `client_id` - (Required) (String) 

* `client_secret_scope` - (Required) (String) 

* `initialize_file_system` - (Required) (Bool) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the azure adls gen2 mount.

* `source` - (String) 


## Import

The resource azure adls gen2 mount can be imported using the `object`, e.g.

```bash
$ terraform import databricks_azure_adls_gen2_mount.object <azure adls gen2 mount id>
```