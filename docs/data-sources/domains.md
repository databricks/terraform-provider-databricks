---
subcategory: "Domains"
---
# databricks_domains Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/domains)

Lists the Discover domains in the account, returning each domain's governed tag key and presentation metadata (subtitle, description, owners, icon).


## Example Usage
```hcl
data "databricks_domains" "all" {}

output "domain_names" {
  value = [for d in data.databricks_domains.all.domains : d.name]
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional)
* `parent_domain_id` (string, optional) - Filter by parent domain.
  - Absent: return all domains regardless of hierarchy.
  - Present: return only direct children of the specified domain
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `domains`. It is a list of resources, each with the following attributes:
* `business_owner_ids` (list of integer) - Principal IDs of the business owners (users, groups, or service principals)
* `create_time` (string) - Timestamp when the domain was created
* `description` (string) - Full description (max 4096 chars)
* `domain_id` (string) - Unique identifier for the domain. If omitted at Create, the server
  generates one
* `draft` (boolean) - Whether to mark the domain as a draft. If omitted on Create, the server
  applies a default; the resolved value is returned in `effective_draft`
* `effective_draft` (boolean) - Resolved draft state of the domain
* `icon` (DomainIcon) - Icon to display for the domain
* `name` (string) - Full resource name of the domain. The primary identifier for this resource.
  Format: `domains/{domain_id}`
  Identifies the domain on get, update, and delete. Not an input on
  create — to choose the id, set `CreateDomainRequest.domain_id`
* `parent_domain_id` (string) - Domain ID of the parent. If absent, this is a top-level domain.
  If present, this domain is a subdomain of the specified parent
* `subtitle` (string) - Short description (max 280 chars)
* `tag_key` (string) - Governed tag key associated with this domain
* `technical_owner_ids` (list of integer) - Principal IDs of the technical owners (users, groups, or service principals)
* `update_time` (string) - Timestamp when the domain was last updated

### DomainIcon
* `color` (string) - Hex color code with # prefix (e.g., "#FF5733")
* `name` (string) - Possible values are: `ADDRESS_BOOK`, `ALARM`, `ARROWS_IN`, `ATOM`, `BALLOON`, `BANK`, `BARRICADE`, `BASKET`, `BRIDGE`, `CACTUS`, `CALL_BELL`, `CARROT`, `CHART_PIE_SLICE`, `CITY`, `CLOUD`, `COINS`, `COMPASS_ROSE`, `CRANE_TOWER`, `CROWN`, `CUBE_TRANSPARENT`, `FADERS`, `FLAG_BANNER_FOLD`, `FLAG_CHECKERED`, `GAVEL`, `HAMBURGER`, `HEAD_CIRCUIT`, `HOURGLASS_HIGH`, `INTERSECT_THREE`, `MICROSCOPE`, `MOON_STARS`, `PACKAGE`, `PARACHUTE`, `PEPPER`, `PIGGY_BANK`, `PILL`, `PLANET`, `PLANT`, `PLUGS_CONNECTED`, `POPCORN`, `PRESENTATION_CHART`, `PUZZLE_PIECE`, `RAINBOW`, `RANKING`, `RECEIPT`, `ROCKET`, `RULER`, `SAILBOAT`, `SCALES`, `SCAN_SMILEY`, `SCROLL`, `SHIELD_CHECKERED`, `SNEAKER`, `SNOWFLAKE`, `SOLAR_ROOF`, `SPEEDOMETER`, `STAMP`, `STEPS`, `STRATEGY`, `SWORD`, `TELEVISION_SIMPLE`, `TENT`, `TICKET`, `TRACTOR`, `TRAFFIC_CONE`, `TRAIN`, `TREE_EVERGREEN`, `TREE_STRUCTURE`, `TROLLEY_SUITCASE`, `TROPHY`, `TRUCK_TRAILER`, `USERS_THREE`, `VECTOR_THREE`