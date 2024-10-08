---
subcategory: "Delta Sharing"
---
# databricks_provider Resource

-> This resource can only be used with a workspace-level provider!

In Delta Sharing, a provider is an entity that shares data with a recipient. Within a metastore, Unity Catalog provides the ability to create a provider which contains a list of shares that have been shared with you.

A `databricks_provider` is contained within [databricks_metastore](metastore.md) and can contain a list of shares that have been shared with you.

-> Databricks to Databricks sharing automatically creates the provider.

## Example Usage

```hcl
resource "databricks_provider" "dbprovider" {
  name                = "terraform-test-provider"
  comment             = "made by terraform 2"
  authentication_type = "TOKEN"
  recipient_profile_str = jsonencode(
    {
      "shareCredentialsVersion" : 1,
      "bearerToken" : "token",
      "endpoint" : "endpoint",
      "expirationTime" : "expiration-time"
    }
  )
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of provider. Change forces creation of a new resource.
* `comment` - (Optional) Description about the provider.
* `authentication_type` - (Optional) The delta sharing authentication type. Valid values are `TOKEN`.
* `recipient_profile_str` - (Optional) This is the json file that is created from a recipient url.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this provider - same as the `name`.

## Related Resources

The following resources are used in the same context:

* [databricks_tables](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schemas](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalogs](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
