# databricks_token Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `creation_time` - (Computed) (Integer) 

* `expiry_time` - (Computed) (Integer) 

* `lifetime_seconds` - (Optional) (Integer) 

* `comment` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the token.

* `token_value` - (String) 


## Import

The resource token can be imported using the `object`, e.g.

```bash
$ terraform import databricks_token.object <token id>
```