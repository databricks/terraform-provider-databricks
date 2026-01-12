---
subcategory: "OAuth"
---
# databricks_service_principal_federation_policy Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

Service principal federation policies allow automated workloads running outside of Databricks to access Databricks APIs without the need for Databricks secrets. Your application (workload) authenticates to Databricks as a Databricks service principal using tokens issued by the workload runtime, for example Github Actions.

A service principal federation policy is associated with a service principal in your Databricks account, and specifies:
* The identity provider (or issuer) from which the service principal can authenticate.
* The workload identity (or subject) that is permitted to authenticate as the Databricks service principal.


## Example Usage
```hcl
resource "databricks_service_principal_federation_policy" "this" {
  service_principal_id = 1234
  policy_id = "my-policy"
  oidc_policy = {
    issuer = "https://myidp.example.com"
    subject_claim = "sub"
    subject = "subject-in-token-from-myidp"
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
  id = "service_principal_id,policy_id"
  to = databricks_service_principal_federation_policy.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_service_principal_federation_policy.this "service_principal_id,policy_id"
```