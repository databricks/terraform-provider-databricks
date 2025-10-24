---
subcategory: "OAuth"
---
# databricks_account_federation_policy Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

Account federation policies allow users and service principals in your Databricks account to securely access Databricks APIs using tokens from your trusted identity providers (IdPs).

Token federation policies eliminate the need to manage Databricks secrets, and allow you to centralize management of token issuance policies in your IdP. Databricks token federation policies are typically used in combination with [SCIM](/admin/users-groups/scim/index.html), so users in your IdP are synchronized into your Databricks account.

An account federation policy specifies:
* which IdP, or issuer, your Databricks account should accept tokens from
* how to determine which Databricks user, or subject, a token is issued for

## Example Usage
```hcl
resource "databricks_account_federation_policy" "this" {
  policy_id = "my-policy"
  oidc_policy = {
    issuer = "https://myidp.example.com"
    subject_claim = "sub"
  }
}
```

## Arguments
The following arguments are supported:
* `description` (string, optional) - Description of the federation policy
* `oidc_policy` (OidcFederationPolicy, optional)

### OidcFederationPolicy
* `audiences` (list of string, optional) - The allowed token audiences, as specified in the 'aud' claim of federated tokens.
  The audience identifier is intended to represent the recipient of the token.
  Can be any non-empty string value. As long as the audience in the token matches
  at least one audience in the policy, the token is considered a match. If audiences
  is unspecified, defaults to your Databricks account id
* `issuer` (string, optional) - The required token issuer, as specified in the 'iss' claim of federated tokens
* `jwks_json` (string, optional) - The public keys used to validate the signature of federated tokens, in JWKS format.
  Most use cases should not need to specify this field. If jwks_uri and jwks_json
  are both unspecified (recommended), Databricks automatically fetches the public
  keys from your issuer’s well known endpoint. Databricks strongly recommends
  relying on your issuer’s well known endpoint for discovering public keys
* `jwks_uri` (string, optional) - URL of the public keys used to validate the signature of federated tokens, in
  JWKS format. Most use cases should not need to specify this field. If jwks_uri
  and jwks_json are both unspecified (recommended), Databricks automatically
  fetches the public keys from your issuer’s well known endpoint. Databricks
  strongly recommends relying on your issuer’s well known endpoint for discovering
  public keys
* `subject` (string, optional) - The required token subject, as specified in the subject claim of federated tokens.
  Must be specified for service principal federation policies. Must not be specified
  for account federation policies
* `subject_claim` (string, optional) - The claim that contains the subject of the token. If unspecified, the default value
  is 'sub'

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Creation time of the federation policy
* `name` (string) - Resource name for the federation policy. Example values include
  `accounts/<account-id>/federationPolicies/my-federation-policy` for Account Federation Policies, and
  `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
  for Service Principal Federation Policies. Typically an output parameter, which does not need to be
  specified in create or update requests. If specified in a request, must match the value in the
  request URL
* `policy_id` (string) - The ID of the federation policy. Output only
* `service_principal_id` (integer) - The service principal ID that this federation policy applies to. Output only. Only set for service principal federation policies
* `uid` (string) - Unique, immutable id of the federation policy
* `update_time` (string) - Last update time of the federation policy

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "policy_id"
  to = databricks_account_federation_policy.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_federation_policy.this "policy_id"
```