# databricks_secret Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `string_value` - (Required) (String) 

* `scope` - (Required) (String) 

* `key` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the secret.

* `last_updated_timestamp` - (Integer) 


## Import

The resource secret can be imported using the `object`, e.g.

```bash
$ terraform import databricks_secret.object <secret id>
```