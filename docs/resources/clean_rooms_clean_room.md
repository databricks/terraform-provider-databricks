---
subcategory: "Clean Rooms"
---
# databricks_clean_rooms_clean_room Resource
# Clean Room Resource

A Clean Room is a secure environment for data collaboration that enables multiple organizations to analyze their data together while maintaining privacy and security. Clean Rooms provide a controlled environment where data can be shared and analyzed without exposing the underlying raw data.


## Example Usage
# Example: Clean Room Resource

```hcl
resource "databricks_clean_rooms_clean_room" "this" {
    name = "example-clean-room"
    owner = "example@databricks.com"
    remote_detailed_info = {
        cloud_vendor = "aws"
        region = "us-west-2"
        collaborators = [
            {
                collaborator_alias = "collaborator"
                global_metastore_id = "aws:us-east-1:12345678-1234-1234-1234-123456789012"
                invite_recipient_email: "example@databricks.com"
                invite_recipient_workspace_id = "123456789012345"
            },
            {
                global_metastore_id = "aws:us-east-1:12345678-1234-1234-1234-123456789012"
                collaborator_alias = "creator"
            }
        ]
        egress_network_policy = {
            internet_access = {
                restriction_mode = "RESTRICTED_ACCESS"
            }
        }
    }
}
``` 

## Arguments
The following arguments are supported:
* `comment` (string, optional)
* `name` (string, optional) - The name of the clean room.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements)
* `owner` (string, optional) - This is Databricks username of the owner of the local clean room securable for permission management
* `remote_detailed_info` (CleanRoomRemoteDetail, optional) - Central clean room details. During creation, users need to specify
  cloud_vendor, region, and collaborators.global_metastore_id.
  This field will not be filled in the ListCleanRooms call

### CleanRoomCollaborator
* `collaborator_alias` (string, required) - Collaborator alias specified by the clean room creator. It is unique across all collaborators of this clean room, and used to derive
  multiple values internally such as catalog alias and clean room name for single metastore clean rooms.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements)
* `global_metastore_id` (string, optional) - The global Unity Catalog metastore id of the collaborator. The identifier is of format cloud:region:metastore-uuid
* `invite_recipient_email` (string, optional) - Email of the user who is receiving the clean room "invitation". It should be empty
  for the creator of the clean room, and non-empty for the invitees of the clean room.
  It is only returned in the output when clean room creator calls GET
* `invite_recipient_workspace_id` (integer, optional) - Workspace ID of the user who is receiving the clean room "invitation". Must be specified if
  invite_recipient_email is specified.
  It should be empty when the collaborator is the creator of the clean room

### CleanRoomOutputCatalog
* `catalog_name` (string, optional) - The name of the output catalog in UC.
  It should follow [UC securable naming requirements](https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements).
  The field will always exist if status is CREATED

### CleanRoomRemoteDetail
* `cloud_vendor` (string, optional) - Cloud vendor (aws,azure,gcp) of the central clean room
* `collaborators` (list of CleanRoomCollaborator, optional) - Collaborators in the central clean room. There should one and only one collaborator
  in the list that satisfies the owner condition:
  
  1. It has the creator's global_metastore_id (determined by caller of CreateCleanRoom).
  
  2. Its invite_recipient_email is empty
* `egress_network_policy` (EgressNetworkPolicy, optional) - Egress network policy to apply to the central clean room workspace
* `region` (string, optional) - Region of the central clean room

### ComplianceSecurityProfile
* `compliance_standards` (list of ComplianceStandard, optional) - The list of compliance standards that the compliance security profile is configured to enforce
* `is_enabled` (boolean, optional) - Whether the compliance security profile is enabled

### EgressNetworkPolicy
* `internet_access` (EgressNetworkPolicyInternetAccessPolicy, optional) - The access policy enforced for egress traffic to the internet

### EgressNetworkPolicyInternetAccessPolicy
* `allowed_internet_destinations` (list of EgressNetworkPolicyInternetAccessPolicyInternetDestination, optional)
* `allowed_storage_destinations` (list of EgressNetworkPolicyInternetAccessPolicyStorageDestination, optional)
* `log_only_mode` (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, optional) - Optional. If not specified, assume the policy is enforced for all workloads
* `restriction_mode` (string, optional) - . Possible values are: `FULL_ACCESS`, `PRIVATE_ACCESS_ONLY`, `RESTRICTED_ACCESS`

### EgressNetworkPolicyInternetAccessPolicyInternetDestination
* `destination` (string, optional)
* `protocol` (string, optional) - . Possible values are: `TCP`
* `type` (string, optional) - . Possible values are: `FQDN`

### EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
* `log_only_mode_type` (string, optional) - . Possible values are: `ALL_SERVICES`, `SELECTED_SERVICES`
* `workloads` (list of , optional)

### EgressNetworkPolicyInternetAccessPolicyStorageDestination
* `allowed_paths` (list of string, optional)
* `azure_container` (string, optional)
* `azure_dns_zone` (string, optional)
* `azure_storage_account` (string, optional)
* `azure_storage_service` (string, optional)
* `bucket_name` (string, optional)
* `region` (string, optional)
* `type` (string, optional) - . Possible values are: `AWS_S3`, `AZURE_STORAGE`, `CLOUDFLARE_R2`, `GOOGLE_CLOUD_STORAGE`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `access_restricted` (string) - Whether clean room access is restricted due to [CSP](https://docs.databricks.com/en/security/privacy/security-profile.html). Possible values are: `CSP_MISMATCH`, `NO_RESTRICTION`
* `created_at` (integer) - When the clean room was created, in epoch milliseconds
* `local_collaborator_alias` (string) - The alias of the collaborator tied to the local clean room
* `output_catalog` (CleanRoomOutputCatalog) - Output catalog of the clean room. It is an output only field. Output catalog is manipulated
  using the separate CreateCleanRoomOutputCatalog API
* `status` (string) - Clean room status. Possible values are: `ACTIVE`, `DELETED`, `FAILED`, `PROVISIONING`
* `updated_at` (integer) - When the clean room was last updated, in epoch milliseconds

### CleanRoomCollaborator
* `display_name` (string) - Generated display name for the collaborator. In the case of a single metastore clean room, it is the clean
  room name. For x-metastore clean rooms, it is the organization name of the metastore. It is not restricted to
  these values and could change in the future
* `organization_name` (string) - [Organization name](:method:metastores/list#metastores-delta_sharing_organization_name)
  configured in the metastore

### CleanRoomOutputCatalog
* `status` (string) - . Possible values are: `CREATED`, `NOT_CREATED`, `NOT_ELIGIBLE`

### CleanRoomRemoteDetail
* `central_clean_room_id` (string) - Central clean room ID
* `compliance_security_profile` (ComplianceSecurityProfile)
* `creator` (CleanRoomCollaborator) - Collaborator who creates the clean room

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_clean_rooms_clean_room.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_clean_rooms_clean_room name
```