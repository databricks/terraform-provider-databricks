---
subcategory: "Delta Sharing"
---
# databricks_recipient_federation_policy Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Name of the federation policy. A recipient can have multiple policies with different names.
  The name must contain only lowercase alphanumeric characters, numbers, and hyphens

## Attributes
The following attributes are exported:
* `comment` (string) - Description of the policy. This is a user-provided description
* `create_time` (string) - System-generated timestamp indicating when the policy was created
* `id` (string) - Unique, immutable system-generated identifier for the federation policy
* `name` (string) - Name of the federation policy. A recipient can have multiple policies with different names.
  The name must contain only lowercase alphanumeric characters, numbers, and hyphens
* `oidc_policy` (OidcFederationPolicy) - Specifies the policy to use for validating OIDC claims in the federated tokens
* `update_time` (string) - System-generated timestamp indicating when the policy was last updated

### OidcFederationPolicy
* `audiences` (list of string) - The allowed token audiences, as specified in the 'aud' claim of federated tokens.
  The audience identifier is intended to represent the recipient of the token.
  Can be any non-empty string value. As long as the audience in the token matches at least one audience in the policy,
* `issuer` (string) - The required token issuer, as specified in the 'iss' claim of federated tokens
* `subject` (string) - The required token subject, as specified in the subject claim of federated tokens.
  The subject claim identifies the identity of the user or machine accessing the resource.
  Examples for Entra ID (AAD):
  - U2M flow (group access): If the subject claim is `groups`, this must be the Object ID of the group in Entra ID.
  - U2M flow (user access): If the subject claim is `oid`, this must be the Object ID of the user in Entra ID.
  - M2M flow (OAuth App access): If the subject claim is `azp`, this must be the client ID of the OAuth app registered in Entra ID
* `subject_claim` (string) - The claim that contains the subject of the token.
  Depending on the identity provider and the use case (U2M or M2M), this can vary:
  - For Entra ID (AAD):
  * U2M flow (group access): Use `groups`.
  * U2M flow (user access): Use `oid`.
  * M2M flow (OAuth App access): Use `azp`.
  - For other IdPs, refer to the specific IdP documentation.
  
  Supported `subject_claim` values are:
  - `oid`: Object ID of the user.
  - `azp`: Client ID of the OAuth app.
  - `groups`: Object ID of the group.
  - `sub`: Subject identifier for other use cases