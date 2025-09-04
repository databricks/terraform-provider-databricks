---
subcategory: "OAuth"
---
# databricks_account_federation_policy Data Source
This data source can be used to get a single account federation policy.

-> **Note** This data source can only be used with an account-level provider!

## Example Usage
Referring to an account federation policy by id:

```hcl
data "databricks_account_federation_policy" "my_policy" {
  policy_id = "my-policy"
  oidc_policy = {
    issuer = "https://myidp.example.com"
    subject_claim = "sub"
  }
}
```

## Arguments
The following arguments are supported:
* `policy_id` (string, required) - The ID of the federation policy

## Attributes
The following attributes are exported:
* `create_time` (string) - Creation time of the federation policy
* `description` (string) - Description of the federation policy
* `name` (string) - Resource name for the federation policy. Example values include
  `accounts/<account-id>/federationPolicies/my-federation-policy` for Account Federation Policies, and
  `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
  for Service Principal Federation Policies. Typically an output parameter, which does not need to be
  specified in create or update requests. If specified in a request, must match the value in the
  request URL
* `oidc_policy` (OidcFederationPolicy)
* `policy_id` (string) - The ID of the federation policy
* `service_principal_id` (integer) - The service principal ID that this federation policy applies to. Only set for service principal federation policies
* `uid` (string) - Unique, immutable id of the federation policy
* `update_time` (string) - Last update time of the federation policy

### OidcFederationPolicy
* `audiences` (list of string) - The allowed token audiences, as specified in the 'aud' claim of federated tokens.
  The audience identifier is intended to represent the recipient of the token.
  Can be any non-empty string value. As long as the audience in the token matches
  at least one audience in the policy, the token is considered a match. If audiences
  is unspecified, defaults to your Databricks account id
* `issuer` (string) - The required token issuer, as specified in the 'iss' claim of federated tokens
* `jwks_json` (string) - The public keys used to validate the signature of federated tokens, in JWKS format.
  Most use cases should not need to specify this field. If jwks_uri and jwks_json
  are both unspecified (recommended), Databricks automatically fetches the public
  keys from your issuer’s well known endpoint. Databricks strongly recommends
  relying on your issuer’s well known endpoint for discovering public keys
* `jwks_uri` (string) - URL of the public keys used to validate the signature of federated tokens, in
  JWKS format. Most use cases should not need to specify this field. If jwks_uri
  and jwks_json are both unspecified (recommended), Databricks automatically
  fetches the public keys from your issuer’s well known endpoint. Databricks
  strongly recommends relying on your issuer’s well known endpoint for discovering
  public keys
* `subject` (string) - The required token subject, as specified in the subject claim of federated tokens.
  Must be specified for service principal federation policies. Must not be specified
  for account federation policies
* `subject_claim` (string) - The claim that contains the subject of the token. If unspecified, the default value
  is 'sub'