# databricks_permissions Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `notebook_path` - (Optional) (String) 

* `cluster_id` - (Optional) (String) 

* `job_id` - (Optional) (String) 

* `notebook_id` - (Optional) (String) 

* `directory_id` - (Optional) (String) 

* `directory_path` - (Optional) (String) 

* `access_control` - (Required) (List)  This field is a block and is documented below.

* `cluster_policy_id` - (Optional) (String) 

* `instance_pool_id` - (Optional) (String) 



### access_control Configuration Block


* `permission_level` - (Required) (String) 

* `user_name` - (Optional) (String) 

* `group_name` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the permissions.

* `object_type` - (String) 


## Import

The resource permissions can be imported using the `object`, e.g.

```bash
$ terraform import databricks_permissions.object <permissions id>
```