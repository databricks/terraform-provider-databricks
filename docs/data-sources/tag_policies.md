---
subcategory: "Tags"
---
# databricks_tag_policies Data Source
This data source can be used to list all tag policies in the account.

-> **Note** This resource can only be used with an account-level provider!

## Example Usage
```hcl
```

## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return in this request. Fewer results may be returned than requested. If
  unspecified or set to 0, this defaults to 1000. The maximum value is 1000; values above 1000 will be coerced down
  to 1000



## Attributes
This data source exports a single attribute, `tag_policies`. It is a list of resources, each with the following attributes:
* `description` (string)
* `id` (string)
* `tag_key` (string)
* `values` (list of Value)

### Value
* `name` (string)