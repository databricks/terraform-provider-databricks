---
subcategory: "Domains"
---
# databricks_domain Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/domains)

This resource allows you to create, read, update, and delete Discover domains. A domain is an organizational container that groups a governed tag key with presentation metadata (subtitle, description, owners, and an icon), letting teams curate and browse data assets by domain.

The `tag_key` must reference a governed tag that already exists in the default tag-policy namespace, and it is immutable once the domain is created.


## Example Usage
```hcl
resource "databricks_domain" "finance" {
  tag_key     = "finance"
  subtitle    = "Financial data assets"
  description = "Curated finance datasets for the analytics org."

  icon = {
    name  = "BANK"
    color = "#1B5E20"
  }
}
```


## Arguments
The following arguments are supported:
* `tag_key` (string, required) - Governed tag key associated with this domain
* `business_owner_ids` (list of integer, optional) - Principal IDs of the business owners (users, groups, or service principals)
* `description` (string, optional) - Full description (max 4096 chars)
* `draft` (boolean, optional) - Whether to mark the domain as a draft. If omitted on Create, the server
  applies a default; the resolved value is returned in `effective_draft`
* `icon` (DomainIcon, optional) - Icon to display for the domain
* `parent_domain_id` (string, optional) - Domain ID of the parent. If absent, this is a top-level domain.
  If present, this domain is a subdomain of the specified parent
* `subtitle` (string, optional) - Short description (max 280 chars)
* `technical_owner_ids` (list of integer, optional) - Principal IDs of the technical owners (users, groups, or service principals)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DomainIcon
* `color` (string, optional) - Hex color code with # prefix (e.g., "#FF5733")
* `name` (string, optional) - Possible values are: `ADDRESS_BOOK`, `ALARM`, `ARROWS_IN`, `ATOM`, `BALLOON`, `BANK`, `BARRICADE`, `BASKET`, `BRIDGE`, `CACTUS`, `CALL_BELL`, `CARROT`, `CHART_PIE_SLICE`, `CITY`, `CLOUD`, `COINS`, `COMPASS_ROSE`, `CRANE_TOWER`, `CROWN`, `CUBE_TRANSPARENT`, `FADERS`, `FLAG_BANNER_FOLD`, `FLAG_CHECKERED`, `GAVEL`, `HAMBURGER`, `HEAD_CIRCUIT`, `HOURGLASS_HIGH`, `INTERSECT_THREE`, `MICROSCOPE`, `MOON_STARS`, `PACKAGE`, `PARACHUTE`, `PEPPER`, `PIGGY_BANK`, `PILL`, `PLANET`, `PLANT`, `PLUGS_CONNECTED`, `POPCORN`, `PRESENTATION_CHART`, `PUZZLE_PIECE`, `RAINBOW`, `RANKING`, `RECEIPT`, `ROCKET`, `RULER`, `SAILBOAT`, `SCALES`, `SCAN_SMILEY`, `SCROLL`, `SHIELD_CHECKERED`, `SNEAKER`, `SNOWFLAKE`, `SOLAR_ROOF`, `SPEEDOMETER`, `STAMP`, `STEPS`, `STRATEGY`, `SWORD`, `TELEVISION_SIMPLE`, `TENT`, `TICKET`, `TRACTOR`, `TRAFFIC_CONE`, `TRAIN`, `TREE_EVERGREEN`, `TREE_STRUCTURE`, `TROLLEY_SUITCASE`, `TROPHY`, `TRUCK_TRAILER`, `USERS_THREE`, `VECTOR_THREE`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Timestamp when the domain was created
* `domain_id` (string) - Unique identifier for the domain. If omitted at Create, the server
  generates one
* `effective_draft` (boolean) - Resolved draft state of the domain
* `name` (string) - Full resource name of the domain. The primary identifier for this resource.
  Format: `domains/{domain_id}`
  Identifies the domain on get, update, and delete. Not an input on
  create — to choose the id, set `CreateDomainRequest.domain_id`
* `update_time` (string) - Timestamp when the domain was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_domain.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_domain.this "name"
```