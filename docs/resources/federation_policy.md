---
subcategory: "Delta Sharing"
---
# databricks_federation_policy Resource


## Example Usage


## Arguments
The following arguments are supported:
* `comment` (string, optional) - Description of the policy. This is a user-provided description
* `create_time` (string, optional) - System-generated timestamp indicating when the policy was created
* `id` (string, optional) - Unique, immutable system-generated identifier for the federation policy
* `name` (string, optional) - Name of the federation policy. A recipient can have multiple policies with different names
* `oidc_policy` (OidcFederationPolicy, optional) - Specifies the policy to use for validating OIDC claims in the federated tokens
* `update_time` (string, optional) - System-generated timestamp indicating when the policy was last updated

### OidcFederationPolicy
* `issuer` (string, required) - The required token issuer, as specified in the 'iss' claim of federated tokens
* `subject` (string, required) - The required token subject, as specified in the subject claim of federated tokens.
  The value of subject claim identifies the identity of the user or machine that is accessing the resource.
  For example for Entra ID (AAD)
  - For U2M flow, when allowing a group of users to access a resource and the subject claim is `groups`, this must be the Object ID of the group in Entra ID
  - For U2M flow, when allowing a user to access a resource and the subject claim is `oid`, this must be the Object ID of the user in Entra Id.
  - For M2M flow, when allowing an OAuth App registered to access a resource and the subject claim is `azp`, this must be the client-id of the oauth app registered in Entra ID
* `subject_claim` (string, required) - The claim that contains the subject of the token.
  Depending on the identity provider and the use case U2M or M2M, this can be different.
  For example for Entra ID (AAD)
  - For U2M flow, when allowing a group of users to access a resource, this must be `groups`
  - For U2M flow, when allowing a user to access a resource, this must be `oid`
  - For M2M flow, when allowing an OAuth App registered to access a resource, this must be `azp`
* `audiences` (list of string, optional) - The allowed token audiences, as specified in the 'aud' claim of federated tokens.
  The audience identifier is intended to represent the recipient of the token.
  Can be any non-empty string value. As long as the audience in the token matches at least one audience in the policy,

## Attributes
In addition to the above arguments, the following attributes are exported:

## Import
As of terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_federation_policy.this
}
```

If you are using an older version of terraform, you can import the resource using cli as follows:
```sh
$ terraform import databricks_federation_policy name
```