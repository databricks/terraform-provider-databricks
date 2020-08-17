# databricks_mws_credentials Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `account_id` - (Required) (String) 

* `credentials_name` - (Required) (String) 

* `role_arn` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws credentials.

* `creation_time` - (Integer) 

* `external_id` - (String) 

* `credentials_id` - (String) 


## Import

The resource mws credentials can be imported using the `object`, e.g.

```bash
$ terraform import databricks_mws_credentials.object <mws credentials id>
```