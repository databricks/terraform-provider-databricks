# databricks_cluster_policy Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `name` - (Required) (String) Cluster policy name. This must be unique.
Length must be between 1 and 100 characters.

* `definition` - (Optional) (String) Policy definition JSON document expressed in
Databricks Policy Definition Language.




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster policy.

* `policy_id` - (String) 


## Import

The resource cluster policy can be imported using the `object`, e.g.

```bash
$ terraform import databricks_cluster_policy.object <cluster policy id>
```