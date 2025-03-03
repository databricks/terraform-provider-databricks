---
subcategory: "FinOps"
---
# databricks_budget_policy Resource

Administrators can use budget policies to ensure that the correct tags appear automatically on serverless resources without depending on users to attach tags manually, allowing for customized cost reporting and chargebacks. Budget policies consist of tags that are applied to any serverless compute activity incurred by a user assigned to the policy. The tags are logged in your billing records, allowing you to attribute serverless usage to specific budgets.

## Example Usage

```hcl
resource "databricks_budget_policy" "this" {
  policy_name = "my-budget-policy"
  custom_tags = [{
    key = "mykey"
    value = "myvalue"
  }]
}
```

## Argument Reference

The following arguments are available:
* `policy_name` - (Required) The name of the Budget Policy. String
* `custom_tags` - (Optional) Custom tags specifying list of key and value pairs. String

### custom_tags Configuration Block (Required)
* `key` - Key of the tag. Must be unique among all custom tags of same policy. String
* `value` - Value of the tag. String


## Attribute Reference

In addition to all arguments above, the following attributes are exported:
* `policy_id` - The ID of the budget policy 


## Import

This resource can be imported by ID.

```sh
terraform import databricks_budget_policy.this policy_id
```