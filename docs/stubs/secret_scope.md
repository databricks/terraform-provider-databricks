# databricks_secret_scope Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `name` - (Required) (String) 

* `initial_manage_principal` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the secret scope.

* `backend_type` - (String) 


## Import

The resource secret scope can be imported using the `object`, e.g.

```bash
$ terraform import databricks_secret_scope.object <secret scope id>
```