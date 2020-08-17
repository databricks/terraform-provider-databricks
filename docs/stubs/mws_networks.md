# databricks_mws_networks Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `account_id` - (Required) (String) 

* `error_messages` - (Computed) (List)  This field is a block and is documented below.

* `network_name` - (Required) (String) 

* `vpc_id` - (Required) (String) 

* `subnet_ids` - (Required) (Set) 

* `security_group_ids` - (Required) (Set) 



### error_messages Configuration Block


* `error_type` - (Computed) (String) 

* `error_message` - (Computed) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws networks.

* `creation_time` - (Integer) 

* `network_id` - (String) 

* `vpc_status` - (String) 

* `workspace_id` - (Integer) 


## Import

The resource mws networks can be imported using the `object`, e.g.

```bash
$ terraform import databricks_mws_networks.object <mws networks id>
```