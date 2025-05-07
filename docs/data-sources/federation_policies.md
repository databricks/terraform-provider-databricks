---
subcategory: "Delta Sharing"
---
# databricks_federation_policies Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `recipient_name` (string, required) - Name of the recipient. This is the name of the recipient for which the policies are being listed
* `max_results` (integer, optional) - 



## Attributes
This data source exports a single attribute, policies. It is a list of resources, each with the following attributes:
* `comment` (string) - Description of the policy. This is a user-provided description
* `create_time` (string) - System-generated timestamp indicating when the policy was created
* `id` (string) - Unique, immutable system-generated identifier for the federation policy
* `name` (string) - Name of the federation policy. A recipient can have multiple policies with different names
* `oidc_policy` (OidcFederationPolicy) - Specifies the policy to use for validating OIDC claims in the federated tokens
* `update_time` (string) - System-generated timestamp indicating when the policy was last updated

### OidcFederationPolicy
* `audiences` (list of string) - The allowed token audiences, as specified in the 'aud' claim of federated tokens.
  The audience identifier is intended to represent the recipient of the token.
  Can be any non-empty string value. As long as the audience in the token matches at least one audience in the policy,
* `issuer` (string) - The required token issuer, as specified in the 'iss' claim of federated tokens
* `subject` (string) - The required token subject, as specified in the subject claim of federated tokens.
  The value of subject claim identifies the identity of the user or machine that is accessing the resource.
  For example for Entra ID (AAD)
  - For U2M flow, when allowing a group of users to access a resource and the subject claim is `groups`, this must be the Object ID of the group in Entra ID
  - For U2M flow, when allowing a user to access a resource and the subject claim is `oid`, this must be the Object ID of the user in Entra Id.
  - For M2M flow, when allowing an OAuth App registered to access a resource and the subject claim is `azp`, this must be the client-id of the oauth app registered in Entra ID
* `subject_claim` (string) - The claim that contains the subject of the token.
  Depending on the identity provider and the use case U2M or M2M, this can be different.
  For example for Entra ID (AAD)
  - For U2M flow, when allowing a group of users to access a resource, this must be `groups`
  - For U2M flow, when allowing a user to access a resource, this must be `oid`
  - For M2M flow, when allowing an OAuth App registered to access a resource, this must be `azp`