# databricks_azure_adls_gen1_mount Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `storage_resource_name` - (Required) (String) 

* `client_id` - (Required) (String) 

* `client_secret_key` - (Required) (String) 

* `cluster_id` - (Optional) (String) 

* `mount_name` - (Required) (String) 

* `directory` - (Computed) (String) 

* `tenant_id` - (Required) (String) 

* `client_secret_scope` - (Required) (String) 

* `spark_conf_prefix` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the azure adls gen1 mount.

* `source` - (String) 


## Import

The resource azure adls gen1 mount can be imported using the `object`, e.g.

```bash
$ terraform import databricks_azure_adls_gen1_mount.object <azure adls gen1 mount id>
```