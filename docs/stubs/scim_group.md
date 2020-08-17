# databricks_scim_group Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `display_name` - (Required) (String) 

* `members` - (Optional) (Set) 

* `roles` - (Optional) (Set) 

* `entitlements` - (Optional) (Set) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the scim group.

* `inherited_roles` - (Set) 


## Import

The resource scim group can be imported using the `object`, e.g.

```bash
$ terraform import databricks_scim_group.object <scim group id>
```