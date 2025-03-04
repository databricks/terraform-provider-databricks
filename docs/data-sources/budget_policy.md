---
subcategory: "FinOps"
---
# databricks_budget_policy Data Source
This data source can be used to get a single budget policy. 

-> **Note** This data source can only be used with an account-level provider!

## Example Usage

Referring to a budget policy by id:

```hcl
data "databricks_budget_policy" "this" {
  policy_id = "test"
}
```

## Argument Reference

Data source allows you to get a budget policy by the following attribute

- `policy_id` - ID of the budget policy.

## Attribute Reference

Data source exposes the following attributes:

- `policy_id` - The id of the budget policy.
- `policy_name` - The name of the budget policy.