# databricks_notebook Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `mkdirs` - (Optional) (Bool) 

* `format` - (Optional) (String) 

* `content` - (Required) (String) 

* `path` - (Required) (String) 

* `language` - (Optional) (String) 

* `overwrite` - (Optional) (Bool) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the notebook.

* `object_type` - (String) 

* `object_id` - (Integer) 


## Import

The resource notebook can be imported using the `object`, e.g.

```bash
$ terraform import databricks_notebook.object <notebook id>
```