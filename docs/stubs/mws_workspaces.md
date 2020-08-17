# databricks_mws_workspaces Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `customer_managed_key_id` - (Optional) (String) 

* `network_id` - (Optional) (String) 

* `is_no_public_ip_enabled` - (Optional) (Bool) 

* `account_id` - (Required) (String) 

* `credentials_id` - (Required) (String) 

* `network_error_messages` - (Computed) (List)  This field is a block and is documented below.

* `deployment_name` - (Required) (String) 

* `workspace_name` - (Required) (String) 

* `aws_region` - (Required) (String) 

* `storage_configuration_id` - (Required) (String) 

* `verify_workspace_runnning` - (Required) (Bool) 



### network_error_messages Configuration Block


* `error_message` - (Computed) (String) 

* `error_type` - (Computed) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws workspaces.

* `workspace_status_message` - (String) 

* `workspace_status` - (String) 

* `creation_time` - (Integer) 

* `workspace_url` - (String) 

* `workspace_id` - (Integer) 


## Import

The resource mws workspaces can be imported using the `object`, e.g.

```bash
$ terraform import databricks_mws_workspaces.object <mws workspaces id>
```