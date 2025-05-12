---
subcategory: "FinOps"
---
# databricks_budget_policy Resource

Administrators can use budget policies to ensure that the correct tags appear automatically on serverless resources without depending on users to attach tags manually, allowing for customized cost reporting and chargebacks. Budget policies consist of tags that are applied to any serverless compute activity incurred by a user assigned to the policy. The tags are logged in your billing records, allowing you to attribute serverless usage to specific budgets.

-> This resource can only be used with an account-level provider!

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

* `policy_name` - (Required) The name of the policy. Must be unique among active policies. Can contain only characters from the ISO 8859-1 (latin1) set.
* `custom_tags` - (Optional) A list of tags defined by the customer. At most 20 entries are allowed per policy. 

### custom_tags Configuration Block

* `key` - The key of the tag. - Must be unique among all custom tags of the same policy. Cannot be "budget-policy-name", "budget-policy-id" or "budget-policy-resolution-result" as these tags are preserved. 
* `value` - The value of the tag. 

## Attribute Reference

In addition to all arguments above, the following attribute is exported:

* `policy_id` - ID of the budget policy 

## Access Control

* [databricks_access_control_rule_set](access_control_rule_set.md) can control which groups or individual users can manage or use the given budget policy.

## Import

This resource can be imported by ID.

```hcl
import {
  to = databricks_budget_policy.this
  id = "policy_id"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```sh
terraform import databricks_budget_policy.this policy_id
```
