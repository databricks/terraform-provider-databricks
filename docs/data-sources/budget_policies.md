---
subcategory: "FinOps"
---
# databricks_budget_policies Data Source
This data source can be used to fetch the list of budget policies. 

## Example Usage

Referring to a budget policy by name:

```hcl
data "databricks_budget_policies" "all" {
}
```

## Argument Reference


## Attribute Reference

Data source exposes the following attributes:

- `budget_policies` - The list of the budget policy.