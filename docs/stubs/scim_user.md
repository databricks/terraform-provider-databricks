# databricks_scim_user Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `roles` - (Optional) (Set) 

* `entitlements` - (Optional) (Set) 

* `default_roles` - (Required) (Set) 

* `set_admin` - (Optional) (Bool) 

* `user_name` - (Required) (String) 

* `display_name` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the scim user.

* `inherited_roles` - (Set) 


## Import

The resource scim user can be imported using the `object`, e.g.

```bash
$ terraform import databricks_scim_user.object <scim user id>
```