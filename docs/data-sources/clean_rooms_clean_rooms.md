---
subcategory: "Clean Rooms"
---
# databricks_clean_rooms_clean_rooms Data Source
# Datasource (Plural) Artifact

This data source can be used to fetch the list of clean rooms.

## Example Usage
# Example: Clean Room Datasource (Plural)

```hcl
data "databricks_clean_rooms_clean_room" "all" {}
```

## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of clean rooms to return (i.e., the page length). Defaults to 100



## Attributes
This data source exports a single attribute, `clean_rooms`. It is a list of resources, each with the following attributes:
* `access_restricted` (string) - Whether clean room access is restricted due to [CSP](https://docs.databricks.com/en/security/privacy/security-profile.html). Possible values are: `CSP_MISMATCH`, `NO_RESTRICTION`
* `comment` (string) - 
* `created_at` (integer) - When the clean room was created, in epoch milliseconds
* `local_collaborator_alias` (string) - The alias of the collaborator tied to the local clean room
* `name` (string) - The name of the clean room.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements)
* `output_catalog` (CleanRoomOutputCatalog) - Output catalog of the clean room. It is an output only field. Output catalog is manipulated
  using the separate CreateCleanRoomOutputCatalog API
* `owner` (string) - This is Databricks username of the owner of the local clean room securable for permission management
* `remote_detailed_info` (CleanRoomRemoteDetail) - Central clean room details. During creation, users need to specify
  cloud_vendor, region, and collaborators.global_metastore_id.
  This field will not be filled in the ListCleanRooms call
* `status` (string) - Clean room status. Possible values are: `ACTIVE`, `DELETED`, `FAILED`, `PROVISIONING`
* `updated_at` (integer) - When the clean room was last updated, in epoch milliseconds

### CleanRoomCollaborator
* `collaborator_alias` (string) - Collaborator alias specified by the clean room creator. It is unique across all collaborators of this clean room, and used to derive
  multiple values internally such as catalog alias and clean room name for single metastore clean rooms.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements)
* `display_name` (string) - Generated display name for the collaborator. In the case of a single metastore clean room, it is the clean
  room name. For x-metastore clean rooms, it is the organization name of the metastore. It is not restricted to
  these values and could change in the future
* `global_metastore_id` (string) - The global Unity Catalog metastore id of the collaborator. The identifier is of format cloud:region:metastore-uuid
* `invite_recipient_email` (string) - Email of the user who is receiving the clean room "invitation". It should be empty
  for the creator of the clean room, and non-empty for the invitees of the clean room.
  It is only returned in the output when clean room creator calls GET
* `invite_recipient_workspace_id` (integer) - Workspace ID of the user who is receiving the clean room "invitation". Must be specified if
  invite_recipient_email is specified.
  It should be empty when the collaborator is the creator of the clean room
* `organization_name` (string) - [Organization name](:method:metastores/list#metastores-delta_sharing_organization_name)
  configured in the metastore

### CleanRoomOutputCatalog
* `catalog_name` (string) - The name of the output catalog in UC.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements).
  The field will always exist if status is CREATED
* `status` (string) - . Possible values are: `CREATED`, `NOT_CREATED`, `NOT_ELIGIBLE`

### CleanRoomRemoteDetail
* `central_clean_room_id` (string) - Central clean room ID
* `cloud_vendor` (string) - Cloud vendor (aws,azure,gcp) of the central clean room
* `collaborators` (list of CleanRoomCollaborator) - Collaborators in the central clean room. There should one and only one collaborator
  in the list that satisfies the owner condition:
  
  1. It has the creator's global_metastore_id (determined by caller of CreateCleanRoom).
  
  2. Its invite_recipient_email is empty
* `compliance_security_profile` (ComplianceSecurityProfile) - 
* `creator` (CleanRoomCollaborator) - Collaborator who creates the clean room
* `egress_network_policy` (EgressNetworkPolicy) - Egress network policy to apply to the central clean room workspace
* `region` (string) - Region of the central clean room

### ComplianceSecurityProfile
* `compliance_standards` (list of ComplianceStandard) - The list of compliance standards that the compliance security profile is configured to enforce
* `is_enabled` (boolean) - Whether the compliance security profile is enabled

### EgressNetworkPolicy
* `internet_access` (EgressNetworkPolicyInternetAccessPolicy) - The access policy enforced for egress traffic to the internet

### EgressNetworkPolicyInternetAccessPolicy
* `allowed_internet_destinations` (list of EgressNetworkPolicyInternetAccessPolicyInternetDestination) - 
* `allowed_storage_destinations` (list of EgressNetworkPolicyInternetAccessPolicyStorageDestination) - 
* `log_only_mode` (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) - Optional. If not specified, assume the policy is enforced for all workloads
* `restriction_mode` (string) - . Possible values are: `FULL_ACCESS`, `PRIVATE_ACCESS_ONLY`, `RESTRICTED_ACCESS`

### EgressNetworkPolicyInternetAccessPolicyInternetDestination
* `destination` (string) - 
* `protocol` (string) - . Possible values are: `TCP`
* `type` (string) - . Possible values are: `FQDN`

### EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
* `log_only_mode_type` (string) - . Possible values are: `ALL_SERVICES`, `SELECTED_SERVICES`
* `workloads` (list of ) - 

### EgressNetworkPolicyInternetAccessPolicyStorageDestination
* `allowed_paths` (list of string) - 
* `azure_container` (string) - 
* `azure_dns_zone` (string) - 
* `azure_storage_account` (string) - 
* `azure_storage_service` (string) - 
* `bucket_name` (string) - 
* `region` (string) - 
* `type` (string) - . Possible values are: `AWS_S3`, `AZURE_STORAGE`, `CLOUDFLARE_R2`, `GOOGLE_CLOUD_STORAGE`