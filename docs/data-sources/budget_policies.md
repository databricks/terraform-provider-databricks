---
subcategory: "FinOps"
---
# databricks_budget_policies Data Source
This data source can be used to fetch the list of budget policies. 

-> **Note** This data source can only be used with an account-level provider!


## Example Usage

Getting a list of all budget policies:

```hcl
data "databricks_budget_policies" "all" {
}
```

## Argument Reference
- `filter_by` - A filter to apply to the list of policies. This consists of: 
    - `policy_name` - The partial name of policies to be filtered on. If unspecified, all policies will be returned.
    - `creator_user_id` - The policy creator user id to be filtered on. If unspecified, all policies will be returned.
    - `creator_user_name` - The policy creator user name to be filtered on. If unspecified, all policies will be returned.

## Attribute Reference

Data source exposes the following attributes:

- `budget_policies` - The list of budget policy.