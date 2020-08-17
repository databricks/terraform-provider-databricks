# databricks_mws_storage_configurations Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `bucket_name` - (Required) (String) 

* `account_id` - (Required) (String) 

* `storage_configuration_name` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws storage configurations.

* `creation_time` - (Integer) 

* `storage_configuration_id` - (String) 


## Import

The resource mws storage configurations can be imported using the `object`, e.g.

```bash
$ terraform import databricks_mws_storage_configurations.object <mws storage configurations id>
```